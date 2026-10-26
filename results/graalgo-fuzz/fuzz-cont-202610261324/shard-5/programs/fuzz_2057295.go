package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 bool }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = -9223372036854775808
    var v3 int32 = -79
    var v4 int = 42
    v5 := make(map[int]bool)
    v4 -= (v4 % (((v4 ^ 928)) | 1))
    if ((42 <= 224) && false) {
        var v6 int32 = (^(v3))
        _ = v6
    }
    v7 := GenS1{F0: 4294967295, F1: 501, F2: (!false)}
    fmt.Println(v2)
    fmt.Println(v3, v3)
    var v8 uint = 65535
    v9 := []int{v4, (v4 & v4), (^(v4))}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 string = ("\n\t" + "foo bar")
        var v12 float32 = (1e10 - (0.0 / float32(3.14159)))
        v13 := GenS1{F0: 4294967295, F1: 42, F2: (4294967295 <= fi10)}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    v3 -= (^(v3))
    var v14 int64 = (v2 - 42)
    _ = len(v5)
    switch (v4 / (((v4 - int(v3))) | 1)) {
    case -888:
        if (false || ((0.1 == 1e10) || false)) {
            for fi15 := 0; fi15 < 18; fi15++ {
                if ((v14 + (v14 | v14)) >= v14) {
                    var v16 float32 = -2.25
                    v3 += 0
                    fmt.Println(v16, v14, v2)
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 11; fi17++ {
                        v3 = -1
                        _ = fi17
                    }
                    fmt.Println(v8)
                }
                _ = fi15
            }
            if (235 < 42) {
                for fi18 := 0; fi18 < 19; fi18++ {
                    if (-32768 <= 1000) {
                        var v19 float32 = 1e10
                        v8 = v8
                        for fi20 := 0; fi20 < 11; fi20++ {
                            v21 := GenS1{F0: 18446744073709551615, F1: 42, F2: true}
                            fmt.Println(v4)
                            _ = fi20
                            _ = v21
                        }
                        _ = v19
                    } else {
                        delete(v5, (fi18 << (((v8 << ((v8) & 63))) & 63)))
                        var v22 float64 = (0.0 * 0.1)
                        _ = v22
                    }
                    _ = fi18
                }
            } else {
                var v23 string = "\n\t"
                if (42 != 255) {
                    switch v4 {
                    case -2147483648:
                        var v24 int8 = 42
                        _ = v24
                    case 1:
                        var v25 int16 = 42
                        _ = v25
                    case 2147483647:
                        if false {
                            v14 -= 926
                        }
                    default:
                        switch (^(v4)) {
                        case 128:
                            for fi26 := 0; fi26 < 19; fi26++ {
                                _ = len(v9)
                                v27 := v5[(-(v4))]
                                _ = fi26
                                _ = v27
                            }
                        default:
                            if (!true) {
                                fmt.Println(v2)
                                fmt.Println(v14, v8, v4)
                            } else {
                                var v28 bool = true
                                v8 = v8
                                _ = v28
                            }
                        }
                    }
                    v3 = v3
                    if false {
                        var v29 int64 = v14
                        _ = v29
                    }
                } else {
                    v30 := v5[(v4 | v4)]
                    fmt.Println(v23)
                    _ = v30
                }
                v23 = ("foo bar" + v23)
                _ = v23
            }
            v9 = append(v9, v4)
        }
    case 2147483647:
        for fi31 := 0; fi31 < 13; fi31++ {
            var v32 uint32 = 42
            _ = fi31
            _ = v32
        }
    case 32768:
        fmt.Println(v4)
    }
    v33 := GenS1{F0: 2, F1: 4, F2: false}
    v2 -= v14
    switch "hello" {
    case "hello":
        var v34 uint64 = 0
        _ = v34
    case "a":
        var v35 uint16 = 878
        _ = v35
    }
    for fi36 := 0; fi36 < 4; fi36++ {
        var v37 int16 = 1
        v38 := GenS1{F0: 53, F1: 2, F2: true}
        _ = fi36
        _ = v37
        _ = v38
    }
    var v39 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v39
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
