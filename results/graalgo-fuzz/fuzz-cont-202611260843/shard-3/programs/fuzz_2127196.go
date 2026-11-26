package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
func genFn0() uint8 {
    v5 := []int{}
    if (float32((0.1 * 0.1)) == (float32((3.14159 + 0.1)) / 1e10)) {
        if (694 <= 255) {
            for fi6 := 0; fi6 < 11; fi6++ {
                var v7 int64 = 9223372036854775807
                for fi8 := 0; fi8 < 18; fi8++ {
                    var v9 float32 = 0.1
                    if len(v5) > 0 {
                        v5[0] = fi8
                    }
                    _ = fi8
                    _ = v9
                }
                for fi10 := 0; fi10 < 3; fi10++ {
                    var v11 uint16 = 0
                    var v12 uint16 = (v11 + (v11 ^ (^(v11))))
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = fi6
                _ = v7
            }
            switch "hello" {
            case "\n\t":
                var v13 uint32 = 2147483648
                _ = v13
            default:
                var v14 uint64 = 2
                _ = v14
            }
        } else {
            for fi15 := 0; fi15 < 15; fi15++ {
                v16 := GenS1{F0: (-(fi15)), F1: 18446744073709551615}
                var v17 uint32 = 227
                _ = fi15
                _ = v16
                _ = v17
            }
            v5 = append(v5, 1)
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 int64 = -9223372036854775808
            _ = fi18
            _ = v19
        }
    }
    var v20 uint8 = 2
    var v21 float64 = -2.25
    var v22 int = 32768
    _ = v20
    _ = v21
    _ = v22
    _ = v5
    return 2
}

func main() {
    var v1 int = -1
    v2 := make(map[uint32]string)
    var v3 uint32 = 65535
    v4 := make(map[uint32]string)
    var v23 float32 = float32(0.1)
    var v24 uint64 = 4294967295
    var v25 int32 = -588
    for fi26 := 0; fi26 < 3; fi26++ {
        v27 := []int{v1, (v1 >> ((255) & 63)), fi26}
        v27 = append(v27, (v1 << ((295) & 63)))
        v25 = (v25 - (v25 - v25))
        _ = fi26
        _ = v27
    }
    v1 = -9223372036854775808
    var v28 int16 = 1000
    if ((v24 & 0) != v24) {
        v1 = (^(v1))
    } else {
        v23 += ((0.1 + v23) - v23)
        for fi29 := 0; fi29 < 11; fi29++ {
            var v30 int = fi29
            _ = fi29
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 19; fi31++ {
        for fi32 := 0; fi32 < 14; fi32++ {
            v33 := []int{(v1 % (((v1 % ((fi31) | 1))) | 1)), fi31, 32768}
            v34 := GenS1{F0: fi32, F1: 18446744073709551615}
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = fi31
    }
    if true {
        var v35 int8 = 100
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
}
