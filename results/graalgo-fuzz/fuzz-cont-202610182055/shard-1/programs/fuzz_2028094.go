package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int32 }

func main() {
    var v1 uint32 = 1
    var v2 bool = (false && ((v1 << ((65535) & 63)) != v1))
    var v3 uint32 = (v1 + (v1 >> ((20) & 63)))
    var v4 uint64 = uint64(v1)
    v5 := []uint32{2147483648, (v1 - 2), v1}
    v6 := []uint32{v1, (v1 % ((v3) | 1)), 0}
    fmt.Println(v2)
    v1 = 4294967295
    var v7 uint64 = 42
    fmt.Println(v4, v3, v2)
    var v8 bool = ((126 < 1) || true)
    var v9 int64 = -1
    if (!(1 >= -32768)) {
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := GenS1{F0: -21, F1: v4}
            for fi12 := 0; fi12 < 17; fi12++ {
                fmt.Println(v1)
                _ = fi12
            }
            switch ("" + (("hello" + ((("" + ("a" + (("a" + "a") + ("\n\t" + "")))) + "a") + ("hello" + "foo bar"))) + ("" + (("a" + ("hello" + ("hello" + ""))) + "a")))) {
            case "":
                var v13 float64 = ((float64(((float32(1.5) - 1e10) / (0.1 / 1.5))) / 1.0) - (float64(-2.25) / -2.25))
                _ = v13
            case "hello":
                if len(v6) > 0 {
                    v6[0] = v3
                }
            }
            _ = fi10
            _ = v11
        }
        switch -1 {
        case 203:
            for fi14 := 0; fi14 < 8; fi14++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v15 int16 = 1
                _ = fi14
                _ = v15
            }
        case 383:
            for fi16 := 0; fi16 < 4; fi16++ {
                v17 := GenS1{F0: -128, F1: 18446744073709551615}
                _ = fi16
                _ = v17
            }
        }
    } else {
        _ = len(v6)
        var v18 bool = v8
        if v18 {
            var v19 int8 = 100
            _ = v19
        }
        _ = v18
    }
    v20 := GenS1{F0: -128, F1: (v7 & (v4 / ((v4) | 1)))}
    var v21 uint32 = 65535
    if false {
        var v22 uint = 317
        v2 = v2
        _ = v22
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        var v24 uint8 = 255
        v4 -= v7
        v21 += v21
        _ = fi23
        _ = v24
    }
    v25 := GenS2{F0: 161, F1: -312}
    v8 = (0 <= 255)
    var v26 string = ("\n\t" + "\n\t")
    switch -519 {
    case 0:
        if v8 {
            v21 -= v3
        } else {
            var v27 int8 = 83
            fmt.Println(v9)
            if true {
                if (v9 < (v9 + v9)) {
                    var v28 int32 = -549
                    var v29 bool = false
                    v2 = v2
                    _ = v28
                    _ = v29
                }
                v9 = -9223372036854775808
            }
            _ = v27
        }
    case 32767:
        var v30 int = -9223372036854775808
        _ = v30
    default:
        for fi31 := 0; fi31 < 13; fi31++ {
            fmt.Println(v26)
            v3 += v3
            _ = fi31
        }
    }
    v2 = (v8 && false)
    v32 := GenS1{F0: 0, F1: uint64(v21)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
