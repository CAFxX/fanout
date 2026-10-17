package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 bool }
type GenS2 struct { F0 uint64 }
func genFn0() string {
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := []int{fi6, (fi6 + fi6), fi6, 9223372036854775807}
        _ = fi6
        _ = v7
    }
    var v8 int = -9223372036854775808
    _ = v8
    return "\n\t"
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []int{1}
    var v3 bool = (!false)
    var v4 int64 = 0
    var v5 string = ("" + "\n\t")
    v9 := GenS1{F0: int32(v4), F1: v4, F2: (v1 != (v1 | (v1 - v1)))}
    if (!true) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        for fi10 := 0; fi10 < 20; fi10++ {
            if ((fi10 << ((339) & 63)) <= fi10) {
                var v11 string = v5
                _ = v11
            }
            _ = fi10
        }
        fmt.Println(v1, v1)
        var v12 uint64 = 4294967295
        _ = v12
    }
    v5 = v5
    var v13 uint8 = 128
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v14 string = ((v5 + "a") + "hello")
    var v15 uint64 = 18446744073709551615
    if (!v3) {
        for fi16 := 0; fi16 < 9; fi16++ {
            v15 = 40
            _ = fi16
        }
        var v17 uint8 = (v13 << ((757) & 63))
        _ = v17
    } else {
        switch (v5 + v14) {
        case "":
            if true {
                v18 := GenS1{F0: -2147483648, F1: (v1 + (-(v4))), F2: (v3 || v3)}
                _ = len(v2)
                _ = v18
            } else {
                v19 := GenS1{F0: 2147483647, F1: 26, F2: false}
                if (v3 && (-2147483648 == 4294967295)) {
                    v13 = (^(v13))
                }
                v20 := GenS1{F0: -362, F1: 650, F2: (v3 && (1 == -1))}
                _ = v19
                _ = v20
            }
        case "foo bar":
            v21 := GenS2{F0: 956}
            _ = v21
        case "hello":
            if ((2 <= 4294967295) || ((!(v3 || false)) && (65535 < 42))) {
                for fi22 := 0; fi22 < 6; fi22++ {
                    var v23 int64 = (v1 | v4)
                    switch "a" {
                    case "hello":
                        var v24 uint32 = 4294967295
                        _ = v24
                    case "\n\t":
                        v15 = v15
                    case "a":
                        switch v14 {
                        case "foo bar":
                            fmt.Println(v15, v3)
                        default:
                            if (!((!(v3 && true)) || v3)) {
                                v2 = append(v2, 127)
                                v3 = ((42 <= int32(fi22)) || false)
                            } else {
                                v23 -= v4
                                v5 = "a"
                            }
                        }
                    default:
                        switch -1 {
                        case -789:
                            fmt.Println(v3, v3, v15)
                        case 32767:
                            var v25 uint32 = 2
                            _ = v25
                        }
                    }
                    switch int(v13) {
                    case 295:
                        var v26 uint64 = 42
                        _ = v26
                    case -525:
                        var v27 uint32 = 942
                        _ = v27
                    default:
                        v28 := GenS2{F0: (v15 / (((^(v15))) | 1))}
                        _ = v28
                    }
                    _ = fi22
                    _ = v23
                }
            }
        default:
            switch 4294967295 {
            case -137:
                var v29 float64 = -2.25
                _ = v29
            case -9223372036854775808:
                var v30 uint8 = 0
                _ = v30
            case 128:
                if ((^(v13)) == (v13 / (((v13 | v13)) | 1))) {
                    var v31 uint64 = 2
                    var v32 uint8 = (^(v13))
                    _ = v31
                    _ = v32
                } else {
                    for fi33 := 0; fi33 < 14; fi33++ {
                        for fi34 := 0; fi34 < 18; fi34++ {
                            v14 = ""
                            if len(v2) > 0 {
                                v2[0] = fi33
                            }
                            _ = fi34
                        }
                        var v35 int16 = 0
                        switch "a" {
                        case "a":
                            var v36 float32 = (3.14159 / 1.0)
                            _ = v36
                        default:
                            var v37 int = 2147483647
                            _ = v37
                        }
                        _ = fi33
                        _ = v35
                    }
                    v5 = "a"
                }
            default:
                var v38 uint8 = (v13 / (((v13 + (^(v13)))) | 1))
                _ = v38
            }
        }
        switch v5 {
        case "hello":
            v15 = v15
        case "a":
            v5 = v5
        }
    }
    _ = len(v2)
    v15 = v15
    if false {
        var v39 uint8 = (v13 % ((v13) | 1))
        _ = v39
    }
    var v40 int16 = -505
    fmt.Println(v3)
    if v3 {
        v4 += (v4 - (v1 & (v1 - v4)))
    } else {
        var v41 float64 = 1e10
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v9
}
