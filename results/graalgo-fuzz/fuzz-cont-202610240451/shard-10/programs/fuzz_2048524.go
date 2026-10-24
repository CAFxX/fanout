package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int16; F1 string; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 uint8; F1 int; F2 int64 }
func genFn0() int8 {
    if (!(996 >= -189)) {
        var v7 int16 = 1
        v7 = v7
        v7 -= v7
        _ = v7
    } else {
        switch 0 {
        case 4294967295:
            v8 := []int{-1, -24}
            _ = v8
        case 2147483647:
            var v9 uint8 = 255
            _ = v9
        default:
            var v10 bool = false
            _ = v10
        }
        var v11 int16 = -32768
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        if ((fi12 * fi12) == fi12) {
            var v13 uint16 = 346
            v14 := []int{fi12}
            if true {
                var v15 uint64 = 2
                switch fi12 {
                case 127:
                    fmt.Println(v13)
                case 128:
                    if (!false) {
                        if false {
                            fmt.Println(v13)
                        }
                        var v16 uint8 = 42
                        _ = v16
                    }
                case -2147483648:
                    for fi17 := 0; fi17 < 10; fi17++ {
                        if len(v14) > 0 {
                            v14[0] = fi17
                        }
                        fmt.Println(v15)
                        switch "\n\t" {
                        case "foo bar":
                            v13 += (v13 >> ((1) & 63))
                        }
                        _ = fi17
                    }
                }
                _ = v15
            } else {
                v13 = (^(v13))
            }
            _ = v13
            _ = v14
        } else {
            var v18 int64 = 546
            _ = v18
        }
        _ = fi12
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        if ((659 >= 255) && (9223372036854775807 < 853)) {
            fmt.Println(fi19)
            v20 := GenS1{F0: 65535, F1: -1}
            var v21 int8 = -1
            _ = v20
            _ = v21
        }
        var v22 int32 = -788
        _ = fi19
        _ = v22
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        var v24 string = (("\n\t" + "") + "foo bar")
        v25 := GenS2{F0: 32767, F1: v24, F2: 0}
        v24 = v24
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 int8 = -126
    _ = v26
    return v26
}
func genFn1() int64 {
    var v27 uint32 = 65535
    fmt.Println(v27)
    for fi28 := 0; fi28 < 4; fi28++ {
        for fi29 := 0; fi29 < 11; fi29++ {
            v30 := GenS3{F0: 148, F1: 341}
            v27 += (v27 << ((4294967295) & 63))
            v31 := GenS4{F0: 1, F1: (985 < 1)}
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi28
    }
    _ = v27
    return 42
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]bool)
    var v3 uint32 = 2147483648
    var v4 uint = 65535
    var v5 int32 = 640
    v6 := make(map[int]int)
    fmt.Println(v4)
    var v32 uint64 = 2
    var v33 bool = (false && (-9223372036854775808 != 42))
    if (32768 != 4294967295) {
        if (v33 && (!(1 < -1))) {
            for fi34 := 0; fi34 < 13; fi34++ {
                v4 += (v1 & 18446744073709551615)
                var v35 int64 = genFn1()
                _ = fi34
                _ = v35
            }
        }
    } else {
        v36 := GenS1{F0: 0, F1: (v5 / ((v5) | 1))}
        v3 = (v3 ^ (v3 - (v3 % ((v3) | 1))))
        _ = v36
    }
    fmt.Println(v32, v33, v5)
    if v33 {
        v1 = 627
        switch 128 {
        case 127:
            var v37 int8 = 1
            _ = v37
        default:
            _ = len(v6)
        }
    } else {
        if (uint(v3) < 255) {
            fmt.Println(v1)
            if v33 {
                for fi38 := 0; fi38 < 11; fi38++ {
                    v39 := GenS3{F0: uint8(v1), F1: int64(v32)}
                    v4 += (v1 - (v4 | (^(v1))))
                    for fi40 := 0; fi40 < 19; fi40++ {
                        v41 := GenS2{F0: 1000, F1: (("" + ("a" + ("" + ""))) + "\n\t"), F2: 204}
                        var v42 int = 489
                        var v43 int16 = -1
                        _ = fi40
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    _ = fi38
                    _ = v39
                }
            }
            fmt.Println(v4)
        }
    }
    for fi44 := 0; fi44 < 5; fi44++ {
        for fi45 := 0; fi45 < 20; fi45++ {
            v46 := GenS4{F0: -1, F1: ("foo bar" < "hello")}
            switch -9223372036854775808 {
            case 128:
                var v47 int32 = (v5 / ((0) | 1))
                _ = v47
            default:
                fmt.Println(v5, v32)
            }
            _ = fi45
            _ = v46
        }
        switch ("\n\t" + ("hello" + "")) {
        case "foo bar":
            var v48 uint16 = 32767
            _ = v48
        case "\n\t":
            var v49 int32 = v5
            _ = v49
        case "":
            v50 := []int{(fi44 / ((fi44) | 1)), 881, 318, (fi44 / (((^(fi44))) | 1))}
            _ = v50
        default:
            v51 := GenS5{F0: 255, F1: (fi44 >> ((uint(v32)) & 63)), F2: -1}
            _ = v51
        }
        var v52 int16 = 975
        _ = fi44
        _ = v52
    }
    var v53 bool = (!v33)
    fmt.Println(v3, v33)
    if v53 {
        fmt.Println(v33)
    } else {
        var v54 uint16 = 32767
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v53
    _ = v6
}
