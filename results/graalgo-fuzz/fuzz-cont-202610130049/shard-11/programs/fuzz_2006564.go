package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 uint8; F1 bool }

func main() {
    var v1 uint = 18446744073709551615
    var v2 int16 = 1000
    var v3 uint = v1
    v4 := []int{-478, 42, 128}
    var v5 int16 = v2
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: v3, F1: -233, F2: (v1 & (v1 / ((v3) | 1)))}
        var v8 float64 = 1.5
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: 32767, F1: 2}
    switch 32767 {
    case 4294967295:
        v3 += v3
    case -27:
        v3 -= v3
    }
    if ((3.14159 / 0.1) > 3.14159) {
        if (!false) {
            switch 0 {
            case 32767:
                v10 := GenS1{F0: v1, F1: 14, F2: 977}
                _ = v10
            case -720:
                var v11 float32 = 1.5
                _ = v11
            case 32768:
                if false {
                    var v12 int16 = (v2 * (-(v2)))
                    if true {
                        v13 := GenS1{F0: (v1 / ((2) | 1)), F1: 0, F2: (v3 / (((v3 / ((v3) | 1))) | 1))}
                        _ = v13
                    } else {
                        v1 = v1
                        if (130 <= 255) {
                            var v14 uint64 = 2
                            _ = v14
                        } else {
                            for fi15 := 0; fi15 < 19; fi15++ {
                                v1 -= (v3 % ((v1) | 1))
                                v2 -= (v2 ^ (v2 * v2))
                                switch 128 {
                                case 626:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                default:
                                    v12 = v5
                                }
                                _ = fi15
                            }
                            v1 = (^(v3))
                            v16 := GenS2{F0: 2, F1: 2}
                            _ = v16
                        }
                    }
                    _ = v12
                } else {
                    v5 -= (v2 % ((1) | 1))
                    v17 := GenS1{F0: (v1 << ((32) & 63)), F1: 0, F2: 2}
                    _ = v17
                }
            }
            var v18 uint16 = 32767
            _ = v18
        } else {
            v19 := GenS2{F0: 566, F1: 1}
            var v20 int8 = 1
            _ = v19
            _ = v20
        }
        v3 = v3
    }
    if (!(2 > 1)) {
        v3 = (^(v1))
    } else {
        if ("hello" > ("hello" + "foo bar")) {
            v2 -= 62
            v4 = append(v4, 379)
        }
        if len(v4) > 0 {
            v4[0] = 9223372036854775807
        }
        v3 -= 255
    }
    v4 = append(v4, -380)
    for fi21 := 0; fi21 < 14; fi21++ {
        for fi22 := 0; fi22 < 9; fi22++ {
            if (!false) {
                v23 := GenS2{F0: 179, F1: 32767}
                if (903 == 83) {
                    v1 += (v1 * (v3 / ((v1) | 1)))
                }
                _ = v23
            } else {
                var v24 uint16 = 581
                v2 = v2
                _ = v24
            }
            var v25 int16 = v2
            _ = fi22
            _ = v25
        }
        v26 := GenS3{F0: 0, F1: (v1 / ((42) | 1))}
        if (!false) {
            if false {
                var v27 uint = v3
                _ = v27
            }
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        } else {
            if ((!false) || (false || false)) {
                switch ((("foo bar" + "hello") + "") + ("foo bar" + (("a" + (("a" + "\n\t") + "")) + "\n\t"))) {
                case "foo bar":
                    if false {
                        for fi28 := 0; fi28 < 19; fi28++ {
                            v1 -= v3
                            _ = fi28
                        }
                        v3 += 65535
                        v2 += v5
                    } else {
                        v1 -= 849
                    }
                case "\n\t":
                    fmt.Println(v1, v2)
                case "a":
                    v29 := GenS4{F0: 198, F1: true}
                    _ = v29
                }
            }
        }
        _ = fi21
        _ = v26
    }
    v1 = (^(v1))
    var v30 int16 = 756
    var v31 int64 = -454
    v32 := GenS3{F0: -755, F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v9
}
