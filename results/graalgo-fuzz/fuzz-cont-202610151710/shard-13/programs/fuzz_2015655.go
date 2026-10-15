package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint64; F1 int32; F2 string }
type GenS4 struct { F0 int }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(fi6)
        var v7 int64 = 0
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 1}
    v9 := GenS1{F0: -1}
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 uint64 = 599
        _ = fi10
        _ = v11
    }
    _ = v8
    _ = v9
    return 14
}
func genFn1(p0 uint8) int16 {
    var v12 string = "\n\t"
    v12 = v12
    if (false && ((-32768 != -1) || true)) {
        v12 = v12
        v12 = (v12 + ((v12 + v12) + "hello"))
        var v13 int8 = -89
        _ = v13
    } else {
        switch 972 {
        case -1:
            if (!(1 <= 42)) {
                var v14 uint8 = genFn0()
                p0 += (v14 | (p0 ^ 42))
                _ = v14
            } else {
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 uint8 = (p0 * p0)
                    v17 := GenS1{F0: 127}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                if (578 > 1) {
                    var v18 uint64 = 575
                    var v19 uint64 = v18
                    var v20 int8 = 0
                    _ = v18
                    _ = v19
                    _ = v20
                } else {
                    var v21 float64 = 1e10
                    var v22 bool = true
                    _ = v21
                    _ = v22
                }
                v23 := GenS1{F0: 48}
                _ = v23
            }
        default:
            var v24 int = 0
            _ = v24
        }
        fmt.Println(v12)
    }
    _ = p0
    _ = v12
    return 42
}

func main() {
    var v1 int16 = -1
    v2 := make(map[string]string)
    var v3 int8 = -66
    var v4 float32 = 3.14159
    var v5 uint32 = 738
    v3 = (v3 << ((255) & 63))
    switch "foo bar" {
    case "":
        if (false || (1.5 != float64(0.1))) {
            v3 = v3
            fmt.Println(v1, v3)
            for fi25 := 0; fi25 < 6; fi25++ {
                var v26 int = 669
                fmt.Println(v4, fi25, v3)
                _ = fi25
                _ = v26
            }
        }
    }
    v1 = (v1 / (((v1 + v1)) | 1))
    var v27 uint32 = (v5 >> ((255) & 63))
    v28 := v2["\n\t"]
    for fi29 := 0; fi29 < 12; fi29++ {
        if ((!false) || true) {
            switch ((v28 + (v28 + ("\n\t" + v28))) + ((v28 + "") + "hello")) {
            case "a":
                for fi30 := 0; fi30 < 12; fi30++ {
                    v31 := GenS1{F0: (v3 / ((0) | 1))}
                    switch 993 {
                    case 127:
                        v4 += 0.1
                    case 128:
                        if ((v1 | (-(v1))) <= v1) {
                            switch 4294967295 {
                            case 128:
                                v2[v28] = v28
                            case 42:
                                if (-484 <= 904) {
                                    v32 := GenS1{F0: v3}
                                    var v33 uint16 = 65535
                                    v4 -= v4
                                    var v34 bool = false
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                } else {
                                    var v35 string = "foo bar"
                                    v36 := GenS1{F0: (v3 >> ((1) & 63))}
                                    var v37 uint = 506
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                            case 4294967295:
                                v38 := GenS1{F0: v3}
                                _ = v38
                            }
                            var v39 uint64 = 2
                            v5 = v27
                            _ = v39
                        }
                    case 939:
                        switch v28 {
                        case "\n\t":
                            if true {
                                var v40 float64 = float64(((v4 / (-2.25 - ((1.5 / 1.0) * v4))) * 0.0))
                                _ = v40
                            } else {
                                var v41 int16 = (v1 - (v1 / (((v1 ^ (^(v1)))) | 1)))
                                v42 := []int{fi30, fi29}
                                v43 := GenS2{F0: fi30, F1: 42, F2: v3}
                                if (v1 != v41) {
                                    fmt.Println(v27)
                                }
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                            var v44 string = (("hello" + ((v28 + "hello") + ("" + "\n\t"))) + v28)
                            _ = v44
                        case "hello":
                            if (float64(-2.25) < -2.25) {
                                var v45 float32 = ((v4 - (1.5 + v4)) - (float32((1.5 - float64(3.14159))) + float32(0.1)))
                                v3 -= (v3 - v3)
                                v46 := GenS3{F0: 42, F1: 0, F2: v28}
                                _ = v45
                                _ = v46
                            }
                        default:
                            v5 -= 929
                        }
                    }
                    _ = fi30
                    _ = v31
                }
            case "\n\t":
                var v47 int = fi29
                _ = v47
            case "":
                var v48 uint8 = 1
                _ = v48
            default:
                var v49 uint = 255
                _ = v49
            }
            v50 := GenS4{F0: fi29}
            _ = v50
        }
        _ = fi29
    }
    var v51 int32 = -2147483648
    fmt.Println(v51, v3)
    switch ((v28 + v28) + v28) {
    case "":
        fmt.Println(v3, v51, v1)
    case "hello":
        switch 128 {
        case -2147483648:
            for fi52 := 0; fi52 < 4; fi52++ {
                v1 = v1
                v3 -= 100
                v27 -= v5
                _ = fi52
            }
        }
    case "\n\t":
        var v53 bool = ((v1 / (((^(v1))) | 1)) != (v1 + (-(v1))))
        _ = v53
    default:
        v27 = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v51
}
