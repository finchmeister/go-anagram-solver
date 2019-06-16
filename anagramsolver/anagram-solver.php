<?php

class AnagramSolver
{
    /**
     * @var int
     */
    private $minWordLength;
    /** @var array */
    private $dictionary;
    private $lettersCount;
    private $anagrams;

    public function __construct(
        int $minWordLength
    ) {
        $this->minWordLength = $minWordLength;
        $this->loadDictionary();
    }

    private function loadDictionary()
    {
        $this->dictionary = explode(PHP_EOL, file_get_contents(__DIR__.'/dictionary.txt'));
    }

    public function nextPermutation(&$input)
    {
        $inputCount = count($input);
        // the head of the suffix
        $i = $inputCount - 1;
        // find longest suffix
        while ($i > 0 && $input[$i] <= $input[$i - 1]) {
            $i--;
        }
        //are we at the last permutation already?
        if ($i <= 0) {
            return false;
        }
        // get the pivot
        $pivotIndex = $i - 1;
        // find rightmost element that exceeds the pivot
        $j = $inputCount - 1;
        while ($input[$j] <= $input[$pivotIndex]) {
            $j--;
        }

        // swap the pivot with j
        $temp = $input[$pivotIndex];
        $input[$pivotIndex] = $input[$j];
        $input[$j] = $temp;
        // reverse the suffix
        $j = $inputCount - 1;
        while ($i < $j) {
            $temp = $input[$i];
            $input[$i] = $input[$j];
            $input[$j] = $temp;
            $i++;
            $j--;
        }
        return true;
    }

    public function anagramSolver(string $letters)
    {
        $start = microtime(true);

        $letters = str_split(strtolower($letters));
        sort($letters);
        $this->lettersCount = count($letters);
        $this->anagrams = [];

        foreach ($this->dictionary as $word) {
            $wordLength = strlen($word);
            do {
                $retryDictionaryWord = false;
                for ($i = 0; $i < $wordLength; $i++) {
                    if ($wordLength < $this->minWordLength || $wordLength > $this->lettersCount) {
                        // next word (word is too short or long)
                        continue 2;
                    }

                    if ($word[$i] < $letters[$i]) {
                        // next word (word is ordered before)
                        continue 2;
                    }

                    if ($word[$i] === $letters[$i]) {
                        if ($i === $wordLength - 1) {
                            $this->anagrams[] = $word;
                            // next word (we've found an anagram)
                            continue 2;
                        }
                        continue;
                    }
                    // keep permutating word until it is greater than dictionary
                    $lettersBefore = $letters;
                    $this->nextPermutation($letters);
                    if ($letters === $lettersBefore) {
                        $timeTaken = microtime(true) - $start;
                        echo $timeTaken*1000 .PHP_EOL;

                        return $this->anagrams;
                    }
                    $retryDictionaryWord = true;
                    break;
                }
            } while ($retryDictionaryWord);
        }

        $timeTaken = microtime(true) - $start;
        echo $timeTaken*1000 .PHP_EOL;


        return $this->anagrams;
    }

    public function anagramSolver2(string $letters)
    {
        $start = microtime(true);

        $letters = str_split(strtolower($letters));
        sort($letters);
        $this->lettersCount = count($letters);
        $this->anagrams = [];

        foreach ($this->dictionary as $word) {
            if ($this->skipWord($word)) {
                continue;
            }
            if ($this->isWordInLetters($word, $letters)) {
                $this->anagrams[] = $word;
            }
        }

        $timeTaken = microtime(true) - $start;

        echo $timeTaken*1000 .PHP_EOL;

        return $this->anagrams;
    }

    private function skipWord(string $word): bool
    {
        $wordLength = strlen($word);
        return $wordLength < $this->minWordLength || $wordLength > $this->lettersCount;
    }

    private function isWordInLetters(string $word, array $letters)
    {
        foreach (str_split($word) as $letter) {
            $key = array_search($letter, $letters);
            if ($key === false) {
                return false;
            }
            unset($letters[$key]);
        }

        return true;
    }
}




$expected = array (
    0 => 'aardvark',
    1 => 'ada',
    2 => 'adar',
    3 => 'aka',
    4 => 'aka',
    5 => 'akra',
    6 => 'akra',
    7 => 'ara',
    8 => 'arad',
    9 => 'arar',
    10 => 'arara',
    11 => 'arara',
    12 => 'ark',
    13 => 'ava',
    14 => 'avar',
    15 => 'dak',
    16 => 'dar',
    17 => 'dark',
    18 => 'darr',
    19 => 'kava',
    20 => 'kra',
    21 => 'kra',
    22 => 'raad',
    23 => 'rad',
    24 => 'rada',
    25 => 'radar',
    26 => 'vara',
    27 => 'varda',
);

$a = new AnagramSolver(3);
//$start = microtime(true);
//$result = $a->anagramSolver2('aardvarksab');
//$timeTaken = microtime(true) - $start;

$a->anagramSolver("aar");
$a->anagramSolver("aard");
$a->anagramSolver("aardv");
$a->anagramSolver("aardva");
$a->anagramSolver("aardvar");
$a->anagramSolver("aardvark");
$a->anagramSolver("aardvarks");
$a->anagramSolver("aardvarksa");
$a->anagramSolver("aardvarksab");

//if ($result === $expected) {
//    echo 'SUCCESS';
//} else {
//    echo 'FAIL';
//    print_r($result);
//}
//print_r($result);
//echo $timeTaken.PHP_EOL;