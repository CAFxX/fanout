package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 string; F1 int16 }
func genFn0() int16 {
    v5 := []int{}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = v5
    return 1
}

func main() {
    var v1 float32 = -2.25
    v2 := []string{("\n\t" + "a"), "foo bar"}
    var v3 float64 = float64(v1)
    var v4 int32 = -1
    var v6 int64 = -71
    v7 := GenS1{F0: 0, F1: 160}
    fmt.Println(v3, v1)
    v8 := GenS1{F0: (-(v6)), F1: 128}
    for fi9 := 0; fi9 < 14; fi9++ {
        if false {
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 bool = (((("hello" + "a") != ("hello" + "foo bar")) || ("hello" < "hello")) && (1e10 >= 1.5))
                if true {
                    switch "" {
                    case "foo bar":
                        if v11 {
                            v2 = append(v2, "foo bar")
                        } else {
                            if ((0 <= fi9) || true) {
                                if (v1 != 0.1) {
                                    var v12 uint32 = 2147483648
                                    v13 := GenS1{F0: v6, F1: 1}
                                    _ = v12
                                    _ = v13
                                } else {
                                    switch "\n\t" {
                                    case "":
                                        fmt.Println(v1)
                                    case "\n\t":
                                        v14 := GenS1{F0: v6, F1: 42}
                                        _ = v14
                                    }
                                    var v15 uint8 = 204
                                    _ = v15
                                }
                                v2 = append(v2, "a")
                                v16 := GenS1{F0: v6, F1: 123}
                                _ = v16
                            } else {
                                var v17 bool = true
                                v2 = append(v2, ((("\n\t" + "foo bar") + "") + ""))
                                _ = v17
                            }
                        }
                    case "hello":
                        _ = len(v2)
                    case "a":
                        v3 = 0.0
                        v6 = (v6 | v6)
                    }
                } else {
                    v1 = v1
                    v18 := GenS2{F0: 18446744073709551615, F1: (v6 / ((-9223372036854775808) | 1))}
                    var v19 uint = 255
                    _ = v18
                    _ = v19
                }
                var v20 uint64 = 11
                _ = fi10
                _ = v11
                _ = v20
            }
            var v21 float32 = v1
            _ = v21
        } else {
            var v22 int32 = 0
            _ = len(v2)
            _ = v22
        }
        switch -1 {
        case 487:
            var v23 int8 = -128
            _ = v23
        case 650:
            fmt.Println(v6, v3, v6)
        }
        fmt.Println(v6, v4)
        _ = fi9
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        v2 = append(v2, "")
        _ = fi24
    }
    _ = len(v2)
    v4 += v4
    switch ("a" + "a") {
    case "\n\t":
        fmt.Println(v1)
    case "hello":
        v4 = v4
    case "a":
        v25 := GenS2{F0: 65535, F1: (v6 % (((v6 % ((-1) | 1))) | 1))}
        _ = v25
    }
    v26 := GenS3{F0: ("foo bar" + "\n\t"), F1: 1000}
    v2 = append(v2, "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
