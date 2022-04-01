// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// SysQuestion is the golang structure for table sys_question.
type SysQuestion struct {
    QuestionId int64       `orm:"question_id,primary" json:"questionId"` //                                
    UserId     int64       `orm:"user_id"             json:"userId"`     //                                
    GroupId    int64       `orm:"group_id"            json:"groupId"`    //                                
    Title      string      `orm:"title"               json:"title"`      //                                
    Content    string      `orm:"content"             json:"content"`    //                                
    Hots       int64       `orm:"hots"                json:"hots"`       //                                
    Favorites  int64       `orm:"favorites"           json:"favorites"`  //                                
    Likes      int64       `orm:"likes"               json:"likes"`      //                                
    Views      int64       `orm:"views"               json:"views"`      //                                
    Anonymous  int         `orm:"anonymous"           json:"anonymous"`  // 是否匿名                       
    Status     int         `orm:"status"              json:"status"`     // 状态：0全部,1待审核 ，2已发布  
    CreateTime *gtime.Time `orm:"create_time"         json:"createTime"` //                                
    UpdateTime *gtime.Time `orm:"update_time"         json:"updateTime"` //                                
    DeleteTime *gtime.Time `orm:"delete_time"         json:"deleteTime"` //                                
    Remark     string      `orm:"remark"              json:"remark"`     //                                
}