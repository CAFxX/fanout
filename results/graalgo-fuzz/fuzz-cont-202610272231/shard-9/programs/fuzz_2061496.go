package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int32; F1 uint64; F2 int32 }
func genFn0() string {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 bool = ("\n\t" == ("" + ("a" + ("hello" + ""))))
        switch 2147483647 {
        case 42:
            if (719 == 853) {
                var v7 int32 = 0
                for fi8 := 0; fi8 < 7; fi8++ {
                    switch "a" {
                    case "\n\t":
                        var v9 uint8 = 1
                        _ = v9
                    default:
                        switch fi5 {
                        case 127:
                            v7 -= 1
                        }
                    }
                    v10 := []int{}
                    fmt.Println(fi8)
                    _ = fi8
                    _ = v10
                }
                _ = v7
            }
        case 32767:
            switch "\n\t" {
            case "hello":
                fmt.Println(v6, fi5)
            case "\n\t":
                var v11 uint8 = 86
                _ = v11
            case "a":
                if (42 == -1) {
                    v12 := []int{9223372036854775807}
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 19; fi13++ {
                        v14 := []int{-673, fi13}
                        _ = fi13
                        _ = v14
                    }
                    v15 := GenS1{F0: 77, F1: 127}
                    v6 = (1000 >= 1000)
                    _ = v15
                }
            }
        }
        _ = fi5
        _ = v6
    }
    var v16 int16 = 1
    v17 := GenS1{F0: 13, F1: -128}
    fmt.Println(v16)
    v16 = 0
    _ = v16
    _ = v17
    return "foo bar"
}

func main() {
    var v1 uint64 = 1
    var v2 uint64 = (v1 * (v1 << ((424) & 63)))
    var v3 int8 = 127
    var v4 int32 = 2147483647
    fmt.Println(v3)
    switch -120 {
    case 32767:
        v18 := GenS2{F0: v4, F1: (v1 & 0), F2: (v4 << ((0) & 63))}
        _ = v18
    default:
        var v19 uint8 = 255
        _ = v19
    }
    if true {
        fmt.Println(v4, v4)
        fmt.Println(v4, v1)
        switch "a" {
        case "":
            v20 := GenS2{F0: (v4 - int32(v3)), F1: (^(v2)), F2: (^(v4))}
            _ = v20
        default:
            var v21 int16 = -1
            _ = v21
        }
    }
    var v22 string = "a"
    var v23 bool = (1e10 < 3.14159)
    var v24 uint16 = 255
    var v25 uint = 4294967295
    v26 := GenS1{F0: v3, F1: -40}
    v27 := GenS1{F0: v3, F1: v3}
    switch "" {
    case "\n\t":
        var v28 int8 = 100
        _ = v28
    default:
        v23 = (!(!v23))
    }
    v29 := []int{128, 101, 32767}
    var v30 float64 = 1.5
    if len(v29) >= 2 {
        v29 = v29[1:len(v29)]
    }
    v30 += v30
    v31 := GenS2{F0: 1, F1: (v1 * v2), F2: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
