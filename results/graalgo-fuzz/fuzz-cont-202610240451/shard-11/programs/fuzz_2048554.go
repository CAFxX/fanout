package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 bool; F2 bool }
type GenS5 struct { F0 uint8; F1 uint32; F2 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string; F1 bool; F2 uint16 }
type GenS8 struct { F0 int64; F1 uint32; F2 int32 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint64 }
func genFn0(p0 bool, p1 uint64) uint64 {
    v4 := GenS1{F0: 9223372036854775807, F1: 393, F2: 2}
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: fi5, F1: 1000, F2: 2}
        p0 = (true || (!p0))
        if p0 {
            var v7 uint16 = 65535
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    if true {
        v8 := GenS2{F0: 128, F1: "a"}
        var v9 int32 = -1
        var v10 float64 = 1.5
        _ = v10
        _ = v8
        _ = v9
    } else {
        v11 := []int{127}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v4
    return (p1 * (p1 - 18446744073709551615))
}

func main() {
    var v1 uint = 796
    v2 := make(map[uint32]bool)
    v3 := []uint32{1, 204, 447}
    if (!true) {
        switch ("a" + "foo bar") {
        case "a":
            v12 := GenS3{F0: -9223372036854775808}
            _ = v12
        case "\n\t":
            v1 = (v1 - (v1 - (v1 ^ v1)))
        default:
            v1 -= v1
        }
        var v13 uint = v1
        _ = v13
    }
    var v14 int = 2147483647
    var v15 float32 = 0.1
    if false {
        switch "a" {
        case "a":
            if len(v3) > 0 {
                v3[0] = 1
            }
        case "\n\t":
            var v16 int = (v14 + v14)
            _ = v16
        }
        v15 = 1e10
        if ((v1 + v1) > (v1 >> (((v1 & v1)) & 63))) {
            switch "hello" {
            case "foo bar":
                v17 := GenS3{F0: 642}
                _ = v17
            case "hello":
                if true {
                    var v18 int64 = 1
                    fmt.Println(v15, v15, v14)
                    v14 = v14
                    _ = v18
                } else {
                    var v19 int32 = -455
                    v20 := GenS4{F0: 0, F1: false, F2: (128 == 128)}
                    for fi21 := 0; fi21 < 1; fi21++ {
                        if (!(455 < 1)) {
                            fmt.Println(v14)
                            v22 := v2[2147483647]
                            switch (v14 / ((fi21) | 1)) {
                            case 1:
                                var v23 uint64 = 18446744073709551615
                                _ = v23
                            case 4294967295:
                                v1 = v1
                            default:
                                fmt.Println(v15, v15)
                            }
                            _ = v22
                        } else {
                            var v24 int8 = -128
                            v24 += 100
                            _ = v24
                        }
                        _ = fi21
                    }
                    _ = v19
                    _ = v20
                }
            default:
                v1 = (v1 + 4294967295)
            }
            fmt.Println(v14, v14, v14)
        } else {
            v3 = append(v3, uint32(v1))
            v15 = (1.5 - v15)
        }
    } else {
        var v25 float32 = v15
        _ = v25
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        if (!true) {
            switch "\n\t" {
            case "\n\t":
                _ = len(v3)
            case "a":
                if len(v3) > 0 {
                    v3[0] = 771
                }
            case "foo bar":
                for fi27 := 0; fi27 < 8; fi27++ {
                    fmt.Println(v1, fi26)
                    var v28 int32 = 2147483647
                    _ = fi27
                    _ = v28
                }
            default:
                if true {
                    for fi29 := 0; fi29 < 18; fi29++ {
                        var v30 uint = (v1 - v1)
                        v31 := GenS5{F0: 255, F1: 2147483648, F2: (-(fi26))}
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    v32 := GenS6{F0: (fi26 << ((v1) & 63))}
                    for fi33 := 0; fi33 < 15; fi33++ {
                        switch v14 {
                        case 9223372036854775807:
                            var v34 int64 = 9223372036854775807
                            _ = v34
                        default:
                            v35 := GenS1{F0: fi33, F1: 42, F2: 281}
                            _ = v35
                        }
                        _ = fi33
                    }
                    _ = v32
                } else {
                    if false {
                        if len(v3) > 0 {
                            v3[0] = 0
                        }
                        v1 -= 404
                    }
                    var v36 int16 = 0
                    for fi37 := 0; fi37 < 20; fi37++ {
                        v14 -= (-(fi26))
                        v36 = v36
                        switch v14 {
                        case 32767:
                            fmt.Println(fi26, v1)
                        }
                        _ = fi37
                    }
                    _ = v36
                }
            }
        }
        _ = fi26
    }
    delete(v2, 377)
    for fi38 := 0; fi38 < 16; fi38++ {
        var v39 float64 = 0.1
        _ = fi38
        _ = v39
    }
    v14 = 32768
    if ((((-2.25 / -2.25) / v15) != v15) || false) {
        var v40 string = "hello"
        var v41 uint32 = 65535
        switch v14 {
        case -998:
            v1 -= (^(v1))
        case 4294967295:
            var v42 int64 = -1
            _ = v42
        default:
            delete(v2, (^(v41)))
        }
        _ = v40
        _ = v41
    } else {
        v2[0] = false
        if false {
            if (2147483647 < 65535) {
                for fi43 := 0; fi43 < 8; fi43++ {
                    v44 := GenS7{F0: ((((("hello" + "hello") + ("" + ("foo bar" + "hello"))) + "") + ("\n\t" + ("\n\t" + ""))) + ""), F1: (!false), F2: 647}
                    v45 := GenS8{F0: 0, F1: 4294967295, F2: 2147483647}
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                v46 := GenS6{F0: v14}
                switch ("\n\t" + "hello") {
                case "":
                    var v47 float64 = 0.0
                    _ = v47
                default:
                    v48 := GenS9{F0: 1000}
                    _ = v48
                }
                _ = v46
            }
            var v49 float32 = 1.5
            for fi50 := 0; fi50 < 3; fi50++ {
                v49 -= v49
                _ = fi50
            }
            _ = v49
        } else {
            fmt.Println(v14)
            if true {
                var v51 uint8 = 190
                if (false || false) {
                    v52 := GenS5{F0: 127, F1: 65535, F2: v14}
                    _ = v52
                } else {
                    v51 = 0
                    v53 := GenS1{F0: 128, F1: -587, F2: 596}
                    v15 = v15
                    _ = v53
                }
                var v54 float32 = 0.1
                _ = v51
                _ = v54
            }
        }
        v55 := GenS10{F0: 2}
        _ = v55
    }
    var v56 uint32 = 65535
    fmt.Println(v1)
    var v57 uint32 = v56
    for fi58 := 0; fi58 < 20; fi58++ {
        var v59 uint32 = 2147483648
        v1 -= (v1 - v1)
        if len(v3) > 0 {
            v3[0] = (v57 + v56)
        }
        _ = fi58
        _ = v59
    }
    for fi60 := 0; fi60 < 18; fi60++ {
        var v61 int64 = 284
        var v62 float32 = v15
        var v63 int16 = int16(fi60)
        _ = fi60
        _ = v61
        _ = v62
        _ = v63
    }
    for fi64 := 0; fi64 < 8; fi64++ {
        v57 = (^(v57))
        var v65 int64 = 9223372036854775807
        _ = len(v3)
        _ = fi64
        _ = v65
    }
    var v66 int8 = 26
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v56
    _ = v57
    _ = v66
}
