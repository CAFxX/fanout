package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 bool; F2 int16 }

func main() {
    var v1 float32 = float32(float64((0.1 + 1.5)))
    var v2 uint32 = 2
    v3 := make(map[string]bool)
    var v4 float64 = 0.0
    var v5 uint32 = (^(v2))
    for fi6 := 0; fi6 < 11; fi6++ {
        if (!(65535 < 18446744073709551615)) {
            if (!true) {
                var v7 float64 = (0.1 * v4)
                _ = v7
            }
            var v8 uint8 = 120
            _ = v8
        } else {
            v2 += v2
        }
        _ = fi6
    }
    v4 = v4
    v9 := []int{42}
    fmt.Println(v4, v2, v1)
    fmt.Println(v1)
    switch "\n\t" {
    case "\n\t":
        v2 += v2
    }
    var v10 uint16 = 1
    v11 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615, F2: (^(v10))}
    switch -168 {
    case -358:
        v12 := GenS2{F0: (0 != -314)}
        _ = v12
    case 9223372036854775807:
        if false {
            for fi13 := 0; fi13 < 20; fi13++ {
                fmt.Println(v1, v1, v1)
                _ = fi13
            }
            var v14 uint32 = 2147483647
            _ = v14
        } else {
            v5 = v5
            var v15 uint64 = 42
            _ = v15
        }
    case 32768:
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    default:
        _ = len(v3)
    }
    if true {
        if ((v4 / v4) >= (v4 - float64(0.1))) {
            for fi16 := 0; fi16 < 11; fi16++ {
                if len(v9) > 0 {
                    v9[0] = fi16
                }
                _ = fi16
            }
        }
        var v17 bool = true
        switch "foo bar" {
        case "a":
            switch 32767 {
            case 115:
                v10 = 255
            }
        case "":
            v5 += (v5 ^ 863)
        default:
            switch -831 {
            case -1:
                v5 = 4294967295
            case -2147483648:
                v18 := v3[""]
                _ = v18
            case 127:
                switch 32768 {
                case 2147483647:
                    switch "hello" {
                    case "hello":
                        for fi19 := 0; fi19 < 14; fi19++ {
                            if len(v9) > 0 {
                                v9[0] = 0
                            }
                            _ = fi19
                        }
                    case "":
                        v20 := v3[("" + "a")]
                        _ = v20
                    case "a":
                        for fi21 := 0; fi21 < 20; fi21++ {
                            v2 += 684
                            v22 := GenS1{F0: 110, F1: 42, F2: v10}
                            v4 -= 1.5
                            _ = fi21
                            _ = v22
                        }
                    default:
                        if (1.5 < v1) {
                            for fi23 := 0; fi23 < 20; fi23++ {
                                fmt.Println(v5)
                                for fi24 := 0; fi24 < 16; fi24++ {
                                    v4 += v4
                                    _ = fi24
                                }
                                _ = fi23
                            }
                            var v25 float64 = (v4 + (v4 + v4))
                            v26 := GenS3{F0: (128 >= 2147483647)}
                            _ = v25
                            _ = v26
                        }
                    }
                case -384:
                    v27 := GenS4{F0: "a", F1: (false && ((false || (!v17)) && (!v17))), F2: -411}
                    _ = v27
                }
            default:
                v10 += 1
            }
        }
        _ = v17
    } else {
        v3[""] = false
        for fi28 := 0; fi28 < 8; fi28++ {
            var v29 uint = 9
            _ = fi28
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
