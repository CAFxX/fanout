package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int64 }
func genFn0() uint16 {
    if (((false || true) && (false || ((("hello" + ("" + "\n\t")) + ("\n\t" + "")) != "\n\t"))) && (!(-904 != 1))) {
        v6 := GenS1{F0: 32767, F1: 0}
        switch -354 {
        case 127:
            var v7 float64 = (1.5 / (-2.25 + 1e10))
            _ = v7
        default:
            v8 := []int{-625}
            _ = v8
        }
        var v9 int = -779
        _ = v6
        _ = v9
    } else {
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := []int{fi10, fi10, (fi10 << ((1) & 63))}
            _ = fi10
            _ = v11
        }
        var v12 uint16 = 1
        _ = v12
    }
    if false {
        var v13 uint32 = 4294967295
        var v14 uint = 701
        if true {
            v15 := GenS1{F0: 721, F1: 449}
            fmt.Println(v13, v14)
            v16 := GenS1{F0: 1, F1: v13}
            _ = v15
            _ = v16
        }
        _ = v13
        _ = v14
    } else {
        var v17 uint64 = 42
        _ = v17
    }
    switch "" {
    case "a":
        if (true || true) {
            var v18 int32 = 42
            _ = v18
        } else {
            v19 := GenS1{F0: 255, F1: 709}
            _ = v19
        }
    case "foo bar":
        var v20 uint8 = 255
        _ = v20
    default:
        switch -197 {
        case 128:
            if (724 >= -32768) {
                switch "foo bar" {
                case "a":
                    v21 := GenS2{F0: -273, F1: 1}
                    _ = v21
                case "hello":
                    for fi22 := 0; fi22 < 17; fi22++ {
                        v23 := []int{}
                        var v24 uint16 = 42
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                case "\n\t":
                    for fi25 := 0; fi25 < 10; fi25++ {
                        var v26 int = fi25
                        v27 := []int{150, 128}
                        var v28 bool = false
                        _ = fi25
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                }
                for fi29 := 0; fi29 < 16; fi29++ {
                    for fi30 := 0; fi30 < 8; fi30++ {
                        fmt.Println(fi29, fi29)
                        v31 := []int{946, (^(fi29))}
                        var v32 string = ("hello" + "")
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    switch fi29 {
                    case 1:
                        fmt.Println(fi29)
                    case 87:
                        if (700 < fi29) {
                            var v33 int8 = -128
                            _ = v33
                        } else {
                            var v34 int32 = 544
                            var v35 uint32 = 441
                            for fi36 := 0; fi36 < 12; fi36++ {
                                var v37 int32 = v34
                                v34 -= (v37 * (v34 & v34))
                                v34 += (v37 | v34)
                                _ = fi36
                                _ = v37
                            }
                            _ = v34
                            _ = v35
                        }
                    }
                    _ = fi29
                }
            } else {
                var v38 int64 = -182
                for fi39 := 0; fi39 < 3; fi39++ {
                    v40 := GenS1{F0: 255, F1: 1}
                    for fi41 := 0; fi41 < 2; fi41++ {
                        v38 += 635
                        _ = fi41
                    }
                    v42 := []int{(fi39 % (((fi39 - (fi39 / (((fi39 - fi39)) | 1)))) | 1)), fi39, fi39, fi39}
                    _ = fi39
                    _ = v40
                    _ = v42
                }
                fmt.Println(v38)
                _ = v38
            }
        default:
            var v43 string = "hello"
            _ = v43
        }
    }
    v44 := []int{-2147483648, 1, -9223372036854775808, -2147483648}
    v45 := GenS3{F0: "", F1: 107}
    _ = v44
    _ = v45
    return 42
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 255
    var v3 float32 = 0.0
    var v4 int64 = 0
    var v5 float32 = (1e10 / v3)
    _ = len(v1)
    v2 = 2
    switch "foo bar" {
    case "hello":
        var v46 bool = false
        _ = v46
    }
    switch 32767 {
    case 1:
        for fi47 := 0; fi47 < 19; fi47++ {
            v48 := []int{(fi47 * fi47), (fi47 & (-(fi47))), (fi47 << ((18446744073709551615) & 63))}
            v49 := GenS4{F0: fi47}
            _ = fi47
            _ = v48
            _ = v49
        }
    case 103:
        v50 := []int{}
        _ = v50
    case 0:
        v51 := GenS5{F0: 65535}
        _ = v51
    default:
        var v52 float64 = float64((v3 + v5))
        _ = v52
    }
    v53 := GenS3{F0: "\n\t", F1: 173}
    if ((!true) || true) {
        var v54 uint = 65535
        var v55 int8 = 42
        _ = v54
        _ = v55
    }
    v2 += v2
    v56 := GenS6{F0: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v56
}
