package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 uint16; F2 uint8 }
type GenS5 struct { F0 bool; F1 int16; F2 uint16 }
func genFn0() uint8 {
    if false {
        for fi5 := 0; fi5 < 16; fi5++ {
            for fi6 := 0; fi6 < 18; fi6++ {
                if true {
                    v7 := GenS1{F0: 42, F1: 949}
                    var v8 bool = true
                    _ = v7
                    _ = v8
                }
                var v9 int32 = 88
                v10 := GenS2{F0: 2, F1: 42}
                _ = fi6
                _ = v10
                _ = v9
            }
            v11 := GenS2{F0: 255, F1: 595}
            var v12 uint64 = 700
            _ = fi5
            _ = v11
            _ = v12
        }
    }
    v13 := []int{-1, 42, 972}
    _ = v13
    return 127
}
func genFn1(p0 int8, p1 uint16) uint8 {
    if ("a" <= "\n\t") {
        v14 := []int{274, 2147483647, 42, 262}
        _ = v14
    }
    if (127 >= 466) {
        switch 4294967295 {
        case -837:
            for fi15 := 0; fi15 < 17; fi15++ {
                p0 -= p0
                p1 = 6
                var v16 int = fi15
                _ = fi15
                _ = v16
            }
        default:
            p1 = (p1 >> ((921) & 63))
        }
        p0 = p0
    } else {
        if ((-1 >= -2147483648) && (!true)) {
            v17 := GenS3{F0: 100}
            _ = v17
        }
        switch "" {
        case "a":
            for fi18 := 0; fi18 < 11; fi18++ {
                var v19 uint16 = p1
                var v20 int = fi18
                v21 := []int{v20, (v20 & v20), (-(fi18))}
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        case "":
            var v22 int32 = -2147483648
            _ = v22
        }
        for fi23 := 0; fi23 < 16; fi23++ {
            var v24 int64 = 42
            if false {
                if (2147483648 >= 0) {
                    p1 = 856
                    v25 := []int{fi23, 32767}
                    _ = v25
                }
                for fi26 := 0; fi26 < 19; fi26++ {
                    switch ("foo bar" + ("foo bar" + "hello")) {
                    case "a":
                        var v27 uint32 = 2147483647
                        _ = v27
                    case "foo bar":
                        switch (fi26 / (((fi26 - (fi26 - (fi26 >> ((2) & 63))))) | 1)) {
                        case 1:
                            v28 := []int{(^(fi26)), (fi23 + (fi26 - fi26)), fi26}
                            _ = v28
                        case -2147483648:
                            var v29 uint8 = 248
                            _ = v29
                        default:
                            fmt.Println(fi26, v24, p0)
                        }
                    default:
                        switch (fi23 * fi26) {
                        case 2147483647:
                            var v30 bool = true
                            _ = v30
                        case 127:
                            fmt.Println(v24, fi23)
                        case -384:
                            fmt.Println(fi23, fi26, fi23)
                        default:
                            if (true && true) {
                                switch "a" {
                                case "foo bar":
                                    var v31 float64 = float64(0.0)
                                    _ = v31
                                }
                            } else {
                                var v32 float32 = 1e10
                                p0 = p0
                                var v33 float32 = 1e10
                                v32 -= 1e10
                                fmt.Println(v33, v33)
                                _ = v32
                                _ = v33
                            }
                        }
                    }
                    v34 := []int{(fi26 - (fi26 << ((575) & 63)))}
                    _ = fi26
                    _ = v34
                }
                var v35 uint8 = 255
                _ = v35
            } else {
                if (true || (float32((((0.1 * 1.5) * 0.0) + 1.5)) <= 0.1)) {
                    v36 := []int{-248}
                    fmt.Println(v24, p1)
                    if (0 < p0) {
                        var v37 bool = ((3.14159 - (0.0 * 1.5)) != 0.1)
                        v38 := GenS4{F0: "", F1: (p1 * 42), F2: 217}
                        _ = v37
                        _ = v38
                    } else {
                        if len(v36) >= 2 {
                            v36 = v36[1:len(v36)]
                        }
                        v24 -= (-(v24))
                    }
                    _ = v36
                } else {
                    v39 := GenS4{F0: ("foo bar" + ("\n\t" + "")), F1: p1, F2: uint8(p1)}
                    v24 += 224
                    for fi40 := 0; fi40 < 18; fi40++ {
                        v41 := []int{fi23}
                        var v42 int32 = -1
                        p0 -= (p0 % (((p0 - (p0 % ((42) | 1)))) | 1))
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    _ = v39
                }
                p0 = p0
            }
            _ = fi23
            _ = v24
        }
    }
    _ = p0
    _ = p1
    return genFn0()
}

func main() {
    v1 := []string{((("a" + "hello") + "hello") + ("hello" + ("a" + "a"))), "foo bar"}
    v2 := make(map[uint32]uint32)
    var v3 int64 = -104
    var v4 int64 = v3
    if ((v3 & v4) <= v4) {
        v43 := GenS1{F0: 255, F1: 32767}
        _ = v43
    } else {
        switch 32768 {
        case -1:
            for fi44 := 0; fi44 < 5; fi44++ {
                var v45 float64 = 0.1
                _ = fi44
                _ = v45
            }
        case 128:
            for fi46 := 0; fi46 < 6; fi46++ {
                for fi47 := 0; fi47 < 17; fi47++ {
                    var v48 uint = 18446744073709551615
                    fmt.Println(fi46)
                    fmt.Println(v48, v48)
                    _ = fi47
                    _ = v48
                }
                v4 = v3
                var v49 uint16 = 65535
                _ = fi46
                _ = v49
            }
        case 32768:
            switch "foo bar" {
            case "":
                v3 = v3
            default:
                var v50 int8 = 21
                _ = v50
            }
        }
        v51 := GenS1{F0: 4294967295, F1: -803}
        _ = v51
    }
    if (!true) {
        v4 -= (v3 | (v4 << ((42) & 63)))
        v2[2147483648] = 269
        var v52 int32 = 42
        _ = v52
    } else {
        var v53 string = ""
        _ = v53
    }
    var v54 int64 = v4
    v55 := GenS2{F0: 255, F1: 2}
    v4 = (v4 ^ 0)
    v2[42] = 42
    for fi56 := 0; fi56 < 20; fi56++ {
        delete(v2, 42)
        if ((4294967295 >= 2) || (255 > 800)) {
            var v57 uint8 = 0
            _ = v57
        }
        if (4 == 2) {
            for fi58 := 0; fi58 < 19; fi58++ {
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
                v59 := GenS5{F0: false, F1: 1, F2: 65535}
                _ = fi58
                _ = v59
            }
        }
        _ = fi56
    }
    fmt.Println(v4)
    var v60 int8 = 0
    delete(v2, 2147483648)
    var v61 int64 = (v4 * v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v55
    _ = v60
    _ = v61
}
