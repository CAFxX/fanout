package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int; F1 int8; F2 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint64) uint16 {
    v4 := GenS1{F0: 133, F1: 2147483648}
    var v5 bool = false
    _ = p0
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 int64 = 42
    var v2 float64 = 0.0
    var v3 int16 = 42
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 float64 = v2
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v2 += v2
    v3 = (v3 / (((v3 % (((v3 - (v3 & 1000))) | 1))) | 1))
    if (!true) {
        v9 := []int{2147483647, -802, -2147483648}
        var v10 float64 = v2
        _ = v10
        _ = v9
    } else {
        switch ("hello" + ("" + (("foo bar" + (("foo bar" + "a") + "")) + "a"))) {
        case "":
            var v11 uint64 = 988
            _ = v11
        case "hello":
            for fi12 := 0; fi12 < 18; fi12++ {
                v2 -= (-2.25 * v2)
                v2 += v2
                v13 := GenS2{F0: fi12, F1: -128, F2: v1}
                _ = fi12
                _ = v13
            }
        case "foo bar":
            switch "\n\t" {
            case "":
                v1 = v1
            case "\n\t":
                fmt.Println(v1)
            }
        }
        var v14 string = "foo bar"
        if false {
            v14 = (v14 + v14)
        } else {
            v15 := []int{118, -2147483648, 32767}
            v16 := GenS3{F0: (v1 - -978)}
            for fi17 := 0; fi17 < 18; fi17++ {
                if true {
                    v15 = append(v15, fi17)
                }
                var v18 uint = 4294967295
                _ = fi17
                _ = v18
            }
            _ = v15
            _ = v16
        }
        _ = v14
    }
    fmt.Println(v2, v2)
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := GenS1{F0: (v3 >> ((645) & 63)), F1: 607}
        for fi21 := 0; fi21 < 6; fi21++ {
            switch "a" {
            case "hello":
                var v22 int8 = 39
                _ = v22
            case "foo bar":
                v23 := []int{(-(fi21))}
                _ = v23
            default:
                fmt.Println(v2, fi21)
                if false {
                    if false {
                        v2 += v2
                    }
                } else {
                    switch fi21 {
                    case 32768:
                        v24 := GenS3{F0: v1}
                        _ = v24
                    case -1:
                        var v25 uint32 = 2147483647
                        _ = v25
                    }
                    v2 -= 1.5
                    var v26 int8 = 100
                    v27 := []int{(fi19 + 799), fi19}
                    var v28 int16 = v3
                    if false {
                        v26 += (v26 & (v26 & 48))
                        if len(v27) >= 2 {
                            v27 = v27[1:len(v27)]
                        }
                        if ((0.0 * float32(v2)) >= (-2.25 - 1e10)) {
                            var v29 uint8 = 1
                            v3 += v3
                            v3 = (^(v3))
                            _ = v29
                        } else {
                            v28 = (v28 | v28)
                            switch fi19 {
                            case 1:
                                var v30 uint16 = 255
                                _ = v30
                            case 32767:
                                v2 = v2
                            }
                            v2 = v2
                        }
                    } else {
                        if false {
                            var v31 int32 = -887
                            var v32 uint = 1
                            var v33 uint32 = uint32(v32)
                            _ = len(v27)
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                        if false {
                            switch fi19 {
                            case -898:
                                fmt.Println(fi21)
                            }
                            if (255 != 42) {
                                if (float32((-2.25 + 0.1)) <= -2.25) {
                                    if false {
                                        v28 = (v28 << ((0) & 63))
                                    } else {
                                        v28 += v28
                                    }
                                    v3 -= v3
                                    v27 = append(v27, fi19)
                                } else {
                                    if false {
                                        switch (-(fi19)) {
                                        case 127:
                                            var v34 string = "foo bar"
                                            _ = v34
                                        default:
                                            var v35 int = 1
                                            _ = v35
                                        }
                                        fmt.Println(v2, fi21, fi21)
                                    } else {
                                        v28 -= v3
                                        if len(v27) > 0 {
                                            v27[0] = fi19
                                        }
                                    }
                                    if len(v27) >= 2 {
                                        v27 = v27[1:len(v27)]
                                    }
                                }
                            }
                            if true {
                                if len(v27) >= 2 {
                                    v27 = v27[1:len(v27)]
                                }
                            }
                        } else {
                            switch 799 {
                            case -941:
                                fmt.Println(v1)
                            case 88:
                                var v36 string = "a"
                                _ = v36
                            default:
                                fmt.Println(fi19, v28)
                            }
                            if (true || true) {
                                v37 := GenS2{F0: fi21, F1: 0, F2: v1}
                                var v38 uint = 2
                                _ = v37
                                _ = v38
                            } else {
                                fmt.Println(v1, fi19)
                            }
                        }
                    }
                    v1 += v1
                    var v39 uint8 = 51
                    v40 := GenS4{F0: 2147483647, F1: (v39 * v39)}
                    _ = v26
                    _ = v27
                    _ = v28
                    _ = v39
                    _ = v40
                }
                switch fi21 {
                case 4294967295:
                    var v41 uint = 1
                    _ = v41
                case -2147483648:
                    v42 := []int{}
                    _ = v42
                default:
                    switch fi21 {
                    case 359:
                        var v43 int32 = 1
                        _ = v43
                    }
                    switch "hello" {
                    case "":
                        v3 = (^(v3))
                    default:
                        var v44 bool = (!((false || false) && (!false)))
                        _ = v44
                    }
                    v2 = v2
                }
                v45 := []int{(fi19 / ((32767) | 1)), 32768, fi19}
                _ = v45
            }
            _ = fi21
        }
        var v46 uint64 = 903
        _ = fi19
        _ = v20
        _ = v46
    }
    v1 -= (v1 / ((-1) | 1))
    if ("a" >= "foo bar") {
        v1 = -9223372036854775808
    }
    if false {
        v47 := GenS1{F0: v3, F1: 1}
        _ = v47
    } else {
        var v48 int64 = 1
        fmt.Println(v3)
        var v49 uint32 = 2147483647
        _ = v48
        _ = v49
    }
    var v50 int64 = (v1 + (v1 ^ v1))
    if (v50 < 9223372036854775807) {
        v51 := GenS1{F0: (v3 % (((-(v3))) | 1)), F1: 4294967295}
        v52 := GenS3{F0: -656}
        var v53 bool = (338 > 42)
        _ = v51
        _ = v52
        _ = v53
    } else {
        fmt.Println(v2, v3, v50)
    }
    for fi54 := 0; fi54 < 11; fi54++ {
        switch "a" {
        case "hello":
            for fi55 := 0; fi55 < 10; fi55++ {
                v50 -= (v50 + (v1 % ((v50) | 1)))
                var v56 uint16 = 119
                _ = fi55
                _ = v56
            }
        }
        _ = fi54
    }
    var v57 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v50
    _ = v57
}
