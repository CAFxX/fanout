package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint8; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 uint8; F2 uint8 }

func main() {
    var v1 uint32 = 652
    var v2 int = 103
    var v3 uint64 = 2
    var v4 string = "\n\t"
    var v5 uint64 = 42
    v4 = v4
    for fi6 := 0; fi6 < 11; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            v8 := []int{(fi6 ^ (fi7 - fi6)), 270, (fi7 & (v2 / ((v2) | 1)))}
            var v9 uint64 = (v3 & (v5 % (((v3 ^ 1)) | 1)))
            if (!true) {
                var v10 float64 = float64(3.14159)
                if false {
                    v11 := GenS1{F0: 127, F1: -71, F2: 4294967295}
                    fmt.Println(fi7)
                    var v12 int64 = 0
                    _ = v11
                    _ = v12
                } else {
                    v8 = append(v8, fi7)
                    var v13 float64 = float64((((float32(v10) / float32(v10)) - (0.0 + 1e10)) / (0.1 + 0.0)))
                    fmt.Println(v4, v10)
                    _ = v13
                }
                _ = v10
            } else {
                if false {
                    v8 = append(v8, (v2 & (^(v2))))
                    v9 += v3
                    fmt.Println(v1)
                }
                v14 := GenS2{F0: false, F1: 2, F2: -661}
                _ = v14
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        var v15 int16 = 1000
        _ = fi6
        _ = v15
    }
    v16 := []int{(v2 % ((v2) | 1))}
    v17 := GenS2{F0: (true && (42 <= 255)), F1: 243, F2: -9223372036854775808}
    var v18 float64 = 0.0
    v3 -= 42
    var v19 uint16 = 42
    for fi20 := 0; fi20 < 9; fi20++ {
        for fi21 := 0; fi21 < 17; fi21++ {
            if (false || (127 < 101)) {
                fmt.Println(v5, v2)
                if false {
                    var v22 uint16 = 42
                    var v23 int8 = 1
                    _ = v22
                    _ = v23
                }
                switch v4 {
                case "":
                    v3 = (v3 * v5)
                case "\n\t":
                    v19 += 985
                default:
                    switch v2 {
                    case -361:
                        var v24 uint = 65535
                        _ = v24
                    case -9223372036854775808:
                        fmt.Println(v2)
                    default:
                        var v25 uint64 = (v3 << ((4294967295) & 63))
                        _ = v25
                    }
                }
                var v26 uint16 = v19
                _ = v26
            } else {
                var v27 uint16 = v19
                if (128 > 255) {
                    var v28 int16 = 1000
                    v4 = ""
                    _ = v28
                } else {
                    v16 = append(v16, -31)
                }
                switch v4 {
                case "a":
                    var v29 uint = 465
                    _ = v29
                default:
                    v30 := GenS3{F0: 0}
                    _ = v30
                }
                switch (fi21 >> ((1) & 63)) {
                case 42:
                    v31 := GenS3{F0: 127}
                    _ = v31
                }
                v18 = 0.1
                fmt.Println(v3, v4, v4)
                _ = v27
            }
            _ = fi21
        }
        _ = fi20
    }
    v2 -= int(v1)
    var v32 uint = 42
    if (-550 >= 0) {
        v33 := GenS3{F0: 127}
        for fi34 := 0; fi34 < 5; fi34++ {
            v32 = 1
            _ = fi34
        }
        _ = v33
    }
    switch (v2 % ((v2) | 1)) {
    case 32767:
        v19 = 32767
    case -1:
        var v35 int64 = 0
        _ = v35
    case 799:
        v36 := GenS4{F0: false, F1: 1, F2: 211}
        _ = v36
    default:
        var v37 uint = (v32 * 42)
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
