package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int32; F2 int16 }

func main() {
    var v1 int8 = -1
    var v2 uint16 = 0
    var v3 float64 = 1e10
    v4 := make(map[string]uint32)
    var v5 float64 = ((v3 * v3) - (v3 - (float64((1.5 / 3.14159)) + v3)))
    v6 := []string{("" + "hello")}
    var v7 int64 = -9223372036854775808
    if (0 >= 1) {
        v6 = append(v6, "hello")
    } else {
        v7 = v7
        v8 := GenS1{F0: 117}
        for fi9 := 0; fi9 < 10; fi9++ {
            if true {
                var v10 int16 = 32767
                _ = v10
            } else {
                for fi11 := 0; fi11 < 13; fi11++ {
                    var v12 int8 = -128
                    var v13 uint64 = 221
                    var v14 uint64 = v13
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
                var v15 uint32 = 65535
                v1 += (^(v1))
                _ = v15
            }
            switch ("a" + "a") {
            case "":
                for fi16 := 0; fi16 < 12; fi16++ {
                    fmt.Println(fi16, v5)
                    var v17 bool = (!((2147483647 < 0) && false))
                    switch "" {
                    case "foo bar":
                        if v17 {
                            if len(v6) > 0 {
                                v6[0] = "hello"
                            }
                            v3 += v5
                            if false {
                                switch fi16 {
                                case 1:
                                    fmt.Println(v17)
                                case -727:
                                    v7 -= v7
                                case -914:
                                    v17 = v17
                                default:
                                    fmt.Println(v7)
                                }
                                fmt.Println(v2, fi9)
                            } else {
                                var v18 int8 = v1
                                if v17 {
                                    _ = len(v6)
                                    var v19 uint16 = v2
                                    v5 += float64(0.1)
                                    _ = v19
                                } else {
                                    var v20 int16 = 739
                                    v21 := GenS1{F0: 255}
                                    delete(v4, "")
                                    _ = v20
                                    _ = v21
                                }
                                v1 = 0
                                v3 = v5
                                _ = v18
                            }
                        } else {
                            var v22 uint64 = 0
                            fmt.Println(fi16)
                            _ = v22
                        }
                    case "\n\t":
                        v23 := GenS1{F0: 0}
                        _ = v23
                    case "hello":
                        fmt.Println(v5)
                    default:
                        v7 -= 1
                    }
                    if v17 {
                        if ((float32(v5) - 0.1) != 1.5) {
                            v2 = 934
                        }
                    }
                    v6 = append(v6, ("" + ("a" + "\n\t")))
                    var v24 int64 = 9223372036854775807
                    _ = fi16
                    _ = v17
                    _ = v24
                }
            case "\n\t":
                var v25 int = 128
                _ = v25
            }
            _ = fi9
        }
        _ = v8
    }
    v5 = -2.25
    switch "a" {
    case "foo bar":
        v26 := v4[""]
        _ = v26
    case "":
        fmt.Println(v3, v7, v7)
    case "hello":
        v2 = v2
    }
    switch 1 {
    case 0:
        v3 += float64(3.14159)
    case 42:
        var v27 uint = 65535
        _ = v27
    }
    fmt.Println(v7)
    v28 := GenS2{F0: "\n\t", F1: 1, F2: -32768}
    if len(v6) > 0 {
        v6[0] = "a"
    }
    v5 += (0.0 * v5)
    var v29 int32 = -2147483648
    var v30 int16 = 0
    var v31 uint64 = 1
    var v32 int16 = 1
    for fi33 := 0; fi33 < 15; fi33++ {
        v34 := GenS1{F0: 65535}
        _ = fi33
        _ = v34
    }
    var v35 uint = 2
    for fi36 := 0; fi36 < 1; fi36++ {
        v4["foo bar"] = 4294967295
        v4[("hello" + "hello")] = 787
        var v37 uint8 = 1
        _ = fi36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
