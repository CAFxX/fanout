package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint32; F1 string; F2 int }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 int16; F1 int32 }
type GenS8 struct { F0 bool }
func genFn0(p0 uint16, p1 uint64, p2 bool) int {
    v4 := GenS1{F0: -2147483648, F1: -1}
    p1 += 198
    var v5 float64 = -2.25
    v6 := []int{32768, 647, 32767}
    var v7 uint64 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 string = "\n\t"
    var v2 uint8 = 127
    var v3 uint64 = 0
    if (v2 >= (v2 << ((uint(v3)) & 63))) {
        var v8 bool = (1 == 32767)
        v2 = (v2 + v2)
        v1 = "foo bar"
        _ = v8
    } else {
        v3 += v3
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := GenS2{F0: v2, F1: 42, F2: (v2 ^ 255)}
            v3 = v3
            var v11 uint8 = v2
            _ = fi9
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 20; fi12++ {
            if false {
                switch fi12 {
                case -124:
                    fmt.Println(fi12, v1)
                case 492:
                    var v13 int64 = 42
                    _ = v13
                case 1:
                    if (true && false) {
                        v14 := []int{}
                        v2 = (v2 / ((v2) | 1))
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        _ = v14
                    }
                default:
                    if true {
                        var v15 int64 = 9223372036854775807
                        switch (^(fi12)) {
                        case 32768:
                            switch ((v1 + "\n\t") + v1) {
                            case "a":
                                v16 := []int{100, 127, (fi12 ^ genFn0(255, 408, false))}
                                _ = v16
                            case "":
                                v17 := GenS2{F0: v2, F1: 230, F2: v2}
                                _ = v17
                            case "\n\t":
                                var v18 int = genFn0(42, v3, true)
                                _ = v18
                            }
                        case 767:
                            v15 = v15
                        case 4294967295:
                            var v19 int = (fi12 << ((65535) & 63))
                            _ = v19
                        default:
                            for fi20 := 0; fi20 < 15; fi20++ {
                                v21 := GenS3{F0: 267, F1: "foo bar", F2: (fi20 / (((^(fi12))) | 1))}
                                v22 := []int{fi12, fi12, fi20, (fi20 << ((4294967295) & 63))}
                                if true {
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v23 float64 = -2.25
                                        v24 := GenS3{F0: 1, F1: "", F2: (fi12 ^ fi12)}
                                        if true {
                                            var v25 uint32 = 0
                                            var v26 uint16 = 255
                                            _ = len(v22)
                                            if (uint(v26) == 255) {
                                                if true {
                                                    fmt.Println(v23, v15)
                                                    var v27 float64 = v23
                                                    if ((-385 != 999) && true) {
                                                        var v28 uint16 = v26
                                                        _ = v28
                                                    } else {
                                                        if (!false) {
                                                            if len(v22) >= 2 {
                                                                v22 = v22[1:len(v22)]
                                                            }
                                                        } else {
                                                            switch "hello" {
                                                            case "":
                                                                if false {
                                                                    v1 = "\n\t"
                                                                    var v29 int64 = -1
                                                                    _ = v29
                                                                }
                                                            case "hello":
                                                                var v30 int = fi20
                                                                _ = v30
                                                            }
                                                            var v31 uint64 = (v3 << ((1) & 63))
                                                            if len(v22) > 0 {
                                                                v22[0] = (^(fi12))
                                                            }
                                                            var v32 int64 = v15
                                                            var v33 int = (fi20 / ((genFn0((^(v26)), v31, (!false))) | 1))
                                                            _ = v31
                                                            _ = v32
                                                            _ = v33
                                                        }
                                                        v34 := GenS4{F0: 0, F1: 0}
                                                        _ = v34
                                                    }
                                                    _ = v27
                                                } else {
                                                    v23 += v23
                                                }
                                            } else {
                                                v35 := GenS1{F0: 288, F1: 1}
                                                _ = v35
                                            }
                                            _ = v25
                                            _ = v26
                                        } else {
                                            v2 = v2
                                            v36 := GenS5{F0: 991}
                                            v22 = append(v22, fi12)
                                            _ = v36
                                        }
                                        _ = v23
                                        _ = v24
                                    case "\n\t":
                                        switch fi12 {
                                        case -2147483648:
                                            fmt.Println(fi12, fi20, v15)
                                        case 93:
                                            switch (v1 + "a") {
                                            case "\n\t":
                                                if len(v22) > 0 {
                                                    v22[0] = (fi20 % ((fi12) | 1))
                                                }
                                            case "foo bar":
                                                fmt.Println(v15, v15, fi12)
                                            case "a":
                                                if true {
                                                    fmt.Println(fi20, fi12)
                                                    var v37 uint64 = v3
                                                    _ = v37
                                                }
                                            default:
                                                v2 = 255
                                            }
                                        }
                                    default:
                                        v38 := GenS5{F0: 178}
                                        _ = v38
                                    }
                                    v2 = v2
                                    var v39 uint8 = (v2 << ((65535) & 63))
                                    switch 128 {
                                    case 32767:
                                        var v40 uint64 = 26
                                        _ = v40
                                    case -162:
                                        v41 := GenS2{F0: v2, F1: 65535, F2: v2}
                                        _ = v41
                                    default:
                                        v2 += (v39 & v39)
                                    }
                                    v1 = ((v1 + "") + (("a" + "foo bar") + v1))
                                    _ = v39
                                } else {
                                    if (v15 <= (v15 % ((v15) | 1))) {
                                        var v42 int32 = -2147483648
                                        fmt.Println(v1, fi12, fi20)
                                        fmt.Println(v3, v42, v3)
                                        _ = v42
                                    } else {
                                        var v43 float32 = 1e10
                                        _ = v43
                                    }
                                    v44 := GenS3{F0: 65535, F1: "a", F2: fi20}
                                    v2 += v2
                                    fmt.Println(v15)
                                    _ = v44
                                }
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        }
                        _ = v15
                    }
                }
                var v45 bool = (v1 <= v1)
                _ = v45
            } else {
                var v46 int16 = -32768
                switch ("\n\t" + (v1 + "")) {
                case "":
                    v47 := GenS1{F0: -1, F1: -128}
                    _ = v47
                case "foo bar":
                    v1 = "\n\t"
                case "a":
                    var v48 float64 = (1e10 * 1e10)
                    _ = v48
                default:
                    v2 = 42
                }
                if false {
                    var v49 uint = 2
                    v50 := []int{fi12, (fi12 + fi12), 127}
                    _ = v49
                    _ = v50
                }
                _ = v46
            }
            fmt.Println(v3, v2)
            v3 = (v3 / ((v3) | 1))
            _ = fi12
        }
    }
    switch v1 {
    case "":
        var v51 uint16 = 32767
        _ = v51
    case "a":
        v1 = "foo bar"
    case "\n\t":
        v52 := GenS6{F0: 202, F1: -1}
        _ = v52
    default:
        var v53 uint32 = 1
        _ = v53
    }
    fmt.Println(v3, v3)
    for fi54 := 0; fi54 < 8; fi54++ {
        v55 := []int{(fi54 / ((127) | 1)), (fi54 >> ((887) & 63)), (fi54 | genFn0(65535, (v3 << ((568) & 63)), (-62 <= 64)))}
        v56 := GenS7{F0: 1, F1: 0}
        var v57 int8 = 1
        _ = fi54
        _ = v55
        _ = v56
        _ = v57
    }
    v58 := []int{31}
    v3 = v3
    v58 = append(v58, -632)
    var v59 uint32 = 65535
    _ = len(v58)
    var v60 int32 = -2147483648
    v61 := GenS3{F0: v59, F1: "\n\t", F2: -9223372036854775808}
    if (v60 == v60) {
        for fi62 := 0; fi62 < 9; fi62++ {
            v63 := GenS5{F0: 0}
            var v64 int64 = -1
            for fi65 := 0; fi65 < 8; fi65++ {
                if false {
                    var v66 uint32 = 1
                    if len(v58) >= 2 {
                        v58 = v58[1:len(v58)]
                    }
                    var v67 uint8 = v2
                    _ = v66
                    _ = v67
                } else {
                    v68 := GenS7{F0: 32767, F1: (v60 % (((v60 + (v60 >> ((125) & 63)))) | 1))}
                    v59 -= 371
                    _ = v68
                }
                v69 := GenS8{F0: false}
                _ = fi65
                _ = v69
            }
            _ = fi62
            _ = v63
            _ = v64
        }
        v60 -= v60
    } else {
        if (((v60 & (v60 - v60)) >= int32(v59)) || true) {
            var v70 uint16 = 1
            for fi71 := 0; fi71 < 17; fi71++ {
                var v72 int32 = v60
                _ = fi71
                _ = v72
            }
            v3 = uint64(v2)
            _ = v70
        }
        var v73 uint8 = v2
        for fi74 := 0; fi74 < 19; fi74++ {
            var v75 int32 = (v60 % ((int32(fi74)) | 1))
            _ = fi74
            _ = v75
        }
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v58:", len(v58), v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v58
    _ = v59
    _ = v60
    _ = v61
}
