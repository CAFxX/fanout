package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int8; F2 int32 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 uint8 = 2
    var v2 uint64 = 2
    v3 := []int{127, 32768, -9223372036854775808}
    var v4 float32 = 0.0
    v5 := make(map[int]string)
    v6 := make(map[string]string)
    if (uint64(v1) >= (v2 % (((^(v2))) | 1))) {
        var v7 int8 = 1
        _ = v7
    }
    if (false && (((false && (!(!false))) && (!(false && (2 > 2)))) || true)) {
        v4 = v4
        fmt.Println(v2, v2, v4)
    } else {
        delete(v5, 194)
    }
    v8 := GenS1{F0: "hello", F1: true}
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 int = fi10
            switch (v11 << ((2) & 63)) {
            case -2147483648:
                fmt.Println(fi10)
            case 303:
                v11 = (fi9 | v11)
            }
            _ = fi10
            _ = v11
        }
        v2 += v2
        var v12 uint32 = 42
        _ = fi9
        _ = v12
    }
    var v13 float64 = -2.25
    v14 := GenS1{F0: "\n\t", F1: true}
    v13 = (v13 / (0.1 * (3.14159 + v13)))
    switch (("hello" + "a") + ("a" + "a")) {
    case "hello":
        for fi15 := 0; fi15 < 11; fi15++ {
            v16 := v6["foo bar"]
            v1 += v1
            _ = fi15
            _ = v16
        }
    }
    switch 2147483647 {
    case -1:
        v3 = append(v3, 1)
    case 32767:
        if true {
            v17 := v6["hello"]
            if true {
                v18 := GenS1{F0: ("a" + v17), F1: false}
                _ = v18
            }
            var v19 uint8 = v1
            _ = v17
            _ = v19
        }
    }
    v4 = v4
    var v20 uint = 969
    var v21 int = 128
    switch "hello" {
    case "foo bar":
        if true {
            fmt.Println(v4, v21)
            var v22 string = ("\n\t" + "hello")
            _ = v22
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    case "\n\t":
        if (67 != (v2 >> ((v20) & 63))) {
            for fi23 := 0; fi23 < 16; fi23++ {
                switch 831 {
                case -1:
                    switch (fi23 - fi23) {
                    case -1:
                        v13 += float64(0.0)
                    case -2147483648:
                        var v24 int64 = -9223372036854775808
                        _ = v24
                    case -116:
                        v2 -= v2
                    }
                case -9223372036854775808:
                    v21 = fi23
                default:
                    if true {
                        v1 += v1
                        for fi25 := 0; fi25 < 18; fi25++ {
                            v6["hello"] = "foo bar"
                            v1 += (v1 >> ((v20) & 63))
                            if (v20 > (v20 >> ((2) & 63))) {
                                var v26 string = "foo bar"
                                if true {
                                    _ = len(v6)
                                    _ = len(v3)
                                    _ = len(v5)
                                } else {
                                    if true {
                                        if len(v3) > 0 {
                                            v3[0] = 127
                                        }
                                        v6[v26] = (v26 + v26)
                                        if false {
                                            v27 := v5[fi23]
                                            v27 = ((v26 + (("a" + v26) + (v26 + v26))) + v26)
                                            var v28 uint64 = (v2 & (v2 ^ (v2 << ((v20) & 63))))
                                            _ = v27
                                            _ = v28
                                        } else {
                                            fmt.Println(v21)
                                            if false {
                                                v29 := GenS2{F0: ("foo bar" + "hello")}
                                                _ = v29
                                            } else {
                                                var v30 uint16 = 42
                                                v31 := GenS1{F0: ("\n\t" + (v26 + ((("a" + (("a" + v26) + v26)) + "") + ""))), F1: (-1 > 623)}
                                                var v32 uint = 927
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                            }
                                        }
                                    } else {
                                        var v33 uint16 = 2
                                        if false {
                                            fmt.Println(v4, v26)
                                        }
                                        if false {
                                            var v34 uint32 = 1
                                            fmt.Println(v2, v33)
                                            _ = v34
                                        }
                                        _ = v33
                                    }
                                    v35 := GenS1{F0: v26, F1: (v26 <= ("" + "hello"))}
                                    fmt.Println(v2, v21)
                                    v5[(fi25 * (v21 << ((v20) & 63)))] = ((v26 + ((v26 + ((v26 + (v26 + "a")) + v26)) + "foo bar")) + v26)
                                    _ = v35
                                }
                                v36 := GenS1{F0: ((v26 + v26) + v26), F1: ((fi23 % ((v21) | 1)) > -2147483648)}
                                _ = v26
                                _ = v36
                            }
                            _ = fi25
                        }
                    }
                }
                v5[-2147483648] = ("\n\t" + "\n\t")
                _ = fi23
            }
            switch (v21 & v21) {
            case 0:
                _ = len(v6)
            }
        } else {
            for fi37 := 0; fi37 < 18; fi37++ {
                fmt.Println(v1, v4)
                _ = fi37
            }
        }
    }
    fmt.Println(v4, v20)
    switch (v21 & v21) {
    case 128:
        for fi38 := 0; fi38 < 5; fi38++ {
            v2 = (v2 * 0)
            if (9223372036854775807 != -80) {
                v20 -= v20
            } else {
                v39 := GenS3{F0: true, F1: -128, F2: -2147483648}
                _ = v39
            }
            _ = fi38
        }
    default:
        delete(v5, v21)
    }
    if false {
        v21 = -184
        v3 = append(v3, v21)
        switch (v21 >> (((^(v20))) & 63)) {
        case 770:
            v40 := GenS4{F0: -92}
            _ = v40
        default:
            v4 = ((-2.25 / v4) + v4)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
