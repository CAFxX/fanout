package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint; F1 int; F2 string }
type GenS8 struct { F0 int; F1 int16 }
func genFn0() uint16 {
    var v7 uint32 = 42
    fmt.Println(v7)
    v8 := GenS1{F0: ""}
    _ = v7
    _ = v8
    return 228
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = true
    v3 := []uint32{535, 2, 1, 65535}
    v4 := make(map[string]int)
    var v5 int = 32768
    var v6 float64 = (float64(1.5) + 1e10)
    fmt.Println(v6)
    v5 = 42
    v9 := v1["\n\t"]
    if (v2 && (((v6 * (((v6 / v6) / v6) * float64(-2.25))) / 1.5) == v6)) {
        v2 = (v2 || (float64(0.1) <= 3.14159))
    } else {
        var v10 int = v5
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 int = fi11
            v13 := GenS2{F0: ("hello" + ("" + "")), F1: 42}
            var v14 uint64 = 2
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        v15 := GenS1{F0: "foo bar"}
        _ = v10
        _ = v15
    }
    if v2 {
        switch "a" {
        case "\n\t":
            for fi16 := 0; fi16 < 20; fi16++ {
                for fi17 := 0; fi17 < 7; fi17++ {
                    var v18 uint16 = 255
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
        case "hello":
            switch "hello" {
            case "hello":
                if v2 {
                    v2 = v2
                    var v19 uint64 = 2
                    v5 += -206
                    _ = v19
                } else {
                    v2 = false
                    v20 := v1[""]
                    _ = v20
                }
            case "":
                v6 = v6
            default:
                v2 = v2
            }
        }
        _ = len(v4)
        var v21 uint16 = 0
        _ = v21
    }
    var v22 int32 = 0
    fmt.Println(v2, v22, v2)
    delete(v1, "")
    for fi23 := 0; fi23 < 17; fi23++ {
        switch fi23 {
        case 9223372036854775807:
            v6 = (v6 + float64(0.1))
        case 127:
            for fi24 := 0; fi24 < 12; fi24++ {
                v25 := GenS3{F0: 486}
                switch fi24 {
                case 32768:
                    if false {
                        var v26 uint32 = 4294967295
                        v9 += int(v26)
                        v27 := GenS4{F0: 1, F1: 0}
                        _ = v26
                        _ = v27
                    }
                }
                v28 := GenS5{F0: 2}
                _ = fi24
                _ = v25
                _ = v28
            }
        default:
            v29 := GenS4{F0: 1, F1: 588}
            _ = v29
        }
        v9 = (v5 + (^(v5)))
        fmt.Println(v2)
        _ = fi23
    }
    var v30 int8 = 100
    v1["\n\t"] = v9
    v3 = append(v3, 515)
    var v31 int = v9
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (924 == 1) {
        switch "\n\t" {
        case "\n\t":
            for fi32 := 0; fi32 < 1; fi32++ {
                if v2 {
                    if false {
                        if len(v3) > 0 {
                            v3[0] = 44
                        }
                        v33 := GenS5{F0: 0}
                        _ = v33
                    }
                    fmt.Println(v6, v6)
                } else {
                    if v2 {
                        var v34 uint = 4294967295
                        var v35 uint64 = 2
                        _ = v34
                        _ = v35
                    }
                    var v36 uint16 = 2
                    _ = v36
                }
                if (2147483647 <= (v9 * 9223372036854775807)) {
                    for fi37 := 0; fi37 < 9; fi37++ {
                        fmt.Println(v2, fi37)
                        _ = fi37
                    }
                    v5 += (fi32 * v9)
                } else {
                    for fi38 := 0; fi38 < 15; fi38++ {
                        v3 = append(v3, 193)
                        v3 = append(v3, 2)
                        v31 += (v5 * fi38)
                        _ = fi38
                    }
                    v39 := GenS6{F0: fi32}
                    for fi40 := 0; fi40 < 19; fi40++ {
                        v41 := GenS7{F0: 369, F1: (^(fi40)), F2: "hello"}
                        v22 = (v22 + 0)
                        v30 = v30
                        _ = fi40
                        _ = v41
                    }
                    _ = v39
                }
                v2 = v2
                _ = fi32
            }
        default:
            v42 := GenS8{F0: (v5 / ((v9) | 1)), F1: 0}
            _ = v42
        }
        switch v31 {
        case 9223372036854775807:
            v4["a"] = (v5 - v31)
        }
        var v43 uint16 = 255
        _ = v43
    }
    var v44 float64 = ((v6 - -2.25) + -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v9
}
