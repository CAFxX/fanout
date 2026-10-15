package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int64; F2 int64 }
type GenS3 struct { F0 uint16; F1 uint16; F2 int32 }
func genFn0() int64 {
    if (!false) {
        v4 := GenS1{F0: 1}
        _ = v4
    } else {
        var v5 int = -921
        if (66 >= 488) {
            switch "hello" {
            case "\n\t":
                v5 += v5
            default:
                if ("a" < "hello") {
                    v5 -= v5
                } else {
                    v5 = (v5 - 713)
                    v5 = 32767
                }
            }
        } else {
            v6 := GenS2{F0: 584, F1: -1, F2: 1}
            var v7 int16 = 0
            _ = v6
            _ = v7
        }
        for fi8 := 0; fi8 < 2; fi8++ {
            switch (v5 - (^(v5))) {
            case 32767:
                var v9 int8 = -88
                _ = v9
            case -1:
                var v10 int32 = 0
                _ = v10
            default:
                if true {
                    for fi11 := 0; fi11 < 12; fi11++ {
                        v12 := GenS3{F0: 42, F1: 255, F2: -404}
                        v13 := []int{}
                        var v14 int = (v5 * fi8)
                        if (true && (false || true)) {
                            var v15 uint8 = 255
                            var v16 uint8 = 127
                            switch "\n\t" {
                            case "a":
                                v14 -= fi8
                            default:
                                v15 -= v16
                            }
                            _ = v15
                            _ = v16
                        }
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                } else {
                    v17 := GenS2{F0: 1000, F1: -1, F2: 9223372036854775807}
                    var v18 int8 = -1
                    _ = v17
                    _ = v18
                }
            }
            v5 = v5
            _ = fi8
        }
        _ = v5
    }
    if true {
        if false {
            var v19 int64 = 0
            _ = v19
        } else {
            var v20 uint8 = 127
            v21 := []int{-276}
            _ = v20
            _ = v21
        }
        switch "a" {
        case "hello":
            var v22 uint64 = 1
            _ = v22
        }
    }
    var v23 int16 = -32768
    _ = v23
    return -524
}

func main() {
    var v1 uint64 = 484
    var v2 int16 = 715
    var v3 bool = (!true)
    if v3 {
        var v24 int64 = -9223372036854775808
        var v25 uint64 = v1
        var v26 uint8 = 42
        _ = v24
        _ = v25
        _ = v26
    }
    if v3 {
        var v27 float64 = 1e10
        v28 := []int{42, -2147483648, 4294967295}
        _ = v27
        _ = v28
    } else {
        v1 -= 42
        var v29 string = (("\n\t" + "a") + "hello")
        _ = v29
    }
    v3 = true
    if v3 {
        fmt.Println(v1)
        if (v3 || v3) {
            for fi30 := 0; fi30 < 16; fi30++ {
                var v31 uint64 = v1
                _ = fi30
                _ = v31
            }
            for fi32 := 0; fi32 < 10; fi32++ {
                v2 = v2
                _ = fi32
            }
        } else {
            v33 := []int{1, 128, 1, -2147483648}
            var v34 float64 = 0.0
            if len(v33) > 0 {
                v33[0] = -174
            }
            _ = v33
            _ = v34
        }
        var v35 int = 1
        _ = v35
    } else {
        var v36 uint16 = 744
        for fi37 := 0; fi37 < 12; fi37++ {
            switch fi37 {
            case -2147483648:
                v38 := []int{(^(fi37)), (fi37 % ((2147483647) | 1)), (fi37 << ((42) & 63)), fi37}
                _ = v38
            case 32767:
                v39 := GenS1{F0: 895}
                _ = v39
            case 32768:
                fmt.Println(v1)
            default:
                v40 := []int{}
                _ = v40
            }
            fmt.Println(v3, v2, v2)
            _ = fi37
        }
        v36 = v36
        _ = v36
    }
    var v41 float64 = float64(3.14159)
    v42 := GenS3{F0: 2, F1: 65535, F2: 874}
    switch 438 {
    case 328:
        for fi43 := 0; fi43 < 5; fi43++ {
            v41 = (v41 + v41)
            var v44 uint64 = 1
            _ = fi43
            _ = v44
        }
    case 712:
        fmt.Println(v1)
    case -358:
        var v45 uint16 = 669
        _ = v45
    }
    v46 := []int{-2147483648, 658, -2147483648, 32768}
    var v47 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v46
    _ = v47
}
