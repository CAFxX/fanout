package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 string; F2 int16 }
type GenS4 struct { F0 int32; F1 int16; F2 uint8 }
type GenS5 struct { F0 int64; F1 bool }
func genFn0(p0 bool, p1 int64, p2 int8) uint {
    fmt.Println(p0, p0)
    p1 += (p1 + (p1 >> ((255) & 63)))
    if p0 {
        fmt.Println(p1)
        var v4 float64 = ((1e10 / 1.5) / 1.0)
        if ((0.0 / -2.25) < v4) {
            v4 -= (v4 - -2.25)
        } else {
            switch 2147483647 {
            case 63:
                fmt.Println(v4, p0, v4)
            case 128:
                fmt.Println(p2)
            case 127:
                v5 := []int{}
                _ = v5
            default:
                switch 1 {
                case 32767:
                    p2 = (p2 << ((113) & 63))
                case 127:
                    p0 = (-1 == 2147483647)
                default:
                    fmt.Println(p2)
                }
            }
        }
        _ = v4
    }
    for fi6 := 0; fi6 < 20; fi6++ {
        p2 = p2
        if (100 <= (p2 | int8(p1))) {
            var v7 uint16 = 163
            fmt.Println(fi6, p1, fi6)
            _ = v7
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]uint32)
    var v3 int16 = -32768
    switch -658 {
    case 1:
        fmt.Println(v3)
    case 106:
        fmt.Println(v3)
    default:
        for fi8 := 0; fi8 < 5; fi8++ {
            switch ((("" + "\n\t") + "\n\t") + "foo bar") {
            case "":
                v9 := []int{-1}
                _ = v9
            case "a":
                if true {
                    var v10 uint = 892
                    v10 = (v10 + genFn0(true, -9223372036854775808, 35))
                    fmt.Println(v10, fi8, v10)
                    _ = v10
                }
            default:
                v3 += (v3 % ((v3) | 1))
            }
            _ = fi8
        }
    }
    delete(v2, 23)
    fmt.Println(v3)
    if false {
        delete(v2, 4294967295)
        if false {
            var v11 uint = 255
            var v12 uint8 = 1
            _ = v11
            _ = v12
        } else {
            if (100 >= -128) {
                v13 := GenS1{F0: 1}
                var v14 float64 = (1.5 - ((0.1 + 1e10) * 1.5))
                switch ("foo bar" + ("" + ("a" + "\n\t"))) {
                case "a":
                    v3 = v3
                case "foo bar":
                    var v15 uint8 = 128
                    _ = v15
                case "\n\t":
                    fmt.Println(v3)
                default:
                    v16 := GenS2{F0: -9223372036854775808}
                    _ = v16
                }
                _ = v13
                _ = v14
            }
        }
        var v17 uint16 = 514
        _ = v17
    } else {
        var v18 uint64 = 0
        fmt.Println(v3)
        v18 += v18
        _ = v18
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 string = "a"
        _ = fi19
        _ = v20
    }
    var v21 uint64 = 18446744073709551615
    for fi22 := 0; fi22 < 6; fi22++ {
        switch (fi22 | (fi22 % ((32768) | 1))) {
        case -2147483648:
            var v23 uint16 = 42
            _ = v23
        default:
            if true {
                v24 := GenS3{F0: 0, F1: (("a" + "foo bar") + ("a" + ((("foo bar" + ((("a" + "\n\t") + ((("\n\t" + "a") + ("foo bar" + "\n\t")) + (("foo bar" + ("hello" + "a")) + ("hello" + "hello")))) + "foo bar")) + "hello") + "foo bar"))), F2: (v3 | (v3 / ((1) | 1)))}
                v3 = (v3 | 1)
                _ = v24
            }
        }
        v25 := GenS1{F0: 42}
        v26 := GenS4{F0: 97, F1: (^(v3)), F2: 2}
        _ = fi22
        _ = v25
        _ = v26
    }
    v27 := GenS5{F0: 0, F1: true}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
}
