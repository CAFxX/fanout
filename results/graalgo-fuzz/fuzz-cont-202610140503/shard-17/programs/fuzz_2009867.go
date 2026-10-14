package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int64 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 int64; F1 string; F2 int32 }
func genFn0() int16 {
    for fi6 := 0; fi6 < 11; fi6++ {
        fmt.Println(fi6)
        var v7 uint16 = 65535
        _ = fi6
        _ = v7
    }
    v8 := []int{1, 956, 32767}
    var v9 uint32 = 2
    fmt.Println(v9)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = v8
    _ = v9
    return 42
}
func genFn1() uint32 {
    var v10 float32 = -2.25
    switch "foo bar" {
    case "":
        v11 := GenS1{F0: 2, F1: "hello", F2: 648}
        _ = v11
    }
    fmt.Println(v10)
    v12 := []int{42}
    _ = v10
    _ = v12
    return 1
}

func main() {
    var v1 bool = (true && true)
    var v2 float64 = float64(1e10)
    var v3 float64 = (((1.5 * (-2.25 - (v2 - v2))) + -2.25) + v2)
    var v4 uint32 = 4294967295
    var v5 uint16 = 0
    var v13 float32 = -2.25
    var v14 int = 32767
    v3 -= v2
    if v1 {
        if (v3 <= 1.5) {
            v3 = -2.25
        } else {
            var v15 uint16 = (v5 << ((619) & 63))
            v14 -= v14
            _ = v15
        }
        v16 := GenS1{F0: (v5 + 32767), F1: "", F2: -1}
        var v17 uint16 = (v5 * (v5 / ((65535) | 1)))
        _ = v16
        _ = v17
    }
    v4 -= genFn1()
    v2 = v2
    for fi18 := 0; fi18 < 17; fi18++ {
        v19 := []int{fi18}
        for fi20 := 0; fi20 < 5; fi20++ {
            if (!v1) {
                v21 := GenS1{F0: (v5 + 2), F1: "hello", F2: -615}
                if v1 {
                    var v22 uint16 = (v5 * (v5 % (((v5 | (^(v5)))) | 1)))
                    _ = v22
                } else {
                    v1 = true
                    fmt.Println(v2, v2, v14)
                    fmt.Println(v2, v5, v14)
                    var v23 uint64 = 42
                    _ = v23
                }
                switch (v14 + -867) {
                case -212:
                    if v1 {
                        v24 := GenS1{F0: (^(v5)), F1: ("foo bar" + ("hello" + "a")), F2: 1}
                        v25 := GenS2{F0: v5, F1: "hello"}
                        v19 = append(v19, fi18)
                        fmt.Println(v2)
                        fmt.Println(v3)
                        _ = v24
                        _ = v25
                    }
                case 4294967295:
                    if v1 {
                        if v1 {
                            if (v5 >= v5) {
                                fmt.Println(v3, fi18)
                            } else {
                                v26 := GenS3{F0: 0}
                                fmt.Println(fi20, v1, fi18)
                                fmt.Println(v3, v2, v4)
                                _ = v26
                            }
                            if false {
                                v27 := GenS3{F0: 4294967295}
                                if v1 {
                                    switch fi18 {
                                    case -2147483648:
                                        var v28 bool = false
                                        _ = v28
                                    case 2147483647:
                                        if v1 {
                                            switch "" {
                                            case "foo bar":
                                                if false {
                                                    if ((0 == 42) || (((^(v4)) <= 42) && v1)) {
                                                        if v1 {
                                                            var v29 int32 = 0
                                                            v30 := GenS4{F0: 42}
                                                            _ = v29
                                                            _ = v30
                                                        }
                                                    } else {
                                                        v19 = append(v19, fi20)
                                                        v19 = append(v19, (fi20 & fi20))
                                                    }
                                                    fmt.Println(v3, v14, v2)
                                                } else {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v31 uint = 511
                                                        _ = v31
                                                    case "":
                                                        var v32 string = ("\n\t" + "")
                                                        _ = v32
                                                    case "hello":
                                                        if v1 {
                                                            switch "a" {
                                                            case "a":
                                                                v33 := GenS1{F0: (v5 ^ v5), F1: "foo bar", F2: 6}
                                                                _ = v33
                                                            default:
                                                                v5 = (v5 / (((v5 & 65535)) | 1))
                                                            }
                                                        }
                                                    default:
                                                        fmt.Println(v3, v3, v2)
                                                    }
                                                    _ = len(v19)
                                                }
                                            case "hello":
                                                switch "\n\t" {
                                                case "":
                                                    if ((v3 * ((1e10 - v2) / 3.14159)) <= v2) {
                                                        var v34 int8 = -128
                                                        _ = v34
                                                    }
                                                }
                                            case "":
                                                switch fi20 {
                                                case 0:
                                                    v35 := GenS5{F0: 9223372036854775807}
                                                    _ = v35
                                                }
                                            }
                                            var v36 float32 = 1.5
                                            v13 = 0.1
                                            v19 = append(v19, (fi18 / ((0) | 1)))
                                            if (-1 <= 1) {
                                                fmt.Println(fi18, v3, fi18)
                                            }
                                            _ = v36
                                        }
                                        if (0.0 <= v2) {
                                            v37 := GenS1{F0: 1, F1: ("a" + ("foo bar" + "foo bar")), F2: -1}
                                            var v38 int8 = 127
                                            _ = v37
                                            _ = v38
                                        } else {
                                            switch 32767 {
                                            case 127:
                                                v3 += v3
                                            }
                                            fmt.Println(fi18, v14)
                                        }
                                    case 233:
                                        v39 := GenS6{F0: "", F1: -1}
                                        v2 -= -2.25
                                        if (1 > 25) {
                                            v40 := GenS7{F0: -1, F1: (("" + ("\n\t" + "foo bar")) + ("" + "")), F2: 636}
                                            v19 = append(v19, (v14 * v14))
                                            var v41 float64 = 0.0
                                            _ = v40
                                            _ = v41
                                        } else {
                                            v5 = 126
                                        }
                                        _ = v39
                                    default:
                                        fmt.Println(v3, v5)
                                    }
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                } else {
                                    var v42 int64 = 257
                                    v13 -= 1.5
                                    _ = v42
                                }
                                _ = v27
                            } else {
                                var v43 uint = 0
                                fmt.Println(v5, v43)
                                var v44 int16 = -280
                                _ = v43
                                _ = v44
                            }
                            v14 = (fi18 / ((fi18) | 1))
                        }
                        var v45 int16 = 1000
                        _ = v45
                    }
                case 127:
                    var v46 float32 = v13
                    _ = v46
                }
                _ = v21
            }
            v47 := GenS3{F0: 87}
            _ = fi20
            _ = v47
        }
        for fi48 := 0; fi48 < 5; fi48++ {
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            if (v13 == ((float32(v2) - 1.5) * (v13 - (v13 - v13)))) {
                v49 := GenS7{F0: 229, F1: "", F2: 964}
                _ = v49
            }
            if (((v3 - 1e10) + v2) <= ((3.14159 * v2) / (v3 + 3.14159))) {
                var v50 int = fi48
                _ = v50
            }
            _ = len(v19)
            _ = fi48
        }
        _ = fi18
        _ = v19
    }
    v51 := []int{}
    switch (v14 ^ (v14 | int(v5))) {
    case -824:
        fmt.Println(v5, v13, v4)
    case -1:
        v51 = append(v51, v14)
    default:
        var v52 float64 = (v3 / 3.14159)
        _ = v52
    }
    for fi53 := 0; fi53 < 4; fi53++ {
        for fi54 := 0; fi54 < 10; fi54++ {
            fmt.Println(v1)
            fmt.Println(v13, v5, fi53)
            _ = fi54
        }
        fmt.Println(fi53)
        _ = fi53
    }
    v4 -= 0
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v51:", len(v51), v51)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
}
