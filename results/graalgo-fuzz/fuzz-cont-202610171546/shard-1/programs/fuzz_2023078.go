package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint }

func main() {
    var v1 int64 = 702
    var v2 uint8 = 16
    v3 := make(map[int]string)
    v4 := []int{9223372036854775807}
    _ = len(v3)
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 uint32 = 0
        switch fi5 {
        case -2147483648:
            v1 = (v1 - (v1 | 9223372036854775807))
        case -857:
            if true {
                for fi7 := 0; fi7 < 15; fi7++ {
                    var v8 uint64 = 18446744073709551615
                    switch -2147483648 {
                    case 127:
                        switch "\n\t" {
                        case "hello":
                            if true {
                                v9 := GenS1{F0: (v6 / (((v6 + v6)) | 1)), F1: v8}
                                fmt.Println(v6, v6)
                                v2 = 127
                                _ = v9
                            } else {
                                if true {
                                    var v10 uint64 = v8
                                    _ = v10
                                }
                                var v11 uint = 0
                                _ = v11
                            }
                        case "a":
                            if false {
                                v12 := v3[(^(fi7))]
                                var v13 int = (fi5 & (fi5 & 32767))
                                var v14 uint = 2
                                _ = v12
                                _ = v13
                                _ = v14
                            }
                        case "\n\t":
                            v6 += 0
                        default:
                            v2 += (^(v2))
                        }
                    }
                    if false {
                        var v15 uint64 = v8
                        fmt.Println(v1, v15, fi7)
                        var v16 string = "a"
                        _ = v15
                        _ = v16
                    }
                    _ = fi7
                    _ = v8
                }
                for fi17 := 0; fi17 < 3; fi17++ {
                    v18 := GenS1{F0: v6, F1: 4294967295}
                    v19 := GenS1{F0: (v6 % (((v6 % (((v6 >> ((18446744073709551615) & 63))) | 1))) | 1)), F1: 42}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
        case -273:
            for fi20 := 0; fi20 < 6; fi20++ {
                var v21 float64 = 1e10
                _ = fi20
                _ = v21
            }
        default:
            if (v1 >= 1) {
                v6 = v6
                var v22 uint = 0
                v23 := GenS1{F0: (v6 << ((v22) & 63)), F1: 407}
                _ = v22
                _ = v23
            } else {
                var v24 string = ("\n\t" + "hello")
                _ = v24
            }
        }
        v2 = 127
        _ = fi5
        _ = v6
    }
    if true {
        var v25 uint8 = (v2 & v2)
        for fi26 := 0; fi26 < 13; fi26++ {
            fmt.Println(v2, v2)
            v27 := GenS1{F0: 2147483647, F1: 42}
            var v28 int = fi26
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = v25
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        var v30 uint32 = 2147483647
        _ = fi29
        _ = v30
    }
    if false {
        for fi31 := 0; fi31 < 4; fi31++ {
            v32 := GenS1{F0: 348, F1: 0}
            _ = fi31
            _ = v32
        }
        v2 += (v2 << ((4294967295) & 63))
    } else {
        var v33 string = ("" + ("foo bar" + (("a" + "foo bar") + "hello")))
        _ = v33
    }
    var v34 int32 = 1
    var v35 int32 = (v34 - (v34 | (v34 + (-(v34)))))
    v36 := GenS1{F0: 2147483648, F1: 42}
    v37 := GenS1{F0: 4294967295, F1: 2}
    fmt.Println(v35, v2)
    if false {
        switch "foo bar" {
        case "\n\t":
            if false {
                var v38 bool = false
                _ = v38
            } else {
                fmt.Println(v1)
                _ = len(v3)
                switch "" {
                case "hello":
                    v39 := v3[-341]
                    _ = v39
                }
            }
        case "hello":
            v40 := GenS2{F0: 0}
            _ = v40
        default:
            var v41 uint64 = 18446744073709551615
            _ = v41
        }
        var v42 uint64 = 18446744073709551615
        _ = v42
    }
    v34 = 520
    v2 += (v2 | (v2 | (v2 - v2)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
