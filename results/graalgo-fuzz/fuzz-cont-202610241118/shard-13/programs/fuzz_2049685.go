package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint16 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64 }
func genFn0() int32 {
    var v7 uint32 = 1
    for fi8 := 0; fi8 < 6; fi8++ {
        v7 = v7
        var v9 int64 = -461
        var v10 uint64 = 1
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v7
    return -2147483648
}

func main() {
    v1 := []uint32{0}
    v2 := make(map[int]bool)
    var v3 float64 = 0.1
    var v4 int16 = -32768
    var v5 float32 = (float32(1.5) / 1.0)
    var v6 string = ("hello" + (((("" + (("a" + "hello") + "hello")) + "hello") + "a") + ""))
    v3 = (-2.25 / 1.0)
    if (102 != 255) {
        _ = len(v1)
        var v11 uint16 = 32767
        fmt.Println(v6, v6, v5)
        _ = v11
    }
    if true {
        if (0 == 4294967295) {
            if true {
                delete(v2, -2147483648)
                v3 -= (float64(float32((v3 + -2.25))) * v3)
                v3 = v3
            } else {
                var v12 uint16 = 134
                _ = v12
            }
            fmt.Println(v3)
        } else {
            var v13 int64 = 9223372036854775807
            _ = v13
        }
    } else {
        var v14 int = 30
        for fi15 := 0; fi15 < 13; fi15++ {
            v2[(v14 + 127)] = true
            v4 += v4
            _ = fi15
        }
        _ = v14
    }
    if true {
        for fi16 := 0; fi16 < 2; fi16++ {
            v4 -= v4
            delete(v2, fi16)
            var v17 int32 = -606
            _ = fi16
            _ = v17
        }
    } else {
        for fi18 := 0; fi18 < 10; fi18++ {
            delete(v2, (fi18 | (fi18 + (-(fi18)))))
            if true {
                fmt.Println(v5)
                switch 1 {
                case 32767:
                    var v19 int16 = v4
                    _ = v19
                default:
                    v20 := GenS1{F0: 2, F1: "a", F2: 226}
                    _ = v20
                }
                v21 := GenS2{F0: 0, F1: 9223372036854775807}
                _ = v21
            }
            v22 := GenS3{F0: 857}
            _ = fi18
            _ = v22
        }
        if false {
            v23 := GenS3{F0: 1}
            var v24 uint8 = 128
            _ = v23
            _ = v24
        }
    }
    var v25 uint64 = 4294967295
    if (!true) {
        v4 = v4
    } else {
        var v26 bool = true
        _ = v26
    }
    v27 := GenS2{F0: 1, F1: 0}
    switch "\n\t" {
    case "foo bar":
        v1 = append(v1, 65535)
    default:
        v28 := GenS2{F0: -75, F1: 985}
        _ = v28
    }
    fmt.Println(v6)
    v6 = v6
    v25 = (v25 ^ 42)
    v29 := GenS2{F0: -100, F1: 426}
    if (((v4 ^ 32767) == (v4 % ((32767) | 1))) && (((!true) && (1 > 65535)) && (false || true))) {
        if (!false) {
            v30 := v2[866]
            fmt.Println(v5, v6)
            for fi31 := 0; fi31 < 18; fi31++ {
                var v32 float64 = 0.1
                var v33 int8 = -100
                v34 := GenS3{F0: -493}
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
            }
            _ = v30
        }
        v35 := GenS1{F0: 65535, F1: "foo bar", F2: 269}
        _ = v35
    }
    var v36 int8 = 42
    if false {
        v5 += v5
        v37 := GenS4{F0: 0}
        _ = v37
    }
    fmt.Println(v5, v6, v36)
    v3 -= (3.14159 + (v3 * float64(1e10)))
    switch 32767 {
    case 56:
        v25 = 1
    case -2147483648:
        delete(v2, 2147483647)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v29
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
