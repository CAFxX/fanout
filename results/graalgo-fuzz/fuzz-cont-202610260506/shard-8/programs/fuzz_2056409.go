package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 uint }
func genFn0(p0 int) uint8 {
    p0 = (p0 << ((1) & 63))
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(fi6)
        p0 = (^(fi6))
        for fi7 := 0; fi7 < 8; fi7++ {
            v8 := GenS1{F0: 229}
            var v9 int64 = -1
            var v10 uint8 = 191
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    _ = p0
    return 212
}
func genFn1(p0 bool, p1 int8) bool {
    var v11 int64 = 1
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 string = ("foo bar" + ("\n\t" + (("\n\t" + "foo bar") + "foo bar")))
            _ = fi13
            _ = v14
        }
        v15 := []int{}
        _ = fi12
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v11
    return (1 >= p1)
}

func main() {
    var v1 int8 = 127
    v2 := []string{"foo bar", "", "", "a"}
    var v3 int32 = 1
    v4 := []string{}
    v5 := make(map[uint32]bool)
    var v16 int16 = -500
    switch "" {
    case "\n\t":
        v3 = 42
    case "hello":
        if (229 > 2) {
            _ = len(v2)
        }
    }
    v16 += (v16 & (^(v16)))
    if false {
        var v17 uint16 = 42
        fmt.Println(v16, v3, v1)
        fmt.Println(v16, v17)
        _ = v17
    }
    if true {
        var v18 float64 = -2.25
        for fi19 := 0; fi19 < 14; fi19++ {
            for fi20 := 0; fi20 < 11; fi20++ {
                v1 += v1
                fmt.Println(v16, v3)
                _ = fi20
            }
            var v21 uint64 = 2
            fmt.Println(v21, v16)
            _ = fi19
            _ = v21
        }
        if false {
            v16 = (v16 ^ (v16 - v16))
        }
        _ = v18
    }
    v16 += (v16 & (v16 | v16))
    var v22 uint8 = 255
    v1 += v1
    var v23 bool = false
    v24 := GenS2{F0: 42, F1: -1}
    if (v3 < (v3 & v3)) {
        for fi25 := 0; fi25 < 6; fi25++ {
            v1 += v1
            if (-640 == int64(v16)) {
                var v26 int64 = 42
                _ = len(v4)
                _ = v26
            }
            _ = fi25
        }
        v27 := GenS1{F0: (v22 % ((42) | 1))}
        _ = v27
    }
    v28 := GenS2{F0: 0, F1: v16}
    if v23 {
        var v29 string = "a"
        if (false && v23) {
            v3 = (v3 & (v3 >> ((65535) & 63)))
            v22 -= (v22 ^ (^(v22)))
        }
        _ = v29
    }
    v30 := GenS3{F0: 65535}
    switch 635 {
    case -9223372036854775808:
        var v31 float64 = float64(-2.25)
        _ = v31
    case 9223372036854775807:
        fmt.Println(v16, v22)
    case -1:
        _ = len(v4)
    default:
        if ((18446744073709551615 == 18446744073709551615) || (-1 > -1)) {
            for fi32 := 0; fi32 < 3; fi32++ {
                var v33 int32 = -1
                fmt.Println(v22, v33, v33)
                _ = fi32
                _ = v33
            }
            v34 := GenS4{F0: ((v1 | v1) < (v1 - -10)), F1: 0}
            _ = v34
        } else {
            var v35 uint32 = 559
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
