package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint64 }

func main() {
    var v1 uint8 = 128
    var v2 bool = (255 == (v1 & 214))
    v3 := []string{"hello", "foo bar", "\n\t", "foo bar"}
    var v4 uint64 = 2
    var v5 int64 = 1
    if v2 {
        v6 := GenS1{F0: 419, F1: 153, F2: 32767}
        v3 = append(v3, "\n\t")
        fmt.Println(v2, v4)
        _ = v6
    } else {
        if v2 {
            if ((float64(0.0) > (((0.0 - 0.1) / 1e10) / 1.0)) && (true || v2)) {
                v4 -= v4
                v3 = append(v3, ("" + "foo bar"))
                if (v4 >= (v4 >> ((2) & 63))) {
                    fmt.Println(v5, v2, v4)
                    _ = len(v3)
                }
            }
            var v7 float64 = (3.14159 + 1e10)
            v3 = append(v3, ("" + "a"))
            _ = v7
        }
        if v2 {
            var v8 int = -724
            switch "hello" {
            case "\n\t":
                v5 = v5
            }
            if ((!(-78 < 1)) || v2) {
                switch "" {
                case "\n\t":
                    var v9 int16 = -32768
                    _ = v9
                }
                v4 += v4
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v8 += int(v5)
                var v10 bool = (688 < 65535)
                _ = v10
            }
            _ = v8
        } else {
            _ = len(v3)
            v5 = 42
            var v11 uint64 = uint64(v1)
            _ = v11
        }
        fmt.Println(v1, v2, v2)
    }
    _ = len(v3)
    fmt.Println(v1)
    var v12 bool = (103 == 65535)
    v13 := GenS2{F0: -2147483648, F1: 2}
    if (-1 <= -744) {
        var v14 uint64 = 42
        _ = v14
    }
    var v15 int16 = -32768
    for fi16 := 0; fi16 < 16; fi16++ {
        fmt.Println(v4)
        switch (fi16 - 95) {
        case -747:
            if (65535 == 0) {
                if v2 {
                    for fi17 := 0; fi17 < 19; fi17++ {
                        var v18 bool = false
                        v4 = v4
                        _ = fi17
                        _ = v18
                    }
                }
                var v19 int16 = v15
                _ = v19
            } else {
                v15 = int16(v1)
            }
        case 585:
            var v20 uint16 = 42
            _ = v20
        case 209:
            v4 -= v4
        default:
            switch "foo bar" {
            case "":
                var v21 float64 = 1.5
                _ = v21
            case "hello":
                _ = len(v3)
            case "a":
                var v22 int16 = v15
                _ = v22
            }
        }
        _ = fi16
    }
    var v23 uint64 = v4
    var v24 float32 = (((1.5 - 0.1) + (1e10 - (float32(1e10) / 1.0))) * 1.5)
    var v25 int = 42
    fmt.Println(v24)
    fmt.Println(v1, v15, v1)
    v12 = (188 < 4294967295)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
