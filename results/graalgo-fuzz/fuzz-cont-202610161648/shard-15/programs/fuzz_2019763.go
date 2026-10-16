package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 int; F1 bool; F2 int32 }
type GenS6 struct { F0 uint64; F1 int16; F2 uint }
type GenS7 struct { F0 uint }
func genFn0(p0 string, p1 uint64) int {
    if (false || false) {
        var v4 int8 = 42
        v5 := []int{42, -1, -546}
        switch "" {
        case "foo bar":
            fmt.Println(v4, p1, v4)
        case "\n\t":
            for fi6 := 0; fi6 < 14; fi6++ {
                fmt.Println(p1, p1, p1)
                _ = fi6
            }
        case "a":
            v7 := GenS1{F0: true}
            _ = v7
        }
        _ = v4
        _ = v5
    } else {
        p1 += 0
    }
    p0 = (((("foo bar" + ((p0 + p0) + p0)) + "") + "hello") + (p0 + p0))
    p0 = ("\n\t" + "hello")
    v8 := GenS1{F0: ((((0.1 - -2.25) + 1e10) - ((((1e10 + ((float64((-2.25 - float32(3.14159))) - ((3.14159 + float64(((1.5 - 3.14159) * 3.14159))) / 1.0)) + 1e10)) * 0.0) / 0.1) - float64(float32(0.1)))) >= 0.0)}
    if true {
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = v8
    return -9223372036854775808
}
func genFn1(p0 int32, p1 int64) uint {
    p1 -= 1
    v9 := GenS2{F0: 4294967295}
    switch "hello" {
    case "\n\t":
        switch 356 {
        case 127:
            p0 = (p0 ^ -1)
        }
    case "":
        if false {
            var v10 float64 = 1.5
            if true {
                p1 -= (-(p1))
            }
            v11 := GenS2{F0: genFn0(("hello" + ""), 42)}
            _ = v10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = v9
    return 65535
}

func main() {
    var v1 uint16 = 663
    var v2 bool = (((float64(1e10) - (1e10 * float64(1.5))) * 1e10) > 3.14159)
    v3 := make(map[uint32]int)
    v2 = true
    if true {
        v12 := GenS1{F0: false}
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 int = fi13
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 15; fi15++ {
            v1 -= v1
            for fi16 := 0; fi16 < 6; fi16++ {
                delete(v3, 65535)
                _ = fi16
            }
            _ = fi15
        }
        _ = v12
    }
    var v17 uint16 = (v1 * 1)
    fmt.Println(v1, v17)
    var v18 int = 134
    var v19 int16 = -32768
    var v20 uint64 = 42
    v21 := GenS3{F0: (v20 >> ((0) & 63)), F1: (v18 << ((570) & 63))}
    v17 = 32767
    v22 := GenS4{F0: v2, F1: 0}
    v23 := v3[266]
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 int16 = v19
        v26 := GenS5{F0: v18, F1: false, F2: 0}
        v27 := []int{fi24, genFn0("hello", 423), v18, (fi24 / ((v23) | 1))}
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    if (!(!((true || (0 > 0)) || v2))) {
        var v28 int32 = 1
        var v29 int64 = 42
        var v30 uint8 = 127
        _ = v28
        _ = v29
        _ = v30
    } else {
        switch 2147483647 {
        case -501:
            v31 := GenS6{F0: 18446744073709551615, F1: (v19 - -32768), F2: 286}
            _ = v31
        default:
            var v32 int8 = -128
            _ = v32
        }
        v18 += -662
    }
    _ = len(v3)
    switch ("a" + "a") {
    case "foo bar":
        switch genFn0("\n\t", 1) {
        case 127:
            var v33 uint = 613
            _ = v33
        case -1:
            fmt.Println(v23)
        case 1:
            switch "foo bar" {
            case "a":
                v34 := GenS7{F0: 490}
                _ = v34
            case "":
                for fi35 := 0; fi35 < 12; fi35++ {
                    v3[309] = genFn0("", 797)
                    _ = fi35
                }
            case "hello":
                var v36 uint32 = 120
                _ = v36
            }
        default:
            v37 := []int{}
            _ = v37
        }
    case "\n\t":
        v3[42] = (v23 ^ v18)
    default:
        v38 := GenS1{F0: false}
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
