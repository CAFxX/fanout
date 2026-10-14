package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 uint8; F2 int16 }
func genFn0(p0 bool) int {
    if p0 {
        p0 = (!p0)
        p0 = (65535 == 433)
        var v6 uint32 = 851
        _ = v6
    } else {
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 uint = 2
            v9 := []int{(fi7 * 127), (fi7 | fi7), fi7, -204}
            _ = fi7
            _ = v8
            _ = v9
        }
        if ((0.0 - 0.0) == 0.1) {
            var v10 int = 32768
            if true {
                v11 := []int{(v10 >> ((0) & 63)), (v10 & v10), (v10 * -137), (v10 << ((715) & 63))}
                switch -1 {
                case -596:
                    _ = len(v11)
                case 128:
                    p0 = false
                case 42:
                    for fi12 := 0; fi12 < 9; fi12++ {
                        var v13 uint8 = 0
                        if p0 {
                            switch ("\n\t" + ("a" + ("" + "foo bar"))) {
                            case "hello":
                                fmt.Println(v13, v10)
                            case "":
                                p0 = (100 != -128)
                            }
                        }
                        if false {
                            for fi14 := 0; fi14 < 20; fi14++ {
                                var v15 int = (fi12 / (((v10 + fi12)) | 1))
                                _ = fi14
                                _ = v15
                            }
                            if (!false) {
                                var v16 int64 = -580
                                _ = v16
                            }
                            for fi17 := 0; fi17 < 20; fi17++ {
                                if (!(((3.14159 * 1.5) == (1.5 * (3.14159 / (0.0 - float64(float32(0.1)))))) || (1000 != 32767))) {
                                    p0 = (!(("\n\t" >= ("hello" + "foo bar")) && true))
                                    var v18 string = "hello"
                                    _ = v18
                                }
                                _ = fi17
                            }
                        }
                        _ = fi12
                        _ = v13
                    }
                default:
                    v19 := GenS1{F0: p0}
                    _ = v19
                }
                _ = v11
            } else {
                v20 := GenS1{F0: (199 != 4294967295)}
                var v21 int32 = 2147483647
                _ = v20
                _ = v21
            }
            _ = v10
        }
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        p0 = p0
        v23 := []int{-616}
        _ = fi22
        _ = v23
    }
    switch ((("hello" + "hello") + "a") + (("" + "a") + "foo bar")) {
    case "foo bar":
        fmt.Println(p0)
    }
    var v24 uint64 = 18446744073709551615
    if (-1 == -1) {
        var v25 uint16 = 874
        p0 = false
        _ = v25
    }
    _ = p0
    _ = v24
    return 4294967295
}
func genFn1(p0 int16) int {
    var v26 uint8 = 62
    v27 := GenS1{F0: true}
    if false {
        if true {
            v26 = v26
            for fi28 := 0; fi28 < 7; fi28++ {
                v26 += 40
                v29 := []int{fi28, fi28, 128}
                _ = fi28
                _ = v29
            }
            var v30 bool = false
            _ = v30
        } else {
            switch 32767 {
            case 9223372036854775807:
                fmt.Println(v26, p0)
            case 128:
                if ((!(false && false)) || true) {
                    fmt.Println(v26, p0)
                    v31 := GenS2{F0: (p0 | 0)}
                    _ = v31
                }
            default:
                var v32 float64 = 0.1
                _ = v32
            }
        }
    } else {
        var v33 uint16 = 2
        _ = v33
    }
    if (p0 == -32768) {
        fmt.Println(v26, v26)
        for fi34 := 0; fi34 < 1; fi34++ {
            v35 := GenS2{F0: p0}
            var v36 int32 = -1
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    fmt.Println(v26)
    _ = p0
    _ = v26
    _ = v27
    return 422
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    var v4 float32 = 1.5
    v5 := []uint32{0, 21, 42, 42}
    for fi37 := 0; fi37 < 5; fi37++ {
        v38 := GenS3{F0: 127}
        var v39 int8 = 127
        _ = fi37
        _ = v38
        _ = v39
    }
    if (0 < 1) {
        if (957 <= -120) {
            var v40 bool = (!true)
            v1 -= (^(v1))
            _ = v40
        }
    } else {
        v1 = (v1 + 255)
        if false {
            fmt.Println(v1, v1)
            var v41 int16 = -200
            _ = v41
        } else {
            switch -9223372036854775808 {
            case 9223372036854775807:
                for fi42 := 0; fi42 < 15; fi42++ {
                    v5 = append(v5, 42)
                    var v43 uint32 = 2147483648
                    _ = fi42
                    _ = v43
                }
            default:
                var v44 uint = 93
                _ = v44
            }
        }
    }
    v45 := GenS4{F0: 843, F1: (v1 / ((v1) | 1)), F2: -1}
    for fi46 := 0; fi46 < 2; fi46++ {
        v47 := GenS3{F0: 100}
        _ = fi46
        _ = v47
    }
    switch -1 {
    case -9223372036854775808:
        for fi48 := 0; fi48 < 17; fi48++ {
            var v49 int16 = 1
            v50 := GenS2{F0: (v49 % ((32767) | 1))}
            _ = fi48
            _ = v49
            _ = v50
        }
    default:
        var v51 int8 = -122
        _ = v51
    }
    var v52 string = ("a" + "\n\t")
    delete(v3, -533)
    v53 := GenS3{F0: -128}
    for fi54 := 0; fi54 < 17; fi54++ {
        var v55 uint16 = 467
        v55 -= v55
        _ = fi54
        _ = v55
    }
    var v56 int8 = -110
    v1 += 2
    if len(v5) > 0 {
        v5[0] = 2
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v52
    _ = v53
    _ = v56
}
