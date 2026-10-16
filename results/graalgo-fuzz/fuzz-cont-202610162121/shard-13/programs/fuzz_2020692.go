package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int16; F1 int; F2 uint64 }
type GenS4 struct { F0 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 bool = true
    var v3 int = -9223372036854775808
    v4 := []int{(v3 - v3)}
    v5 := GenS1{F0: 143}
    if v2 {
        var v6 int32 = -2147483648
        v7 := GenS1{F0: 199}
        var v8 uint16 = 65535
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 uint64 = 217
    v3 -= v3
    if (1 <= 0) {
        if v2 {
            for fi10 := 0; fi10 < 5; fi10++ {
                fmt.Println(v9, fi10, v3)
                _ = fi10
            }
            v9 -= v9
            if (v9 >= (v9 / (((v9 / ((v9) | 1))) | 1))) {
                for fi11 := 0; fi11 < 14; fi11++ {
                    if ("foo bar" == "foo bar") {
                        v3 = (fi11 ^ (^(fi11)))
                        if (v9 != v9) {
                            var v12 float64 = (0.0 - 1e10)
                            _ = v12
                        } else {
                            v1["\n\t"] = ("hello" + ("foo bar" + "\n\t"))
                            switch "foo bar" {
                            case "a":
                                fmt.Println(v2, v3)
                            case "\n\t":
                                for fi13 := 0; fi13 < 19; fi13++ {
                                    if (v2 && v2) {
                                        var v14 int16 = 1
                                        fmt.Println(v3, v2)
                                        if v2 {
                                            var v15 int = (fi11 | (fi13 << ((1) & 63)))
                                            _ = v15
                                        } else {
                                            if len(v4) > 0 {
                                                v4[0] = 9223372036854775807
                                            }
                                        }
                                        _ = v14
                                    } else {
                                        v16 := GenS2{F0: 972, F1: "hello"}
                                        switch ("a" + "hello") {
                                        case "foo bar":
                                            v17 := GenS2{F0: 1, F1: "foo bar"}
                                            _ = v17
                                        }
                                        var v18 uint8 = 255
                                        v19 := GenS1{F0: (v18 / ((v18) | 1))}
                                        var v20 uint32 = 2147483648
                                        _ = v16
                                        _ = v18
                                        _ = v19
                                        _ = v20
                                    }
                                    _ = fi13
                                }
                            default:
                                var v21 int32 = 2147483647
                                _ = v21
                            }
                            if (920 <= 863) {
                                v9 = (v9 * v9)
                            }
                        }
                    } else {
                        delete(v1, "")
                        var v22 float32 = 1.5
                        if len(v4) > 0 {
                            v4[0] = (fi11 ^ (v3 | fi11))
                        }
                        _ = v22
                    }
                    v2 = false
                    _ = fi11
                }
                var v23 int64 = 0
                _ = v23
            } else {
                var v24 float32 = 3.14159
                switch -65 {
                case 127:
                    v25 := GenS1{F0: 1}
                    _ = v25
                case 4294967295:
                    v9 += (v9 >> ((935) & 63))
                }
                _ = v24
            }
        } else {
            v26 := GenS2{F0: 255, F1: "a"}
            fmt.Println(v3)
            switch "a" {
            case "hello":
                v3 -= v3
            case "a":
                var v27 uint32 = 2147483648
                _ = v27
            case "foo bar":
                for fi28 := 0; fi28 < 16; fi28++ {
                    _ = len(v4)
                    v2 = true
                    v29 := GenS3{F0: 42, F1: v3, F2: v9}
                    _ = fi28
                    _ = v29
                }
            default:
                v30 := GenS2{F0: 0, F1: "foo bar"}
                _ = v30
            }
            _ = v26
        }
        var v31 int = (v3 | -9223372036854775808)
        _ = v31
    }
    v4 = append(v4, v3)
    v2 = false
    var v32 int16 = -1
    if false {
        if v2 {
            var v33 int16 = v32
            var v34 int64 = 9223372036854775807
            for fi35 := 0; fi35 < 1; fi35++ {
                v3 = (fi35 / ((fi35) | 1))
                v32 += (v32 * v33)
                v36 := GenS3{F0: 1000, F1: (-(fi35)), F2: v9}
                _ = fi35
                _ = v36
            }
            _ = v33
            _ = v34
        }
    }
    v37 := GenS4{F0: 9223372036854775807}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v37
    _ = v4
    _ = v5
    _ = v9
}
