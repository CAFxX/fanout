package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int8; F1 uint16; F2 int8 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 int8; F1 int64; F2 int32 }
func genFn0() int64 {
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint = 1
        v9 := []int{42, fi7, (fi7 / ((fi7) | 1))}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 bool = false
    if true {
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := GenS1{F0: 42, F1: -1}
            if false {
                v13 := GenS1{F0: 127, F1: 42}
                if v10 {
                    if v10 {
                        v10 = v10
                        v10 = (344 > 0)
                    } else {
                        if true {
                            for fi14 := 0; fi14 < 13; fi14++ {
                                fmt.Println(fi11, v10, fi11)
                                fmt.Println(v10, fi11, fi14)
                                fmt.Println(fi14, fi14)
                                v15 := []int{fi11, -1}
                                _ = fi14
                                _ = v15
                            }
                        }
                        if v10 {
                            var v16 uint8 = 0
                            var v17 int8 = 1
                            v17 += -1
                            _ = v16
                            _ = v17
                        } else {
                            switch "hello" {
                            case "hello":
                                if ((833 <= 65535) || (!((-1 < -630) && (fi11 != (-(fi11)))))) {
                                    for fi18 := 0; fi18 < 9; fi18++ {
                                        v19 := GenS2{F0: 255, F1: 1}
                                        _ = fi18
                                        _ = v19
                                    }
                                    switch "a" {
                                    case "a":
                                        fmt.Println(fi11)
                                    }
                                    v20 := []int{fi11, (^(fi11)), 0}
                                    _ = v20
                                } else {
                                    v10 = true
                                    var v21 string = ""
                                    fmt.Println(v21, v21, v21)
                                    _ = v21
                                }
                            case "foo bar":
                                switch ("" + ("a" + (("" + "") + "a"))) {
                                case "foo bar":
                                    var v22 float32 = 0.0
                                    _ = v22
                                case "a":
                                    v23 := GenS1{F0: 42, F1: -70}
                                    _ = v23
                                default:
                                    var v24 float64 = float64(1e10)
                                    _ = v24
                                }
                            case "a":
                                v25 := GenS1{F0: 0, F1: -1}
                                _ = v25
                            default:
                                v26 := []int{fi11, fi11, fi11}
                                _ = v26
                            }
                        }
                    }
                }
                v10 = ((!((true && (v10 || (1 == 1))) || (0.0 >= 0.0))) && true)
                _ = v13
            } else {
                v27 := []int{fi11, (fi11 & fi11), 128, (fi11 << ((786) & 63))}
                var v28 uint16 = 1
                v29 := GenS2{F0: v28, F1: 2}
                _ = v27
                _ = v28
                _ = v29
            }
            _ = fi11
            _ = v12
        }
        v30 := GenS1{F0: -128, F1: 0}
        var v31 uint8 = 255
        _ = v30
        _ = v31
    } else {
        var v32 float64 = (float64(0.1) - 0.1)
        fmt.Println(v32)
        _ = v32
    }
    var v33 uint8 = 8
    v10 = (366 != 42)
    _ = v10
    _ = v33
    return -1
}
func genFn1(p0 uint16) int16 {
    fmt.Println(p0)
    var v34 uint16 = p0
    switch "\n\t" {
    case "\n\t":
        p0 = 65535
    case "a":
        v34 -= p0
    default:
        v34 -= (p0 & p0)
    }
    _ = p0
    _ = v34
    return 32767
}

func main() {
    v1 := []int{-1, -982, 829}
    var v2 int = 270
    var v3 int8 = 1
    v4 := make(map[string]string)
    var v5 uint = 988
    var v6 int8 = v3
    v35 := v4[((("hello" + "") + "foo bar") + "foo bar")]
    switch (^(v2)) {
    case 639:
        v1 = append(v1, (v2 + (v2 % ((v2) | 1))))
    case -687:
        fmt.Println(v3)
    case 42:
        var v36 uint32 = 65535
        _ = v36
    }
    fmt.Println(v35)
    for fi37 := 0; fi37 < 16; fi37++ {
        for fi38 := 0; fi38 < 16; fi38++ {
            switch fi38 {
            case 627:
                switch v35 {
                case "foo bar":
                    if false {
                        var v39 bool = false
                        v2 += fi37
                        fmt.Println(fi38, v5, fi38)
                        _ = v39
                    }
                    v3 += v6
                case "hello":
                    var v40 uint16 = 2
                    _ = v40
                case "":
                    v41 := GenS3{F0: (v6 << ((v5) & 63)), F1: 978, F2: 127}
                    _ = v41
                default:
                    switch (v35 + "a") {
                    case "\n\t":
                        v42 := GenS4{F0: 1, F1: -764}
                        _ = v42
                    case "":
                        if len(v1) > 0 {
                            v1[0] = (fi37 % ((360) | 1))
                        }
                    case "a":
                        var v43 bool = (true && false)
                        _ = v43
                    }
                }
            }
            v2 -= (v2 >> (((v5 % (((v5 % ((v5) | 1))) | 1))) & 63))
            var v44 uint8 = 0
            _ = fi38
            _ = v44
        }
        delete(v4, "\n\t")
        var v45 int8 = (v3 % ((v6) | 1))
        _ = fi37
        _ = v45
    }
    v2 -= v2
    v3 = v3
    v46 := GenS5{F0: v2, F1: 128}
    var v47 uint = 0
    v35 = v35
    v5 = 42
    _ = len(v4)
    if false {
        var v48 uint8 = 128
        v49 := GenS4{F0: 2147483647, F1: -1}
        _ = v48
        _ = v49
    }
    v50 := GenS2{F0: 65535, F1: 997}
    v35 = v35
    fmt.Println(v47, v35)
    fmt.Println(v35, v2)
    v51 := GenS6{F0: 0, F1: 42, F2: -351}
    var v52 uint64 = 1000
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
}
