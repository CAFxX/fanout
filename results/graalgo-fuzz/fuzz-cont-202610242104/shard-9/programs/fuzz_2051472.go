package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 bool; F1 string }

func main() {
    var v1 uint16 = 1
    var v2 uint32 = 361
    var v3 uint = 89
    var v4 uint8 = 2
    switch "\n\t" {
    case "hello":
        var v5 uint16 = (v1 * 42)
        _ = v5
    case "\n\t":
        fmt.Println(v1, v2)
    case "a":
        switch 4294967295 {
        case 1:
            if (true || (-1 != 364)) {
                for fi6 := 0; fi6 < 17; fi6++ {
                    if (((1e10 + float32(1e10)) + ((1.5 - (1e10 + float32(0.1))) * 1.5)) == 1.5) {
                        var v7 float64 = 1e10
                        v1 -= 42
                        _ = v7
                    }
                    _ = fi6
                }
                if true {
                    var v8 uint8 = v4
                    _ = v8
                } else {
                    fmt.Println(v2)
                }
            } else {
                var v9 int8 = 1
                fmt.Println(v3)
                for fi10 := 0; fi10 < 17; fi10++ {
                    v4 += v4
                    for fi11 := 0; fi11 < 5; fi11++ {
                        var v12 int = fi10
                        _ = fi11
                        _ = v12
                    }
                    _ = fi10
                }
                _ = v9
            }
        case -2147483648:
            var v13 string = "a"
            _ = v13
        }
    }
    var v14 uint = (v3 << (((^(v3))) & 63))
    fmt.Println(v14, v4)
    v15 := GenS1{F0: ("\n\t" + "foo bar"), F1: true}
    var v16 int = 723
    var v17 int64 = 0
    var v18 int = v16
    fmt.Println(v3, v18, v16)
    for fi19 := 0; fi19 < 4; fi19++ {
        v14 = uint(v2)
        var v20 int32 = 1
        v4 += (^(v4))
        _ = fi19
        _ = v20
    }
    switch "" {
    case "hello":
        v21 := []int{128, 9223372036854775807}
        _ = v21
    case "a":
        v1 = 896
    case "foo bar":
        switch "" {
        case "":
            switch v18 {
            case 2147483647:
                for fi22 := 0; fi22 < 1; fi22++ {
                    for fi23 := 0; fi23 < 5; fi23++ {
                        fmt.Println(v16, v14, v18)
                        _ = fi23
                    }
                    var v24 uint = v3
                    _ = fi22
                    _ = v24
                }
            case -9223372036854775808:
                var v25 float32 = -2.25
                _ = v25
            default:
                var v26 uint = (v3 % ((0) | 1))
                _ = v26
            }
        case "foo bar":
            if (uint16(v4) != 1) {
                v27 := []int{959, 32768, int(v2)}
                v18 += v16
                _ = v27
            } else {
                for fi28 := 0; fi28 < 18; fi28++ {
                    v3 = uint(v2)
                    fmt.Println(v16, v4, v3)
                    var v29 string = "hello"
                    _ = fi28
                    _ = v29
                }
                var v30 int32 = 0
                v31 := GenS2{F0: (!false), F1: "hello"}
                _ = v30
                _ = v31
            }
        }
    }
    var v32 int16 = -32768
    fmt.Println(v18, v1, v18)
    v33 := []int{0}
    var v34 uint16 = 23
    var v35 bool = false
    if v35 {
        var v36 string = "foo bar"
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
