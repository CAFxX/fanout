package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 string }
type GenS2 struct { F0 uint }
func genFn0(p0 uint8, p1 int8, p2 uint8) int {
    switch "a" {
    case "foo bar":
        switch 128 {
        case 42:
            if ((p1 | int8(p2)) <= 0) {
                for fi6 := 0; fi6 < 12; fi6++ {
                    switch "" {
                    case "":
                        v7 := GenS1{F0: -2147483648, F1: -2147483648, F2: (("a" + "a") + ("hello" + ("hello" + "hello")))}
                        _ = v7
                    default:
                        var v8 uint16 = 0
                        _ = v8
                    }
                    var v9 int8 = (p1 & p1)
                    fmt.Println(p0)
                    _ = fi6
                    _ = v9
                }
            } else {
                v10 := GenS2{F0: 65535}
                v11 := []int{-36}
                p2 -= p2
                _ = v10
                _ = v11
            }
        case -1:
            p2 += 247
        case -9223372036854775808:
            var v12 int = 81
            _ = v12
        }
    }
    p2 -= (p2 / (((p2 % (((p0 * p2)) | 1))) | 1))
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 uint = 676
    var v2 uint8 = 42
    var v3 int8 = -27
    var v4 uint32 = 4294967295
    var v5 int16 = -323
    v4 += (v4 % (((^(v4))) | 1))
    v13 := GenS1{F0: -2147483648, F1: -84, F2: ""}
    var v14 uint16 = 7
    v15 := []int{-1}
    var v16 uint8 = (v2 >> ((uint(v2)) & 63))
    v3 -= v3
    if false {
        v17 := GenS1{F0: 1, F1: -2147483648, F2: ("hello" + "a")}
        var v18 uint64 = 2
        v4 += v4
        _ = v17
        _ = v18
    } else {
        if true {
            v15 = append(v15, 987)
            v14 += v14
            switch 753 {
            case 42:
                v19 := GenS2{F0: 18446744073709551615}
                _ = v19
            case 32768:
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
            case 9223372036854775807:
                v14 = v14
            }
        }
    }
    if (793 > 4294967295) {
        v2 -= (v16 * v16)
        for fi20 := 0; fi20 < 9; fi20++ {
            fmt.Println(v1, v4)
            _ = fi20
        }
        var v21 int32 = -417
        _ = v21
    }
    v16 += (v16 | v16)
    v22 := GenS2{F0: (^(v1))}
    v15 = append(v15, -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
