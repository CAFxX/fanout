package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint32; F1 int; F2 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 int32; F1 uint32; F2 bool }
type GenS8 struct { F0 uint32; F1 uint32; F2 bool }
func genFn0() int64 {
    if false {
        switch -622 {
        case 2147483647:
            v5 := GenS1{F0: 127}
            _ = v5
        case 32768:
            var v6 int32 = 2147483647
            _ = v6
        case -1:
            var v7 uint = 65535
            _ = v7
        default:
            switch ((("" + "") + "") + ("foo bar" + "foo bar")) {
            case "":
                var v8 uint = 708
                _ = v8
            case "\n\t":
                v9 := GenS1{F0: 1}
                _ = v9
            default:
                var v10 int32 = 981
                _ = v10
            }
        }
    }
    if (128 != 116) {
        if false {
            switch 32768 {
            case -1:
                for fi11 := 0; fi11 < 12; fi11++ {
                    v12 := []int{-911, (fi11 * 128), (-(fi11))}
                    _ = fi11
                    _ = v12
                }
            }
        } else {
            v13 := GenS2{F0: 1, F1: 2, F2: 4294967295}
            _ = v13
        }
        var v14 uint64 = 18446744073709551615
        switch 22 {
        case -884:
            v14 -= 2
        case 319:
            fmt.Println(v14)
        case -1:
            fmt.Println(v14)
        }
        _ = v14
    } else {
        var v15 int8 = -128
        if (!true) {
            v16 := []int{-1, -2147483648}
            if ((127 > -774) && true) {
                switch "\n\t" {
                case "\n\t":
                    for fi17 := 0; fi17 < 18; fi17++ {
                        var v18 uint32 = 42
                        _ = fi17
                        _ = v18
                    }
                case "a":
                    if len(v16) > 0 {
                        v16[0] = -2147483648
                    }
                case "":
                    v19 := GenS2{F0: 916, F1: 206, F2: 18446744073709551615}
                    _ = v19
                }
                v20 := GenS3{F0: 42, F1: 1}
                _ = v20
            }
            var v21 bool = (!true)
            _ = v16
            _ = v21
        }
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 bool = true
            _ = fi22
            _ = v23
        }
        _ = v15
    }
    v24 := []int{4294967295}
    var v25 bool = false
    var v26 int16 = 32767
    _ = v24
    _ = v25
    _ = v26
    return 1
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]int)
    var v3 uint32 = v1
    var v4 uint8 = 128
    fmt.Println(v1)
    for fi27 := 0; fi27 < 20; fi27++ {
        var v28 uint = 1
        if true {
            v29 := v2[fi27]
            for fi30 := 0; fi30 < 5; fi30++ {
                fmt.Println(fi30, v28, v4)
                _ = fi30
            }
            var v31 float32 = 0.0
            _ = v29
            _ = v31
        } else {
            for fi32 := 0; fi32 < 17; fi32++ {
                if false {
                    fmt.Println(v4)
                    switch fi27 {
                    case 2147483647:
                        if true {
                            var v33 int32 = 150
                            delete(v2, (fi32 & (fi32 << (((^(v28))) & 63))))
                            switch "a" {
                            case "a":
                                v4 += 128
                            case "\n\t":
                                var v34 int32 = v33
                                _ = v34
                            case "":
                                switch "\n\t" {
                                case "foo bar":
                                    _ = len(v2)
                                case "hello":
                                    var v35 uint64 = 1
                                    _ = v35
                                }
                            }
                            var v36 int16 = -923
                            var v37 int32 = 1
                            _ = v33
                            _ = v36
                            _ = v37
                        } else {
                            var v38 int8 = -1
                            if ((!(-1 > (fi32 / (((fi32 / ((42) | 1))) | 1)))) || (42 != (v38 * 100))) {
                                var v39 int16 = 453
                                _ = v39
                            } else {
                                fmt.Println(v38, v28, fi27)
                                var v40 uint = uint(v1)
                                var v41 uint32 = (v3 & 1)
                                _ = v40
                                _ = v41
                            }
                            v38 = (-(v38))
                            _ = v38
                        }
                    default:
                        switch (-(fi32)) {
                        case 32768:
                            if (uint16(v28) < 42) {
                                v1 += v3
                                _ = len(v2)
                                if false {
                                    v42 := v2[(fi32 | -293)]
                                    _ = v42
                                }
                            }
                        }
                    }
                    var v43 int16 = -32768
                    _ = v43
                } else {
                    var v44 int8 = -61
                    v28 = (v28 & uint(v1))
                    v1 -= v3
                    _ = v44
                }
                var v45 bool = ("\n\t" != "hello")
                _ = fi32
                _ = v45
            }
            v28 = 0
            _ = len(v2)
        }
        _ = fi27
        _ = v28
    }
    for fi46 := 0; fi46 < 18; fi46++ {
        for fi47 := 0; fi47 < 14; fi47++ {
            var v48 uint8 = (v4 | 2)
            v2[fi47] = fi47
            v3 = 720
            _ = fi47
            _ = v48
        }
        _ = fi46
    }
    v4 += (v4 / ((127) | 1))
    v4 += 2
    v49 := GenS4{F0: v1, F1: -2147483648, F2: 2}
    v50 := GenS5{F0: -1}
    if ((v3 - 1) != (v3 - v1)) {
        v51 := []int{128, 1, -220, 32767}
        if false {
            for fi52 := 0; fi52 < 6; fi52++ {
                if (false && true) {
                    v1 += 792
                    var v53 bool = (42 == 102)
                    _ = v53
                }
                for fi54 := 0; fi54 < 20; fi54++ {
                    var v55 uint = 65535
                    _ = fi54
                    _ = v55
                }
                _ = fi52
            }
            var v56 int = 853
            fmt.Println(v4)
            _ = v56
        } else {
            if false {
                fmt.Println(v3, v1)
                if (!false) {
                    for fi57 := 0; fi57 < 9; fi57++ {
                        if len(v51) > 0 {
                            v51[0] = (fi57 >> ((42) & 63))
                        }
                        _ = fi57
                    }
                } else {
                    if len(v51) > 0 {
                        v51[0] = -1
                    }
                    v58 := GenS6{F0: 81, F1: 547}
                    _ = v58
                }
                for fi59 := 0; fi59 < 8; fi59++ {
                    v51 = append(v51, (fi59 * fi59))
                    switch ("foo bar" + "a") {
                    case "":
                        v1 += (v1 - v1)
                    default:
                        v4 -= 42
                    }
                    _ = fi59
                }
            } else {
                v1 -= 2
                for fi60 := 0; fi60 < 12; fi60++ {
                    fmt.Println(v3, fi60, fi60)
                    _ = fi60
                }
            }
            var v61 uint16 = 2
            fmt.Println(v61)
            _ = v61
        }
        var v62 uint16 = uint16(v4)
        _ = v51
        _ = v62
    }
    for fi63 := 0; fi63 < 11; fi63++ {
        v64 := []int{(fi63 - fi63), fi63}
        _ = fi63
        _ = v64
    }
    v65 := GenS7{F0: -816, F1: 4294967295, F2: false}
    fmt.Println(v1)
    if (42 < -128) {
        v66 := []int{755, -2147483648}
        v4 -= v4
        switch 9223372036854775807 {
        case 0:
            v1 = (^(v3))
        case 568:
            v4 = (^(v4))
        }
        _ = v66
    } else {
        var v67 uint64 = 1
        _ = v67
    }
    v68 := GenS8{F0: v3, F1: (^(v3)), F2: false}
    for fi69 := 0; fi69 < 18; fi69++ {
        v3 = v3
        v2[int(v4)] = -402
        _ = fi69
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v50
    _ = v65
    _ = v68
}
