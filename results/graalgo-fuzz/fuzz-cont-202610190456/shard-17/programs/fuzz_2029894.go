package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 string; F2 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 int; F2 uint }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 uint64; F1 int16 }
type GenS7 struct { F0 int; F1 bool }
func genFn0(p0 string) int32 {
    v6 := GenS1{F0: -1}
    v7 := GenS1{F0: -2147483648}
    if (255 != 52) {
        p0 = "\n\t"
        fmt.Println(p0)
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 uint16 = 1
            v9 = v9
            switch (fi8 | fi8) {
            case 4294967295:
                fmt.Println(v9, p0)
            case -816:
                v9 += (v9 % ((v9) | 1))
            }
            _ = fi8
            _ = v9
        }
    } else {
        v10 := []int{941, 4294967295, -10}
        if (255 <= 530) {
            switch ((((p0 + p0) + (((p0 + (p0 + (p0 + p0))) + p0) + "foo bar")) + p0) + (p0 + p0)) {
            case "hello":
                p0 = ("\n\t" + (p0 + p0))
            case "foo bar":
                var v11 string = p0
                _ = v11
            case "\n\t":
                if (1 < 85) {
                    v12 := GenS1{F0: 42}
                    _ = v12
                } else {
                    fmt.Println(p0)
                }
            }
            for fi13 := 0; fi13 < 13; fi13++ {
                var v14 float32 = float32(-2.25)
                v15 := GenS2{F0: 4294967295, F1: p0, F2: 2}
                switch (-(fi13)) {
                case -558:
                    p0 = (p0 + (p0 + p0))
                case 32768:
                    var v16 float64 = 3.14159
                    _ = v16
                }
                _ = fi13
                _ = v14
                _ = v15
            }
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 uint64 = 12
                switch fi17 {
                case -1:
                    if false {
                        var v19 float32 = 0.1
                        if true {
                            switch fi17 {
                            case -9223372036854775808:
                                for fi20 := 0; fi20 < 2; fi20++ {
                                    v18 += 0
                                    var v21 bool = ((0.1 + 3.14159) >= (1.5 + (float64(1.5) - 0.0)))
                                    var v22 uint = 0
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                            case 127:
                                fmt.Println(v19, p0)
                            case -628:
                                fmt.Println(v18)
                            }
                        }
                        _ = v19
                    }
                default:
                    v23 := GenS3{F0: 4294967295}
                    _ = v23
                }
                v24 := GenS4{F0: fi17, F1: fi17, F2: 948}
                _ = fi17
                _ = v18
                _ = v24
            }
        } else {
            for fi25 := 0; fi25 < 4; fi25++ {
                var v26 int64 = -338
                _ = fi25
                _ = v26
            }
        }
        for fi27 := 0; fi27 < 16; fi27++ {
            for fi28 := 0; fi28 < 6; fi28++ {
                if len(v10) > 0 {
                    v10[0] = (fi27 % ((fi27) | 1))
                }
                if true {
                    v29 := GenS5{F0: 151, F1: 1}
                    _ = v29
                }
                var v30 int8 = 1
                _ = fi28
                _ = v30
            }
            _ = fi27
        }
        _ = v10
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        p0 = "a"
        _ = fi31
    }
    _ = p0
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint8 = 212
    var v2 uint = 18446744073709551615
    v3 := make(map[uint32]string)
    var v4 int64 = -1
    v5 := []uint32{877, 2147483648, 4294967295}
    var v32 uint32 = 65535
    v33 := GenS6{F0: 4294967295, F1: 1000}
    v34 := GenS7{F0: -2147483648, F1: false}
    v4 += (v4 % ((v4) | 1))
    var v35 int16 = 0
    fmt.Println(v4, v35, v32)
    var v36 uint16 = 2
    var v37 float64 = 0.0
    _ = len(v3)
    v32 -= v32
    fmt.Println(v2, v35)
    v38 := GenS5{F0: 1, F1: (v4 + v4)}
    fmt.Println(v37)
    v37 += v37
    var v39 uint8 = (^(v1))
    switch "\n\t" {
    case "a":
        if len(v5) > 0 {
            v5[0] = 569
        }
    case "foo bar":
        v32 = v32
    default:
        v4 += (v4 % (((v4 * v4)) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
