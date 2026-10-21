package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int16; F1 int16; F2 uint }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 int32; F1 int32 }
func genFn0(p0 int32) bool {
    fmt.Println(p0)
    var v6 uint32 = 65535
    v7 := GenS1{F0: p0, F1: ("a" + ("hello" + "a"))}
    _ = p0
    _ = v6
    _ = v7
    return false
}

func main() {
    var v1 int16 = 1
    var v2 int8 = 100
    var v3 uint8 = 0
    var v4 uint64 = 4294967295
    var v5 uint32 = 42
    v8 := GenS1{F0: 42, F1: ("" + ("hello" + ("" + "")))}
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint8 = v3
        v11 := GenS2{F0: (v1 - v1), F1: 0, F2: 814}
        fmt.Println(v3, v3, v3)
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS3{F0: int(v5), F1: 9223372036854775807}
    if ((v3 / ((42) | 1)) >= 118) {
        v13 := GenS1{F0: 42, F1: "foo bar"}
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 float32 = 1.5
            switch "foo bar" {
            case "foo bar":
                switch fi14 {
                case 2:
                    if (true || true) {
                        v16 := GenS1{F0: 0, F1: ("\n\t" + "\n\t")}
                        _ = v16
                    } else {
                        if false {
                            var v17 string = ""
                            _ = v17
                        } else {
                            v18 := []int{579}
                            v5 = (v5 << ((543) & 63))
                            _ = v18
                        }
                    }
                default:
                    switch fi14 {
                    case -447:
                        for fi19 := 0; fi19 < 6; fi19++ {
                            if (int32(fi14) == -1) {
                                v4 -= v4
                            }
                            var v20 uint32 = (v5 + v5)
                            fmt.Println(v2, v20, v2)
                            _ = fi19
                            _ = v20
                        }
                    case 127:
                        v21 := []int{(fi14 / ((fi14) | 1))}
                        _ = v21
                    }
                }
            case "hello":
                for fi22 := 0; fi22 < 12; fi22++ {
                    v5 -= 1
                    v1 = 1
                    if ((v3 != v3) && true) {
                        var v23 int8 = (v2 | v2)
                        _ = v23
                    } else {
                        v2 = -52
                        v24 := []int{fi14, (fi14 >> ((472) & 63)), -1, (fi22 << ((2) & 63))}
                        if (v15 <= -2.25) {
                            v4 = (v4 - uint64(v3))
                            if true {
                                v25 := GenS4{F0: int8(v1), F1: v3}
                                fmt.Println(v15)
                                _ = v25
                            } else {
                                if false {
                                    var v26 bool = false
                                    _ = v26
                                }
                                v27 := GenS5{F0: 1, F1: (v1 + (v1 >> ((931) & 63)))}
                                _ = v27
                            }
                        }
                        _ = v24
                    }
                    _ = fi22
                }
            default:
                v28 := GenS6{F0: 490, F1: 255}
                _ = v28
            }
            var v29 float64 = (float64(1.5) - 1e10)
            _ = fi14
            _ = v15
            _ = v29
        }
        switch "" {
        case "foo bar":
            var v30 uint32 = 42
            _ = v30
        }
        _ = v13
    }
    if (!false) {
        v2 += (v2 << ((4294967295) & 63))
        v3 -= (v3 * 0)
        v31 := []int{}
        _ = v31
    } else {
        var v32 uint = 255
        fmt.Println(v5, v3)
        _ = v32
    }
    v3 = (^(v3))
    var v33 uint16 = 61
    switch ("a" + "\n\t") {
    case "hello":
        if false {
            v34 := GenS2{F0: 1, F1: 355, F2: 1}
            v35 := []int{762, -496, -9223372036854775808, 127}
            _ = v34
            _ = v35
        } else {
            if (!false) {
                v36 := GenS5{F0: 1, F1: 0}
                var v37 string = ("\n\t" + (("hello" + ("hello" + "foo bar")) + (("\n\t" + "a") + "\n\t")))
                _ = v36
                _ = v37
            }
            v38 := []int{32767, 32767, -9223372036854775808, 1}
            v38 = append(v38, -2147483648)
            _ = v38
        }
    case "\n\t":
        v4 = v4
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        v40 := []int{(fi39 ^ 128), (-(fi39))}
        v41 := GenS1{F0: 1, F1: "\n\t"}
        var v42 uint = 946
        _ = fi39
        _ = v40
        _ = v41
        _ = v42
    }
    switch 32768 {
    case -2147483648:
        v3 += v3
    }
    v43 := GenS4{F0: 42, F1: (v3 / (((v3 - 42)) | 1))}
    switch -283 {
    case 320:
        v44 := []int{int(v33), 42, 42, int(v33)}
        _ = v44
    }
    switch "" {
    case "a":
        switch 32768 {
        case 32768:
            var v45 float32 = float32(0.1)
            _ = v45
        default:
            v46 := []int{603, 127, 4294967295, 1}
            _ = v46
        }
    case "hello":
        if ((!(1 >= -1)) || (!true)) {
            var v47 uint = 42
            for fi48 := 0; fi48 < 3; fi48++ {
                for fi49 := 0; fi49 < 15; fi49++ {
                    v2 = 100
                    fmt.Println(v5)
                    v3 += v3
                    v33 = v33
                    _ = fi49
                }
                _ = fi48
            }
            _ = v47
        }
    case "\n\t":
        var v50 float32 = (1.5 - (0.1 - 0.0))
        _ = v50
    }
    v5 = v5
    v3 = (v3 - (v3 << ((0) & 63)))
    if false {
        var v51 float32 = (-2.25 - 1e10)
        var v52 uint8 = v3
        _ = v51
        _ = v52
    } else {
        if ((v2 * (v2 << ((731) & 63))) <= (v2 % ((-1) | 1))) {
            var v53 float64 = 1.5
            _ = v53
        } else {
            v4 -= v4
            v54 := []int{-1, 9223372036854775807, int(v3), 1}
            if (v3 > v3) {
                switch "foo bar" {
                case "":
                    if len(v54) >= 2 {
                        v54 = v54[1:len(v54)]
                    }
                default:
                    var v55 float32 = 0.1
                    _ = v55
                }
                var v56 int32 = -1
                _ = v56
            }
            _ = v54
        }
        if (0 >= -1) {
            if (false && true) {
                v4 -= (v4 * v4)
                v57 := []int{0}
                var v58 int32 = -2147483648
                _ = v57
                _ = v58
            }
            if (!(false || (v5 >= v5))) {
                v59 := []int{1, -9223372036854775808, -2147483648}
                _ = v59
            }
        }
        v60 := []int{2147483647, 128, 1, 9223372036854775807}
        _ = v60
    }
    v61 := GenS7{F0: 1, F1: 112}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v43
    _ = v5
    _ = v61
    _ = v8
}
