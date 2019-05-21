/*
   conflux - Distributed database synchronization library
	Based on the algorithm described in
		"Set Reconciliation with Nearly Optimal	Communication Complexity",
			Yaron Minsky, Ari Trachtenberg, and Richard Zippel, 2004.

   Copyright (c) 2012-2015  Casey Marshall <cmars@cmarstech.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, version 3.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package testing

import (
	cf "hockeypuck/conflux"
)

var PtreeSplits85 = []*cf.Zp{
	cf.Zs(cf.P_SKS, "54945054303302140323349777569652159744"),
	cf.Zs(cf.P_SKS, "301824390735659941098168552847110967299"),
	cf.Zs(cf.P_SKS, "244727299682701342167768131406454086662"),
	cf.Zs(cf.P_SKS, "246090505779456321483693264299682785547"),
	cf.Zs(cf.P_SKS, "132666079786438034357736690869858972940"),
	cf.Zs(cf.P_SKS, "77943753696469936247570506393661277454"),
	cf.Zs(cf.P_SKS, "23196332603806579862361573649796228117"),
	cf.Zs(cf.P_SKS, "4716213446676942567518507102644048922"),
	cf.Zs(cf.P_SKS, "296958268935570641352142910860288566305"),
	cf.Zs(cf.P_SKS, "68302849918166164850536779468406147620"),
	cf.Zs(cf.P_SKS, "218289647857890898753469351137063165732"),
	cf.Zs(cf.P_SKS, "49493891938832871357367248144831830823"),
	cf.Zs(cf.P_SKS, "48291393928127338850139259452183259432"),
	cf.Zs(cf.P_SKS, "86222875303132381523404440898289042729"),
	cf.Zs(cf.P_SKS, "141808492182097342190378424004993438762"),
	cf.Zs(cf.P_SKS, "335945044879925574404988722388729177130"),
	cf.Zs(cf.P_SKS, "236652262168326400305360829310456888110"),
	cf.Zs(cf.P_SKS, "251126436454877830379605469450395348293"),
	cf.Zs(cf.P_SKS, "326459174278123784017559830530989911369"),
	cf.Zs(cf.P_SKS, "33880488004771349616788442031172397900"),
	cf.Zs(cf.P_SKS, "52260317194205405422862037231440764495"),
	cf.Zs(cf.P_SKS, "111966997212661832286471884984785895763"),
	cf.Zs(cf.P_SKS, "65346596413500442902105433236945513815"),
	cf.Zs(cf.P_SKS, "336173260736631828030752156996804317528"),
	cf.Zs(cf.P_SKS, "12549734565921827638994232318473088346"),
	cf.Zs(cf.P_SKS, "236987695883303042196824159729755015004"),
	cf.Zs(cf.P_SKS, "323119199726242591092367701897211489888"),
	cf.Zs(cf.P_SKS, "41584217555254899291679977414635303265"),
	cf.Zs(cf.P_SKS, "118348940079828526226895879854757946209"),
	cf.Zs(cf.P_SKS, "76747152908027187555920496063053721192"),
	cf.Zs(cf.P_SKS, "141520615381718155000336085124075411305"),
	cf.Zs(cf.P_SKS, "12644319991398156892717186701440448617"),
	cf.Zs(cf.P_SKS, "6412217881018912834671730377842624107"),
	cf.Zs(cf.P_SKS, "40272626613761839643433023788839454318"),
	cf.Zs(cf.P_SKS, "317014674893435005836306380764370532983"),
	cf.Zs(cf.P_SKS, "311101531782501702629491213596113502586"),
	cf.Zs(cf.P_SKS, "188716858420292079269415903308294938757"),
	cf.Zs(cf.P_SKS, "59683129049585326195019115368974177413"),
	cf.Zs(cf.P_SKS, "120554134733908208956936621136810387334"),
	cf.Zs(cf.P_SKS, "118284951518112845084965537305254483334"),
	cf.Zs(cf.P_SKS, "334377276241926936152366343232870849927"),
	cf.Zs(cf.P_SKS, "77310576102193849850198786368460166024"),
	cf.Zs(cf.P_SKS, "48414329405458959169482475701469749386"),
	cf.Zs(cf.P_SKS, "250015874231309260492838193181798330515"),
	cf.Zs(cf.P_SKS, "159774163460744271274851987570681581206"),
	cf.Zs(cf.P_SKS, "143198666233767249362774217511755743896"),
	cf.Zs(cf.P_SKS, "240912385064508366145523910482350846106"),
	cf.Zs(cf.P_SKS, "137051103418540437753971021562365500061"),
	cf.Zs(cf.P_SKS, "327716597715155576166988551942487950757"),
	cf.Zs(cf.P_SKS, "183616583934478418670273555233235445158"),
	cf.Zs(cf.P_SKS, "205743196898220792903183384302974895272"),
	cf.Zs(cf.P_SKS, "186204001228737102171059261656423711656"),
	cf.Zs(cf.P_SKS, "252570316860507925711101109364354724777"),
	cf.Zs(cf.P_SKS, "244999449118738179307733396365253793195"),
	cf.Zs(cf.P_SKS, "8050877573458140058716086496555437487"),
	cf.Zs(cf.P_SKS, "53287660286265266199637380204079727539"),
	cf.Zs(cf.P_SKS, "20018245473587526137270198427261386422"),
	cf.Zs(cf.P_SKS, "171075676243784157336190498877480108992"),
	cf.Zs(cf.P_SKS, "325970951673190825762721854817698148039"),
	cf.Zs(cf.P_SKS, "248360922124393115642405238390171529928"),
	cf.Zs(cf.P_SKS, "285728769080344321852508208337475846600"),
	cf.Zs(cf.P_SKS, "142870176296995226271711928223105046985"),
	cf.Zs(cf.P_SKS, "46264038433717262828185678231991802569"),
	cf.Zs(cf.P_SKS, "259502011074668048058351875079612836042"),
	cf.Zs(cf.P_SKS, "38755373377015903993868338342907167182"),
	cf.Zs(cf.P_SKS, "211463311137941721160830773248262076623"),
	cf.Zs(cf.P_SKS, "220668689716423815468107223827565625040"),
	cf.Zs(cf.P_SKS, "278495691295316092131115910823992635857"),
	cf.Zs(cf.P_SKS, "290108658999174932224227283677350406611"),
	cf.Zs(cf.P_SKS, "49288527855085557669219340700286100693"),
	cf.Zs(cf.P_SKS, "208585055893782380269169696681176208600"),
	cf.Zs(cf.P_SKS, "175869353869268402931895766549932311261"),
	cf.Zs(cf.P_SKS, "52452224898726559179038736954156221661"),
	cf.Zs(cf.P_SKS, "106012643650337984639895849375969307870"),
	cf.Zs(cf.P_SKS, "105919911020868424250639020684116302309"),
	cf.Zs(cf.P_SKS, "105424832931958607019255672786388535013"),
	cf.Zs(cf.P_SKS, "43048960869965803184334424644920615143"),
	cf.Zs(cf.P_SKS, "332850420638398765868853355107487298026"),
	cf.Zs(cf.P_SKS, "274284651981540161475649614117607611627"),
	cf.Zs(cf.P_SKS, "305496626634797764740487407455046400238"),
	cf.Zs(cf.P_SKS, "104161670190828326738629254941040826870"),
	cf.Zs(cf.P_SKS, "233126962571760965470316713449937861113"),
	cf.Zs(cf.P_SKS, "128034845779765006459409039712734820092"),
	cf.Zs(cf.P_SKS, "29360578736782464072123892399520131324"),
	cf.Zs(cf.P_SKS, "32934118080763265448914909383880829695"),
}
