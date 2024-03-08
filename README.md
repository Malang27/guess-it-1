 This project aims to output in *output* estimates of the next value that will be given in *input*.

For this we used the law of normal distribution which is often used in statistics.The value 1.28 is used because it corresponds to the number of standard deviations away from the mean that encompasses 90% of the area under a normal distribution curve. In other words, if the values in the sequence follow a normal distribution, we can reasonably expect that about 90% of the next values will fall within the range defined by the mean plus or minus 1.28 times the standard deviation.

This value comes from a statistical property of normal distributions known as the empirical rule or the 68-95-99.7 rule. The empirical rule states that for any normal distribution, approximately:

    68% of the data falls within one standard deviation of the mean
    95% of the data falls within two standard deviations of the mean
    99.7% of the data falls within three standard deviations of the mean Using this rule, we can derive the value 1.28 as the number of standard deviations that encompasses the middle 90% of the data, or equivalently, the number of standard deviations that leaves out the outer 5% on each side of the distribution.

