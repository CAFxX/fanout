package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]int)
    var v3 int32 = 42
    var v4 uint32 = 65535
    v5 := make(map[int]int)
    var v6 int64 = -9223372036854775808
    switch "\n\t" {
    case "hello":
        v5[4294967295] = 9223372036854775807
    case "a":
        v7 := []int{-1, 729}
        _ = v7
    case "\n\t":
        v3 = (v3 ^ (v3 >> ((647) & 63)))
    }
    fmt.Println(v4, v1)
    delete(v5, 32768)
    v8 := GenS1{F0: 2, F1: 42, F2: 42}
    switch 505 {
    case -909:
        v4 = (v4 | 2)
    }
    if (v6 <= v6) {
        if ((!((^(v3)) != v3)) || false) {
            if (3.14159 >= 0.0) {
                v6 = -9223372036854775808
                v1 -= 0.1
                v9 := GenS1{F0: 42, F1: 1000, F2: 255}
                _ = v9
            }
            var v10 uint16 = 255
            for fi11 := 0; fi11 < 9; fi11++ {
                v4 -= v4
                v12 := []int{(fi11 * (fi11 << ((0) & 63))), -1}
                _ = fi11
                _ = v12
            }
            _ = v10
        }
    }
    var v13 int8 = 127
    var v14 uint = 42
    v1 = v1
    if true {
        if (0.0 <= 0.0) {
            v14 = 2
            var v15 uint64 = 42
            _ = v15
        }
        for fi16 := 0; fi16 < 2; fi16++ {
            fmt.Println(v14, v4)
            _ = fi16
        }
    }
    v17 := GenS1{F0: (v14 << ((v14) & 63)), F1: -190, F2: 2}
    switch "\n\t" {
    case "a":
        for fi18 := 0; fi18 < 8; fi18++ {
            switch ("\n\t" + "\n\t") {
            case "":
                v19 := []int{(fi18 % ((fi18) | 1)), (fi18 * int(v14))}
                _ = v19
            }
            _ = fi18
        }
    case "foo bar":
        for fi20 := 0; fi20 < 19; fi20++ {
            if false {
                for fi21 := 0; fi21 < 14; fi21++ {
                    if (!false) {
                        v1 += 1.5
                        var v22 uint16 = 2
                        delete(v5, fi21)
                        _ = v22
                    }
                    var v23 bool = (uint(v4) == (^(v14)))
                    _ = fi21
                    _ = v23
                }
            }
            switch "" {
            case "a":
                v24 := []int{(fi20 ^ (fi20 ^ -748)), (-(fi20))}
                _ = v24
            case "\n\t":
                switch (fi20 ^ (fi20 >> (((v14 * (^(v14)))) & 63))) {
                case -2147483648:
                    var v25 int16 = -32768
                    _ = v25
                case -299:
                    if (32767 != uint16(v14)) {
                        for fi26 := 0; fi26 < 13; fi26++ {
                            var v27 int64 = 1
                            v28 := GenS1{F0: (^(v14)), F1: -32768, F2: v14}
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        v29 := []int{fi20, fi20, fi20}
                        _ = v29
                    } else {
                        var v30 uint8 = 0
                        _ = v30
                    }
                }
            default:
                fmt.Println(fi20, v6, v1)
            }
            v31 := v5[(fi20 | (fi20 + (fi20 + -9223372036854775808)))]
            _ = fi20
            _ = v31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
