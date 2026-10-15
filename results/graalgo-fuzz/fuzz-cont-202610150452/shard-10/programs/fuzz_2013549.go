package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint32, p1 string, p2 int8) string {
    if true {
        p2 += p2
    }
    var v4 int64 = 0
    var v5 uint64 = 18446744073709551615
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (p1 + "")
}
func genFn1(p0 int64) uint16 {
    var v6 int64 = (p0 % ((p0) | 1))
    v7 := []int{128}
    switch 466 {
    case 32767:
        fmt.Println(p0, p0)
    case 980:
        v8 := GenS1{F0: 255, F1: 18, F2: 4294967295}
        _ = v8
    case 0:
        if true {
            _ = len(v7)
            v6 += (v6 - 815)
        }
    }
    var v9 int = 42
    _ = p0
    _ = v6
    _ = v7
    _ = v9
    return 65535
}

func main() {
    var v1 uint = 4294967295
    var v2 uint64 = 332
    v3 := make(map[uint32]bool)
    v1 = v1
    var v10 int64 = -9223372036854775808
    var v11 uint64 = 1
    for fi12 := 0; fi12 < 15; fi12++ {
        v1 = (v1 << ((231) & 63))
        var v13 bool = (true && (genFn1(v10) != uint16(v1)))
        v14 := GenS1{F0: 0, F1: 255, F2: fi12}
        _ = fi12
        _ = v13
        _ = v14
    }
    v10 = (v10 ^ (v10 * 0))
    for fi15 := 0; fi15 < 16; fi15++ {
        v11 -= (v11 + (v2 ^ (v11 << ((v1) & 63))))
        _ = fi15
    }
    v11 = (^(v11))
    var v16 uint8 = 15
    fmt.Println(v11, v1, v10)
    v17 := GenS2{F0: 65535, F1: v10}
    var v18 uint = v1
    var v19 bool = ((v18 % ((v1) | 1)) > v18)
    v20 := GenS1{F0: 0, F1: (v16 ^ v16), F2: 1}
    v18 -= 1
    _ = len(v3)
    for fi21 := 0; fi21 < 4; fi21++ {
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 int64 = (v10 * int64(fi22))
            _ = fi22
            _ = v23
        }
        switch -1 {
        case -213:
            for fi24 := 0; fi24 < 13; fi24++ {
                var v25 int8 = 0
                _ = fi24
                _ = v25
            }
        case 1:
            switch "\n\t" {
            case "hello":
                fmt.Println(v1, v2, v10)
            case "foo bar":
                var v26 int8 = -128
                _ = v26
            }
        }
        if (v19 || (4294967295 != 1)) {
            var v27 int = 1
            var v28 int8 = -1
            _ = v27
            _ = v28
        } else {
            for fi29 := 0; fi29 < 4; fi29++ {
                switch 2147483647 {
                case -2147483648:
                    fmt.Println(v2)
                default:
                    v30 := GenS3{F0: fi29}
                    _ = v30
                }
                v31 := GenS1{F0: 338, F1: (v16 | 42), F2: (fi29 ^ -9223372036854775808)}
                fmt.Println(v19, v16, v16)
                _ = fi29
                _ = v31
            }
        }
        _ = fi21
    }
    fmt.Println(v11)
    v32 := GenS4{F0: true, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
}
