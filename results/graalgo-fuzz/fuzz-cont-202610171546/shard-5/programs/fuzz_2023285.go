package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int32; F1 uint8; F2 bool }
type GenS5 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS6 struct { F0 uint8; F1 int16; F2 int64 }
type GenS7 struct { F0 int8; F1 int32; F2 uint16 }
type GenS8 struct { F0 bool; F1 uint; F2 string }
type GenS9 struct { F0 int16 }
func genFn0(p0 string, p1 int32, p2 uint8) uint64 {
    fmt.Println(p0)
    var v4 uint8 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 0.0
    v3 := make(map[string]int)
    v5 := GenS1{F0: 42}
    v2 = (0.0 / v2)
    v6 := GenS1{F0: 1}
    v7 := GenS1{F0: 4294967295}
    if true {
        switch 4294967295 {
        case 0:
            v8 := GenS2{F0: 503, F1: 42}
            _ = v8
        case 4294967295:
            v9 := GenS3{F0: -486, F1: -792}
            _ = v9
        default:
            switch "\n\t" {
            case "\n\t":
                v2 = 3.14159
            case "a":
                v10 := GenS3{F0: -1, F1: -979}
                _ = v10
            case "hello":
                var v11 uint8 = 165
                _ = v11
            }
        }
        for fi12 := 0; fi12 < 1; fi12++ {
            for fi13 := 0; fi13 < 4; fi13++ {
                v14 := []int{fi13, (-(fi13))}
                _ = len(v3)
                _ = fi13
                _ = v14
            }
            var v15 float64 = float64(float32(3.14159))
            _ = fi12
            _ = v15
        }
        switch 128 {
        case -1:
            if true {
                if ((!false) && (!(true && ("foo bar" > ("a" + ""))))) {
                    var v16 uint = 18446744073709551615
                    _ = v16
                } else {
                    v17 := []int{9223372036854775807, 2147483647, 32767, -1}
                    _ = v17
                }
            } else {
                v3["hello"] = 32767
                if false {
                    v2 = 3.14159
                }
                for fi18 := 0; fi18 < 15; fi18++ {
                    var v19 string = "a"
                    for fi20 := 0; fi20 < 19; fi20++ {
                        v21 := GenS1{F0: genFn0(v19, 42, 54)}
                        _ = fi20
                        _ = v21
                    }
                    _ = fi18
                    _ = v19
                }
            }
        case -2147483648:
            switch 2147483647 {
            case 128:
                v22 := GenS1{F0: 468}
                _ = v22
            default:
                var v23 uint16 = 0
                _ = v23
            }
        case -9223372036854775808:
            for fi24 := 0; fi24 < 13; fi24++ {
                if true {
                    var v25 int16 = -32768
                    if (!(!true)) {
                        v26 := GenS4{F0: 1, F1: 255, F2: (0.1 < (1.5 / 1.5))}
                        if (!true) {
                            var v27 int = (-(fi24))
                            v28 := GenS2{F0: 2, F1: 42}
                            _ = v27
                            _ = v28
                        } else {
                            fmt.Println(v2, v2)
                        }
                        _ = v26
                    } else {
                        v29 := []int{fi24, (-(fi24)), (^(fi24)), fi24}
                        if false {
                            v25 = -32768
                        } else {
                            v25 = (v25 - v25)
                        }
                        var v30 int8 = -128
                        _ = v29
                        _ = v30
                    }
                    v2 = 1e10
                    _ = v25
                } else {
                    v31 := GenS5{F0: 465, F1: 42, F2: 2147483647}
                    var v32 int8 = 100
                    if false {
                        v32 -= (v32 ^ -128)
                        var v33 uint = 660
                        _ = v33
                    } else {
                        if false {
                            v34 := GenS6{F0: 1, F1: 0, F2: 817}
                            _ = v34
                        } else {
                            v2 = v2
                        }
                        fmt.Println(v32)
                    }
                    _ = v31
                    _ = v32
                }
                var v35 float64 = -2.25
                _ = fi24
                _ = v35
            }
        default:
            switch -1 {
            case -1:
                v36 := v1[2147483648]
                _ = v36
            case 32768:
                var v37 uint = 253
                _ = v37
            }
        }
    }
    for fi38 := 0; fi38 < 13; fi38++ {
        v2 = (v2 * 1e10)
        _ = fi38
    }
    var v39 uint8 = 0
    var v40 float32 = float32(0.1)
    for fi41 := 0; fi41 < 20; fi41++ {
        v40 = v40
        v40 = 3.14159
        _ = fi41
    }
    v42 := GenS1{F0: 0}
    var v43 int32 = -287
    var v44 float64 = float64(1.5)
    for fi45 := 0; fi45 < 1; fi45++ {
        v44 = -2.25
        var v46 int64 = -33
        _ = fi45
        _ = v46
    }
    for fi47 := 0; fi47 < 12; fi47++ {
        var v48 int64 = -9223372036854775808
        if false {
            var v49 uint64 = genFn0("", (v43 % (((^(v43))) | 1)), 42)
            if false {
                v50 := []int{fi47, fi47}
                v40 -= float32(v44)
                v51 := GenS7{F0: -42, F1: (v43 & (v43 << ((65535) & 63))), F2: 0}
                _ = v50
                _ = v51
            } else {
                v52 := GenS3{F0: -340, F1: 42}
                v53 := GenS8{F0: (!false), F1: 2, F2: ("a" + "a")}
                _ = v52
                _ = v53
            }
            var v54 float32 = v40
            _ = v49
            _ = v54
        }
        if false {
            v55 := v3[("hello" + "hello")]
            _ = v55
        } else {
            var v56 uint = 65535
            switch "\n\t" {
            case "":
                for fi57 := 0; fi57 < 9; fi57++ {
                    var v58 uint = v56
                    _ = fi57
                    _ = v58
                }
            case "foo bar":
                v59 := GenS8{F0: true, F1: 4294967295, F2: "foo bar"}
                _ = v59
            case "hello":
                v56 -= uint(v39)
            default:
                if false {
                    switch (fi47 ^ (fi47 >> ((v56) & 63))) {
                    case -2147483648:
                        switch fi47 {
                        case 128:
                            for fi60 := 0; fi60 < 13; fi60++ {
                                switch "a" {
                                case "hello":
                                    v61 := GenS9{F0: -1}
                                    _ = v61
                                }
                                var v62 float64 = (float64(v40) / float64(1e10))
                                _ = fi60
                                _ = v62
                            }
                        case -1:
                            for fi63 := 0; fi63 < 17; fi63++ {
                                v64 := GenS9{F0: 1}
                                _ = fi63
                                _ = v64
                            }
                        }
                    default:
                        v65 := []int{-455, fi47, fi47, -487}
                        _ = v65
                    }
                    v66 := GenS1{F0: 4294967295}
                    v48 -= v48
                    _ = v66
                }
            }
            _ = v56
        }
        _ = fi47
        _ = v48
    }
    v67 := []int{-421, 127, 42, 128}
    var v68 uint = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v67:", len(v67), v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v67
    _ = v68
    _ = v7
}
