package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint64 }
func genFn0() uint {
    var v7 bool = false
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int = (-(fi8))
        var v10 float64 = 0.1
        fmt.Println(fi8)
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 0}
    v7 = (42 == 547)
    _ = v11
    _ = v7
    return 255
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    v4 := make(map[int]bool)
    var v5 float32 = (1.5 + (1e10 + 1e10))
    v6 := make(map[uint32]int)
    var v12 uint8 = 1
    for fi13 := 0; fi13 < 11; fi13++ {
        switch fi13 {
        case 2147483647:
            var v14 uint = 255
            _ = v14
        default:
            v15 := GenS1{F0: -128}
            _ = v15
        }
        v16 := []int{}
        _ = fi13
        _ = v16
    }
    var v17 string = ("hello" + (("hello" + "") + ((("hello" + "foo bar") + "foo bar") + ("\n\t" + ("foo bar" + "a")))))
    v12 += v12
    var v18 int8 = 0
    fmt.Println(v17, v17, v5)
    if true {
        delete(v1, 4294967295)
        var v19 bool = ((float32(3.14159) == (float32(1.5) * (v5 * v5))) || (!true))
        _ = v19
    } else {
        if true {
            switch -396 {
            case 42:
                v5 += (v5 * v5)
            case -9223372036854775808:
                if ((^(v12)) < (v12 * v12)) {
                    var v20 uint = 47
                    for fi21 := 0; fi21 < 1; fi21++ {
                        v22 := GenS2{F0: 283, F1: 2}
                        _ = fi21
                        _ = v22
                    }
                    var v23 float64 = (0.1 + 0.0)
                    _ = v20
                    _ = v23
                } else {
                    delete(v4, 5)
                }
            default:
                var v24 int = 9223372036854775807
                _ = v24
            }
            var v25 float32 = 0.0
            _ = v25
        } else {
            v12 -= (^(v12))
        }
        for fi26 := 0; fi26 < 14; fi26++ {
            var v27 int32 = -1
            fmt.Println(fi26)
            var v28 uint64 = 1
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    if true {
        v5 = float32((float64(v5) / 1e10))
    }
    v5 -= (((v5 / 1e10) - v5) / (v5 - ((v5 / (((float32(-2.25) * v5) / v5) - ((3.14159 * 0.0) + v5))) * (-2.25 + v5))))
    if false {
        for fi29 := 0; fi29 < 3; fi29++ {
            var v30 uint8 = v12
            _ = fi29
            _ = v30
        }
        v31 := GenS1{F0: (v18 >> ((genFn0()) & 63))}
        _ = v31
    } else {
        for fi32 := 0; fi32 < 5; fi32++ {
            switch 1 {
            case 742:
                for fi33 := 0; fi33 < 13; fi33++ {
                    v34 := []int{fi33, fi33}
                    switch (fi33 % ((fi33) | 1)) {
                    case 108:
                        v12 = v12
                    default:
                        switch (v17 + (((v17 + "hello") + (v17 + "foo bar")) + ("" + v17))) {
                        case "":
                            var v35 uint32 = uint32(v12)
                            _ = v35
                        }
                    }
                    _ = fi33
                    _ = v34
                }
            case -2147483648:
                v36 := GenS2{F0: uint16(v12), F1: 380}
                _ = v36
            case -9223372036854775808:
                var v37 int32 = -1
                _ = v37
            default:
                var v38 uint8 = v12
                _ = v38
            }
            switch fi32 {
            case 293:
                v5 = v5
            }
            _ = fi32
        }
    }
    var v39 string = v17
    for fi40 := 0; fi40 < 12; fi40++ {
        fmt.Println(fi40, v12, v17)
        _ = fi40
    }
    v41 := []int{4294967295, 4294967295, -2147483648}
    switch -287 {
    case 127:
        var v42 int = 127
        _ = v42
    case 416:
        v18 += v18
    }
    if ((int32(v12) != 0) || false) {
        switch int(v18) {
        case 128:
            for fi43 := 0; fi43 < 9; fi43++ {
                var v44 float32 = 3.14159
                _ = fi43
                _ = v44
            }
        case -2147483648:
            fmt.Println(v5, v39)
        default:
            v45 := GenS2{F0: 42, F1: 425}
            _ = v45
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
