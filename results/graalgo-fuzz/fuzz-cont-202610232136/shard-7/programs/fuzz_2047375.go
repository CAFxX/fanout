package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint8; F2 int64 }

func main() {
    var v1 int = 42
    v2 := []int{(v1 & v1)}
    var v3 float32 = 0.1
    v4 := []string{"a", "", "a", "\n\t"}
    v5 := make(map[uint32]string)
    if (true || false) {
        v6 := GenS1{F0: v1}
        _ = v6
    } else {
        switch "foo bar" {
        case "a":
            v7 := v5[382]
            _ = v7
        case "\n\t":
            if (-9223372036854775808 >= 42) {
                var v8 uint8 = 31
                v1 += v1
                _ = v8
            }
        default:
            fmt.Println(v1, v3)
        }
        fmt.Println(v3)
        if len(v4) > 0 {
            v4[0] = ("hello" + "hello")
        }
    }
    var v9 float64 = 1e10
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 int8 = -116
        switch (^(v1)) {
        case 128:
            switch "a" {
            case "a":
                switch "" {
                case "\n\t":
                    var v12 int32 = 42
                    _ = v12
                case "foo bar":
                    for fi13 := 0; fi13 < 9; fi13++ {
                        var v14 uint8 = 151
                        if false {
                            v3 = v3
                            if false {
                                if false {
                                    v15 := GenS1{F0: fi10}
                                    v11 -= v11
                                    var v16 uint16 = 2
                                    _ = v15
                                    _ = v16
                                }
                                fmt.Println(v1, v14)
                            }
                            var v17 uint64 = 1
                            _ = v17
                        }
                        v11 = (v11 * (v11 + (v11 * v11)))
                        _ = fi13
                        _ = v14
                    }
                default:
                    if false {
                        v18 := GenS1{F0: 9223372036854775807}
                        _ = v18
                    } else {
                        var v19 int = (v1 + (v1 * (^(v1))))
                        _ = v19
                    }
                }
            case "":
                v20 := v5[2147483647]
                _ = v20
            case "foo bar":
                switch (("foo bar" + "hello") + "foo bar") {
                case "foo bar":
                    var v21 int64 = 9223372036854775807
                    _ = v21
                case "":
                    var v22 bool = false
                    _ = v22
                }
            default:
                for fi23 := 0; fi23 < 7; fi23++ {
                    fmt.Println(fi23, v9)
                    fmt.Println(v1, fi23)
                    _ = len(v4)
                    _ = fi23
                }
            }
        case 0:
            v1 -= fi10
        case 9223372036854775807:
            if true {
                for fi24 := 0; fi24 < 8; fi24++ {
                    v3 += v3
                    _ = len(v5)
                    _ = len(v4)
                    v25 := GenS1{F0: 4294967295}
                    fmt.Println(v1)
                    _ = fi24
                    _ = v25
                }
                if len(v4) > 0 {
                    v4[0] = ""
                }
                v26 := GenS1{F0: (fi10 | v1)}
                _ = v26
            } else {
                fmt.Println(v11, v11, v3)
            }
        }
        v4 = append(v4, "")
        _ = fi10
        _ = v11
    }
    v27 := GenS2{F0: 2, F1: 119, F2: -1}
    for fi28 := 0; fi28 < 7; fi28++ {
        v29 := v5[282]
        _ = fi28
        _ = v29
    }
    v9 = v9
    var v30 uint32 = 641
    v30 += 896
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v9
}
