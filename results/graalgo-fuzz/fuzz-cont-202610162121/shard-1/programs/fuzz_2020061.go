package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int64, p1 uint32, p2 uint8) string {
    var v5 uint = 994
    var v6 uint32 = p1
    v7 := GenS1{F0: 1, F1: -9223372036854775808}
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return "a"
}

func main() {
    v1 := []int{-601}
    var v2 int16 = 116
    var v3 float32 = 1e10
    var v4 uint16 = 0
    var v9 uint = 1
    var v10 int32 = -2147483648
    _ = len(v1)
    v4 -= 65535
    var v11 uint16 = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = -1
    }
    v9 += v9
    var v12 uint = 42
    if true {
        var v13 uint = 315
        switch "\n\t" {
        case "foo bar":
            var v14 string = (((("a" + ("a" + ("\n\t" + (((("hello" + "hello") + ("foo bar" + "hello")) + "hello") + "a")))) + "hello") + ("" + "\n\t")) + ("\n\t" + "hello"))
            _ = v14
        case "hello":
            var v15 string = (("hello" + "") + ("hello" + ("" + "hello")))
            _ = v15
        }
        v10 += v10
        _ = v13
    } else {
        if true {
            var v16 int = 1
            _ = v16
        } else {
            var v17 uint32 = 499
            var v18 int64 = -115
            fmt.Println(v2, v10)
            _ = v17
            _ = v18
        }
    }
    if true {
        var v19 float64 = 0.0
        fmt.Println(v12, v10, v11)
        _ = len(v1)
        _ = v19
    } else {
        v3 += 3.14159
        v20 := GenS1{F0: (v10 >> ((uint(v11)) & 63)), F1: -9223372036854775808}
        if true {
            for fi21 := 0; fi21 < 20; fi21++ {
                for fi22 := 0; fi22 < 4; fi22++ {
                    v23 := GenS2{F0: (v10 % (((v10 + v10)) | 1))}
                    _ = fi22
                    _ = v23
                }
                v2 = 32767
                _ = fi21
            }
            v11 -= v11
        } else {
            v3 = (0.0 * (1e10 - (1e10 * v3)))
        }
        _ = v20
    }
    v1 = append(v1, 603)
    var v24 bool = false
    var v25 int64 = -922
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v9
}
