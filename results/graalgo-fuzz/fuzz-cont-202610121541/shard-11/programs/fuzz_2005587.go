package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 string; F1 uint8; F2 bool }
func genFn0() uint {
    var v5 string = (("a" + "foo bar") + ("hello" + ("" + ("hello" + ("a" + "a")))))
    v5 = v5
    if false {
        var v6 int8 = 100
        _ = v6
    }
    v5 = ((v5 + ("foo bar" + "\n\t")) + v5)
    _ = v5
    return 42
}
func genFn1(p0 int8) int32 {
    v7 := GenS1{F0: -1, F1: 430, F2: 1}
    var v8 int16 = -32768
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := []int{fi9, (fi9 - -921), (fi9 ^ (fi9 % ((fi9) | 1)))}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    return 349
}

func main() {
    var v1 bool = false
    v2 := []int{}
    var v3 int64 = -5
    var v4 uint8 = 72
    for fi11 := 0; fi11 < 5; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            fmt.Println(v3, fi12)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi12
        }
        for fi13 := 0; fi13 < 12; fi13++ {
            v14 := GenS2{F0: "\n\t"}
            if ((("a" + "foo bar") < "hello") && true) {
                v4 -= (v4 * (v4 + 187))
                switch 195 {
                case -337:
                    var v15 float64 = (float64((1.5 - -2.25)) + 1e10)
                    _ = v15
                case 4294967295:
                    var v16 uint64 = 500
                    _ = v16
                default:
                    if v1 {
                        fmt.Println(fi13)
                        fmt.Println(fi11, v4, fi13)
                    } else {
                        var v17 float32 = 0.0
                        if false {
                            if len(v2) > 0 {
                                v2[0] = -108
                            }
                            v17 += ((v17 - ((3.14159 - (v17 - (v17 * v17))) - v17)) / float32(3.14159))
                            var v18 float32 = v17
                            _ = v18
                        } else {
                            var v19 uint32 = 931
                            v4 -= 127
                            var v20 bool = ((!v1) || true)
                            _ = v19
                            _ = v20
                        }
                        _ = len(v2)
                        _ = v17
                    }
                }
            }
            _ = fi13
            _ = v14
        }
        _ = fi11
    }
    var v21 uint64 = 4294967295
    fmt.Println(v1, v3, v1)
    for fi22 := 0; fi22 < 14; fi22++ {
        for fi23 := 0; fi23 < 15; fi23++ {
            var v24 uint16 = 373
            if len(v2) > 0 {
                v2[0] = (fi23 - (fi22 >> ((80) & 63)))
            }
            _ = fi23
            _ = v24
        }
        switch fi22 {
        case 42:
            var v25 int8 = 0
            _ = v25
        case 41:
            if (100 != 37) {
                v2 = append(v2, -156)
                if (v1 && (v4 >= v4)) {
                    if true {
                        for fi26 := 0; fi26 < 4; fi26++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v27 uint16 = 42
                            _ = len(v2)
                            fmt.Println(v3)
                            _ = fi26
                            _ = v27
                        }
                        if (255 != 1) {
                            v28 := GenS2{F0: ("a" + "a")}
                            _ = v28
                        } else {
                            v3 -= v3
                        }
                        v4 = v4
                    }
                }
                for fi29 := 0; fi29 < 2; fi29++ {
                    if v1 {
                        var v30 int16 = 42
                        fmt.Println(fi29, v21, v30)
                        _ = v30
                    } else {
                        var v31 uint32 = uint32(v4)
                        if false {
                            v31 -= v31
                            switch 2147483647 {
                            case -1:
                                v31 = v31
                            }
                        } else {
                            var v32 float64 = -2.25
                            _ = v32
                        }
                        if (0 >= -1) {
                            if v1 {
                                var v33 int = (fi29 * (^(fi29)))
                                v33 = (^(fi22))
                                v34 := GenS2{F0: "foo bar"}
                                _ = v33
                                _ = v34
                            }
                        }
                        _ = v31
                    }
                    _ = fi29
                }
            }
        }
        _ = fi22
    }
    var v35 uint16 = 2
    for fi36 := 0; fi36 < 2; fi36++ {
        v3 = 42
        for fi37 := 0; fi37 < 19; fi37++ {
            _ = len(v2)
            if (!(v4 > v4)) {
                switch (("foo bar" + "hello") + ("\n\t" + "hello")) {
                case "\n\t":
                    var v38 float32 = 3.14159
                    _ = v38
                case "hello":
                    v3 = (v3 + int64(v21))
                case "a":
                    v39 := GenS2{F0: ("\n\t" + "a")}
                    if len(v2) > 0 {
                        v2[0] = (fi36 & 2147483647)
                    }
                    v4 -= 0
                    _ = v39
                }
                if ((((v4 * 71) != v4) && true) && (!v1)) {
                    v40 := GenS2{F0: "\n\t"}
                    _ = v40
                }
                fmt.Println(v35, v21, v4)
                switch ("\n\t" + (("" + "a") + "")) {
                case "a":
                    switch "hello" {
                    case "":
                        v41 := GenS1{F0: (v3 * 198), F1: 606, F2: 1}
                        _ = v41
                    }
                default:
                    fmt.Println(fi37, v4, v4)
                }
            } else {
                v21 = 902
                var v42 int = 32767
                var v43 uint16 = v35
                _ = v42
                _ = v43
            }
            v1 = (true || v1)
            _ = fi37
        }
        var v44 bool = v1
        _ = fi36
        _ = v44
    }
    var v45 uint64 = (v21 - (v21 << ((2) & 63)))
    for fi46 := 0; fi46 < 5; fi46++ {
        if v1 {
            v45 += v45
            if len(v2) > 0 {
                v2[0] = 32768
            }
            if (v1 || (v1 || v1)) {
                _ = len(v2)
                switch "foo bar" {
                case "foo bar":
                    var v47 bool = (v1 && ((float32(0.1) - (0.1 * (3.14159 + -2.25))) < 1.5))
                    _ = v47
                }
            } else {
                if v1 {
                    var v48 int = (fi46 ^ fi46)
                    fmt.Println(v4, v35)
                    fmt.Println(v1, v3)
                    _ = v48
                } else {
                    v2 = append(v2, (fi46 + (fi46 | -1)))
                }
                v21 = v21
                if (v1 || (638 >= 2147483648)) {
                    v49 := GenS3{F0: 2147483647, F1: "\n\t"}
                    v1 = v1
                    v50 := GenS4{F0: "a", F1: (^(v4)), F2: (!v1)}
                    _ = v49
                    _ = v50
                } else {
                    var v51 uint32 = 151
                    var v52 uint16 = v35
                    _ = v51
                    _ = v52
                }
            }
        }
        if len(v2) > 0 {
            v2[0] = (fi46 >> ((1) & 63))
        }
        _ = fi46
    }
    if (!v1) {
        fmt.Println(v3, v21, v21)
    }
    v21 = v45
    for fi53 := 0; fi53 < 5; fi53++ {
        fmt.Println(v35, v21, v3)
        _ = fi53
    }
    switch 968 {
    case -785:
        var v54 bool = v1
        _ = v54
    default:
        switch 642 {
        case 4294967295:
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
        case 32767:
            _ = len(v2)
        default:
            if false {
                var v55 float32 = (3.14159 + 0.1)
                v3 += 42
                _ = v55
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v35
    _ = v4
    _ = v45
}
