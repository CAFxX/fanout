package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint16 }

func main() {
    var v1 string = ""
    var v2 uint8 = 92
    v3 := make(map[int]int)
    v4 := make(map[string]int)
    var v5 uint = 42
    if (766 != 42) {
        for fi6 := 0; fi6 < 8; fi6++ {
            v2 -= v2
            _ = fi6
        }
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            v2 += v2
            v9 := []int{fi7, (fi8 >> ((18446744073709551615) & 63)), -2147483648}
            fmt.Println(v2, v2)
            _ = fi8
            _ = v9
        }
        var v10 int64 = 1
        var v11 float64 = 3.14159
        _ = fi7
        _ = v10
        _ = v11
    }
    if false {
        switch 32767 {
        case 4294967295:
            var v12 uint64 = 2
            _ = v12
        case 2147483647:
            if true {
                switch 588 {
                case 4294967295:
                    if ((^(v5)) > 2) {
                        for fi13 := 0; fi13 < 19; fi13++ {
                            v5 += v5
                            var v14 int8 = 23
                            _ = fi13
                            _ = v14
                        }
                        var v15 uint16 = 0
                        var v16 uint32 = 741
                        _ = v15
                        _ = v16
                    }
                case 1:
                    v17 := GenS1{F0: 0}
                    _ = v17
                default:
                    _ = len(v4)
                }
                if (true && false) {
                    v18 := v3[42]
                    for fi19 := 0; fi19 < 19; fi19++ {
                        var v20 int16 = 1
                        _ = fi19
                        _ = v20
                    }
                    _ = v18
                } else {
                    var v21 uint8 = v2
                    switch -2147483648 {
                    case -1:
                        var v22 bool = ((false && false) && true)
                        _ = v22
                    case 0:
                        for fi23 := 0; fi23 < 20; fi23++ {
                            fmt.Println(fi23, fi23)
                            v24 := []int{fi23, fi23}
                            _ = fi23
                            _ = v24
                        }
                    case 2147483647:
                        if true {
                            v25 := GenS2{F0: int64(v2), F1: 32767}
                            var v26 uint32 = 65535
                            switch v1 {
                            case "":
                                if true {
                                    var v27 uint16 = 744
                                    _ = v27
                                }
                            case "hello":
                                if (0 > -128) {
                                    v28 := GenS1{F0: -9223372036854775808}
                                    v2 = (v21 | v2)
                                    _ = v28
                                }
                            case "\n\t":
                                v29 := GenS2{F0: -9223372036854775808, F1: 42}
                                _ = v29
                            default:
                                var v30 uint32 = (v26 | (v26 & v26))
                                _ = v30
                            }
                            _ = v25
                            _ = v26
                        } else {
                            v21 -= 127
                            v1 = ("" + "\n\t")
                        }
                    }
                    if (2 > 916) {
                        fmt.Println(v21, v21, v1)
                        var v31 int16 = 32767
                        v5 -= (^(v5))
                        _ = v31
                    }
                    _ = v21
                }
                for fi32 := 0; fi32 < 12; fi32++ {
                    v1 = v1
                    v33 := []int{fi32}
                    _ = fi32
                    _ = v33
                }
            } else {
                fmt.Println(v2, v5)
            }
        default:
            v2 -= uint8(v5)
        }
        switch -306 {
        case 42:
            v34 := []int{}
            _ = v34
        case -353:
            for fi35 := 0; fi35 < 5; fi35++ {
                v5 = 342
                v2 += v2
                _ = fi35
            }
        case 4294967295:
            _ = len(v4)
        default:
            v36 := GenS2{F0: 0, F1: 32767}
            _ = v36
        }
    } else {
        v3[-9223372036854775808] = int(v5)
        for fi37 := 0; fi37 < 13; fi37++ {
            _ = len(v3)
            v5 = 65535
            for fi38 := 0; fi38 < 10; fi38++ {
                fmt.Println(v2, v2)
                v39 := []int{32767, fi37, fi38, (fi37 >> (((v5 + v5)) & 63))}
                _ = fi38
                _ = v39
            }
            _ = fi37
        }
    }
    if (!true) {
        fmt.Println(v5)
    }
    var v40 int = 9223372036854775807
    v41 := []int{}
    v5 = (v5 << ((v5) & 63))
    var v42 uint = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
