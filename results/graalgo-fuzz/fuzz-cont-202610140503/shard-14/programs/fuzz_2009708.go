package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 uint64; F1 int; F2 uint8 }
type GenS6 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS7 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int, p1 uint16, p2 int64) int8 {
    v4 := GenS1{F0: 1, F1: -581}
    if (((-2.25 * ((1e10 - 0.1) * (-2.25 * float32((-2.25 * ((0.0 + -2.25) * 3.14159)))))) / 1.0) >= (1.5 * 1e10)) {
        v5 := []int{p0}
        var v6 uint32 = 0
        _ = v5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -1
}

func main() {
    v1 := []string{"", "hello", (("a" + ("foo bar" + "foo bar")) + "\n\t")}
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    if (255 >= 0) {
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 uint32 = 765
            v8 += v8
            if false {
                v9 := GenS1{F0: v8, F1: 2147483647}
                fmt.Println(v8)
                _ = v9
            }
            _ = fi7
            _ = v8
        }
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 uint8 = 128
            _ = fi10
            _ = v11
        }
        var v12 uint16 = 65535
        _ = v12
    } else {
        var v13 string = "foo bar"
        _ = v13
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v14 uint32 = 120
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 bool = (1 > -83)
        _ = len(v1)
        _ = fi15
        _ = v16
    }
    var v17 int64 = 0
    v14 -= v14
    for fi18 := 0; fi18 < 18; fi18++ {
        fmt.Println(v17)
        _ = fi18
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        for fi20 := 0; fi20 < 8; fi20++ {
            var v21 int64 = v17
            v22 := GenS2{F0: -247}
            _ = fi20
            _ = v21
            _ = v22
        }
        switch fi19 {
        case 170:
            fmt.Println(v14, v14, fi19)
        case 8:
            _ = len(v3)
        default:
            if ((v14 | v14) == v14) {
                v17 = 901
            }
        }
        var v23 uint = 18446744073709551615
        _ = fi19
        _ = v23
    }
    v14 -= v14
    switch 42 {
    case 547:
        var v24 bool = (!true)
        _ = v24
    default:
        v14 += v14
    }
    switch 42 {
    case 2147483647:
        var v25 int8 = 127
        _ = v25
    case -9223372036854775808:
        v14 -= 992
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        if true {
            v27 := GenS3{F0: 100, F1: v17}
            v14 = (v14 % ((468) | 1))
            if (-128 == 1) {
                v17 += 9223372036854775807
                if ((1e10 / (0.1 - 1.5)) <= -2.25) {
                    var v28 bool = true
                    if false {
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                        fmt.Println(v17, v17, v17)
                        if false {
                            var v29 uint = 680
                            var v30 float64 = 3.14159
                            v31 := GenS4{F0: -1, F1: -9223372036854775808}
                            _ = v29
                            _ = v30
                            _ = v31
                        } else {
                            var v32 bool = (v28 && (v28 || false))
                            var v33 int8 = 100
                            v34 := GenS5{F0: 4294967295, F1: int(v33), F2: 1}
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                    }
                    _ = v28
                }
            }
            _ = v27
        } else {
            switch "a" {
            case "foo bar":
                if ((((("\n\t" + "hello") + "\n\t") + ("" + ("a" + "foo bar"))) + "hello") == "") {
                    v17 -= (v17 & (v17 << ((18446744073709551615) & 63)))
                    v35 := GenS2{F0: 0}
                    var v36 int16 = 692
                    _ = v35
                    _ = v36
                } else {
                    fmt.Println(fi26)
                    if true {
                        var v37 int16 = 1000
                        v38 := GenS5{F0: 2, F1: fi26, F2: 128}
                        _ = v37
                        _ = v38
                    } else {
                        v39 := GenS6{F0: 1, F1: 2, F2: 709}
                        switch fi26 {
                        case -9223372036854775808:
                            if true {
                                fmt.Println(v17, v17, fi26)
                                fmt.Println(v17)
                            } else {
                                var v40 int32 = 1
                                _ = v40
                            }
                        case 4294967295:
                            var v41 float32 = (1.5 + 3.14159)
                            _ = v41
                        }
                        v42 := GenS7{F0: 0, F1: (^(v14))}
                        _ = v39
                        _ = v42
                    }
                    var v43 uint16 = 32767
                    _ = v43
                }
            }
            v44 := GenS7{F0: 18446744073709551615, F1: (^(v14))}
            v17 = v17
            _ = v44
        }
        _ = fi26
    }
    delete(v2, -916)
    var v45 int8 = genFn0(4294967295, 1, (^(v17)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v45:", v45)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v45
}
