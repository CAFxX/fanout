package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int16; F2 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 uint8; F2 int8 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 uint16 }
func genFn0(p0 int32, p1 uint8) uint32 {
    v7 := GenS1{F0: 127, F1: 42}
    var v8 int8 = 127
    if true {
        v8 -= (^(v8))
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 894
}
func genFn1(p0 int8, p1 uint16, p2 int32) int16 {
    v9 := []int{-988, 499, 618, -2147483648}
    var v10 string = "a"
    p0 -= p0
    v11 := GenS2{F0: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return -160
}

func main() {
    var v1 int8 = -128
    v2 := make(map[int]int)
    v3 := []uint32{42, 42, 2, 2}
    var v4 int8 = 127
    v5 := []int{741}
    var v6 uint32 = 4294967295
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS3{F0: (v1 % ((42) | 1)), F1: 0, F2: 42}
        for fi14 := 0; fi14 < 7; fi14++ {
            fmt.Println(v4, v1)
            v6 = 2147483648
            _ = len(v3)
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    var v15 int = -1
    var v16 int32 = -70
    var v17 bool = ((0.1 / 0.1) == float64(0.1))
    if ((^(v16)) != -1) {
        if ("hello" < "a") {
            v18 := v2[-674]
            _ = v18
        }
        switch "\n\t" {
        case "":
            for fi19 := 0; fi19 < 2; fi19++ {
                if (genFn0(v16, 227) <= 762) {
                    v1 -= -128
                    if (float32(-2.25) > -2.25) {
                        v20 := GenS4{F0: v6}
                        _ = v20
                    } else {
                        var v21 bool = ("foo bar" >= "\n\t")
                        fmt.Println(v4)
                        v4 -= (v1 & (v1 | -106))
                        _ = v21
                    }
                } else {
                    var v22 uint32 = v6
                    _ = v22
                }
                switch (fi19 ^ -1) {
                case 334:
                    _ = len(v3)
                case -2147483648:
                    v23 := GenS5{F0: fi19, F1: 29, F2: (v1 * (-(v4)))}
                    _ = v23
                case 32768:
                    if false {
                        v2[42] = 806
                        for fi24 := 0; fi24 < 5; fi24++ {
                            var v25 string = "\n\t"
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi24
                            _ = v25
                        }
                    }
                default:
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v6, v16)
                    }
                }
                _ = fi19
            }
        case "\n\t":
            fmt.Println(v17, v16)
        }
    }
    if false {
        if len(v3) > 0 {
            v3[0] = (v6 % ((42) | 1))
        }
    }
    var v26 bool = true
    if (true && v17) {
        fmt.Println(v1, v16, v17)
        v16 += v16
        var v27 uint = 226
        _ = v27
    }
    _ = len(v5)
    if (-2147483648 <= (v16 & (-(v16)))) {
        if (!v26) {
            for fi28 := 0; fi28 < 4; fi28++ {
                for fi29 := 0; fi29 < 7; fi29++ {
                    if true {
                        delete(v2, (v15 & v15))
                    } else {
                        var v30 bool = true
                        var v31 int16 = -32768
                        v32 := v2[(-(v15))]
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    _ = fi29
                }
                for fi33 := 0; fi33 < 12; fi33++ {
                    var v34 float64 = -2.25
                    var v35 int16 = 32767
                    switch -620 {
                    case 4294967295:
                        var v36 uint = 1
                        _ = v36
                    case 128:
                        var v37 float64 = (-2.25 + -2.25)
                        _ = v37
                    case 1:
                        v34 -= (v34 * 1e10)
                    default:
                        v16 += v16
                    }
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                v4 = (^(v4))
                _ = fi28
            }
            _ = len(v2)
            switch "foo bar" {
            case "hello":
                var v38 int = (v15 + -434)
                _ = v38
            case "":
                var v39 bool = (!true)
                _ = v39
            case "\n\t":
                if (255 <= 326) {
                    v16 = (v16 - (-(v16)))
                } else {
                    v40 := GenS6{F0: (v17 && false), F1: 65535}
                    fmt.Println(v16, v6, v15)
                    var v41 float32 = -2.25
                    _ = v40
                    _ = v41
                }
            default:
                if (float32(((0.1 - 1.5) / (1.5 * ((1e10 / float64(float32((float64(float32((((3.14159 / 1e10) / 1.0) * 3.14159))) / 1.0)))) + 3.14159)))) == 1e10) {
                    v42 := GenS4{F0: (v6 + (v6 >> ((1) & 63)))}
                    fmt.Println(v16, v6)
                    var v43 int8 = 0
                    _ = v42
                    _ = v43
                } else {
                    var v44 string = "foo bar"
                    _ = v44
                }
            }
        }
    }
    switch (("\n\t" + ((("a" + "hello") + "\n\t") + "")) + "\n\t") {
    case "hello":
        var v45 int32 = 2147483647
        _ = v45
    case "foo bar":
        v46 := GenS6{F0: false, F1: 4294967295}
        _ = v46
    case "":
        v47 := GenS7{F0: 65535}
        _ = v47
    }
    v26 = (v26 && ((v4 * 0) == v4))
    fmt.Println(v17)
    v16 -= 2147483647
    for fi48 := 0; fi48 < 8; fi48++ {
        fmt.Println(v15, v4)
        _ = fi48
    }
    v17 = (true || v17)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
