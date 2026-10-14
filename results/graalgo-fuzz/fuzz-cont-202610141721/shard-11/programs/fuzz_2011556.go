package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint }
func genFn0(p0 bool, p1 uint16) uint16 {
    v5 := GenS1{F0: -250}
    p1 -= 251
    v6 := []int{}
    p0 = false
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint16 = p1
        if p0 {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            p0 = (p1 != (^(p1)))
            var v9 int64 = 337
            _ = v9
        } else {
            v6 = append(v6, 128)
        }
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 int8 = -57
            v12 := GenS1{F0: 42}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (p1 | (p1 % ((2) | 1)))
}
func genFn1(p0 string) uint {
    var v13 uint8 = 1
    var v14 int64 = -627
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 int64 = v14
        v17 := []int{}
        _ = fi15
        _ = v16
        _ = v17
    }
    switch 32768 {
    case 760:
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := GenS1{F0: (v14 << ((2) & 63))}
            _ = fi18
            _ = v19
        }
    case -305:
        for fi20 := 0; fi20 < 7; fi20++ {
            v21 := GenS1{F0: v14}
            if true {
                if ((!(2 <= 616)) || (fi20 > 42)) {
                    for fi22 := 0; fi22 < 16; fi22++ {
                        var v23 int16 = int16(v13)
                        _ = fi22
                        _ = v23
                    }
                    v24 := GenS1{F0: (^(v14))}
                    _ = v24
                }
                v25 := GenS2{F0: fi20, F1: 1000}
                var v26 bool = true
                _ = v25
                _ = v26
            }
            var v27 int8 = -1
            _ = fi20
            _ = v21
            _ = v27
        }
    default:
        for fi28 := 0; fi28 < 14; fi28++ {
            switch "hello" {
            case "":
                v29 := []int{(fi28 * (fi28 - (fi28 ^ (fi28 << ((1) & 63)))))}
                _ = v29
            case "foo bar":
                v30 := GenS3{F0: 1}
                _ = v30
            case "\n\t":
                if false {
                    if (!false) {
                        if true {
                            for fi31 := 0; fi31 < 15; fi31++ {
                                v32 := GenS1{F0: (v14 % ((v14) | 1))}
                                fmt.Println(p0)
                                v33 := []int{}
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                        }
                    } else {
                        if (0 > 42) {
                            var v34 uint32 = 2147483647
                            v34 = 46
                            _ = v34
                        }
                    }
                }
            default:
                var v35 int = 821
                _ = v35
            }
            _ = fi28
        }
    }
    switch int(v13) {
    case 2147483647:
        for fi36 := 0; fi36 < 17; fi36++ {
            var v37 int32 = 1
            _ = fi36
            _ = v37
        }
    default:
        var v38 int32 = 2147483647
        _ = v38
    }
    _ = p0
    _ = v13
    _ = v14
    return 2
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{"", ("hello" + "a")}
    v3 := make(map[uint32]string)
    var v4 int8 = -128
    var v39 float32 = -2.25
    var v40 int = -993
    for fi41 := 0; fi41 < 4; fi41++ {
        v39 += v39
        _ = fi41
    }
    v4 += v4
    v39 = 3.14159
    v42 := v3[42]
    fmt.Println(v39, v40, v42)
    var v43 int8 = v4
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v2 = append(v2, v42)
    for fi44 := 0; fi44 < 15; fi44++ {
        if true {
            v1[fi44] = (-(v40))
            var v45 int32 = 2147483647
            _ = v45
        }
        v46 := GenS1{F0: 1}
        _ = fi44
        _ = v46
    }
    var v47 string = v42
    for fi48 := 0; fi48 < 17; fi48++ {
        for fi49 := 0; fi49 < 10; fi49++ {
            var v50 int8 = (^(v43))
            delete(v1, fi49)
            delete(v3, 2147483647)
            _ = fi49
            _ = v50
        }
        _ = fi48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v43
    _ = v47
}
