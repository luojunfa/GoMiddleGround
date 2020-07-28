package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/go-xweb/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)
var(
	wxGameDBConfig = DBConfig{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "qwer123",
		DBName:   "wxgame",
	}
	wxgameDB *gorm.DB//数据库定义为全局变量
	myLog    *log.Logger
	CheckTime=flag.String("CheckTime",time.Now().AddDate(0,0,-4).Format("2006-01-02"),"date")//获取前一天的时间,运行就计算
)
// 数据库连接信息
type DBConfig struct {
	Host     string
	Port     int32
	UserName string
	Password string
	DBName   string
}
type AllMatchUserNum struct {
	LogDate string `json:"log_date" gorm:"column:log_date"`
	MatchId int `json:"match_id" gorm:"column:match_id"`
	SmallMatchId int `json:"small_match_id"gorm:"column:small_match_id"`
	RobotUserNum int `json:"robot_user_num" gorm:"column:robot_user_num"`
	TrueUserNum int `json:"true_user_num" gorm:"column:true_user_num"`
	MatchCount int`json:"match_count"gorm:"column:match_count"`
	LogTime string `json:"log_time" gorm:"column:log_time"`
}
func main()  {

	wxgameDB=ConnectDB(wxGameDBConfig)
	MatchFilterDataAndInsert()
	SmallMatchFilterDataAndInsert()
}
//数据初筛，前一天的数据match_id,small_match_id,robot_user_num,true_user_num,match_count
func ALLPeopleEnterGame(Robot bool,SmallMatch bool) []AllMatchUserNum {
	a := make([]AllMatchUserNum, 0)
	sql := wxgameDB.Debug().Table("match_user_info")
	if !SmallMatch{
		if Robot {
			sql= sql.Select("match_id, count(distinct user_id) as robot_user_num, log_date, count(small_match_id) as match_count")
		}
		// 不是机器人
		if !Robot {
			sql = sql.Select("match_id, count(distinct user_id) as true_user_num, log_date, count(small_match_id) as match_count")
		}
		if err := sql.Where("robot=?and log_date=?", Robot,CheckTime).Group("match_id,log_date").Find(&a).Error;
		err != nil {
			myLog.Printf("match_user_info,ALLPeopleEnterGame,查询出错")
		}
		return a
	}else{
		if Robot{
			sql.Select("match_id, small_match_id,count(distinct user_id) as robot_user_num, min(log_time) as log_time")
		}else{
			sql.Select("match_id, small_match_id,count(distinct user_id) as true_user_num, min(log_time) as log_time")
		}
		if err:=sql.Where("robot=?and log_date=?",Robot,CheckTime).Find(&a).Error;
		 err != nil {
			myLog.Printf("查询出错")
		}
		return a
	}
	//}
}
//对数据做过滤和插入表操作
func MatchFilterDataAndInsert() {
	// 最后数据
	AllUserLog := make([]AllMatchUserNum, 0)
	AllUserLogdateChanged:=make([]AllMatchUserNum,0)
	RobotUserNum := ALLPeopleEnterGame(true,false)   // 机器人数量
	TrueUserNum:=  ALLPeopleEnterGame(false,false)  // 真实用户数量
	var SpecialNumI []int
	var SpecialNumJ []int
	for i:=0;i<len(RobotUserNum);i++{
		for j:=0;j<len(TrueUserNum);j++{
			if RobotUserNum[i].MatchId==TrueUserNum[j].MatchId && RobotUserNum[i].LogDate==TrueUserNum[j].LogDate{
				AllUserLog=append(AllUserLog,AllMatchUserNum{
					MatchId: RobotUserNum[i].MatchId,
					LogDate: RobotUserNum[i].LogDate,
					RobotUserNum: RobotUserNum[i].RobotUserNum,
					TrueUserNum: TrueUserNum[j].TrueUserNum,
					MatchCount: RobotUserNum[i].MatchCount+TrueUserNum[j].MatchCount,
				})
				SpecialNumI=append(SpecialNumI,i)
				SpecialNumJ=append(SpecialNumJ,j)
			}
		}
	}
	for i:=0;i<len(RobotUserNum);i++{
		if IsContain(i,SpecialNumI){
		}else {
			AllUserLog=append(AllUserLog,AllMatchUserNum{
				MatchId: RobotUserNum[i].MatchId,
				LogDate: RobotUserNum[i].LogDate,
				RobotUserNum: RobotUserNum[i].RobotUserNum,
				MatchCount: RobotUserNum[i].MatchCount,
			})
		}
	}
	for j:=0;j<len(TrueUserNum);j++{
		if IsContain(j,SpecialNumJ){
		}else {
			AllUserLog=append(AllUserLog,AllMatchUserNum{
				MatchId: TrueUserNum[j].MatchId,
				LogDate: TrueUserNum[j].LogDate,
				TrueUserNum: TrueUserNum[j].TrueUserNum,
				MatchCount: RobotUserNum[j].MatchCount,
			})
		}
	}
	AllUserLogdateChanged=Replace(AllUserLog)
	fmt.Println(AllUserLogdateChanged)
	var buffer bytes.Buffer
	sql := "insert into match_user_counts (match_id, robot_user_num, true_user_num, log_date, match_count) values "
	if _, err := buffer.WriteString(sql); err != nil {
		fmt.Println("")
	}
	for i, e := range AllUserLogdateChanged {
		if i == len(AllUserLogdateChanged)-1 {
			buffer.WriteString(fmt.Sprintf("(%d,%d,%d,%s,%d);", e.MatchId, e.RobotUserNum,e.TrueUserNum,e.LogDate,e.MatchCount))
		} else {
			buffer.WriteString(fmt.Sprintf("(%d,%d,%d,%s,%d),", e.MatchId, e.RobotUserNum,e.TrueUserNum,e.LogDate,e.MatchCount))
		}
	}
	if err:=wxgameDB.Debug().Exec(buffer.String()).Error;
		err!=nil{
		fmt.Println("插入数据错误:",err)
	}
}
//small-match
func SmallMatchFilterDataAndInsert()  {
	//AllUserLogdateChanged:=make([]AllMatchUserNum,0)
	AllUserLog:=make([]AllMatchUserNum,0)
	RobotUserNum:=make([]AllMatchUserNum,0)
	TrueUserNum:=make([]AllMatchUserNum,0)
	RobotUserNum=ALLPeopleEnterGame(true,true)
	TrueUserNum=ALLPeopleEnterGame(false,true)
	var SpecialNumI []int
	var SpecialNumJ []int
	for i:=0;i<len(RobotUserNum);i++{
		for j:=0;j<len(TrueUserNum);j++{
			if RobotUserNum[i].MatchId==TrueUserNum[j].MatchId && RobotUserNum[i].SmallMatchId==TrueUserNum[j].SmallMatchId{
				AllUserLog=append(AllUserLog,AllMatchUserNum{
					MatchId: RobotUserNum[i].MatchId,
					SmallMatchId: RobotUserNum[i].SmallMatchId,
					LogTime: RobotUserNum[i].LogTime,
					RobotUserNum: RobotUserNum[i].RobotUserNum,
					TrueUserNum: TrueUserNum[j].TrueUserNum,
				})
				SpecialNumI=append(SpecialNumI,i)
				SpecialNumJ=append(SpecialNumJ,j)
			}
		}
	}
	for i:=0;i<len(RobotUserNum);i++{
		if IsContain(i,SpecialNumI){
		}else {
			AllUserLog=append(AllUserLog,AllMatchUserNum{
				MatchId: RobotUserNum[i].MatchId,
				SmallMatchId: RobotUserNum[i].SmallMatchId,
				LogTime: RobotUserNum[i].LogTime,
				RobotUserNum: RobotUserNum[i].RobotUserNum,
			})
		}
	}
	for j:=0;j<len(TrueUserNum);j++{
		if IsContain(j,SpecialNumJ){
		}else {
			AllUserLog=append(AllUserLog,AllMatchUserNum{
				MatchId: TrueUserNum[j].MatchId,
				SmallMatchId: TrueUserNum[j].SmallMatchId,
				LogTime: TrueUserNum[j].LogTime,
				TrueUserNum: TrueUserNum[j].TrueUserNum,
			})
		}
	}
	fmt.Println(AllUserLog)
	//AllUserLogdateChanged=Replace(AllUserLog)
	//fmt.Println(AllUserLogdateChanged)
	//var buffer bytes.Buffer
	//sql := "insert into match_user_counts (match_id, robot_user_num, true_user_num, log_date, match_count) values "
	//if _, err := buffer.WriteString(sql); err != nil {
	//	fmt.Println("")
	//}
	//for i, e := range AllUserLogdateChanged {
	//	if i == len(AllUserLogdateChanged)-1 {
	//		buffer.WriteString(fmt.Sprintf("(%d,%d,%d,%s,%d);", e.MatchId, e.RobotUserNum,e.TrueUserNum,e.LogDate,e.MatchCount))
	//	} else {
	//		buffer.WriteString(fmt.Sprintf("(%d,%d,%d,%s,%d),", e.MatchId, e.RobotUserNum,e.TrueUserNum,e.LogDate,e.MatchCount))
	//	}
	//}
	//if err:=wxgameDB.Debug().Exec(buffer.String()).Error;
	//	err!=nil{
	//	fmt.Println("插入数据错误:",err)
	//}
}
//做判断
func IsContain(tar int, slice[]int)bool  {
	for i:=0;i<len(slice);i++ {
		if tar==slice[i]{
			return true
		}
	}
	return false
}
//ToString 表示连接数据库的字符串
func (c DBConfig) ToString() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", c.Host, c.Port,
		c.UserName, c.Password, c.DBName,
	)
}
//替换
func Replace( StructNeedToBeReplace []AllMatchUserNum) []AllMatchUserNum {
	AllUserLogdateChanged:=make([]AllMatchUserNum,0)
	for _,s:=range StructNeedToBeReplace{
		logdate:=fmt.Sprintf("'%s'",s.LogDate[0:10])
		AllUserLogdateChanged=append(AllUserLogdateChanged,AllMatchUserNum{
			MatchId: s.MatchId,
			LogDate: logdate,
			TrueUserNum: s.TrueUserNum,
			RobotUserNum: s.RobotUserNum,
			MatchCount: s.MatchCount,
		})
	}
	return AllUserLogdateChanged
}
//连接数据库
func ConnectDB(dbconfig DBConfig) *gorm.DB {
	db, err := gorm.Open("postgres", dbconfig.ToString())
	if err != nil {
		panic(fmt.Sprintf("gorm.Open: err:%v", err))
	}else{
		fmt.Println("Connected Successfully!")
	}
	// 设置最大链接数
	db.DB().SetMaxOpenConns(1000)
	// 开启gorm日志模块
	db.LogMode(true)
	return db
}

