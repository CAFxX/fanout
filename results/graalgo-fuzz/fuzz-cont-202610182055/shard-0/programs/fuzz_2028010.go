package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint }

func main() {
    var v1 uint = 65535
    var v2 int64 = -1
    var v3 uint32 = 2147483648
    var v4 int32 = -1
    switch "\n\t" {
    case "\n\t":
        v1 = 4294967295
    case "a":
        if true {
            var v5 int8 = 127
            v3 -= (v3 * 4294967295)
            v4 = (v4 + -1)
            _ = v5
        } else {
            for fi6 := 0; fi6 < 13; fi6++ {
                var v7 int8 = 1
                _ = fi6
                _ = v7
            }
            v8 := GenS1{F0: 1, F1: (v4 % (((v4 << (((v1 ^ v1)) & 63))) | 1)), F2: 164}
            fmt.Println(v1, v3)
            _ = v8
        }
    case "":
        for fi9 := 0; fi9 < 5; fi9++ {
            switch 128 {
            case 0:
                v2 += v2
            }
            var v10 float64 = 0.1
            _ = fi9
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        if false {
            var v12 uint = 918
            _ = v12
        } else {
            v1 = (v1 | (v1 >> ((v1) & 63)))
            if false {
                v3 -= v3
                if false {
                    if false {
                        for fi13 := 0; fi13 < 8; fi13++ {
                            v14 := GenS1{F0: -1, F1: (v4 | 2147483647), F2: 0}
                            v15 := []int{fi11, 4294967295, fi13}
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                        var v16 int32 = 42
                        _ = v16
                    }
                    var v17 int32 = -526
                    _ = v17
                }
            }
        }
        _ = fi11
    }
    var v18 uint32 = (v3 - (^(v3)))
    v1 += 103
    switch "foo bar" {
    case "hello":
        var v19 uint8 = 77
        _ = v19
    case "foo bar":
        for fi20 := 0; fi20 < 15; fi20++ {
            v4 -= -809
            _ = fi20
        }
    case "a":
        v21 := []int{int(v3), -39, -1}
        _ = v21
    default:
        v2 = (v2 | (v2 ^ v2))
    }
    switch 32767 {
    case -201:
        fmt.Println(v3, v2)
    case -63:
        v22 := GenS2{F0: 4294967295, F1: (v1 - v1)}
        _ = v22
    case 2147483647:
        var v23 string = ("a" + "hello")
        _ = v23
    default:
        for fi24 := 0; fi24 < 17; fi24++ {
            fmt.Println(fi24, v18)
            _ = fi24
        }
    }
    v25 := []int{4294967295, 0, -78, 394}
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
