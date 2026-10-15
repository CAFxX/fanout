package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int16; F1 int32; F2 int16 }
type GenS4 struct { F0 int8; F1 int64; F2 uint }
type GenS5 struct { F0 bool }
func genFn0() uint64 {
    v5 := []int{301}
    switch "hello" {
    case "":
        var v6 float32 = 1.5
        _ = v6
    case "hello":
        v7 := GenS1{F0: 0, F1: 255}
        _ = v7
    }
    _ = v5
    return 180
}

func main() {
    var v1 string = ""
    v2 := []int{-25, 4294967295, 42}
    var v3 int32 = -2147483648
    var v4 int64 = 1
    v4 = 551
    var v8 int32 = (v3 >> ((42) & 63))
    switch 24 {
    case -876:
        v2 = append(v2, 4294967295)
    case 2147483647:
        v4 += (v4 % (((v4 << ((4294967295) & 63))) | 1))
    default:
        for fi9 := 0; fi9 < 16; fi9++ {
            _ = len(v2)
            _ = fi9
        }
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        fmt.Println(v3, v8)
        if (2147483647 == 2) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            if (42 < 118) {
                for fi11 := 0; fi11 < 9; fi11++ {
                    var v12 bool = true
                    _ = fi11
                    _ = v12
                }
                if false {
                    fmt.Println(v8, v3)
                }
                var v13 int64 = (v4 / ((v4) | 1))
                _ = v13
            } else {
                fmt.Println(v4)
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    v14 := GenS2{F0: (v3 | 0), F1: ((v1 + v1) + v1)}
                    _ = v14
                }
            }
        }
        var v15 int64 = v4
        _ = fi10
        _ = v15
    }
    switch (v1 + v1) {
    case "hello":
        switch 1 {
        case -2147483648:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case -177:
            if (18446744073709551615 >= 786) {
                if (2 <= 0) {
                    var v16 float32 = -2.25
                    v2 = append(v2, 1)
                    _ = v16
                } else {
                    v3 -= v8
                    v1 = (v1 + (v1 + v1))
                    v4 = (v4 | 9223372036854775807)
                }
                v17 := GenS2{F0: int32(v4), F1: ("\n\t" + "foo bar")}
                v4 = v4
                _ = v17
            }
        case 32768:
            for fi18 := 0; fi18 < 2; fi18++ {
                if (128 < 10) {
                    for fi19 := 0; fi19 < 6; fi19++ {
                        v20 := GenS2{F0: v3, F1: v1}
                        v2 = append(v2, 32768)
                        _ = fi19
                        _ = v20
                    }
                    switch -1 {
                    case 128:
                        var v21 uint32 = 2147483648
                        _ = v21
                    case 32768:
                        if true {
                            v22 := GenS1{F0: (v3 & v8), F1: 545}
                            _ = v22
                        } else {
                            var v23 uint = 65535
                            _ = len(v2)
                            _ = v23
                        }
                    default:
                        v24 := GenS3{F0: 1000, F1: v3, F2: -1}
                        _ = v24
                    }
                    var v25 int8 = -128
                    _ = v25
                } else {
                    switch v1 {
                    case "a":
                        fmt.Println(v8)
                    case "foo bar":
                        fmt.Println(v3, v3)
                    case "\n\t":
                        v8 = (v3 << ((602) & 63))
                    }
                    var v26 uint = 255
                    _ = v26
                }
                switch ("" + "hello") {
                case "hello":
                    var v27 float32 = float32(float64(float32(1.5)))
                    _ = v27
                case "\n\t":
                    fmt.Println(v1)
                }
                _ = fi18
            }
        }
    case "\n\t":
        v28 := GenS4{F0: int8(v4), F1: v4, F2: 42}
        _ = v28
    case "":
        if false {
            var v29 int = 2147483647
            _ = v29
        } else {
            if false {
                if false {
                    v4 -= 42
                    v4 += 0
                    if false {
                        v30 := GenS3{F0: -1, F1: v8, F2: 1}
                        switch 127 {
                        case -2147483648:
                            v2 = append(v2, -268)
                        case 4294967295:
                            if (1e10 <= 0.1) {
                                if true {
                                    var v31 string = (v1 + v1)
                                    fmt.Println(v3)
                                    if false {
                                        v3 = v8
                                        v31 = "hello"
                                        var v32 int16 = 0
                                        _ = v32
                                    } else {
                                        v33 := GenS3{F0: 1, F1: (v8 ^ (^(v8))), F2: 741}
                                        _ = v33
                                    }
                                    _ = v31
                                }
                                v1 = v1
                            }
                        }
                        _ = v30
                    }
                }
                fmt.Println(v3, v4)
                if false {
                    for fi34 := 0; fi34 < 1; fi34++ {
                        v2 = append(v2, 2147483647)
                        if len(v2) > 0 {
                            v2[0] = fi34
                        }
                        _ = fi34
                    }
                }
            }
            v35 := GenS5{F0: false}
            var v36 uint16 = 255
            _ = v35
            _ = v36
        }
    default:
        fmt.Println(v4, v8, v3)
    }
    switch ((v1 + v1) + ("hello" + v1)) {
    case "":
        v3 = -928
    }
    for fi37 := 0; fi37 < 12; fi37++ {
        v1 = "hello"
        var v38 float64 = -2.25
        _ = fi37
        _ = v38
    }
    if (1 == 127) {
        v39 := GenS1{F0: (v3 * v3), F1: 255}
        v4 += 0
        _ = v39
    } else {
        v40 := GenS3{F0: 0, F1: (v3 | v8), F2: -180}
        var v41 uint32 = 4294967295
        _ = v40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
