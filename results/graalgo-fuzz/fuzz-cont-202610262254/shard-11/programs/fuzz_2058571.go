package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int64; F1 uint; F2 uint }
type GenS5 struct { F0 int32; F1 bool; F2 int }
type GenS6 struct { F0 int64; F1 uint }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 string = (("a" + "hello") + "\n\t")
    v3 := make(map[uint32]bool)
    var v4 uint64 = 488
    var v5 int32 = -391
    delete(v3, 2147483647)
    if true {
        v6 := GenS1{F0: -511}
        if true {
            for fi7 := 0; fi7 < 13; fi7++ {
                var v8 uint32 = 2147483647
                v8 = 4294967295
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 11; fi9++ {
                v4 -= 18446744073709551615
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 int32 = v5
                    _ = fi10
                    _ = v11
                }
                _ = fi9
            }
            v5 += v5
        } else {
            switch ("\n\t" + "\n\t") {
            case "foo bar":
                v2 = v2
            }
            var v12 uint = 18446744073709551615
            _ = v12
        }
        _ = v6
    }
    v5 = (v5 & -1)
    if (((!(4294967295 < 42)) && (!((763 != 32767) && (1 >= 127)))) && ((1.5 >= -2.25) || true)) {
        var v13 string = "\n\t"
        v14 := GenS1{F0: 292}
        _ = v13
        _ = v14
    } else {
        if false {
            for fi15 := 0; fi15 < 6; fi15++ {
                var v16 uint64 = (v4 & v4)
                var v17 uint8 = 255
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        if true {
            var v18 uint64 = v4
            for fi19 := 0; fi19 < 16; fi19++ {
                var v20 int64 = 9223372036854775807
                for fi21 := 0; fi21 < 17; fi21++ {
                    v22 := GenS1{F0: 1}
                    v23 := []int{fi19}
                    var v24 float32 = -2.25
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
                _ = fi19
                _ = v20
            }
            _ = v18
        }
    }
    v25 := GenS2{F0: (^(v5)), F1: 2, F2: 610}
    v26 := GenS1{F0: -248}
    var v27 uint64 = (v4 & (^(v4)))
    switch -9223372036854775808 {
    case 127:
        v28 := []int{-661}
        _ = v28
    case -699:
        var v29 float64 = (0.1 + 3.14159)
        _ = v29
    }
    var v30 uint64 = v27
    switch v2 {
    case "":
        var v31 bool = false
        _ = v31
    case "hello":
        v32 := GenS1{F0: 607}
        _ = v32
    default:
        if (true || true) {
            if (!true) {
                for fi33 := 0; fi33 < 9; fi33++ {
                    v34 := GenS2{F0: 11, F1: 331, F2: 432}
                    if true {
                        v30 -= (^(v4))
                        v35 := GenS3{F0: 0, F1: 42, F2: (v27 / ((v4) | 1))}
                        _ = v35
                    } else {
                        v36 := GenS4{F0: -129, F1: 42, F2: 255}
                        v30 += v30
                        _ = v36
                    }
                    if (!((true || false) || true)) {
                        v2 = v2
                    } else {
                        for fi37 := 0; fi37 < 11; fi37++ {
                            v38 := GenS5{F0: (v5 % ((v5) | 1)), F1: false, F2: (fi37 / ((42) | 1))}
                            var v39 int = 4294967295
                            if (false && false) {
                                var v40 uint16 = 0
                                v41 := []int{fi37}
                                _ = v40
                                _ = v41
                            } else {
                                v27 += (v30 + 0)
                                v30 += 2
                            }
                            v42 := GenS4{F0: -9223372036854775808, F1: 1, F2: 255}
                            _ = fi37
                            _ = v38
                            _ = v39
                            _ = v42
                        }
                        fmt.Println(fi33, v30, v2)
                    }
                    _ = fi33
                    _ = v34
                }
                v27 = 18446744073709551615
            } else {
                for fi43 := 0; fi43 < 16; fi43++ {
                    var v44 uint8 = 202
                    v45 := []int{776, (fi43 | 32767)}
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                v46 := GenS6{F0: 9223372036854775807, F1: 1}
                _ = v46
            }
            switch (v2 + v2) {
            case "\n\t":
                var v47 int16 = 1000
                _ = v47
            case "foo bar":
                if false {
                    for fi48 := 0; fi48 < 13; fi48++ {
                        if true {
                            if true {
                                for fi49 := 0; fi49 < 14; fi49++ {
                                    v50 := GenS3{F0: 979, F1: 0, F2: v27}
                                    _ = fi49
                                    _ = v50
                                }
                                v51 := v1[68]
                                _ = v51
                            } else {
                                var v52 float32 = 3.14159
                                _ = v52
                            }
                            fmt.Println(v4, fi48, v30)
                        } else {
                            for fi53 := 0; fi53 < 2; fi53++ {
                                v54 := GenS3{F0: 42, F1: 106, F2: v30}
                                _ = fi53
                                _ = v54
                            }
                            v55 := GenS3{F0: 1, F1: 42, F2: (^(v30))}
                            _ = v55
                        }
                        _ = fi48
                    }
                    v56 := GenS7{F0: 1}
                    _ = v56
                }
            default:
                for fi57 := 0; fi57 < 12; fi57++ {
                    v58 := []int{}
                    v59 := GenS7{F0: 1}
                    _ = fi57
                    _ = v58
                    _ = v59
                }
            }
            var v60 int16 = 1000
            _ = v60
        } else {
            fmt.Println(v5, v2, v27)
            v61 := GenS4{F0: -867, F1: 856, F2: 72}
            _ = v61
        }
    }
    for fi62 := 0; fi62 < 9; fi62++ {
        if true {
            v63 := GenS8{F0: fi62}
            if (979 > 42) {
                for fi64 := 0; fi64 < 16; fi64++ {
                    var v65 uint8 = 1
                    _ = fi64
                    _ = v65
                }
            } else {
                v66 := GenS9{F0: 2}
                v67 := GenS6{F0: 790, F1: 4294967295}
                var v68 int32 = (v5 ^ v5)
                _ = v66
                _ = v67
                _ = v68
            }
            v69 := GenS5{F0: v5, F1: false, F2: fi62}
            _ = v63
            _ = v69
        } else {
            v2 = v2
            for fi70 := 0; fi70 < 13; fi70++ {
                switch fi62 {
                case -156:
                    fmt.Println(fi62, v5)
                }
                _ = fi70
            }
            v71 := GenS1{F0: -968}
            _ = v71
        }
        _ = fi62
    }
    fmt.Println(v27, v30, v5)
    for fi72 := 0; fi72 < 1; fi72++ {
        v73 := []int{fi72, 4294967295, (fi72 << ((436) & 63))}
        _ = fi72
        _ = v73
    }
    for fi74 := 0; fi74 < 17; fi74++ {
        var v75 int32 = v5
        for fi76 := 0; fi76 < 10; fi76++ {
            var v77 string = v2
            v78 := GenS9{F0: 65535}
            _ = fi76
            _ = v77
            _ = v78
        }
        _ = fi74
        _ = v75
    }
    v27 = v30
    var v79 uint64 = (v30 % (((v4 - 4294967295)) | 1))
    var v80 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v79
    _ = v80
}
