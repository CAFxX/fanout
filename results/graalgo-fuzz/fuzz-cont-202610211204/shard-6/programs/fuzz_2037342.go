package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 bool }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint8 = 2
    var v3 uint32 = 830
    var v4 float32 = (3.14159 - 0.1)
    v5 := make(map[int]bool)
    v4 = v4
    v6 := GenS1{F0: "hello", F1: true, F2: true}
    fmt.Println(v4)
    for fi7 := 0; fi7 < 20; fi7++ {
        if true {
            v8 := GenS1{F0: (("a" + "hello") + ("hello" + "foo bar")), F1: false, F2: true}
            if false {
                for fi9 := 0; fi9 < 3; fi9++ {
                    if true {
                        switch "a" {
                        case "foo bar":
                            var v10 int = (fi7 & (fi9 - (fi9 + fi7)))
                            _ = v10
                        case "\n\t":
                            switch 0 {
                            case 127:
                                switch fi9 {
                                case 127:
                                    switch "hello" {
                                    case "foo bar":
                                        var v11 int64 = 9223372036854775807
                                        _ = v11
                                    case "hello":
                                        if false {
                                            v12 := GenS1{F0: ((("\n\t" + "") + ("" + ("" + "foo bar"))) + "a"), F1: false, F2: true}
                                            if false {
                                                if false {
                                                    v13 := []int{fi7, (-(fi7)), (fi9 % (((fi9 | fi9)) | 1)), 128}
                                                    _ = v13
                                                }
                                                var v14 uint64 = 1
                                                _ = v14
                                            }
                                            var v15 int16 = 1
                                            _ = v12
                                            _ = v15
                                        }
                                    }
                                case -2147483648:
                                    v16 := GenS2{F0: v1}
                                    _ = v16
                                default:
                                    v17 := GenS1{F0: "", F1: false, F2: true}
                                    _ = v17
                                }
                            }
                        }
                        v3 -= v3
                    }
                    _ = fi9
                }
            } else {
                var v18 uint8 = 0
                var v19 uint32 = (v3 - v3)
                for fi20 := 0; fi20 < 9; fi20++ {
                    v2 = (v2 / (((^(v18))) | 1))
                    v2 = 42
                    _ = fi20
                }
                _ = v18
                _ = v19
            }
            _ = v8
        }
        v21 := v5[fi7]
        _ = fi7
        _ = v21
    }
    var v22 uint64 = 441
    var v23 float64 = (0.0 + 1.5)
    v24 := GenS2{F0: uint64(v3)}
    if (-1 < -1) {
        var v25 bool = true
        fmt.Println(v2, v4, v4)
        _ = v25
    } else {
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
