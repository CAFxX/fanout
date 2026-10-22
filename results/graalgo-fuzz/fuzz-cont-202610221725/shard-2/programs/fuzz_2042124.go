package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 bool }
func genFn0(p0 string) uint16 {
    var v5 uint = 4294967295
    for fi6 := 0; fi6 < 11; fi6++ {
        if true {
            for fi7 := 0; fi7 < 20; fi7++ {
                switch 4294967295 {
                case 1:
                    v8 := []int{}
                    _ = v8
                case 4294967295:
                    switch int(v5) {
                    case -9223372036854775808:
                        fmt.Println(fi7, v5, fi7)
                    case 127:
                        switch (fi6 | fi6) {
                        case -9223372036854775808:
                            if ((v5 * (v5 / (((v5 / ((v5) | 1))) | 1))) > v5) {
                                var v9 uint = (^(v5))
                                if false {
                                    fmt.Println(fi6)
                                } else {
                                    fmt.Println(p0, p0)
                                    if ((false || false) || true) {
                                        if (!(!(1 >= 0))) {
                                            if false {
                                                v10 := GenS1{F0: 1000, F1: -1}
                                                v11 := GenS1{F0: -876, F1: 0}
                                                switch fi6 {
                                                case 1:
                                                    v12 := GenS1{F0: -32768, F1: 42}
                                                    _ = v12
                                                case -1:
                                                    switch (p0 + "hello") {
                                                    case "":
                                                        var v13 bool = (true || (true || (!false)))
                                                        _ = v13
                                                    default:
                                                        fmt.Println(fi6, v5)
                                                    }
                                                case 0:
                                                    v14 := GenS2{F0: 918, F1: -101}
                                                    _ = v14
                                                }
                                                _ = v10
                                                _ = v11
                                            }
                                            if true {
                                                v5 += v5
                                            }
                                            p0 = ""
                                            v5 -= v5
                                        }
                                        fmt.Println(fi6, v5, fi6)
                                        fmt.Println(fi7)
                                    }
                                    var v15 int8 = -128
                                    var v16 int = fi7
                                    _ = v15
                                    _ = v16
                                }
                                v17 := GenS3{F0: (305 != 1)}
                                _ = v17
                                _ = v9
                            } else {
                                p0 = ((p0 + p0) + (p0 + p0))
                                v5 = v5
                            }
                        case 9223372036854775807:
                            if true {
                                fmt.Println(p0)
                            }
                        case -1:
                            v18 := []int{}
                            _ = v18
                        }
                    case 32767:
                        v19 := GenS2{F0: 881, F1: 0}
                        _ = v19
                    }
                case 2147483647:
                    v5 -= 427
                default:
                    var v20 uint32 = uint32(v5)
                    _ = v20
                }
                _ = fi7
            }
            p0 = p0
            v21 := GenS3{F0: true}
            _ = v21
        } else {
            v5 += v5
            v22 := GenS2{F0: 0, F1: 1}
            fmt.Println(v5, fi6)
            _ = v22
        }
        v23 := GenS2{F0: 255, F1: -1}
        _ = fi6
        _ = v23
    }
    v24 := []int{}
    _ = p0
    _ = v24
    _ = v5
    return 65535
}

func main() {
    var v1 int = 0
    var v2 string = "foo bar"
    v3 := make(map[uint32]uint32)
    var v4 int = (v1 / ((v1) | 1))
    v2 = v2
    var v25 uint32 = 2
    for fi26 := 0; fi26 < 13; fi26++ {
        v27 := []int{(fi26 / ((fi26) | 1)), (v4 << ((789) & 63)), fi26}
        v25 = (v25 | (v25 | 2))
        _ = fi26
        _ = v27
    }
    v28 := []int{(-(v4)), v1, -2147483648}
    var v29 int = 4294967295
    v30 := GenS1{F0: -1, F1: 0}
    var v31 float32 = 3.14159
    _ = len(v28)
    v25 = 1
    var v32 uint64 = 812
    fmt.Println(v25, v31)
    if ((v25 + (v25 >> ((129) & 63))) == v25) {
        v32 -= 4294967295
        fmt.Println(v1, v1)
        v33 := GenS2{F0: genFn0(v2), F1: 927}
        _ = v33
    }
    v34 := GenS3{F0: (v32 <= (v32 % ((618) | 1)))}
    var v35 uint64 = 918
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v4
}
