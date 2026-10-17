package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS8 struct { F0 uint; F1 string }
func genFn0(p0 uint32) int64 {
    var v4 float32 = ((3.14159 / 1e10) + 1.5)
    v5 := []int{}
    _ = p0
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 uint) uint {
    v6 := []int{-744, 2147483647, 0}
    v7 := GenS1{F0: "a", F1: 1}
    v8 := GenS1{F0: "hello", F1: (p0 % ((2) | 1))}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return p0
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 16
    var v3 uint16 = 819
    v3 = (v3 - 752)
    v3 = v3
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS1{F0: "\n\t", F1: uint(v2)}
        v11 := GenS2{F0: "a"}
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        if (!(!(false || (((1 == 702) && (!(uint8(v2) == 2))) && (true && true))))) {
            for fi13 := 0; fi13 < 11; fi13++ {
                delete(v1, (fi13 + fi12))
                if ((fi12 >> ((749) & 63)) != (fi13 + -578)) {
                    v3 = (v3 - 255)
                } else {
                    var v14 uint64 = (v2 * v2)
                    if true {
                        var v15 bool = true
                        if (9223372036854775807 <= fi13) {
                            v15 = v15
                        }
                        _ = v15
                    } else {
                        v16 := v1[fi12]
                        var v17 int8 = 104
                        if true {
                            if true {
                                var v18 uint64 = (v14 << ((18446744073709551615) & 63))
                                var v19 uint64 = (^(v14))
                                var v20 float64 = 0.0
                                _ = v18
                                _ = v19
                                _ = v20
                            } else {
                                switch fi13 {
                                case 2147483647:
                                    var v21 float64 = 0.1
                                    _ = v21
                                }
                                v16 += fi12
                                var v22 int8 = -128
                                _ = v22
                            }
                        } else {
                            var v23 uint = 4294967295
                            var v24 uint16 = (v3 % ((v3) | 1))
                            switch 128 {
                            case 42:
                                v25 := GenS1{F0: "a", F1: 18446744073709551615}
                                _ = v25
                            case -887:
                                var v26 int16 = -32768
                                _ = v26
                            default:
                                if ((1.5 + 1e10) != -2.25) {
                                    switch 127 {
                                    case -1:
                                        fmt.Println(fi12)
                                    case 946:
                                        switch 4294967295 {
                                        case 128:
                                            v27 := []int{(fi13 - (fi13 | v16)), (-(fi13)), fi13}
                                            _ = v27
                                        default:
                                            if true {
                                                v28 := GenS3{F0: 127}
                                                fmt.Println(fi13)
                                                var v29 int32 = 1
                                                _ = v28
                                                _ = v29
                                            } else {
                                                var v30 bool = false
                                                fmt.Println(fi13, v24)
                                                var v31 uint = v23
                                                _ = v30
                                                _ = v31
                                            }
                                        }
                                        v32 := []int{fi12}
                                        v33 := GenS1{F0: ("a" + "\n\t"), F1: (^(v23))}
                                        _ = v32
                                        _ = v33
                                    default:
                                        var v34 uint8 = 0
                                        _ = v34
                                    }
                                    v35 := []int{128, 4294967295, 4294967295, -857}
                                    v1[int(v23)] = (fi13 & fi12)
                                    _ = v35
                                } else {
                                    if false {
                                        switch (v16 | (^(fi12))) {
                                        case 32768:
                                            v2 = v2
                                        }
                                    } else {
                                        _ = len(v1)
                                    }
                                    v36 := GenS1{F0: ("hello" + ""), F1: 1}
                                    switch ("foo bar" + (("a" + "hello") + "")) {
                                    case "":
                                        if (!false) {
                                            _ = len(v1)
                                        } else {
                                            var v37 int = (^(v16))
                                            v38 := GenS4{F0: (v3 - (v24 | (^(v3)))), F1: -992, F2: v24}
                                            _ = v37
                                            _ = v38
                                        }
                                    default:
                                        if true {
                                            v39 := GenS5{F0: int64(v24)}
                                            fmt.Println(fi12)
                                            var v40 float64 = -2.25
                                            _ = v39
                                            _ = v40
                                        }
                                        v24 += 836
                                        v41 := GenS1{F0: "hello", F1: (v23 + (v23 % ((v23) | 1)))}
                                        _ = v41
                                    }
                                    _ = v36
                                }
                            }
                            _ = v23
                            _ = v24
                        }
                        _ = v16
                        _ = v17
                    }
                    v42 := GenS6{F0: 0, F1: (v3 | v3)}
                    v1[(fi12 % ((fi12) | 1))] = fi12
                    var v43 uint32 = 2147483648
                    _ = v14
                    _ = v42
                    _ = v43
                }
                var v44 int = fi13
                _ = fi13
                _ = v44
            }
        }
        _ = fi12
    }
    var v45 int32 = -170
    fmt.Println(v3)
    v46 := GenS3{F0: 2}
    for fi47 := 0; fi47 < 17; fi47++ {
        v1[fi47] = (fi47 & -9223372036854775808)
        _ = fi47
    }
    if (42 != 42) {
        for fi48 := 0; fi48 < 4; fi48++ {
            if (!(1 != 65535)) {
                switch ("hello" + (("hello" + ("a" + "a")) + "a")) {
                case "":
                    v49 := GenS6{F0: 255, F1: (v3 * v3)}
                    _ = v49
                default:
                    if (!false) {
                        switch (^(fi48)) {
                        case 42:
                            fmt.Println(v3, v2, v2)
                        }
                    } else {
                        if false {
                            if (-128 <= 1) {
                                fmt.Println(fi48)
                                delete(v1, (fi48 ^ fi48))
                                for fi50 := 0; fi50 < 20; fi50++ {
                                    v51 := GenS1{F0: ("" + "foo bar"), F1: 42}
                                    fmt.Println(v2, v45, v3)
                                    var v52 uint = 4294967295
                                    _ = fi50
                                    _ = v51
                                    _ = v52
                                }
                            } else {
                                var v53 string = ("foo bar" + "hello")
                                _ = v53
                            }
                        }
                    }
                }
                v1[127] = (-(fi48))
                v54 := GenS2{F0: "\n\t"}
                _ = v54
            }
            v55 := []int{fi48}
            _ = fi48
            _ = v55
        }
        v2 += (v2 << ((uint(v3)) & 63))
    } else {
        fmt.Println(v45, v3)
    }
    if ((^(v3)) != 65535) {
        v56 := GenS7{F0: 42, F1: int64(v45), F2: 255}
        v3 += (v3 + 32767)
        v57 := GenS8{F0: 412, F1: "foo bar"}
        _ = v56
        _ = v57
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
}
