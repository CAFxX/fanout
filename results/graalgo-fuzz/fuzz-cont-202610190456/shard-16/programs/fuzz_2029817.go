package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int32; F2 int16 }
type GenS3 struct { F0 uint; F1 bool; F2 uint64 }
type GenS4 struct { F0 int16; F1 string; F2 string }
type GenS5 struct { F0 bool; F1 int64; F2 int8 }
func genFn0(p0 string, p1 int, p2 bool) uint8 {
    p2 = true
    fmt.Println(p0, p0)
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 67
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    var v3 uint = 255
    var v4 float64 = (3.14159 + (float64(3.14159) - 0.0))
    var v5 int = -248
    var v6 int = v5
    v7 := GenS1{F0: (^(v3))}
    var v8 int64 = -9223372036854775808
    if false {
        if (685 >= (v3 | v3)) {
            v4 = (v4 / ((0.0 - v4) - (v4 - 0.1)))
            v9 := []int{v6, (v6 / ((v6) | 1)), v6, v5}
            _ = len(v9)
            _ = v9
        } else {
            for fi10 := 0; fi10 < 20; fi10++ {
                switch fi10 {
                case -1:
                    v4 -= v4
                default:
                    var v11 uint = (v3 / ((535) | 1))
                    _ = v11
                }
                v5 += fi10
                _ = fi10
            }
        }
        if false {
            if false {
                v3 -= (v3 * (v3 - (v3 ^ (v3 + (v3 * v3)))))
            } else {
                if (1e10 == 3.14159) {
                    var v12 uint16 = 1
                    for fi13 := 0; fi13 < 7; fi13++ {
                        v3 += (^(v3))
                        var v14 int = fi13
                        var v15 uint = (v3 + v3)
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    _ = v12
                } else {
                    if (0 != 128) {
                        switch v5 {
                        case -380:
                            switch (v5 / ((635) | 1)) {
                            case 266:
                                fmt.Println(v8, v4, v8)
                            case 780:
                                fmt.Println(v6, v8)
                            case -335:
                                for fi16 := 0; fi16 < 2; fi16++ {
                                    v4 = v4
                                    v17 := GenS2{F0: 42, F1: -271, F2: -32768}
                                    if (!true) {
                                        v4 = 3.14159
                                        v18 := GenS3{F0: 260, F1: true, F2: 0}
                                        _ = v18
                                    }
                                    _ = fi16
                                    _ = v17
                                }
                            }
                        case 32768:
                            v3 -= v3
                        }
                    } else {
                        fmt.Println(v8, v5)
                    }
                }
                v19 := []int{}
                _ = v19
            }
        } else {
            for fi20 := 0; fi20 < 20; fi20++ {
                v8 += 0
                _ = fi20
            }
            v5 = (v6 << (((v3 * 1)) & 63))
            fmt.Println(v8)
        }
        v21 := GenS4{F0: 32767, F1: (("foo bar" + "\n\t") + "hello"), F2: "hello"}
        _ = v21
    }
    v22 := GenS1{F0: v3}
    v6 += (v6 ^ int(v3))
    if true {
        var v23 int64 = v8
        _ = v23
    } else {
        var v24 uint64 = 42
        var v25 float64 = (1.5 * v4)
        if (((("foo bar" + "") + "hello") + "foo bar") == ("\n\t" + ("a" + "a"))) {
            for fi26 := 0; fi26 < 2; fi26++ {
                v3 += (v3 >> (((v3 | (^(v3)))) & 63))
                for fi27 := 0; fi27 < 18; fi27++ {
                    var v28 float32 = (1e10 / 0.1)
                    v5 = (v6 | (fi26 - v6))
                    _ = fi27
                    _ = v28
                }
                var v29 float32 = 3.14159
                _ = fi26
                _ = v29
            }
            if false {
                v30 := []int{}
                var v31 uint8 = 128
                _ = v30
                _ = v31
            }
        }
        _ = v24
        _ = v25
    }
    v32 := []int{}
    for fi33 := 0; fi33 < 8; fi33++ {
        var v34 int8 = 42
        for fi35 := 0; fi35 < 17; fi35++ {
            if ((float64(1e10) + ((-2.25 / (((v4 + v4) * (v4 * v4)) * v4)) * (v4 - ((0.1 / 1.0) / float64(-2.25))))) > 1.5) {
                v36 := GenS2{F0: int32(v8), F1: int32(v34), F2: -906}
                _ = v36
            } else {
                if len(v32) >= 2 {
                    v32 = v32[1:len(v32)]
                }
                if ((v34 ^ (v34 - v34)) >= -1) {
                    fmt.Println(v4)
                    v34 -= (v34 - (-(v34)))
                    v34 = v34
                }
                v37 := GenS4{F0: 42, F1: "\n\t", F2: (("" + "foo bar") + ("" + ((("hello" + "hello") + "hello") + ("" + "hello"))))}
                _ = v37
            }
            _ = fi35
        }
        _ = fi33
        _ = v34
    }
    var v38 int32 = -443
    v6 = 42
    var v39 bool = false
    _ = len(v1)
    switch (v6 + -586) {
    case 686:
        for fi40 := 0; fi40 < 2; fi40++ {
            if (0 > 255) {
                fmt.Println(fi40)
                fmt.Println(v8, v5, v3)
                v41 := GenS5{F0: v39, F1: v8, F2: -128}
                _ = v41
            }
            v6 = (fi40 + int(v3))
            var v42 int64 = v8
            _ = fi40
            _ = v42
        }
    }
    v43 := GenS1{F0: 42}
    switch (((((((("" + "\n\t") + (("hello" + "\n\t") + "\n\t")) + "hello") + "a") + ("hello" + "hello")) + "hello") + (("hello" + ("" + ("a" + (((("a" + ("foo bar" + (("" + "foo bar") + "a"))) + "hello") + "\n\t") + "\n\t")))) + "\n\t")) + "hello") {
    case "foo bar":
        v44 := GenS5{F0: true, F1: (v8 | (v8 - (v8 ^ (v8 ^ (v8 & (-(v8))))))), F2: 127}
        _ = v44
    }
    if true {
        if true {
            var v45 int = v5
            v46 := GenS1{F0: (v3 | v3)}
            if len(v32) >= 2 {
                v32 = v32[1:len(v32)]
            }
            _ = v45
            _ = v46
        } else {
            var v47 uint = (^(v3))
            for fi48 := 0; fi48 < 7; fi48++ {
                var v49 int8 = -1
                var v50 uint64 = 42
                _ = len(v1)
                _ = fi48
                _ = v49
                _ = v50
            }
            _ = v47
        }
        var v51 float64 = float64(float32(v4))
        _ = v51
    }
    var v52 float32 = float32(((v4 * 1e10) / (v4 - (v4 + v4))))
    fmt.Println(v3, v38, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v5
    _ = v52
    _ = v6
    _ = v7
    _ = v8
}
