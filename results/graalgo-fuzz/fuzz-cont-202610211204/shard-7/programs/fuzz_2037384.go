package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int64; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint; F2 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16 }
func genFn0() uint32 {
    switch "a" {
    case "a":
        v7 := []int{}
        _ = v7
    case "":
        var v8 uint = 256
        _ = v8
    }
    var v9 int = 1
    for fi10 := 0; fi10 < 8; fi10++ {
        fmt.Println(v9)
        _ = fi10
    }
    fmt.Println(v9)
    var v11 int64 = -110
    _ = v11
    _ = v9
    return 1
}
func genFn1(p0 int32, p1 bool) int16 {
    if (42 >= 0) {
        fmt.Println(p0, p0)
    } else {
        var v12 int8 = 100
        _ = v12
    }
    if (p1 || ((p1 && (127 >= 42)) || true)) {
        var v13 uint32 = 2
        v14 := GenS1{F0: 445, F1: -32768}
        _ = v13
        _ = v14
    } else {
        v15 := GenS1{F0: 882, F1: 1000}
        v16 := []int{}
        for fi17 := 0; fi17 < 6; fi17++ {
            if (p1 || false) {
                v16 = append(v16, fi17)
                for fi18 := 0; fi18 < 8; fi18++ {
                    switch (fi18 ^ (fi17 * int(p0))) {
                    case 569:
                        var v19 uint = 89
                        _ = v19
                    case -710:
                        fmt.Println(fi17, p0, fi18)
                    default:
                        v16 = append(v16, (fi18 & (fi17 % ((356) | 1))))
                    }
                    _ = fi18
                }
                _ = len(v16)
            } else {
                switch (fi17 << ((763) & 63)) {
                case 9223372036854775807:
                    if true {
                        for fi20 := 0; fi20 < 2; fi20++ {
                            v16 = append(v16, (-(fi17)))
                            var v21 int32 = -1
                            v22 := GenS1{F0: 1, F1: 1}
                            v23 := GenS1{F0: 42, F1: -857}
                            var v24 uint64 = 0
                            p1 = (0 == 806)
                            _ = fi20
                            _ = v21
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                    }
                }
                p1 = (float64(((float32(((3.14159 / 1.0) - 1e10)) / 1.5) - 1.5)) < ((((((1e10 * ((-2.25 + (-2.25 * -2.25)) + (3.14159 - 0.0))) - 1.5) / 0.1) - (3.14159 * 1e10)) * float64(1.5)) + 3.14159))
            }
            if p1 {
                for fi25 := 0; fi25 < 19; fi25++ {
                    var v26 int = (fi17 ^ fi17)
                    _ = fi25
                    _ = v26
                }
            } else {
                if false {
                    fmt.Println(p0, p0, p1)
                    var v27 int64 = 1
                    fmt.Println(fi17)
                    _ = v27
                }
            }
            _ = fi17
        }
        _ = v15
        _ = v16
    }
    p1 = (!(0 > 106))
    var v28 int8 = 6
    _ = p0
    _ = p1
    _ = v28
    return -1
}

func main() {
    v1 := []uint32{65535, 0}
    v2 := make(map[int]uint32)
    var v3 uint8 = 42
    var v4 uint = uint(v3)
    var v5 uint64 = 18446744073709551615
    var v6 int32 = -604
    for fi29 := 0; fi29 < 4; fi29++ {
        v4 = v4
        var v30 uint = (v4 % ((v4) | 1))
        _ = fi29
        _ = v30
    }
    v6 += v6
    for fi31 := 0; fi31 < 16; fi31++ {
        if (false || (fi31 >= (fi31 ^ fi31))) {
            _ = len(v2)
            v32 := GenS2{F0: 42, F1: 65535, F2: -702}
            switch fi31 {
            case 560:
                if false {
                    switch "hello" {
                    case "":
                        v33 := GenS1{F0: (v4 - 743), F1: -1}
                        _ = v33
                    case "hello":
                        if true {
                            v4 -= v4
                        } else {
                            var v34 int8 = -62
                            for fi35 := 0; fi35 < 18; fi35++ {
                                var v36 uint16 = uint16(v4)
                                if (-2.25 != (1e10 - 0.0)) {
                                    switch fi31 {
                                    case 32768:
                                        v5 += v5
                                    default:
                                        v4 += 255
                                    }
                                }
                                _ = fi35
                                _ = v36
                            }
                            _ = v34
                        }
                    case "a":
                        fmt.Println(v5, v3)
                    default:
                        if (false && (!true)) {
                            v37 := GenS3{F0: v3, F1: (v4 ^ v4), F2: (v4 / ((v4) | 1))}
                            v4 = v4
                            var v38 uint = uint(v5)
                            _ = v37
                            _ = v38
                        }
                    }
                    var v39 int64 = 0
                    for fi40 := 0; fi40 < 18; fi40++ {
                        v41 := GenS4{F0: 4294967295}
                        _ = fi40
                        _ = v41
                    }
                    _ = v39
                } else {
                    v3 += (v3 + 255)
                    if true {
                        v42 := GenS5{F0: 42}
                        var v43 uint8 = v3
                        _ = v42
                        _ = v43
                    } else {
                        v44 := GenS2{F0: 534, F1: 42, F2: -1}
                        v45 := GenS1{F0: uint(v5), F1: 1}
                        _ = v44
                        _ = v45
                    }
                    for fi46 := 0; fi46 < 10; fi46++ {
                        var v47 uint16 = 65535
                        v48 := GenS5{F0: 32767}
                        var v49 uint = uint(v5)
                        _ = fi46
                        _ = v47
                        _ = v48
                        _ = v49
                    }
                }
            case 4294967295:
                v5 = (v5 | v5)
            }
            _ = v32
        } else {
            v50 := GenS5{F0: -613}
            _ = v50
        }
        var v51 uint16 = 0
        _ = fi31
        _ = v51
    }
    fmt.Println(v5, v4, v6)
    switch "hello" {
    case "foo bar":
        for fi52 := 0; fi52 < 17; fi52++ {
            var v53 uint32 = 4294967295
            var v54 float32 = ((0.1 - 0.0) - (1e10 - 3.14159))
            var v55 int16 = 1
            _ = fi52
            _ = v53
            _ = v54
            _ = v55
        }
    default:
        var v56 uint32 = uint32(v4)
        _ = v56
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v57 int = -391
    for fi58 := 0; fi58 < 19; fi58++ {
        if (!true) {
            fmt.Println(v6)
            v3 = 128
        } else {
            v4 -= 67
            v59 := GenS4{F0: v5}
            v2[(v57 & (fi58 >> (((v4 | 2)) & 63)))] = genFn0()
            _ = v59
        }
        _ = fi58
    }
    var v60 string = ("a" + "")
    v60 = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v60
}
