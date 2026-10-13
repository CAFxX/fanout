package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 int }
func genFn0(p0 int32, p1 int64) uint16 {
    v5 := GenS1{F0: p0, F1: 620}
    var v6 float32 = ((float32(((1.5 * 1e10) / -2.25)) + (3.14159 + 0.0)) / -2.25)
    p0 = 1
    v7 := []int{0}
    v7 = append(v7, 128)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 int, p1 bool, p2 string) bool {
    v8 := GenS2{F0: 1}
    var v9 int = (p0 - -1)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return false
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 871
    v3 := make(map[uint32]string)
    v4 := make(map[int]string)
    v2 = v2
    fmt.Println(v2)
    v10 := GenS2{F0: 0}
    var v11 float32 = 3.14159
    v12 := GenS3{F0: 2147483647}
    var v13 int = -9223372036854775808
    v14 := GenS4{F0: (v13 >> ((18446744073709551615) & 63))}
    switch ("a" + ("hello" + "a")) {
    case "":
        var v15 int8 = -128
        _ = v15
    case "\n\t":
        delete(v1, 2)
    case "a":
        for fi16 := 0; fi16 < 7; fi16++ {
            v3[218] = "foo bar"
            _ = fi16
        }
    default:
        for fi17 := 0; fi17 < 20; fi17++ {
            for fi18 := 0; fi18 < 14; fi18++ {
                fmt.Println(fi17, fi18)
                var v19 int64 = 42
                _ = fi18
                _ = v19
            }
            if (100 < 74) {
                v11 += -2.25
                for fi20 := 0; fi20 < 9; fi20++ {
                    fmt.Println(v11, fi17, fi17)
                    if (!(!false)) {
                        v2 += 4294967295
                        var v21 uint8 = 128
                        _ = v21
                    }
                    v13 -= (fi20 >> ((204) & 63))
                    _ = fi20
                }
            }
            switch ("\n\t" + "\n\t") {
            case "a":
                switch "a" {
                case "a":
                    fmt.Println(v11, v2, v13)
                case "\n\t":
                    var v22 uint32 = 2147483648
                    _ = v22
                default:
                    var v23 uint32 = 65535
                    _ = v23
                }
            case "":
                var v24 float64 = (0.0 / 1.5)
                _ = v24
            }
            _ = fi17
        }
    }
    var v25 int = (^(v13))
    for fi26 := 0; fi26 < 6; fi26++ {
        fmt.Println(v25)
        _ = fi26
    }
    var v27 int16 = -1
    switch v13 {
    case 128:
        if true {
            v25 += v25
            var v28 int16 = -644
            _ = v28
        } else {
            var v29 uint64 = v2
            _ = v29
        }
    case -365:
        v25 += v25
    default:
        fmt.Println(v11)
    }
    v30 := GenS3{F0: -1}
    v31 := GenS5{F0: true, F1: (v13 + -2147483648)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
