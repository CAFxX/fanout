package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint64; F1 uint }
func genFn0(p0 uint8) uint {
    p0 -= 42
    switch 1 {
    case 86:
        fmt.Println(p0)
    case 0:
        for fi4 := 0; fi4 < 9; fi4++ {
            p0 = p0
            fmt.Println(fi4)
            _ = fi4
        }
    case 9223372036854775807:
        var v5 int32 = -1
        _ = v5
    default:
        p0 = (p0 - 127)
    }
    var v6 int64 = -1
    _ = p0
    _ = v6
    return 18446744073709551615
}
func genFn1() uint32 {
    v7 := GenS1{F0: 0}
    var v8 uint64 = 908
    switch "\n\t" {
    case "\n\t":
        if true {
            switch "foo bar" {
            case "foo bar":
                var v9 uint32 = 42
                _ = v9
            case "a":
                v10 := []int{-625}
                _ = v10
            default:
                for fi11 := 0; fi11 < 15; fi11++ {
                    var v12 int64 = 9223372036854775807
                    switch ("" + "\n\t") {
                    case "\n\t":
                        switch (("foo bar" + "") + "hello") {
                        case "\n\t":
                            if true {
                                v13 := GenS2{F0: -81, F1: 42}
                                v14 := GenS1{F0: (v12 * v12)}
                                v12 = v12
                                _ = v13
                                _ = v14
                            }
                        default:
                            var v15 int64 = v12
                            _ = v15
                        }
                    case "hello":
                        var v16 uint32 = 4294967295
                        _ = v16
                    default:
                        switch fi11 {
                        case 2147483647:
                            v17 := []int{(fi11 >> ((0) & 63)), fi11, 1, fi11}
                            _ = v17
                        }
                    }
                    _ = fi11
                    _ = v12
                }
            }
            fmt.Println(v8)
            for fi18 := 0; fi18 < 4; fi18++ {
                switch ("\n\t" + "foo bar") {
                case "a":
                    v8 += v8
                case "foo bar":
                    v19 := []int{128, (fi18 ^ (fi18 ^ 127))}
                    _ = v19
                default:
                    switch 32768 {
                    case 9223372036854775807:
                        fmt.Println(fi18)
                    }
                }
                _ = fi18
            }
        }
    case "":
        var v20 float32 = -2.25
        _ = v20
    default:
        var v21 uint32 = 65535
        _ = v21
    }
    _ = v7
    _ = v8
    return 100
}

func main() {
    v1 := []int{187}
    var v2 int16 = 731
    var v3 int = 127
    if true {
        for fi22 := 0; fi22 < 10; fi22++ {
            v1 = append(v1, 4294967295)
            v23 := GenS1{F0: -834}
            _ = fi22
            _ = v23
        }
        v2 += v2
    }
    v24 := GenS1{F0: -9223372036854775808}
    v3 -= v3
    v2 = v2
    var v25 uint64 = 67
    var v26 int8 = 104
    if true {
        v27 := GenS2{F0: (v26 ^ 0), F1: 42}
        _ = len(v1)
        v28 := GenS3{F0: 76, F1: 4294967295}
        _ = v27
        _ = v28
    }
    var v29 uint32 = 0
    for fi30 := 0; fi30 < 9; fi30++ {
        fmt.Println(v29)
        _ = fi30
    }
    v31 := GenS3{F0: (v25 >> ((752) & 63)), F1: 697}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v31
}
