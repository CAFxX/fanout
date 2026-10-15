package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
func genFn0(p0 int64) uint {
    v4 := []int{}
    v5 := GenS1{F0: 461, F1: ""}
    v6 := GenS1{F0: 42, F1: ""}
    switch "\n\t" {
    case "":
        p0 = (p0 << ((9) & 63))
    case "hello":
        var v7 float64 = -2.25
        _ = v7
    case "a":
        v4 = append(v4, -179)
    default:
        v4 = append(v4, 128)
    }
    fmt.Println(p0)
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 uint8) uint32 {
    v8 := []int{762, -893}
    var v9 uint32 = 23
    var v10 int8 = 42
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 917
}

func main() {
    var v1 int64 = 415
    var v2 uint32 = 2147483648
    var v3 int32 = 677
    v11 := []int{32767, -316, 880}
    v1 = v1
    switch 723 {
    case 127:
        v11 = append(v11, -721)
    case 128:
        if (true || (32767 >= -32768)) {
            var v12 string = (("a" + "") + "hello")
            if (!false) {
                fmt.Println(v1, v3)
                v13 := GenS1{F0: 0, F1: v12}
                _ = v13
            }
            fmt.Println(v12)
            _ = v12
        }
    case 1:
        if (!true) {
            v14 := GenS1{F0: 4294967295, F1: ("" + (("\n\t" + "") + ""))}
            _ = v14
        } else {
            v2 -= v2
        }
    default:
        v15 := GenS1{F0: 1, F1: "hello"}
        _ = v15
    }
    v1 += -1
    v2 = (v2 * v2)
    var v16 uint32 = genFn1(153)
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 uint64 = 18446744073709551615
        v2 -= v16
        for fi19 := 0; fi19 < 6; fi19++ {
            switch fi17 {
            case 127:
                var v20 int16 = -32768
                _ = v20
            case 9223372036854775807:
                if len(v11) > 0 {
                    v11[0] = (fi17 << ((42) & 63))
                }
            }
            var v21 int64 = v1
            if ((^(v18)) <= (v18 >> ((65535) & 63))) {
                v3 -= (v3 ^ 42)
                v3 = -165
                fmt.Println(fi19, v21)
            }
            _ = fi19
            _ = v21
        }
        _ = fi17
        _ = v18
    }
    if ((v16 % ((v16) | 1)) >= (v2 + v16)) {
        v16 += 758
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
}
