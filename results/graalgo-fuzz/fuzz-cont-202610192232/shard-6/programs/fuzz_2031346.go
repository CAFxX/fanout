package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint16; F1 bool; F2 int8 }
func genFn0() string {
    v4 := []int{}
    var v5 int16 = 1
    _ = v4
    _ = v5
    return (("hello" + "hello") + ("foo bar" + ""))
}
func genFn1() int {
    if true {
        var v6 float32 = 3.14159
        var v7 uint16 = 255
        if (65535 < 310) {
            var v8 uint = 42
            _ = v8
        }
        _ = v6
        _ = v7
    } else {
        if (255 == 1) {
            var v9 uint16 = 1
            v9 -= v9
            _ = v9
        } else {
            var v10 string = ""
            if (2 < 128) {
                for fi11 := 0; fi11 < 2; fi11++ {
                    if (-860 <= -1) {
                        v12 := GenS1{F0: 304}
                        var v13 string = ((v10 + v10) + (((v10 + (v10 + (v10 + v10))) + v10) + ("a" + "hello")))
                        var v14 bool = true
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    var v15 int8 = 100
                    _ = fi11
                    _ = v15
                }
                v10 = v10
            }
            _ = v10
        }
        var v16 bool = (1 <= -1)
        _ = v16
    }
    var v17 int64 = -634
    v18 := []int{2147483647, -122}
    var v19 float32 = 3.14159
    _ = v17
    _ = v18
    _ = v19
    return 129
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint32 = 65535
    v3 := make(map[uint32]string)
    if (!(1000 >= 0)) {
        v20 := GenS1{F0: v2}
        _ = v20
    }
    if false {
        for fi21 := 0; fi21 < 9; fi21++ {
            v1 -= (v1 & v1)
            _ = fi21
        }
        v2 = 2
        v1 -= v1
    } else {
        v22 := GenS2{F0: -1, F1: 0, F2: -128}
        _ = v22
    }
    if false {
        v23 := []int{2147483647, 128, 0, -591}
        _ = v23
    }
    v1 = v1
    if true {
        switch 829 {
        case -1:
            var v24 int32 = -2147483648
            _ = v24
        default:
            v25 := []int{32767, 32768, 2147483647}
            _ = v25
        }
        v26 := []int{}
        v27 := v3[v2]
        _ = v26
        _ = v27
    } else {
        v2 = 42
        for fi28 := 0; fi28 < 19; fi28++ {
            switch genFn1() {
            case 32768:
                v3[(v2 ^ 2)] = "\n\t"
            case 9223372036854775807:
                var v29 uint8 = 226
                _ = v29
            default:
                if ((1.5 * 0.1) < 3.14159) {
                    v2 = (v2 & (^(v2)))
                    var v30 float32 = 3.14159
                    if false {
                        delete(v3, 467)
                        switch "hello" {
                        case "foo bar":
                            v2 = (v2 ^ (v2 - 870))
                        case "a":
                            var v31 uint = 426
                            _ = v31
                        default:
                            v2 = v2
                        }
                        v32 := v3[v2]
                        _ = v32
                    } else {
                        v33 := []int{128, (fi28 >> ((0) & 63))}
                        for fi34 := 0; fi34 < 17; fi34++ {
                            var v35 int8 = 42
                            v1 = (v1 & v1)
                            _ = fi34
                            _ = v35
                        }
                        _ = v33
                    }
                    _ = v30
                }
            }
            _ = fi28
        }
    }
    v1 = v1
    switch 1 {
    case 1:
        v2 -= 65535
    case -29:
        v2 -= 696
    default:
        for fi36 := 0; fi36 < 18; fi36++ {
            switch -2147483648 {
            case 42:
                var v37 string = ("\n\t" + "hello")
                _ = v37
            case -817:
                v38 := GenS2{F0: -43, F1: (^(v2)), F2: 1}
                _ = v38
            case 471:
                switch fi36 {
                case -373:
                    _ = len(v3)
                case 127:
                    v39 := []int{fi36, genFn1()}
                    _ = v39
                case 32767:
                    var v40 int8 = -34
                    _ = v40
                default:
                    for fi41 := 0; fi41 < 13; fi41++ {
                        v1 += (v1 ^ uint64(v2))
                        if (!(0 > 471)) {
                            v2 = (v2 + v2)
                            switch "a" {
                            case "foo bar":
                                var v42 float32 = -2.25
                                _ = v42
                            case "hello":
                                if (-1 == 42) {
                                    v2 -= (^(v2))
                                    if true {
                                        var v43 int = fi41
                                        switch "foo bar" {
                                        case "a":
                                            v44 := []int{v43, (fi41 | v43), fi41, genFn1()}
                                            _ = v44
                                        case "hello":
                                            fmt.Println(fi36, fi36, v1)
                                        default:
                                            v45 := GenS2{F0: -1, F1: v2, F2: 90}
                                            _ = v45
                                        }
                                        if false {
                                            var v46 string = (("\n\t" + "foo bar") + "")
                                            v47 := GenS3{F0: 65535, F1: (false && (((^(v2)) == 75) || (!false))), F2: -1}
                                            _ = v46
                                            _ = v47
                                        }
                                        _ = v43
                                    }
                                    v1 -= (v1 >> ((265) & 63))
                                    v48 := []int{(fi36 * fi41)}
                                    _ = v48
                                } else {
                                    v1 = v1
                                }
                            case "":
                                fmt.Println(fi36)
                            }
                            switch fi36 {
                            case -2147483648:
                                var v49 float32 = 0.1
                                _ = v49
                            default:
                                if (2 >= 122) {
                                    var v50 float64 = -2.25
                                    _ = len(v3)
                                    var v51 bool = (65535 >= 65535)
                                    _ = v50
                                    _ = v51
                                } else {
                                    fmt.Println(fi36, fi36, v1)
                                    v2 = v2
                                    delete(v3, 2)
                                    v52 := []int{fi41, (fi41 / ((fi36) | 1))}
                                    var v53 string = "\n\t"
                                    _ = v52
                                    _ = v53
                                }
                            }
                        }
                        fmt.Println(fi36, v1, fi36)
                        _ = fi41
                    }
                }
            }
            _ = fi36
        }
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
}
