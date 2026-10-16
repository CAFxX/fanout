package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 int32; F1 uint; F2 uint8 }

func main() {
    v1 := []int{32767, 862, 127, 453}
    var v2 float32 = -2.25
    v3 := make(map[uint32]uint32)
    switch (("a" + ("hello" + ("a" + "foo bar"))) + "a") {
    case "foo bar":
        var v4 int8 = 127
        _ = v4
    case "":
        fmt.Println(v2)
    }
    v5 := GenS1{F0: 1, F1: 1}
    switch (("\n\t" + "foo bar") + "a") {
    case "":
        v2 = v2
    case "hello":
        var v6 int16 = -1
        _ = v6
    case "a":
        var v7 float64 = 3.14159
        _ = v7
    }
    var v8 int16 = 342
    var v9 float32 = v2
    if len(v1) > 0 {
        v1[0] = 42
    }
    v10 := GenS2{F0: -915, F1: -2147483648}
    v9 = v9
    if true {
        var v11 float64 = float64(-2.25)
        if true {
            v9 = (v2 / float32(v11))
            v9 += (v2 - v2)
        } else {
            if (!false) {
                var v12 string = "a"
                fmt.Println(v12, v9, v11)
                for fi13 := 0; fi13 < 15; fi13++ {
                    var v14 bool = (fi13 < 0)
                    if v14 {
                        for fi15 := 0; fi15 < 20; fi15++ {
                            if len(v1) > 0 {
                                v1[0] = 0
                            }
                            _ = fi15
                        }
                    }
                    var v16 int32 = 42
                    _ = fi13
                    _ = v14
                    _ = v16
                }
                _ = v12
            }
            for fi17 := 0; fi17 < 20; fi17++ {
                if true {
                    v11 = (float64(-2.25) - ((1e10 - v11) * v11))
                    var v18 int16 = 691
                    _ = v18
                } else {
                    switch -1 {
                    case -880:
                        v11 = ((0.1 * (0.0 + 0.1)) / 1.0)
                    case 9223372036854775807:
                        var v19 int32 = -1
                        _ = v19
                    case -1:
                        v11 -= (v11 / (v11 / ((1e10 * (v11 * (0.1 / float64((((v9 - (1e10 / float32(1e10))) + 1.5) / v2))))) / v11)))
                    default:
                        v9 = v9
                    }
                    v11 = (v11 / v11)
                }
                v8 = 1
                _ = fi17
            }
            v20 := GenS3{F0: 1, F1: 65535}
            _ = v20
        }
        if (42 <= 1) {
            for fi21 := 0; fi21 < 8; fi21++ {
                fmt.Println(v8, v11, v11)
                _ = fi21
            }
            if true {
                var v22 string = "hello"
                var v23 int64 = -1
                v22 = (("" + (v22 + v22)) + (v22 + v22))
                _ = v22
                _ = v23
            }
        } else {
            var v24 int32 = -2147483648
            v25 := v3[0]
            var v26 int64 = -9223372036854775808
            _ = v24
            _ = v25
            _ = v26
        }
        _ = v11
    } else {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi27 := 0; fi27 < 12; fi27++ {
                fmt.Println(v9, v8)
                var v28 int = (fi27 ^ 0)
                fmt.Println(v2)
                _ = fi27
                _ = v28
            }
            if true {
                v2 = 3.14159
                switch "\n\t" {
                case "foo bar":
                    if true {
                        v9 = float32(1.5)
                    }
                case "hello":
                    if false {
                        v1 = append(v1, 42)
                        _ = len(v3)
                    } else {
                        var v29 uint = 255
                        _ = v29
                    }
                case "\n\t":
                    for fi30 := 0; fi30 < 16; fi30++ {
                        v8 = 1
                        _ = fi30
                    }
                }
            }
        } else {
            if (!true) {
                fmt.Println(v9)
            } else {
                for fi31 := 0; fi31 < 5; fi31++ {
                    var v32 uint = 42
                    v33 := v3[122]
                    var v34 bool = (!(0.0 >= 1.5))
                    _ = fi31
                    _ = v32
                    _ = v33
                    _ = v34
                }
                var v35 int = -9223372036854775808
                var v36 uint64 = 2
                _ = v35
                _ = v36
            }
            v37 := GenS1{F0: -1, F1: 0}
            _ = v37
        }
        if (!(!(!(19 != 618)))) {
            if true {
                for fi38 := 0; fi38 < 7; fi38++ {
                    var v39 float64 = (0.1 / 1.0)
                    _ = fi38
                    _ = v39
                }
            } else {
                v8 = v8
                v2 += float32(0.0)
                _ = len(v3)
            }
            if false {
                var v40 float64 = 3.14159
                _ = v40
            } else {
                v9 += ((0.0 / v2) + (v2 + (v9 - (float32(0.0) + (v9 * v2)))))
            }
        } else {
            if (-2147483648 == 2147483647) {
                if false {
                    if (!false) {
                        v2 -= (-2.25 + v9)
                        v41 := GenS4{F0: 1, F1: -2147483648}
                        _ = v41
                    }
                    if (595 < 1) {
                        var v42 bool = (!(int32(v8) >= 2147483647))
                        _ = v42
                    }
                } else {
                    for fi43 := 0; fi43 < 5; fi43++ {
                        var v44 uint8 = 21
                        delete(v3, 627)
                        _ = fi43
                        _ = v44
                    }
                    v3[2147483648] = 0
                }
                fmt.Println(v2, v9, v9)
            } else {
                v2 -= (v2 / v9)
            }
            fmt.Println(v8)
        }
        if (100 != 0) {
            v8 += (-(v8))
            if len(v1) > 0 {
                v1[0] = 462
            }
        }
    }
    v45 := GenS5{F0: -41, F1: 2, F2: 105}
    fmt.Println(v9, v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v45
    _ = v5
    _ = v8
    _ = v9
}
