package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 uint32; F1 int64 }
func genFn0(p0 int32, p1 int64, p2 bool) bool {
    switch int(p0) {
    case 32768:
        if p2 {
            var v4 uint64 = 1
            for fi5 := 0; fi5 < 12; fi5++ {
                var v6 uint64 = v4
                p0 -= -1
                _ = fi5
                _ = v6
            }
            if (18446744073709551615 > 4294967295) {
                switch 747 {
                case -1:
                    for fi7 := 0; fi7 < 1; fi7++ {
                        for fi8 := 0; fi8 < 6; fi8++ {
                            if true {
                                fmt.Println(p0, p2)
                                fmt.Println(p1)
                            } else {
                                if p2 {
                                    v9 := []int{(fi8 << ((255) & 63)), (fi7 % (((fi8 ^ 32768)) | 1))}
                                    _ = v9
                                } else {
                                    var v10 bool = false
                                    switch int(p0) {
                                    case 0:
                                        v11 := GenS1{F0: p1, F1: 42, F2: fi8}
                                        _ = v11
                                    }
                                    v12 := GenS1{F0: p1, F1: 42, F2: fi8}
                                    _ = v10
                                    _ = v12
                                }
                                var v13 bool = (p2 || p2)
                                _ = v13
                            }
                            var v14 uint = 1
                            var v15 int = fi7
                            _ = fi8
                            _ = v14
                            _ = v15
                        }
                        for fi16 := 0; fi16 < 18; fi16++ {
                            v17 := []int{(^(fi16)), fi16, (fi16 >> ((18446744073709551615) & 63)), (^(fi16))}
                            _ = fi16
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 6; fi18++ {
                            if (-879 > (-(p0))) {
                                v19 := []int{(^(fi18)), (-(fi7))}
                                _ = v19
                            }
                            fmt.Println(fi7, p1, fi18)
                            fmt.Println(p2)
                            _ = fi18
                        }
                        _ = fi7
                    }
                }
                v20 := []int{42, -2147483648, -382}
                var v21 string = "hello"
                _ = v20
                _ = v21
            } else {
                for fi22 := 0; fi22 < 3; fi22++ {
                    v23 := GenS1{F0: (p1 | -9223372036854775808), F1: 127, F2: -9223372036854775808}
                    v24 := GenS1{F0: p1, F1: -1, F2: -73}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                v25 := []int{-9223372036854775808}
                _ = v25
            }
            _ = v4
        }
    case 42:
        fmt.Println(p1, p0, p1)
    case -2147483648:
        for fi26 := 0; fi26 < 15; fi26++ {
            fmt.Println(fi26)
            fmt.Println(fi26)
            var v27 uint = 0
            _ = fi26
            _ = v27
        }
    }
    var v28 int64 = p1
    for fi29 := 0; fi29 < 19; fi29++ {
        for fi30 := 0; fi30 < 4; fi30++ {
            if (42 < -1) {
                if (((1e10 + 3.14159) == ((0.1 - -2.25) - 0.1)) && (p0 >= 0)) {
                    var v31 uint32 = 774
                    v32 := []int{fi29, (fi30 - 0), fi30, fi30}
                    var v33 uint = 1
                    _ = v31
                    _ = v32
                    _ = v33
                }
                var v34 float64 = float64((((3.14159 / float32(1.5)) * 0.1) * ((3.14159 - float32((-2.25 * 1.5))) - (1.5 + 1.5))))
                _ = v34
            } else {
                v35 := GenS2{F0: 42}
                _ = v35
            }
            _ = fi30
        }
        _ = fi29
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    return p2
}

func main() {
    var v1 uint = 467
    var v2 uint32 = 893
    v3 := make(map[uint32]bool)
    for fi36 := 0; fi36 < 13; fi36++ {
        v1 += 618
        _ = fi36
    }
    switch ("a" + "") {
    case "hello":
        switch ("" + "\n\t") {
        case "":
            var v37 int8 = -128
            _ = v37
        case "hello":
            v38 := []int{955, 0, -538, 32767}
            _ = v38
        default:
            v39 := GenS3{F0: -1, F1: 65535}
            _ = v39
        }
    case "foo bar":
        v40 := GenS4{F0: v2, F1: ((false || (("a" + "") != "")) || true)}
        _ = v40
    }
    for fi41 := 0; fi41 < 3; fi41++ {
        _ = len(v3)
        v42 := []int{(fi41 / ((9223372036854775807) | 1)), 374, fi41}
        _ = fi41
        _ = v42
    }
    var v43 string = "hello"
    var v44 string = ((("\n\t" + (("" + "") + v43)) + (v43 + v43)) + v43)
    fmt.Println(v43, v43, v43)
    var v45 float32 = 1e10
    v46 := GenS4{F0: v2, F1: (((v2 & (^(v2))) >= 802) && false)}
    v47 := GenS5{F0: 42, F1: 23}
    v3[(v2 + v2)] = (2 < 42)
    var v48 uint16 = 274
    for fi49 := 0; fi49 < 14; fi49++ {
        v50 := []int{(-(fi49))}
        v51 := GenS4{F0: (v2 | uint32(v1)), F1: true}
        v52 := GenS6{F0: v2, F1: 42}
        _ = fi49
        _ = v50
        _ = v51
        _ = v52
    }
    var v53 uint32 = 982
    switch -139 {
    case 676:
        var v54 int32 = -783
        _ = v54
    }
    v55 := GenS6{F0: v2, F1: 1}
    var v56 int = 32768
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v53:", v53)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v53
    _ = v55
    _ = v56
}
