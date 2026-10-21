package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 string; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint; F1 int16; F2 int32 }
type GenS5 struct { F0 int64; F1 int32 }
func genFn0() bool {
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 uint = 255
            if (((!(!true)) && false) || (0 > -1)) {
                var v9 bool = true
                if (2147483647 == 2147483648) {
                    v10 := []int{fi6, (fi7 + (fi7 << ((v8) & 63))), (fi6 * -165)}
                    if len(v10) > 0 {
                        v10[0] = 4294967295
                    }
                    if (!true) {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        var v11 int32 = -2147483648
                        fmt.Println(v11)
                        _ = v11
                    } else {
                        var v12 int = fi7
                        _ = v12
                    }
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        if len(v10) > 0 {
                            v10[0] = fi7
                        }
                    case "foo bar":
                        v13 := GenS1{F0: v9, F1: ("foo bar" > ("hello" + ""))}
                        _ = v13
                    default:
                        v8 -= v8
                    }
                    var v14 int16 = 347
                    var v15 int64 = -9223372036854775808
                    v15 = int64(v14)
                    var v16 bool = v9
                    if len(v10) > 0 {
                        v10[0] = (fi6 / (((fi6 & (fi7 >> (((v8 & (v8 * v8))) & 63)))) | 1))
                    }
                    var v17 uint32 = 2147483648
                    _ = v10
                    _ = v14
                    _ = v15
                    _ = v16
                    _ = v17
                }
                v18 := GenS2{F0: "\n\t", F1: 4294967295, F2: 9223372036854775807}
                _ = v18
                _ = v9
            } else {
                v19 := []int{fi6}
                switch fi6 {
                case 42:
                    v8 += 4294967295
                case 0:
                    var v20 int64 = -819
                    _ = v20
                default:
                    var v21 int8 = 4
                    _ = v21
                }
                _ = v19
            }
            var v22 uint16 = 979
            if (!false) {
                if ((0 != 127) || (!(true && (!(!false))))) {
                    v23 := GenS1{F0: false, F1: (60 < 2147483648)}
                    var v24 int64 = 9223372036854775807
                    _ = v23
                    _ = v24
                } else {
                    var v25 float64 = 0.1
                    v26 := GenS2{F0: ((("" + "\n\t") + ("" + "")) + "a"), F1: 0, F2: 9223372036854775807}
                    var v27 uint = 2
                    _ = v25
                    _ = v26
                    _ = v27
                }
            }
            _ = fi7
            _ = v22
            _ = v8
        }
        _ = fi6
    }
    var v28 uint32 = 65535
    if (true && true) {
        v29 := []int{}
        _ = v29
    }
    _ = v28
    return (67 <= 4294967295)
}

func main() {
    var v1 bool = ("hello" < "a")
    var v2 string = "a"
    var v3 float32 = 1e10
    var v4 int64 = 0
    v5 := make(map[uint32]uint32)
    switch v2 {
    case "hello":
        v3 = (((v3 + float32(0.1)) * (v3 + float32(1.5))) + float32(1.5))
    case "":
        fmt.Println(v2, v3)
    default:
        var v30 uint32 = 1
        _ = v30
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        if v1 {
            for fi32 := 0; fi32 < 1; fi32++ {
                var v33 float32 = -2.25
                if v1 {
                    var v34 uint32 = 0
                    v35 := []int{(^(fi31)), -1}
                    var v36 int8 = 0
                    _ = v34
                    _ = v35
                    _ = v36
                } else {
                    switch ((("foo bar" + "") + v2) + v2) {
                    case "\n\t":
                        var v37 uint8 = 0
                        _ = v37
                    default:
                        v2 = (("a" + (v2 + (((v2 + "hello") + v2) + ("a" + "a")))) + v2)
                    }
                    fmt.Println(v2)
                    var v38 int8 = 19
                    _ = v38
                }
                _ = fi32
                _ = v33
            }
            for fi39 := 0; fi39 < 20; fi39++ {
                delete(v5, 2)
                v40 := GenS1{F0: v1, F1: v1}
                v41 := v5[42]
                v4 = (-(v4))
                fmt.Println(fi39)
                _ = fi39
                _ = v40
                _ = v41
            }
        }
        var v42 int = -9223372036854775808
        _ = fi31
        _ = v42
    }
    v43 := GenS2{F0: (v2 + v2), F1: 490, F2: (v4 & v4)}
    switch v2 {
    case "a":
        for fi44 := 0; fi44 < 9; fi44++ {
            for fi45 := 0; fi45 < 20; fi45++ {
                v46 := GenS3{F0: 521, F1: 65535}
                _ = fi45
                _ = v46
            }
            _ = fi44
        }
    case "hello":
        if (v4 >= v4) {
            v2 = "hello"
        }
    case "":
        v47 := GenS2{F0: v2, F1: 901, F2: (v4 << ((0) & 63))}
        _ = v47
    default:
        var v48 string = v2
        _ = v48
    }
    v49 := []int{4294967295}
    if false {
        var v50 string = (("foo bar" + (v2 + v2)) + v2)
        v1 = v1
        var v51 int16 = -1
        _ = v50
        _ = v51
    } else {
        fmt.Println(v2, v4)
    }
    var v52 int64 = v4
    v53 := GenS4{F0: 18446744073709551615, F1: 0, F2: 2147483647}
    v54 := GenS5{F0: (v52 >> ((65535) & 63)), F1: 484}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v49
    _ = v5
    _ = v52
    _ = v53
    _ = v54
}
