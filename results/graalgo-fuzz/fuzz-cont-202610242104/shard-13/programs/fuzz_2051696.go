package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint8; F2 int64 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 int16; F1 int32; F2 bool }
type GenS7 struct { F0 uint16; F1 int64 }
func genFn0(p0 int8) int8 {
    var v4 uint = 1
    fmt.Println(v4, p0)
    v5 := GenS1{F0: ("\n\t" + ("foo bar" + "\n\t"))}
    _ = p0
    _ = v4
    _ = v5
    return 126
}
func genFn1(p0 uint64) uint64 {
    if (!true) {
        var v6 int64 = 9223372036854775807
        var v7 uint8 = 128
        v8 := GenS1{F0: ("\n\t" + "\n\t")}
        _ = v6
        _ = v7
        _ = v8
    }
    if (1.5 <= 1e10) {
        fmt.Println(p0)
    } else {
        p0 -= (^(p0))
        if false {
            v9 := []int{}
            v10 := GenS1{F0: "\n\t"}
            _ = v10
            _ = v9
        }
    }
    _ = p0
    return (p0 - (p0 << ((65535) & 63)))
}

func main() {
    var v1 int = 42
    v2 := make(map[string]bool)
    var v3 uint32 = 2
    switch (-(v1)) {
    case 128:
        fmt.Println(v3, v3)
    case -1:
        var v11 int32 = 1
        _ = v11
    case 0:
        var v12 float32 = 1.5
        _ = v12
    default:
        fmt.Println(v3, v3)
    }
    var v13 float32 = 1.5
    v14 := []int{(v1 >> ((0) & 63)), v1, (v1 / ((v1) | 1))}
    fmt.Println(v13)
    switch (v1 | v1) {
    case 178:
        var v15 float32 = float32(0.0)
        _ = v15
    default:
        v16 := GenS2{F0: 42, F1: 131, F2: -9223372036854775808}
        _ = v16
    }
    v17 := GenS2{F0: 100, F1: 128, F2: 42}
    var v18 bool = true
    fmt.Println(v13)
    _ = len(v14)
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := GenS2{F0: genFn0(127), F1: 0, F2: -512}
        if ((-3 >= 100) && v18) {
            if false {
                v21 := GenS3{F0: (-(v1)), F1: 69}
                var v22 int64 = -1
                switch fi19 {
                case 819:
                    v3 -= (^(v3))
                }
                _ = v21
                _ = v22
            }
        }
        v23 := GenS4{F0: v3, F1: "\n\t"}
        _ = fi19
        _ = v20
        _ = v23
    }
    if (289 <= 9223372036854775807) {
        switch ("a" + (("foo bar" + (("hello" + ("" + "foo bar")) + "")) + ("a" + ("" + ("\n\t" + ("" + ("\n\t" + "a"))))))) {
        case "foo bar":
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 bool = (!true)
                var v26 int8 = 0
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        var v27 int16 = -761
        for fi28 := 0; fi28 < 20; fi28++ {
            if v18 {
                for fi29 := 0; fi29 < 9; fi29++ {
                    var v30 uint64 = 18446744073709551615
                    v14 = append(v14, (v1 & fi28))
                    _ = fi29
                    _ = v30
                }
                var v31 int16 = 0
                _ = v31
            } else {
                v32 := GenS5{F0: 4294967295, F1: 2}
                v13 -= v13
                _ = v32
            }
            _ = fi28
        }
        _ = v27
    } else {
        var v33 float64 = ((0.1 * 0.1) - 0.0)
        _ = v33
    }
    var v34 float64 = 1e10
    for fi35 := 0; fi35 < 14; fi35++ {
        if false {
            fmt.Println(fi35)
            v1 -= fi35
            _ = len(v2)
        }
        for fi36 := 0; fi36 < 7; fi36++ {
            if false {
                var v37 uint32 = (v3 * (v3 << ((2) & 63)))
                switch fi35 {
                case -2147483648:
                    if true {
                        fmt.Println(v34, v18)
                    } else {
                        v38 := v2["\n\t"]
                        switch "a" {
                        case "foo bar":
                            v39 := GenS6{F0: -916, F1: 0, F2: v38}
                            _ = v39
                        default:
                            if true {
                                if (9 >= 87) {
                                    var v40 int8 = 1
                                    var v41 int64 = -981
                                    _ = v40
                                    _ = v41
                                }
                                v18 = (v38 && false)
                            } else {
                                var v42 uint8 = uint8(v37)
                                _ = v42
                            }
                        }
                        _ = v38
                    }
                case -9223372036854775808:
                    fmt.Println(v13, v18, v37)
                case -642:
                    v43 := GenS7{F0: 255, F1: -952}
                    _ = v43
                default:
                    fmt.Println(v13)
                }
                _ = v37
            }
            _ = fi36
        }
        fmt.Println(v3, v13)
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v34
}
