package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint16; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint32; F2 bool }
func genFn0(p0 int64) int32 {
    for fi6 := 0; fi6 < 16; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            fmt.Println(fi6, fi6)
            _ = fi7
        }
        _ = fi6
    }
    v8 := GenS1{F0: "a"}
    var v9 float64 = float64(-2.25)
    var v10 float32 = (0.1 / 1.0)
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 bool = (3.14159 > -2.25)
    v2 := make(map[uint32]uint32)
    var v3 float32 = 3.14159
    var v4 uint64 = 825
    v5 := []uint32{2147483648, 65535, 285, 65535}
    if true {
        if (-1 != -921) {
            v1 = v1
            v3 = v3
            v4 -= v4
        }
        v3 = ((v3 + v3) + float32(3.14159))
    }
    v3 += float32(1.5)
    fmt.Println(v4, v1, v3)
    fmt.Println(v1, v3)
    var v11 int = -663
    if (!v1) {
        switch (v11 << ((141) & 63)) {
        case 32767:
            v12 := GenS1{F0: ("" + "hello")}
            _ = v12
        case 127:
            v13 := GenS2{F0: (("hello" + "a") > "foo bar"), F1: uint16(v4), F2: 4294967295}
            _ = v13
        }
        var v14 uint64 = v4
        v4 = (v4 << ((255) & 63))
        _ = v14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        switch ("\n\t" + ("a" + "")) {
        case "a":
            v2[2] = 813
        default:
            var v16 uint8 = 127
            _ = v16
        }
        for fi17 := 0; fi17 < 2; fi17++ {
            if v1 {
                var v18 int16 = int16(fi15)
                _ = len(v2)
                if (((v4 >= v4) && (127 > -128)) && (264 > 255)) {
                    var v19 bool = (v1 && false)
                    var v20 int32 = 42
                    _ = v19
                    _ = v20
                } else {
                    v18 = -704
                    v21 := GenS3{F0: v11}
                    _ = v21
                }
                _ = v18
            } else {
                var v22 float64 = ((3.14159 + 1e10) * 1.5)
                var v23 uint8 = 76
                _ = v22
                _ = v23
            }
            v24 := GenS4{F0: 65535}
            v25 := GenS5{F0: fi17, F1: 1, F2: v1}
            delete(v2, 1)
            _ = fi17
            _ = v24
            _ = v25
        }
        var v26 uint8 = 103
        _ = fi15
        _ = v26
    }
    v3 = v3
    v3 = v3
    _ = len(v2)
    var v27 uint64 = (v4 ^ 2)
    switch (("hello" + "") + (("hello" + "foo bar") + "hello")) {
    case "foo bar":
        v27 = (^(v4))
    case "hello":
        var v28 int32 = 1
        _ = v28
    case "":
        var v29 int8 = 1
        _ = v29
    }
    var v30 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
