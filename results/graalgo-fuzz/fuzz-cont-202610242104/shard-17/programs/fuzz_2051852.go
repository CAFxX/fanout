package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int8 }
func genFn0() uint8 {
    var v4 uint32 = 2147483648
    v5 := GenS1{F0: 1000, F1: 42, F2: 128}
    v6 := []int{-22, 127}
    if true {
        if (2 == 255) {
            _ = len(v6)
        }
        if true {
            switch 0 {
            case -2147483648:
                if false {
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v4)
                    case "":
                        var v7 bool = (("hello" <= "a") || (2 > 962))
                        _ = v7
                    default:
                        var v8 uint8 = 0
                        _ = v8
                    }
                } else {
                    var v9 string = "\n\t"
                    _ = v9
                }
            }
        }
        var v10 uint64 = 18446744073709551615
        _ = v10
    }
    if len(v6) > 0 {
        v6[0] = 1
    }
    _ = v4
    _ = v5
    _ = v6
    return 39
}

func main() {
    var v1 int64 = 0
    var v2 int64 = (v1 % (((^(v1))) | 1))
    v3 := make(map[uint32]int)
    v11 := []int{-357, 4294967295, 32767, 32767}
    v2 = 9223372036854775807
    v1 = (^(v2))
    if false {
        if true {
            if true {
                if true {
                    v1 = (v1 ^ v2)
                    for fi12 := 0; fi12 < 20; fi12++ {
                        fmt.Println(fi12, fi12)
                        if false {
                            switch "" {
                            case "\n\t":
                                var v13 int8 = 100
                                _ = v13
                            case "":
                                v14 := GenS1{F0: 1, F1: 1, F2: 42}
                                _ = v14
                            }
                        } else {
                            _ = len(v11)
                            switch fi12 {
                            case 9223372036854775807:
                                v15 := GenS2{F0: 0, F1: (1 == 2), F2: 42}
                                _ = v15
                            case 42:
                                v16 := GenS2{F0: 4294967295, F1: true, F2: 1}
                                _ = v16
                            }
                        }
                        fmt.Println(v1, fi12)
                        _ = fi12
                    }
                }
            } else {
                switch ("" + "hello") {
                case "hello":
                    if false {
                        var v17 uint64 = 42
                        _ = v17
                    } else {
                        delete(v3, 21)
                        v2 -= 0
                        var v18 uint = 2
                        _ = v18
                    }
                case "foo bar":
                    if false {
                        switch "\n\t" {
                        case "foo bar":
                            var v19 uint = 868
                            _ = v19
                        default:
                            v2 -= (v2 / (((^(v1))) | 1))
                        }
                        v20 := GenS2{F0: 4294967295, F1: (!(-414 >= (^(v2)))), F2: 0}
                        if len(v11) > 0 {
                            v11[0] = -399
                        }
                        _ = v20
                    }
                case "a":
                    var v21 int64 = v1
                    _ = v21
                default:
                    v1 = v2
                }
                switch "a" {
                case "":
                    var v22 uint32 = 2147483647
                    _ = v22
                case "hello":
                    var v23 int16 = 32767
                    _ = v23
                case "\n\t":
                    var v24 int16 = 1
                    _ = v24
                }
            }
        }
        var v25 int8 = -128
        var v26 uint16 = 418
        _ = v25
        _ = v26
    } else {
        v27 := GenS3{F0: 0, F1: -97}
        _ = v27
    }
    var v28 int64 = (v1 * (-(v2)))
    fmt.Println(v1)
    v29 := GenS2{F0: 2, F1: true, F2: 4294967295}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v2
    _ = v28
    _ = v29
    _ = v3
}
