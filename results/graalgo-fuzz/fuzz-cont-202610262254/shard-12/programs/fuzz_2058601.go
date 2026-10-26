package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint16; F1 uint; F2 uint32 }
func genFn0(p0 uint16, p1 bool, p2 uint) int64 {
    p0 += 2
    var v6 bool = (!true)
    v6 = ((!(127 <= 127)) && false)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -809
}
func genFn1(p0 uint64, p1 int8, p2 uint8) uint16 {
    for fi7 := 0; fi7 < 14; fi7++ {
        if (-2147483648 > 545) {
            for fi8 := 0; fi8 < 13; fi8++ {
                if ((65535 >= 32767) && (true && true)) {
                    v9 := []int{(fi7 ^ -965)}
                    p1 -= p1
                    _ = v9
                }
                _ = fi8
            }
            var v10 uint32 = 811
            _ = v10
        }
        fmt.Println(p0, p2, fi7)
        _ = fi7
    }
    fmt.Println(p2, p1, p2)
    switch "" {
    case "hello":
        v11 := GenS1{F0: p0, F1: ""}
        _ = v11
    default:
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 1
    v3 := []int{9223372036854775807, -478}
    var v4 uint8 = 139
    var v5 uint8 = (v4 + (v4 ^ (v4 + (v4 - (v4 * 1)))))
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := GenS1{F0: 642, F1: "\n\t"}
        v14 := GenS2{F0: (v2 % ((genFn0(2, ((v2 > (v2 >> ((663) & 63))) && (("foo bar" <= "foo bar") || true)), 4294967295)) | 1)), F1: v4}
        _ = fi12
        _ = v13
        _ = v14
    }
    delete(v1, "")
    v15 := GenS3{F0: -1, F1: 18446744073709551615}
    fmt.Println(v5)
    v16 := GenS3{F0: 61, F1: 1}
    _ = len(v3)
    v4 += v4
    v5 += 42
    for fi17 := 0; fi17 < 2; fi17++ {
        fmt.Println(v5, fi17, v5)
        _ = fi17
    }
    v4 = (v5 / ((139) | 1))
    v18 := GenS4{F0: 370, F1: 172, F2: 4294967295}
    if (1e10 != (-2.25 * (1.5 / 1e10))) {
        if (!false) {
            if (4294967295 < 65535) {
                var v19 uint = uint(v5)
                v19 = v19
                _ = v19
            }
        } else {
            if (2147483647 <= 1) {
                var v20 int16 = 0
                _ = v20
            }
            var v21 int = 32768
            _ = v21
        }
        fmt.Println(v4, v4)
    }
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    v4 = v5
    var v22 int8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
