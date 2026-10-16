package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int16; F1 uint8; F2 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32; F1 string }
func genFn0(p0 bool) string {
    if p0 {
        var v5 uint16 = 2
        _ = v5
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(fi6)
        p0 = p0
        p0 = (42 < 42)
        _ = fi6
    }
    v7 := []int{32768, 32768, 201, 1}
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: -342, F1: true}
        if p0 {
            fmt.Println(p0, fi8)
            var v10 uint64 = 4294967295
            p0 = (0 == 42)
            _ = v10
        }
        v7 = append(v7, fi8)
        _ = fi8
        _ = v9
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 10; fi12++ {
            fmt.Println(fi11, fi12, fi11)
            var v13 uint = 1
            p0 = p0
            var v14 float32 = (1.5 - (float32(((0.0 - 1.5) - 1.5)) * 3.14159))
            _ = fi12
            _ = v13
            _ = v14
        }
        if ((-2147483648 >= -368) || (310 < 1)) {
            var v15 int16 = 1000
            v16 := GenS2{F0: 100}
            for fi17 := 0; fi17 < 13; fi17++ {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                if true {
                    switch (("" + ("" + "foo bar")) + ("hello" + ("foo bar" + ("" + "\n\t")))) {
                    case "foo bar":
                        var v18 uint = 42
                        _ = v18
                    case "\n\t":
                        var v19 int16 = v15
                        _ = v19
                    default:
                        v15 -= v15
                    }
                } else {
                    switch 633 {
                    case 42:
                        var v20 uint32 = 1
                        _ = v20
                    case 32768:
                        v15 = (v15 >> ((65535) & 63))
                    case 9223372036854775807:
                        fmt.Println(fi17, fi11, v15)
                    default:
                        p0 = true
                    }
                    v15 -= (v15 << ((662) & 63))
                }
                p0 = p0
                if (65535 != 4294967295) {
                    switch "foo bar" {
                    case "hello":
                        var v21 int64 = -1
                        _ = v21
                    case "":
                        if len(v7) > 0 {
                            v7[0] = fi17
                        }
                    }
                    _ = len(v7)
                }
                p0 = p0
                v22 := GenS3{F0: "foo bar", F1: 0}
                if (false && p0) {
                    p0 = p0
                } else {
                    var v23 float32 = 0.0
                    v23 += 1e10
                    _ = v23
                }
                v24 := GenS4{F0: 0, F1: 255, F2: (768 > 4294967295)}
                _ = fi17
                _ = v22
                _ = v24
            }
            _ = v15
            _ = v16
        }
        p0 = p0
        _ = fi11
    }
    _ = p0
    _ = v7
    return "hello"
}
func genFn1() int64 {
    var v25 int8 = 93
    fmt.Println(v25)
    fmt.Println(v25)
    var v26 int64 = 0
    _ = v25
    _ = v26
    return 0
}

func main() {
    v1 := []string{}
    v2 := make(map[int]string)
    v3 := []uint32{572, 4294967295, 27, 1}
    var v4 bool = false
    switch "\n\t" {
    case "a":
        v27 := v2[32767]
        _ = v27
    case "":
        var v28 int32 = -219
        _ = v28
    }
    if v4 {
        v4 = (((-128 == 1) && v4) || (float32(-2.25) != 0.1))
    }
    var v29 int64 = -463
    v29 += (^(v29))
    var v30 uint32 = 65535
    _ = len(v2)
    if v4 {
        if v4 {
            v29 -= v29
        }
    }
    if v4 {
        switch 1 {
        case 2147483647:
            var v31 uint = 42
            _ = v31
        case 32768:
            v32 := GenS5{F0: 80}
            _ = v32
        case -914:
            for fi33 := 0; fi33 < 7; fi33++ {
                v34 := GenS6{F0: 42, F1: "a"}
                _ = fi33
                _ = v34
            }
        }
    } else {
        var v35 int8 = -24
        _ = v35
    }
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
