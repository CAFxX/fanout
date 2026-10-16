package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 bool; F1 string }

func main() {
    var v1 string = ""
    v2 := []uint32{2147483647, 65535}
    var v3 int32 = 42
    var v4 int32 = -64
    v5 := []uint32{396}
    v3 += v4
    v1 = ("" + ("hello" + v1))
    if (3.14159 >= 1e10) {
        v4 = (^(v3))
    } else {
        var v6 string = ("foo bar" + (v1 + "hello"))
        v3 = -198
        if false {
            fmt.Println(v4)
        } else {
            var v7 uint8 = 128
            _ = v7
        }
        _ = v6
    }
    switch v1 {
    case "":
        var v8 int32 = v3
        _ = v8
    case "\n\t":
        if (!true) {
            v1 = ("hello" + "a")
            var v9 int64 = -9223372036854775808
            v3 = (v4 | v3)
            _ = v9
        }
    case "hello":
        var v10 int8 = 42
        _ = v10
    default:
        for fi11 := 0; fi11 < 8; fi11++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = len(v5)
            fmt.Println(v1, v1)
            _ = fi11
        }
    }
    if (-2147483648 < -1) {
        for fi12 := 0; fi12 < 15; fi12++ {
            v5 = append(v5, 4294967295)
            v13 := GenS1{F0: true}
            v1 = v1
            _ = fi12
            _ = v13
        }
        v5 = append(v5, 986)
    }
    v3 += 2147483647
    v14 := GenS2{F0: 255, F1: 111}
    var v15 uint64 = 463
    var v16 int = 2147483647
    if (127 < -57) {
        v17 := GenS3{F0: true, F1: (v1 + v1)}
        _ = len(v5)
        _ = v17
    } else {
        _ = len(v2)
        var v18 uint64 = (^(v15))
        _ = v18
    }
    v4 += v3
    if (true && (false && (-8 <= 100))) {
        for fi19 := 0; fi19 < 13; fi19++ {
            fmt.Println(v4, fi19, v1)
            _ = fi19
        }
    }
    var v20 uint64 = v15
    v20 += v20
    v1 = v1
    v21 := GenS2{F0: 0, F1: 1}
    if len(v5) > 0 {
        v5[0] = 725
    }
    switch (v16 - (v16 << ((0) & 63))) {
    case -350:
        v20 -= v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
