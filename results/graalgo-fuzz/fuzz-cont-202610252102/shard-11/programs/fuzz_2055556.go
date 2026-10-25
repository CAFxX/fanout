package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }

func main() {
    var v1 float64 = 0.1
    v2 := make(map[string]bool)
    var v3 uint = 42
    var v4 string = ("hello" + ("\n\t" + ""))
    v2[v4] = (5 >= 469)
    switch ("a" + "hello") {
    case "foo bar":
        for fi5 := 0; fi5 < 9; fi5++ {
            fmt.Println(v3, v1)
            _ = fi5
        }
    case "hello":
        for fi6 := 0; fi6 < 9; fi6++ {
            v1 = float64(0.0)
            var v7 uint16 = 2
            fmt.Println(v3, v3, fi6)
            _ = fi6
            _ = v7
        }
    case "\n\t":
        var v8 uint64 = 586
        _ = v8
    }
    var v9 string = v4
    for fi10 := 0; fi10 < 17; fi10++ {
        switch "" {
        case "a":
            var v11 uint64 = 42
            _ = v11
        case "foo bar":
            if false {
                var v12 int = fi10
                _ = v12
            } else {
                v1 = 0.0
            }
        case "hello":
            var v13 bool = (false || (-104 < 1))
            _ = v13
        default:
            v14 := []int{(-(fi10)), -2147483648, fi10}
            _ = v14
        }
        v1 -= v1
        var v15 string = v4
        _ = fi10
        _ = v15
    }
    if (1 < 1) {
        v16 := []int{-2147483648, 32768}
        switch 42 {
        case -9223372036854775808:
            switch "" {
            case "foo bar":
                if (true && (!true)) {
                    v1 = v1
                    var v17 int8 = -128
                    switch v4 {
                    case "hello":
                        if false {
                            if ((0.1 <= (0.1 + (1.5 * float32(0.1)))) || (-2147483648 >= 42)) {
                                for fi18 := 0; fi18 < 11; fi18++ {
                                    v17 += v17
                                    fmt.Println(v3)
                                    _ = fi18
                                }
                                v19 := GenS1{F0: 4294967295, F1: "foo bar"}
                                _ = v19
                            }
                            fmt.Println(v9, v1, v9)
                        } else {
                            if true {
                                _ = len(v2)
                                v20 := GenS1{F0: 4294967295, F1: (("" + (v4 + v9)) + (("" + v9) + v4))}
                                v17 = (v17 + v17)
                                _ = v20
                            } else {
                                v1 = ((v1 / (v1 / (((v1 + float64(0.1)) + v1) - v1))) / v1)
                                v21 := GenS1{F0: 2, F1: (v4 + "\n\t")}
                                _ = v21
                            }
                        }
                    case "a":
                        delete(v2, v4)
                    case "foo bar":
                        var v22 uint = 64
                        _ = v22
                    }
                    _ = v17
                } else {
                    var v23 uint16 = uint16(v3)
                    for fi24 := 0; fi24 < 2; fi24++ {
                        for fi25 := 0; fi25 < 12; fi25++ {
                            v2["hello"] = true
                            _ = fi25
                        }
                        _ = fi24
                    }
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    _ = v23
                }
            case "a":
                if true {
                    v26 := GenS2{F0: 0}
                    v4 = v9
                    fmt.Println(v9)
                    _ = v26
                } else {
                    v27 := GenS1{F0: 0, F1: v4}
                    _ = v27
                }
            default:
                v1 -= v1
            }
        case 484:
            if false {
                v28 := GenS3{F0: (v4 + (v9 + "a"))}
                _ = v28
            }
        case 2147483647:
            var v29 int32 = 588
            _ = v29
        }
        _ = v16
    } else {
        _ = len(v2)
        v30 := v2[v9]
        _ = v30
    }
    v9 = "hello"
    var v31 int16 = 338
    var v32 float32 = 0.0
    v4 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v9
}
