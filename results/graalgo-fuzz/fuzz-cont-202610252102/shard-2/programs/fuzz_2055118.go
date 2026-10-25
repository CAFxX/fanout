package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int64; F1 int64; F2 bool }
type GenS3 struct { F0 int32; F1 uint32; F2 int }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 bool; F1 int16; F2 int8 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 string, p1 bool, p2 int16) int {
    if p1 {
        var v6 string = p0
        v6 = ""
        var v7 int32 = -317
        _ = v6
        _ = v7
    }
    p0 = ((p0 + p0) + "\n\t")
    v8 := GenS1{F0: 724, F1: 42}
    var v9 bool = (false && p1)
    v10 := GenS1{F0: -9223372036854775808, F1: (p2 / ((p2) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 244
}
func genFn1() string {
    switch 9223372036854775807 {
    case 1:
        if false {
            v11 := GenS1{F0: 0, F1: 32767}
            if ((!true) || (true || false)) {
                v12 := GenS2{F0: -1, F1: -1, F2: true}
                var v13 uint64 = 4294967295
                _ = v12
                _ = v13
            } else {
                if (1 == 588) {
                    v14 := []int{-234}
                    var v15 float32 = float32(-2.25)
                    var v16 string = "\n\t"
                    _ = v14
                    _ = v15
                    _ = v16
                } else {
                    if (1.5 < (0.1 - 0.1)) {
                        if false {
                            var v17 uint32 = 1
                            v18 := []int{-2147483648, genFn0("a", false, 0), 128}
                            _ = v17
                            _ = v18
                        }
                    } else {
                        var v19 int = -529
                        v20 := []int{1, genFn0("a", true, 105)}
                        v21 := GenS2{F0: 1, F1: -858, F2: (true && ((int64(v19) >= -1) || false))}
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    v22 := []int{}
                    _ = v22
                }
                var v23 float32 = 3.14159
                if true {
                    for fi24 := 0; fi24 < 20; fi24++ {
                        v25 := []int{fi24}
                        v25 = append(v25, (fi24 - fi24))
                        _ = fi24
                        _ = v25
                    }
                    fmt.Println(v23)
                } else {
                    for fi26 := 0; fi26 < 9; fi26++ {
                        v27 := []int{(-(fi26)), (fi26 / (((-(fi26))) | 1))}
                        if ((((0.1 - (float64((v23 - v23)) + 0.0)) * 3.14159) <= float64((v23 / v23))) && true) {
                            if (1 < 133) {
                                var v28 float64 = (1.5 / 3.14159)
                                v29 := GenS3{F0: 1, F1: 2, F2: fi26}
                                for fi30 := 0; fi30 < 1; fi30++ {
                                    fmt.Println(v28, v23, fi26)
                                    _ = fi30
                                }
                                _ = v28
                                _ = v29
                            }
                            for fi31 := 0; fi31 < 10; fi31++ {
                                v23 -= ((v23 - 0.0) / 1.5)
                                var v32 uint8 = 1
                                _ = fi31
                                _ = v32
                            }
                            if len(v27) > 0 {
                                v27[0] = fi26
                            }
                        }
                        if false {
                            if len(v27) >= 2 {
                                v27 = v27[1:len(v27)]
                            }
                            var v33 string = "a"
                            _ = v33
                        }
                        _ = fi26
                        _ = v27
                    }
                    v34 := []int{32767, -2147483648, 1}
                    if false {
                        v23 -= (v23 / v23)
                        if len(v34) >= 2 {
                            v34 = v34[1:len(v34)]
                        }
                    }
                    _ = v34
                }
                _ = v23
            }
            _ = v11
        } else {
            v35 := []int{42, 61, -1}
            _ = v35
        }
    }
    v36 := GenS4{F0: false, F1: -241}
    _ = v36
    return "foo bar"
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    var v4 float64 = -2.25
    var v5 int32 = -1
    v37 := GenS5{F0: (70 < 433), F1: 32767, F2: -1}
    v38 := []int{-2147483648, -846, 9223372036854775807, 32767}
    fmt.Println(v5, v5)
    fmt.Println(v5, v4)
    var v39 int8 = 62
    if ((false || (v4 < (v4 * v4))) || (!true)) {
        var v40 int32 = v5
        v41 := v2[4294967295]
        _ = v40
        _ = v41
    } else {
        var v42 float64 = (v4 - 1.5)
        switch -9223372036854775808 {
        case 127:
            var v43 int8 = -1
            _ = v43
        case 843:
            for fi44 := 0; fi44 < 18; fi44++ {
                if false {
                    switch ((("hello" + "a") + ("a" + "foo bar")) + (("a" + "\n\t") + "\n\t")) {
                    case "hello":
                        if (false || ((((0 != 2) && (744 >= 255)) || true) || false)) {
                            v45 := GenS3{F0: v5, F1: 898, F2: (^(fi44))}
                            _ = v45
                        }
                    }
                }
                delete(v3, 0)
                _ = fi44
            }
        default:
            switch "foo bar" {
            case "a":
                v46 := GenS3{F0: (v5 & v5), F1: 2147483647, F2: 0}
                _ = v46
            case "":
                v1 = v1
            default:
                delete(v3, 2)
            }
        }
        _ = v42
    }
    delete(v3, 65535)
    if false {
        delete(v3, 42)
    }
    v47 := GenS6{F0: 127}
    var v48 int64 = v1
    if (1 < 42) {
        if true {
            for fi49 := 0; fi49 < 11; fi49++ {
                var v50 float64 = 0.1
                var v51 int64 = v1
                v4 = 1.5
                _ = fi49
                _ = v50
                _ = v51
            }
            v52 := GenS1{F0: 127, F1: -311}
            _ = v52
        }
        var v53 int8 = v39
        _ = v53
    }
    fmt.Println(v39, v39, v39)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v47
    _ = v48
    _ = v5
}
