package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 uint8; F1 uint8; F2 bool }
func genFn0(p0 int) int {
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: 356, F1: -67}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := []int{(p0 % (((fi9 % ((-1) | 1))) | 1)), p0}
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: 707, F1: 940}
    switch ("foo bar" + "hello") {
    case "\n\t":
        var v12 uint = 0
        _ = v12
    case "":
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 uint = 155
            v15 := []int{}
            v16 := GenS1{F0: 2, F1: -85}
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
    case "a":
        if ((0 == 42) && (!true)) {
            fmt.Println(p0)
            for fi17 := 0; fi17 < 7; fi17++ {
                switch "foo bar" {
                case "":
                    v18 := GenS1{F0: 42, F1: -128}
                    _ = v18
                case "foo bar":
                    var v19 float64 = -2.25
                    _ = v19
                default:
                    p0 = -1
                }
                _ = fi17
            }
        }
    default:
        for fi20 := 0; fi20 < 9; fi20++ {
            if true {
                p0 = -824
                switch (p0 - 4294967295) {
                case 4294967295:
                    var v21 string = "\n\t"
                    _ = v21
                default:
                    v22 := GenS1{F0: 963, F1: -128}
                    _ = v22
                }
            }
            fmt.Println(p0, fi20)
            v23 := GenS3{F0: 0}
            _ = fi20
            _ = v23
        }
    }
    if true {
        p0 += (p0 | p0)
    }
    _ = p0
    _ = v11
    return (p0 & (^(p0)))
}

func main() {
    var v1 float32 = 1.5
    var v2 string = ""
    var v3 string = (("\n\t" + v2) + v2)
    v4 := make(map[string]int)
    var v5 int64 = 9223372036854775807
    v6 := make(map[uint32]int)
    v24 := GenS3{F0: 1}
    v2 = v3
    var v25 uint16 = 2
    for fi26 := 0; fi26 < 1; fi26++ {
        v2 = v3
        delete(v6, 65535)
        switch (-(fi26)) {
        case 9223372036854775807:
            v5 += (v5 & 1)
        case 42:
            if (0 != 127) {
                v1 = ((((0.0 / 3.14159) / v1) / v1) + -2.25)
                v5 += int64(fi26)
            }
        }
        _ = fi26
    }
    v2 = ("" + (("hello" + v3) + v2))
    fmt.Println(v5, v3)
    v27 := GenS4{F0: 357, F1: 42}
    switch ("" + v3) {
    case "":
        for fi28 := 0; fi28 < 8; fi28++ {
            fmt.Println(v25, fi28, v5)
            if false {
                var v29 uint32 = 2
                v29 -= (v29 << ((739) & 63))
                _ = v29
            } else {
                if false {
                    var v30 uint8 = 42
                    _ = v30
                } else {
                    v31 := v6[2]
                    v31 = (v31 ^ genFn0(int(v5)))
                    v32 := GenS2{F0: 216, F1: 639}
                    _ = v31
                    _ = v32
                }
            }
            _ = fi28
        }
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        fmt.Println(v25)
        _ = fi33
    }
    v34 := GenS1{F0: 42, F1: 42}
    for fi35 := 0; fi35 < 13; fi35++ {
        if (!(8 == 1)) {
            switch v3 {
            case "":
                var v36 int8 = -1
                _ = v36
            case "hello":
                var v37 int = (fi35 << ((167) & 63))
                _ = v37
            default:
                v38 := []int{}
                _ = v38
            }
            switch 1 {
            case -9223372036854775808:
                _ = len(v4)
            }
        }
        var v39 int8 = 127
        var v40 uint64 = 729
        _ = fi35
        _ = v39
        _ = v40
    }
    v41 := GenS5{F0: 255, F1: 128, F2: false}
    v25 = v25
    if false {
        for fi42 := 0; fi42 < 20; fi42++ {
            v3 = (v2 + v3)
            _ = fi42
        }
    } else {
        v43 := []int{}
        _ = v43
    }
    var v44 float64 = (1e10 / 1.0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v41
    _ = v44
    _ = v5
    _ = v6
}
