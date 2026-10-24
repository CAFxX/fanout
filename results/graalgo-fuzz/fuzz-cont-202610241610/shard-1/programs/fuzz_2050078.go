package main

import (
    "fmt"
)


func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = -32768
    var v3 int16 = 42
    v4 := []string{"foo bar", "foo bar"}
    var v5 float64 = -2.25
    var v6 string = "hello"
    var v7 string = v6
    var v8 string = v7
    var v9 uint32 = 2147483648
    var v10 int8 = 127
    v7 = ((((v7 + v6) + v8) + "") + (((("\n\t" + (v8 + v7)) + (v6 + v8)) + (v7 + v6)) + (("" + (v8 + v7)) + v7)))
    if false {
        var v11 float32 = (float32(0.1) - 0.1)
        _ = v11
    } else {
        var v12 uint32 = (v9 << ((999) & 63))
        var v13 int = 4294967295
        for fi14 := 0; fi14 < 18; fi14++ {
            fmt.Println(v7, v9)
            _ = fi14
        }
        _ = v12
        _ = v13
    }
    var v15 uint8 = 241
    v4 = append(v4, "\n\t")
    if len(v4) > 0 {
        v4[0] = (v7 + "hello")
    }
    v15 += (v15 - v15)
    for fi16 := 0; fi16 < 11; fi16++ {
        for fi17 := 0; fi17 < 17; fi17++ {
            switch "hello" {
            case "foo bar":
                v4 = append(v4, "\n\t")
            case "hello":
                v5 -= -2.25
            case "a":
                var v18 float64 = 1.5
                _ = v18
            default:
                var v19 uint32 = (v9 / ((uint32(v15)) | 1))
                _ = v19
            }
            switch 1 {
            case 32767:
                var v20 uint = 0
                _ = v20
            case -744:
                v5 = float64(3.14159)
            case 4294967295:
                var v21 int = fi17
                _ = v21
            default:
                var v22 uint = 42
                if (72 > (^(v9))) {
                    if len(v4) > 0 {
                        v4[0] = "a"
                    }
                    fmt.Println(v3, v8, fi17)
                    var v23 uint16 = 0
                    switch int(v15) {
                    case -9223372036854775808:
                        var v24 uint32 = 22
                        _ = v24
                    case -1:
                        v8 = v7
                    case 1:
                        var v25 uint = 0
                        _ = v25
                    }
                    fmt.Println(v3, v3)
                    _ = v23
                } else {
                    v15 = (v15 << ((v22) & 63))
                }
                _ = v22
            }
            _ = fi17
        }
        _ = fi16
    }
    switch 32768 {
    case -399:
        if true {
            fmt.Println(v1, v1)
        }
    case -9223372036854775808:
        var v26 bool = false
        _ = v26
    case -608:
        switch (v7 + ("\n\t" + "\n\t")) {
        case "":
            var v27 int64 = -1
            _ = v27
        case "\n\t":
            switch ((v8 + "") + ("foo bar" + "\n\t")) {
            case "\n\t":
                v7 = ("" + v8)
            case "foo bar":
                v9 = (^(v9))
            case "a":
                v4 = append(v4, v6)
            default:
                fmt.Println(v3, v8, v10)
            }
        case "a":
            v6 = v8
        default:
            switch "hello" {
            case "hello":
                var v28 uint32 = v9
                _ = v28
            }
        }
    }
    v7 = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
