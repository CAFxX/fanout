package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint; F1 bool }
func genFn0() bool {
    var v5 uint = 65535
    var v6 int8 = 42
    switch "foo bar" {
    case "\n\t":
        var v7 uint32 = 2147483647
        _ = v7
    case "foo bar":
        v6 = v6
    case "":
        for fi8 := 0; fi8 < 6; fi8++ {
            for fi9 := 0; fi9 < 15; fi9++ {
                if false {
                    v5 += 0
                    switch (fi9 >> ((v5) & 63)) {
                    case 127:
                        v5 -= (v5 + v5)
                    case 32768:
                        v10 := GenS1{F0: (fi9 - fi9), F1: 126}
                        _ = v10
                    case 2147483647:
                        var v11 bool = false
                        _ = v11
                    default:
                        var v12 int64 = -856
                        _ = v12
                    }
                }
                fmt.Println(fi8, v6)
                switch "" {
                case "":
                    v13 := GenS1{F0: (fi9 * (fi8 % ((fi9) | 1))), F1: 0}
                    if true {
                        v14 := []int{-2147483648, 32767, fi8, 32767}
                        var v15 float32 = 0.1
                        fmt.Println(v6, v15, v6)
                        _ = v14
                        _ = v15
                    }
                    _ = v13
                case "a":
                    fmt.Println(v5, fi8, fi8)
                case "\n\t":
                    if (true && true) {
                        var v16 bool = ("foo bar" == "hello")
                        v6 = 100
                        v5 += v5
                        _ = v16
                    } else {
                        v17 := GenS2{F0: 42, F1: 748}
                        v18 := []int{(^(fi9)), -605, (fi9 / ((fi8) | 1)), fi8}
                        if (true || (((true && (0 > 42)) && false) || ((1.5 / (((float32((float64((1e10 * 1e10)) - 1e10)) + 1e10) * 0.1) - 3.14159)) < 1.5))) {
                            switch 9223372036854775807 {
                            case -749:
                                v6 += (v6 % ((v6) | 1))
                            case 128:
                                var v19 float64 = 0.1
                                _ = v19
                            case 147:
                                var v20 int = (fi8 | fi9)
                                _ = v20
                            default:
                                v6 = v6
                            }
                            v6 = 42
                        }
                        v21 := GenS1{F0: (^(fi8)), F1: 2}
                        _ = v17
                        _ = v18
                        _ = v21
                    }
                }
                _ = fi9
            }
            fmt.Println(v6, fi8, fi8)
            _ = fi8
        }
    }
    _ = v5
    _ = v6
    return true
}

func main() {
    var v1 string = "foo bar"
    var v2 int64 = 1
    var v3 uint = 1
    v4 := []string{"a"}
    var v22 float32 = -2.25
    fmt.Println(v22, v1, v3)
    v22 -= 1e10
    fmt.Println(v2, v3, v2)
    v2 = (v2 | v2)
    v23 := GenS3{F0: (^(v3)), F1: (!(v3 != (v3 % (((v3 + (v3 << (((v3 << ((v3) & 63))) & 63)))) | 1))))}
    if false {
        var v24 int8 = 42
        v24 = -2
        _ = v24
    } else {
        switch 0 {
        case -2147483648:
            v25 := GenS1{F0: 32768, F1: 32767}
            _ = v25
        }
    }
    if len(v4) > 0 {
        v4[0] = ("\n\t" + "foo bar")
    }
    fmt.Println(v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
