package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int16; F1 string; F2 uint32 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int32; F1 string; F2 uint8 }
func genFn0() uint16 {
    var v6 int32 = 815
    fmt.Println(v6)
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 bool = (!((((true || (989 > 681)) && ((!(2 < 0)) || ("hello" < "hello"))) || false) || false))
            fmt.Println(v6, v6, fi7)
            v9 = false
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 int64 = 1
    _ = v10
    _ = v6
    return 1
}

func main() {
    var v1 int32 = 1
    var v2 uint = 844
    var v3 uint = 994
    var v4 uint8 = 223
    var v5 uint32 = 2
    v11 := GenS1{F0: 622, F1: uint8(v2)}
    var v12 float32 = 1.5
    if false {
        var v13 uint8 = (v4 - (v4 ^ v4))
        _ = v13
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        if true {
            for fi15 := 0; fi15 < 5; fi15++ {
                v4 += 194
                fmt.Println(fi14, v12)
                _ = fi15
            }
            fmt.Println(v1, v1)
        }
        var v16 uint16 = 255
        for fi17 := 0; fi17 < 6; fi17++ {
            v18 := GenS2{F0: 0, F1: v1}
            v19 := GenS1{F0: 0, F1: v4}
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi14
        _ = v16
    }
    var v20 int16 = -32768
    var v21 uint = (v2 << ((v2) & 63))
    v1 += (-(v1))
    var v22 int32 = (v1 << ((v2) & 63))
    v23 := []int{4294967295, -90, 595, 762}
    v24 := GenS1{F0: 9223372036854775807, F1: v4}
    if true {
        for fi25 := 0; fi25 < 10; fi25++ {
            v26 := GenS3{F0: 491, F1: "foo bar", F2: (v5 * v5)}
            v27 := GenS4{F0: v20, F1: 100}
            fmt.Println(v2, v1)
            _ = fi25
            _ = v26
            _ = v27
        }
        if false {
            var v28 float32 = 3.14159
            _ = v28
        } else {
            var v29 float64 = 0.0
            _ = v29
        }
    }
    if false {
        v22 += (-(v1))
        switch "" {
        case "\n\t":
            v20 -= (v20 - (v20 / ((v20) | 1)))
        case "foo bar":
            v30 := GenS3{F0: v20, F1: "\n\t", F2: v5}
            _ = v30
        case "hello":
            switch ((("foo bar" + "foo bar") + "") + "foo bar") {
            case "foo bar":
                v31 := GenS5{F0: v1, F1: ("foo bar" + "hello"), F2: (v4 & (v4 % ((v4) | 1)))}
                _ = v31
            default:
                v22 -= int32(v5)
            }
        default:
            var v32 uint32 = 790
            _ = v32
        }
        var v33 string = "hello"
        _ = v33
    } else {
        v1 = v1
        var v34 uint8 = v4
        _ = v34
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        if len(v23) > 0 {
            v23[0] = (fi35 << ((v2) & 63))
        }
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
