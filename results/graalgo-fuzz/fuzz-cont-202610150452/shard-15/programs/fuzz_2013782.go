package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint16 }

func main() {
    var v1 uint = 706
    var v2 int = 4294967295
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    switch (^(v2)) {
    case -1:
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 uint8 = 255
            v6 = (v6 % ((v6) | 1))
            _ = fi5
            _ = v6
        }
    case 9223372036854775807:
        v1 -= v1
    case 4294967295:
        if true {
            if true {
                var v7 int32 = 2147483647
                if (v2 > (^(v2))) {
                    var v8 int16 = -1
                    var v9 uint64 = 42
                    var v10 int8 = 0
                    _ = v10
                    _ = v8
                    _ = v9
                }
                var v11 uint32 = 42
                _ = v11
                _ = v7
            } else {
                fmt.Println(v1, v2)
                v12 := GenS1{F0: -1}
                var v13 int = 32768
                _ = v12
                _ = v13
            }
            var v14 uint64 = 0
            _ = v14
        }
    default:
        v1 = v1
    }
    fmt.Println(v1, v1)
    fmt.Println(v2)
    var v15 uint32 = 65535
    switch "hello" {
    case "a":
        switch "\n\t" {
        case "foo bar":
            if true {
                var v16 float32 = (1e10 - 1.5)
                v16 -= float32((0.1 + 0.0))
                for fi17 := 0; fi17 < 18; fi17++ {
                    for fi18 := 0; fi18 < 11; fi18++ {
                        if (1e10 >= ((1.5 - 1e10) * 3.14159)) {
                            if false {
                                delete(v4, 42)
                                if true {
                                    v15 -= v15
                                } else {
                                    fmt.Println(fi17, v15)
                                    v19 := v3[516]
                                    var v20 int64 = 832
                                    _ = v19
                                    _ = v20
                                }
                            } else {
                                v21 := GenS1{F0: 339}
                                v22 := v3[(v15 >> (((^(v1))) & 63))]
                                _ = v21
                                _ = v22
                            }
                        }
                        v15 -= (v15 % ((v15) | 1))
                        _ = fi18
                    }
                    v2 = (v2 ^ -567)
                    _ = fi17
                }
                _ = v16
            }
        case "a":
            if true {
                var v23 uint16 = 42
                _ = v23
            } else {
                if (false || (false || (false || ((!false) || false)))) {
                    switch (v2 - v2) {
                    case -9223372036854775808:
                        var v24 int32 = -1
                        _ = v24
                    default:
                        var v25 float64 = float64(0.1)
                        _ = v25
                    }
                    v1 -= v1
                    v26 := GenS1{F0: -1}
                    _ = v26
                } else {
                    var v27 uint32 = 468
                    _ = v27
                }
                if (v1 != (v1 ^ v1)) {
                    var v28 int8 = -128
                    _ = v28
                } else {
                    var v29 int = (v2 >> ((v1) & 63))
                    _ = v29
                }
                if true {
                    delete(v3, (v15 * v15))
                    fmt.Println(v1)
                } else {
                    if (!(-100 != 100)) {
                        v30 := GenS2{F0: (v15 / ((v15) | 1)), F1: 2}
                        switch "hello" {
                        case "hello":
                            var v31 int32 = 2147483647
                            _ = v31
                        }
                        _ = v30
                    }
                }
            }
        case "hello":
            if true {
                v1 -= (^(v1))
            }
        default:
            var v32 int8 = 42
            _ = v32
        }
    case "foo bar":
        v33 := []int{4294967295, 0, v2}
        _ = v33
    }
    if false {
        fmt.Println(v1, v2, v2)
        var v34 int32 = 0
        _ = v34
    } else {
        var v35 uint32 = (v15 | 4294967295)
        _ = v35
    }
    v36 := []int{}
    v1 += (^(v1))
    delete(v3, (v15 * 2147483648))
    for fi37 := 0; fi37 < 12; fi37++ {
        var v38 int16 = 32767
        _ = fi37
        _ = v38
    }
    v1 = (v1 | 18446744073709551615)
    var v39 int8 = 100
    v2 += -318
    if false {
        v2 = v2
    }
    if len(v36) >= 2 {
        v36 = v36[1:len(v36)]
    }
    switch "hello" {
    case "foo bar":
        v40 := GenS2{F0: (v15 & 2147483647), F1: 1}
        _ = v40
    default:
        v15 -= 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
}
