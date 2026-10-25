package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint8; F2 bool }
func genFn0() int8 {
    switch 1 {
    case 9223372036854775807:
        var v5 uint32 = 1
        _ = v5
    default:
        var v6 uint32 = 65535
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: (3.14159 > 0.0), F1: "\n\t"}
        var v9 uint16 = 439
        fmt.Println(v9)
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: (-932 <= 0), F1: "foo bar"}
    var v11 uint64 = 79
    for fi12 := 0; fi12 < 16; fi12++ {
        v11 = (v11 % ((v11) | 1))
        if (117 != 1) {
            if true {
                switch "a" {
                case "foo bar":
                    if ("" != ("hello" + "")) {
                        var v13 float32 = (-2.25 - (0.0 - 1.5))
                        if (true || true) {
                            var v14 string = ("foo bar" + "foo bar")
                            _ = v14
                        } else {
                            v11 = v11
                            v13 = (float32(float64(1e10)) * v13)
                            fmt.Println(v13, fi12)
                        }
                        for fi15 := 0; fi15 < 16; fi15++ {
                            var v16 float64 = (((1.5 + ((3.14159 * (((((1.5 - -2.25) * 0.0) + -2.25) * (1.5 - -2.25)) + (1.5 - 1e10))) * (-2.25 - 3.14159))) / (1e10 - -2.25)) / 3.14159)
                            _ = fi15
                            _ = v16
                        }
                        _ = v13
                    } else {
                        var v17 int64 = -178
                        _ = v17
                    }
                default:
                    fmt.Println(fi12)
                }
                v11 += v11
                v18 := GenS1{F0: false, F1: "\n\t"}
                _ = v18
            } else {
                var v19 uint8 = 128
                var v20 uint64 = v11
                _ = v19
                _ = v20
            }
            v21 := GenS2{F0: 104, F1: 42, F2: 905}
            _ = v21
        } else {
            v11 = (v11 ^ v11)
            if (fi12 >= fi12) {
                fmt.Println(fi12, v11)
                if (((^(fi12)) <= (fi12 - 956)) && (false && true)) {
                    v22 := []int{}
                    _ = v22
                }
            } else {
                v23 := []int{fi12}
                _ = v23
            }
            var v24 uint64 = 1
            _ = v24
        }
        _ = fi12
    }
    _ = v10
    _ = v11
    return -84
}

func main() {
    var v1 uint = 4294967295
    var v2 int16 = 1000
    var v3 int = 9223372036854775807
    var v4 int32 = -748
    var v25 string = ""
    var v26 float64 = float64((1.5 / float32(float64(float32((-2.25 - 0.1))))))
    var v27 int16 = (^(v2))
    v28 := []int{(v3 - (v3 * 92)), (v3 << (((v1 / (((v1 % ((v1) | 1))) | 1))) & 63))}
    var v29 uint8 = 0
    v1 += v1
    var v30 uint16 = 1
    if ((-(v3)) == v3) {
        v31 := GenS2{F0: 1, F1: v30, F2: 42}
        if len(v28) > 0 {
            v28[0] = 2147483647
        }
        if (v1 == v1) {
            var v32 int16 = (v2 ^ 1000)
            v30 += (v30 << (((v1 >> ((215) & 63))) & 63))
            var v33 float32 = 1.5
            _ = v32
            _ = v33
        }
        _ = v31
    }
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    switch v25 {
    case "\n\t":
        for fi34 := 0; fi34 < 12; fi34++ {
            for fi35 := 0; fi35 < 9; fi35++ {
                var v36 int64 = -831
                _ = fi35
                _ = v36
            }
            v27 = v2
            if true {
                switch (fi34 ^ fi34) {
                case 128:
                    var v37 bool = true
                    _ = v37
                default:
                    var v38 int32 = v4
                    _ = v38
                }
                var v39 uint = v1
                v26 -= v26
                _ = v39
            } else {
                for fi40 := 0; fi40 < 16; fi40++ {
                    v41 := GenS1{F0: true, F1: ""}
                    v42 := GenS3{F0: v2, F1: uint8(v1), F2: false}
                    v4 = -1
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                var v43 int64 = 1
                _ = v43
            }
            _ = fi34
        }
    }
    var v44 uint = (v1 << (((v1 >> ((v1) & 63))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v44
}
