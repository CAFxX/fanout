package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64; F1 bool; F2 uint }
func genFn0() uint32 {
    v4 := []int{9223372036854775807, 664}
    for fi5 := 0; fi5 < 9; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 int8 = 100
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 17; fi8++ {
            fmt.Println(fi5)
            var v9 bool = (42 >= -64)
            switch "hello" {
            case "a":
                v10 := GenS1{F0: 804}
                _ = v10
            case "":
                v9 = v9
            case "foo bar":
                var v11 int16 = 42
                _ = v11
            }
            _ = fi8
            _ = v9
        }
        _ = fi5
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        fmt.Println(fi12)
        for fi13 := 0; fi13 < 10; fi13++ {
            switch (fi12 & 32767) {
            case 32767:
                var v14 uint64 = 144
                _ = v14
            case -538:
                var v15 uint = 225
                _ = v15
            case 0:
                v16 := GenS2{F0: 2, F1: 1}
                _ = v16
            default:
                v17 := GenS3{F0: -128, F1: true}
                _ = v17
            }
            v18 := GenS4{F0: 42}
            if true {
                switch "a" {
                case "":
                    _ = len(v4)
                }
                if true {
                    var v19 float64 = (1e10 / (0.1 * 3.14159))
                    _ = len(v4)
                    _ = v19
                }
            }
            _ = fi13
            _ = v18
        }
        if (float64(3.14159) != (-2.25 - -2.25)) {
            if false {
                var v20 float32 = 1e10
                _ = v20
            } else {
                var v21 uint64 = 4294967295
                _ = v21
            }
            v22 := GenS5{F0: -832, F1: true, F2: 18446744073709551615}
            v4 = append(v4, fi12)
            _ = v22
        } else {
            var v23 uint64 = 18446744073709551615
            _ = v23
        }
        _ = fi12
    }
    var v24 int16 = 42
    _ = v24
    _ = v4
    return 46
}
func genFn1() int64 {
    for fi25 := 0; fi25 < 19; fi25++ {
        v26 := GenS2{F0: 31, F1: -128}
        if (true && false) {
            for fi27 := 0; fi27 < 15; fi27++ {
                var v28 uint16 = 255
                v28 = v28
                fmt.Println(v28, fi27, fi27)
                switch "a" {
                case "hello":
                    v28 -= v28
                case "\n\t":
                    if ((v28 < (v28 & v28)) && false) {
                        var v29 uint32 = 0
                        _ = v29
                    }
                    var v30 uint8 = 42
                    _ = v30
                case "":
                    var v31 uint8 = 214
                    _ = v31
                default:
                    v28 -= (v28 << ((899) & 63))
                }
                fmt.Println(fi27)
                v32 := []int{fi27}
                _ = fi27
                _ = v28
                _ = v32
            }
        } else {
            var v33 string = ((("hello" + "") + "") + "\n\t")
            var v34 uint64 = 42
            for fi35 := 0; fi35 < 7; fi35++ {
                fmt.Println(fi25, fi35)
                v36 := []int{(fi25 & (^(fi25))), fi35}
                v36 = append(v36, fi25)
                _ = fi35
                _ = v36
            }
            _ = v33
            _ = v34
        }
        var v37 int64 = 499
        _ = fi25
        _ = v26
        _ = v37
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        if false {
            var v39 int64 = 1
            _ = v39
        }
        _ = fi38
    }
    return -519
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    var v3 float64 = 3.14159
    var v40 string = "\n\t"
    v3 = float64(1e10)
    var v41 bool = ("\n\t" < v40)
    var v42 uint = 1
    v43 := GenS5{F0: -889, F1: (v41 || (!v41)), F2: v42}
    for fi44 := 0; fi44 < 19; fi44++ {
        v42 = v42
        for fi45 := 0; fi45 < 4; fi45++ {
            v40 = "\n\t"
            v46 := []int{-632, (fi44 | 1)}
            _ = fi45
            _ = v46
        }
        v3 -= ((v3 / v3) + 3.14159)
        _ = fi44
    }
    fmt.Println(v40)
    var v47 int64 = -9223372036854775808
    v47 = (^(v47))
    fmt.Println(v42)
    fmt.Println(v3)
    var v48 float32 = (0.1 - (3.14159 + (((3.14159 + 0.1) * 1.5) / 1e10)))
    v48 = float32((1.5 * ((0.1 - v3) / (3.14159 + ((v3 / (((v3 / 1e10) + ((v3 * v3) - 1e10)) - v3)) * 0.0)))))
    v49 := GenS4{F0: 32767}
    fmt.Println(v48, v41, v47)
    var v50 uint16 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
