package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
func genFn0() int16 {
    v5 := GenS1{F0: -635, F1: 65535, F2: 1}
    var v6 int8 = -128
    switch 32768 {
    case 0:
        if ((!false) || false) {
            var v7 uint8 = 2
            if false {
                for fi8 := 0; fi8 < 4; fi8++ {
                    v7 = (v7 ^ 42)
                    _ = fi8
                }
                var v9 int64 = 959
                switch "hello" {
                case "":
                    var v10 float64 = float64(1e10)
                    _ = v10
                case "foo bar":
                    if false {
                        var v11 uint64 = 18446744073709551615
                        _ = v11
                    }
                }
                _ = v9
            }
            switch ("foo bar" + (("\n\t" + "foo bar") + ("" + (("a" + ("hello" + (("" + ("a" + "")) + "foo bar"))) + ("\n\t" + "\n\t"))))) {
            case "hello":
                fmt.Println(v7, v7)
            case "a":
                v6 = 0
            case "\n\t":
                if true {
                    var v12 float64 = 1.5
                    if false {
                        var v13 float64 = ((v12 + v12) - ((-2.25 + 0.0) * v12))
                        for fi14 := 0; fi14 < 1; fi14++ {
                            v7 += v7
                            var v15 uint16 = 2
                            _ = fi14
                            _ = v15
                        }
                        v16 := GenS1{F0: -550, F1: 201, F2: 2}
                        _ = v13
                        _ = v16
                    } else {
                        switch (("\n\t" + "") + "hello") {
                        case "hello":
                            if ((!true) || ((v7 / ((v7) | 1)) > 255)) {
                                v12 -= ((v12 * v12) + (v12 * v12))
                            } else {
                                var v17 int32 = 2147483647
                                _ = v17
                            }
                        case "foo bar":
                            v18 := GenS1{F0: int32(v7), F1: 18446744073709551615, F2: 1}
                            _ = v18
                        }
                        if (true || (1.5 >= 0.0)) {
                            for fi19 := 0; fi19 < 6; fi19++ {
                                v20 := []int{-855, (fi19 << ((18446744073709551615) & 63)), (fi19 / (((fi19 * fi19)) | 1))}
                                if (false && (v6 == v6)) {
                                    if ((v6 >> ((83) & 63)) <= (v6 % (((v6 * v6)) | 1))) {
                                        for fi21 := 0; fi21 < 17; fi21++ {
                                            v6 = -1
                                            _ = fi21
                                        }
                                        v7 = (v7 / (((v7 ^ (v7 >> ((915) & 63)))) | 1))
                                    } else {
                                        switch (-(fi19)) {
                                        case 32768:
                                            if true {
                                                var v22 uint32 = 65535
                                                _ = v22
                                            }
                                        }
                                        v6 += 1
                                    }
                                    for fi23 := 0; fi23 < 11; fi23++ {
                                        var v24 string = (("foo bar" + "\n\t") + ("\n\t" + "a"))
                                        _ = fi23
                                        _ = v24
                                    }
                                    v12 = float64(1e10)
                                }
                                for fi25 := 0; fi25 < 3; fi25++ {
                                    v12 += 0.1
                                    if (!(42 == 2)) {
                                        if true {
                                            v26 := GenS1{F0: 42, F1: 321, F2: 42}
                                            if (128 != v7) {
                                                switch (fi19 - fi19) {
                                                case 822:
                                                    fmt.Println(fi25, v7)
                                                default:
                                                    fmt.Println(fi19, v7, fi19)
                                                }
                                                if false {
                                                    fmt.Println(v6, fi19)
                                                    if len(v20) > 0 {
                                                        v20[0] = fi19
                                                    }
                                                }
                                            }
                                            _ = v26
                                        }
                                        if len(v20) > 0 {
                                            v20[0] = fi25
                                        }
                                    }
                                    v27 := GenS2{F0: 18446744073709551615}
                                    _ = fi25
                                    _ = v27
                                }
                                _ = fi19
                                _ = v20
                            }
                            v28 := []int{4294967295}
                            v29 := GenS2{F0: 1}
                            _ = v28
                            _ = v29
                        }
                    }
                    v30 := GenS2{F0: 27}
                    _ = v12
                    _ = v30
                }
            default:
                v31 := []int{128, 128}
                _ = v31
            }
            _ = v7
        }
    case -871:
        for fi32 := 0; fi32 < 6; fi32++ {
            var v33 uint64 = 296
            v33 = v33
            if false {
                v34 := []int{fi32, (fi32 | fi32), (fi32 >> ((255) & 63))}
                _ = v34
            }
            _ = fi32
            _ = v33
        }
    case 128:
        v35 := []int{-9223372036854775808, 710, 924}
        _ = v35
    }
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int32, p1 int8, p2 uint) uint {
    v36 := GenS3{F0: 128}
    p2 = (p2 << (((p2 << (((p2 << ((p2) & 63))) & 63))) & 63))
    var v37 uint = (^(p2))
    p1 += p1
    for fi38 := 0; fi38 < 9; fi38++ {
        switch "" {
        case "foo bar":
            v39 := GenS2{F0: 80}
            _ = v39
        case "\n\t":
            p0 = (p0 << ((v37) & 63))
        case "hello":
            fmt.Println(v37, p1)
        }
        for fi40 := 0; fi40 < 11; fi40++ {
            fmt.Println(p0, fi40)
            p2 = 42
            _ = fi40
        }
        _ = fi38
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v36
    _ = v37
    return (p2 ^ 18446744073709551615)
}

func main() {
    v1 := []string{"\n\t", "hello", "foo bar"}
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    var v4 int64 = 42
    v41 := GenS2{F0: 42}
    var v42 float32 = ((0.0 + (-2.25 / 3.14159)) + float32(0.1))
    fmt.Println(v42)
    for fi43 := 0; fi43 < 17; fi43++ {
        fmt.Println(fi43, v42)
        var v44 uint32 = 148
        _ = fi43
        _ = v44
    }
    for fi45 := 0; fi45 < 2; fi45++ {
        if len(v1) > 0 {
            v1[0] = ""
        }
        var v46 float64 = -2.25
        _ = fi45
        _ = v46
    }
    var v47 bool = (!false)
    if (v47 || true) {
        if len(v1) > 0 {
            v1[0] = ""
        }
    } else {
        if false {
            var v48 int8 = 42
            _ = v48
        }
        var v49 uint = genFn1(1, 100, 18446744073709551615)
        v49 = 18446744073709551615
        _ = v49
    }
    v42 = v42
    switch 4294967295 {
    case -1:
        for fi50 := 0; fi50 < 6; fi50++ {
            for fi51 := 0; fi51 < 4; fi51++ {
                var v52 float64 = (3.14159 - 0.0)
                if false {
                    v53 := GenS2{F0: 4294967295}
                    var v54 int8 = -128
                    v42 -= (v42 - -2.25)
                    _ = v53
                    _ = v54
                }
                _ = fi51
                _ = v52
            }
            switch fi50 {
            case -2147483648:
                var v55 uint = 2
                _ = v55
            case 0:
                v47 = (-1 <= v4)
            }
            _ = fi50
        }
    case 128:
        v2[-912] = 4294967295
    case -9223372036854775808:
        if len(v1) > 0 {
            v1[0] = ""
        }
    }
    for fi56 := 0; fi56 < 3; fi56++ {
        var v57 int8 = -128
        fmt.Println(v57, v4)
        _ = fi56
        _ = v57
    }
    var v58 uint16 = 32767
    v47 = v47
    v59 := GenS3{F0: 127}
    if (((v42 + 0.0) / ((v42 * float32(0.1)) - float32((0.1 - 0.1)))) > ((float32((((1.5 / (-2.25 + -2.25)) + 1e10) - (float64(-2.25) / float64(1.5)))) + (float32(0.1) + (v42 + 1e10))) * v42)) {
        var v60 bool = v47
        v61 := GenS1{F0: -2147483648, F1: 360, F2: 18446744073709551615}
        v47 = false
        _ = v60
        _ = v61
    }
    var v62 uint8 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("v58:", v58)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v47
    _ = v58
    _ = v59
    _ = v62
}
