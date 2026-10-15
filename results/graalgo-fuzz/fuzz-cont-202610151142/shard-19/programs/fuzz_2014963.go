package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint8; F1 bool }

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 0
    v3 := make(map[uint32]uint32)
    var v4 int8 = 100
    var v5 int64 = (-(v2))
    v5 = v2
    if (false && false) {
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := GenS1{F0: (v4 + -12)}
            var v8 float32 = (3.14159 * ((float32(1.5) - 0.0) * float32(3.14159)))
            switch "" {
            case "":
                v2 -= -909
            case "a":
                var v9 uint = 255
                _ = v9
            case "\n\t":
                delete(v1, (fi6 << ((961) & 63)))
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        v10 := []int{2147483647, 782, 0, 32768}
        _ = v10
    }
    if true {
        delete(v1, -9223372036854775808)
        v11 := GenS1{F0: (^(v4))}
        v5 -= (v2 ^ (v2 | v5))
        _ = v11
    }
    v12 := GenS1{F0: (v4 >> ((1) & 63))}
    var v13 uint64 = 0
    v4 = v4
    if (false && (!(!(44 < 2)))) {
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := GenS1{F0: v4}
            v16 := GenS2{F0: v4}
            delete(v1, fi14)
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    v13 -= (v13 * 18446744073709551615)
    fmt.Println(v5)
    for fi17 := 0; fi17 < 4; fi17++ {
        switch "a" {
        case "foo bar":
            var v18 int64 = v2
            _ = v18
        case "hello":
            if true {
                _ = len(v1)
                v5 = v5
            } else {
                switch ("a" + "hello") {
                case "hello":
                    v19 := GenS3{F0: 4294967295, F1: 1, F2: 127}
                    _ = v19
                }
                var v20 uint32 = 2
                if (!(int32(v4) <= 42)) {
                    fmt.Println(v4)
                    for fi21 := 0; fi21 < 4; fi21++ {
                        fmt.Println(v13, v13)
                        _ = fi21
                    }
                    var v22 uint = uint(v20)
                    _ = v22
                } else {
                    fmt.Println(v2, v2, v13)
                }
                _ = v20
            }
        case "\n\t":
            var v23 uint8 = 191
            _ = v23
        }
        _ = fi17
    }
    v24 := GenS4{F0: 127, F1: ((v2 / ((-9223372036854775808) | 1)) != -1)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
