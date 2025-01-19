package db

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Body   string `json:"body"`
}

var Quotes = []Quote{
	{ID: 1, Author: "Shayar Jakhmi", Body: "Koi samjhe toh sahi, koi na samjhe toh bhi sahi, ye raaz-e-mohabbat hai, jise samjhe woh kya samjhe, jise na samjhe woh kya samjhe."},
	{ID: 2, Author: "Shayar Bedardi", Body: "Dil ki baat zuban pe aaye to kya kare, yun hi bechaini hai dhadkan me, ishq laaye to kya kare."},
	{ID: 1, Author: "Kavi Albela", Body: "Koi deewana kehta hai, koi pagal samajhta hai, magar dharti ki bechaini ko bas badal samajhta hai."},
	{ID: 1, Author: "Vyas Chailla", Body: "Doob kar jisme marr jaayein woh nadiyan hi achhi hain, baaki sab toh bekaar hain, jo bas paani pe chal jaayein."},
	{ID: 1, Author: "Vyas Rasiya", Body: "Kabhi kabhi kuch jeetne ke liye kuch haarna bhi padta hai, aur haar kar jeetne wale ko baazigar kehte hain."},
}
