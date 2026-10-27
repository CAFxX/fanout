package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int; F2 int8 }
type GenS4 struct { F0 int8; F1 uint; F2 uint }
type GenS5 struct { F0 uint64; F1 bool }
func genFn0() int64 {
    switch ("hello" + "") {
    case "\n\t":
        switch ("hello" + "") {
        case "":
            var v4 float32 = ((1.5 - 0.0) * ((3.14159 - 0.0) - (((0.0 / 1.5) + 3.14159) * (0.1 * 0.0))))
            _ = v4
        case "hello":
            switch ("\n\t" + "a") {
            case "\n\t":
                for fi5 := 0; fi5 < 16; fi5++ {
                    var v6 string = ""
                    _ = fi5
                    _ = v6
                }
            case "":
                for fi7 := 0; fi7 < 8; fi7++ {
                    var v8 string = "foo bar"
                    _ = fi7
                    _ = v8
                }
            case "a":
                for fi9 := 0; fi9 < 11; fi9++ {
                    v10 := []int{(fi9 ^ fi9), -475, fi9, 0}
                    if (false || false) {
                        v10 = append(v10, -268)
                        if true {
                            var v11 int = fi9
                            _ = v11
                        }
                        fmt.Println(fi9)
                    } else {
                        var v12 string = ""
                        _ = v12
                    }
                    if (false || (207 <= 42)) {
                        for fi13 := 0; fi13 < 1; fi13++ {
                            var v14 int = -442
                            v14 = -953
                            _ = fi13
                            _ = v14
                        }
                        fmt.Println(fi9)
                    } else {
                        fmt.Println(fi9)
                        var v15 string = "a"
                        _ = v15
                    }
                    _ = fi9
                    _ = v10
                }
            default:
                v16 := GenS1{F0: 2147483648}
                _ = v16
            }
        case "\n\t":
            v17 := GenS1{F0: 4294967295}
            _ = v17
        }
    case "a":
        v18 := GenS2{F0: 1}
        _ = v18
    default:
        v19 := []int{-90, -415}
        _ = v19
    }
    var v20 uint16 = 42
    v20 += 0
    _ = v20
    return 0
}

func main() {
    var v1 uint = 2
    v2 := make(map[int]bool)
    var v3 uint32 = 65535
    fmt.Println(v3)
    if false {
        if (-700 <= -32768) {
            if false {
                var v21 int16 = 1000
                if true {
                    v22 := GenS2{F0: 148}
                    switch 503 {
                    case 32768:
                        for fi23 := 0; fi23 < 8; fi23++ {
                            var v24 uint = 255
                            v25 := GenS3{F0: uint16(v1), F1: fi23, F2: 100}
                            for fi26 := 0; fi26 < 4; fi26++ {
                                v27 := GenS1{F0: (v3 / (((^(v3))) | 1))}
                                v24 = v1
                                delete(v2, fi26)
                                _ = fi26
                                _ = v27
                            }
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                    }
                    v28 := GenS3{F0: 930, F1: -948, F2: 40}
                    _ = v22
                    _ = v28
                }
                _ = v21
            } else {
                if (1e10 <= float32(1.5)) {
                    var v29 uint16 = 2
                    _ = v29
                }
                delete(v2, 1)
            }
            switch "hello" {
            case "\n\t":
                v1 += (v1 << (((v1 * (v1 << (((v1 | (^(v1)))) & 63)))) & 63))
            default:
                var v30 uint8 = 94
                _ = v30
            }
        }
        var v31 float32 = -2.25
        switch 700 {
        case 32768:
            v3 = v3
        default:
            var v32 uint32 = v3
            _ = v32
        }
        _ = v31
    } else {
        v33 := []int{42, -672, 127}
        _ = v33
    }
    v34 := GenS1{F0: v3}
    v3 = (v3 >> ((v1) & 63))
    v35 := GenS4{F0: -128, F1: (v1 & v1), F2: (v1 ^ v1)}
    v1 = 132
    v1 = (v1 ^ 0)
    if true {
        var v36 uint64 = 0
        var v37 int32 = 2147483647
        _ = v36
        _ = v37
    }
    v38 := GenS5{F0: 1, F1: true}
    if (!(65535 != 65535)) {
        var v39 int16 = -855
        v39 = (^(v39))
        var v40 int8 = 29
        _ = v39
        _ = v40
    } else {
        var v41 int64 = -9223372036854775808
        _ = v41
    }
    v42 := GenS5{F0: uint64(v3), F1: true}
    v43 := v2[-6]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v42
    _ = v43
}
