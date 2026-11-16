package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 0.1
    var v3 uint32 = 2147483647
    v4 := make(map[int]bool)
    var v5 uint32 = v3
    v6 := GenS1{F0: 32768}
    switch "\n\t" {
    case "\n\t":
        for fi7 := 0; fi7 < 11; fi7++ {
            if true {
                if (2 > 42) {
                    switch 645 {
                    case -2147483648:
                        v5 = v5
                    case -1:
                        delete(v1, 9223372036854775807)
                    case 9223372036854775807:
                        var v8 int8 = 43
                        _ = v8
                    }
                    for fi9 := 0; fi9 < 1; fi9++ {
                        var v10 float32 = 0.0
                        _ = fi9
                        _ = v10
                    }
                    for fi11 := 0; fi11 < 4; fi11++ {
                        if false {
                            var v12 string = "foo bar"
                            v12 = (v12 + (v12 + v12))
                            if false {
                                if true {
                                    v13 := GenS1{F0: -9223372036854775808}
                                    var v14 bool = (true && false)
                                    v15 := []int{-867}
                                    _ = v13
                                    _ = v14
                                    _ = v15
                                } else {
                                    fmt.Println(v12, v12)
                                }
                                v12 = ""
                                switch ("" + ((v12 + (((v12 + "a") + ("a" + "")) + v12)) + (v12 + (v12 + "foo bar")))) {
                                case "hello":
                                    if true {
                                        if false {
                                            v16 := GenS1{F0: fi7}
                                            _ = v16
                                        }
                                        var v17 int64 = 248
                                        _ = v17
                                    }
                                case "":
                                    v2 = v2
                                case "foo bar":
                                    var v18 uint32 = v3
                                    var v19 uint64 = 34
                                    v18 = v18
                                    _ = v18
                                    _ = v19
                                }
                                v2 = 0.1
                            } else {
                                v20 := []int{fi11, (^(fi11)), -1, 9223372036854775807}
                                v21 := GenS1{F0: fi7}
                                _ = len(v4)
                                _ = v20
                                _ = v21
                            }
                            _ = v12
                        }
                        _ = fi11
                    }
                } else {
                    switch (fi7 ^ (fi7 ^ fi7)) {
                    case 2147483647:
                        v1[fi7] = ""
                    case -448:
                        v22 := GenS1{F0: (-(fi7))}
                        _ = v22
                    default:
                        v2 -= 0.0
                    }
                }
            } else {
                var v23 bool = (false || true)
                _ = v23
            }
            fmt.Println(v3)
            switch "\n\t" {
            case "a":
                var v24 uint16 = 32767
                _ = v24
            }
            _ = fi7
        }
    case "hello":
        var v25 float64 = (v2 / (1.5 - (v2 + v2)))
        _ = v25
    default:
        var v26 uint64 = 4294967295
        _ = v26
    }
    fmt.Println(v5, v2)
    fmt.Println(v2, v5)
    v5 = (v5 >> ((668) & 63))
    switch 391 {
    case 42:
        for fi27 := 0; fi27 < 8; fi27++ {
            switch (fi27 + -9223372036854775808) {
            case 42:
                v3 = v5
            }
            v28 := GenS2{F0: "\n\t", F1: 395}
            _ = fi27
            _ = v28
        }
    }
    fmt.Println(v3, v5, v5)
    v2 = ((((v2 - v2) / v2) * 0.1) + ((float64(1e10) / (((v2 * v2) - (1e10 * (((v2 * 1e10) / v2) - ((float64(0.0) - 0.1) * 1.5)))) / ((v2 * v2) / v2))) * v2))
    if false {
        var v29 float32 = (0.1 + 3.14159)
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
