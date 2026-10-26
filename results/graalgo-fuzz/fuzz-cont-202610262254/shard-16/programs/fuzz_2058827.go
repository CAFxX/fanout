package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint16) int {
    v5 := GenS1{F0: 42, F1: 459, F2: 2147483647}
    v6 := GenS1{F0: 907, F1: 4294967295, F2: 1}
    _ = p0
    _ = v5
    _ = v6
    return 2147483647
}

func main() {
    v1 := []uint32{2147483648, 0, 2, 42}
    var v2 int64 = 1
    var v3 float64 = 1e10
    var v4 int64 = v2
    switch 1 {
    case 128:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 1:
        for fi7 := 0; fi7 < 18; fi7++ {
            fmt.Println(v4)
            _ = fi7
        }
    case -9223372036854775808:
        var v8 string = ""
        _ = v8
    }
    var v9 bool = (genFn0(882) <= 42)
    if (2 < 0) {
        v1 = append(v1, 2147483647)
        switch -1 {
        case 42:
            v3 = (v3 - -2.25)
        case 9223372036854775807:
            v10 := GenS1{F0: 42, F1: 18446744073709551615, F2: -577}
            _ = v10
        case -1:
            _ = len(v1)
        }
        var v11 uint16 = 223
        _ = v11
    }
    v4 -= v4
    var v12 uint = 0
    if len(v1) > 0 {
        v1[0] = 2147483648
    }
    switch "\n\t" {
    case "hello":
        for fi13 := 0; fi13 < 12; fi13++ {
            switch "a" {
            case "foo bar":
                var v14 uint64 = 767
                _ = v14
            case "hello":
                for fi15 := 0; fi15 < 18; fi15++ {
                    v16 := GenS1{F0: 195, F1: 1, F2: -1}
                    var v17 float64 = v3
                    switch "foo bar" {
                    case "foo bar":
                        switch (fi13 & fi13) {
                        case 2147483647:
                            var v18 int = genFn0(553)
                            _ = v18
                        case 1:
                            switch (((("hello" + (("" + "hello") + "\n\t")) + "\n\t") + "\n\t") + "\n\t") {
                            case "hello":
                                v1 = append(v1, 2147483648)
                            }
                        }
                    case "a":
                        var v19 uint32 = 2147483648
                        _ = v19
                    case "\n\t":
                        v1 = append(v1, 1)
                    }
                    v12 += (^(v12))
                    var v20 string = ""
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v20
                }
            case "a":
                var v21 int32 = -20
                _ = v21
            }
            for fi22 := 0; fi22 < 18; fi22++ {
                v23 := GenS2{F0: 18446744073709551615}
                if (34 <= 255) {
                    fmt.Println(fi13, v3, v4)
                    var v24 uint16 = 255
                    switch fi22 {
                    case 42:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    _ = v24
                } else {
                    if v9 {
                        v25 := GenS1{F0: 2147483647, F1: 0, F2: 1}
                        if (453 != 4294967295) {
                            var v26 int8 = 79
                            var v27 uint8 = 0
                            _ = v26
                            _ = v27
                        } else {
                            fmt.Println(v4, fi13, v4)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        _ = v25
                    }
                    var v28 float64 = float64(3.14159)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(fi13)
                    fmt.Println(fi13)
                    _ = v28
                }
                _ = fi22
                _ = v23
            }
            _ = fi13
        }
    case "foo bar":
        v2 = -93
    default:
        var v29 uint16 = 536
        _ = v29
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v30 bool = v9
    if v9 {
        fmt.Println(v30)
        var v31 int = 128
        v32 := GenS2{F0: (v12 | (v12 | v12))}
        _ = v31
        _ = v32
    }
    fmt.Println(v9)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v33 uint16 = 1
    v30 = v9
    v34 := GenS1{F0: 1, F1: 0, F2: -979}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v4
    _ = v9
}
