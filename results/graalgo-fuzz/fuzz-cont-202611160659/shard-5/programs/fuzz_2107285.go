package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int16; F2 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    v4 := []int{128, 42, 1}
    switch 420 {
    case 2147483647:
        switch -1 {
        case 127:
            var v5 int = 1
            _ = v5
        case 307:
            var v6 uint32 = 0
            _ = v6
        case 532:
            var v7 uint8 = 0
            _ = v7
        default:
            _ = len(v2)
        }
    case -2147483648:
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := GenS1{F0: 228}
            _ = fi8
            _ = v9
        }
    case 128:
        switch 386 {
        case -423:
            var v10 string = "a"
            _ = v10
        }
    default:
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := GenS1{F0: 503}
            if false {
                _ = len(v4)
                switch 56 {
                case 699:
                    var v13 int = (fi11 - fi11)
                    _ = v13
                }
            } else {
                fmt.Println(fi11)
            }
            v14 := GenS1{F0: 255}
            _ = fi11
            _ = v12
            _ = v14
        }
    }
    var v15 uint64 = 92
    if (1 <= 597) {
        if false {
            for fi16 := 0; fi16 < 5; fi16++ {
                for fi17 := 0; fi17 < 12; fi17++ {
                    v18 := GenS1{F0: 640}
                    v15 += v15
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
        }
        if ((!(567 <= 706)) || true) {
            v15 -= v15
            v19 := GenS1{F0: 1}
            for fi20 := 0; fi20 < 3; fi20++ {
                var v21 float64 = 3.14159
                _ = fi20
                _ = v21
            }
            _ = v19
        }
    }
    if (("a" + ((((("\n\t" + ("foo bar" + "\n\t")) + "a") + "a") + "hello") + ("\n\t" + ""))) == "\n\t") {
        switch "\n\t" {
        case "a":
            delete(v2, ("a" + "foo bar"))
        case "":
            v22 := GenS1{F0: 879}
            _ = v22
        case "\n\t":
            v4 = append(v4, 0)
        default:
            _ = len(v1)
        }
    } else {
        v15 = v15
        v23 := GenS1{F0: 1}
        _ = v23
    }
    var v24 float64 = 1e10
    v25 := GenS1{F0: 18446744073709551615}
    for fi26 := 0; fi26 < 13; fi26++ {
        _ = len(v4)
        if (!(!false)) {
            v15 = (v15 >> ((0) & 63))
            var v27 uint32 = 65535
            _ = v27
        }
        _ = fi26
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        var v29 uint8 = 127
        if len(v4) > 0 {
            v4[0] = -80
        }
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        v31 := GenS2{F0: -127, F1: 229, F2: 128}
        fmt.Println(v24, fi30, v15)
        var v32 int8 = 42
        _ = fi30
        _ = v31
        _ = v32
    }
    v24 = (((v24 - 0.0) * 0.1) * float64((0.0 / (float32(v24) + ((((0.0 * 0.1) * ((3.14159 / 3.14159) + 0.0)) + (1e10 * 0.0)) * (0.1 + float32((1e10 / (v24 * 0.0)))))))))
    delete(v1, 4294967295)
    v15 -= v15
    var v33 string = ("foo bar" + (("foo bar" + (("hello" + ("foo bar" + "")) + "a")) + "foo bar"))
    v24 = v24
    v33 = ("" + ("a" + v33))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v4
}
