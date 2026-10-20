package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS3 struct { F0 bool; F1 int64; F2 uint32 }
type GenS4 struct { F0 int32; F1 uint16; F2 uint }
type GenS5 struct { F0 int32 }
func genFn0(p0 int32) int32 {
    var v4 uint64 = 4294967295
    fmt.Println(v4)
    if true {
        var v5 float64 = -2.25
        v4 = v4
        _ = v5
    } else {
        v4 = (v4 ^ 745)
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := GenS1{F0: -1}
            _ = fi6
            _ = v7
        }
    }
    if false {
        var v8 bool = false
        p0 -= p0
        for fi9 := 0; fi9 < 9; fi9++ {
            fmt.Println(v4)
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 uint32 = 42
                switch (-(fi9)) {
                case -471:
                    var v12 int64 = -299
                    _ = v12
                case 4294967295:
                    v13 := []int{fi9, fi9}
                    _ = v13
                }
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        _ = v8
    } else {
        var v14 int = -855
        _ = v14
    }
    _ = p0
    _ = v4
    return -976
}
func genFn1(p0 int32) uint {
    v15 := GenS2{F0: -1, F1: 0, F2: 42}
    if (42 > 2) {
        v16 := []int{}
        var v17 int = 935
        _ = v16
        _ = v17
    } else {
        var v18 uint32 = 4294967295
        var v19 float32 = -2.25
        p0 = p0
        _ = v18
        _ = v19
    }
    v20 := []int{-224, -2147483648}
    var v21 float64 = float64(0.0)
    _ = p0
    _ = v15
    _ = v20
    _ = v21
    return 2
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{}
    v3 := make(map[int]bool)
    var v22 int32 = genFn0(genFn0(42))
    for fi23 := 0; fi23 < 19; fi23++ {
        fmt.Println(fi23)
        v24 := v3[(fi23 | fi23)]
        _ = fi23
        _ = v24
    }
    if false {
        delete(v1, 32768)
        switch 32768 {
        case 1:
            var v25 int8 = -1
            _ = v25
        default:
            v26 := GenS1{F0: 42}
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        var v28 uint32 = 65535
        for fi29 := 0; fi29 < 13; fi29++ {
            v28 -= 946
            switch fi29 {
            case 470:
                var v30 int16 = -810
                _ = v30
            case 32768:
                if true {
                    v22 -= (v22 / (((v22 + v22)) | 1))
                    var v31 uint8 = 191
                    v32 := GenS1{F0: 42}
                    _ = v31
                    _ = v32
                }
            }
            if (0 < 42) {
                var v33 int16 = 42
                _ = v33
            } else {
                var v34 string = "\n\t"
                fmt.Println(v22)
                fmt.Println(v34)
                _ = v34
            }
            v22 = -2147483648
            v35 := GenS2{F0: 1000, F1: 2, F2: 42}
            _ = fi29
            _ = v35
        }
        v36 := GenS3{F0: true, F1: -951, F2: v28}
        _ = fi27
        _ = v28
        _ = v36
    }
    if true {
        var v37 uint8 = 104
        switch 0 {
        case 0:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case -824:
            var v38 uint16 = 255
            _ = v38
        case -1:
            fmt.Println(v22)
        }
        v1[0] = (true && (genFn1(v22) >= 2))
        _ = v37
    }
    if false {
        var v39 float32 = 0.1
        v22 += (-(v22))
        _ = v39
    } else {
        var v40 uint = 255
        v41 := GenS4{F0: (v22 << ((genFn1((v22 - 1))) & 63)), F1: 883, F2: v40}
        v42 := GenS5{F0: v22}
        _ = v40
        _ = v41
        _ = v42
    }
    v22 = (v22 % ((v22) | 1))
    if (2147483647 > 1) {
        for fi43 := 0; fi43 < 6; fi43++ {
            var v44 string = ("\n\t" + "")
            v44 = v44
            _ = fi43
            _ = v44
        }
        for fi45 := 0; fi45 < 9; fi45++ {
            for fi46 := 0; fi46 < 12; fi46++ {
                switch ("\n\t" + "") {
                case "a":
                    var v47 int32 = v22
                    _ = v47
                case "\n\t":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    var v48 int8 = 7
                    _ = v48
                }
                _ = fi46
            }
            var v49 int8 = 53
            _ = fi45
            _ = v49
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
}
