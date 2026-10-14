package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 bool; F2 uint8 }

func main() {
    v1 := []uint32{965, 1}
    var v2 uint16 = 255
    v3 := make(map[uint32]uint32)
    var v4 int32 = 485
    var v5 float64 = -2.25
    var v6 uint16 = v2
    v7 := GenS1{F0: 42}
    v2 = 2
    var v8 uint32 = 4294967295
    var v9 int8 = 65
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            if (!(!false)) {
                v6 -= (^(v6))
            }
            v8 = v8
            var v12 uint64 = 1
            _ = fi11
            _ = v12
        }
        v13 := GenS2{F0: false}
        _ = fi10
        _ = v13
    }
    v14 := GenS1{F0: 9223372036854775807}
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 uint8 = 255
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        if ("\n\t" < ("" + "\n\t")) {
            var v18 uint64 = 42
            fmt.Println(v2, v4, v5)
            _ = len(v1)
            _ = v18
        } else {
            if (true || (false || (128 <= 42))) {
                if (((!false) && ((v8 ^ 987) <= 860)) || ((fi17 | (fi17 & fi17)) >= (fi17 % ((fi17) | 1)))) {
                    switch "\n\t" {
                    case "foo bar":
                        v2 -= v6
                    }
                    switch (^(fi17)) {
                    case 9223372036854775807:
                        v5 += v5
                    case -1:
                        var v19 uint8 = 128
                        _ = v19
                    default:
                        v20 := GenS2{F0: true}
                        _ = v20
                    }
                } else {
                    if true {
                        v21 := GenS1{F0: 4294967295}
                        var v22 uint16 = 838
                        _ = len(v1)
                        _ = v21
                        _ = v22
                    }
                    switch 42 {
                    case 32767:
                        fmt.Println(fi17, v8, fi17)
                    }
                    for fi23 := 0; fi23 < 17; fi23++ {
                        var v24 uint8 = 251
                        v2 -= (^(v2))
                        var v25 int = 128
                        if (((fi17 == (fi17 >> ((144) & 63))) || false) || false) {
                            var v26 int8 = v9
                            v25 = v25
                            _ = v26
                        }
                        if true {
                            v27 := GenS1{F0: v25}
                            var v28 string = "hello"
                            _ = v27
                            _ = v28
                        }
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                }
            }
            if (true || true) {
                v29 := GenS3{F0: v2, F1: (v5 < float64((float32(v5) * -2.25))), F2: 128}
                v8 += v8
                _ = v29
            }
            v30 := GenS2{F0: (false || (1 == 2))}
            _ = v30
        }
        fmt.Println(v6, v2)
        _ = fi17
    }
    v3[65535] = (^(v8))
    var v31 int = -2147483648
    if (!true) {
        if false {
            v1 = append(v1, (v8 & v8))
            var v32 int = v31
            _ = v32
        }
    } else {
        switch v31 {
        case 0:
            var v33 int8 = (v9 - v9)
            _ = v33
        }
        var v34 int = v31
        _ = v34
    }
    fmt.Println(v8)
    v35 := GenS2{F0: false}
    if (0 > 42) {
        v36 := GenS1{F0: v31}
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
