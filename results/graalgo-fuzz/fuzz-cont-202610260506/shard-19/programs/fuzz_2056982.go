package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint8 }
func genFn0() uint64 {
    switch ("\n\t" + "") {
    case "hello":
        for fi6 := 0; fi6 < 2; fi6++ {
            if true {
                for fi7 := 0; fi7 < 3; fi7++ {
                    v8 := []int{-586, fi6, fi7}
                    var v9 int32 = -1
                    var v10 int = 2147483647
                    switch fi7 {
                    case 32768:
                        fmt.Println(fi7, v10, v10)
                    default:
                        switch fi6 {
                        case 127:
                            fmt.Println(fi6)
                        case 765:
                            switch fi6 {
                            case 32768:
                                var v11 bool = true
                                _ = v11
                            case -9223372036854775808:
                                v10 = (^(fi6))
                            }
                        default:
                            _ = len(v8)
                        }
                    }
                    fmt.Println(fi7, v10)
                    v10 = (^(fi7))
                    v8 = append(v8, fi7)
                    _ = fi7
                    _ = v10
                    _ = v8
                    _ = v9
                }
            }
            v12 := []int{}
            _ = fi6
            _ = v12
        }
    default:
        v13 := GenS1{F0: 62, F1: 300, F2: 42}
        _ = v13
    }
    var v14 uint64 = 583
    for fi15 := 0; fi15 < 11; fi15++ {
        v16 := []int{}
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: 42, F1: 42, F2: 157}
    v14 = v14
    _ = v14
    _ = v17
    return (v14 / ((v14) | 1))
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 2
    var v3 uint8 = 0
    v4 := make(map[uint32]bool)
    v5 := make(map[string]string)
    switch -492 {
    case 4294967295:
        var v18 int32 = -992
        _ = v18
    }
    if (0 <= 2) {
        switch "foo bar" {
        case "":
            for fi19 := 0; fi19 < 7; fi19++ {
                v2 = (v2 % (((v2 * (^(v2)))) | 1))
                fmt.Println(fi19, fi19)
                var v20 bool = (false && (!true))
                _ = fi19
                _ = v20
            }
        }
    }
    v1[991] = false
    _ = len(v4)
    var v21 uint8 = (v3 / (((v3 ^ (v3 | (v3 + v3)))) | 1))
    v22 := []int{-1, -450, 42}
    if (4294967295 == 2) {
        var v23 uint16 = v2
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        fmt.Println(v21)
        _ = v23
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        switch (^(fi24)) {
        case 42:
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 int32 = -269
                _ = fi25
                _ = v26
            }
        }
        if (true && true) {
            v2 -= v2
            switch -526 {
            case -437:
                if false {
                    v3 += v21
                } else {
                    var v27 uint16 = (v2 * 255)
                    _ = v27
                }
            default:
                fmt.Println(v21)
            }
            fmt.Println(v2, fi24)
        } else {
            v28 := GenS1{F0: 1, F1: -1, F2: (v3 & (v21 << ((42) & 63)))}
            _ = v28
        }
        if (1 != 323) {
            v29 := v1[999]
            v3 = (v3 & (v21 >> ((18446744073709551615) & 63)))
            _ = len(v22)
            _ = v29
        } else {
            v3 = (v3 / ((v3) | 1))
            var v30 uint16 = (^(v2))
            _ = v30
        }
        _ = fi24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
