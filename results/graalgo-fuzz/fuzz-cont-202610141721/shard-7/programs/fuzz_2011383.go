package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint }
type GenS2 struct { F0 uint8; F1 int8; F2 int }
type GenS3 struct { F0 int16 }

func main() {
    var v1 uint64 = 42
    var v2 int64 = 1
    v3 := make(map[string]string)
    for fi4 := 0; fi4 < 20; fi4++ {
        delete(v3, "foo bar")
        v5 := GenS1{F0: "", F1: (("foo bar" + ((("hello" + "foo bar") + (("a" + "hello") + (("hello" + ((((("hello" + "a") + ("" + ("hello" + "hello"))) + "a") + "hello") + "")) + "foo bar"))) + (("hello" + "foo bar") + "foo bar"))) + (("hello" + (("hello" + (((("" + "foo bar") + "\n\t") + "foo bar") + ((("a" + ((("hello" + "hello") + "hello") + "hello")) + "\n\t") + "a"))) + ("" + "a"))) + "a")), F2: 798}
        switch "foo bar" {
        case "\n\t":
            v1 = 18446744073709551615
        case "hello":
            for fi6 := 0; fi6 < 9; fi6++ {
                v7 := GenS1{F0: "\n\t", F1: "hello", F2: 0}
                v8 := GenS1{F0: "a", F1: "hello", F2: 764}
                _ = fi6
                _ = v7
                _ = v8
            }
        case "a":
            var v9 int32 = -2147483648
            _ = v9
        default:
            if ((!(false || false)) && false) {
                fmt.Println(fi4, fi4)
                for fi10 := 0; fi10 < 7; fi10++ {
                    switch "" {
                    case "":
                        if (!(42 <= 42)) {
                            v11 := v3[(("\n\t" + "hello") + (("" + ("hello" + "hello")) + "hello"))]
                            var v12 uint8 = 51
                            _ = v11
                            _ = v12
                        }
                    }
                    var v13 uint64 = 2
                    _ = fi10
                    _ = v13
                }
                v1 += (^(v1))
            }
        }
        _ = fi4
        _ = v5
    }
    var v14 string = ""
    var v15 int32 = 42
    v16 := []int{-2147483648, -1, 128}
    if false {
        v17 := GenS2{F0: 2, F1: 127, F2: -9223372036854775808}
        var v18 bool = ((v2 << ((1) & 63)) >= (v2 << ((647) & 63)))
        fmt.Println(v2)
        _ = v17
        _ = v18
    } else {
        switch v14 {
        case "foo bar":
            v19 := GenS1{F0: "", F1: v14, F2: 2}
            _ = v19
        case "a":
            var v20 int32 = v15
            _ = v20
        default:
            fmt.Println(v14, v14)
        }
        _ = len(v16)
    }
    switch (("hello" + "foo bar") + (v14 + "foo bar")) {
    case "hello":
        fmt.Println(v14, v1)
    case "a":
        switch 127 {
        case -581:
            v2 = (v2 + (v2 / ((42) | 1)))
        case 609:
            for fi21 := 0; fi21 < 2; fi21++ {
                var v22 uint64 = (v1 - v1)
                for fi23 := 0; fi23 < 19; fi23++ {
                    _ = len(v3)
                    v24 := GenS2{F0: 181, F1: -24, F2: 4294967295}
                    v25 := GenS2{F0: 2, F1: -37, F2: fi23}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = fi21
                _ = v22
            }
        case 128:
            for fi26 := 0; fi26 < 16; fi26++ {
                v27 := GenS1{F0: ("hello" + v14), F1: v14, F2: 1}
                _ = fi26
                _ = v27
            }
        }
    case "":
        v15 += (v15 ^ 853)
    default:
        v2 += (v2 ^ (v2 << ((18446744073709551615) & 63)))
    }
    if (318 != (v15 | int32(v2))) {
        v28 := GenS1{F0: v14, F1: v14, F2: 534}
        v29 := GenS3{F0: 253}
        switch -9223372036854775808 {
        case -514:
            v30 := GenS2{F0: 2, F1: 42, F2: -9223372036854775808}
            _ = v30
        case -627:
            for fi31 := 0; fi31 < 20; fi31++ {
                v3["\n\t"] = (v14 + v14)
                var v32 float32 = 0.0
                _ = fi31
                _ = v32
            }
        default:
            v33 := GenS1{F0: (("hello" + "\n\t") + "\n\t"), F1: v14, F2: 65535}
            _ = v33
        }
        _ = v28
        _ = v29
    }
    if (712 > 0) {
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        v34 := GenS3{F0: -1}
        var v35 int16 = -124
        _ = v34
        _ = v35
    }
    var v36 int32 = 517
    fmt.Println(v15, v14, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v36
}
