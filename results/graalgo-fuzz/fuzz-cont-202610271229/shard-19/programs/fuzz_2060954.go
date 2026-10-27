package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
func genFn0(p0 uint64, p1 uint8, p2 uint16) uint8 {
    fmt.Println(p2, p1)
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 float64 = (1.5 - (0.1 + 0.1))
    var v2 float32 = 3.14159
    v3 := []int{9223372036854775807}
    v4 := make(map[string]bool)
    var v5 int32 = 2147483647
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 int8 = -116
        for fi8 := 0; fi8 < 2; fi8++ {
            switch "hello" {
            case "foo bar":
                var v9 uint64 = 1
                _ = v9
            }
            fmt.Println(v5, v1)
            v2 = float32((float64(v2) + float64((v2 / v2))))
            v1 += v1
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    var v10 int8 = -128
    delete(v4, "hello")
    var v11 int = 127
    fmt.Println(v5, v2, v10)
    if (786 < 65535) {
        v4["a"] = false
    } else {
        v1 -= v1
    }
    switch 4294967295 {
    case 2147483647:
        if ((-(v11)) != v11) {
            if false {
                v12 := GenS1{F0: 103, F1: 605}
                v3 = append(v3, (v11 - v11))
                var v13 string = ""
                _ = v12
                _ = v13
            }
            if true {
                if false {
                    v11 += 32767
                    if true {
                        var v14 uint64 = 18446744073709551615
                        delete(v4, "\n\t")
                        if false {
                            if false {
                                fmt.Println(v1)
                                v11 = v11
                                var v15 float32 = v2
                                _ = v15
                            }
                        }
                        _ = v14
                    } else {
                        var v16 int16 = 661
                        if (255 >= 383) {
                            var v17 float32 = (v2 - 0.1)
                            _ = v17
                        } else {
                            switch 499 {
                            case 32768:
                                v4[""] = (false && (42 >= 4294967295))
                            default:
                                var v18 uint32 = 0
                                _ = v18
                            }
                            var v19 int64 = -1
                            _ = v19
                        }
                        _ = v16
                    }
                    var v20 uint = 1
                    _ = v20
                } else {
                    if (!false) {
                        var v21 uint8 = 84
                        _ = v21
                    }
                }
                for fi22 := 0; fi22 < 2; fi22++ {
                    v5 = 1
                    v10 = (v10 * int8(v5))
                    _ = fi22
                }
            } else {
                v10 = v10
            }
            for fi23 := 0; fi23 < 16; fi23++ {
                var v24 uint16 = 42
                _ = fi23
                _ = v24
            }
        }
    case 9223372036854775807:
        if false {
            if true {
                var v25 bool = false
                _ = v25
            } else {
                v26 := GenS1{F0: 1, F1: 4294967295}
                _ = v26
            }
            v27 := GenS1{F0: -9223372036854775808, F1: 18446744073709551615}
            if (32767 <= 25) {
                fmt.Println(v11, v10, v11)
                v2 = v2
                switch "hello" {
                case "a":
                    fmt.Println(v2, v1)
                }
            } else {
                var v28 uint8 = 127
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v29 uint32 = 799
                _ = v28
                _ = v29
            }
            _ = v27
        }
    }
    v2 = v2
    v30 := GenS1{F0: 1, F1: 4294967295}
    v2 = 0.1
    var v31 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
