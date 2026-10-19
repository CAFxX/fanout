package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int64; F1 int8 }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = -9223372036854775808
    var v3 float64 = float64((1.5 / 3.14159))
    v4 := make(map[uint32]int)
    var v5 string = "\n\t"
    v6 := v1[(v5 + "hello")]
    v7 := []int{}
    switch ((v6 + v6) + (("a" + v5) + "\n\t")) {
    case "\n\t":
        v5 = (v6 + v6)
    case "a":
        var v8 string = v5
        _ = v8
    }
    v5 = "\n\t"
    if (3.14159 < (1.5 + 0.1)) {
        var v9 int32 = 0
        _ = v9
    } else {
        v3 -= (float64(3.14159) - -2.25)
        var v10 uint16 = 0
        if (83 > 131) {
            v6 = ("a" + v5)
            fmt.Println(v5, v3)
            v11 := GenS1{F0: 127, F1: -32768}
            _ = v11
        }
        _ = v10
    }
    v3 = v3
    var v12 uint16 = 32767
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if (float32(v3) >= 1.5) {
        var v13 float32 = -2.25
        var v14 int64 = v2
        _ = v13
        _ = v14
    } else {
        v3 -= (3.14159 / 1e10)
        switch 0 {
        case 1:
            v6 = v6
        }
        var v15 float32 = ((0.0 - 1e10) * 1.5)
        _ = v15
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        v17 := GenS1{F0: 0, F1: 42}
        _ = fi16
        _ = v17
    }
    var v18 int64 = v2
    v19 := GenS1{F0: 100, F1: -32768}
    if (-262 == 254) {
        v20 := GenS2{F0: 1, F1: (v5 + "a")}
        var v21 float64 = v3
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 string = v5
            fmt.Println(v2)
            _ = fi22
            _ = v23
        }
        _ = v20
        _ = v21
    } else {
        v12 = v12
        for fi24 := 0; fi24 < 18; fi24++ {
            fmt.Println(v3, v3)
            _ = fi24
        }
        for fi25 := 0; fi25 < 3; fi25++ {
            v5 = ""
            v2 = v18
            _ = fi25
        }
    }
    switch -1 {
    case -9223372036854775808:
        v26 := GenS2{F0: -118, F1: v6}
        _ = v26
    case 128:
        switch v5 {
        case "hello":
            fmt.Println(v12, v3, v12)
        case "":
            fmt.Println(v12)
        default:
            var v27 int8 = 127
            _ = v27
        }
    case 32768:
        var v28 float64 = ((((v3 / v3) * 0.0) + v3) + v3)
        _ = v28
    default:
        switch "foo bar" {
        case "\n\t":
            v29 := v4[uint32(v12)]
            _ = v29
        case "hello":
            v30 := GenS3{F0: v18, F1: -27}
            _ = v30
        case "a":
            var v31 float32 = 3.14159
            _ = v31
        default:
            v7 = append(v7, -2147483648)
        }
    }
    v18 = (v2 + 9223372036854775807)
    v3 = (3.14159 * -2.25)
    if false {
        switch 921 {
        case -313:
            switch ("foo bar" + (v5 + v5)) {
            case "hello":
                var v32 float32 = 1e10
                _ = v32
            default:
                for fi33 := 0; fi33 < 13; fi33++ {
                    delete(v1, v5)
                    v12 = (v12 | 32767)
                    v34 := GenS3{F0: v2, F1: 42}
                    _ = fi33
                    _ = v34
                }
            }
        case 1:
            v35 := GenS3{F0: (v18 % ((v18) | 1)), F1: 100}
            _ = v35
        default:
            fmt.Println(v6, v3, v12)
        }
        if true {
            fmt.Println(v3, v18, v12)
            if false {
                delete(v4, 1)
            } else {
                fmt.Println(v5, v3, v6)
                if ((0 != -491) || false) {
                    v6 = (v6 + v5)
                    fmt.Println(v2, v3, v3)
                }
                var v36 int = -1
                _ = v36
            }
            var v37 float64 = v3
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 1; fi38++ {
        var v39 uint64 = 4294967295
        if (313 == 344) {
            _ = len(v1)
        } else {
            v12 += (v12 & v12)
            var v40 int32 = 42
            v3 = (v3 / v3)
            _ = v40
        }
        fmt.Println(v12, v39, v3)
        _ = fi38
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
