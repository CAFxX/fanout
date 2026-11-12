package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint64) uint8 {
    var v5 uint32 = 779
    v6 := []int{0, 2147483647, 32768}
    v5 = 282
    _ = p0
    _ = v5
    _ = v6
    return 127
}
func genFn1(p0 bool, p1 uint64) uint {
    var v7 uint = 42
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: 820, F1: 0, F2: 128}
        v10 := GenS1{F0: fi8, F1: 65535, F2: 42}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint = (v7 * v7)
    fmt.Println(p1, p0, p0)
    v12 := []int{}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v7
    return (v7 + v7)
}

func main() {
    var v1 int32 = 546
    v2 := make(map[int]bool)
    v3 := []string{("\n\t" + "foo bar"), ("foo bar" + ("a" + ("" + "hello"))), "\n\t"}
    var v4 float64 = -2.25
    v13 := GenS1{F0: 502, F1: 271, F2: 1}
    var v14 int16 = 0
    if false {
        v2[662] = false
        if true {
            fmt.Println(v1)
        }
    }
    delete(v2, 32768)
    _ = len(v2)
    if true {
        delete(v2, -882)
        fmt.Println(v1, v14)
    }
    v15 := GenS2{F0: 2147483647, F1: (^(v14))}
    var v16 int8 = 0
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 uint32 = 854
        _ = fi17
        _ = v18
    }
    _ = len(v2)
    var v19 float32 = 1e10
    var v20 int16 = v14
    if true {
        var v21 int32 = v1
        for fi22 := 0; fi22 < 8; fi22++ {
            v14 += (v14 << ((2) & 63))
            switch fi22 {
            case 127:
                v23 := GenS3{F0: -779, F1: 118, F2: 1}
                _ = v23
            }
            _ = fi22
        }
        _ = v21
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        v14 -= (^(v14))
        _ = fi24
    }
    v19 = v19
    switch -563 {
    case 1:
        v25 := GenS2{F0: 2, F1: v14}
        _ = v25
    case -2147483648:
        switch -2147483648 {
        case -440:
            switch 32768 {
            case -1:
                if ((v14 & v20) > (-(v14))) {
                    for fi26 := 0; fi26 < 3; fi26++ {
                        v14 = v14
                        var v27 int32 = v1
                        _ = fi26
                        _ = v27
                    }
                    v28 := GenS1{F0: 1, F1: 2, F2: 127}
                    for fi29 := 0; fi29 < 3; fi29++ {
                        v30 := GenS1{F0: fi29, F1: 42, F2: 255}
                        _ = fi29
                        _ = v30
                    }
                    _ = v28
                }
            case -551:
                v31 := GenS3{F0: v1, F1: 0, F2: 4294967295}
                _ = v31
            case 32768:
                var v32 float64 = float64(v19)
                _ = v32
            default:
                var v33 uint8 = 110
                _ = v33
            }
        default:
            v34 := GenS4{F0: 0}
            _ = v34
        }
    case -1:
        if true {
            var v35 uint8 = 1
            var v36 bool = true
            v35 = genFn0(779)
            _ = v35
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        if true {
            for fi38 := 0; fi38 < 20; fi38++ {
                v39 := GenS1{F0: 4294967295, F1: 42, F2: 1}
                if true {
                    _ = len(v3)
                    var v40 uint8 = 146
                    v41 := GenS1{F0: (fi37 % ((4294967295) | 1)), F1: 0, F2: genFn0(uint64(v40))}
                    if false {
                        v42 := GenS4{F0: 8}
                        _ = v42
                    }
                    _ = v40
                    _ = v41
                } else {
                    var v43 bool = (v20 != v20)
                    if v43 {
                        if (66 == v1) {
                            var v44 uint64 = 479
                            switch fi38 {
                            case 4294967295:
                                if (32767 <= 2) {
                                    switch (fi38 - (fi38 / ((-2147483648) | 1))) {
                                    case 9223372036854775807:
                                        v19 += (float32((0.1 - (v4 + (v4 - v4)))) + 3.14159)
                                    case 4294967295:
                                        fmt.Println(v1, v1)
                                    case -1:
                                        v16 = 0
                                    }
                                    v19 = v19
                                }
                            case 42:
                                var v45 uint32 = 0
                                var v46 uint = 2
                                _ = v45
                                _ = v46
                            }
                            _ = v44
                        }
                        v47 := GenS3{F0: v1, F1: 127, F2: 1}
                        fmt.Println(v19)
                        _ = v47
                    } else {
                        _ = len(v3)
                        var v48 int32 = 0
                        if (42 != 400) {
                            v1 += (v48 | 0)
                            fmt.Println(v48)
                            _ = len(v2)
                        } else {
                            fmt.Println(v16, v48)
                            if (v19 <= v19) {
                                v1 += v1
                                var v49 int32 = 0
                                _ = v49
                            } else {
                                var v50 int = (fi38 ^ fi37)
                                var v51 uint = 255
                                fmt.Println(v50)
                                _ = v50
                                _ = v51
                            }
                        }
                        _ = v48
                    }
                    _ = v43
                }
                _ = fi38
                _ = v39
            }
            var v52 uint8 = 255
            if len(v3) > 0 {
                v3[0] = ("a" + "foo bar")
            }
            _ = v52
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
