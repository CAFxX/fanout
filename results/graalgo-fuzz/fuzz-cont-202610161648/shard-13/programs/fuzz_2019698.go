package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 int32; F2 int32 }
func genFn0(p0 uint32, p1 int64) bool {
    p0 += 65535
    v7 := GenS1{F0: 128, F1: -1}
    p1 -= -9223372036854775808
    v8 := GenS2{F0: "\n\t"}
    var v9 uint8 = 0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (int(p1) > 32768)
}
func genFn1() int32 {
    var v10 int64 = 677
    v10 = v10
    v10 = 879
    _ = v10
    return 523
}

func main() {
    var v1 int32 = 1
    v2 := []uint32{2, 898}
    var v3 string = "hello"
    v4 := make(map[int]string)
    var v5 uint8 = 0
    v6 := []int{2147483647, 2147483647}
    v11 := GenS1{F0: (v5 >> ((65535) & 63)), F1: int(v5)}
    var v12 uint = 0
    _ = len(v4)
    v13 := v4[120]
    v14 := v4[-9223372036854775808]
    var v15 int16 = 1000
    v12 = v12
    var v16 uint64 = 2
    switch v3 {
    case "a":
        var v17 int64 = -1
        _ = v17
    case "\n\t":
        fmt.Println(v1, v1, v16)
    case "foo bar":
        fmt.Println(v14, v15, v16)
    }
    v18 := GenS3{F0: 0}
    for fi19 := 0; fi19 < 3; fi19++ {
        v20 := GenS2{F0: ""}
        _ = fi19
        _ = v20
    }
    v1 -= (v1 ^ v1)
    for fi21 := 0; fi21 < 7; fi21++ {
        switch (fi21 + (^(fi21))) {
        case -1:
            v22 := GenS1{F0: uint8(v16), F1: 128}
            _ = v22
        case -700:
            if false {
                var v23 uint16 = 2
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v23
            } else {
                var v24 uint32 = 158
                v14 = "foo bar"
                v25 := GenS4{F0: v12, F1: 2147483647, F2: genFn1()}
                _ = v24
                _ = v25
            }
        case 396:
            var v26 uint = v12
            _ = v26
        default:
            var v27 int8 = 0
            _ = v27
        }
        v28 := v4[(fi21 << (((v12 ^ (v12 ^ 344))) & 63))]
        var v29 string = v28
        _ = fi21
        _ = v28
        _ = v29
    }
    _ = len(v6)
    switch "" {
    case "foo bar":
        v4[2147483647] = v14
    case "a":
        for fi30 := 0; fi30 < 13; fi30++ {
            if ((1.5 + (1.5 - 1.5)) != (((0.0 * 0.0) - 0.1) * (1e10 - (-2.25 - float64(1.5))))) {
                var v31 int64 = 1
                _ = v31
            } else {
                v4[9223372036854775807] = "a"
                v2 = append(v2, 42)
                var v32 string = v3
                _ = v32
            }
            for fi33 := 0; fi33 < 11; fi33++ {
                if (!true) {
                    switch v13 {
                    case "\n\t":
                        fmt.Println(v5, v15)
                    case "a":
                        var v34 uint64 = (v16 & v16)
                        _ = v34
                    }
                    v3 = "hello"
                }
                _ = fi33
            }
            var v35 int16 = (v15 % ((-1) | 1))
            _ = fi30
            _ = v35
        }
    case "hello":
        switch "" {
        case "":
            for fi36 := 0; fi36 < 1; fi36++ {
                var v37 int = fi36
                _ = fi36
                _ = v37
            }
        default:
            delete(v4, 42)
        }
    default:
        if true {
            v16 -= (v16 & (v16 ^ 0))
            v13 = v3
        } else {
            if (v16 == (v16 / (((v16 ^ 595)) | 1))) {
                fmt.Println(v1, v14)
                v2 = append(v2, 65535)
            }
            for fi38 := 0; fi38 < 17; fi38++ {
                switch (^(fi38)) {
                case 42:
                    _ = len(v4)
                default:
                    v39 := GenS4{F0: (^(v12)), F1: int32(v15), F2: v1}
                    _ = v39
                }
                _ = fi38
            }
            var v40 int64 = 9223372036854775807
            _ = v40
        }
    }
    var v41 float64 = ((0.1 * -2.25) + float64(float32((3.14159 - -2.25))))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
