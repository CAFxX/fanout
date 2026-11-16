package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int16) int8 {
    p0 -= p0
    for fi6 := 0; fi6 < 8; fi6++ {
        p0 += p0
        _ = fi6
    }
    _ = p0
    return 42
}
func genFn1() uint32 {
    v7 := GenS1{F0: 34, F1: -2147483648}
    if true {
        var v8 uint64 = 979
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 float64 = 3.14159
            for fi11 := 0; fi11 < 8; fi11++ {
                fmt.Println(v10, v8, fi9)
                fmt.Println(fi11, fi9, v8)
                _ = fi11
            }
            var v12 string = "\n\t"
            _ = fi9
            _ = v10
            _ = v12
        }
        switch "a" {
        case "foo bar":
            if true {
                v8 += 1
                for fi13 := 0; fi13 < 19; fi13++ {
                    fmt.Println(fi13)
                    v8 = 4294967295
                    v8 += (v8 % ((1) | 1))
                    _ = fi13
                }
            } else {
                var v14 string = "a"
                if false {
                    v15 := GenS1{F0: genFn0(1), F1: -970}
                    switch -1 {
                    case 127:
                        v14 = v14
                    case -2147483648:
                        for fi16 := 0; fi16 < 3; fi16++ {
                            fmt.Println(v14, v8, v14)
                            switch -177 {
                            case 4294967295:
                                v17 := GenS1{F0: 127, F1: 0}
                                _ = v17
                            case -841:
                                v18 := []int{}
                                _ = v18
                            case 42:
                                v8 = 582
                            }
                            _ = fi16
                        }
                    case 32768:
                        switch 42 {
                        case 496:
                            if true {
                                var v19 bool = true
                                var v20 uint16 = 2
                                _ = v19
                                _ = v20
                            } else {
                                if (1 >= 588) {
                                    var v21 string = v14
                                    var v22 uint64 = (v8 << ((255) & 63))
                                    for fi23 := 0; fi23 < 6; fi23++ {
                                        v24 := []int{(fi23 * (fi23 * (fi23 << ((65535) & 63)))), (fi23 << ((4294967295) & 63)), fi23}
                                        fmt.Println(v22, v8, fi23)
                                        _ = fi23
                                        _ = v24
                                    }
                                    _ = v21
                                    _ = v22
                                }
                                var v25 int32 = 998
                                for fi26 := 0; fi26 < 1; fi26++ {
                                    switch fi26 {
                                    case -138:
                                        switch fi26 {
                                        case -9223372036854775808:
                                            v14 = v14
                                        case -395:
                                            v27 := GenS2{F0: 162, F1: v8, F2: (v25 + (v25 | 1))}
                                            _ = v27
                                        }
                                    default:
                                        switch (fi26 / ((-287) | 1)) {
                                        case -2147483648:
                                            for fi28 := 0; fi28 < 4; fi28++ {
                                                var v29 uint16 = 42
                                                v30 := GenS1{F0: 42, F1: 2147483647}
                                                _ = fi28
                                                _ = v29
                                                _ = v30
                                            }
                                        case 943:
                                            v31 := GenS2{F0: 9, F1: (v8 & v8), F2: -32}
                                            _ = v31
                                        case 128:
                                            fmt.Println(v25)
                                        default:
                                            var v32 string = ((v14 + (v14 + v14)) + "hello")
                                            _ = v32
                                        }
                                    }
                                    v8 = v8
                                    _ = fi26
                                }
                                _ = v25
                            }
                        default:
                            v33 := []int{}
                            _ = v33
                        }
                    }
                    _ = v15
                } else {
                    if false {
                        fmt.Println(v14, v14)
                        v34 := GenS2{F0: 1, F1: 18446744073709551615, F2: -327}
                        _ = v34
                    } else {
                        for fi35 := 0; fi35 < 13; fi35++ {
                            v36 := []int{fi35, fi35, (fi35 >> ((65535) & 63))}
                            if (!true) {
                                v14 = "foo bar"
                            }
                            _ = fi35
                            _ = v36
                        }
                        fmt.Println(v8, v8)
                    }
                }
                var v37 bool = false
                _ = v14
                _ = v37
            }
        case "a":
            v8 -= 514
        case "":
            v38 := []int{}
            _ = v38
        default:
            v39 := []int{32767, 715, 2147483647}
            _ = v39
        }
        _ = v8
    }
    if true {
        var v40 uint64 = 18446744073709551615
        fmt.Println(v40)
        _ = v40
    }
    for fi41 := 0; fi41 < 5; fi41++ {
        v42 := []int{878, fi41, fi41}
        v43 := GenS3{F0: 582}
        _ = fi41
        _ = v42
        _ = v43
    }
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 2
    v3 := make(map[string]bool)
    v4 := []int{2147483647, 128}
    v5 := make(map[string]bool)
    var v44 uint16 = 646
    v5[("a" + (("foo bar" + ("a" + "a")) + "foo bar"))] = true
    v44 += 1
    var v45 int8 = int8(v2)
    v44 = uint16(v2)
    switch -455 {
    case 127:
        v2 += 1
    }
    var v46 int8 = 1
    fmt.Println(v46)
    v1[""] = 412
    v44 = v44
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v45)
    v45 = (v45 << ((1) & 63))
    var v47 uint64 = 18446744073709551615
    for fi48 := 0; fi48 < 9; fi48++ {
        var v49 uint8 = 119
        switch 128 {
        case 128:
            if (-32768 <= 1) {
                var v50 int8 = int8(v44)
                _ = v50
            } else {
                if false {
                    var v51 string = "foo bar"
                    if (!(1 == 2147483647)) {
                        var v52 int8 = (v46 ^ 127)
                        if true {
                            var v53 float64 = 1.5
                            v45 += (v45 * v45)
                            _ = v53
                        }
                        _ = v52
                    }
                    _ = v51
                }
                v49 += 135
                for fi54 := 0; fi54 < 5; fi54++ {
                    var v55 uint = 450
                    v47 = (^(v47))
                    var v56 uint = 1
                    _ = fi54
                    _ = v55
                    _ = v56
                }
            }
        }
        var v57 int64 = -151
        _ = fi48
        _ = v49
        _ = v57
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
