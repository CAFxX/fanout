package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int32; F1 uint8; F2 uint }
type GenS3 struct { F0 int16; F1 bool; F2 uint8 }
func genFn0(p0 bool, p1 uint, p2 int) bool {
    fmt.Println(p1, p1, p2)
    var v7 uint8 = 225
    if p0 {
        var v8 int16 = -95
        for fi9 := 0; fi9 < 13; fi9++ {
            if (160 != 2) {
                v10 := []int{p2, 32768}
                fmt.Println(fi9)
                _ = v10
            }
            for fi11 := 0; fi11 < 19; fi11++ {
                v7 = 255
                _ = fi11
            }
            _ = fi9
        }
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 int8 = 127
            _ = fi12
            _ = v13
        }
        _ = v8
    } else {
        fmt.Println(p1)
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}
func genFn1(p0 uint32, p1 int16, p2 string) uint64 {
    var v14 uint8 = 100
    var v15 uint8 = v14
    v16 := GenS1{F0: (p0 ^ (p0 - uint32(v14))), F1: 18446744073709551615}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    return 42
}

func main() {
    v1 := []string{((("hello" + "") + ("foo bar" + "foo bar")) + "foo bar"), "hello"}
    v2 := make(map[int]bool)
    var v3 int = 2147483647
    var v4 int32 = -1
    v5 := make(map[uint32]string)
    var v6 uint32 = 946
    v3 = (-(v3))
    var v17 int64 = 42
    v18 := GenS1{F0: v6, F1: 0}
    var v19 int16 = 980
    if (!((v17 * v17) > v17)) {
        for fi20 := 0; fi20 < 6; fi20++ {
            v3 = 2147483647
            v4 -= v4
            _ = len(v1)
            _ = fi20
        }
        if false {
            if true {
                for fi21 := 0; fi21 < 6; fi21++ {
                    var v22 int32 = v4
                    _ = fi21
                    _ = v22
                }
                v1 = append(v1, "a")
                for fi23 := 0; fi23 < 19; fi23++ {
                    v24 := v2[fi23]
                    v25 := GenS1{F0: 1, F1: 0}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            } else {
                for fi26 := 0; fi26 < 20; fi26++ {
                    v27 := GenS1{F0: v6, F1: 255}
                    var v28 int8 = 0
                    _ = len(v5)
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
            for fi29 := 0; fi29 < 5; fi29++ {
                var v30 uint64 = 4294967295
                _ = fi29
                _ = v30
            }
        }
        var v31 float64 = float64(0.0)
        _ = v31
    } else {
        switch ("\n\t" + ("\n\t" + "\n\t")) {
        case "hello":
            if (true && false) {
                var v32 string = ("foo bar" + ((("\n\t" + "a") + "") + ""))
                if (!false) {
                    var v33 uint64 = 18446744073709551615
                    _ = v33
                } else {
                    switch v32 {
                    case "a":
                        var v34 int64 = v17
                        _ = v34
                    case "\n\t":
                        var v35 bool = ((-(v17)) >= 42)
                        _ = v35
                    default:
                        switch v3 {
                        case 32768:
                            v6 = 1
                        case 96:
                            if (-3 < 127) {
                                v36 := GenS1{F0: v6, F1: 255}
                                _ = v36
                            }
                        }
                    }
                    v17 += (v17 << ((65535) & 63))
                }
                v3 -= (v3 + v3)
                _ = v32
            }
        case "foo bar":
            if (1e10 < 0.0) {
                v6 = (v6 ^ v6)
                var v37 uint = 65535
                _ = v37
            } else {
                if (84 >= 127) {
                    fmt.Println(v17)
                    if ((1 >= 127) && (!((false && false) && ((float64(1e10) <= 3.14159) && false)))) {
                        for fi38 := 0; fi38 < 19; fi38++ {
                            if len(v1) > 0 {
                                v1[0] = "a"
                            }
                            _ = fi38
                        }
                        fmt.Println(v6, v3, v4)
                        fmt.Println(v4, v3)
                    } else {
                        v6 += v6
                        v4 = v4
                    }
                    for fi39 := 0; fi39 < 17; fi39++ {
                        v3 += (fi39 * v3)
                        _ = fi39
                    }
                }
                switch "\n\t" {
                case "hello":
                    var v40 int8 = -113
                    _ = v40
                default:
                    v3 += v3
                }
                fmt.Println(v3, v3, v17)
            }
        case "a":
            for fi41 := 0; fi41 < 5; fi41++ {
                for fi42 := 0; fi42 < 2; fi42++ {
                    v19 -= -951
                    _ = fi42
                }
                _ = fi41
            }
        }
    }
    var v43 int = (v3 >> ((42) & 63))
    _ = len(v2)
    v44 := GenS1{F0: v6, F1: 255}
    v45 := GenS1{F0: 0, F1: 42}
    if false {
        v19 -= v19
        for fi46 := 0; fi46 < 6; fi46++ {
            fmt.Println(v19)
            for fi47 := 0; fi47 < 5; fi47++ {
                _ = len(v1)
                _ = fi47
            }
            switch "a" {
            case "":
                for fi48 := 0; fi48 < 11; fi48++ {
                    switch ("a" + "") {
                    case "a":
                        v49 := GenS2{F0: (v4 * v4), F1: 255, F2: 482}
                        _ = v49
                    case "foo bar":
                        var v50 int64 = (v17 / (((v17 >> ((4294967295) & 63))) | 1))
                        _ = v50
                    case "hello":
                        _ = len(v1)
                    default:
                        var v51 float32 = float32(1e10)
                        _ = v51
                    }
                    var v52 float32 = 0.1
                    if true {
                        var v53 uint64 = 1
                        v2[32767] = (!true)
                        _ = v53
                    } else {
                        var v54 uint16 = 42
                        v3 = fi46
                        _ = v54
                    }
                    _ = fi48
                    _ = v52
                }
            case "\n\t":
                for fi55 := 0; fi55 < 1; fi55++ {
                    if true {
                        v17 += (v17 & (^(v17)))
                    }
                    _ = fi55
                }
            default:
                v56 := GenS3{F0: (^(v19)), F1: (3.14159 != 0.0), F2: 2}
                _ = v56
            }
            _ = fi46
        }
    }
    delete(v2, int(v19))
    switch ("\n\t" + ("hello" + (("" + "a") + ""))) {
    case "a":
        for fi57 := 0; fi57 < 12; fi57++ {
            var v58 float32 = 1e10
            switch v3 {
            case 4294967295:
                for fi59 := 0; fi59 < 20; fi59++ {
                    fmt.Println(v3, v17)
                    delete(v5, (v6 ^ v6))
                    var v60 float32 = v58
                    _ = fi59
                    _ = v60
                }
            }
            var v61 int64 = -683
            _ = fi57
            _ = v58
            _ = v61
        }
    case "":
        v3 = v3
    case "\n\t":
        fmt.Println(v19, v3)
    default:
        _ = len(v1)
    }
    v19 -= v19
    _ = len(v1)
    v17 = v17
    v1 = append(v1, ("foo bar" + ("foo bar" + "hello")))
    v43 = v43
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
