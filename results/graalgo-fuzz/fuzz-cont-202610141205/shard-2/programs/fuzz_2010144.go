package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 bool; F1 int16 }
func genFn0(p0 uint64, p1 uint64, p2 int8) int8 {
    switch "\n\t" {
    case "hello":
        p2 -= -123
    }
    var v4 uint16 = 125
    v5 := GenS1{F0: 42, F1: (("a" + ("foo bar" + (("" + "") + ("" + "\n\t")))) + ((("\n\t" + ((("a" + "") + "foo bar") + "hello")) + "hello") + ("a" + "a")))}
    for fi6 := 0; fi6 < 10; fi6++ {
        p1 = 800
        fmt.Println(p1)
        _ = fi6
    }
    p0 -= (p1 % ((uint64(v4)) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p2
}
func genFn1(p0 int16, p1 uint32) uint8 {
    switch ("\n\t" + ("hello" + "foo bar")) {
    case "\n\t":
        if false {
            p1 = p1
            var v7 uint64 = 4294967295
            _ = v7
        } else {
            var v8 int8 = 17
            if false {
                for fi9 := 0; fi9 < 17; fi9++ {
                    var v10 int8 = (v8 >> ((0) & 63))
                    v11 := []int{(fi9 ^ fi9), -207, fi9}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                for fi12 := 0; fi12 < 10; fi12++ {
                    var v13 int = (fi12 >> ((255) & 63))
                    v8 -= genFn0(4294967295, 42, genFn0(440, 18446744073709551615, genFn0(18446744073709551615, 1, v8)))
                    v14 := GenS1{F0: (-(v8)), F1: "hello"}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                v15 := []int{}
                _ = v15
            }
            _ = v8
        }
    default:
        v16 := []int{1}
        _ = v16
    }
    var v17 float64 = 0.1
    _ = p0
    _ = p1
    _ = v17
    return 235
}

func main() {
    var v1 int8 = 100
    var v2 int32 = 2147483647
    var v3 uint8 = 0
    fmt.Println(v3, v3)
    v18 := []int{32767, 0, 4294967295}
    for fi19 := 0; fi19 < 4; fi19++ {
        for fi20 := 0; fi20 < 13; fi20++ {
            _ = len(v18)
            v1 -= -1
            if true {
                v1 += v1
            }
            var v21 uint8 = v3
            v2 -= v2
            _ = fi20
            _ = v21
        }
        var v22 uint = 1
        _ = fi19
        _ = v22
    }
    var v23 int16 = int16(v2)
    v24 := GenS1{F0: v1, F1: ("hello" + ("" + ("hello" + (("a" + "\n\t") + (("" + "foo bar") + "foo bar")))))}
    v3 -= (v3 + 42)
    for fi25 := 0; fi25 < 10; fi25++ {
        for fi26 := 0; fi26 < 8; fi26++ {
            _ = len(v18)
            _ = fi26
        }
        _ = fi25
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        var v28 string = ("" + "foo bar")
        if true {
            var v29 int = (fi27 + (-(fi27)))
            v3 += 127
            v29 = fi27
            _ = v29
        } else {
            fmt.Println(v23)
            for fi30 := 0; fi30 < 4; fi30++ {
                var v31 int8 = v1
                v18 = append(v18, fi27)
                v3 = (v3 << ((1) & 63))
                _ = fi30
                _ = v31
            }
        }
        v32 := GenS2{F0: (1000 >= (v23 * v23)), F1: (^(v23))}
        _ = fi27
        _ = v28
        _ = v32
    }
    v3 = v3
    fmt.Println(v3)
    if (false && (65535 < 409)) {
        if len(v18) > 0 {
            v18[0] = 0
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
