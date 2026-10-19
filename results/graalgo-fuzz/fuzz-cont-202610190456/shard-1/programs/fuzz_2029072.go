package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 int32; F1 uint8; F2 uint8 }
func genFn0() uint8 {
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: 467}
        if false {
            var v9 int = fi7
            v9 += 1
            var v10 float32 = float32(float64(-2.25))
            _ = v10
            _ = v9
        } else {
            var v11 uint8 = 42
            for fi12 := 0; fi12 < 9; fi12++ {
                v11 -= 255
                v11 = 197
                switch "" {
                case "foo bar":
                    var v13 bool = ("" < ("" + ("a" + ("hello" + "hello"))))
                    _ = v13
                case "":
                    var v14 uint32 = 2147483647
                    _ = v14
                case "hello":
                    v11 = (v11 >> ((255) & 63))
                }
                v11 = 210
                _ = fi12
            }
            for fi15 := 0; fi15 < 6; fi15++ {
                v16 := GenS2{F0: (-(fi15))}
                v17 := GenS1{F0: 18446744073709551615}
                var v18 int32 = -467
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    if (0 <= 0) {
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 float32 = -2.25
            fmt.Println(fi19)
            _ = fi19
            _ = v20
        }
        var v21 int = 2147483647
        var v22 uint8 = 127
        _ = v21
        _ = v22
    }
    var v23 uint64 = 409
    fmt.Println(v23)
    if true {
        switch 0 {
        case -9223372036854775808:
            if (!(42 <= 128)) {
                var v24 int8 = -1
                _ = v24
            }
        }
    }
    _ = v23
    return 42
}

func main() {
    var v1 int = 2147483647
    v2 := make(map[uint32]int)
    var v3 int16 = 32767
    v4 := []uint32{42, 65535, 815, 2}
    var v5 uint = 292
    var v6 float64 = (0.1 / 1.0)
    v25 := GenS3{F0: 18446744073709551615}
    if ((v3 & v3) < v3) {
        fmt.Println(v1, v6)
        fmt.Println(v3, v1, v6)
        var v26 int64 = -1
        _ = v26
    }
    v27 := GenS4{F0: true, F1: 2}
    v6 -= 0.1
    var v28 int32 = int32(v5)
    switch 1 {
    case 0:
        fmt.Println(v5, v6, v3)
    case 305:
        v1 = (^(v1))
    case 32767:
        v29 := v2[1]
        _ = v29
    default:
        v30 := GenS4{F0: ("" <= "\n\t"), F1: 42}
        _ = v30
    }
    _ = len(v2)
    v6 += v6
    if true {
        for fi31 := 0; fi31 < 11; fi31++ {
            _ = len(v2)
            _ = fi31
        }
    }
    var v32 uint64 = 181
    v32 += (v32 % ((v32) | 1))
    var v33 uint64 = v32
    for fi34 := 0; fi34 < 8; fi34++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v3, v5)
        if len(v4) > 0 {
            v4[0] = 2147483648
        }
        _ = fi34
    }
    _ = len(v2)
    if false {
        if false {
            var v35 float32 = -2.25
            _ = v35
        }
        var v36 int16 = v3
        for fi37 := 0; fi37 < 4; fi37++ {
            var v38 float32 = 1e10
            _ = fi37
            _ = v38
        }
        _ = v36
    } else {
        v28 -= (v28 >> ((628) & 63))
        v39 := GenS4{F0: ((v32 | 2) != 682), F1: 42}
        for fi40 := 0; fi40 < 16; fi40++ {
            v1 -= (fi40 + v1)
            v5 = 18446744073709551615
            _ = fi40
        }
        _ = v39
    }
    v3 = (v3 >> (((v5 & (v5 >> ((v5) & 63)))) & 63))
    if true {
        v41 := GenS5{F0: int32(v3), F1: 128, F2: 25}
        _ = v41
    }
    v42 := GenS5{F0: v28, F1: 47, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
