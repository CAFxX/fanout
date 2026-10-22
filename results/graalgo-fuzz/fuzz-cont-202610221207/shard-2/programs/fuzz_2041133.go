package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int32; F2 int64 }
type GenS3 struct { F0 int8; F1 string }
func genFn0(p0 uint32, p1 uint32) int32 {
    v5 := []int{9223372036854775807}
    v6 := GenS1{F0: 0}
    var v7 int = 1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := []string{"a", "\n\t", "a", "a"}
    var v2 float32 = 1.5
    var v3 string = "a"
    var v4 uint64 = 224
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        v4 = 601
        var v8 uint = 76
        switch (("\n\t" + v3) + v3) {
        case "a":
            v9 := GenS1{F0: (v4 * (v4 & 42))}
            _ = v9
        default:
            if false {
                _ = len(v1)
                v4 = 2
                _ = len(v1)
            } else {
                v8 = (v8 + v8)
            }
        }
        _ = v8
    }
    v2 = (3.14159 + -2.25)
    var v10 float64 = 0.1
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 string = v3
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int16 = 736
        var v15 uint64 = v4
        if (3.14159 == ((0.0 / v2) / (float32(v10) - (((v2 - v2) / 3.14159) * (v2 * (((v2 - 3.14159) / -2.25) - v2)))))) {
            var v16 int = (fi13 % (((fi13 - fi13)) | 1))
            v17 := GenS2{F0: -1, F1: 720, F2: -9223372036854775808}
            if (34 < 2) {
                if (false || (true || false)) {
                    if (!false) {
                        var v18 uint = uint(v15)
                        for fi19 := 0; fi19 < 12; fi19++ {
                            v20 := GenS1{F0: (v15 % ((v15) | 1))}
                            _ = fi19
                            _ = v20
                        }
                        _ = v18
                    }
                    v21 := GenS3{F0: 25, F1: v3}
                    _ = v21
                } else {
                    if true {
                        v15 -= (v15 | v4)
                        if false {
                            if (v14 > v14) {
                                v22 := GenS1{F0: 457}
                                _ = v22
                            } else {
                                switch v3 {
                                case "foo bar":
                                    if (false && true) {
                                        var v23 uint64 = (v15 * 18446744073709551615)
                                        v4 += v15
                                        _ = v23
                                    }
                                case "hello":
                                    var v24 uint16 = 1
                                    _ = v24
                                case "":
                                    for fi25 := 0; fi25 < 1; fi25++ {
                                        fmt.Println(v10)
                                        _ = fi25
                                    }
                                }
                                for fi26 := 0; fi26 < 1; fi26++ {
                                    if true {
                                        if false {
                                            var v27 uint64 = (v4 - v4)
                                            switch (v3 + v3) {
                                            case "hello":
                                                v16 += -9223372036854775808
                                            case "":
                                                v27 += v27
                                            case "a":
                                                v15 = (v27 >> ((0) & 63))
                                            }
                                            _ = v27
                                        }
                                    } else {
                                        var v28 uint64 = 2
                                        _ = v28
                                    }
                                    _ = fi26
                                }
                            }
                        } else {
                            fmt.Println(v10)
                            _ = len(v1)
                        }
                        var v29 bool = true
                        _ = v29
                    }
                    fmt.Println(v4, v15, v14)
                    v14 += v14
                }
                v1 = append(v1, (v3 + ""))
            }
            _ = v16
            _ = v17
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        var v30 int64 = -1
        _ = v30
    case "a":
        if len(v1) > 0 {
            v1[0] = (("a" + v3) + (v3 + "foo bar"))
        }
    case "hello":
        v10 -= (float64((v2 / v2)) * (v10 + (float64(v2) / v10)))
    }
    v2 = (((v2 - v2) * (v2 + v2)) - 1.5)
    switch "" {
    case "a":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v4 += v4
    v10 -= v10
    v3 = v3
    v4 += v4
    if (((v10 - v10) * (((v10 - v10) / 3.14159) + 3.14159)) > v10) {
        fmt.Println(v2, v2)
        _ = len(v1)
    }
    var v31 uint64 = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}
