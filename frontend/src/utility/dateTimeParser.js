export function DateTimeParser(dateTime) {
  const currentDateTime = new Date(Date.now());
  if (currentDateTime - dateTime < 1000 * 60 * 60 * 24) {
    const hours = dateTime.getHours();
    const minutes = dateTime.getMinutes();
    return (hours < 10 ? "0" + hours : hours.toString()) + ":" + (minutes < 10 ? "0" + minutes : minutes);
  } else {
    const day = dateTime.getDate();
    const month = dateTime.getMonth() + 1;
    let res = (day < 10 ? "0" + day : day.toString()) + "." + (month < 10 ? "0" + month : month);
    if (currentDateTime - dateTime > 1000 * 60 * 60 * 24 * 365) {
      res += "." + dateTime.getFullYear();
    }
    return res;
  }
}
