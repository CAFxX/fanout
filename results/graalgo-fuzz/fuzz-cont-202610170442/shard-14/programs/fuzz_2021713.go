package main

import (
    "fmt"
)


func main() {
    var v1 uint8 = 1
    var v2 int32 = -2147483648
    v3 := []uint32{692, 65535, 1, 42}
    v4 := make(map[int]string)
    var v5 uint64 = 18446744073709551615
    var v6 bool = true
    var v7 float64 = 3.14159
    var v8 int = -756
    v1 = (v1 * 255)
    switch (("hello" + "hello") + "") {
    case "\n\t":
        if (65535 < 4294967295) {
            var v9 string = "hello"
            _ = v9
        } else {
            v2 = v2
            v5 -= (^(v5))
        }
    case "hello":
        if v6 {
            v6 = (float64((0.1 + 3.14159)) > (v7 / ((-2.25 * 3.14159) / 1e10)))
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 uint16 = 65535
                var v12 int32 = 42
                if ((0.0 / (1e10 + (((0.0 / (float32(v7) + -2.25)) / 1.0) + (float32(v7) + 0.0)))) >= (-2.25 - 0.0)) {
                    _ = len(v3)
                } else {
                    delete(v4, 32767)
                }
                _ = fi10
                _ = v11
                _ = v12
            }
        } else {
            var v13 float64 = v7
            v7 = (v13 * v7)
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint = 65535
        for fi16 := 0; fi16 < 3; fi16++ {
            if v6 {
                var v17 uint64 = v5
                var v18 int = 4294967295
                v6 = (65535 != 909)
                var v19 uint32 = 0
                _ = v17
                _ = v18
                _ = v19
            }
            if v6 {
                v2 -= v2
                v2 += (v2 & (v2 ^ (v2 & -662)))
            } else {
                var v20 uint = uint(v1)
                v5 -= 2
                _ = v20
            }
            if true {
                v5 = (^(v5))
                switch "hello" {
                case "foo bar":
                    if (!((2147483647 > 42) && (-1 < -448))) {
                        delete(v4, fi16)
                    } else {
                        var v21 int32 = (v2 << (((^(v15))) & 63))
                        fmt.Println(v21)
                        _ = v21
                    }
                case "a":
                    v4[fi16] = ("" + ("a" + ""))
                case "hello":
                    fmt.Println(v15, v15)
                }
            } else {
                fmt.Println(v2)
                var v22 bool = ((310 <= int16(fi14)) || (76 < 42))
                _ = v22
            }
            _ = fi16
        }
        v2 = (v2 >> ((uint(v1)) & 63))
        _ = fi14
        _ = v15
    }
    var v23 int16 = -32768
    v3 = append(v3, 574)
    switch (v8 | v8) {
    case -819:
        fmt.Println(v23, v6)
    default:
        fmt.Println(v7, v6, v1)
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
