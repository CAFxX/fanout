package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16 }
func genFn0(p0 int64, p1 string) int {
    var v7 string = p1
    p1 = v7
    fmt.Println(p1)
    p1 = v7
    _ = p0
    _ = p1
    _ = v7
    return 2147483647
}
func genFn1(p0 int32) int16 {
    v8 := GenS1{F0: 128}
    if (4294967295 != 18446744073709551615) {
        v9 := []int{1, genFn0(0, ((("foo bar" + "\n\t") + ("" + "a")) + "hello")), genFn0(1, ("hello" + "")), -99}
        v10 := GenS2{F0: false, F1: 867}
        _ = v10
        _ = v9
    }
    var v11 int16 = -1
    var v12 uint = 18446744073709551615
    v11 = v11
    _ = p0
    _ = v11
    _ = v12
    _ = v8
    return v11
}

func main() {
    var v1 uint8 = 223
    v2 := make(map[int]uint32)
    var v3 uint8 = v1
    var v4 int32 = -1
    var v5 int8 = 127
    var v6 int8 = int8(v4)
    var v13 string = ("" + (("\n\t" + "a") + ""))
    var v14 int = genFn0(9223372036854775807, "hello")
    fmt.Println(v13, v5)
    fmt.Println(v6, v1, v1)
    var v15 uint64 = 451
    var v16 string = ("\n\t" + v13)
    fmt.Println(v14, v4)
    v15 = v15
    for fi17 := 0; fi17 < 1; fi17++ {
        if (!(0 >= 2)) {
            for fi18 := 0; fi18 < 3; fi18++ {
                if (uint32(v3) != 2147483648) {
                    v19 := GenS2{F0: ((v3 * (v1 & v1)) < v3), F1: 42}
                    var v20 string = v13
                    v21 := v2[(v14 / ((fi17) | 1))]
                    v16 = ("" + ("a" + v20))
                    var v22 uint8 = (v3 | v3)
                    _ = v19
                    _ = v20
                    _ = v21
                    _ = v22
                } else {
                    if true {
                        v5 = (v5 >> ((uint(v1)) & 63))
                        if false {
                            v14 = (fi17 ^ -9223372036854775808)
                        }
                        v4 = v4
                    } else {
                        v23 := GenS3{F0: 255, F1: 65535, F2: v4}
                        delete(v2, (v14 ^ fi17))
                        _ = v23
                    }
                    var v24 uint = 600
                    v15 += v15
                    _ = v24
                }
                switch ((v13 + "") + v16) {
                case "a":
                    if true {
                        if (42 > 2147483647) {
                            if ((18446744073709551615 < 2) && ((!(v15 != (v15 & (^(v15))))) && (float32(-2.25) != 0.1))) {
                                if (genFn0(42, (v13 + (v16 + (("foo bar" + "a") + v13)))) >= fi18) {
                                    switch (fi18 | (fi18 % ((325) | 1))) {
                                    case 127:
                                        v25 := []int{(v14 - 9223372036854775807)}
                                        _ = v25
                                    case 42:
                                        var v26 uint64 = v15
                                        _ = v26
                                    case 9223372036854775807:
                                        v13 = v13
                                    }
                                    if false {
                                        var v27 float64 = float64((0.1 + 0.0))
                                        _ = v27
                                    } else {
                                        v28 := GenS1{F0: 190}
                                        _ = v28
                                    }
                                    if ((fi18 ^ fi18) < (fi17 % ((v14) | 1))) {
                                        var v29 uint16 = 288
                                        var v30 string = ((v16 + (v16 + v13)) + v13)
                                        if (!true) {
                                            var v31 int16 = 42
                                            _ = v31
                                        } else {
                                            v32 := []int{(-(fi17)), genFn0(9223372036854775807, ((v30 + "\n\t") + "")), genFn0(-1, v30), genFn0(-657, v30)}
                                            v33 := GenS4{F0: (^(fi17)), F1: 255}
                                            _ = v32
                                            _ = v33
                                        }
                                        _ = v29
                                        _ = v30
                                    }
                                } else {
                                    v34 := GenS3{F0: 32767, F1: 2147483648, F2: v4}
                                    if (847 >= 42) {
                                        v4 = v4
                                    }
                                    if (!true) {
                                        fmt.Println(v5)
                                    } else {
                                        v16 = (("\n\t" + ("a" + v16)) + v16)
                                        fmt.Println(v5, v6, fi17)
                                    }
                                    _ = v34
                                }
                                v35 := GenS4{F0: fi18, F1: 1}
                                var v36 string = v13
                                var v37 uint32 = 2
                                _ = v35
                                _ = v36
                                _ = v37
                            } else {
                                v3 -= (v3 - (v1 * 127))
                                if (!((-(v4)) < (v4 ^ v4))) {
                                    fmt.Println(v5)
                                    v3 += v3
                                    _ = len(v2)
                                } else {
                                    v14 += (fi18 - (fi17 - (fi17 ^ (^(fi18)))))
                                    fmt.Println(v6, v15, v15)
                                    var v38 uint16 = uint16(v1)
                                    _ = v38
                                }
                            }
                            var v39 bool = (-9223372036854775808 > -9223372036854775808)
                            _ = v39
                        }
                    }
                case "foo bar":
                    if false {
                        v40 := GenS5{F0: v6, F1: (100 >= (v6 % ((-1) | 1)))}
                        if true {
                            v41 := GenS5{F0: -128, F1: false}
                            v1 += v3
                            _ = v41
                        } else {
                            _ = len(v2)
                            v42 := GenS4{F0: 42, F1: uint(v3)}
                            v14 += (v14 ^ (fi17 + (fi17 - (v14 ^ (fi17 * v14)))))
                            _ = v42
                        }
                        v43 := v2[fi17]
                        _ = v40
                        _ = v43
                    }
                }
                var v44 uint16 = 42
                switch genFn0(-9223372036854775808, v16) {
                case 0:
                    if (42 == 42) {
                        fmt.Println(fi18, fi17)
                        v45 := []int{(v14 << ((uint(v1)) & 63)), genFn0(212, "hello")}
                        _ = v45
                    } else {
                        var v46 bool = ((v4 << ((4294967295) & 63)) >= (v4 >> ((42) & 63)))
                        fmt.Println(v13, v1, v1)
                        fmt.Println(fi17)
                        _ = v46
                    }
                case 9223372036854775807:
                    var v47 float64 = float64(1.5)
                    _ = v47
                case -9223372036854775808:
                    v13 = v13
                    v48 := GenS5{F0: v6, F1: false}
                    _ = v48
                default:
                    v49 := GenS1{F0: (v3 | (v3 + (v1 - (v3 * (v3 - (v1 * v1))))))}
                    _ = v49
                }
                v50 := []int{(fi17 - (fi17 - (^(fi17)))), (fi18 << ((18446744073709551615) & 63)), 128}
                v15 -= 1
                if false {
                    var v51 uint = 228
                    _ = v51
                }
                v52 := GenS3{F0: 0, F1: 42, F2: (v4 + -269)}
                var v53 uint = 1
                v54 := GenS5{F0: v6, F1: (v16 >= (v13 + "hello"))}
                var v55 uint8 = v1
                _ = fi18
                _ = v44
                _ = v50
                _ = v52
                _ = v53
                _ = v54
                _ = v55
            }
            v15 = (v15 - v15)
        } else {
            var v56 float32 = (1e10 * 1e10)
            _ = v56
        }
        _ = fi17
    }
    var v57 bool = false
    if (((1.5 + 1e10) / ((0.1 - ((1.5 + 1e10) - -2.25)) * -2.25)) >= 0.0) {
        switch ((((("hello" + v16) + (v16 + (v16 + "a"))) + "hello") + ("foo bar" + v13)) + v13) {
        case "hello":
            if (false || (false && v57)) {
                var v58 int64 = 78
                _ = v58
            }
        default:
            v5 += int8(v4)
        }
        if (42 == -34) {
            v59 := GenS6{F0: v13}
            v60 := GenS7{F0: -32768}
            _ = v59
            _ = v60
        }
        var v61 string = v13
        _ = v61
    } else {
        v3 = v1
        var v62 bool = v57
        _ = v62
    }
    for fi63 := 0; fi63 < 15; fi63++ {
        delete(v2, (fi63 % (((v14 << ((uint(v15)) & 63))) | 1)))
        _ = fi63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
}
