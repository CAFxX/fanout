package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint32; F1 string }

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 42
    v3 := make(map[string]uint32)
    v4 := []string{("a" + ""), "", ((("" + "foo bar") + ("" + ("hello" + "hello"))) + "a"), "\n\t"}
    var v5 string = ""
    var v6 int16 = 0
    var v7 int = 42
    if len(v4) > 0 {
        v4[0] = v5
    }
    v6 += v6
    switch (v7 ^ v7) {
    case 550:
        v5 = v5
    case 607:
        for fi8 := 0; fi8 < 4; fi8++ {
            v5 = v5
            _ = fi8
        }
    case -427:
        switch ((v5 + "a") + v5) {
        case "foo bar":
            v4 = append(v4, "\n\t")
        }
    default:
        for fi9 := 0; fi9 < 17; fi9++ {
            if false {
                v2 = 65535
            }
            _ = fi9
        }
    }
    var v10 int32 = 2147483647
    if (((float64(0.0) - 0.1) > 0.1) || (0.0 == ((1.5 - float32(-2.25)) / 1.5))) {
        if len(v4) > 0 {
            v4[0] = v5
        }
    } else {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint64 = 1
            var v13 uint32 = 2147483648
            for fi14 := 0; fi14 < 5; fi14++ {
                fmt.Println(fi14, fi11)
                _ = fi14
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        var v15 float32 = 3.14159
        v7 = (v7 & v7)
        _ = v15
    }
    switch (^(v7)) {
    case -925:
        switch (v5 + (v5 + (v5 + v5))) {
        case "\n\t":
            v16 := GenS1{F0: 68}
            _ = v16
        case "foo bar":
            if (true && true) {
                switch v5 {
                case "\n\t":
                    switch 418 {
                    case 252:
                        v1[-862] = ((!false) && true)
                    case -540:
                        _ = len(v4)
                    case -426:
                        if len(v4) > 0 {
                            v4[0] = "a"
                        }
                    default:
                        var v17 uint16 = v2
                        _ = v17
                    }
                case "a":
                    v6 = -32768
                case "foo bar":
                    v18 := GenS1{F0: 255}
                    _ = v18
                }
                delete(v1, v7)
            }
        case "":
            v6 = (v6 ^ (v6 << ((255) & 63)))
        default:
            if (255 > 255) {
                v19 := GenS2{F0: v2}
                if true {
                    v7 -= (v7 << ((65535) & 63))
                } else {
                    v20 := GenS3{F0: 18446744073709551615, F1: -782}
                    fmt.Println(v2, v7)
                    v21 := GenS3{F0: 101, F1: -9223372036854775808}
                    _ = v20
                    _ = v21
                }
                v5 = (v5 + v5)
                _ = v19
            } else {
                v5 = v5
            }
        }
    case -826:
        v22 := GenS4{F0: 2147483648, F1: v5}
        _ = v22
    }
    fmt.Println(v5, v5)
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
