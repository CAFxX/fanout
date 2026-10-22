package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int16; F1 bool }

func main() {
    var v1 int = 338
    var v2 uint32 = 0
    var v3 int64 = 42
    var v4 int = v1
    var v5 uint8 = 42
    if ((v2 == 2147483648) && false) {
        if ((!(true && false)) || ("\n\t" >= ((((("\n\t" + (("\n\t" + "hello") + "hello")) + ("a" + "a")) + "foo bar") + "a") + "foo bar"))) {
            v6 := []int{}
            _ = v6
        }
        v3 = 739
    }
    fmt.Println(v2)
    if (127 != (^(v4))) {
        v4 -= (-(v1))
        switch v1 {
        case 127:
            var v7 uint8 = (v5 + v5)
            _ = v7
        case 4294967295:
            v3 = -9223372036854775808
        }
    }
    switch 2147483647 {
    case -637:
        switch (v1 + (v4 + (v4 / (((v4 / (((v4 ^ (v1 - v4))) | 1))) | 1)))) {
        case 42:
            v8 := GenS1{F0: 1000, F1: (^(v2))}
            _ = v8
        case 32767:
            switch "" {
            case "\n\t":
                v4 = 2147483647
            case "a":
                var v9 uint16 = 32767
                _ = v9
            default:
                v3 -= v3
            }
        default:
            v2 -= v2
        }
    case -1:
        v1 += 9223372036854775807
    }
    var v10 uint64 = 42
    v4 = (v4 | (^(v4)))
    v11 := GenS1{F0: 0, F1: (v2 - v2)}
    if false {
        v12 := GenS2{F0: 1, F1: true}
        if (v1 > 0) {
            v13 := []int{42}
            var v14 uint8 = 4
            _ = v13
            _ = v14
        }
        _ = v12
    }
    if false {
        switch (v1 % (((v4 << ((2) & 63))) | 1)) {
        case 32768:
            v3 = int64(v2)
        case 42:
            v15 := []int{v1, v1, v1}
            _ = v15
        case 9223372036854775807:
            v4 -= (v4 - 127)
        default:
            for fi16 := 0; fi16 < 4; fi16++ {
                fmt.Println(v3, v4, v4)
                var v17 bool = false
                _ = fi16
                _ = v17
            }
        }
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 string = "hello"
            v5 += v5
            _ = fi18
            _ = v19
        }
        v20 := GenS1{F0: 0, F1: (v2 / (((v2 | 200)) | 1))}
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
