package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int16 }
type GenS2 struct { F0 int16; F1 bool; F2 int16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int }
func genFn0() uint64 {
    switch (("" + "foo bar") + ("a" + "hello")) {
    case "\n\t":
        if false {
            var v7 float64 = (3.14159 / 1.5)
            switch 32768 {
            case 32767:
                v7 -= v7
            case 42:
                var v8 int8 = 125
                _ = v8
            case 222:
                var v9 int8 = -44
                _ = v9
            default:
                v10 := []int{-677}
                _ = v10
            }
            if (2 >= 0) {
                v7 = 1.5
            } else {
                v11 := []int{4294967295, -306}
                var v12 float64 = 3.14159
                _ = v11
                _ = v12
            }
            _ = v7
        } else {
            v13 := GenS1{F0: 1, F1: -32768, F2: 1}
            _ = v13
        }
    case "a":
        var v14 uint16 = 2
        _ = v14
    }
    var v15 float64 = (0.0 / 1e10)
    var v16 float64 = -2.25
    v17 := GenS1{F0: 44, F1: 223, F2: -808}
    for fi18 := 0; fi18 < 8; fi18++ {
        for fi19 := 0; fi19 < 19; fi19++ {
            v20 := []int{fi18, 127, -295, (fi18 * -2147483648)}
            _ = fi19
            _ = v20
        }
        v21 := GenS2{F0: 1, F1: false, F2: -712}
        for fi22 := 0; fi22 < 20; fi22++ {
            if false {
                fmt.Println(fi22, v15, fi22)
            } else {
                if (!(!(false || ("" >= "foo bar")))) {
                    fmt.Println(fi22)
                    v23 := []int{32767}
                    var v24 int64 = 0
                    _ = v23
                    _ = v24
                }
                switch "foo bar" {
                case "a":
                    if false {
                        if (2 != 42) {
                            v15 = (v15 - 0.0)
                        }
                        var v25 uint16 = 1
                        _ = v25
                    }
                case "hello":
                    switch ("" + "") {
                    case "\n\t":
                        v26 := GenS1{F0: 42, F1: -32768, F2: 32767}
                        if false {
                            fmt.Println(fi22, fi18, fi22)
                            v27 := []int{(fi18 - (fi18 * (fi22 % ((-1) | 1)))), (fi18 - (fi22 << ((355) & 63))), fi22}
                            v15 = float64(((float32(v16) + (float32(v16) / 1e10)) * -2.25))
                            _ = v27
                        } else {
                            var v28 uint8 = 166
                            v29 := []int{(fi22 | fi18), fi22}
                            v15 = (1e10 / 3.14159)
                            _ = v28
                            _ = v29
                        }
                        _ = v26
                    case "":
                        fmt.Println(v15, fi22)
                    default:
                        v30 := []int{fi18, 128, (fi18 ^ fi18)}
                        _ = v30
                    }
                case "":
                    var v31 uint8 = 78
                    _ = v31
                }
                var v32 int16 = 42
                _ = v32
            }
            var v33 uint64 = 0
            _ = fi22
            _ = v33
        }
        _ = fi18
        _ = v21
    }
    _ = v15
    _ = v16
    _ = v17
    return 42
}

func main() {
    var v1 uint8 = 189
    var v2 uint = 4294967295
    var v3 float64 = (1.5 * 0.1)
    v4 := make(map[int]uint32)
    var v5 int64 = 681
    var v6 uint16 = 255
    if false {
        v34 := GenS3{F0: (^(v1)), F1: v6, F2: -9223372036854775808}
        _ = v34
    }
    if (v6 != v6) {
        v35 := []int{127}
        _ = v35
    } else {
        var v36 uint = v2
        v6 -= 1
        _ = v36
    }
    for fi37 := 0; fi37 < 20; fi37++ {
        var v38 uint16 = 1
        _ = fi37
        _ = v38
    }
    var v39 uint8 = (^(v1))
    fmt.Println(v39, v3)
    for fi40 := 0; fi40 < 8; fi40++ {
        for fi41 := 0; fi41 < 9; fi41++ {
            v42 := GenS1{F0: 65535, F1: 0, F2: 341}
            v43 := GenS1{F0: 2147483647, F1: -32768, F2: 32767}
            _ = fi41
            _ = v42
            _ = v43
        }
        _ = fi40
    }
    var v44 int8 = 127
    var v45 int64 = (v5 & v5)
    v46 := []int{}
    var v47 float32 = 1e10
    var v48 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
