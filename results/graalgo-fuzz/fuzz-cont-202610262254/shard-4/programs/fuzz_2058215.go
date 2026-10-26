package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint }
func genFn0(p0 uint32) uint {
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int = (fi7 | -808)
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 uint32 = p0
            v8 -= fi7
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    v11 := GenS1{F0: 127, F1: 0, F2: 18446744073709551615}
    var v12 float64 = -2.25
    if false {
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 int = fi13
            v15 := []int{}
            v16 := GenS1{F0: 0, F1: -802, F2: 1}
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
    } else {
        if false {
            for fi17 := 0; fi17 < 10; fi17++ {
                if ((^(p0)) < p0) {
                    var v18 float64 = (1e10 * v12)
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 9; fi19++ {
                        fmt.Println(p0, fi19, v12)
                        var v20 float32 = 1e10
                        var v21 uint64 = 0
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    if true {
                        p0 -= p0
                    } else {
                        v12 = v12
                        fmt.Println(p0, v12, v12)
                    }
                    var v22 int = (fi17 / (((-(fi17))) | 1))
                    _ = v22
                }
                if true {
                    fmt.Println(v12)
                    v12 -= float64(0.0)
                }
                var v23 uint = 685
                _ = fi17
                _ = v23
            }
        }
    }
    _ = p0
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 uint16 = 1
    var v2 int16 = 32767
    var v3 uint64 = 42
    v4 := make(map[uint32]string)
    var v5 uint16 = v1
    var v6 uint8 = uint8(v3)
    if true {
        var v24 uint64 = (^(v3))
        var v25 uint = 539
        _ = v24
        _ = v25
    } else {
        fmt.Println(v2)
        v26 := GenS1{F0: v6, F1: 693, F2: 18446744073709551615}
        for fi27 := 0; fi27 < 16; fi27++ {
            v28 := v4[42]
            delete(v4, 552)
            v29 := []int{-304, fi27, fi27}
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = v26
    }
    var v30 uint = uint(v3)
    for fi31 := 0; fi31 < 9; fi31++ {
        v32 := []int{}
        _ = fi31
        _ = v32
    }
    switch -2147483648 {
    case 9223372036854775807:
        if false {
            var v33 bool = true
            _ = v33
        } else {
            var v34 int8 = 97
            _ = v34
        }
    case 4294967295:
        var v35 int16 = int16(v6)
        _ = v35
    case 697:
        fmt.Println(v2, v2)
    }
    var v36 uint64 = (^(v3))
    var v37 int16 = 1
    v3 += 336
    for fi38 := 0; fi38 < 19; fi38++ {
        var v39 int32 = 622
        if false {
            v40 := []int{fi38, 1, fi38, fi38}
            if len(v40) > 0 {
                v40[0] = (fi38 - 42)
            }
            _ = v40
        }
        _ = fi38
        _ = v39
    }
    v41 := GenS1{F0: 128, F1: 274, F2: (v30 >> ((genFn0(2147483648)) & 63))}
    var v42 int = -1
    for fi43 := 0; fi43 < 16; fi43++ {
        for fi44 := 0; fi44 < 5; fi44++ {
            if (1 >= 2147483647) {
                v45 := GenS1{F0: v6, F1: v42, F2: (v30 + (v30 ^ v30))}
                fmt.Println(v30)
                var v46 uint64 = 42
                v47 := GenS1{F0: 42, F1: (fi43 + (fi44 * -2147483648)), F2: (v30 | (v30 & v30))}
                _ = v45
                _ = v46
                _ = v47
            } else {
                v48 := []int{(v42 % (((fi44 << (((v30 % ((genFn0(725)) | 1))) & 63))) | 1))}
                v6 = (^(v6))
                if (!true) {
                    var v49 uint16 = (v5 + (v1 & 0))
                    var v50 bool = (0 == 67)
                    v51 := GenS1{F0: (v6 * 255), F1: fi43, F2: v30}
                    _ = v49
                    _ = v50
                    _ = v51
                } else {
                    switch -9223372036854775808 {
                    case -1:
                        var v52 uint16 = v5
                        _ = v52
                    case 2147483647:
                        var v53 uint8 = 1
                        _ = v53
                    default:
                        if (!false) {
                            var v54 int8 = -128
                            switch int(v6) {
                            case 128:
                                var v55 int = 32768
                                _ = v55
                            case 127:
                                if len(v48) > 0 {
                                    v48[0] = 748
                                }
                            }
                            _ = v54
                        }
                    }
                }
                _ = v48
            }
            v1 += (v1 + (v5 << ((v30) & 63)))
            _ = fi44
        }
        var v56 uint16 = (v5 >> (((v30 | v30)) & 63))
        _ = fi43
        _ = v56
    }
    v57 := GenS1{F0: (^(v6)), F1: (v42 % (((^(v42))) | 1)), F2: (v30 / ((v30) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v57
    _ = v6
}
