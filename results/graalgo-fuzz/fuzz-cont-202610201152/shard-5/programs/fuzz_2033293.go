package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int32; F2 uint16 }

func main() {
    v1 := []int{-1, -651, 9223372036854775807}
    var v2 uint8 = 0
    var v3 float32 = 1.5
    v4 := make(map[string]string)
    var v5 uint32 = 1
    var v6 uint = 0
    var v7 int32 = -1
    switch 0 {
    case -2147483648:
        v6 = v6
    case -194:
        if true {
            if ((^(v7)) != (v7 + 565)) {
                if len(v1) > 0 {
                    v1[0] = -482
                }
                v7 = v7
                if ((v7 <= -253) && true) {
                    var v8 uint32 = 2147483647
                    _ = v8
                }
            }
            _ = len(v1)
            var v9 uint = 943
            _ = v9
        } else {
            fmt.Println(v7)
            var v10 uint32 = (v5 ^ v5)
            _ = v10
        }
    default:
        v1 = append(v1, -731)
    }
    v2 = (v2 % ((92) | 1))
    fmt.Println(v5, v3)
    if true {
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 string = "a"
            v7 = int32(v2)
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: 1}
        v3 -= (0.1 - ((v3 * ((float32(1e10) + v3) - float32(-2.25))) - (v3 / (1.5 - 0.1))))
        _ = v13
    } else {
        var v14 bool = false
        if true {
            if (v14 && (("" >= ("\n\t" + "")) && v14)) {
                var v15 uint32 = v5
                for fi16 := 0; fi16 < 18; fi16++ {
                    var v17 float64 = 3.14159
                    v18 := GenS2{F0: (v7 | (v7 ^ (v7 * -841)))}
                    var v19 int32 = (v7 >> (((^(v6))) & 63))
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
                _ = v15
            }
            v14 = ("foo bar" >= "hello")
        } else {
            var v20 float64 = float64(float32((1.5 / -2.25)))
            _ = v20
        }
        v21 := GenS3{F0: -9223372036854775808, F1: int32(v5), F2: 65535}
        _ = v14
        _ = v21
    }
    fmt.Println(v3, v7)
    var v22 int64 = 0
    v22 -= -323
    v4["a"] = ("\n\t" + "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
