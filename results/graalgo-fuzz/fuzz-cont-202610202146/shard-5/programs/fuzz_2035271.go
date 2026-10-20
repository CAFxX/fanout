package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 int; F1 uint64; F2 uint16 }
func genFn0() uint16 {
    var v6 float64 = 3.14159
    if true {
        v6 += float64((1.5 / 0.1))
    } else {
        for fi7 := 0; fi7 < 10; fi7++ {
            v6 -= v6
            _ = fi7
        }
        switch "foo bar" {
        case "hello":
            fmt.Println(v6)
        }
        switch -2147483648 {
        case 0:
            v8 := []int{128, -918, -1, 32767}
            _ = v8
        case -245:
            v9 := []int{9223372036854775807, 0}
            _ = v9
        default:
            v10 := GenS1{F0: 2}
            _ = v10
        }
    }
    if (-1 != 63) {
        if (0 == 68) {
            var v11 int16 = 746
            _ = v11
        } else {
            var v12 string = "foo bar"
            _ = v12
        }
        v6 += v6
        for fi13 := 0; fi13 < 18; fi13++ {
            if true {
                v14 := GenS2{F0: (fi13 + fi13), F1: -32768}
                var v15 int16 = 0
                if (0 >= 0) {
                    for fi16 := 0; fi16 < 20; fi16++ {
                        var v17 float64 = v6
                        v18 := []int{}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    v15 = v15
                    for fi19 := 0; fi19 < 18; fi19++ {
                        v6 = ((v6 / float64((float32((float64(((0.0 - (1e10 - float32(v6))) / (3.14159 - 0.1))) + ((((1e10 + v6) + 0.0) - v6) - -2.25))) - float32((v6 / v6))))) + v6)
                        v20 := []int{(^(fi13)), 856, fi19}
                        _ = fi19
                        _ = v20
                    }
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        switch "hello" {
                        case "foo bar":
                            for fi21 := 0; fi21 < 1; fi21++ {
                                fmt.Println(fi21)
                                v22 := GenS2{F0: fi13, F1: 1000}
                                switch "foo bar" {
                                case "hello":
                                    var v23 int16 = v15
                                    _ = v23
                                case "foo bar":
                                    v24 := GenS1{F0: 1}
                                    _ = v24
                                case "a":
                                    var v25 uint16 = 32767
                                    _ = v25
                                default:
                                    fmt.Println(fi21, v6)
                                }
                                if ((!false) || (-2147483648 <= 0)) {
                                    var v26 string = ("" + "\n\t")
                                    if false {
                                        switch fi21 {
                                        case 9223372036854775807:
                                            v27 := []int{fi21, (^(fi21))}
                                            _ = v27
                                        case 83:
                                            fmt.Println(v26, fi21, v6)
                                        default:
                                            switch (fi13 & 1) {
                                            case -19:
                                                switch (fi13 + (-(fi13))) {
                                                case -2147483648:
                                                    v28 := []int{(fi13 / (((fi21 << ((230) & 63))) | 1)), (fi13 ^ (^(fi21))), fi21}
                                                    _ = v28
                                                case 9223372036854775807:
                                                    var v29 uint8 = 250
                                                    _ = v29
                                                }
                                            case 9223372036854775807:
                                                var v30 uint8 = 2
                                                _ = v30
                                            case 0:
                                                var v31 uint64 = 18446744073709551615
                                                _ = v31
                                            default:
                                                v15 = (-(v15))
                                            }
                                        }
                                        if true {
                                            if false {
                                                v6 -= (v6 * 1.5)
                                            } else {
                                                var v32 bool = false
                                                v32 = (151 < 42)
                                                _ = v32
                                            }
                                        }
                                    } else {
                                        fmt.Println(v15)
                                        fmt.Println(v15, fi21, v15)
                                        v33 := []int{fi13, fi21}
                                        _ = v33
                                    }
                                    v34 := []int{29}
                                    _ = v26
                                    _ = v34
                                }
                                _ = fi21
                                _ = v22
                            }
                        case "a":
                            for fi35 := 0; fi35 < 3; fi35++ {
                                v36 := []int{}
                                v37 := GenS1{F0: 42}
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                        }
                    case "":
                        v38 := GenS3{F0: 775}
                        _ = v38
                    case "a":
                        switch ("\n\t" + "") {
                        case "":
                            v39 := GenS4{F0: 830, F1: 100}
                            _ = v39
                        default:
                            var v40 uint32 = 2147483647
                            _ = v40
                        }
                    default:
                        for fi41 := 0; fi41 < 1; fi41++ {
                            v42 := []int{127}
                            var v43 string = ((("a" + (("\n\t" + "") + "foo bar")) + ("" + "hello")) + ("" + "a"))
                            v44 := GenS1{F0: 65535}
                            fmt.Println(fi41)
                            _ = fi41
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                    }
                    var v45 uint8 = 207
                    v45 -= (^(v45))
                    _ = v45
                }
                _ = v14
                _ = v15
            } else {
                if (("a" + "a") <= ("a" + "a")) {
                    fmt.Println(fi13, v6)
                    v46 := []int{fi13, 833, fi13}
                    _ = v46
                }
                fmt.Println(v6, v6)
                var v47 uint16 = 42
                _ = v47
            }
            _ = fi13
        }
    } else {
        v48 := []int{0, 127, 128, 4294967295}
        if len(v48) > 0 {
            v48[0] = 671
        }
        _ = v48
    }
    _ = v6
    return 65535
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 0
    v3 := make(map[uint32]bool)
    v4 := make(map[string]uint32)
    var v5 uint16 = 2
    fmt.Println(v5, v5)
    v2 = 918
    v5 = (v5 >> (((v2 % ((1) | 1))) & 63))
    _ = len(v4)
    v49 := GenS4{F0: -1, F1: -1}
    var v50 int8 = -1
    if (genFn0() > v5) {
        var v51 uint32 = 0
        fmt.Println(v5, v2)
        _ = v51
    }
    v52 := GenS1{F0: 4294967295}
    v53 := []int{2147483647, 128}
    v50 = 100
    if false {
        var v54 uint64 = 18446744073709551615
        v54 = v54
        _ = v54
    } else {
        v55 := GenS5{F0: 42, F1: 4294967295, F2: 855}
        v5 = v5
        fmt.Println(v50)
        _ = v55
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v53:", len(v53), v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v52
    _ = v53
}
