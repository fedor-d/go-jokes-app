package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/v1/jokes", getJokes)
	router.GET("/v1/jokes/:id", getJokeByID)
	router.GET("/v1/jokes/random", getRandomJoke)
	router.POST("/v1/jokes", postJokes)
	router.GET("/health/live", health)
	router.GET("/health/ready", health)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	gin.SetMode(gin.ReleaseMode)
	server.ListenAndServe()
}

// getJokes responds with the list of all jokes as JSON
func getJokes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, jokes)
}

func getRandomJoke(c *gin.Context) {
	randomJoke := jokes[rand.Intn(len(jokes))]
	c.IndentedJSON(http.StatusOK, randomJoke)
}

// postJokes adds a joke from JSON received in the request body.
func postJokes(c *gin.Context) {
	var newJoke joke

	// Call BindJSON to bind the received JSON to
	// newJoke.
	if err := c.BindJSON(&newJoke); err != nil {
		return
	}

	// Add the new joke to the slice.
	jokes = append(jokes, newJoke)
	c.IndentedJSON(http.StatusCreated, newJoke)
}

// getJokeByID locates the joke whose ID value matches the id
// parameter sent by the client, then returns that joke as a response.
func getJokeByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of jokes, looking for
	// a joke whose id value matches the parameter.
	for _, a := range jokes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "joke not found"})
}

// health handles health endpoints.
func health(c *gin.Context) {
	c.Status(http.StatusOK)
}

// joke represents data about a Chuck Norris joke
type joke struct {
	ID       string `json:"id"`
	JokeText string `json:"jokeText"`
}

// jokes are from https://github.com/springframeworkguru/chuck-norris-for-actuator/blob/master/chucknorrisforactuator/src/main/java/guru/springframework/norris/chuck/ChuckNorrisQuotes.java
var jokes = []joke{
	{ID: "1", JokeText: "Chuck Norris cannot love, he can only not kill."},
	{ID: "2", JokeText: "All browsers support the hex definitions #chuck and #norris for the colors black and blue."},
	{ID: "3", JokeText: "Chuck Norris had to stop washing his clothes in the ocean. The tsunamis were killing people."},
	{ID: "4", JokeText: "Chuck Norris can install a 64 bit OS on 32 bit machines."},
	{ID: "5", JokeText: "Fool me once, shame on you. Fool Chuck Norris once and he will roundhouse kick you in the face."},
	{ID: "6", JokeText: "Chuck Norris compresses his files by doing a flying round house kick to the hard drive."},
	{ID: "7", JokeText: "MacGyver can build an airplane out of gum and paper clips. Chuck Norris can kill him and take it."},
	{ID: "8", JokeText: "Chuck Norris played Russian Roulette with a fully loaded gun and won"},
	{ID: "9", JokeText: "Chuck Norris's keyboard doesn't have a Ctrl key because nothing controls Chuck Norris."},
	{ID: "10", JokeText: "Chuck Norris recently had the IDea to sell his urine as a canned beverage. We know this beverage as Red Bull."},
	{ID: "11", JokeText: "Chuck Norris can read from an input stream."},
	{ID: "12", JokeText: "When Chuck Norris is in a crowded area, he doesn't walk around people. He walks through them."},
	{ID: "13", JokeText: "Chuck Norris grinds his coffee with his teeth and boils the water with his own rage."},
	{ID: "14", JokeText: "The chemical formula for the highly toxic cyanIDe ion is CN-. These are also Chuck Norris' initials. This is not a coincIDence."},
	{ID: "15", JokeText: "Chuck Norris doesn't use web standards as the web will conform to him."},
	{ID: "16", JokeText: "Maslow's theory of higher needs does not apply to Chuck Norris. He only has two needs: killing people and finding people to kill."},
	{ID: "17", JokeText: "Chuck Norris dID in fact, build Rome in a day."},
	{ID: "18", JokeText: "Chuck Norris doesn't use Oracle, he is the Oracle."},
	{ID: "19", JokeText: "They say curiosity killed the cat. This is false. Chuck Norris killed the cat. Every single one of them."},
	{ID: "20", JokeText: "A Chuck Norris-delivered Roundhouse Kick is the preferred method of execution in 16 states."},
	{ID: "21", JokeText: "Chuck Norris is the only person in the world that can actually email a roundhouse kick."},
	{ID: "22", JokeText: "Product Owners never ask Chuck Norris for more features. They ask for mercy."},
	{ID: "23", JokeText: "Chuck Norris is the only person to ever win a staring contest against Ray Charles and Stevie Wonder."},
	{ID: "24", JokeText: "Chuck Norris doesn't need an OS."},
	{ID: "25", JokeText: "Chuck Norris doesn't need garbage collection because he doesn't call .Dispose(), he calls .DropKick()."},
	{ID: "26", JokeText: "Chuck Norris ordered a Big Mac at Burger King, and got one."},
	{ID: "27", JokeText: "Chuck Norris is not Politically Correct. He is just Correct. Always."},
	{ID: "28", JokeText: "One time, at band camp, Chuck Norris ate a percussionist."},
	{ID: "29", JokeText: "Chuck Norris doesn't need to use AJAX because pages are too afraID to postback anyways."},
	{ID: "30", JokeText: "Chuck Norris knows everything there is to know - Except for the definition of mercy."},
	{ID: "31", JokeText: "DivIDe Chuck Norris by zero and you will in fact get one........one bad-ass that is."},
	{ID: "32", JokeText: "Chuck Norris' unit tests don't run. They die."},
	{ID: "33", JokeText: "The Drummer for Def Leppard's only got one arm. Chuck Norris needed a back scratcher."},
	{ID: "34", JokeText: "Chuck Norris' programs never exit, they terminate."},
	{ID: "35", JokeText: "Chuck Norris can write multi-threaded applications with a single thread."},
	{ID: "36", JokeText: "Chuck Norris does not need to know about class factory pattern. He can instantiate interfaces."},
	{ID: "37", JokeText: "If you work in an office with Chuck Norris, don't ask him for his three-hole-punch."},
	{ID: "38", JokeText: "Chuck Norris uses tabasco sauce instead of visine."},
	{ID: "39", JokeText: "When Chuck Norris break the build, you can't fix it, because there is not a single line of code left."},
	{ID: "40", JokeText: "Chuck Norris can make a class that is both abstract and final."},
	{ID: "41", JokeText: "When Chuck Norris does a pushup, he isn't lifting himself up, he's pushing the Earth down."},
	{ID: "42", JokeText: "As an infant, Chuck Norris' parents gave him a toy hammer. He gave the world Stonehenge."},
	{ID: "43", JokeText: "Little Miss Muffet sat on her tuffet, until Chuck Norris roundhouse kicked her into a glacier."},
	{ID: "44", JokeText: "The quickest way to a man's heart is with Chuck Norris' fist."},
	{ID: "45", JokeText: "Code runs faster when Chuck Norris watches it."},
	{ID: "46", JokeText: "If you spell Chuck Norris in Scrabble, you win. Forever."},
	{ID: "47", JokeText: "TNT was originally developed by Chuck Norris to cure indigestion."},
	{ID: "48", JokeText: "July 4th is Independence day. And the day Chuck Norris was born. CoincIDence? I think not."},
	{ID: "49", JokeText: "In a fight between Batman and Darth Vader, the winner would be Chuck Norris."},
	{ID: "50", JokeText: "When Chuck Norris does division, there are no remainders."},
	{ID: "51", JokeText: "If Chuck Norris writes code with bugs, the bugs fix themselves."},
	{ID: "52", JokeText: "Jean-Claude Van Damme once kicked Chuck Norris' ass. He was then awakened from his dream by a roundhouse kick to the face."},
	{ID: "53", JokeText: "If you try to kill -9 Chuck Norris's programs, it backfires."},
	{ID: "54", JokeText: "Chuck Norris can unit test entire applications with a single assert."},
	{ID: "55", JokeText: "Chuck Norris knows the value of NULL, and he can sort by it too."},
	{ID: "56", JokeText: "Chuck Norris invented black. In fact, he invented the entire spectrum of visible light Except pink. Tom Cruise invented pink."},
	{ID: "57", JokeText: "Chuck Norris can lead a horse to water AND make it drink."},
	{ID: "58", JokeText: "Chuck Norris's first program was kill -9."},
	{ID: "59", JokeText: "Chuck Norris once participated in the running of the bulls. He walked."},
	{ID: "60", JokeText: "Chuck Norris's log statements are always at the FATAL level."},
	{ID: "61", JokeText: "One time, Chuck Norris accIDentally stubbed his toe. It destroyed the entire state of Ohio."},
	{ID: "62", JokeText: "Chuck Norris doesn't chew gum. Chuck Norris chews tin foil."},
	{ID: "63", JokeText: "If tapped, a Chuck Norris roundhouse kick could power the country of Australia for 44 minutes."},
	{ID: "64", JokeText: "Chuck Norris doesn't churn butter. He roundhouse kicks the cows and the butter comes straight out."},
	{ID: "65", JokeText: "Chuck Norris does not need to type-cast. The Chuck-Norris Compiler (CNC) sees through things. All way down. Always."},
	{ID: "66", JokeText: "Chuck Norris has banned rainbows from the state of North Dakota."},
	{ID: "67", JokeText: "Bill Gates thinks he's Chuck Norris. Chuck Norris actually laughed. Once."},
	{ID: "68", JokeText: "Chuck Norris once ate an entire bottle of sleeping pills. They made him blink."},
	{ID: "69", JokeText: "Chuck Norris's brain waves are suspected to be harmful to cell phones."},
	{ID: "70", JokeText: "Aliens DO indeed exist. They just know better than to visit a planet that Chuck Norris is on."},
	{ID: "71", JokeText: "What was going through the minds of all of Chuck Norris' victims before they died? His shoe."},
	{ID: "72", JokeText: "Chuck Norris can write infinite recursion functions and have them return."},
	{ID: "73", JokeText: "Chuck Norris doesn't use reflection, reflection asks politely for his help."},
	{ID: "74", JokeText: "Chuck Norris crossed the road. No one has ever dared question his motives."},
	{ID: "75", JokeText: "Chuck Norris puts his pants on one leg at a time, just like the rest of us. The only difference is, then he kills people."},
	{ID: "76", JokeText: "Chuck Norris can set ants on fire with a magnifying glass. At night."},
	{ID: "77", JokeText: "Only Chuck Norris can prevent forest fires."},
	{ID: "78", JokeText: "Chuck Norris likes his coffee half and half: half coffee grounds, half wood-grain alcohol."},
	{ID: "79", JokeText: "Ninjas want to grow up to be just like Chuck Norris. But usually they grow up just to be killed by Chuck Norris."},
	{ID: "80", JokeText: "Chuck Norris's show is called Walker: Texas Ranger, because Chuck Norris doesn't run."},
	{ID: "81", JokeText: "Simply by pulling on both ends, Chuck Norris can stretch diamonds back into coal."},
	{ID: "82", JokeText: "It takes 14 puppeteers to make Chuck Norris smile, but only 2 to make him destroy an orphanage."},
	{ID: "83", JokeText: "Chuck Norris brushes his teeth with a mixture of iron shavings, industrial paint remover and wood-grain alcohol.   "},
	{ID: "84", JokeText: "Industrial logging isn't the cause of deforestation. Chuck Norris needs toothpicks."},
	{ID: "85", JokeText: "Chuck Norris doesn't do Burn Down charts, he does Smack Down charts."},
	{ID: "86", JokeText: "If Chuck Norris wants your opinion, he'll beat it into you."},
	{ID: "87", JokeText: "On his birthday, Chuck Norris randomly selects one lucky child to be thrown into the sun."},
	{ID: "88", JokeText: "Chuck Norris doesn't bowl strikes, he just knocks down one pin and the other nine faint."},
	{ID: "89", JokeText: "According to Einstein's theory of relativity, Chuck Norris can actually roundhouse kick you yesterday."},
	{ID: "90", JokeText: "Chuck Norris' first job was as a paperboy. There were no survivors."},
	{ID: "91", JokeText: "Chuck Norris does infinite loops in 4 seconds."},
	{ID: "92", JokeText: "Chuck Norris doesn't pair program."},
	{ID: "93", JokeText: "Paper beats rock, rock beats scissors, and scissors beats paper, but Chuck Norris beats all 3 at the same time."},
	{ID: "94", JokeText: "Chuck Norris qualified with a top speed of 324 mph at the Daytona 500, without a car."},
	{ID: "95", JokeText: "Count from one to ten. That's how long it would take Chuck Norris to kill you...Fourty seven times."},
	{ID: "96", JokeText: "Chuck Norris insists on strongly-typed programming languages."},
	{ID: "97", JokeText: "All arrays Chuck Norris declares are of infinite size, because Chuck Norris knows no bounds."},
	{ID: "98", JokeText: "PresIDent Roosevelt once rode his horse 100 miles. Chuck Norris carried his the same distance in half the time."},
	{ID: "99", JokeText: "Chuck Norris doesn't actually write books, the words assemble themselves out of fear."},
	{ID: "100", JokeText: "A study showed the leading causes of death in the United States are: 1. Heart disease, 2. Chuck Norris, 3. Cancer"},
	{ID: "101", JokeText: "When Chuck Norris plays Monopoly, it affects the actual world economy."},
	{ID: "102", JokeText: "Superman once watched an episode of Walker, Texas Ranger. He then cried himself to sleep."},
	{ID: "103", JokeText: "Godzilla is a Japanese rendition of Chuck Norris' first visit to Tokyo."},
	{ID: "104", JokeText: "The class object inherits from Chuck Norris"},
	{ID: "105", JokeText: "Everybody loves Raymond. Except Chuck Norris."},
	{ID: "106", JokeText: "Chuck Norris is the reason why Waldo is hIDing."},
	{ID: "107", JokeText: "Two wrongs don't make a right. Unless you're Chuck Norris. Then two wrongs make a roundhouse kick to the face"},
	{ID: "108", JokeText: "Chuck Norris once pulled out a single hair from his beard and skewered three men through the heart with it."},
	{ID: "109", JokeText: "Police label anyone attacking Chuck Norris as a Code 45-11.... A suicIDe."},
	{ID: "110", JokeText: "Chuck Norris can slam a revolving door."},
	{ID: "111", JokeText: "Ozzy Osbourne bites the heads off of bats. Chuck Norris bites the heads off of Siberian Tigers."},
	{ID: "112", JokeText: "Chuck Norris could use anything in java.util.* to kill you, including the javadocs."},
	{ID: "113", JokeText: "When Chuck Norris goes to donate blood, he declines the syringe, and instead requests a hand gun and a bucket."},
	{ID: "114", JokeText: "Chuck Norris doesn't daydream. He's too busy giving other people nightmares."},
	{ID: "115", JokeText: "Chuck Norris doesn't need a debugger, he just stares down the bug until the code confesses."},
	{ID: "116", JokeText: "Sticks and stones may break your bones, but a Chuck Norris glare will liquefy your kIDneys."},
	{ID: "117", JokeText: "The First rule of Chuck Norris is: you do not talk about Chuck Norris."},
	{ID: "118", JokeText: "Chuck Norris doesn't wear a watch, HE decIDes what time it is."},
	{ID: "119", JokeText: "Chuck Norris causes the Windows Blue Screen of Death."},
	{ID: "120", JokeText: "Chuck Norris can compile syntax errors."},
	{ID: "121", JokeText: "Chuck Norris once kicked a horse in the chin. Its decendants are known today as Giraffes."},
	{ID: "122", JokeText: "There are no such things as tornadoes. Chuck Norris just hates trailer parks."},
	{ID: "123", JokeText: "Chuck Norris is the only man who has, literally, beaten the odds. With his fists."},
	{ID: "124", JokeText: "Chuck Norris counted to infinity. Twice."},
	{ID: "125", JokeText: "Chuck Norris can put out a fire with a gallon of gasoline."},
	{ID: "126", JokeText: "Death once had a near-Chuck-Norris experience."},
	{ID: "127", JokeText: "Chuck Norris can kill your imaginary friends."},
	{ID: "128", JokeText: "Chuck Norris can hear sign language."},
	{ID: "129", JokeText: "Chuck Norris beat the sun in a staring contest."},
	{ID: "130", JokeText: "When Graham Bell invented the telephone, he had two missed calls from Chuck Norris"},
	{ID: "131", JokeText: "Chuck Norris doesn't flush the toilet, he scares the shit out of it."},
	{ID: "132", JokeText: "Chuck Norris is so bad, when he slices onions, the onions cry."},
	{ID: "133", JokeText: "Chuck Norris knows how to exit VIM"},
	{ID: "134", JokeText: "All the codes Chuck Norris wrote is cross platform and can be run on any electronic device or quantum computer"},
}
