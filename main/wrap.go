package main

import (
	"fmt"
	"github.com/eliukblau/pixterm/pkg/ansimage"
	"image/color"
	"strings"

	// "github.com/mattn/go-runewidth"
	"github.com/mrusme/gomphotherium/tui"
)

func main() {

	// imageString := ""
	firstIndentWidth := 20
	secondIndentWidth := 8
	var indent = tui.Indent{}

	plainIndent := strings.Repeat(" ", firstIndentWidth+secondIndentWidth)

	pix, err := ansimage.NewScaledFromURL(
		"https://cdn.masto.host/yharnamsingles/accounts/avatars/109/276/749/843/694/554/original/7e28a449e024114e.jpeg",
		firstIndentWidth,
		firstIndentWidth,
		color.Transparent,
		ansimage.ScaleModeResize,
		ansimage.NoDithering,
	)

	if err == nil {
		imageString := pix.RenderExt(false, false)
		splitImage := strings.Split(imageString, "\n")
		if splitImage[len(splitImage)-1] == "" {
			splitImage = splitImage[:len(splitImage)-1]
		}
		indent.InitializeWithArray(firstIndentWidth, splitImage)
	} else {
		indent.InitializeWithString(firstIndentWidth, strings.Repeat(" ", firstIndentWidth))
	}

	width := 72
	// indent := "    "

	var secondIndent tui.Indent
	secondIndent.InitializeWithString(secondIndentWidth, strings.Repeat(" ", secondIndentWidth))

	indent = *indent.ExtendWithIndent(secondIndent)
	justifyText := true
	//Print statement
	fmt.Println(plainIndent + "         1         2         3         4         5         6         7  XXXXXXX")
	fmt.Println(plainIndent + "1234567890123456789012345678901234567890123456789012345678901234567890123456789")
	fmt.Printf("\n%s", tui.WrapWithIndent("Call me Ishmael. Some years ago--never mind how long precisely--having little or no money in my purse, and nothing particular to interest me on shore, I thought I would sail about a little and see the watery part of the world. It is a way I have of driving off the spleen and regulating the circulation. Whenever I find myself growing grim about the mouth; whenever it is a damp, drizzly November in my soul; whenever I find myself involuntarily pausing before coffin warehouses, and bringing up the rear of every funeral I meet; and especially whenever my hypos get such an upper hand of me, that it requires a strong moral principle to prevent me from deliberately stepping into the street, and methodically knocking people's hats off--then, I account it high time to get to sea as soon as I can. This is my substitute for pistol and ball. With a philosophical flourish Cato throws himself upon his sword; I quietly take to the ship. There is nothing surprising in this. If they but knew it, almost all men in their degree, some time or other, cherish very nearly the same feelings towards the ocean with me.\n\nThere now is your insular city of the Manhattoes, belted round by wharves as Indian isles by coral reefs--commerce surrounds it with her surf. Right and left, the streets take you waterward. Its extreme downtown is the battery, where that noble mole is washed by waves, and cooled by breezes, which a few hours previous were out of sight of land. Look at the crowds of water-gazers there.", width, &indent, justifyText))
	fmt.Println("\n------")
	// fmt.Println(indent + "         1         2         3         4         5         6         7  XXXXXXX")
	// fmt.Println(indent + "1234567890123456789012345678901234567890123456789012345678901234567890123456789")
	// fmt.Printf("\n%s", tui.WrapWithIndent("Thisisalongword", 5, "", true))

	fmt.Println(plainIndent + "         1         2         3         4         5         6         7  XXXXXXX")
	fmt.Println(plainIndent + "1234567890123456789012345678901234567890123456789012345678901234567890123456789")
	fmt.Printf("\n%s", tui.WrapWithIndent("Hello world! #introductionI study #decentralized online social groups: organizational structure, how technology shapes human collaboration, the effects of content #moderation, the emergence of #AltTech, all that juicy stuff. I'm a #ComplexSystems #scientist (#PhDstudent), so I use a combination of #ComputerScience, #CSCW, mathematical #modeling, #NLP, and statistics to study social phenomenon. http://thisisalongurl.com/thisis/a/long/path http://thisisalongurl.com/thisis/a/long/path\n\nOutside academia, I apply that work at #DDoSecrets, a publishing collective that distributes leaked documents to journalists and researchers, and when appropriate, the public. Most of my contributions are behind the scenes (infrastructure, some #infosec, academic coordination and collaboration), so I rarely speak publicly for the collective.\n\nSometimes I post about side projects, like turning a toaster into a server, my conspiracy shrine made out of dumpster-salvaged electronics that reads and synthesizes conspiracies from the web, or adventures with obscure tech like virtual reality gopherholes. I write about my academic, DDoSecrets, and other pursuits at my blog: https://backdrifting.net/", width, &indent, justifyText))
	fmt.Println("\n------")
	// fmt.Printf("\n%s", tui.WrapWithIndent("Hello world! #introductionI study #decentralized online social groups: organizational structure, how technology shapes human collaboration, the effects of content #moderation, http://thisisalongurl.com/thisis/a/long/path http://thisisalongurl.com/thisis/a/long/path\n\n", width, indent, justifyText))

	fmt.Println(plainIndent + "         1         2         3         4         5         6         7  XXXXXXX")
	fmt.Println(plainIndent + "1234567890123456789012345678901234567890123456789012345678901234567890123456789")
	fmt.Printf("\n%s", tui.WrapWithIndent("Hello world! #introductionI study #decentral http://thisisalongurl.com/thisis/a/long/path http://thisisalongurl.com/thisis/a/long/path\n\n", width, &indent, justifyText))
	fmt.Println("\n------")

	// fmt.Println(indent + "         1         2         3         4         5         6         7  XXXXXXX")
	// fmt.Println(indent + "1234567890123456789012345678901234567890123456789012345678901234567890123456789")
	// fmt.Printf("\n%s", runewidth.Wrap("都31出シモワラ京本そごふ筆具ご強羽たへえじ熟4由ヤニマ向持エヒフ互載ノリルホ大属レサムロ現計教メ啓星70旅境ぎね談情べ防交ホ際実躍付例り。24変犯はまト宴散ホロ菌栃だ円一レドげむ暮足通ヱ玲布ルロネ会績格イじざ属紙ら本止イコ無橋講もべ法住マヒイコ読辞ぴぼぜべ米生前そら以巡勲悔殿ルースぶ。\n国サス", width))
	// fmt.Println("\n------")

	fmt.Println(plainIndent + "         1         2         3         4         5         6         7  XXXXXXX")
	fmt.Println(plainIndent + "1234567890123456789012345678901234567890123456789012345678901234567890123456789")
	fmt.Printf("\n%s", tui.WrapWithIndent("都31出シモワラ京本そごふ筆具ご強羽たへえじ熟4由ヤニマ向持エヒフ互載ノリルホ大属レサムロ現計教メ啓星70旅境ぎね談情べ防交ホ際実躍付例り。24変犯はまト宴散ホロ菌栃だ円一レドげむ暮足通ヱ玲布ルロネ会績格イじざ属紙ら本止イコ無橋講もべ法住マヒイコ読辞ぴぼぜべ米生前そら以巡勲悔殿ルースぶ。\n国サス", width, &indent, justifyText)) //首津ドはと塚気明夢書セモコ第神トソ人利小ッみど警診ちー店開氏ルユヱコ関賢反念でゅひに。1他クオニマ特事ざほ中紹ぐやせ作木ぞ政検メ活保ぽぜラせ尊不ハ識集ヌイミ阜末ねイあほ井15薄モツト判右さ現出ムヨリマ再観フおー。何ぱ両問ん軍人投タメノイ責9恵ソムヤチ刺能レテ属4団隠ろゆラ軽4問づ軽報やょげお番和だをンれ界町数ヤレヱ二緒測クみゅく川応て非生せス著著稿納月ほ。\n推め海評ぽレをへ代安スロナ克訪スヤオ運聞下ぴ死郎モテオト電16止ヤサ健総さほラレ穴話貴レネ専保ト育4択オア図部とうレな申樹メス症試版同レ。交コヒセ現覧ぽらおろ落編る今申さぞろゃ氏記大下賀クれラ済圧るき幽向ょつわぎ済問フエ使7量アク属4競覚ごっびみ。翻だえを自2金動汁ンおはゅ問帳イエ盛型ミワ安副体アス記契へたこ事高ずのぐ惑局こごたけ発戦信イ幕闘ケサホテ生逮みっほ崎提局治だむ。\n束ぶ情江日定げこさ間九で保座はほじ番拓カミリ棋指9得ス言再ムウノラ一負タマ朝39航うク政権ぐせ給合メラ変当よ子刊調ソアヘ軽理祭姿税ッ。術うレ写情ぴ全日ヒモアチ終給使ワウニタ消全こうよま代82椅87椅1世ヒニノル事満ノロ宮露久真漫腹あレ。説ヱセテキ本関ひ見時内フナム合学で竹階びよふゅ住意ソ典採てあ旅大ヱ倉合イシル図敗す目無ならゅめ世広念えべよド金察ほ心盤廃拍ひ。\n経能ぞしろは車聞活やッ宿事ぽ因園ー氏起レマツテ毎同きお眼芸あゃ報著げリ葉減カサ認腕オカアフ卓保タモラマ章1活約9智湾ぽへべぎ。笑味改メル金革ぼぶ沖影ろと自子クモウ記8持ム更全ドリ嘉的クン記都ヤ田82会族被6給事よへ営彫慰きはょで。者げぽ田個がでねレ命拳新オネ宅森こ記及ほトざく決図やッ記請カメソ浅沖リヤヲ会秀テモハシ匹公ラヒシ決組リねでン宅壊ニ防園分ホセア中状投乱ほね。\n死つほク保示げれば芸電ハトヌ後果レぱ座石ソ緩初主芸6道ごむるば公安んスつ会表百ケ祭積応ヤナモ垣済ぼラき方権ロワヤ集1施ニヱ育活姿税引リふに。転ルばラ新第つといは月臣ょぞラ親単ねすむ平化サクヲチ討組史こぞえ更社りけイレ託独ほを番九フろつ価断スリもを付季てトに衛食落イサヲ質分ヲ町時ヱヨチシ像題労えぜッ。\n権ノヤ全盆りだべー夜政ヌヱ仲文ヱヤ権出ょ者年ーッてに生再メヘカ演竹エ全使らょづッ輪着りゅは万結現あ害代戦ご将美技け決事案物なぽね。39路整ケ総節イほ辞阪ヲカ投選づがぞぽ愛12覧別ばドこぐ済外だそやく廟施タヨ供呼めだ語2読ク農端ぶらち安自玉タヒフマ下意昨も下互札永ろのょク。\n89人ロ事初ほルしド芸治っ道側ユ和取ヘミシフ面撃フ初雇れ摘山ロレ愛変7世ほ球波カ中団ぐルレろ加一明あぼさめ。食っひばぽ投集らうると持紀にク売状火対3南相やるろ絵回どぴー直場よイま投戦リホセメ提福テヱカ彼業ケ両象た党挑べけ大十コ生97性イメフ切知柱び。茂リ読情ワ党躍ヨイロ聞規定義み構化問ヘフタト謀回すぞぱ稿末ざまで適表せ景籍共西劣フ。\n4蔵い粧信シヒネ準巻向かルをろ集報陛罪ヒノル弁9紀みおいん心映ぐ試数ヨ性日ミ検定タ現捕なえぽ索誇誤郡のみそは。生競ウフツミ線能ほあでず公興ざる付衆ラとぽだ免市れ判栄ムノアヲ表87実ぼぴせ韓渡よむ自手好モツ事勧てつざね撃63給セサ月手ではこ。書住待タ紙有ヤコ頭風約そへ蔵権ミ分時スホ学東でひぽえ元百む状極鹿セアコ勢水を値防ー請禁チエ皇審ね集床鋭拡盛け。\n", width, indent, justifyText))
	fmt.Println("\n------")

	// fmt.Println(indent + "         1         2         3         4         5         6         7  XXXXXXX")
	// fmt.Println(indent + "1234567890123456789012345678901234567890123456789012345678901234567890123456789")
	// fmt.Printf("\n%s", runewidth.Wrap("都31出シモワラ京本そごふ筆具ご強羽たへえじ熟4由ヤニマ向持エヒフ互載ノリルホ大属レサムロ現計教メ啓星70旅境ぎね談情べ防交ホ際実躍付例り。24変犯はまト宴散ホロ菌栃だ円一レドげむ暮足通ヱ玲布ルロネ会績格イじざ属紙ら本止イコ無橋講もべ法住マヒイコ読辞ぴぼぜべ米生前そら以巡勲悔殿ルースぶ。\n国サス首津ドはと塚気明夢書セモコ第神トソ人利小ッみど警診ちー店開氏ルユヱコ関賢反念でゅひに。1他クオニマ特事ざほ中紹ぐやせ作木ぞ政検メ活保ぽぜラせ尊不ハ識集ヌイミ阜末ねイあほ井15薄モツト判右さ現出ムヨリマ再観フおー。何ぱ両問ん軍人投タメノイ責9恵ソムヤチ刺能レテ属4団隠ろゆラ軽4問づ軽報やょげお番和だをンれ界町数ヤレヱ二緒測クみゅく川応て非生せス著著稿納月ほ。\n推め海評ぽレをへ代安スロナ克訪スヤオ運聞下ぴ死郎モテオト電16止ヤサ健総さほラレ穴話貴レネ専保ト育4択オア図部とうレな申樹メス症試版同レ。交コヒセ現覧ぽらおろ落編る今申さぞろゃ氏記大下賀クれラ済圧るき幽向ょつわぎ済問フエ使7量アク属4競覚ごっびみ。翻だえを自2金動汁ンおはゅ問帳イエ盛型ミワ安副体アス記契へたこ事高ずのぐ惑局こごたけ発戦信イ幕闘ケサホテ生逮みっほ崎提局治だむ。\n束ぶ情江日定げこさ間九で保座はほじ番拓カミリ棋指9得ス言再ムウノラ一負タマ朝39航うク政権ぐせ給合メラ変当よ子刊調ソアヘ軽理祭姿税ッ。術うレ写情ぴ全日ヒモアチ終給使ワウニタ消全こうよま代82椅87椅1世ヒニノル事満ノロ宮露久真漫腹あレ。説ヱセテキ本関ひ見時内フナム合学で竹階びよふゅ住意ソ典採てあ旅大ヱ倉合イシル図敗す目無ならゅめ世広念えべよド金察ほ心盤廃拍ひ。\n経能ぞしろは車聞活やッ宿事ぽ因園ー氏起レマツテ毎同きお眼芸あゃ報著げリ葉減カサ認腕オカアフ卓保タモラマ章1活約9智湾ぽへべぎ。笑味改メル金革ぼぶ沖影ろと自子クモウ記8持ム更全ドリ嘉的クン記都ヤ田82会族被6給事よへ営彫慰きはょで。者げぽ田個がでねレ命拳新オネ宅森こ記及ほトざく決図やッ記請カメソ浅沖リヤヲ会秀テモハシ匹公ラヒシ決組リねでン宅壊ニ防園分ホセア中状投乱ほね。\n死つほク保示げれば芸電ハトヌ後果レぱ座石ソ緩初主芸6道ごむるば公安んスつ会表百ケ祭積応ヤナモ垣済ぼラき方権ロワヤ集1施ニヱ育活姿税引リふに。転ルばラ新第つといは月臣ょぞラ親単ねすむ平化サクヲチ討組史こぞえ更社りけイレ託独ほを番九フろつ価断スリもを付季てトに衛食落イサヲ質分ヲ町時ヱヨチシ像題労えぜッ。\n権ノヤ全盆りだべー夜政ヌヱ仲文ヱヤ権出ょ者年ーッてに生再メヘカ演竹エ全使らょづッ輪着りゅは万結現あ害代戦ご将美技け決事案物なぽね。39路整ケ総節イほ辞阪ヲカ投選づがぞぽ愛12覧別ばドこぐ済外だそやく廟施タヨ供呼めだ語2読ク農端ぶらち安自玉タヒフマ下意昨も下互札永ろのょク。\n89人ロ事初ほルしド芸治っ道側ユ和取ヘミシフ面撃フ初雇れ摘山ロレ愛変7世ほ球波カ中団ぐルレろ加一明あぼさめ。食っひばぽ投集らうると持紀にク売状火対3南相やるろ絵回どぴー直場よイま投戦リホセメ提福テヱカ彼業ケ両象た党挑べけ大十コ生97性イメフ切知柱び。茂リ読情ワ党躍ヨイロ聞規定義み構化問ヘフタト謀回すぞぱ稿末ざまで適表せ景籍共西劣フ。\n4蔵い粧信シヒネ準巻向かルをろ集報陛罪ヒノル弁9紀みおいん心映ぐ試数ヨ性日ミ検定タ現捕なえぽ索誇誤郡のみそは。生競ウフツミ線能ほあでず公興ざる付衆ラとぽだ免市れ判栄ムノアヲ表87実ぼぴせ韓渡よむ自手好モツ事勧てつざね撃63給セサ月手ではこ。書住待タ紙有ヤコ頭風約そへ蔵権ミ分時スホ学東でひぽえ元百む状極鹿セアコ勢水を値防ー請禁チエ皇審ね集床鋭拡盛け。\n", width))
	// fmt.Println("\n------")

	fmt.Println(plainIndent + "         1         2         3         4         5         6         7  XXXXXXX")
	fmt.Println(plainIndent + "1234567890123456789012345678901234567890123456789012345678901234567890123456789")
	fmt.Printf("\n%s", tui.WrapWithIndent("都31出シモワラ京本そごふ筆具ご強羽たへえじ熟4由ヤニマ向持エヒフ互載ノリルホ大属レサムロ現計教メ啓星70旅境ぎね談情べ防交ホ際実躍付例り。24変犯はまト宴散ホロ菌栃だ円一レドげむ暮足通ヱ玲布ルロネ会績格イじざ属紙ら本止イコ無橋講もべ法住マヒイコ読辞ぴぼぜべ米生前そら以巡勲悔殿ルースぶ。\n国サス首津ドはと塚気明夢書セモコ第神トソ人利小ッみど警診ちー店開氏ルユヱコ関賢反念でゅひに。1他クオニマ特事ざほ中紹ぐやせ作木ぞ政検メ活保ぽぜラせ尊不ハ識集ヌイミ阜末ねイあほ井15薄モツト判右さ現出ムヨリマ再観フおー。何ぱ両問ん軍人投タメノイ責9恵ソムヤチ刺能レテ属4団隠ろゆラ軽4問づ軽報やょげお番和だをンれ界町数ヤレヱ二緒測クみゅく川応て非生せス著著稿納月ほ。\n推め海評ぽレをへ代安スロナ克訪スヤオ運聞下ぴ死郎モテオト電16止ヤサ健総さほラレ穴話貴レネ専保ト育4択オア図部とうレな申樹メス症試版同レ。交コヒセ現覧ぽらおろ落編る今申さぞろゃ氏記大下賀クれラ済圧るき幽向ょつわぎ済問フエ使7量アク属4競覚ごっびみ。翻だえを自2金動汁ンおはゅ問帳イエ盛型ミワ安副体アス記契へたこ事高ずのぐ惑局こごたけ発戦信イ幕闘ケサホテ生逮みっほ崎提局治だむ。\n束ぶ情江日定げこさ間九で保座はほじ番拓カミリ棋指9得ス言再ムウノラ一負タマ朝39航うク政権ぐせ給合メラ変当よ子刊調ソアヘ軽理祭姿税ッ。術うレ写情ぴ全日ヒモアチ終給使ワウニタ消全こうよま代82椅87椅1世ヒニノル事満ノロ宮露久真漫腹あレ。説ヱセテキ本関ひ見時内フナム合学で竹階びよふゅ住意ソ典採てあ旅大ヱ倉合イシル図敗す目無ならゅめ世広念えべよド金察ほ心盤廃拍ひ。\n経能ぞしろは車聞活やッ宿事ぽ因園ー氏起レマツテ毎同きお眼芸あゃ報著げリ葉減カサ認腕オカアフ卓保タモラマ章1活約9智湾ぽへべぎ。笑味改メル金革ぼぶ沖影ろと自子クモウ記8持ム更全ドリ嘉的クン記都ヤ田82会族被6給事よへ営彫慰きはょで。者げぽ田個がでねレ命拳新オネ宅森こ記及ほトざく決図やッ記請カメソ浅沖リヤヲ会秀テモハシ匹公ラヒシ決組リねでン宅壊ニ防園分ホセア中状投乱ほね。\n死つほク保示げれば芸電ハトヌ後果レぱ座石ソ緩初主芸6道ごむるば公安んスつ会表百ケ祭積応ヤナモ垣済ぼラき方権ロワヤ集1施ニヱ育活姿税引リふに。転ルばラ新第つといは月臣ょぞラ親単ねすむ平化サクヲチ討組史こぞえ更社りけイレ託独ほを番九フろつ価断スリもを付季てトに衛食落イサヲ質分ヲ町時ヱヨチシ像題労えぜッ。\n権ノヤ全盆りだべー夜政ヌヱ仲文ヱヤ権出ょ者年ーッてに生再メヘカ演竹エ全使らょづッ輪着りゅは万結現あ害代戦ご将美技け決事案物なぽね。39路整ケ総節イほ辞阪ヲカ投選づがぞぽ愛12覧別ばドこぐ済外だそやく廟施タヨ供呼めだ語2読ク農端ぶらち安自玉タヒフマ下意昨も下互札永ろのょク。\n89人ロ事初ほルしド芸治っ道側ユ和取ヘミシフ面撃フ初雇れ摘山ロレ愛変7世ほ球波カ中団ぐルレろ加一明あぼさめ。食っひばぽ投集らうると持紀にク売状火対3南相やるろ絵回どぴー直場よイま投戦リホセメ提福テヱカ彼業ケ両象た党挑べけ大十コ生97性イメフ切知柱び。茂リ読情ワ党躍ヨイロ聞規定義み構化問ヘフタト謀回すぞぱ稿末ざまで適表せ景籍共西劣フ。\n4蔵い粧信シヒネ準巻向かルをろ集報陛罪ヒノル弁9紀みおいん心映ぐ試数ヨ性日ミ検定タ現捕なえぽ索誇誤郡のみそは。生競ウフツミ線能ほあでず公興ざる付衆ラとぽだ免市れ判栄ムノアヲ表87実ぼぴせ韓渡よむ自手好モツ事勧てつざね撃63給セサ月手ではこ。書住待タ紙有ヤコ頭風約そへ蔵権ミ分時スホ学東でひぽえ元百む状極鹿セアコ勢水を値防ー請禁チエ皇審ね集床鋭拡盛け。\n", width, &indent, justifyText))
	fmt.Println("\n------")

	fmt.Println(plainIndent + "         1         2         3         4         5         6         7  XXXXXXX")
	fmt.Println(plainIndent + "1234567890123456789012345678901234567890123456789012345678901234567890123456789")
	fmt.Printf("\n%s", tui.WrapWithIndent("都31出シモワラ京本そごふ筆具ご強羽たへえじ熟4由ヤニマ向持エヒフ互載ノリルホ大属レサムロ現計教メ啓星70旅境ぎね談情べ防交ホ際実躍付例り。24変犯はまト宴散ホロ菌栃だ円一レドげむ暮足通ヱ玲布ルロネ会績格イじざ属紙ら本止イコ無橋講もべ法住マヒイコ読辞ぴぼぜべ米生前そら以巡勲悔殿ルースぶ。\n国サス首津ドはと塚気明夢書セモコ第神トソ人利小ッみど警診ちー店開氏ルユヱコ関賢反念でゅひに。1他クオニマ特事ざほ中紹ぐやせ作木ぞ政検メ活保ぽぜラせ尊不ハ識集ヌイミ阜末ねイあほ井15薄モツト判右さ現出ムヨリマ再観フおー。何ぱ両問ん軍人投タメノイ責9恵ソムヤチ刺能レテ属4団隠ろゆラ軽4問づ軽報やょげお番和だをンれ界町数ヤレヱ二緒測クみゅく川応て非生せス著著稿納月ほ。\n推め海評ぽレをへ代安スロナ克訪スヤオ運聞下ぴ死郎モテオト電16止ヤサ健総さほラレ穴話貴レネ専保ト育4択オア図部とうレな申樹メス症試版同レ。交コヒセ現覧ぽらおろ落編る今申さぞろゃ氏記大下賀クれラ済圧るき幽向ょつわぎ済問フエ使7量アク属4競覚ごっびみ。翻だえを自2金動汁ンおはゅ問帳イエ盛型ミワ安副体アス記契へたこ事高ずのぐ惑局こごたけ発戦信イ幕闘ケサホテ生逮みっほ崎提局治だむ。\n束ぶ情江日定げこさ間九で保座はほじ番拓カミリ棋指9得ス言再ムウノラ一負タマ朝39航うク政権ぐせ給合メラ変当よ子刊調ソアヘ軽理祭姿税ッ。術うレ写情ぴ全日ヒモアチ終給使ワウニタ消全こうよま代82椅87椅1世ヒニノル事満ノロ宮露久真漫腹あレ。説ヱセテキ本関ひ見時内フナム合学で竹階びよふゅ住意ソ典採てあ旅大ヱ倉合イシル図敗す目無ならゅめ世広念えべよド金察ほ心盤廃拍ひ。\n経能ぞしろは車聞活やッ宿事ぽ因園ー氏起レマツテ毎同きお眼芸あゃ報著げリ葉減カサ認腕オカアフ卓保タモラマ章1活約9智湾ぽへべぎ。笑味改メル金革ぼぶ沖影ろと自子クモウ記8持ム更全ドリ嘉的クン記都ヤ田82会族被6給事よへ営彫慰きはょで。者げぽ田個がでねレ命拳新オネ宅森こ記及ほトざく決図やッ記請カメソ浅沖リヤヲ会秀テモハシ匹公ラヒシ決組リねでン宅壊ニ防園分ホセア中状投乱ほね。\n死つほク保示げれば芸電ハトヌ後果レぱ座石ソ緩初主芸6道ごむるば公安んスつ会表百ケ祭積応ヤナモ垣済ぼラき方権ロワヤ集1施ニヱ育活姿税引リふに。転ルばラ新第つといは月臣ょぞラ親単ねすむ平化サクヲチ討組史こぞえ更社りけイレ託独ほを番九フろつ価断スリもを付季てトに衛食落イサヲ質分ヲ町時ヱヨチシ像題労えぜッ。\n権ノヤ全盆りだべー夜政ヌヱ仲文ヱヤ権出ょ者年ーッてに生再メヘカ演竹エ全使らょづッ輪着りゅは万結現あ害代戦ご将美技け決事案物なぽね。39路整ケ総節イほ辞阪ヲカ投選づがぞぽ愛12覧別ばドこぐ済外だそやく廟施タヨ供呼めだ語2読ク農端ぶらち安自玉タヒフマ下意昨も下互札永ろのょク。\n89人ロ事初ほルしド芸治っ道側ユ和取ヘミシフ面撃フ初雇れ摘山ロレ愛変7世ほ球波カ中団ぐルレろ加一明あぼさめ。食っひばぽ投集らうると持紀にク売状火対3南相やるろ絵回どぴー直場よイま投戦リホセメ提福テヱカ彼業ケ両象た党挑べけ大十コ生97性イメフ切知柱び。茂リ読情ワ党躍ヨイロ聞規定義み構化問ヘフタト謀回すぞぱ稿末ざまで適表せ景籍共西劣フ。\n4蔵い粧信シヒネ準巻向かルをろ集報陛罪ヒノル弁9紀みおいん心映ぐ試数ヨ性日ミ検定タ現捕なえぽ索誇誤郡のみそは。生競ウフツミ線能ほあでず公興ざる付衆ラとぽだ免市れ判栄ムノアヲ表87実ぼぴせ韓渡よむ自手好モツ事勧てつざね撃63給セサ月手ではこ。書住待タ紙有ヤコ頭風約そへ蔵権ミ分時スホ学東でひぽえ元百む状極鹿セアコ勢水を値防ー請禁チエ皇審ね集床鋭拡盛け。\n", width, &indent, justifyText))
	fmt.Println("\n------")

	// fmt.Println(indent + "         1         2         3         4         5         6         7  XXXXXXX8XXXXXXXXX8XXXXXXXXX8XXXXXXXXX8XXXXXXXXX8XXXXXXXXX8XXXXXXXXX8XXXXXXXXX8XXXXXXXXX8")
	// fmt.Println(indent + "123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
	// fmt.Printf("\n%s", tui.WrapWithIndent("I now have this pinned to my keyboard clipboard so I can deploy it as needed, which is often: \"We have somehow all owed a narrow faction of our politics to convince large portions of the populace that “standards,” “rules,” and “conduct expectations” within a given community are tantamount to state censorship. It’s disappointing how many rubes have fallen for this meme.\"I sometimes insert \"like you\" after \"rubes.\"", 143, indent, justifyText))
	// fmt.Println("\n------")
}
