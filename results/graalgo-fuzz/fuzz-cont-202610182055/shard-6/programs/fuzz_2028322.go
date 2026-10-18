package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint64; F2 int16 }
func genFn0(p0 string, p1 int16, p2 int) string {
    for fi4 := 0; fi4 < 12; fi4++ {
        if true {
            p2 -= 32768
        } else {
            fmt.Println(p0)
            var v5 uint64 = 18446744073709551615
            _ = v5
        }
        fmt.Println(fi4)
        _ = fi4
    }
    var v6 uint8 = 1
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int64 = 1000
        _ = fi7
        _ = v8
    }
    var v9 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return "foo bar"
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    var v3 uint = 255
    v10 := []int{2147483647, 1}
    fmt.Println(v3)
    var v11 uint64 = 4294967295
    v10 = append(v10, 4294967295)
    for fi12 := 0; fi12 < 11; fi12++ {
        if ((v11 % ((18446744073709551615) | 1)) < (v11 + 42)) {
            if true {
                var v13 int64 = -43
                _ = v13
            }
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        } else {
            delete(v2, 42)
            switch -458 {
            case 1:
                var v14 uint16 = 2
                _ = v14
            }
            var v15 float64 = float64((0.0 - 0.0))
            _ = v15
        }
        _ = fi12
    }
    v16 := v2[1]
    v11 += (v11 - 1)
    if false {
        switch 1 {
        case -9223372036854775808:
            for fi17 := 0; fi17 < 11; fi17++ {
                for fi18 := 0; fi18 < 2; fi18++ {
                    if ((!((3.14159 >= 3.14159) || (0 < 666))) || true) {
                        var v19 float64 = float64(0.0)
                        _ = v19
                    } else {
                        var v20 int = fi17
                        if ("\n\t" != "hello") {
                            var v21 int16 = 1000
                            _ = v21
                        }
                        delete(v2, v16)
                        _ = v20
                    }
                    switch fi18 {
                    case 637:
                        v11 = (^(v11))
                    case -813:
                        fmt.Println(fi18)
                    case 633:
                        var v22 bool = ((v3 ^ (^(v3))) == (v3 << (((v3 ^ uint(v11))) & 63)))
                        _ = v22
                    default:
                        var v23 uint8 = 253
                        _ = v23
                    }
                    v3 -= 658
                    _ = fi18
                }
                var v24 int16 = 32767
                v2[v16] = 4294967295
                _ = fi17
                _ = v24
            }
        case 127:
            var v25 uint8 = 118
            _ = v25
        case -730:
            fmt.Println(v11, v11, v3)
        }
        v11 = uint64(v3)
        fmt.Println(v16, v16)
    } else {
        for fi26 := 0; fi26 < 16; fi26++ {
            v16 -= v16
            v16 -= (v16 * v16)
            switch "hello" {
            case "":
                v11 += v11
            case "hello":
                delete(v2, (v16 / ((2147483648) | 1)))
            case "\n\t":
                v11 -= v11
            }
            _ = fi26
        }
        var v27 float32 = ((0.0 + -2.25) * -2.25)
        var v28 bool = (!false)
        _ = v27
        _ = v28
    }
    v29 := GenS1{F0: (^(v11)), F1: 0}
    switch "foo bar" {
    case "hello":
        for fi30 := 0; fi30 < 6; fi30++ {
            v16 = 65535
            _ = fi30
        }
    case "a":
        v3 -= (^(v3))
    default:
        for fi31 := 0; fi31 < 3; fi31++ {
            v3 = v3
            _ = fi31
        }
    }
    switch "\n\t" {
    case "hello":
        for fi32 := 0; fi32 < 3; fi32++ {
            switch "hello" {
            case "hello":
                v2[(^(v16))] = v16
            }
            _ = fi32
        }
    case "":
        var v33 int8 = 1
        _ = v33
    case "a":
        v34 := GenS2{F0: v3, F1: 0, F2: uint32(v11)}
        _ = v34
    }
    if (!false) {
        var v35 bool = (true || (-219 > 42))
        var v36 int8 = -57
        for fi37 := 0; fi37 < 20; fi37++ {
            v16 = (v16 ^ 959)
            v38 := GenS3{F0: (v36 % (((v36 << ((1) & 63))) | 1)), F1: (^(v11)), F2: -1}
            _ = fi37
            _ = v38
        }
        _ = v35
        _ = v36
    }
    v3 = v3
    v11 = (^(v11))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v29
    _ = v3
}
