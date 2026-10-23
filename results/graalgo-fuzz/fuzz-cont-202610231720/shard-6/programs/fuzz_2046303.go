package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 string }
func genFn0(p0 int32, p1 int32) string {
    v5 := []int{127, 32767}
    var v6 uint32 = 605
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return "foo bar"
}
func genFn1(p0 int) string {
    v7 := []int{p0}
    var v8 int = (p0 / ((p0) | 1))
    p0 += v8
    _ = p0
    _ = v7
    _ = v8
    return ("a" + "hello")
}

func main() {
    var v1 bool = false
    var v2 int32 = 42
    var v3 float32 = ((0.1 + 3.14159) - 1e10)
    var v4 int32 = -975
    var v9 bool = v1
    v3 = 0.0
    v10 := GenS1{F0: 2147483647}
    var v11 int64 = -40
    v12 := []int{}
    fmt.Println(v4, v1)
    v13 := GenS1{F0: -299}
    var v14 bool = v1
    for fi15 := 0; fi15 < 1; fi15++ {
        v16 := GenS2{F0: 42}
        v3 = (v3 + v3)
        if (((v4 >> ((255) & 63)) < (v2 / ((v4) | 1))) && v9) {
            fmt.Println(v11)
            var v17 bool = (0 <= -128)
            _ = v17
        } else {
            if false {
                var v18 string = (("a" + "") + "\n\t")
                fmt.Println(v2, v14)
                if (((((v3 / -2.25) - v3) * (0.0 * ((0.1 / v3) * v3))) >= v3) || (v11 == (^(v11)))) {
                    var v19 uint16 = 65535
                    var v20 int = (fi15 + fi15)
                    _ = v19
                    _ = v20
                } else {
                    if (v14 || ((fi15 << ((4294967295) & 63)) <= 128)) {
                        v12 = append(v12, -9223372036854775808)
                    }
                    v21 := GenS2{F0: 32}
                    _ = v21
                }
                _ = v18
            }
            fmt.Println(v1, v14, v2)
        }
        _ = fi15
        _ = v16
    }
    var v22 int64 = 382
    for fi23 := 0; fi23 < 5; fi23++ {
        var v24 string = "hello"
        if (1 == -128) {
            switch v24 {
            case "":
                v12 = append(v12, fi23)
            default:
                var v25 int8 = 127
                _ = v25
            }
        }
        _ = fi23
        _ = v24
    }
    fmt.Println(v4, v3)
    var v26 uint16 = 42
    if (42 == int(v26)) {
        v4 -= (v4 & (-(v2)))
    } else {
        v27 := GenS3{F0: (v11 ^ v22), F1: ""}
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v9
}
