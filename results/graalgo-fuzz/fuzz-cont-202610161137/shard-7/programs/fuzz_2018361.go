package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 uint64; F1 string; F2 uint8 }
type GenS8 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint32) int {
    p0 = p0
    v7 := GenS1{F0: 930, F1: p0}
    _ = p0
    _ = v7
    return 32768
}
func genFn1(p0 uint16, p1 int32, p2 int8) int8 {
    var v8 float64 = 1e10
    var v9 int16 = -818
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 uint32 = 1
    var v2 uint16 = 313
    var v3 bool = false
    var v4 uint = 2
    var v5 int16 = 1000
    var v6 int32 = 42
    switch 32767 {
    case 2147483647:
        var v10 uint32 = (v1 & (v1 / ((48) | 1)))
        _ = v10
    case 128:
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS1{F0: 45, F1: (v1 ^ (^(v1)))}
            var v13 uint8 = 128
            _ = fi11
            _ = v12
            _ = v13
        }
    case -332:
        for fi14 := 0; fi14 < 15; fi14++ {
            if ((v6 < v6) || v3) {
                v15 := []int{genFn0((v1 >> (((v4 ^ 2)) & 63)))}
                _ = v15
            } else {
                v16 := GenS1{F0: (v6 | -2147483648), F1: (v1 << ((v4) & 63))}
                _ = v16
            }
            _ = fi14
        }
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        v4 = v4
        if (1e10 <= 1.5) {
            switch fi17 {
            case -2147483648:
                fmt.Println(v2, fi17)
            case -642:
                v18 := GenS2{F0: (fi17 & 1)}
                _ = v18
            case 32767:
                if (127 > 42) {
                    for fi19 := 0; fi19 < 4; fi19++ {
                        var v20 int = fi19
                        if v3 {
                            switch (fi17 % ((32768) | 1)) {
                            case -2147483648:
                                v21 := GenS3{F0: uint32(v2)}
                                _ = v21
                            default:
                                var v22 int16 = (v5 >> (((v4 << ((0) & 63))) & 63))
                                v22 += v5
                                v5 = 637
                                v23 := []int{}
                                _ = v22
                                _ = v23
                            }
                        } else {
                            v24 := []int{(v20 ^ (fi17 + (fi17 - (v20 - (^(fi17)))))), -1, (fi17 + (fi19 % ((fi19) | 1)))}
                            _ = v24
                        }
                        fmt.Println(v4, v3, v20)
                        _ = fi19
                        _ = v20
                    }
                    v2 += v2
                    if v3 {
                        for fi25 := 0; fi25 < 6; fi25++ {
                            v26 := GenS2{F0: -1}
                            var v27 uint32 = v1
                            if true {
                                switch (fi17 << ((v4) & 63)) {
                                case 42:
                                    var v28 uint32 = (v27 - v27)
                                    _ = v28
                                case 2147483647:
                                    var v29 string = ""
                                    _ = v29
                                case 362:
                                    var v30 uint8 = 0
                                    _ = v30
                                default:
                                    if v3 {
                                        v31 := GenS4{F0: 128, F1: 1000, F2: 42}
                                        var v32 string = "foo bar"
                                        _ = v31
                                        _ = v32
                                    }
                                }
                                v33 := GenS5{F0: 0}
                                _ = v33
                            } else {
                                v3 = v3
                                v5 -= v5
                                v27 -= 403
                            }
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    }
                }
            default:
                var v34 uint32 = v1
                _ = v34
            }
        } else {
            var v35 int = 759
            _ = v35
        }
        v1 -= (v1 & v1)
        _ = fi17
    }
    var v36 int16 = int16(v2)
    for fi37 := 0; fi37 < 4; fi37++ {
        switch -1 {
        case 128:
            fmt.Println(v3, v5)
        case 785:
            v38 := []int{fi37, 140, fi37}
            _ = v38
        }
        _ = fi37
    }
    for fi39 := 0; fi39 < 3; fi39++ {
        for fi40 := 0; fi40 < 6; fi40++ {
            if ((^(v1)) > v1) {
                v2 -= (v2 << (((v4 + 65535)) & 63))
            } else {
                var v41 string = ("" + ("foo bar" + ("hello" + "hello")))
                v42 := GenS6{F0: 42, F1: genFn1((v2 + 65535), v6, int8(v2))}
                if ((!v3) && v3) {
                    var v43 string = v41
                    _ = v43
                }
                if false {
                    fmt.Println(v41)
                    switch genFn0(v1) {
                    case 128:
                        var v44 uint8 = 2
                        _ = v44
                    case 662:
                        v36 += (v5 ^ 32767)
                    case 127:
                        v41 = ((v41 + "hello") + v41)
                        v4 += v4
                    default:
                        fmt.Println(v5, v41, v6)
                    }
                    v45 := []int{genFn0(65535), fi39}
                    _ = v45
                }
                _ = v41
                _ = v42
            }
            var v46 uint = (^(v4))
            v2 += (v2 & v2)
            _ = fi40
            _ = v46
        }
        var v47 uint8 = 131
        for fi48 := 0; fi48 < 17; fi48++ {
            v49 := GenS7{F0: 4294967295, F1: "\n\t", F2: (^(v47))}
            v50 := []int{(fi39 ^ genFn0((v1 | (v1 + v1)))), (fi39 * int(v6)), -199, (fi39 % (((fi48 / (((fi48 ^ fi48)) | 1))) | 1))}
            v6 -= (v6 ^ v6)
            _ = fi48
            _ = v49
            _ = v50
        }
        _ = fi39
        _ = v47
    }
    var v51 uint32 = 1
    v52 := GenS5{F0: 42}
    for fi53 := 0; fi53 < 8; fi53++ {
        var v54 int8 = genFn1(v2, v6, 0)
        for fi55 := 0; fi55 < 20; fi55++ {
            v56 := []int{fi55}
            _ = fi55
            _ = v56
        }
        var v57 int32 = (v6 & 42)
        _ = fi53
        _ = v54
        _ = v57
    }
    v4 += 2
    if (v3 && ((true && (!v3)) && (!false))) {
        var v58 uint64 = 982
        var v59 int64 = -217
        _ = v58
        _ = v59
    }
    for fi60 := 0; fi60 < 10; fi60++ {
        v61 := GenS8{F0: 0, F1: 0}
        _ = fi60
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
