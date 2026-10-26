package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 string; F1 uint64 }
type GenS9 struct { F0 int16; F1 uint8; F2 int64 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint16; F1 int8; F2 uint32 }
func genFn0() uint64 {
    for fi6 := 0; fi6 < 9; fi6++ {
        switch fi6 {
        case 0:
            var v7 uint64 = 0
            _ = v7
        case -607:
            var v8 string = ""
            _ = v8
        }
        var v9 int8 = -47
        _ = fi6
        _ = v9
    }
    var v10 int8 = 100
    v11 := []int{-209, 128, -1}
    v11 = append(v11, 9223372036854775807)
    _ = v10
    _ = v11
    return 2
}
func genFn1(p0 int8) int32 {
    fmt.Println(p0)
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 int16 = -1
        if true {
            v14 := []int{2147483647, fi12, fi12}
            _ = v14
        }
        v15 := GenS1{F0: 32767, F1: v13}
        _ = fi12
        _ = v13
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        fmt.Println(fi16, p0)
        var v17 int32 = int32(p0)
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 uint8 = 128
        _ = fi18
        _ = v19
    }
    var v20 uint64 = 290
    _ = p0
    _ = v20
    return 2147483647
}

func main() {
    v1 := []string{"\n\t", ("foo bar" + "")}
    var v2 float64 = 0.0
    v3 := []uint32{1, 640, 0}
    var v4 float64 = v2
    v5 := []int{}
    v4 += v2
    v2 = (v4 - (1e10 + float64(((0.1 / 1.5) * float32(v2)))))
    v2 = (0.0 / 1.0)
    fmt.Println(v4, v4)
    var v21 int16 = -1
    v4 += float64((float32((float64(float32(3.14159)) * (v2 + 0.1))) + 3.14159))
    v4 = v4
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        v4 = (v4 * 0.1)
        switch fi22 {
        case 32768:
            v21 = v21
        case 127:
            v2 -= 1e10
        default:
            for fi23 := 0; fi23 < 12; fi23++ {
                _ = len(v3)
                var v24 uint = 0
                var v25 int16 = (v21 % ((1000) | 1))
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        _ = fi22
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        for fi27 := 0; fi27 < 20; fi27++ {
            var v28 uint = 0
            v2 += v4
            if false {
                switch ("a" + "") {
                case "foo bar":
                    if false {
                        var v29 int32 = genFn1(71)
                        switch "\n\t" {
                        case "hello":
                            v21 -= v21
                        case "":
                            v21 -= v21
                        default:
                            v30 := GenS2{F0: -111}
                            _ = v30
                        }
                        v3 = append(v3, 4294967295)
                        fmt.Println(fi26, v21, v2)
                        var v31 int8 = 100
                        v32 := GenS3{F0: 42}
                        _ = v29
                        _ = v31
                        _ = v32
                    } else {
                        v21 += v21
                        v33 := GenS2{F0: 0}
                        switch "\n\t" {
                        case "":
                            var v34 uint8 = 127
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v34
                        case "hello":
                            v35 := GenS4{F0: 42, F1: 4294967295, F2: -141}
                            _ = v35
                        default:
                            switch (fi26 << ((2) & 63)) {
                            case 29:
                                v2 -= v4
                            }
                        }
                        _ = v33
                    }
                case "\n\t":
                    fmt.Println(v4)
                case "hello":
                    var v36 int = 823
                    _ = v36
                default:
                    if false {
                        var v37 bool = true
                        _ = v37
                    } else {
                        if false {
                            if ("a" < "a") {
                                var v38 int32 = 0
                                if (!true) {
                                    switch fi26 {
                                    case 2147483647:
                                        switch (("\n\t" + "hello") + ("\n\t" + ("a" + "foo bar"))) {
                                        case "hello":
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        case "a":
                                            v21 += (v21 ^ v21)
                                        case "":
                                            v39 := GenS5{F0: 4294967295}
                                            _ = v39
                                        default:
                                            var v40 uint64 = 2
                                            var v41 int32 = (v38 >> ((v28) & 63))
                                            v3 = append(v3, 0)
                                            _ = v40
                                            _ = v41
                                        }
                                    }
                                }
                                _ = v38
                            } else {
                                var v42 int32 = genFn1(-20)
                                var v43 uint32 = 2
                                _ = v42
                                _ = v43
                            }
                            if false {
                                if (-128 < 0) {
                                    fmt.Println(v28, v21)
                                    v21 = v21
                                    v4 = float64(float32(v2))
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                v44 := GenS6{F0: 584}
                                switch 17 {
                                case 32767:
                                    v45 := GenS7{F0: false, F1: int32(v28)}
                                    _ = v45
                                }
                                _ = v44
                            }
                            v46 := GenS8{F0: "", F1: 42}
                            _ = v46
                        }
                        _ = len(v1)
                        var v47 int = (fi27 | (-(fi27)))
                        _ = v47
                    }
                }
            } else {
                var v48 uint32 = 2147483647
                if true {
                    var v49 uint64 = 4294967295
                    v28 = v28
                    _ = v49
                }
                _ = v48
            }
            _ = fi27
            _ = v28
        }
        v21 = (v21 | 938)
        v50 := GenS9{F0: (^(v21)), F1: 2, F2: -9223372036854775808}
        _ = fi26
        _ = v50
    }
    for fi51 := 0; fi51 < 12; fi51++ {
        for fi52 := 0; fi52 < 13; fi52++ {
            var v53 int8 = 72
            v54 := GenS10{F0: v21}
            if true {
                var v55 int32 = -1
                _ = v55
            }
            _ = fi52
            _ = v53
            _ = v54
        }
        for fi56 := 0; fi56 < 13; fi56++ {
            switch fi51 {
            case 2147483647:
                if true {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v1 = append(v1, "foo bar")
                    if true {
                        v1 = append(v1, "")
                        if len(v5) > 0 {
                            v5[0] = fi56
                        }
                    }
                } else {
                    _ = len(v3)
                    if (136 <= 2) {
                        switch ("a" + ("foo bar" + "a")) {
                        case "a":
                            var v57 uint16 = 1
                            _ = v57
                        case "\n\t":
                            v2 += float64(0.0)
                        case "hello":
                            switch -426 {
                            case 32767:
                                if ((v21 * v21) < (v21 - 32767)) {
                                    v58 := GenS11{F0: 65535, F1: 91, F2: 761}
                                    v5 = append(v5, (fi51 >> ((4294967295) & 63)))
                                    _ = v58
                                }
                            case 32768:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case 913:
                                v4 += (v4 + 1.5)
                            default:
                                v5 = append(v5, fi56)
                            }
                        }
                    } else {
                        var v59 int64 = 9223372036854775807
                        _ = v59
                    }
                    switch (fi56 / ((fi51) | 1)) {
                    case -9223372036854775808:
                        v4 = v4
                    case 1:
                        v2 -= -2.25
                    }
                    v60 := GenS1{F0: 128, F1: (v21 ^ (v21 ^ v21))}
                    _ = v60
                }
            case -9223372036854775808:
                var v61 float32 = 3.14159
                _ = v61
            case 515:
                if len(v1) > 0 {
                    v1[0] = (("hello" + "") + "a")
                }
            }
            _ = fi56
        }
        _ = fi51
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
