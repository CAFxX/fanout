package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 uint16; F2 int32 }
type GenS3 struct { F0 int8; F1 int; F2 uint32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    var v3 int64 = -828
    v4 := make(map[string]uint32)
    v5 := make(map[string]bool)
    v6 := GenS1{F0: (((32767 >= 32767) && true) || (1.5 < 3.14159))}
    if (-2.25 == (-2.25 / (0.0 - 1e10))) {
        v3 += v3
    } else {
        if (0.0 >= 3.14159) {
            delete(v4, "\n\t")
        } else {
            v3 -= v3
            var v7 int32 = 308
            fmt.Println(v7, v3)
            _ = v7
        }
    }
    v3 += v3
    if (1 < 515) {
        v8 := []int{127}
        var v9 int32 = 0
        _ = v8
        _ = v9
    } else {
        var v10 float32 = ((0.0 * ((-2.25 * 1.5) / (3.14159 + 0.1))) / 0.1)
        v11 := GenS2{F0: 9223372036854775807, F1: 0, F2: 42}
        if false {
            var v12 bool = (3.14159 == 0.0)
            var v13 int64 = v3
            if ((!(!v12)) && (-571 >= 4294967295)) {
                for fi14 := 0; fi14 < 16; fi14++ {
                    for fi15 := 0; fi15 < 2; fi15++ {
                        v16 := GenS3{F0: 1, F1: fi14, F2: 740}
                        v12 = (((float64(1.5) != ((1.5 / 0.1) * 1e10)) && (!v12)) || v12)
                        _ = fi15
                        _ = v16
                    }
                    _ = fi14
                }
                v12 = (((-1 != -876) && ((-2147483648 > 378) && (false || v12))) && false)
                fmt.Println(v13)
            } else {
                v17 := []int{-233}
                for fi18 := 0; fi18 < 2; fi18++ {
                    switch fi18 {
                    case 1:
                        var v19 float64 = -2.25
                        _ = v19
                    case 2147483647:
                        if (!(true && true)) {
                            for fi20 := 0; fi20 < 19; fi20++ {
                                v13 = (v3 - (v13 * (v13 % (((v13 >> ((0) & 63))) | 1))))
                                v10 = v10
                                v12 = (!(0 <= -1))
                                _ = fi20
                            }
                            var v21 uint64 = 0
                            var v22 uint = 1
                            _ = v21
                            _ = v22
                        } else {
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                        }
                    case -548:
                        var v23 string = ("\n\t" + "foo bar")
                        _ = v23
                    default:
                        var v24 int32 = 2147483647
                        _ = v24
                    }
                    _ = fi18
                }
                _ = v17
            }
            _ = v12
            _ = v13
        } else {
            if false {
                v3 = v3
            } else {
                if true {
                    v25 := GenS1{F0: false}
                    delete(v2, 2147483647)
                    _ = v25
                }
                var v26 uint64 = 568
                v10 = v10
                _ = v26
            }
        }
        _ = v10
        _ = v11
    }
    switch -9223372036854775808 {
    case 32768:
        v3 -= v3
    case 467:
        v27 := []int{9223372036854775807, 32768}
        _ = v27
    }
    _ = len(v2)
    v28 := []int{32768, -69, -701, int(v3)}
    if len(v28) > 0 {
        v28[0] = 9223372036854775807
    }
    fmt.Println(v3)
    v29 := GenS3{F0: 42, F1: -2147483648, F2: 697}
    v3 = (v3 * v3)
    for fi30 := 0; fi30 < 16; fi30++ {
        fmt.Println(v3, v3)
        fmt.Println(v3)
        v1[42] = fi30
        _ = fi30
    }
    if (!true) {
        if true {
            var v31 uint64 = 2
            fmt.Println(v31, v31)
            switch "a" {
            case "foo bar":
                v32 := GenS3{F0: 1, F1: 865, F2: 2147483647}
                _ = v32
            case "a":
                fmt.Println(v3, v3)
            case "hello":
                delete(v5, "")
            default:
                _ = len(v28)
            }
            _ = v31
        } else {
            v3 = (^(v3))
        }
    }
    var v33 uint32 = 4294967295
    for fi34 := 0; fi34 < 8; fi34++ {
        var v35 int64 = v3
        _ = fi34
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
