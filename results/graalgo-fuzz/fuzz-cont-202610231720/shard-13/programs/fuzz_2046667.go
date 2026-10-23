package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 uint8; F2 uint8 }
type GenS6 struct { F0 uint32; F1 uint16; F2 int32 }
func genFn0(p0 uint16) uint32 {
    p0 = 32767
    var v5 int = -1
    v6 := GenS1{F0: (p0 ^ (p0 - p0)), F1: "foo bar"}
    switch (v5 | (v5 % ((v5) | 1))) {
    case 32768:
        v7 := []int{v5, (v5 + (v5 << ((207) & 63)))}
        _ = v7
    case 128:
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := []int{}
            if (v5 > v5) {
                switch fi8 {
                case 128:
                    for fi10 := 0; fi10 < 17; fi10++ {
                        if true {
                            var v11 uint16 = (p0 % (((p0 + 32767)) | 1))
                            var v12 int = fi8
                            _ = v11
                            _ = v12
                        } else {
                            if len(v9) > 0 {
                                v9[0] = -960
                            }
                            fmt.Println(fi8)
                            if ((128 != 128) && true) {
                                p0 += p0
                                var v13 uint32 = 65535
                                if true {
                                    fmt.Println(v5, v13)
                                }
                                v14 := GenS1{F0: 0, F1: "hello"}
                                switch "" {
                                case "\n\t":
                                    v13 -= v13
                                }
                                _ = v13
                                _ = v14
                            }
                        }
                        var v15 uint = 18446744073709551615
                        _ = fi10
                        _ = v15
                    }
                }
                p0 = p0
                v16 := GenS1{F0: p0, F1: "a"}
                _ = v16
            }
            p0 += (p0 | (^(p0)))
            _ = fi8
            _ = v9
        }
    case 42:
        for fi17 := 0; fi17 < 11; fi17++ {
            v18 := []int{178, fi17}
            if (fi17 <= (v5 % ((v5) | 1))) {
                var v19 int = 2147483647
                _ = v19
            }
            switch "a" {
            case "a":
                var v20 uint64 = 2
                _ = v20
            case "\n\t":
                if true {
                    if (!false) {
                        switch "foo bar" {
                        case "foo bar":
                            v5 = 127
                        case "a":
                            for fi21 := 0; fi21 < 10; fi21++ {
                                v22 := GenS1{F0: 32767, F1: ("hello" + "a")}
                                _ = fi21
                                _ = v22
                            }
                        case "":
                            for fi23 := 0; fi23 < 7; fi23++ {
                                v5 += -2147483648
                                var v24 int32 = 42
                                v25 := GenS1{F0: (p0 >> ((1) & 63)), F1: "foo bar"}
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                        default:
                            switch v5 {
                            case 2147483647:
                                var v26 uint64 = 1
                                _ = v26
                            case 127:
                                switch (^(fi17)) {
                                case 2147483647:
                                    fmt.Println(v5, v5, p0)
                                }
                            default:
                                for fi27 := 0; fi27 < 4; fi27++ {
                                    v28 := GenS2{F0: 402, F1: fi17}
                                    _ = fi27
                                    _ = v28
                                }
                            }
                        }
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                        if len(v18) > 0 {
                            v18[0] = fi17
                        }
                    } else {
                        if (!((!(807 < 890)) && (390 >= 4294967295))) {
                            fmt.Println(v5)
                            v29 := GenS3{F0: 42}
                            _ = v29
                        }
                    }
                } else {
                    for fi30 := 0; fi30 < 3; fi30++ {
                        fmt.Println(p0)
                        _ = len(v18)
                        _ = fi30
                    }
                }
            default:
                switch 2147483647 {
                case 4294967295:
                    p0 -= p0
                case 1:
                    var v31 uint = 0
                    _ = v31
                }
            }
            _ = fi17
            _ = v18
        }
    default:
        for fi32 := 0; fi32 < 3; fi32++ {
            for fi33 := 0; fi33 < 16; fi33++ {
                if (!(((p0 & (p0 * 2)) < (p0 << ((255) & 63))) || (-995 <= -915))) {
                    var v34 int16 = 0
                    _ = v34
                }
                var v35 float32 = (-2.25 * (((0.0 / (1.5 + 1.5)) - float32((3.14159 - float64(((1e10 * 0.1) / 1.0))))) / 3.14159))
                var v36 uint8 = 151
                _ = fi33
                _ = v35
                _ = v36
            }
            v37 := GenS3{F0: 2147483647}
            for fi38 := 0; fi38 < 3; fi38++ {
                var v39 int8 = 42
                var v40 int8 = (v39 % ((v39) | 1))
                _ = fi38
                _ = v39
                _ = v40
            }
            _ = fi32
            _ = v37
        }
    }
    _ = p0
    _ = v5
    _ = v6
    return 909
}
func genFn1() uint32 {
    var v41 int32 = 1
    v42 := []int{-1, 4294967295}
    _ = v41
    _ = v42
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 4294967295
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    var v43 int64 = 42
    v44 := GenS1{F0: 1, F1: "\n\t"}
    var v45 string = ""
    for fi46 := 0; fi46 < 11; fi46++ {
        v45 = "foo bar"
        var v47 uint8 = 2
        switch "hello" {
        case "foo bar":
            v48 := []int{32767, (fi46 ^ (fi46 + int(v47)))}
            _ = v48
        case "":
            var v49 int8 = 42
            _ = v49
        case "hello":
            if false {
                if false {
                    v47 = v47
                    for fi50 := 0; fi50 < 18; fi50++ {
                        v47 = uint8(v2)
                        v51 := GenS4{F0: 7}
                        _ = fi50
                        _ = v51
                    }
                    var v52 int8 = -128
                    _ = v52
                } else {
                    var v53 string = ""
                    _ = v53
                }
                if (false || true) {
                    v47 = (v47 << ((18446744073709551615) & 63))
                    switch (fi46 | (-(fi46))) {
                    case -2147483648:
                        v47 -= v47
                    }
                    v45 = ((v45 + v45) + v45)
                }
            }
        default:
            if false {
                for fi54 := 0; fi54 < 19; fi54++ {
                    v55 := GenS5{F0: v45, F1: (^(v47)), F2: (v47 % ((v47) | 1))}
                    _ = fi54
                    _ = v55
                }
            }
        }
        _ = fi46
        _ = v47
    }
    if (-2147483648 != 598) {
        fmt.Println(v2, v45, v45)
        var v56 uint64 = (v2 | (v2 >> ((157) & 63)))
        _ = v56
    } else {
        v43 += v43
    }
    if ("a" > (((v45 + ((v45 + v45) + ("" + v45))) + "a") + ("\n\t" + v45))) {
        var v57 int8 = 111
        v43 -= -9223372036854775808
        fmt.Println(v43)
        _ = v57
    } else {
        if false {
            var v58 int64 = 42
            var v59 int16 = 42
            _ = v58
            _ = v59
        } else {
            for fi60 := 0; fi60 < 4; fi60++ {
                for fi61 := 0; fi61 < 13; fi61++ {
                    if false {
                        if ((2 != 0) && (true && true)) {
                            var v62 uint = 4294967295
                            _ = v62
                        }
                        v2 += 0
                    } else {
                        var v63 int32 = 42
                        _ = v63
                    }
                    v64 := GenS3{F0: int32(v43)}
                    v43 = 1
                    _ = fi61
                    _ = v64
                }
                if false {
                    v3[fi60] = 2147483647
                    for fi65 := 0; fi65 < 13; fi65++ {
                        var v66 int32 = 2147483647
                        _ = fi65
                        _ = v66
                    }
                }
                _ = fi60
            }
            fmt.Println(v45, v45, v2)
            fmt.Println(v2, v43)
        }
        v67 := GenS6{F0: 2147483648, F1: 779, F2: 1}
        _ = v67
    }
    if true {
        if false {
            v45 = v45
            v45 = v45
            var v68 int = 127
            _ = v68
        }
    } else {
        v43 -= -1
        var v69 uint32 = 4294967295
        _ = v69
    }
    if true {
        var v70 uint64 = (v2 | v2)
        var v71 int8 = 0
        _ = v70
        _ = v71
    }
    v45 = v45
    v72 := []int{32767, 4294967295, -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v72:", len(v72), v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v72
}
