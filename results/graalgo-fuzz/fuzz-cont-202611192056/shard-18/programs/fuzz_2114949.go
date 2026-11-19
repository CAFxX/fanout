package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 bool }
type GenS2 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8 }

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 42
    var v3 bool = (true && false)
    var v4 float32 = 0.1
    v5 := make(map[string]int)
    switch 128 {
    case 956:
        for fi6 := 0; fi6 < 13; fi6++ {
            var v7 uint32 = 476
            v8 := GenS1{F0: 18446744073709551615, F1: 42, F2: v3}
            v7 = v7
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    v4 += ((v4 + 0.0) * 0.1)
    v9 := []int{0, 1, 1}
    if true {
        v10 := GenS2{F0: (!(v4 >= (v4 * v4))), F1: 2147483647, F2: 708}
        fmt.Println(v2, v2)
        v11 := GenS1{F0: 0, F1: 507, F2: false}
        _ = v10
        _ = v11
    } else {
        switch "a" {
        case "":
            if (!(1e10 < 1.5)) {
                if v3 {
                    var v12 string = "a"
                    var v13 int = 502
                    v14 := v5[v12]
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    switch 42 {
                    case 128:
                        switch 1 {
                        case 1:
                            if (167 != 129) {
                                v2 = v2
                            }
                        default:
                            fmt.Println(v3)
                        }
                    case 32767:
                        var v15 int32 = -2147483648
                        _ = v15
                    case 1:
                        switch 2147483647 {
                        case 32768:
                            _ = len(v9)
                        case 2147483647:
                            fmt.Println(v4)
                        case 42:
                            fmt.Println(v3, v4)
                        default:
                            if (!((32767 > -1) && ((2147483647 <= -996) && v3))) {
                                for fi16 := 0; fi16 < 14; fi16++ {
                                    if ((v2 >> ((455) & 63)) == (^(v2))) {
                                        var v17 uint64 = 381
                                        _ = v17
                                    }
                                    _ = fi16
                                }
                            } else {
                                for fi18 := 0; fi18 < 16; fi18++ {
                                    v3 = (!true)
                                    v3 = v3
                                    _ = fi18
                                }
                            }
                        }
                    }
                }
                v19 := v5["hello"]
                _ = v19
            }
        case "a":
            var v20 float32 = (v4 + v4)
            _ = v20
        case "hello":
            var v21 float32 = v4
            _ = v21
        }
        var v22 uint = 42
        v5[("" + "a")] = 0
        _ = v22
    }
    v2 -= 32767
    v4 -= 0.0
    v23 := GenS2{F0: v3, F1: 2147483648, F2: 42}
    delete(v5, "hello")
    for fi24 := 0; fi24 < 4; fi24++ {
        v25 := GenS2{F0: v3, F1: 395, F2: 65535}
        for fi26 := 0; fi26 < 16; fi26++ {
            if v3 {
                switch (fi26 | fi26) {
                case 4294967295:
                    var v27 uint16 = (^(v2))
                    _ = v27
                default:
                    v28 := GenS2{F0: (2 != 416), F1: 42, F2: 387}
                    _ = v28
                }
            }
            _ = fi26
        }
        _ = fi24
        _ = v25
    }
    v29 := GenS3{F0: (v2 & v2)}
    v30 := GenS4{F0: -128}
    var v31 int8 = 113
    fmt.Println(v4, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v9
}
