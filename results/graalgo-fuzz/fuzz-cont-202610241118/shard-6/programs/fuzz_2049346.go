package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int }
func genFn0(p0 uint16, p1 string) string {
    fmt.Println(p1, p0)
    v7 := GenS1{F0: (p0 | 32767), F1: 1}
    v8 := GenS1{F0: p0, F1: 367}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p1
}

func main() {
    var v1 string = "a"
    var v2 uint64 = 413
    var v3 int8 = 1
    var v4 string = (v1 + "a")
    v5 := []int{-826, 42}
    var v6 uint64 = v2
    v2 += 961
    var v9 int64 = -427
    if false {
        fmt.Println(v2, v3)
    } else {
        fmt.Println(v4)
    }
    switch v4 {
    case "a":
        if false {
            v10 := GenS1{F0: 0, F1: v6}
            var v11 uint8 = 110
            v2 = (v6 - v2)
            _ = v10
            _ = v11
        }
    case "hello":
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 int64 = v9
            if false {
                switch v1 {
                case "":
                    var v14 int16 = 290
                    _ = v14
                case "\n\t":
                    if false {
                        switch (v4 + "hello") {
                        case "\n\t":
                            v13 += 9223372036854775807
                        case "hello":
                            var v15 uint = 40
                            _ = v15
                        case "a":
                            v9 = v13
                        default:
                            v16 := GenS1{F0: 573, F1: v6}
                            _ = v16
                        }
                    } else {
                        v17 := GenS1{F0: 603, F1: 42}
                        _ = v17
                    }
                case "hello":
                    v1 = "foo bar"
                default:
                    var v18 uint16 = 0
                    _ = v18
                }
                if true {
                    v2 += v2
                    if (1e10 <= 3.14159) {
                        switch (fi12 * fi12) {
                        case -2147483648:
                            v6 -= v6
                        case 127:
                            v19 := GenS2{F0: fi12}
                            _ = v19
                        }
                        var v20 uint8 = 42
                        _ = v20
                    } else {
                        if true {
                            if true {
                                v5 = append(v5, -2147483648)
                            }
                            var v21 uint8 = 42
                            var v22 uint32 = 2147483648
                            _ = v21
                            _ = v22
                        }
                        fmt.Println(v4, v2)
                        var v23 int8 = v3
                        _ = v23
                    }
                    v5 = append(v5, fi12)
                }
            }
            v3 += (v3 * v3)
            _ = fi12
            _ = v13
        }
    case "foo bar":
        if len(v5) > 0 {
            v5[0] = 32768
        }
    default:
        if (255 > 4294967295) {
            for fi24 := 0; fi24 < 3; fi24++ {
                var v25 int = (fi24 * (fi24 | 32767))
                var v26 float32 = (0.1 * (1.5 / (float32(3.14159) - 1.5)))
                _ = fi24
                _ = v25
                _ = v26
            }
            var v27 uint32 = 837
            if true {
                fmt.Println(v2, v6, v2)
            } else {
                fmt.Println(v6)
            }
            _ = v27
        } else {
            _ = len(v5)
            if false {
                fmt.Println(v4, v1)
                var v28 int8 = 1
                switch ((v4 + "foo bar") + "") {
                case "":
                    for fi29 := 0; fi29 < 14; fi29++ {
                        var v30 int = fi29
                        switch 32767 {
                        case -9223372036854775808:
                            if false {
                                v31 := GenS1{F0: 255, F1: (^(v2))}
                                fmt.Println(v2, v1)
                                v2 = (v6 + 766)
                                _ = v31
                            }
                        case 128:
                            v28 = v28
                        }
                        _ = fi29
                        _ = v30
                    }
                case "a":
                    fmt.Println(v9, v3, v6)
                case "foo bar":
                    var v32 string = "\n\t"
                    _ = v32
                }
                _ = v28
            }
            _ = len(v5)
        }
    }
    switch ((v4 + v4) + v4) {
    case "hello":
        fmt.Println(v2, v6)
    case "\n\t":
        v4 = ("foo bar" + ("a" + ""))
    }
    v6 = (v2 + 786)
    v6 = v6
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v3, v3, v2)
    if true {
        v1 = v1
        _ = len(v5)
    }
    var v33 int16 = -1
    v3 += (v3 | -128)
    if len(v5) > 0 {
        v5[0] = -109
    }
    v1 = (v1 + v1)
    for fi34 := 0; fi34 < 5; fi34++ {
        var v35 uint = 42
        var v36 int32 = 205
        _ = fi34
        _ = v35
        _ = v36
    }
    if (true && (!((^(v33)) < v33))) {
        for fi37 := 0; fi37 < 5; fi37++ {
            if (16 <= 2147483647) {
                if len(v5) > 0 {
                    v5[0] = (fi37 / ((42) | 1))
                }
                v5 = append(v5, fi37)
                for fi38 := 0; fi38 < 17; fi38++ {
                    var v39 int16 = (v33 & v33)
                    _ = fi38
                    _ = v39
                }
            } else {
                v6 = v2
                fmt.Println(v33)
                switch "hello" {
                case "\n\t":
                    var v40 uint8 = 0
                    _ = v40
                default:
                    var v41 uint8 = 2
                    _ = v41
                }
            }
            v42 := GenS1{F0: 2, F1: v6}
            v3 -= 51
            _ = fi37
            _ = v42
        }
        var v43 uint16 = uint16(v6)
        for fi44 := 0; fi44 < 19; fi44++ {
            v2 = (^(v6))
            v4 = v4
            _ = fi44
        }
        _ = v43
    }
    v5 = append(v5, int(v2))
    var v45 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v9
}
