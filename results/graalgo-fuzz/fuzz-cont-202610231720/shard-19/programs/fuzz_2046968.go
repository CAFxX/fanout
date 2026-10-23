package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 bool; F1 uint16 }
func genFn0() int {
    for fi7 := 0; fi7 < 15; fi7++ {
        switch fi7 {
        case 127:
            var v8 uint16 = 2
            _ = v8
        case 397:
            v9 := []int{fi7, fi7}
            _ = v9
        default:
            v10 := []int{(fi7 - fi7), fi7, -1, -152}
            _ = v10
        }
        _ = fi7
    }
    switch -550 {
    case -9223372036854775808:
        var v11 uint64 = 0
        _ = v11
    case -1:
        var v12 float64 = 1.5
        _ = v12
    default:
        v13 := GenS1{F0: (!false), F1: 2147483648, F2: 2}
        _ = v13
    }
    if (false && (100 < 0)) {
        if true {
            var v14 uint = 1
            v14 -= v14
            switch 32767 {
            case -305:
                v14 = (^(v14))
            default:
                if ("" != "") {
                    fmt.Println(v14)
                }
            }
            _ = v14
        }
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := []int{}
            var v17 int8 = 100
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    return 1
}
func genFn1() uint64 {
    var v18 float64 = float64(0.0)
    if false {
        for fi19 := 0; fi19 < 8; fi19++ {
            if false {
                v20 := GenS1{F0: true, F1: 42, F2: 128}
                for fi21 := 0; fi21 < 19; fi21++ {
                    fmt.Println(fi19, v18, fi21)
                    v22 := []int{(fi21 * fi21), fi21, fi21}
                    var v23 float64 = (v18 + (v18 + float64(-2.25)))
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                v18 = float64(1.5)
                _ = v20
            }
            var v24 int16 = -865
            switch "" {
            case "":
                for fi25 := 0; fi25 < 5; fi25++ {
                    v26 := GenS2{F0: 44}
                    v27 := []int{(fi25 >> ((4294967295) & 63)), 127, (fi19 & 4294967295), fi25}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            case "foo bar":
                if true {
                    if false {
                        v18 = v18
                        switch -1 {
                        case 42:
                            for fi28 := 0; fi28 < 20; fi28++ {
                                var v29 float32 = -2.25
                                _ = fi28
                                _ = v29
                            }
                        case -9223372036854775808:
                            switch "foo bar" {
                            case "hello":
                                for fi30 := 0; fi30 < 15; fi30++ {
                                    v31 := []int{42, fi19, fi30, (fi19 ^ (fi30 * (fi19 & fi30)))}
                                    _ = fi30
                                    _ = v31
                                }
                            case "":
                                v32 := []int{fi19, fi19, genFn0(), fi19}
                                _ = v32
                            default:
                                v33 := []int{}
                                _ = v33
                            }
                        case 32768:
                            if (1 > 32767) {
                                v34 := []int{9223372036854775807, fi19, (fi19 - genFn0())}
                                v18 = v18
                                v18 -= 1e10
                                _ = v34
                            }
                        default:
                            for fi35 := 0; fi35 < 10; fi35++ {
                                v18 = v18
                                _ = fi35
                            }
                        }
                        v36 := []int{}
                        _ = v36
                    } else {
                        var v37 bool = (393 <= 0)
                        v24 = (v24 & v24)
                        _ = v37
                    }
                    if false {
                        v38 := GenS3{F0: 65535, F1: 774, F2: 65535}
                        v39 := []int{}
                        fmt.Println(fi19)
                        _ = v38
                        _ = v39
                    }
                } else {
                    v40 := []int{}
                    v24 -= (v24 ^ -586)
                    _ = v40
                }
            case "a":
                v24 = v24
            }
            _ = fi19
            _ = v24
        }
        fmt.Println(v18)
    }
    v18 -= v18
    v41 := GenS4{F0: "", F1: -387}
    _ = v18
    _ = v41
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 1
    v3 := []uint32{561}
    var v4 uint = 652
    v5 := []string{"a", "\n\t", "hello"}
    v6 := make(map[int]string)
    var v42 int64 = -812
    v2 += 127
    v43 := v1[788]
    if (!true) {
        fmt.Println(v43)
        fmt.Println(v2)
    }
    v44 := GenS1{F0: (!false), F1: 0, F2: 255}
    v45 := GenS4{F0: (("" + v43) + ("a" + "a")), F1: v42}
    fmt.Println(v4, v4)
    if (float64(0.0) >= (3.14159 * float64(1e10))) {
        for fi46 := 0; fi46 < 14; fi46++ {
            if true {
                var v47 bool = false
                fmt.Println(v42)
                _ = v47
            } else {
                var v48 uint16 = 2
                _ = v48
            }
            v49 := GenS5{F0: 255, F1: 0}
            _ = fi46
            _ = v49
        }
        var v50 int = v2
        for fi51 := 0; fi51 < 2; fi51++ {
            for fi52 := 0; fi52 < 3; fi52++ {
                v4 = v4
                _ = fi52
            }
            v53 := GenS6{F0: true, F1: 0}
            _ = fi51
            _ = v53
        }
        _ = v50
    } else {
        switch (v2 - 9223372036854775807) {
        case 9223372036854775807:
            switch (v2 + v2) {
            case 32768:
                var v54 int64 = v42
                _ = v54
            case 42:
                var v55 int32 = 0
                _ = v55
            case -1:
                var v56 float32 = 0.1
                _ = v56
            }
        case 127:
            fmt.Println(v43)
        case 42:
            v2 += v2
        default:
            for fi57 := 0; fi57 < 20; fi57++ {
                if (!(42 != 42)) {
                    v58 := GenS4{F0: (v43 + v43), F1: (v42 / ((0) | 1))}
                    _ = v58
                }
                _ = fi57
            }
        }
    }
    v2 -= -198
    v1[0] = v43
    switch v43 {
    case "foo bar":
        v43 = ("\n\t" + "foo bar")
    }
    var v59 int64 = (v42 % ((42) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v59
    _ = v6
}
