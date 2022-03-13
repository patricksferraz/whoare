function CountWords(s) {
  var words = document.getElementById('words');
  words.innerText = s.getAttribute('maxlength') - s.value.length;
}
