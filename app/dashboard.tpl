<!--
=========================================================
* Argon Dashboard - v1.2.0
=========================================================
* Product Page: https://www.creative-tim.com/product/argon-dashboard


* Copyright  Creative Tim (http://www.creative-tim.com)
* Coded by www.creative-tim.com



=========================================================
* The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
-->
<!DOCTYPE html>
<html>

<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
  <meta name="description" content="Start your development with a Dashboard for Bootstrap 4.">
  <meta name="author" content="Creative Tim">
  <title>MSCH</title>
  <!-- Favicon -->
  <link rel="icon" href="assets/img/brand/favicon.png" type="image/png">
  <!-- Fonts -->
  <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Open+Sans:300,400,600,700">
  <!-- Icons -->
  <link rel="stylesheet" href="assets/vendor/nucleo/css/nucleo.css" type="text/css">
  <link rel="stylesheet" href="assets/vendor/@fortawesome/fontawesome-free/css/all.min.css" type="text/css">
  <!-- Page plugins -->
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
  <!-- <link rel="stylesheet" href="assets/css/bootstrap/bootstrap-datetimepicker.min.css" type="text/css"> -->
  <link rel="stylesheet" href="assets/css/bootstrap/bootstrap.css" type="text/css">
  <link rel="stylesheet" href="assets/css/argon.css?v=1.2.0" type="text/css">

</head>

<body>
  <!-- Sidenav -->
  <nav class="sidenav navbar navbar-vertical  fixed-left  navbar-expand-xs navbar-light bg-white" id="sidenav-main">
    <div class="scrollbar-inner">
      <!-- Brand -->
      <div class="sidenav-header  align-items-center">
        <a class="navbar-brand" href="javascript:void(0)">
          <img src="assets/img/logo.png" class="navbar-brand-img" alt="...">
        </a>
      </div>
      <div class="navbar-inner">
        <!-- Collapse -->
        <div class="collapse navbar-collapse" id="sidenav-collapse-main">
          <!-- Nav items -->
          <ul class="navbar-nav">
            <li class="nav-item">
              <a class="nav-link active" href="/dashboard">
                <i class="ni ni-tv-2 text-primary"></i>
                <span class="nav-link-text">Dashboard</span>
              </a>
            </li>
            <a class="nav-link logout" href="#">
            <li class="nav-item">
                <i class="ni ni-key-25 text-info"></i>
                <span class="nav-link-text">Logout</span>
              </a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </nav>
  <!-- Main content -->
  <div class="main-content" id="panel">
    <!-- Topnav -->
    <nav class="navbar navbar-top navbar-expand navbar-dark bg-primary border-bottom">
      <div class="container-fluid">
        <div class="collapse navbar-collapse" id="navbarSupportedContent">
          <!-- Search form -->
          <form class="navbar-search navbar-search-light form-inline mr-sm-3" id="navbar-search-main">
            <div class="form-group mb-0">
              <div class="input-group input-group-alternative input-group-merge">
                <div class="input-group-prepend">
                  <span class="input-group-text"><i class="fas fa-search"></i></span>
                </div>
                <input class="form-control" placeholder="Search" type="text">
              </div>
            </div>
            <button type="button" class="close" data-action="search-close" data-target="#navbar-search-main" aria-label="Close">
              <span aria-hidden="true">×</span>
            </button>
          </form>
        </div>
      </div>
    </nav>
    <!-- Header -->
    <!-- Header -->
    <div class="header bg-primary pb-6">
      <div class="container-fluid">
        <div class="header-body">
          <div class="row align-items-center py-4">
            <div class="col-lg-6 col-7">
              <h6 class="h2 text-white d-inline-block mb-0">Dashboard</h6>
            </div>
          </div>
          
        </div>
      </div>
    </div>
    <!-- Page content -->
    <div class="container-fluid mt--6">
      <div class="row">
        <div class="col-xl-12 upcoming-meetings">
          <div class="card">
            <div class="card-header bg-transparent">
              <div class="row align-items-center">
                <div class="col">
                  <h6 class="text-uppercase text-muted ls-1 mb-1"></h6>
                  <h5 class="h3 mb-0">Upcoming Meetings</h5>
                </div>
                
                <div class="col text-right">
                  <button type="button" class="btn btn-primary" data-toggle="modal" data-target="#exampleModal">
                    Schedule Meeting
                  </button>
                </div>
              </div>
            </div>
            <div class="card-body">
              <div class="table-responsive">
                <div>
                  <table class="table align-items-center">
                    <thead class="thead-light">
                      <tr>
                        <th scope="col" class="sort" data-sort="name">ID</th>
                        <th scope="col" class="sort" data-sort="budget">Title</th>
                        <th scope="col">Users</th>
                        <th scope="col" class="sort" data-sort="completion">Time</th>
                        <th scope="col"></th>
                      </tr>
                    </thead>

                    <tbody class="list upcoming-meetings--table">
                     
                    </tbody>
                  </table>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="col-xl-12 recent-meetings">
          <div class="card">
            <div class="card-header bg-transparent">
              <div class="row align-items-center">
                <div class="col">
                  <h6 class="text-uppercase text-muted ls-1 mb-1"></h6>
                  <h5 class="h3 mb-0">Recent Meetings</h5>
                </div>
              </div>
            </div>
            <div class="card-body">
              <div class="table-responsive">
                <div>
                  <table class="table align-items-center">
                    <thead class="thead-light">
                      <tr>
                        <th scope="col" class="sort" data-sort="name">ID</th>
                        <th scope="col" class="sort" data-sort="budget">Title</th>
                        <th scope="col">Users</th>
                        <th scope="col" class="sort" data-sort="completion">Time</th>
                        <th scope="col"></th>
                      </tr>
                    </thead>

                    <tbody class="list recent-meetings--table">
                      
                    </tbody>
                  </table>
                </div>
              </div>
            </div>
          </div>
        </div>

      </div>
    </div>

    
  </div>


  <div class="modal fade" id="exampleModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" style="display: none;" aria-hidden="true">
    <div class="modal-dialog  modal-lg " role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="exampleModalLabel">Schedule meeting</h5>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true">×</span>
          </button>
        </div>
        <div class="modal-body">
          <div class="form-group">
            <label for="">Meeting title</label>
            <input type="text" class="form-control field-schedule--title" placeholder="Meeting title">
          </div>
          <div class="form-group">
            <label for="">Meeting Description</label>
            <textarea type="text" class="form-control field-schedule--description" rows="5" placeholder="Meeting description"></textarea>
          </div>
          <div class="form-group">
            <label for="exampleFormControlSelect2">Participants</label>
            <select multiple="" class="form-control field-schedule--participants" id="exampleFormControlSelect2">
            </select>
          </div>

          <div class="form-group">
            <label for="meeting__starttime" class="form-control-label">Start time</label>
            <input class="form-control field-schedule--starttime" type="datetime-local" value="" id="meeting__starttime">
          </div>
          

          <div class="form-group">
            <label for="meeting__starttime" class="form-control-label">End time</label>
            <input class="form-control field-schedule--endtime" type="datetime-local" value="" id="meeting__starttime">
          </div>

          <div class="form-group error-display" style="color:red">
          </div>

        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
          <button type="button" class="btn btn-primary schedule-meeting-submit">Schedule</button>
        </div>
      </div>
    </div>
  </div>

 

  </div>
  <!-- Argon Scripts -->
  <!-- Core -->
  <script src="assets/vendor/jquery/dist/jquery.min.js"></script>
  <script src="assets/vendor/bootstrap/dist/js/bootstrap.bundle.min.js"></script>
  <script src="assets/js/plugins/moment.min.js"></script>
  <script src="assets/vendor/bootstrap-datetimepicker.js"></script>
  <script src="assets/vendor/js-cookie/js.cookie.js"></script>
  <!-- <script src="assets/vendor/bootstrap-datepicker/dist/js/bootstrap-datepicker.min.js"></script> -->
  <script src="assets/js/argon.js?v=1.2.0"></script>
  <script src="assets/js/api.js"></script>
  <script type="text/javascript">
    $(function() {
      // $('#meeting__starttime').datetimepicker({});
      // $('#meeting__endtime').datetimepicker({});
    });
  </script>

</body>

</html>
