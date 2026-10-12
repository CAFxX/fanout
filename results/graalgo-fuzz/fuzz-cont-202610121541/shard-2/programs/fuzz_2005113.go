package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint32 }

func main() {
    var v1 string = ("\n\t" + ("\n\t" + ("foo bar" + "a")))
    var v2 int = -895
    var v3 string = v1
    var v4 uint = 42
    var v5 uint64 = 532
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 int16 = 427
        for fi8 := 0; fi8 < 19; fi8++ {
            v4 = (v4 ^ 0)
            var v9 bool = (-998 < 1)
            _ = fi8
            _ = v9
        }
        switch v2 {
        case -2147483648:
            v5 = v5
        }
        _ = fi6
        _ = v7
    }
    switch (v2 & v2) {
    case -9223372036854775808:
        fmt.Println(v5, v4)
    case 1:
        v10 := []int{0, v2, v2, -2147483648}
        _ = v10
    default:
        v4 += v4
    }
    var v11 uint16 = 2
    var v12 int32 = -1
    v12 = v12
    v1 = v3
    if (42 == 0) {
        var v13 uint32 = 1
        switch (^(v2)) {
        case -40:
            if ((v4 % (((v4 * (^(v4)))) | 1)) != v4) {
                fmt.Println(v12)
                var v14 uint64 = (^(v5))
                v15 := []int{}
                _ = v14
                _ = v15
            }
        default:
            v2 -= 128
        }
        _ = v13
    }
    if true {
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 float32 = 3.14159
            _ = fi16
            _ = v17
        }
        fmt.Println(v4)
        var v18 bool = ((v5 << ((uint(v11)) & 63)) != v5)
        _ = v18
    } else {
        for fi19 := 0; fi19 < 9; fi19++ {
            v20 := []int{(fi19 ^ 32767), fi19, -275}
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 17; fi21++ {
            for fi22 := 0; fi22 < 16; fi22++ {
                var v23 uint16 = (v11 - (v11 ^ (v11 % (((v11 << ((v4) & 63))) | 1))))
                v24 := GenS1{F0: 79, F1: 18446744073709551615, F2: uint32(v23)}
                if (true || true) {
                    if (42 < 1) {
                        fmt.Println(v11, v11, v2)
                        v25 := GenS1{F0: 127, F1: 42, F2: 0}
                        _ = v25
                    } else {
                        v26 := []int{fi22}
                        _ = v26
                    }
                    var v27 uint = (v4 * (^(v4)))
                    v28 := []int{(fi22 & (fi22 * 833)), fi21}
                    _ = v27
                    _ = v28
                } else {
                    var v29 int = fi21
                    var v30 int32 = (v12 % ((v12) | 1))
                    fmt.Println(fi21)
                    _ = v29
                    _ = v30
                }
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = fi21
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
