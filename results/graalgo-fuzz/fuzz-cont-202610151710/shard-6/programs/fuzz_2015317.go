package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 string }
func genFn0(p0 uint, p1 uint32, p2 uint32) uint16 {
    var v7 float32 = 1.5
    var v8 string = ("" + "hello")
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 int = (fi9 << ((4294967295) & 63))
        if (42 >= 727) {
            switch (^(v10)) {
            case -2147483648:
                v8 = "hello"
            case 9223372036854775807:
                for fi11 := 0; fi11 < 16; fi11++ {
                    var v12 int32 = 0
                    v13 := GenS1{F0: p0}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            case 4294967295:
                v7 += v7
            default:
                if (2 >= 1) {
                    v14 := []int{(v10 << ((p0) & 63)), (v10 & (^(fi9)))}
                    fmt.Println(v8)
                    _ = v14
                } else {
                    v10 = 32767
                    var v15 int16 = -720
                    for fi16 := 0; fi16 < 9; fi16++ {
                        v7 = v7
                        v17 := []int{(v10 >> (((p0 ^ (p0 >> ((4294967295) & 63)))) & 63)), (fi9 | 1), 4294967295, v10}
                        if false {
                            v10 = 9223372036854775807
                            p0 -= (^(p0))
                        } else {
                            var v18 int32 = 1
                            if true {
                                var v19 uint64 = 0
                                var v20 float64 = 3.14159
                                _ = len(v17)
                                _ = v19
                                _ = v20
                            }
                            fmt.Println(fi9)
                            _ = v18
                        }
                        _ = fi16
                        _ = v17
                    }
                    _ = v15
                }
            }
            v21 := []int{(fi9 >> (((p0 << ((114) & 63))) & 63)), (fi9 / ((fi9) | 1)), (fi9 * (fi9 - (fi9 ^ -25)))}
            _ = v21
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 966
}

func main() {
    var v1 int64 = -318
    var v2 float64 = float64(0.0)
    var v3 uint = 870
    var v4 string = "foo bar"
    var v5 int32 = 1
    var v6 float64 = 1.5
    fmt.Println(v4, v3)
    var v22 int = 127
    switch "hello" {
    case "hello":
        var v23 bool = false
        _ = v23
    default:
        v22 -= (-(v22))
    }
    switch 2147483647 {
    case -29:
        fmt.Println(v4, v5)
    case -1:
        var v24 int8 = 0
        _ = v24
    default:
        v4 = v4
    }
    var v25 uint32 = 2147483648
    v25 = (v25 * v25)
    var v26 uint16 = 32767
    v1 += (v1 - v1)
    if true {
        if false {
            if (!(v5 > (v5 | v5))) {
                for fi27 := 0; fi27 < 12; fi27++ {
                    v5 = (-(v5))
                    var v28 bool = (false || true)
                    _ = fi27
                    _ = v28
                }
                switch "\n\t" {
                case "hello":
                    if false {
                        var v29 float32 = (0.0 + 3.14159)
                        _ = v29
                    }
                case "\n\t":
                    for fi30 := 0; fi30 < 9; fi30++ {
                        v31 := []int{(fi30 | fi30), (v22 % ((v22) | 1))}
                        _ = fi30
                        _ = v31
                    }
                }
            } else {
                for fi32 := 0; fi32 < 18; fi32++ {
                    v22 = 127
                    v33 := []int{32768, 962}
                    _ = fi32
                    _ = v33
                }
                switch (v22 | v22) {
                case 0:
                    switch v4 {
                    case "":
                        v6 = 0.1
                    }
                }
                for fi34 := 0; fi34 < 17; fi34++ {
                    v6 -= 0.1
                    _ = fi34
                }
            }
        }
        if (!((-32768 <= 178) && ((v5 * v5) == v5))) {
            v1 = (v1 - v1)
            v26 = (v26 | (v26 + genFn0(v3, v25, (^(v25)))))
            v35 := GenS2{F0: 867}
            _ = v35
        } else {
            v36 := []int{127, v22, (v22 % (((v22 * (v22 + 823))) | 1)), (v22 + (v22 / ((int(v26)) | 1)))}
            if len(v36) >= 2 {
                v36 = v36[1:len(v36)]
            }
            v37 := GenS1{F0: v3}
            _ = v36
            _ = v37
        }
    } else {
        for fi38 := 0; fi38 < 20; fi38++ {
            var v39 uint64 = 200
            v40 := GenS2{F0: int16(v5)}
            _ = fi38
            _ = v39
            _ = v40
        }
        if (!false) {
            v1 = (-(v1))
            if false {
                v41 := GenS3{F0: (-(v5)), F1: v4}
                _ = v41
            } else {
                fmt.Println(v2)
            }
        } else {
            switch int(v1) {
            case 896:
                if false {
                    v25 -= (^(v25))
                }
            case 2147483647:
                fmt.Println(v3)
            case 1:
                switch -994 {
                case 0:
                    switch v4 {
                    case "foo bar":
                        for fi42 := 0; fi42 < 3; fi42++ {
                            switch "a" {
                            case "foo bar":
                                v43 := GenS4{F0: ""}
                                _ = v43
                            default:
                                v44 := []int{v22}
                                _ = v44
                            }
                            _ = fi42
                        }
                    case "hello":
                        for fi45 := 0; fi45 < 2; fi45++ {
                            var v46 bool = ((1e10 > (3.14159 + (3.14159 * 0.1))) && true)
                            var v47 uint32 = 4294967295
                            fmt.Println(v26, v2)
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                    case "":
                        for fi48 := 0; fi48 < 7; fi48++ {
                            v6 = 1e10
                            if false {
                                if true {
                                    fmt.Println(v2, v22)
                                    var v49 uint = v3
                                    fmt.Println(v25, v26, v3)
                                    _ = v49
                                } else {
                                    v22 -= v22
                                }
                            }
                            _ = fi48
                        }
                    default:
                        fmt.Println(v5, v26, v4)
                    }
                default:
                    for fi50 := 0; fi50 < 7; fi50++ {
                        if ((1e10 + float32(3.14159)) != 0.0) {
                            for fi51 := 0; fi51 < 8; fi51++ {
                                var v52 int16 = 1
                                v4 = ""
                                _ = fi51
                                _ = v52
                            }
                            v2 = 3.14159
                        } else {
                            var v53 int64 = -324
                            var v54 int16 = 0
                            _ = v53
                            _ = v54
                        }
                        fmt.Println(v2, v2, v4)
                        for fi55 := 0; fi55 < 11; fi55++ {
                            var v56 int16 = -532
                            fmt.Println(v26, fi55, v3)
                            var v57 int = fi50
                            _ = fi55
                            _ = v56
                            _ = v57
                        }
                        _ = fi50
                    }
                }
            default:
                if (true && (-32768 > -150)) {
                    switch v4 {
                    case "foo bar":
                        for fi58 := 0; fi58 < 5; fi58++ {
                            if (false && ((true || (true || (-2.25 == (0.0 - (3.14159 - 0.1))))) && (false || (65535 < (^(v26)))))) {
                                v4 = v4
                            } else {
                                fmt.Println(v1)
                                v59 := []int{fi58, v22, (fi58 * v22)}
                                switch (v4 + v4) {
                                case "foo bar":
                                    for fi60 := 0; fi60 < 20; fi60++ {
                                        v4 = ("" + v4)
                                        _ = fi60
                                    }
                                case "a":
                                    v3 -= (v3 ^ 965)
                                default:
                                    var v61 float64 = (3.14159 / (v2 * (0.0 - float64(3.14159))))
                                    _ = v61
                                }
                                _ = v59
                            }
                            v6 += v6
                            _ = fi58
                        }
                    }
                    var v62 int = 32767
                    _ = v62
                }
            }
        }
        for fi63 := 0; fi63 < 7; fi63++ {
            var v64 float64 = ((float64(float32(v2)) + 1.5) * float64(-2.25))
            _ = fi63
            _ = v64
        }
    }
    for fi65 := 0; fi65 < 2; fi65++ {
        if true {
            v66 := []int{}
            v67 := GenS2{F0: -735}
            switch fi65 {
            case 163:
                if true {
                    var v68 string = v4
                    v69 := GenS2{F0: 0}
                    var v70 string = v68
                    _ = v68
                    _ = v69
                    _ = v70
                }
            default:
                v6 -= (((float64(1.5) / v2) / v6) + float64(0.1))
            }
            _ = v66
            _ = v67
        } else {
            for fi71 := 0; fi71 < 14; fi71++ {
                v72 := GenS2{F0: 1}
                var v73 uint16 = 65535
                v1 += (v1 >> (((v3 << ((uint(v26)) & 63))) & 63))
                _ = fi71
                _ = v72
                _ = v73
            }
            v2 -= v2
        }
        if true {
            v74 := []int{}
            _ = v74
        }
        _ = fi65
    }
    var v75 uint32 = (v25 + v25)
    switch "\n\t" {
    case "":
        for fi76 := 0; fi76 < 14; fi76++ {
            var v77 uint8 = 0
            fmt.Println(v77, v25)
            fmt.Println(fi76, v25)
            _ = fi76
            _ = v77
        }
    }
    fmt.Println(v4, v25)
    for fi78 := 0; fi78 < 8; fi78++ {
        var v79 int8 = 100
        v80 := []int{}
        _ = fi78
        _ = v79
        _ = v80
    }
    for fi81 := 0; fi81 < 16; fi81++ {
        v82 := []int{v22, (v22 >> (((v3 << ((uint(v25)) & 63))) & 63)), fi81}
        _ = fi81
        _ = v82
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
}
