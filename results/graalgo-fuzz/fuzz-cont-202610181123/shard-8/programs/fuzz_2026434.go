package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int }
type GenS2 struct { F0 int8; F1 uint64; F2 bool }

func main() {
    var v1 string = "\n\t"
    var v2 int = -155
    var v3 uint = 284
    fmt.Println(v2, v3, v3)
    switch (v1 + "\n\t") {
    case "foo bar":
        var v4 uint8 = 206
        _ = v4
    default:
        var v5 int8 = int8(v2)
        _ = v5
    }
    var v6 string = (v1 + "a")
    v7 := []int{(v2 % ((-9223372036854775808) | 1))}
    v3 += v3
    v1 = "hello"
    v2 += (^(v2))
    if (255 == 32767) {
        v6 = (("\n\t" + ("hello" + "hello")) + v1)
        switch (v2 << ((537) & 63)) {
        case -2147483648:
            var v8 uint8 = 255
            _ = v8
        }
        if (float64(float32(float64(0.0))) >= float64(3.14159)) {
            var v9 uint64 = 42
            var v10 string = (v1 + v6)
            _ = v10
            _ = v9
        } else {
            v3 = (^(v3))
            var v11 uint8 = 28
            _ = v11
        }
    } else {
        v2 = v2
        if false {
            v12 := GenS1{F0: 0, F1: 2147483648, F2: (v2 * v2)}
            for fi13 := 0; fi13 < 15; fi13++ {
                v2 += (^(fi13))
                _ = fi13
            }
            v2 = 4294967295
            _ = v12
        }
        for fi14 := 0; fi14 < 11; fi14++ {
            v6 = (v1 + v6)
            var v15 bool = false
            v7 = append(v7, (v2 + (-(fi14))))
            _ = fi14
            _ = v15
        }
    }
    switch v2 {
    case 255:
        v16 := GenS2{F0: 127, F1: 539, F2: false}
        _ = v16
    case 128:
        fmt.Println(v1)
    case -2147483648:
        for fi17 := 0; fi17 < 14; fi17++ {
            v6 = ((v1 + (("foo bar" + v6) + (((v1 + ("\n\t" + v6)) + (v6 + "hello")) + ("hello" + v6)))) + v6)
            for fi18 := 0; fi18 < 18; fi18++ {
                v7 = append(v7, fi18)
                var v19 int64 = 0
                fmt.Println(v3, v19, fi18)
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
    default:
        switch -2147483648 {
        case 42:
            for fi20 := 0; fi20 < 13; fi20++ {
                if true {
                    var v21 bool = ((!false) || true)
                    fmt.Println(v3, v1, v3)
                    _ = v21
                }
                if len(v7) > 0 {
                    v7[0] = (fi20 >> ((704) & 63))
                }
                _ = fi20
            }
        default:
            v1 = v1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
