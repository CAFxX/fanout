package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int64 }
func genFn0(p0 int16) int16 {
    fmt.Println(p0)
    p0 += p0
    _ = p0
    return p0
}
func genFn1() int {
    var v5 int64 = 767
    if (((42 == 42) && false) && true) {
        v5 += v5
        fmt.Println(v5)
    } else {
        v5 += v5
        fmt.Println(v5)
        v6 := GenS1{F0: 0}
        _ = v6
    }
    v5 -= v5
    var v7 uint32 = 90
    v8 := GenS1{F0: -128}
    _ = v5
    _ = v7
    _ = v8
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -93
    v3 := make(map[uint32]int)
    var v4 string = ((("a" + ("hello" + ("\n\t" + ("hello" + "\n\t")))) + ("hello" + "\n\t")) + "foo bar")
    v2 += v2
    var v9 bool = (false && (100 != -128))
    if ((v9 || v9) || (!v9)) {
        v9 = v9
        if v9 {
            if ((v2 >> ((255) & 63)) >= (v2 / ((v2) | 1))) {
                var v10 float64 = -2.25
                v11 := GenS1{F0: 0}
                _ = v10
                _ = v11
            }
        }
        if (!v9) {
            v4 = ((v4 + "foo bar") + "")
        }
    }
    v9 = (68 > 164)
    switch genFn1() {
    case 128:
        v12 := GenS1{F0: 100}
        _ = v12
    case 127:
        v13 := GenS1{F0: 42}
        _ = v13
    case -2147483648:
        if (335 < 2147483647) {
            for fi14 := 0; fi14 < 2; fi14++ {
                switch (fi14 - genFn1()) {
                case 9223372036854775807:
                    v4 = v4
                default:
                    for fi15 := 0; fi15 < 20; fi15++ {
                        v16 := GenS1{F0: 42}
                        v17 := GenS1{F0: -128}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                }
                if true {
                    v18 := v3[1]
                    _ = v18
                } else {
                    var v19 int64 = 1
                    var v20 float64 = 0.1
                    _ = v19
                    _ = v20
                }
                for fi21 := 0; fi21 < 3; fi21++ {
                    switch genFn1() {
                    case -2147483648:
                        v2 += -326
                    case -624:
                        fmt.Println(v9, v2)
                    }
                    if v9 {
                        v22 := []int{fi21, 128}
                        _ = v22
                    } else {
                        if v9 {
                            switch "\n\t" {
                            case "foo bar":
                                v23 := []int{(fi21 * (fi14 - int(v2))), (fi21 / ((-1) | 1))}
                                _ = v23
                            case "":
                                v24 := GenS1{F0: int8(v2)}
                                _ = v24
                            case "\n\t":
                                var v25 uint32 = 2
                                _ = v25
                            default:
                                v26 := []int{}
                                _ = v26
                            }
                            if (!v9) {
                                switch fi14 {
                                case 127:
                                    fmt.Println(v4)
                                default:
                                    v27 := GenS2{F0: v2}
                                    _ = v27
                                }
                                var v28 uint16 = 42
                                v29 := GenS2{F0: int64(fi14)}
                                v2 = 9223372036854775807
                                _ = v28
                                _ = v29
                            } else {
                                v1[2] = 0
                                v30 := GenS2{F0: v2}
                                var v31 int16 = 42
                                _ = v30
                                _ = v31
                            }
                        } else {
                            var v32 bool = v9
                            _ = v32
                        }
                        v4 = "\n\t"
                    }
                    v2 -= v2
                    v33 := []int{(fi14 % ((genFn1()) | 1))}
                    _ = len(v33)
                    var v34 uint32 = 0
                    fmt.Println(v2, fi21)
                    _ = fi21
                    _ = v33
                    _ = v34
                }
                _ = fi14
            }
        } else {
            var v35 uint = 2
            v36 := GenS2{F0: (v2 & v2)}
            var v37 int16 = 1000
            _ = v35
            _ = v36
            _ = v37
        }
    }
    var v38 float64 = 3.14159
    fmt.Println(v4)
    for fi39 := 0; fi39 < 9; fi39++ {
        v40 := []int{fi39}
        _ = fi39
        _ = v40
    }
    switch v4 {
    case "a":
        if (v2 <= v2) {
            v41 := GenS3{F0: 4294967295, F1: (v2 & v2)}
            for fi42 := 0; fi42 < 1; fi42++ {
                v38 -= (float64((1e10 * 3.14159)) / (0.0 - v38))
                var v43 uint16 = 42
                fmt.Println(fi42)
                _ = fi42
                _ = v43
            }
            _ = v41
        }
    case "foo bar":
        fmt.Println(v2, v38)
    }
    if true {
        v44 := GenS2{F0: v2}
        v4 = v4
        var v45 uint16 = 788
        _ = v44
        _ = v45
    } else {
        v46 := GenS3{F0: 42, F1: v2}
        var v47 int = -1
        _ = v46
        _ = v47
    }
    var v48 float64 = v38
    var v49 uint32 = 2147483648
    if v9 {
        switch -2147483648 {
        case 128:
            fmt.Println(v49)
        case 0:
            v50 := []int{}
            _ = v50
        case 4294967295:
            fmt.Println(v49)
        default:
            v4 = ""
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v48
    _ = v49
    _ = v9
}
