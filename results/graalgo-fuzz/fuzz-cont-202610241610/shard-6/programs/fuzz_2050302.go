package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint32 }

func main() {
    var v1 uint64 = 0
    v2 := []uint32{}
    var v3 uint16 = 342
    var v4 int8 = 94
    var v5 int64 = 9223372036854775807
    if true {
        var v6 float64 = -2.25
        _ = v6
    }
    _ = len(v2)
    var v7 float32 = -2.25
    var v8 uint64 = (^(v1))
    switch 9223372036854775807 {
    case -219:
        fmt.Println(v7)
    case 968:
        for fi9 := 0; fi9 < 1; fi9++ {
            if len(v2) > 0 {
                v2[0] = 465
            }
            for fi10 := 0; fi10 < 14; fi10++ {
                v4 += v4
                _ = fi10
            }
            _ = fi9
        }
    case -1:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        v11 := GenS1{F0: 65535, F1: v5}
        _ = v11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            v14 := GenS2{F0: v8, F1: v8, F2: 1}
            if true {
                switch "\n\t" {
                case "hello":
                    v4 = v4
                default:
                    if true {
                        v15 := GenS1{F0: 0, F1: (-(v5))}
                        var v16 int = int(v8)
                        _ = v15
                        _ = v16
                    }
                }
                var v17 float64 = 1e10
                _ = v17
            }
            _ = fi13
            _ = v14
        }
        v3 = (v3 / ((v3) | 1))
        _ = fi12
    }
    v5 += (v5 - v5)
    fmt.Println(v7, v1)
    switch "\n\t" {
    case "foo bar":
        v7 = 0.0
    default:
        if (-894 < 9223372036854775807) {
            if len(v2) > 0 {
                v2[0] = 1
            }
            v5 = 42
            var v18 uint = 0
            _ = v18
        }
    }
    v8 -= (v1 % ((v8) | 1))
    var v19 int = 32768
    for fi20 := 0; fi20 < 8; fi20++ {
        switch ("\n\t" + (("foo bar" + ("" + "\n\t")) + ("hello" + "\n\t"))) {
        case "hello":
            if len(v2) > 0 {
                v2[0] = 65535
            }
        case "foo bar":
            var v21 int32 = 0
            _ = v21
        default:
            _ = len(v2)
        }
        if (-1 > (v4 + (v4 / ((-128) | 1)))) {
            fmt.Println(v3, v5, v4)
            if false {
                var v22 uint64 = (v1 % ((uint64(v3)) | 1))
                _ = v22
            } else {
                var v23 uint32 = 2147483647
                var v24 int8 = (v4 % (((^(v4))) | 1))
                if true {
                    for fi25 := 0; fi25 < 14; fi25++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi25
                    }
                }
                _ = v23
                _ = v24
            }
            v3 = v3
        }
        var v26 float64 = 3.14159
        _ = fi20
        _ = v26
    }
    if false {
        var v27 uint32 = 2147483648
        _ = v27
    } else {
        v28 := GenS2{F0: 590, F1: (v1 + (^(v1))), F2: 42}
        var v29 int = v19
        _ = v28
        _ = v29
    }
    if (1 == 947) {
        v19 += v19
        if (!true) {
            v3 = v3
        } else {
            var v30 string = "foo bar"
            _ = v30
        }
    }
    switch v19 {
    case 4294967295:
        switch (v19 | v19) {
        case 42:
            fmt.Println(v1, v5)
        default:
            fmt.Println(v7)
        }
    }
    v19 += int(v3)
    if (!true) {
        var v31 int64 = v5
        _ = v31
    }
    if ((((false && (v4 >= v4)) && (true && false)) && (1.5 <= float64((float32(3.14159) / (((v7 - float32(((3.14159 / 1.5) * float64(v7)))) + (v7 / v7)) * v7))))) && false) {
        v5 -= 126
    } else {
        var v32 uint16 = (v3 - v3)
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
