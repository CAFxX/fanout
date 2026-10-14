package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint8; F2 uint16 }

func main() {
    var v1 string = ("" + "hello")
    var v2 int64 = 1
    var v3 int64 = (v2 ^ v2)
    v4 := make(map[int]string)
    v5 := make(map[uint32]string)
    var v6 uint32 = 2147483647
    v3 += (v3 & v3)
    v7 := []int{-1}
    if (-32768 >= 42) {
        if false {
            for fi8 := 0; fi8 < 19; fi8++ {
                fmt.Println(fi8)
                _ = fi8
            }
            var v9 uint8 = 2
            _ = v9
        } else {
            switch 4294967295 {
            case 660:
                for fi10 := 0; fi10 < 11; fi10++ {
                    v6 = (v6 - v6)
                    v7 = append(v7, fi10)
                    _ = fi10
                }
            case 32768:
                var v11 bool = ((v6 / ((v6) | 1)) > (v6 & v6))
                _ = v11
            }
            v1 = ((v1 + ((("foo bar" + v1) + "") + v1)) + v1)
            if false {
                switch "foo bar" {
                case "\n\t":
                    if true {
                        var v12 uint8 = 42
                        v13 := GenS1{F0: -1}
                        _ = v12
                        _ = v13
                    } else {
                        var v14 uint64 = 0
                        _ = len(v7)
                        _ = v14
                    }
                case "hello":
                    var v15 bool = true
                    _ = v15
                case "foo bar":
                    var v16 int8 = 89
                    _ = v16
                }
            } else {
                v2 = 9223372036854775807
                if (-2.25 > (float64(1e10) * 0.0)) {
                    var v17 int16 = 32767
                    for fi18 := 0; fi18 < 11; fi18++ {
                        v2 -= (^(v2))
                        for fi19 := 0; fi19 < 12; fi19++ {
                            fmt.Println(fi18)
                            if false {
                                var v20 float64 = 3.14159
                                _ = v20
                            }
                            var v21 bool = (-2.25 <= 1.5)
                            _ = fi19
                            _ = v21
                        }
                        v3 += v3
                        _ = fi18
                    }
                    var v22 uint8 = 0
                    _ = v17
                    _ = v22
                } else {
                    for fi23 := 0; fi23 < 2; fi23++ {
                        switch fi23 {
                        case -2147483648:
                            v5[v6] = v1
                        case 385:
                            _ = len(v4)
                        case 42:
                            v24 := GenS2{F0: (v3 == -1)}
                            _ = v24
                        }
                        _ = len(v7)
                        _ = fi23
                    }
                    var v25 uint32 = (v6 % (((v6 * (v6 << ((0) & 63)))) | 1))
                    if false {
                        fmt.Println(v2, v6, v6)
                    } else {
                        switch 32768 {
                        case -195:
                            v5[v6] = ("foo bar" + v1)
                        case 0:
                            for fi26 := 0; fi26 < 1; fi26++ {
                                for fi27 := 0; fi27 < 17; fi27++ {
                                    v3 -= v3
                                    _ = fi27
                                }
                                if len(v7) > 0 {
                                    v7[0] = fi26
                                }
                                v7 = append(v7, 4294967295)
                                _ = fi26
                            }
                        }
                        v3 += (v2 - (v2 % (((^(v2))) | 1)))
                        fmt.Println(v25)
                    }
                    _ = v25
                }
            }
        }
        v28 := GenS1{F0: -2147483648}
        _ = v28
    }
    v3 += v2
    var v29 string = "a"
    switch -1 {
    case -1:
        for fi30 := 0; fi30 < 3; fi30++ {
            if false {
                fmt.Println(v2)
            }
            if false {
                v29 = "foo bar"
            }
            _ = fi30
        }
    case 882:
        switch ("foo bar" + (v1 + v29)) {
        case "":
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        default:
            var v31 float64 = 0.0
            _ = v31
        }
    case 128:
        v32 := GenS1{F0: -1}
        _ = v32
    default:
        for fi33 := 0; fi33 < 3; fi33++ {
            if true {
                v6 = v6
                v34 := GenS2{F0: true}
                var v35 int8 = -27
                _ = v34
                _ = v35
            }
            var v36 int32 = 0
            fmt.Println(v6, v36, fi33)
            _ = fi33
            _ = v36
        }
    }
    if ((255 < 0) && false) {
        delete(v4, 675)
    }
    for fi37 := 0; fi37 < 6; fi37++ {
        switch fi37 {
        case 4294967295:
            for fi38 := 0; fi38 < 5; fi38++ {
                v29 = (v1 + (v1 + v1))
                _ = fi38
            }
        case 117:
            v39 := GenS2{F0: (!(((fi37 == (fi37 + 32768)) && (true && (-625 <= fi37))) || true))}
            _ = v39
        }
        for fi40 := 0; fi40 < 9; fi40++ {
            if false {
                var v41 float64 = 1.5
                v42 := GenS1{F0: -2147483648}
                _ = v41
                _ = v42
            }
            _ = fi40
        }
        _ = fi37
    }
    v43 := GenS3{F0: 42, F1: 42, F2: 2}
    v44 := GenS2{F0: (18446744073709551615 <= 2)}
    var v45 int64 = v3
    v46 := GenS3{F0: -128, F1: 210, F2: 384}
    v7 = append(v7, 0)
    var v47 bool = true
    v29 = "hello"
    for fi48 := 0; fi48 < 2; fi48++ {
        v47 = (497 != 1)
        var v49 uint = 592
        _ = fi48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
