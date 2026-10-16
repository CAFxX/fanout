package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 string; F1 uint; F2 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32; F1 int64; F2 int32 }
type GenS7 struct { F0 uint16; F1 uint64 }
type GenS8 struct { F0 uint32 }
func genFn0(p0 int8, p1 int, p2 uint8) string {
    fmt.Println(p0, p0)
    if true {
        var v4 uint32 = 2147483647
        var v5 bool = false
        fmt.Println(p0, v5)
        _ = v4
        _ = v5
    } else {
        var v6 bool = true
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return "a"
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 0
    var v3 uint32 = 4294967295
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 650, F1: -1, F2: 0}
        var v9 int32 = 2147483647
        if ((v2 + 4294967295) != (v2 / ((2) | 1))) {
            v9 -= (v9 + v9)
            for fi10 := 0; fi10 < 15; fi10++ {
                switch (fi10 & fi10) {
                case 2147483647:
                    v11 := GenS2{F0: (v3 & v3)}
                    _ = v11
                case -906:
                    v12 := GenS3{F0: -29, F1: (v9 + (v9 | (v9 << ((42) & 63))))}
                    _ = v12
                }
                var v13 string = "\n\t"
                _ = fi10
                _ = v13
            }
            v14 := GenS1{F0: uint16(v2), F1: (v9 + v9), F2: 4294967295}
            _ = v14
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v15 float64 = (1.5 / 0.1)
    v16 := GenS4{F0: ("\n\t" + "a"), F1: 612, F2: 0}
    if (!false) {
        v1[421] = 128
        for fi17 := 0; fi17 < 9; fi17++ {
            switch fi17 {
            case -1:
                switch fi17 {
                case 1:
                    v18 := GenS1{F0: 65535, F1: -2147483648, F2: 18446744073709551615}
                    _ = v18
                default:
                    var v19 uint64 = (v2 ^ (v2 | v2))
                    _ = v19
                }
            case 42:
                v20 := GenS1{F0: 65535, F1: 42, F2: uint(v2)}
                _ = v20
            case -9223372036854775808:
                v2 = (^(v2))
            default:
                v3 = (^(v3))
            }
            _ = fi17
        }
    } else {
        var v21 int64 = 42
        if false {
            _ = len(v1)
            v22 := GenS5{F0: v3}
            _ = v22
        } else {
            v23 := GenS6{F0: (v3 | v3), F1: -246, F2: 42}
            for fi24 := 0; fi24 < 17; fi24++ {
                v25 := []int{}
                switch fi24 {
                case 9223372036854775807:
                    var v26 uint16 = 28
                    _ = v26
                case 42:
                    var v27 uint32 = (v3 / (((v3 << ((264) & 63))) | 1))
                    _ = v27
                case 0:
                    if len(v25) >= 2 {
                        v25 = v25[1:len(v25)]
                    }
                }
                _ = fi24
                _ = v25
            }
            _ = v23
        }
        fmt.Println(v15)
        _ = v21
    }
    v28 := GenS7{F0: 1, F1: v2}
    for fi29 := 0; fi29 < 14; fi29++ {
        if false {
            var v30 string = ""
            _ = v30
        }
        _ = fi29
    }
    _ = len(v1)
    v31 := GenS1{F0: 0, F1: -1, F2: 255}
    var v32 string = "a"
    var v33 uint32 = 989
    for fi34 := 0; fi34 < 3; fi34++ {
        if ((true && (!(false || (-92 > 42)))) || ((-1 == -1) && true)) {
            fmt.Println(v15, v15)
            _ = len(v1)
            if false {
                var v35 string = v32
                _ = v35
            } else {
                var v36 string = (v32 + v32)
                v37 := GenS8{F0: v33}
                if (-1 >= 211) {
                    fmt.Println(v15, v2, v32)
                } else {
                    v33 = 2147483647
                    var v38 int8 = 100
                    for fi39 := 0; fi39 < 15; fi39++ {
                        switch fi39 {
                        case 32768:
                            var v40 uint16 = 892
                            _ = v40
                        }
                        _ = len(v1)
                        _ = fi39
                    }
                    _ = v38
                }
                _ = v36
                _ = v37
            }
        } else {
            v15 -= v15
            var v41 int = fi34
            if false {
                v3 = v33
                var v42 bool = false
                var v43 float32 = 1.5
                _ = v42
                _ = v43
            } else {
                for fi44 := 0; fi44 < 14; fi44++ {
                    v45 := []int{2147483647}
                    v2 = (v2 | (v2 % (((v2 % (((v2 / (((^(v2))) | 1))) | 1))) | 1)))
                    _ = fi44
                    _ = v45
                }
                v2 = (^(v2))
            }
            _ = v41
        }
        v15 = (v15 * v15)
        _ = fi34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}
