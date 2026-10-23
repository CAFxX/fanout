package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int64 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[string]uint32)
    v3 := make(map[int]uint32)
    var v4 uint8 = 127
    v5 := make(map[int]uint32)
    v6 := make(map[string]int)
    v1 = (v1 * (3.14159 - ((0.0 - v1) / 0.1)))
    var v7 uint64 = 42
    v7 -= v7
    v7 += (v7 + v7)
    if ("" <= "foo bar") {
        if (244 <= 2147483647) {
            fmt.Println(v4, v4, v7)
        }
        var v8 int32 = int32(v4)
        _ = v8
    }
    v4 -= 42
    var v9 uint64 = v7
    v4 += (v4 * 2)
    v10 := GenS1{F0: 0, F1: 671, F2: 9223372036854775807}
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 int8 = 100
            v14 := GenS1{F0: -9223372036854775808, F1: -1, F2: -1}
            _ = fi12
            _ = v13
            _ = v14
        }
        if (fi11 != (fi11 ^ 127)) {
            v7 = 1
        }
        var v15 int8 = -114
        _ = fi11
        _ = v15
    }
    v7 = (^(v9))
    var v16 int32 = 1
    fmt.Println(v4, v7, v1)
    switch 0 {
    case 32767:
        v16 = v16
    case -9223372036854775808:
        if false {
            v4 = (v4 / (((v4 ^ v4)) | 1))
        }
    }
    var v17 uint8 = (^(v4))
    for fi18 := 0; fi18 < 14; fi18++ {
        switch (fi18 / ((-1) | 1)) {
        case 32768:
            v19 := []int{fi18, -9223372036854775808, fi18}
            _ = v19
        }
        for fi20 := 0; fi20 < 6; fi20++ {
            v21 := GenS2{F0: ("foo bar" + ("foo bar" + ("a" + "hello"))), F1: v17}
            if (2147483647 > 0) {
                v17 = 131
                var v22 int64 = 9223372036854775807
                var v23 int32 = v16
                _ = v22
                _ = v23
            } else {
                v24 := []int{(fi20 % ((-716) | 1)), (fi18 / ((fi18) | 1)), fi20, fi18}
                _ = v24
            }
            var v25 int = fi18
            _ = fi20
            _ = v21
            _ = v25
        }
        _ = fi18
    }
    v9 = 150
    v26 := GenS3{F0: 76}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
