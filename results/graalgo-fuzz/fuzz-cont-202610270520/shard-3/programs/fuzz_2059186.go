package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint16 = 0
    var v2 float32 = 1.5
    var v3 uint = 18446744073709551615
    v4 := make(map[uint32]bool)
    delete(v4, 1)
    v2 += v2
    if true {
        var v5 uint64 = 389
        _ = v5
    } else {
        switch "a" {
        case "a":
            for fi6 := 0; fi6 < 7; fi6++ {
                v2 -= (-2.25 - v2)
                _ = fi6
            }
        }
        var v7 uint16 = (v1 - (v1 % ((v1) | 1)))
        _ = v7
    }
    v8 := []int{827}
    if (false || true) {
        var v9 int64 = int64(v3)
        _ = v9
    } else {
        var v10 uint = v3
        v10 = 82
        _ = v10
    }
    var v11 string = ("\n\t" + ("\n\t" + ""))
    fmt.Println(v2, v3, v3)
    fmt.Println(v1)
    if (42 != 373) {
        if (uint64(v3) >= 0) {
            v12 := GenS1{F0: 1}
            v13 := GenS1{F0: -886}
            if (!false) {
                v1 -= (v1 % ((2) | 1))
                if (1e10 <= 1.5) {
                    if (441 < 0) {
                        var v14 int = 32768
                        var v15 uint32 = 98
                        _ = v14
                        _ = v15
                    }
                    var v16 int32 = 42
                    _ = v16
                }
            } else {
                if true {
                    v11 = v11
                } else {
                    for fi17 := 0; fi17 < 1; fi17++ {
                        if true {
                            var v18 int16 = 32767
                            if true {
                                _ = len(v8)
                            } else {
                                v19 := GenS1{F0: 240}
                                var v20 int8 = 102
                                switch v11 {
                                case "\n\t":
                                    if true {
                                        var v21 float32 = 1e10
                                        var v22 int16 = 42
                                        if true {
                                            if false {
                                                var v23 bool = false
                                                _ = v23
                                            } else {
                                                for fi24 := 0; fi24 < 1; fi24++ {
                                                    var v25 int64 = -9223372036854775808
                                                    _ = len(v4)
                                                    v21 = v2
                                                    _ = fi24
                                                    _ = v25
                                                }
                                                _ = len(v4)
                                                var v26 string = "\n\t"
                                                _ = v26
                                            }
                                        }
                                        _ = v21
                                        _ = v22
                                    } else {
                                        fmt.Println(v20)
                                        v11 = v11
                                    }
                                case "hello":
                                    v27 := GenS1{F0: 117}
                                    _ = v27
                                case "":
                                    fmt.Println(v1, v1, v18)
                                default:
                                    var v28 float32 = v2
                                    _ = v28
                                }
                                _ = v19
                                _ = v20
                            }
                            _ = v18
                        }
                        fmt.Println(fi17)
                        switch fi17 {
                        case 91:
                            fmt.Println(v1, v3)
                        case 32767:
                            v2 = (1e10 / v2)
                        case 128:
                            var v29 int = 2147483647
                            _ = v29
                        }
                        _ = fi17
                    }
                    if (100 == 0) {
                        v30 := GenS1{F0: 375}
                        _ = v30
                    }
                    fmt.Println(v11, v2)
                }
            }
            _ = v12
            _ = v13
        }
        var v31 uint = (v3 / ((v3) | 1))
        v31 = 2
        _ = v31
    }
    var v32 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v8
}
