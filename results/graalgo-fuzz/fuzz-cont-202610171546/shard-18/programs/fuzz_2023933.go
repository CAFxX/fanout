package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int32 }
type GenS2 struct { F0 int; F1 uint64; F2 bool }
type GenS3 struct { F0 string; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int; F1 uint64 }

func main() {
    var v1 uint16 = 65535
    v2 := make(map[string]string)
    var v3 uint = 42
    var v4 int32 = 1
    var v5 uint = (v3 % (((v3 % ((v3) | 1))) | 1))
    v6 := make(map[uint32]uint32)
    var v7 string = ""
    var v8 int64 = -9223372036854775808
    v2[""] = "\n\t"
    if (false || false) {
        v1 = (v1 * (v1 & v1))
        var v9 uint = (v3 | v5)
        var v10 float64 = float64(0.0)
        _ = v10
        _ = v9
    }
    var v11 int16 = 485
    v4 = 42
    if true {
        v1 -= v1
        v12 := []int{-716, 0, 4294967295}
        v13 := v6[484]
        _ = v12
        _ = v13
    }
    switch 0 {
    case 32768:
        v14 := GenS1{F0: (((("hello" + v7) + (v7 + ("a" + v7))) + "a") + ("" + (v7 + v7))), F1: -1, F2: int32(v1)}
        _ = v14
    case 127:
        switch 0 {
        case 676:
            if (v8 <= (v8 & v8)) {
                fmt.Println(v5)
                v1 = uint16(v5)
                var v15 uint = (v5 - 65535)
                _ = v15
            } else {
                for fi16 := 0; fi16 < 12; fi16++ {
                    for fi17 := 0; fi17 < 11; fi17++ {
                        var v18 string = v7
                        var v19 int = fi17
                        v20 := GenS1{F0: v7, F1: (v11 * (v11 ^ -7)), F2: (v4 % ((v4) | 1))}
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    v11 = v11
                    _ = fi16
                }
            }
        case 69:
            for fi21 := 0; fi21 < 19; fi21++ {
                for fi22 := 0; fi22 < 18; fi22++ {
                    if (797 <= 0) {
                        v23 := GenS2{F0: fi21, F1: 0, F2: false}
                        v4 = 1
                        v4 += 42
                        _ = v23
                    }
                    var v24 bool = ((^(v5)) > (v3 % (((^(v5))) | 1)))
                    v7 = ((v7 + ("a" + "")) + (v7 + v7))
                    _ = fi22
                    _ = v24
                }
                if false {
                    for fi25 := 0; fi25 < 11; fi25++ {
                        switch v7 {
                        case "\n\t":
                            v5 -= 691
                        case "hello":
                            if false {
                                var v26 int64 = v8
                                _ = v26
                            }
                        }
                        v5 = 42
                        _ = len(v2)
                        if true {
                            v27 := GenS2{F0: fi21, F1: uint64(v1), F2: true}
                            switch 0 {
                            case 171:
                                var v28 uint16 = v1
                                _ = v28
                            case 127:
                                var v29 uint = (v5 % (((v3 * (^(v3)))) | 1))
                                _ = v29
                            case 42:
                                if true {
                                    fmt.Println(v1, v5)
                                    var v30 uint16 = 32767
                                    _ = v30
                                }
                            default:
                                v11 = v11
                            }
                            _ = v27
                        }
                        delete(v6, 649)
                        _ = fi25
                    }
                    for fi31 := 0; fi31 < 4; fi31++ {
                        var v32 string = "\n\t"
                        switch v7 {
                        case "\n\t":
                            v33 := GenS3{F0: "a", F1: 2147483648, F2: 1}
                            _ = v33
                        case "hello":
                            switch 32767 {
                            case 42:
                                fmt.Println(v7, fi21)
                            case 124:
                                var v34 uint8 = 42
                                _ = v34
                            }
                        }
                        v8 = (v8 | -652)
                        var v35 uint16 = v1
                        v36 := []int{fi21, fi31}
                        _ = fi31
                        _ = v32
                        _ = v35
                        _ = v36
                    }
                    if (-1 > 42) {
                        v11 += (-(v11))
                    }
                }
                switch "foo bar" {
                case "hello":
                    fmt.Println(v7, fi21)
                }
                _ = fi21
            }
        default:
            if (!(((3.14159 + 1e10) * 0.1) <= (0.1 / ((1.5 * 1e10) - (-2.25 / 1e10))))) {
                v37 := GenS1{F0: (v7 + ""), F1: (v11 - v11), F2: v4}
                _ = v37
            } else {
                v38 := GenS4{F0: -490, F1: 0}
                var v39 int64 = int64(v1)
                _ = v38
                _ = v39
            }
        }
    }
    fmt.Println(v1)
    var v40 int = 2147483647
    v41 := []int{120, v40}
    v1 = 619
    v42 := GenS1{F0: "", F1: (v11 >> (((^(v3))) & 63)), F2: v4}
    var v43 uint = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
