package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint32; F2 uint }
type GenS3 struct { F0 uint; F1 int64; F2 int16 }
type GenS4 struct { F0 string }
func genFn0(p0 int32) int8 {
    var v7 uint8 = 255
    for fi8 := 0; fi8 < 15; fi8++ {
        switch fi8 {
        case -9223372036854775808:
            v9 := GenS1{F0: (false && true)}
            _ = v9
        case -974:
            var v10 int8 = -128
            _ = v10
        }
        p0 += 1
        _ = fi8
    }
    _ = p0
    _ = v7
    return -128
}
func genFn1(p0 int32, p1 int8, p2 int32) bool {
    switch 24 {
    case -244:
        v11 := GenS1{F0: false}
        _ = v11
    case -278:
        v12 := GenS2{F0: 42, F1: 4294967295, F2: 696}
        _ = v12
    default:
        v13 := []int{434, 127, -1}
        _ = v13
    }
    switch 128 {
    case -662:
        for fi14 := 0; fi14 < 12; fi14++ {
            fmt.Println(p2, p0, fi14)
            v15 := GenS3{F0: 255, F1: int64(fi14), F2: -828}
            _ = fi14
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return true
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]int)
    var v3 int = -994
    var v4 float32 = ((0.0 / 1.5) * (1.5 / (0.0 - 1.5)))
    var v5 string = "\n\t"
    var v6 string = "hello"
    v5 = (v5 + "\n\t")
    for fi16 := 0; fi16 < 12; fi16++ {
        v17 := []int{v3, v3, v3, (fi16 << ((855) & 63))}
        fmt.Println(v6, v5, v6)
        _ = fi16
        _ = v17
    }
    switch (((((v6 + v5) + (v5 + "")) + "") + "hello") + "\n\t") {
    case "\n\t":
        var v18 string = "hello"
        _ = v18
    default:
        if (true && (!true)) {
            v5 = v5
        }
    }
    if false {
        v3 = 31
        v19 := []int{(^(v3)), v3, v3}
        v20 := GenS3{F0: 4294967295, F1: -9223372036854775808, F2: 1000}
        _ = v19
        _ = v20
    } else {
        v21 := []int{}
        var v22 int16 = -735
        if false {
            fmt.Println(v6, v6, v1)
        } else {
            for fi23 := 0; fi23 < 6; fi23++ {
                v2[32767] = (v3 * (fi23 % ((v3) | 1)))
                _ = fi23
            }
            for fi24 := 0; fi24 < 5; fi24++ {
                var v25 int16 = 32767
                switch v6 {
                case "a":
                    for fi26 := 0; fi26 < 5; fi26++ {
                        _ = len(v2)
                        _ = fi26
                    }
                case "foo bar":
                    if len(v21) > 0 {
                        v21[0] = v3
                    }
                case "":
                    for fi27 := 0; fi27 < 14; fi27++ {
                        var v28 float32 = float32(1e10)
                        switch "\n\t" {
                        case "a":
                            v29 := GenS1{F0: (!false)}
                            _ = v29
                        case "\n\t":
                            v30 := GenS1{F0: ((true || (95 == 127)) || true)}
                            fmt.Println(v25, v1, v1)
                            _ = v30
                        case "hello":
                            if ((-2147483648 < (fi27 % (((^(v3))) | 1))) && true) {
                                v2[0] = fi27
                            } else {
                                v31 := GenS1{F0: false}
                                var v32 float32 = v4
                                var v33 int16 = (v22 + 1)
                                if true {
                                    var v34 uint8 = 0
                                    var v35 int8 = 104
                                    delete(v2, fi27)
                                    _ = v34
                                    _ = v35
                                } else {
                                    var v36 uint16 = 42
                                    _ = v36
                                }
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                        default:
                            if len(v21) > 0 {
                                v21[0] = -868
                            }
                        }
                        _ = fi27
                        _ = v28
                    }
                }
                _ = fi24
                _ = v25
            }
        }
        _ = v21
        _ = v22
    }
    if true {
        if false {
            var v37 float64 = float64(((v4 - v4) - float32(3.14159)))
            v6 = ""
            _ = v37
        } else {
            v6 = "foo bar"
            if true {
                for fi38 := 0; fi38 < 9; fi38++ {
                    v39 := GenS1{F0: (139 == 42)}
                    v40 := GenS1{F0: (!true)}
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                if (!(325 < 672)) {
                    var v41 float64 = (float64(0.0) * 1e10)
                    var v42 int = 4294967295
                    switch (v3 % ((v42) | 1)) {
                    case 4294967295:
                        switch ((v6 + (v5 + "a")) + v5) {
                        case "foo bar":
                            var v43 uint64 = 1
                            _ = v43
                        }
                    case -9223372036854775808:
                        var v44 bool = (!true)
                        _ = v44
                    case -2147483648:
                        if ("a" < "\n\t") {
                            fmt.Println(v1)
                            for fi45 := 0; fi45 < 18; fi45++ {
                                var v46 uint64 = 42
                                fmt.Println(v3)
                                var v47 uint8 = 13
                                _ = fi45
                                _ = v46
                                _ = v47
                            }
                        } else {
                            v48 := GenS4{F0: (v6 + v5)}
                            _ = v48
                        }
                    }
                    _ = v41
                    _ = v42
                }
            } else {
                fmt.Println(v1)
                var v49 int16 = 622
                _ = v49
            }
            if (9223372036854775807 != 0) {
                var v50 uint32 = 2147483647
                _ = v50
            }
        }
    } else {
        v1 += v1
        v51 := GenS2{F0: 1000, F1: 42, F2: 255}
        fmt.Println(v4, v3, v4)
        _ = v51
    }
    var v52 uint64 = 18446744073709551615
    for fi53 := 0; fi53 < 7; fi53++ {
        v52 = v52
        v54 := GenS2{F0: -32768, F1: 427, F2: 65535}
        _ = fi53
        _ = v54
    }
    var v55 uint32 = 2147483647
    var v56 int16 = 586
    switch ("\n\t" + v5) {
    case "a":
        v57 := GenS3{F0: 18446744073709551615, F1: 9223372036854775807, F2: (v56 & (v56 | v56))}
        _ = v57
    case "hello":
        v5 = v5
    }
    v58 := GenS4{F0: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v55
    _ = v56
    _ = v58
    _ = v6
}
