package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }
func genFn0() int8 {
    if true {
        for fi4 := 0; fi4 < 9; fi4++ {
            v5 := []int{}
            _ = fi4
            _ = v5
        }
        v6 := []int{127, -9223372036854775808}
        v6 = append(v6, 9223372036854775807)
        _ = v6
    } else {
        var v7 uint8 = 229
        v7 = 245
        for fi8 := 0; fi8 < 19; fi8++ {
            v7 -= 255
            var v9 uint16 = 255
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    var v10 int16 = 655
    v11 := GenS1{F0: 2147483647, F1: 255, F2: 65535}
    var v12 string = "hello"
    _ = v10
    _ = v11
    _ = v12
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 157
    v3 = (v3 * v3)
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 14; fi14++ {
            switch fi14 {
            case 1:
                v3 = v3
            case 635:
                v3 = (v3 * 1)
            case 0:
                v15 := []int{-1}
                _ = v15
            default:
                var v16 uint16 = 341
                _ = v16
            }
            v17 := []int{}
            if true {
                var v18 int = (-(fi14))
                if true {
                    if false {
                        var v19 string = "foo bar"
                        var v20 int64 = 1
                        if true {
                            fmt.Println(fi14)
                            if ((-1 == -75) && false) {
                                var v21 string = ("foo bar" + (("a" + v19) + ("hello" + "a")))
                                _ = v21
                            } else {
                                var v22 int32 = 1
                                v17 = append(v17, v18)
                                _ = v22
                            }
                            v19 = (v19 + (v19 + v19))
                            if (false && true) {
                                fmt.Println(v19, v18, v19)
                                v3 += (v3 << ((4294967295) & 63))
                            } else {
                                if true {
                                    v23 := GenS1{F0: 2147483648, F1: 0, F2: 2147483648}
                                    var v24 float64 = 0.1
                                    _ = v23
                                    _ = v24
                                } else {
                                    v19 = v19
                                    fmt.Println(fi13, v20, v18)
                                }
                                v19 = ("" + v19)
                                var v25 float32 = float32(((0.0 + -2.25) / 1.0))
                                if (0 <= 1) {
                                    var v26 uint64 = 800
                                    fmt.Println(fi13, fi13)
                                    v27 := GenS2{F0: "hello"}
                                    if (!false) {
                                        fmt.Println(fi14)
                                        var v28 string = v19
                                        _ = v28
                                    } else {
                                        var v29 uint32 = 1
                                        _ = v29
                                    }
                                    _ = v26
                                    _ = v27
                                } else {
                                    v19 = v19
                                    v19 = "hello"
                                    var v30 int16 = 1000
                                    _ = v30
                                }
                                fmt.Println(v3, fi14, v18)
                                _ = v25
                            }
                        }
                        _ = v19
                        _ = v20
                    }
                    var v31 int = (fi13 - (fi13 & (fi14 % (((^(fi14))) | 1))))
                    _ = len(v17)
                    _ = v31
                }
                _ = v18
            }
            _ = fi14
            _ = v17
        }
        var v32 bool = false
        _ = fi13
        _ = v32
    }
    switch "foo bar" {
    case "":
        var v33 bool = (762 < 559)
        _ = v33
    case "\n\t":
        v3 = (v3 << ((1) & 63))
    case "a":
        v3 = v3
    }
    v3 = v3
    if false {
        var v34 uint16 = 32767
        var v35 float64 = -2.25
        v36 := GenS3{F0: 42}
        _ = v34
        _ = v35
        _ = v36
    }
    var v37 uint64 = 641
    var v38 int8 = -111
    var v39 uint8 = v3
    var v40 int = -1
    for fi41 := 0; fi41 < 10; fi41++ {
        v37 += (v37 << ((42) & 63))
        _ = fi41
    }
    switch "hello" {
    case "a":
        fmt.Println(v40, v39)
    case "\n\t":
        v42 := GenS4{F0: -1}
        _ = v42
    case "":
        var v43 string = "a"
        _ = v43
    }
    var v44 int16 = int16(v40)
    if false {
        v45 := []int{(v40 * (v40 & v40)), (v40 / ((v40) | 1)), 985, 1}
        fmt.Println(v3)
        _ = v45
    }
    var v46 uint64 = 2
    v47 := v1[859]
    v48 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v44
    _ = v46
    _ = v47
    _ = v48
}
