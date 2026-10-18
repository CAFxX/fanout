package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 string; F1 uint64; F2 bool }
type GenS3 struct { F0 int32 }
func genFn0(p0 string) bool {
    p0 = "a"
    p0 = (p0 + p0)
    for fi4 := 0; fi4 < 11; fi4++ {
        p0 = (p0 + p0)
        _ = fi4
    }
    v5 := GenS1{F0: 53, F1: "foo bar"}
    _ = p0
    _ = v5
    return true
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = (0.1 / 1.0)
    v3 := make(map[string]uint32)
    fmt.Println(v2)
    v1["a"] = 42
    fmt.Println(v2)
    for fi6 := 0; fi6 < 15; fi6++ {
        v3["a"] = 854
        switch "\n\t" {
        case "":
            v2 = v2
        case "hello":
            var v7 int = 128
            _ = v7
        case "\n\t":
            var v8 float32 = 0.1
            _ = v8
        }
        switch (^(fi6)) {
        case 32767:
            v2 = ((v2 - v2) - (1e10 + 1e10))
        default:
            var v9 int16 = 32767
            _ = v9
        }
        _ = fi6
    }
    v10 := GenS1{F0: 258, F1: "hello"}
    for fi11 := 0; fi11 < 9; fi11++ {
        v2 -= 0.1
        if false {
            v12 := []int{4294967295, -9223372036854775808, -734, (-(fi11))}
            var v13 uint16 = 551
            _ = len(v3)
            _ = v12
            _ = v13
        }
        if true {
            v14 := GenS1{F0: 0, F1: ("hello" + ("" + ""))}
            _ = v14
        }
        _ = fi11
    }
    if (0 <= 279) {
        for fi15 := 0; fi15 < 16; fi15++ {
            v16 := []int{}
            switch "\n\t" {
            case "a":
                if true {
                    switch fi15 {
                    case 32767:
                        v2 -= 3.14159
                    case -785:
                        var v17 uint64 = 0
                        _ = v17
                    }
                } else {
                    switch -1 {
                    case 640:
                        switch (fi15 % (((^(fi15))) | 1)) {
                        case 128:
                            v18 := GenS1{F0: 0, F1: "\n\t"}
                            _ = v18
                        case 4294967295:
                            switch "" {
                            case "":
                                v2 += 1.5
                            case "\n\t":
                                if false {
                                    fmt.Println(fi15, fi15)
                                    var v19 int8 = -128
                                    _ = v19
                                }
                            case "hello":
                                fmt.Println(v2)
                            }
                        case 32767:
                            switch ("a" + "\n\t") {
                            case "a":
                                var v20 int8 = -128
                                _ = v20
                            case "\n\t":
                                var v21 string = ""
                                _ = v21
                            }
                        default:
                            v2 = float64(1.5)
                        }
                    case -129:
                        switch "a" {
                        case "a":
                            v22 := GenS2{F0: "foo bar", F1: 689, F2: false}
                            _ = v22
                        case "hello":
                            v23 := v1["a"]
                            _ = v23
                        default:
                            if (2147483647 == 2147483647) {
                                var v24 uint64 = 810
                                v24 = 563
                                _ = v24
                            } else {
                                v2 = v2
                                if ((1.5 < 0.0) || (!(-60 < 1))) {
                                    var v25 uint8 = 42
                                    for fi26 := 0; fi26 < 11; fi26++ {
                                        var v27 int = (fi26 & (fi15 & (fi26 * int(v25))))
                                        _ = fi26
                                        _ = v27
                                    }
                                    if (!(true || (!((fi15 % ((fi15) | 1)) >= (fi15 % ((9223372036854775807) | 1)))))) {
                                        _ = len(v16)
                                        var v28 uint8 = 0
                                        _ = v28
                                    }
                                    _ = v25
                                }
                            }
                        }
                    default:
                        _ = len(v16)
                    }
                    v2 = -2.25
                }
            default:
                switch -2147483648 {
                case -2147483648:
                    v29 := GenS2{F0: "\n\t", F1: 2, F2: true}
                    _ = v29
                case 1:
                    var v30 float64 = (1.5 + v2)
                    _ = v30
                default:
                    v31 := GenS2{F0: "hello", F1: 42, F2: true}
                    _ = v31
                }
            }
            var v32 float32 = -2.25
            _ = fi15
            _ = v16
            _ = v32
        }
        fmt.Println(v2)
    } else {
        v2 = (float64(((0.0 * 1e10) - 1.5)) + v2)
        v2 -= (float64((0.1 + 1.5)) / -2.25)
    }
    v33 := GenS3{F0: -584}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
}
