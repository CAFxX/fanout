package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint16 }

func main() {
    var v1 bool = (!(!false))
    v2 := make(map[uint32]uint32)
    var v3 bool = (-213 == 9223372036854775807)
    var v4 int32 = 42
    v5 := []string{}
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 uint8 = 127
        _ = fi6
        _ = v7
    }
    switch -537 {
    case 191:
        v8 := GenS1{F0: -2147483648, F1: 65535}
        _ = v8
    case -1:
        v3 = (0.1 == (((1.5 * 1.5) / 0.1) - 0.1))
    case 32767:
        _ = len(v5)
    }
    var v9 uint32 = 65535
    if v3 {
        v1 = (457 != 32767)
        fmt.Println(v1)
    }
    switch "hello" {
    case "":
        if false {
            fmt.Println(v3, v1, v9)
        } else {
            var v10 int64 = 481
            _ = v10
        }
    case "\n\t":
        if v1 {
            v11 := GenS2{F0: 42, F1: 4294967295, F2: 1}
            _ = v11
        }
    default:
        v4 = (v4 + v4)
    }
    v12 := GenS3{F0: -547, F1: 0}
    var v13 uint64 = 458
    v2[(v9 / ((v9) | 1))] = v9
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 string = (("\n\t" + ("" + "")) + "")
        if (!(!true)) {
            _ = len(v5)
        }
        switch fi14 {
        case 4294967295:
            if (78 < 127) {
                var v16 uint32 = (v9 + v9)
                var v17 uint64 = 0
                _ = v16
                _ = v17
            } else {
                if ((v4 * (v4 >> ((306) & 63))) == (v4 << ((693) & 63))) {
                    switch (fi14 / (((fi14 << ((34) & 63))) | 1)) {
                    case -9223372036854775808:
                        fmt.Println(v13, v4, v9)
                    case 2147483647:
                        var v18 bool = (v4 < v4)
                        _ = v18
                    }
                    v19 := GenS3{F0: -1, F1: 42}
                    _ = v19
                } else {
                    if false {
                        delete(v2, (^(v9)))
                        _ = len(v5)
                    }
                }
            }
        }
        _ = fi14
        _ = v15
    }
    var v20 float64 = 0.1
    v4 -= v4
    v20 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
