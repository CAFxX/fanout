package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint16, p1 uint, p2 uint32) int32 {
    v7 := []int{1, 1, -260}
    p1 = (p1 ^ p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 uint8 = 255
    var v2 uint = 255
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    var v5 int = -1
    v6 := []int{(^(v5)), (v5 + 127), v5, -2147483648}
    if (true && true) {
        var v8 uint32 = 2
        _ = v8
    } else {
        v2 = (v2 - 2)
    }
    var v9 int16 = -32768
    if true {
        var v10 float64 = 3.14159
        fmt.Println(v1, v2, v2)
        var v11 int16 = -32768
        _ = v10
        _ = v11
    }
    v5 += v5
    for fi12 := 0; fi12 < 9; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            v2 -= (^(v2))
            _ = fi13
        }
        fmt.Println(v2, v9, v1)
        if len(v6) > 0 {
            v6[0] = v5
        }
        _ = fi12
    }
    v5 -= v5
    if true {
        if true {
            var v14 int64 = 1
            fmt.Println(v14, v14)
            fmt.Println(v5, v9)
            _ = v14
        } else {
            v15 := GenS1{F0: (v5 * 42), F1: v9, F2: v1}
            var v16 uint = 19
            _ = v15
            _ = v16
        }
    } else {
        v17 := v4[(v5 & (-(v5)))]
        var v18 int8 = 100
        if v17 {
            var v19 bool = v17
            if (!((v19 && v19) || (!v17))) {
                v20 := GenS1{F0: 426, F1: -1, F2: 62}
                v21 := GenS1{F0: (-(v5)), F1: v9, F2: 128}
                _ = v20
                _ = v21
            } else {
                v22 := GenS2{F0: 64}
                var v23 int8 = (v18 + (v18 << ((v2) & 63)))
                _ = v22
                _ = v23
            }
            v1 = v1
            _ = v19
        }
        _ = v17
        _ = v18
    }
    var v24 uint64 = 897
    fmt.Println(v2, v24)
    v5 = (v5 / (((v5 | v5)) | 1))
    var v25 uint16 = 178
    for fi26 := 0; fi26 < 2; fi26++ {
        switch v5 {
        case -268:
            if (fi26 <= (fi26 % ((625) | 1))) {
                switch (v5 << (((v2 & 2)) & 63)) {
                case 632:
                    for fi27 := 0; fi27 < 11; fi27++ {
                        if len(v6) > 0 {
                            v6[0] = fi26
                        }
                        switch "\n\t" {
                        case "foo bar":
                            v1 = (v1 % ((v1) | 1))
                        case "hello":
                            v5 = (fi27 / ((fi27) | 1))
                        case "a":
                            var v28 string = ("hello" + "foo bar")
                            _ = v28
                        }
                        v6 = append(v6, (fi27 ^ (-(v5))))
                        switch (^(fi27)) {
                        case -649:
                            if len(v6) > 0 {
                                v6[0] = v5
                            }
                        case 2147483647:
                            v25 += 32767
                        default:
                            delete(v4, int(v24))
                        }
                        var v29 int = (fi27 / ((fi27) | 1))
                        _ = fi27
                        _ = v29
                    }
                case 127:
                    v30 := GenS3{F0: 1, F1: 1}
                    _ = v30
                case 42:
                    _ = len(v6)
                }
                v24 += (v24 % (((^(v24))) | 1))
                if false {
                    for fi31 := 0; fi31 < 7; fi31++ {
                        v32 := GenS4{F0: (v1 << (((v2 - v2)) & 63))}
                        var v33 uint8 = (v1 / ((10) | 1))
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                }
            } else {
                var v34 uint64 = 66
                v25 += uint16(v1)
                if true {
                    for fi35 := 0; fi35 < 12; fi35++ {
                        switch 0 {
                        case 32768:
                            if false {
                                var v36 uint16 = 950
                                _ = v36
                            } else {
                                if (("a" + ("hello" + "foo bar")) == "foo bar") {
                                    v5 += int(v1)
                                    switch fi35 {
                                    case 128:
                                        v25 = v25
                                    case 2147483647:
                                        var v37 bool = true
                                        _ = v37
                                    default:
                                        switch (fi26 << (((v2 / ((18446744073709551615) | 1))) & 63)) {
                                        case 2147483647:
                                            if len(v6) > 0 {
                                                v6[0] = int(v9)
                                            }
                                        case -170:
                                            switch int(v1) {
                                            case -619:
                                                v25 = v25
                                            default:
                                                if len(v6) > 0 {
                                                    v6[0] = v5
                                                }
                                            }
                                        }
                                    }
                                    _ = len(v4)
                                    v34 = (v24 % (((v24 + uint64(v25))) | 1))
                                } else {
                                    v38 := GenS1{F0: fi35, F1: v9, F2: 42}
                                    _ = v38
                                }
                                switch (fi35 / (((fi26 | fi35)) | 1)) {
                                case 52:
                                    switch fi26 {
                                    case -902:
                                        v39 := GenS2{F0: v34}
                                        _ = v39
                                    default:
                                        var v40 string = "hello"
                                        _ = v40
                                    }
                                case 720:
                                    _ = len(v3)
                                case -860:
                                    v1 = 42
                                default:
                                    fmt.Println(v24)
                                }
                                v34 -= (v24 ^ 18446744073709551615)
                            }
                            var v41 float32 = (-2.25 - float32(0.0))
                            _ = v41
                        case 128:
                            v9 -= v9
                        case -9223372036854775808:
                            fmt.Println(v1, v1)
                        default:
                            if false {
                                v25 = 558
                            } else {
                                v25 -= v25
                            }
                        }
                        _ = fi35
                    }
                    if (true || false) {
                        var v42 int8 = 127
                        for fi43 := 0; fi43 < 12; fi43++ {
                            var v44 uint32 = 65535
                            switch "" {
                            case "foo bar":
                                var v45 int32 = 0
                                _ = v45
                            case "hello":
                                v25 += (v25 & (^(v25)))
                            case "a":
                                v4[v5] = ((false && true) && (!false))
                            default:
                                v34 -= (^(v24))
                            }
                            v44 = (^(v44))
                            _ = fi43
                            _ = v44
                        }
                        _ = v42
                    }
                } else {
                    v46 := v3[1]
                    _ = v46
                }
                _ = v34
            }
        }
        for fi47 := 0; fi47 < 8; fi47++ {
            var v48 uint32 = 282
            _ = fi47
            _ = v48
        }
        _ = fi26
    }
    if (v24 == 573) {
        switch v5 {
        case 0:
            var v49 int64 = -1
            _ = v49
        }
        v2 = uint(v1)
        v1 += v1
    }
    var v50 uint32 = 146
    for fi51 := 0; fi51 < 10; fi51++ {
        for fi52 := 0; fi52 < 4; fi52++ {
            var v53 uint8 = 2
            _ = fi52
            _ = v53
        }
        for fi54 := 0; fi54 < 5; fi54++ {
            fmt.Println(fi54)
            _ = fi54
        }
        _ = fi51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
    _ = v9
}
