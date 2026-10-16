package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint64 }
func genFn0() uint {
    v4 := GenS1{F0: "hello", F1: 451}
    var v5 float64 = float64(0.1)
    v5 -= v5
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 int16, p1 uint16, p2 int8) uint {
    p1 = 255
    var v6 uint = genFn0()
    if true {
        for fi7 := 0; fi7 < 13; fi7++ {
            p2 += (^(p2))
            var v8 float64 = 3.14159
            v9 := GenS2{F0: p2, F1: (p0 ^ (p0 * 479))}
            _ = fi7
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 14; fi10++ {
            switch -611 {
            case -9223372036854775808:
                p1 -= p1
            case 205:
                p2 = int8(v6)
            }
            _ = fi10
        }
    } else {
        for fi11 := 0; fi11 < 16; fi11++ {
            if ((p1 / ((0) | 1)) == (^(p1))) {
                if (false || (!false)) {
                    p2 = p2
                    var v12 int64 = -9223372036854775808
                    switch (^(fi11)) {
                    case 42:
                        v12 += v12
                    case 32767:
                        var v13 uint32 = 65535
                        _ = v13
                    case 128:
                        for fi14 := 0; fi14 < 6; fi14++ {
                            v15 := []int{fi11, fi11, fi11, (^(fi11))}
                            var v16 int = 454
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    default:
                        v17 := GenS2{F0: 1, F1: (p0 << (((v6 % ((13) | 1))) & 63))}
                        _ = v17
                    }
                    _ = v12
                } else {
                    var v18 int = -2147483648
                    v6 = 284
                    _ = v18
                }
                v6 = v6
            } else {
                switch (fi11 << ((0) & 63)) {
                case -9223372036854775808:
                    switch (^(fi11)) {
                    case -181:
                        if true {
                            v6 -= genFn0()
                            if true {
                                for fi19 := 0; fi19 < 19; fi19++ {
                                    v20 := GenS1{F0: (("" + ("a" + ("a" + ("hello" + "")))) + ""), F1: 42}
                                    _ = fi19
                                    _ = v20
                                }
                                if true {
                                    v21 := GenS2{F0: p2, F1: p0}
                                    var v22 uint16 = (p1 - (p1 + p1))
                                    for fi23 := 0; fi23 < 3; fi23++ {
                                        var v24 uint = (v6 + genFn0())
                                        var v25 uint16 = 65535
                                        var v26 int16 = (p0 * p0)
                                        _ = fi23
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    }
                                    _ = v21
                                    _ = v22
                                } else {
                                    v6 = v6
                                }
                                p2 -= int8(v6)
                            }
                        } else {
                            if ((true && true) && false) {
                                if true {
                                    for fi27 := 0; fi27 < 10; fi27++ {
                                        var v28 int32 = -2147483648
                                        _ = fi27
                                        _ = v28
                                    }
                                    v29 := GenS2{F0: (p2 * p2), F1: p0}
                                    var v30 int64 = -271
                                    _ = v29
                                    _ = v30
                                } else {
                                    v31 := []int{}
                                    var v32 uint16 = p1
                                    fmt.Println(v6, fi11, v32)
                                    _ = v31
                                    _ = v32
                                }
                                p1 += (^(p1))
                                var v33 uint32 = 2
                                _ = v33
                            }
                            v34 := GenS1{F0: ("a" + (("\n\t" + "a") + "")), F1: 1}
                            _ = v34
                        }
                    case -2147483648:
                        v6 = v6
                    case 1:
                        v35 := []int{fi11}
                        _ = v35
                    }
                case 1:
                    p1 = (p1 & (^(p1)))
                }
                var v36 int = -745
                _ = v36
            }
            for fi37 := 0; fi37 < 13; fi37++ {
                var v38 int32 = -2147483648
                _ = fi37
                _ = v38
            }
            _ = fi11
        }
        if false {
            v39 := GenS3{F0: int(v6), F1: 2147483648}
            var v40 int32 = int32(p1)
            _ = v39
            _ = v40
        }
        v6 = v6
    }
    if (0 > 1) {
        if true {
            v41 := []int{42, 9223372036854775807, 2147483647, -9223372036854775808}
            var v42 int32 = -277
            var v43 uint = v6
            _ = v41
            _ = v42
            _ = v43
        }
        var v44 uint32 = 2147483648
        v45 := GenS1{F0: ("" + "a"), F1: -114}
        _ = v44
        _ = v45
    } else {
        var v46 uint = 255
        p2 = 1
        for fi47 := 0; fi47 < 15; fi47++ {
            var v48 int8 = p2
            _ = fi47
            _ = v48
        }
        _ = v46
    }
    for fi49 := 0; fi49 < 16; fi49++ {
        v50 := GenS2{F0: p2, F1: p0}
        v51 := GenS4{F0: p1, F1: 570, F2: 2}
        p1 = p1
        _ = fi49
        _ = v50
        _ = v51
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return genFn0()
}

func main() {
    var v1 int16 = 336
    var v2 int32 = 1
    v3 := []int{}
    fmt.Println(v2, v2)
    v52 := GenS2{F0: 0, F1: (^(v1))}
    if false {
        _ = len(v3)
    }
    for fi53 := 0; fi53 < 5; fi53++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v54 int8 = -8
        v54 += (v54 * 23)
        _ = fi53
        _ = v54
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v55 := GenS3{F0: 114, F1: 2147483647}
    for fi56 := 0; fi56 < 1; fi56++ {
        if (true && (0.1 == 1.5)) {
            var v57 uint = 0
            switch (^(fi56)) {
            case 1:
                if ((fi56 | (fi56 ^ fi56)) <= 1) {
                    for fi58 := 0; fi58 < 16; fi58++ {
                        v2 = v2
                        v3 = append(v3, fi58)
                        _ = fi58
                    }
                    var v59 bool = false
                    _ = v59
                }
            }
            _ = v57
        }
        _ = fi56
    }
    if len(v3) > 0 {
        v3[0] = -77
    }
    v1 += v1
    var v60 int32 = (-(v2))
    v2 = v60
    if false {
        if false {
            var v61 uint64 = 4294967295
            v60 = int32(v1)
            v62 := GenS4{F0: 32767, F1: v61, F2: v61}
            _ = v61
            _ = v62
        }
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
    _ = v55
    _ = v60
}
