package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint64; F1 int; F2 bool }
type GenS4 struct { F0 int32 }
func genFn0() int32 {
    var v4 uint16 = 84
    v5 := GenS1{F0: -1}
    v4 = (v4 ^ (v4 % (((^(v4))) | 1)))
    _ = v4
    _ = v5
    return -1
}
func genFn1() int {
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := []int{32767, (fi6 ^ (-(fi6)))}
        var v8 int64 = 160
        _ = fi6
        _ = v7
        _ = v8
    }
    if (!false) {
        var v9 uint16 = 42
        v9 += (v9 * 2)
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := GenS1{F0: 668}
            var v12 float64 = float64(1e10)
            if true {
                if (true && true) {
                    v12 += v12
                    var v13 uint = 65535
                    _ = v13
                }
                switch fi10 {
                case 127:
                    fmt.Println(v12, v12)
                case 128:
                    if (2 >= 42) {
                        var v14 uint64 = 0
                        _ = v14
                    } else {
                        v12 += float64((1e10 - ((1.5 * 1.5) - -2.25)))
                        var v15 int32 = -2147483648
                        _ = v15
                    }
                case 430:
                    var v16 int32 = genFn0()
                    _ = v16
                }
                var v17 uint32 = 406
                _ = v17
            } else {
                var v18 int32 = -825
                switch fi10 {
                case -2147483648:
                    v9 = (v9 ^ v9)
                case 128:
                    switch "" {
                    case "foo bar":
                        var v19 int8 = 32
                        _ = v19
                    case "":
                        for fi20 := 0; fi20 < 5; fi20++ {
                            var v21 uint8 = 36
                            v22 := GenS2{F0: 39, F1: 42}
                            var v23 uint16 = v9
                            _ = fi20
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                    case "a":
                        v24 := GenS3{F0: 0, F1: 32767, F2: (true && false)}
                        _ = v24
                    }
                case -9223372036854775808:
                    v12 -= ((0.1 * v12) * (((1e10 * (v12 - v12)) - v12) * (v12 - 1.5)))
                }
                if false {
                    v18 = genFn0()
                } else {
                    fmt.Println(v9, v12)
                    switch fi10 {
                    case 2147483647:
                        var v25 int64 = 9223372036854775807
                        _ = v25
                    case 32768:
                        var v26 int32 = v18
                        _ = v26
                    }
                    v27 := []int{(fi10 & (-(fi10)))}
                    _ = v27
                }
                _ = v18
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v9
    }
    v28 := []int{128}
    var v29 float64 = ((3.14159 * 1.5) - float64((float32((0.0 / 1.5)) - (float32(0.1) - 1e10))))
    fmt.Println(v29)
    _ = v28
    _ = v29
    return -281
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    var v3 uint8 = 255
    for fi30 := 0; fi30 < 9; fi30++ {
        v31 := []int{-30}
        for fi32 := 0; fi32 < 4; fi32++ {
            if ((fi32 >> ((0) & 63)) <= genFn1()) {
                if true {
                    v3 += 29
                }
                v33 := GenS2{F0: 771, F1: 1}
                var v34 uint16 = 65535
                _ = v33
                _ = v34
            } else {
                v3 += 127
            }
            fmt.Println(v3, fi32)
            _ = fi32
        }
        _ = fi30
        _ = v31
    }
    if (18446744073709551615 < 171) {
        v1[2147483647] = true
        v3 -= (v3 & 42)
    }
    if true {
        if (!false) {
            v35 := v1[984]
            fmt.Println(v35, v3)
            _ = v35
        } else {
            if true {
                fmt.Println(v3)
                switch -402 {
                case -268:
                    if ((true && (!(false && true))) || (1 < -292)) {
                        fmt.Println(v3)
                    }
                case 42:
                    var v36 int64 = 9223372036854775807
                    _ = v36
                default:
                    delete(v1, 748)
                }
                var v37 bool = (127 < genFn1())
                _ = v37
            }
            v3 = (v3 + v3)
        }
        v3 += 127
    } else {
        var v38 int = 32768
        var v39 float64 = float64((1e10 + (0.0 * 3.14159)))
        if ((!true) || false) {
            for fi40 := 0; fi40 < 7; fi40++ {
                v38 += (v38 & genFn1())
                _ = fi40
            }
        } else {
            if true {
                for fi41 := 0; fi41 < 10; fi41++ {
                    v42 := []int{42, (v38 + genFn1()), genFn1(), 2147483647}
                    fmt.Println(fi41)
                    _ = fi41
                    _ = v42
                }
                for fi43 := 0; fi43 < 11; fi43++ {
                    fmt.Println(v3, v3, v3)
                    v44 := GenS3{F0: 0, F1: v38, F2: (!(!true))}
                    v45 := GenS1{F0: -9223372036854775808}
                    _ = fi43
                    _ = v44
                    _ = v45
                }
            } else {
                v3 = 128
                delete(v2, v38)
            }
        }
        _ = v38
        _ = v39
    }
    for fi46 := 0; fi46 < 1; fi46++ {
        if true {
            v47 := GenS4{F0: 0}
            if (!true) {
                switch (fi46 / (((fi46 | (fi46 / ((88) | 1)))) | 1)) {
                case 4294967295:
                    for fi48 := 0; fi48 < 8; fi48++ {
                        var v49 bool = (4294967295 > 42)
                        if false {
                            var v50 int = (fi46 << ((42) & 63))
                            _ = v50
                        }
                        _ = fi48
                        _ = v49
                    }
                case 824:
                    if (42 >= 2) {
                        var v51 uint32 = 20
                        for fi52 := 0; fi52 < 16; fi52++ {
                            v53 := GenS3{F0: 42, F1: 0, F2: true}
                            var v54 uint8 = v3
                            _ = fi52
                            _ = v53
                            _ = v54
                        }
                        for fi55 := 0; fi55 < 8; fi55++ {
                            v56 := GenS1{F0: -460}
                            _ = fi55
                            _ = v56
                        }
                        _ = v51
                    } else {
                        fmt.Println(fi46, fi46)
                        for fi57 := 0; fi57 < 8; fi57++ {
                            var v58 uint32 = 1
                            v58 = v58
                            _ = fi57
                            _ = v58
                        }
                    }
                }
                v59 := []int{(fi46 << ((255) & 63))}
                _ = v59
            } else {
                v3 = 42
                for fi60 := 0; fi60 < 20; fi60++ {
                    var v61 uint64 = 2
                    v3 = v3
                    var v62 bool = false
                    delete(v1, 2147483648)
                    var v63 bool = ((0 < 255) || (!v62))
                    _ = fi60
                    _ = v61
                    _ = v62
                    _ = v63
                }
            }
            if (!(true || (161 != 42))) {
                var v64 uint = 2
                v3 -= (^(v3))
                _ = v64
            }
            _ = v47
        }
        if true {
            var v65 int32 = 137
            _ = v65
        }
        _ = fi46
    }
    var v66 int = 127
    var v67 int64 = 42
    delete(v1, 670)
    v66 = (v66 / ((596) | 1))
    v68 := v2[(v66 >> ((764) & 63))]
    v69 := v1[(v68 ^ (v68 ^ (v68 / ((2147483647) | 1))))]
    var v70 bool = false
    var v71 float64 = float64(1e10)
    for fi72 := 0; fi72 < 12; fi72++ {
        var v73 string = ((("a" + "hello") + "foo bar") + "\n\t")
        var v74 uint16 = 65535
        _ = fi72
        _ = v73
        _ = v74
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v70
    _ = v71
}
