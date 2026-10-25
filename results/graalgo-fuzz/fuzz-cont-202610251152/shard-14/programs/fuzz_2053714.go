package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 string }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint64; F1 uint; F2 uint32 }
type GenS4 struct { F0 int8; F1 int32 }
func genFn0(p0 int32, p1 string, p2 uint64) uint {
    var v5 string = (p1 + p1)
    var v6 uint8 = 42
    p1 = (("hello" + p1) + p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 255
}
func genFn1(p0 bool) int32 {
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint = 1
        var v9 int16 = 0
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{-1, 2147483647, -512}
    var v11 uint8 = 127
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 bool = ((((3.14159 / (((3.14159 - 0.1) / 1e10) / float64(float32(float64((1e10 * 3.14159)))))) + 3.14159) - 3.14159) != 0.0)
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = v10
    _ = v11
    return -744
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 uint64 = 1
    v14 := GenS1{F0: 0, F1: -542, F2: "a"}
    if (!false) {
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 bool = (false && (true && (0 < 124)))
            v2[fi15] = "foo bar"
            fmt.Println(v4, v16)
            _ = fi15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := []int{(fi17 | fi17), 9223372036854775807, (fi17 + (-(fi17))), (fi17 << ((977) & 63))}
        var v19 int8 = 0
        v4 += (^(v4))
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 int16 = 0
    v21 := GenS2{F0: v4, F1: 32767}
    if true {
        var v22 int = -9223372036854775808
        fmt.Println(v20, v22)
        v23 := GenS2{F0: (v4 ^ (v4 & v4)), F1: 127}
        _ = v22
        _ = v23
    } else {
        for fi24 := 0; fi24 < 5; fi24++ {
            var v25 uint16 = 65535
            v26 := GenS1{F0: -99, F1: (v20 | v20), F2: "\n\t"}
            switch 120 {
            case 128:
                var v27 int8 = 100
                _ = v27
            case 2147483647:
                v28 := v3[fi24]
                _ = v28
            }
            _ = fi24
            _ = v25
            _ = v26
        }
        var v29 int = 2147483647
        _ = v29
    }
    if (2 <= 2) {
        if (((1 < 857) || true) || (!false)) {
            v4 = 753
        }
        if (1 < 0) {
            var v30 int8 = -128
            v4 = 2
            for fi31 := 0; fi31 < 5; fi31++ {
                for fi32 := 0; fi32 < 17; fi32++ {
                    v33 := GenS1{F0: (v30 + v30), F1: (v20 >> ((255) & 63)), F2: ""}
                    fmt.Println(fi31)
                    _ = fi32
                    _ = v33
                }
                var v34 int = (fi31 - (^(fi31)))
                _ = fi31
                _ = v34
            }
            _ = v30
        }
    }
    v35 := GenS3{F0: v4, F1: 18446744073709551615, F2: 2147483647}
    v36 := GenS4{F0: -73, F1: -2147483648}
    var v37 uint64 = (v4 | v4)
    var v38 int16 = (-(v20))
    v20 = v20
    v4 = v37
    var v39 float32 = 1e10
    var v40 uint = genFn0(-2147483648, ("a" + "foo bar"), v4)
    v41 := GenS3{F0: 1, F1: 18446744073709551615, F2: 28}
    var v42 int = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
