package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 int16; F2 int16 }
type GenS3 struct { F0 uint8; F1 int; F2 uint32 }
type GenS4 struct { F0 int; F1 int8; F2 string }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint32, p1 uint64, p2 uint8) int8 {
    var v7 uint32 = 4294967295
    p2 += (p2 << ((42) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 100
}

func main() {
    var v1 int = -460
    var v2 int16 = -1
    v3 := []int{(v1 | 42), (v1 - 32768), v1, v1}
    v4 := make(map[int]int)
    v5 := []uint32{}
    var v6 int16 = v2
    for fi8 := 0; fi8 < 11; fi8++ {
        fmt.Println(v1, v2)
        _ = fi8
    }
    if false {
        v2 += 1
        v6 += (^(v6))
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 uint64 = 978
            _ = fi9
            _ = v10
        }
    }
    if true {
        for fi11 := 0; fi11 < 12; fi11++ {
            if len(v5) > 0 {
                v5[0] = 65535
            }
            v12 := GenS1{F0: 1}
            _ = fi11
            _ = v12
        }
        var v13 int64 = -9223372036854775808
        _ = v13
    }
    var v14 uint64 = 4294967295
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 uint16 = 0
        fmt.Println(v6, fi15)
        var v17 string = ((((("foo bar" + (("" + "a") + "foo bar")) + "a") + "foo bar") + "foo bar") + (("hello" + "a") + ("foo bar" + "\n\t")))
        _ = fi15
        _ = v16
        _ = v17
    }
    v14 += v14
    v18 := GenS1{F0: 494}
    if true {
        if (!(!((!false) && (false || (false || false))))) {
            if true {
                v1 = v1
                if false {
                    var v19 uint32 = 2147483648
                    v20 := GenS1{F0: 2}
                    _ = v19
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 11; fi21++ {
                        if (42 != 311) {
                            if (!(!(v2 == v2))) {
                                fmt.Println(fi21, v6, fi21)
                                v14 = (v14 / (((v14 << ((270) & 63))) | 1))
                            } else {
                                v14 = v14
                                if (936 <= 591) {
                                    v22 := GenS1{F0: 928}
                                    for fi23 := 0; fi23 < 13; fi23++ {
                                        _ = len(v3)
                                        var v24 float32 = ((3.14159 + 0.1) - 0.1)
                                        var v25 uint16 = 1
                                        _ = fi23
                                        _ = v24
                                        _ = v25
                                    }
                                    var v26 int16 = v6
                                    _ = v22
                                    _ = v26
                                } else {
                                    var v27 int = (fi21 + 4294967295)
                                    if (!true) {
                                        fmt.Println(fi21)
                                        v1 += 42
                                    }
                                    for fi28 := 0; fi28 < 12; fi28++ {
                                        var v29 float64 = 1.5
                                        _ = fi28
                                        _ = v29
                                    }
                                    _ = v27
                                }
                                v30 := GenS2{F0: 127, F1: v2, F2: v2}
                                _ = v30
                            }
                            var v31 uint64 = 0
                            if (1 == 2) {
                                var v32 uint16 = 640
                                for fi33 := 0; fi33 < 7; fi33++ {
                                    switch (-(fi33)) {
                                    case 32767:
                                        if false {
                                            v31 = v31
                                            var v34 int8 = -61
                                            fmt.Println(v31)
                                            _ = v34
                                        }
                                    case 32768:
                                        delete(v4, (v1 << ((352) & 63)))
                                    case -2147483648:
                                        var v35 float64 = 1.5
                                        _ = v35
                                    }
                                    var v36 int = (fi33 + (fi33 - (v1 & 32767)))
                                    v31 = v31
                                    var v37 uint64 = 488
                                    fmt.Println(v6)
                                    _ = fi33
                                    _ = v36
                                    _ = v37
                                }
                                var v38 int64 = 0
                                _ = v32
                                _ = v38
                            }
                            _ = v31
                        }
                        var v39 string = "hello"
                        _ = fi21
                        _ = v39
                    }
                }
            } else {
                switch "\n\t" {
                case "hello":
                    switch int(v14) {
                    case -387:
                        if true {
                            for fi40 := 0; fi40 < 20; fi40++ {
                                fmt.Println(fi40, v6, fi40)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v41 := GenS2{F0: 100, F1: (v6 ^ 1), F2: (v6 * -294)}
                                _ = fi40
                                _ = v41
                            }
                            fmt.Println(v6)
                            v14 -= 0
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v42 := GenS2{F0: -1, F1: (-(v2)), F2: (^(v6))}
                            _ = v42
                        }
                    default:
                        var v43 uint32 = 0
                        _ = v43
                    }
                case "\n\t":
                    if false {
                        v6 = v2
                    } else {
                        v6 = (v2 * v2)
                        var v44 float32 = 3.14159
                        _ = v44
                    }
                case "":
                    var v45 bool = ((v14 >> ((1) & 63)) >= v14)
                    _ = v45
                }
            }
        }
        v46 := GenS2{F0: 77, F1: 42, F2: (v2 + (v2 << ((42) & 63)))}
        _ = v46
    }
    if false {
        for fi47 := 0; fi47 < 8; fi47++ {
            v48 := GenS2{F0: 100, F1: (v2 - (v6 * (^(v2)))), F2: v6}
            _ = fi47
            _ = v48
        }
        for fi49 := 0; fi49 < 7; fi49++ {
            var v50 int8 = 100
            v14 = (v14 << ((42) & 63))
            for fi51 := 0; fi51 < 1; fi51++ {
                _ = len(v3)
                v52 := GenS3{F0: 236, F1: fi51, F2: 669}
                _ = fi51
                _ = v52
            }
            _ = fi49
            _ = v50
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi53 := 0; fi53 < 2; fi53++ {
        for fi54 := 0; fi54 < 19; fi54++ {
            if (true && ((((v1 << ((319) & 63)) < fi54) && false) || ((-2.25 > float64(float32(1.5))) || true))) {
                var v55 uint32 = 2147483648
                _ = v55
            }
            _ = fi54
        }
        for fi56 := 0; fi56 < 17; fi56++ {
            switch -9223372036854775808 {
            case 32767:
                if (726 <= 515) {
                    v57 := GenS1{F0: 347}
                    v2 = (v2 - (v2 + (v2 * (v6 / ((v2) | 1)))))
                    if false {
                        fmt.Println(v1, v2)
                        v58 := GenS4{F0: (fi53 >> ((2) & 63)), F1: 102, F2: "a"}
                        v2 = v6
                        _ = v58
                    }
                    _ = v57
                } else {
                    if false {
                        var v59 int32 = 1
                        switch fi53 {
                        case 32768:
                            var v60 float64 = 1.5
                            _ = v60
                        default:
                            v59 = v59
                        }
                        if (1.5 == (((float32(((0.1 - 1.5) * (3.14159 - 1.5))) - 1e10) + (0.1 * 0.0)) + (((0.0 / 1.5) + ((1.5 + 0.0) + 0.0)) + float32(-2.25)))) {
                            var v61 uint64 = v14
                            v6 -= 32767
                            _ = v61
                        } else {
                            var v62 uint = 255
                            _ = v62
                        }
                        _ = v59
                    }
                    v63 := GenS3{F0: 0, F1: -2147483648, F2: 2}
                    _ = v63
                }
            case -123:
                if false {
                    v2 += 493
                    if true {
                        _ = len(v3)
                    }
                }
            case 32768:
                switch "hello" {
                case "a":
                    var v64 uint32 = 42
                    _ = v64
                case "\n\t":
                    if (804 >= 9223372036854775807) {
                        _ = len(v5)
                        var v65 string = ("hello" + "a")
                        delete(v4, v1)
                        v14 = v14
                        _ = v65
                    }
                }
            }
            var v66 int32 = 0
            _ = fi56
            _ = v66
        }
        _ = fi53
    }
    v6 = 224
    v67 := GenS5{F0: v14}
    v14 = (v14 ^ v14)
    delete(v4, (^(v1)))
    var v68 float32 = float32((((0.0 - 0.1) + 3.14159) * 1.5))
    v6 = (v2 / (((v2 | 1)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v68
}
