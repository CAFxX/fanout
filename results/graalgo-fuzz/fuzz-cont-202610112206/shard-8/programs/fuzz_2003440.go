package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint16; F1 int32; F2 uint16 }
func genFn0(p0 uint8, p1 uint32) int32 {
    var v7 uint32 = (^(p1))
    var v8 int32 = -697
    v8 += v8
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 595
}
func genFn1(p0 uint64, p1 uint16) int {
    p1 += 2
    v9 := []int{-419, 127, -573, -377}
    v9 = append(v9, 42)
    _ = p0
    _ = p1
    _ = v9
    return -198
}

func main() {
    var v1 uint64 = 1
    var v2 int32 = 0
    v3 := make(map[string]bool)
    var v4 bool = true
    var v5 uint64 = (v1 | 511)
    var v6 bool = v4
    var v10 bool = ((-32768 >= -641) && (1 != 42))
    var v11 int8 = 0
    if ((false || ((1e10 - float32(1.5)) != (0.1 + 3.14159))) && (!v6)) {
        v12 := v3["foo bar"]
        v13 := GenS1{F0: false, F1: "a"}
        _ = v12
        _ = v13
    }
    if (false || (540 < 807)) {
        v14 := GenS1{F0: (!((v11 | v11) < (v11 | (-(v11))))), F1: "a"}
        _ = v14
    } else {
        if ((v2 % ((v2) | 1)) > -2147483648) {
            if ((v1 + 18446744073709551615) != v5) {
                switch 4294967295 {
                case -31:
                    switch -122 {
                    case -1:
                        if ((!((1 == 605) && (!false))) || true) {
                            if v4 {
                                v2 = v2
                                _ = len(v3)
                            }
                            if ("foo bar" >= "\n\t") {
                                v15 := []int{511, 0, genFn1((v1 - (v5 % (((v1 - v1)) | 1))), 1), 2147483647}
                                for fi16 := 0; fi16 < 1; fi16++ {
                                    for fi17 := 0; fi17 < 16; fi17++ {
                                        var v18 float32 = ((-2.25 + ((1.5 + 1e10) - 0.0)) * float32(0.1))
                                        if (!false) {
                                            v10 = (2147483648 > 42)
                                        }
                                        _ = fi17
                                        _ = v18
                                    }
                                    if v6 {
                                        var v19 string = ("" + "")
                                        if (42 > int16(fi16)) {
                                            var v20 int32 = genFn0(86, 2)
                                            for fi21 := 0; fi21 < 14; fi21++ {
                                                var v22 uint = 210
                                                _ = fi21
                                                _ = v22
                                            }
                                            v23 := GenS1{F0: v10, F1: (v19 + "hello")}
                                            _ = v20
                                            _ = v23
                                        }
                                        switch fi16 {
                                        case 1:
                                            v24 := GenS2{F0: 42, F1: 1}
                                            _ = v24
                                        case -1:
                                            fmt.Println(v4)
                                        }
                                        _ = v19
                                    }
                                    fmt.Println(v2, v2)
                                    _ = fi16
                                }
                                _ = v15
                            }
                            if v4 {
                                v4 = (2 != 0)
                                _ = len(v3)
                            }
                        }
                    case 32768:
                        switch 1 {
                        case 32768:
                            if (v6 && false) {
                                for fi25 := 0; fi25 < 8; fi25++ {
                                    v1 = v1
                                    _ = fi25
                                }
                                var v26 int32 = (v2 * v2)
                                _ = v26
                            }
                        default:
                            fmt.Println(v1)
                        }
                    case 2147483647:
                        _ = len(v3)
                    }
                case -924:
                    var v27 int64 = -9223372036854775808
                    _ = v27
                }
            } else {
                var v28 uint32 = 754
                _ = v28
            }
            v10 = true
            var v29 int32 = (v2 ^ (v2 << ((2) & 63)))
            _ = v29
        } else {
            for fi30 := 0; fi30 < 19; fi30++ {
                v31 := []int{}
                var v32 string = ""
                var v33 int32 = (v2 ^ genFn0(127, 2147483647))
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
            v34 := v3[""]
            fmt.Println(v34)
            _ = v34
        }
    }
    var v35 float64 = float64((0.1 * -2.25))
    v3[("foo bar" + ("" + (((("foo bar" + "hello") + ("foo bar" + "foo bar")) + "hello") + "")))] = v6
    _ = len(v3)
    var v36 int8 = v11
    v4 = true
    if false {
        if (255 > 34) {
            for fi37 := 0; fi37 < 15; fi37++ {
                v38 := []int{fi37}
                v4 = v10
                _ = fi37
                _ = v38
            }
        } else {
            v39 := GenS3{F0: -240}
            for fi40 := 0; fi40 < 18; fi40++ {
                var v41 uint32 = 4294967295
                _ = len(v3)
                _ = fi40
                _ = v41
            }
            _ = v39
        }
        v42 := GenS4{F0: v1, F1: 127}
        v43 := GenS2{F0: 4294967295, F1: 255}
        _ = v42
        _ = v43
    }
    var v44 int = 128
    v5 = (v5 << ((428) & 63))
    v3["hello"] = (((3.14159 + 1e10) - ((float32(float64((1.5 * -2.25))) / (-2.25 - 1.5)) * 0.0)) <= (0.0 / (3.14159 * 1e10)))
    if true {
        var v45 float64 = v35
        v10 = (!false)
        v46 := GenS5{F0: 65535, F1: 2, F2: v1}
        _ = v45
        _ = v46
    } else {
        var v47 float64 = v35
        for fi48 := 0; fi48 < 19; fi48++ {
            if v6 {
                if (1 < 934) {
                    switch "" {
                    case "a":
                        v49 := []int{}
                        _ = v49
                    case "hello":
                        if true {
                            v6 = (857 > 32767)
                            v35 += 1.5
                            fmt.Println(fi48)
                        }
                    case "\n\t":
                        v50 := GenS6{F0: 643, F1: (v2 * v2), F2: 0}
                        _ = v50
                    }
                } else {
                    var v51 int8 = (v11 % ((-128) | 1))
                    _ = v51
                }
            }
            var v52 string = ("a" + "hello")
            _ = fi48
            _ = v52
        }
        _ = v47
    }
    v53 := []int{v44, (v44 % ((int(v11)) | 1)), (^(v44)), (v44 + v44)}
    var v54 uint16 = uint16(v5)
    if v10 {
        v55 := GenS6{F0: 634, F1: (v2 * (v2 % (((^(v2))) | 1))), F2: (v54 * 65535)}
        switch "hello" {
        case "\n\t":
            if (!((^(v54)) == (v54 * 255))) {
                for fi56 := 0; fi56 < 3; fi56++ {
                    v53 = append(v53, -9223372036854775808)
                    _ = fi56
                }
                var v57 int8 = v36
                var v58 uint64 = (v1 << ((4294967295) & 63))
                _ = v57
                _ = v58
            } else {
                v59 := GenS5{F0: 18446744073709551615, F1: v54, F2: v1}
                fmt.Println(v11)
                v60 := GenS2{F0: 18446744073709551615, F1: 4294967295}
                _ = v59
                _ = v60
            }
        default:
            var v61 int16 = 948
            _ = v61
        }
        _ = v55
    }
    if len(v53) >= 2 {
        v53 = v53[1:len(v53)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v44
    _ = v5
    _ = v53
    _ = v54
    _ = v6
}
