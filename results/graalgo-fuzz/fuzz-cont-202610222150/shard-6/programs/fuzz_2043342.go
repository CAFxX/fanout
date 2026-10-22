package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 uint64; F1 int64 }
func genFn0(p0 int64, p1 bool) uint {
    switch 32768 {
    case 4294967295:
        for fi4 := 0; fi4 < 5; fi4++ {
            for fi5 := 0; fi5 < 15; fi5++ {
                if ((1 < 255) && p1) {
                    p0 = p0
                    p0 = 772
                    fmt.Println(p1)
                    p1 = (p1 && p1)
                    v6 := GenS1{F0: -51}
                    _ = v6
                } else {
                    v7 := []int{388, fi5, (fi4 * 32767), fi5}
                    _ = v7
                }
                v8 := GenS1{F0: 127}
                v9 := GenS1{F0: 127}
                _ = fi5
                _ = v8
                _ = v9
            }
            fmt.Println(p0, p0)
            _ = fi4
        }
    }
    switch 42 {
    case 128:
        v10 := []int{127}
        _ = v10
    case 42:
        v11 := []int{-2147483648, 597, 795}
        _ = v11
    case 859:
        switch -352 {
        case 4294967295:
            p0 -= (p0 % ((p0) | 1))
        case 128:
            if (4 <= -75) {
                if p1 {
                    p0 = (-(p0))
                }
            }
        }
    }
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 int16, p1 string, p2 string) uint8 {
    if (!false) {
        var v12 int = -185
        v13 := []int{v12, (v12 << ((254) & 63)), (v12 & v12), (v12 / ((v12) | 1))}
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 0
}

func main() {
    var v1 int16 = 1
    var v2 uint8 = 0
    var v3 string = ""
    var v15 int32 = -1
    fmt.Println(v3, v3)
    var v16 uint16 = 42
    v17 := []int{}
    var v18 uint8 = 244
    if true {
        var v19 float64 = 0.0
        if false {
            switch v3 {
            case "a":
                var v20 float64 = (float64((1.5 + float32(1.5))) / (-2.25 * (((float64(0.1) - v19) + v19) / v19)))
                _ = v20
            case "":
                switch -811 {
                case 825:
                    fmt.Println(v1)
                default:
                    switch v3 {
                    case "foo bar":
                        var v21 float64 = float64(1.5)
                        _ = v21
                    case "":
                        var v22 bool = (!(true || true))
                        _ = v22
                    }
                }
            case "hello":
                v2 = (v18 * v18)
            default:
                if false {
                    for fi23 := 0; fi23 < 16; fi23++ {
                        v16 += v16
                        _ = fi23
                    }
                    v3 = "a"
                    if ((v15 * -2147483648) < (^(v15))) {
                        for fi24 := 0; fi24 < 3; fi24++ {
                            v25 := GenS1{F0: 1}
                            _ = fi24
                            _ = v25
                        }
                    } else {
                        fmt.Println(v16, v2, v3)
                        switch v3 {
                        case "hello":
                            for fi26 := 0; fi26 < 18; fi26++ {
                                for fi27 := 0; fi27 < 3; fi27++ {
                                    v16 = uint16(v2)
                                    if (v16 > 0) {
                                        v28 := GenS3{F0: false}
                                        v29 := GenS2{F0: v16}
                                        v30 := GenS3{F0: true}
                                        _ = v28
                                        _ = v29
                                        _ = v30
                                    } else {
                                        var v31 uint64 = 1
                                        v18 = v2
                                        var v32 float64 = v19
                                        _ = v31
                                        _ = v32
                                    }
                                    _ = fi27
                                }
                                v17 = append(v17, (fi26 << ((18446744073709551615) & 63)))
                                v33 := GenS4{F0: ("" + (v3 + (((v3 + v3) + v3) + (v3 + "hello")))), F1: ("" + "a")}
                                _ = fi26
                                _ = v33
                            }
                        case "\n\t":
                            if true {
                                fmt.Println(v2, v15, v19)
                                v34 := GenS5{F0: 94, F1: -1}
                                v15 += (-(v15))
                                _ = v34
                            } else {
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                var v35 int32 = v15
                                _ = v35
                            }
                        }
                        fmt.Println(v2)
                    }
                } else {
                    if len(v17) > 0 {
                        v17[0] = 559
                    }
                }
            }
        }
        v16 -= v16
        _ = v19
    } else {
        var v36 float32 = (1e10 + float32(((1.5 - 0.1) * float64(-2.25))))
        v18 = v18
        var v37 float64 = 1.5
        _ = v36
        _ = v37
    }
    var v38 int32 = v15
    if len(v17) > 0 {
        v17[0] = 4294967295
    }
    v3 = "foo bar"
    v39 := GenS1{F0: 95}
    v40 := GenS3{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v40
}
