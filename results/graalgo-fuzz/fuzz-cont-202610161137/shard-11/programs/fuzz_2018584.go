package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int64; F2 int16 }
type GenS4 struct { F0 int; F1 uint; F2 int }
func genFn0(p0 uint32) string {
    if false {
        v7 := GenS1{F0: 1, F1: 42}
        var v8 bool = (-63 < 45)
        var v9 int64 = 881
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := []int{-476}
    var v11 int16 = 1000
    _ = p0
    _ = v10
    _ = v11
    return (((("\n\t" + "hello") + "\n\t") + "hello") + "\n\t")
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float32 = 0.0
    v3 := []string{(("" + (("\n\t" + "\n\t") + ("a" + "foo bar"))) + "foo bar"), "hello", (("hello" + ("" + "foo bar")) + "hello"), "foo bar"}
    v4 := make(map[int]int)
    var v5 bool = (false && (!(!(1.5 != 1e10))))
    var v6 int16 = -1
    if v5 {
        var v12 int64 = int64(v6)
        v12 = 9223372036854775807
        v13 := GenS2{F0: -1}
        _ = v12
        _ = v13
    } else {
        fmt.Println(v1, v5, v5)
        v2 = (0.1 * (((float32(3.14159) - v2) * ((v2 - 0.1) * 1e10)) / v2))
    }
    v14 := GenS2{F0: -1}
    var v15 float32 = ((v2 * (float32(1.5) / (v2 * 0.0))) + v2)
    v2 = -2.25
    var v16 int64 = 0
    for fi17 := 0; fi17 < 1; fi17++ {
        switch "hello" {
        case "hello":
            v4[-1] = (fi17 + (fi17 & 4294967295))
        case "":
            switch "hello" {
            case "hello":
                for fi18 := 0; fi18 < 12; fi18++ {
                    var v19 uint8 = 255
                    v2 = v2
                    _ = fi18
                    _ = v19
                }
            default:
                v20 := GenS1{F0: int64(fi17), F1: 1}
                _ = v20
            }
        default:
            for fi21 := 0; fi21 < 8; fi21++ {
                v22 := GenS3{F0: v16, F1: 0, F2: (v6 | v6)}
                _ = fi21
                _ = v22
            }
        }
        v16 = v16
        _ = fi17
    }
    var v23 int = 4294967295
    var v24 int64 = v16
    if true {
        if true {
            var v25 uint64 = v1
            v26 := GenS2{F0: 1}
            if v5 {
                var v27 int64 = 813
                var v28 uint16 = 255
                for fi29 := 0; fi29 < 1; fi29++ {
                    switch "foo bar" {
                    case "\n\t":
                        v4[fi29] = (-(v23))
                    default:
                        switch v23 {
                        case 2147483647:
                            if (!v5) {
                                if false {
                                    fmt.Println(v28, v23)
                                    v30 := GenS1{F0: -910, F1: 139}
                                    var v31 float64 = -2.25
                                    _ = v30
                                    _ = v31
                                }
                                _ = len(v4)
                            } else {
                                var v32 float32 = 1.5
                                if (0 != 127) {
                                    v33 := GenS3{F0: v24, F1: v27, F2: (v6 - v6)}
                                    _ = v33
                                }
                                _ = v32
                            }
                        case -2147483648:
                            fmt.Println(v6, v25)
                        case 128:
                            for fi34 := 0; fi34 < 12; fi34++ {
                                var v35 float32 = 3.14159
                                v24 -= (v27 % ((v27) | 1))
                                _ = len(v4)
                                switch "\n\t" {
                                case "a":
                                    v28 += (v28 + (^(v28)))
                                case "foo bar":
                                    switch v23 {
                                    case 32768:
                                        var v36 int = v23
                                        _ = v36
                                    default:
                                        var v37 int64 = (v27 - v16)
                                        _ = v37
                                    }
                                case "":
                                    var v38 int8 = -33
                                    _ = v38
                                }
                                v6 = (v6 << ((0) & 63))
                                _ = fi34
                                _ = v35
                            }
                        default:
                            var v39 int16 = 42
                            _ = v39
                        }
                    }
                    delete(v4, v23)
                    v15 -= -2.25
                    _ = fi29
                }
                _ = v27
                _ = v28
            }
            _ = v25
            _ = v26
        } else {
            var v40 bool = false
            v41 := GenS4{F0: 821, F1: 717, F2: v23}
            _ = v40
            _ = v41
        }
    } else {
        for fi42 := 0; fi42 < 9; fi42++ {
            if true {
                for fi43 := 0; fi43 < 20; fi43++ {
                    var v44 int8 = 92
                    v45 := GenS2{F0: (v44 % (((v44 / (((^(v44))) | 1))) | 1))}
                    v46 := GenS4{F0: (v23 | v23), F1: 74, F2: fi42}
                    _ = fi43
                    _ = v44
                    _ = v45
                    _ = v46
                }
                v47 := GenS3{F0: v16, F1: (^(v24)), F2: (v6 & v6)}
                _ = v47
            } else {
                var v48 uint8 = 22
                v15 -= v2
                _ = v48
            }
            _ = fi42
        }
    }
    var v49 bool = (((-128 > 100) && (860 <= 2147483648)) && v5)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
