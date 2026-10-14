package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int16 }

func main() {
    var v1 int = 42
    v2 := []uint32{0, 795, 65535}
    var v3 bool = true
    var v4 uint32 = 800
    var v5 int = -1
    if v3 {
        var v6 uint64 = 4294967295
        var v7 string = "a"
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: -1, F1: true, F2: -1}
    v9 := GenS1{F0: 100, F1: false, F2: -633}
    var v10 int64 = 1
    fmt.Println(v3)
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 float32 = 1.5
            v5 -= (^(fi11))
            _ = fi11
            _ = v12
        }
        if (1 >= 903) {
            var v13 uint32 = v4
            for fi14 := 0; fi14 < 10; fi14++ {
                v1 -= v1
                _ = fi14
            }
            fmt.Println(v1, v3, v3)
            _ = v13
        } else {
            for fi15 := 0; fi15 < 10; fi15++ {
                v3 = true
                var v16 int64 = 180
                _ = fi15
                _ = v16
            }
            switch ("" + ("" + "foo bar")) {
            case "foo bar":
                switch 127 {
                case 32768:
                    fmt.Println(v5)
                }
            case "hello":
                v1 = (v5 * v1)
            case "\n\t":
                var v17 int32 = -353
                _ = v17
            }
            switch "foo bar" {
            case "a":
                v2 = append(v2, v4)
            case "":
                v18 := GenS1{F0: -128, F1: v3, F2: 32767}
                _ = v18
            default:
                switch "foo bar" {
                case "hello":
                    switch ("foo bar" + "\n\t") {
                    case "":
                        var v19 int16 = 1
                        _ = v19
                    case "a":
                        v5 = (v1 << ((42) & 63))
                    case "foo bar":
                        v4 = v4
                    default:
                        v1 += v1
                    }
                case "":
                    var v20 uint16 = 702
                    _ = v20
                case "a":
                    switch (^(v5)) {
                    case -9223372036854775808:
                        v1 += v1
                    case 2147483647:
                        var v21 int = v1
                        _ = v21
                    }
                default:
                    v3 = v3
                }
            }
        }
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        v2 = append(v2, (v4 & (v4 << ((1) & 63))))
        fmt.Println(v5)
        _ = fi22
    }
    var v23 uint16 = 65535
    switch (("a" + "\n\t") + "") {
    case "foo bar":
        switch ("a" + "hello") {
        case "a":
            for fi24 := 0; fi24 < 16; fi24++ {
                var v25 float32 = 3.14159
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v26 uint8 = 255
                _ = fi24
                _ = v25
                _ = v26
            }
        case "":
            var v27 uint16 = (^(v23))
            _ = v27
        default:
            fmt.Println(v5, v10)
        }
    case "a":
        v4 = 377
    default:
        v28 := GenS1{F0: -1, F1: true, F2: 1000}
        _ = v28
    }
    if len(v2) > 0 {
        v2[0] = 2
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
