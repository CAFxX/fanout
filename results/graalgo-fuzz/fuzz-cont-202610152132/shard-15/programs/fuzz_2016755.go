package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 uint16; F1 int8; F2 int }
func genFn0(p0 int8, p1 int, p2 int16) uint8 {
    var v4 uint = 65535
    switch int(v4) {
    case 86:
        v5 := []int{p1, (p1 * (p1 + p1))}
        _ = v5
    default:
        var v6 uint16 = 1
        _ = v6
    }
    v4 -= 305
    fmt.Println(p0)
    v7 := GenS1{F0: 997, F1: 2147483647, F2: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return 223
}

func main() {
    var v1 uint16 = 0
    v2 := []string{"\n\t", ("a" + "a"), "foo bar", (("a" + "a") + "hello")}
    var v3 uint64 = 620
    if false {
        var v8 float32 = -2.25
        fmt.Println(v1, v3, v8)
        _ = v8
    }
    v2 = append(v2, "")
    var v9 int32 = 1
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 bool = (int32(v3) == -1)
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: (v3 ^ 805), F1: 42, F2: -1}
    v1 = 336
    for fi13 := 0; fi13 < 12; fi13++ {
        v2 = append(v2, ("foo bar" + ((("hello" + "a") + "a") + "a")))
        for fi14 := 0; fi14 < 9; fi14++ {
            switch "" {
            case "\n\t":
                v9 = (v9 >> ((65535) & 63))
            case "hello":
                if (9223372036854775807 == 376) {
                    var v15 int64 = -1
                    v16 := GenS1{F0: v3, F1: v9, F2: v15}
                    _ = v15
                    _ = v16
                } else {
                    var v17 uint32 = 819
                    fmt.Println(fi14, v1)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(v3, v9, v9)
                    _ = v17
                }
            case "a":
                v18 := GenS2{F0: 340}
                _ = v18
            default:
                v19 := GenS3{F0: 554, F1: int8(v3)}
                _ = v19
            }
            switch fi14 {
            case -1:
                _ = len(v2)
            case -860:
                v3 = v3
            default:
                _ = len(v2)
            }
            var v20 float64 = 1.5
            var v21 int8 = -128
            _ = fi14
            _ = v20
            _ = v21
        }
        v22 := GenS4{F0: 2147483648, F1: v9}
        _ = fi13
        _ = v22
    }
    switch -1 {
    case -9223372036854775808:
        v9 += 2147483647
    default:
        v9 = 1
    }
    v23 := GenS2{F0: 65535}
    v3 += 4294967295
    v24 := GenS5{F0: v1, F1: 1, F2: 205}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v9
}
