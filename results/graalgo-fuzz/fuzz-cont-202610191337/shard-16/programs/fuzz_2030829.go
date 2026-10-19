package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool; F1 int; F2 uint8 }
func genFn0(p0 int16, p1 int16, p2 uint32) uint16 {
    var v5 int8 = 0
    v6 := GenS1{F0: -32768}
    v5 = v5
    v7 := []int{}
    if (-9223372036854775808 == -9223372036854775808) {
        var v8 int16 = 1000
        p1 += (p0 - (p1 << ((716) & 63)))
        if len(v7) > 0 {
            v7[0] = 128
        }
        _ = v8
    } else {
        if (!(!(29 == 0))) {
            p2 = p2
        }
        switch (("\n\t" + "") + "hello") {
        case "\n\t":
            var v9 float64 = 0.0
            _ = v9
        case "foo bar":
            _ = len(v7)
        }
        switch 1 {
        case 2147483647:
            var v10 int32 = 721
            _ = v10
        case 970:
            if len(v7) > 0 {
                v7[0] = -436
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 bool) uint8 {
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := []int{445, fi11, fi11, fi11}
        v13 := GenS1{F0: 1000}
        var v14 int64 = -390
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        v16 := GenS2{F0: 42, F1: 0, F2: 227}
        var v17 uint32 = 525
        if p0 {
            if false {
                if (1 == 745) {
                    v18 := []int{-9223372036854775808, fi15, 9223372036854775807}
                    var v19 int32 = 2147483647
                    _ = v18
                    _ = v19
                }
                v17 += 42
                v20 := []int{}
                _ = v20
            }
            var v21 int8 = 1
            var v22 int64 = -9223372036854775808
            _ = v21
            _ = v22
        } else {
            v23 := GenS2{F0: 2, F1: 519, F2: 2}
            var v24 bool = false
            _ = v23
            _ = v24
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        switch "foo bar" {
        case "foo bar":
            for fi26 := 0; fi26 < 11; fi26++ {
                switch fi26 {
                case 178:
                    if p0 {
                        var v27 int = -961
                        v28 := GenS1{F0: 0}
                        _ = v27
                        _ = v28
                    } else {
                        v29 := GenS3{F0: -1, F1: -898}
                        _ = v29
                    }
                case 32767:
                    var v30 float32 = float32(-2.25)
                    _ = v30
                case 127:
                    v31 := GenS4{F0: 1, F1: "\n\t"}
                    _ = v31
                }
                p0 = (!((-128 < 0) || false))
                var v32 string = "\n\t"
                _ = fi26
                _ = v32
            }
        case "a":
            var v33 uint16 = 583
            _ = v33
        case "hello":
            if ((p0 || p0) || (!p0)) {
                var v34 int8 = 103
                v35 := []int{(fi25 << ((94) & 63))}
                _ = v34
                _ = v35
            }
        }
        for fi36 := 0; fi36 < 20; fi36++ {
            if ("foo bar" >= "a") {
                var v37 uint64 = 42
                var v38 bool = (false || p0)
                v38 = p0
                fmt.Println(v38)
                v37 -= 2
                _ = v37
                _ = v38
            }
            var v39 int64 = 9223372036854775807
            _ = fi36
            _ = v39
        }
        _ = fi25
    }
    _ = p0
    return 127
}

func main() {
    var v1 bool = (!true)
    var v2 int = 128
    var v3 uint64 = 0
    v4 := make(map[int]string)
    var v40 int16 = 1000
    if true {
        for fi41 := 0; fi41 < 12; fi41++ {
            var v42 float64 = (1e10 - (1e10 / 1.5))
            var v43 uint32 = 42
            delete(v4, v2)
            _ = fi41
            _ = v42
            _ = v43
        }
        var v44 bool = (42 > 2147483647)
        var v45 uint8 = 1
        _ = v44
        _ = v45
    } else {
        v46 := []int{(v2 / ((128) | 1)), (v2 / ((v2) | 1))}
        if v1 {
            v47 := GenS2{F0: 255, F1: 0, F2: v3}
            switch int(v3) {
            case 363:
                v40 = v40
            case 32768:
                if v1 {
                    fmt.Println(v40, v2, v40)
                    v2 = 945
                } else {
                    v3 = v3
                    v2 = (v2 / ((127) | 1))
                    v40 -= (-(v40))
                }
            case 9223372036854775807:
                var v48 bool = ("" != "a")
                _ = v48
            }
            _ = v47
        } else {
            fmt.Println(v40, v2)
            if (true || v1) {
                if v1 {
                    var v49 float64 = 0.0
                    switch (v2 ^ (v2 | (-(v2)))) {
                    case 4294967295:
                        v40 += -1
                    default:
                        var v50 string = "a"
                        _ = v50
                    }
                    _ = v49
                }
            } else {
                if false {
                    v51 := GenS5{F0: genFn1(v1)}
                    var v52 uint16 = 1
                    _ = v51
                    _ = v52
                }
                fmt.Println(v3, v2)
                v53 := GenS6{F0: (113 >= -1), F1: -473, F2: 2}
                _ = v53
            }
            for fi54 := 0; fi54 < 1; fi54++ {
                delete(v4, 2147483647)
                for fi55 := 0; fi55 < 18; fi55++ {
                    if (691 <= 760) {
                        fmt.Println(v2, fi54)
                        v56 := GenS4{F0: 42, F1: "hello"}
                        _ = v56
                    } else {
                        fmt.Println(fi55)
                        v46 = append(v46, (fi55 - fi54))
                        v40 += 42
                    }
                    _ = fi55
                }
                switch (-(fi54)) {
                case 0:
                    switch (v2 / ((v2) | 1)) {
                    case 127:
                        for fi57 := 0; fi57 < 5; fi57++ {
                            var v58 int8 = int8(v2)
                            _ = fi57
                            _ = v58
                        }
                    }
                case 510:
                    var v59 bool = (0.1 > (0.0 / 1e10))
                    _ = v59
                default:
                    if v1 {
                        var v60 uint32 = 65535
                        fmt.Println(v2)
                        v4[(fi54 & 128)] = ("foo bar" + (("a" + "") + "\n\t"))
                        _ = v60
                    } else {
                        var v61 string = "foo bar"
                        var v62 float32 = (float32(0.1) + (3.14159 + float32(float64(1.5))))
                        _ = v61
                        _ = v62
                    }
                }
                _ = fi54
            }
        }
        v3 = (v3 * (v3 + v3))
        _ = v46
    }
    v1 = (v1 || v1)
    v63 := GenS4{F0: 9223372036854775807, F1: "\n\t"}
    switch "\n\t" {
    case "foo bar":
        var v64 int64 = 1
        _ = v64
    default:
        var v65 int64 = 0
        _ = v65
    }
    var v66 uint64 = v3
    v67 := GenS3{F0: -343, F1: 1}
    var v68 int = v2
    var v69 int8 = 100
    v40 = v40
    v40 += v40
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v66:", v66)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v63
    _ = v66
    _ = v67
    _ = v68
    _ = v69
}
