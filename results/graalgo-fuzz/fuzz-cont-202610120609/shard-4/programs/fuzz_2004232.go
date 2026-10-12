package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS3 struct { F0 int8; F1 int8; F2 int32 }
func genFn0() uint {
    if ((1000 != 0) || true) {
        var v4 string = ("" + "a")
        for fi5 := 0; fi5 < 19; fi5++ {
            if (!false) {
                v4 = v4
                v4 = v4
            } else {
                v4 = ("\n\t" + (v4 + (v4 + v4)))
                if true {
                    var v6 float64 = -2.25
                    if true {
                        v7 := []int{}
                        for fi8 := 0; fi8 < 4; fi8++ {
                            fmt.Println(fi8, fi5, fi8)
                            var v9 uint32 = 2147483647
                            _ = fi8
                            _ = v9
                        }
                        _ = v7
                    } else {
                        v6 = v6
                    }
                    _ = v6
                }
            }
            if true {
                v4 = "\n\t"
                for fi10 := 0; fi10 < 15; fi10++ {
                    v11 := []int{(fi5 & fi5), 9223372036854775807, fi10}
                    _ = fi10
                    _ = v11
                }
            } else {
                var v12 uint8 = 0
                _ = v12
            }
            var v13 float32 = (-2.25 / 1.0)
            _ = fi5
            _ = v13
        }
        switch -280 {
        case -888:
            if ((-1 < 42) && ((float32(1.5) + (1e10 - ((0.0 - 3.14159) * (0.0 + 0.1)))) >= (((0.1 - (0.1 - (1e10 + 1e10))) * 1.5) / 1e10))) {
                for fi14 := 0; fi14 < 6; fi14++ {
                    fmt.Println(fi14, fi14)
                    fmt.Println(fi14)
                    _ = fi14
                }
            }
        case -1:
            v15 := GenS1{F0: ("\n\t" + v4), F1: 65535}
            _ = v15
        default:
            v16 := GenS2{F0: 45, F1: 42, F2: 2147483647}
            _ = v16
        }
        _ = v4
    }
    v17 := GenS2{F0: 127, F1: 65535, F2: 1}
    _ = v17
    return 42
}
func genFn1() int8 {
    v18 := GenS2{F0: 225, F1: 621, F2: 1}
    var v19 string = "\n\t"
    if (!false) {
        switch -2147483648 {
        case -274:
            v19 = "a"
        case 935:
            switch 2147483647 {
            case 1:
                var v20 bool = (!false)
                _ = v20
            default:
                var v21 uint32 = 420
                _ = v21
            }
        default:
            fmt.Println(v19)
        }
        for fi22 := 0; fi22 < 10; fi22++ {
            fmt.Println(fi22, fi22)
            v23 := GenS3{F0: 1, F1: -128, F2: -1}
            fmt.Println(v19, v19)
            _ = fi22
            _ = v23
        }
    } else {
        if false {
            var v24 int8 = 83
            _ = v24
        }
        v25 := GenS1{F0: (v19 + "foo bar"), F1: 799}
        _ = v25
    }
    v26 := []int{4294967295, 600, 2147483647}
    _ = v18
    _ = v19
    _ = v26
    return 127
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    var v27 uint8 = v1
    v1 = v27
    v1 += v1
    v27 -= v1
    v1 = 177
    for fi28 := 0; fi28 < 11; fi28++ {
        switch fi28 {
        case 2147483647:
            v29 := []int{(fi28 & (^(fi28))), (fi28 << ((2) & 63)), (-(fi28)), (fi28 + fi28)}
            _ = v29
        }
        _ = fi28
    }
    delete(v2, 2147483648)
    var v30 uint64 = 1
    v27 -= (v1 & (v27 & (^(v27))))
    if (float64(((-2.25 + 0.1) * 1e10)) == 0.1) {
        var v31 uint64 = v30
        v1 = (^(v27))
        for fi32 := 0; fi32 < 19; fi32++ {
            switch -115 {
            case 42:
                switch ((("" + "hello") + "foo bar") + "foo bar") {
                case "foo bar":
                    var v33 int = (fi32 | (fi32 - (fi32 | fi32)))
                    _ = v33
                }
            case 913:
                v34 := GenS3{F0: -128, F1: genFn1(), F2: 1}
                _ = v34
            default:
                for fi35 := 0; fi35 < 10; fi35++ {
                    v31 += v31
                    v27 = 253
                    _ = fi35
                }
            }
            _ = fi32
        }
        _ = v31
    } else {
        v30 -= v30
    }
    v36 := v2[2147483648]
    if (false || true) {
        v37 := v3[v36]
        v38 := []int{829, 4294967295, v36, 4294967295}
        _ = v37
        _ = v38
    }
    v39 := []int{128, v36}
    if len(v39) > 0 {
        v39[0] = (v36 & v36)
    }
    var v40 int64 = 9223372036854775807
    var v41 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v36
    _ = v39
    _ = v40
    _ = v41
}
