package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 uint }

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint8 = 127
    var v3 int32 = 0
    v4 := GenS1{F0: 42}
    for fi5 := 0; fi5 < 13; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            if false {
                if (!((((3.14159 * -2.25) - 0.1) / 1.5) >= 1.5)) {
                    if (!(1.5 != 0.1)) {
                        switch (-(fi6)) {
                        case 130:
                            v1 = 194
                        case 949:
                            v1 = 42
                        }
                        fmt.Println(v3, fi6)
                        fmt.Println(v1)
                    } else {
                        v7 := GenS1{F0: 1}
                        v2 = v2
                        _ = v7
                    }
                    v1 -= v1
                    v2 -= (v2 % ((151) | 1))
                    if (2 > 4294967295) {
                        if (!false) {
                            switch "foo bar" {
                            case "foo bar":
                                v8 := []int{835, (-(fi6)), (fi5 | (fi6 << (((^(v1))) & 63)))}
                                _ = v8
                            case "":
                                v9 := GenS2{F0: int(v1), F1: (v1 % (((v1 * v1)) | 1))}
                                _ = v9
                            }
                            if (32767 != 32767) {
                                var v10 int64 = -9223372036854775808
                                _ = v10
                            } else {
                                v3 += v3
                                var v11 float64 = 1.5
                                fmt.Println(v2, v1)
                                _ = v11
                            }
                            v3 = 0
                        }
                        v2 = v2
                    }
                    var v12 float64 = 1e10
                    _ = v12
                } else {
                    v3 -= (v3 + (v3 ^ 2147483647))
                }
            }
            var v13 uint32 = 65535
            _ = fi6
            _ = v13
        }
        switch ("a" + ("foo bar" + "\n\t")) {
        case "hello":
            switch "a" {
            case "foo bar":
                var v14 int64 = int64(v2)
                _ = v14
            case "a":
                var v15 int16 = 11
                _ = v15
            case "hello":
                v16 := GenS2{F0: fi5, F1: (v1 ^ 434)}
                _ = v16
            default:
                v17 := []int{(fi5 & (fi5 + fi5)), (-(fi5))}
                _ = v17
            }
        case "":
            fmt.Println(fi5)
        case "a":
            v18 := GenS1{F0: 42}
            _ = v18
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            var v20 int32 = 910
            _ = fi19
            _ = v20
        }
        _ = fi5
    }
    var v21 uint8 = (v2 * (v2 | 0))
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 string = ""
        if false {
            for fi24 := 0; fi24 < 6; fi24++ {
                v21 += 138
                _ = fi24
            }
            v1 = (v1 << (((v1 % (((^(v1))) | 1))) & 63))
        }
        _ = fi22
        _ = v23
    }
    v25 := []int{}
    var v26 uint16 = 42
    v27 := GenS2{F0: 42, F1: (v1 & (v1 * (v1 ^ v1)))}
    var v28 float32 = 3.14159
    if (1 <= 42) {
        v29 := GenS2{F0: -2147483648, F1: 616}
        var v30 uint32 = 930
        var v31 string = "foo bar"
        _ = v29
        _ = v30
        _ = v31
    }
    var v32 float64 = 1.5
    v21 = v21
    v2 = (v2 << (((^(v1))) & 63))
    v33 := GenS1{F0: -1}
    var v34 uint8 = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
