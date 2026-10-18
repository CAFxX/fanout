package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int16 }

func main() {
    v1 := []string{("foo bar" + "\n\t")}
    v2 := make(map[string]bool)
    var v3 int16 = 0
    var v4 int8 = -97
    v5 := make(map[int]int)
    var v6 int64 = 9223372036854775807
    if (1 < 846) {
        var v7 int = 128
        for fi8 := 0; fi8 < 1; fi8++ {
            v6 = 602
            v6 = 0
            var v9 float64 = 1e10
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        v6 = v6
        v10 := GenS1{F0: 684, F1: false}
        _ = v10
    }
    switch 0 {
    case 207:
        v4 = (v4 - (-(v4)))
    case 32767:
        var v11 uint = 65535
        _ = v11
    case -1:
        switch ("hello" + "a") {
        case "hello":
            v3 += -118
        case "\n\t":
            var v12 int64 = (v6 - 869)
            _ = v12
        default:
            var v13 uint64 = 42
            _ = v13
        }
    default:
        var v14 float32 = 0.0
        _ = v14
    }
    v3 += v3
    var v15 uint32 = 2
    if len(v1) > 0 {
        v1[0] = "a"
    }
    var v16 float32 = (1e10 - 3.14159)
    fmt.Println(v16, v16, v3)
    fmt.Println(v16)
    var v17 uint64 = 2
    var v18 int16 = (v3 + -32768)
    if (false || ((1 > 375) && false)) {
        v6 -= v6
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
        v6 -= v6
    } else {
        if ((!true) && (!((v4 % ((v4) | 1)) > (v4 / (((v4 + (v4 | (-(v4))))) | 1))))) {
            v4 -= 100
            v6 = v6
        } else {
            for fi19 := 0; fi19 < 14; fi19++ {
                if (!(!true)) {
                    var v20 int32 = 0
                    if false {
                        v20 += (v20 >> ((703) & 63))
                        var v21 uint64 = (v17 - (v17 << ((294) & 63)))
                        _ = v21
                    }
                    var v22 uint = 812
                    _ = v20
                    _ = v22
                }
                _ = len(v1)
                if (-32768 == (v3 ^ (v3 * (v3 | int16(fi19))))) {
                    if (-32768 != v18) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        for fi23 := 0; fi23 < 15; fi23++ {
                            var v24 uint = 0
                            v25 := GenS1{F0: (v24 % ((uint(v17)) | 1)), F1: (false && ((v24 & (v24 | 65535)) >= v24))}
                            v17 += 1
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                    } else {
                        v26 := GenS1{F0: 4294967295, F1: true}
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 3; fi27++ {
                        v3 = (v18 / (((v18 & (v3 ^ (v18 ^ -32768)))) | 1))
                        v15 = v15
                        _ = fi27
                    }
                    switch (-(fi19)) {
                    case 273:
                        switch ("hello" + ((("foo bar" + "hello") + (("foo bar" + ("foo bar" + "\n\t")) + "")) + "hello")) {
                        case "foo bar":
                            v28 := GenS1{F0: 322, F1: (32767 >= 65535)}
                            _ = v28
                        case "a":
                            switch (fi19 - (^(fi19))) {
                            case 32768:
                                v29 := GenS1{F0: 255, F1: (v4 <= v4)}
                                _ = v29
                            default:
                                var v30 uint32 = v15
                                _ = v30
                            }
                        case "\n\t":
                            v1 = append(v1, "hello")
                        default:
                            if false {
                                if (true && (4294967295 > 255)) {
                                    fmt.Println(fi19, fi19, v3)
                                }
                            }
                        }
                    case 266:
                        for fi31 := 0; fi31 < 6; fi31++ {
                            v32 := v2["\n\t"]
                            var v33 int16 = 32767
                            var v34 bool = (!(!false))
                            _ = fi31
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                    case -740:
                        v35 := GenS2{F0: v18}
                        _ = v35
                    default:
                        v36 := GenS2{F0: v18}
                        _ = v36
                    }
                } else {
                    var v37 int32 = -2147483648
                    _ = v37
                }
                _ = fi19
            }
            var v38 float32 = v16
            _ = v38
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
