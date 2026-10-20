package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 bool }
type GenS2 struct { F0 uint8; F1 int; F2 int }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint8; F1 int64; F2 string }
type GenS5 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint16, p1 bool, p2 uint) uint32 {
    fmt.Println(p0, p2)
    if false {
        fmt.Println(p2)
        v5 := []int{}
        var v6 uint = (p2 / ((p2) | 1))
        _ = v5
        _ = v6
    }
    fmt.Println(p0, p0)
    var v7 uint16 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 4294967295
}
func genFn1(p0 uint32, p1 uint) int64 {
    var v8 uint16 = 65535
    fmt.Println(p0)
    if true {
        for fi9 := 0; fi9 < 2; fi9++ {
            if true {
                for fi10 := 0; fi10 < 14; fi10++ {
                    p1 = p1
                    switch 201 {
                    case -1:
                        if true {
                            var v11 int8 = 127
                            p1 = p1
                            _ = v11
                        } else {
                            v12 := GenS1{F0: p0, F1: 37, F2: true}
                            _ = v12
                        }
                        switch fi9 {
                        case -434:
                            var v13 uint8 = 0
                            _ = v13
                        default:
                            p0 = 1
                        }
                    case -2147483648:
                        switch "foo bar" {
                        case "hello":
                            var v14 int32 = 693
                            _ = v14
                        case "":
                            if true {
                                v15 := GenS1{F0: p0, F1: -1, F2: false}
                                _ = v15
                            }
                        case "foo bar":
                            switch "\n\t" {
                            case "":
                                var v16 float32 = 1e10
                                _ = v16
                            }
                        }
                    default:
                        var v17 uint = p1
                        _ = v17
                    }
                    v18 := GenS1{F0: (p0 >> (((p1 >> ((255) & 63))) & 63)), F1: 1, F2: (!(9223372036854775807 != 1))}
                    _ = fi10
                    _ = v18
                }
                p1 = 18446744073709551615
                p0 = p0
            } else {
                for fi19 := 0; fi19 < 17; fi19++ {
                    v20 := []int{(fi19 / ((fi9) | 1)), 32768, fi9}
                    _ = fi19
                    _ = v20
                }
            }
            v21 := []int{(fi9 & fi9), (-(fi9))}
            for fi22 := 0; fi22 < 19; fi22++ {
                if len(v21) > 0 {
                    v21[0] = (-(fi22))
                }
                fmt.Println(v8)
                _ = fi22
            }
            _ = fi9
            _ = v21
        }
        for fi23 := 0; fi23 < 4; fi23++ {
            v24 := GenS2{F0: 79, F1: fi23, F2: (fi23 - int(p1))}
            _ = fi23
            _ = v24
        }
    } else {
        var v25 uint64 = 251
        p1 += (^(p1))
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = -1
    var v3 uint8 = 42
    v4 := make(map[int]string)
    if true {
        if true {
            for fi26 := 0; fi26 < 17; fi26++ {
                if (!(!(v2 < v2))) {
                    for fi27 := 0; fi27 < 16; fi27++ {
                        var v28 int8 = 42
                        v29 := GenS3{F0: v2, F1: false}
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                } else {
                    fmt.Println(v2)
                    v30 := GenS1{F0: 65535, F1: -128, F2: false}
                    _ = v30
                }
                v3 += (v3 + v3)
                if (2 == v3) {
                    var v31 uint32 = 1
                    v32 := GenS2{F0: 42, F1: (fi26 + fi26), F2: fi26}
                    _ = v31
                    _ = v32
                }
                _ = fi26
            }
            v3 -= (v3 ^ (v3 >> ((505) & 63)))
            v2 -= 1000
        }
    }
    if (-9223372036854775808 != 606) {
        var v33 bool = true
        _ = v33
    }
    v34 := GenS4{F0: (v3 % ((v3) | 1)), F1: 0, F2: "foo bar"}
    v35 := GenS5{F0: 65535, F1: v3}
    v36 := GenS1{F0: 2147483647, F1: -1, F2: (false && true)}
    v3 = 87
    v37 := []int{2147483647, -9223372036854775808, 2147483647}
    v38 := GenS5{F0: genFn0(1, false, 1), F1: v3}
    for fi39 := 0; fi39 < 19; fi39++ {
        v2 = (v2 + 0)
        switch (^(fi39)) {
        case 303:
            if len(v37) > 0 {
                v37[0] = (fi39 << ((255) & 63))
            }
        case 127:
            v3 -= (^(v3))
        case 0:
            var v40 int64 = -617
            _ = v40
        default:
            if (!(-638 <= -9223372036854775808)) {
                for fi41 := 0; fi41 < 7; fi41++ {
                    v3 += v3
                    _ = fi41
                }
            }
        }
        _ = fi39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
