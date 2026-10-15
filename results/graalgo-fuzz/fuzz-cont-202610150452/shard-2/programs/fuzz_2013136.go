package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int32 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 0
    v3 := []string{}
    v4 := []string{}
    var v5 bool = (!(!(42 > 2147483648)))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v6 bool = (v2 != v2)
    switch 9223372036854775807 {
    case 42:
        v2 -= 50
    case 826:
        if v6 {
            switch 1 {
            case 9223372036854775807:
                fmt.Println(v5)
            case 449:
                for fi7 := 0; fi7 < 17; fi7++ {
                    var v8 bool = (v6 || v5)
                    v2 = (v2 / (((^(v2))) | 1))
                    _ = fi7
                    _ = v8
                }
            }
            switch ("a" + "\n\t") {
            case "\n\t":
                for fi9 := 0; fi9 < 19; fi9++ {
                    var v10 int8 = -13
                    _ = fi9
                    _ = v10
                }
            case "hello":
                fmt.Println(v5, v5, v2)
            }
        } else {
            var v11 uint16 = 1
            _ = v11
        }
    case 1:
        switch -408 {
        case 267:
            v12 := GenS1{F0: -1, F1: 2147483647, F2: -252}
            _ = v12
        case 9223372036854775807:
            var v13 int8 = v2
            _ = v13
        }
    }
    v14 := GenS2{F0: 522, F1: 2147483648}
    var v15 bool = (!v6)
    for fi16 := 0; fi16 < 16; fi16++ {
        switch ((("foo bar" + ("" + "hello")) + "hello") + (((("" + ("" + "foo bar")) + "") + "hello") + "")) {
        case "hello":
            v5 = (18446744073709551615 > 916)
        case "a":
            var v17 uint64 = 562
            _ = v17
        }
        v2 += (-(v2))
        var v18 float32 = ((3.14159 + 0.1) - 3.14159)
        _ = fi16
        _ = v18
    }
    var v19 uint32 = 2147483647
    v20 := v1["foo bar"]
    switch (("" + "foo bar") + "hello") {
    case "hello":
        var v21 float64 = 1e10
        _ = v21
    case "":
        var v22 float64 = (0.0 + ((-2.25 / 3.14159) * 3.14159))
        _ = v22
    case "a":
        switch (("hello" + "") + (("hello" + "a") + "hello")) {
        case "\n\t":
            if false {
                v23 := GenS2{F0: v20, F1: (v19 / ((2147483648) | 1))}
                var v24 uint64 = 244
                _ = v23
                _ = v24
            }
        case "hello":
            switch -9223372036854775808 {
            case 4294967295:
                var v25 int16 = 0
                _ = v25
            case 128:
                var v26 uint16 = 65535
                _ = v26
            case -2147483648:
                for fi27 := 0; fi27 < 19; fi27++ {
                    var v28 bool = v6
                    _ = fi27
                    _ = v28
                }
            default:
                delete(v1, "\n\t")
            }
        case "foo bar":
            v29 := v1[""]
            _ = v29
        }
    default:
        var v30 float64 = (-2.25 - ((((((((1e10 - 0.0) * (-2.25 + 1.5)) + -2.25) / (1.5 - float64(0.0))) + 1e10) - 1.5) - ((3.14159 * 0.0) * 0.1)) + 0.0))
        _ = v30
    }
    v31 := GenS2{F0: (-(v20)), F1: v19}
    fmt.Println(v6, v20, v15)
    v5 = v5
    v32 := GenS3{F0: -1}
    fmt.Println(v6, v6)
    fmt.Println(v6, v15, v2)
    v19 += (^(v19))
    var v33 int = v20
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
