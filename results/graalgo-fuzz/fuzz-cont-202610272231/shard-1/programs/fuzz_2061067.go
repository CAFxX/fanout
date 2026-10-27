package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int; F2 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 string; F2 int }
func genFn0(p0 uint64, p1 int8, p2 string) uint8 {
    var v4 int16 = -1
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 float32 = 3.14159
        fmt.Println(p1)
        _ = fi5
        _ = v6
    }
    switch -419 {
    case -1:
        var v7 uint8 = 128
        _ = v7
    default:
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS1{F0: 42}
            if true {
                var v10 float32 = (1e10 + 3.14159)
                _ = v10
            }
            v11 := GenS2{F0: 2, F1: fi8, F2: 555}
            _ = fi8
            _ = v11
            _ = v9
        }
    }
    v12 := GenS3{F0: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v4
    return 127
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float64 = 1e10
    v3 := []string{}
    for fi13 := 0; fi13 < 15; fi13++ {
        v1 -= 2
        var v14 bool = false
        fmt.Println(v1)
        _ = fi13
        _ = v14
    }
    _ = len(v3)
    var v15 int32 = 1
    var v16 string = "\n\t"
    fmt.Println(v1)
    var v17 bool = false
    v18 := GenS3{F0: 4294967295}
    v19 := GenS4{F0: (v1 * 0), F1: (v16 + "hello"), F2: -9223372036854775808}
    for fi20 := 0; fi20 < 19; fi20++ {
        if v17 {
            for fi21 := 0; fi21 < 10; fi21++ {
                if (v1 > (v1 & v1)) {
                    fmt.Println(fi20, v1, v16)
                    switch (v16 + v16) {
                    case "\n\t":
                        var v22 uint64 = (v1 * (v1 >> ((0) & 63)))
                        _ = v22
                    }
                }
                _ = fi21
            }
            switch (("\n\t" + v16) + v16) {
            case "foo bar":
                if (true || v17) {
                    if v17 {
                        v2 = (-2.25 * v2)
                    } else {
                        if v17 {
                            var v23 float64 = (-2.25 * (v2 + -2.25))
                            _ = v23
                        } else {
                            v24 := GenS3{F0: 65535}
                            _ = v24
                        }
                        var v25 uint = 4294967295
                        _ = v25
                    }
                } else {
                    v1 -= (v1 + (v1 & (v1 << ((51) & 63))))
                    var v26 uint = 4294967295
                    if v17 {
                        switch (v16 + v16) {
                        case "foo bar":
                            if (255 < 42) {
                                if (v17 && (true && v17)) {
                                    var v27 float32 = 0.1
                                    _ = v27
                                }
                                if ((v17 && true) && true) {
                                    switch v16 {
                                    case "":
                                        var v28 float64 = v2
                                        _ = v28
                                    case "hello":
                                        fmt.Println(v26)
                                    case "foo bar":
                                        var v29 int32 = (^(v15))
                                        _ = v29
                                    }
                                    for fi30 := 0; fi30 < 6; fi30++ {
                                        if (v17 && v17) {
                                            v16 = v16
                                            v26 = (v26 - v26)
                                        }
                                        v31 := GenS3{F0: v26}
                                        _ = fi30
                                        _ = v31
                                    }
                                }
                                fmt.Println(v1, v15, v17)
                            }
                        }
                    } else {
                        switch ("hello" + "foo bar") {
                        case "foo bar":
                            var v32 int64 = -105
                            _ = v32
                        default:
                            _ = len(v3)
                        }
                        v17 = (127 >= 128)
                        var v33 int32 = 2147483647
                        _ = v33
                    }
                    _ = v26
                }
            case "a":
                var v34 int64 = 548
                _ = v34
            case "hello":
                v3 = append(v3, v16)
            default:
                for fi35 := 0; fi35 < 20; fi35++ {
                    var v36 float64 = (v2 + -2.25)
                    v1 = v1
                    _ = fi35
                    _ = v36
                }
            }
            if false {
                if false {
                    for fi37 := 0; fi37 < 4; fi37++ {
                        var v38 uint16 = 255
                        var v39 uint = 18446744073709551615
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                }
                var v40 uint16 = 379
                _ = v40
            }
        }
        _ = fi20
    }
    v1 -= (^(v1))
    var v41 int = 32768
    switch (v41 ^ v41) {
    case -1:
        fmt.Println(v17)
    case -2147483648:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        var v42 float32 = float32(((float64(1e10) + (v2 * v2)) + float64(-2.25)))
        _ = v42
    }
    var v43 int = v41
    for fi44 := 0; fi44 < 8; fi44++ {
        var v45 uint8 = 255
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi44
        _ = v45
    }
    var v46 float64 = 1e10
    fmt.Println(v1, v15)
    fmt.Println(v1, v1, v1)
    if ((42 != 126) && v17) {
        v47 := GenS1{F0: -1}
        _ = v47
    } else {
        v3 = append(v3, v16)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v41
    _ = v43
    _ = v46
}
