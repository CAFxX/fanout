package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }
func genFn0(p0 int64, p1 int16, p2 uint) int8 {
    var v6 int32 = 0
    p2 = p2
    var v7 int16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 127
}

func main() {
    var v1 float32 = -2.25
    v2 := []int{128, 9223372036854775807}
    v3 := []string{("a" + ""), "foo bar", "foo bar"}
    var v4 float32 = 0.0
    var v5 int32 = 42
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 bool = false
        if (115 == 79) {
            v10 := GenS1{F0: 0}
            _ = v10
        }
        switch (("\n\t" + "") + "foo bar") {
        case "\n\t":
            v11 := GenS1{F0: -1}
            _ = v11
        default:
            for fi12 := 0; fi12 < 13; fi12++ {
                fmt.Println(fi8, v4)
                v2 = append(v2, fi12)
                _ = fi12
            }
        }
        _ = fi8
        _ = v9
    }
    v5 -= (v5 / ((v5) | 1))
    if len(v2) > 0 {
        v2[0] = 9223372036854775807
    }
    fmt.Println(v5, v5)
    v1 = 1e10
    v5 += 1
    var v13 int32 = v5
    var v14 uint32 = 2147483647
    fmt.Println(v4, v4)
    switch "\n\t" {
    case "a":
        if (v13 >= v13) {
            var v15 string = ""
            _ = v15
        }
    default:
        var v16 int32 = (v13 ^ (v5 << ((823) & 63)))
        _ = v16
    }
    v17 := GenS2{F0: true}
    var v18 uint64 = uint64(v14)
    var v19 int16 = 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v20 string = ""
    if false {
        var v21 uint = 18446744073709551615
        v22 := GenS3{F0: -128}
        _ = v21
        _ = v22
    }
    var v23 string = (v20 + (((v20 + v20) + ("hello" + (v20 + v20))) + v20))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
