package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }

func main() {
    v1 := []int{127, 0}
    v2 := make(map[string]bool)
    var v3 int = 127
    for fi4 := 0; fi4 < 18; fi4++ {
        switch 1 {
        case -9223372036854775808:
            var v5 float32 = 0.0
            _ = v5
        }
        var v6 float32 = (1.5 * (((-2.25 / 0.1) * 1.5) + 0.0))
        for fi7 := 0; fi7 < 6; fi7++ {
            v3 = fi4
            _ = fi7
        }
        _ = fi4
        _ = v6
    }
    var v8 int16 = 32767
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int16 = v8
        switch 1 {
        case 4294967295:
            if ((false && (!(!false))) && ((100 != -128) && (false && ("" > "hello")))) {
                if false {
                    switch -409 {
                    case 2147483647:
                        var v11 uint16 = 525
                        _ = v11
                    case 785:
                        v8 = v8
                    case -421:
                        var v12 uint = 1
                        _ = v12
                    default:
                        for fi13 := 0; fi13 < 16; fi13++ {
                            v14 := GenS1{F0: (fi9 - 2147483647)}
                            if false {
                                v10 += (v8 & v10)
                            }
                            _ = fi13
                            _ = v14
                        }
                    }
                    fmt.Println(v3)
                } else {
                    v15 := GenS1{F0: (fi9 * (^(v3)))}
                    _ = len(v2)
                    _ = v15
                }
            }
        case -933:
            switch (v3 % (((v3 - v3)) | 1)) {
            case 1000:
                var v16 float32 = 0.1
                _ = v16
            default:
                fmt.Println(v10, fi9, fi9)
            }
        default:
            var v17 float32 = 0.0
            _ = v17
        }
        _ = fi9
        _ = v10
    }
    v18 := GenS1{F0: v3}
    v19 := v2["hello"]
    switch ("" + "\n\t") {
    case "foo bar":
        v1 = append(v1, -97)
    case "":
        var v20 uint8 = 71
        _ = v20
    default:
        if v19 {
            var v21 int = 4294967295
            fmt.Println(v3)
            _ = len(v2)
            _ = v21
        }
    }
    v8 = (v8 >> ((42) & 63))
    fmt.Println(v19, v3, v8)
    var v22 float32 = float32((-2.25 / 1e10))
    for fi23 := 0; fi23 < 11; fi23++ {
        var v24 int64 = -1
        _ = fi23
        _ = v24
    }
    var v25 bool = ((v19 || (v19 || v19)) && (!(18446744073709551615 != 681)))
    if v25 {
        var v26 int16 = v8
        switch v3 {
        case 32768:
            switch 9223372036854775807 {
            case 4294967295:
                for fi27 := 0; fi27 < 3; fi27++ {
                    v25 = (false || v25)
                    v26 = -872
                    var v28 int32 = 2147483647
                    _ = fi27
                    _ = v28
                }
            case 9223372036854775807:
                switch ("foo bar" + "a") {
                case "a":
                    var v29 int16 = v8
                    _ = v29
                case "\n\t":
                    fmt.Println(v3)
                }
            case 706:
                v25 = v25
            }
        default:
            if ((127 > (v3 / (((v3 % ((v3) | 1))) | 1))) && (925 == 0)) {
                switch 9223372036854775807 {
                case 42:
                    var v30 int = (v3 / ((v3) | 1))
                    _ = v30
                default:
                    v2[("" + "a")] = (1 <= 255)
                }
            }
        }
        v31 := GenS2{F0: 522}
        _ = v26
        _ = v31
    }
    v8 += v8
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v8
}
