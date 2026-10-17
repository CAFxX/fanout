package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int32; F1 uint16 }
func genFn0() uint8 {
    v7 := []int{128}
    if false {
        if true {
            var v8 int8 = 127
            v8 = (^(v8))
            var v9 float32 = -2.25
            _ = v8
            _ = v9
        }
    }
    _ = v7
    return 42
}

func main() {
    var v1 uint32 = 2147483647
    var v2 bool = false
    v3 := make(map[string]int)
    v4 := make(map[uint32]bool)
    var v5 bool = false
    var v6 string = "hello"
    if v2 {
        v1 -= 65535
        v10 := []int{682, -9223372036854775808}
        _ = v10
    } else {
        v11 := []int{127, -9223372036854775808}
        switch 127 {
        case 32767:
            v12 := GenS1{F0: -1, F1: -83}
            _ = v12
        }
        fmt.Println(v5, v5)
        _ = v11
    }
    var v13 string = (((((v6 + v6) + ((v6 + "hello") + v6)) + (v6 + v6)) + "") + "foo bar")
    v14 := []int{}
    v1 = 42
    switch (v6 + v13) {
    case "a":
        var v15 float64 = (0.0 / -2.25)
        _ = v15
    case "foo bar":
        if v5 {
            var v16 int64 = 0
            _ = len(v4)
            if (false || ((v2 || (v5 && (0 > 1))) || true)) {
                switch v6 {
                case "\n\t":
                    v17 := GenS1{F0: -970, F1: 67}
                    _ = v17
                case "":
                    switch v6 {
                    case "foo bar":
                        var v18 int8 = 42
                        _ = v18
                    case "hello":
                        switch 4294967295 {
                        case 2147483647:
                            v13 = "\n\t"
                        case 904:
                            v19 := GenS1{F0: 127, F1: -1}
                            _ = v19
                        case 401:
                            v3[v13] = 513
                        }
                    default:
                        v20 := GenS1{F0: 128, F1: -25}
                        _ = v20
                    }
                case "foo bar":
                    var v21 uint16 = 848
                    _ = v21
                default:
                    v22 := GenS1{F0: 2147483647, F1: 0}
                    _ = v22
                }
                v16 = 9223372036854775807
            }
            _ = v16
        } else {
            if v5 {
                switch "hello" {
                case "":
                    var v23 uint16 = 572
                    _ = v23
                case "foo bar":
                    v1 += v1
                }
            } else {
                var v24 uint = 421
                v25 := GenS2{F0: -915, F1: 650}
                _ = v24
                _ = v25
            }
            for fi26 := 0; fi26 < 15; fi26++ {
                var v27 uint64 = 42
                _ = fi26
                _ = v27
            }
            var v28 uint16 = 0
            _ = v28
        }
    }
    switch ("\n\t" + "") {
    case "":
        if v2 {
            fmt.Println(v2, v13, v5)
            var v29 int32 = 1
            v2 = (0.0 <= (1.5 - (float32(1.5) / 1.5)))
            _ = v29
        } else {
            v30 := GenS2{F0: 42, F1: 171}
            _ = v30
        }
    case "hello":
        switch v13 {
        case "foo bar":
            v14 = append(v14, 1)
        case "\n\t":
            var v31 int32 = 42
            _ = v31
        }
    }
    fmt.Println(v5, v5, v1)
    v13 = v13
    var v32 uint32 = (^(v1))
    _ = len(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
