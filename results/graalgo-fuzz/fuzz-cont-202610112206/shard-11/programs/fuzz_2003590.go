package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 uint; F1 uint32; F2 uint64 }
func genFn0(p0 int64, p1 uint32) int64 {
    v5 := GenS1{F0: 2147483647, F1: 2}
    for fi6 := 0; fi6 < 14; fi6++ {
        if true {
            for fi7 := 0; fi7 < 19; fi7++ {
                p0 -= (p0 * p0)
                _ = fi7
            }
            if true {
                var v8 int64 = (p0 % ((p0) | 1))
                var v9 uint64 = 18446744073709551615
                _ = v8
                _ = v9
            } else {
                var v10 string = ""
                for fi11 := 0; fi11 < 11; fi11++ {
                    var v12 int8 = 0
                    _ = fi11
                    _ = v12
                }
                _ = v10
            }
            v13 := []int{fi6, 4294967295, (-(fi6))}
            _ = v13
        }
        if (false && ("hello" >= (("" + "a") + "a"))) {
            for fi14 := 0; fi14 < 9; fi14++ {
                v15 := []int{}
                switch "\n\t" {
                case "foo bar":
                    if false {
                        v16 := GenS2{F0: 1000}
                        if true {
                            v17 := GenS1{F0: p1, F1: 33}
                            v15 = append(v15, (^(fi14)))
                            _ = v17
                        }
                        _ = v16
                    }
                default:
                    if (1 != 0) {
                        fmt.Println(p1, p0)
                    }
                }
                switch fi14 {
                case 41:
                    fmt.Println(p0)
                case 970:
                    var v18 int64 = p0
                    v19 := GenS1{F0: (p1 & p1), F1: 128}
                    v20 := GenS1{F0: p1, F1: 189}
                    _ = v18
                    _ = v19
                    _ = v20
                case 1:
                    v21 := GenS1{F0: p1, F1: 255}
                    _ = v21
                default:
                    if false {
                        if (!((^(p0)) <= p0)) {
                            p1 -= (^(p1))
                            var v22 float64 = float64(-2.25)
                            _ = v22
                        }
                    }
                }
                _ = fi14
                _ = v15
            }
        }
        fmt.Println(fi6, p0)
        _ = fi6
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v5
    return (-(p0))
}
func genFn1(p0 bool, p1 int) uint {
    fmt.Println(p0, p0)
    if (!(128 == 42)) {
        if ((p0 && true) || (!p0)) {
            switch (^(p1)) {
            case 127:
                if p0 {
                    if (!(!(!p0))) {
                        p1 = (p1 | (^(p1)))
                    }
                } else {
                    p1 = (^(p1))
                }
            case -1:
                for fi23 := 0; fi23 < 10; fi23++ {
                    v24 := GenS2{F0: 42}
                    var v25 int = -9223372036854775808
                    for fi26 := 0; fi26 < 10; fi26++ {
                        if (0 < 255) {
                            var v27 int = p1
                            fmt.Println(p1, fi23, fi23)
                            var v28 uint = 0
                            _ = v27
                            _ = v28
                        }
                        var v29 int8 = 100
                        switch "foo bar" {
                        case "foo bar":
                            v30 := GenS3{F0: (v29 | -1), F1: 500}
                            _ = v30
                        case "a":
                            v31 := GenS3{F0: v29, F1: 113}
                            _ = v31
                        case "hello":
                            var v32 int8 = 42
                            _ = v32
                        }
                        _ = fi26
                        _ = v29
                    }
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
            v33 := GenS2{F0: 59}
            _ = v33
        } else {
            var v34 float32 = -2.25
            var v35 float64 = (float64(((float32(0.0) - 1e10) * 3.14159)) * float64(((((v34 + (float32(1e10) * v34)) * (v34 / (0.1 * (v34 - 0.1)))) + 3.14159) / (v34 * (0.1 - v34)))))
            _ = v34
            _ = v35
        }
    } else {
        fmt.Println(p0, p0)
    }
    for fi36 := 0; fi36 < 11; fi36++ {
        v37 := []int{}
        v37 = append(v37, (-(fi36)))
        var v38 int16 = -229
        _ = fi36
        _ = v37
        _ = v38
    }
    if p0 {
        switch (p1 + (p1 - (p1 >> ((726) & 63)))) {
        case 586:
            switch ("a" + "\n\t") {
            case "hello":
                fmt.Println(p1)
            case "":
                v39 := []int{p1, 4294967295, p1}
                _ = v39
            case "a":
                var v40 uint32 = 65535
                _ = v40
            default:
                v41 := GenS4{F0: (p1 * p1), F1: (p1 / ((p1) | 1))}
                _ = v41
            }
        default:
            if p0 {
                v42 := GenS4{F0: p1, F1: p1}
                v43 := []int{1}
                switch 63 {
                case 0:
                    v44 := GenS1{F0: 65535, F1: 42}
                    _ = v44
                case 222:
                    p1 = 974
                default:
                    var v45 uint64 = 0
                    _ = v45
                }
                _ = v42
                _ = v43
            }
        }
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    return 255
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 186
    v3 := make(map[string]string)
    var v4 float64 = float64((((3.14159 * (-2.25 - -2.25)) * float32(((float64(((1.5 + 1.5) - 0.1)) / (1e10 + 3.14159)) / 3.14159))) / ((((float32(3.14159) - 1e10) + ((float32(-2.25) / (1e10 + 3.14159)) * 1.5)) - 1e10) - float32((0.1 * float64((0.0 * float32(0.1))))))))
    var v46 int64 = 9223372036854775807
    fmt.Println(v2, v4, v46)
    var v47 float64 = 1.5
    v48 := GenS5{F0: 342, F1: 445, F2: 1}
    v47 -= 3.14159
    var v49 uint = 4294967295
    v47 += 0.0
    var v50 int16 = v2
    if false {
        for fi51 := 0; fi51 < 10; fi51++ {
            v46 = (^(v46))
            _ = fi51
        }
        v50 -= -849
        var v52 float32 = ((float32(v47) * 1.5) * 0.1)
        _ = v52
    } else {
        v47 = v47
        v46 = (v46 % ((-751) | 1))
        switch "" {
        case "\n\t":
            fmt.Println(v4, v2)
        case "a":
            switch -759 {
            case 42:
                for fi53 := 0; fi53 < 18; fi53++ {
                    v47 = float64(float32((v4 * float64(((float32(v4) / float32((((0.1 + v47) / 1.0) * (v47 * float64(1.5))))) + (((0.0 * 1e10) + 0.0) - -2.25))))))
                    _ = fi53
                }
            case 4294967295:
                fmt.Println(v47)
            default:
                var v54 float64 = ((v47 * v4) + float64(1.5))
                _ = v54
            }
        default:
            var v55 int64 = 42
            _ = v55
        }
    }
    switch ("" + "foo bar") {
    case "":
        v4 = (0.1 / v4)
    case "hello":
        fmt.Println(v4)
    case "foo bar":
        v2 += (v50 >> ((0) & 63))
    default:
        var v56 uint8 = 1
        _ = v56
    }
    var v57 uint = v49
    var v58 uint64 = 42
    v59 := v3["hello"]
    v60 := GenS1{F0: 0, F1: 128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
    _ = v57
    _ = v58
    _ = v59
    _ = v60
}
