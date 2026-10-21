package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 bool; F1 int8; F2 bool }
type GenS7 struct { F0 uint64 }
func genFn0() int16 {
    v6 := []int{9223372036854775807, 42, 42}
    v7 := GenS1{F0: false, F1: 42}
    _ = len(v6)
    var v8 int = 0
    if (true || false) {
        v8 -= -9223372036854775808
    } else {
        v8 -= (v8 + 128)
        v9 := GenS1{F0: false, F1: 850}
        _ = v9
    }
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []string{"hello", "\n\t"}
    v2 := make(map[uint32]bool)
    var v3 int8 = 1
    v4 := make(map[uint32]bool)
    var v5 bool = true
    if v5 {
        var v10 int64 = 9223372036854775807
        _ = v10
    } else {
        v11 := GenS1{F0: (936 == 1), F1: 42}
        _ = v11
    }
    var v12 int8 = (-(v3))
    v12 += (v12 << ((65535) & 63))
    v5 = v5
    if ((((0.1 * 1.5) >= (float32(float64(((float32(1e10) * -2.25) + 1e10))) / float32(1e10))) && (780 > 1)) || v5) {
        switch "foo bar" {
        case "\n\t":
            if (1 < 141) {
                for fi13 := 0; fi13 < 1; fi13++ {
                    fmt.Println(v12, v3, v5)
                    var v14 int64 = 9223372036854775807
                    v14 = (v14 >> ((255) & 63))
                    _ = fi13
                    _ = v14
                }
                v5 = ("hello" <= ((("foo bar" + "foo bar") + "") + (("foo bar" + "hello") + "hello")))
            }
        case "a":
            switch 0 {
            case 32768:
                v15 := GenS2{F0: 799, F1: 42}
                _ = v15
            case 42:
                for fi16 := 0; fi16 < 13; fi16++ {
                    v17 := GenS3{F0: 1, F1: 559}
                    _ = fi16
                    _ = v17
                }
            case 295:
                if v5 {
                    for fi18 := 0; fi18 < 14; fi18++ {
                        if v5 {
                            fmt.Println(v12, v3)
                            fmt.Println(v3, v5)
                            if false {
                                v19 := GenS4{F0: 520, F1: 65535, F2: 1}
                                _ = v19
                            }
                        }
                        _ = fi18
                    }
                } else {
                    var v20 float64 = 0.1
                    var v21 uint32 = 2147483647
                    _ = v20
                    _ = v21
                }
            }
        }
    }
    fmt.Println(v5, v3)
    v22 := GenS5{F0: 65535, F1: 578}
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 uint64 = 2
        v1 = append(v1, "foo bar")
        if (v5 && (!v5)) {
            var v25 int16 = 32767
            for fi26 := 0; fi26 < 19; fi26++ {
                switch (-(fi26)) {
                case -2147483648:
                    var v27 float32 = 0.0
                    _ = v27
                case 2147483647:
                    if v5 {
                        if (0 >= 42) {
                            v28 := v4[348]
                            v28 = v28
                            _ = v28
                        }
                        var v29 string = "hello"
                        _ = v29
                    }
                }
                _ = fi26
            }
            v30 := GenS2{F0: 32767, F1: 127}
            _ = v25
            _ = v30
        } else {
            if (2 < 127) {
                var v31 bool = v5
                switch "" {
                case "hello":
                    v32 := GenS6{F0: (4294967295 == 326), F1: v3, F2: v31}
                    _ = v32
                default:
                    for fi33 := 0; fi33 < 17; fi33++ {
                        var v34 uint32 = 2147483647
                        if (v31 || v31) {
                            var v35 bool = v31
                            var v36 int8 = 27
                            _ = v35
                            _ = v36
                        } else {
                            var v37 uint64 = v24
                            _ = v37
                        }
                        _ = fi33
                        _ = v34
                    }
                }
                _ = v31
            } else {
                if (-32768 <= 1000) {
                    switch ("" + "\n\t") {
                    case "":
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            var v38 string = "foo bar"
                            v39 := GenS2{F0: 1, F1: 248}
                            switch "" {
                            case "a":
                                v40 := GenS3{F0: uint(v24), F1: 4294967295}
                                _ = v40
                            default:
                                v41 := GenS3{F0: 0, F1: v24}
                                _ = v41
                            }
                            _ = v38
                            _ = v39
                        }
                    }
                    delete(v4, 2)
                    switch "a" {
                    case "":
                        if v5 {
                            var v42 int64 = 9223372036854775807
                            v3 = -1
                            v5 = (false && (!v5))
                            _ = v42
                        } else {
                            var v43 uint64 = v24
                            var v44 uint16 = 42
                            _ = v43
                            _ = v44
                        }
                    case "foo bar":
                        v3 = v3
                    default:
                        v12 += (v3 | -1)
                    }
                } else {
                    v12 -= v12
                    v45 := GenS7{F0: v24}
                    v5 = (16 == 42)
                    _ = v45
                }
            }
        }
        _ = fi23
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
