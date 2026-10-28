package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint16 }

func main() {
    var v1 int64 = 1
    var v2 string = "\n\t"
    var v3 uint32 = 42
    var v4 int32 = 0
    switch -448 {
    case -9223372036854775808:
        if false {
            fmt.Println(v4, v3)
        } else {
            for fi5 := 0; fi5 < 2; fi5++ {
                v1 += v1
                _ = fi5
            }
        }
    default:
        switch -2147483648 {
        case 889:
            v6 := []int{42, 0, -9223372036854775808, 32767}
            _ = v6
        case -855:
            v7 := GenS1{F0: 114, F1: 1}
            _ = v7
        case 0:
            switch ("hello" + "foo bar") {
            case "\n\t":
                v8 := GenS1{F0: 0, F1: 248}
                _ = v8
            default:
                switch int(v4) {
                case 4294967295:
                    v9 := []int{}
                    _ = v9
                case 2147483647:
                    v3 = (v3 % (((v3 + (v3 + (^(v3))))) | 1))
                }
            }
        }
    }
    var v10 int32 = 42
    var v11 int8 = -128
    var v12 uint = 4294967295
    for fi13 := 0; fi13 < 17; fi13++ {
        fmt.Println(v11)
        v2 = ("hello" + v2)
        v12 = v12
        _ = fi13
    }
    var v14 uint64 = 42
    v15 := GenS1{F0: 27, F1: 406}
    fmt.Println(v2, v2)
    v16 := []int{4294967295, 1, 1}
    if true {
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 int = (fi17 % ((-670) | 1))
            switch v2 {
            case "\n\t":
                var v19 int64 = v1
                _ = v19
            default:
                var v20 uint16 = 32767
                _ = v20
            }
            _ = fi17
            _ = v18
        }
        fmt.Println(v12)
        var v21 uint8 = 0
        _ = v21
    }
    if (!true) {
        v22 := GenS1{F0: 22, F1: 255}
        var v23 uint16 = 42
        for fi24 := 0; fi24 < 5; fi24++ {
            v3 = (v3 & (v3 * (v3 | (v3 / (((v3 << ((42) & 63))) | 1)))))
            _ = fi24
        }
        _ = v22
        _ = v23
    } else {
        v11 = (v11 ^ v11)
        var v25 float64 = -2.25
        v26 := GenS2{F0: -1, F1: 92}
        _ = v25
        _ = v26
    }
    fmt.Println(v12, v12)
    var v27 float64 = (-2.25 - -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
