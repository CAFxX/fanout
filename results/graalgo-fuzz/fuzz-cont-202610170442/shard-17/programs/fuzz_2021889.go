package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int8; F2 string }

func main() {
    var v1 bool = false
    var v2 int8 = 127
    var v3 int8 = 100
    var v4 int32 = 42
    v5 := []uint32{65535, 1, 2147483647, 0}
    var v6 int64 = 42
    var v7 int8 = v3
    v8 := GenS1{F0: (4294967295 == 921)}
    if (0 > 895) {
        v1 = ((("" != ("\n\t" + ("foo bar" + "a"))) || false) || (9223372036854775807 != (v6 >> ((255) & 63))))
        if (0 == 2) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            fmt.Println(v3, v1)
        }
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint8 = 42
        var v11 int = 1
        _ = fi9
        _ = v10
        _ = v11
    }
    if (485 <= 145) {
        v12 := GenS2{F0: (!v1), F1: v7, F2: "hello"}
        v13 := GenS1{F0: (((float64(((1.5 + float32(-2.25)) * (((1e10 / 1e10) + 3.14159) * 1.5))) == (-2.25 - 1.5)) && v1) && ((0 < 18446744073709551615) || true))}
        _ = v12
        _ = v13
    } else {
        fmt.Println(v7)
        for fi14 := 0; fi14 < 12; fi14++ {
            v5 = append(v5, 42)
            _ = fi14
        }
    }
    v4 = (v4 + (v4 << ((979) & 63)))
    v2 += (v7 & v2)
    _ = len(v5)
    var v15 int64 = (v6 * v6)
    v2 = (-(v7))
    fmt.Println(v7, v6, v3)
    fmt.Println(v4, v7, v15)
    for fi16 := 0; fi16 < 4; fi16++ {
        if (255 >= 40) {
            if v1 {
                var v17 bool = true
                var v18 float64 = 0.1
                _ = v17
                _ = v18
            } else {
                var v19 int = 0
                switch "foo bar" {
                case "a":
                    v5 = append(v5, 42)
                default:
                    var v20 int64 = -1
                    _ = v20
                }
                _ = v19
            }
        }
        var v21 int64 = 1
        _ = fi16
        _ = v21
    }
    v1 = true
    for fi22 := 0; fi22 < 16; fi22++ {
        for fi23 := 0; fi23 < 3; fi23++ {
            if v1 {
                fmt.Println(fi23)
                var v24 int32 = v4
                switch (fi23 / (((fi22 + fi22)) | 1)) {
                case -1:
                    v5 = append(v5, 65535)
                }
                _ = len(v5)
                _ = v24
            }
            v3 = v7
            _ = fi23
        }
        switch "a" {
        case "hello":
            for fi25 := 0; fi25 < 4; fi25++ {
                var v26 int32 = (v4 % ((v4) | 1))
                v27 := GenS1{F0: true}
                _ = fi25
                _ = v26
                _ = v27
            }
        case "a":
            var v28 uint8 = 204
            _ = v28
        case "foo bar":
            fmt.Println(fi22, v1)
        default:
            switch (^(fi22)) {
            case -869:
                for fi29 := 0; fi29 < 16; fi29++ {
                    v5 = append(v5, 882)
                    v7 += (v7 / ((v3) | 1))
                    var v30 uint32 = 4294967295
                    _ = fi29
                    _ = v30
                }
            }
        }
        v7 = v7
        _ = fi22
    }
    _ = len(v5)
    if false {
        v15 -= 0
    }
    if v1 {
        if (v15 != (^(v15))) {
            switch ((((("\n\t" + ("\n\t" + ("\n\t" + ("" + "a")))) + "a") + "\n\t") + "a") + "") {
            case "":
                switch ("" + ("\n\t" + "\n\t")) {
                case "\n\t":
                    var v31 uint64 = 42
                    _ = v31
                case "a":
                    v15 = v6
                default:
                    v6 = v6
                }
            }
            v6 += -1
        } else {
            var v32 int32 = (^(v4))
            v15 = (v15 << ((0) & 63))
            _ = v32
        }
    } else {
        fmt.Println(v15)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
