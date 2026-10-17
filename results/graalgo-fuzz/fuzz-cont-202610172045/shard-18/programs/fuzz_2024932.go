package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int8; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int32; F2 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 uint; F1 string; F2 uint64 }
func genFn0(p0 int64) int8 {
    v5 := []int{4294967295, 1, 309}
    v6 := GenS1{F0: (^(p0)), F1: false}
    v7 := GenS2{F0: -128, F1: -169, F2: 201}
    switch 965 {
    case 807:
        for fi8 := 0; fi8 < 16; fi8++ {
            fmt.Println(p0, p0)
            _ = fi8
        }
    }
    var v9 uint64 = 4294967295
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v9
    return -128
}

func main() {
    var v1 bool = false
    v2 := []string{("hello" + ("\n\t" + ("foo bar" + "a"))), (("" + "foo bar") + "a"), "a", "hello"}
    v3 := make(map[string]int)
    var v4 int64 = 943
    if (!v1) {
        v4 = v4
        v4 += (v4 - -9223372036854775808)
        var v10 int64 = v4
        _ = v10
    } else {
        if false {
            var v11 int = 2147483647
            _ = v11
        }
        var v12 uint32 = 911
        v13 := v3["a"]
        _ = v12
        _ = v13
    }
    v14 := GenS3{F0: 42, F1: -1, F2: 127}
    delete(v3, "foo bar")
    v15 := GenS3{F0: 0, F1: 772, F2: genFn0((v4 & 1))}
    var v16 uint64 = 4294967295
    v16 -= (v16 >> ((4294967295) & 63))
    var v17 int32 = 0
    v4 = v4
    var v18 int = 1
    var v19 uint = 1
    v18 = v18
    if (!v1) {
        fmt.Println(v18, v4, v18)
    } else {
        fmt.Println(v17)
        var v20 float64 = ((0.1 * (float64((0.0 * (-2.25 * ((3.14159 / (0.1 - 1e10)) - 1.5)))) * 1e10)) * 3.14159)
        if (true && (true || (v18 <= v18))) {
            switch (v18 | 32768) {
            case -1:
                v21 := GenS4{F0: 127}
                _ = v21
            case -2147483648:
                _ = len(v2)
            case -9223372036854775808:
                v1 = false
            }
            switch ("" + "") {
            case "a":
                switch 4294967295 {
                case -796:
                    v4 = (v4 % ((v4) | 1))
                case 128:
                    v1 = true
                case -2147483648:
                    v4 = v4
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            for fi22 := 0; fi22 < 10; fi22++ {
                switch "hello" {
                case "a":
                    v1 = v1
                default:
                    for fi23 := 0; fi23 < 18; fi23++ {
                        fmt.Println(v17, v1)
                        switch "a" {
                        case "foo bar":
                            switch "\n\t" {
                            case "hello":
                                if v1 {
                                    var v24 uint32 = 1
                                    v24 = v24
                                    var v25 int = fi22
                                    _ = v24
                                    _ = v25
                                } else {
                                    var v26 uint32 = 2
                                    _ = v26
                                }
                                v1 = false
                            case "a":
                                var v27 int64 = v4
                                _ = v27
                            case "\n\t":
                                v28 := GenS5{F0: 0, F1: v1}
                                _ = v28
                            default:
                                fmt.Println(v4, v1, v20)
                            }
                        }
                        _ = fi23
                    }
                }
                for fi29 := 0; fi29 < 7; fi29++ {
                    var v30 uint8 = 127
                    fmt.Println(v19, v1)
                    switch fi22 {
                    case -79:
                        fmt.Println(v4, v20)
                    case 127:
                        v31 := GenS6{F0: (v19 << ((v19) & 63)), F1: ("\n\t" + "a"), F2: (v16 & (v16 << (((v19 >> (((v19 >> (((v19 & 980)) & 63))) & 63))) & 63)))}
                        _ = v31
                    case 42:
                        switch ("\n\t" + (("" + "hello") + ("" + (("\n\t" + ("a" + "")) + "a")))) {
                        case "":
                            _ = len(v3)
                        case "\n\t":
                            if (-32768 < 42) {
                                v3["foo bar"] = (fi22 ^ int(v4))
                            } else {
                                var v32 uint8 = (v30 | v30)
                                _ = v32
                            }
                        default:
                            var v33 int8 = 1
                            _ = v33
                        }
                    }
                    _ = fi29
                    _ = v30
                }
                v19 = v19
                _ = fi22
            }
            var v34 int16 = 1000
            _ = v34
        }
        _ = v20
    }
    v2 = append(v2, "a")
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
