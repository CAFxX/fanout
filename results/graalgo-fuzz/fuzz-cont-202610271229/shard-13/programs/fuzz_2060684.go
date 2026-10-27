package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 int32; F1 int16 }

func main() {
    var v1 string = "hello"
    var v2 uint8 = 61
    var v3 int = 128
    v4 := make(map[string]bool)
    v5 := make(map[uint32]uint32)
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 int64 = 42
        _ = fi6
        _ = v7
    }
    if (2 < 0) {
        v8 := GenS1{F0: (3.14159 == 1e10)}
        switch "a" {
        case "":
            fmt.Println(v3, v3, v2)
        }
        for fi9 := 0; fi9 < 4; fi9++ {
            v1 = "\n\t"
            v3 += (v3 / ((fi9) | 1))
            _ = fi9
        }
        _ = v8
    } else {
        v2 = v2
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := GenS1{F0: false}
            fmt.Println(fi10, v2, v2)
            v12 := GenS2{F0: 42, F1: (v2 + v2)}
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    var v13 int = v3
    var v14 uint16 = 42
    if true {
        for fi15 := 0; fi15 < 19; fi15++ {
            if false {
                v3 += (fi15 + 9223372036854775807)
                switch ("a" + "") {
                case "foo bar":
                    var v16 uint32 = 1
                    _ = v16
                default:
                    var v17 uint16 = v14
                    _ = v17
                }
                var v18 int16 = 1000
                _ = v18
            } else {
                v19 := []int{4294967295, (fi15 ^ (-(fi15))), (v13 >> ((65535) & 63)), 128}
                v1 = "a"
                _ = v19
            }
            v20 := GenS2{F0: v14, F1: (v2 << ((18446744073709551615) & 63))}
            for fi21 := 0; fi21 < 14; fi21++ {
                if (65535 >= 0) {
                    v22 := GenS1{F0: true}
                    _ = v22
                } else {
                    switch (v1 + "hello") {
                    case "hello":
                        fmt.Println(v13)
                        v1 = "a"
                        if false {
                            v23 := GenS1{F0: false}
                            var v24 float32 = (0.1 - 3.14159)
                            _ = v23
                            _ = v24
                        }
                    case "\n\t":
                        switch (v13 % ((v13) | 1)) {
                        case -194:
                            v2 = v2
                        case 1:
                            fmt.Println(v2, v13, v3)
                        }
                    default:
                        _ = len(v5)
                    }
                    fmt.Println(fi15, v1, v14)
                    v25 := GenS3{F0: v1, F1: 4294967295}
                    _ = v25
                }
                v26 := GenS4{F0: 42, F1: 4294967295}
                _ = fi21
                _ = v26
            }
            _ = fi15
            _ = v20
        }
        fmt.Println(v2)
    } else {
        v27 := GenS5{F0: 192, F1: 1}
        var v28 int32 = -2147483648
        var v29 int32 = v28
        _ = v27
        _ = v28
        _ = v29
    }
    var v30 uint64 = 1
    if (0.0 == float64(float32(((3.14159 / 1.0) - 1e10)))) {
        v31 := v5[0]
        for fi32 := 0; fi32 < 19; fi32++ {
            var v33 int8 = 127
            _ = fi32
            _ = v33
        }
        _ = v31
    }
    var v34 uint32 = 2147483647
    if ((!(1 != 26)) || true) {
        v14 = v14
        var v35 float32 = 0.1
        _ = v35
    } else {
        v36 := []int{}
        _ = v36
    }
    var v37 uint8 = (v2 % (((v2 * v2)) | 1))
    var v38 uint32 = v34
    v3 = -494
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
