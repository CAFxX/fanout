package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
func genFn0() uint8 {
    if (!false) {
        var v5 int32 = -724
        if (-325 < v5) {
            if (true || (149 != 127)) {
                if (2147483647 >= 2) {
                    var v6 float32 = -2.25
                    _ = v6
                }
                var v7 int32 = (v5 / ((v5) | 1))
                _ = v7
            }
            var v8 float64 = 0.1
            _ = v8
        }
        _ = v5
    } else {
        v9 := GenS1{F0: 76, F1: 255, F2: 4294967295}
        v10 := []int{}
        var v11 uint8 = 127
        _ = v10
        _ = v11
        _ = v9
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 uint16 = 165
        v14 := []int{342}
        var v15 int16 = -91
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 float64 = -2.25
    if ((true || true) && (1 != -236)) {
        var v17 string = ("foo bar" + "a")
        _ = v17
    }
    _ = v16
    return 128
}
func genFn1(p0 uint64, p1 int, p2 uint8) string {
    switch "a" {
    case "foo bar":
        p0 += p0
    }
    var v18 uint8 = (^(p2))
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return "hello"
}

func main() {
    var v1 float32 = 0.0
    var v2 uint8 = 42
    v3 := []int{-2147483648, 1}
    var v4 uint = 4294967295
    for fi19 := 0; fi19 < 3; fi19++ {
        fmt.Println(v2, v2)
        v3 = append(v3, fi19)
        v2 = genFn0()
        _ = fi19
    }
    v2 = (^(v2))
    fmt.Println(v1)
    switch -1 {
    case 1:
        switch "foo bar" {
        case "":
            v20 := GenS2{F0: 255}
            _ = v20
        case "foo bar":
            v2 += (v2 << ((326) & 63))
        case "a":
            if true {
                v1 += ((v1 + -2.25) - float32((3.14159 / 1.5)))
                if true {
                    for fi21 := 0; fi21 < 11; fi21++ {
                        v4 += v4
                        v1 = ((v1 + v1) * v1)
                        v2 -= 0
                        _ = fi21
                    }
                    v1 = (float32(0.1) - v1)
                }
                if len(v3) > 0 {
                    v3[0] = 808
                }
            }
        }
    }
    switch 42 {
    case 2147483647:
        v4 = v4
    case 42:
        for fi22 := 0; fi22 < 2; fi22++ {
            for fi23 := 0; fi23 < 14; fi23++ {
                v4 -= v4
                if false {
                    v1 = ((v1 - float32(0.0)) * float32(1.5))
                    if len(v3) > 0 {
                        v3[0] = fi23
                    }
                    if ((3.14159 / (v1 * v1)) < (1e10 - float32(float64((((v1 / v1) * (0.1 / 1.5)) + v1))))) {
                        if len(v3) > 0 {
                            v3[0] = fi23
                        }
                        fmt.Println(fi22, v2)
                        switch fi22 {
                        case 4294967295:
                            var v24 uint16 = 558
                            _ = v24
                        case 32:
                            var v25 uint = v4
                            _ = v25
                        }
                        var v26 uint8 = (v2 & 42)
                        var v27 int8 = -1
                        _ = v26
                        _ = v27
                    }
                }
                if true {
                    v28 := GenS1{F0: 1, F1: uint(v2), F2: v4}
                    fmt.Println(v1, v4)
                    _ = v28
                }
                _ = fi23
            }
            switch ("hello" + "foo bar") {
            case "hello":
                v1 = (float32((1.5 + 1e10)) - 3.14159)
            case "foo bar":
                var v29 int8 = 1
                _ = v29
            case "\n\t":
                if (false || (1 != 581)) {
                    if (303 <= 42) {
                        v1 -= 0.1
                        for fi30 := 0; fi30 < 12; fi30++ {
                            v31 := GenS2{F0: v2}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi30
                            _ = v31
                        }
                        var v32 uint = 1
                        _ = v32
                    } else {
                        for fi33 := 0; fi33 < 13; fi33++ {
                            var v34 float32 = ((0.0 + 1e10) + float32(0.1))
                            v3 = append(v3, (fi33 * fi33))
                            var v35 bool = ((fi33 - fi22) != 9223372036854775807)
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                        var v36 uint32 = 749
                        var v37 float64 = (3.14159 - 0.1)
                        _ = v36
                        _ = v37
                    }
                    var v38 float64 = float64(-2.25)
                    _ = v38
                }
            }
            v2 = (^(v2))
            _ = fi22
        }
    }
    var v39 string = "a"
    switch 32767 {
    case 42:
        var v40 int16 = -32768
        _ = v40
    case -782:
        var v41 uint = (v4 ^ 1)
        _ = v41
    case 127:
        v4 -= v4
    }
    var v42 int = 857
    v1 = v1
    var v43 uint32 = 42
    for fi44 := 0; fi44 < 14; fi44++ {
        v43 += v43
        v45 := GenS1{F0: (v2 | genFn0()), F1: (v4 << ((v4) & 63)), F2: (v4 ^ 1)}
        _ = fi44
        _ = v45
    }
    switch (v39 + v39) {
    case "foo bar":
        var v46 int64 = 964
        _ = v46
    default:
        v47 := GenS3{F0: 127}
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v42
    _ = v43
}
