package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint32 }
func genFn0(p0 bool) uint8 {
    fmt.Println(p0)
    if p0 {
        v7 := GenS1{F0: 714, F1: -32768}
        var v8 uint64 = 1
        _ = v7
        _ = v8
    } else {
        p0 = false
        p0 = false
        v9 := []int{-2147483648}
        _ = v9
    }
    var v10 string = "\n\t"
    for fi11 := 0; fi11 < 17; fi11++ {
        v10 = "\n\t"
        _ = fi11
    }
    v12 := GenS1{F0: 0, F1: -1}
    _ = p0
    _ = v10
    _ = v12
    return 128
}
func genFn1(p0 uint8) uint64 {
    for fi13 := 0; fi13 < 12; fi13++ {
        p0 = genFn0(false)
        for fi14 := 0; fi14 < 3; fi14++ {
            if true {
                switch (fi13 & fi14) {
                case 42:
                    switch (fi13 & 32767) {
                    case 128:
                        p0 += p0
                    case -9223372036854775808:
                        fmt.Println(fi13, p0, fi13)
                    }
                    switch 398 {
                    case -1:
                        v15 := []int{128, (fi13 << ((340) & 63)), fi14}
                        _ = v15
                    case 0:
                        if false {
                            fmt.Println(fi14, fi14, p0)
                            p0 = p0
                            var v16 int16 = int16(fi13)
                            _ = v16
                        }
                    case 1:
                        switch ("a" + (("\n\t" + "a") + "hello")) {
                        case "a":
                            v17 := []int{(fi13 - 705), fi13, -117}
                            _ = v17
                        default:
                            var v18 uint = 604
                            _ = v18
                        }
                    }
                    var v19 float32 = 1.5
                    _ = v19
                case 633:
                    if (42 > 1) {
                        var v20 int64 = -592
                        _ = v20
                    }
                case 128:
                    v21 := []int{fi14, 903, 4294967295, -1}
                    _ = v21
                }
                p0 = (p0 * (p0 % ((p0) | 1)))
                var v22 string = ""
                _ = v22
            }
            _ = fi14
        }
        _ = fi13
    }
    switch "" {
    case "hello":
        for fi23 := 0; fi23 < 4; fi23++ {
            if false {
                if (!false) {
                    p0 += (p0 / ((1) | 1))
                } else {
                    v24 := []int{fi23, fi23, fi23, (-(fi23))}
                    v25 := GenS1{F0: 0, F1: 42}
                    if len(v24) > 0 {
                        v24[0] = fi23
                    }
                    _ = v24
                    _ = v25
                }
                for fi26 := 0; fi26 < 3; fi26++ {
                    var v27 string = "hello"
                    _ = fi26
                    _ = v27
                }
            } else {
                var v28 int64 = -153
                var v29 int64 = 42
                for fi30 := 0; fi30 < 13; fi30++ {
                    if false {
                        v31 := []int{2147483647, fi23, (fi30 / ((fi30) | 1)), (fi23 + 2147483647)}
                        v28 -= v28
                        switch ("\n\t" + "\n\t") {
                        case "foo bar":
                            v32 := GenS1{F0: 0, F1: 42}
                            _ = v32
                        case "\n\t":
                            v33 := GenS1{F0: -1, F1: -1}
                            _ = v33
                        case "a":
                            fmt.Println(fi30, v29)
                        default:
                            v28 += (v28 << ((57) & 63))
                        }
                        fmt.Println(v28, fi30, fi30)
                        _ = v31
                    } else {
                        var v34 int16 = 42
                        _ = v34
                    }
                    _ = fi30
                }
                _ = v28
                _ = v29
            }
            _ = fi23
        }
    default:
        var v35 int32 = 2147483647
        _ = v35
    }
    var v36 int32 = -2147483648
    _ = p0
    _ = v36
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    var v5 int64 = 419
    v6 := make(map[int]bool)
    if ((("hello" + ("foo bar" + (("foo bar" + "\n\t") + ("hello" + "hello")))) + ("hello" + ("\n\t" + "hello"))) <= "") {
        if (false || (!false)) {
            var v37 bool = ((v1 >> ((0) & 63)) >= (v1 << ((4294967295) & 63)))
            v5 = v5
            _ = v37
        } else {
            v38 := []int{53}
            if true {
                var v39 int8 = -80
                _ = v39
            }
            _ = v38
        }
    }
    var v40 int8 = -33
    for fi41 := 0; fi41 < 13; fi41++ {
        v42 := []int{fi41, -9223372036854775808, fi41}
        fmt.Println(v1, fi41)
        if (42 >= -115) {
            v43 := v2["\n\t"]
            v44 := GenS2{F0: v43}
            v45 := GenS2{F0: (v1 <= 65535)}
            _ = v43
            _ = v44
            _ = v45
        }
        _ = fi41
        _ = v42
    }
    var v46 float32 = (3.14159 + -2.25)
    switch 128 {
    case 17:
        if ((0.0 / (1e10 * ((float64((3.14159 / v46)) * float64(v46)) * (1.5 - 0.1)))) < 0.0) {
            for fi47 := 0; fi47 < 3; fi47++ {
                v48 := v4[(fi47 & fi47)]
                var v49 int8 = v40
                v50 := GenS1{F0: 1, F1: 1000}
                _ = fi47
                _ = v48
                _ = v49
                _ = v50
            }
            var v51 int64 = (^(v5))
            _ = v51
        }
    case 32768:
        v52 := GenS3{F0: v40, F1: (v1 & (v1 | 2147483647))}
        _ = v52
    case 320:
        v1 = (v1 | 2)
    }
    if false {
        v46 -= 1e10
        for fi53 := 0; fi53 < 17; fi53++ {
            for fi54 := 0; fi54 < 17; fi54++ {
                fmt.Println(v46, v46)
                v46 += 3.14159
                v55 := []int{}
                _ = fi54
                _ = v55
            }
            var v56 int64 = (v5 & v5)
            var v57 int16 = 42
            _ = fi53
            _ = v56
            _ = v57
        }
        var v58 int = 127
        _ = v58
    }
    switch "" {
    case "":
        v5 -= v5
    default:
        fmt.Println(v5, v40)
    }
    if true {
        delete(v2, "")
    } else {
        for fi59 := 0; fi59 < 18; fi59++ {
            var v60 float64 = 0.0
            switch 605 {
            case 2147483647:
                switch (fi59 ^ (fi59 / (((^(fi59))) | 1))) {
                case 2147483647:
                    fmt.Println(fi59, v5, v46)
                case -9223372036854775808:
                    v40 = v40
                case -1:
                    v61 := []int{}
                    _ = v61
                }
            case 32768:
                switch (-(fi59)) {
                case 9223372036854775807:
                    if true {
                        v60 = v60
                        switch "a" {
                        case "foo bar":
                            for fi62 := 0; fi62 < 18; fi62++ {
                                v60 += (((((((v60 / v60) + 1.5) * ((v60 * v60) * (0.1 + v60))) * float64(3.14159)) / v60) / 0.1) - ((v60 * v60) - v60))
                                var v63 uint8 = 2
                                _ = fi62
                                _ = v63
                            }
                        default:
                            fmt.Println(fi59, v60)
                        }
                        v64 := []int{}
                        _ = v64
                    }
                }
            case -405:
                v65 := GenS1{F0: -2147483648, F1: 1000}
                _ = v65
            default:
                fmt.Println(v60, v40, v46)
            }
            v66 := v4[(fi59 * 9223372036854775807)]
            _ = fi59
            _ = v60
            _ = v66
        }
        fmt.Println(v40)
    }
    delete(v3, (^(v1)))
    v67 := v3[v1]
    v68 := []int{-9223372036854775808, 42, -2147483648, -2147483648}
    v67 = v67
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v67:", v67)
    fmt.Println("slicelen v68:", len(v68), v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v46
    _ = v5
    _ = v6
    _ = v67
    _ = v68
}
