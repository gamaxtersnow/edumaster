// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	class "edumaster/internal/handler/class"
	classabsencenotification "edumaster/internal/handler/classabsencenotification"
	classattendancerecord "edumaster/internal/handler/classattendancerecord"
	classattendancerecordtimeout "edumaster/internal/handler/classattendancerecordtimeout"
	classschedule "edumaster/internal/handler/classschedule"
	course "edumaster/internal/handler/course"
	customer "edumaster/internal/handler/customer"
	customerrelationship "edumaster/internal/handler/customerrelationship"
	customerrelationshipgroup "edumaster/internal/handler/customerrelationshipgroup"
	financial "edumaster/internal/handler/financial"
	leaveapplication "edumaster/internal/handler/leaveapplication"
	noticemanage "edumaster/internal/handler/noticemanage"
	notification "edumaster/internal/handler/notification"
	order "edumaster/internal/handler/order"
	student "edumaster/internal/handler/student"
	studentschedule "edumaster/internal/handler/studentschedule"
	teacher "edumaster/internal/handler/teacher"
	teacherschedule "edumaster/internal/handler/teacherschedule"
	timeschedule "edumaster/internal/handler/timeschedule"
	"edumaster/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/class/add",
				Handler: class.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/class/delete",
				Handler: class.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/class/info",
				Handler: class.InfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/class/list",
				Handler: class.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/class/modify",
				Handler: class.ModifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/class/stop",
				Handler: class.ClassStopHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/class/absence/notification/list",
				Handler: classabsencenotification.ListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/class/attendance/record/add",
				Handler: classattendancerecord.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/class/attendance/record/delete",
				Handler: classattendancerecord.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/class/attendance/record/info",
				Handler: classattendancerecord.InfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/class/attendance/record/list",
				Handler: classattendancerecord.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/class/attendance/record/modify",
				Handler: classattendancerecord.ModifyHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/class/attendance/record/timeout/list",
				Handler: classattendancerecordtimeout.ListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/classschedule/add",
				Handler: classschedule.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/classschedule/delete",
				Handler: classschedule.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/classschedule/info",
				Handler: classschedule.InfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/classschedule/list",
				Handler: classschedule.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/classschedule/modify",
				Handler: classschedule.ModifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/classschedule/stop",
				Handler: classschedule.StopHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/course/add",
				Handler: course.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/course/delete",
				Handler: course.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/course/info",
				Handler: course.InfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/course/list",
				Handler: course.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/course/modify",
				Handler: course.ModifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/course/stop",
				Handler: course.StopHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/customer/add",
				Handler: customer.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/customer/delete",
				Handler: customer.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/customer/info",
				Handler: customer.InfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/customer/list",
				Handler: customer.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/customer/modify",
				Handler: customer.ModifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/customer/stop",
				Handler: customer.StopHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/customer/relationship/add",
				Handler: customerrelationship.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/customer/relationship/delete",
				Handler: customerrelationship.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/customer/relationship/info",
				Handler: customerrelationship.InfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/customer/relationship/list",
				Handler: customerrelationship.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/customer/relationship/modify",
				Handler: customerrelationship.ModifyHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/customer/relationship/group/add",
				Handler: customerrelationshipgroup.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/customer/relationship/group/delete",
				Handler: customerrelationshipgroup.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/customer/relationship/group/info",
				Handler: customerrelationshipgroup.InfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/customer/relationship/group/list",
				Handler: customerrelationshipgroup.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/customer/relationship/group/modify",
				Handler: customerrelationshipgroup.ModifyHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/financial/salary/settlement",
				Handler: financial.SalarysettlementHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/leave/application/list",
				Handler: leaveapplication.ListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/NoticeManage/add",
				Handler: noticemanage.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/NoticeManage/delete",
				Handler: noticemanage.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/NoticeManage/info",
				Handler: noticemanage.InfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/NoticeManage/list",
				Handler: noticemanage.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/NoticeManage/modify",
				Handler: noticemanage.ModifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/NoticeManage/stop",
				Handler: noticemanage.StopHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/notification/add",
				Handler: notification.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notification/delete",
				Handler: notification.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notification/info",
				Handler: notification.InfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notification/list",
				Handler: notification.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notification/modify",
				Handler: notification.ModifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notification/stop",
				Handler: notification.StopHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/order/list",
				Handler: order.ListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/student/add",
				Handler: student.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/student/delete",
				Handler: student.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/student/enrollment/situation",
				Handler: student.StudentenrollmentsituationHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/student/info",
				Handler: student.InfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/student/list",
				Handler: student.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/student/modify",
				Handler: student.ModifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/student/stop",
				Handler: student.StopHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/studentschedule/add",
				Handler: studentschedule.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/studentschedule/delete",
				Handler: studentschedule.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/studentschedule/info",
				Handler: studentschedule.InfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/studentschedule/list",
				Handler: studentschedule.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/studentschedule/modify",
				Handler: studentschedule.ModifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/studentschedule/stop",
				Handler: studentschedule.StopHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/teacher/add",
				Handler: teacher.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/teacher/delete",
				Handler: teacher.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/teacher/info",
				Handler: teacher.InfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/teacher/list",
				Handler: teacher.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/teacher/modify",
				Handler: teacher.ModifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/teacher/stop",
				Handler: teacher.StopHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/teacherschedule/add",
				Handler: teacherschedule.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/teacherschedule/delete",
				Handler: teacherschedule.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/teacherschedule/info",
				Handler: teacherschedule.InfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/teacherschedule/list",
				Handler: teacherschedule.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/teacherschedule/modify",
				Handler: teacherschedule.ModifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/teacherschedule/stop",
				Handler: teacherschedule.StopHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/timeschedule/add",
				Handler: timeschedule.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/timeschedule/delete",
				Handler: timeschedule.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/timeschedule/info",
				Handler: timeschedule.InfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/timeschedule/list",
				Handler: timeschedule.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/timeschedule/modify",
				Handler: timeschedule.ModifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/timeschedule/stop",
				Handler: timeschedule.StopHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)
}
