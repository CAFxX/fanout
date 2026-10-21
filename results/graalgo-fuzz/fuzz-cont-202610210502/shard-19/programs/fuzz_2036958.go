package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = (1.5 + 1.5)
    var v3 int16 = 42
    var v4 float64 = (0.0 / 3.14159)
    var v5 int = -234
    fmt.Println(v5, v2, v2)
    v5 += v5
    for fi6 := 0; fi6 < 16; fi6++ {
        v1[2] = "foo bar"
        v3 = 1
        fmt.Println(v3, v3, v3)
        _ = fi6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint32 = 263
        var v9 uint64 = 97
        fmt.Println(fi7)
        _ = fi7
        _ = v8
        _ = v9
    }
    v5 = v5
    if (false && false) {
        for fi10 := 0; fi10 < 14; fi10++ {
            if false {
                for fi11 := 0; fi11 < 1; fi11++ {
                    if (42 < 4294967295) {
                        if true {
                            v12 := GenS1{F0: ("" + ("foo bar" + "")), F1: 137, F2: 1}
                            _ = v12
                        } else {
                            v2 += 3.14159
                        }
                        v13 := v1[2147483647]
                        _ = v13
                    }
                    var v14 uint64 = 0
                    fmt.Println(fi11, v3, v14)
                    if true {
                        v3 = v3
                    } else {
                        var v15 int32 = -235
                        if false {
                            v15 += -1
                        } else {
                            if (false && false) {
                                switch 32768 {
                                case 4294967295:
                                    fmt.Println(v5, v14)
                                case 844:
                                    v16 := GenS1{F0: ("foo bar" + "\n\t"), F1: 42, F2: 1}
                                    _ = v16
                                default:
                                    var v17 uint16 = 317
                                    _ = v17
                                }
                            } else {
                                v18 := GenS1{F0: "foo bar", F1: 255, F2: 1}
                                var v19 bool = false
                                _ = v18
                                _ = v19
                            }
                            v20 := v1[564]
                            _ = v20
                        }
                        v21 := GenS1{F0: "foo bar", F1: 0, F2: 1}
                        _ = v15
                        _ = v21
                    }
                    _ = fi11
                    _ = v14
                }
            } else {
                for fi22 := 0; fi22 < 19; fi22++ {
                    v5 += v5
                    _ = fi22
                }
            }
            _ = fi10
        }
        var v23 uint = 2
        _ = v23
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        fmt.Println(v5, v3)
        switch -405 {
        case 0:
            v2 = (0.1 - v4)
        case 4294967295:
            for fi25 := 0; fi25 < 2; fi25++ {
                if true {
                    v26 := []int{fi24, (fi24 >> ((2) & 63)), -278}
                    v27 := v1[2]
                    _ = v26
                    _ = v27
                }
                if (((-2.25 + (float64(1e10) + v2)) * 0.1) >= float64(0.0)) {
                    var v28 bool = (!(-259 > 539))
                    v28 = v28
                    _ = v28
                }
                _ = fi25
            }
        default:
            v4 = (float64(0.1) + (v4 / v4))
        }
        _ = fi24
    }
    v29 := []int{2147483647, (v5 << ((255) & 63))}
    v30 := GenS1{F0: (("foo bar" + "") + ("foo bar" + "hello")), F1: 181, F2: 1}
    for fi31 := 0; fi31 < 8; fi31++ {
        var v32 string = "\n\t"
        _ = fi31
        _ = v32
    }
    if (((v4 * (v2 / (v2 - 1e10))) / (v2 + 0.0)) != (v4 / v4)) {
        for fi33 := 0; fi33 < 5; fi33++ {
            v4 -= 3.14159
            var v34 uint8 = 128
            _ = fi33
            _ = v34
        }
    }
    fmt.Println(v5)
    var v35 uint32 = 2147483648
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v5
}
