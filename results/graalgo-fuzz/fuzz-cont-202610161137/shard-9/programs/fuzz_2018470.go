package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 bool }
func genFn0(p0 int8, p1 uint32, p2 uint32) uint16 {
    p2 = 177
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int8 = p0
        var v8 int64 = 0
        for fi9 := 0; fi9 < 9; fi9++ {
            fmt.Println(v8, p0, p1)
            p1 = p2
            v10 := GenS1{F0: fi6}
            _ = fi9
            _ = v10
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v11 := []int{9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 0
}
func genFn1() int {
    var v12 uint = 701
    for fi13 := 0; fi13 < 19; fi13++ {
        if true {
            switch fi13 {
            case 571:
                for fi14 := 0; fi14 < 15; fi14++ {
                    if false {
                        var v15 float32 = (((3.14159 + 1.5) - float32((((0.1 * 0.0) + -2.25) / 0.1))) + ((0.0 - (1.5 * 0.0)) - 1.5))
                        var v16 float64 = (1e10 - float64(v15))
                        fmt.Println(v16, fi14, fi13)
                        _ = v15
                        _ = v16
                    } else {
                        if true {
                            fmt.Println(fi13, fi14)
                            v12 = (v12 & 42)
                        } else {
                            v17 := []int{}
                            _ = v17
                        }
                    }
                    var v18 string = "\n\t"
                    switch "a" {
                    case "hello":
                        var v19 string = "a"
                        _ = v19
                    case "":
                        var v20 int32 = 42
                        v20 = v20
                        _ = v20
                    }
                    _ = fi14
                    _ = v18
                }
            case -442:
                v21 := []int{2147483647}
                _ = v21
            case 0:
                for fi22 := 0; fi22 < 6; fi22++ {
                    v23 := []int{(fi13 ^ fi13), fi13}
                    if len(v23) > 0 {
                        v23[0] = 128
                    }
                    v12 -= v12
                    _ = fi22
                    _ = v23
                }
            }
            if false {
                switch "foo bar" {
                case "foo bar":
                    if (278 >= fi13) {
                        if (!(((42 == 513) || true) && (("hello" + ((("\n\t" + "hello") + ("" + (("foo bar" + "") + "hello"))) + "hello")) < "a"))) {
                            if false {
                                var v24 float64 = (3.14159 + (((float64(1.5) + 3.14159) + 1e10) / -2.25))
                                _ = v24
                            }
                        }
                        v25 := []int{}
                        if (218 > 1) {
                            var v26 float32 = 3.14159
                            for fi27 := 0; fi27 < 1; fi27++ {
                                var v28 float32 = 3.14159
                                _ = fi27
                                _ = v28
                            }
                            _ = v26
                        } else {
                            v29 := GenS2{F0: 42, F1: -9223372036854775808}
                            var v30 int32 = -821
                            v30 -= v30
                            _ = v29
                            _ = v30
                        }
                        _ = v25
                    }
                default:
                    v12 = (v12 % (((v12 * (v12 | v12))) | 1))
                }
                var v31 int64 = 42
                v32 := GenS1{F0: (fi13 >> (((v12 + v12)) & 63))}
                _ = v31
                _ = v32
            } else {
                var v33 int8 = 123
                v34 := []int{fi13, fi13, fi13, fi13}
                _ = v33
                _ = v34
            }
            v12 = (^(v12))
        }
        _ = fi13
    }
    var v35 float32 = 0.1
    var v36 uint64 = 516
    _ = v12
    _ = v35
    _ = v36
    return -2147483648
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[string]string)
    var v3 int = 128
    var v4 string = "a"
    v5 := make(map[string]uint32)
    for fi37 := 0; fi37 < 7; fi37++ {
        for fi38 := 0; fi38 < 8; fi38++ {
            fmt.Println(fi38)
            v39 := GenS2{F0: (v1 - 127), F1: 0}
            _ = fi38
            _ = v39
        }
        for fi40 := 0; fi40 < 14; fi40++ {
            v41 := []int{fi37, v3, 4294967295}
            fmt.Println(v1)
            v3 += (v3 ^ (fi40 + (v3 + -721)))
            _ = fi40
            _ = v41
        }
        _ = fi37
    }
    for fi42 := 0; fi42 < 17; fi42++ {
        v43 := []int{(v3 - (v3 << ((0) & 63)))}
        _ = fi42
        _ = v43
    }
    for fi44 := 0; fi44 < 9; fi44++ {
        var v45 string = v4
        var v46 uint = 42
        _ = fi44
        _ = v45
        _ = v46
    }
    switch "foo bar" {
    case "hello":
        var v47 int32 = 1
        _ = v47
    case "a":
        var v48 float32 = 0.1
        _ = v48
    }
    v1 += 128
    var v49 uint = 1
    if true {
        var v50 uint = v49
        _ = v50
    } else {
        if true {
            for fi51 := 0; fi51 < 20; fi51++ {
                v52 := GenS3{F0: true}
                _ = fi51
                _ = v52
            }
            var v53 float32 = (((1e10 / ((((3.14159 * 0.0) - 0.0) + float32(3.14159)) / ((1e10 / 0.1) + (0.1 + 3.14159)))) * 3.14159) / 3.14159)
            for fi54 := 0; fi54 < 5; fi54++ {
                v1 += (^(v1))
                for fi55 := 0; fi55 < 19; fi55++ {
                    fmt.Println(v4, v49)
                    _ = fi55
                }
                _ = fi54
            }
            _ = v53
        }
    }
    var v56 uint16 = 1
    fmt.Println(v4)
    v57 := []int{genFn1(), (v3 * (^(v3)))}
    v58 := GenS2{F0: (v1 | 0), F1: -9223372036854775808}
    var v59 uint32 = uint32(v56)
    var v60 float64 = 1.5
    if false {
        v56 += v56
        for fi61 := 0; fi61 < 9; fi61++ {
            v62 := GenS1{F0: (fi61 | v3)}
            var v63 string = "\n\t"
            v64 := GenS1{F0: fi61}
            _ = fi61
            _ = v62
            _ = v63
            _ = v64
        }
        v59 = v59
    }
    if true {
        fmt.Println(v60, v56)
    } else {
        v4 = ("hello" + v4)
        v1 -= v1
        v56 = (v56 >> ((v49) & 63))
    }
    v60 += -2.25
    if (540 < -375) {
        if false {
            for fi65 := 0; fi65 < 10; fi65++ {
                v66 := GenS3{F0: false}
                _ = fi65
                _ = v66
            }
        } else {
            _ = len(v5)
        }
        fmt.Println(v56, v59)
        v56 = genFn0(42, v59, 4294967295)
    }
    var v67 uint32 = v59
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v60
    _ = v67
}
