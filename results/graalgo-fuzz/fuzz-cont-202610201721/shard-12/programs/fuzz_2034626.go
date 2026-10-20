package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int32; F2 int8 }
type GenS4 struct { F0 uint }
func genFn0(p0 int16, p1 uint, p2 int16) uint64 {
    v4 := []int{344, 643, 334}
    var v5 int16 = p0
    v6 := GenS1{F0: 710, F1: 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int16 = 0
    v2 := []uint32{112, 1, 783, 311}
    var v3 bool = true
    if len(v2) > 0 {
        v2[0] = 0
    }
    fmt.Println(v1, v1)
    switch "foo bar" {
    case "\n\t":
        v3 = (-1 <= -461)
    case "a":
        var v7 int8 = -128
        _ = v7
    }
    if false {
        switch ("" + "a") {
        case "":
            v8 := GenS1{F0: 18446744073709551615, F1: -128}
            _ = v8
        default:
            switch 127 {
            case 1:
                if len(v2) > 0 {
                    v2[0] = 1
                }
            }
        }
        for fi9 := 0; fi9 < 17; fi9++ {
            v3 = v3
            _ = fi9
        }
    } else {
        if v3 {
            v1 -= v1
            v10 := GenS2{F0: (v1 >> ((173) & 63))}
            _ = v10
        }
        if len(v2) > 0 {
            v2[0] = 322
        }
        v11 := GenS1{F0: 399, F1: -90}
        _ = v11
    }
    var v12 uint = 65535
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 int64 = -762
        if (!v3) {
            fmt.Println(v12, v14)
            switch ("\n\t" + "a") {
            case "a":
                v3 = v3
            default:
                fmt.Println(v1)
            }
        }
        fmt.Println(v12, v12)
        _ = fi13
        _ = v14
    }
    v12 -= (v12 + v12)
    var v15 uint32 = 2
    if v3 {
        var v16 int64 = -9223372036854775808
        for fi17 := 0; fi17 < 4; fi17++ {
            v3 = (0.1 != 1e10)
            var v18 string = "\n\t"
            _ = fi17
            _ = v18
        }
        _ = v16
    } else {
        v19 := GenS1{F0: 4294967295, F1: 42}
        var v20 int = 9223372036854775807
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        var v22 uint32 = (v15 - 517)
        for fi23 := 0; fi23 < 1; fi23++ {
            switch (fi23 << ((v12) & 63)) {
            case 630:
                if v3 {
                    var v24 uint = 858
                    switch (^(fi21)) {
                    case 869:
                        var v25 int = 232
                        _ = v25
                    default:
                        v12 -= v24
                    }
                    v24 += v24
                    _ = v24
                } else {
                    switch fi21 {
                    case 2147483647:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case 4294967295:
                        switch (fi23 >> ((v12) & 63)) {
                        case -838:
                            v26 := GenS3{F0: 2, F1: -2147483648, F2: 42}
                            _ = v26
                        }
                    default:
                        v12 = (v12 - v12)
                    }
                    if v3 {
                        if true {
                            if v3 {
                                v12 = v12
                            }
                        }
                    } else {
                        v1 += v1
                    }
                }
            case 444:
                v27 := GenS4{F0: (v12 * 65535)}
                _ = v27
            case -2147483648:
                v28 := GenS4{F0: v12}
                _ = v28
            }
            _ = fi23
        }
        var v29 string = "foo bar"
        _ = fi21
        _ = v22
        _ = v29
    }
    var v30 int32 = -1
    for fi31 := 0; fi31 < 4; fi31++ {
        var v32 uint8 = 0
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi31
        _ = v32
    }
    if len(v2) > 0 {
        v2[0] = 449
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v30
}
