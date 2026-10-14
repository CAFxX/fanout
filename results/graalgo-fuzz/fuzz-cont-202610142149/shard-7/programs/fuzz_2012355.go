package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 string }
type GenS2 struct { F0 int; F1 bool; F2 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := []int{-107, -2147483648, 1}
    var v2 uint = 237
    v3 := make(map[string]string)
    var v4 float64 = 0.0
    var v5 int64 = 42
    switch 32768 {
    case 128:
        fmt.Println(v2, v2, v5)
    default:
        if (((0.1 - 3.14159) * -2.25) == 0.0) {
            var v6 int32 = 1
            _ = v6
        } else {
            if (!(((!false) && true) && (100 == 23))) {
                v2 += (v2 * (v2 - (v2 | v2)))
                if true {
                    var v7 uint32 = 763
                    var v8 int32 = 0
                    v7 -= 2
                    _ = v7
                    _ = v8
                } else {
                    if true {
                        var v9 int64 = 0
                        var v10 uint = v2
                        _ = v10
                        _ = v9
                    } else {
                        _ = len(v3)
                        var v11 int32 = 2147483647
                        fmt.Println(v11)
                        _ = v11
                    }
                }
                for fi12 := 0; fi12 < 3; fi12++ {
                    if false {
                        for fi13 := 0; fi13 < 15; fi13++ {
                            switch fi13 {
                            case -1:
                                fmt.Println(fi13)
                            case -41:
                                var v14 uint64 = 18446744073709551615
                                _ = v14
                            case 762:
                                var v15 int8 = -1
                                _ = v15
                            }
                            v16 := GenS1{F0: ("hello" + "\n\t"), F1: ("foo bar" + ("a" + "hello")), F2: "a"}
                            if false {
                                var v17 int = fi12
                                var v18 int32 = -2147483648
                                fmt.Println(v18, v2)
                                _ = v17
                                _ = v18
                            }
                            _ = fi13
                            _ = v16
                        }
                    }
                    fmt.Println(v5, v2)
                    v19 := GenS2{F0: (fi12 ^ fi12), F1: true, F2: ((0.0 - (0.1 + float64((1e10 + 0.0)))) == v4)}
                    _ = fi12
                    _ = v19
                }
            }
        }
    }
    switch "\n\t" {
    case "a":
        var v20 uint32 = 0
        _ = v20
    case "\n\t":
        if false {
            if (!(4294967295 >= uint32(v2))) {
                var v21 int16 = -652
                fmt.Println(v5, v4)
                var v22 string = ("\n\t" + "")
                _ = v21
                _ = v22
            } else {
                v5 = v5
                if false {
                    v23 := GenS3{F0: 78}
                    _ = len(v1)
                    _ = len(v3)
                    _ = v23
                } else {
                    v24 := GenS4{F0: 0}
                    _ = v24
                }
            }
            var v25 uint16 = 276
            v26 := v3[(("\n\t" + (("" + "\n\t") + "\n\t")) + "")]
            _ = v25
            _ = v26
        }
    }
    v2 += 634
    if (1 > -1) {
        v27 := v3[""]
        var v28 int32 = -1
        _ = v27
        _ = v28
    }
    var v29 uint = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        for fi31 := 0; fi31 < 19; fi31++ {
            if true {
                v29 += v29
                v5 -= v5
                v32 := GenS3{F0: 0}
                _ = v32
            }
            fmt.Println(v4, fi31)
            _ = fi31
        }
        _ = fi30
    }
    var v33 float64 = -2.25
    for fi34 := 0; fi34 < 13; fi34++ {
        v33 = float64(3.14159)
        v35 := GenS2{F0: fi34, F1: false, F2: (0.0 >= 0.0)}
        _ = fi34
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
