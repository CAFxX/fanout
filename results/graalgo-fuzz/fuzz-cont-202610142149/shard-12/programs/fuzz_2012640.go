package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int32; F1 uint64; F2 int }
type GenS4 struct { F0 int8; F1 int16; F2 bool }
type GenS5 struct { F0 int8 }

func main() {
    var v1 uint16 = 346
    v2 := []int{42}
    var v3 string = "a"
    var v4 int64 = -1
    var v5 float64 = float64(((float32(float64(1.5)) * 0.0) * 0.1))
    var v6 string = "\n\t"
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: 2147483647}
        _ = fi7
        _ = v8
    }
    if false {
        var v9 int64 = v4
        if false {
            var v10 uint = 0
            v11 := GenS1{F0: -219}
            v12 := GenS1{F0: 73}
            _ = v10
            _ = v11
            _ = v12
        } else {
            if len(v2) > 0 {
                v2[0] = 32768
            }
            v1 = 65535
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v9
    } else {
        var v13 int64 = (v4 - -1)
        var v14 uint8 = 255
        _ = v13
        _ = v14
    }
    v3 = v3
    var v15 int = 611
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((390 != 2) || ((1 > 106) || (false || true))) {
        _ = len(v2)
        if true {
            v16 := GenS2{F0: 2, F1: 4294967295}
            _ = v16
        }
    }
    v3 = ""
    if true {
        if false {
            for fi17 := 0; fi17 < 9; fi17++ {
                if (2 == 65535) {
                    for fi18 := 0; fi18 < 10; fi18++ {
                        v1 = 32767
                        if false {
                            v2 = append(v2, fi18)
                        } else {
                            _ = len(v2)
                            if false {
                                v2 = append(v2, fi18)
                                var v19 int64 = v4
                                v5 = float64(1e10)
                                _ = v19
                            } else {
                                if (42 > 0) {
                                    v6 = "\n\t"
                                    _ = len(v2)
                                }
                            }
                            v4 += (v4 & (-(v4)))
                        }
                        var v20 float32 = 0.0
                        _ = fi18
                        _ = v20
                    }
                    switch "hello" {
                    case "a":
                        fmt.Println(v6, v15)
                    case "foo bar":
                        v3 = v6
                    }
                } else {
                    var v21 int32 = 475
                    _ = v21
                }
                var v22 int = -1
                _ = fi17
                _ = v22
            }
        }
        v23 := GenS3{F0: 2147483647, F1: 559, F2: (-(v15))}
        var v24 uint32 = 2147483647
        _ = v23
        _ = v24
    } else {
        var v25 int8 = 1
        for fi26 := 0; fi26 < 6; fi26++ {
            v27 := GenS2{F0: 0, F1: 2147483648}
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    var v28 uint32 = 1
    v1 -= v1
    var v29 float64 = (v5 - (v5 * (v5 + 0.1)))
    v15 += (v15 << ((uint(v1)) & 63))
    if ((v15 % ((v15) | 1)) == (v15 << ((4294967295) & 63))) {
        for fi30 := 0; fi30 < 7; fi30++ {
            for fi31 := 0; fi31 < 7; fi31++ {
                if true {
                    v1 = v1
                    if (78 < 1) {
                        if true {
                            var v32 int32 = 2147483647
                            var v33 uint8 = 255
                            v1 += v1
                            _ = v32
                            _ = v33
                        } else {
                            v34 := GenS1{F0: -196}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v34
                        }
                        var v35 uint16 = 745
                        fmt.Println(v35, v4, v5)
                        fmt.Println(v29, v3)
                        _ = v35
                    }
                }
                if len(v2) > 0 {
                    v2[0] = fi31
                }
                v3 = "hello"
                _ = fi31
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v2 = append(v2, v15)
            _ = fi30
        }
        var v36 int = v15
        v1 = v1
        _ = v36
    } else {
        if false {
            v37 := GenS4{F0: -1, F1: 1, F2: (303 > 392)}
            var v38 uint = 65535
            _ = v37
            _ = v38
        } else {
            fmt.Println(v5, v5)
            fmt.Println(v5)
            var v39 int = (v15 | 4294967295)
            _ = v39
        }
        v2 = append(v2, (v15 ^ -9223372036854775808))
        v40 := GenS2{F0: v28, F1: 2}
        _ = v40
    }
    v41 := GenS5{F0: 100}
    var v42 int8 = int8(v1)
    if (uint32(v1) == v28) {
        v43 := GenS1{F0: 2147483647}
        _ = v43
    } else {
        var v44 int8 = 0
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
