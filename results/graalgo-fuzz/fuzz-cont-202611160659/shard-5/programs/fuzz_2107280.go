package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int16 }
func genFn0(p0 int16, p1 int8, p2 uint) int16 {
    p0 -= 42
    var v6 float32 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p0
}
func genFn1() string {
    var v7 float32 = float32((-2.25 / 3.14159))
    var v8 uint = 65535
    _ = v7
    _ = v8
    return (("\n\t" + "hello") + "\n\t")
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 2
    var v3 int16 = 42
    v4 := []uint32{42, 2147483647, 141, 2147483647}
    v5 := make(map[string]int)
    v9 := GenS1{F0: genFn0(v3, 100, 42), F1: 4294967295}
    if false {
        _ = len(v1)
    }
    v10 := GenS1{F0: (v3 << ((4294967295) & 63)), F1: 2}
    _ = len(v4)
    if ((("foo bar" + "a") < "") && (false && false)) {
        _ = len(v1)
    }
    var v11 uint8 = 167
    v12 := GenS2{F0: 1}
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 int32 = -2147483648
        var v15 string = ("a" + (("foo bar" + "foo bar") + ((((("hello" + "foo bar") + ("\n\t" + "a")) + "") + ("foo bar" + "a")) + ("hello" + ""))))
        _ = fi13
        _ = v14
        _ = v15
    }
    if (65535 > 561) {
        var v16 int32 = 243
        _ = v16
    } else {
        v2 += v2
        if true {
            v4 = append(v4, 163)
            var v17 int8 = -125
            var v18 uint16 = 32767
            _ = v17
            _ = v18
        } else {
            v19 := GenS3{F0: 127, F1: (-(v3))}
            v3 = (v3 >> ((1) & 63))
            _ = v19
        }
        v3 += v3
    }
    v2 = v2
    switch 2147483647 {
    case 349:
        var v20 int64 = -56
        _ = v20
    case -9223372036854775808:
        v11 -= (^(v11))
    case 127:
        v21 := GenS3{F0: 100, F1: v3}
        _ = v21
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
