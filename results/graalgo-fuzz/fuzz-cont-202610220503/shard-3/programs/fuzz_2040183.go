package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint8 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint = 4294967295
    var v2 int32 = 0
    var v3 int32 = v2
    var v4 uint64 = 42
    var v5 uint32 = 534
    var v6 uint = (v1 | (v1 >> (((v1 % (((v1 + (v1 | v1))) | 1))) & 63)))
    fmt.Println(v5, v6)
    if true {
        v7 := GenS1{F0: 255, F1: 32768, F2: 128}
        v1 += v1
        v8 := []int{-617, 9223372036854775807, 4294967295, 1}
        _ = v7
        _ = v8
    }
    var v9 int16 = 1
    v1 += (v1 & 0)
    v10 := GenS1{F0: 65535, F1: -2147483648, F2: 0}
    v11 := GenS2{F0: 966}
    v3 += int32(v9)
    if true {
        var v12 uint64 = v4
        v13 := []int{-9223372036854775808, 1, -206, 554}
        if (!(!(!false))) {
            switch "foo bar" {
            case "a":
                if (!(true && (true && (false && ((!true) && true))))) {
                    switch "foo bar" {
                    case "\n\t":
                        v6 += (v1 ^ (v1 * (^(v1))))
                    case "foo bar":
                        for fi14 := 0; fi14 < 6; fi14++ {
                            var v15 bool = true
                            var v16 uint8 = 255
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    case "a":
                        v17 := GenS1{F0: 42, F1: -2147483648, F2: 255}
                        _ = v17
                    default:
                        v18 := GenS1{F0: 65535, F1: 431, F2: 128}
                        _ = v18
                    }
                    fmt.Println(v9, v1, v1)
                }
            }
            _ = len(v13)
            if len(v13) > 0 {
                v13[0] = 2147483647
            }
        }
        _ = v12
        _ = v13
    } else {
        for fi19 := 0; fi19 < 5; fi19++ {
            v20 := GenS1{F0: 861, F1: -1, F2: 42}
            v3 -= 1
            _ = fi19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        if (true && false) {
            for fi22 := 0; fi22 < 6; fi22++ {
                fmt.Println(fi21, fi22)
                _ = fi22
            }
            if (false && ("hello" == "\n\t")) {
                fmt.Println(v1)
                var v23 uint8 = uint8(v5)
                var v24 int64 = -9223372036854775808
                _ = v23
                _ = v24
            }
            v3 = v3
        }
        _ = fi21
    }
    v2 = (v2 + (-(v3)))
    fmt.Println(v9)
    v25 := GenS2{F0: 947}
    fmt.Println(v5)
    var v26 uint32 = (v5 + (^(v5)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
