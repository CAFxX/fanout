package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint8; F1 bool; F2 string }

func main() {
    var v1 uint16 = 2
    v2 := make(map[uint32]string)
    v3 := make(map[int]string)
    var v4 uint16 = 754
    v5 := make(map[string]bool)
    var v6 int8 = 97
    v7 := []int{202, -9223372036854775808, 2147483647}
    var v8 int16 = -32768
    v9 := GenS1{F0: -128, F1: 28, F2: (v1 - v1)}
    fmt.Println(v4, v8, v8)
    var v10 uint16 = (v1 - 255)
    switch "" {
    case "\n\t":
        delete(v5, "a")
    case "":
        fmt.Println(v6, v10)
    case "foo bar":
        if false {
            var v11 uint8 = 1
            v4 = v10
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 17; fi12++ {
            if false {
                for fi13 := 0; fi13 < 2; fi13++ {
                    v14 := v2[439]
                    if true {
                        switch -2147483648 {
                        case 32767:
                            var v15 string = (("a" + "foo bar") + "\n\t")
                            _ = v15
                        case -9223372036854775808:
                            v7 = append(v7, fi12)
                        default:
                            var v16 uint16 = v1
                            _ = v16
                        }
                        v7 = append(v7, fi13)
                    } else {
                        v17 := GenS1{F0: v6, F1: v6, F2: 1}
                        _ = v17
                    }
                    _ = fi13
                    _ = v14
                }
            } else {
                for fi18 := 0; fi18 < 1; fi18++ {
                    v19 := GenS2{F0: 255, F1: (false || false), F2: "foo bar"}
                    v10 = v10
                    var v20 bool = true
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                var v21 uint32 = 0
                if true {
                    v22 := v5["hello"]
                    _ = v22
                }
                _ = v21
            }
            var v23 int32 = 2147483647
            v8 = 32767
            _ = fi12
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        v25 := GenS2{F0: 128, F1: true, F2: (("" + ("foo bar" + ("" + ("" + ("" + "a"))))) + "\n\t")}
        _ = fi24
        _ = v25
    }
    fmt.Println(v10, v4)
    var v26 uint8 = 128
    fmt.Println(v26, v8)
    switch -990 {
    case 4294967295:
        v10 += 32767
    case 32767:
        _ = len(v7)
    default:
        switch "\n\t" {
        case "hello":
            var v27 float64 = (0.1 / ((3.14159 / 0.1) * -2.25))
            _ = v27
        case "\n\t":
            v2[869] = ("\n\t" + "a")
        default:
            v8 += -1
        }
    }
    var v28 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
