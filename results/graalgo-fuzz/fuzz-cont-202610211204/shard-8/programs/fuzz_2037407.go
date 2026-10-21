package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int8; F1 int64; F2 uint }
func genFn0(p0 uint16) int {
    var v7 int8 = 1
    var v8 int32 = -305
    v9 := []int{2147483647, 0}
    if (!(!false)) {
        _ = len(v9)
    } else {
        if (0 < 1) {
            switch ((("foo bar" + "foo bar") + "\n\t") + ("" + "")) {
            case "hello":
                switch -9223372036854775808 {
                case -244:
                    if ((p0 / ((p0) | 1)) <= (p0 / (((p0 >> ((255) & 63))) | 1))) {
                        v8 = (v8 + v8)
                        for fi10 := 0; fi10 < 4; fi10++ {
                            var v11 float64 = (-2.25 - 0.1)
                            _ = fi10
                            _ = v11
                        }
                    } else {
                        if ("a" == ("hello" + "hello")) {
                            var v12 float32 = float32(0.1)
                            p0 -= (^(p0))
                            switch 2147483647 {
                            case 125:
                                var v13 int32 = 2147483647
                                _ = v13
                            default:
                                var v14 uint = 0
                                _ = v14
                            }
                            _ = v12
                        }
                    }
                case 1:
                    var v15 int32 = -2147483648
                    _ = v15
                }
            case "":
                fmt.Println(p0, v7, v8)
            }
        }
    }
    p0 -= (^(p0))
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 32767
}
func genFn1() uint16 {
    if (!false) {
        switch -924 {
        case 127:
            if true {
                if true {
                    if (127 != 127) {
                        v16 := []int{-236}
                        var v17 uint8 = 255
                        _ = v16
                        _ = v17
                    } else {
                        v18 := GenS1{F0: ("\n\t" + "\n\t")}
                        switch "\n\t" {
                        case "":
                            for fi19 := 0; fi19 < 3; fi19++ {
                                var v20 float32 = ((0.0 - -2.25) + 1.5)
                                var v21 string = ("a" + "foo bar")
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                        case "hello":
                            v22 := GenS1{F0: "hello"}
                            _ = v22
                        case "a":
                            for fi23 := 0; fi23 < 10; fi23++ {
                                if (2147483647 <= 0) {
                                    if true {
                                        var v24 int32 = -2147483648
                                        v24 = v24
                                        _ = v24
                                    }
                                    v25 := GenS2{F0: (0 > 2147483647), F1: 249}
                                    v26 := GenS3{F0: 1, F1: 218}
                                    _ = v25
                                    _ = v26
                                } else {
                                    var v27 int64 = -9223372036854775808
                                    _ = v27
                                }
                                for fi28 := 0; fi28 < 2; fi28++ {
                                    var v29 float32 = 1e10
                                    var v30 string = ("a" + "a")
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                                _ = fi23
                            }
                        }
                        v31 := GenS3{F0: 65535, F1: 119}
                        _ = v18
                        _ = v31
                    }
                } else {
                    if (false && (4294967295 < 42)) {
                        for fi32 := 0; fi32 < 9; fi32++ {
                            var v33 string = ("foo bar" + "a")
                            if false {
                                var v34 float32 = -2.25
                                fmt.Println(v34)
                                _ = v34
                            }
                            _ = fi32
                            _ = v33
                        }
                    }
                }
                var v35 float64 = float64((0.1 - 1e10))
                v35 = v35
                _ = v35
            }
        case 2147483647:
            var v36 string = ("\n\t" + ("foo bar" + "hello"))
            _ = v36
        default:
            v37 := []int{}
            _ = v37
        }
        var v38 int16 = -32768
        if false {
            v39 := GenS4{F0: 42, F1: 48, F2: 474}
            _ = v39
        }
        _ = v38
    }
    var v40 uint8 = 2
    switch "hello" {
    case "hello":
        if true {
            v41 := GenS3{F0: 590, F1: 65535}
            v42 := []int{32768}
            _ = v41
            _ = v42
        }
    default:
        if true {
            v40 += v40
            var v43 uint16 = 42
            v43 += (v43 - (v43 * v43))
            _ = v43
        }
    }
    _ = v40
    return 32767
}

func main() {
    var v1 int32 = -833
    var v2 float32 = 1e10
    v3 := make(map[int]int)
    var v4 string = ("foo bar" + "foo bar")
    var v5 bool = (v4 >= v4)
    v6 := []string{(("hello" + "") + "a"), v4, "hello"}
    var v44 uint16 = 1
    v44 -= genFn1()
    if (!(v2 >= v2)) {
        var v45 uint = 1
        _ = v45
    } else {
        v4 = ((("" + v4) + (v4 + "hello")) + v4)
        if (true || (v2 == (v2 / (v2 - float32(((-2.25 * 1.5) * (1e10 + (1.5 / (float64(1e10) / 1.0))))))))) {
            v46 := GenS3{F0: 65535, F1: 2}
            _ = v46
        } else {
            v3[42] = 42
            v4 = ("\n\t" + v4)
            v2 = v2
        }
    }
    var v47 uint32 = 65535
    switch "" {
    case "":
        v48 := GenS3{F0: genFn1(), F1: 65535}
        _ = v48
    }
    v6 = append(v6, v4)
    v2 -= v2
    v49 := GenS1{F0: (v4 + v4)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v47
    _ = v49
    _ = v5
    _ = v6
}
