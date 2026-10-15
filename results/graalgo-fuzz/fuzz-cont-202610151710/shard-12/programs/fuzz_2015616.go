package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int8; F1 int8; F2 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool; F1 uint64; F2 uint8 }

func main() {
    var v1 float32 = float32((0.1 - (((3.14159 / 1.0) - float64(1.5)) / (float64(1e10) * (float64(-2.25) * ((float64((-2.25 + 0.1)) / -2.25) + -2.25))))))
    var v2 uint32 = 4294967295
    v3 := make(map[int]int)
    var v4 string = ("" + "\n\t")
    for fi5 := 0; fi5 < 6; fi5++ {
        for fi6 := 0; fi6 < 5; fi6++ {
            v7 := v3[fi6]
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint16 = 42
            v10 := GenS1{F0: (((v9 / ((v9) | 1)) == 32767) || true)}
            v9 = (v9 - 32767)
            if (166 <= 1) {
                var v11 int = -9223372036854775808
                v9 = v9
                _ = v11
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi5
    }
    v12 := v3[636]
    if true {
        var v13 int64 = 138
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 int8 = 16
            for fi16 := 0; fi16 < 18; fi16++ {
                if true {
                    v4 = "a"
                }
                _ = fi16
            }
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    var v17 uint32 = (v2 - v2)
    for fi18 := 0; fi18 < 20; fi18++ {
        fmt.Println(v12, v12, v12)
        v19 := GenS1{F0: true}
        for fi20 := 0; fi20 < 13; fi20++ {
            if (false || false) {
                var v21 int8 = 127
                v1 += float32(float64((v1 / v1)))
                fmt.Println(v2, v1, v2)
                var v22 int = fi18
                var v23 uint = 42
                switch 4294967295 {
                case 42:
                    var v24 uint32 = v2
                    _ = v24
                }
                if (0 == -32768) {
                    var v25 float32 = (v1 + 3.14159)
                    _ = v25
                } else {
                    v26 := []int{v12}
                    var v27 uint8 = 67
                    v28 := GenS1{F0: false}
                    _ = v26
                    _ = v27
                    _ = v28
                }
                var v29 uint8 = 128
                _ = v21
                _ = v22
                _ = v23
                _ = v29
            } else {
                if true {
                    v1 += 3.14159
                    _ = len(v3)
                } else {
                    v4 = "foo bar"
                    var v30 float32 = v1
                    if ((fi20 ^ fi18) == 1) {
                        fmt.Println(v1, v2, v1)
                        if false {
                            var v31 uint16 = 65535
                            _ = v31
                        }
                    } else {
                        v2 += v17
                        if (4294967295 > 4294967295) {
                            var v32 int32 = 0
                            fmt.Println(v4)
                            _ = v32
                        }
                        if (!((v4 + "") >= v4)) {
                            v33 := []int{}
                            fmt.Println(fi20)
                            var v34 uint8 = 123
                            _ = v33
                            _ = v34
                        } else {
                            v3[-532] = fi20
                            v35 := v3[fi18]
                            if (65535 != v17) {
                                var v36 uint8 = 4
                                fmt.Println(v2, fi20, v30)
                                v37 := GenS1{F0: (-2147483648 != 2147483647)}
                                _ = v36
                                _ = v37
                            } else {
                                var v38 float32 = (1.5 / ((v1 * (v1 / 0.1)) * (float32(3.14159) - (((1e10 * 0.0) - (float32(-2.25) - v1)) / 1.0))))
                                if true {
                                    v17 = 0
                                    if (-2.25 > 3.14159) {
                                        fmt.Println(v12, v38, fi20)
                                        fmt.Println(fi20)
                                        v1 = ((v38 + v30) - v38)
                                        if true {
                                            v38 -= ((v1 + ((v38 * 1.5) / ((((v30 / (v1 + (v30 + 1e10))) + (((v1 * v30) * v38) - 0.0)) + 0.0) - float32(1e10)))) * 3.14159)
                                            var v39 uint64 = 4294967295
                                            _ = v39
                                        }
                                    } else {
                                        if false {
                                            v40 := GenS1{F0: true}
                                            var v41 int8 = -1
                                            _ = v40
                                            _ = v41
                                        } else {
                                            if (0.0 != (1.5 - ((float64(1.5) * (1e10 / ((float64((float32(1e10) / v38)) * 1e10) * ((0.1 - 0.0) - 0.1)))) / 1e10))) {
                                                if (0.0 != -2.25) {
                                                    var v42 uint16 = 2
                                                    v43 := GenS2{F0: false, F1: -94}
                                                    _ = v42
                                                    _ = v43
                                                }
                                            }
                                            v44 := []int{v35}
                                            v17 -= (v17 * v17)
                                            _ = v44
                                        }
                                        v35 -= fi20
                                    }
                                }
                                v30 -= v38
                                _ = v38
                            }
                            _ = v35
                        }
                    }
                    _ = v30
                }
                var v45 float32 = v1
                v17 = 292
                _ = v45
            }
            if ((v2 << ((2) & 63)) < (v17 ^ (v17 << ((65535) & 63)))) {
                v46 := GenS3{F0: 127, F1: 0, F2: (fi18 / (((v12 | (fi20 % ((4294967295) | 1)))) | 1))}
                var v47 int8 = 100
                var v48 int8 = 100
                _ = v46
                _ = v47
                _ = v48
            } else {
                if (!true) {
                    v49 := GenS4{F0: 533}
                    _ = v49
                } else {
                    if false {
                        fmt.Println(v17, fi20, v17)
                        v2 = (v17 | v2)
                        var v50 uint16 = 65535
                        var v51 uint32 = (v2 & v17)
                        var v52 uint32 = (^(v17))
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                }
                if (!(!true)) {
                    if (((fi20 >> ((4294967295) & 63)) < (v12 ^ 42)) && (386 > (^(v2)))) {
                        fmt.Println(v1, v17, v2)
                        if false {
                            v53 := GenS5{F0: 225}
                            v54 := GenS3{F0: -1, F1: 107, F2: 42}
                            if ((1e10 / 1.5) == 1.5) {
                                var v55 int32 = 42
                                fmt.Println(v2, fi18, v17)
                                _ = v55
                            } else {
                                v56 := []int{fi20, (fi20 % ((-2147483648) | 1))}
                                _ = v56
                            }
                            _ = v53
                            _ = v54
                        } else {
                            v17 += (v2 * (^(v17)))
                        }
                        fmt.Println(v17, v17, v1)
                    }
                    v4 = (v4 + v4)
                    fmt.Println(v17, v17)
                } else {
                    v57 := GenS5{F0: 42}
                    _ = v57
                }
                var v58 int16 = 42
                _ = v58
            }
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v12, v1, v2)
    var v59 uint = 18446744073709551615
    var v60 int64 = -9223372036854775808
    switch (v12 ^ v12) {
    case -2147483648:
        for fi61 := 0; fi61 < 13; fi61++ {
            var v62 float32 = (3.14159 * ((0.0 - v1) - (v1 / 1.5)))
            if true {
                switch fi61 {
                case 32767:
                    var v63 uint16 = 2
                    _ = v63
                case -106:
                    fmt.Println(v4)
                case 128:
                    for fi64 := 0; fi64 < 17; fi64++ {
                        var v65 uint32 = v17
                        v66 := []int{128, -266}
                        fmt.Println(v12)
                        _ = fi64
                        _ = v65
                        _ = v66
                    }
                }
            }
            for fi67 := 0; fi67 < 16; fi67++ {
                v2 -= 964
                _ = fi67
            }
            _ = fi61
            _ = v62
        }
    case 321:
        if false {
            switch v12 {
            case 128:
                v68 := GenS3{F0: -35, F1: int8(v2), F2: v12}
                _ = v68
            }
            v4 = v4
        }
    case -70:
        var v69 bool = true
        _ = v69
    }
    v59 += 423
    v70 := []int{}
    var v71 uint32 = v17
    fmt.Println(v2, v2, v59)
    for fi72 := 0; fi72 < 20; fi72++ {
        if (42 <= 1) {
            for fi73 := 0; fi73 < 4; fi73++ {
                v60 += int64(fi72)
                v17 += v17
                _ = fi73
            }
        }
        v74 := GenS2{F0: (v4 > v4), F1: 0}
        var v75 uint8 = 2
        _ = fi72
        _ = v74
        _ = v75
    }
    _ = len(v3)
    v76 := GenS6{F0: (v17 >= v17), F1: 0, F2: 2}
    v17 -= 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v70:", len(v70), v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v59
    _ = v60
    _ = v70
    _ = v71
    _ = v76
}
