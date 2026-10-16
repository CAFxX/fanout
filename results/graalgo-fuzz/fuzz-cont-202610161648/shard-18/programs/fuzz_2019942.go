package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 bool; F1 int8; F2 int16 }
type GenS5 struct { F0 int32; F1 string; F2 int }
type GenS6 struct { F0 uint8; F1 int; F2 int64 }
func genFn0() uint {
    v5 := []int{-2147483648, 32768}
    if true {
        var v6 int = -9223372036854775808
        if false {
            v5 = append(v5, 772)
            v6 += (v6 & 2147483647)
        }
        var v7 uint8 = 1
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: 100, F1: 2}
        for fi9 := 0; fi9 < 11; fi9++ {
            if (9223372036854775807 >= 0) {
                var v10 uint32 = 65535
                v11 := GenS2{F0: (fi9 << ((258) & 63))}
                _ = v10
                _ = v11
            } else {
                var v12 int = (fi9 + (fi9 << ((255) & 63)))
                v13 := GenS3{F0: 553, F1: 4294967295}
                _ = v12
                _ = v13
            }
            v14 := GenS4{F0: (!true), F1: 42, F2: -710}
            _ = fi9
            _ = v14
        }
        _ = v8
    }
    var v15 string = "foo bar"
    if (2147483647 >= -21) {
        if false {
            v15 = ""
            v15 = v15
        } else {
            var v16 float32 = 0.0
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 uint = 4294967295
                v19 := GenS1{F0: 42, F1: 470}
                _ = fi17
                _ = v18
                _ = v19
            }
            v16 = ((0.0 - float32((float64((float32(1.5) / v16)) - 1e10))) / v16)
            _ = v16
        }
    }
    _ = v15
    _ = v5
    return 2
}
func genFn1(p0 uint64, p1 bool) uint8 {
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 bool = (p1 || (1e10 == 1.5))
        for fi22 := 0; fi22 < 5; fi22++ {
            var v23 uint16 = 2
            _ = fi22
            _ = v23
        }
        var v24 int = fi20
        _ = fi20
        _ = v21
        _ = v24
    }
    p0 += 1
    for fi25 := 0; fi25 < 5; fi25++ {
        v26 := []int{fi25, fi25, (fi25 + (fi25 >> ((901) & 63))), 128}
        _ = fi25
        _ = v26
    }
    v27 := []int{}
    _ = p0
    _ = p1
    _ = v27
    return uint8(p0)
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    var v3 float32 = float32(-2.25)
    var v4 int32 = 418
    for fi28 := 0; fi28 < 1; fi28++ {
        delete(v1, 1)
        v29 := []int{(^(fi28)), 128, fi28}
        var v30 int64 = 1
        _ = fi28
        _ = v29
        _ = v30
    }
    switch ("hello" + "\n\t") {
    case "":
        var v31 float64 = 3.14159
        _ = v31
    case "\n\t":
        v3 -= v3
    }
    if (128 != 32768) {
        v32 := GenS2{F0: 0}
        _ = v32
    } else {
        v4 = (v4 % ((42) | 1))
        v2[4294967295] = (0 != 255)
    }
    if true {
        v33 := GenS5{F0: v4, F1: "foo bar", F2: -406}
        v4 = 42
        _ = v33
    } else {
        for fi34 := 0; fi34 < 6; fi34++ {
            var v35 int32 = v4
            v36 := GenS4{F0: (0.1 == float64((0.0 / (1.5 / 1.5)))), F1: 1, F2: -294}
            _ = fi34
            _ = v35
            _ = v36
        }
        v4 -= -708
    }
    v37 := GenS6{F0: genFn1(42, (v4 > (v4 - (v4 | (-(v4)))))), F1: -660, F2: 9223372036854775807}
    if ((790 < 1) && (1e10 == 1.5)) {
        v3 = float32(1e10)
        var v38 uint = 503
        fmt.Println(v3, v4, v38)
        _ = v38
    }
    var v39 uint8 = 127
    v40 := []int{}
    switch ("foo bar" + ("\n\t" + "hello")) {
    case "":
        v39 += 255
    case "a":
        fmt.Println(v39, v4, v39)
    default:
        fmt.Println(v3, v39)
    }
    v39 += 102
    switch ("\n\t" + ("hello" + "foo bar")) {
    case "a":
        if len(v40) >= 2 {
            v40 = v40[1:len(v40)]
        }
    case "\n\t":
        fmt.Println(v4, v39)
    default:
        if false {
            var v41 int16 = -92
            _ = v41
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v39
    _ = v4
    _ = v40
}
