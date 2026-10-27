package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS7 struct { F0 int; F1 uint8; F2 bool }
type GenS8 struct { F0 uint }
func genFn0(p0 int8, p1 int, p2 int64) int8 {
    v5 := []int{}
    v6 := GenS1{F0: (-(p0)), F1: 2, F2: 1}
    v7 := GenS1{F0: p0, F1: 4294967295, F2: -91}
    var v8 int8 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 100
}
func genFn1() int {
    if true {
        var v9 uint16 = 956
        v9 = v9
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := []int{fi10, (fi10 % ((fi10) | 1))}
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    var v12 int = -2147483648
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := GenS1{F0: 127, F1: 0, F2: int16(v12)}
        if (false || (false && false)) {
            v12 -= v12
            if (!true) {
                switch ("" + "hello") {
                case "":
                    if false {
                        v15 := GenS2{F0: 637}
                        if false {
                            var v16 uint16 = 667
                            v17 := GenS3{F0: true}
                            fmt.Println(v16, fi13)
                            _ = v16
                            _ = v17
                        } else {
                            v18 := GenS3{F0: false}
                            var v19 uint8 = 2
                            var v20 uint32 = 186
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        v12 -= 4294967295
                        _ = v15
                    }
                case "a":
                    if true {
                        v21 := GenS3{F0: (!((!false) && ((0 <= 1) && (0 >= 1))))}
                        v12 += fi13
                        _ = v21
                    } else {
                        fmt.Println(v12, v12)
                        var v22 int8 = 26
                        v23 := []int{(v12 + fi13), (-(fi13)), 643, fi13}
                        _ = v22
                        _ = v23
                    }
                case "\n\t":
                    for fi24 := 0; fi24 < 4; fi24++ {
                        v25 := GenS1{F0: 88, F1: 55, F2: 1000}
                        _ = fi24
                        _ = v25
                    }
                default:
                    var v26 int16 = -785
                    _ = v26
                }
                var v27 bool = false
                _ = v27
            } else {
                for fi28 := 0; fi28 < 12; fi28++ {
                    fmt.Println(fi13)
                    v29 := []int{}
                    _ = fi28
                    _ = v29
                }
            }
        } else {
            var v30 int64 = -60
            v31 := GenS4{F0: true, F1: 2}
            _ = v30
            _ = v31
        }
        for fi32 := 0; fi32 < 8; fi32++ {
            var v33 int64 = -1
            fmt.Println(v33, v33, fi32)
            _ = fi32
            _ = v33
        }
        _ = fi13
        _ = v14
    }
    for fi34 := 0; fi34 < 12; fi34++ {
        v35 := GenS5{F0: 1}
        _ = fi34
        _ = v35
    }
    v12 = (v12 >> ((829) & 63))
    _ = v12
    return (v12 - v12)
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 0
    var v3 int8 = -128
    var v4 float32 = float32(((-2.25 / 3.14159) - ((((3.14159 - 1e10) - (1e10 + (1e10 * 0.0))) - 1e10) - 1.5)))
    var v36 bool = true
    for fi37 := 0; fi37 < 11; fi37++ {
        v2 = (v2 << ((2) & 63))
        var v38 int32 = 0
        _ = fi37
        _ = v38
    }
    var v39 uint8 = 2
    v40 := GenS4{F0: true, F1: 150}
    fmt.Println(v2, v3)
    for fi41 := 0; fi41 < 16; fi41++ {
        fmt.Println(v36, v2, fi41)
        switch "foo bar" {
        case "hello":
            v4 = 0.1
        case "":
            v42 := GenS2{F0: 0}
            _ = v42
        case "foo bar":
            switch "" {
            case "\n\t":
                for fi43 := 0; fi43 < 7; fi43++ {
                    v36 = (255 < 65535)
                    if v36 {
                        fmt.Println(v4)
                        if (!(11 >= -1)) {
                            if (false || ((v3 & (v3 + v3)) < (v3 + genFn0((v3 & (v3 >> ((1) & 63))), genFn1(), 0)))) {
                                v4 += 1.5
                                var v44 uint64 = v2
                                if v36 {
                                    fmt.Println(fi41, v39)
                                } else {
                                    if (v2 >= (v2 | v2)) {
                                        v39 = 240
                                    } else {
                                        v39 = 127
                                        v3 += (v3 * v3)
                                    }
                                    var v45 int8 = (v3 & (v3 & 1))
                                    if ((fi41 << ((1) & 63)) == genFn1()) {
                                        v45 += 0
                                        var v46 float32 = (v4 - v4)
                                        _ = v46
                                    }
                                    _ = v45
                                }
                                _ = v44
                            } else {
                                if ("a" != ("hello" + "")) {
                                    v47 := GenS4{F0: (fi41 == (fi41 / (((fi43 << ((324) & 63))) | 1))), F1: v39}
                                    _ = v47
                                }
                            }
                        } else {
                            v48 := GenS2{F0: 4294967295}
                            _ = v48
                        }
                        v2 += v2
                    } else {
                        v49 := []int{}
                        var v50 uint32 = 65535
                        _ = v49
                        _ = v50
                    }
                    switch "a" {
                    case "hello":
                        if (v3 <= (v3 >> ((0) & 63))) {
                            var v51 uint16 = uint16(v2)
                            v52 := []int{}
                            v51 = 0
                            v4 -= (-2.25 - 0.0)
                            var v53 float32 = v4
                            if v36 {
                                v51 -= uint16(v39)
                                var v54 uint8 = 127
                                _ = len(v52)
                                v54 = v54
                                _ = v54
                            }
                            _ = v51
                            _ = v52
                            _ = v53
                        }
                    case "foo bar":
                        v55 := GenS6{F0: 1, F1: (!(-32768 == 32767)), F2: 128}
                        _ = v55
                    }
                    switch (fi43 ^ -672) {
                    case 473:
                        v56 := []int{fi41, 127, (fi43 * (fi41 ^ -1))}
                        _ = v56
                    }
                    _ = fi43
                }
            default:
                v3 -= 42
            }
        default:
            var v57 int = 127
            _ = v57
        }
        v3 -= (v3 >> ((228) & 63))
        _ = fi41
    }
    v58 := GenS7{F0: 369, F1: v39, F2: (2 > 4294967295)}
    if true {
        v59 := GenS8{F0: 255}
        v60 := []int{-150, 128, -718}
        _ = v59
        _ = v60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v58
}
