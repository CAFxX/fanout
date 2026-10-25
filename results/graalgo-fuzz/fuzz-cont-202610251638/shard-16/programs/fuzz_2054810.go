package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 int }
func genFn0(p0 int, p1 int32) bool {
    v6 := []int{(p0 ^ p0), p0, p0, p0}
    fmt.Println(p1, p1)
    v7 := GenS1{F0: 9223372036854775807, F1: 169}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (4 >= 609)
}

func main() {
    var v1 uint = 428
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    var v4 uint16 = 0
    v5 := []uint32{1, 2147483648, 2, 1}
    v8 := GenS2{F0: 2147483647, F1: 237, F2: int64(v1)}
    if true {
        v4 += 433
        v4 = 32767
        for fi9 := 0; fi9 < 15; fi9++ {
            v4 = v4
            if true {
                _ = len(v2)
                var v10 int32 = -63
                var v11 uint = 42
                _ = v10
                _ = v11
            } else {
                if (42 <= 127) {
                    v4 = (^(v4))
                }
                v1 = (v1 & v1)
            }
            _ = fi9
        }
    } else {
        var v12 uint = (^(v1))
        switch ("a" + "\n\t") {
        case "a":
            v12 -= v1
        case "":
            var v13 uint = 884
            _ = v13
        case "hello":
            switch 4294967295 {
            case 9223372036854775807:
                fmt.Println(v1, v4)
            case 242:
                delete(v2, 504)
            case 127:
                v12 += 2
            }
        }
        switch "a" {
        case "\n\t":
            for fi14 := 0; fi14 < 6; fi14++ {
                v4 -= 1
                var v15 float64 = 3.14159
                v5 = append(v5, 4294967295)
                _ = fi14
                _ = v15
            }
        case "foo bar":
            switch -1 {
            case -533:
                var v16 int8 = -123
                _ = v16
            default:
                if (!("a" != "foo bar")) {
                    var v17 int = 152
                    v18 := GenS1{F0: 1, F1: 42}
                    _ = v17
                    _ = v18
                }
            }
        case "hello":
            var v19 float64 = (((((-2.25 / -2.25) + (float64(0.0) / float64((-2.25 - 1e10)))) - -2.25) * 0.1) + 1.5)
            _ = v19
        default:
            for fi20 := 0; fi20 < 12; fi20++ {
                v21 := v2[42]
                _ = fi20
                _ = v21
            }
        }
        _ = v12
    }
    v22 := GenS3{F0: 104}
    v23 := GenS3{F0: 1}
    _ = len(v5)
    var v24 string = "a"
    for fi25 := 0; fi25 < 11; fi25++ {
        v4 = v4
        for fi26 := 0; fi26 < 2; fi26++ {
            v27 := GenS4{F0: 1, F1: fi25}
            fmt.Println(v24)
            if true {
                v28 := v2[fi26]
                var v29 int64 = 0
                v1 -= (v1 - 520)
                _ = v28
                _ = v29
            } else {
                v4 -= (v4 << ((693) & 63))
                var v30 uint = (^(v1))
                var v31 bool = true
                _ = v30
                _ = v31
            }
            _ = fi26
            _ = v27
        }
        switch fi25 {
        case -9223372036854775808:
            v24 = v24
        case 1:
            fmt.Println(v24, v1, v1)
        }
        _ = fi25
    }
    v4 -= 65535
    v5 = append(v5, 2147483648)
    v32 := GenS3{F0: 0}
    fmt.Println(v1)
    var v33 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v8
}
