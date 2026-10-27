package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 string; F1 int16 }
func genFn0(p0 uint) uint32 {
    if (("a" + "foo bar") == "foo bar") {
        if (0.0 < 3.14159) {
            p0 = p0
        } else {
            var v6 uint16 = 32767
            v7 := GenS1{F0: (v6 * (v6 | (^(v6))))}
            p0 = (p0 | (p0 ^ p0))
            _ = v6
            _ = v7
        }
    } else {
        if (695 > 261) {
            var v8 int64 = 9223372036854775807
            fmt.Println(p0, v8)
            v9 := GenS2{F0: 763}
            _ = v8
            _ = v9
        }
        var v10 float64 = 1e10
        var v11 uint8 = 19
        _ = v10
        _ = v11
    }
    var v12 int = -756
    switch ("a" + "") {
    case "":
        v13 := GenS2{F0: 474}
        _ = v13
    default:
        fmt.Println(p0)
    }
    v14 := GenS2{F0: 351}
    var v15 int8 = 100
    _ = p0
    _ = v12
    _ = v14
    _ = v15
    return 651
}

func main() {
    var v1 string = ""
    var v2 float32 = -2.25
    var v3 float64 = (1.5 * (0.0 * ((1.5 / float64(v2)) + ((0.1 / 1e10) - 0.0))))
    var v4 uint32 = 2147483647
    var v5 uint8 = 2
    var v16 int = 42
    var v17 bool = true
    v4 -= v4
    v18 := GenS3{F0: genFn0(1), F1: 0}
    var v19 uint8 = (v5 % ((142) | 1))
    v20 := []int{v16, 2147483647}
    var v21 int32 = -1
    var v22 float64 = ((v3 - ((((0.0 - v3) / (v3 * v3)) * v3) * float64(0.0))) / -2.25)
    for fi23 := 0; fi23 < 13; fi23++ {
        var v24 float64 = v3
        if v17 {
            var v25 float32 = v2
            _ = v25
        } else {
            fmt.Println(v22, v4)
            v26 := GenS1{F0: 65535}
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    v16 = (v16 / (((v16 ^ v16)) | 1))
    switch v1 {
    case "\n\t":
        fmt.Println(v1, v3)
    }
    if v17 {
        fmt.Println(v19)
    }
    v27 := GenS4{F0: ("a" + (v1 + ("hello" + ("foo bar" + v1)))), F1: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
