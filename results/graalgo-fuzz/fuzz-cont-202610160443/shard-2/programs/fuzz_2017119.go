package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 string; F1 uint8; F2 int32 }
type GenS7 struct { F0 uint; F1 uint64; F2 int16 }
func genFn0(p0 int8, p1 int64, p2 string) int8 {
    switch 42 {
    case 1:
        if ((690 > 380) || true) {
            if false {
                for fi7 := 0; fi7 < 6; fi7++ {
                    if true {
                        fmt.Println(p0, p2)
                        fmt.Println(fi7, p0)
                    }
                    v8 := GenS1{F0: 42, F1: (("foo bar" + p2) + p2), F2: 42}
                    if (false && (((fi7 + 9223372036854775807) != (fi7 << ((126) & 63))) || false)) {
                        for fi9 := 0; fi9 < 12; fi9++ {
                            var v10 float32 = 0.1
                            _ = fi9
                            _ = v10
                        }
                        for fi11 := 0; fi11 < 6; fi11++ {
                            p2 = p2
                            _ = fi11
                        }
                        var v12 string = p2
                        _ = v12
                    }
                    _ = fi7
                    _ = v8
                }
                var v13 int8 = p0
                var v14 int8 = (v13 / ((v13) | 1))
                _ = v13
                _ = v14
            }
            p2 = p2
        }
    case -1:
        v15 := GenS2{F0: 2, F1: 127, F2: 694}
        _ = v15
    default:
        for fi16 := 0; fi16 < 19; fi16++ {
            p1 += p1
            fmt.Println(p0)
            _ = fi16
        }
    }
    p2 = "\n\t"
    p0 -= p0
    switch 42 {
    case 32767:
        v17 := GenS2{F0: 1, F1: 127, F2: 128}
        _ = v17
    default:
        var v18 int16 = 42
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 | p0)
}
func genFn1() uint32 {
    if true {
        if (-34 >= -1) {
            v19 := GenS3{F0: 42, F1: -9223372036854775808}
            _ = v19
        } else {
            v20 := GenS4{F0: 2}
            for fi21 := 0; fi21 < 15; fi21++ {
                var v22 int = (^(fi21))
                _ = fi21
                _ = v22
            }
            switch -1 {
            case 32767:
                if false {
                    v23 := GenS3{F0: 32767, F1: 32767}
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 14; fi24++ {
                        var v25 float64 = 3.14159
                        _ = fi24
                        _ = v25
                    }
                }
            case 4294967295:
                for fi26 := 0; fi26 < 1; fi26++ {
                    var v27 uint16 = 42
                    _ = fi26
                    _ = v27
                }
            case 9223372036854775807:
                var v28 uint = 410
                _ = v28
            default:
                var v29 uint64 = 1
                _ = v29
            }
            _ = v20
        }
        var v30 bool = true
        _ = v30
    } else {
        for fi31 := 0; fi31 < 19; fi31++ {
            var v32 uint64 = 2
            v33 := []int{fi31, fi31}
            v32 = (v32 << ((65535) & 63))
            _ = fi31
            _ = v32
            _ = v33
        }
        var v34 uint32 = 0
        _ = v34
    }
    var v35 uint32 = 169
    v35 += (v35 * (v35 ^ (v35 + v35)))
    v35 = v35
    _ = v35
    return (v35 & v35)
}

func main() {
    var v1 uint32 = 42
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]bool)
    var v4 string = "a"
    var v5 int64 = 42
    var v6 uint16 = 42
    if false {
        if false {
            fmt.Println(v1)
            v36 := GenS3{F0: -791, F1: 861}
            _ = len(v2)
            _ = v36
        } else {
            v37 := GenS4{F0: 2}
            _ = v37
        }
    }
    if true {
        for fi38 := 0; fi38 < 1; fi38++ {
            for fi39 := 0; fi39 < 4; fi39++ {
                v1 -= v1
                if false {
                    switch fi38 {
                    case 868:
                        fmt.Println(v4)
                    case -1:
                        v4 = v4
                    }
                    if false {
                        if true {
                            var v40 string = v4
                            v41 := []int{fi38, fi39}
                            if true {
                                v6 = (v6 - v6)
                                switch fi38 {
                                case 0:
                                    v42 := v3[v1]
                                    _ = v42
                                case 32767:
                                    var v43 uint64 = 1
                                    _ = v43
                                case 9223372036854775807:
                                    v44 := GenS3{F0: -1, F1: (fi39 - fi38)}
                                    _ = v44
                                }
                            }
                            _ = v40
                            _ = v41
                        }
                        _ = len(v3)
                        v45 := GenS1{F0: 1, F1: ("a" + v4), F2: -253}
                        _ = v45
                    } else {
                        var v46 int32 = int32(fi38)
                        if (true && (!(true || true))) {
                            v47 := v2[(^(v1))]
                            v48 := GenS3{F0: 528, F1: (^(fi39))}
                            if true {
                                v6 += (v6 >> ((181) & 63))
                                v5 += v5
                            }
                            _ = v47
                            _ = v48
                        } else {
                            v49 := GenS5{F0: fi39, F1: 32767}
                            _ = v49
                        }
                        v50 := GenS6{F0: "hello", F1: 0, F2: (^(v46))}
                        _ = v46
                        _ = v50
                    }
                } else {
                    var v51 float32 = (0.1 + -2.25)
                    _ = v51
                }
                if (1e10 == 0.0) {
                    var v52 int8 = int8(v6)
                    _ = v52
                } else {
                    if false {
                        v53 := []int{(^(fi38)), (fi38 ^ (-(fi38))), fi39, (fi39 / ((fi39) | 1))}
                        v54 := GenS4{F0: 2}
                        var v55 int64 = v5
                        _ = v53
                        _ = v54
                        _ = v55
                    }
                }
                _ = fi39
            }
            v4 = ("" + v4)
            _ = fi38
        }
        for fi56 := 0; fi56 < 1; fi56++ {
            v5 = v5
            var v57 uint32 = v1
            _ = fi56
            _ = v57
        }
        var v58 uint64 = 42
        _ = v58
    }
    for fi59 := 0; fi59 < 4; fi59++ {
        v1 = 2147483648
        var v60 uint = 154
        _ = fi59
        _ = v60
    }
    v6 += (v6 * (^(v6)))
    var v61 int16 = -32768
    var v62 uint8 = uint8(v1)
    if (!(1 != (v5 & 0))) {
        fmt.Println(v4)
        for fi63 := 0; fi63 < 15; fi63++ {
            if false {
                switch v4 {
                case "":
                    if (v61 <= v61) {
                        switch (fi63 | fi63) {
                        case -825:
                            var v64 uint8 = (v62 | v62)
                            _ = v64
                        case 422:
                            v65 := GenS7{F0: 59, F1: 2, F2: v61}
                            _ = v65
                        default:
                            var v66 int = 32768
                            _ = v66
                        }
                        v5 = -741
                    }
                case "\n\t":
                    v6 = 2
                }
                v67 := []int{(fi63 << ((4294967295) & 63)), fi63, fi63, (fi63 & fi63)}
                _ = v67
            }
            _ = fi63
        }
    }
    v4 = "hello"
    v68 := []int{4294967295, 744}
    var v69 int8 = 51
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("slicelen v68:", len(v68), v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
    _ = v68
    _ = v69
}
