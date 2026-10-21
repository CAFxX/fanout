package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int8; F2 uint16 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 uint16; F1 int8; F2 string }
type GenS6 struct { F0 int16; F1 bool }
func genFn0(p0 int, p1 uint16) int32 {
    v4 := []int{(p0 + (p0 + (p0 << ((18446744073709551615) & 63)))), (p0 & p0), (p0 - p0), -9223372036854775808}
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: -107}
        _ = fi5
        _ = v6
    }
    if false {
        switch ("hello" + ("hello" + (("a" + "hello") + ""))) {
        case "foo bar":
            v7 := GenS1{F0: -128}
            _ = v7
        case "\n\t":
            var v8 int16 = -32768
            _ = v8
        default:
            v9 := GenS2{F0: 32767, F1: 0, F2: 32767}
            _ = v9
        }
        v4 = append(v4, p0)
        fmt.Println(p0, p1)
    } else {
        p1 = (p1 | 600)
    }
    p1 = p1
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v4
    return 284
}
func genFn1() int8 {
    var v10 string = "a"
    v11 := GenS1{F0: 19}
    v12 := GenS1{F0: 42}
    if (!(false && (259 <= -1))) {
        v13 := []int{486, 32767, 32767}
        var v14 bool = (319 >= 18446744073709551615)
        _ = len(v13)
        _ = v13
        _ = v14
    }
    _ = v10
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 string = "foo bar"
    var v2 bool = ((18446744073709551615 <= 773) || (1.5 < float32((0.0 - (1.5 * (1e10 / 1.0))))))
    var v3 int32 = 42
    v15 := GenS3{F0: -1, F1: (^(v3))}
    switch v1 {
    case "foo bar":
        fmt.Println(v2, v1)
    }
    var v16 uint16 = 2
    var v17 uint = 65535
    fmt.Println(v16)
    var v18 uint8 = 129
    v19 := []int{-1}
    var v20 uint16 = v16
    switch v1 {
    case "\n\t":
        for fi21 := 0; fi21 < 1; fi21++ {
            switch ("a" + "\n\t") {
            case "foo bar":
                v17 += 0
            case "":
                v22 := GenS3{F0: fi21, F1: genFn0(0, (^(v20)))}
                _ = v22
            case "a":
                fmt.Println(fi21, fi21)
            }
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 uint16 = v20
                if (v1 >= v1) {
                    var v25 int8 = -1
                    if true {
                        v1 = v1
                        v20 += uint16(v17)
                        if ((v17 ^ v17) >= v17) {
                            v18 += 255
                            v20 = (v16 % (((^(v24))) | 1))
                        } else {
                            v26 := GenS3{F0: fi21, F1: genFn0((fi21 << (((v17 | v17)) & 63)), 2)}
                            _ = v26
                        }
                    }
                    fmt.Println(fi21, fi23, v3)
                    _ = v25
                }
                _ = fi23
                _ = v24
            }
            switch ((v1 + ((v1 + (v1 + ("foo bar" + (v1 + (v1 + "a"))))) + v1)) + "foo bar") {
            case "foo bar":
                if true {
                    for fi27 := 0; fi27 < 20; fi27++ {
                        _ = len(v19)
                        if (!false) {
                            _ = len(v19)
                        } else {
                            v18 += v18
                            v16 -= 42
                            v17 -= v17
                        }
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        _ = fi27
                    }
                    for fi28 := 0; fi28 < 17; fi28++ {
                        v17 += v17
                        _ = fi28
                    }
                    for fi29 := 0; fi29 < 20; fi29++ {
                        v2 = (42 >= -128)
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        _ = fi29
                    }
                }
            case "":
                for fi30 := 0; fi30 < 20; fi30++ {
                    switch fi21 {
                    case -88:
                        v18 += v18
                    }
                    _ = fi30
                }
            case "\n\t":
                for fi31 := 0; fi31 < 11; fi31++ {
                    switch ("hello" + "") {
                    case "a":
                        v20 += (v20 >> ((1) & 63))
                    case "\n\t":
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                    case "hello":
                        switch fi31 {
                        case 128:
                            switch v1 {
                            case "":
                                if v2 {
                                    fmt.Println(fi31, v3)
                                    switch v1 {
                                    case "foo bar":
                                        var v32 int16 = -1
                                        v17 -= v17
                                        var v33 string = ("\n\t" + v1)
                                        _ = v32
                                        _ = v33
                                    case "":
                                        v16 = v20
                                    case "a":
                                        if v2 {
                                            var v34 int = int(v17)
                                            _ = v34
                                        } else {
                                            v35 := GenS2{F0: 963, F1: 100, F2: 0}
                                            v17 = (^(v17))
                                            switch v1 {
                                            case "a":
                                                v2 = (100 != genFn1())
                                            case "\n\t":
                                                v17 += 128
                                            case "foo bar":
                                                var v36 string = (v1 + "")
                                                _ = v36
                                            }
                                            var v37 uint16 = v16
                                            if (!(1.5 >= 3.14159)) {
                                                var v38 float64 = 1e10
                                                v19 = append(v19, fi21)
                                                _ = v38
                                            }
                                            _ = v35
                                            _ = v37
                                        }
                                    }
                                    v39 := GenS3{F0: fi31, F1: v3}
                                    _ = v39
                                }
                            default:
                                v40 := GenS1{F0: 0}
                                _ = v40
                            }
                        case 9223372036854775807:
                            switch (fi21 >> ((v17) & 63)) {
                            case 4294967295:
                                v1 = v1
                            default:
                                fmt.Println(v17)
                                fmt.Println(v17, v3, v18)
                                var v41 int = fi31
                                var v42 int = fi21
                                _ = len(v19)
                                if (!v2) {
                                    _ = len(v19)
                                    fmt.Println(v41)
                                } else {
                                    var v43 float64 = (1.5 - (3.14159 - 1.5))
                                    var v44 int = (v42 >> (((^(v17))) & 63))
                                    _ = v43
                                    _ = v44
                                }
                                if len(v19) > 0 {
                                    v19[0] = (v42 / (((v42 / ((fi31) | 1))) | 1))
                                }
                                _ = v41
                                _ = v42
                            }
                        default:
                            v17 = v17
                        }
                    }
                    if len(v19) > 0 {
                        v19[0] = fi31
                    }
                    var v45 string = v1
                    v18 += 255
                    if (int16(v18) < 1) {
                        v46 := GenS4{F0: false, F1: 268}
                        v47 := GenS2{F0: 32767, F1: -12, F2: v16}
                        _ = v46
                        _ = v47
                    }
                    _ = fi31
                    _ = v45
                }
            }
            _ = fi21
        }
    case "":
        v48 := GenS5{F0: (v16 ^ 255), F1: 10, F2: v1}
        _ = v48
    case "a":
        var v49 uint8 = (v18 ^ v18)
        _ = v49
    default:
        if v2 {
            for fi50 := 0; fi50 < 7; fi50++ {
                var v51 uint8 = 0
                _ = fi50
                _ = v51
            }
            var v52 uint16 = (v16 << (((^(v17))) & 63))
            _ = v52
        } else {
            var v53 int64 = -9223372036854775808
            _ = v53
        }
    }
    v54 := GenS1{F0: 127}
    var v55 uint64 = 42
    if len(v19) > 0 {
        v19[0] = 32767
    }
    var v56 float64 = -2.25
    for fi57 := 0; fi57 < 12; fi57++ {
        if true {
            if v2 {
                if (0 > 0) {
                    fmt.Println(v3, v16)
                } else {
                    var v58 string = (("\n\t" + v1) + "a")
                    _ = v58
                }
                var v59 uint8 = 1
                _ = v59
            }
            var v60 uint64 = (^(v55))
            for fi61 := 0; fi61 < 10; fi61++ {
                v62 := GenS6{F0: -1, F1: v2}
                var v63 uint16 = v16
                _ = fi61
                _ = v62
                _ = v63
            }
            _ = v60
        }
        _ = fi57
    }
    var v64 uint8 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v54
    _ = v55
    _ = v56
    _ = v64
}
