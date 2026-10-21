package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS5 struct { F0 int16; F1 uint }
func genFn0() uint16 {
    for fi4 := 0; fi4 < 4; fi4++ {
        fmt.Println(fi4)
        var v5 bool = (true && (false && (!(126 >= 32767))))
        _ = fi4
        _ = v5
    }
    var v6 bool = (!false)
    v7 := []int{1, 23, -118, 127}
    for fi8 := 0; fi8 < 9; fi8++ {
        _ = len(v7)
        v9 := GenS1{F0: -128}
        fmt.Println(v6, fi8)
        _ = fi8
        _ = v9
    }
    switch -844 {
    case -9223372036854775808:
        var v10 int64 = 1
        _ = v10
    case -2147483648:
        v6 = (-9223372036854775808 < 0)
    case 2147483647:
        v6 = v6
    }
    _ = v6
    _ = v7
    return 260
}

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 100
    var v3 uint8 = 1
    if (true && (-1 != 32767)) {
        if true {
            if (2147483648 == 415) {
                if (!(!false)) {
                    var v11 uint = 65535
                    _ = v11
                } else {
                    var v12 int = 42
                    _ = v12
                }
                if (-2147483648 > -9223372036854775808) {
                    if (2147483647 < -626) {
                        fmt.Println(v3, v3)
                    }
                    fmt.Println(v3, v3)
                } else {
                    v3 = (v3 | v3)
                    var v13 uint16 = 501
                    if true {
                        v3 = (v3 << ((1) & 63))
                        var v14 uint16 = (v13 + v13)
                        for fi15 := 0; fi15 < 17; fi15++ {
                            fmt.Println(fi15, v3)
                            v16 := GenS1{F0: -75}
                            for fi17 := 0; fi17 < 7; fi17++ {
                                if true {
                                    var v18 float64 = 0.1
                                    if (4294967295 < 0) {
                                        if false {
                                            var v19 float32 = float32(1e10)
                                            _ = v19
                                        }
                                        var v20 int64 = 42
                                        _ = v20
                                    } else {
                                        switch (fi15 - (fi17 >> ((0) & 63))) {
                                        case 2147483647:
                                            v14 = v14
                                            switch (fi17 - 9223372036854775807) {
                                            case 0:
                                                v13 = 1
                                                var v21 uint = 822
                                                _ = v21
                                            case 4294967295:
                                                fmt.Println(fi17, v13)
                                            }
                                        case -1:
                                            delete(v1, "")
                                        case 663:
                                            var v22 uint64 = 18446744073709551615
                                            _ = v22
                                        }
                                        var v23 uint16 = v14
                                        _ = v23
                                    }
                                    v24 := GenS1{F0: v2}
                                    _ = v18
                                    _ = v24
                                } else {
                                    var v25 string = "\n\t"
                                    v26 := []int{fi17, (fi17 | 42), fi17, fi15}
                                    _ = v25
                                    _ = v26
                                }
                                v3 = 42
                                v3 -= v3
                                _ = fi17
                            }
                            _ = fi15
                            _ = v16
                        }
                        _ = v14
                    } else {
                        v27 := []int{0, -38, 128, 4294967295}
                        _ = v27
                    }
                    _ = v13
                }
                var v28 int16 = -978
                _ = v28
            }
            switch "hello" {
            case "":
                if (630 == 0) {
                    v2 = -1
                    _ = len(v1)
                } else {
                    if true {
                        v2 = v2
                        var v29 string = "\n\t"
                        _ = v29
                    } else {
                        fmt.Println(v3, v3)
                        v2 = -128
                        v30 := GenS1{F0: (-(v2))}
                        _ = v30
                    }
                    if true {
                        v3 -= 53
                        v3 -= v3
                    } else {
                        v2 -= v2
                        v31 := []int{}
                        v3 = (v3 * (^(v3)))
                        _ = v31
                    }
                }
            case "hello":
                fmt.Println(v2, v3)
            }
        } else {
            v32 := []int{-561, -524, 9223372036854775807}
            _ = v32
        }
        fmt.Println(v3, v2)
        v3 = (v3 % (((v3 / (((v3 * 97)) | 1))) | 1))
    }
    v33 := GenS1{F0: v2}
    v34 := v1["\n\t"]
    v35 := GenS1{F0: v2}
    v36 := []int{744}
    v37 := GenS1{F0: (v2 % ((v2) | 1))}
    if true {
        for fi38 := 0; fi38 < 7; fi38++ {
            v2 = (v2 / (((^(v2))) | 1))
            v39 := GenS2{F0: v34}
            _ = fi38
            _ = v39
        }
        v40 := GenS3{F0: (false || (!(65535 == 65535))), F1: 42, F2: 2}
        _ = v40
    }
    v41 := GenS4{F0: 2147483647, F1: 0, F2: v2}
    switch v34 {
    case "hello":
        for fi42 := 0; fi42 < 18; fi42++ {
            v34 = "\n\t"
            _ = fi42
        }
    case "foo bar":
        _ = len(v36)
    }
    v34 = ("a" + v34)
    if (float64(float32(1.5)) > (((0.1 + 0.0) * (1.5 * 0.1)) / 1.0)) {
        v43 := GenS5{F0: 134, F1: 4294967295}
        var v44 float32 = float32(1.5)
        for fi45 := 0; fi45 < 6; fi45++ {
            var v46 float32 = ((((float32(1.5) / (v44 - (v44 - v44))) - -2.25) - v44) / (v44 + (v44 - (3.14159 / 3.14159))))
            _ = fi45
            _ = v46
        }
        _ = v43
        _ = v44
    } else {
        fmt.Println(v2, v34, v3)
        v3 = 42
        if true {
            for fi47 := 0; fi47 < 6; fi47++ {
                v48 := GenS2{F0: v34}
                if len(v36) > 0 {
                    v36[0] = 1
                }
                _ = fi47
                _ = v48
            }
        }
    }
    if (!(42 > 0)) {
        v2 -= (v2 / ((v2) | 1))
    }
    v49 := v1["hello"]
    v50 := GenS5{F0: 0, F1: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v41
    _ = v49
    _ = v50
}
