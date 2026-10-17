package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 uint8; F2 string }
func genFn0(p0 int8) uint8 {
    var v4 int32 = 931
    var v5 uint8 = 79
    var v6 uint64 = 0
    fmt.Println(v6)
    v7 := GenS1{F0: true}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return v5
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{("foo bar" + "hello"), "\n\t", "a", ("\n\t" + "")}
    var v3 float64 = 3.14159
    v3 = (v3 - -2.25)
    var v8 float64 = 3.14159
    switch "hello" {
    case "hello":
        var v9 uint = 255
        _ = v9
    default:
        v3 += v8
    }
    if false {
        for fi10 := 0; fi10 < 7; fi10++ {
            fmt.Println(fi10)
            var v11 uint64 = 42
            _ = fi10
            _ = v11
        }
    } else {
        fmt.Println(v3)
        v8 = (float64(1.5) + 0.1)
    }
    switch 127 {
    case -9223372036854775808:
        var v12 int = 2147483647
        _ = v12
    case 42:
        for fi13 := 0; fi13 < 11; fi13++ {
            v3 -= v3
            v14 := GenS2{F0: -2147483648, F1: -604}
            _ = fi13
            _ = v14
        }
    case 9223372036854775807:
        for fi15 := 0; fi15 < 20; fi15++ {
            v8 -= v3
            _ = fi15
        }
    default:
        if true {
            switch 42 {
            case 42:
                v8 -= (3.14159 / (((float64(0.1) * v8) - v3) - v8))
            case 427:
                v16 := GenS1{F0: (4294967295 == -385)}
                _ = v16
            case 4294967295:
                switch 9223372036854775807 {
                case 0:
                    if (!(1 < -1)) {
                        for fi17 := 0; fi17 < 11; fi17++ {
                            if (true && (1.5 != 1.5)) {
                                var v18 int8 = 127
                                var v19 uint16 = 335
                                var v20 int16 = 42
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                            v21 := GenS1{F0: false}
                            var v22 int = fi17
                            _ = fi17
                            _ = v21
                            _ = v22
                        }
                    }
                case -621:
                    v23 := GenS1{F0: true}
                    _ = v23
                default:
                    v3 = float64(((1.5 * (((0.1 / 1.0) * 1.5) + (0.0 - ((3.14159 - 0.0) / -2.25)))) / (3.14159 - (((((float32((0.1 * (v3 + (float64(1e10) + v3)))) - 0.0) + 3.14159) + 3.14159) / (1e10 + (1e10 / -2.25))) - 0.0))))
                }
            default:
                v2 = append(v2, ((("a" + "\n\t") + ("a" + "")) + "hello"))
            }
        }
    }
    var v24 uint32 = 903
    var v25 int32 = 1
    var v26 string = "hello"
    if (205 > 2) {
        v3 -= (v8 - ((0.1 + v3) - v8))
        v27 := GenS2{F0: -2147483648, F1: -7}
        _ = v27
    } else {
        var v28 uint = 42
        if true {
            v29 := GenS1{F0: ((v28 ^ v28) >= (v28 * v28))}
            _ = v29
        } else {
            if false {
                _ = len(v1)
                for fi30 := 0; fi30 < 15; fi30++ {
                    v31 := GenS2{F0: -94, F1: -32768}
                    _ = fi30
                    _ = v31
                }
            }
        }
        for fi32 := 0; fi32 < 4; fi32++ {
            fmt.Println(v3, v3)
            v25 = (^(v25))
            _ = fi32
        }
        _ = v28
    }
    if (!(656 == 42)) {
        var v33 float64 = -2.25
        _ = v33
    } else {
        if true {
            var v34 uint16 = 65535
            if (0 > (v34 % ((v34) | 1))) {
                v34 += (^(v34))
            }
            _ = v34
        } else {
            if false {
                for fi35 := 0; fi35 < 6; fi35++ {
                    _ = len(v2)
                    _ = fi35
                }
            } else {
                v26 = (v26 + "\n\t")
                for fi36 := 0; fi36 < 4; fi36++ {
                    for fi37 := 0; fi37 < 18; fi37++ {
                        v3 = v3
                        var v38 int = fi37
                        _ = fi37
                        _ = v38
                    }
                    _ = fi36
                }
                for fi39 := 0; fi39 < 17; fi39++ {
                    fmt.Println(v3)
                    _ = fi39
                }
            }
            switch 9223372036854775807 {
            case -743:
                v40 := GenS3{F0: (v25 % (((v25 << ((497) & 63))) | 1))}
                _ = v40
            case -322:
                v41 := GenS4{F0: 1, F1: 23, F2: v26}
                _ = v41
            }
            if (0.0 < (-2.25 - ((1e10 / -2.25) - 3.14159))) {
                switch v26 {
                case "\n\t":
                    v42 := v1[42]
                    _ = v42
                case "":
                    v43 := GenS4{F0: 104, F1: genFn0(-128), F2: "hello"}
                    _ = v43
                }
                v3 = (v3 * v8)
            } else {
                for fi44 := 0; fi44 < 4; fi44++ {
                    v26 = v26
                    fmt.Println(v26, v3, fi44)
                    _ = fi44
                }
            }
        }
    }
    if true {
        if (-9223372036854775808 < -9223372036854775808) {
            for fi45 := 0; fi45 < 11; fi45++ {
                switch 157 {
                case -889:
                    fmt.Println(fi45, v26, v3)
                default:
                    v25 = (v25 - 1)
                }
                var v46 float64 = 1e10
                _ = fi45
                _ = v46
            }
            if false {
                v25 = v25
            }
        }
    }
    switch ("\n\t" + v26) {
    case "a":
        fmt.Println(v26, v26, v25)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v8
}
