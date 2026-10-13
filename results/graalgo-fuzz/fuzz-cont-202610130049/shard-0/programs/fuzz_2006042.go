package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 string; F2 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int }
func genFn0(p0 bool, p1 uint16) int {
    v4 := GenS1{F0: "", F1: -835, F2: -1}
    v5 := []int{0}
    fmt.Println(p0, p0)
    var v6 string = "hello"
    for fi7 := 0; fi7 < 8; fi7++ {
        if (!(-128 < 0)) {
            if (p1 > (p1 + 404)) {
                fmt.Println(p0)
            } else {
                var v8 int32 = -2147483648
                p0 = (!(!p0))
                v9 := GenS1{F0: ((v6 + "foo bar") + (v6 + v6)), F1: -32768, F2: -1}
                _ = v8
                _ = v9
            }
            var v10 int8 = 127
            _ = v10
        }
        var v11 float32 = (3.14159 + ((3.14159 * (-2.25 + (1.5 + (float32(1.5) / 3.14159)))) * -2.25))
        _ = fi7
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 32767
}
func genFn1(p0 int64, p1 uint64, p2 int8) int16 {
    v12 := GenS1{F0: ("a" + "a"), F1: 1000, F2: (p2 & 43)}
    var v13 int64 = p0
    for fi14 := 0; fi14 < 2; fi14++ {
        v13 = (v13 & 42)
        if false {
            v15 := []int{}
            var v16 uint64 = 0
            if (1 != 42) {
                if (!true) {
                    _ = len(v15)
                    if false {
                        v17 := GenS2{F0: (-(fi14))}
                        var v18 int32 = 1
                        var v19 int32 = v18
                        _ = v17
                        _ = v18
                        _ = v19
                    } else {
                        fmt.Println(p1, fi14)
                        v13 += v13
                        fmt.Println(v13, fi14, p2)
                    }
                }
            }
            _ = v15
            _ = v16
        } else {
            if (-1 != (^(v13))) {
                switch (fi14 * (fi14 >> ((42) & 63))) {
                case 128:
                    if true {
                        for fi20 := 0; fi20 < 15; fi20++ {
                            switch -2147483648 {
                            case 42:
                                if true {
                                    p1 += (p1 + 42)
                                    if (!true) {
                                        if true {
                                            p0 = (p0 + v13)
                                            if false {
                                                if (("\n\t" != "\n\t") || false) {
                                                    fmt.Println(p0, p2, p0)
                                                } else {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        fmt.Println(p0, p1, p0)
                                                    }
                                                    fmt.Println(p1, fi14)
                                                    p2 -= (p2 & p2)
                                                }
                                                v21 := []int{-1}
                                                if len(v21) >= 2 {
                                                    v21 = v21[1:len(v21)]
                                                }
                                                _ = v21
                                            } else {
                                                switch "" {
                                                case "":
                                                    fmt.Println(p1, fi14)
                                                case "hello":
                                                    var v22 string = "a"
                                                    _ = v22
                                                }
                                                v23 := GenS2{F0: (fi20 - int(p1))}
                                                var v24 uint32 = 1
                                                _ = v23
                                                _ = v24
                                            }
                                            v25 := GenS2{F0: genFn0(false, 255)}
                                            _ = v25
                                        }
                                        fmt.Println(fi20)
                                        fmt.Println(fi20)
                                    }
                                    v26 := []int{(fi20 % (((fi20 ^ (fi14 / (((fi14 / ((fi14) | 1))) | 1)))) | 1)), (fi14 & fi14), 1}
                                    _ = v26
                                }
                            }
                            v27 := []int{42, (-(fi14)), -1, (fi14 * genFn0(true, 2))}
                            _ = fi20
                            _ = v27
                        }
                        p0 = v13
                    }
                default:
                    var v28 int8 = p2
                    _ = v28
                }
                v29 := []int{(fi14 % ((28) | 1)), fi14}
                _ = v29
            } else {
                var v30 bool = true
                for fi31 := 0; fi31 < 5; fi31++ {
                    v32 := GenS1{F0: "a", F1: int16(v13), F2: (^(p2))}
                    var v33 bool = (false || v30)
                    fmt.Println(v30)
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                _ = v30
            }
            var v34 uint32 = 0
            v35 := []int{fi14}
            _ = v34
            _ = v35
        }
        if (42 != 24) {
            v36 := GenS2{F0: genFn0(false, 255)}
            var v37 uint64 = (p1 & (p1 % (((p1 | 291)) | 1)))
            var v38 int8 = (p2 << ((65535) & 63))
            _ = v36
            _ = v37
            _ = v38
        } else {
            var v39 string = "\n\t"
            _ = v39
        }
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 uint16 = 0
    v2 := []uint32{}
    var v3 uint32 = 2147483648
    switch 42 {
    case -2147483648:
        var v40 int8 = 127
        _ = v40
    case 197:
        switch 718 {
        case 2147483647:
            for fi41 := 0; fi41 < 18; fi41++ {
                if (false || ("" != "hello")) {
                    if ("hello" != (("hello" + ("" + "a")) + "hello")) {
                        v42 := GenS3{F0: v1, F1: ((("\n\t" + "\n\t") + "a") + "hello"), F2: "\n\t"}
                        v43 := GenS4{F0: v1}
                        _ = v42
                        _ = v43
                    }
                }
                _ = fi41
            }
        case 127:
            fmt.Println(v1)
        default:
            v1 = v1
        }
    default:
        v1 = v1
    }
    var v44 float32 = 0.0
    var v45 string = "hello"
    v3 += 131
    var v46 uint16 = 1
    switch -1 {
    case 1:
        if (!(!(255 < 1))) {
            v46 += 284
            _ = len(v2)
            v47 := GenS5{F0: 0}
            _ = v47
        } else {
            v45 = v45
        }
    case -628:
        var v48 float32 = v44
        _ = v48
    default:
        var v49 string = (v45 + (v45 + ((("a" + v45) + (((v45 + ("\n\t" + v45)) + v45) + "\n\t")) + v45)))
        _ = v49
    }
    var v50 uint16 = v1
    var v51 int = -9223372036854775808
    if false {
        for fi52 := 0; fi52 < 2; fi52++ {
            v53 := GenS3{F0: v1, F1: "\n\t", F2: v45}
            _ = fi52
            _ = v53
        }
        v45 = v45
    }
    if false {
        if false {
            v46 = v50
        } else {
            if (-32768 == genFn1(-9223372036854775808, 1, 0)) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v54 int8 = -128
                _ = v54
            } else {
                var v55 float32 = v44
                var v56 int64 = 9223372036854775807
                _ = v55
                _ = v56
            }
            for fi57 := 0; fi57 < 1; fi57++ {
                v44 = (((float32(1.5) / (0.0 * ((v44 + float32(1e10)) + float32(-2.25)))) * v44) - v44)
                _ = fi57
            }
        }
    } else {
        var v58 float32 = float32(0.1)
        _ = v58
    }
    var v59 float64 = 0.0
    var v60 uint32 = v3
    v46 -= (v1 & 1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v46
    _ = v50
    _ = v51
    _ = v59
    _ = v60
}
