package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int16; F1 int32; F2 uint }
type GenS3 struct { F0 uint16 }
func genFn0() int8 {
    var v5 uint8 = 1
    var v6 int8 = 127
    v6 += v6
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint32 = 845
        _ = fi7
        _ = v8
    }
    _ = v5
    _ = v6
    return int8(v5)
}
func genFn1(p0 int32, p1 uint64, p2 uint32) uint64 {
    p1 = (p1 & 0)
    switch ("" + "a") {
    case "foo bar":
        switch 32768 {
        case 128:
            fmt.Println(p0, p0, p1)
        }
    case "a":
        p1 += (^(p1))
    default:
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    return (^(p1))
}

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]int)
    v3 := []int{}
    var v4 uint8 = 255
    if (("\n\t" + (("\n\t" + ("" + (("\n\t" + "\n\t") + (((("a" + "") + "\n\t") + ((("foo bar" + "foo bar") + "foo bar") + "a")) + "")))) + "hello")) != (("" + ("hello" + "a")) + (("\n\t" + "foo bar") + "a"))) {
        if false {
            for fi9 := 0; fi9 < 2; fi9++ {
                v1 -= -1
                _ = fi9
            }
            var v10 int = 448
            var v11 float32 = ((0.1 + 0.0) / 1.0)
            _ = v10
            _ = v11
        }
        if true {
            _ = len(v2)
        } else {
            fmt.Println(v1)
            var v12 uint16 = 32767
            fmt.Println(v4, v12)
            _ = v12
        }
    } else {
        v13 := GenS1{F0: v4, F1: 314}
        v1 += 100
        _ = v13
    }
    if (!true) {
        switch ("a" + "a") {
        case "foo bar":
            v14 := GenS2{F0: 42, F1: 0, F2: 42}
            _ = v14
        case "a":
            fmt.Println(v1, v4)
        case "hello":
            v15 := v2[2147483647]
            _ = v15
        }
        _ = len(v2)
        v16 := GenS3{F0: 255}
        _ = v16
    }
    fmt.Println(v4)
    v4 += v4
    v1 = genFn0()
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        delete(v2, 2147483647)
    }
    if false {
        v17 := GenS3{F0: 965}
        _ = v17
    } else {
        switch 836 {
        case 1:
            var v18 int64 = -9223372036854775808
            _ = v18
        case -9223372036854775808:
            v19 := GenS3{F0: 65535}
            _ = v19
        default:
            var v20 uint16 = 255
            _ = v20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
