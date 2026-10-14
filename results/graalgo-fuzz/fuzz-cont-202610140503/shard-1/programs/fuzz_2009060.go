package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 uint; F1 uint; F2 int32 }

func main() {
    v1 := []int{}
    var v2 uint = 255
    var v3 int16 = -1
    v4 := make(map[string]bool)
    var v5 float64 = 1e10
    v2 = (^(v2))
    switch "" {
    case "\n\t":
        if false {
            v3 += v3
        } else {
            var v6 int = -2147483648
            v6 -= v6
            _ = v6
        }
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        for fi7 := 0; fi7 < 5; fi7++ {
            v3 = v3
            var v8 bool = false
            _ = fi7
            _ = v8
        }
    }
    fmt.Println(v5, v2)
    v9 := GenS1{F0: -1, F1: -128, F2: 127}
    for fi10 := 0; fi10 < 2; fi10++ {
        _ = len(v4)
        fmt.Println(v3, fi10, v5)
        _ = fi10
    }
    switch "" {
    case "\n\t":
        v11 := GenS2{F0: (v2 & (v2 + 1))}
        _ = v11
    case "":
        _ = len(v1)
    }
    if len(v1) > 0 {
        v1[0] = -945
    }
    if (2147483647 >= 408) {
        var v12 int16 = -791
        _ = v12
    } else {
        v13 := GenS3{F0: -854, F1: (("a" + ("" + "foo bar")) + ("" + ("\n\t" + "")))}
        var v14 uint64 = 42
        _ = v13
        _ = v14
    }
    v2 = 255
    for fi15 := 0; fi15 < 3; fi15++ {
        v3 = v3
        _ = fi15
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 int16 = (v3 & v3)
        v18 := GenS4{F0: v2, F1: (v2 & 255), F2: 1}
        _ = fi16
        _ = v17
        _ = v18
    }
    if (!(!false)) {
        if true {
            v19 := GenS3{F0: -1, F1: "foo bar"}
            _ = v19
        } else {
            v3 -= v3
            if false {
                switch 4294967295 {
                case 1:
                    fmt.Println(v3, v3)
                case 568:
                    var v20 int64 = -1
                    _ = v20
                }
                fmt.Println(v3)
            } else {
                switch "\n\t" {
                case "hello":
                    for fi21 := 0; fi21 < 9; fi21++ {
                        var v22 uint = 42
                        _ = fi21
                        _ = v22
                    }
                case "\n\t":
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                }
                for fi23 := 0; fi23 < 3; fi23++ {
                    v5 -= v5
                    _ = fi23
                }
                if len(v1) > 0 {
                    v1[0] = -906
                }
            }
            var v24 string = ("a" + "")
            _ = v24
        }
    }
    v2 -= v2
    var v25 uint8 = 134
    var v26 int8 = 0
    fmt.Println(v5, v26)
    fmt.Println(v3, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
