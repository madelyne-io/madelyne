package matcher

import (
	"errors"
	"testing"
)

func TestMatchStringValues(t *testing.T) {
	tests := []struct {
		value          interface{}
		patternOrValue interface{}
		expectedError  error
	}{
		{
			value:          "Bonjour !",
			patternOrValue: "@string@",
			expectedError:  nil,
		},
		{
			value:          "Bonjour !",
			patternOrValue: "@number@",
			expectedError:  ErrNotNumber,
		},
		{
			value:          "Bonjour !",
			patternOrValue: "Bonjour !",
			expectedError:  nil,
		},
		{
			value:          "Bonjour !",
			patternOrValue: "Bonsoir !",
			expectedError:  ErrInvalidValue,
		},
		{
			value:          "Bonjour !",
			patternOrValue: 17.0,
			expectedError:  ErrInvalidValue,
		},
		{
			value:          "Bonjour !",
			patternOrValue: "@UNKNOWN_TYPE@",
			expectedError:  ErrInvalidPattern,
		},
		{
			value:          "B(onjour, comment ça va ?",
			patternOrValue: "@string@.startsWith('B(onjour')",
			expectedError:  nil,
		},
		{
			value:          "@B(onjour, comment ça va ?",
			patternOrValue: "@string@.startsWith('@')",
			expectedError:  nil,
		},
		{
			value:          "Bonjour, comment ça va ?",
			patternOrValue: "@string@.startsWith('bonjour')",
			expectedError:  ErrNotStartsWith,
		},
		{
			value:          "Bonjour, comment ça va ?",
			patternOrValue: "@string@.startsWith('blablabla')",
			expectedError:  ErrNotStartsWith,
		},
		{
			value:          "Bonjour, comment ça va ?",
			patternOrValue: "@string@.endsWith('ça va ?')",
			expectedError:  nil,
		},
		{
			value:          "Bonjour, comment ça va ?",
			patternOrValue: "@string@.endsWith('ça VA ?')",
			expectedError:  ErrNotEndsWith,
		},
		{
			value:          "Bonjour, comment ça va ?",
			patternOrValue: "@string@.endsWith('blablabla')",
			expectedError:  ErrNotEndsWith,
		},
		{
			value:          "Bonjour, comment ça va ?",
			patternOrValue: "@string@.startsWith('Bonjour').endsWith('ça va ?')",
			expectedError:  nil,
		},
		{
			value:          "Bonjour, comment ça va ?",
			patternOrValue: "@string@.contains('comment')",
			expectedError:  nil,
		},
		{
			value:          "Bonjour, comment ça va ?",
			patternOrValue: "@string@.contains('COMMENT')",
			expectedError:  ErrNotContains,
		},
		{
			value:          "Bonjour, comment ça va ?",
			patternOrValue: "@string@.contains('bonsoir')",
			expectedError:  ErrNotContains,
		},
		{
			value:          "Bonjour, comment ça va ?",
			patternOrValue: "@string@.notContains('Bonsoir')",
			expectedError:  nil,
		},
		{
			value:          "Bonjour, comment ça va ?",
			patternOrValue: "@string@.notContains('comment')",
			expectedError:  ErrContains,
		},
		{
			value:          "Bonjour, comment ça va ?",
			patternOrValue: "@string@.notContains('COMMENT')",
			expectedError:  nil,
		},
		{
			value:          "Bonjour, comment ça va ?",
			patternOrValue: "@string@.contains('Bonjour').notContains('Bonsoir')",
			expectedError:  nil,
		},
		{
			value:          "https://www.everycheck.com/",
			patternOrValue: "@string@.isUrl()",
			expectedError:  nil,
		},
		{
			value:          "everycheck",
			patternOrValue: "@string@.isUrl()",
			expectedError:  ErrNotUrl,
		},
		{
			value:          "https://www.everycheck.com/",
			patternOrValue: "@string@.isUrl().contains('everycheck')",
			expectedError:  nil,
		},
		{
			value:          "2020-07-24T08:11:55.537Z",
			patternOrValue: "@string@.isDateTime()",
			expectedError:  nil,
		},
		{
			value:          "2020-07-24",
			patternOrValue: "@string@.isDateTime()",
			expectedError:  ErrNotDateTime,
		},
		{
			value:          "2020-07-24T08:11:55.537Z",
			patternOrValue: "@string@.isDateTime().contains('2020')",
			expectedError:  nil,
		},
		{
			value:          "raphael.alves@everycheck.fr",
			patternOrValue: "@string@.isEmail()",
			expectedError:  nil,
		},
		{
			value:          "raphael.alves@everycheck",
			patternOrValue: "@string@.isEmail()",
			expectedError:  ErrNotEmail,
		},
		{
			value:          "",
			patternOrValue: "@string@.isEmpty()",
			expectedError:  nil,
		},
		{
			value:          "str",
			patternOrValue: "@string@.isEmpty()",
			expectedError:  ErrNotEmpty,
		},
		{
			value:          "",
			patternOrValue: "@string@.isNotEmpty()",
			expectedError:  ErrEmpty,
		},
		{
			value:          "str",
			patternOrValue: "@string@.isNotEmpty()",
			expectedError:  nil,
		},
		{
			value:          "123456",
			patternOrValue: "@string@.matchRegex('\\d')",
			expectedError:  nil,
		},
		{
			value:          "AGR_AGD46DGZT6D7",
			patternOrValue: "@string@.matchRegex('^AGR_[A-Z0-9]{12}$')",
			expectedError:  nil,
		},
		{
			value:          "teste",
			patternOrValue: "@string@.oneOf(contains('test'), notContains('abc'))",
			expectedError:  nil,
		},
		{
			value:          "teste",
			patternOrValue: "@string@.oneOf(contains('abc'), notContains('teste'), startsWith('abc'))",
			expectedError:  ErrOneOf,
		},
		{
			value:          "#FFFFFF",
			patternOrValue: "#FFFFFF",
			expectedError:  nil,
		},
		{
			value:          1.6,
			patternOrValue: "@number@",
			expectedError:  nil,
		},
		{
			value:          "2.6",
			patternOrValue: "@number@",
			expectedError:  ErrNotNumber,
		},
		{
			value:          3.6,
			patternOrValue: 3.6,
			expectedError:  nil,
		},
		{
			value:          4.7,
			patternOrValue: 4.6,
			expectedError:  ErrInvalidValue,
		},
		{
			value:          5.7,
			patternOrValue: "5.6",
			expectedError:  ErrInvalidValue,
		},
		{
			value:          6.6,
			patternOrValue: "@UNKNOWN_TYPE@",
			expectedError:  ErrInvalidPattern,
		},
		{
			value:          7.6,
			patternOrValue: "@number@.greaterThan(5)",
			expectedError:  nil,
		},
		{
			value:          8.6,
			patternOrValue: "@number@.greaterThan(9)",
			expectedError:  ErrGreaterThan,
		},
		{
			value:          9.6,
			patternOrValue: "@number@.lowerThan(10)",
			expectedError:  nil,
		},
		{
			value:          10.6,
			patternOrValue: "@number@.lowerThan(5)",
			expectedError:  ErrLowerThan,
		},
		{
			value:          11.6,
			patternOrValue: "@number@.lowerThan(15).greaterThan(5)",
			expectedError:  nil,
		},
		{
			value:          12.6,
			patternOrValue: "@number@.oneOf(lowerThan(7), greaterThan(9))",
			expectedError:  nil,
		},
		{
			value:          12.6,
			patternOrValue: "@number@.oneOf(lowerThan(7), greaterThan(19))",
			expectedError:  ErrOneOf,
		},
		{
			value:          true,
			patternOrValue: "@boolean@",
			expectedError:  nil,
		},
		{
			value:          true,
			patternOrValue: "@number@",
			expectedError:  ErrNotNumber,
		},
		{
			value:          true,
			patternOrValue: true,
			expectedError:  nil,
		},
		{
			value:          true,
			patternOrValue: false,
			expectedError:  ErrInvalidValue,
		},
		{
			value:          "true",
			patternOrValue: true,
			expectedError:  ErrInvalidValue,
		},
		{
			value:          true,
			patternOrValue: "@UNKNOWN_TYPE@",
			expectedError:  ErrInvalidPattern,
		},
		{
			value:          "2cbd1211-aa7b-4251-97f5-a895af9f6002",
			patternOrValue: "@uuid@",
			expectedError:  nil,
		},
		{
			value:          "b16865b2-7bd5-4214-9566-5160ebd39640",
			patternOrValue: "@uuid@",
			expectedError:  nil,
		},
		{
			value:          true,
			patternOrValue: "@uuid@",
			expectedError:  ErrNotUuid,
		},
		{
			value:          "This is not a uuid",
			patternOrValue: "@uuid@",
			expectedError:  ErrNotUuid,
		},
		{
			value:          "c45bf1a5-ea78-4091-b4d9-f9ed94c760f4",
			patternOrValue: "c45bf1a5-ea78-4091-b4d9-f9ed94c760f4",
			expectedError:  nil,
		},
	}
	for i, test := range tests {

		err := Match(test.value, test.patternOrValue)
		if !errors.Is(err, test.expectedError) {
			t.Fatalf(
				"[%d] Failed. Value: %s ; Pattern/Value : %s\n want %v got: %v\n",
				i, test.value, test.patternOrValue, test.expectedError, err,
			)
		}
	}
}
