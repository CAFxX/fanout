package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 string }
type GenS2 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint8 }
func genFn0() int64 {
    v4 := []int{9223372036854775807, 791}
    var v5 float64 = (0.0 * (((-2.25 * 1e10) * 1.5) * ((0.1 + 0.0) / 1e10)))
    v5 += (1e10 / 0.1)
    _ = v4
    _ = v5
    return 42
}
func genFn1() int {
    if (("\n\t" + "\n\t") == "a") {
        v6 := []int{}
        var v7 int16 = 1000
        _ = v6
        _ = v7
    }
    var v8 int = -9223372036854775808
    if (!false) {
        var v9 int16 = 607
        _ = v9
    } else {
        var v10 float64 = 1e10
        var v11 uint16 = 367
        _ = v10
        _ = v11
    }
    switch (v8 / (((v8 | v8)) | 1)) {
    case 42:
        var v12 int16 = -70
        _ = v12
    case 2147483647:
        var v13 uint32 = 1
        _ = v13
    }
    _ = v8
    return v8
}

func main() {
    v1 := []string{}
    var v2 string = (((("foo bar" + ((("foo bar" + ("a" + "a")) + ("" + "")) + "\n\t")) + "foo bar") + ("a" + "hello")) + "\n\t")
    v3 := make(map[string]uint32)
    v14 := GenS1{F0: -2147483648, F1: "", F2: "foo bar"}
    v2 = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        switch v2 {
        case "hello":
            v2 = ("\n\t" + "foo bar")
        default:
            for fi16 := 0; fi16 < 7; fi16++ {
                if true {
                    if (fi15 == 2147483647) {
                        v17 := GenS2{F0: genFn0(), F1: 128, F2: 42}
                        var v18 int = fi16
                        v19 := v3[v2]
                        fmt.Println(fi16, fi15)
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                }
                _ = fi16
            }
        }
        _ = fi15
    }
    fmt.Println(v2)
    switch "\n\t" {
    case "\n\t":
        for fi20 := 0; fi20 < 12; fi20++ {
            v21 := GenS3{F0: 0}
            _ = len(v1)
            _ = fi20
            _ = v21
        }
    case "a":
        fmt.Println(v2)
    default:
        if ((v2 + "foo bar") == "\n\t") {
            v22 := GenS2{F0: 42, F1: 4, F2: 412}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch "" {
            case "\n\t":
                var v23 float64 = 1e10
                _ = v23
            case "hello":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                var v24 float32 = 1.5
                _ = v24
            }
            _ = v22
        } else {
            for fi25 := 0; fi25 < 6; fi25++ {
                switch v2 {
                case "a":
                    var v26 uint = 65535
                    _ = v26
                case "":
                    var v27 int8 = 21
                    _ = v27
                }
                _ = fi25
            }
            if true {
                switch ("a" + (v2 + v2)) {
                case "hello":
                    v2 = "a"
                case "foo bar":
                    v28 := GenS1{F0: 1, F1: v2, F2: ""}
                    _ = v28
                default:
                    if true {
                        fmt.Println(v2)
                        var v29 int8 = 1
                        _ = v29
                    }
                }
                v30 := GenS3{F0: 1}
                _ = v30
            }
        }
    }
    for fi31 := 0; fi31 < 19; fi31++ {
        var v32 uint = 782
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi31
        _ = v32
    }
    var v33 uint = 628
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v33
}
