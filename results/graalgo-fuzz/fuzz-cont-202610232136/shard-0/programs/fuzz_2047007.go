package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int; F1 uint32 }
func genFn0(p0 uint64, p1 uint32, p2 int8) uint32 {
    p0 = 42
    var v5 int8 = 100
    var v6 int32 = 582
    p2 = (v5 << ((0) & 63))
    p0 += (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (p1 + uint32(p0))
}
func genFn1() int16 {
    var v7 bool = false
    var v8 string = ("\n\t" + "a")
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint16 = 228
    var v2 string = ("hello" + "a")
    v3 := make(map[uint32]int)
    var v4 uint8 = 1
    switch -9223372036854775808 {
    case -6:
        v4 = v4
    case 42:
        var v9 uint16 = (^(v1))
        _ = v9
    case -1:
        _ = len(v3)
    default:
        for fi10 := 0; fi10 < 11; fi10++ {
            fmt.Println(v4, v4)
            var v11 int32 = -268
            _ = fi10
            _ = v11
        }
    }
    var v12 int32 = 42
    if true {
        for fi13 := 0; fi13 < 12; fi13++ {
            v14 := GenS1{F0: 128}
            v2 = (v2 + "foo bar")
            _ = fi13
            _ = v14
        }
        v1 -= (v1 | v1)
        v15 := GenS1{F0: (v4 & v4)}
        _ = v15
    } else {
        v16 := []int{4294967295, 32768}
        v12 -= int32(v1)
        _ = v16
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        v18 := []int{(fi17 * fi17)}
        _ = fi17
        _ = v18
    }
    if true {
        if false {
            if (1e10 < 3.14159) {
                for fi19 := 0; fi19 < 20; fi19++ {
                    var v20 uint8 = v4
                    if (!(("a" + ("foo bar" + ("foo bar" + v2))) < (("\n\t" + v2) + ""))) {
                        var v21 int64 = -554
                        _ = v21
                    }
                    _ = fi19
                    _ = v20
                }
            } else {
                v22 := GenS2{F0: 42, F1: 1}
                if false {
                    if (!(760 < 0)) {
                        var v23 uint16 = v1
                        _ = v23
                    }
                }
                _ = v22
            }
            if false {
                var v24 float64 = (float64(float32(0.1)) / (1.5 / 1e10))
                v2 = ((v2 + v2) + ("\n\t" + v2))
                switch -122 {
                case 4294967295:
                    var v25 int16 = 740
                    _ = v25
                }
                _ = v24
            }
            for fi26 := 0; fi26 < 20; fi26++ {
                if true {
                    var v27 uint32 = 2147483648
                    var v28 int8 = 1
                    for fi29 := 0; fi29 < 19; fi29++ {
                        v4 = 42
                        if false {
                            v12 = v12
                            v2 = "hello"
                        }
                        _ = fi29
                    }
                    _ = v27
                    _ = v28
                } else {
                    v30 := []int{}
                    v31 := GenS3{F0: true, F1: 329, F2: 0}
                    _ = v30
                    _ = v31
                }
                if (65535 == uint32(v1)) {
                    if false {
                        v32 := GenS1{F0: (v4 % (((v4 ^ 0)) | 1))}
                        var v33 int = fi26
                        var v34 int16 = 824
                        _ = v32
                        _ = v33
                        _ = v34
                    } else {
                        var v35 bool = true
                        v36 := GenS1{F0: (v4 % (((^(v4))) | 1))}
                        v35 = (v12 == (^(v12)))
                        _ = v35
                        _ = v36
                    }
                } else {
                    for fi37 := 0; fi37 < 5; fi37++ {
                        v38 := GenS1{F0: v4}
                        _ = fi37
                        _ = v38
                    }
                }
                _ = fi26
            }
        } else {
            v39 := v3[615]
            fmt.Println(v4, v2)
            switch "foo bar" {
            case "hello":
                v40 := []int{}
                _ = v40
            case "foo bar":
                v41 := GenS2{F0: 1, F1: -128}
                _ = v41
            case "a":
                for fi42 := 0; fi42 < 3; fi42++ {
                    for fi43 := 0; fi43 < 17; fi43++ {
                        var v44 uint8 = v4
                        fmt.Println(fi42, v4, v12)
                        switch ((v2 + "") + "a") {
                        case "":
                            v45 := GenS1{F0: 51}
                            _ = v45
                        case "a":
                            switch v2 {
                            case "foo bar":
                                var v46 uint64 = 2
                                _ = v46
                            }
                        case "\n\t":
                            var v47 int8 = -128
                            var v48 string = "\n\t"
                            _ = v47
                            _ = v48
                        default:
                            var v49 uint16 = 2
                            _ = v49
                        }
                        _ = fi43
                        _ = v44
                    }
                    _ = fi42
                }
            default:
                v50 := GenS3{F0: true, F1: v1, F2: (^(v4))}
                _ = v50
            }
            _ = v39
        }
        v1 -= v1
    } else {
        var v51 int8 = 127
        var v52 uint = uint(v1)
        _ = v51
        _ = v52
    }
    switch 435 {
    case 701:
        for fi53 := 0; fi53 < 7; fi53++ {
            fmt.Println(v12, v1)
            switch (v2 + "\n\t") {
            case "\n\t":
                v3[42] = fi53
            case "a":
                var v54 uint8 = (v4 | v4)
                _ = v54
            case "foo bar":
                v1 = 255
            }
            v4 = 127
            _ = fi53
        }
    }
    _ = len(v3)
    fmt.Println(v1)
    v55 := GenS3{F0: (true || true), F1: (v1 * (v1 % ((v1) | 1))), F2: (v4 & v4)}
    var v56 uint32 = 760
    switch v2 {
    case "\n\t":
        var v57 uint = 42
        _ = v57
    case "":
        v58 := GenS1{F0: v4}
        _ = v58
    case "foo bar":
        if true {
            v2 = v2
        } else {
            v59 := GenS2{F0: -354, F1: 0}
            var v60 uint64 = 2
            v61 := []int{0, -285, 9223372036854775807}
            _ = v59
            _ = v60
            _ = v61
        }
    default:
        var v62 int8 = 0
        _ = v62
    }
    switch "hello" {
    case "foo bar":
        v63 := GenS4{F0: 0, F1: uint32(v1)}
        _ = v63
    case "hello":
        v64 := GenS2{F0: int64(v56), F1: 78}
        _ = v64
    case "\n\t":
        for fi65 := 0; fi65 < 12; fi65++ {
            v56 -= v56
            _ = fi65
        }
    }
    if (v4 != (v4 << ((2) & 63))) {
        var v66 float64 = (0.1 - ((((float64(-2.25) * 0.0) + 1e10) * 0.1) * (((1.5 - -2.25) - (-2.25 - -2.25)) * -2.25)))
        _ = v66
    }
    v67 := []int{685, -756, 127, 128}
    fmt.Println(v2, v1, v2)
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v67:", len(v67), v67)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v67
}
