package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int16; F2 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32; F1 int; F2 string }
func genFn0(p0 int) int32 {
    var v4 int64 = -9223372036854775808
    v5 := []int{p0, p0, (p0 + 127)}
    _ = p0
    _ = v4
    _ = v5
    return -427
}
func genFn1(p0 int, p1 int32) int32 {
    var v6 int64 = 0
    var v7 uint16 = 0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p1
}

func main() {
    var v1 int32 = 1
    var v2 float32 = 1.5
    var v3 int = 1
    var v8 uint8 = 128
    var v9 int64 = -1
    fmt.Println(v8, v8)
    var v10 int16 = -906
    if (v3 >= -9223372036854775808) {
        var v11 int8 = 90
        v3 -= 931
        _ = v11
    } else {
        switch (v3 >> ((42) & 63)) {
        case 32767:
            v10 = -32768
        case -627:
            v12 := GenS1{F0: "foo bar"}
            _ = v12
        case -2147483648:
            v13 := GenS1{F0: ""}
            _ = v13
        }
        v14 := GenS1{F0: (("\n\t" + (("hello" + "foo bar") + "\n\t")) + "\n\t")}
        for fi15 := 0; fi15 < 6; fi15++ {
            fmt.Println(v2, v9, v8)
            _ = fi15
        }
        _ = v14
    }
    if false {
        var v16 uint = 255
        _ = v16
    } else {
        v1 = (-(v1))
        switch v3 {
        case 592:
            v10 -= (-(v10))
        case -940:
            var v17 bool = true
            _ = v17
        case -9223372036854775808:
            if (240 < 4294967295) {
                if (42 <= 255) {
                    var v18 int16 = v10
                    v1 = (^(v1))
                    _ = v18
                }
                if (-123 > 0) {
                    for fi19 := 0; fi19 < 17; fi19++ {
                        v3 = v3
                        if true {
                            v20 := []int{}
                            _ = v20
                        } else {
                            v9 += v9
                        }
                        v21 := []int{v3, 9223372036854775807, (v3 + 9223372036854775807), (v3 << ((1) & 63))}
                        _ = fi19
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 15; fi22++ {
                        for fi23 := 0; fi23 < 15; fi23++ {
                            switch "" {
                            case "foo bar":
                                if (!(255 >= 881)) {
                                    if false {
                                        var v24 int = fi23
                                        _ = v24
                                    }
                                    var v25 int32 = v1
                                    _ = v25
                                } else {
                                    if true {
                                        v26 := []int{fi22, 42}
                                        v1 -= 905
                                        _ = v26
                                    }
                                    fmt.Println(fi22)
                                    if false {
                                        var v27 uint8 = (v8 | v8)
                                        _ = v27
                                    } else {
                                        switch (("hello" + "foo bar") + "\n\t") {
                                        case "foo bar":
                                            var v28 uint16 = 2
                                            _ = v28
                                        case "a":
                                            if (273 == 598) {
                                                v29 := []int{(fi22 * (^(fi22))), (^(fi22)), fi23, (v3 << ((1) & 63))}
                                                _ = v29
                                            } else {
                                                v30 := []int{32767, fi22}
                                                if true {
                                                    v31 := GenS1{F0: "foo bar"}
                                                    _ = v31
                                                } else {
                                                    fmt.Println(v8, v1)
                                                    v3 += fi23
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        fmt.Println(v8)
                                                    case "foo bar":
                                                        switch (fi22 ^ (^(fi23))) {
                                                        case -2147483648:
                                                            v1 = (v1 << ((42) & 63))
                                                        default:
                                                            fmt.Println(v9, v1, v2)
                                                        }
                                                    default:
                                                        v8 = v8
                                                    }
                                                }
                                                v1 = (v1 << ((2) & 63))
                                                _ = v30
                                            }
                                        case "\n\t":
                                            v32 := []int{}
                                            _ = v32
                                        default:
                                            v9 += (v9 / (((^(v9))) | 1))
                                        }
                                    }
                                    switch ("\n\t" + "a") {
                                    case "\n\t":
                                        v9 = 1
                                    case "a":
                                        var v33 string = "hello"
                                        _ = v33
                                    case "foo bar":
                                        if true {
                                            var v34 uint32 = 271
                                            fmt.Println(fi22, fi22)
                                            _ = v34
                                        }
                                    }
                                    v3 -= fi23
                                }
                            }
                            _ = fi23
                        }
                        _ = fi22
                    }
                    v2 -= ((3.14159 - 3.14159) - 1.5)
                }
            }
        }
        fmt.Println(v10, v2, v2)
    }
    v1 = (-(v1))
    v35 := GenS1{F0: (("\n\t" + "\n\t") + (("\n\t" + "foo bar") + ("\n\t" + "\n\t")))}
    if true {
        v36 := []int{v3}
        if false {
            if true {
                fmt.Println(v8)
                v37 := GenS1{F0: "a"}
                _ = v37
            }
        } else {
            fmt.Println(v10, v3, v9)
        }
        _ = v36
    } else {
        fmt.Println(v9, v9)
        if true {
            if false {
                var v38 int64 = 1
                _ = v38
            }
            var v39 float32 = 1.5
            if true {
                v40 := GenS2{F0: 0, F1: (^(v10)), F2: genFn1(-836, (v1 * 0))}
                v41 := []int{}
                _ = v40
                _ = v41
            }
            _ = v39
        } else {
            var v42 string = "foo bar"
            v43 := GenS1{F0: "hello"}
            _ = v42
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 14; fi44++ {
        v45 := GenS3{F0: (v9 % ((1) | 1))}
        v46 := []int{v3, fi44, 42}
        _ = fi44
        _ = v45
        _ = v46
    }
    switch (v3 ^ v3) {
    case -9223372036854775808:
        for fi47 := 0; fi47 < 8; fi47++ {
            fmt.Println(fi47)
            for fi48 := 0; fi48 < 14; fi48++ {
                v49 := []int{(fi48 - (fi48 / ((2147483647) | 1))), (fi48 - (^(fi47))), fi48, 9223372036854775807}
                v10 = v10
                _ = fi48
                _ = v49
            }
            _ = fi47
        }
    }
    if false {
        v8 += (v8 & (v8 * v8))
    } else {
        if (-373 < (v3 ^ (v3 >> ((0) & 63)))) {
            v50 := GenS4{F0: (v1 * v1), F1: false}
            switch (v3 << ((625) & 63)) {
            case 4294967295:
                if ((!false) && ("hello" != "\n\t")) {
                    v51 := GenS5{F0: (v1 << ((0) & 63))}
                    v3 += (-(v3))
                    _ = v51
                }
            case -9223372036854775808:
                if ((v8 >> ((501) & 63)) != 255) {
                    var v52 int8 = -1
                    _ = v52
                } else {
                    var v53 uint = 4294967295
                    v54 := GenS6{F0: (v10 >> ((v53) & 63))}
                    _ = v53
                    _ = v54
                }
            }
            _ = v50
        } else {
            v9 = (v9 * int64(v8))
            v3 += 1
        }
    }
    if true {
        var v55 int8 = 42
        v56 := GenS7{F0: 0, F1: (v3 % ((v3) | 1)), F2: "\n\t"}
        _ = v55
        _ = v56
    }
    switch (v3 | v3) {
    case 128:
        var v57 uint64 = 734
        _ = v57
    case 353:
        var v58 uint64 = 1
        _ = v58
    case 32768:
        for fi59 := 0; fi59 < 16; fi59++ {
            for fi60 := 0; fi60 < 11; fi60++ {
                v10 -= (v10 * (v10 & (v10 / (((-(v10))) | 1))))
                _ = fi60
            }
            _ = fi59
        }
    default:
        if (true || false) {
            v9 = 42
            v61 := []int{v3}
            _ = len(v61)
            _ = v61
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v35
    _ = v8
    _ = v9
}
