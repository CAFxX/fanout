package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint }
type GenS2 struct { F0 int8; F1 bool; F2 bool }
type GenS3 struct { F0 uint16; F1 uint16; F2 string }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 uint64; F1 int8; F2 uint32 }
func genFn0() int {
    var v6 int = -401
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: "\n\t", F1: 42, F2: 65535}
        _ = fi7
        _ = v8
    }
    var v9 float32 = ((1.5 - -2.25) / 1.5)
    v6 -= (v6 * (v6 + v6))
    for fi10 := 0; fi10 < 17; fi10++ {
        switch "a" {
        case "\n\t":
            if false {
                for fi11 := 0; fi11 < 7; fi11++ {
                    v12 := []int{(fi10 & fi11), (v6 * fi10), v6}
                    if ((fi10 | (fi11 | (fi11 / (((v6 * fi10)) | 1)))) == (v6 << ((0) & 63))) {
                        if (0.1 <= 0.0) {
                            v9 = v9
                        } else {
                            switch (("hello" + ("\n\t" + ("\n\t" + ("foo bar" + "foo bar")))) + "") {
                            case "":
                                var v13 uint = 42
                                _ = v13
                            case "foo bar":
                                var v14 int8 = 100
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                v15 := GenS2{F0: v14, F1: true, F2: false}
                                fmt.Println(v6)
                                v16 := GenS1{F0: (("\n\t" + "hello") + ("hello" + ("foo bar" + ""))), F1: 18446744073709551615, F2: 2}
                                var v17 string = (((("\n\t" + "\n\t") + ("\n\t" + "foo bar")) + (("" + "\n\t") + "")) + "hello")
                                _ = v14
                                _ = v15
                                _ = v16
                                _ = v17
                            case "a":
                                switch "" {
                                case "\n\t":
                                    if true {
                                        var v18 string = ""
                                        _ = v18
                                    } else {
                                        if false {
                                            v9 = v9
                                            v19 := GenS1{F0: "foo bar", F1: 778, F2: 255}
                                            if true {
                                                var v20 uint8 = 191
                                                _ = v20
                                            }
                                            _ = v19
                                        } else {
                                            v9 += 0.1
                                        }
                                        if false {
                                            v9 = (3.14159 + (float32(0.0) - float32(3.14159)))
                                            var v21 float64 = float64(float32((1.5 - -2.25)))
                                            var v22 uint8 = 28
                                            _ = v21
                                            _ = v22
                                        }
                                        fmt.Println(v9, fi10, v6)
                                        switch "hello" {
                                        case "hello":
                                            var v23 int16 = 32767
                                            _ = v23
                                        }
                                        var v24 uint32 = 2147483647
                                        _ = v24
                                    }
                                }
                            }
                            var v25 float64 = float64(0.0)
                            var v26 uint16 = 389
                            _ = v25
                            _ = v26
                        }
                    } else {
                        if (-2147483648 >= 1) {
                            var v27 string = ((("foo bar" + (("\n\t" + "a") + (((("foo bar" + "\n\t") + "") + "a") + "foo bar"))) + "a") + "a")
                            var v28 int16 = 0
                            _ = len(v12)
                            _ = v27
                            _ = v28
                        } else {
                            v29 := GenS1{F0: "hello", F1: 807, F2: 786}
                            _ = v29
                        }
                        v9 -= float32((0.1 - 3.14159))
                        var v30 int = 127
                        _ = v30
                    }
                    _ = fi11
                    _ = v12
                }
                fmt.Println(fi10, v6, v6)
            } else {
                for fi31 := 0; fi31 < 5; fi31++ {
                    switch v6 {
                    case 32768:
                        v32 := GenS2{F0: 36, F1: true, F2: false}
                        _ = v32
                    }
                    _ = fi31
                }
                v6 = fi10
                if true {
                    for fi33 := 0; fi33 < 2; fi33++ {
                        v34 := []int{fi33}
                        _ = fi33
                        _ = v34
                    }
                    v35 := GenS1{F0: "a", F1: 709, F2: 42}
                    var v36 int8 = int8(v6)
                    _ = v35
                    _ = v36
                }
            }
        case "hello":
            switch "foo bar" {
            case "foo bar":
                v6 = v6
            case "hello":
                var v37 uint64 = 424
                _ = v37
            default:
                if ((v9 - 1e10) > v9) {
                    v9 += v9
                    v38 := GenS1{F0: "hello", F1: 4294967295, F2: 18446744073709551615}
                    var v39 uint64 = 42
                    _ = v38
                    _ = v39
                }
            }
        case "foo bar":
            v6 += (-(v6))
        default:
            v6 = v6
        }
        _ = fi10
    }
    _ = v6
    _ = v9
    return v6
}
func genFn1(p0 string, p1 uint8, p2 bool) uint32 {
    for fi40 := 0; fi40 < 14; fi40++ {
        v41 := GenS3{F0: 42, F1: 92, F2: (((("a" + p0) + ((("hello" + "hello") + (((p0 + ((p0 + p0) + ((p0 + p0) + p0))) + p0) + p0)) + "hello")) + ("" + ("foo bar" + "\n\t"))) + p0)}
        v42 := []int{fi40, (fi40 << ((814) & 63)), genFn0(), (fi40 >> ((713) & 63))}
        _ = fi40
        _ = v41
        _ = v42
    }
    v43 := GenS3{F0: 255, F1: 0, F2: "a"}
    var v44 int16 = -32768
    switch p0 {
    case "":
        switch "" {
        case "a":
            v44 = v44
        case "\n\t":
            v45 := GenS4{F0: 583, F1: 351}
            _ = v45
        case "hello":
            var v46 uint64 = 558
            _ = v46
        default:
            var v47 uint16 = 65535
            _ = v47
        }
    case "\n\t":
        v48 := []int{300, 357, 32768, -2147483648}
        _ = v48
    case "hello":
        p1 = (p1 ^ p1)
    default:
        var v49 uint = 885
        _ = v49
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v43
    _ = v44
    return 4294967295
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 float32 = float32(0.1)
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]int)
    var v50 int = 695
    fmt.Println(v50)
    var v51 uint64 = 866
    v52 := GenS5{F0: 2147483647, F1: 127}
    var v53 int8 = 0
    var v54 int64 = -673
    var v55 float32 = v3
    if false {
        var v56 uint = 16
        for fi57 := 0; fi57 < 4; fi57++ {
            v58 := GenS6{F0: (v51 << ((v56) & 63)), F1: v53, F2: 2147483648}
            var v59 float64 = 0.1
            v3 = v3
            _ = fi57
            _ = v58
            _ = v59
        }
        var v60 float64 = 0.1
        _ = v56
        _ = v60
    }
    delete(v5, (("a" + "foo bar") + "a"))
    var v61 uint32 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v61
}
