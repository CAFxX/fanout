package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int16 }
type GenS2 struct { F0 int; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint32; F1 string; F2 int16 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 bool }
func genFn0(p0 int8, p1 int32, p2 bool) int32 {
    fmt.Println(p2)
    v7 := []int{int(p0), 4294967295, 0, 42}
    fmt.Println(p1, p0, p0)
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (^(p1))
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{}
    v3 := []string{("" + "\n\t"), "", "foo bar", (("hello" + "foo bar") + "a")}
    var v4 int16 = -1
    v5 := make(map[int]bool)
    var v6 int32 = 42
    if true {
        var v8 uint64 = 18446744073709551615
        v6 = genFn0(-128, v6, (387 != 255))
        _ = v8
    } else {
        var v9 uint8 = 0
        var v10 uint16 = 1
        if true {
            v11 := GenS1{F0: 1, F1: v4, F2: (v4 << ((4294967295) & 63))}
            _ = v11
        } else {
            v12 := GenS1{F0: 525, F1: v4, F2: 313}
            switch ("\n\t" + "hello") {
            case "a":
                fmt.Println(v9, v6, v9)
            case "foo bar":
                if true {
                    for fi13 := 0; fi13 < 8; fi13++ {
                        if ((!true) && true) {
                            v6 = v6
                            v14 := GenS1{F0: 246, F1: 32767, F2: v4}
                            v9 = 255
                            _ = v14
                        }
                        _ = fi13
                    }
                    if (v4 == v4) {
                        var v15 uint64 = 1
                        fmt.Println(v6, v4, v4)
                        v6 += 0
                        _ = v15
                    }
                    if ((42 >= -53) && (2147483647 != 1)) {
                        for fi16 := 0; fi16 < 7; fi16++ {
                            var v17 uint16 = (^(v10))
                            v9 -= (v9 - v9)
                            if false {
                                fmt.Println(fi16)
                            }
                            _ = fi16
                            _ = v17
                        }
                        var v18 int32 = (v6 & (^(v6)))
                        v19 := GenS1{F0: 255, F1: int16(v9), F2: 0}
                        _ = v18
                        _ = v19
                    }
                }
            case "hello":
                var v20 string = "hello"
                _ = v20
            default:
                for fi21 := 0; fi21 < 7; fi21++ {
                    v10 -= v10
                    v4 = -552
                    _ = fi21
                }
            }
            _ = v12
        }
        _ = v10
        _ = v9
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        switch (^(fi22)) {
        case 128:
            if false {
                v4 = int16(fi22)
                var v23 float64 = 1e10
                delete(v5, -854)
                _ = v23
            }
        }
        _ = fi22
    }
    switch -2147483648 {
    case 127:
        switch "hello" {
        case "\n\t":
            v24 := GenS1{F0: 0, F1: v4, F2: v4}
            _ = v24
        case "foo bar":
            var v25 int64 = 0
            _ = v25
        case "hello":
            v26 := GenS2{F0: 4294967295, F1: v4, F2: 2147483647}
            _ = v26
        default:
            if (true || (!true)) {
                var v27 uint32 = 2147483647
                _ = v27
            } else {
                var v28 uint32 = 0
                _ = v28
            }
        }
    case 2147483647:
        for fi29 := 0; fi29 < 3; fi29++ {
            if (true || (!(float64(1e10) != (float64(0.0) * (1.5 * float64(((-2.25 + ((0.1 / (0.1 + 0.0)) / -2.25)) * ((float32(3.14159) * (1.5 + 3.14159)) / 1.0)))))))) {
                v30 := GenS2{F0: 260, F1: -1, F2: 12}
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    switch "" {
                    case "a":
                        _ = len(v1)
                    case "":
                        for fi31 := 0; fi31 < 4; fi31++ {
                            _ = len(v1)
                            v32 := v5[fi29]
                            var v33 string = "foo bar"
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                    }
                    var v34 float64 = ((((0.0 / 1.5) - 0.1) + 3.14159) / (1e10 * (float64(((float32(0.1) * 1.5) + (3.14159 + (-2.25 / (float32(-2.25) - (1e10 / 0.1)))))) - (0.1 + 1.5))))
                    _ = v34
                }
                _ = v30
            } else {
                for fi35 := 0; fi35 < 10; fi35++ {
                    var v36 uint32 = 4294967295
                    _ = fi35
                    _ = v36
                }
            }
            v4 += (v4 ^ v4)
            v4 = v4
            _ = fi29
        }
    case 0:
        var v37 uint = 0
        _ = v37
    }
    if true {
        for fi38 := 0; fi38 < 10; fi38++ {
            v4 -= (^(v4))
            v39 := GenS3{F0: 0, F1: "\n\t", F2: (v4 << ((65535) & 63))}
            for fi40 := 0; fi40 < 6; fi40++ {
                v6 = genFn0(-1, (v6 % ((v6) | 1)), ((1 <= 248) && (!false)))
                _ = fi40
            }
            _ = fi38
            _ = v39
        }
        v6 = (v6 / ((-890) | 1))
    }
    for fi41 := 0; fi41 < 10; fi41++ {
        for fi42 := 0; fi42 < 18; fi42++ {
            delete(v1, "")
            _ = fi42
        }
        switch (fi41 >> ((891) & 63)) {
        case 4294967295:
            var v43 string = ("foo bar" + "a")
            _ = v43
        default:
            var v44 int64 = -1
            _ = v44
        }
        _ = fi41
    }
    var v45 bool = (!true)
    if v45 {
        v45 = v45
        var v46 int32 = v6
        _ = v46
    }
    v45 = (737 == 2)
    switch "hello" {
    case "":
        for fi47 := 0; fi47 < 15; fi47++ {
            if (!v45) {
                if (v4 <= (^(v4))) {
                    v6 = v6
                }
                delete(v5, 2147483647)
                v3 = append(v3, "\n\t")
            } else {
                var v48 int64 = 9223372036854775807
                var v49 int64 = (v48 >> ((255) & 63))
                _ = v48
                _ = v49
            }
            _ = len(v3)
            var v50 int16 = v4
            _ = fi47
            _ = v50
        }
    }
    var v51 string = "\n\t"
    var v52 float64 = float64(-2.25)
    for fi53 := 0; fi53 < 11; fi53++ {
        for fi54 := 0; fi54 < 17; fi54++ {
            v4 = v4
            var v55 bool = v45
            _ = fi54
            _ = v55
        }
        fmt.Println(v45, v51)
        _ = fi53
    }
    v56 := GenS4{F0: 1, F1: (-(v4))}
    switch -9223372036854775808 {
    case 9223372036854775807:
        if (v45 || (!v45)) {
            switch v51 {
            case "":
                v52 = float64(float32((((3.14159 + (3.14159 + (float64(float32(v52)) + v52))) - (-2.25 * v52)) * ((v52 + v52) / float64(((((0.1 - 1e10) + ((3.14159 - -2.25) / ((0.1 * 3.14159) + 1e10))) - (-2.25 - -2.25)) + (0.1 - -2.25)))))))
            case "\n\t":
                var v57 uint32 = 2147483647
                _ = v57
            default:
                switch 9223372036854775807 {
                case 1:
                    v5[32768] = (0 >= (-(v6)))
                }
            }
            delete(v1, "a")
        } else {
            for fi58 := 0; fi58 < 19; fi58++ {
                v59 := GenS2{F0: (fi58 << ((4294967295) & 63)), F1: 1, F2: 4294967295}
                v60 := GenS5{F0: true}
                v51 = v51
                _ = fi58
                _ = v59
                _ = v60
            }
            var v61 string = v51
            var v62 uint64 = 4294967295
            _ = v61
            _ = v62
        }
    case 32767:
        for fi63 := 0; fi63 < 2; fi63++ {
            v6 = (v6 | v6)
            v6 = -139
            _ = len(v5)
            _ = fi63
        }
    default:
        var v64 string = ("\n\t" + (v51 + ("\n\t" + v51)))
        _ = v64
    }
    for fi65 := 0; fi65 < 14; fi65++ {
        var v66 int8 = 1
        v67 := v1[v51]
        fmt.Println(v51, v66, v67)
        _ = fi65
        _ = v66
        _ = v67
    }
    var v68 uint8 = 170
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v51
    _ = v52
    _ = v56
    _ = v6
    _ = v68
}
