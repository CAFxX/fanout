package main

import (
    "fmt"
)


func main() {
    var v1 int8 = -128
    v2 := []uint32{0, 2, 428, 4294967295}
    var v3 int64 = -1
    var v4 int = 128
    var v5 float64 = ((0.1 + 0.1) * float64(-2.25))
    var v6 float32 = 3.14159
    if false {
        var v7 uint32 = 42
        _ = v7
    }
    v6 += 0.0
    if ((v1 / ((v1) | 1)) > -1) {
        if true {
            v3 = (v3 - v3)
            switch "foo bar" {
            case "foo bar":
                var v8 int32 = 583
                _ = v8
            }
        } else {
            var v9 int16 = 32767
            for fi10 := 0; fi10 < 10; fi10++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v11 int = (^(fi10))
                _ = fi10
                _ = v11
            }
            v6 -= v6
            _ = v9
        }
        var v12 uint = 445
        v12 = (^(v12))
        _ = v12
    } else {
        for fi13 := 0; fi13 < 2; fi13++ {
            v5 += v5
            _ = fi13
        }
        fmt.Println(v1, v1)
        if ((1 <= 2) && (18446744073709551615 > 4294967295)) {
            var v14 string = ((("a" + ("\n\t" + "")) + ("foo bar" + "foo bar")) + "foo bar")
            for fi15 := 0; fi15 < 14; fi15++ {
                v6 = v6
                var v16 bool = true
                _ = fi15
                _ = v16
            }
            var v17 bool = (!(255 >= 42))
            _ = v14
            _ = v17
        } else {
            var v18 uint = 255
            _ = v18
        }
    }
    var v19 int16 = -1
    var v20 int16 = (v19 >> ((18446744073709551615) & 63))
    var v21 float32 = -2.25
    _ = len(v2)
    for fi22 := 0; fi22 < 6; fi22++ {
        var v23 int16 = v19
        for fi24 := 0; fi24 < 16; fi24++ {
            switch 127 {
            case 296:
                v21 = float32(((((v5 / (-2.25 / float64(0.1))) - 1e10) / 1.0) + float64(v6)))
            case 2147483647:
                var v25 float64 = float64(float32(((1e10 * v5) / 1e10)))
                _ = v25
            case -483:
                v3 = -9223372036854775808
            default:
                var v26 uint64 = 522
                _ = v26
            }
            _ = fi24
        }
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
