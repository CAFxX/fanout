package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int64; F1 int8; F2 int }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 bool; F1 int16; F2 int8 }

func main() {
    var v1 int = -9223372036854775808
    v2 := []uint32{}
    var v3 uint64 = 128
    var v4 bool = (778 > 18446744073709551615)
    fmt.Println(v1, v4)
    switch v1 {
    case 127:
        v5 := GenS1{F0: "hello", F1: ("" + "")}
        _ = v5
    case 42:
        fmt.Println(v1)
    case 4294967295:
        v1 += v1
    }
    fmt.Println(v1, v1, v3)
    var v6 int = (v1 << ((917) & 63))
    if false {
        v4 = (0 != -37)
        var v7 int16 = 365
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        _ = v7
    } else {
        fmt.Println(v6, v6)
        switch v6 {
        case -122:
            var v8 string = ("\n\t" + "foo bar")
            _ = v8
        }
        if v4 {
            _ = len(v2)
        }
    }
    var v9 string = "\n\t"
    if (-2 >= -123) {
        v9 = (v9 + "")
        v10 := GenS2{F0: 0, F1: 1, F2: v1}
        _ = v10
    }
    v11 := GenS3{F0: 1, F1: -9223372036854775808}
    switch (v9 + (((v9 + v9) + "a") + (((v9 + v9) + "hello") + (("\n\t" + "\n\t") + v9)))) {
    case "foo bar":
        switch v9 {
        case "":
            var v12 bool = v4
            _ = v12
        case "hello":
            for fi13 := 0; fi13 < 17; fi13++ {
                v4 = (fi13 >= 32767)
                _ = fi13
            }
        default:
            var v14 uint = 1
            _ = v14
        }
    case "hello":
        fmt.Println(v6)
    case "a":
        v3 = (^(v3))
    }
    var v15 float32 = ((0.1 / 1.0) / 1.5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v15 += (v15 * v15)
    fmt.Println(v4, v3)
    if (42 > 0) {
        v9 = v9
        v3 = v3
    }
    if true {
        var v16 uint16 = 9
        v2 = append(v2, 42)
        _ = v16
    }
    switch (v1 | v1) {
    case 1:
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 uint8 = 42
            _ = fi17
            _ = v18
        }
    case 0:
        var v19 int16 = 597
        _ = v19
    case 976:
        v20 := GenS3{F0: 18446744073709551615, F1: 0}
        _ = v20
    default:
        for fi21 := 0; fi21 < 5; fi21++ {
            if len(v2) > 0 {
                v2[0] = 18
            }
            switch v6 {
            case -614:
                for fi22 := 0; fi22 < 6; fi22++ {
                    switch "a" {
                    case "":
                        v23 := GenS2{F0: 9223372036854775807, F1: -123, F2: -9223372036854775808}
                        _ = v23
                    }
                    var v24 uint8 = 1
                    _ = fi22
                    _ = v24
                }
            case 4294967295:
                v15 += (v15 / v15)
            case 42:
                switch (v6 % ((4294967295) | 1)) {
                case 478:
                    var v25 int64 = 9223372036854775807
                    _ = v25
                case 4294967295:
                    switch (^(v6)) {
                    case 768:
                        v15 = float32(0.0)
                    case 1:
                        if len(v2) > 0 {
                            v2[0] = 65535
                        }
                    }
                case 32768:
                    v26 := GenS4{F0: (v4 && v4), F1: -316, F2: -1}
                    _ = v26
                }
            }
            _ = fi21
        }
    }
    v4 = (!v4)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
