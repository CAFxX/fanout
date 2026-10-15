package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int8; F1 string; F2 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64 }
func genFn0() int {
    var v7 float64 = 0.0
    var v8 uint16 = 32767
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 1
    v3 := make(map[uint32]string)
    var v4 int = -2147483648
    var v5 uint = 2
    var v6 float32 = float32((1e10 + 0.1))
    var v9 int = (v4 * v4)
    if (551 >= 65535) {
        v10 := GenS1{F0: 32767, F1: -700}
        if (-30 != 124) {
            v11 := GenS1{F0: uint16(v2), F1: 0}
            _ = v11
        } else {
            v12 := GenS2{F0: -8, F1: v1, F2: -32768}
            fmt.Println(v4)
            _ = v12
        }
        _ = v10
    } else {
        v4 -= v4
        for fi13 := 0; fi13 < 7; fi13++ {
            if true {
                var v14 bool = (true && (!false))
                v9 = 514
                v4 += 0
                _ = v14
            }
            var v15 int16 = 32767
            v2 = (v2 | (v2 - v2))
            _ = fi13
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 bool = true
        _ = fi16
        _ = v17
    }
    fmt.Println(v1, v1, v1)
    for fi18 := 0; fi18 < 9; fi18++ {
        for fi19 := 0; fi19 < 9; fi19++ {
            if false {
                switch (-(fi18)) {
                case 0:
                    fmt.Println(v4, v6)
                case 32767:
                    if (-32768 < 42) {
                        switch v1 {
                        case "":
                            v20 := GenS3{F0: 2147483647}
                            var v21 uint = (^(v5))
                            v5 = (v21 + (v5 % (((v21 & (v21 << ((v5) & 63)))) | 1)))
                            _ = v20
                            _ = v21
                        case "hello":
                            if (false || true) {
                                if (42 == 32767) {
                                    if (!false) {
                                        v22 := []int{genFn0(), (-(fi19)), fi19}
                                        v9 = fi19
                                        _ = v22
                                    }
                                }
                                var v23 int = (v4 % (((fi19 >> (((v5 << ((v5) & 63))) & 63))) | 1))
                                var v24 int16 = -548
                                _ = v23
                                _ = v24
                            }
                        case "a":
                            if true {
                                if (true && false) {
                                    v25 := []int{128, v4, (fi19 % ((-944) | 1))}
                                    if true {
                                        _ = len(v3)
                                        fmt.Println(v6, v2)
                                        v5 -= (v5 & 562)
                                    }
                                    fmt.Println(v5)
                                    v1 = "a"
                                    _ = v25
                                } else {
                                    v26 := GenS4{F0: 646, F1: 0}
                                    v5 = (v5 - v5)
                                    _ = v26
                                }
                                switch (fi18 % ((771) | 1)) {
                                case -266:
                                    var v27 uint64 = 0
                                    _ = v27
                                case -10:
                                    if (-9223372036854775808 == 1) {
                                        v28 := GenS5{F0: (("" + v1) + v1)}
                                        if false {
                                            var v29 int8 = -1
                                            if (42 < (-(v29))) {
                                                switch v1 {
                                                case "\n\t":
                                                    v30 := GenS3{F0: 398}
                                                    _ = v30
                                                case "":
                                                    fmt.Println(fi19, v5, v6)
                                                    if true {
                                                        v31 := []int{(-(fi19)), -789, (fi18 ^ (fi18 | (v4 >> ((v5) & 63))))}
                                                        fmt.Println(fi19, v4)
                                                        var v32 uint64 = 18446744073709551615
                                                        v29 += (v29 ^ 0)
                                                        v1 = (v1 + v1)
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                case "a":
                                                    v33 := []int{(fi18 + v9), int(v5)}
                                                    _ = v33
                                                }
                                            } else {
                                                v1 = v1
                                                v6 = (0.1 + v6)
                                                v9 -= fi19
                                            }
                                            _ = v29
                                        }
                                        _ = v28
                                    }
                                    var v34 int8 = 42
                                    _ = v34
                                case 4294967295:
                                    v35 := GenS6{F0: 0}
                                    _ = v35
                                default:
                                    var v36 uint32 = 594
                                    _ = v36
                                }
                                var v37 uint32 = 65535
                                _ = v37
                            } else {
                                v9 = genFn0()
                            }
                        default:
                            if false {
                                var v38 float64 = 1.5
                                v3[0] = (v1 + v1)
                                _ = len(v3)
                                _ = v38
                            } else {
                                var v39 uint64 = uint64(v5)
                                v40 := GenS3{F0: 588}
                                v9 -= v9
                                _ = v39
                                _ = v40
                            }
                        }
                    } else {
                        if false {
                            v2 += (v2 >> ((v5) & 63))
                            var v41 int16 = 898
                            v42 := GenS6{F0: 341}
                            _ = v41
                            _ = v42
                        }
                        var v43 uint8 = 42
                        v44 := []int{fi18, v9}
                        _ = v43
                        _ = v44
                    }
                default:
                    v9 = 1
                }
                var v45 bool = false
                var v46 string = (v1 + (("foo bar" + v1) + v1))
                _ = v45
                _ = v46
            }
            switch v1 {
            case "\n\t":
                var v47 bool = ((v2 & uint8(v5)) == 255)
                _ = v47
            case "hello":
                v48 := v3[2147483647]
                _ = v48
            case "a":
                var v49 uint = 574
                _ = v49
            }
            _ = fi19
        }
        _ = fi18
    }
    var v50 uint32 = uint32(v2)
    v51 := []int{(v4 & v9), 47, (v9 ^ int(v50))}
    v3[v50] = v1
    if len(v51) >= 2 {
        v51 = v51[1:len(v51)]
    }
    if true {
        var v52 int32 = 42
        _ = v52
    } else {
        v53 := GenS1{F0: 65535, F1: -1}
        var v54 uint16 = uint16(v50)
        _ = v53
        _ = v54
    }
    for fi55 := 0; fi55 < 16; fi55++ {
        _ = len(v3)
        if (false || true) {
            fmt.Println(v1)
            for fi56 := 0; fi56 < 1; fi56++ {
                fmt.Println(fi55, v50)
                var v57 uint32 = v50
                fmt.Println(fi55, v1)
                _ = fi56
                _ = v57
            }
        } else {
            switch "a" {
            case "hello":
                if (-9223372036854775808 > 42) {
                    delete(v3, v50)
                    delete(v3, (v50 | v50))
                }
            default:
                if (v50 > (v50 >> ((v5) & 63))) {
                    var v58 int16 = -32768
                    fmt.Println(v9, v58, v6)
                    _ = v58
                }
            }
            var v59 uint16 = uint16(v5)
            _ = v59
        }
        _ = fi55
    }
    v60 := v3[(^(v50))]
    var v61 uint16 = 65535
    v60 = (("a" + v1) + v60)
    if true {
        if false {
            var v62 bool = false
            _ = v62
        } else {
            v6 += v6
            fmt.Println(v2, v60, v50)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v60
    _ = v61
    _ = v9
}
