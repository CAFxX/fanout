package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int64 }
func genFn0(p0 uint) int {
    for fi4 := 0; fi4 < 9; fi4++ {
        p0 = 255
        fmt.Println(fi4)
        for fi5 := 0; fi5 < 2; fi5++ {
            if ((((false || true) || false) || (!false)) || (255 < 128)) {
                v6 := []int{(fi5 / (((fi5 | 127)) | 1)), fi5, (-(fi5)), fi5}
                p0 += 1
                fmt.Println(fi5, fi5, fi5)
                _ = v6
            }
            switch ("hello" + "\n\t") {
            case "\n\t":
                p0 -= 42
            default:
                p0 -= (^(p0))
            }
            _ = fi5
        }
        _ = fi4
    }
    v7 := []int{-2147483648, -527, -943, -2147483648}
    if (351 < 623) {
        var v8 uint64 = 42
        _ = v8
    }
    if len(v7) > 0 {
        v7[0] = 127
    }
    _ = p0
    _ = v7
    return 128
}

func main() {
    var v1 uint64 = 0
    var v2 uint32 = 65535
    var v3 float64 = 0.0
    for fi9 := 0; fi9 < 2; fi9++ {
        v1 += (v1 & 61)
        var v10 int64 = 0
        _ = fi9
        _ = v10
    }
    v2 = 65535
    var v11 uint8 = 42
    for fi12 := 0; fi12 < 1; fi12++ {
        switch ("a" + "\n\t") {
        case "a":
            fmt.Println(v2, v2)
        case "foo bar":
            switch "" {
            case "":
                for fi13 := 0; fi13 < 18; fi13++ {
                    v11 += v11
                    var v14 float64 = v3
                    fmt.Println(fi12, v11)
                    var v15 uint64 = (^(v1))
                    var v16 uint8 = 184
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
            case "hello":
                var v17 bool = ((float32(((v3 - ((0.0 - (v3 * 3.14159)) * 0.1)) + v3)) != 3.14159) || ((((("foo bar" + "\n\t") + "\n\t") + ("\n\t" + ("" + "hello"))) <= "foo bar") && true))
                _ = v17
            default:
                var v18 uint64 = (v1 / ((v1) | 1))
                _ = v18
            }
        case "hello":
            for fi19 := 0; fi19 < 10; fi19++ {
                v11 = v11
                fmt.Println(v2, v3, v2)
                _ = fi19
            }
        }
        v1 = v1
        v2 = (v2 - (^(v2)))
        _ = fi12
    }
    if (true && false) {
        var v20 int8 = -10
        var v21 int8 = int8(v1)
        switch 2147483647 {
        case -341:
            var v22 uint = 639
            _ = v22
        case 2147483647:
            v3 -= v3
        }
        _ = v20
        _ = v21
    } else {
        v23 := GenS1{F0: -836, F1: -32768, F2: 42}
        _ = v23
    }
    v3 = (0.1 - (v3 + float64((((float32(v3) - -2.25) - ((1.5 * (0.0 - (-2.25 * float32(v3)))) * float32(v3))) / 0.1))))
    var v24 bool = (-473 != -701)
    v2 = v2
    var v25 uint32 = v2
    switch "a" {
    case "":
        if v24 {
            fmt.Println(v25)
            var v26 uint64 = (v1 & v1)
            v1 -= v1
            _ = v26
        } else {
            var v27 bool = v24
            var v28 int16 = 1
            _ = v27
            _ = v28
        }
    case "hello":
        v2 = 794
    case "foo bar":
        var v29 uint16 = 1
        _ = v29
    default:
        v25 = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
