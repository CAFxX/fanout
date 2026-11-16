package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int64; F1 uint64; F2 bool }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 uint8; F1 uint8; F2 int }
func genFn0(p0 uint32, p1 uint64) uint16 {
    var v6 uint64 = (p1 | p1)
    var v7 int8 = 100
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 uint, p1 uint32, p2 int8) int16 {
    fmt.Println(p2)
    fmt.Println(p2, p0, p1)
    switch -1 {
    case -328:
        v8 := GenS1{F0: 546}
        _ = v8
    case 2147483647:
        if true {
            var v9 bool = (45 < 0)
            _ = v9
        } else {
            v10 := GenS2{F0: true}
            var v11 float64 = 3.14159
            _ = v10
            _ = v11
        }
    case 1:
        if true {
            var v12 uint = (p0 + p0)
            p0 = v12
            if true {
                fmt.Println(p1, p0)
            }
            _ = v12
        }
    }
    p2 = (p2 >> ((p0) & 63))
    switch "" {
    case "foo bar":
        var v13 uint64 = 2
        _ = v13
    case "a":
        for fi14 := 0; fi14 < 8; fi14++ {
            p1 -= (p1 % ((p1) | 1))
            v15 := []int{fi14, 4294967295, -1, -2147483648}
            switch -789 {
            case 550:
                var v16 uint32 = p1
                _ = v16
            default:
                if false {
                    if true {
                        switch "" {
                        case "hello":
                            if (0 >= 418) {
                                p2 = p2
                            } else {
                                switch "\n\t" {
                                case "foo bar":
                                    for fi17 := 0; fi17 < 3; fi17++ {
                                        var v18 int = (fi17 % (((fi14 >> ((p0) & 63))) | 1))
                                        _ = fi17
                                        _ = v18
                                    }
                                case "a":
                                    var v19 uint16 = genFn0(987, 42)
                                    _ = v19
                                }
                            }
                        case "\n\t":
                            var v20 int8 = (p2 << (((p0 + p0)) & 63))
                            _ = v20
                        }
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                    } else {
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        var v21 uint8 = 128
                        _ = v21
                    }
                    fmt.Println(p0)
                    var v22 uint64 = 548
                    _ = v22
                } else {
                    v23 := GenS3{F0: p1, F1: p2}
                    var v24 float64 = (float64(float32(1e10)) * 0.0)
                    _ = v23
                    _ = v24
                }
            }
            _ = fi14
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 1000
}

func main() {
    v1 := []string{"hello"}
    v2 := []int{9223372036854775807, 1, 127, 2147483647}
    v3 := []int{504}
    var v4 int = 9223372036854775807
    v5 := make(map[string]string)
    for fi25 := 0; fi25 < 10; fi25++ {
        var v26 bool = true
        _ = fi25
        _ = v26
    }
    var v27 float32 = 0.1
    if true {
        v28 := GenS1{F0: 1}
        v4 = v4
        _ = v28
    } else {
        for fi29 := 0; fi29 < 18; fi29++ {
            if true {
                fmt.Println(v27, v27, fi29)
            } else {
                for fi30 := 0; fi30 < 4; fi30++ {
                    fmt.Println(v4, v27, v27)
                    _ = fi30
                }
                for fi31 := 0; fi31 < 8; fi31++ {
                    v32 := v5["a"]
                    v32 = v32
                    v33 := GenS2{F0: true}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                fmt.Println(v4)
            }
            switch ("" + "hello") {
            case "foo bar":
                v34 := GenS2{F0: false}
                _ = v34
            case "\n\t":
                v35 := GenS2{F0: true}
                _ = v35
            case "":
                var v36 uint32 = 2
                _ = v36
            default:
                switch "a" {
                case "foo bar":
                    _ = len(v2)
                case "hello":
                    for fi37 := 0; fi37 < 9; fi37++ {
                        v5[("a" + ("hello" + "hello"))] = ("\n\t" + "\n\t")
                        switch "" {
                        case "":
                            v27 -= float32((float64(v27) / 1.0))
                        default:
                            var v38 int32 = -2147483648
                            _ = v38
                        }
                        _ = fi37
                    }
                case "a":
                    switch "a" {
                    case "a":
                        var v39 uint16 = 733
                        _ = v39
                    case "":
                        var v40 float32 = (v27 + v27)
                        _ = v40
                    case "\n\t":
                        var v41 uint16 = 0
                        _ = v41
                    default:
                        fmt.Println(fi29)
                    }
                }
            }
            _ = fi29
        }
        var v42 int32 = 0
        _ = v42
    }
    v27 = v27
    v4 = v4
    for fi43 := 0; fi43 < 16; fi43++ {
        for fi44 := 0; fi44 < 19; fi44++ {
            fmt.Println(fi43, v4)
            _ = len(v2)
            var v45 float64 = 3.14159
            _ = fi44
            _ = v45
        }
        fmt.Println(v4, fi43)
        _ = fi43
    }
    for fi46 := 0; fi46 < 19; fi46++ {
        fmt.Println(fi46, fi46, fi46)
        fmt.Println(v4, v27)
        _ = fi46
    }
    switch "foo bar" {
    case "\n\t":
        var v47 uint16 = 32767
        _ = v47
    case "":
        if true {
            v5[("" + "foo bar")] = ("foo bar" + ("\n\t" + "foo bar"))
            switch (v4 - (v4 | (^(v4)))) {
            case 9223372036854775807:
                if (v4 >= 32767) {
                    v27 = (float32(0.1) * v27)
                    if true {
                        v48 := v5["\n\t"]
                        _ = v48
                    }
                } else {
                    v49 := GenS4{F0: -9223372036854775808, F1: 864, F2: (!(4294967295 >= 896))}
                    for fi50 := 0; fi50 < 17; fi50++ {
                        var v51 int64 = 42
                        for fi52 := 0; fi52 < 16; fi52++ {
                            fmt.Println(fi52, v4, v27)
                            fmt.Println(v27, fi50)
                            switch v4 {
                            case 1:
                                v53 := GenS5{F0: (1 > 42), F1: 1}
                                _ = v53
                            case 127:
                                var v54 uint = 636
                                _ = v54
                            case -2147483648:
                                if true {
                                    v5["hello"] = "foo bar"
                                    fmt.Println(v4)
                                } else {
                                    _ = len(v5)
                                    v4 = fi52
                                }
                            }
                            switch v4 {
                            case -2147483648:
                                v4 = fi50
                            case 1:
                                switch (-(fi52)) {
                                case -1:
                                    var v55 int = -1
                                    _ = v55
                                }
                            }
                            _ = fi52
                        }
                        _ = fi50
                        _ = v51
                    }
                    _ = v49
                }
            default:
                if (1 < 441) {
                    if len(v3) > 0 {
                        v3[0] = -662
                    }
                    var v56 int64 = 42
                    _ = v56
                } else {
                    v57 := GenS4{F0: 9223372036854775807, F1: 4294967295, F2: ((v27 / v27) < ((v27 - v27) / (0.1 + 1e10)))}
                    v58 := GenS6{F0: 0, F1: 2, F2: v4}
                    _ = v57
                    _ = v58
                }
            }
            var v59 int32 = 2147483647
            _ = v59
        }
    }
    v27 = v27
    v27 = 0.0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
