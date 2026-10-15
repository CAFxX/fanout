package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint64, p1 int) uint64 {
    var v6 uint = 288
    p1 = int(p0)
    v7 := []int{(p1 ^ -2147483648), (p1 % ((-1) | 1)), p1, (p1 >> (((v6 & v6)) & 63))}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 476
}
func genFn1(p0 bool, p1 int16) int32 {
    v8 := []int{-89, 668, 32767}
    if p0 {
        p1 = (p1 | p1)
        if (-9223372036854775808 > 42) {
            v8 = append(v8, 32767)
            if false {
                var v9 int64 = 1
                fmt.Println(p1)
                p0 = (!p0)
                _ = v9
            }
            for fi10 := 0; fi10 < 13; fi10++ {
                v11 := GenS1{F0: 32, F1: ("a" + ("hello" + "\n\t")), F2: 2}
                p1 += 42
                var v12 int64 = 9223372036854775807
                _ = fi10
                _ = v11
                _ = v12
            }
        } else {
            var v13 int16 = (-(p1))
            p1 += (p1 | 1)
            switch "\n\t" {
            case "hello":
                p0 = p0
            }
            _ = v13
        }
        v14 := GenS2{F0: 42}
        _ = v14
    } else {
        v15 := GenS1{F0: 1, F1: "a", F2: 2}
        var v16 uint = 4294967295
        _ = v15
        _ = v16
    }
    var v17 uint16 = 791
    for fi18 := 0; fi18 < 17; fi18++ {
        p1 = p1
        _ = fi18
    }
    if len(v8) > 0 {
        v8[0] = 128
    }
    _ = p0
    _ = p1
    _ = v17
    _ = v8
    return 1
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]string)
    var v3 uint = 65535
    var v4 uint64 = (v1 >> ((451) & 63))
    var v5 bool = true
    v5 = ((((0.0 - -2.25) > 1e10) && v5) && v5)
    v2[("a" + "hello")] = ("\n\t" + "a")
    v4 = (v1 % ((v4) | 1))
    var v19 uint16 = 2
    fmt.Println(v4, v19)
    v19 = v19
    if (((v5 || (!true)) || (127 != 922)) && true) {
        v20 := GenS1{F0: 0, F1: "\n\t", F2: (v19 ^ 255)}
        _ = v20
    } else {
        switch 32768 {
        case 2147483647:
            v19 -= 65535
        case -2147483648:
            var v21 uint32 = 622
            _ = v21
        default:
            var v22 int8 = -39
            _ = v22
        }
        for fi23 := 0; fi23 < 12; fi23++ {
            v1 = (v4 % ((v1) | 1))
            _ = fi23
        }
        switch -9223372036854775808 {
        case 135:
            switch 637 {
            case 1:
                var v24 int64 = -242
                _ = v24
            case 32767:
                for fi25 := 0; fi25 < 14; fi25++ {
                    v4 = v1
                    v4 -= (^(v4))
                    _ = fi25
                }
            case 2147483647:
                if (!(false || (-2.25 == (1e10 + ((1.5 - (-2.25 + float64(-2.25))) / 0.1))))) {
                    if v5 {
                        var v26 uint32 = 65535
                        _ = v26
                    } else {
                        var v27 uint16 = 2
                        v28 := []int{127, 2147483647, -778, 128}
                        for fi29 := 0; fi29 < 6; fi29++ {
                            var v30 uint8 = 1
                            v2["a"] = "foo bar"
                            _ = len(v28)
                            _ = fi29
                            _ = v30
                        }
                        _ = v27
                        _ = v28
                    }
                    var v31 int64 = -9223372036854775808
                    _ = v31
                }
            }
        case 0:
            var v32 bool = v5
            _ = v32
        case -1:
            fmt.Println(v4)
        default:
            v33 := []int{42, 128, -9223372036854775808, -480}
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        v35 := GenS1{F0: 128, F1: "\n\t", F2: v19}
        if (((-2.25 + (1.5 + (1e10 / 1.0))) * 1e10) > (1e10 * 0.0)) {
            v36 := []int{}
            _ = v36
        }
        var v37 int16 = 1000
        _ = fi34
        _ = v35
        _ = v37
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        if true {
            v19 = (^(v19))
        }
        v39 := GenS1{F0: 1, F1: ("" + "a"), F2: (v19 / ((v19) | 1))}
        _ = len(v2)
        _ = fi38
        _ = v39
    }
    for fi40 := 0; fi40 < 16; fi40++ {
        v4 = (v4 + v4)
        fmt.Println(v3, v5)
        var v41 float64 = 3.14159
        _ = fi40
        _ = v41
    }
    for fi42 := 0; fi42 < 19; fi42++ {
        if false {
            v43 := []int{(-(fi42)), fi42, fi42}
            v44 := GenS2{F0: -1}
            _ = v43
            _ = v44
        }
        _ = fi42
    }
    var v45 float64 = 1e10
    for fi46 := 0; fi46 < 17; fi46++ {
        v5 = true
        _ = fi46
    }
    v47 := []int{469, 205, 42}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v47
    _ = v5
}
