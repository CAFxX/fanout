package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 uint32; F2 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int }
func genFn0(p0 int8) int32 {
    for fi4 := 0; fi4 < 16; fi4++ {
        for fi5 := 0; fi5 < 15; fi5++ {
            v6 := []int{}
            v7 := GenS1{F0: 9223372036854775807, F1: 42, F2: fi4}
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    p0 = p0
    _ = p0
    return 2147483647
}

func main() {
    var v1 uint = 4294967295
    var v2 int = -9223372036854775808
    var v3 uint8 = 42
    v8 := GenS2{F0: ((!(true || (false || (65535 <= 267)))) && (-9223372036854775808 > 9223372036854775807))}
    if (127 < 42) {
        var v9 int32 = 1
        for fi10 := 0; fi10 < 12; fi10++ {
            v2 += v2
            v2 = fi10
            var v11 int8 = 0
            _ = fi10
            _ = v11
        }
        var v12 uint64 = 4294967295
        _ = v12
        _ = v9
    }
    fmt.Println(v3)
    var v13 uint16 = 1
    var v14 string = "hello"
    v13 = 65535
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := GenS3{F0: true, F1: 2147483648, F2: int(v1)}
        v2 = -818
        _ = fi15
        _ = v16
    }
    v2 = (^(v2))
    for fi17 := 0; fi17 < 9; fi17++ {
        v18 := []int{fi17, v2, v2, 9223372036854775807}
        if (58 < 65535) {
            if (v1 != (v1 * v1)) {
                if (2147483648 < 273) {
                    var v19 int32 = 0
                    if true {
                        var v20 uint16 = v13
                        _ = v20
                    }
                    var v21 uint32 = 42
                    _ = v19
                    _ = v21
                } else {
                    switch 32768 {
                    case 338:
                        for fi22 := 0; fi22 < 14; fi22++ {
                            v13 -= uint16(v3)
                            _ = fi22
                        }
                    case 127:
                        v23 := GenS4{F0: v14}
                        _ = v23
                    case -2147483648:
                        switch "hello" {
                        case "a":
                            v24 := GenS4{F0: (v14 + (v14 + v14))}
                            _ = v24
                        default:
                            v14 = "foo bar"
                        }
                    }
                    v13 += v13
                }
                var v25 uint = v1
                v26 := GenS5{F0: (^(fi17))}
                _ = v25
                _ = v26
            } else {
                var v27 string = v14
                _ = v27
            }
            v13 = (v13 ^ v13)
        } else {
            var v28 bool = false
            _ = v28
        }
        _ = fi17
        _ = v18
    }
    v29 := GenS4{F0: v14}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v8
}
