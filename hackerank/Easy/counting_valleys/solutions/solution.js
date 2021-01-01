/**
 *
 * @param {number} n  Number of steps taken for the hikes
 * @param {string} paths sequence of characters denoting the upward and downward movement above and below sea-level
 */
const countingValleys = (n, paths) => {
  let seaLevel = 0,
    valleyCount = 0,
    depthTracker = 0;

  paths.split("").forEach((path) => {
    if (path === "U") {
      seaLevel++;
    } else {
      seaLevel--;
    }

    //for a deep below sea-level
    if (seaLevel === -1) {
      depthTracker = -1;
    }

    if (depthTracker === -1 && seaLevel === 0) {
      valleyCount++;
      depthTracker = 0;
    }
  });

  return valleyCount;
};

module.exports = countingValleys;
