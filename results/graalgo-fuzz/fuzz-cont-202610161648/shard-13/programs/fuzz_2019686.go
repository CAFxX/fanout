package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 bool; F1 int8 }
type GenS9 struct { F0 uint16; F1 uint8; F2 int8 }
func genFn0() int64 {
    for fi4 := 0; fi4 < 7; fi4++ {
        if false {
            var v5 bool = true
            v6 := []int{9223372036854775807, fi4, (fi4 - 128)}
            for fi7 := 0; fi7 < 2; fi7++ {
                fmt.Println(fi7)
                _ = fi7
            }
            _ = v5
            _ = v6
        } else {
            var v8 int = fi4
            if true {
                v8 = fi4
            }
            var v9 uint64 = 42
            _ = v8
            _ = v9
        }
        var v10 uint64 = 2
        v11 := GenS1{F0: 32767}
        _ = fi4
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 uint8 = 245
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 string = "a"
            switch (fi14 % ((1) | 1)) {
            case 4294967295:
                var v16 int8 = 0
                _ = v16
            case 42:
                var v17 uint8 = (v13 - 128)
                _ = v17
            default:
                if ((-(fi14)) >= (^(fi12))) {
                    v18 := GenS1{F0: 32767}
                    _ = v18
                }
            }
            var v19 int32 = 0
            _ = fi14
            _ = v15
            _ = v19
        }
        _ = fi12
        _ = v13
    }
    return 1
}

func main() {
    var v1 int8 = 0
    var v2 float32 = (1e10 * 0.0)
    v3 := make(map[int]string)
    var v20 bool = (false && ((-474 >= 32) || false))
    v21 := GenS2{F0: 1, F1: 688}
    for fi22 := 0; fi22 < 3; fi22++ {
        fmt.Println(fi22)
        _ = fi22
    }
    fmt.Println(v1, v20, v1)
    for fi23 := 0; fi23 < 8; fi23++ {
        for fi24 := 0; fi24 < 3; fi24++ {
            switch ((("hello" + ("foo bar" + "")) + "hello") + "hello") {
            case "a":
                delete(v3, fi23)
            case "":
                fmt.Println(fi24, fi24)
                v25 := GenS3{F0: 4294967295, F1: "foo bar"}
                v2 += v2
                _ = v25
            }
            delete(v3, (fi23 << ((2) & 63)))
            if false {
                v2 -= v2
                var v26 uint8 = 2
                _ = v26
            } else {
                v3[fi23] = ""
                v27 := GenS4{F0: -769}
                var v28 int32 = 42
                _ = v27
                _ = v28
            }
            _ = fi24
        }
        fmt.Println(v1, v2)
        fmt.Println(v20, fi23)
        _ = fi23
    }
    v2 = float32(float64(float32(3.14159)))
    for fi29 := 0; fi29 < 17; fi29++ {
        switch ("hello" + "hello") {
        case "\n\t":
            switch fi29 {
            case 213:
                v30 := []int{(fi29 ^ 42), 977, (fi29 >> ((713) & 63))}
                _ = v30
            case -202:
                fmt.Println(v2)
            default:
                if v20 {
                    fmt.Println(v20, v2)
                } else {
                    var v31 uint16 = 527
                    _ = v31
                }
            }
        case "hello":
            var v32 uint = 0
            _ = v32
        case "":
            switch "hello" {
            case "\n\t":
                if v20 {
                    delete(v3, 0)
                    if (-664 < -19) {
                        var v33 int = (fi29 & 127)
                        var v34 uint8 = 1
                        _ = v33
                        _ = v34
                    } else {
                        v20 = (!false)
                        var v35 int64 = 306
                        _ = v35
                    }
                    v36 := GenS5{F0: 204, F1: 4294967295}
                    _ = v36
                } else {
                    v20 = (2147483648 != 1)
                }
            case "foo bar":
                if (("foo bar" + "") < "hello") {
                    for fi37 := 0; fi37 < 13; fi37++ {
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(v2, v1, v20)
                        default:
                            switch (^(fi29)) {
                            case 4294967295:
                                if (301 >= 2) {
                                    v2 -= ((float32((float64((v2 - -2.25)) - 0.1)) * ((v2 * v2) / v2)) * float32(3.14159))
                                    v38 := GenS3{F0: 255, F1: "a"}
                                    _ = v38
                                } else {
                                    v39 := []int{1, fi29, fi29}
                                    _ = v39
                                }
                            case 127:
                                var v40 int = 486
                                _ = v40
                            case 906:
                                var v41 uint32 = 1
                                _ = v41
                            }
                        }
                        fmt.Println(fi29)
                        v1 += (v1 % (((v1 - (-(v1)))) | 1))
                        fmt.Println(v1)
                        v42 := []int{276}
                        _ = fi37
                        _ = v42
                    }
                    var v43 uint16 = 531
                    _ = v43
                }
            case "hello":
                if (-1 < -395) {
                    v44 := GenS2{F0: 0, F1: 2147483647}
                    fmt.Println(v20, fi29)
                    _ = v44
                } else {
                    fmt.Println(v1, v20)
                    for fi45 := 0; fi45 < 11; fi45++ {
                        var v46 uint32 = 65535
                        if v20 {
                            var v47 uint8 = 42
                            if true {
                                _ = len(v3)
                                if (v2 > v2) {
                                    switch fi29 {
                                    case -529:
                                        v48 := GenS6{F0: (^(fi45)), F1: 835}
                                        _ = v48
                                    }
                                } else {
                                    v49 := GenS7{F0: 451, F1: 392}
                                    v50 := []int{(^(fi45)), 812}
                                    _ = v49
                                    _ = v50
                                }
                                var v51 int = (fi29 * (fi45 - fi45))
                                var v52 bool = (uint(v46) < 562)
                                _ = v51
                                _ = v52
                            }
                            if (1.5 != 1.5) {
                                v47 = 127
                                fmt.Println(v1)
                            } else {
                                v53 := []int{(fi45 << ((680) & 63)), 824, fi45}
                                _ = v53
                            }
                            _ = v47
                        } else {
                            var v54 float32 = float32((0.0 + 0.1))
                            _ = len(v3)
                            v55 := GenS1{F0: 129}
                            _ = v54
                            _ = v55
                        }
                        _ = fi45
                        _ = v46
                    }
                }
            default:
                v56 := GenS7{F0: 1, F1: 753}
                _ = v56
            }
        }
        for fi57 := 0; fi57 < 14; fi57++ {
            v58 := GenS8{F0: ("" <= ""), F1: (v1 - v1)}
            v1 -= (v1 - v1)
            fmt.Println(v20)
            _ = fi57
            _ = v58
        }
        _ = fi29
    }
    var v59 uint32 = 474
    for fi60 := 0; fi60 < 3; fi60++ {
        var v61 int32 = 616
        _ = fi60
        _ = v61
    }
    if ((!v20) && (!(!(v20 || ((genFn0() <= 0) && (!(!(833 == 996)))))))) {
        var v62 uint16 = 65535
        var v63 uint16 = (v62 / ((uint16(v59)) | 1))
        for fi64 := 0; fi64 < 18; fi64++ {
            v63 -= 42
            var v65 float32 = (float32(1.5) * v2)
            _ = fi64
            _ = v65
        }
        _ = v62
        _ = v63
    } else {
        v66 := GenS9{F0: 65535, F1: 42, F2: v1}
        _ = v66
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v59
}
