package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 bool) int32 {
    if (0 < 2) {
        var v4 uint16 = 754
        if (!(((9223372036854775807 > 1) || ((42 >= 685) || (32767 == 9223372036854775807))) || true)) {
            p0 = p0
            switch -9223372036854775808 {
            case 9223372036854775807:
                for fi5 := 0; fi5 < 15; fi5++ {
                    for fi6 := 0; fi6 < 7; fi6++ {
                        if false {
                            v7 := GenS1{F0: -921, F1: 0}
                            _ = v7
                        }
                        v4 = v4
                        p0 = p0
                        _ = fi6
                    }
                    v8 := []int{(-(fi5)), 206}
                    _ = fi5
                    _ = v8
                }
            case 1:
                for fi9 := 0; fi9 < 14; fi9++ {
                    for fi10 := 0; fi10 < 10; fi10++ {
                        var v11 int16 = 1000
                        _ = fi10
                        _ = v11
                    }
                    _ = fi9
                }
            case 32767:
                if false {
                    var v12 int8 = 127
                    _ = v12
                } else {
                    v13 := GenS1{F0: 42, F1: -743}
                    if false {
                        var v14 int16 = 0
                        var v15 int8 = 0
                        fmt.Println(v4, v4, v15)
                        _ = v14
                        _ = v15
                    } else {
                        var v16 uint16 = v4
                        _ = v16
                    }
                    _ = v13
                }
            default:
                v17 := GenS1{F0: -32768, F1: -9223372036854775808}
                _ = v17
            }
        }
        _ = v4
    } else {
        if (42 > -1) {
            var v18 uint16 = 106
            _ = v18
        } else {
            v19 := GenS2{F0: 127, F1: 100}
            fmt.Println(p0)
            if ("a" != ("a" + "hello")) {
                fmt.Println(p0)
                var v20 uint8 = 1
                v21 := []int{1}
                _ = v20
                _ = v21
            } else {
                v22 := []int{-460, -870}
                v22 = append(v22, 128)
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                _ = v22
            }
            _ = v19
        }
        p0 = p0
    }
    var v23 uint = 2
    _ = p0
    _ = v23
    return -309
}

func main() {
    var v1 string = ("foo bar" + "foo bar")
    v2 := []int{}
    var v3 float32 = float32((((3.14159 * 3.14159) - (float64((0.0 * 1e10)) - float64(1.5))) * 1.5))
    v1 = v1
    switch "a" {
    case "a":
        if false {
            var v24 int64 = 9223372036854775807
            _ = v24
        }
    case "":
        var v25 int16 = 1
        _ = v25
    }
    var v26 uint64 = 712
    v3 -= 3.14159
    fmt.Println(v3, v26, v1)
    v26 = v26
    var v27 int = 4294967295
    var v28 float64 = 0.0
    if (false && (1.5 > (-2.25 - v28))) {
        v3 += v3
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println(v26)
    var v29 int64 = 564
    var v30 bool = false
    v26 = (v26 >> ((117) & 63))
    var v31 string = ((v1 + v1) + "foo bar")
    for fi32 := 0; fi32 < 13; fi32++ {
        var v33 float32 = (v3 * 0.1)
        _ = fi32
        _ = v33
    }
    switch v31 {
    case "\n\t":
        fmt.Println(v28, v1, v30)
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = -2147483648
        }
    case "":
        v1 = (("\n\t" + v1) + "a")
    }
    v26 = v26
    v34 := GenS3{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
}
