import os

from typing import List


_stop_words = None


def stopwords() -> set:
    """ Returns a list of stopwords. """
    global _stop_words
    if _stop_words is not None:
        return _stop_words
    stop_file = os.path.join(os.path.dirname(__file__), 'stopwords.txt')
    _stop_words = set()
    with open(stop_file, 'r', encoding='utf-8') as f:
        for line in f:
            _stop_words.add(line.strip())
    return _stop_words


def words(text: str) -> List[str]:
    """Extract the words from the given text."""
    words = []
    buffer = ""
    for char in text.lower():
        if char.isalnum() or char == "-":
            buffer = buffer + char
            continue
        if buffer == "":
            continue
        words.append(buffer)
        buffer = ""
    if buffer != "":
        words.append(buffer)
    return words


def remove_duplicates(words: List[str]) -> List[str]:
    """Remove duplicate words but keep the order of the words."""
    r = []
    for word in words:
        if word in r:
            continue
        r.append(word)
    return r


def remove_stopwords(words: List[str]) -> List[str]:
    r = []
    stops = stopwords()
    for word in words:
        if word in stops:
            continue
        r.append(word)
    return r


def jaccard_index(words1: List[str], words2: List[str]) -> float:
    """Calculates the Jaccard index for the given lists of words;
       see https://en.wikipedia.org/wiki/Jaccard_index"""
    intersection = set(words1).intersection(set(words2))
    union = set(words1).union(set(words2))
    return len(intersection) / len(union)


if __name__ == '__main__':
    s1 = 'AI is our friend and it has been friendly.'
    s2 = 'AI and humans have always been friendly'
    print(jaccard_index(words(s1), words(s2)))
