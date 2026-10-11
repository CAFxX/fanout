package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int64 }
func genFn0(p0 int32, p1 uint, p2 uint16) uint16 {
    v4 := []int{}
    var v5 int32 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p2
}
func genFn1() int8 {
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := GenS1{F0: 2, F1: 65535, F2: 1}
        switch fi6 {
        case 1:
            v8 := []int{}
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    var v9 int32 = -1
    v10 := GenS1{F0: 837, F1: 4294967295, F2: -46}
    if false {
        v11 := []int{1, 4294967295, 1, -960}
        if len(v11) > 0 {
            v11[0] = 128
        }
        v9 = (v9 - v9)
        _ = v11
    } else {
        switch -9223372036854775808 {
        case -821:
            var v12 int = 2147483647
            _ = v12
        case -879:
            v13 := GenS1{F0: 483, F1: 0, F2: 0}
            _ = v13
        case 127:
            v9 = (v9 / ((-300) | 1))
        }
    }
    _ = v10
    _ = v9
    return -121
}

func main() {
    var v1 bool = (false && true)
    var v2 uint16 = 42
    var v3 int8 = 26
    switch 0 {
    case 32767:
        v14 := []int{2147483647, 1, 0}
        _ = v14
    case 9223372036854775807:
        v2 -= 2
    }
    var v15 uint = 4294967295
    fmt.Println(v2)
    v2 = (v2 - v2)
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 uint16 = 1
        for fi18 := 0; fi18 < 7; fi18++ {
            v17 -= (v2 + genFn0(-628, 4294967295, (v2 << ((v15) & 63))))
            _ = fi18
        }
        for fi19 := 0; fi19 < 15; fi19++ {
            var v20 float64 = float64(1.5)
            if v1 {
                if false {
                    fmt.Println(v17)
                    v21 := GenS1{F0: 18446744073709551615, F1: (v15 ^ (v15 ^ v15)), F2: 9223372036854775807}
                    if ((-505 == 2147483647) && (962 == 42)) {
                        v17 += 42
                        v22 := []int{9223372036854775807, 2147483647, (-(fi19))}
                        _ = v22
                    }
                    _ = v21
                }
            }
            v23 := GenS1{F0: 18446744073709551615, F1: v15, F2: 737}
            _ = fi19
            _ = v20
            _ = v23
        }
        _ = fi16
        _ = v17
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        v1 = false
        if false {
            v25 := GenS1{F0: 4294967295, F1: v15, F2: 9223372036854775807}
            v3 += (v3 % ((v3) | 1))
            _ = v25
        } else {
            var v26 int = int(v3)
            fmt.Println(fi24)
            _ = v26
        }
        _ = fi24
    }
    v3 = v3
    v27 := []int{1, 1, 42, 32767}
    v27 = append(v27, 2147483647)
    var v28 float64 = 3.14159
    fmt.Println(v28, v2, v28)
    var v29 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
