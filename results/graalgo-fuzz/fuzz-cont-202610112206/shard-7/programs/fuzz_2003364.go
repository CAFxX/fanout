package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = -1
    var v3 int32 = -1
    var v4 int = 1
    var v5 int32 = (v3 / (((v3 | (v3 >> ((65535) & 63)))) | 1))
    var v6 int = int(v2)
    if true {
        var v7 bool = (42 < -122)
        fmt.Println(v5, v4)
        v8 := GenS1{F0: (v5 >> ((18446744073709551615) & 63)), F1: 6}
        _ = v7
        _ = v8
    } else {
        switch (v6 * (v4 >> ((2) & 63))) {
        case 42:
            for fi9 := 0; fi9 < 20; fi9++ {
                switch v6 {
                case 153:
                    switch fi9 {
                    case 127:
                        for fi10 := 0; fi10 < 7; fi10++ {
                            if (!(!true)) {
                                v11 := GenS1{F0: v3, F1: 185}
                                var v12 float32 = 0.1
                                v13 := GenS1{F0: (v3 >> ((4294967295) & 63)), F1: 0}
                                _ = v11
                                _ = v12
                                _ = v13
                            }
                            v2 -= -299
                            _ = fi10
                        }
                    }
                case 32767:
                    var v14 bool = true
                    _ = v14
                case -9223372036854775808:
                    var v15 int32 = v3
                    _ = v15
                }
                v5 += (-(v5))
                _ = fi9
            }
        }
        var v16 int32 = (-(v3))
        v6 = (^(v4))
        _ = v16
    }
    v17 := []int{(v6 << ((18446744073709551615) & 63)), (v4 / ((v4) | 1)), -2147483648, (v6 << ((731) & 63))}
    var v18 int32 = 958
    v4 = v6
    v18 = (v18 & (v18 + v18))
    switch (^(v4)) {
    case -2147483648:
        fmt.Println(v6, v4)
    case -464:
        var v19 uint = 273
        _ = v19
    }
    v4 += (v4 | 42)
    v20 := GenS1{F0: (v18 << ((65535) & 63)), F1: 241}
    var v21 uint32 = 975
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
