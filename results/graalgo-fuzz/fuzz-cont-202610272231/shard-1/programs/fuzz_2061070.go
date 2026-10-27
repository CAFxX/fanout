package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 uint16; F2 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32; F1 int32; F2 int8 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 uint8, p1 int8) int8 {
    switch ("a" + "\n\t") {
    case "\n\t":
        var v7 uint = 0
        _ = v7
    case "foo bar":
        var v8 int16 = -1
        _ = v8
    case "hello":
        fmt.Println(p1, p1)
    }
    fmt.Println(p0, p1)
    var v9 int8 = 127
    _ = p0
    _ = p1
    _ = v9
    return 42
}
func genFn1() int16 {
    var v10 uint16 = 255
    if false {
        v10 -= v10
        for fi11 := 0; fi11 < 19; fi11++ {
            v12 := []int{(fi11 >> ((132) & 63)), (^(fi11)), fi11, fi11}
            if ((true && (!false)) && (0.0 <= 0.1)) {
                for fi13 := 0; fi13 < 11; fi13++ {
                    var v14 uint16 = v10
                    fmt.Println(fi11, fi13, v14)
                    _ = fi13
                    _ = v14
                }
            } else {
                if (true || false) {
                    for fi15 := 0; fi15 < 3; fi15++ {
                        var v16 int = (^(fi15))
                        v10 -= v10
                        var v17 float64 = 0.0
                        v18 := GenS1{F0: uint64(v10), F1: genFn0(42, -41)}
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    v10 -= (v10 - (v10 << ((18446744073709551615) & 63)))
                    var v19 float64 = 0.0
                    _ = v19
                } else {
                    var v20 int32 = -983
                    v21 := GenS1{F0: 904, F1: 44}
                    _ = v20
                    _ = v21
                }
                switch (^(fi11)) {
                case 127:
                    var v22 uint = 698
                    _ = v22
                case 710:
                    if false {
                        for fi23 := 0; fi23 < 13; fi23++ {
                            var v24 uint32 = 2147483648
                            _ = fi23
                            _ = v24
                        }
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        fmt.Println(fi11, v10)
                    } else {
                        for fi25 := 0; fi25 < 15; fi25++ {
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            var v26 int32 = 1
                            _ = fi25
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 15; fi27++ {
                            v10 = 720
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            v10 -= (v10 & v10)
                            if (!false) {
                                fmt.Println(fi27, fi11)
                                v28 := GenS1{F0: 521, F1: 100}
                                if false {
                                    fmt.Println(fi27, v10)
                                    var v29 float64 = 0.0
                                    fmt.Println(fi27, v10, fi27)
                                    v30 := GenS2{F0: -9, F1: 23}
                                    v29 = v29
                                    _ = v29
                                    _ = v30
                                }
                                _ = v28
                            } else {
                                fmt.Println(fi27, v10, fi11)
                                v12 = append(v12, (fi11 & 127))
                                if false {
                                    if (!true) {
                                        v31 := GenS3{F0: 160}
                                        fmt.Println(v10, fi11)
                                        fmt.Println(fi11)
                                        if true {
                                            if (1 > 199) {
                                                var v32 int32 = -2147483648
                                                _ = v32
                                            } else {
                                                v10 -= v10
                                                v12 = append(v12, int(v10))
                                                v10 = 32767
                                            }
                                        }
                                        _ = v31
                                    }
                                    if true {
                                        var v33 bool = true
                                        _ = len(v12)
                                        _ = v33
                                    } else {
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                        v12 = append(v12, (fi27 >> ((18446744073709551615) & 63)))
                                        v10 = v10
                                    }
                                }
                                if len(v12) > 0 {
                                    v12[0] = (fi11 ^ 4294967295)
                                }
                            }
                            var v34 uint8 = 128
                            v35 := GenS4{F0: (-(fi11)), F1: (^(v10)), F2: (v10 ^ v10)}
                            if false {
                                v36 := GenS5{F0: 163}
                                var v37 uint8 = v34
                                _ = v36
                                _ = v37
                            }
                            _ = fi27
                            _ = v34
                            _ = v35
                        }
                    }
                default:
                    v38 := GenS5{F0: 2}
                    _ = v38
                }
            }
            v39 := GenS1{F0: 42, F1: 100}
            _ = fi11
            _ = v12
            _ = v39
        }
    }
    var v40 string = ("" + ("hello" + "\n\t"))
    _ = v10
    _ = v40
    return -1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 0
    var v3 uint64 = 42
    var v4 uint16 = 65535
    var v5 bool = ((("hello" + "foo bar") + "foo bar") == (("" + ((("" + "a") + "\n\t") + (("hello" + "\n\t") + ("foo bar" + "a")))) + "foo bar"))
    var v6 string = (((("a" + "") + "") + "\n\t") + "")
    for fi41 := 0; fi41 < 2; fi41++ {
        if true {
            switch 686 {
            case 18:
                var v42 uint8 = 255
                _ = v42
            case -2147483648:
                v43 := v1["\n\t"]
                _ = v43
            }
            switch ("a" + v6) {
            case "hello":
                v44 := []int{}
                _ = v44
            case "":
                v4 += v4
            }
            for fi45 := 0; fi45 < 6; fi45++ {
                v2 = (^(v2))
                _ = fi45
            }
        } else {
            v5 = (255 != 128)
            var v46 int64 = 953
            v3 += v3
            _ = v46
        }
        var v47 int32 = v2
        for fi48 := 0; fi48 < 13; fi48++ {
            v3 += v3
            _ = fi48
        }
        _ = fi41
        _ = v47
    }
    var v49 string = v6
    if v5 {
        v6 = ((v49 + v6) + (("a" + (("hello" + ((v6 + ((((v6 + v6) + "foo bar") + "") + v6)) + "")) + "")) + ""))
    } else {
        var v50 int64 = 1
        _ = v50
    }
    v2 += 0
    v51 := GenS6{F0: 325, F1: (v2 ^ (-(v2))), F2: 100}
    var v52 bool = false
    for fi53 := 0; fi53 < 8; fi53++ {
        var v54 float32 = float32(3.14159)
        switch int(v2) {
        case -302:
            v55 := GenS3{F0: 1}
            _ = v55
        case 128:
            var v56 string = ((v6 + v6) + "")
            _ = v56
        default:
            var v57 string = ((v49 + v6) + (v49 + v49))
            _ = v57
        }
        _ = fi53
        _ = v54
    }
    v4 = v4
    var v58 float64 = float64(float32(((0.0 - 0.0) * 1.5)))
    fmt.Println(v3, v3)
    v5 = (!v52)
    if true {
        for fi59 := 0; fi59 < 20; fi59++ {
            if v5 {
                if v52 {
                    v6 = v6
                    if v5 {
                        if (v2 >= (v2 * -2147483648)) {
                            v60 := GenS7{F0: v4}
                            var v61 bool = v52
                            _ = v60
                            _ = v61
                        } else {
                            v62 := GenS1{F0: (^(v3)), F1: 100}
                            var v63 float64 = v58
                            _ = v62
                            _ = v63
                        }
                        v64 := GenS2{F0: int32(v3), F1: genFn0(255, -1)}
                        _ = v64
                    } else {
                        for fi65 := 0; fi65 < 16; fi65++ {
                            var v66 int8 = -1
                            var v67 uint16 = (v4 & v4)
                            if true {
                                v68 := GenS5{F0: 42}
                                switch v49 {
                                case "hello":
                                    fmt.Println(v66, v3, v6)
                                }
                                var v69 uint8 = 0
                                _ = v68
                                _ = v69
                            }
                            _ = fi65
                            _ = v66
                            _ = v67
                        }
                        var v70 float64 = (v58 + v58)
                        _ = v70
                    }
                    v71 := []int{(fi59 & fi59), fi59}
                    _ = v71
                }
            }
            _ = fi59
        }
    } else {
        switch v49 {
        case "\n\t":
            v72 := GenS4{F0: -595, F1: (^(v4)), F2: v4}
            _ = v72
        case "foo bar":
            fmt.Println(v5, v4, v6)
        default:
            v73 := []int{2147483647, 128, -2147483648, -986}
            _ = v73
        }
        fmt.Println(v58, v58)
    }
    fmt.Println(v58)
    v6 = v49
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v51
    _ = v52
    _ = v58
    _ = v6
}
