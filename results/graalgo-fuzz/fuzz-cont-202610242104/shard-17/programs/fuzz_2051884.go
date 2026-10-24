package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int; F1 uint }

func main() {
    var v1 uint8 = 255
    v2 := []string{"", "\n\t"}
    v3 := make(map[uint32]bool)
    var v4 string = "hello"
    var v5 int8 = 127
    if (42 < 250) {
        _ = len(v3)
        if ((1000 >= 1) || ((-2.25 - -2.25) != (float32(1e10) + 1.5))) {
            for fi6 := 0; fi6 < 16; fi6++ {
                if true {
                    if true {
                        var v7 int64 = 42
                        v4 = v4
                        _ = v7
                    } else {
                        var v8 int = 9223372036854775807
                        switch (fi6 / (((-(fi6))) | 1)) {
                        case 4294967295:
                            v9 := GenS1{F0: 139, F1: 795, F2: 1}
                            _ = v9
                        case 292:
                            v8 = 9223372036854775807
                        default:
                            fmt.Println(v4)
                        }
                        for fi10 := 0; fi10 < 15; fi10++ {
                            var v11 float64 = 0.0
                            _ = fi10
                            _ = v11
                        }
                        _ = v8
                    }
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                    fmt.Println(v1, fi6)
                }
                var v12 bool = (((1e10 - 1.5) + 0.0) > 0.1)
                _ = fi6
                _ = v12
            }
            if len(v2) > 0 {
                v2[0] = (v4 + v4)
            }
            fmt.Println(v4, v1)
        } else {
            var v13 int16 = 1000
            var v14 float32 = (float32(0.1) * ((0.0 * 0.1) + (0.0 + 1e10)))
            _ = v13
            _ = v14
        }
    }
    var v15 int8 = v5
    v16 := GenS2{F0: 32767, F1: 18446744073709551615, F2: v5}
    if false {
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 uint16 = 1
            if false {
                fmt.Println(v1, v15)
            }
            _ = len(v3)
            _ = fi17
            _ = v18
        }
    } else {
        _ = len(v3)
        if true {
            v5 = 127
            var v19 int64 = 1
            _ = v19
        }
        for fi20 := 0; fi20 < 6; fi20++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi20
        }
    }
    if true {
        var v21 bool = false
        switch v4 {
        case "a":
            if (65535 != 305) {
                switch v4 {
                case "":
                    if (!(float32((3.14159 * (float64(-2.25) + 0.0))) != float32(1e10))) {
                        v22 := GenS3{F0: 32767}
                        _ = v22
                    }
                case "hello":
                    v1 -= (v1 << ((632) & 63))
                }
                if v21 {
                    var v23 string = ("hello" + v4)
                    for fi24 := 0; fi24 < 7; fi24++ {
                        _ = len(v3)
                        v25 := GenS2{F0: 0, F1: 0, F2: 0}
                        if (0 == 487) {
                            v26 := GenS4{F0: 397}
                            if v21 {
                                for fi27 := 0; fi27 < 12; fi27++ {
                                    var v28 uint32 = 42
                                    var v29 uint32 = (v28 >> ((1) & 63))
                                    v4 = v4
                                    _ = fi27
                                    _ = v28
                                    _ = v29
                                }
                                var v30 uint64 = 42
                                _ = v30
                            } else {
                                v31 := GenS1{F0: -1, F1: -889, F2: 474}
                                if (0 == 18446744073709551615) {
                                    fmt.Println(v5, v23)
                                    v2 = append(v2, v4)
                                }
                                _ = v31
                            }
                            v15 -= v15
                            _ = v26
                        }
                        _ = fi24
                        _ = v25
                    }
                    _ = v23
                } else {
                    v32 := v3[384]
                    for fi33 := 0; fi33 < 6; fi33++ {
                        delete(v3, 981)
                        _ = fi33
                    }
                    v34 := GenS5{F0: 42, F1: 1}
                    _ = v32
                    _ = v34
                }
                v35 := GenS3{F0: 1}
                _ = v35
            }
        }
        for fi36 := 0; fi36 < 5; fi36++ {
            v5 = (v5 * v15)
            var v37 bool = v21
            for fi38 := 0; fi38 < 13; fi38++ {
                v15 = -1
                v4 = v4
                v2 = append(v2, "\n\t")
                _ = fi38
            }
            _ = fi36
            _ = v37
        }
        _ = v21
    } else {
        var v39 int8 = (-(v15))
        for fi40 := 0; fi40 < 5; fi40++ {
            v5 -= -68
            v4 = ""
            _ = fi40
        }
        _ = v39
    }
    var v41 uint16 = 255
    v41 = v41
    var v42 uint16 = (^(v41))
    v4 = ""
    delete(v3, 51)
    v43 := GenS5{F0: 32767, F1: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
