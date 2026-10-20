package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int8; F1 uint16; F2 int16 }
type GenS3 struct { F0 int16; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint32; F1 int; F2 bool }
func genFn0(p0 int8, p1 int) int16 {
    v7 := GenS1{F0: 128, F1: 9223372036854775807}
    v8 := []int{32768, p1, (^(p1))}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return int16(p1)
}

func main() {
    var v1 string = "foo bar"
    var v2 uint16 = 255
    v3 := make(map[string]int)
    var v4 string = v1
    v5 := make(map[string]bool)
    var v6 int = 42
    v9 := GenS2{F0: -70, F1: (v2 & v2), F2: -364}
    fmt.Println(v4, v4)
    v10 := []int{-1, 4294967295}
    var v11 float64 = (0.1 + 0.1)
    v12 := GenS3{F0: -1, F1: -128, F2: 343}
    v2 += (^(v2))
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 uint16 = (^(v2))
        v15 := GenS3{F0: int16(fi13), F1: 42, F2: 172}
        _ = fi13
        _ = v14
        _ = v15
    }
    switch "" {
    case "hello":
        v2 += v2
    case "a":
        for fi16 := 0; fi16 < 4; fi16++ {
            switch ("a" + "a") {
            case "hello":
                for fi17 := 0; fi17 < 8; fi17++ {
                    v1 = v4
                    var v18 uint64 = 1
                    v11 = (3.14159 / (((0.1 - 1e10) - v11) / (0.0 - float64(((float32(1e10) * 1e10) - -2.25)))))
                    fmt.Println(fi16)
                    switch (-(fi17)) {
                    case 1:
                        if len(v10) > 0 {
                            v10[0] = fi16
                        }
                    case -2147483648:
                        v4 = ((v1 + v4) + v1)
                    }
                    _ = fi17
                    _ = v18
                }
            default:
                var v19 uint16 = v2
                _ = v19
            }
            var v20 int = fi16
            _ = fi16
            _ = v20
        }
    default:
        v1 = (v4 + "\n\t")
    }
    var v21 int8 = -128
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 uint8 = 0
        if (true && true) {
            if ((!true) && (false || true)) {
                var v24 int16 = 1000
                var v25 int8 = v21
                _ = v24
                _ = v25
            }
        } else {
            var v26 uint8 = (^(v23))
            v27 := GenS4{F0: 1, F1: fi22, F2: false}
            if (!(!false)) {
                v11 += v11
                var v28 int8 = (v21 << ((2) & 63))
                v1 = v4
                _ = v28
            }
            _ = v26
            _ = v27
        }
        _ = fi22
        _ = v23
    }
    if (-32768 == -715) {
        var v29 bool = false
        v4 = v4
        for fi30 := 0; fi30 < 12; fi30++ {
            switch v4 {
            case "foo bar":
                var v31 int32 = 2147483647
                _ = v31
            case "":
                v32 := GenS3{F0: 556, F1: v21, F2: 4294967295}
                _ = v32
            }
            var v33 uint = 41
            _ = fi30
            _ = v33
        }
        _ = v29
    } else {
        if (int8(v6) >= (^(v21))) {
            v1 = ""
            fmt.Println(v1, v4)
            v6 += v6
        } else {
            v6 = (^(v6))
        }
        v21 = (-(v21))
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        for fi35 := 0; fi35 < 6; fi35++ {
            if (!(18446744073709551615 <= 2)) {
                _ = len(v10)
            }
            v11 = v11
            v36 := GenS3{F0: genFn0((v21 ^ v21), (v6 - fi34)), F1: v21, F2: 18446744073709551615}
            _ = fi35
            _ = v36
        }
        v37 := v5[v4]
        v11 = v11
        _ = fi34
        _ = v37
    }
    v6 -= (v6 / (((-(v6))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
