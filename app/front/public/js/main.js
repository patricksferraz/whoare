function countWords(elem) {
  let words = document.getElementById('words');
  words.innerText = elem.getAttribute('maxlength') - elem.value.length;
}
