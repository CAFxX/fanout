package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 bool; F2 uint16 }
type GenS3 struct { F0 int16; F1 uint }
func genFn0(p0 uint16, p1 int64) string {
    fmt.Println(p0)
    if true {
        for fi4 := 0; fi4 < 4; fi4++ {
            var v5 string = "a"
            var v6 int16 = 0
            _ = fi4
            _ = v5
            _ = v6
        }
        fmt.Println(p0, p0)
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := []int{int(p0), (fi7 | (fi7 | (fi7 >> ((255) & 63))))}
            _ = fi7
            _ = v8
        }
    } else {
        if true {
            v9 := []int{-590, 4294967295, 2147483647, -9223372036854775808}
            var v10 float64 = 0.0
            _ = v10
            _ = v9
        }
        p0 += 565
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    return "hello"
}
func genFn1(p0 int32, p1 uint) int64 {
    v11 := []int{42, 128, 32767}
    p1 += p1
    _ = p0
    _ = p1
    _ = v11
    return -1
}

func main() {
    var v1 int = -656
    var v2 uint8 = 52
    var v3 uint8 = (v2 * (v2 * 1))
    v2 -= (v2 % (((v2 | (v2 >> ((590) & 63)))) | 1))
    var v12 float64 = (0.0 / float64((0.1 - 3.14159)))
    v13 := []int{}
    if len(v13) > 0 {
        v13[0] = v1
    }
    v14 := GenS1{F0: genFn1(0, 2)}
    if (false || true) {
        switch (v1 >> ((662) & 63)) {
        case -118:
            v2 = (v3 + v3)
        default:
            var v15 uint32 = 42
            _ = v15
        }
    }
    if (1e10 != -2.25) {
        for fi16 := 0; fi16 < 14; fi16++ {
            v12 = 3.14159
            _ = fi16
        }
    }
    v17 := GenS2{F0: 1, F1: ((float64(1.5) - ((v12 / 1.5) - v12)) == v12), F2: 979}
    v1 += 950
    switch "\n\t" {
    case "\n\t":
        var v18 int32 = -380
        _ = v18
    default:
        var v19 float64 = v12
        _ = v19
    }
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    fmt.Println(v3)
    var v20 bool = (false && (!false))
    v20 = false
    switch (^(v1)) {
    case 4294967295:
        v3 = (v3 ^ (v2 | v2))
    case 32768:
        for fi21 := 0; fi21 < 2; fi21++ {
            switch "foo bar" {
            case "a":
                fmt.Println(fi21, v20)
            case "foo bar":
                var v22 uint8 = 42
                _ = v22
            case "":
                switch "" {
                case "\n\t":
                    v23 := GenS3{F0: -1, F1: 88}
                    _ = v23
                case "":
                    v2 = v2
                default:
                    _ = len(v13)
                }
            default:
                for fi24 := 0; fi24 < 9; fi24++ {
                    var v25 float64 = (1e10 * v12)
                    var v26 string = "a"
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
            for fi27 := 0; fi27 < 7; fi27++ {
                var v28 uint64 = 18446744073709551615
                _ = fi27
                _ = v28
            }
            var v29 uint = 0
            _ = fi21
            _ = v29
        }
    default:
        v20 = (((0 >= 131) || ("a" > (("hello" + "") + ("\n\t" + "hello")))) && (!(75 == 0)))
    }
    _ = len(v13)
    var v30 bool = ("\n\t" <= ("" + "\n\t"))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v30
}
