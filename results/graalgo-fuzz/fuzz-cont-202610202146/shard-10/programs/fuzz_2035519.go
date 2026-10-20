package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }

func main() {
    v1 := []string{}
    var v2 uint = 0
    v3 := make(map[uint32]string)
    var v4 int32 = 445
    v5 := GenS1{F0: true, F1: 0}
    if false {
        v4 += v4
        var v6 uint = v2
        var v7 float64 = 3.14159
        _ = v6
        _ = v7
    }
    var v8 uint = v2
    fmt.Println(v8)
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 float32 = 1e10
        var v11 uint32 = 4294967295
        var v12 int16 = 0
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    switch -51 {
    case -2147483648:
        v2 = 42
    case 32767:
        delete(v3, 195)
    case 9223372036854775807:
        var v13 uint64 = 775
        _ = v13
    default:
        switch 127 {
        case -2147483648:
            v14 := GenS1{F0: true, F1: 128}
            _ = v14
        }
    }
    if false {
        if (false || (65535 == 0)) {
            v15 := GenS1{F0: true, F1: 127}
            v16 := GenS1{F0: false, F1: 128}
            _ = v15
            _ = v16
        } else {
            for fi17 := 0; fi17 < 12; fi17++ {
                v4 = (v4 % ((2147483647) | 1))
                _ = fi17
            }
            var v18 int32 = v4
            v18 = v4
            _ = v18
        }
    }
    if true {
        switch -401 {
        case 32768:
            _ = len(v3)
        case -2147483648:
            var v19 uint8 = 253
            _ = v19
        case 0:
            for fi20 := 0; fi20 < 10; fi20++ {
                fmt.Println(v4, v4)
                switch ("\n\t" + "\n\t") {
                case "hello":
                    _ = len(v3)
                }
                var v21 uint16 = 65535
                _ = fi20
                _ = v21
            }
        default:
            var v22 int32 = v4
            _ = v22
        }
    } else {
        for fi23 := 0; fi23 < 12; fi23++ {
            var v24 uint64 = 4294967295
            var v25 uint32 = 65535
            switch ("foo bar" + ("\n\t" + ("\n\t" + ((((("" + "\n\t") + "a") + "hello") + "") + "a")))) {
            case "":
                v26 := GenS1{F0: (42 == 65535), F1: 255}
                _ = v26
            case "a":
                v1 = append(v1, "foo bar")
            }
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    if true {
        v4 += (^(v4))
    } else {
        v4 += v4
    }
    var v27 uint8 = 0
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
