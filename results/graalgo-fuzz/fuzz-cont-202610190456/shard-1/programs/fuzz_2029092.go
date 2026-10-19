package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 uint16; F2 uint }
type GenS4 struct { F0 int64; F1 uint32 }

func main() {
    var v1 int32 = 860
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    var v4 int = 32767
    var v5 int = (v4 % ((32768) | 1))
    if true {
        for fi6 := 0; fi6 < 18; fi6++ {
            v1 = (v1 + v1)
            var v7 uint8 = 1
            if true {
                v5 -= (fi6 << ((2) & 63))
                v1 = (v1 & 0)
            }
            _ = fi6
            _ = v7
        }
        fmt.Println(v1, v4, v5)
    } else {
        var v8 uint16 = 571
        v5 -= (v5 << ((42) & 63))
        _ = v8
    }
    v4 -= v5
    v5 += v4
    switch (v4 >> ((55) & 63)) {
    case 32768:
        fmt.Println(v5, v1)
    case -2147483648:
        if true {
            v1 -= (^(v1))
            var v9 int64 = 0
            var v10 bool = (((("foo bar" + "\n\t") + "a") != (("foo bar" + "\n\t") + "a")) && true)
            _ = v10
            _ = v9
        } else {
            v5 = v5
            for fi11 := 0; fi11 < 3; fi11++ {
                v12 := GenS1{F0: 2}
                switch (fi11 << ((65535) & 63)) {
                case -1:
                    var v13 uint64 = 334
                    _ = v13
                case 763:
                    v4 = (^(v5))
                }
                _ = fi11
                _ = v12
            }
            switch (("foo bar" + "hello") + "\n\t") {
            case "":
                v14 := []int{1, -352, (v5 >> ((552) & 63)), 9223372036854775807}
                _ = v14
            default:
                switch "a" {
                case "\n\t":
                    v2[v5] = ""
                case "hello":
                    fmt.Println(v4, v4)
                case "foo bar":
                    for fi15 := 0; fi15 < 17; fi15++ {
                        v4 -= v5
                        fmt.Println(v5)
                        for fi16 := 0; fi16 < 9; fi16++ {
                            v1 = v1
                            if (true && true) {
                                if true {
                                    var v17 uint64 = 1
                                    v4 += v5
                                    v18 := GenS2{F0: 1}
                                    _ = v17
                                    _ = v18
                                }
                                if true {
                                    var v19 int8 = 0
                                    switch (fi15 + (^(v5))) {
                                    case 1:
                                        fmt.Println(v19, v19)
                                    case 42:
                                        v20 := []int{(fi15 ^ fi15)}
                                        _ = v20
                                    default:
                                        var v21 uint32 = 2
                                        _ = v21
                                    }
                                    switch ("foo bar" + "hello") {
                                    case "a":
                                        v5 -= fi16
                                    default:
                                        v22 := []int{int(v19), (v5 / (((fi15 / ((128) | 1))) | 1))}
                                        if len(v22) > 0 {
                                            v22[0] = v4
                                        }
                                        _ = v22
                                    }
                                    v4 = fi15
                                    var v23 string = ""
                                    _ = v19
                                    _ = v23
                                } else {
                                    v24 := GenS2{F0: -9223372036854775808}
                                    if true {
                                        var v25 float64 = 3.14159
                                        v26 := GenS3{F0: (fi16 % (((v4 >> ((0) & 63))) | 1)), F1: 2, F2: 18446744073709551615}
                                        _ = v25
                                        _ = v26
                                    } else {
                                        v5 += int(v1)
                                    }
                                    fmt.Println(fi16, v5, v5)
                                    _ = v24
                                }
                                fmt.Println(v1)
                                var v27 uint32 = 354
                                if ((!false) || false) {
                                    var v28 int16 = 1
                                    _ = v28
                                }
                                _ = v27
                            }
                            v2[fi16] = ("a" + "hello")
                            fmt.Println(fi15, v5, v4)
                            _ = fi16
                        }
                        _ = fi15
                    }
                }
            }
        }
    case 9223372036854775807:
        v29 := GenS3{F0: v5, F1: 0, F2: 2}
        _ = v29
    }
    v5 += v4
    switch "" {
    case "\n\t":
        v5 -= v4
    case "":
        for fi30 := 0; fi30 < 9; fi30++ {
            v1 = v1
            v31 := []int{v5, v5}
            _ = fi30
            _ = v31
        }
    case "foo bar":
        if (false || false) {
            fmt.Println(v1, v5, v1)
            delete(v3, ((("foo bar" + "") + (("foo bar" + "foo bar") + ("" + "hello"))) + (("\n\t" + ("" + "a")) + "")))
        }
    }
    v32 := v3["\n\t"]
    v4 = (-(v4))
    var v33 uint8 = 2
    v34 := v3["a"]
    v35 := GenS4{F0: -555, F1: 1}
    if (0 > (v33 % (((v33 ^ (v33 * v33))) | 1))) {
        v1 += 789
        switch v5 {
        case 128:
            if (true && false) {
                fmt.Println(v4)
            }
        case 1:
            v34 += 0
        case 813:
            v3["\n\t"] = v34
        }
        v34 -= (-(v34))
    }
    var v36 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
