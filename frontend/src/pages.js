/**
 * Created by nwu on 9/16/17.
 */
import HomeContent from './components/HomeComponent'
import AboutContent from './components/AboutComponent'

var pageList = [
  {
    url: "/",
    name: "Home",
    comp: HomeContent,
    exact: true
  },
  {
    url:"/about",
    name: "About",
    comp: AboutContent
  }
];

export default pageList;