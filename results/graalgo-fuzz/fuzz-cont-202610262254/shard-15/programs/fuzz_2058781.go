package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int8; F2 int8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint8, p1 uint8) int8 {
    p0 = (p0 | p0)
    var v4 int64 = 0
    var v5 uint8 = p1
    switch 9223372036854775807 {
    case 4294967295:
        v4 = (-(v4))
    default:
        var v6 int16 = 32767
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 uint, p1 int32) uint32 {
    fmt.Println(p0, p0)
    p1 += (p1 << (((p0 | (p0 % ((p0) | 1)))) & 63))
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: 138}
            v10 := []int{(fi8 ^ (^(fi7))), fi7, 0, (fi8 % (((fi7 ^ fi7)) | 1))}
            _ = fi8
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 8; fi11++ {
            switch 288 {
            case 69:
                v12 := []int{(fi11 % (((-(fi11))) | 1)), 318}
                _ = v12
            case 0:
                var v13 int8 = 100
                _ = v13
            }
            var v14 string = (("foo bar" + "hello") + "a")
            switch v14 {
            case "\n\t":
                switch 32767 {
                case 2147483647:
                    var v15 int16 = int16(p1)
                    _ = v15
                default:
                    v16 := GenS2{F0: (^(p0)), F1: -1, F2: 42}
                    _ = v16
                }
            default:
                v17 := GenS2{F0: (p0 << (((p0 * 4294967295)) & 63)), F1: 59, F2: -128}
                _ = v17
            }
            _ = fi11
            _ = v14
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    return 373
}

func main() {
    var v1 uint = 19
    var v2 uint16 = 65535
    var v3 uint32 = 2147483647
    for fi18 := 0; fi18 < 15; fi18++ {
        switch -2147483648 {
        case -9223372036854775808:
            for fi19 := 0; fi19 < 19; fi19++ {
                v20 := []int{int(v2), fi18, (fi18 * fi18)}
                var v21 int8 = int8(fi18)
                switch ("" + (("hello" + "foo bar") + ("" + ("hello" + ("a" + ((("\n\t" + "a") + "") + "hello")))))) {
                case "foo bar":
                    v2 += v2
                case "hello":
                    var v22 uint16 = v2
                    _ = v22
                case "\n\t":
                    var v23 float32 = 1.5
                    _ = v23
                default:
                    v21 = genFn0(32, 188)
                }
                fmt.Println(fi19, fi18)
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        v24 := []int{fi18}
        _ = fi18
        _ = v24
    }
    v25 := []int{-2147483648}
    if true {
        v1 -= v1
        v26 := GenS3{F0: 0}
        _ = v26
    } else {
        v27 := GenS1{F0: 42}
        _ = len(v25)
        var v28 int = -38
        _ = v27
        _ = v28
    }
    if (true && true) {
        v3 = v3
    } else {
        for fi29 := 0; fi29 < 3; fi29++ {
            fmt.Println(fi29, v1)
            if true {
                switch "" {
                case "a":
                    if true {
                        if true {
                            switch 2147483647 {
                            case -493:
                                if true {
                                    switch fi29 {
                                    case 0:
                                        if len(v25) >= 2 {
                                            v25 = v25[1:len(v25)]
                                        }
                                    case -525:
                                        _ = len(v25)
                                    }
                                } else {
                                    if len(v25) > 0 {
                                        v25[0] = fi29
                                    }
                                    v3 -= (v3 - (v3 % (((v3 >> (((v1 << (((v1 % ((v1) | 1))) & 63))) & 63))) | 1)))
                                }
                            }
                            var v30 int32 = -2147483648
                            var v31 uint8 = 255
                            _ = v30
                            _ = v31
                        } else {
                            var v32 int16 = -457
                            _ = v32
                        }
                    }
                case "":
                    if false {
                        var v33 float64 = 1e10
                        v34 := GenS1{F0: 0}
                        var v35 float32 = 3.14159
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                }
                v36 := GenS1{F0: uint8(v3)}
                _ = len(v25)
                _ = v36
            } else {
                for fi37 := 0; fi37 < 12; fi37++ {
                    var v38 uint = v1
                    v38 = (v38 / ((v38) | 1))
                    _ = fi37
                    _ = v38
                }
            }
            _ = fi29
        }
        fmt.Println(v3, v2)
    }
    var v39 string = ("\n\t" + ("a" + "\n\t"))
    var v40 int = 128
    var v41 uint16 = v2
    _ = len(v25)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v39
    _ = v40
    _ = v41
}
