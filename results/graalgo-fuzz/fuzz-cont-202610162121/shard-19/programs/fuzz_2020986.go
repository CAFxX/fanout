package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 int }
func genFn0(p0 uint64, p1 int, p2 uint) int64 {
    if (!(true && (106 >= 42))) {
        for fi5 := 0; fi5 < 8; fi5++ {
            var v6 int64 = -1
            for fi7 := 0; fi7 < 11; fi7++ {
                fmt.Println(fi5, v6, p2)
                p0 = (p0 ^ (p0 >> (((p2 ^ (p2 / (((p2 | (p2 / (((p2 / ((18446744073709551615) | 1))) | 1)))) | 1)))) & 63)))
                var v8 uint8 = 32
                _ = fi7
                _ = v8
            }
            _ = fi5
            _ = v6
        }
    } else {
        if (p1 >= p1) {
            for fi9 := 0; fi9 < 12; fi9++ {
                for fi10 := 0; fi10 < 2; fi10++ {
                    v11 := []int{}
                    _ = fi10
                    _ = v11
                }
                _ = fi9
            }
            fmt.Println(p2, p1, p0)
            for fi12 := 0; fi12 < 3; fi12++ {
                switch "hello" {
                case "":
                    v13 := []int{(fi12 | (fi12 + (fi12 | fi12)))}
                    _ = v13
                }
                if (3.14159 >= (1.5 / 1.0)) {
                    var v14 int16 = 32767
                    v15 := []int{-9223372036854775808}
                    _ = v14
                    _ = v15
                }
                var v16 int64 = 42
                _ = fi12
                _ = v16
            }
        }
        p2 = p2
    }
    if (p0 <= 18446744073709551615) {
        for fi17 := 0; fi17 < 6; fi17++ {
            p0 = (p0 / ((p0) | 1))
            _ = fi17
        }
        fmt.Println(p2, p1, p0)
    } else {
        if (100 >= 1) {
            var v18 uint32 = 4294967295
            _ = v18
        } else {
            v19 := []int{-526, (^(p1)), p1, (p1 % ((p1) | 1))}
            var v20 int = (p1 * 9223372036854775807)
            var v21 uint = p2
            _ = v19
            _ = v20
            _ = v21
        }
    }
    var v22 uint64 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return -9223372036854775808
}

func main() {
    var v1 float64 = ((0.1 * (0.1 * 1.5)) + 0.0)
    var v2 uint64 = 168
    var v3 uint16 = 32767
    v4 := []int{0, -2147483648}
    v3 = v3
    if (128 <= 127) {
        v23 := GenS1{F0: -128}
        for fi24 := 0; fi24 < 11; fi24++ {
            v1 = 0.0
            var v25 uint8 = 227
            var v26 float64 = (v1 / (float64(3.14159) / 0.1))
            _ = fi24
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 16; fi27++ {
            switch "\n\t" {
            case "":
                var v28 bool = false
                _ = v28
            }
            _ = fi27
        }
        _ = v23
    }
    if (-870 > genFn0(v2, -9223372036854775808, 42)) {
        if true {
            var v29 int8 = 42
            v30 := GenS2{F0: v2}
            var v31 int32 = 2147483647
            _ = v29
            _ = v30
            _ = v31
        }
    }
    fmt.Println(v2)
    var v32 uint64 = v2
    for fi33 := 0; fi33 < 19; fi33++ {
        for fi34 := 0; fi34 < 6; fi34++ {
            v35 := GenS1{F0: 27}
            fmt.Println(fi33, v3)
            v3 -= v3
            _ = fi34
            _ = v35
        }
        for fi36 := 0; fi36 < 20; fi36++ {
            var v37 int32 = -1
            var v38 float64 = ((3.14159 - v1) / ((v1 / 0.1) - v1))
            _ = fi36
            _ = v37
            _ = v38
        }
        var v39 uint = 4294967295
        _ = fi33
        _ = v39
    }
    v40 := GenS3{F0: (v2 >> ((18446744073709551615) & 63)), F1: 128}
    _ = len(v4)
    v41 := GenS1{F0: -43}
    v3 = 286
    var v42 int64 = -565
    v4 = append(v4, 665)
    var v43 float32 = 1.5
    v43 = (float32(float64(v43)) * float32((float64((v43 / v43)) - (v1 - 0.0))))
    for fi44 := 0; fi44 < 3; fi44++ {
        _ = len(v4)
        _ = fi44
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        _ = len(v4)
        var v45 float32 = -2.25
        fmt.Println(v32, v1, v32)
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
