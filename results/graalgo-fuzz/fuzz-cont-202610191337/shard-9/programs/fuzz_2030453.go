package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 bool; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint16; F1 string }
func genFn0(p0 uint64) uint64 {
    v6 := GenS1{F0: true}
    if false {
        if true {
            v7 := []int{-733}
            var v8 uint16 = 42
            v9 := GenS2{F0: 97, F1: (-765 <= -1), F2: 18446744073709551615}
            _ = v7
            _ = v8
            _ = v9
        }
        v10 := []int{}
        _ = v10
    } else {
        var v11 uint16 = 65535
        switch ("\n\t" + ("\n\t" + (("a" + ((("hello" + "a") + "") + "foo bar")) + (("foo bar" + "\n\t") + "")))) {
        case "hello":
            v11 = (v11 | (v11 / ((v11) | 1)))
        case "foo bar":
            var v12 string = "\n\t"
            _ = v12
        }
        _ = v11
    }
    _ = p0
    _ = v6
    return p0
}

func main() {
    v1 := []uint32{871}
    v2 := make(map[string]uint32)
    var v3 int16 = -299
    var v4 uint = 119
    var v5 int64 = -975
    if (((1e10 + (0.1 * 0.0)) >= 3.14159) && false) {
        for fi13 := 0; fi13 < 17; fi13++ {
            fmt.Println(v5)
            _ = fi13
        }
    }
    fmt.Println(v3, v5, v4)
    v14 := GenS3{F0: (!(0 < -9223372036854775808))}
    switch "" {
    case "\n\t":
        var v15 uint64 = 0
        _ = v15
    case "":
        if (2 < 42) {
            if false {
                v16 := GenS2{F0: 127, F1: true, F2: (v4 ^ v4)}
                var v17 uint64 = 0
                _ = v16
                _ = v17
            }
            switch -2147483648 {
            case 32768:
                v2[("" + "foo bar")] = 2147483648
            case 9223372036854775807:
                var v18 uint32 = 901
                _ = v18
            default:
                v3 -= (v3 | v3)
            }
        }
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        v5 = (v5 | (v5 & v5))
        var v20 int64 = (v5 << ((2) & 63))
        _ = fi19
        _ = v20
    }
    switch "a" {
    case "":
        var v21 int16 = v3
        _ = v21
    case "\n\t":
        v2["\n\t"] = 2147483647
    case "foo bar":
        switch "hello" {
        case "a":
            for fi22 := 0; fi22 < 8; fi22++ {
                var v23 float64 = 1.5
                _ = fi22
                _ = v23
            }
        default:
            v4 -= 1
        }
    }
    v24 := GenS4{F0: 65535, F1: (-(v5))}
    if (("foo bar" + "a") > "hello") {
        v1 = append(v1, uint32(v4))
    } else {
        fmt.Println(v3)
        if false {
            v5 = (v5 + (^(v5)))
            if (!((false || (false && true)) && ((v4 ^ (v4 ^ v4)) == (v4 >> ((0) & 63))))) {
                _ = len(v1)
                fmt.Println(v5, v3)
            }
        } else {
            v25 := v2[(("hello" + "\n\t") + "hello")]
            v2[("" + "a")] = 2147483647
            _ = v25
        }
    }
    switch -1 {
    case -1:
        _ = len(v1)
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        if false {
            v3 -= (-(v3))
            fmt.Println(fi26)
        }
        var v27 float64 = 1e10
        _ = len(v2)
        _ = fi26
        _ = v27
    }
    var v28 int32 = 1
    switch ((("hello" + "foo bar") + ("\n\t" + "foo bar")) + "") {
    case "foo bar":
        delete(v2, "a")
    case "":
        v29 := GenS5{F0: 42, F1: "a"}
        _ = v29
    case "\n\t":
        switch -807 {
        case 4294967295:
            var v30 float64 = (-2.25 - 3.14159)
            _ = v30
        case 2147483647:
            var v31 string = (("" + (("" + "hello") + ("a" + ""))) + "")
            _ = v31
        case 32768:
            var v32 float64 = (0.1 / -2.25)
            _ = v32
        }
    default:
        var v33 uint64 = 18446744073709551615
        _ = v33
    }
    fmt.Println(v3, v5, v5)
    var v34 float32 = 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
