package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16; F1 bool }
func genFn0(p0 uint8, p1 int64, p2 uint64) uint {
    fmt.Println(p1)
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := GenS1{F0: p1}
        v8 := GenS2{F0: 2147483648, F1: 255}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint8 = 32
    fmt.Println(v9, p1, p0)
    v10 := []int{int(p2), 135}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 508
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = false
    var v3 uint = 42
    var v4 int = -610
    var v5 string = ("foo bar" + ("\n\t" + ("" + "a")))
    v2 = (115 < 255)
    v11 := GenS1{F0: 100}
    switch (v4 << ((v3) & 63)) {
    case 4294967295:
        var v12 uint8 = 231
        _ = v12
    case 536:
        for fi13 := 0; fi13 < 18; fi13++ {
            v4 = v4
            _ = fi13
        }
    case 127:
        switch "\n\t" {
        case "\n\t":
            v14 := []int{42}
            _ = v14
        }
    }
    if (v5 == "") {
        switch v4 {
        case -189:
            var v15 int16 = 1000
            _ = v15
        }
    } else {
        v16 := []int{v4, (v4 % (((-(v4))) | 1))}
        v1[65535] = (163 < -1)
        _ = v16
    }
    v17 := GenS1{F0: -864}
    if v2 {
        var v18 float32 = -2.25
        switch 322 {
        case 0:
            v4 = (^(v4))
        case -1:
            var v19 bool = v2
            _ = v19
        }
        switch 4294967295 {
        case 128:
            switch v4 {
            case -1:
                fmt.Println(v5, v3, v2)
            case -510:
                if (42 == -861) {
                    var v20 uint16 = 688
                    var v21 float32 = (float32((((((0.1 / ((0.0 + ((0.1 * 0.1) + 0.0)) * 0.1)) / float64(v18)) * ((1.5 * 1.5) / (float64(3.14159) * -2.25))) * (3.14159 - float64(v18))) / 1.0)) / ((-2.25 / v18) + float32(3.14159)))
                    _ = v20
                    _ = v21
                } else {
                    v4 = v4
                    switch (v5 + (v5 + ("\n\t" + v5))) {
                    case "foo bar":
                        if v2 {
                            if true {
                                for fi22 := 0; fi22 < 2; fi22++ {
                                    for fi23 := 0; fi23 < 5; fi23++ {
                                        var v24 bool = (42 == -1)
                                        v4 = (^(v4))
                                        fmt.Println(fi23, v3)
                                        _ = fi23
                                        _ = v24
                                    }
                                    _ = fi22
                                }
                                if (!(!false)) {
                                    var v25 uint16 = 32767
                                    for fi26 := 0; fi26 < 13; fi26++ {
                                        if v2 {
                                            if (v18 < (0.0 * (v18 + v18))) {
                                                v27 := GenS2{F0: 321, F1: (v3 >> (((v3 | v3)) & 63))}
                                                for fi28 := 0; fi28 < 18; fi28++ {
                                                    switch (((v5 + "\n\t") + (v5 + "a")) + "foo bar") {
                                                    case "":
                                                        fmt.Println(v25, fi26, v25)
                                                    }
                                                    if false {
                                                        v29 := []int{fi26, (fi28 / ((fi26) | 1))}
                                                        fmt.Println(v25)
                                                        fmt.Println(fi28, v3)
                                                        _ = v29
                                                    } else {
                                                        switch -2147483648 {
                                                        case 127:
                                                            v5 = ("" + "hello")
                                                        case 0:
                                                            v30 := GenS1{F0: 1}
                                                            _ = v30
                                                        case 32767:
                                                            v18 -= ((v18 + (-2.25 + v18)) - float32(1.5))
                                                        default:
                                                            v31 := v1[835]
                                                            _ = v31
                                                        }
                                                        if v2 {
                                                            v32 := GenS3{F0: v25}
                                                            var v33 int32 = 0
                                                            _ = v32
                                                            _ = v33
                                                        } else {
                                                            fmt.Println(v25)
                                                        }
                                                    }
                                                    var v34 uint8 = 0
                                                    _ = fi28
                                                    _ = v34
                                                }
                                                _ = v27
                                            }
                                            switch v4 {
                                            case 2147483647:
                                                v35 := v1[2147483647]
                                                _ = v35
                                            case 127:
                                                if false {
                                                    v3 = (v3 | v3)
                                                    v36 := GenS3{F0: 878}
                                                    _ = v36
                                                } else {
                                                    v4 = 32767
                                                    var v37 string = v5
                                                    _ = v37
                                                }
                                            case 815:
                                                var v38 int = (v4 + (fi26 % ((4294967295) | 1)))
                                                _ = v38
                                            default:
                                                var v39 string = (v5 + v5)
                                                _ = v39
                                            }
                                        }
                                        _ = fi26
                                    }
                                    v25 = 638
                                    _ = v25
                                } else {
                                    v40 := GenS4{F0: 467, F1: 42}
                                    _ = v40
                                }
                            }
                            fmt.Println(v2)
                            v41 := []int{(v4 & 2147483647), (v4 & (-(v4))), v4, v4}
                            _ = v41
                        }
                    case "a":
                        if v2 {
                            v42 := GenS4{F0: 255, F1: -1}
                            v4 += v4
                            _ = v42
                        }
                    default:
                        var v43 uint = v3
                        _ = v43
                    }
                    v2 = true
                }
            case 4294967295:
                if true {
                    for fi44 := 0; fi44 < 8; fi44++ {
                        fmt.Println(v3, v2)
                        _ = fi44
                    }
                } else {
                    if ((!(!(722 >= 558))) || (v4 >= (v4 - 179))) {
                        v3 = 69
                    } else {
                        var v45 int64 = 42
                        if v2 {
                            v45 += v45
                            v46 := []int{v4, v4, v4, (v4 / (((v4 | (v4 >> ((118) & 63)))) | 1))}
                            fmt.Println(v45, v3, v4)
                            _ = v46
                        }
                        v2 = (0.1 <= (1e10 * 1e10))
                        _ = v45
                    }
                    for fi47 := 0; fi47 < 16; fi47++ {
                        v48 := GenS5{F0: 0}
                        _ = len(v1)
                        _ = fi47
                        _ = v48
                    }
                }
            default:
                fmt.Println(v3)
            }
        default:
            v49 := GenS5{F0: -913}
            _ = v49
        }
        _ = v18
    } else {
        v50 := GenS1{F0: 42}
        for fi51 := 0; fi51 < 6; fi51++ {
            for fi52 := 0; fi52 < 2; fi52++ {
                v53 := []int{-565, v4, (v4 - 9223372036854775807), (v4 / ((fi51) | 1))}
                _ = len(v1)
                v54 := v1[640]
                _ = fi52
                _ = v53
                _ = v54
            }
            var v55 string = v5
            _ = fi51
            _ = v55
        }
        _ = v50
    }
    delete(v1, 65535)
    switch (v4 + v4) {
    case -9223372036854775808:
        var v56 bool = (true || false)
        _ = v56
    case 132:
        var v57 float64 = 0.1
        _ = v57
    case 4294967295:
        v5 = ("hello" + (v5 + (("\n\t" + ("\n\t" + (("" + v5) + v5))) + v5)))
    default:
        v3 += 538
    }
    var v58 uint = 18446744073709551615
    v5 = (("hello" + "\n\t") + (v5 + ("" + v5)))
    if true {
        var v59 bool = v2
        v58 = (v3 * 4294967295)
        _ = v59
    } else {
        v60 := []int{(v4 << (((^(v58))) & 63)), -1}
        _ = v60
    }
    var v61 uint32 = 939
    if ((-2.25 - 1e10) >= 1e10) {
        var v62 float64 = 0.1
        for fi63 := 0; fi63 < 5; fi63++ {
            for fi64 := 0; fi64 < 13; fi64++ {
                fmt.Println(v4, v62, v61)
                var v65 int8 = int8(v61)
                _ = fi64
                _ = v65
            }
            var v66 string = ("foo bar" + (v5 + (v5 + v5)))
            _ = fi63
            _ = v66
        }
        _ = v62
    } else {
        v5 = ""
        if v2 {
            for fi67 := 0; fi67 < 10; fi67++ {
                v68 := GenS6{F0: -128}
                fmt.Println(v4, v58)
                v69 := GenS7{F0: 0, F1: true}
                _ = fi67
                _ = v68
                _ = v69
            }
            v2 = v2
            if (!(0 > 2)) {
                v70 := []int{127, -2147483648, v4, v4}
                v61 -= (^(v61))
                if len(v70) >= 2 {
                    v70 = v70[1:len(v70)]
                }
                _ = v70
            } else {
                var v71 int32 = -2147483648
                fmt.Println(v71, v2)
                _ = v71
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v61
}
