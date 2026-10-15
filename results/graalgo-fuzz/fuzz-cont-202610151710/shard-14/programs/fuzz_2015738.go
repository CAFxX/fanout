package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{663, 2147483647}
    var v2 float64 = 1e10
    var v3 uint64 = 1
    var v4 float32 = float32(float64(((1.5 - 0.1) / 1.0)))
    var v5 uint8 = 42
    if true {
        v4 += (v4 * -2.25)
    }
    switch "a" {
    case "":
        if ("a" <= "") {
            if (v5 == (v5 | (v5 % ((128) | 1)))) {
                v1 = append(v1, 389)
            } else {
                v4 += 0.1
            }
        } else {
            for fi6 := 0; fi6 < 2; fi6++ {
                var v7 string = ("\n\t" + ((("" + "hello") + "a") + "foo bar"))
                v2 = ((float64(0.0) + v2) * (float64(v4) / (float64(((3.14159 - ((float32(((v2 - v2) * float64((1e10 + 1.5)))) - v4) - v4)) - (float32(((v2 / (float64(1.5) + 1e10)) - ((((v2 + v2) - v2) + v2) * v2))) * ((v4 + v4) / v4)))) + 1.5)))
                fmt.Println(v3)
                _ = fi6
                _ = v7
            }
            switch ("hello" + ("hello" + ("\n\t" + ("hello" + "foo bar")))) {
            case "hello":
                v3 += 4294967295
            case "foo bar":
                fmt.Println(v4)
            }
            var v8 int64 = 9223372036854775807
            _ = v8
        }
    }
    v4 += ((((float32(-2.25) - (v4 / (((v4 - 3.14159) + v4) - v4))) - v4) * (1.5 * (1e10 - (v4 * float32(3.14159))))) / 3.14159)
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(v2, v3)
        var v10 int8 = 1
        var v11 uint16 = 812
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint16 = 778
    var v13 uint = 4294967295
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 uint32 = 387
        for fi16 := 0; fi16 < 20; fi16++ {
            if ("\n\t" != "hello") {
                switch int(v13) {
                case 2147483647:
                    v5 = (^(v5))
                default:
                    var v17 uint8 = v5
                    _ = v17
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            switch 32767 {
            case 128:
                fmt.Println(fi14, v5, v12)
            default:
                if (65535 > 255) {
                    v1 = append(v1, 2147483647)
                } else {
                    fmt.Println(v5, v5, fi14)
                }
            }
            if len(v1) > 0 {
                v1[0] = (v15 % ((v15) | 1))
            }
            fmt.Println(v12, v13)
            var v18 uint16 = v12
            _ = fi16
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    v2 = 1e10
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
