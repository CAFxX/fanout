package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string }

func main() {
    var v1 uint32 = 2
    var v2 uint16 = 692
    var v3 bool = true
    v4 := make(map[string]string)
    var v5 uint32 = (^(v1))
    v6 := make(map[uint32]int)
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: 32767, F1: -1, F2: 838}
        _ = fi7
        _ = v8
    }
    fmt.Println(v5, v2)
    var v9 int = -2147483648
    v5 = 2147483647
    v2 += 1
    v10 := GenS2{F0: (v3 || v3)}
    var v11 uint32 = (v5 % ((v5) | 1))
    var v12 uint16 = v2
    v13 := GenS3{F0: ((v1 != v5) || ((v9 | (^(v9))) >= -867))}
    if (0 >= 42) {
        for fi14 := 0; fi14 < 5; fi14++ {
            if v3 {
                if (!v3) {
                    switch v9 {
                    case 128:
                        var v15 string = (("a" + "hello") + ("" + "hello"))
                        _ = v15
                    case -571:
                        var v16 uint8 = 248
                        _ = v16
                    default:
                        v17 := GenS2{F0: ((v2 & 227) < v12)}
                        _ = v17
                    }
                    var v18 uint64 = 1
                    _ = v18
                }
                if false {
                    var v19 float32 = (-2.25 - 1.5)
                    v20 := []int{0, fi14}
                    _ = v19
                    _ = v20
                }
            } else {
                var v21 uint8 = 229
                switch 9223372036854775807 {
                case 32767:
                    for fi22 := 0; fi22 < 15; fi22++ {
                        fmt.Println(v3, fi14)
                        fmt.Println(v9)
                        _ = fi22
                    }
                case 524:
                    v5 = v1
                }
                for fi23 := 0; fi23 < 18; fi23++ {
                    var v24 int64 = 9223372036854775807
                    v5 -= v1
                    _ = fi23
                    _ = v24
                }
                _ = v21
            }
            _ = fi14
        }
        var v25 bool = (127 != -128)
        v3 = ((42 < -1) && v25)
        _ = v25
    } else {
        for fi26 := 0; fi26 < 14; fi26++ {
            for fi27 := 0; fi27 < 6; fi27++ {
                var v28 uint64 = 798
                if v3 {
                    var v29 uint16 = v2
                    _ = v29
                } else {
                    if false {
                        v11 += (v11 / ((v11) | 1))
                        delete(v6, (v11 & v5))
                    } else {
                        fmt.Println(v9)
                        var v30 float32 = ((0.0 - (3.14159 - 1.5)) - (0.1 * ((1.5 - (3.14159 / 1.5)) - -2.25)))
                        v31 := GenS4{F0: "foo bar"}
                        var v32 int8 = 76
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    v33 := []int{4294967295, (fi27 / (((v9 + (^(v9)))) | 1)), fi26}
                    switch (fi27 + (-(fi26))) {
                    case 32768:
                        fmt.Println(v11, v28, v28)
                    default:
                        var v34 int64 = 42
                        _ = v34
                    }
                    _ = v33
                }
                fmt.Println(fi27, fi27, fi26)
                _ = fi27
                _ = v28
            }
            if v3 {
                var v35 int8 = -128
                _ = v35
            }
            _ = fi26
        }
        v36 := GenS2{F0: v3}
        var v37 float32 = 0.1
        _ = v36
        _ = v37
    }
    switch v9 {
    case 510:
        v3 = v3
    }
    var v38 uint8 = 42
    var v39 string = "foo bar"
    var v40 uint16 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v9
}
