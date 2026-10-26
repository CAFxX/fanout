package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int64; F1 int8; F2 bool }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int8; F1 string; F2 uint }
func genFn0(p0 int32, p1 uint) string {
    var v4 string = "hello"
    v5 := []int{0}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return v4
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]string)
    var v3 uint32 = 224
    if (((v3 - v3) >= v3) || true) {
        v3 -= (v3 + v3)
    } else {
        v3 = 2147483647
        v6 := GenS1{F0: "a", F1: 255}
        for fi7 := 0; fi7 < 9; fi7++ {
            switch "\n\t" {
            case "a":
                fmt.Println(v3, fi7, fi7)
            case "hello":
                v8 := GenS2{F0: 0, F1: -86, F2: true}
                _ = v8
            default:
                v1 = -2.25
            }
            v9 := GenS3{F0: 1, F1: "\n\t"}
            _ = fi7
            _ = v9
        }
        _ = v6
    }
    var v10 int32 = -308
    delete(v2, (((("foo bar" + ("\n\t" + "hello")) + "foo bar") + ("foo bar" + "a")) + "foo bar"))
    v3 -= (v3 / (((v3 ^ (v3 ^ 591))) | 1))
    var v11 uint = 39
    var v12 uint32 = 65535
    var v13 int16 = 42
    v14 := GenS3{F0: (v10 << ((v11) & 63)), F1: ("foo bar" + "foo bar")}
    if false {
        var v15 int8 = 1
        v15 = -96
        _ = v15
    } else {
        var v16 string = "hello"
        v10 += (v10 ^ (v10 & 0))
        if (!(!false)) {
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 uint64 = 1
                _ = fi17
                _ = v18
            }
        } else {
            v19 := GenS4{F0: 42, F1: v16, F2: (^(v11))}
            fmt.Println(v3, v16, v10)
            v16 = v16
            _ = v19
        }
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
