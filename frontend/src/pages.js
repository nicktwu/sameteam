/**
 * Created by nwu on 9/16/17.
 */
import HomeContent from './components/HomeComponent'
import AboutContent from './components/AboutComponent'
import SignupContent from './components/SignupComponent'
import LoginContent from './components/LoginComponent'
import ProfileContent from './components/ProfileComponent'
import LogoutContent from './components/LogoutComponent'

var pageList = [
  {
    url: "/",
    name: "Home",
    comp: HomeContent,
    exact: true,
    show: true
  },
  {
    url:"/about",
    name: "About",
    comp: AboutContent,
    show: true
  },
  {
    url:"/signup",
    name:"Sign Up",
    comp: SignupContent,
  },
  {
    url:"/login",
    name:"Login",
    comp: LoginContent,
  },
  {
    url:"/profile",
    name:"Profile",
    comp: ProfileContent,
    auth: true,
  },
  {
    url:"/logout",
    name:"Logout",
    comp: LogoutContent,
    auth: true,
  }
];



export default pageList;