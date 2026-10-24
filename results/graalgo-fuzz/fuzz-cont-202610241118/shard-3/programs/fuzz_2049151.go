package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 bool; F1 int; F2 uint }
type GenS4 struct { F0 uint64; F1 int; F2 int16 }
type GenS5 struct { F0 uint; F1 uint32 }
func genFn0(p0 uint8) uint32 {
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint8 = (p0 | (^(p0)))
        if ((v8 | 42) <= (p0 & v8)) {
            v9 := []int{-69, (fi7 | fi7), fi7, (fi7 >> ((2) & 63))}
            for fi10 := 0; fi10 < 1; fi10++ {
                v8 = v8
                var v11 string = "hello"
                _ = fi10
                _ = v11
            }
            _ = v9
        }
        fmt.Println(v8, p0, fi7)
        _ = fi7
        _ = v8
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 12; fi13++ {
            fmt.Println(fi12, fi12)
            _ = fi13
        }
        for fi14 := 0; fi14 < 10; fi14++ {
            fmt.Println(fi12, fi12, p0)
            p0 = p0
            if false {
                v15 := []int{}
                p0 = (p0 | p0)
                _ = v15
            } else {
                var v16 uint8 = p0
                switch fi12 {
                case -977:
                    var v17 float64 = 1.5
                    _ = v17
                default:
                    var v18 int = (fi14 / ((fi12) | 1))
                    _ = v18
                }
                var v19 int32 = 42
                if false {
                    fmt.Println(p0, fi12, v19)
                }
                _ = v16
                _ = v19
            }
            _ = fi14
        }
        _ = fi12
    }
    v20 := GenS1{F0: 1, F1: -128, F2: 169}
    _ = p0
    _ = v20
    return 2147483647
}

func main() {
    var v1 float64 = (-2.25 / 1e10)
    var v2 uint64 = 0
    var v3 string = ""
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]string)
    v6 := make(map[int]bool)
    switch 128 {
    case 2147483647:
        v21 := GenS1{F0: 255, F1: 11, F2: 42}
        _ = v21
    case 32768:
        for fi22 := 0; fi22 < 3; fi22++ {
            fmt.Println(fi22, fi22)
            var v23 uint = 512
            if (true && true) {
                for fi24 := 0; fi24 < 4; fi24++ {
                    var v25 uint64 = 0
                    v26 := v5[(-(fi24))]
                    if true {
                        fmt.Println(v23, fi24)
                        var v27 bool = false
                        var v28 int = (fi22 + 127)
                        _ = v27
                        _ = v28
                    } else {
                        v25 = 42
                        v29 := []int{(fi24 / (((fi22 >> ((1) & 63))) | 1)), fi22, 4294967295}
                        switch fi22 {
                        case 60:
                            var v30 int64 = 316
                            v3 = v26
                            fmt.Println(v25, v30)
                            _ = v30
                        }
                        v1 += v1
                        _ = v29
                    }
                    v31 := []int{fi22}
                    fmt.Println(v23, v23)
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v31
                }
            }
            _ = fi22
            _ = v23
        }
    case 128:
        v32 := GenS2{F0: 2, F1: 272}
        _ = v32
    default:
        fmt.Println(v2, v1, v3)
    }
    if false {
        var v33 uint = uint(v2)
        v2 += v2
        fmt.Println(v33, v33)
        _ = v33
    } else {
        var v34 uint32 = 2147483647
        if (!true) {
            v6[-205] = true
            v34 = 309
            v34 = (v34 | 42)
        }
        _ = v34
    }
    v3 = v3
    for fi35 := 0; fi35 < 3; fi35++ {
        if false {
            switch (fi35 - fi35) {
            case 0:
                switch ("foo bar" + (v3 + "")) {
                case "\n\t":
                    if (1 < 42) {
                        for fi36 := 0; fi36 < 18; fi36++ {
                            v3 = (v3 + (v3 + v3))
                            v1 = v1
                            _ = fi36
                        }
                    } else {
                        v37 := GenS2{F0: 65535, F1: genFn0(74)}
                        if (796 < 42) {
                            v38 := GenS2{F0: 2, F1: genFn0(127)}
                            switch v3 {
                            case "":
                                v39 := GenS1{F0: 65535, F1: -1, F2: 38}
                                _ = v39
                            default:
                                if (!(v2 < (v2 / (((v2 ^ (v2 + (v2 * (^(v2)))))) | 1)))) {
                                    v40 := GenS3{F0: false, F1: 0, F2: 1}
                                    v41 := GenS2{F0: 0, F1: 2147483647}
                                    _ = v40
                                    _ = v41
                                } else {
                                    var v42 uint64 = (v2 - 272)
                                    _ = v42
                                }
                            }
                            _ = v38
                        }
                        if true {
                            v43 := GenS4{F0: (^(v2)), F1: -9223372036854775808, F2: -826}
                            _ = v43
                        }
                        _ = v37
                    }
                default:
                    delete(v4, 2147483648)
                }
            }
        } else {
            var v44 string = v3
            if (true || false) {
                v2 += (v2 & v2)
                var v45 float64 = v1
                v46 := GenS5{F0: 2, F1: 725}
                _ = v45
                _ = v46
            } else {
                fmt.Println(v44)
                v47 := v5[fi35]
                v3 = "a"
                _ = v47
            }
            _ = v44
        }
        _ = fi35
    }
    v1 = 0.0
    var v48 uint8 = uint8(v2)
    v49 := v4[genFn0((v48 / ((v48) | 1)))]
    switch v3 {
    case "a":
        if false {
            v48 = v48
            delete(v6, -980)
            var v50 uint64 = (v2 & v2)
            _ = v50
        } else {
            v51 := v5[32767]
            v49 = (v49 * v49)
            _ = v51
        }
    }
    var v52 float32 = 3.14159
    v48 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v52
    _ = v6
}
