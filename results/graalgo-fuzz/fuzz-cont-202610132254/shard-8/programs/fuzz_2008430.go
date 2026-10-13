package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 uint64; F2 int64 }
func genFn0(p0 int, p1 uint64) int32 {
    fmt.Println(p1, p0)
    p1 = (p1 ^ (p1 * p1))
    v5 := []int{(p0 | (p0 >> ((548) & 63))), p0, (p0 >> ((446) & 63)), p0}
    v6 := GenS1{F0: 624}
    v7 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 525
}

func main() {
    var v1 uint = 42
    var v2 uint32 = 4294967295
    v3 := make(map[int]bool)
    v4 := make(map[int]string)
    var v8 uint8 = 183
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 int = fi9
        if false {
            var v11 int16 = 1000
            _ = v11
        }
        switch ("hello" + "\n\t") {
        case "foo bar":
            var v12 uint64 = 1
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    v13 := GenS2{F0: 128}
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := GenS3{F0: 18446744073709551615, F1: 2, F2: 9223372036854775807}
        var v16 uint8 = (v8 | 128)
        v1 += (^(v1))
        _ = fi14
        _ = v15
        _ = v16
    }
    if (0.1 > ((-2.25 * 0.1) + -2.25)) {
        var v17 int32 = 1
        fmt.Println(v8, v8, v1)
        v2 += 4294967295
        _ = v17
    } else {
        v1 = (v1 % (((v1 / (((v1 * v1)) | 1))) | 1))
    }
    if (127 >= 127) {
        var v18 int8 = -78
        v18 -= (^(v18))
        _ = v18
    }
    v4[0] = "hello"
    fmt.Println(v8)
    fmt.Println(v8, v1)
    var v19 string = "hello"
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
