package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint32; F1 int8; F2 int }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8; F1 uint32; F2 string }
type GenS7 struct { F0 int16; F1 uint64 }
type GenS8 struct { F0 uint32; F1 string }
func genFn0(p0 int64) uint64 {
    var v5 uint32 = 2
    p0 += p0
    p0 = p0
    if (false || true) {
        switch "foo bar" {
        case "hello":
            var v6 int8 = 72
            _ = v6
        case "a":
            v7 := []int{-1, -540, 32768, -2147483648}
            _ = v7
        case "":
            v8 := GenS1{F0: ((p0 >= p0) || false), F1: 1}
            _ = v8
        default:
            var v9 int64 = p0
            _ = v9
        }
        v10 := []int{}
        _ = v10
    } else {
        p0 = (p0 & 1)
    }
    _ = p0
    _ = v5
    return 18446744073709551615
}

func main() {
    var v1 int16 = 32767
    var v2 uint32 = 4294967295
    v3 := make(map[int]uint32)
    var v4 int32 = 2147483647
    v4 -= (v4 ^ (v4 << ((65535) & 63)))
    v1 = 0
    fmt.Println(v4, v1, v1)
    switch -9223372036854775808 {
    case 9223372036854775807:
        v4 = v4
    }
    var v11 int64 = 1
    for fi12 := 0; fi12 < 9; fi12++ {
        switch ("a" + (("a" + "") + "hello")) {
        case "":
            for fi13 := 0; fi13 < 17; fi13++ {
                v4 -= v4
                _ = fi13
            }
        }
        var v14 float32 = 1.5
        v15 := []int{}
        _ = fi12
        _ = v14
        _ = v15
    }
    v16 := v3[-1]
    for fi17 := 0; fi17 < 8; fi17++ {
        for fi18 := 0; fi18 < 13; fi18++ {
            if false {
                v19 := []int{(fi17 >> ((0) & 63)), 2147483647, (fi17 & (fi18 - fi17)), int(v1)}
                if (-128 != 42) {
                    var v20 float64 = 0.1
                    if (true || (1 != 65535)) {
                        if false {
                            v2 += (v2 << ((797) & 63))
                            v21 := GenS2{F0: 73, F1: v4}
                            _ = v21
                        } else {
                            if true {
                                switch (fi18 ^ 32767) {
                                case 843:
                                    fmt.Println(v2)
                                default:
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v22 int32 = v4
                                        _ = v22
                                    case "":
                                        if true {
                                            v23 := GenS3{F0: (v16 + 2147483647), F1: -30, F2: (-(fi17))}
                                            _ = v23
                                        }
                                    case "a":
                                        _ = len(v19)
                                    }
                                }
                            } else {
                                v24 := GenS3{F0: v16, F1: 1, F2: fi17}
                                if (((^(v16)) >= (v2 ^ (^(v2)))) && (498 != (v1 >> ((65535) & 63)))) {
                                    var v25 int16 = (-(v1))
                                    _ = v25
                                } else {
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                    switch fi18 {
                                    case 32768:
                                        v26 := GenS1{F0: true, F1: 255}
                                        _ = v26
                                    case 42:
                                        if false {
                                            v16 = 2147483647
                                            v27 := GenS2{F0: 1, F1: v4}
                                            v4 = v4
                                            v28 := v3[-1]
                                            _ = v27
                                            _ = v28
                                        } else {
                                            v20 -= (((v20 + v20) * v20) - v20)
                                            v16 += (v16 & v16)
                                            var v29 uint64 = 1
                                            _ = v29
                                        }
                                    case 626:
                                        var v30 bool = false
                                        _ = v30
                                    }
                                    fmt.Println(v16, v20, v16)
                                }
                                _ = v24
                            }
                            if false {
                                v20 = (float64(1e10) - v20)
                            }
                        }
                        v19 = append(v19, 0)
                    } else {
                        var v31 float64 = v20
                        _ = v31
                    }
                    v16 += v2
                    _ = v20
                }
                v16 = 964
                _ = v19
            } else {
                delete(v3, (-(fi17)))
                v32 := GenS4{F0: 750, F1: fi17}
                _ = v32
            }
            v33 := []int{9223372036854775807, -9223372036854775808}
            v3[fi17] = v16
            if false {
                v4 -= v4
            } else {
                v1 += 32767
            }
            if false {
                v4 -= 42
                fmt.Println(v16)
                v34 := GenS5{F0: false}
                _ = v34
            } else {
                if true {
                    v35 := GenS5{F0: (!(v1 <= -1))}
                    v36 := GenS5{F0: (65535 == 0)}
                    _ = v35
                    _ = v36
                } else {
                    switch int(v16) {
                    case -348:
                        fmt.Println(v1, fi18)
                    case 76:
                        v37 := GenS6{F0: -41, F1: v2, F2: ("\n\t" + ("hello" + ""))}
                        _ = v37
                    case 0:
                        delete(v3, -2147483648)
                    default:
                        _ = len(v3)
                    }
                    v2 = v16
                }
            }
            if (int8(v11) < -128) {
                if (18446744073709551615 <= 0) {
                    fmt.Println(fi18, v1, v4)
                    if false {
                        if false {
                            if (fi18 < (fi18 >> ((0) & 63))) {
                                fmt.Println(v4, v11)
                            }
                        }
                        v11 += v11
                        v2 = v2
                    }
                } else {
                    v4 -= (^(v4))
                    var v38 int = 946
                    v39 := v3[int(v1)]
                    if ((2147483647 < 0) || true) {
                        var v40 uint = 18446744073709551615
                        v41 := GenS1{F0: (4294967295 == (v40 << (((v40 - 862)) & 63))), F1: 913}
                        var v42 uint32 = 2147483647
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                    _ = v38
                    _ = v39
                }
                if false {
                    v3[-2147483648] = (^(v16))
                } else {
                    if len(v33) > 0 {
                        v33[0] = fi18
                    }
                    var v43 float64 = 1.5
                    _ = v43
                }
            } else {
                if false {
                    var v44 uint64 = 18446744073709551615
                    _ = v44
                }
                v45 := GenS7{F0: v1, F1: 2}
                v11 = (v11 % (((v11 >> ((42) & 63))) | 1))
                _ = v45
            }
            _ = fi18
            _ = v33
        }
        _ = fi17
    }
    if (float32((3.14159 * -2.25)) < -2.25) {
        var v46 float32 = 1.5
        if (899 < v11) {
            var v47 float64 = float64(float32((float64(v46) - (0.0 / (-2.25 - 3.14159)))))
            v16 = (v16 % (((v16 / ((v2) | 1))) | 1))
            _ = v47
        }
        _ = v46
    } else {
        if true {
            fmt.Println(v11, v11, v16)
            var v48 int16 = 376
            _ = v48
        }
        v16 += (v2 ^ 0)
        for fi49 := 0; fi49 < 1; fi49++ {
            for fi50 := 0; fi50 < 18; fi50++ {
                var v51 string = ("foo bar" + ("" + "\n\t"))
                switch (v51 + v51) {
                case "\n\t":
                    v52 := GenS3{F0: (v2 / (((v16 / ((455) | 1))) | 1)), F1: 88, F2: (fi49 * fi49)}
                    _ = v52
                }
                var v53 int = (fi50 - (-(fi50)))
                var v54 uint64 = 4294967295
                var v55 int = (v53 * fi50)
                v56 := []int{(fi50 ^ (fi50 << ((474) & 63))), (fi49 >> ((4294967295) & 63)), (v55 % ((-510) | 1))}
                if true {
                    v57 := v3[fi50]
                    _ = len(v3)
                    _ = v57
                }
                var v58 int8 = 1
                _ = len(v56)
                _ = fi50
                _ = v51
                _ = v53
                _ = v54
                _ = v55
                _ = v56
                _ = v58
            }
            _ = fi49
        }
    }
    for fi59 := 0; fi59 < 5; fi59++ {
        if true {
            v1 -= (v1 * v1)
        }
        _ = fi59
    }
    var v60 uint64 = 498
    v61 := []int{0, 549, 1}
    for fi62 := 0; fi62 < 13; fi62++ {
        var v63 bool = false
        v64 := GenS8{F0: v2, F1: ((("hello" + "a") + "hello") + "foo bar")}
        switch "\n\t" {
        case "foo bar":
            v63 = (v63 || (v4 <= (-(v4))))
        case "\n\t":
            v60 += v60
        case "":
            if v63 {
                var v65 int16 = v1
                v65 -= v1
                v66 := v3[32768]
                _ = v65
                _ = v66
            } else {
                var v67 int8 = 0
                for fi68 := 0; fi68 < 3; fi68++ {
                    var v69 int8 = v67
                    fmt.Println(v60, v63)
                    if len(v61) >= 2 {
                        v61 = v61[1:len(v61)]
                    }
                    _ = fi68
                    _ = v69
                }
                v11 -= 300
                _ = v67
            }
        }
        _ = fi62
        _ = v63
        _ = v64
    }
    var v70 uint32 = uint32(v60)
    _ = len(v61)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v61:", len(v61), v61)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v60
    _ = v61
    _ = v70
}
