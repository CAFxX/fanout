package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 int16; F1 uint32 }
func genFn0(p0 int32, p1 bool, p2 uint) uint64 {
    p0 += (p0 * (-(p0)))
    var v7 bool = (1 > 42)
    var v8 int16 = 0
    v9 := GenS1{F0: (v8 - v8), F1: 32767}
    if ((v7 || true) || p1) {
        var v10 uint16 = 0
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 562
}

func main() {
    var v1 int16 = 1000
    var v2 int8 = -39
    v3 := make(map[int]int)
    var v4 int64 = 1
    var v5 int8 = 127
    var v6 int8 = (v2 + v2)
    var v11 int64 = (v4 << ((1) & 63))
    switch 4294967295 {
    case -9223372036854775808:
        for fi12 := 0; fi12 < 14; fi12++ {
            fmt.Println(v1)
            fmt.Println(v4, v5)
            _ = fi12
        }
    }
    v1 = (v1 / ((-554) | 1))
    fmt.Println(v2, v5)
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := []int{(fi13 ^ 127)}
        var v15 bool = true
        _ = fi13
        _ = v14
        _ = v15
    }
    if false {
        v16 := v3[2147483647]
        v1 -= v1
        _ = v16
    }
    switch 128 {
    case 128:
        for fi17 := 0; fi17 < 20; fi17++ {
            fmt.Println(fi17, v5)
            v18 := GenS2{F0: 42, F1: -112}
            var v19 int8 = 127
            _ = fi17
            _ = v18
            _ = v19
        }
    case 127:
        var v20 int64 = 42
        _ = v20
    default:
        if false {
            for fi21 := 0; fi21 < 16; fi21++ {
                var v22 int16 = (v1 << ((42) & 63))
                var v23 uint16 = 110
                fmt.Println(v1, v1)
                _ = fi21
                _ = v22
                _ = v23
            }
            var v24 int8 = -34
            for fi25 := 0; fi25 < 2; fi25++ {
                if (!false) {
                    _ = len(v3)
                    if (!true) {
                        switch (fi25 - fi25) {
                        case 9223372036854775807:
                            var v26 float32 = (1e10 + (float32(0.0) * -2.25))
                            _ = v26
                        case 32767:
                            v2 = v24
                        default:
                            v27 := v3[(fi25 & -9223372036854775808)]
                            _ = v27
                        }
                        v5 -= v5
                        for fi28 := 0; fi28 < 2; fi28++ {
                            v29 := GenS3{F0: ("a" + (("a" + "foo bar") + "a")), F1: 179}
                            _ = fi28
                            _ = v29
                        }
                    }
                }
                switch fi25 {
                case 127:
                    v30 := GenS1{F0: (v1 / ((v1) | 1)), F1: 255}
                    _ = v30
                case 479:
                    if true {
                        delete(v3, (^(fi25)))
                        var v31 bool = (825 != 2)
                        if v31 {
                            for fi32 := 0; fi32 < 16; fi32++ {
                                v11 += 659
                                v4 = (v11 + (v11 / (((v11 % ((v4) | 1))) | 1)))
                                _ = fi32
                            }
                            var v33 float64 = 1e10
                            _ = v33
                        } else {
                            v2 += (^(v6))
                            if v31 {
                                fmt.Println(fi25)
                                v24 += v24
                                for fi34 := 0; fi34 < 6; fi34++ {
                                    var v35 int8 = v5
                                    switch fi25 {
                                    case -1:
                                        var v36 int16 = v1
                                        _ = v36
                                    default:
                                        var v37 int32 = -2147483648
                                        _ = v37
                                    }
                                    v5 -= v5
                                    var v38 int32 = -2147483648
                                    var v39 float64 = 0.0
                                    _ = fi34
                                    _ = v35
                                    _ = v38
                                    _ = v39
                                }
                            }
                        }
                        _ = v31
                    } else {
                        var v40 bool = false
                        fmt.Println(fi25, v1)
                        _ = v40
                    }
                default:
                    v41 := []int{fi25, (fi25 * 4294967295), fi25}
                    _ = v41
                }
                var v42 string = "foo bar"
                _ = fi25
                _ = v42
            }
            _ = v24
        }
    }
    for fi43 := 0; fi43 < 16; fi43++ {
        v6 = (v5 % ((v5) | 1))
        var v44 uint64 = 2
        for fi45 := 0; fi45 < 11; fi45++ {
            var v46 int16 = v1
            _ = fi45
            _ = v46
        }
        _ = fi43
        _ = v44
    }
    v47 := GenS4{F0: -1, F1: 0}
    v48 := []int{180}
    switch "foo bar" {
    case "foo bar":
        v11 += 1
    case "":
        var v49 int64 = (v4 + (v11 | 42))
        _ = v49
    }
    v2 += (-(v2))
    var v50 int = 0
    var v51 int64 = v4
    for fi52 := 0; fi52 < 18; fi52++ {
        v3[(v50 % ((fi52) | 1))] = v50
        var v53 string = "a"
        fmt.Println(fi52, v11, v1)
        _ = fi52
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
