package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 string }
type GenS2 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int64, p1 uint8, p2 string) int8 {
    p2 = (p2 + p2)
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return -128
}

func main() {
    var v1 float32 = (1.5 - 0.0)
    var v2 uint = 42
    var v3 uint64 = 1
    switch "foo bar" {
    case "":
        v4 := GenS1{F0: 0, F1: 1, F2: ""}
        _ = v4
    case "foo bar":
        var v5 uint32 = uint32(v2)
        _ = v5
    default:
        var v6 int = 32767
        _ = v6
    }
    v3 = v3
    v2 -= (v2 >> ((v2) & 63))
    v7 := []int{42, 0, 9223372036854775807}
    if ("\n\t" > "hello") {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        fmt.Println(v2, v2, v3)
        for fi8 := 0; fi8 < 9; fi8++ {
            v3 = (^(v3))
            var v9 float64 = float64((-2.25 * v1))
            v10 := GenS1{F0: genFn0(1, 1, ("hello" + ("" + "a"))), F1: 9, F2: "hello"}
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v11 := GenS2{F0: 1, F1: -2147483648, F2: (^(v3))}
    switch "a" {
    case "\n\t":
        v12 := GenS1{F0: 100, F1: 1000, F2: ""}
        _ = v12
    default:
        if false {
            if true {
                for fi13 := 0; fi13 < 10; fi13++ {
                    var v14 bool = ((fi13 / ((fi13) | 1)) == 219)
                    var v15 float64 = (0.0 - float64(float32(1e10)))
                    var v16 int64 = 9223372036854775807
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
            } else {
                var v17 bool = false
                if len(v7) > 0 {
                    v7[0] = -1
                }
                _ = v17
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
    }
    v18 := GenS3{F0: 42}
    var v19 uint32 = 732
    fmt.Println(v3, v19, v3)
    if true {
        v3 += v3
        if ((false && ((-222 <= 1) && true)) || false) {
            fmt.Println(v3)
            var v20 bool = (!true)
            _ = v20
        } else {
            v21 := GenS2{F0: v19, F1: -2147483648, F2: (v3 * (v3 ^ (^(v3))))}
            _ = v21
        }
        if (0 == 369) {
            var v22 uint8 = 255
            _ = v22
        }
    } else {
        if (2 < 674) {
            for fi23 := 0; fi23 < 5; fi23++ {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                for fi24 := 0; fi24 < 18; fi24++ {
                    v3 = 4294967295
                    v1 = (float32((float64(float32(1e10)) - 3.14159)) + (v1 + ((v1 + (float32(float64((v1 * v1))) + 3.14159)) / v1)))
                    _ = fi24
                }
                _ = fi23
            }
            v2 = (^(v2))
            fmt.Println(v3)
        } else {
            v25 := GenS1{F0: -93, F1: 963, F2: "a"}
            var v26 float64 = 1e10
            if (!false) {
                var v27 float32 = 1e10
                _ = v27
            } else {
                v7 = append(v7, 42)
            }
            _ = v25
            _ = v26
        }
        if (!true) {
            for fi28 := 0; fi28 < 10; fi28++ {
                for fi29 := 0; fi29 < 1; fi29++ {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    var v30 int = fi28
                    var v31 float64 = 0.1
                    if true {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    } else {
                        fmt.Println(v31)
                        switch (v30 / (((fi29 - (fi28 ^ 9223372036854775807))) | 1)) {
                        case 42:
                            v32 := GenS3{F0: 939}
                            _ = v32
                        case 128:
                            v33 := GenS4{F0: 798}
                            _ = v33
                        case 59:
                            switch v30 {
                            case 127:
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            }
                            v19 = v19
                            var v34 string = ((("" + (("foo bar" + "") + "foo bar")) + "hello") + ("foo bar" + "hello"))
                            v35 := GenS5{F0: -128}
                            _ = v34
                            _ = v35
                        }
                    }
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                _ = fi28
            }
            if len(v7) > 0 {
                v7[0] = -576
            }
            v1 += (v1 / (0.1 * (v1 + v1)))
        }
        v1 += v1
    }
    fmt.Println(v19, v1, v2)
    fmt.Println(v2, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v7
}
