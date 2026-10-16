package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint64) uint16 {
    v5 := GenS1{F0: uint8(p0)}
    v6 := GenS1{F0: 0}
    v7 := []int{1, 127, -9223372036854775808, -2147483648}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 32767
}

func main() {
    var v1 string = (("" + ("" + "foo bar")) + ("\n\t" + ((("a" + "hello") + "") + ("" + ("hello" + ("hello" + "foo bar"))))))
    var v2 float64 = 0.0
    v3 := make(map[uint32]int)
    var v4 int16 = 1
    v1 = v1
    if false {
        v4 = v4
        v8 := GenS2{F0: 100, F1: 18446744073709551615, F2: 100}
        v2 = (((v2 / (v2 / (float64((0.1 / (1e10 * 0.1))) + v2))) * (v2 * 3.14159)) - 1.5)
        _ = v8
    } else {
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 int32 = -1
            _ = fi9
            _ = v10
        }
    }
    fmt.Println(v1)
    v2 -= v2
    switch 9223372036854775807 {
    case 651:
        v1 = v1
    case 32768:
        v2 = (0.1 + float64(((0.1 * 0.1) + float32(v2))))
    case -9223372036854775808:
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 string = ("a" + v1)
            switch v1 {
            case "foo bar":
                var v13 uint16 = genFn0(42)
                _ = v13
            }
            _ = fi11
            _ = v12
        }
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            if (545 != 255) {
                if false {
                    v1 = ((v1 + v1) + v1)
                }
                v16 := GenS3{F0: 623, F1: 73}
                _ = v16
            }
            v17 := []int{128, fi15}
            v3[2147483648] = 0
            _ = fi15
            _ = v17
        }
        for fi18 := 0; fi18 < 5; fi18++ {
            if false {
                var v19 uint16 = 1
                var v20 int16 = (v4 >> ((255) & 63))
                if false {
                    v21 := GenS1{F0: 128}
                    v4 += (v4 << ((65535) & 63))
                    switch fi14 {
                    case 4294967295:
                        var v22 uint8 = 2
                        _ = v22
                    case 2147483647:
                        v23 := GenS1{F0: 1}
                        _ = v23
                    case -9223372036854775808:
                        var v24 uint16 = (v19 << ((1) & 63))
                        _ = v24
                    }
                    _ = v21
                }
                _ = v19
                _ = v20
            } else {
                var v25 int32 = -2147483648
                _ = v25
            }
            v26 := v3[638]
            var v27 uint8 = 235
            _ = fi18
            _ = v26
            _ = v27
        }
        _ = fi14
    }
    fmt.Println(v4, v1, v4)
    v28 := []int{-102}
    v2 = -2.25
    v29 := GenS4{F0: 448}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
