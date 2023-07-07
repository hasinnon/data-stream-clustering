// Code generated by hero.
// source: D:\project\GO\gowork\src\datastream\view\template\services.html
// DO NOT EDIT!
package gotemplate

import (
	"my/ar/399/datastream/controller/utility/ptime"
	"my/ar/399/datastream/datalayer"
	"fmt"
	"io"

	"github.com/shiyanhui/hero"
)

func ServicesHandler(u datalayer.UserLogin, serv []datalayer.ServiceInfo, w io.Writer) {
	_buffer := hero.GetBuffer()
	defer hero.PutBuffer(_buffer)
	_buffer.WriteString(`<!DOCTYPE html>
<html lang="fa" dir="rtl">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <!-- Meta, title, CSS, favicons, etc. -->
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="icon" href="/StaticUI/Images/icon.ico" type="image/ico" />

    <title>بـاداده | 
      `)
	_buffer.WriteString(`
  سرویس ها
`)

	_buffer.WriteString(`
    </title>

    <!-- Bootstrap -->
    <link href="/StaticUI/vendors/bootstrap/dist/css/bootstrap.min.css" rel="stylesheet">
    <!-- Font Awesome -->
    <link href="/StaticUI/vendors/font-awesome/css/font-awesome.min.css" rel="stylesheet">
    <!-- NProgress -->
    <link href="/StaticUI/vendors/nprogress/nprogress.css" rel="stylesheet">
    `)
	_buffer.WriteString(`
    <!-- iCheck -->
    <link href="/StaticUI/vendors/iCheck/skins/flat/green.css" rel="stylesheet">
    <style>
    #linedr:before, #linedr:after { 
      content: ""; 
      flex: 1 1; 
      border-bottom: 1px solid #e6e9ed; 
      margin: 0px 7px;
    } 
    </style>
`)

	_buffer.WriteString(`

    <!-- Custom Theme Style -->
    <link href="/StaticUI/build/css/custom.min.css" rel="stylesheet">
  </head>

  <body class="nav-md">
    <div class="container body">
      <div class="main_container">
        <div class="col-md-3 right_col menu_fixed">
          <div class="right_col scroll-view">
            <div class="navbar nav_title" style="border: 0;">
              <a href="/Dashboard" class="site_title">
                <img style="padding: 6px 10px 6px 7px;height: 38px; width: auto;display: inline;" src="/StaticUI/Images/logo.png" >
                <span> بـاداده </span></a>
            </div>
            <div class="clearfix"></div>
            <br />

            <!-- sidebar menu -->
            <div id="sidebar-menu" class="main_menu_side hidden-print main_menu">
              <div class="menu_section">
                <ul class="nav side-menu">
                  <li><a href="/Dashboard"><i class="fa fa-home"></i>پیش خوان</a>
                  </li>
                  <li><a><i class="fa fa-tasks"></i>سرویس ها<span class="fa fa-chevron-down"></span></a>
                    <ul class="nav child_menu">
                      <li><a href="/Services">مشاهده سرویس های ایجاد شده</a></li>
                      <li><a href="/NewService">ایجاد سرویس جدید</a></li>
                    </ul>
                  </li>
                  <li><a><i class="fa fa-dollar"></i>اعتبار حساب<span class="fa fa-chevron-down"></span></a>
                    <ul class="nav child_menu">
                      <li><a href="/Credit">مشاهده اعتبار حساب</a></li>
                      <li><a href="/IncreaseCredit">افزایش اعتبار حساب</a></li>
                    </ul>
                  </li>
                  <li><a><i class="fa fa-envelope"></i>پیام ها<span class="fa fa-chevron-down"></span></a>
                    <ul class="nav child_menu">
                      <li><a href="/Messages">مشاهده اعلان ها</a></li>
                      <li><a href="/ContactUs">ارتباط با ما</a></li>
                    </ul>
                  </li>
                  <li><a><i class="fa fa-user"></i>حساب کاربری<span class="fa fa-chevron-down"></span></a>
                    <ul class="nav child_menu">
                      <li><a href="/Settings">تنظیمات حساب کاربری</a></li>
                      <li><a href="/Logout">خروج</a></li>
                    </ul>
                  </li>
                </ul>
              </div>

            </div>
            <!-- /sidebar menu -->

          </div>
        </div>

        <!-- top navigation -->
        <div class="top_nav">
          <div class="nav_menu">
            <nav>
              <div class="nav toggle">
                <a id="menu_toggle"><i class="fa fa-bars"></i></a>
              </div>

              <ul class="nav navbar-nav navbar-left">
                <li class="">
                  <a href="javascript:;" class="user-profile dropdown-toggle" data-toggle="dropdown" aria-expanded="false">
                    `)
	hero.EscapeHTML(u.Fname+" "+u.Lname, _buffer)

	_buffer.WriteString(`            
                    <span class=" fa fa-angle-down"></span>
                  </a>
                  <ul class="dropdown-menu dropdown-usermenu pull-right">                     
                  </ul>
                </li>
              </ul>
            </nav>
          </div>
        </div>
        <!-- /top navigation -->

        <!-- page content -->
        <div class="left_col" role="main">
          `)
	_buffer.WriteString(`
  <div class="">
      <div class="page-title">
        <div class="title_right">
          <h3>سرویس‌ ها</h3>
        </div>

        <div class="title_left">
          <div class="form-group pull-left" style="display: inline;">
            <div class="btn-group">
              <a href="/NewService">
              <button type="button" class="btn btn-default"> ایجاد سرویس جدید </button> 
              </a>
              <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-expanded="false">
                <span class="caret"></span>
                <span class="sr-only">#</span>
              </button>
              <ul class="dropdown-menu" role="menu">
                <li><a href="/AddStructuredService">سرویس ساختاریافته</a>
                </li>
                <li><a href="/AddUnStructuredService">سرویس بدون ساختار</a>
                </li>
                <li class="divider"></li>
                <li><a href="/NewService">ایجاد سرویس جدید</a>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>
      
      <div class="clearfix"></div>

      <div class="row">
        <div class="col-md-12 col-sm-12 col-xs-12">
          <div class="x_panel">
            <div class="x_title">
              <h2>سرویس‌ های ایجاد شده</h2>
              <ul class="nav navbar-left panel_toolbox">    
                <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>
                </li>
              </ul>
              <div class="clearfix"></div>
            </div>
            <div class="x_content">
              <!-- start services list -->
              `)
	if len(serv) != 0 {
		_buffer.WriteString(`
                <div class="table-responsive">
                  <table class="table table-striped projects">
                    <thead>
                      <tr>
                        <th style="width: 1%">#</th>
                        <th style="width: 15%">نام سرویس</th>
                        <th>تاریخ&nbsp;ایجاد</th>
                        <th style="text-align: center;">تاریخ&nbsp;انقضا</th>
                        <th style="width: 15%; text-align: center;">وضعیت</th>
                        <th style="text-align: center;">اعتبار&nbsp;سرویس</th>
                        <th style="width: 25%; text-align: center;">ویرایش</th>
                      </tr>
                    </thead>                
                    <tbody>
                      `)
		num := 0
		for _, s := range serv {
			num++
			if s.BServ.Type.Bool {
				_buffer.WriteString(`
                          <tr>
                            <td>`)
				hero.EscapeHTML(fmt.Sprint(num), _buffer)
				_buffer.WriteString(`</td>
                            <td>
                              <a dir="ltr">`)
				hero.EscapeHTML(s.BServ.Name.String, _buffer)
				_buffer.WriteString(`</a>   
                            </td>
                            <td>
                              <a dir="ltr">`)
				hero.EscapeHTML(ptime.ToPersian(s.BServ.Created.Time), _buffer)
				_buffer.WriteString(` </a>
                              <br />
                              <small>`)
				hero.EscapeHTML(s.BServ.Created.Time.Format("15:04:05"), _buffer)
				_buffer.WriteString(`</small>
                            </td>
                            <td>
                              <a dir="ltr">`)
				hero.EscapeHTML(ptime.ToPersian(s.SServ.ExpScheduled.Time), _buffer)
				_buffer.WriteString(` </a>
                              <br />
                              <small>`)
				hero.EscapeHTML(s.SServ.ExpScheduled.Time.Format("15:04:05"), _buffer)
				_buffer.WriteString(`</small>
                            </td>
                            <td>
                              `)
				if s.SServ.Status.Bool == true {
					_buffer.WriteString(`
                                <button type="button" class="btn btn-success btn-block btn-xs">فعال</button>
                              `)
				} else {
					_buffer.WriteString(`
                                <button type="button" class="btn btn-danger btn-block btn-xs">غیرفعال</button>
                              `)
				}
				_buffer.WriteString(`
                            </td>
                            <td class="project_progress">
                              <div style="text-align: center;"><a><i class="fa fa-ticket"></i> `)
				hero.EscapeHTML(fmt.Sprint(s.BServ.Credit.Int32), _buffer)
				_buffer.WriteString(` (ساعت)</a></div>
                              <div class="progress progress_sm">
                                <div class="progress-bar bg-green" role="progressbar" data-transitiongoal="`)
				hero.EscapeHTML(fmt.Sprint(s.SServ.Used), _buffer)
				_buffer.WriteString(`"></div>
                              </div>                              
                              <small>`)
				hero.EscapeHTML(fmt.Sprint(s.SServ.Used)+"%", _buffer)
				_buffer.WriteString(` استفاده شده</small>
                            </td>   
                            <td style="align-items: center; display: flex; justify-content: center;">
                              <a href="/Service/`)
				hero.EscapeHTML(fmt.Sprint(s.BServ.Sid.Int32), _buffer)
				_buffer.WriteString(`/View" class="btn btn-primary btn-xs"><i class="fa fa-info"></i>  مشاهده  </a>
                              <a href="/Service/`)
				hero.EscapeHTML(fmt.Sprint(s.BServ.Sid.Int32), _buffer)
				_buffer.WriteString(`/Management" class="btn btn-info btn-xs"><i class="fa fa-gears"></i>  مدیریت  </a>
                            </td>
                          </tr>
                        `)
			} else {
				_buffer.WriteString(`
                          <tr>
                            <td>`)
				hero.EscapeHTML(fmt.Sprint(num), _buffer)
				_buffer.WriteString(`</td>
                            <td>
                              <a dir="ltr">`)
				hero.EscapeHTML(s.BServ.Name.String, _buffer)
				_buffer.WriteString(`</a>                    
                            </td>
                            <td>
                              <a dir="ltr">`)
				hero.EscapeHTML(ptime.ToPersian(s.BServ.Created.Time), _buffer)
				_buffer.WriteString(` </a>
                              <br />
                              <small>`)
				hero.EscapeHTML(s.BServ.Created.Time.Format("15:04:05"), _buffer)
				_buffer.WriteString(`</small>
                            </td>
                            <td colspan="2"  >
                              <div id="linedr" style="align-items: center; display: flex; justify-content: center;">
                                <small>سرویس بدون ساختار</small>
                              </div>                            
                            </td>
                            <td >
                              <div style="align-items: center; display: flex; justify-content: center;">
                              `)
				if s.BServ.Credit.Int32 == -1 {
					_buffer.WriteString(`
                                <a><i class="fa fa-ticket"></i>  رایگان </a>
                              `)
				} else {
					_buffer.WriteString(`
                                <a><i class="fa fa-ticket"></i> `)
					hero.EscapeHTML(fmt.Sprint(s.SServ.Credit.Int32), _buffer)
					_buffer.WriteString(` </a>
                              `)
				}
				_buffer.WriteString(`
                              </div>
                            </td>    
              
                            <td style="align-items: center; display: flex; justify-content: center;">
                              <a href="/Service/`)
				hero.EscapeHTML(fmt.Sprint(s.BServ.Sid.Int32), _buffer)
				_buffer.WriteString(`/Result" class="btn btn-primary btn-xs"><i class="fa fa-info"></i>  مشاهده  </a>
                            </td>
                          </tr>
                        `)
			}
		}
		_buffer.WriteString(`
                    </tbody>
                  </table>
                </div>
                `)
	} else {
		_buffer.WriteString(`
                  <div id="singuptxt" class="form-subnote verfication-subnote text-center">
                    <span class="form-subnote__label">
                        <h3>تا کنون سرویسی را ایجاد نکردید</h3>
                        <br>
                        <h2>برای ایجاد سرویس جدید <a href="/NewService"><strong>کلیک</strong></a> کنید</h2>
                    </span>
                </div> 
              `)
	}
	_buffer.WriteString(`
              
              <!-- end services list -->

            </div>
          </div>
        </div>
      </div>
  </div>  
`)

	_buffer.WriteString(`
        </div>
        <!-- /page content -->

        <!-- footer content -->
        <footer>
          <div class="pull-left">
            بـاداده » سامانه برخط خوشه‌بندی داده‌های جریانی
          </div>
          <div class="clearfix"></div>
        </footer>
        <!-- /footer content -->

      </div>
    </div>

    <!-- jQuery -->
    <script src="/StaticUI/vendors/jquery/dist/jquery.min.js"></script>
    <!-- Bootstrap -->
    <script src="/StaticUI/vendors/bootstrap/dist/js/bootstrap.min.js"></script>
    <!-- FastClick -->
    <script src="/StaticUI/vendors/fastclick/lib/fastclick.js"></script>
    <!-- NProgress -->
    <script src="/StaticUI/vendors/nprogress/nprogress.js"></script>
    `)
	_buffer.WriteString(`
  <!-- bootstrap-progressbar -->
  <script src="/StaticUI/vendors/bootstrap-progressbar/bootstrap-progressbar.min.js"></script>
`)

	_buffer.WriteString(`

    <!-- Custom Theme Scripts -->
    <script src="/StaticUI/build/js/custom.min.js"></script>
    `)
	_buffer.WriteString(`
    
  </body>
</html>`)
	w.Write(_buffer.Bytes())

}
