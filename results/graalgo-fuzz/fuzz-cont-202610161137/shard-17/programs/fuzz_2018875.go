package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }

func main() {
    v1 := []int{}
    v2 := []uint32{4294967295}
    var v3 uint32 = 2
    var v4 int32 = -1
    var v5 float32 = 1e10
    v6 := make(map[int]uint32)
    v7 := v6[32768]
    var v8 float32 = v5
    _ = len(v1)
    v3 += 2147483647
    if len(v1) > 0 {
        v1[0] = -567
    }
    var v9 float64 = 1.5
    var v10 int8 = -61
    var v11 float64 = 3.14159
    switch 2147483647 {
    case 4294967295:
        switch "" {
        case "foo bar":
            fmt.Println(v7, v11)
        case "a":
            if false {
                fmt.Println(v4, v7)
            } else {
                v9 = v11
                v5 += (-2.25 * -2.25)
                var v12 uint16 = 0
                _ = v12
            }
        }
    default:
        v5 = v5
    }
    var v13 uint16 = 0
    if false {
        if false {
            for fi14 := 0; fi14 < 12; fi14++ {
                fmt.Println(v4)
                v5 = -2.25
                _ = len(v1)
                _ = fi14
            }
        }
        var v15 uint64 = 42
        var v16 uint32 = 575
        _ = v15
        _ = v16
    } else {
        switch (("" + "a") + "hello") {
        case "hello":
            if false {
                if (!(!(!true))) {
                    var v17 float64 = v11
                    var v18 int16 = 785
                    _ = v17
                    _ = v18
                }
            } else {
                v19 := GenS1{F0: "hello", F1: v7, F2: 18446744073709551615}
                var v20 float64 = (v9 + 1e10)
                v21 := GenS1{F0: "\n\t", F1: (v7 * v3), F2: 2}
                _ = v19
                _ = v20
                _ = v21
            }
        case "a":
            v11 = float64((v8 + v8))
        case "foo bar":
            v22 := GenS1{F0: "\n\t", F1: v3, F2: 0}
            _ = v22
        default:
            var v23 string = "\n\t"
            _ = v23
        }
        v10 -= (^(v10))
        var v24 int = 909
        _ = v24
    }
    var v25 float32 = ((0.1 * v5) + (float32((v11 * float64(1e10))) / ((v8 / ((v8 / 3.14159) / v8)) + ((((((0.1 + v5) + 1e10) + 1.5) / (v5 * 0.0)) / ((v5 + ((v8 * v5) - v8)) / (v5 + (1e10 / v8)))) - float32(float64(v8))))))
    switch 127 {
    case 42:
        if false {
            fmt.Println(v4)
            if (v4 < 739) {
                v26 := GenS2{F0: -435}
                v27 := GenS1{F0: ("a" + "foo bar"), F1: 4294967295, F2: 375}
                v28 := GenS3{F0: 134}
                _ = v26
                _ = v27
                _ = v28
            } else {
                var v29 string = ("foo bar" + "\n\t")
                for fi30 := 0; fi30 < 5; fi30++ {
                    v31 := GenS3{F0: 1}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi30
                    _ = v31
                }
                switch -1 {
                case 2147483647:
                    fmt.Println(v11, v25, v11)
                case 42:
                    for fi32 := 0; fi32 < 10; fi32++ {
                        for fi33 := 0; fi33 < 6; fi33++ {
                            v34 := GenS4{F0: 0}
                            _ = fi33
                            _ = v34
                        }
                        _ = fi32
                    }
                }
                _ = v29
            }
            fmt.Println(v4, v11, v9)
        } else {
            v3 += (v7 % (((v3 << ((82) & 63))) | 1))
            if len(v1) > 0 {
                v1[0] = 1
            }
        }
    default:
        v9 -= v11
    }
    _ = len(v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
