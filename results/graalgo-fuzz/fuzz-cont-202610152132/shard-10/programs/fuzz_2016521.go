package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 int8; F2 uint32 }
func genFn0(p0 int32, p1 bool) uint64 {
    v7 := []int{127, 9223372036854775807, 9223372036854775807, -494}
    switch "\n\t" {
    case "\n\t":
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 uint16 = 42
            _ = fi8
            _ = v9
        }
    case "foo bar":
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := GenS1{F0: 644, F1: 65535, F2: 255}
            _ = fi10
            _ = v11
        }
    case "":
        var v12 uint8 = 128
        _ = v12
    }
    fmt.Println(p0, p0)
    for fi13 := 0; fi13 < 15; fi13++ {
        p0 = (p0 << ((65535) & 63))
        _ = fi13
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1(p0 int8, p1 uint64) int {
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 int8 = p0
        v16 := []int{fi14}
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 uint = 1
        _ = fi17
        _ = v18
    }
    fmt.Println(p1)
    for fi19 := 0; fi19 < 20; fi19++ {
        p0 = p0
        _ = fi19
    }
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 uint32 = 42
    var v2 float32 = -2.25
    var v3 string = ((("\n\t" + "\n\t") + "hello") + ((("\n\t" + (("foo bar" + "a") + "")) + ("a" + "a")) + "\n\t"))
    var v4 int8 = 0
    var v5 int64 = 0
    var v6 uint64 = 1
    switch 2147483647 {
    case -47:
        var v20 int8 = (-(v4))
        _ = v20
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        fmt.Println(v5)
        var v22 uint32 = v1
        v23 := []int{(fi21 << ((810) & 63)), (fi21 % (((fi21 ^ genFn1((v4 % ((v4) | 1)), v6))) | 1)), fi21, (fi21 | fi21)}
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 string = ""
    if false {
        if (!(false && false)) {
            for fi25 := 0; fi25 < 11; fi25++ {
                v24 = ("hello" + (v24 + "hello"))
                _ = fi25
            }
            switch (("a" + (("hello" + v24) + ((((v24 + ("" + v24)) + ("\n\t" + v24)) + "hello") + ""))) + ("a" + v24)) {
            case "hello":
                for fi26 := 0; fi26 < 17; fi26++ {
                    switch 167 {
                    case -2147483648:
                        v5 += v5
                    case 0:
                        fmt.Println(v6)
                    case 9223372036854775807:
                        v27 := GenS1{F0: v1, F1: 822, F2: 65535}
                        _ = v27
                    }
                    _ = fi26
                }
            case "":
                if true {
                    if true {
                        fmt.Println(v3, v24, v3)
                        v28 := GenS1{F0: v1, F1: 65535, F2: 1}
                        _ = v28
                    }
                    var v29 bool = (42 <= 42)
                    _ = v29
                } else {
                    if true {
                        var v30 float64 = 1.5
                        if true {
                            var v31 float32 = v2
                            var v32 float64 = (v30 + (v30 / -2.25))
                            fmt.Println(v30, v30, v2)
                            _ = v31
                            _ = v32
                        }
                        _ = v30
                    } else {
                        var v33 int = 42
                        _ = v33
                    }
                    switch v24 {
                    case "foo bar":
                        v1 -= 867
                    }
                    if (255 <= 2) {
                        var v34 uint16 = 32767
                        if true {
                            for fi35 := 0; fi35 < 3; fi35++ {
                                for fi36 := 0; fi36 < 5; fi36++ {
                                    v37 := GenS1{F0: v1, F1: 2, F2: 0}
                                    v6 -= 2
                                    _ = fi36
                                    _ = v37
                                }
                                _ = fi35
                            }
                        }
                        _ = v34
                    } else {
                        var v38 uint = 1
                        _ = v38
                    }
                }
            default:
                for fi39 := 0; fi39 < 3; fi39++ {
                    var v40 int64 = (v5 | v5)
                    _ = fi39
                    _ = v40
                }
            }
        } else {
            var v41 uint32 = v1
            var v42 uint = 2
            var v43 uint32 = v41
            _ = v41
            _ = v42
            _ = v43
        }
        v44 := []int{-1, 32767}
        _ = v44
    }
    v45 := []int{-2147483648, -1, -308}
    v46 := GenS2{F0: 127}
    var v47 int64 = v5
    if len(v45) > 0 {
        v45[0] = 32767
    }
    if (-1 < 32768) {
        v45 = append(v45, genFn1(-128, 113))
        var v48 uint32 = 2147483648
        var v49 float32 = 0.0
        _ = v48
        _ = v49
    }
    fmt.Println(v4, v1, v1)
    var v50 uint32 = v1
    switch v24 {
    case "hello":
        _ = len(v45)
    case "\n\t":
        if (true || true) {
            for fi51 := 0; fi51 < 2; fi51++ {
                for fi52 := 0; fi52 < 6; fi52++ {
                    if len(v45) >= 2 {
                        v45 = v45[1:len(v45)]
                    }
                    switch (v24 + v3) {
                    case "hello":
                        var v53 int16 = -545
                        _ = v53
                    case "a":
                        switch 127 {
                        case 9223372036854775807:
                            if ((1 != 2) || false) {
                                var v54 float64 = float64(((v2 - (-2.25 - v2)) / 1.5))
                                switch (fi51 ^ (fi51 * 4294967295)) {
                                case 128:
                                    switch -9223372036854775808 {
                                    case 636:
                                        v6 = 42
                                    case 229:
                                        var v55 uint32 = v50
                                        _ = v55
                                    }
                                case -9223372036854775808:
                                    fmt.Println(fi52)
                                default:
                                    if false {
                                        var v56 bool = (!true)
                                        v56 = (false || false)
                                        _ = v56
                                    } else {
                                        _ = len(v45)
                                        v54 += v54
                                    }
                                }
                                _ = v54
                            }
                        case 32767:
                            var v57 uint = 218
                            _ = v57
                        default:
                            fmt.Println(fi51, fi51)
                        }
                        var v58 uint16 = 592
                        if (!((!false) || (2147483647 != 45))) {
                            if len(v45) >= 2 {
                                v45 = v45[1:len(v45)]
                            }
                            v45 = append(v45, (fi52 / ((genFn1(v4, (v6 % ((v6) | 1)))) | 1)))
                            v59 := GenS3{F0: v50}
                            _ = v59
                        }
                        v60 := GenS3{F0: 42}
                        _ = v58
                        _ = v60
                    }
                    var v61 uint32 = 0
                    if false {
                        fmt.Println(v5, fi52)
                        var v62 int = 42
                        _ = v62
                    } else {
                        var v63 int32 = -2147483648
                        v2 += 1e10
                        _ = v63
                    }
                    var v64 int8 = v4
                    _ = fi52
                    _ = v61
                    _ = v64
                }
                for fi65 := 0; fi65 < 8; fi65++ {
                    if true {
                        var v66 int = genFn1(-121, (v6 | (v6 + 18446744073709551615)))
                        v66 = (fi65 & (-(fi51)))
                        var v67 int8 = 1
                        _ = v66
                        _ = v67
                    }
                    var v68 float32 = v2
                    var v69 uint16 = 65535
                    v70 := GenS3{F0: v1}
                    _ = fi65
                    _ = v68
                    _ = v69
                    _ = v70
                }
                _ = fi51
            }
            v71 := GenS4{F0: 1, F1: 95, F2: (v1 * (v50 & v1))}
            _ = v71
        } else {
            v1 -= (v50 ^ v50)
            _ = len(v45)
        }
    case "a":
        v72 := GenS1{F0: 2, F1: 0, F2: 0}
        _ = v72
    default:
        fmt.Println(v5, v4, v3)
    }
    if (v4 != v4) {
        if ((v50 - v50) >= 0) {
            fmt.Println(v2)
            fmt.Println(v47, v6, v3)
            var v73 int32 = 400
            _ = v73
        }
    } else {
        var v74 string = "a"
        _ = v74
    }
    if len(v45) > 0 {
        v45[0] = 128
    }
    v45 = append(v45, genFn1((v4 + (v4 | (v4 ^ (v4 ^ v4)))), v6))
    fmt.Println(v24)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v50
    _ = v6
}
