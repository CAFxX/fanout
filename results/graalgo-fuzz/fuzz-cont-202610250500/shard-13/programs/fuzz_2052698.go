package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint16, p1 int8) uint8 {
    var v7 int32 = 0
    var v8 uint8 = uint8(p0)
    v7 = (-(v7))
    v9 := GenS1{F0: -341}
    if (!true) {
        switch -484 {
        case 32767:
            if true {
                for fi10 := 0; fi10 < 20; fi10++ {
                    p1 = (p1 / (((-(p1))) | 1))
                    if (104 <= 227) {
                        var v11 int8 = 42
                        var v12 uint16 = 0
                        v8 = 128
                        _ = v11
                        _ = v12
                    } else {
                        v13 := GenS2{F0: (p0 ^ (p0 >> ((uint(v8)) & 63))), F1: 42}
                        switch int(p1) {
                        case -289:
                            p0 += p0
                        case -2147483648:
                            v14 := GenS3{F0: (p0 << ((1) & 63)), F1: v8}
                            _ = v14
                        case -9223372036854775808:
                            var v15 int64 = -1
                            _ = v15
                        }
                        _ = v13
                    }
                    v8 -= uint8(p0)
                    _ = fi10
                }
                var v16 uint16 = 985
                if true {
                    fmt.Println(v7)
                } else {
                    v17 := GenS3{F0: (^(v16)), F1: 1}
                    v18 := []int{9223372036854775807, 920, -9223372036854775808}
                    var v19 int64 = 42
                    _ = v17
                    _ = v18
                    _ = v19
                }
                _ = v16
            } else {
                fmt.Println(v7, p1)
            }
        case -9223372036854775808:
            fmt.Println(v8, p0)
        default:
            v20 := GenS1{F0: 32767}
            _ = v20
        }
        switch "a" {
        case "foo bar":
            var v21 uint = 422
            _ = v21
        case "a":
            for fi22 := 0; fi22 < 9; fi22++ {
                v23 := []int{1, fi22}
                _ = fi22
                _ = v23
            }
        case "\n\t":
            var v24 uint32 = 1
            _ = v24
        }
        fmt.Println(p1, v8, v8)
    } else {
        switch 1 {
        case -662:
            v8 = 95
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (v8 + (v8 % (((v8 + v8)) | 1)))
}
func genFn1(p0 uint8) uint {
    switch 32767 {
    case 2147483647:
        var v25 int64 = -1
        _ = v25
    default:
        var v26 int32 = 985
        _ = v26
    }
    switch ("" + ("a" + "hello")) {
    case "foo bar":
        for fi27 := 0; fi27 < 3; fi27++ {
            if false {
                for fi28 := 0; fi28 < 9; fi28++ {
                    var v29 int16 = 43
                    v29 -= (v29 / (((v29 * v29)) | 1))
                    _ = fi28
                    _ = v29
                }
                p0 -= p0
                if false {
                    for fi30 := 0; fi30 < 19; fi30++ {
                        var v31 int64 = 1
                        _ = fi30
                        _ = v31
                    }
                    v32 := []int{}
                    _ = v32
                } else {
                    if false {
                        for fi33 := 0; fi33 < 19; fi33++ {
                            v34 := GenS1{F0: fi27}
                            fmt.Println(fi27, fi33, fi27)
                            fmt.Println(fi27, p0, fi27)
                            _ = fi33
                            _ = v34
                        }
                    }
                }
            }
            for fi35 := 0; fi35 < 10; fi35++ {
                fmt.Println(p0, fi35, fi27)
                switch fi27 {
                case 707:
                    v36 := []int{(fi27 / ((fi27) | 1))}
                    _ = v36
                default:
                    v37 := GenS1{F0: (fi27 << ((2) & 63))}
                    _ = v37
                }
                _ = fi35
            }
            v38 := []int{fi27, fi27}
            _ = fi27
            _ = v38
        }
    case "\n\t":
        if true {
            v39 := []int{-911, 128}
            p0 = (^(p0))
            _ = v39
        }
    }
    _ = p0
    return 377
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 0
    var v3 int16 = -600
    var v4 string = ((("foo bar" + "\n\t") + ("" + (("hello" + (("foo bar" + (("foo bar" + "\n\t") + "\n\t")) + (("\n\t" + "foo bar") + ("" + "")))) + "\n\t"))) + "")
    var v5 int = 1
    var v6 float32 = 1e10
    v5 = 334
    fmt.Println(v3)
    v40 := GenS2{F0: 1, F1: 4294967295}
    if true {
        for fi41 := 0; fi41 < 17; fi41++ {
            if (0 != (v2 - -1)) {
                fmt.Println(v2)
                for fi42 := 0; fi42 < 20; fi42++ {
                    v1[("hello" + "")] = 2147483648
                    fmt.Println(fi41)
                    _ = fi42
                }
                var v43 uint8 = 42
                _ = v43
            } else {
                for fi44 := 0; fi44 < 16; fi44++ {
                    v4 = v4
                    v45 := GenS1{F0: 32768}
                    var v46 uint64 = 4294967295
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                var v47 int8 = -60
                _ = v47
            }
            var v48 uint32 = 42
            fmt.Println(v4)
            _ = fi41
            _ = v48
        }
        var v49 bool = true
        _ = v49
    } else {
        v3 = v3
    }
    var v50 string = ("hello" + "\n\t")
    v51 := GenS2{F0: 32767, F1: 2}
    var v52 string = v4
    v52 = v50
    for fi53 := 0; fi53 < 18; fi53++ {
        for fi54 := 0; fi54 < 8; fi54++ {
            var v55 float32 = v6
            v56 := GenS4{F0: 892, F1: -264}
            v6 = v6
            v6 += v6
            _ = fi54
            _ = v55
            _ = v56
        }
        fmt.Println(fi53, v3, v5)
        v57 := []int{(^(fi53)), (v5 / ((v5) | 1))}
        _ = fi53
        _ = v57
    }
    v58 := GenS1{F0: int(v2)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v58
    _ = v6
}
