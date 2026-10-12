package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint; F2 int32 }
func genFn0(p0 int) int {
    v4 := GenS1{F0: 4294967295, F1: 18446744073709551615, F2: p0}
    p0 = 128
    if (0 == 0) {
        switch (p0 ^ p0) {
        case 56:
            fmt.Println(p0)
        case 32768:
            var v5 uint32 = 2147483648
            _ = v5
        case 9223372036854775807:
            switch "hello" {
            case "\n\t":
                var v6 uint32 = 2
                _ = v6
            case "a":
                var v7 float64 = ((3.14159 / (3.14159 + -2.25)) - float64(0.0))
                _ = v7
            }
        default:
            v8 := GenS1{F0: 18446744073709551615, F1: 42, F2: p0}
            _ = v8
        }
        v9 := []int{p0, (p0 + p0), (p0 | -1)}
        _ = v9
    } else {
        if (true && ((((true || false) || (65535 == 2)) && false) || true)) {
            fmt.Println(p0)
        }
        for fi10 := 0; fi10 < 2; fi10++ {
            if (!(!true)) {
                v11 := GenS1{F0: 672, F1: 613, F2: (p0 ^ 910)}
                for fi12 := 0; fi12 < 9; fi12++ {
                    if false {
                        var v13 uint = 2
                        if true {
                            v14 := []int{4294967295, p0}
                            _ = v14
                        }
                        _ = v13
                    }
                    _ = fi12
                }
                p0 = (fi10 >> ((623) & 63))
                _ = v11
            } else {
                fmt.Println(p0, p0)
            }
            for fi15 := 0; fi15 < 7; fi15++ {
                v16 := []int{p0}
                fmt.Println(fi10)
                _ = fi15
                _ = v16
            }
            p0 = (^(fi10))
            _ = fi10
        }
    }
    v17 := GenS2{F0: 18446744073709551615}
    v18 := []int{p0, p0, 706, (p0 ^ (p0 | p0))}
    _ = p0
    _ = v17
    _ = v18
    _ = v4
    return p0
}
func genFn1(p0 uint8) uint {
    p0 = (^(p0))
    v19 := []int{128, 4294967295, 342}
    p0 = (p0 ^ (p0 - p0))
    _ = p0
    _ = v19
    return 4294967295
}

func main() {
    var v1 int64 = 1
    var v2 float64 = (-2.25 / 1.0)
    var v3 int32 = -2147483648
    switch "\n\t" {
    case "foo bar":
        v20 := []int{2147483647, genFn0(-9223372036854775808), genFn0(128)}
        _ = v20
    case "":
        var v21 int = genFn0(808)
        _ = v21
    case "hello":
        v3 = 40
    }
    v2 -= v2
    fmt.Println(v2)
    v1 -= v1
    v22 := GenS3{F0: (v1 & 255), F1: 4294967295, F2: (v3 % (((v3 * v3)) | 1))}
    var v23 int32 = (v3 & 2147483647)
    v2 = (v2 - -2.25)
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
