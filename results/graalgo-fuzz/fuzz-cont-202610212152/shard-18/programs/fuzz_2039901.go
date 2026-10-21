package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 int8; F2 bool }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    var v3 uint8 = 128
    var v4 string = ""
    var v5 int32 = -1
    for fi6 := 0; fi6 < 17; fi6++ {
        if true {
            if (false && true) {
                var v7 uint16 = 2
                v8 := []int{fi6, fi6, 1, (fi6 & fi6)}
                fmt.Println(v4)
                _ = v7
                _ = v8
            }
            v9 := []int{9223372036854775807, 128}
            if (false && (!true)) {
                v4 = ((v4 + ("foo bar" + "\n\t")) + "")
                v10 := GenS1{F0: v4}
                if ((-2.25 * 1e10) <= (float32((((1e10 * (((((float64(0.1) - 1e10) / (3.14159 - 0.0)) * (1e10 / float64(0.1))) * 1e10) - float64(1.5))) - 0.0) + 1e10)) - 0.0)) {
                    fmt.Println(v3, v3)
                    var v11 float64 = 0.0
                    _ = v11
                }
                _ = v10
            } else {
                v12 := GenS1{F0: v4}
                _ = v12
            }
            _ = v9
        } else {
            if true {
                v13 := []int{fi6, fi6, (fi6 & fi6), fi6}
                _ = v13
            } else {
                fmt.Println(v3, v4, v5)
                if true {
                    for fi14 := 0; fi14 < 19; fi14++ {
                        var v15 uint16 = 65535
                        if ((^(fi14)) <= (fi6 * 128)) {
                            v16 := []int{}
                            _ = v16
                        }
                        _ = fi14
                        _ = v15
                    }
                } else {
                    switch 470 {
                    case -9223372036854775808:
                        v5 -= v5
                    case -2147483648:
                        if ("" < v4) {
                            var v17 bool = true
                            var v18 int64 = -480
                            _ = v17
                            _ = v18
                        }
                    case 32768:
                        v19 := GenS1{F0: v4}
                        _ = v19
                    default:
                        v20 := []int{127, fi6, fi6}
                        _ = v20
                    }
                    v3 += v3
                    if (!false) {
                        if false {
                            var v21 uint = 243
                            _ = v21
                        }
                        fmt.Println(fi6)
                        v3 += v3
                    }
                }
                var v22 int8 = 113
                _ = v22
            }
            for fi23 := 0; fi23 < 17; fi23++ {
                v24 := GenS2{F0: 694, F1: v5}
                _ = fi23
                _ = v24
            }
            fmt.Println(v4)
        }
        v25 := []int{fi6, fi6, fi6, 128}
        v5 -= v5
        _ = fi6
        _ = v25
    }
    v4 = v4
    var v26 int8 = 33
    for fi27 := 0; fi27 < 1; fi27++ {
        var v28 int64 = 1
        _ = fi27
        _ = v28
    }
    var v29 int = -672
    v29 += v29
    switch v29 {
    case 741:
        var v30 string = "a"
        _ = v30
    case -698:
        var v31 uint8 = v3
        _ = v31
    default:
        switch 916 {
        case 368:
            for fi32 := 0; fi32 < 6; fi32++ {
                for fi33 := 0; fi33 < 7; fi33++ {
                    var v34 int64 = 1
                    var v35 uint16 = 255
                    var v36 uint32 = 904
                    _ = fi33
                    _ = v34
                    _ = v35
                    _ = v36
                }
                var v37 float32 = (1e10 - (0.0 / 3.14159))
                _ = fi32
                _ = v37
            }
        case -9223372036854775808:
            var v38 uint16 = 906
            _ = v38
        case 32767:
            v5 = (-(v5))
        }
    }
    v26 -= v26
    var v39 int = 1
    var v40 uint64 = 4294967295
    var v41 int8 = v26
    for fi42 := 0; fi42 < 5; fi42++ {
        for fi43 := 0; fi43 < 15; fi43++ {
            v3 = (v3 * v3)
            var v44 float64 = 1e10
            v3 -= (v3 + 42)
            v45 := GenS1{F0: ("" + (v4 + v4))}
            _ = fi43
            _ = v44
            _ = v45
        }
        v46 := GenS1{F0: v4}
        fmt.Println(v29)
        _ = fi42
        _ = v46
    }
    if (!(-1 < 1)) {
        v47 := GenS3{F0: 570}
        v39 += (^(v39))
        v48 := GenS4{F0: 1, F1: v26, F2: false}
        _ = v47
        _ = v48
    }
    var v49 uint8 = (v3 - uint8(v40))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v49
    _ = v5
}
