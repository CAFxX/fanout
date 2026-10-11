package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 uint; F2 int16 }
type GenS4 struct { F0 uint32; F1 int64; F2 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8 }

func main() {
    var v1 uint8 = 1
    var v2 float64 = -2.25
    var v3 string = ("hello" + (("hello" + "foo bar") + "a"))
    for fi4 := 0; fi4 < 12; fi4++ {
        v5 := GenS1{F0: 42, F1: (!(!false)), F2: 1}
        fmt.Println(v2, v2, fi4)
        v6 := GenS2{F0: (fi4 >> ((0) & 63))}
        _ = fi4
        _ = v5
        _ = v6
    }
    fmt.Println(v3, v2, v1)
    fmt.Println(v1, v2, v1)
    for fi7 := 0; fi7 < 11; fi7++ {
        v2 = 1e10
        v8 := []int{fi7, (fi7 << ((243) & 63)), (fi7 % ((fi7) | 1)), fi7}
        _ = fi7
        _ = v8
    }
    v9 := []int{}
    var v10 float32 = -2.25
    var v11 uint8 = 89
    _ = len(v9)
    fmt.Println(v1, v1, v2)
    _ = len(v9)
    for fi12 := 0; fi12 < 14; fi12++ {
        switch (fi12 + 32767) {
        case 127:
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
        _ = len(v9)
        v13 := GenS3{F0: -9223372036854775808, F1: 16, F2: 32767}
        _ = fi12
        _ = v13
    }
    var v14 int8 = -15
    v3 = v3
    v15 := GenS1{F0: -700, F1: false, F2: 1}
    switch -9223372036854775808 {
    case 239:
        var v16 string = v3
        _ = v16
    case 437:
        if false {
            v17 := GenS1{F0: 51, F1: ("foo bar" > (v3 + v3)), F2: 42}
            v1 = v11
            _ = v17
        } else {
            if false {
                for fi18 := 0; fi18 < 8; fi18++ {
                    var v19 int16 = -32768
                    if len(v9) > 0 {
                        v9[0] = fi18
                    }
                    var v20 bool = true
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                for fi21 := 0; fi21 < 10; fi21++ {
                    v10 -= float32(v2)
                    if ((true && (4294967295 != 1)) && true) {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    }
                    _ = fi21
                }
            }
            if false {
                if false {
                    var v22 int8 = -80
                    _ = v22
                } else {
                    v10 -= 1e10
                    fmt.Println(v10)
                }
                for fi23 := 0; fi23 < 19; fi23++ {
                    var v24 uint32 = 2147483648
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    _ = fi23
                    _ = v24
                }
                var v25 uint32 = 2147483648
                _ = v25
            } else {
                var v26 float64 = (-2.25 * v2)
                var v27 int32 = 0
                switch (("\n\t" + (v3 + v3)) + ("\n\t" + "\n\t")) {
                case "\n\t":
                    v11 -= 159
                case "":
                    v2 = (0.0 * float64(v10))
                case "a":
                    v2 -= (v26 / 0.1)
                default:
                    v26 = v2
                }
                _ = v26
                _ = v27
            }
            for fi28 := 0; fi28 < 19; fi28++ {
                if (-412 != 1) {
                    var v29 int32 = 806
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    if (v29 < (v29 >> ((144) & 63))) {
                        fmt.Println(v3, v1, v10)
                    } else {
                        v30 := GenS3{F0: -9223372036854775808, F1: 65535, F2: 1}
                        _ = v30
                    }
                    _ = v29
                } else {
                    for fi31 := 0; fi31 < 5; fi31++ {
                        if len(v9) > 0 {
                            v9[0] = 0
                        }
                        _ = fi31
                    }
                }
                fmt.Println(v2, fi28, fi28)
                for fi32 := 0; fi32 < 14; fi32++ {
                    v33 := GenS2{F0: (fi32 / ((fi32) | 1))}
                    switch ((v3 + "\n\t") + v3) {
                    case "a":
                        if false {
                            var v34 uint32 = 42
                            _ = v34
                        }
                    case "foo bar":
                        v35 := GenS4{F0: 42, F1: -725, F2: 0}
                        _ = v35
                    case "hello":
                        switch v3 {
                        case "":
                            if len(v9) > 0 {
                                v9[0] = fi32
                            }
                        default:
                            var v36 int16 = 886
                            _ = v36
                        }
                    }
                    var v37 uint32 = 0
                    _ = fi32
                    _ = v33
                    _ = v37
                }
                _ = fi28
            }
        }
    case 128:
        if (true || false) {
            var v38 int8 = (v14 + v14)
            _ = v38
        } else {
            if false {
                fmt.Println(v11)
                var v39 uint = 480
                if (!(0 < 0)) {
                    if len(v9) > 0 {
                        v9[0] = 0
                    }
                    fmt.Println(v2, v3)
                } else {
                    _ = len(v9)
                    var v40 int64 = 42
                    _ = v40
                }
                _ = v39
            } else {
                var v41 float32 = 1e10
                for fi42 := 0; fi42 < 10; fi42++ {
                    if (3.14159 != ((0.0 / (0.1 / (1e10 + 0.0))) / float32((v2 + 1.5)))) {
                        fmt.Println(v11, v14)
                        var v43 uint8 = v11
                        _ = v43
                    } else {
                        v10 -= 3.14159
                        for fi44 := 0; fi44 < 9; fi44++ {
                            var v45 int = 127
                            var v46 uint16 = 42
                            fmt.Println(fi42, v14, v1)
                            v47 := GenS1{F0: 1, F1: (9223372036854775807 > -1), F2: 42}
                            _ = fi44
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                        if false {
                            _ = len(v9)
                            switch fi42 {
                            case 0:
                                v11 = (^(v11))
                            case 9223372036854775807:
                                v48 := GenS5{F0: 18446744073709551615}
                                _ = v48
                            default:
                                fmt.Println(v3)
                            }
                            v49 := GenS1{F0: 2147483647, F1: true, F2: 9223372036854775807}
                            _ = v49
                        } else {
                            for fi50 := 0; fi50 < 19; fi50++ {
                                if true {
                                    switch (fi42 + 1) {
                                    case 852:
                                        var v51 float32 = v41
                                        _ = v51
                                    case -1:
                                        var v52 int32 = int32(v11)
                                        _ = v52
                                    case 965:
                                        if (435 != 0) {
                                            var v53 int32 = 0
                                            _ = v53
                                        }
                                    }
                                    if (1.5 == 1.5) {
                                        var v54 int16 = 1
                                        _ = v54
                                    }
                                }
                                _ = fi50
                            }
                            v55 := GenS6{F0: v14}
                            if len(v9) > 0 {
                                v9[0] = fi42
                            }
                            _ = v55
                        }
                    }
                    var v56 bool = true
                    var v57 string = (v3 + v3)
                    _ = fi42
                    _ = v56
                    _ = v57
                }
                for fi58 := 0; fi58 < 19; fi58++ {
                    v9 = append(v9, fi58)
                    _ = fi58
                }
                _ = v41
            }
        }
    }
    v14 += (^(v14))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v9
}
