package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int16; F1 int16; F2 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 bool; F2 uint8 }
type GenS7 struct { F0 int8 }
func genFn0() int64 {
    var v6 int32 = -2147483648
    var v7 uint32 = 2147483648
    if (!true) {
        var v8 int32 = 42
        v9 := []int{}
        v7 = (v7 << ((18446744073709551615) & 63))
        _ = v8
        _ = v9
    }
    var v10 int = -571
    v7 = 0
    _ = v10
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint64) uint32 {
    v11 := GenS1{F0: "hello", F1: uint16(p0), F2: 2147483647}
    p0 = (p0 / ((p0) | 1))
    for fi12 := 0; fi12 < 5; fi12++ {
        fmt.Println(fi12)
        _ = fi12
    }
    v13 := GenS2{F0: -65, F1: 751}
    v14 := []int{-2147483648}
    _ = p0
    _ = v11
    _ = v13
    _ = v14
    return 608
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float32 = ((((float32(0.1) + ((0.0 * 3.14159) - -2.25)) - 0.1) - ((float32(float64(3.14159)) + float32(0.1)) * ((((0.1 * ((0.0 / float32(0.1)) / -2.25)) - 0.0) + (0.1 - -2.25)) + -2.25))) / (-2.25 / 1.0))
    v3 := make(map[string]int)
    var v4 int = 4
    var v5 float32 = (float32(1e10) * ((0.1 - (1.5 * v2)) + ((3.14159 + v2) - v2)))
    v15 := []int{9223372036854775807, v4}
    switch (v4 >> ((42) & 63)) {
    case -2147483648:
        var v16 int16 = int16(v4)
        _ = v16
    case 4294967295:
        fmt.Println(v2)
    case 902:
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 string = ("" + "foo bar")
            fmt.Println(v4)
            delete(v3, (v18 + v18))
            _ = fi17
            _ = v18
        }
    }
    fmt.Println(v1, v5)
    v19 := GenS3{F0: 2, F1: 2, F2: 4294967295}
    v20 := GenS4{F0: -35, F1: 0, F2: 1}
    if false {
        fmt.Println(v4, v2)
        if len(v15) > 0 {
            v15[0] = v4
        }
    }
    v5 = 1.5
    v4 = (^(v4))
    v5 += (1e10 + -2.25)
    for fi21 := 0; fi21 < 12; fi21++ {
        if (false || true) {
            v2 -= v5
            delete(v3, ((((("a" + "\n\t") + "hello") + "\n\t") + ((("foo bar" + "") + ("foo bar" + "")) + "hello")) + "foo bar"))
            v5 = -2.25
        } else {
            v1 = 18446744073709551615
            v4 -= (fi21 | 2147483647)
        }
        if false {
            v22 := GenS5{F0: genFn0()}
            fmt.Println(v1)
            fmt.Println(fi21, v2)
            _ = v22
        } else {
            switch v4 {
            case -363:
                fmt.Println(v1)
            default:
                for fi23 := 0; fi23 < 4; fi23++ {
                    if false {
                        v3[("" + "\n\t")] = (fi21 >> ((581) & 63))
                        v4 = (fi23 + (^(fi21)))
                    }
                    var v24 uint = 918
                    _ = fi23
                    _ = v24
                }
            }
            for fi25 := 0; fi25 < 2; fi25++ {
                var v26 string = ""
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                if true {
                    switch (v4 >> ((255) & 63)) {
                    case 812:
                        v27 := GenS1{F0: v26, F1: 42, F2: genFn1((v1 * 2))}
                        _ = v27
                    case 647:
                        delete(v3, ((("hello" + v26) + ("\n\t" + "a")) + "a"))
                    default:
                        v2 = v2
                    }
                    v4 = v4
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    v28 := GenS6{F0: 1, F1: ("a" != v26), F2: 128}
                    _ = v28
                }
                _ = fi25
                _ = v26
            }
        }
        _ = fi21
    }
    v29 := v3[""]
    fmt.Println(v2, v2)
    var v30 int16 = -335
    for fi31 := 0; fi31 < 19; fi31++ {
        v32 := GenS7{F0: -128}
        delete(v3, "")
        for fi33 := 0; fi33 < 3; fi33++ {
            v4 += v29
            _ = fi33
        }
        _ = fi31
        _ = v32
    }
    if (v5 >= (v2 + v5)) {
        var v34 uint64 = (v1 ^ v1)
        _ = v34
    } else {
        switch (v4 & v4) {
        case 0:
            var v35 int32 = 2147483647
            _ = v35
        }
    }
    v4 += 128
    fmt.Println(v5, v1)
    var v36 uint8 = 239
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v5
}
