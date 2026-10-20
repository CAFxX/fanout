package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int8 }
type GenS2 struct { F0 string; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int16 }
func genFn0() int {
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: -154, F1: 722, F2: -128}
        switch 32768 {
        case 0:
            if ((false || true) || (!true)) {
                for fi8 := 0; fi8 < 6; fi8++ {
                    var v9 int32 = -2147483648
                    _ = fi8
                    _ = v9
                }
            } else {
                v10 := GenS1{F0: -171, F1: 1, F2: -61}
                _ = v10
            }
        case -9223372036854775808:
            var v11 float64 = 1.5
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    var v12 uint32 = 4294967295
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 uint8 = 42
        v15 := []int{}
        v14 = (v14 - v14)
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = v12
    return 32767
}

func main() {
    v1 := []int{0, -9223372036854775808, 933}
    var v2 bool = false
    var v3 int16 = -1
    var v4 uint64 = 0
    var v5 uint64 = (v4 + 87)
    fmt.Println(v3, v4)
    v16 := GenS2{F0: "a", F1: 65535, F2: 220}
    fmt.Println(v3)
    var v17 uint16 = 255
    for fi18 := 0; fi18 < 1; fi18++ {
        _ = len(v1)
        v17 += v17
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi18
    }
    fmt.Println(v5)
    switch "a" {
    case "\n\t":
        if v2 {
            v17 = v17
        } else {
            v19 := GenS1{F0: 268, F1: 592, F2: 1}
            _ = v19
        }
    case "foo bar":
        fmt.Println(v2)
    }
    v20 := GenS2{F0: (("a" + "hello") + ("\n\t" + ("\n\t" + (("\n\t" + "hello") + "foo bar")))), F1: uint16(v4), F2: (v17 >> ((18446744073709551615) & 63))}
    v2 = (0 == 42)
    var v21 uint16 = 65535
    var v22 uint = 591
    for fi23 := 0; fi23 < 7; fi23++ {
        v22 = v22
        if len(v1) > 0 {
            v1[0] = 1
        }
        if false {
            if len(v1) > 0 {
                v1[0] = (fi23 ^ fi23)
            }
            v21 -= (v17 & 42)
            v24 := GenS3{F0: 255, F1: v3}
            _ = v24
        }
        _ = fi23
    }
    v22 = v22
    for fi25 := 0; fi25 < 9; fi25++ {
        v17 -= (v21 >> (((v22 | v22)) & 63))
        v26 := GenS2{F0: "foo bar", F1: uint16(v4), F2: v21}
        var v27 uint32 = 612
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 uint16 = (v17 ^ (^(v17)))
    for fi29 := 0; fi29 < 14; fi29++ {
        for fi30 := 0; fi30 < 19; fi30++ {
            if v2 {
                v31 := GenS3{F0: 204, F1: 1}
                if len(v1) > 0 {
                    v1[0] = fi29
                }
                v28 += 187
                _ = v31
            }
            _ = fi30
        }
        if true {
            switch ("foo bar" + ("a" + "")) {
            case "a":
                for fi32 := 0; fi32 < 6; fi32++ {
                    if ((!(-32768 <= (v3 ^ (v3 >> ((255) & 63))))) || (int32(v28) > -1)) {
                        if len(v1) > 0 {
                            v1[0] = fi29
                        }
                    }
                    switch "\n\t" {
                    case "\n\t":
                        var v33 int32 = 1
                        _ = v33
                    case "":
                        v34 := GenS2{F0: "a", F1: (^(v17)), F2: uint16(v5)}
                        _ = v34
                    case "hello":
                        fmt.Println(v2, v22)
                    }
                    if true {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v35 string = ("" + ("a" + ("foo bar" + (("\n\t" + ("\n\t" + "hello")) + (("a" + "hello") + "hello")))))
                        if v2 {
                            v2 = v2
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            var v36 uint32 = 2147483648
                            _ = v36
                        }
                        _ = v35
                    }
                    _ = fi32
                }
            case "hello":
                v37 := GenS3{F0: 42, F1: (v3 | -184)}
                _ = v37
            case "":
                _ = len(v1)
            default:
                for fi38 := 0; fi38 < 10; fi38++ {
                    v4 = v4
                    var v39 int32 = -211
                    _ = fi38
                    _ = v39
                }
            }
            var v40 string = "\n\t"
            _ = len(v1)
            _ = v40
        }
        if v2 {
            if false {
                var v41 uint8 = 42
                if len(v1) > 0 {
                    v1[0] = fi29
                }
                _ = v41
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v5, v21, v21)
        }
        _ = fi29
    }
    var v42 uint64 = (v4 ^ 483)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
