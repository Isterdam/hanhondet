function getWordInfo(urlParams) {
  var txt = "\"" + urlParams["ord"][0] + "\" är";
  // can not find word if "Error" is in the url
  if (urlParams["ord"][0] === "Error") {
    txt = "Kunde inte hitta ordet :(";
  } else {
    var nBool, mBool, fBool, oBool = false;
    var txts = [];
    for (i = 0; i < urlParams["ord"].length; i++) {
      switch (urlParams["genus"][i]) {
        // Find genders of word
        case "n":
          if (!nBool) txts.push(" ett \"det\""); nBool = true;
          break;
        case "m":
          if (!mBool) txts.push(" en han"); mBool = true;
          break;
        case "f":
          if (!fBool) txts.push(" en hon"); fBool = true;
          break;
        default:
          if (!oBool) txts.push(" en okänd"); oBool = true;
      }
    }
    txt += txts[0];
    for (i = 1; i < txts.length; i++) {
      if (txts.length - i === 1) {
        txt += " eller" + txts[i];
      } else {
        txt += "," + txts[i];
      }
    }
  }

  return txt;
}
