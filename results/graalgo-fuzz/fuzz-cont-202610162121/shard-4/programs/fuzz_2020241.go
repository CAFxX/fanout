package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS3 struct { F0 bool; F1 string }

func main() {
    v1 := []int{1}
    v2 := []string{"hello"}
    var v3 string = "a"
    v4 := make(map[int]string)
    v5 := GenS1{F0: -634}
    if false {
        v6 := GenS1{F0: -9223372036854775808}
        _ = len(v4)
        v3 = (v3 + (v3 + v3))
        _ = v6
    }
    var v7 bool = (95 >= 100)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if v7 {
        switch 4294967295 {
        case 128:
            v3 = v3
        case 127:
            if false {
                if v7 {
                    v7 = (0 < 723)
                }
                var v8 int64 = 42
                _ = v8
            } else {
                v7 = false
            }
        default:
            v9 := GenS2{F0: 127, F1: 1, F2: 819}
            _ = v9
        }
        v7 = (!(-128 <= -83))
    }
    v3 = v3
    switch -9223372036854775808 {
    case 0:
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 bool = (2 != 1)
            _ = fi10
            _ = v11
        }
    case 9223372036854775807:
        if v7 {
            for fi12 := 0; fi12 < 8; fi12++ {
                for fi13 := 0; fi13 < 12; fi13++ {
                    fmt.Println(v3)
                    if v7 {
                        var v14 float64 = 0.1
                        v14 -= v14
                        switch "foo bar" {
                        case "a":
                            var v15 bool = false
                            v3 = v3
                            _ = v15
                        default:
                            v3 = ("foo bar" + (((v3 + "") + "a") + ""))
                        }
                        v16 := v4[fi13]
                        _ = len(v4)
                        _ = v14
                        _ = v16
                    } else {
                        v7 = v7
                        v3 = ("" + v3)
                        v17 := GenS1{F0: fi12}
                        _ = v17
                    }
                    _ = fi13
                }
                if v7 {
                    fmt.Println(fi12, fi12, v3)
                    v7 = (!true)
                    v18 := GenS3{F0: ((1 != -2147483648) && v7), F1: v3}
                    _ = v18
                } else {
                    if true {
                        v7 = (0 != 1)
                        var v19 float64 = -2.25
                        if (0 == -1) {
                            fmt.Println(v3, fi12)
                            switch ((v3 + "hello") + v3) {
                            case "hello":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case "":
                                var v20 float32 = float32(v19)
                                _ = v20
                            case "\n\t":
                                switch v3 {
                                case "\n\t":
                                    v1 = append(v1, (fi12 | 202))
                                }
                            }
                            var v21 int8 = -83
                            _ = v21
                        }
                        _ = v19
                    }
                    v22 := GenS2{F0: 1, F1: 2147483648, F2: 42}
                    _ = v22
                }
                _ = fi12
            }
            v3 = v3
            var v23 float64 = 1.5
            _ = v23
        }
    case 404:
        v3 = ""
    default:
        for fi24 := 0; fi24 < 10; fi24++ {
            var v25 uint64 = 453
            _ = fi24
            _ = v25
        }
    }
    v3 = (v3 + v3)
    var v26 float64 = 1e10
    v26 = v26
    var v27 bool = ((v3 + v3) < (v3 + "hello"))
    for fi28 := 0; fi28 < 1; fi28++ {
        fmt.Println(fi28, v26, v3)
        _ = fi28
    }
    fmt.Println(v3, v7)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
