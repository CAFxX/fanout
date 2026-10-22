package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int }
type GenS2 struct { F0 int; F1 uint8; F2 uint32 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]string)
    var v3 int64 = 0
    var v4 int64 = (v3 << ((v1) & 63))
    v5 := v2[571]
    fmt.Println(v1)
    for fi6 := 0; fi6 < 3; fi6++ {
        if (0 == 1) {
            v7 := []int{2147483647, (fi6 << (((v1 ^ v1)) & 63)), (fi6 % (((fi6 + -173)) | 1)), (fi6 / (((fi6 % ((fi6) | 1))) | 1))}
            _ = v7
        } else {
            var v8 uint16 = 65535
            _ = v8
        }
        for fi9 := 0; fi9 < 8; fi9++ {
            fmt.Println(v4)
            v3 = (-(v4))
            var v10 int32 = 996
            if (3 <= 4294967295) {
                v11 := GenS1{F0: -9223372036854775808, F1: 94, F2: fi9}
                var v12 uint32 = 65535
                _ = v11
                _ = v12
            }
            if false {
                v13 := []int{fi6, 361}
                var v14 uint16 = 32767
                switch v5 {
                case "hello":
                    v15 := GenS1{F0: v3, F1: 2147483648, F2: (fi9 | (fi6 | -381))}
                    _ = v15
                default:
                    v16 := GenS2{F0: 127, F1: 0, F2: 2147483647}
                    _ = v16
                }
                fmt.Println(v5, fi6, v4)
                _ = v13
                _ = v14
            } else {
                v10 = 1
                v17 := GenS1{F0: (v3 >> ((2) & 63)), F1: 0, F2: fi6}
                _ = v17
            }
            _ = fi9
            _ = v10
        }
        _ = fi6
    }
    v5 = ("hello" + (v5 + v5))
    v4 += v3
    var v18 int32 = 0
    switch "a" {
    case "\n\t":
        var v19 uint8 = 0
        _ = v19
    case "a":
        for fi20 := 0; fi20 < 6; fi20++ {
            delete(v2, fi20)
            _ = fi20
        }
    default:
        for fi21 := 0; fi21 < 18; fi21++ {
            if (false && true) {
                v22 := []int{}
                v5 = ""
                if (-43 <= 1) {
                    v4 += 42
                    delete(v2, 128)
                }
                _ = v22
            } else {
                var v23 int16 = int16(fi21)
                _ = v23
            }
            _ = fi21
        }
    }
    v24 := v2[-881]
    for fi25 := 0; fi25 < 5; fi25++ {
        switch fi25 {
        case 4294967295:
            v18 -= (v18 | (v18 << ((42) & 63)))
        }
        var v26 uint = (v1 - (v1 - 42))
        _ = fi25
        _ = v26
    }
    switch "\n\t" {
    case "":
        if false {
            var v27 float32 = 1.5
            _ = v27
        }
    case "a":
        v5 = v5
    }
    var v28 int64 = (v4 - v4)
    if ((417 == 4294967295) || (42 == 0)) {
        v1 = v1
        var v29 int8 = -128
        for fi30 := 0; fi30 < 11; fi30++ {
            v1 = 42
            _ = fi30
        }
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
