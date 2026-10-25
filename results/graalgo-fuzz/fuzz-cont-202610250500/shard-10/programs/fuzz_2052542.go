package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 bool }

func main() {
    var v1 string = ("" + ("a" + ""))
    var v2 int8 = -1
    v3 := make(map[int]uint32)
    var v4 int16 = 42
    for fi5 := 0; fi5 < 17; fi5++ {
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 uint16 = 65535
            v3[4294967295] = 42
            var v8 string = "foo bar"
            _ = fi6
            _ = v7
            _ = v8
        }
        if false {
            switch (-(fi5)) {
            case 708:
                if (!false) {
                    var v9 int8 = (v2 + v2)
                    for fi10 := 0; fi10 < 11; fi10++ {
                        if (65535 == 584) {
                            if true {
                                if false {
                                    var v11 int8 = -63
                                    var v12 int16 = (v4 | (v4 << ((18446744073709551615) & 63)))
                                    v13 := GenS1{F0: "\n\t", F1: 679}
                                    switch fi5 {
                                    case -2147483648:
                                        var v14 uint64 = 18446744073709551615
                                        _ = v14
                                    default:
                                        if true {
                                            if false {
                                                if (18446744073709551615 < 18446744073709551615) {
                                                    v12 = (v4 - 42)
                                                    var v15 int32 = 0
                                                    _ = v15
                                                }
                                                var v16 float64 = 3.14159
                                                _ = v16
                                            } else {
                                                if ((true || (!true)) || (!true)) {
                                                    v17 := GenS2{F0: 4294967295}
                                                    switch 128 {
                                                    case -166:
                                                        switch v1 {
                                                        case "":
                                                            var v18 int64 = -9223372036854775808
                                                            _ = v18
                                                        case "hello":
                                                            v2 = -1
                                                        }
                                                    case 0:
                                                        v11 += 8
                                                    }
                                                    v9 += (v2 + (-(v2)))
                                                    _ = v17
                                                }
                                                switch v1 {
                                                case "hello":
                                                    var v19 uint8 = 1
                                                    _ = v19
                                                }
                                            }
                                        }
                                    }
                                    _ = v11
                                    _ = v12
                                    _ = v13
                                } else {
                                    var v20 uint32 = 4294967295
                                    if false {
                                        var v21 int8 = -119
                                        v1 = v1
                                        _ = v21
                                    } else {
                                        if false {
                                            if true {
                                                _ = len(v3)
                                                if (v1 < "\n\t") {
                                                    var v22 int32 = -1
                                                    var v23 int32 = (v22 % (((v22 ^ (-(v22)))) | 1))
                                                    v4 -= v4
                                                    _ = v22
                                                    _ = v23
                                                }
                                                v24 := GenS3{F0: uint(v20), F1: ("\n\t" < (v1 + "\n\t"))}
                                                _ = v24
                                            }
                                        } else {
                                            v25 := GenS3{F0: 65535, F1: false}
                                            _ = v25
                                        }
                                    }
                                    _ = v20
                                }
                            } else {
                                switch fi5 {
                                case 4294967295:
                                    var v26 uint = 65535
                                    _ = v26
                                case 128:
                                    v2 -= (v2 / ((11) | 1))
                                    v27 := GenS3{F0: 65535, F1: true}
                                    v3[fi5] = 163
                                    _ = v27
                                }
                                var v28 int32 = 895
                                switch (("\n\t" + (v1 + v1)) + v1) {
                                case "a":
                                    switch v1 {
                                    case "\n\t":
                                        v29 := []int{128, (^(fi5)), 1, fi5}
                                        _ = v29
                                    case "hello":
                                        var v30 int = fi10
                                        _ = v30
                                    default:
                                        if false {
                                            var v31 int16 = v4
                                            _ = v31
                                        }
                                        var v32 uint = 0
                                        _ = v32
                                    }
                                case "\n\t":
                                    var v33 uint = 241
                                    _ = v33
                                case "foo bar":
                                    v3[fi5] = 413
                                default:
                                    var v34 int8 = v9
                                    _ = v34
                                }
                                _ = v28
                            }
                        } else {
                            var v35 int16 = (v4 & v4)
                            _ = v35
                        }
                        fmt.Println(fi5)
                        _ = fi10
                    }
                    var v36 string = ((v1 + "") + v1)
                    _ = v36
                    _ = v9
                }
            case 816:
                v37 := GenS1{F0: "", F1: 857}
                _ = v37
            default:
                var v38 uint8 = 37
                _ = v38
            }
            var v39 int = (-(fi5))
            v2 = 16
            _ = v39
        }
        delete(v3, (fi5 % ((fi5) | 1)))
        _ = fi5
    }
    var v40 uint64 = 948
    v4 = 0
    fmt.Println(v1, v2, v2)
    for fi41 := 0; fi41 < 11; fi41++ {
        var v42 float32 = (0.1 / 1e10)
        _ = fi41
        _ = v42
    }
    v3[9223372036854775807] = 0
    v4 -= 1
    for fi43 := 0; fi43 < 2; fi43++ {
        v40 += (v40 & v40)
        var v44 float64 = -2.25
        v1 = ""
        _ = fi43
        _ = v44
    }
    delete(v3, 547)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
}
