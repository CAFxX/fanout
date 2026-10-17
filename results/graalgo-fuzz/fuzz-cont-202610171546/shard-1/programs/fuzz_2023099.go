package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := []int{9223372036854775807, 4294967295, 94}
    var v2 int64 = -9223372036854775808
    v3 := make(map[string]int)
    var v4 string = "hello"
    var v5 uint = 147
    var v6 int8 = 42
    if ((!("" > (v4 + v4))) || true) {
        v2 += (v2 | (v2 * (v2 - (v2 / (((v2 % ((v2) | 1))) | 1)))))
        v2 = v2
    }
    switch v4 {
    case "":
        v6 = -1
    }
    if false {
        v7 := GenS1{F0: 32767}
        _ = v7
    }
    switch v4 {
    case "hello":
        var v8 int32 = -740
        _ = v8
    case "a":
        var v9 uint64 = 143
        _ = v9
    default:
        delete(v3, v4)
    }
    v1 = append(v1, 4294967295)
    v10 := GenS1{F0: 1000}
    v6 -= 100
    var v11 int32 = 0
    for fi12 := 0; fi12 < 6; fi12++ {
        v5 += v5
        var v13 int32 = v11
        _ = fi12
        _ = v13
    }
    if len(v1) > 0 {
        v1[0] = 32768
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        if false {
            fmt.Println(v2)
            var v15 bool = false
            var v16 uint = (v5 % (((v5 | 703)) | 1))
            _ = v15
            _ = v16
        } else {
            v4 = "a"
            v5 += (v5 + (v5 * (v5 & v5)))
        }
        _ = fi14
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
