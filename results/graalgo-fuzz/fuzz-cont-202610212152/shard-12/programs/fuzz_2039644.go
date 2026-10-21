package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint16 }
func genFn0() int32 {
    var v6 float32 = (0.1 - (((0.1 + 0.0) / 0.1) + -2.25))
    switch -9223372036854775808 {
    case -1:
        var v7 uint = 820
        _ = v7
    case -31:
        var v8 int32 = 1
        _ = v8
    }
    _ = v6
    return 479
}

func main() {
    var v1 bool = (-128 != -128)
    var v2 uint = 42
    var v3 bool = true
    var v4 uint32 = 2147483647
    var v5 string = ("a" + ((("\n\t" + "foo bar") + "foo bar") + ("\n\t" + (("hello" + "\n\t") + "\n\t"))))
    var v9 float32 = -2.25
    v5 = ((("" + v5) + ("\n\t" + "hello")) + v5)
    v9 += 0.0
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint8 = 128
        switch fi10 {
        case 4294967295:
            v12 := GenS1{F0: 127, F1: v11, F2: 0}
            _ = v12
        case -1:
            v13 := []int{int(v4)}
            _ = v13
        case 32768:
            fmt.Println(v1, v3, v1)
        }
        _ = fi10
        _ = v11
    }
    if v3 {
        if false {
            v14 := []int{32767}
            switch int(v4) {
            case 9223372036854775807:
                v14 = append(v14, -1)
            default:
                fmt.Println(v3, v3)
            }
            _ = v14
        } else {
            if true {
                if (556 != 1) {
                    v2 -= 616
                    v15 := GenS2{F0: -913, F1: 262}
                    _ = v15
                }
                var v16 uint = (v2 + (^(v2)))
                fmt.Println(v9, v5, v3)
                _ = v16
            } else {
                v3 = ((v4 < (^(v4))) || (!v3))
                v17 := []int{42}
                var v18 uint32 = v4
                _ = v17
                _ = v18
            }
        }
        v19 := GenS2{F0: 107, F1: 65535}
        v4 = 482
        _ = v19
    } else {
        v20 := []int{4294967295, -2147483648, 4294967295}
        _ = v20
    }
    if false {
        var v21 uint64 = 42
        var v22 int8 = 47
        v4 = 42
        _ = v21
        _ = v22
    }
    v23 := []int{32767, -2147483648, 0, -249}
    var v24 uint64 = 1
    var v25 uint32 = v4
    for fi26 := 0; fi26 < 6; fi26++ {
        switch v5 {
        case "":
            var v27 int8 = -128
            _ = v27
        case "\n\t":
            v24 = v24
        default:
            v4 = v4
        }
        _ = fi26
    }
    v24 = uint64(v2)
    var v28 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
