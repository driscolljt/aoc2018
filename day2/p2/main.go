package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// given 'data'
	data := data()

	for _, d := range data {
		tmp := rmvFromSlice(data, d)
		for _, t := range tmp {
			diffs := compareStrings(d, t)
			if diffs == 1 {
				var buf bytes.Buffer
				a := []rune(d)
				b := []rune(t)
				for i := 0; i < len(d); i++ {
					if a[i] == b[i] {
						buf.WriteString(string(a[i]))
					}
				}
				fmt.Println(buf.String())
			}
		}
	}
	// compare values to find
	// a pair with all but 1 char
	// in the same place
	// Ex. 'abcde' + 'abxde'
	// return the common chars 'abde'
}

// return count of diff chars
// Ex. 'abcde' + 'xbcye' returns 2
func compareStrings(s1 string, s2 string) int {
	a := len(s1)
	b := len(s2)
	shortest := 0
	if a > b {
		shortest = b
	} else {
		shortest = a
	}

	s1r := []rune(s1)
	s2r := []rune(s2)
	countDiff := 0
	for i := 0; i < shortest; i++ {
		if s1r[i] != s2r[i] {
			countDiff++
		}
	}
	return countDiff
}

func rmvFromSlice(slice []string, s string) []string {
	for i, v := range slice {
		if v == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func data() []string {
	// return data as formatted slice of strings
	s := "wlpiogsvdfecjdqmnxakudrhbz\r\nwbpioesvyfecjuqmlxaktdrhbz\r\nblviogavyfecjuqmnxaktdrhbz\r\nwlpiogsvydecjuqmnipktdrhbz\r\nwlwiogsvyfmcjuqmoxaktdrhbz\r\nwlpiogsvphecjuqmnxaktdrzbz\r\nwlpiogsvyfecjuqmnkakhdrkbz\r\nwlpiogsvyfhcjuqmnxxktddhbz\r\nwlpiogsvyfccfuqmnxgktdrhbz\r\nwlpiogsvhmecjvqmnxaktdrhbz\r\nwlpiogsvyfecjdqqnxaktdrhyz\r\nwlpiogyvycecjaqmnxaktdrhbz\r\nwlpiogsvyfecjfqmnxaktybhbz\r\nwkpiogsvyfmcauqmnxaktdrhbz\r\nwlmiolsvyfecjuqmnxaktdrhbn\r\nwlpioksvyfecjuqmnxaktdrhgs\r\nwlpiogsvyflcjuvmnxsktdrhbz\r\nwgziogsvyfecjuqmnxaktdrhoz\r\nwhpingsvyeecjuqmnxaktdrhbz\r\nwlpiogsvyfecjuqgnxaktdvhlz\r\nwlpioasvtfecjuqmnxaktdahbz\r\nwlpihgsvefeceuqmnxaktdrhbz\r\nwlpiogsvyfecyuqwnxaktdghbz\r\nwlpfsgsvyfhcjuqmnxaktdrhbz\r\nwlpiogyvyfecjuqmnxalpdrhbz\r\nwlpiogsvyfescsqmnxaktdrhbz\r\nwluiogsvyfecytqmnxaktdrhbz\r\nwltiorsvyfecjuqmoxaktdrhbz\r\nwlmiogwvyfejjuqmnxaktdrhbz\r\nwlpiogsvyfycjuumnxvktdrhbz\r\nwlkiogsqyfecjqqmnxaktdrhbz\r\nwlpiogsvyfecouqmnxaktfrubz\r\nhupiogsvlfecjuqmnxaktdrhbz\r\nwlpiogsvpxecjuqmnxaksdrhbz\r\nwlpiogsvyfkcjfqmnxxktdrhbz\r\nwlpiogsjyfecjuqnnxakthrhbz\r\nwlpiogsvyfecjuqmnxaktddhdk\r\nwlpipgsvyfecjuqmnhaktdrubz\r\nwlpiogsoyfenjpqmnxaktdrhbz\r\nwlpiogsvyfecnuqmnxaxtdrmbz\r\nwlpiggsvyfjcjuumnxaktdrhbz\r\nwlppogsvyfecjuqmnxautdhhbz\r\nwlpiovbvyfecjuqmnxaktcrhbz\r\nwlpiogsvyfecjoqmnxaktdrobu\r\nwlpiohsvyfecjugmnxakthrhbz\r\nwvpiogovyfecjuqmnxakwdrhbz\r\nwlpiogsbyfecjuqmnxdktdrtbz\r\nwlpnogsvyfecjuqmnxakykrhbz\r\nwlpiogpvyfecjuqmnxvktdrhbs\r\nwlpiogsvkvecjuqmnxadtdrhbz\r\nwlkihgsvyfecjuqmnxlktdrhbz\r\nwlpilgsvyfhcjuqmnxakudrhbz\r\nwlpioksvysgcjuqmnxaktdrhbz\r\nwlpiogsvyfecorqmoxaktdrhbz\r\nwlpiogsvyfectzlmnxaktdrhbz\r\nwlpiogsvywecjuqwnqaktdrhbz\r\nwlpiowsvyfecjuqrnxaftdrhbz\r\nwlpiogsuyfecjutmnxaktnrhbz\r\nwepiowsvyfqcjuqmnxaktdrhbz\r\nwlpirgssyfecjuqmvxaktdrhbz\r\nwlyiogstyfecjuqmnxaktdrhbw\r\nwlpiogseyfecauqmnxaktdjhbz\r\nwlpioisvyfenjuqmnxakrdrhbz\r\nwlpiogsvyfrgjfqmnxaktdrhbz\r\nwlpionsvyfecjmqmjxaktdrhbz\r\nalpiggsvyfecjuqmnxaktdrhkz\r\nwlphngsvyfecjuqmnxfktdrhbz\r\nwlpiogsvyferjuqmlxakttrhbz\r\nwlniogsvefecjuqsnxaktdrhbz\r\nwlpiogsvyfncjurmnxakudrhbz\r\nwlpiogsvyfecjuqmnxaktprlaz\r\nwlpiocsvyfecjupmkxaktdrhbz\r\nwlpihgsvyfecjeqfnxaktdrhbz\r\nwlwioasvyfjcjuqmnxaktdrhbz\r\nwlpifgsvyfecjuqsnxaktdshbz\r\nwlxiogsvyrechuqmnxaktdrhbz\r\nwlpiogovyfxcjuqmnxakkdrhbz\r\nwlpiogsvyfecjkqmdxaktmrhbz\r\nwlpiogsvyfecjwqmntaktdhhbz\r\nwlpiogsvdfecjuqmmxaktbrhbz\r\nwlpiogsvyfecauqmnxaksdrhwz\r\nwlpiogsvwfecjuqznxaktorhbz\r\nwlpiogtvyfecjuqhnxakidrhbz\r\nwlpiogsvyyecjuqmnxaktdrhwt\r\nwljiogsvyfecfuqbnxaktdrhbz\r\nwlpiogsvybecjuqmnxagtdrjbz\r\nwrpiogsvyfecjuqmnuaktdrhbd\r\nwlpiogsvyfecjurmnxnltdrhbz\r\nblpvogsvyaecjuqmnxaktdrhbz\r\nbfpiogyvyfecjuqmnxaktdrhbz\r\nwlpiogsvyfecjuqinxaknddhbz\r\nwlpizgsvvfecjuqxnxaktdrhbz\r\nglpiogsvyrecjuqmnxaktdrhbr\r\nwlpiogskhfecjutmnxaktdrhbz\r\nwlpiogsvyfecmuqmnxaktdribc\r\nwlpioesvwfecjuqmnxakkdrhbz\r\nwlpionsrafecjuqmnxaktdrhbz\r\nwlsiogsvyfecjuqmnaaktdrhvz\r\nbloiogsvyfecjuqmnxakjdrhbz\r\nwlpiogsvyfecjuqenmastdrhbz\r\nwlpiogyvyfecjuqmuxakldrhbz\r\nplpiogovyfecjuvmnxaktdrhbz\r\nwlpiogsvyfebjuqmnkakvdrhbz\r\nwlziogsvyfhcjuqmngaktdrhbz\r\nwlsiogsvyfecjuqmnxaktdrzjz\r\nplbiogsvyfecfuqmnxaktdrhbz\r\nwfpiogsvyfecjuqknxaktdrhiz\r\nwlpiogjbyfecjuqmnxaktprhbz\r\nwmpiogsvyrecjcqmnxaktdrhbz\r\nwlpiogsyyfecjuqmqxaktdrbbz\r\nwlpiogsvyfecjuqknxaetdehbz\r\nwlpiogsvyfezjuqmnxakkdhhbz\r\nwlpiogsvyfecjjqvnxaktdrhiz\r\nwkpiogsvyfucjuqmnxaktdrhbd\r\nlliiogsvyfecjuqmnxaktdrhoz\r\nwlpiogsvyfecjuqmsxdktdshbz\r\nwlprogtvyfecjuqmnxaktvrhbz\r\nwlpizgssyffcjuqmnxaktdrhbz\r\nwlpioasvyfvcjuqmnxakldrhbz\r\nwlpoogsvyyecjuqmnxastdrhbz\r\nwlpiognvyfecjuqmnsaktdrhbr\r\nwlpiogsoyfecjuqmnxaktdrhho\r\nwfpiogsvydecjuqmnxaotdrhbz\r\nwlpiogsvqhecjuqmnxaktdrhhz\r\nwkpiogsvyfeojuqmnxaktdrqbz\r\nwlpiogsvyfeveuqmnxaktdshbz\r\nwlpiogbvyfecjuqmexaktdrcbz\r\nwlpxogsvyfehjsqmnxaktdrhbz\r\nwlpcogsvyfecjuqmnjakttrhbz\r\nwlpiogsvvkecjuqmnxaftdrhbz\r\nwlpiogsvffecnuqmnxaktdnhbz\r\nwlpiogsvyfecjupjnxaktdrhbr\r\nwlpqogsvyfecjuqmnxlktdphbz\r\nwlpxogsvyfecjvqmnxaktirhbz\r\nelpiogsvyfecjuqlnxaqtdrhbz\r\nwspiogsvrfecjuqmnxakadrhbz\r\nwlpiogsmyfecbuqmnxactdrhbz\r\nwlpiogsvyfecauqmnyakzdrhbz\r\nwlsiogsvyfecjuqmnxakvdrnbz\r\nwlpiogsxyfeijuqmnxakndrhbz\r\nwlpiogsvyfecjuumnxakbyrhbz\r\nwlpiogsvyfecjuqmnxhktdyhbo\r\nwlpiogsvyfecjuqqnxnjtdrhbz\r\nwapiogsvyfecjuqmnxoktdrmbz\r\nwlpiogsvyfeejvqmnxaktdrubz\r\nwlpitgsvyeectuqmnxaktdrhbz\r\nalpiogsvyfecjulmnxaktdchbz\r\nwlpiogsvyfecjuqmuxoktdrwbz\r\nwlpiogsvyfzgjuhmnxaktdrhbz\r\nwlpnogsvyfecjuqmdxaktyrhbz\r\nwlpiogsvyfecjuqmnxakthrhra\r\nwliiogsvyfecluqmnxaktdhhbz\r\nwlpiogsvyfecjuymnxaltdrhwz\r\nwlpiogsvyfeljuqmnxaktyrhbd\r\nwlpiygsvvfecjuqmfxaktdrhbz\r\nwlpiogihsfecjuqmnxaktdrhbz\r\nwlpiogjvyfecjuqmnhuktdrhbz\r\nwldiogsvyfecjiqmwxaktdrhbz\r\nwlpiogsvjfecjuqmnxaktdrgbr\r\nwlpioisvyfecjuqwnxabtdrhbz\r\nwlviogsvyfscjuqmnxqktdrhbz\r\nwlpiogsvyfecjuqmuxakbdrubz\r\nwlpiogsvyfecjuqmnxmatdrhqz\r\nwlpiogsvyfbcjuqwmxaktdrhbz\r\nwlpiogsvyfexjuqmnxuxtdrhbz\r\nwljiogsvbfecjuqmnxartdrhbz\r\nwlpvogsvyfeujuqmnxaktdmhbz\r\nwnpiogsvyfekjuqanxaktdrhbz\r\nwlprogsvyfecjuqmzxvktdrhbz\r\nwkpiogvvyfecjuqmnxaktdrabz\r\nwlpiogsvwfecjuqmnxaktkbhbz\r\nwlpiogsvyfecjlqmnxtttdrhbz\r\nwlpioqsvyfecjuqznxaktyrhbz\r\nwlpiogsvyfecjuqmnxnethrhbz\r\nwlpiogsyyfgcjuqmnxaktdrhbm\r\nwlpiopsvbfecjuqmnxaktdlhbz\r\nwloqogsvyfucjuqmnxaktdrhbz\r\nwlpiogsvmfecjuqmnxmktdrhtz\r\nwlhiogsvyfecjuhmnxaktsrhbz\r\nwlpioggvpfecjufmnxaktdrhbz\r\nwlpiogsvyfbcjuomnxaktdrhbh\r\nwlpmogsvyfecyuqmnxoktdrhbz\r\nwlpiogslyfecjuqmnxaptyrhbz\r\ntlpiogsnyfecguqmnxaktdrhbz\r\nwlpiogsvyfecjuqmnxwktwehbz\r\nwlpiogsvgfecjuqmnxaktdrabu\r\nwbpiogsvyfecjuqmnxagtdrhoz\r\nwlwipgsvyfecjuqmnxaktdrhbu\r\nwlpwogsvykeczuqmnxaktdrhbz\r\nwlpgogsvwfecjuqmnxcktdrhbz\r\nwlpiogsqyfecjuqmrxaktdrrbz\r\nwlpiogsvyfecjuqmnxakthrafz\r\nwypicgseyfecjuqmnxaktdrhbz\r\nwlpiogcvqfecjuqmnxaktdrhzz\r\nwlriogsvyfecouqmnkaktdrhbz\r\nwlpiogsvyfemjulmnxaktdrhdz\r\nflpiogadyfecjuqmnxaktdrhbz\r\nwupiogsvyfbvjuqmnxaktdrhbz\r\nwlpiogsvyfebjummnxaktdrrbz\r\nwjpiogsvyfecjuqmnxaktprybz\r\nwlpirgsvyfecjiqmnxaatdrhbz\r\nbvpiogsvyfecjuqmnxaktdrhez\r\nwlpiogsvyfxcjuqmnxykzdrhbz\r\nwlkiwgsqyfecjqqmnxaktdrhbz\r\nwepaogsvyfecjxqmnxaktdrhbz\r\nwlpiovsvyfecjjqmnxaktdmhbz\r\nwlpioysryfecjuqmnxaktdrhiz\r\nwlpizjsvyfecjuvmnxaktdrhbz\r\ndlpiogsvyfecjucmnxakbdrhbz\r\nwlpiogsccfecjrqmnxaktdrhbz\r\nwlpioggvyfecpuqmnxagtdrhbz\r\nwlpiogsvyfvcjuumlxaktdrhbz\r\nwwpiogsryfjcjuqmnxaktdrhbz\r\nwlpiogsvyfecjuqynxaktdrogz\r\nwlpiogsvyfecjpqmnxskbdrhbz\r\nwlpiogsvyfecjuhmfxaktvrhbz\r\nwlpiogevyfecjrqmnwaktdrhbz\r\nwlpiigsvyfemjuqmnxaktdrhtz\r\nwlpcogsvyfecjuqhnxakgdrhbz\r\nwupiogsvyfxcjuqmnxaktdrhgz\r\nwlsiogsvyfecjuqenxuktdrhbz\r\nwlpioglvyfecjujmexaktdrhbz\r\nwlriogsvyfeljuqmnxattdrhbz\r\nwlpiogsvyfecfuqmhxaktkrhbz\r\nwlppogsvyfecjuqmxxabtdrhbz\r\nwlniogsvyfevjuqwnxaktdrhbz\r\nwlhiogsvyfecjuqmnxactxrhbz\r\nilpiogivyflcjuqmnxaktdrhbz\r\nwlpmogsvyfecjuqmnxaktdrlbs\r\nwipiogsvyfeqjuqmnxaktrrhbz\r\nwvpiogsvyfecjuqknxaktdrrbz\r\nwwpioguvyfecxuqmnxaktdrhbz\r\nwlpiogsvkfdcjuqmnxaktdzhbz\r\nwlpiogfvyfecjuqmnxadtdrhbg\r\nwlpiogsvyzefjuqfnxaktdrhbz\r\nwlpiogstyfechuqmnxaktdchbz\r\nwlpiogszyfedjuqmnxsktdrhbz\r\nwzpiozsvyfncjuqmnxaktdrhbz\r\nxlpiogsvyfefjuqmnmaktdrhbz\r\nwlpiogsvyfebxummnxaktdrhbz\r\nwlpiogsgyfecfurmnxaktdrhbz\r\nwlpqogsvyfecjuomnxaktdrhbi\r\nwlpiogjvufecjuqmnxaktdrhbd\r\nwlpiolsvyfecduqmnxaktrrhbz\r\nwlpxogsvyfecjuqmnxaktgrhbk\r\nwlpiogsfyfncjuqmnxsktdrhbz\r\nwlpioggvyfecjufmnxaktdrebz\r\nwlpiogsvyfecfujmnxaktdrwbz\r\nrlpiogsvyfecjlqmnxaktdqhbz\r\nwlpfogsvyfecjuimnxaktfrhbz"
	return strings.Split(s, "\r\n")
}
