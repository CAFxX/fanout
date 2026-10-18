package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int; F2 uint16 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 int16; F1 int; F2 int16 }

func main() {
    v1 := []int{128}
    var v2 int32 = -2147483648
    var v3 int64 = -9223372036854775808
    v1 = append(v1, 1)
    var v4 uint32 = 4294967295
    v5 := GenS1{F0: 32767, F1: -697, F2: ("\n\t" + (("a" + "") + "\n\t"))}
    if true {
        fmt.Println(v3)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 := GenS2{F0: v4}
    switch "a" {
    case "a":
        for fi7 := 0; fi7 < 3; fi7++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch "\n\t" {
            case "\n\t":
                v3 -= v3
            case "":
                var v8 int8 = -1
                _ = v8
            case "hello":
                var v9 float32 = (3.14159 + 1e10)
                _ = v9
            }
            for fi10 := 0; fi10 < 18; fi10++ {
                v2 = v2
                var v11 int16 = 1000
                _ = fi10
                _ = v11
            }
            _ = fi7
        }
    default:
        v2 = 1
    }
    v1 = append(v1, 9223372036854775807)
    v3 -= (v3 / ((v3) | 1))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v12 := GenS3{F0: 1, F1: -349, F2: 42}
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 float32 = -2.25
        fmt.Println(v2, v14)
        _ = fi13
        _ = v14
    }
    if false {
        v4 -= v4
        for fi15 := 0; fi15 < 2; fi15++ {
            for fi16 := 0; fi16 < 17; fi16++ {
                var v17 string = ("a" + (("hello" + ("\n\t" + "foo bar")) + "foo bar"))
                if false {
                    if (fi16 != 1) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if (0 >= v4) {
                            switch -9223372036854775808 {
                            case -1:
                                v18 := GenS1{F0: -286, F1: -32768, F2: v17}
                                _ = v18
                            default:
                                fmt.Println(v4, v17, fi15)
                            }
                            if (2 != 469) {
                                _ = len(v1)
                                v4 = v4
                                if (757 >= 2) {
                                    v4 -= (v4 ^ v4)
                                    switch fi16 {
                                    case -669:
                                        if (!((!(28 < 21)) && (3.14159 != 0.1))) {
                                            var v19 uint16 = 1
                                            v20 := GenS1{F0: 320, F1: 42, F2: v17}
                                            var v21 int64 = 139
                                            _ = v19
                                            _ = v20
                                            _ = v21
                                        } else {
                                            var v22 uint16 = 42
                                            _ = v22
                                        }
                                    case 127:
                                        if len(v1) > 0 {
                                            v1[0] = (fi15 >> ((18446744073709551615) & 63))
                                        }
                                    default:
                                        switch 159 {
                                        case 32767:
                                            if (uint16(v4) < 32767) {
                                                v1 = append(v1, (fi16 + fi16))
                                                var v23 int16 = 642
                                                _ = v23
                                            } else {
                                                var v24 uint = 255
                                                _ = v24
                                            }
                                        case 42:
                                            fmt.Println(v3, v17, v4)
                                        default:
                                            v3 = (v3 << ((1) & 63))
                                        }
                                    }
                                    v17 = (v17 + ("hello" + v17))
                                } else {
                                    switch (fi16 * (fi15 | (^(fi16)))) {
                                    case 32767:
                                        v3 -= v3
                                    case 4294967295:
                                        if (-1 <= -291) {
                                            if true {
                                                fmt.Println(v2, v2, v3)
                                                v1 = append(v1, fi16)
                                                if false {
                                                    v25 := GenS3{F0: 255, F1: fi15, F2: 341}
                                                    v26 := GenS3{F0: 173, F1: fi16, F2: 801}
                                                    _ = v25
                                                    _ = v26
                                                }
                                            }
                                            switch 9223372036854775807 {
                                            case -9223372036854775808:
                                                if (v17 <= (v17 + v17)) {
                                                    v3 += 1
                                                    _ = len(v1)
                                                    var v27 uint32 = (v4 - (v4 + (v4 | v4)))
                                                    fmt.Println(fi15, v17, v2)
                                                    _ = v27
                                                }
                                            default:
                                                fmt.Println(fi16)
                                            }
                                        }
                                        switch v17 {
                                        case "a":
                                            fmt.Println(v4, fi16)
                                        case "":
                                            var v28 uint = 615
                                            _ = v28
                                        case "foo bar":
                                            switch (^(fi15)) {
                                            case -2147483648:
                                                v4 += v4
                                            default:
                                                if (18446744073709551615 > 2) {
                                                    v29 := GenS4{F0: 4294967295, F1: int64(v2)}
                                                    var v30 int64 = (v3 + v3)
                                                    _ = v29
                                                    _ = v30
                                                }
                                            }
                                            _ = len(v1)
                                        default:
                                            switch (fi16 & 4294967295) {
                                            case 42:
                                                v3 -= 0
                                            }
                                        }
                                    }
                                }
                            } else {
                                var v31 int64 = v3
                                _ = v31
                            }
                        }
                        if len(v1) > 0 {
                            v1[0] = 1
                        }
                    } else {
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v32 := GenS4{F0: 4294967295, F1: v3}
                            var v33 float32 = -2.25
                            var v34 bool = (false && false)
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                    }
                } else {
                    var v35 string = "\n\t"
                    if false {
                        v36 := GenS5{F0: -32768, F1: fi15, F2: -877}
                        v37 := GenS1{F0: 1000, F1: 0, F2: v35}
                        var v38 uint64 = 1
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    _ = v35
                }
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
