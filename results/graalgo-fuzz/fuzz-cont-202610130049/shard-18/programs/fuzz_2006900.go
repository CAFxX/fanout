package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 uint64; F2 int }
func genFn0(p0 uint64, p1 uint16) uint8 {
    v4 := GenS1{F0: 42}
    v5 := []int{808, 592, int(p0)}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 127
}
func genFn1(p0 uint16, p1 uint32, p2 int8) int8 {
    fmt.Println(p0, p0, p0)
    var v6 uint32 = (p1 % ((p1) | 1))
    p0 = 255
    var v7 int8 = (p2 | (-(p2)))
    switch -214 {
    case 32767:
        p0 += (p0 << ((230) & 63))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -128
}

func main() {
    var v1 float32 = (0.1 - float32(1e10))
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    fmt.Println(v1)
    if ((0.1 < 0.0) && (true && (!true))) {
        v1 -= 1.5
        if (false && (!false)) {
            v1 = (v1 * ((v1 / float32((((float64(float32(1.5)) * 3.14159) / ((-2.25 - 0.1) + (float64(1e10) + 1.5))) * 1.5))) - v1))
            switch "" {
            case "":
                var v8 uint32 = 42
                _ = v8
            case "a":
                _ = len(v3)
            case "foo bar":
                if false {
                    fmt.Println(v1)
                    v9 := v2["hello"]
                    v3[(v9 / ((190) | 1))] = ("" + "")
                    _ = v9
                }
            }
        }
        fmt.Println(v1)
    }
    v1 = (v1 + v1)
    v1 -= v1
    v1 = v1
    v1 = 1e10
    var v10 uint8 = 223
    v11 := GenS1{F0: 1}
    v10 = (v10 % ((2) | 1))
    v10 -= genFn0(946, 255)
    var v12 int = 127
    v10 += 1
    v12 = 9223372036854775807
    var v13 bool = true
    switch "a" {
    case "hello":
        for fi14 := 0; fi14 < 14; fi14++ {
            if v13 {
                if v13 {
                    var v15 int8 = -59
                    fmt.Println(v12, v10)
                    fmt.Println(fi14)
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 12; fi16++ {
                        v17 := GenS1{F0: 100}
                        var v18 int32 = -2147483648
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    if (false && v13) {
                        fmt.Println(v13, v13)
                        v19 := v2["a"]
                        _ = v19
                    }
                }
            }
            for fi20 := 0; fi20 < 3; fi20++ {
                v10 = v10
                switch ("hello" + "hello") {
                case "\n\t":
                    switch -2147483648 {
                    case 32768:
                        v21 := GenS1{F0: 0}
                        _ = v21
                    default:
                        v10 = 127
                    }
                case "foo bar":
                    if v13 {
                        v10 = v10
                        delete(v2, "foo bar")
                        v22 := GenS2{F0: 4294967295, F1: -9223372036854775808, F2: (v12 * fi14)}
                        _ = v22
                    }
                case "hello":
                    var v23 int8 = 41
                    _ = v23
                default:
                    v24 := []int{fi14}
                    _ = v24
                }
                v25 := GenS2{F0: 615, F1: fi14, F2: fi14}
                _ = fi20
                _ = v25
            }
            _ = fi14
        }
    case "a":
        var v26 int32 = int32(v10)
        _ = v26
    default:
        switch "a" {
        case "a":
            delete(v3, 4294967295)
        }
    }
    v27 := GenS3{F0: (v10 % ((2) | 1))}
    var v28 uint16 = 1
    v29 := GenS4{F0: -1, F1: 309, F2: v12}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
