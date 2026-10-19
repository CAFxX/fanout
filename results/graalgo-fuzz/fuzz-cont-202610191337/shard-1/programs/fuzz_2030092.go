package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint32; F1 int8; F2 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 uint32 }
func genFn0(p0 int32, p1 bool) int32 {
    v5 := []int{32768, -1, 197}
    var v6 uint = 0
    p0 = (p0 * (p0 + (p0 << ((28) & 63))))
    for fi7 := 0; fi7 < 12; fi7++ {
        if false {
            if len(v5) > 0 {
                v5[0] = (-(fi7))
            }
            v5 = append(v5, (fi7 - 4294967295))
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (-(p0))
}
func genFn1() bool {
    var v8 uint16 = 255
    var v9 uint = 0
    v10 := []int{}
    if false {
        fmt.Println(v8, v8)
        switch 128 {
        case -1:
            v11 := GenS1{F0: 42, F1: 16}
            _ = v11
        }
    }
    v8 -= (v8 ^ (v8 % ((v8) | 1)))
    _ = v10
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 int64 = 329
    v2 := []uint32{0, 21, 1}
    var v3 uint8 = 128
    var v4 int16 = 0
    if true {
        v4 = v4
        var v12 int16 = (v4 ^ v4)
        _ = v12
    }
    var v13 float32 = 0.0
    v14 := GenS2{F0: 255, F1: 1, F2: v4}
    v3 = (v3 + (v3 << ((4294967295) & 63)))
    v1 = 0
    if (!(v3 >= v3)) {
        var v15 int8 = -1
        v15 += (v15 & (-(v15)))
        _ = v15
    }
    switch ((("" + ("" + "hello")) + ("a" + ("hello" + "a"))) + ("" + "foo bar")) {
    case "hello":
        var v16 string = "\n\t"
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        switch (fi17 - (fi17 & 128)) {
        case 1:
            var v18 uint32 = 777
            _ = v18
        case -1:
            v13 = float32(-2.25)
        case -229:
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
        }
        _ = fi17
    }
    v19 := GenS3{F0: 211, F1: 6, F2: (v1 << ((367) & 63))}
    for fi20 := 0; fi20 < 11; fi20++ {
        var v21 int32 = 40
        var v22 uint64 = 15
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := GenS4{F0: 100}
            var v25 int16 = -1
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi20
        _ = v21
        _ = v22
    }
    v26 := GenS5{F0: -2147483648, F1: 980}
    v3 = v3
    var v27 int8 = 127
    v2 = append(v2, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
