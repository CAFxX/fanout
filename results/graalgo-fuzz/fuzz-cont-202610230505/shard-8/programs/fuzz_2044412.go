package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int64; F2 uint16 }
func genFn0() int {
    var v6 uint16 = 65535
    v6 = v6
    fmt.Println(v6)
    _ = v6
    return 32768
}
func genFn1(p0 uint, p1 uint32) bool {
    fmt.Println(p1, p1)
    var v7 uint64 = 1
    p0 = 340
    var v8 float64 = float64((1e10 + (float32(-2.25) / 0.1)))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return true
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{"foo bar", "hello", (((((("a" + "hello") + "") + ("" + (("" + "") + "\n\t"))) + ("foo bar" + "hello")) + ("\n\t" + "hello")) + ("hello" + "hello"))}
    var v3 float64 = -2.25
    var v4 uint64 = 104
    v5 := make(map[string]int)
    switch 4294967295 {
    case -2147483648:
        if false {
            var v9 int32 = 36
            switch "hello" {
            case "foo bar":
                fmt.Println(v3, v4)
            default:
                switch "foo bar" {
                case "hello":
                    var v10 uint64 = 951
                    _ = v10
                case "foo bar":
                    delete(v5, "hello")
                case "":
                    var v11 uint32 = 0
                    _ = v11
                default:
                    var v12 int16 = 32767
                    _ = v12
                }
            }
            _ = v9
        } else {
            if ((!(-921 >= -1)) || false) {
                fmt.Println(v3, v3)
                var v13 int16 = -695
                var v14 uint32 = 2147483647
                _ = v13
                _ = v14
            } else {
                v15 := v5["hello"]
                _ = v15
            }
        }
    case 127:
        v1[(("\n\t" + "\n\t") + "foo bar")] = genFn0()
    case 0:
        v16 := GenS1{F0: 1000, F1: 32767, F2: ""}
        _ = v16
    default:
        fmt.Println(v3)
    }
    if false {
        fmt.Println(v4, v4)
        var v17 uint32 = 0
        _ = v17
    }
    var v18 int = 32767
    v4 -= 18446744073709551615
    var v19 float32 = float32(0.0)
    if (!((32767 > 65535) || ((v18 - (v18 + v18)) == 358))) {
        v19 = 1.5
        if false {
            v18 += v18
            v20 := GenS2{F0: 42}
            v2 = append(v2, "foo bar")
            _ = v20
        }
        var v21 uint8 = 2
        _ = v21
    } else {
        fmt.Println(v3, v19, v18)
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        v23 := GenS1{F0: 0, F1: 2, F2: "foo bar"}
        _ = fi22
        _ = v23
    }
    v3 = (1.5 - v3)
    v24 := GenS3{F0: 973, F1: 928, F2: 32767}
    v5["foo bar"] = (v18 ^ v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
