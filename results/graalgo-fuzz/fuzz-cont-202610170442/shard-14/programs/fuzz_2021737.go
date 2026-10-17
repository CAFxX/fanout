package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS4 struct { F0 int16; F1 int64; F2 int64 }
func genFn0() int64 {
    if (2147483648 == 1) {
        var v6 int32 = -1
        var v7 string = "foo bar"
        var v8 int8 = 127
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 int8 = 127
    v10 := GenS1{F0: (v9 | 79), F1: 42, F2: v9}
    v11 := GenS1{F0: v9, F1: 0, F2: (v9 * -1)}
    _ = v10
    _ = v11
    _ = v9
    return 42
}
func genFn1() uint64 {
    v12 := GenS1{F0: 28, F1: 823, F2: 0}
    for fi13 := 0; fi13 < 17; fi13++ {
        for fi14 := 0; fi14 < 4; fi14++ {
            v15 := GenS2{F0: -842}
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    switch "foo bar" {
    case "hello":
        if (707 > 91) {
            for fi16 := 0; fi16 < 17; fi16++ {
                v17 := []int{-9223372036854775808, fi16, (-(fi16))}
                for fi18 := 0; fi18 < 1; fi18++ {
                    fmt.Println(fi16, fi16)
                    fmt.Println(fi16, fi18)
                    v19 := GenS1{F0: 42, F1: 42, F2: -61}
                    _ = fi18
                    _ = v19
                }
                var v20 int = fi16
                _ = fi16
                _ = v17
                _ = v20
            }
        } else {
            var v21 uint32 = 0
            v21 = 65535
            _ = v21
        }
    }
    _ = v12
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 2
    var v4 int16 = 1000
    var v5 float32 = 3.14159
    if (false || true) {
        var v22 uint8 = 255
        _ = v22
    } else {
        delete(v2, 2147483647)
        var v23 float32 = 1e10
        _ = v23
    }
    switch "a" {
    case "\n\t":
        for fi24 := 0; fi24 < 3; fi24++ {
            v3 -= 2
            _ = fi24
        }
    case "foo bar":
        if (!true) {
            v3 -= genFn1()
        } else {
            switch -2147483648 {
            case 0:
                v3 -= (v3 >> ((2) & 63))
            }
            var v25 int8 = 1
            v5 = float32(1.5)
            _ = v25
        }
    case "":
        if (-422 == 42) {
            v26 := GenS3{F0: (^(v3)), F1: 0, F2: 278}
            v27 := GenS1{F0: 1, F1: 65535, F2: -128}
            v3 = (v3 % (((v3 | (v3 % ((2) | 1)))) | 1))
            _ = v26
            _ = v27
        } else {
            v4 = int16(v3)
            for fi28 := 0; fi28 < 12; fi28++ {
                if (v5 > v5) {
                    var v29 int16 = -658
                    v5 = float32(((3.14159 - (0.0 + 3.14159)) + 1e10))
                    _ = v29
                } else {
                    switch fi28 {
                    case 902:
                        switch "foo bar" {
                        case "foo bar":
                            fmt.Println(fi28)
                        case "hello":
                            v30 := GenS1{F0: -89, F1: 0, F2: 42}
                            _ = v30
                        case "":
                            if true {
                                var v31 bool = (2147483647 != 0)
                                v32 := []int{(-(fi28)), fi28}
                                v33 := GenS4{F0: (v4 % ((0) | 1)), F1: -9223372036854775808, F2: genFn0()}
                                _ = v31
                                _ = v32
                                _ = v33
                            } else {
                                var v34 int16 = v4
                                var v35 float64 = (1e10 - (0.0 - (0.1 / 1.5)))
                                var v36 uint32 = 2147483647
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                        }
                    default:
                        fmt.Println(fi28)
                    }
                    var v37 uint32 = 4294967295
                    for fi38 := 0; fi38 < 4; fi38++ {
                        if (float64(1e10) >= 0.1) {
                            v39 := GenS2{F0: fi28}
                            _ = v39
                        } else {
                            var v40 uint64 = (v3 >> ((657) & 63))
                            _ = v40
                        }
                        _ = fi38
                    }
                    _ = v37
                }
                var v41 float32 = v5
                _ = fi28
                _ = v41
            }
        }
    default:
        switch 1 {
        case -286:
            switch "hello" {
            case "":
                v5 -= -2.25
            default:
                var v42 string = ""
                _ = v42
            }
        case 32767:
            var v43 uint64 = v3
            _ = v43
        case 909:
            v3 = 832
        default:
            var v44 int8 = 85
            _ = v44
        }
    }
    for fi45 := 0; fi45 < 11; fi45++ {
        v5 -= float32(3.14159)
        _ = fi45
    }
    if true {
        var v46 int16 = (v4 % ((v4) | 1))
        var v47 int = 32767
        _ = v46
        _ = v47
    } else {
        if true {
            v3 += (v3 - (v3 | 1))
            if (0 != 0) {
                v4 -= v4
                v48 := v2[1]
                _ = v48
            }
            v49 := GenS2{F0: 731}
            _ = v49
        }
        v50 := GenS4{F0: 1, F1: -9223372036854775808, F2: -9223372036854775808}
        _ = v50
    }
    var v51 int8 = -1
    if false {
        switch 0 {
        case 32767:
            for fi52 := 0; fi52 < 15; fi52++ {
                v4 -= 1
                _ = fi52
            }
        }
    } else {
        var v53 bool = false
        var v54 bool = (!(42 != 32767))
        var v55 uint64 = genFn1()
        _ = v53
        _ = v54
        _ = v55
    }
    fmt.Println(v5, v5, v4)
    var v56 float64 = 0.0
    v57 := []int{127, 0, 2147483647}
    v56 = v56
    v5 = (v5 + float32((v56 / 0.1)))
    var v58 float64 = v56
    var v59 int = 0
    fmt.Println(v3)
    fmt.Println(v4, v3)
    var v60 uint64 = (v3 + (^(v3)))
    v61 := GenS3{F0: v3, F1: -1, F2: 32767}
    v56 = ((((v56 / v56) + float64(1.5)) - (v58 + v56)) * (0.1 + 0.1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v60
    _ = v61
}
