package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 uint }

func main() {
    v1 := []string{}
    var v2 int64 = -9223372036854775808
    var v3 string = ("" + "foo bar")
    v4 := []string{"", v3, "", ("foo bar" + (v3 + v3))}
    if (1e10 < float64(float32(-2.25))) {
        v5 := GenS1{F0: v2}
        var v6 uint64 = 95
        _ = v5
        _ = v6
    } else {
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v7 string = "hello"
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 int32 = -2147483648
                _ = fi8
                _ = v9
            }
            _ = v7
        }
        if false {
            switch -487 {
            case 289:
                var v10 uint16 = 32767
                _ = v10
            case -9223372036854775808:
                v4 = append(v4, v3)
            case 42:
                v3 = v3
            }
            var v11 uint = 51
            _ = v11
        }
        var v12 uint64 = 42
        _ = v12
    }
    v3 = ((v3 + v3) + "\n\t")
    _ = len(v4)
    fmt.Println(v2)
    switch 32767 {
    case 9223372036854775807:
        v13 := GenS2{F0: 0, F1: 2, F2: 365}
        _ = v13
    default:
        var v14 float64 = -2.25
        _ = v14
    }
    var v15 float64 = 0.0
    v16 := GenS3{F0: 620}
    for fi17 := 0; fi17 < 5; fi17++ {
        v18 := GenS3{F0: 2147483647}
        switch ((v3 + v3) + ("" + v3)) {
        case "":
            for fi19 := 0; fi19 < 14; fi19++ {
                if true {
                    if (3.14159 != -2.25) {
                        var v20 float32 = 0.1
                        v21 := GenS3{F0: 0}
                        var v22 uint64 = 2
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                } else {
                    var v23 uint16 = 1
                    switch ("a" + ("\n\t" + v3)) {
                    case "foo bar":
                        fmt.Println(v3)
                    case "hello":
                        v24 := GenS1{F0: v2}
                        _ = v24
                    }
                    var v25 int8 = 76
                    _ = v23
                    _ = v25
                }
                v2 = 1
                _ = len(v1)
                _ = fi19
            }
        case "a":
            var v26 float64 = v15
            _ = v26
        case "foo bar":
            var v27 int8 = 1
            _ = v27
        default:
            fmt.Println(v2, fi17, v3)
        }
        var v28 uint16 = 255
        _ = fi17
        _ = v18
        _ = v28
    }
    var v29 uint32 = 290
    var v30 uint32 = v29
    var v31 bool = (547 > 0)
    switch v3 {
    case "foo bar":
        fmt.Println(v29, v3, v3)
    case "hello":
        v3 = ((("foo bar" + "foo bar") + v3) + (v3 + v3))
    default:
        var v32 int64 = v2
        _ = v32
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v33 float32 = 0.1
    v34 := GenS1{F0: (v2 / ((v2) | 1))}
    if v31 {
        v35 := GenS4{F0: 100, F1: 4294967295}
        _ = v35
    }
    for fi36 := 0; fi36 < 20; fi36++ {
        fmt.Println(v3)
        fmt.Println(v15, v29, v15)
        var v37 int16 = -343
        _ = fi36
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v4
}
