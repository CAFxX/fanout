package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint32; F1 uint16; F2 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64; F1 uint64; F2 uint32 }
func genFn0(p0 uint64) uint32 {
    if false {
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 int8 = 1
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 uint8 = 42
            for fi10 := 0; fi10 < 2; fi10++ {
                v9 = 1
                _ = fi10
            }
            var v11 bool = (((1e10 / (float64(0.0) - -2.25)) + float64(3.14159)) == 0.0)
            _ = fi8
            _ = v11
            _ = v9
        }
    }
    var v12 float32 = ((float32(1.5) - (-2.25 + float32(((1e10 - 1.5) / 1.0)))) / 1.5)
    var v13 uint16 = 792
    var v14 uint8 = 0
    switch "a" {
    case "hello":
        if (658 >= (p0 << ((2) & 63))) {
            v15 := GenS1{F0: v14}
            if true {
                for fi16 := 0; fi16 < 13; fi16++ {
                    fmt.Println(v14, v13, v12)
                    v17 := []int{(fi16 >> ((42) & 63)), (fi16 >> ((28) & 63)), 32767, -2147483648}
                    var v18 string = "hello"
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                if (127 >= v14) {
                    v19 := GenS1{F0: v14}
                    _ = v19
                }
            } else {
                v20 := GenS1{F0: (v14 & 228)}
                for fi21 := 0; fi21 < 19; fi21++ {
                    if ((!true) && true) {
                        var v22 int64 = 42
                        if true {
                            v23 := GenS1{F0: v14}
                            _ = v23
                        }
                        _ = v22
                    } else {
                        for fi24 := 0; fi24 < 17; fi24++ {
                            v12 -= -2.25
                            var v25 bool = false
                            v13 -= (v13 + 142)
                            _ = fi24
                            _ = v25
                        }
                        v12 = (1.5 * 1.5)
                    }
                    _ = fi21
                }
                v26 := []int{2147483647, 783}
                _ = v20
                _ = v26
            }
            _ = v15
        } else {
            p0 = 4294967295
            for fi27 := 0; fi27 < 10; fi27++ {
                var v28 bool = (v14 >= 128)
                for fi29 := 0; fi29 < 10; fi29++ {
                    v30 := GenS1{F0: 80}
                    v31 := []int{(^(fi29)), (fi27 + fi29), (^(fi29)), (fi27 & fi29)}
                    var v32 uint = 311
                    _ = fi29
                    _ = v30
                    _ = v31
                    _ = v32
                }
                v33 := []int{}
                _ = fi27
                _ = v28
                _ = v33
            }
        }
    case "a":
        switch 0 {
        case -350:
            if true {
                fmt.Println(v14, v12)
                v34 := GenS1{F0: v14}
                _ = v34
            }
        }
    }
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 42
}
func genFn1(p0 bool, p1 int8) int {
    if (!(!p0)) {
        p1 += p1
        for fi35 := 0; fi35 < 7; fi35++ {
            v36 := GenS1{F0: 255}
            p0 = (696 >= 65535)
            _ = fi35
            _ = v36
        }
    } else {
        for fi37 := 0; fi37 < 11; fi37++ {
            v38 := GenS1{F0: 127}
            _ = fi37
            _ = v38
        }
        p0 = p0
        for fi39 := 0; fi39 < 12; fi39++ {
            if true {
                fmt.Println(p0, p1, p1)
                for fi40 := 0; fi40 < 13; fi40++ {
                    var v41 int16 = 32767
                    switch (fi39 << ((0) & 63)) {
                    case -1:
                        v42 := []int{}
                        if (0.1 != (((1.5 * (-2.25 + (float64(3.14159) + 0.0))) * 1.5) + -2.25)) {
                            v43 := GenS1{F0: 128}
                            _ = v43
                        } else {
                            switch "hello" {
                            case "":
                                p1 -= p1
                            case "hello":
                                var v44 int64 = 1
                                _ = v44
                            }
                        }
                        v41 -= (v41 + v41)
                        _ = v42
                    }
                    v45 := GenS1{F0: 136}
                    _ = fi40
                    _ = v41
                    _ = v45
                }
            } else {
                for fi46 := 0; fi46 < 12; fi46++ {
                    var v47 string = "hello"
                    _ = fi46
                    _ = v47
                }
                for fi48 := 0; fi48 < 7; fi48++ {
                    var v49 uint32 = 65535
                    p1 -= (p1 + (p1 | 100))
                    v50 := []int{(fi39 >> ((18446744073709551615) & 63))}
                    _ = fi48
                    _ = v49
                    _ = v50
                }
                p0 = p0
            }
            _ = fi39
        }
    }
    p0 = true
    v51 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = v51
    return 32767
}

func main() {
    var v1 string = ""
    var v2 uint32 = 937
    v3 := []int{}
    var v4 uint8 = 42
    var v5 int8 = 1
    for fi52 := 0; fi52 < 11; fi52++ {
        for fi53 := 0; fi53 < 3; fi53++ {
            _ = len(v3)
            var v54 uint32 = v2
            var v55 uint8 = v4
            _ = fi53
            _ = v54
            _ = v55
        }
        if true {
            v4 = v4
            for fi56 := 0; fi56 < 18; fi56++ {
                v4 = v4
                var v57 int64 = 42
                _ = fi56
                _ = v57
            }
        }
        _ = fi52
    }
    v4 = v4
    var v58 uint16 = 97
    fmt.Println(v2, v2)
    switch "\n\t" {
    case "hello":
        if false {
            v59 := GenS1{F0: v4}
            _ = len(v3)
            _ = v59
        } else {
            v5 -= 100
        }
    }
    v4 -= v4
    _ = len(v3)
    var v60 uint16 = 65535
    if (!true) {
        v3 = append(v3, -786)
    }
    _ = len(v3)
    v5 -= v5
    var v61 uint16 = 191
    for fi62 := 0; fi62 < 15; fi62++ {
        if false {
            for fi63 := 0; fi63 < 17; fi63++ {
                v64 := GenS2{F0: false, F1: 1000}
                if (0.1 > 0.0) {
                    v61 -= (v61 & v60)
                    fmt.Println(v4, fi62)
                }
                _ = fi63
                _ = v64
            }
            if len(v3) > 0 {
                v3[0] = fi62
            }
        } else {
            for fi65 := 0; fi65 < 13; fi65++ {
                fmt.Println(v2, v5)
                if true {
                    v61 = v58
                    fmt.Println(fi62)
                }
                _ = fi65
            }
        }
        v61 -= v60
        v2 -= 303
        _ = fi62
    }
    if ((^(v2)) > genFn0(4294967295)) {
        var v66 int8 = (v5 & (v5 ^ (v5 % ((-128) | 1))))
        var v67 int32 = 42
        if true {
            v68 := GenS1{F0: v4}
            v69 := GenS3{F0: genFn0(18446744073709551615), F1: v60, F2: 4294967295}
            v60 = v60
            _ = v68
            _ = v69
        } else {
            v70 := GenS4{F0: v66}
            v3 = append(v3, -1)
            var v71 bool = true
            _ = v70
            _ = v71
        }
        _ = v66
        _ = v67
    } else {
        v60 -= v60
        v58 = v60
        for fi72 := 0; fi72 < 6; fi72++ {
            var v73 uint8 = 35
            if (-626 < 42) {
                if (false || ((!false) || true)) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    _ = len(v3)
                    var v74 uint64 = uint64(v4)
                    v75 := GenS5{F0: 997, F1: v74, F2: (v2 / (((v2 | uint32(v60))) | 1))}
                    _ = v74
                    _ = v75
                }
            }
            _ = fi72
            _ = v73
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v60
    _ = v61
}
