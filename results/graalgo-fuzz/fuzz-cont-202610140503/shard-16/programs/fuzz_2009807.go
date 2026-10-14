package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint; F2 uint32 }
func genFn0(p0 int16, p1 int, p2 int16) string {
    var v5 uint = 553
    v6 := GenS1{F0: 255}
    v5 -= v5
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return ("" + "a")
}

func main() {
    var v1 uint = 972
    var v2 uint32 = 0
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]bool)
    switch ((("" + ("foo bar" + "foo bar")) + "\n\t") + "\n\t") {
    case "hello":
        v2 = (^(v2))
    default:
        for fi7 := 0; fi7 < 3; fi7++ {
            for fi8 := 0; fi8 < 3; fi8++ {
                fmt.Println(v1, v1, fi7)
                _ = fi8
            }
            var v9 bool = ((v1 % (((v1 % (((v1 * 1)) | 1))) | 1)) < (v1 / ((862) | 1)))
            _ = fi7
            _ = v9
        }
    }
    var v10 int32 = -52
    v11 := []int{}
    v1 -= (^(v1))
    if (false || ((!(792 == (^(v2)))) && false)) {
        var v12 bool = false
        var v13 int16 = 32767
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: 2}
        _ = v14
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    var v15 float64 = ((0.0 / 1e10) - 0.0)
    if false {
        if len(v11) > 0 {
            v11[0] = 2147483647
        }
        var v16 int = 869
        _ = v16
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        v18 := GenS2{F0: v1, F1: (v2 % ((v2) | 1))}
        for fi19 := 0; fi19 < 2; fi19++ {
            fmt.Println(fi17)
            _ = fi19
        }
        var v20 int8 = 42
        _ = fi17
        _ = v18
        _ = v20
    }
    switch 1 {
    case 4294967295:
        for fi21 := 0; fi21 < 4; fi21++ {
            fmt.Println(v15, fi21)
            _ = fi21
        }
    case 127:
        var v22 uint64 = 18446744073709551615
        _ = v22
    case 591:
        v15 += float64(0.1)
    default:
        for fi23 := 0; fi23 < 18; fi23++ {
            v4[(^(fi23))] = (124 >= 127)
            if (false || (!true)) {
                if false {
                    v24 := GenS3{F0: (563 < 0), F1: (v1 * (^(v1))), F2: 0}
                    _ = v24
                }
                var v25 int8 = -1
                v4[fi23] = true
                _ = v25
            }
            fmt.Println(fi23, fi23)
            _ = fi23
        }
    }
    fmt.Println(v2)
    var v26 int16 = -215
    v1 = v1
    v4[4294967295] = (745 >= 0)
    for fi27 := 0; fi27 < 1; fi27++ {
        var v28 int8 = 100
        if (v10 <= v10) {
            v26 -= (v26 << (((v1 - (v1 ^ (^(v1))))) & 63))
        }
        _ = fi27
        _ = v28
    }
    v10 += -2147483648
    for fi29 := 0; fi29 < 15; fi29++ {
        switch "\n\t" {
        case "a":
            v1 -= 18446744073709551615
        case "\n\t":
            v26 -= v26
        case "hello":
            var v30 int16 = v26
            _ = v30
        default:
            fmt.Println(fi29, v15)
        }
        for fi31 := 0; fi31 < 8; fi31++ {
            if ((873 != 2) && (true || false)) {
                var v32 int16 = v26
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                v32 += v32
                _ = v32
            }
            if len(v11) > 0 {
                v11[0] = -909
            }
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = fi31
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
