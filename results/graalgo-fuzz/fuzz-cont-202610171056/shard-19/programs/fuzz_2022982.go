package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 int64; F2 uint8 }
func genFn0(p0 int16) int64 {
    var v4 int64 = 0
    if false {
        switch "a" {
        case "\n\t":
            v5 := []int{-9223372036854775808, 0, 385, 32768}
            _ = v5
        case "":
            var v6 uint64 = 0
            _ = v6
        default:
            for fi7 := 0; fi7 < 5; fi7++ {
                switch fi7 {
                case 526:
                    for fi8 := 0; fi8 < 12; fi8++ {
                        v4 -= 1
                        _ = fi8
                    }
                case -682:
                    switch "" {
                    case "":
                        p0 -= (p0 * p0)
                    case "\n\t":
                        v9 := []int{}
                        _ = v9
                    case "foo bar":
                        if (!(false && true)) {
                            v10 := GenS1{F0: 509}
                            for fi11 := 0; fi11 < 7; fi11++ {
                                switch "foo bar" {
                                case "foo bar":
                                    v12 := GenS2{F0: (fi7 % ((fi11) | 1))}
                                    _ = v12
                                }
                                v4 -= v4
                                var v13 float32 = -2.25
                                _ = fi11
                                _ = v13
                            }
                            _ = v10
                        } else {
                            if true {
                                v14 := GenS3{F0: "foo bar"}
                                _ = v14
                            } else {
                                for fi15 := 0; fi15 < 1; fi15++ {
                                    if (!true) {
                                        var v16 uint = 65535
                                        v17 := []int{}
                                        _ = v16
                                        _ = v17
                                    }
                                    _ = fi15
                                }
                                var v18 uint32 = 65535
                                _ = v18
                            }
                            for fi19 := 0; fi19 < 19; fi19++ {
                                p0 += 0
                                v20 := []int{(^(fi19)), (fi7 - 2147483647), (-(fi7)), (fi19 >> ((2) & 63))}
                                _ = fi19
                                _ = v20
                            }
                            v21 := GenS1{F0: 693}
                            _ = v21
                        }
                    default:
                        v22 := GenS2{F0: fi7}
                        _ = v22
                    }
                }
                _ = fi7
            }
        }
        if true {
            v4 += (-(v4))
        } else {
            for fi23 := 0; fi23 < 18; fi23++ {
                v24 := GenS2{F0: fi23}
                var v25 int32 = 1
                v26 := GenS4{F0: 628, F1: (v4 - v4), F2: 255}
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
        }
    } else {
        v27 := GenS3{F0: ""}
        p0 += p0
        _ = v27
    }
    var v28 int = 2147483647
    for fi29 := 0; fi29 < 11; fi29++ {
        fmt.Println(fi29, p0)
        v4 = (v4 % ((1) | 1))
        _ = fi29
    }
    if ((1.5 >= ((-2.25 * 1.5) - 3.14159)) || ((false || (!(true || ((2 == 4294967295) && (float64((float32(0.0) / 1.0)) >= 3.14159))))) && true)) {
        for fi30 := 0; fi30 < 3; fi30++ {
            switch "hello" {
            case "a":
                v31 := []int{}
                _ = v31
            }
            v32 := GenS1{F0: 1}
            var v33 int = (fi30 & 0)
            _ = fi30
            _ = v32
            _ = v33
        }
    } else {
        if false {
            fmt.Println(v4, p0, v4)
            fmt.Println(p0, v4)
            if (!true) {
                fmt.Println(p0, v4)
                fmt.Println(v4, v4)
            }
        }
    }
    _ = p0
    _ = v28
    _ = v4
    return v4
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 42
    var v3 float32 = (1.5 - (0.0 - 3.14159))
    v34 := []int{v2, (-(v2)), (v2 & (v2 >> ((1) & 63)))}
    if len(v34) >= 2 {
        v34 = v34[1:len(v34)]
    }
    var v35 int8 = 100
    var v36 uint = 334
    var v37 uint64 = 2
    v36 -= (^(v36))
    v35 += -1
    v38 := GenS2{F0: (v2 << ((2) & 63))}
    for fi39 := 0; fi39 < 8; fi39++ {
        fmt.Println(v3, v36)
        delete(v1, "a")
        if true {
            var v40 int = fi39
            if (!(v3 != (v3 - float32((((3.14159 / 1.0) + (0.0 + float64(v3))) / 1.5))))) {
                v41 := GenS2{F0: v2}
                fmt.Println(v3, v40, v36)
                _ = v41
            } else {
                _ = len(v1)
            }
            delete(v1, "foo bar")
            _ = v40
        }
        _ = fi39
    }
    v35 = v35
    var v42 string = "\n\t"
    var v43 int64 = 1
    fmt.Println(v2, v3, v3)
    v42 = (v42 + ("\n\t" + "\n\t"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v42
    _ = v43
}
