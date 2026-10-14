package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 string }

func main() {
    var v1 int8 = 42
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    fmt.Println(v1)
    fmt.Println(v1)
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := GenS1{F0: fi4}
        _ = fi4
        _ = v5
    }
    fmt.Println(v1)
    if true {
        fmt.Println(v1)
    }
    v6 := GenS2{F0: int16(v1)}
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS3{F0: 602, F1: (fi7 >> ((0) & 63))}
        _ = fi7
        _ = v8
    }
    var v9 float64 = float64(0.0)
    if false {
        var v10 float32 = 1e10
        switch 413 {
        case 4294967295:
            v10 = ((v10 * -2.25) + float32(v9))
        }
        for fi11 := 0; fi11 < 18; fi11++ {
            v9 = v9
            v9 += float64((3.14159 + v10))
            if (!(-2147483648 >= -528)) {
                v12 := GenS4{F0: "foo bar"}
                v10 = (v10 + 0.0)
                _ = v12
            } else {
                if (!true) {
                    var v13 int16 = 844
                    v14 := []int{2147483647, -105, fi11}
                    if (!(true || (4294967295 == 1))) {
                        if (v1 == v1) {
                            var v15 int8 = v1
                            v15 += -103
                            if true {
                                if ((93 != 42) && (((3.14159 + 3.14159) <= (float32((1e10 - (v9 - v9))) / 1e10)) || (!("foo bar" == (("hello" + "hello") + ""))))) {
                                    for fi16 := 0; fi16 < 11; fi16++ {
                                        v17 := GenS1{F0: 614}
                                        _ = fi16
                                        _ = v17
                                    }
                                } else {
                                    v3[575] = 65535
                                    var v18 int8 = (v1 % ((127) | 1))
                                    _ = v18
                                }
                            }
                            _ = v15
                        } else {
                            var v19 float64 = (v9 + (v9 * 1e10))
                            for fi20 := 0; fi20 < 14; fi20++ {
                                v9 -= v19
                                if true {
                                    if (42 >= 9223372036854775807) {
                                        var v21 uint32 = 878
                                        v22 := GenS3{F0: 512, F1: -1}
                                        _ = v21
                                        _ = v22
                                    }
                                }
                                _ = fi20
                            }
                            _ = v19
                        }
                    } else {
                        v23 := GenS4{F0: "hello"}
                        for fi24 := 0; fi24 < 12; fi24++ {
                            v25 := GenS1{F0: fi11}
                            fmt.Println(v9, fi24)
                            _ = fi24
                            _ = v25
                        }
                        _ = v23
                    }
                    _ = v13
                    _ = v14
                } else {
                    if false {
                        var v26 uint64 = 4294967295
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 17; fi27++ {
                        fmt.Println(v1)
                        _ = fi27
                    }
                    var v28 uint8 = 42
                    _ = v28
                }
                v29 := []int{32768}
                if (-32768 != -32768) {
                    v30 := GenS3{F0: 56, F1: (fi11 & fi11)}
                    if true {
                        for fi31 := 0; fi31 < 6; fi31++ {
                            v32 := v3[2147483647]
                            _ = fi31
                            _ = v32
                        }
                    } else {
                        v9 = v9
                    }
                    var v33 uint64 = 2
                    _ = v30
                    _ = v33
                }
                _ = v29
            }
            _ = fi11
        }
        _ = v10
    } else {
        switch 42 {
        case 32768:
            switch 1 {
            case 0:
                var v34 uint8 = 255
                _ = v34
            }
        }
        var v35 int = -245
        _ = v35
    }
    v2[("" + "a")] = -434
    fmt.Println(v1)
    var v36 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v6
    _ = v9
}
