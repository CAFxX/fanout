package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint64; F2 int8 }

func main() {
    var v1 int64 = -1
    var v2 bool = false
    v3 := make(map[uint32]int)
    v4 := make(map[string]string)
    var v5 uint16 = 1
    v2 = (!v2)
    for fi6 := 0; fi6 < 16; fi6++ {
        if false {
            if v2 {
                v1 = v1
                v1 = int64(fi6)
                for fi7 := 0; fi7 < 5; fi7++ {
                    var v8 string = "\n\t"
                    var v9 uint8 = 2
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            }
        } else {
            var v10 string = "hello"
            v11 := v4[v10]
            _ = v10
            _ = v11
        }
        if (2147483647 == 1) {
            v5 = 1
        } else {
            for fi12 := 0; fi12 < 2; fi12++ {
                v13 := GenS1{F0: 0}
                switch int(v5) {
                case 0:
                    fmt.Println(v5)
                }
                var v14 uint8 = 86
                _ = fi12
                _ = v13
                _ = v14
            }
            fmt.Println(v5, fi6, v5)
        }
        _ = fi6
    }
    switch "hello" {
    case "hello":
        for fi15 := 0; fi15 < 1; fi15++ {
            v1 -= v1
            if (!v2) {
                switch 1 {
                case 32767:
                    if (v2 || (v2 || (445 == -1))) {
                        v1 = (v1 + (v1 / ((604) | 1)))
                    }
                default:
                    if false {
                        for fi16 := 0; fi16 < 5; fi16++ {
                            switch fi16 {
                            case 32767:
                                var v17 uint32 = 65535
                                _ = v17
                            case 464:
                                v18 := GenS1{F0: 754}
                                _ = v18
                            case -9223372036854775808:
                                if (!true) {
                                    switch "a" {
                                    case "":
                                        if (!(!(1.5 >= 0.1))) {
                                            if (false || v2) {
                                                var v19 uint16 = v5
                                                _ = v19
                                            } else {
                                                v20 := GenS1{F0: -32768}
                                                _ = v20
                                            }
                                            v5 += v5
                                            var v21 uint32 = 2147483647
                                            _ = v21
                                        }
                                    case "\n\t":
                                        fmt.Println(fi16, fi16)
                                    case "foo bar":
                                        fmt.Println(fi16, fi15)
                                        var v22 string = "hello"
                                        var v23 string = (v22 + "")
                                        _ = v22
                                        _ = v23
                                    default:
                                        v24 := GenS2{F0: 155, F1: v5, F2: 202}
                                        _ = v24
                                    }
                                } else {
                                    v2 = true
                                    v25 := GenS3{F0: uint(v5), F1: 42, F2: 0}
                                    _ = v25
                                }
                            }
                            v2 = (255 < 2)
                            fmt.Println(v2, v5)
                            _ = fi16
                        }
                        for fi26 := 0; fi26 < 15; fi26++ {
                            v1 = 649
                            var v27 uint16 = v5
                            _ = fi26
                            _ = v27
                        }
                    }
                }
            }
            _ = fi15
        }
    case "\n\t":
        v28 := []int{-57, -9223372036854775808}
        _ = v28
    case "a":
        fmt.Println(v5, v1)
    }
    v29 := []int{-2147483648}
    v30 := v3[249]
    var v31 uint8 = 42
    if false {
        if v2 {
            v30 = (v30 % (((v30 / (((-(v30))) | 1))) | 1))
        }
        v4[("" + ("hello" + "\n\t"))] = ""
        v31 -= v31
    }
    v32 := GenS3{F0: 2, F1: 25, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
