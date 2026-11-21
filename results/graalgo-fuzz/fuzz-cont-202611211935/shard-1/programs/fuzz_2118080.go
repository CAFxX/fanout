package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int; F1 uint8; F2 int8 }

func main() {
    var v1 int64 = -1
    v2 := make(map[uint32]string)
    var v3 uint32 = 0
    var v4 float64 = (0.0 - (1e10 * 3.14159))
    var v5 uint = 0
    v6 := make(map[uint32]int)
    fmt.Println(v4)
    if true {
        var v7 int = 128
        _ = v7
    } else {
        switch -240 {
        case 128:
            if true {
                if false {
                    var v8 int8 = -128
                    for fi9 := 0; fi9 < 19; fi9++ {
                        v3 = v3
                        v8 = (v8 * (v8 << (((v5 << (((v5 | 0)) & 63))) & 63)))
                        _ = fi9
                    }
                    _ = v8
                } else {
                    for fi10 := 0; fi10 < 7; fi10++ {
                        var v11 bool = (!((!true) || (true && (255 < v5))))
                        var v12 uint8 = 194
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                }
                switch "a" {
                case "hello":
                    var v13 int32 = -585
                    _ = v13
                case "foo bar":
                    if true {
                        v14 := []int{128, -296, 128}
                        _ = v14
                    } else {
                        _ = len(v6)
                        var v15 int16 = -349
                        _ = v15
                    }
                }
            } else {
                for fi16 := 0; fi16 < 18; fi16++ {
                    v17 := GenS1{F0: -1, F1: 1, F2: 0}
                    var v18 uint64 = 18446744073709551615
                    v18 -= v18
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
        case 0:
            for fi19 := 0; fi19 < 6; fi19++ {
                fmt.Println(v5)
                var v20 uint64 = 18446744073709551615
                if (!false) {
                    if true {
                        switch "" {
                        case "\n\t":
                            switch "hello" {
                            case "foo bar":
                                var v21 int = (-(fi19))
                                _ = v21
                            case "":
                                var v22 float64 = (v4 * (1e10 * v4))
                                _ = v22
                            case "a":
                                for fi23 := 0; fi23 < 3; fi23++ {
                                    var v24 int = fi23
                                    _ = fi23
                                    _ = v24
                                }
                            default:
                                fmt.Println(v3)
                            }
                        }
                    }
                    v25 := GenS2{F0: uint(v3), F1: -46}
                    var v26 float64 = float64(((1.5 - 1.5) / (3.14159 + -2.25)))
                    _ = v25
                    _ = v26
                } else {
                    v27 := v6[(^(v3))]
                    for fi28 := 0; fi28 < 7; fi28++ {
                        if (false || (float32(float64(0.1)) == 3.14159)) {
                            v29 := []int{}
                            v1 -= v1
                            _ = v29
                        }
                        v30 := GenS3{F0: fi19, F1: 255, F2: 100}
                        var v31 uint64 = v20
                        _ = fi28
                        _ = v30
                        _ = v31
                    }
                    _ = v27
                }
                _ = fi19
                _ = v20
            }
        default:
            var v32 uint8 = 2
            _ = v32
        }
        switch 313 {
        case 42:
            for fi33 := 0; fi33 < 11; fi33++ {
                fmt.Println(v5, v3, v4)
                fmt.Println(v3)
                v34 := []int{(fi33 >> (((v5 + (v5 / ((18446744073709551615) | 1)))) & 63)), fi33, fi33, (fi33 + (fi33 + fi33))}
                _ = fi33
                _ = v34
            }
        case 32767:
            v4 = -2.25
        case -9223372036854775808:
            var v35 string = "hello"
            _ = v35
        default:
            var v36 float32 = float32(float64((0.0 + 3.14159)))
            _ = v36
        }
        v37 := []int{127, -9223372036854775808, -2147483648}
        _ = v37
    }
    v38 := GenS1{F0: 1, F1: 0, F2: (v3 >> ((400) & 63))}
    v39 := GenS3{F0: -162, F1: 1, F2: -1}
    v1 = v1
    v40 := GenS1{F0: 1, F1: 255, F2: uint32(v5)}
    for fi41 := 0; fi41 < 7; fi41++ {
        switch fi41 {
        case -1:
            _ = len(v6)
        }
        var v42 bool = ((0.0 / -2.25) < 0.1)
        _ = fi41
        _ = v42
    }
    var v43 string = ("" + "")
    var v44 int32 = -623
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
