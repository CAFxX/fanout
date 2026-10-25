package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64; F1 int32; F2 int16 }
func genFn0(p0 uint) uint64 {
    p0 = (p0 % ((p0) | 1))
    p0 = 185
    var v7 uint64 = 42
    _ = p0
    _ = v7
    return v7
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 int = -1
    var v4 uint = 42
    var v5 int64 = 1
    v6 := make(map[string]int)
    v3 = (v3 ^ v3)
    v8 := GenS1{F0: false}
    if false {
        v9 := GenS1{F0: (false && true)}
        v5 -= -1
        v5 -= v5
        _ = v9
    }
    if true {
        var v10 uint8 = 1
        v4 += v4
        _ = v10
    }
    var v11 int16 = -1
    for fi12 := 0; fi12 < 5; fi12++ {
        v5 = (v5 ^ v5)
        if (!(v11 != (^(v11)))) {
            var v13 uint = (v4 % (((v4 * (v4 | v4))) | 1))
            _ = v13
        } else {
            delete(v1, v3)
            if ((v5 * (v5 >> ((65535) & 63))) > v5) {
                for fi14 := 0; fi14 < 20; fi14++ {
                    if false {
                        var v15 int8 = 1
                        _ = v15
                    }
                    _ = fi14
                }
                for fi16 := 0; fi16 < 19; fi16++ {
                    var v17 int16 = (v11 ^ v11)
                    if (883 < 1) {
                        if true {
                            if (1 == 42) {
                                if false {
                                    if true {
                                        switch 128 {
                                        case -759:
                                            v11 = 194
                                        case -2147483648:
                                            var v18 int32 = 0
                                            _ = v18
                                        }
                                        var v19 uint32 = 65535
                                        v2["\n\t"] = (false && false)
                                        _ = v19
                                    } else {
                                        if false {
                                            var v20 bool = true
                                            v5 -= (-(v5))
                                            _ = v20
                                        } else {
                                            _ = len(v6)
                                            switch -425 {
                                            case 42:
                                                v21 := GenS2{F0: 1, F1: v4}
                                                _ = v21
                                            default:
                                                var v22 bool = (!false)
                                                _ = v22
                                            }
                                            switch "hello" {
                                            case "foo bar":
                                                if (v4 <= (v4 << ((430) & 63))) {
                                                    v23 := []int{v3, (^(fi12)), (fi12 ^ fi16)}
                                                    v24 := GenS2{F0: -128, F1: v4}
                                                    if (2147483648 >= 2) {
                                                        fmt.Println(fi12, fi16, v17)
                                                        v17 -= (v11 / (((v11 & 345)) | 1))
                                                        v25 := GenS1{F0: ("" != "\n\t")}
                                                        _ = v25
                                                    } else {
                                                        if true {
                                                            _ = len(v6)
                                                            var v26 int8 = 127
                                                            _ = v26
                                                        }
                                                    }
                                                    _ = v23
                                                    _ = v24
                                                }
                                                v27 := GenS3{F0: 1, F1: 1}
                                                _ = v27
                                            }
                                            v3 = (-(fi16))
                                        }
                                        fmt.Println(v17)
                                    }
                                } else {
                                    v28 := []int{fi12, (v3 | v3), (fi16 + fi16)}
                                    var v29 string = "hello"
                                    var v30 bool = false
                                    v3 = 805
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                                if (false && true) {
                                    var v31 int8 = 100
                                    _ = v31
                                } else {
                                    var v32 uint16 = 42
                                    _ = v32
                                }
                                if true {
                                    v33 := []int{-2147483648}
                                    var v34 uint64 = genFn0(1)
                                    _ = v33
                                    _ = v34
                                } else {
                                    v2[""] = ((v5 - (^(v5))) != -9223372036854775808)
                                    var v35 string = (("foo bar" + (((("\n\t" + "\n\t") + (((("a" + "\n\t") + (("" + "\n\t") + "\n\t")) + "hello") + "a")) + "\n\t") + "foo bar")) + (("foo bar" + "foo bar") + "foo bar"))
                                    _ = v35
                                }
                            }
                        } else {
                            v36 := GenS4{F0: 255, F1: 2}
                            if ((222 == 0) || true) {
                                var v37 int16 = (v17 - (v17 - v11))
                                var v38 uint8 = 255
                                _ = len(v1)
                                _ = v37
                                _ = v38
                            }
                            if false {
                                v1[(fi16 / ((fi16) | 1))] = "\n\t"
                                var v39 int64 = 1
                                _ = v39
                            } else {
                                if (!(2 == 127)) {
                                    v6[""] = fi12
                                }
                            }
                            _ = v36
                        }
                        v17 = v11
                    } else {
                        var v40 float32 = -2.25
                        v3 = (fi16 >> ((v4) & 63))
                        _ = v40
                    }
                    var v41 float64 = (-2.25 / 1.0)
                    _ = fi16
                    _ = v17
                    _ = v41
                }
                if true {
                    if (155 <= 0) {
                        v4 += (^(v4))
                    } else {
                        v42 := GenS2{F0: 0, F1: 530}
                        _ = v42
                    }
                    var v43 uint = 2
                    _ = v43
                } else {
                    for fi44 := 0; fi44 < 14; fi44++ {
                        v45 := GenS5{F0: 4294967295}
                        _ = fi44
                        _ = v45
                    }
                    v6["hello"] = fi12
                }
            } else {
                if (0.0 > float32((3.14159 * 1.5))) {
                    v11 -= v11
                } else {
                    v46 := GenS4{F0: 0, F1: 398}
                    if true {
                        var v47 uint = (v4 & v4)
                        var v48 uint16 = 32767
                        _ = v47
                        _ = v48
                    }
                    v3 += (^(v3))
                    _ = v46
                }
            }
            v49 := GenS4{F0: v4, F1: 923}
            _ = v49
        }
        var v50 float32 = 0.0
        _ = fi12
        _ = v50
    }
    switch 32768 {
    case 128:
        var v51 uint = (v4 - 139)
        _ = v51
    case -530:
        switch v3 {
        case 127:
            for fi52 := 0; fi52 < 5; fi52++ {
                v11 = v11
                var v53 int32 = -2147483648
                v4 = (v4 % (((v4 >> ((18446744073709551615) & 63))) | 1))
                _ = fi52
                _ = v53
            }
        default:
            if false {
                for fi54 := 0; fi54 < 19; fi54++ {
                    v4 = v4
                    v3 = fi54
                    for fi55 := 0; fi55 < 7; fi55++ {
                        var v56 uint16 = 243
                        _ = fi55
                        _ = v56
                    }
                    _ = fi54
                }
                fmt.Println(v5, v3, v4)
            } else {
                var v57 int8 = -50
                v5 = v5
                _ = v57
            }
        }
    }
    if false {
        v58 := GenS2{F0: 16, F1: 2}
        _ = v58
    }
    var v59 string = ("a" + "a")
    v60 := []int{v3, (v3 & 0), 26, v3}
    v61 := GenS6{F0: 0, F1: -1, F2: v11}
    fmt.Println(v59, v59, v59)
    v4 = 10
    v62 := GenS4{F0: (v4 & 2), F1: 820}
    fmt.Println(v11, v3)
    v63 := GenS1{F0: (true || true)}
    for fi64 := 0; fi64 < 19; fi64++ {
        var v65 uint32 = 2
        v65 = (^(v65))
        _ = fi64
        _ = v65
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v60:", len(v60), v60)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v62
    _ = v63
    _ = v8
}
