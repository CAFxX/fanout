package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 string }

func main() {
    var v1 int = 4294967295
    var v2 bool = (1 >= 504)
    v3 := make(map[string]uint32)
    v4 := []string{"a", "a", "\n\t", "\n\t"}
    v5 := make(map[int]string)
    for fi6 := 0; fi6 < 16; fi6++ {
        if ((fi6 / (((v1 >> ((0) & 63))) | 1)) >= -2147483648) {
            v7 := GenS1{F0: 118}
            for fi8 := 0; fi8 < 3; fi8++ {
                v1 += (v1 + (fi8 ^ fi6))
                var v9 int16 = 1
                _ = fi8
                _ = v9
            }
            var v10 float32 = (1e10 + -2.25)
            _ = v10
            _ = v7
        } else {
            v4 = append(v4, "\n\t")
            var v11 int32 = 589
            v1 = -9223372036854775808
            _ = v11
        }
        var v12 int32 = 0
        if v2 {
            for fi13 := 0; fi13 < 16; fi13++ {
                v1 = 32767
                if v2 {
                    if v2 {
                        v14 := GenS1{F0: 126}
                        _ = v14
                    }
                } else {
                    if v2 {
                        if (-1 <= 514) {
                            v2 = v2
                            var v15 uint16 = 255
                            fmt.Println(v2, v15, v15)
                            _ = v15
                        }
                        var v16 uint8 = 200
                        fmt.Println(v16, fi13, v16)
                        _ = v16
                    } else {
                        if (0 != -412) {
                            var v17 int32 = v12
                            v18 := v3["\n\t"]
                            v2 = (1000 >= 42)
                            _ = v17
                            _ = v18
                        } else {
                            v19 := GenS1{F0: 1}
                            _ = v19
                        }
                        if true {
                            if (!v2) {
                                v20 := GenS1{F0: 4294967295}
                                v21 := GenS2{F0: 4294967295, F1: (18446744073709551615 <= 42), F2: 0}
                                _ = v20
                                _ = v21
                            }
                        }
                    }
                    v4 = append(v4, (((("\n\t" + "a") + "foo bar") + "\n\t") + "\n\t"))
                }
                _ = len(v4)
                v22 := GenS2{F0: 1, F1: v2, F2: 128}
                v2 = (334 != 0)
                _ = fi13
                _ = v22
            }
            v23 := GenS1{F0: 0}
            var v24 uint8 = 128
            _ = v23
            _ = v24
        }
        _ = fi6
        _ = v12
    }
    if (100 < 42) {
        for fi25 := 0; fi25 < 14; fi25++ {
            switch "\n\t" {
            case "a":
                if v2 {
                    v26 := GenS1{F0: 1}
                    _ = v26
                }
            case "\n\t":
                delete(v5, v1)
            case "hello":
                var v27 uint8 = 2
                _ = v27
            }
            v1 -= fi25
            _ = fi25
        }
        if v2 {
            _ = len(v5)
        } else {
            var v28 float64 = (float64((3.14159 * (1e10 - ((0.0 / -2.25) - ((0.0 - 1e10) / -2.25))))) + 0.1)
            _ = len(v4)
            v29 := GenS1{F0: 2}
            _ = v28
            _ = v29
        }
    }
    _ = len(v3)
    delete(v3, ("\n\t" + ""))
    v2 = v2
    if (true || (128 >= 255)) {
        for fi30 := 0; fi30 < 10; fi30++ {
            var v31 int32 = 42
            _ = fi30
            _ = v31
        }
        _ = len(v3)
        v1 = v1
    } else {
        switch "" {
        case "foo bar":
            for fi32 := 0; fi32 < 8; fi32++ {
                v33 := v5[(^(fi32))]
                v34 := GenS3{F0: 31, F1: -1}
                _ = fi32
                _ = v33
                _ = v34
            }
        case "a":
            var v35 int8 = 0
            _ = v35
        default:
            var v36 int64 = 343
            _ = v36
        }
        v37 := GenS4{F0: ("hello" + "\n\t")}
        _ = v37
    }
    for fi38 := 0; fi38 < 11; fi38++ {
        var v39 uint32 = 4294967295
        _ = fi38
        _ = v39
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
