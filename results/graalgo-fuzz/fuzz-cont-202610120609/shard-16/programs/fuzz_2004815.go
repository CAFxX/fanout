package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 string }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 uint32; F2 int32 }

func main() {
    var v1 uint32 = 536
    var v2 uint64 = 0
    var v3 uint = 354
    v2 = (v2 & v2)
    v4 := GenS1{F0: -128, F1: 254, F2: ("" + "a")}
    var v5 uint32 = 2147483647
    v3 = 65535
    var v6 float32 = 0.1
    v7 := GenS2{F0: -821, F1: 4294967295}
    switch 127 {
    case -2147483648:
        v2 += (v2 & (v2 / ((v2) | 1)))
    case -9223372036854775808:
        v3 = (^(v3))
    default:
        switch "\n\t" {
        case "\n\t":
            if true {
                v1 = 65535
                v8 := GenS1{F0: 91, F1: 42, F2: "a"}
                _ = v8
            } else {
                var v9 uint16 = 42
                switch 817 {
                case 128:
                    v10 := []int{42, 4294967295, 127, -1}
                    _ = v10
                }
                _ = v9
            }
        case "":
            if (((!((100 < -1) && (-123 <= -128))) || false) && false) {
                var v11 bool = (!false)
                for fi12 := 0; fi12 < 5; fi12++ {
                    if (("hello" != "\n\t") || (v11 || v11)) {
                        v13 := GenS2{F0: 1000, F1: 413}
                        if ((1e10 - 3.14159) == (1e10 + 0.1)) {
                            if v11 {
                                v14 := GenS1{F0: 1, F1: 65535, F2: "foo bar"}
                                var v15 bool = v11
                                v15 = (!v11)
                                _ = v14
                                _ = v15
                            }
                            if false {
                                if (1 == 100) {
                                    var v16 int16 = -1
                                    v17 := GenS1{F0: 0, F1: 0, F2: "a"}
                                    _ = v16
                                    _ = v17
                                }
                                for fi18 := 0; fi18 < 19; fi18++ {
                                    v2 += 42
                                    v3 = (v3 % ((0) | 1))
                                    v3 += (v3 % ((42) | 1))
                                    _ = fi18
                                }
                            }
                            if v11 {
                                switch (fi12 / ((fi12) | 1)) {
                                case 128:
                                    var v19 uint64 = v2
                                    _ = v19
                                case -2147483648:
                                    if (!v11) {
                                        for fi20 := 0; fi20 < 14; fi20++ {
                                            v6 = (float32((0.0 / 1.0)) + v6)
                                            v21 := GenS2{F0: 0, F1: (^(fi20))}
                                            v22 := GenS1{F0: 42, F1: uint16(v5), F2: (("hello" + "hello") + "a")}
                                            _ = fi20
                                            _ = v21
                                            _ = v22
                                        }
                                        v1 = v5
                                    }
                                case -592:
                                    if false {
                                        for fi23 := 0; fi23 < 13; fi23++ {
                                            v1 -= v1
                                            v6 -= v6
                                            _ = fi23
                                        }
                                    }
                                }
                                var v24 uint = (v3 / ((2) | 1))
                                _ = v24
                            } else {
                                if v11 {
                                    v11 = true
                                    if v11 {
                                        v25 := []int{1, fi12, -289, 32767}
                                        var v26 float32 = (0.1 * 1e10)
                                        var v27 int32 = 42
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    }
                                } else {
                                    v28 := []int{888}
                                    v3 = v3
                                    _ = v28
                                }
                                if false {
                                    for fi29 := 0; fi29 < 3; fi29++ {
                                        if (v6 > v6) {
                                            v30 := GenS2{F0: -1, F1: fi29}
                                            var v31 int8 = 0
                                            _ = v30
                                            _ = v31
                                        }
                                        var v32 uint8 = 138
                                        var v33 int = fi29
                                        _ = fi29
                                        _ = v32
                                        _ = v33
                                    }
                                } else {
                                    var v34 float64 = 1.5
                                    _ = v34
                                }
                                var v35 uint64 = 152
                                _ = v35
                            }
                        } else {
                            var v36 int64 = 0
                            _ = v36
                        }
                        _ = v13
                    }
                    _ = fi12
                }
                v37 := GenS1{F0: 0, F1: 506, F2: "a"}
                _ = v11
                _ = v37
            } else {
                if (3.14159 >= 0.1) {
                    var v38 uint32 = (v5 << ((194) & 63))
                    v39 := []int{}
                    _ = v38
                    _ = v39
                } else {
                    for fi40 := 0; fi40 < 15; fi40++ {
                        v41 := GenS1{F0: -128, F1: 42, F2: ("\n\t" + ("a" + ("" + "")))}
                        for fi42 := 0; fi42 < 5; fi42++ {
                            var v43 int = (fi40 + 0)
                            _ = fi42
                            _ = v43
                        }
                        _ = fi40
                        _ = v41
                    }
                    var v44 int8 = 127
                    _ = v44
                }
            }
        case "foo bar":
            v45 := []int{int(v1), 42}
            _ = v45
        default:
            v46 := GenS3{F0: 2}
            _ = v46
        }
    }
    if (1.5 >= ((1.5 / 1.0) * -2.25)) {
        v47 := []int{}
        for fi48 := 0; fi48 < 11; fi48++ {
            fmt.Println(v6, fi48)
            for fi49 := 0; fi49 < 9; fi49++ {
                var v50 int = (fi49 / (((fi49 / ((0) | 1))) | 1))
                v51 := GenS1{F0: -111, F1: 661, F2: "\n\t"}
                _ = fi49
                _ = v50
                _ = v51
            }
            _ = fi48
        }
        v3 = 65535
        _ = v47
    }
    v1 = 792
    var v52 float32 = v6
    v53 := GenS1{F0: -1, F1: 32767, F2: "foo bar"}
    if (1 > 127) {
        var v54 int8 = 127
        var v55 bool = true
        _ = v54
        _ = v55
    }
    v3 = (v3 & (v3 + v3))
    var v56 string = "a"
    v57 := GenS4{F0: 236, F1: (v5 << ((v3) & 63)), F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v56
    _ = v57
    _ = v6
    _ = v7
}
