package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint; F2 int64 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint32, p1 int) int64 {
    p0 -= (p0 & (^(p0)))
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 47
}
func genFn1(p0 int32) int8 {
    var v6 uint16 = 42
    if false {
        var v7 float64 = 0.0
        switch 54 {
        case 32768:
            fmt.Println(v7, p0, p0)
        }
        _ = v7
    } else {
        if true {
            var v8 uint16 = 781
            _ = v8
        }
        var v9 int32 = p0
        if true {
            v10 := []int{}
            for fi11 := 0; fi11 < 15; fi11++ {
                fmt.Println(v6, p0, v9)
                _ = fi11
            }
            _ = v10
        }
        _ = v9
    }
    p0 -= 829
    p0 = (p0 % ((p0) | 1))
    v12 := GenS1{F0: 18446744073709551615, F1: 0}
    _ = p0
    _ = v12
    _ = v6
    return 127
}

func main() {
    var v1 string = ("" + ("a" + "a"))
    v2 := make(map[int]int)
    var v3 int16 = 0
    v4 := []uint32{733, 65535}
    v5 := make(map[uint32]int)
    var v13 float64 = (1.5 * 3.14159)
    v1 = "a"
    var v14 uint8 = 128
    fmt.Println(v3, v14, v3)
    v4 = append(v4, 473)
    for fi15 := 0; fi15 < 7; fi15++ {
        if false {
            if (true && false) {
                if true {
                    v16 := GenS1{F0: 255, F1: 1}
                    v17 := GenS1{F0: 18446744073709551615, F1: 1}
                    _ = v16
                    _ = v17
                }
                if (!(fi15 > (fi15 + fi15))) {
                    var v18 uint64 = 2
                    var v19 int8 = -1
                    switch 0 {
                    case 262:
                        _ = len(v2)
                    case 42:
                        if false {
                            if ((!(genFn0(65535, fi15) == genFn0(uint32(v18), fi15))) || (42 != 127)) {
                                fmt.Println(v19, v14, v14)
                                v13 += 3.14159
                            }
                            for fi20 := 0; fi20 < 19; fi20++ {
                                switch fi15 {
                                case 127:
                                    v14 += 0
                                case 32767:
                                    var v21 int8 = 0
                                    _ = v21
                                case 4294967295:
                                    if false {
                                        if false {
                                            if false {
                                                v22 := GenS1{F0: 0, F1: -199}
                                                _ = v22
                                            }
                                        }
                                        v1 = v1
                                        v19 = v19
                                    } else {
                                        var v23 float32 = ((0.1 * 1e10) / float32((float64(1e10) / v13)))
                                        var v24 uint16 = 808
                                        v3 += (v3 | v3)
                                        _ = v23
                                        _ = v24
                                    }
                                }
                                v25 := GenS1{F0: 255, F1: genFn0(2147483647, (fi20 << ((65535) & 63)))}
                                _ = fi20
                                _ = v25
                            }
                        } else {
                            fmt.Println(v1, v13)
                            var v26 float64 = 1.5
                            switch (fi15 * 32767) {
                            case 42:
                                for fi27 := 0; fi27 < 19; fi27++ {
                                    if (1 < -76) {
                                        var v28 uint32 = 0
                                        var v29 uint = 8
                                        _ = v28
                                        _ = v29
                                    }
                                    switch v1 {
                                    case "foo bar":
                                        v26 = v13
                                    case "hello":
                                        v3 = (v3 << ((42) & 63))
                                    }
                                    var v30 uint = 65535
                                    _ = fi27
                                    _ = v30
                                }
                            case 33:
                                v14 += v14
                            case 333:
                                var v31 int32 = -1
                                _ = v31
                            }
                            _ = v26
                        }
                    }
                    _ = v18
                    _ = v19
                } else {
                    v13 += -2.25
                }
            } else {
                v32 := GenS1{F0: 18446744073709551615, F1: 42}
                var v33 string = ((v1 + (v1 + ("\n\t" + v1))) + "a")
                _ = v32
                _ = v33
            }
            v34 := v2[fi15]
            var v35 int32 = 2147483647
            _ = v34
            _ = v35
        } else {
            if (0 < 576) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v36 int8 = 0
                _ = v36
            }
            if len(v4) > 0 {
                v4[0] = 0
            }
            if (true && false) {
                _ = len(v5)
                v14 = v14
                v37 := GenS2{F0: 0}
                _ = v37
            } else {
                for fi38 := 0; fi38 < 16; fi38++ {
                    v14 += v14
                    var v39 uint = 255
                    var v40 bool = false
                    var v41 bool = v40
                    var v42 uint64 = 663
                    _ = fi38
                    _ = v39
                    _ = v40
                    _ = v41
                    _ = v42
                }
            }
        }
        switch (v1 + "") {
        case "hello":
            v3 = v3
        case "":
            v43 := GenS1{F0: 89, F1: -915}
            _ = v43
        case "foo bar":
            v14 -= (v14 + (^(v14)))
        }
        _ = fi15
    }
    if len(v4) > 0 {
        v4[0] = 261
    }
    v4 = append(v4, 2)
    v44 := GenS2{F0: (^(v3))}
    v13 -= v13
    var v45 int8 = -1
    if (!((v45 | v45) != genFn1(1))) {
        if false {
            var v46 int8 = (v45 + genFn1(-1))
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v46
        } else {
            var v47 string = v1
            delete(v5, 0)
            v1 = ("" + v1)
            _ = v47
        }
        for fi48 := 0; fi48 < 19; fi48++ {
            v49 := GenS1{F0: 75, F1: 42}
            var v50 uint64 = 0
            if true {
                fmt.Println(v13, v50, v3)
                var v51 uint8 = v14
                switch "foo bar" {
                case "a":
                    delete(v5, 1)
                default:
                    var v52 int = 127
                    _ = v52
                }
                _ = v51
            }
            _ = fi48
            _ = v49
            _ = v50
        }
        var v53 float64 = -2.25
        _ = v53
    } else {
        if false {
            v13 += (float64((0.1 * (3.14159 + 3.14159))) - (v13 - v13))
            var v54 int16 = (-(v3))
            _ = len(v4)
            _ = v54
        } else {
            if false {
                v45 -= (v45 >> ((1) & 63))
                if len(v4) > 0 {
                    v4[0] = 87
                }
                v55 := v2[127]
                _ = v55
            } else {
                v56 := GenS2{F0: v3}
                switch 770 {
                case 32767:
                    var v57 uint8 = v14
                    _ = v57
                case -45:
                    switch "" {
                    case "\n\t":
                        var v58 uint32 = 962
                        _ = v58
                    case "foo bar":
                        var v59 uint16 = 65535
                        _ = v59
                    case "":
                        for fi60 := 0; fi60 < 10; fi60++ {
                            if len(v4) > 0 {
                                v4[0] = 900
                            }
                            _ = fi60
                        }
                    }
                case 127:
                    var v61 int32 = -1
                    _ = v61
                }
                _ = v56
            }
            var v62 bool = true
            _ = v62
        }
        v63 := GenS3{F0: 65535, F1: 255, F2: 59}
        var v64 uint8 = v14
        _ = v63
        _ = v64
    }
    var v65 int64 = 825
    var v66 float32 = (0.0 * ((((1e10 - float32(((v13 - float64(3.14159)) - v13))) * 0.1) / -2.25) * 3.14159))
    fmt.Println(v13, v45, v1)
    if true {
        if false {
            v5[0] = 128
            var v67 int16 = 59
            _ = v67
        } else {
            var v68 float32 = -2.25
            for fi69 := 0; fi69 < 15; fi69++ {
                var v70 uint32 = 2
                var v71 float32 = 3.14159
                _ = fi69
                _ = v70
                _ = v71
            }
            if (1 <= -1) {
                switch v1 {
                case "a":
                    var v72 bool = true
                    _ = v72
                }
            } else {
                v66 = float32((v13 - (v13 * (((0.0 * (v13 + float64(v66))) + v13) / v13))))
                v13 = float64(float32(v13))
            }
            _ = v68
        }
    } else {
        fmt.Println(v14, v66, v1)
    }
    v73 := GenS3{F0: 963, F1: 18446744073709551615, F2: -919}
    switch -1 {
    case 790:
        v45 += (^(v45))
    default:
        v74 := GenS4{F0: v14}
        _ = v74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v65
    _ = v66
    _ = v73
}
