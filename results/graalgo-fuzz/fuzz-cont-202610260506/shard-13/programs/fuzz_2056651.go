package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 int32; F2 uint64 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 int; F2 uint32 }
func genFn0() int {
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 bool = false
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 128, F1: 42}
    var v10 bool = true
    switch -126 {
    case 848:
        if true {
            v11 := []int{32768, -317}
            _ = len(v11)
            v10 = true
            _ = v11
        }
    }
    var v12 uint8 = 127
    _ = v10
    _ = v12
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]int)
    var v3 bool = ((627 != 4294967295) || (!(((false && (("foo bar" + "") <= "\n\t")) && (!true)) || false)))
    var v4 uint16 = 2
    var v5 float64 = (1.5 * (1.5 * (0.0 - 1e10)))
    var v6 int64 = -928
    v13 := GenS1{F0: 128, F1: 42}
    var v14 uint32 = 1
    v15 := GenS1{F0: 128, F1: -128}
    v5 = v5
    fmt.Println(v4, v3)
    if v3 {
        v1 -= v1
        switch 42 {
        case 2147483647:
            for fi16 := 0; fi16 < 9; fi16++ {
                switch fi16 {
                case 2147483647:
                    for fi17 := 0; fi17 < 3; fi17++ {
                        var v18 uint8 = 255
                        v19 := GenS2{F0: -1, F1: (v6 << ((255) & 63))}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                default:
                    fmt.Println(v14)
                }
                delete(v2, 966)
                for fi20 := 0; fi20 < 1; fi20++ {
                    v4 = v4
                    v21 := GenS3{F0: 42}
                    _ = fi20
                    _ = v21
                }
                _ = fi16
            }
        case 4294967295:
            for fi22 := 0; fi22 < 16; fi22++ {
                if false {
                    v23 := GenS4{F0: 522, F1: 0, F2: 262}
                    var v24 int32 = 42
                    for fi25 := 0; fi25 < 15; fi25++ {
                        var v26 string = "\n\t"
                        if false {
                            switch fi25 {
                            case -9223372036854775808:
                                v14 -= uint32(v4)
                            case 958:
                                var v27 uint16 = v4
                                _ = v27
                            }
                            v28 := []int{}
                            if v3 {
                                _ = len(v2)
                                var v29 int = fi25
                                fmt.Println(v6, v3)
                                v24 = v24
                                _ = v29
                            }
                            _ = v28
                        } else {
                            var v30 int = (-(fi22))
                            v31 := GenS1{F0: 86, F1: 109}
                            v32 := GenS5{F0: v26, F1: v6}
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                        _ = fi25
                        _ = v26
                    }
                    _ = v23
                    _ = v24
                }
                if v3 {
                    for fi33 := 0; fi33 < 7; fi33++ {
                        if v3 {
                            var v34 int32 = 2147483647
                            var v35 bool = (1 < 42)
                            var v36 int64 = v6
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                        switch (fi22 / (((-(fi22))) | 1)) {
                        case 9223372036854775807:
                            v37 := GenS5{F0: (((("" + ("hello" + "\n\t")) + "\n\t") + "foo bar") + "\n\t"), F1: v6}
                            var v38 uint32 = (v14 + 2)
                            var v39 uint8 = 255
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                        _ = fi33
                    }
                    v14 -= v14
                }
                _ = fi22
            }
        case 127:
            v40 := GenS6{F0: v3}
            _ = v40
        default:
            switch "" {
            case "":
                for fi41 := 0; fi41 < 10; fi41++ {
                    var v42 int64 = 1
                    v43 := []int{fi41, (^(fi41))}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            case "foo bar":
                for fi44 := 0; fi44 < 17; fi44++ {
                    var v45 uint32 = (v14 + (v14 & v14))
                    v3 = (false && false)
                    if false {
                        _ = len(v2)
                    } else {
                        v46 := []int{fi44}
                        _ = v46
                    }
                    _ = fi44
                    _ = v45
                }
            default:
                v47 := []int{938, 32768, genFn0(), 191}
                _ = v47
            }
        }
    } else {
        var v48 int64 = v6
        v49 := GenS3{F0: 0}
        _ = v48
        _ = v49
    }
    for fi50 := 0; fi50 < 9; fi50++ {
        var v51 float64 = v5
        _ = fi50
        _ = v51
    }
    v52 := GenS7{F0: (v1 | (v1 / ((0) | 1))), F1: 4294967295, F2: v14}
    v14 = 386
    fmt.Println(v5, v3)
    var v53 uint = 1
    var v54 uint = v53
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
