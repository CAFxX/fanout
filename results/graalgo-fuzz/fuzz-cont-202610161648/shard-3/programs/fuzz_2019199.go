package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint8; F1 int; F2 uint16 }

func main() {
    var v1 int = 466
    v2 := make(map[string]uint32)
    var v3 uint8 = 42
    var v4 uint64 = 2
    v5 := v2[("\n\t" + "foo bar")]
    for fi6 := 0; fi6 < 10; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 uint32 = uint32(v3)
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    switch -9223372036854775808 {
    case 127:
        var v9 float32 = -2.25
        _ = v9
    default:
        v1 = 32768
    }
    v10 := GenS1{F0: -32768, F1: 1}
    v5 = v5
    var v11 uint8 = v3
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            v1 = fi12
            _ = fi13
        }
        _ = fi12
    }
    delete(v2, ("a" + ""))
    var v14 uint64 = v4
    for fi15 := 0; fi15 < 16; fi15++ {
        for fi16 := 0; fi16 < 18; fi16++ {
            if ((^(v4)) == (^(v14))) {
                if false {
                    v3 = (v3 >> ((533) & 63))
                }
                if (0 != -2147483648) {
                    switch v1 {
                    case 864:
                        var v17 uint16 = 255
                        _ = v17
                    default:
                        fmt.Println(fi15, v11)
                    }
                } else {
                    if (100 == 112) {
                        v18 := []int{int(v5)}
                        _ = v18
                    }
                    fmt.Println(fi15, v4, fi16)
                    if ((((3.14159 + 0.1) * ((3.14159 + 0.0) + 0.1)) - ((-2.25 + (1e10 + 3.14159)) - 1.5)) >= 0.1) {
                        delete(v2, "\n\t")
                        v4 -= (v14 ^ (v14 + (v14 % (((v14 / (((v14 << ((255) & 63))) | 1))) | 1))))
                        v3 = (v3 & (^(v11)))
                        v11 = (v3 * 131)
                    }
                }
                var v19 int32 = -1
                _ = v19
            }
            v2["a"] = 4294967295
            v20 := GenS1{F0: 42, F1: -84}
            _ = fi16
            _ = v20
        }
        _ = fi15
    }
    fmt.Println(v14)
    fmt.Println(v5, v5)
    v21 := GenS2{F0: 4, F1: -9223372036854775808, F2: 744}
    fmt.Println(v1, v3, v14)
    if (false || true) {
        for fi22 := 0; fi22 < 10; fi22++ {
            v14 = v4
            v23 := []int{fi22, (fi22 % (((^(fi22))) | 1)), (fi22 - 624), 127}
            _ = fi22
            _ = v23
        }
    } else {
        v24 := GenS2{F0: v3, F1: v1, F2: 2}
        _ = v24
    }
    if false {
        var v25 float32 = 1.5
        var v26 int = 568
        _ = v25
        _ = v26
    } else {
        for fi27 := 0; fi27 < 15; fi27++ {
            var v28 uint16 = 292
            v11 = v3
            for fi29 := 0; fi29 < 12; fi29++ {
                v4 += v14
                var v30 uint16 = (v28 / ((v28) | 1))
                _ = fi29
                _ = v30
            }
            _ = fi27
            _ = v28
        }
        var v31 string = "a"
        var v32 uint16 = 417
        _ = v31
        _ = v32
    }
    switch v1 {
    case 0:
        v33 := []int{v1, (^(v1)), v1, (v1 + (v1 >> ((731) & 63)))}
        _ = v33
    case -1:
        fmt.Println(v11, v1, v3)
    case 9223372036854775807:
        for fi34 := 0; fi34 < 2; fi34++ {
            v1 += fi34
            v11 += (v11 % ((v3) | 1))
            _ = fi34
        }
    }
    if (!true) {
        var v35 uint8 = uint8(v14)
        _ = v35
    } else {
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
