package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int64; F2 int16 }
func genFn0(p0 uint, p1 uint16, p2 string) uint64 {
    v6 := []int{1, 32767}
    p0 += (p0 ^ 464)
    p0 -= (p0 << (((^(p0))) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 4294967295
}
func genFn1(p0 int8) int {
    var v7 uint16 = 42
    v8 := []int{4294967295, -9223372036854775808}
    fmt.Println(v7, v7)
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 string = ("hello" + "foo bar")
    v2 := make(map[uint32]uint32)
    var v3 float64 = 1e10
    v4 := []int{127, 42}
    v5 := make(map[int]bool)
    v3 = (v3 * v3)
    if true {
        var v9 int8 = -117
        switch v1 {
        case "":
            fmt.Println(v9, v1)
        case "a":
            var v10 int = 32768
            _ = v10
        }
        v11 := GenS1{F0: 0, F1: 993}
        _ = v11
        _ = v9
    }
    v1 = (("hello" + ("foo bar" + ("" + (((v1 + ("hello" + v1)) + v1) + v1)))) + (v1 + ((v1 + v1) + v1)))
    v12 := GenS1{F0: 32767, F1: -9223372036854775808}
    switch v1 {
    case "\n\t":
        var v13 uint64 = 18446744073709551615
        _ = v13
    case "":
        v14 := GenS1{F0: -1, F1: -951}
        _ = v14
    case "hello":
        v15 := GenS2{F0: -1}
        _ = v15
    }
    switch -1 {
    case -2147483648:
        var v16 int8 = -1
        _ = v16
    case 0:
        for fi17 := 0; fi17 < 15; fi17++ {
            v3 = v3
            _ = fi17
        }
    case 32767:
        if (-32768 != 32767) {
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 int = 9223372036854775807
                if false {
                    switch (-(v19)) {
                    case 4294967295:
                        fmt.Println(fi18)
                    default:
                        v5[genFn1(-1)] = (4294967295 >= 0)
                    }
                    v19 = fi18
                }
                _ = fi18
                _ = v19
            }
            fmt.Println(v3)
            var v20 int16 = -271
            _ = v20
        } else {
            v1 = v1
            v21 := GenS2{F0: 1}
            for fi22 := 0; fi22 < 15; fi22++ {
                fmt.Println(fi22, v1, fi22)
                v1 = "foo bar"
                v3 = (-2.25 / 1.0)
                _ = fi22
            }
            _ = v21
        }
    }
    switch 2147483647 {
    case -2147483648:
        var v23 int16 = -1
        _ = v23
    case 4294967295:
        v3 -= 0.1
    default:
        for fi24 := 0; fi24 < 8; fi24++ {
            switch fi24 {
            case 2147483647:
                v1 = v1
            case 0:
                if true {
                    var v25 uint = 255
                    var v26 int = -9223372036854775808
                    _ = v25
                    _ = v26
                } else {
                    v27 := GenS2{F0: 836}
                    v4 = append(v4, (fi24 + genFn1(0)))
                    _ = v27
                }
            }
            _ = fi24
        }
    }
    v3 = ((1e10 / float64(float32(v3))) + float64(float32(v3)))
    v3 = v3
    fmt.Println(v1)
    v28 := GenS2{F0: 462}
    var v29 uint32 = 2147483647
    v30 := GenS3{F0: -1, F1: -1, F2: 42}
    fmt.Println(v29, v1, v3)
    switch (v1 + (v1 + v1)) {
    case "\n\t":
        if true {
            v29 = (v29 ^ v29)
            if len(v4) > 0 {
                v4[0] = -652
            }
            var v31 int8 = 28
            _ = v31
        }
    default:
        if true {
            fmt.Println(v29)
            switch v1 {
            case "a":
                var v32 int16 = 1000
                _ = v32
            case "":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                _ = len(v4)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
