package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 uint64; F1 uint32; F2 int8 }
func genFn0(p0 int64) int8 {
    var v7 uint64 = 18446744073709551615
    v8 := GenS1{F0: 554, F1: 9223372036854775807, F2: 42}
    _ = p0
    _ = v7
    _ = v8
    return -40
}
func genFn1(p0 string, p1 uint8, p2 int8) string {
    v9 := []int{}
    v10 := GenS1{F0: -18, F1: -189, F2: genFn0(-1)}
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 string = "a"
        _ = fi11
        _ = v12
    }
    var v13 float64 = (-2.25 * 0.1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v13
    _ = v9
    return (p0 + p0)
}

func main() {
    var v1 int16 = 814
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    var v4 string = (("foo bar" + ("a" + (("a" + "foo bar") + (("\n\t" + "hello") + "a")))) + ("foo bar" + ("a" + "foo bar")))
    var v5 int32 = -2147483648
    v6 := make(map[uint32]string)
    v4 = v4
    v14 := GenS1{F0: -806, F1: 32768, F2: 1}
    var v15 float64 = (3.14159 / 1.5)
    if ((v15 <= float64((0.0 - 0.1))) || false) {
        v5 += 42
    }
    if (false || (false && ((!(false || false)) && false))) {
        var v16 int8 = genFn0(9223372036854775807)
        _ = v16
    }
    v17 := GenS2{F0: 32767}
    v18 := []int{382}
    v5 = v5
    for fi19 := 0; fi19 < 3; fi19++ {
        for fi20 := 0; fi20 < 15; fi20++ {
            v21 := GenS3{F0: 4294967295}
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    if ((1 <= 0) && true) {
        v4 = "foo bar"
        fmt.Println(v15)
        var v22 int16 = v1
        _ = v22
    }
    if (127 == -2147483648) {
        v23 := GenS1{F0: 32767, F1: 127, F2: genFn0(-9223372036854775808)}
        fmt.Println(v4, v5)
        _ = v23
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 uint32 = 2147483648
        switch v4 {
        case "hello":
            switch 365 {
            case -897:
                for fi26 := 0; fi26 < 8; fi26++ {
                    v27 := GenS4{F0: v4}
                    _ = fi26
                    _ = v27
                }
            default:
                v28 := GenS5{F0: fi24, F1: true}
                _ = v28
            }
        case "":
            v3[(v4 + ("hello" + ((("a" + (v4 + ("\n\t" + v4))) + v4) + "hello")))] = v4
        }
        _ = fi24
        _ = v25
    }
    var v29 int16 = v1
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    v30 := GenS6{F0: 2, F1: 2147483648, F2: 127}
    v31 := GenS6{F0: 4294967295, F1: 0, F2: 54}
    fmt.Println(v15, v29, v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
