package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int) int32 {
    v6 := []int{(^(p0)), (p0 ^ p0)}
    p0 += -134
    v7 := GenS1{F0: "", F1: 9223372036854775807, F2: 9223372036854775807}
    if ((1e10 + -2.25) >= (0.0 + ((0.0 / 3.14159) * 0.0))) {
        v8 := GenS1{F0: "a", F1: 0, F2: int64(p0)}
        var v9 int = 0
        if false {
            v6 = append(v6, (v9 | v9))
            if false {
                for fi10 := 0; fi10 < 20; fi10++ {
                    _ = len(v6)
                    _ = fi10
                }
            }
            var v11 int = (v9 * -2147483648)
            _ = v11
        }
        _ = v8
        _ = v9
    } else {
        if len(v6) > 0 {
            v6[0] = (p0 - p0)
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return -2147483648
}

func main() {
    var v1 uint32 = 2147483648
    v2 := []int{}
    v3 := []uint32{v1, v1, (v1 & (v1 ^ v1))}
    var v4 uint = 255
    v5 := make(map[int]int)
    if false {
        if true {
            var v12 int16 = 42
            if ((809 == uint16(v4)) && (false || (!false))) {
                var v13 uint64 = 558
                _ = v13
            } else {
                if true {
                    v1 += v1
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                v4 = (v4 + (v4 >> ((v4) & 63)))
            }
            v14 := v5[26]
            _ = v12
            _ = v14
        }
        v4 = uint(v1)
    }
    v3 = append(v3, 801)
    var v15 int = 32768
    for fi16 := 0; fi16 < 11; fi16++ {
        v1 = 0
        for fi17 := 0; fi17 < 2; fi17++ {
            if (false || false) {
                switch ("hello" + ("" + "\n\t")) {
                case "a":
                    if (v4 > v4) {
                        v4 = (v4 | (v4 | (v4 | 811)))
                    } else {
                        v4 += v4
                        v1 = (v1 * v1)
                        v18 := GenS1{F0: "a", F1: -733, F2: 0}
                        _ = v18
                    }
                default:
                    v4 = (v4 % (((v4 / ((v4) | 1))) | 1))
                }
            } else {
                if false {
                    v15 -= (fi17 ^ (fi16 << (((v4 + (^(v4)))) & 63)))
                }
            }
            _ = fi17
        }
        v19 := GenS1{F0: ((("hello" + "\n\t") + "foo bar") + ""), F1: 1, F2: -24}
        _ = fi16
        _ = v19
    }
    var v20 int32 = genFn0(v15)
    switch "" {
    case "a":
        var v21 bool = true
        _ = v21
    case "":
        v4 = v4
    }
    v22 := GenS1{F0: ("" + "\n\t"), F1: 9223372036854775807, F2: -1}
    v23 := GenS1{F0: "foo bar", F1: 1, F2: 0}
    for fi24 := 0; fi24 < 3; fi24++ {
        fmt.Println(v1, v4, v1)
        v15 -= (fi24 & (v15 << (((v4 >> ((v4) & 63))) & 63)))
        v2 = append(v2, v15)
        _ = fi24
    }
    v25 := GenS2{F0: 42}
    for fi26 := 0; fi26 < 3; fi26++ {
        v4 += v4
        v1 = v1
        var v27 float64 = 3.14159
        _ = fi26
        _ = v27
    }
    v28 := GenS1{F0: ("hello" + "foo bar"), F1: 708, F2: 42}
    var v29 uint64 = 422
    if (!false) {
        for fi30 := 0; fi30 < 2; fi30++ {
            var v31 uint16 = 65535
            var v32 float32 = 0.1
            _ = fi30
            _ = v31
            _ = v32
        }
        switch "foo bar" {
        case "\n\t":
            switch "\n\t" {
            case "a":
                for fi33 := 0; fi33 < 5; fi33++ {
                    var v34 int = (-(fi33))
                    _ = fi33
                    _ = v34
                }
            case "foo bar":
                v1 -= 0
            case "":
                var v35 int = 32767
                _ = v35
            }
        }
    } else {
        v5[v15] = int(v20)
        v4 -= (v4 & (v4 / (((v4 << ((255) & 63))) | 1)))
        var v36 string = "foo bar"
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
