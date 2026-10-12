package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint }

func main() {
    var v1 string = "hello"
    var v2 int64 = 1
    v3 := []string{((v1 + ((v1 + (v1 + v1)) + v1)) + "hello"), "foo bar"}
    var v4 string = v1
    var v5 uint32 = 826
    for fi6 := 0; fi6 < 2; fi6++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch "a" {
        case "foo bar":
            var v7 int32 = 1
            _ = v7
        case "hello":
            if false {
                for fi8 := 0; fi8 < 3; fi8++ {
                    switch "\n\t" {
                    case "hello":
                        v9 := GenS1{F0: (fi6 / ((-1) | 1))}
                        _ = v9
                    default:
                        var v10 int16 = -790
                        _ = v10
                    }
                    if (-203 < -2147483648) {
                        if len(v3) > 0 {
                            v3[0] = v4
                        }
                        v11 := GenS1{F0: (fi8 ^ fi6)}
                        if (!(!(v5 == v5))) {
                            switch -528 {
                            case -9223372036854775808:
                                var v12 float32 = 0.1
                                _ = v12
                            case 4294967295:
                                var v13 uint64 = 984
                                _ = v13
                            case -534:
                                fmt.Println(v1)
                            default:
                                _ = len(v3)
                            }
                        }
                        _ = v11
                    }
                    var v14 int8 = -128
                    _ = fi8
                    _ = v14
                }
            } else {
                if (2 <= 0) {
                    v5 = v5
                }
            }
        }
        _ = fi6
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := GenS2{F0: 142, F1: (fi15 % (((fi15 << ((4294967295) & 63))) | 1)), F2: 32767}
        v1 = v4
        var v17 int64 = v2
        _ = fi15
        _ = v16
        _ = v17
    }
    v1 = ""
    fmt.Println(v2)
    if (true || (v5 < (v5 / ((v5) | 1)))) {
        var v18 uint8 = 127
        _ = v18
    }
    v1 = v1
    v2 -= v2
    v19 := GenS3{F0: 2147483647, F1: 42}
    var v20 bool = true
    v4 = v4
    v20 = (1 != 794)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
