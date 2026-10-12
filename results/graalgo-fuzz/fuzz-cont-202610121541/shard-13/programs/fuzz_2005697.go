package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool }

func main() {
    v1 := []uint32{42, 65535}
    var v2 uint16 = 1
    var v3 float64 = 1e10
    var v4 string = ("\n\t" + (("a" + "foo bar") + "foo bar"))
    var v5 int32 = -457
    v6 := make(map[string]uint32)
    var v7 uint32 = 2147483647
    if true {
        fmt.Println(v5)
        v7 -= 904
    } else {
        v4 = v4
        v8 := GenS1{F0: 1}
        _ = v8
    }
    v7 -= (v7 / (((^(v7))) | 1))
    var v9 int64 = 42
    if (false && true) {
        v9 -= 42
        if ((^(v7)) <= (v7 ^ (v7 ^ (v7 & 65535)))) {
            fmt.Println(v5, v3, v2)
            var v10 uint32 = v7
            v11 := GenS2{F0: 32768, F1: v4}
            _ = v10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        if len(v1) > 0 {
            v1[0] = 2
        }
        var v13 uint = 335
        v14 := GenS3{F0: v5}
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 uint32 = (v7 + (v7 << ((65535) & 63)))
    if true {
        if true {
            delete(v6, v4)
            switch ("" + (v4 + "a")) {
            case "\n\t":
                if (true || false) {
                    v2 += (^(v2))
                    v16 := GenS4{F0: (true && ((!false) && true))}
                    var v17 uint8 = 169
                    _ = v16
                    _ = v17
                }
            case "foo bar":
                v5 = v5
            }
        }
    }
    var v18 float32 = 1e10
    switch (("a" + (((("foo bar" + v4) + v4) + v4) + (v4 + ((v4 + "a") + v4)))) + (v4 + v4)) {
    case "a":
        v9 -= -1
    case "\n\t":
        for fi19 := 0; fi19 < 14; fi19++ {
            v6[""] = 56
            for fi20 := 0; fi20 < 15; fi20++ {
                v4 = "hello"
                var v21 float32 = 3.14159
                var v22 int64 = v9
                fmt.Println(v15, v15, fi20)
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = fi19
        }
    default:
        v1 = append(v1, v7)
    }
    var v23 int8 = -128
    var v24 uint16 = v2
    fmt.Println(v5, v5, v4)
    if (("hello" != ("foo bar" + v4)) && (int32(v7) <= v5)) {
        var v25 int8 = (^(v23))
        v9 = int64(v5)
        _ = v25
    } else {
        if false {
            if false {
                fmt.Println(v15, v7)
            } else {
                v26 := GenS2{F0: 32767, F1: (((("a" + "foo bar") + v4) + ((v4 + (v4 + ((v4 + v4) + (("hello" + v4) + v4)))) + ("foo bar" + (v4 + "hello")))) + v4)}
                v27 := GenS2{F0: 32768, F1: (v4 + "a")}
                _ = v26
                _ = v27
            }
            fmt.Println(v4, v24)
        } else {
            fmt.Println(v24, v23)
            v1 = append(v1, v15)
        }
    }
    v7 += (v7 % ((v15) | 1))
    for fi28 := 0; fi28 < 16; fi28++ {
        for fi29 := 0; fi29 < 9; fi29++ {
            if (v23 <= -21) {
                var v30 float64 = ((v3 / float64((v18 - v18))) / v3)
                var v31 int32 = 570
                var v32 int16 = 1000
                var v33 int8 = (v23 >> ((0) & 63))
                _ = v30
                _ = v31
                _ = v32
                _ = v33
            } else {
                v5 = v5
                v34 := GenS2{F0: fi28, F1: (v4 + "a")}
                v18 += (v18 * (float32(0.0) - (v18 / v18)))
                _ = v34
            }
            v35 := GenS1{F0: v23}
            var v36 bool = true
            v4 = v4
            _ = fi29
            _ = v35
            _ = v36
        }
        _ = fi28
    }
    if (127 < 0) {
        fmt.Println(v4, v23, v3)
        v24 += 255
    } else {
        fmt.Println(v9)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
