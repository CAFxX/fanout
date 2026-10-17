package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint; F2 int32 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 uint; F1 int }
type GenS8 struct { F0 int16; F1 int }
func genFn0(p0 int, p1 int) string {
    v4 := GenS1{F0: 2, F1: 128}
    if false {
        v5 := GenS1{F0: 925, F1: 42}
        var v6 uint8 = 127
        _ = v5
        _ = v6
    } else {
        var v7 uint32 = 2
        if false {
            var v8 int16 = 1000
            fmt.Println(v7, p1)
            _ = v8
        } else {
            v7 += 65535
            v9 := []int{-434, (p1 / (((p0 + (p0 & (p1 * 4294967295)))) | 1))}
            v10 := GenS2{F0: 800, F1: 4294967295, F2: 2147483647}
            _ = v10
            _ = v9
        }
        _ = v7
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v4
    return "foo bar"
}
func genFn1(p0 int8) int32 {
    v11 := GenS1{F0: 1, F1: 110}
    p0 = (p0 & 117)
    _ = p0
    _ = v11
    return 2147483647
}

func main() {
    var v1 string = "a"
    v2 := make(map[string]int)
    v3 := []string{}
    v1 = v1
    if (true && ((true && (!(113 >= 211))) && true)) {
        fmt.Println(v1)
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        v1 = ("" + v1)
        _ = fi12
    }
    if true {
        v1 = v1
        v13 := GenS1{F0: 619, F1: 109}
        v14 := GenS3{F0: 127, F1: 112}
        _ = v13
        _ = v14
    }
    fmt.Println(v1)
    for fi15 := 0; fi15 < 20; fi15++ {
        fmt.Println(fi15)
        if (4294967295 != 492) {
            fmt.Println(v1, fi15)
        } else {
            var v16 int16 = -997
            _ = v16
        }
        for fi17 := 0; fi17 < 6; fi17++ {
            v1 = v1
            fmt.Println(v1)
            v1 = v1
            v1 = v1
            _ = fi17
        }
        _ = fi15
    }
    var v18 uint64 = 419
    var v19 int64 = 59
    var v20 int64 = (v19 % ((-902) | 1))
    switch 9223372036854775807 {
    case -710:
        v19 += (v20 * (^(v20)))
    case -202:
        switch 42 {
        case 128:
            v21 := GenS4{F0: 527, F1: 142}
            _ = v21
        case -625:
            v22 := GenS5{F0: v20}
            _ = v22
        default:
            v19 -= (v19 | 0)
        }
    }
    v20 = (^(v20))
    var v23 bool = (749 > 2)
    if false {
        fmt.Println(v18, v23, v1)
        var v24 int32 = -1
        v25 := GenS6{F0: ("a" + "a"), F1: 0}
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        v23 = v23
        _ = fi26
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        for fi28 := 0; fi28 < 3; fi28++ {
            var v29 float64 = ((0.1 / 1.0) * 1e10)
            var v30 uint16 = 1
            if (!(1000 != 1)) {
                var v31 float32 = 1e10
                v32 := GenS7{F0: 4294967295, F1: (-(fi28))}
                _ = v31
                _ = v32
            }
            _ = fi28
            _ = v29
            _ = v30
        }
        if true {
            if (float64(1.5) >= (-2.25 * -2.25)) {
                var v33 bool = v23
                v1 = v1
                _ = v33
            } else {
                for fi34 := 0; fi34 < 19; fi34++ {
                    v18 = 1
                    _ = fi34
                }
            }
            v35 := GenS8{F0: 42, F1: fi27}
            v20 -= (v20 >> ((0) & 63))
            _ = v35
        } else {
            v19 = 275
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi27
    }
    for fi36 := 0; fi36 < 16; fi36++ {
        for fi37 := 0; fi37 < 8; fi37++ {
            v1 = v1
            var v38 int64 = v19
            _ = fi37
            _ = v38
        }
        v18 = (v18 ^ v18)
        var v39 uint32 = 65535
        _ = fi36
        _ = v39
    }
    if (!(-32768 >= 954)) {
        v40 := GenS6{F0: v1, F1: 4294967295}
        fmt.Println(v20, v1)
        _ = v40
    }
    fmt.Println(v18, v20)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
}
