package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint64; F1 string; F2 int64 }
type GenS4 struct { F0 int16; F1 int8 }
func genFn0() uint16 {
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := []int{}
        var v7 float32 = -2.25
        if true {
            var v8 int16 = 42
            var v9 bool = (!true)
            if (true && (!(!false))) {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            } else {
                for fi10 := 0; fi10 < 15; fi10++ {
                    var v11 uint64 = 674
                    v12 := GenS1{F0: v8, F1: 125}
                    fmt.Println(fi5, v8, v7)
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                if (!true) {
                    v8 = (v8 + (v8 >> ((369) & 63)))
                    var v13 uint32 = 42
                    for fi14 := 0; fi14 < 2; fi14++ {
                        v15 := GenS2{F0: 100, F1: 0, F2: 617}
                        v13 -= (^(v13))
                        var v16 int64 = -1
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    _ = v13
                }
            }
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    switch -573 {
    case 32768:
        v17 := []int{128, 860}
        _ = v17
    case 484:
        v18 := []int{-791, -485, 1}
        _ = v18
    default:
        var v19 int = 4294967295
        _ = v19
    }
    if true {
        var v20 int32 = -2147483648
        if false {
            v20 -= 840
            var v21 string = "hello"
            _ = v21
        }
        if false {
            v20 = (^(v20))
        }
        _ = v20
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        switch (-(fi22)) {
        case 470:
            var v23 float64 = ((-2.25 / 1.0) * 1e10)
            _ = v23
        }
        var v24 uint8 = 42
        _ = fi22
        _ = v24
    }
    return 769
}
func genFn1(p0 int32, p1 uint, p2 uint8) uint16 {
    var v25 int8 = 0
    for fi26 := 0; fi26 < 18; fi26++ {
        v27 := GenS1{F0: 42, F1: p2}
        _ = fi26
        _ = v27
    }
    v28 := []int{-868, int(v25), 32768}
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v28
    return 2
}

func main() {
    v1 := []string{}
    v2 := make(map[string]bool)
    v3 := []uint32{376, 2147483647, 997, 51}
    var v4 int = 128
    var v29 int64 = 1
    for fi30 := 0; fi30 < 15; fi30++ {
        var v31 uint16 = 2
        _ = fi30
        _ = v31
    }
    switch "foo bar" {
    case "":
        fmt.Println(v29, v29)
    default:
        fmt.Println(v29)
    }
    fmt.Println(v4)
    if ((4294967295 >= 301) && (0.1 != 1e10)) {
        var v32 string = "hello"
        var v33 int64 = (v29 % (((v29 << ((827) & 63))) | 1))
        fmt.Println(v29)
        _ = v32
        _ = v33
    } else {
        v34 := GenS3{F0: 4294967295, F1: "hello", F2: v29}
        for fi35 := 0; fi35 < 19; fi35++ {
            for fi36 := 0; fi36 < 7; fi36++ {
                v4 += fi36
                switch "a" {
                case "\n\t":
                    fmt.Println(fi36, v29)
                default:
                    var v37 float32 = ((float32((0.0 / 1.0)) + -2.25) + 0.1)
                    _ = v37
                }
                if (148 <= 780) {
                    v29 -= v29
                }
                _ = fi36
            }
            _ = fi35
        }
        v38 := GenS2{F0: -1, F1: 65535, F2: genFn0()}
        _ = v34
        _ = v38
    }
    if (genFn1(693, 0, 42) >= 118) {
        v39 := GenS2{F0: 100, F1: 255, F2: genFn1(1, 0, 42)}
        _ = v39
    }
    v4 += v4
    var v40 uint32 = 65535
    var v41 float32 = (3.14159 * 3.14159)
    if (1 == 108) {
        var v42 int32 = 912
        fmt.Println(v4, v40, v4)
        _ = v42
    }
    var v43 uint32 = 2147483647
    var v44 bool = (218 >= -916)
    v40 = (v40 >> ((1) & 63))
    if (v44 || (v43 == (v40 ^ (v40 ^ (v43 ^ 81))))) {
        if v44 {
            if false {
                for fi45 := 0; fi45 < 19; fi45++ {
                    v46 := GenS1{F0: -32768, F1: 251}
                    _ = fi45
                    _ = v46
                }
            }
            fmt.Println(v43)
            switch 32767 {
            case 336:
                var v47 string = ("\n\t" + ("" + ""))
                _ = v47
            case 0:
                v48 := GenS4{F0: -197, F1: 40}
                _ = v48
            default:
                v49 := GenS1{F0: 42, F1: 0}
                _ = v49
            }
        }
    } else {
        var v50 uint32 = 2147483648
        var v51 uint8 = 2
        _ = v50
        _ = v51
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v43
    _ = v44
}
