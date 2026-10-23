package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int32; F1 string; F2 bool }
type GenS4 struct { F0 int8; F1 uint8; F2 uint }
func genFn0() int32 {
    switch "hello" {
    case "\n\t":
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 float32 = 3.14159
            _ = fi6
            _ = v7
        }
    default:
        var v8 int64 = -9223372036854775808
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int32 = 0
        _ = fi9
        _ = v10
    }
    return -1
}
func genFn1(p0 int32, p1 uint8) uint32 {
    v11 := []int{-1}
    var v12 int8 = 1
    v13 := GenS1{F0: 255, F1: false}
    switch "\n\t" {
    case "":
        fmt.Println(v12)
    case "a":
        v12 += 0
    case "\n\t":
        if true {
            var v14 uint = 2
            _ = v14
        }
    }
    if len(v11) > 0 {
        v11[0] = -1
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return 42
}

func main() {
    v1 := []uint32{}
    var v2 float32 = -2.25
    v3 := []string{"a"}
    v4 := []uint32{}
    v5 := make(map[string]string)
    v15 := GenS1{F0: 0, F1: (!(false && (-803 > 42)))}
    var v16 uint = 788
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 string = "foo bar"
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        fmt.Println(v16, v2, v16)
        fmt.Println(v2, v2)
        var v20 uint32 = 2147483648
        _ = fi19
        _ = v20
    }
    var v21 int8 = 127
    v22 := GenS2{F0: 831, F1: (!false)}
    if (255 == v16) {
        fmt.Println(v16)
    }
    v23 := GenS2{F0: 491, F1: (false && true)}
    delete(v5, ("hello" + ""))
    if ((0 <= 924) && true) {
        v24 := GenS3{F0: genFn0(), F1: "hello", F2: (false || false)}
        v16 += (v16 >> ((v16) & 63))
        _ = v24
    }
    v25 := GenS1{F0: (^(v16)), F1: (2 != 414)}
    for fi26 := 0; fi26 < 3; fi26++ {
        v27 := GenS3{F0: genFn0(), F1: "a", F2: (false || false)}
        _ = fi26
        _ = v27
    }
    v5[(("\n\t" + "\n\t") + "hello")] = ""
    for fi28 := 0; fi28 < 15; fi28++ {
        switch fi28 {
        case 128:
            switch fi28 {
            case 412:
                var v29 float32 = v2
                _ = v29
            case 0:
                v16 -= (v16 | (v16 + v16))
            case -9223372036854775808:
                fmt.Println(fi28)
            }
        case 32768:
            var v30 int64 = 0
            _ = v30
        default:
            fmt.Println(v2)
        }
        for fi31 := 0; fi31 < 20; fi31++ {
            if ((!(true || (("hello" + ("" + ("foo bar" + "hello"))) != "\n\t"))) || true) {
                v16 -= 2
                if true {
                    v2 -= (0.0 * 3.14159)
                } else {
                    _ = len(v5)
                }
                var v32 uint32 = 4294967295
                _ = v32
            }
            if (v2 > v2) {
                v2 = 3.14159
                fmt.Println(v16, fi31)
                v33 := GenS3{F0: 2147483647, F1: ("" + "hello"), F2: false}
                _ = v33
            } else {
                v34 := GenS4{F0: v21, F1: 38, F2: (v16 * (v16 + v16))}
                v16 -= (v16 - 65535)
                _ = v34
            }
            v16 = (^(v16))
            _ = fi31
        }
        _ = fi28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
