package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint64; F2 string }

func main() {
    var v1 uint16 = 0
    var v2 uint64 = 2
    var v3 int8 = -1
    v4 := make(map[string]string)
    v5 := []int{778}
    v6 := GenS1{F0: ("foo bar" + (("a" + "a") + (("" + ("\n\t" + ("a" + ("foo bar" + "foo bar")))) + "hello")))}
    fmt.Println(v2, v2, v1)
    var v7 uint32 = 42
    for fi8 := 0; fi8 < 19; fi8++ {
        fmt.Println(v2, fi8, v3)
        switch "a" {
        case "hello":
            if true {
                v1 -= v1
                var v9 int64 = int64(v3)
                _ = v9
            }
        case "\n\t":
            if true {
                v1 = 1
            } else {
                fmt.Println(v3, v1, v3)
                for fi10 := 0; fi10 < 20; fi10++ {
                    v2 += 1
                    var v11 float64 = (3.14159 - (0.0 - float64((0.0 + (((3.14159 / 0.1) * (1e10 + ((0.0 - 1e10) + 1.5))) / ((3.14159 - ((((float32(0.1) * (float32(0.0) + float32((3.14159 + 0.0)))) - -2.25) + 1.5) - 1e10)) + -2.25))))))
                    var v12 uint64 = uint64(v7)
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                var v13 int32 = 42
                _ = v13
            }
        default:
            if len(v5) > 0 {
                v5[0] = 32767
            }
        }
        _ = fi8
    }
    var v14 float32 = float32(0.1)
    var v15 int8 = (-(v3))
    fmt.Println(v15, v3)
    v1 = v1
    fmt.Println(v3, v3, v2)
    v15 -= v15
    if true {
        v16 := GenS2{F0: v15, F1: 4294967295, F2: "foo bar"}
        v14 += v14
        if false {
            v3 += v15
        } else {
            v5 = append(v5, 4294967295)
        }
        _ = v16
    }
    var v17 bool = true
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
