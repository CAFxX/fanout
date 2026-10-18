package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int }
type GenS2 struct { F0 int8; F1 int64; F2 int }
type GenS3 struct { F0 int; F1 uint16 }
func genFn0(p0 int32, p1 bool) int32 {
    if (0 < 575) {
        p0 -= 165
    }
    var v6 int32 = p0
    switch 2147483647 {
    case -2147483648:
        for fi7 := 0; fi7 < 2; fi7++ {
            fmt.Println(p1, fi7)
            p1 = true
            _ = fi7
        }
    case -65:
        v6 = v6
    default:
        v8 := GenS1{F0: 0, F1: -518, F2: 4294967295}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 1
}
func genFn1(p0 uint32) int {
    v9 := []int{-2147483648}
    var v10 int8 = -1
    v10 = (v10 & (v10 % ((-122) | 1)))
    v10 -= v10
    _ = p0
    _ = v10
    _ = v9
    return 464
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[string]bool)
    var v3 int32 = 1
    var v4 uint = 360
    v5 := []string{}
    var v11 float64 = (1e10 - 1e10)
    v12 := GenS2{F0: -1, F1: 0, F2: 4294967295}
    v2["foo bar"] = ("hello" == "\n\t")
    var v13 uint32 = v1
    if false {
        v14 := v2["\n\t"]
        if len(v5) > 0 {
            v5[0] = ("a" + "a")
        }
        _ = v14
    }
    v4 -= (v4 + (v4 + 18446744073709551615))
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 int32 = (v3 >> (((^(v4))) & 63))
        v13 += (^(v1))
        var v17 float64 = (v11 / v11)
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = len(v5)
    var v18 int8 = 0
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = len(v2)
    v5 = append(v5, "a")
    var v19 int32 = v3
    if true {
        switch ("foo bar" + ("" + "foo bar")) {
        case "hello":
            var v20 int8 = (^(v18))
            _ = v20
        }
        var v21 int32 = int32(v13)
        var v22 bool = (-2.25 <= 1e10)
        _ = v21
        _ = v22
    }
    var v23 uint8 = 0
    if (32768 <= -442) {
        v24 := GenS3{F0: 305, F1: 955}
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
