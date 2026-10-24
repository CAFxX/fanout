package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int16; F1 int }
func genFn0(p0 uint) int8 {
    var v7 int8 = 90
    v8 := []int{-9223372036854775808, -2147483648, 0, 32767}
    v7 += 55
    _ = p0
    _ = v7
    _ = v8
    return int8(p0)
}

func main() {
    var v1 int8 = 127
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    var v4 uint = 42
    var v5 int8 = 100
    var v6 int = 0
    v9 := GenS1{F0: 18446744073709551615}
    v5 += v1
    var v10 uint16 = 859
    var v11 int64 = 42
    v12 := GenS1{F0: 884}
    v13 := GenS1{F0: 42}
    var v14 string = "a"
    v15 := GenS2{F0: 1, F1: 679, F2: 2}
    for fi16 := 0; fi16 < 9; fi16++ {
        v1 += int8(fi16)
        v17 := []int{-1, (^(v6)), v6, -2147483648}
        if (-718 > -2147483648) {
            for fi18 := 0; fi18 < 12; fi18++ {
                v1 += (v5 % ((v5) | 1))
                if false {
                    if false {
                        v2[fi18] = (fi18 - fi16)
                        v1 = v1
                        if (!(42 <= 0)) {
                            v10 -= (v10 >> (((v4 & 18446744073709551615)) & 63))
                            switch (fi18 - fi16) {
                            case 637:
                                delete(v3, (((v14 + v14) + "foo bar") + (v14 + v14)))
                            }
                            _ = len(v17)
                        }
                    } else {
                        var v19 int64 = (v11 * -238)
                        _ = v19
                    }
                    v3[v14] = (!(-620 >= fi18))
                    var v20 bool = false
                    _ = v20
                }
                _ = fi18
            }
            var v21 int64 = -9223372036854775808
            fmt.Println(v4, v11)
            _ = v21
        }
        _ = fi16
        _ = v17
    }
    v22 := GenS3{F0: 32767, F1: v6}
    var v23 string = v14
    v4 = (v4 / (((v4 * (v4 & uint(v10)))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
