package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int }

func main() {
    v1 := []uint32{2147483648, 4294967295}
    v2 := make(map[string]string)
    var v3 uint32 = 343
    v4 := make(map[uint32]int)
    v5 := []uint32{v3, v3}
    v3 = v3
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = len(v5)
    var v6 uint8 = 42
    var v7 int8 = 99
    v7 -= (-(v7))
    var v8 uint8 = v6
    v9 := v2[("a" + ("" + "a"))]
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(v3)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (1 <= 1) {
            if false {
                v7 -= v7
            } else {
                var v11 int32 = 2147483647
                switch v9 {
                case "foo bar":
                    var v12 uint = 1
                    _ = v12
                case "\n\t":
                    var v13 bool = (false || (65535 < 355))
                    _ = v13
                case "":
                    if (4294967295 == v3) {
                        var v14 string = (v9 + "foo bar")
                        _ = v14
                    } else {
                        delete(v4, v3)
                        var v15 int64 = -9223372036854775808
                        switch "" {
                        case "":
                            v7 = (v7 * (v7 & (v7 - (v7 >> ((4294967295) & 63)))))
                        case "\n\t":
                            for fi16 := 0; fi16 < 17; fi16++ {
                                _ = len(v2)
                                if false {
                                    fmt.Println(fi16)
                                    v11 = (v11 & v11)
                                } else {
                                    v3 = 988
                                    var v17 int64 = v15
                                    var v18 int32 = v11
                                    _ = v17
                                    _ = v18
                                }
                                _ = fi16
                            }
                        }
                        _ = v15
                    }
                }
                var v19 uint64 = uint64(v3)
                _ = v11
                _ = v19
            }
        } else {
            for fi20 := 0; fi20 < 9; fi20++ {
                v9 = "hello"
                var v21 uint64 = 260
                delete(v4, (v3 * v3))
                _ = fi20
                _ = v21
            }
        }
        _ = fi10
    }
    switch -2147483648 {
    case 9223372036854775807:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if (true || (!false)) {
        var v22 uint32 = v3
        var v23 float32 = 1.5
        _ = v22
        _ = v23
    } else {
        for fi24 := 0; fi24 < 5; fi24++ {
            var v25 uint = 215
            _ = fi24
            _ = v25
        }
        _ = len(v2)
    }
    if (0 >= 1000) {
        var v26 string = v9
        if (true && ((v3 ^ 1) == (v3 % ((v3) | 1)))) {
            var v27 string = v9
            v28 := v2[(v27 + ("hello" + v9))]
            var v29 uint64 = 42
            _ = v27
            _ = v28
            _ = v29
        }
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
        _ = v26
    }
    v6 += v8
    var v30 float32 = 0.1
    fmt.Println(v9, v30, v8)
    var v31 float32 = v30
    v32 := GenS1{F0: (true || true), F1: v9, F2: 4294967295}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
