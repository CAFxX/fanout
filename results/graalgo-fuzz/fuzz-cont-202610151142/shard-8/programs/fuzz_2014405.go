package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 uint32; F1 int }
func genFn0(p0 int, p1 int32) uint32 {
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 uint = 4294967295
        _ = fi6
        _ = v7
    }
    switch "foo bar" {
    case "hello":
        v8 := []int{}
        _ = v8
    }
    if false {
        var v9 float64 = float64(0.0)
        if (!false) {
            v10 := GenS1{F0: 65535}
            _ = v10
        }
        var v11 bool = (true || ((p0 >= p0) || false))
        _ = v11
        _ = v9
    }
    fmt.Println(p1, p1)
    if true {
        if (true && (255 != 255)) {
            var v12 uint32 = 42
            p1 = 2147483647
            v13 := GenS2{F0: 255}
            _ = v12
            _ = v13
        }
        v14 := GenS3{F0: "", F1: 127}
        _ = v14
    } else {
        var v15 uint16 = 63
        var v16 int32 = p1
        v17 := GenS4{F0: 2147483648}
        _ = v15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 uint8, p1 int) uint {
    if (2147483648 >= 4294967295) {
        if (true && true) {
            var v18 int16 = -448
            for fi19 := 0; fi19 < 17; fi19++ {
                v20 := []int{}
                var v21 uint = 87
                _ = fi19
                _ = v20
                _ = v21
            }
            if false {
                switch (p1 ^ -1) {
                case -421:
                    v22 := []int{p1}
                    _ = v22
                case 32768:
                    for fi23 := 0; fi23 < 6; fi23++ {
                        v18 = (^(v18))
                        v18 = v18
                        _ = fi23
                    }
                }
                v18 = 1
            } else {
                switch (p1 >> ((0) & 63)) {
                case 0:
                    var v24 uint16 = 432
                    _ = v24
                case -9223372036854775808:
                    v18 -= (v18 ^ (-(v18)))
                case 4294967295:
                    p0 += (p0 * p0)
                }
                fmt.Println(p1, v18)
                var v25 string = "foo bar"
                _ = v25
            }
            _ = v18
        }
        for fi26 := 0; fi26 < 4; fi26++ {
            switch ("a" + "foo bar") {
            case "hello":
                v27 := GenS5{F0: -32768}
                _ = v27
            case "\n\t":
                p1 += 179
            case "":
                v28 := []int{p1, p1, fi26}
                _ = v28
            default:
                p0 = p0
            }
            for fi29 := 0; fi29 < 13; fi29++ {
                var v30 uint16 = 42
                v31 := []int{(p1 >> ((65535) & 63)), 0, (p1 - fi26), fi29}
                _ = fi29
                _ = v30
                _ = v31
            }
            v32 := GenS6{F0: 2, F1: 1}
            _ = fi26
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 4; fi33++ {
        fmt.Println(p0, fi33)
        _ = fi33
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := []string{"hello", ""}
    v2 := make(map[string]bool)
    var v3 uint64 = 18446744073709551615
    var v4 int = -2147483648
    var v5 string = "foo bar"
    var v34 uint64 = 128
    v35 := GenS5{F0: -687}
    var v36 bool = (168 != 301)
    if (v36 || v36) {
        v34 = 0
    }
    v36 = (!((v5 + "foo bar") < (v5 + v5)))
    _ = len(v2)
    v3 = v3
    _ = len(v1)
    switch 466 {
    case 190:
        var v37 uint32 = genFn0((^(v4)), 0)
        _ = v37
    case 737:
        var v38 uint16 = 1
        _ = v38
    case 32767:
        v39 := GenS5{F0: -511}
        _ = v39
    default:
        for fi40 := 0; fi40 < 1; fi40++ {
            v41 := GenS2{F0: 128}
            v4 = (fi40 + fi40)
            _ = fi40
            _ = v41
        }
    }
    if v36 {
        var v42 int32 = -2147483648
        switch "hello" {
        case "hello":
            v43 := GenS7{F0: genFn0(v4, (v42 ^ v42)), F1: (v4 - v4)}
            _ = v43
        case "a":
            switch (v4 << ((18446744073709551615) & 63)) {
            case -9223372036854775808:
                switch -1 {
                case 32768:
                    v44 := GenS4{F0: 65535}
                    _ = v44
                case 1:
                    _ = len(v2)
                default:
                    v2[(v5 + v5)] = (!(65535 > 4294967295))
                }
            case 1:
                fmt.Println(v34, v5, v36)
            case 127:
                for fi45 := 0; fi45 < 5; fi45++ {
                    fmt.Println(fi45)
                    _ = fi45
                }
            }
        }
        _ = v42
    }
    if (!(!true)) {
        if ((-2.25 / 1.5) <= 0.0) {
            fmt.Println(v4, v36, v5)
            switch v5 {
            case "foo bar":
                v36 = v36
            case "a":
                var v46 float32 = ((((0.0 + 3.14159) / -2.25) * float32((3.14159 + -2.25))) - 1e10)
                _ = v46
            case "hello":
                if (297 != 796) {
                    _ = len(v1)
                    v5 = "foo bar"
                } else {
                    for fi47 := 0; fi47 < 19; fi47++ {
                        v48 := GenS6{F0: 255, F1: -198}
                        fmt.Println(v5, v34, v3)
                        for fi49 := 0; fi49 < 14; fi49++ {
                            v4 -= fi47
                            var v50 uint8 = 132
                            _ = fi49
                            _ = v50
                        }
                        _ = fi47
                        _ = v48
                    }
                }
            }
            if (!true) {
                v51 := GenS2{F0: 226}
                v34 = (^(v3))
                for fi52 := 0; fi52 < 1; fi52++ {
                    v4 -= -304
                    if (!v36) {
                        if false {
                            var v53 float64 = float64(1e10)
                            if (!(0 >= 2)) {
                                _ = len(v1)
                            } else {
                                var v54 bool = v36
                                _ = v54
                            }
                            _ = v53
                        }
                    } else {
                        for fi55 := 0; fi55 < 20; fi55++ {
                            var v56 int32 = -1
                            fmt.Println(v5)
                            var v57 float64 = ((-2.25 / -2.25) + 0.1)
                            _ = fi55
                            _ = v56
                            _ = v57
                        }
                        var v58 int32 = -2147483648
                        _ = v58
                    }
                    _ = fi52
                }
                _ = v51
            } else {
                if len(v1) > 0 {
                    v1[0] = (("foo bar" + "hello") + (v5 + "\n\t"))
                }
            }
        }
        v59 := GenS7{F0: 1, F1: 120}
        v34 += 592
        _ = v59
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
