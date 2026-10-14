package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
func genFn0(p0 string) string {
    if false {
        p0 = p0
    } else {
        var v7 uint = 1
        v7 += v7
        fmt.Println(v7, v7)
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            switch (fi9 / (((fi9 % (((fi9 + 42)) | 1))) | 1)) {
            case 127:
                fmt.Println(fi8)
            }
            _ = fi9
        }
        var v10 int16 = -187
        _ = fi8
        _ = v10
    }
    if (1 != 12) {
        fmt.Println(p0)
        p0 = p0
        p0 = ("hello" + ("" + (p0 + p0)))
    }
    fmt.Println(p0)
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := []int{fi11}
        _ = fi11
        _ = v12
    }
    _ = p0
    return (((("foo bar" + p0) + "\n\t") + p0) + "foo bar")
}

func main() {
    var v1 bool = (!(true || false))
    v2 := make(map[string]string)
    var v3 uint8 = 56
    var v4 uint8 = v3
    var v5 float32 = -2.25
    var v6 int = -9223372036854775808
    var v13 uint = 13
    fmt.Println(v5)
    v14 := GenS1{F0: 1, F1: 0}
    fmt.Println(v5)
    var v15 float32 = float32((float64(v5) + 1e10))
    var v16 float32 = v5
    v1 = true
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
