package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int32; F1 int16 }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 4; fi4++ {
        if (!false) {
            var v5 int64 = 0
            var v6 bool = true
            _ = v5
            _ = v6
        }
        _ = fi4
    }
    var v7 uint64 = 18446744073709551615
    var v8 uint16 = 32767
    for fi9 := 0; fi9 < 1; fi9++ {
        if (-2147483648 == 0) {
            var v10 uint8 = 56
            var v11 string = ("hello" + "")
            fmt.Println(v8, v7, fi9)
            _ = v10
            _ = v11
        }
        _ = fi9
    }
    switch 2147483647 {
    case 32767:
        fmt.Println(v8)
    case -1:
        v7 += v7
    case 42:
        if ((!(0 != 255)) && true) {
            v12 := []int{873, 4294967295, 9223372036854775807, 962}
            v13 := GenS1{F0: -303, F1: 128}
            _ = v12
            _ = v13
        }
    }
    _ = v7
    _ = v8
    return 128
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint = 18446744073709551615
    var v3 int = 0
    var v14 uint = v2
    switch -570 {
    case 2147483647:
        switch "" {
        case "a":
            switch v3 {
            case -985:
                v15 := GenS2{F0: 517, F1: 0}
                _ = v15
            case -862:
                for fi16 := 0; fi16 < 20; fi16++ {
                    for fi17 := 0; fi17 < 20; fi17++ {
                        v18 := []int{(v3 + (-(fi16))), 32768, 2147483647}
                        var v19 uint32 = 0
                        v19 -= 807
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    v20 := GenS1{F0: 32767, F1: 9}
                    var v21 uint8 = 1
                    _ = fi16
                    _ = v20
                    _ = v21
                }
            }
        case "\n\t":
            var v22 int = (v3 & (v3 >> ((v2) & 63)))
            _ = v22
        case "hello":
            fmt.Println(v1, v1)
        default:
            var v23 int64 = -1
            _ = v23
        }
    default:
        var v24 uint = v2
        _ = v24
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        var v26 float32 = v1
        _ = fi25
        _ = v26
    }
    v3 = -9223372036854775808
    var v27 uint32 = 1
    switch "a" {
    case "":
        var v28 float32 = 1.5
        _ = v28
    }
    var v29 uint64 = 4294967295
    v30 := GenS2{F0: -2147483648, F1: 1}
    var v31 float32 = v1
    v14 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v14
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
