package main

import (
    "fmt"
)


func main() {
    var v1 uint64 = 0
    var v2 uint64 = v1
    var v3 int64 = -9223372036854775808
    v4 := []string{"foo bar", (("" + "\n\t") + ("foo bar" + "")), "a"}
    for fi5 := 0; fi5 < 1; fi5++ {
        v2 = v1
        for fi6 := 0; fi6 < 8; fi6++ {
            fmt.Println(fi6, v1)
            switch ("foo bar" + "a") {
            case "\n\t":
                v3 -= v3
            case "":
                var v7 string = "hello"
                _ = len(v4)
                _ = v7
            }
            _ = fi6
        }
        _ = fi5
    }
    var v8 float32 = ((1e10 + 3.14159) * 1e10)
    switch -1 {
    case -285:
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 uint32 = 937
            _ = fi9
            _ = v10
        }
    }
    v8 = (((v8 * ((1e10 / (v8 + float32(float64((((0.1 - 1.5) * -2.25) / (v8 / 1.0)))))) * v8)) / v8) * v8)
    var v11 int16 = -818
    var v12 int = 128
    v2 += 18446744073709551615
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 float32 = ((v8 + 3.14159) / v8)
        fmt.Println(v2, v14)
        v1 += v2
        _ = fi13
        _ = v14
    }
    var v15 uint64 = 42
    v3 -= (^(v3))
    fmt.Println(v11, v15)
    v15 = 1
    var v16 uint = 42
    v2 += (v15 % (((^(v2))) | 1))
    var v17 uint64 = 4294967295
    fmt.Println(v8)
    var v18 int = v12
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
