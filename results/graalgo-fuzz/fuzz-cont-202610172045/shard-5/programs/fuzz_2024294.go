package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int8 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint64; F1 int; F2 uint }
func genFn0(p0 uint32, p1 uint32) int64 {
    p0 = 2147483647
    var v4 int32 = 2147483647
    v5 := []int{}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -9223372036854775808
}
func genFn1() int64 {
    switch -318 {
    case -591:
        var v6 bool = false
        _ = v6
    case 128:
        var v7 float64 = (0.0 / 1.5)
        _ = v7
    case 0:
        switch 32767 {
        case -2147483648:
            for fi8 := 0; fi8 < 15; fi8++ {
                v9 := []int{-165}
                if (!true) {
                    if true {
                        _ = len(v9)
                        for fi10 := 0; fi10 < 2; fi10++ {
                            v9 = append(v9, -9223372036854775808)
                            v11 := GenS1{F0: 62, F1: fi10, F2: 83}
                            switch 4294967295 {
                            case -9223372036854775808:
                                fmt.Println(fi8, fi8)
                            case 32768:
                                _ = len(v9)
                            }
                            _ = fi10
                            _ = v11
                        }
                    } else {
                        var v12 uint16 = 32767
                        v13 := GenS1{F0: 711, F1: 892, F2: -56}
                        _ = v12
                        _ = v13
                    }
                    fmt.Println(fi8)
                } else {
                    var v14 uint64 = 623
                    for fi15 := 0; fi15 < 4; fi15++ {
                        var v16 uint16 = 443
                        v14 += 1
                        _ = fi15
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 20; fi17++ {
                        if true {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                        }
                        v14 -= v14
                        var v18 float64 = 3.14159
                        v14 = (v14 % ((v14) | 1))
                        v18 -= 3.14159
                        if false {
                            if (!false) {
                                if len(v9) > 0 {
                                    v9[0] = (fi17 >> ((42) & 63))
                                }
                            }
                        } else {
                            switch "foo bar" {
                            case "foo bar":
                                _ = len(v9)
                            default:
                                var v19 uint = 955
                                _ = v19
                            }
                            var v20 bool = (4294967295 < 654)
                            var v21 int64 = -1
                            _ = v20
                            _ = v21
                        }
                        v9 = append(v9, fi8)
                        var v22 int16 = -32768
                        v18 -= v18
                        _ = fi17
                        _ = v18
                        _ = v22
                    }
                    _ = v14
                }
                var v23 bool = true
                _ = fi8
                _ = v23
                _ = v9
            }
        }
    }
    v24 := GenS2{F0: 4294967295, F1: true}
    var v25 uint16 = 42
    fmt.Println(v25)
    var v26 float32 = float32((0.0 + (1e10 * 0.1)))
    _ = v24
    _ = v25
    _ = v26
    return genFn0(726, 1)
}

func main() {
    var v1 int32 = 148
    var v2 int64 = 1
    var v3 bool = true
    if v3 {
        v3 = (1000 > -32768)
        v27 := []int{322}
        var v28 bool = v3
        _ = v27
        _ = v28
    } else {
        v29 := GenS2{F0: 955, F1: v3}
        v30 := []int{-2147483648, -9223372036854775808, 128, 127}
        v31 := GenS2{F0: 0, F1: v3}
        _ = v29
        _ = v30
        _ = v31
    }
    if (v3 && true) {
        if v3 {
            switch 734 {
            case 32767:
                v32 := GenS3{F0: 878, F1: 158, F2: 0}
                _ = v32
            }
        } else {
            fmt.Println(v3, v2, v2)
        }
        var v33 int = -411
        switch (v33 % ((int(v1)) | 1)) {
        case -2147483648:
            var v34 int = (-(v33))
            _ = v34
        }
        _ = v33
    } else {
        if v3 {
            var v35 uint8 = 127
            v2 = v2
            _ = v35
        }
    }
    v2 = v2
    var v36 int = 0
    var v37 bool = v3
    v38 := []int{v36}
    v39 := GenS3{F0: 1, F1: (-(v36)), F2: 2}
    for fi40 := 0; fi40 < 6; fi40++ {
        v2 = v2
        fmt.Println(v36, v2, v36)
        v2 = (v2 / ((v2) | 1))
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
}
