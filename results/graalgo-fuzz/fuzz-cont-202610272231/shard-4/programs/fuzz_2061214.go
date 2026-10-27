package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint16 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 int16; F2 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 bool; F1 bool; F2 int16 }
type GenS8 struct { F0 uint16; F1 int8 }
type GenS9 struct { F0 uint32 }
func genFn0(p0 uint, p1 int16, p2 string) int8 {
    v6 := GenS1{F0: 128, F1: 128, F2: 255}
    var v7 int64 = 1
    for fi8 := 0; fi8 < 17; fi8++ {
        switch p2 {
        case "a":
            var v9 float64 = 1.5
            _ = v9
        case "\n\t":
            v10 := []int{fi8}
            _ = v10
        }
        p2 = p2
        _ = fi8
    }
    v11 := []int{1, int(p1), 742, -9223372036854775808}
    switch -912 {
    case -306:
        v12 := GenS2{F0: (4294967295 < -804), F1: p2}
        _ = v12
    case 16:
        var v13 uint8 = 0
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    _ = v7
    return -128
}

func main() {
    var v1 int8 = 42
    var v2 float64 = -2.25
    var v3 uint = 862
    var v4 uint64 = 1
    var v5 uint64 = uint64(v3)
    var v14 int8 = v1
    switch 32767 {
    case 4294967295:
        switch 127 {
        case 32767:
            var v15 int8 = v14
            _ = v15
        case -2147483648:
            v16 := []int{-370, 9223372036854775807}
            _ = v16
        default:
            for fi17 := 0; fi17 < 16; fi17++ {
                if false {
                    switch (("\n\t" + ("a" + "")) + ("a" + "hello")) {
                    case "hello":
                        v18 := GenS2{F0: true, F1: "foo bar"}
                        _ = v18
                    case "foo bar":
                        switch "foo bar" {
                        case "\n\t":
                            v2 = (float64(-2.25) * float64(-2.25))
                        case "":
                            v19 := GenS3{F0: 588}
                            _ = v19
                        case "foo bar":
                            if true {
                                v5 = v4
                                switch fi17 {
                                case 127:
                                    var v20 bool = true
                                    _ = v20
                                default:
                                    switch "hello" {
                                    case "foo bar":
                                        var v21 uint = (v3 | (v3 / ((v3) | 1)))
                                        _ = v21
                                    }
                                }
                            } else {
                                if ((2147483648 < 4294967295) && (!false)) {
                                    v22 := []int{-2147483648}
                                    _ = v22
                                } else {
                                    var v23 uint16 = 1
                                    _ = v23
                                }
                                var v24 uint32 = 0
                                var v25 float64 = v2
                                _ = v24
                                _ = v25
                            }
                        default:
                            if (0 < 2147483647) {
                                v26 := GenS2{F0: (128 > 1), F1: ""}
                                v1 = 42
                                _ = v26
                            } else {
                                v27 := GenS4{F0: fi17, F1: -913, F2: 127}
                                v28 := GenS2{F0: (-2.25 < 3.14159), F1: ("a" + ((("" + ("foo bar" + (("hello" + "hello") + ("" + "\n\t")))) + "foo bar") + "hello"))}
                                v3 += v3
                                _ = v27
                                _ = v28
                            }
                        }
                    default:
                        v29 := GenS5{F0: 42}
                        _ = v29
                    }
                    if (float32(v2) == (float32(v2) + 0.0)) {
                        var v30 uint64 = (v5 & 724)
                        fmt.Println(v5, v2)
                        v31 := GenS2{F0: (0 >= -678), F1: "foo bar"}
                        _ = v30
                        _ = v31
                    }
                }
                switch -747 {
                case -540:
                    v2 = 0.0
                case -955:
                    var v32 int32 = -1
                    _ = v32
                case 319:
                    var v33 string = "foo bar"
                    _ = v33
                default:
                    v2 -= v2
                }
                v34 := []int{fi17, (fi17 / ((-2147483648) | 1))}
                _ = fi17
                _ = v34
            }
        }
    case 0:
        for fi35 := 0; fi35 < 16; fi35++ {
            if true {
                v36 := []int{int(v4), fi35}
                if (0 <= v5) {
                    switch "\n\t" {
                    case "foo bar":
                        _ = len(v36)
                    case "hello":
                        var v37 float64 = v2
                        _ = v37
                    case "\n\t":
                        v14 -= 1
                    }
                    var v38 bool = false
                    var v39 int8 = v1
                    _ = v38
                    _ = v39
                }
                var v40 bool = true
                _ = v36
                _ = v40
            } else {
                for fi41 := 0; fi41 < 1; fi41++ {
                    if false {
                        var v42 uint64 = v4
                        v43 := GenS1{F0: fi41, F1: 1, F2: 1}
                        _ = v42
                        _ = v43
                    } else {
                        var v44 string = "hello"
                        if false {
                            var v45 uint16 = 4
                            _ = v45
                        } else {
                            v46 := []int{4294967295, fi35, fi35, -458}
                            _ = v46
                        }
                        fmt.Println(v2, fi35, v2)
                        var v47 bool = false
                        var v48 uint16 = 543
                        v4 -= v4
                        v5 = 18446744073709551615
                        _ = v44
                        _ = v47
                        _ = v48
                    }
                    v3 -= v3
                    var v49 int32 = int32(v5)
                    _ = fi41
                    _ = v49
                }
            }
            v50 := []int{(fi35 ^ (fi35 >> ((47) & 63))), (fi35 ^ 32767), fi35, fi35}
            var v51 float64 = v2
            _ = fi35
            _ = v50
            _ = v51
        }
    }
    v52 := GenS6{F0: 250, F1: 1}
    var v53 uint32 = 2
    for fi54 := 0; fi54 < 3; fi54++ {
        if true {
            for fi55 := 0; fi55 < 19; fi55++ {
                var v56 uint = v3
                v57 := []int{fi55}
                var v58 int64 = -1
                _ = fi55
                _ = v56
                _ = v57
                _ = v58
            }
            for fi59 := 0; fi59 < 14; fi59++ {
                var v60 float64 = v2
                _ = fi59
                _ = v60
            }
        } else {
            v14 = (^(v14))
        }
        v61 := GenS7{F0: true, F1: (124 <= 136), F2: 42}
        for fi62 := 0; fi62 < 6; fi62++ {
            if false {
                var v63 int32 = 0
                _ = v63
            }
            if true {
                v64 := []int{(fi54 * fi62), fi62}
                _ = v64
            } else {
                if true {
                    v4 = 1
                    v5 += v4
                    v65 := []int{fi62, -9223372036854775808, int(v5), fi54}
                    _ = v65
                } else {
                    var v66 int32 = -1
                    _ = v66
                }
                var v67 float64 = (((v2 + v2) * v2) - ((v2 - float64((0.0 - ((0.1 - 1.5) * ((1e10 / 3.14159) - 0.1))))) + (0.0 + float64((1e10 - -2.25)))))
                v53 = (v53 ^ v53)
                _ = v67
            }
            v2 -= ((1e10 + 1.5) * (((v2 * v2) * v2) * float64(((-2.25 / ((3.14159 + ((float32(v2) * ((0.1 - (float32((v2 - v2)) - 1.5)) * 0.0)) / (float32(1e10) + 0.0))) / float32(((1.5 - (v2 + 0.1)) - v2)))) * float32(v2)))))
            _ = fi62
        }
        _ = fi54
        _ = v61
    }
    v68 := GenS8{F0: 0, F1: (v14 - -128)}
    v69 := GenS9{F0: v53}
    v1 += int8(v5)
    fmt.Println(v53, v3, v5)
    v70 := GenS4{F0: -9223372036854775808, F1: -1, F2: 1}
    var v71 int8 = (v1 << ((v3) & 63))
    v72 := GenS5{F0: 2}
    fmt.Println(v5, v1)
    var v73 uint8 = uint8(v4)
    v3 += (v3 ^ (v3 - 564))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v68
    _ = v69
    _ = v70
    _ = v71
    _ = v72
    _ = v73
}
