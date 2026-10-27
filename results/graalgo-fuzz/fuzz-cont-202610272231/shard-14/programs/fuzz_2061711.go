package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 bool = true
    var v2 string = ((("\n\t" + ("hello" + ("" + "foo bar"))) + "\n\t") + "foo bar")
    v3 := []int{471, -846}
    var v4 uint = 255
    v5 := make(map[int]uint32)
    var v6 float32 = 0.0
    fmt.Println(v6)
    for fi7 := 0; fi7 < 3; fi7++ {
        if v1 {
            if true {
                fmt.Println(v4, v1, v6)
                if (0 >= 1) {
                    if true {
                        v4 += v4
                        fmt.Println(v4, v4, v6)
                    } else {
                        if (v2 >= v2) {
                            var v8 float64 = ((3.14159 + ((float64(0.0) / 1.5) + 0.0)) * 3.14159)
                            v4 = 448
                            _ = v8
                        }
                        fmt.Println(v6, fi7, v1)
                    }
                } else {
                    v9 := GenS1{F0: 1, F1: v1, F2: 0}
                    _ = v9
                }
                var v10 float32 = -2.25
                _ = v10
            }
        } else {
            switch v2 {
            case "hello":
                var v11 int8 = 123
                _ = v11
            case "a":
                v2 = (("\n\t" + v2) + "hello")
            case "foo bar":
                if (v1 && v1) {
                    switch v2 {
                    case "hello":
                        v1 = (!v1)
                    default:
                        switch ("\n\t" + v2) {
                        case "":
                            var v12 float32 = 1e10
                            _ = v12
                        case "\n\t":
                            if false {
                                for fi13 := 0; fi13 < 17; fi13++ {
                                    v1 = (((((v2 + ("a" + ("foo bar" + (v2 + v2)))) + "") + v2) < v2) || (v1 && ((((!(0 == -1)) && v1) || (255 == 127)) && true)))
                                    _ = fi13
                                }
                            } else {
                                v6 -= 0.0
                                for fi14 := 0; fi14 < 2; fi14++ {
                                    v2 = (v2 + "")
                                    _ = fi14
                                }
                            }
                        case "a":
                            if len(v3) > 0 {
                                v3[0] = fi7
                            }
                        }
                    }
                }
            default:
                for fi15 := 0; fi15 < 15; fi15++ {
                    var v16 float64 = float64(3.14159)
                    _ = fi15
                    _ = v16
                }
            }
            var v17 int64 = -386
            _ = v17
        }
        v5[fi7] = 2147483648
        _ = fi7
    }
    var v18 int32 = 2147483647
    switch ((("\n\t" + "foo bar") + v2) + v2) {
    case "a":
        v19 := GenS2{F0: -9223372036854775808}
        _ = v19
    }
    fmt.Println(v18, v2, v18)
    _ = len(v3)
    v6 = float32(1.5)
    var v20 int16 = -484
    var v21 int16 = 32767
    var v22 uint = 360
    for fi23 := 0; fi23 < 18; fi23++ {
        if (87 < 97) {
            var v24 float32 = v6
            _ = v24
        } else {
            v2 = "\n\t"
        }
        v2 = v2
        v25 := GenS2{F0: -888}
        _ = fi23
        _ = v25
    }
    var v26 int8 = int8(v18)
    v27 := GenS3{F0: 128}
    var v28 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
