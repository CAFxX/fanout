package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int64; F1 int16 }
func genFn0(p0 uint8, p1 int32, p2 int64) uint8 {
    p2 = p2
    v4 := GenS1{F0: 127}
    if (false && false) {
        switch 0 {
        case 856:
            var v5 uint = 477
            _ = v5
        default:
            var v6 uint64 = 42
            _ = v6
        }
        var v7 uint16 = 65535
        switch 127 {
        case -9223372036854775808:
            for fi8 := 0; fi8 < 4; fi8++ {
                for fi9 := 0; fi9 < 5; fi9++ {
                    p0 = (p0 >> ((65535) & 63))
                    v10 := GenS1{F0: 42}
                    _ = fi9
                    _ = v10
                }
                var v11 int8 = -11
                for fi12 := 0; fi12 < 18; fi12++ {
                    fmt.Println(p2, p0, fi8)
                    var v13 uint16 = v7
                    p1 = 2147483647
                    var v14 float64 = (1.5 - ((1e10 + (0.0 / float64(-2.25))) - -2.25))
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                _ = fi8
                _ = v11
            }
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p0 % (((^(p0))) | 1))
}

func main() {
    var v1 bool = (!("hello" < ("\n\t" + "\n\t")))
    v2 := []int{128, 0, 32767, 274}
    var v3 uint64 = 42
    v1 = false
    var v15 bool = v1
    var v16 int8 = 100
    switch 127 {
    case 32768:
        if ((^(v3)) <= (^(v3))) {
            fmt.Println(v1, v1, v16)
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 uint16 = 255
                v19 := GenS1{F0: (-(v16))}
                v2 = append(v2, (-(fi17)))
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 int64 = 1
            _ = v20
        }
    case 127:
        fmt.Println(v3)
    case -2147483648:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 int16 = 32767
        for fi23 := 0; fi23 < 14; fi23++ {
            v22 = v22
            switch "a" {
            case "foo bar":
                v1 = (!true)
            case "a":
                fmt.Println(v3, v22, v15)
            case "":
                var v24 float32 = 0.1
                _ = v24
            }
            fmt.Println(v22)
            _ = fi23
        }
        _ = fi21
        _ = v22
    }
    fmt.Println(v15, v16, v3)
    for fi25 := 0; fi25 < 16; fi25++ {
        fmt.Println(v16, v15, v3)
        _ = fi25
    }
    if (uint16(v3) <= 65535) {
        var v26 int32 = 2147483647
        v2 = append(v2, 2147483647)
        _ = v26
    } else {
        for fi27 := 0; fi27 < 16; fi27++ {
            for fi28 := 0; fi28 < 7; fi28++ {
                if len(v2) > 0 {
                    v2[0] = fi27
                }
                v29 := GenS2{F0: "", F1: 765}
                var v30 int16 = 42
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = fi27
        }
        var v31 uint8 = 2
        _ = v31
    }
    switch 128 {
    case 32768:
        var v32 int16 = 32767
        _ = v32
    case 4294967295:
        v15 = v15
    case -2147483648:
        v3 += v3
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        v3 = (v3 << ((0) & 63))
        switch "\n\t" {
        case "a":
            var v34 uint = 42
            _ = v34
        default:
            if false {
                var v35 int64 = 0
                if v1 {
                    v35 -= 278
                }
                var v36 int32 = -2147483648
                _ = v35
                _ = v36
            } else {
                switch "a" {
                case "\n\t":
                    v2 = append(v2, fi33)
                default:
                    fmt.Println(v3)
                }
            }
        }
        _ = fi33
    }
    for fi37 := 0; fi37 < 9; fi37++ {
        var v38 uint16 = 514
        var v39 string = ("a" + ("" + ""))
        v40 := GenS3{F0: 1, F1: 1}
        _ = fi37
        _ = v38
        _ = v39
        _ = v40
    }
    v1 = true
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
