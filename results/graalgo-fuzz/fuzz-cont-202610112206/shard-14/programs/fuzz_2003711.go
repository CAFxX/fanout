package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int16 }

func main() {
    v1 := []int{-171, 32767}
    var v2 uint64 = 0
    var v3 int32 = 42
    v4 := make(map[uint32]bool)
    var v5 uint32 = 31
    var v6 uint64 = (^(v2))
    var v7 uint = 4294967295
    fmt.Println(v5, v5, v5)
    fmt.Println(v2)
    v6 = 1
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS1{F0: 42, F1: 32767, F2: uint16(v5)}
            var v11 float64 = 3.14159
            var v12 int64 = -1
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        delete(v4, (^(v5)))
        if false {
            if false {
                fmt.Println(v3, v7, v2)
            } else {
                v13 := GenS1{F0: 17, F1: 42, F2: uint16(v7)}
                v7 -= v7
                _ = v13
            }
        }
        _ = fi8
    }
    v2 = (v6 ^ v6)
    var v14 uint16 = uint16(v5)
    if false {
        v6 += (v2 * 28)
        switch "a" {
        case "a":
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 uint16 = (v14 * v14)
                var v17 int16 = -32768
                _ = fi15
                _ = v16
                _ = v17
            }
        }
    } else {
        var v18 int = 1
        v6 = 478
        switch 127 {
        case -1:
            if (v3 != (v3 + -1)) {
                for fi19 := 0; fi19 < 17; fi19++ {
                    if true {
                        fmt.Println(v5)
                        for fi20 := 0; fi20 < 20; fi20++ {
                            v21 := GenS1{F0: 42, F1: 0, F2: 65535}
                            _ = fi20
                            _ = v21
                        }
                        var v22 float64 = 0.1
                        _ = v22
                    }
                    switch ("\n\t" + "a") {
                    case "foo bar":
                        v7 = (v7 << (((v7 >> ((v7) & 63))) & 63))
                    case "a":
                        _ = len(v4)
                    }
                    if ((v3 - -1) > v3) {
                        v2 = v2
                        var v23 uint64 = 1
                        _ = v23
                    } else {
                        v3 += 105
                        v6 += (v6 | v6)
                        v24 := GenS2{F0: 65535, F1: (fi19 << ((v7) & 63))}
                        _ = v24
                    }
                    _ = fi19
                }
            } else {
                for fi25 := 0; fi25 < 5; fi25++ {
                    fmt.Println(v18, v3)
                    _ = fi25
                }
                v14 += v14
                for fi26 := 0; fi26 < 12; fi26++ {
                    if (!(true && false)) {
                        v27 := GenS2{F0: v7, F1: 42}
                        if (0.1 != -2.25) {
                            var v28 uint8 = uint8(v6)
                            _ = v28
                        } else {
                            delete(v4, (v5 * v5))
                            var v29 int = 2147483647
                            _ = v29
                        }
                        var v30 float32 = 3.14159
                        _ = v27
                        _ = v30
                    } else {
                        if true {
                            _ = len(v1)
                            if (!(true || ((394 >= -414) && false))) {
                                var v31 float64 = -2.25
                                for fi32 := 0; fi32 < 4; fi32++ {
                                    if ((^(v5)) == v5) {
                                        delete(v4, uint32(v7))
                                        switch fi32 {
                                        case 4294967295:
                                            if false {
                                                if false {
                                                    var v33 int = (v18 & fi32)
                                                    v34 := GenS1{F0: 42, F1: 255, F2: 2}
                                                    v7 = v7
                                                    switch (("foo bar" + ("a" + ("\n\t" + "\n\t"))) + "") {
                                                    case "":
                                                        v4[285] = ((v6 <= v6) || true)
                                                    default:
                                                        var v35 int64 = 0
                                                        _ = v35
                                                    }
                                                    v36 := GenS1{F0: 217, F1: (v14 & v14), F2: 2}
                                                    _ = v33
                                                    _ = v34
                                                    _ = v36
                                                } else {
                                                    v37 := GenS3{F0: -1}
                                                    _ = v37
                                                }
                                                var v38 uint8 = 42
                                                _ = v38
                                            } else {
                                                delete(v4, (v5 / ((65535) | 1)))
                                            }
                                        case 530:
                                            var v39 float64 = (1e10 / (v31 + v31))
                                            _ = v39
                                        case 2147483647:
                                            v3 = v3
                                        default:
                                            var v40 float64 = v31
                                            _ = v40
                                        }
                                        var v41 int64 = 42
                                        _ = v41
                                    } else {
                                        if false {
                                            var v42 string = "hello"
                                            if false {
                                                var v43 bool = (true && true)
                                                var v44 bool = (v31 == (((0.1 / (float64((1.5 / (0.0 + 1.5))) + float64(float32((1.5 - (v31 / (((v31 / -2.25) / v31) + v31))))))) * v31) + (v31 * (v31 + v31))))
                                                fmt.Println(v44, v42, v14)
                                                _ = v43
                                                _ = v44
                                            }
                                            var v45 int64 = -9223372036854775808
                                            _ = v42
                                            _ = v45
                                        }
                                        var v46 string = "hello"
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v46
                                    }
                                    _ = fi32
                                }
                                if false {
                                    var v47 uint = v7
                                    v47 -= v47
                                    v48 := v4[(^(v5))]
                                    _ = v47
                                    _ = v48
                                }
                                _ = v31
                            }
                        } else {
                            var v49 float32 = ((0.0 / 1.0) / 1.5)
                            _ = v49
                        }
                        var v50 uint16 = (^(v14))
                        _ = v50
                    }
                    _ = fi26
                }
            }
        }
        _ = v18
    }
    var v51 string = "hello"
    for fi52 := 0; fi52 < 5; fi52++ {
        fmt.Println(v3, v51, v3)
        v53 := GenS2{F0: 65535, F1: fi52}
        v2 = (^(v2))
        _ = fi52
        _ = v53
    }
    var v54 uint8 = 255
    var v55 float32 = 0.1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v54
    _ = v55
    _ = v6
    _ = v7
}
