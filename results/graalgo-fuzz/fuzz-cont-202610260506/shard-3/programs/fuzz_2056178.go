package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int16; F1 string; F2 int64 }
func genFn0() string {
    var v6 string = ("\n\t" + (("foo bar" + "hello") + ("" + "a")))
    v7 := GenS1{F0: 4294967295, F1: 42}
    var v8 int16 = 32767
    v9 := GenS1{F0: 405, F1: 400}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return v6
}
func genFn1(p0 int16, p1 int) int8 {
    var v10 uint = 924
    switch p1 {
    case 416:
        v11 := GenS1{F0: 18446744073709551615, F1: 36}
        _ = v11
    case 128:
        if (1e10 == ((((((float64(float32((0.1 + 1.5))) + 3.14159) - float64(1e10)) * (1e10 * (((1.5 * 1.5) / 1e10) + (1e10 + (float64(float32(1e10)) + 0.1))))) / 1e10) + (0.0 + 0.0)) - 0.1)) {
            for fi12 := 0; fi12 < 10; fi12++ {
                v10 = (v10 ^ (v10 ^ (v10 % ((660) | 1))))
                if (2147483648 <= 587) {
                    var v13 int8 = 1
                    _ = v13
                }
                var v14 bool = (-1 != 42)
                _ = fi12
                _ = v14
            }
            v15 := []int{(p1 << (((v10 ^ (v10 - (v10 ^ 4294967295)))) & 63)), p1}
            v10 = (^(v10))
            _ = v15
        }
    default:
        var v16 uint16 = 683
        _ = v16
    }
    p1 += p1
    p0 = p0
    p1 += (^(p1))
    _ = p0
    _ = p1
    _ = v10
    return 100
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]int)
    var v3 float64 = 0.0
    var v4 uint8 = 107
    var v5 int64 = 558
    if ((v1 + ((v1 / (1.5 + (v1 + ((v1 + (v1 / 1e10)) - (v1 * (v1 + float32(v3))))))) * 0.1)) >= (v1 + float32(v3))) {
        fmt.Println(v4, v3, v1)
        v3 = (v3 - (v3 + float64(v1)))
    }
    fmt.Println(v4)
    v17 := GenS1{F0: 652, F1: v5}
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 int32 = 282
        for fi20 := 0; fi20 < 12; fi20++ {
            fmt.Println(v19, v4)
            v1 -= 3.14159
            _ = fi20
        }
        v4 += 42
        _ = fi18
        _ = v19
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        switch fi21 {
        case 680:
            var v22 uint32 = 2147483647
            _ = v22
        case -1:
            v23 := GenS1{F0: 863, F1: 1}
            _ = v23
        case 4294967295:
            var v24 float64 = (3.14159 + v3)
            _ = v24
        default:
            v4 += v4
        }
        var v25 int8 = 1
        for fi26 := 0; fi26 < 16; fi26++ {
            _ = len(v2)
            v1 -= v1
            _ = fi26
        }
        _ = fi21
        _ = v25
    }
    switch ("foo bar" + "\n\t") {
    case "a":
        v27 := GenS1{F0: 320, F1: (v5 + (-(v5)))}
        _ = v27
    default:
        v28 := GenS1{F0: 4294967295, F1: (v5 ^ v5)}
        _ = v28
    }
    var v29 uint32 = 2147483648
    var v30 bool = false
    var v31 uint64 = 4294967295
    switch (("\n\t" + ((("hello" + "a") + "hello") + "foo bar")) + "\n\t") {
    case "foo bar":
        for fi32 := 0; fi32 < 18; fi32++ {
            var v33 bool = (!v30)
            var v34 uint8 = v4
            _ = fi32
            _ = v33
            _ = v34
        }
    case "a":
        fmt.Println(v3, v1, v31)
    default:
        v3 += v3
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        v36 := GenS1{F0: (v31 / ((v31) | 1)), F1: v5}
        _ = fi35
        _ = v36
    }
    if v30 {
        v37 := GenS1{F0: (v31 + v31), F1: v5}
        _ = v37
    } else {
        v38 := GenS1{F0: (v31 + uint64(v29)), F1: -1}
        if v30 {
            switch 32768 {
            case 32767:
                for fi39 := 0; fi39 < 4; fi39++ {
                    if (65535 == 300) {
                        var v40 int = fi39
                        _ = v40
                    } else {
                        for fi41 := 0; fi41 < 12; fi41++ {
                            var v42 int32 = -86
                            _ = fi41
                            _ = v42
                        }
                        var v43 uint = 0
                        v44 := GenS2{F0: (v5 >> ((v43) & 63)), F1: -1}
                        _ = v43
                        _ = v44
                    }
                    v29 = (v29 * uint32(v31))
                    _ = fi39
                }
            case 712:
                if v30 {
                    if v30 {
                        v31 -= uint64(v29)
                    } else {
                        for fi45 := 0; fi45 < 12; fi45++ {
                            v46 := GenS1{F0: uint64(v29), F1: v5}
                            var v47 int64 = (v5 & v5)
                            v48 := GenS2{F0: v5, F1: 730}
                            _ = fi45
                            _ = v46
                            _ = v47
                            _ = v48
                        }
                        if ("\n\t" > "") {
                            var v49 uint32 = (v29 - (v29 ^ (v29 * (^(v29)))))
                            v29 = 2
                            _ = v49
                        } else {
                            if v30 {
                                for fi50 := 0; fi50 < 18; fi50++ {
                                    v1 -= 3.14159
                                    var v51 float64 = float64(0.0)
                                    v2[uint32(v4)] = (fi50 * fi50)
                                    _ = fi50
                                    _ = v51
                                }
                                v52 := v2[0]
                                _ = v52
                            } else {
                                for fi53 := 0; fi53 < 10; fi53++ {
                                    v29 -= v29
                                    var v54 string = "\n\t"
                                    fmt.Println(v4, v29)
                                    _ = fi53
                                    _ = v54
                                }
                                var v55 int32 = -385
                                v1 = v1
                                _ = v55
                            }
                            fmt.Println(v1, v30, v1)
                        }
                        v31 = uint64(v29)
                    }
                } else {
                    var v56 int8 = 127
                    var v57 uint8 = uint8(v29)
                    switch (("\n\t" + ("foo bar" + "")) + "\n\t") {
                    case "":
                        for fi58 := 0; fi58 < 14; fi58++ {
                            switch (fi58 & (fi58 / ((fi58) | 1))) {
                            case -348:
                                var v59 bool = (((fi58 - fi58) <= (fi58 * -2147483648)) || v30)
                                _ = v59
                            case 32767:
                                v60 := GenS3{F0: 979, F1: ("hello" + ""), F2: 0}
                                _ = v60
                            case 32768:
                                v61 := GenS3{F0: 1000, F1: "", F2: v5}
                                _ = v61
                            }
                            _ = fi58
                        }
                    case "a":
                        fmt.Println(v56, v5)
                    case "foo bar":
                        v5 -= (v5 >> ((0) & 63))
                    default:
                        v62 := []int{}
                        _ = v62
                    }
                    _ = v56
                    _ = v57
                }
            case -2147483648:
                switch ("foo bar" + "hello") {
                case "":
                    var v63 bool = false
                    _ = v63
                default:
                    var v64 int64 = 694
                    _ = v64
                }
            }
            if (1 <= 1) {
                _ = len(v2)
                v30 = true
            }
            v3 += v3
        } else {
            v65 := v2[v29]
            var v66 uint16 = uint16(v31)
            v67 := GenS3{F0: -768, F1: ("" + "foo bar"), F2: (v5 | -1)}
            _ = v65
            _ = v66
            _ = v67
        }
        _ = v38
    }
    fmt.Println(v4, v5)
    if (42 > uint(v31)) {
        var v68 string = ""
        var v69 int8 = -128
        switch 4294967295 {
        case 38:
            var v70 uint = 0
            _ = v70
        case -2147483648:
            v30 = v30
        }
        _ = v68
        _ = v69
    }
    v5 -= (v5 / ((int64(v29)) | 1))
    fmt.Println(v29)
    v71 := []int{-9223372036854775808, 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v71:", len(v71), v71)
    _ = v1
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v71
}
