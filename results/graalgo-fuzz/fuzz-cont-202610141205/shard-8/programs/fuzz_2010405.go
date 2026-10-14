package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint8; F1 int16; F2 int64 }
func genFn0() uint16 {
    if (1 == 263) {
        v6 := []int{128, -539}
        _ = v6
    } else {
        v7 := []int{}
        switch "hello" {
        case "a":
            for fi8 := 0; fi8 < 1; fi8++ {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                var v9 bool = (false && true)
                _ = fi8
                _ = v9
            }
        case "hello":
            switch -9223372036854775808 {
            case 101:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            case -2147483648:
                if (0 != 1) {
                    var v10 uint8 = 255
                    var v11 bool = ((true && (!(true && (false || (0 > 2))))) && ("foo bar" > "a"))
                    _ = v10
                    _ = v11
                } else {
                    if len(v7) > 0 {
                        v7[0] = 32767
                    }
                    if true {
                        v7 = append(v7, 127)
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        if len(v7) > 0 {
                            v7[0] = 42
                        }
                    }
                }
            default:
                for fi12 := 0; fi12 < 19; fi12++ {
                    v13 := GenS1{F0: 45, F1: true}
                    fmt.Println(fi12)
                    v14 := GenS2{F0: "hello", F1: (2147483648 < 674)}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
        case "\n\t":
            v15 := GenS1{F0: 100, F1: true}
            _ = v15
        default:
            var v16 float32 = 3.14159
            _ = v16
        }
        var v17 string = ((("\n\t" + "\n\t") + "\n\t") + "")
        _ = v17
        _ = v7
    }
    if true {
        var v18 int32 = -2147483648
        v19 := []int{}
        _ = v18
        _ = v19
    }
    v20 := []int{435, 128, -872, 740}
    var v21 int = 1
    switch v21 {
    case -97:
        v21 = (v21 & (^(v21)))
    case -60:
        _ = len(v20)
    case -9223372036854775808:
        for fi22 := 0; fi22 < 7; fi22++ {
            v23 := GenS1{F0: -121, F1: (((float32(1e10) + 1.5) - 1.5) != 3.14159)}
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            v24 := GenS1{F0: 0, F1: true}
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    _ = v20
    _ = v21
    return 205
}

func main() {
    var v1 uint = 489
    var v2 int32 = -516
    var v3 float64 = -2.25
    var v4 uint32 = 4294967295
    var v5 int16 = -695
    var v25 string = ""
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 int64 = 1
        var v28 int8 = -5
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := []int{-807, 605, 2147483647}
    var v30 int32 = (v2 & (v2 >> ((v1) & 63)))
    v31 := GenS2{F0: "hello", F1: true}
    var v32 int32 = 0
    if false {
        var v33 int8 = -1
        _ = v33
    } else {
        switch -1 {
        case 0:
            if (0 >= 1) {
                v34 := GenS3{F0: 0, F1: (v5 % ((v5) | 1)), F2: -9223372036854775808}
                _ = v34
            }
        case -228:
            _ = len(v29)
        default:
            if len(v29) > 0 {
                v29[0] = 533
            }
        }
        var v35 uint = (v1 + (v1 * (v1 | 853)))
        _ = v35
    }
    switch v25 {
    case "":
        var v36 int64 = 42
        _ = v36
    case "hello":
        fmt.Println(v5, v4, v25)
    default:
        var v37 float32 = 1e10
        _ = v37
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
