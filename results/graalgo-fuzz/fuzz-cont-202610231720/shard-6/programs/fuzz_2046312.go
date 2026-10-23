package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
func genFn0(p0 uint64) bool {
    fmt.Println(p0)
    p0 += p0
    _ = p0
    return true
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = true
    var v3 int32 = 0
    v4 := []uint32{42, 2, 1}
    var v5 uint8 = 42
    var v6 uint16 = 42
    var v7 int = -798
    v7 -= int(v5)
    if v2 {
        switch ("\n\t" + ("" + ("" + (("" + "hello") + (((("" + "hello") + "foo bar") + "\n\t") + ("\n\t" + "foo bar")))))) {
        case "foo bar":
            var v8 float64 = 0.0
            _ = v8
        case "hello":
            switch "hello" {
            case "\n\t":
                var v9 int = (v7 << ((2) & 63))
                _ = v9
            default:
                if true {
                    var v10 uint16 = (v6 >> ((0) & 63))
                    _ = v10
                }
            }
        case "\n\t":
            v11 := GenS1{F0: (v3 ^ (v3 / ((-138) | 1))), F1: 18446744073709551615}
            _ = v11
        default:
            var v12 int16 = 333
            _ = v12
        }
    }
    var v13 int32 = v3
    v14 := GenS2{F0: (!(v2 && (65535 != 690)))}
    v7 -= v7
    v15 := GenS3{F0: ""}
    var v16 float64 = float64(1e10)
    var v17 float32 = 1.5
    delete(v1, ((("foo bar" + ("\n\t" + "a")) + ("" + "")) + ""))
    v18 := GenS2{F0: (!(int32(v6) == v3))}
    v7 += v7
    v17 = v17
    v19 := v1[(("\n\t" + "") + "\n\t")]
    v3 = (v13 - v13)
    var v20 int32 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
