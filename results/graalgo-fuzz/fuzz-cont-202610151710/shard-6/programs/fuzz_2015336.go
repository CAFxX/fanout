package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint; F1 bool; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 uint64 }
func genFn0() int16 {
    v7 := GenS1{F0: -1, F1: true, F2: 2147483648}
    var v8 int64 = 960
    v9 := []int{569, 0}
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 int8 = -128
    var v4 uint32 = 65535
    v5 := make(map[int]string)
    var v6 int8 = (v3 >> ((255) & 63))
    var v10 int32 = 201
    for fi11 := 0; fi11 < 9; fi11++ {
        v3 -= (^(v3))
        for fi12 := 0; fi12 < 20; fi12++ {
            v3 = v3
            fmt.Println(v3, v10, fi11)
            _ = fi12
        }
        _ = fi11
    }
    v3 += v3
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS1{F0: 9223372036854775807, F1: true, F2: v4}
        v15 := v2["\n\t"]
        if (false || (v15 >= "hello")) {
            var v16 uint64 = 523
            v17 := GenS1{F0: 1, F1: false, F2: (v4 << ((18446744073709551615) & 63))}
            var v18 uint = 1
            _ = v16
            _ = v17
            _ = v18
        } else {
            if false {
                fmt.Println(v3)
                v5[(^(fi13))] = (v15 + v15)
                v19 := GenS2{F0: 4294967295, F1: false, F2: 1}
                _ = v19
            } else {
                var v20 uint64 = 4294967295
                for fi21 := 0; fi21 < 1; fi21++ {
                    v22 := []int{(fi21 * 128), (-(fi13))}
                    if (!(v4 >= v4)) {
                        if true {
                            if ((fi13 & fi13) < fi21) {
                                if (true || (((-2.25 / 1.0) * (1e10 * (0.0 + float64((1.5 / 1e10))))) > 3.14159)) {
                                    fmt.Println(v10, fi21, fi13)
                                    v4 -= 4294967295
                                    v15 = v15
                                    v3 = 0
                                } else {
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                    var v23 int16 = 398
                                    var v24 int64 = 0
                                    _ = v23
                                    _ = v24
                                }
                                if (!true) {
                                    var v25 uint32 = (v4 + (^(v4)))
                                    switch v15 {
                                    case "":
                                        v6 = (v3 & (^(v3)))
                                    case "hello":
                                        var v26 uint32 = 7
                                        _ = v26
                                    default:
                                        var v27 int = (fi13 << ((511) & 63))
                                        _ = v27
                                    }
                                    if (false && true) {
                                        v15 = v15
                                        var v28 int64 = 9223372036854775807
                                        v25 = v4
                                        if true {
                                            var v29 float32 = 1.5
                                            v4 = (v25 >> ((4294967295) & 63))
                                            _ = v29
                                        }
                                        var v30 int64 = v28
                                        _ = v28
                                        _ = v30
                                    }
                                    _ = v25
                                }
                                v3 += (v6 << ((207) & 63))
                            }
                            v31 := GenS3{F0: -263}
                            _ = v31
                        }
                        v32 := v5[fi21]
                        var v33 uint32 = 149
                        _ = v32
                        _ = v33
                    }
                    var v34 uint = 255
                    if false {
                        var v35 float64 = float64(1e10)
                        fmt.Println(fi21, v20)
                        fmt.Println(v4)
                        var v36 float64 = (3.14159 - 1.5)
                        v37 := GenS4{F0: v6}
                        var v38 string = ""
                        _ = v35
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    fmt.Println(v3, v20)
                    _ = fi21
                    _ = v22
                    _ = v34
                }
                for fi39 := 0; fi39 < 12; fi39++ {
                    var v40 bool = (fi13 != -306)
                    var v41 float32 = -2.25
                    var v42 uint16 = 545
                    _ = fi39
                    _ = v40
                    _ = v41
                    _ = v42
                }
                _ = v20
            }
            var v43 uint64 = 2
            _ = v43
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi44 := 0; fi44 < 11; fi44++ {
        v45 := GenS5{F0: 0, F1: 42}
        v3 += (-(v3))
        _ = fi44
        _ = v45
    }
    for fi46 := 0; fi46 < 8; fi46++ {
        switch "a" {
        case "hello":
            switch 0 {
            case -928:
                var v47 uint32 = (^(v4))
                _ = v47
            case 2147483647:
                fmt.Println(v10)
            default:
                fmt.Println(v6, v4)
            }
        default:
            var v48 float64 = 3.14159
            _ = v48
        }
        v49 := GenS2{F0: 255, F1: (!(1 != 127)), F2: 760}
        _ = fi46
        _ = v49
    }
    var v50 int8 = v3
    var v51 uint64 = 572
    var v52 uint64 = (v51 ^ v51)
    v10 += (v10 + (v10 & v10))
    v50 = 100
    v52 += (v51 & 848)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
}
