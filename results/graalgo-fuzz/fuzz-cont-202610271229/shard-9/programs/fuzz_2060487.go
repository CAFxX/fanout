package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 bool }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint64, p1 uint8, p2 string) uint32 {
    v4 := GenS1{F0: 42, F1: "", F2: false}
    p2 = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1(p0 int32, p1 int, p2 uint8) uint16 {
    p1 -= (p1 << ((42) & 63))
    var v5 int16 = int16(p0)
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int = (fi6 - p1)
        v8 := []int{(^(v7)), fi6, (fi6 << ((65535) & 63))}
        switch "" {
        case "\n\t":
            var v9 uint8 = (p2 >> ((4294967295) & 63))
            _ = v9
        case "hello":
            switch "\n\t" {
            case "":
                v10 := GenS1{F0: 52, F1: "foo bar", F2: (902 >= 4294967295)}
                _ = v10
            case "\n\t":
                p2 += 54
            case "foo bar":
                fmt.Println(p0, fi6, p2)
            default:
                p2 = (^(p2))
            }
        case "":
            _ = len(v8)
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v11 := []int{(p1 / ((-1) | 1)), (p1 * p1)}
    var v12 int = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v5
    return 760
}

func main() {
    var v1 uint8 = 128
    var v2 int32 = -1
    v3 := make(map[string]bool)
    v13 := GenS2{F0: 2}
    v3[("" + ((("\n\t" + "foo bar") + ((("" + ("foo bar" + ("\n\t" + "a"))) + "") + ((("" + "foo bar") + "") + "foo bar"))) + (("\n\t" + "a") + "a")))] = false
    var v14 bool = ((true || (((((("hello" + "a") + "foo bar") < (("\n\t" + "") + "")) || (false || false)) || false) || (!(!(-109 < -97))))) || false)
    var v15 int64 = 770
    if true {
        v2 = v2
        fmt.Println(v15, v1)
        switch "\n\t" {
        case "\n\t":
            v2 -= (-(v2))
        default:
            v15 = v15
        }
    } else {
        var v16 uint32 = 2
        _ = v16
    }
    if (v14 && true) {
        var v17 bool = (-2147483648 >= 32767)
        _ = v17
    } else {
        v18 := v3[(("hello" + (("hello" + ("\n\t" + ("\n\t" + "a"))) + "foo bar")) + (("" + "a") + "\n\t"))]
        if v18 {
            var v19 bool = v18
            var v20 int = 127
            _ = v19
            _ = v20
        } else {
            fmt.Println(v15, v18)
            for fi21 := 0; fi21 < 2; fi21++ {
                fmt.Println(v18, v15)
                var v22 float32 = (0.0 - (0.0 - (0.1 - 1e10)))
                _ = fi21
                _ = v22
            }
        }
        v18 = v14
        _ = v18
    }
    v23 := GenS2{F0: (v1 | v1)}
    fmt.Println(v2, v14)
    var v24 int = -9223372036854775808
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 uint16 = 416
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
