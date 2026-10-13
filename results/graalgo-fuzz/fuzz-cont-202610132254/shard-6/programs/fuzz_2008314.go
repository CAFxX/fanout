package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := []uint32{2, 2147483648}
    var v2 int16 = -32768
    var v3 uint32 = 65535
    var v4 uint32 = 2
    switch 42 {
    case 273:
        fmt.Println(v3, v4)
    case 0:
        if false {
            if (!(false && false)) {
                _ = len(v1)
            } else {
                var v5 int32 = 2147483647
                _ = v5
            }
            for fi6 := 0; fi6 < 4; fi6++ {
                if false {
                    if (!((98 >= 67) && false)) {
                        if (float64((3.14159 / 1.0)) >= 0.0) {
                            var v7 int = (fi6 / ((fi6) | 1))
                            _ = v7
                        } else {
                            if (true || (!true)) {
                                var v8 int = 1
                                var v9 uint8 = 77
                                _ = len(v1)
                                _ = v8
                                _ = v9
                            }
                        }
                    }
                    for fi10 := 0; fi10 < 20; fi10++ {
                        if (!false) {
                            if false {
                                fmt.Println(v4)
                            } else {
                                var v11 int8 = int8(v2)
                                _ = v11
                            }
                            var v12 uint8 = 128
                            var v13 int8 = 68
                            _ = v12
                            _ = v13
                        }
                        v2 = 1
                        _ = fi10
                    }
                    var v14 int8 = -128
                    _ = v14
                }
                var v15 uint8 = 42
                switch "\n\t" {
                case "\n\t":
                    for fi16 := 0; fi16 < 10; fi16++ {
                        if (!true) {
                            v4 -= 65535
                        } else {
                            v2 = v2
                            if (true || true) {
                                var v17 int32 = 0
                                _ = v17
                            } else {
                                fmt.Println(v4, fi6, fi16)
                                v18 := GenS1{F0: (2 >= (v3 & v4))}
                                _ = v18
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        if len(v1) > 0 {
                            v1[0] = (^(v4))
                        }
                        v3 += 2147483648
                        _ = fi16
                    }
                case "foo bar":
                    var v19 uint32 = (v3 / ((v4) | 1))
                    _ = v19
                default:
                    var v20 int = fi6
                    _ = v20
                }
                _ = fi6
                _ = v15
            }
            for fi21 := 0; fi21 < 6; fi21++ {
                v22 := GenS1{F0: (!(true || true))}
                if true {
                    if false {
                        v3 = (v3 / (((v3 - (v3 >> ((18446744073709551615) & 63)))) | 1))
                    }
                } else {
                    if false {
                        v2 -= v2
                        v2 += 42
                        v2 += (v2 + 564)
                    }
                }
                switch (^(fi21)) {
                case 32767:
                    for fi23 := 0; fi23 < 12; fi23++ {
                        var v24 uint64 = 141
                        v3 += 2147483647
                        _ = fi23
                        _ = v24
                    }
                default:
                    v25 := GenS1{F0: ((!false) && false)}
                    _ = v25
                }
                _ = fi21
                _ = v22
            }
        }
    }
    var v26 float64 = 1.5
    _ = len(v1)
    for fi27 := 0; fi27 < 5; fi27++ {
        if false {
            var v28 float32 = 3.14159
            var v29 uint16 = 0
            v30 := GenS2{F0: v2, F1: 907, F2: v2}
            _ = v28
            _ = v29
            _ = v30
        }
        for fi31 := 0; fi31 < 11; fi31++ {
            v3 -= 42
            var v32 float32 = ((((3.14159 - (1.5 + ((0.1 * 1e10) - -2.25))) * 0.1) * float32(v26)) + 1.5)
            var v33 float32 = (-2.25 + (1e10 - 1.5))
            _ = fi31
            _ = v32
            _ = v33
        }
        v34 := GenS2{F0: v2, F1: v3, F2: 32767}
        _ = fi27
        _ = v34
    }
    if len(v1) > 0 {
        v1[0] = v4
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        v36 := GenS3{F0: v3}
        _ = fi35
        _ = v36
    }
    v26 = v26
    var v37 bool = true
    v38 := GenS3{F0: 2147483648}
    var v39 int8 = -128
    if v37 {
        var v40 int16 = v2
        _ = v40
    } else {
        v3 = (^(v4))
        if ((990 < 65535) || (!true)) {
            var v41 uint = 4294967295
            _ = v41
        } else {
            v39 = v39
            fmt.Println(v37, v37, v3)
            var v42 uint = 183
            _ = v42
        }
    }
    fmt.Println(v37)
    for fi43 := 0; fi43 < 16; fi43++ {
        _ = len(v1)
        var v44 uint = 4294967295
        _ = fi43
        _ = v44
    }
    switch 4294967295 {
    case 42:
        if (455 > 1) {
            v26 = v26
        }
    }
    var v45 uint16 = 250
    if len(v1) > 0 {
        v1[0] = 1
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v45
}
