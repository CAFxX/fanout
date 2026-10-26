package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint32; F2 bool }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint16, p1 uint32, p2 int32) int16 {
    var v6 float32 = 1e10
    p0 += p0
    switch -100 {
    case 1:
        if false {
            for fi7 := 0; fi7 < 4; fi7++ {
                v8 := []int{(-(fi7)), fi7, (fi7 & fi7)}
                var v9 bool = ((true && (false && false)) && (((0.1 - (3.14159 / (0.0 / (3.14159 - float64(v6))))) / (-2.25 + ((0.0 / float64((-2.25 / 1.0))) - ((1e10 * 1.5) / -2.25)))) > -2.25))
                for fi10 := 0; fi10 < 16; fi10++ {
                    var v11 int64 = 9223372036854775807
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    _ = fi10
                    _ = v11
                }
                _ = fi7
                _ = v8
                _ = v9
            }
        }
    case -9223372036854775808:
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 uint32 = (p1 * p1)
            v13 -= (^(v13))
            _ = fi12
            _ = v13
        }
    case 9223372036854775807:
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 bool = (65535 > p1)
            _ = fi14
            _ = v15
        }
    default:
        fmt.Println(p2, p0)
    }
    v6 += (((v6 + 1e10) / -2.25) * v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 32767
}

func main() {
    var v1 uint16 = 677
    v2 := []int{4294967295, -792, -9223372036854775808}
    var v3 uint64 = 4294967295
    var v4 uint32 = 47
    var v5 uint16 = (v1 % ((0) | 1))
    v16 := GenS1{F0: -2147483648}
    if ("foo bar" == "") {
        v1 += (^(v5))
        for fi17 := 0; fi17 < 13; fi17++ {
            var v18 uint32 = (v4 / ((65535) | 1))
            fmt.Println(v5, v5)
            fmt.Println(fi17, v5)
            _ = fi17
            _ = v18
        }
        v19 := GenS2{F0: "", F1: (v4 * 1), F2: (true && true)}
        _ = v19
    } else {
        var v20 string = "\n\t"
        var v21 string = "foo bar"
        _ = v20
        _ = v21
    }
    fmt.Println(v3, v5)
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 bool = true
        var v24 uint = 255
        _ = fi22
        _ = v23
        _ = v24
    }
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    var v25 string = ""
    v26 := GenS2{F0: "a", F1: 77, F2: false}
    v27 := GenS2{F0: v25, F1: v4, F2: (1e10 >= 1.5)}
    var v28 uint16 = v1
    fmt.Println(v4, v5)
    var v29 int64 = -9223372036854775808
    fmt.Println(v1, v3, v5)
    for fi30 := 0; fi30 < 9; fi30++ {
        v5 -= (v5 ^ (^(v28)))
        fmt.Println(v28, v1)
        var v31 float64 = 3.14159
        _ = fi30
        _ = v31
    }
    v5 = v28
    v29 = v29
    if (127 >= 32767) {
        var v32 int8 = 109
        _ = v32
    } else {
        v33 := GenS3{F0: -593}
        if len(v2) > 0 {
            v2[0] = 32767
        }
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
