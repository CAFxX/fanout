package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint; F2 int64 }
func genFn0() int32 {
    var v5 float64 = 3.14159
    var v6 int64 = 42
    _ = v5
    _ = v6
    return 0
}
func genFn1() int32 {
    v7 := GenS1{F0: true, F1: 540, F2: -1}
    var v8 int32 = 870
    if true {
        v9 := GenS1{F0: false, F1: 127, F2: -373}
        if false {
            v8 = 970
            fmt.Println(v8)
        } else {
            fmt.Println(v8)
        }
        _ = v9
    } else {
        v8 += (v8 / ((v8) | 1))
        var v10 float32 = (0.0 - -2.25)
        _ = v10
    }
    v8 = 2147483647
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = -1
    var v3 bool = (42 == 128)
    v4 := make(map[int]string)
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 int = fi11
        for fi13 := 0; fi13 < 2; fi13++ {
            v14 := GenS1{F0: v3, F1: (fi11 ^ (^(fi11))), F2: fi11}
            if (!((v3 || true) || (((("\n\t" + ("a" + "a")) + ("a" + "\n\t")) + "") <= ""))) {
                v15 := v4[fi11]
                v16 := v1[(fi11 * v12)]
                _ = v15
                _ = v16
            } else {
                v2 = -860
            }
            switch 9223372036854775807 {
            case 127:
                if v3 {
                    v12 -= -1
                    var v17 uint8 = 42
                    switch (v12 + v12) {
                    case -646:
                        var v18 int64 = 9223372036854775807
                        _ = v18
                    case -2147483648:
                        if v3 {
                            fmt.Println(v17)
                        }
                    }
                    _ = v17
                }
                if v3 {
                    v2 = (v2 ^ v2)
                    if v3 {
                        if v3 {
                            if false {
                                v19 := GenS1{F0: v3, F1: -68, F2: 9223372036854775807}
                                if v3 {
                                    v12 += 127
                                } else {
                                    switch "\n\t" {
                                    case "":
                                        delete(v4, v12)
                                    default:
                                        var v20 uint32 = 1
                                        _ = v20
                                    }
                                    v21 := GenS1{F0: (true || v3), F1: (v12 ^ fi13), F2: fi13}
                                    _ = v21
                                }
                                _ = v19
                            }
                            v12 = fi13
                        } else {
                            v3 = (!v3)
                            if false {
                                var v22 int32 = genFn1()
                                v3 = (v2 <= int64(fi11))
                                _ = v22
                            } else {
                                v3 = (!v3)
                                if (!(v12 != (^(fi13)))) {
                                    var v23 int8 = 100
                                    if (!("\n\t" == "\n\t")) {
                                        v23 = v23
                                        switch "" {
                                        case "a":
                                            switch ("\n\t" + ("" + ("hello" + ""))) {
                                            case "a":
                                                fmt.Println(v3, v23)
                                            default:
                                                if v3 {
                                                    var v24 uint64 = 497
                                                    _ = v24
                                                }
                                            }
                                        case "":
                                            v25 := GenS1{F0: (false && ("hello" > "")), F1: (-(fi13)), F2: fi13}
                                            _ = v25
                                        case "\n\t":
                                            v26 := GenS2{F0: v12, F1: 15}
                                            _ = v26
                                        default:
                                            fmt.Println(v12)
                                        }
                                        v27 := []int{fi13, (fi11 - 2147483647)}
                                        _ = v27
                                    }
                                    v3 = v3
                                    _ = v23
                                }
                                fmt.Println(v12, fi11)
                            }
                        }
                        v2 += -1
                    }
                    delete(v4, 958)
                    v2 += (v2 + v2)
                    switch (-(fi11)) {
                    case -1:
                        v28 := GenS2{F0: (fi13 % (((fi11 * fi11)) | 1)), F1: -10}
                        _ = v28
                    }
                    v2 = 663
                    v4[(fi11 * (v12 | fi13))] = ("foo bar" + ("\n\t" + ("foo bar" + (("a" + ("hello" + "hello")) + "\n\t"))))
                }
                v12 = (fi11 % ((v12) | 1))
            default:
                var v29 bool = ((v12 < (fi13 / ((v12) | 1))) || v3)
                _ = v29
            }
            _ = fi13
            _ = v14
        }
        for fi30 := 0; fi30 < 8; fi30++ {
            v31 := GenS3{F0: 74, F1: 316, F2: (v2 - -1)}
            var v32 float64 = 1.5
            var v33 string = "foo bar"
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v2)
    v3 = (127 > -128)
    v1[2147483647] = (18446744073709551615 <= 65535)
    if (true || (-128 != 100)) {
        var v34 int64 = v2
        _ = v34
    }
    if (824 == 32767) {
        var v35 float32 = ((float32(0.1) + 3.14159) + ((3.14159 + ((1e10 + (-2.25 + -2.25)) * (-2.25 / 1.5))) * 3.14159))
        v3 = v3
        var v36 float64 = float64(-2.25)
        _ = v35
        _ = v36
    }
    v2 += (v2 & (^(v2)))
    v3 = v3
    for fi37 := 0; fi37 < 14; fi37++ {
        if ("foo bar" == ("" + ("a" + (("hello" + "foo bar") + ("hello" + "\n\t"))))) {
            for fi38 := 0; fi38 < 1; fi38++ {
                var v39 uint16 = 457
                _ = fi38
                _ = v39
            }
            v3 = (52 < 14)
            var v40 uint16 = 42
            _ = v40
        }
        _ = fi37
    }
    v41 := GenS2{F0: -113, F1: -128}
    if (v3 && (!(1 != 100))) {
        switch 127 {
        case 32767:
            var v42 uint32 = 4294967295
            _ = v42
        default:
            if (4294967295 <= 780) {
                switch "hello" {
                case "foo bar":
                    if (v3 && v3) {
                        switch 1 {
                        case -486:
                            var v43 bool = ((2147483647 > int32(v2)) && (!(!(433 < 2147483648))))
                            _ = v43
                        default:
                            for fi44 := 0; fi44 < 8; fi44++ {
                                for fi45 := 0; fi45 < 10; fi45++ {
                                    v2 = (-(v2))
                                    var v46 int32 = 439
                                    _ = fi45
                                    _ = v46
                                }
                                v47 := GenS3{F0: 4294967295, F1: 18446744073709551615, F2: -9223372036854775808}
                                v48 := GenS2{F0: (-(fi44)), F1: 42}
                                _ = fi44
                                _ = v47
                                _ = v48
                            }
                        }
                    }
                }
            } else {
                if true {
                    for fi49 := 0; fi49 < 12; fi49++ {
                        v50 := v1[(^(fi49))]
                        v51 := []int{(-(fi49))}
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                    v52 := []int{-9223372036854775808}
                    v52 = append(v52, 32768)
                    _ = v52
                }
                v3 = (!(!v3))
                if true {
                    fmt.Println(v3, v3)
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
}
