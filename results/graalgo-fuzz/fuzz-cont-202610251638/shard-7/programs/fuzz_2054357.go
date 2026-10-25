package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint64; F2 bool }
type GenS5 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint; F1 uint16; F2 uint }
type GenS9 struct { F0 int32; F1 uint64; F2 uint8 }
func genFn0(p0 int16, p1 int64, p2 bool) string {
    v4 := GenS1{F0: p1}
    for fi5 := 0; fi5 < 3; fi5++ {
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := []int{fi5}
            _ = fi6
            _ = v7
        }
        p0 = p0
        var v8 uint64 = 4294967295
        _ = fi5
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := GenS2{F0: ""}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return ("a" + ("foo bar" + ((("hello" + "hello") + "a") + "a")))
}
func genFn1(p0 uint16, p1 int) int16 {
    v11 := GenS3{F0: 1}
    v12 := []int{(^(p1)), p1}
    for fi13 := 0; fi13 < 8; fi13++ {
        if (true || false) {
            var v14 string = ("foo bar" + ("hello" + (("hello" + "\n\t") + "")))
            _ = v14
        }
        if true {
            var v15 int32 = -1
            var v16 int8 = 127
            _ = v15
            _ = v16
        } else {
            for fi17 := 0; fi17 < 3; fi17++ {
                _ = len(v12)
                p0 -= 0
                _ = fi17
            }
        }
        switch "hello" {
        case "foo bar":
            if ((!true) && ((958 > 815) || true)) {
                p0 += 42
                v12 = append(v12, -262)
                var v18 uint = 18446744073709551615
                _ = v18
            } else {
                for fi19 := 0; fi19 < 16; fi19++ {
                    var v20 uint64 = 18446744073709551615
                    var v21 uint16 = (p0 | p0)
                    var v22 uint8 = 128
                    p1 += fi19
                    var v23 int = -155
                    v21 = 42
                    var v24 int8 = int8(p0)
                    v25 := GenS4{F0: 642, F1: 2, F2: (2147483648 >= 2147483647)}
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                    _ = v23
                    _ = v24
                    _ = v25
                }
            }
        default:
            v26 := GenS5{F0: 553, F1: 889, F2: 116}
            _ = v26
        }
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return -433
}

func main() {
    var v1 int64 = -697
    var v2 int64 = (v1 | -955)
    v3 := []string{}
    var v27 uint8 = 2
    switch "a" {
    case "a":
        var v28 int32 = -1
        _ = v28
    case "hello":
        var v29 uint32 = 407
        _ = v29
    default:
        v2 -= (v2 << ((349) & 63))
    }
    if (!true) {
        v30 := GenS6{F0: genFn1(65535, 0)}
        var v31 uint16 = 0
        for fi32 := 0; fi32 < 19; fi32++ {
            v33 := GenS7{F0: v31}
            v34 := GenS7{F0: 255}
            var v35 uint64 = uint64(v27)
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
        _ = v30
        _ = v31
    }
    fmt.Println(v2, v27)
    v2 += v2
    fmt.Println(v2, v2)
    v36 := GenS8{F0: 2, F1: 432, F2: 573}
    for fi37 := 0; fi37 < 19; fi37++ {
        v38 := GenS9{F0: 1, F1: 260, F2: 221}
        var v39 uint = 255
        for fi40 := 0; fi40 < 20; fi40++ {
            var v41 float64 = -2.25
            v2 = 1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi40
            _ = v41
        }
        _ = fi37
        _ = v38
        _ = v39
    }
    v27 -= v27
    switch "" {
    case "a":
        _ = len(v3)
    case "hello":
        var v42 string = "a"
        _ = v42
    }
    v27 = (^(v27))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 -= 1
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    v27 += v27
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v36
}
