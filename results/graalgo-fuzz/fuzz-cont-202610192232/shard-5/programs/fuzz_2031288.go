package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint8; F2 int8 }

func main() {
    var v1 uint8 = 249
    v2 := []int{9223372036854775807, 32768, -415, -9223372036854775808}
    var v3 string = ""
    switch ("hello" + "a") {
    case "foo bar":
        var v4 int = 492
        _ = v4
    case "hello":
        switch v3 {
        case "hello":
            var v5 uint16 = 425
            _ = v5
        case "a":
            for fi6 := 0; fi6 < 16; fi6++ {
                for fi7 := 0; fi7 < 5; fi7++ {
                    switch "foo bar" {
                    case "a":
                        var v8 uint32 = 65535
                        _ = v8
                    case "\n\t":
                        v2 = append(v2, 42)
                    }
                    if len(v2) > 0 {
                        v2[0] = (fi7 >> ((2) & 63))
                    }
                    _ = fi7
                }
                _ = fi6
            }
        default:
            v3 = "\n\t"
        }
    default:
        v3 = (v3 + v3)
    }
    switch (v3 + (v3 + v3)) {
    case "hello":
        v1 = 255
    case "\n\t":
        v1 += (v1 ^ (v1 + (v1 & (^(v1)))))
    case "":
        for fi9 := 0; fi9 < 4; fi9++ {
            if len(v2) > 0 {
                v2[0] = (fi9 * (fi9 / (((fi9 / (((fi9 - fi9)) | 1))) | 1)))
            }
            var v10 int16 = 1000
            v1 += v1
            _ = fi9
            _ = v10
        }
    }
    var v11 int = -9223372036854775808
    var v12 float32 = 1e10
    v1 = (v1 / ((v1) | 1))
    v3 = (v3 + "")
    if true {
        var v13 uint = 2
        _ = v13
    } else {
        for fi14 := 0; fi14 < 16; fi14++ {
            for fi15 := 0; fi15 < 14; fi15++ {
                fmt.Println(v1)
                _ = fi15
            }
            var v16 int64 = 0
            _ = fi14
            _ = v16
        }
    }
    switch v11 {
    case 0:
        v17 := GenS1{F0: 257}
        _ = v17
    case 4294967295:
        switch (v11 * (^(v11))) {
        case 32767:
            for fi18 := 0; fi18 < 18; fi18++ {
                if (!(!(!true))) {
                    v19 := GenS1{F0: 351}
                    _ = v19
                }
                var v20 bool = (158 == (v1 & v1))
                var v21 uint8 = (^(v1))
                _ = fi18
                _ = v20
                _ = v21
            }
        case -9223372036854775808:
            var v22 float32 = ((float32(0.1) - 3.14159) - (float32(1.5) / v12))
            _ = v22
        case 42:
            v23 := GenS2{F0: 0, F1: (v1 * 80), F2: -17}
            _ = v23
        default:
            switch -9223372036854775808 {
            case 192:
                fmt.Println(v12, v1)
            case -9223372036854775808:
                if len(v2) > 0 {
                    v2[0] = -9223372036854775808
                }
            case 128:
                if len(v2) > 0 {
                    v2[0] = 42
                }
            }
        }
    }
    v24 := GenS2{F0: 255, F1: (^(v1)), F2: 72}
    var v25 int64 = 1
    for fi26 := 0; fi26 < 1; fi26++ {
        var v27 bool = (250 > 325)
        for fi28 := 0; fi28 < 10; fi28++ {
            var v29 bool = (!v27)
            var v30 uint32 = 4294967295
            _ = fi28
            _ = v29
            _ = v30
        }
        v2 = append(v2, (fi26 + -9223372036854775808))
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
