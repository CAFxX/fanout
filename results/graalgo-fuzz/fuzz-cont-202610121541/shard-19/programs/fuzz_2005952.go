package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint64; F2 int32 }
func genFn0(p0 int32, p1 int64, p2 int) uint {
    var v6 uint64 = 2
    v7 := []int{p2, (p2 - (p2 << ((255) & 63))), p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 float32 = (3.14159 - float32((3.14159 / 1.0)))
    var v3 int8 = 0
    v4 := []int{32767, 32768, -416, 128}
    var v5 uint32 = 2
    v5 = 943
    v5 = (v5 ^ (v5 ^ (v5 - (^(v5)))))
    switch "a" {
    case "foo bar":
        v4 = append(v4, -9223372036854775808)
    case "\n\t":
        for fi8 := 0; fi8 < 14; fi8++ {
            v3 = (v3 >> ((460) & 63))
            var v9 uint8 = 137
            v2 += float32(0.1)
            _ = fi8
            _ = v9
        }
    case "a":
        switch 128 {
        case 128:
            if (65535 != 496) {
                v10 := GenS1{F0: 127}
                var v11 uint64 = 1
                var v12 string = "hello"
                _ = v10
                _ = v11
                _ = v12
            } else {
                var v13 float32 = (0.1 - ((0.0 - float32(3.14159)) * (v2 * (v2 / v2))))
                _ = v13
            }
        case 566:
            v14 := GenS2{F0: ""}
            _ = v14
        case 127:
            v5 -= 2147483648
        }
    }
    v15 := GenS3{F0: 128, F1: 18446744073709551615, F2: 42}
    var v16 float64 = (float64(float32(-2.25)) * float64((v2 - (v2 / (v2 * v2)))))
    if (18446744073709551615 == 2) {
        var v17 bool = false
        var v18 float64 = v16
        if ((true || (42 != 0)) && (!(519 < -1))) {
            var v19 uint16 = 2
            v20 := GenS1{F0: 128}
            _ = v19
            _ = v20
        }
        _ = v17
        _ = v18
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println(v3, v16, v5)
    v1 = (v1 / (((v1 ^ v1)) | 1))
    var v21 uint32 = (v5 % (((v5 >> ((18446744073709551615) & 63))) | 1))
    v4 = append(v4, 42)
    fmt.Println(v1, v16)
    var v22 uint8 = 40
    fmt.Println(v1, v21)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
