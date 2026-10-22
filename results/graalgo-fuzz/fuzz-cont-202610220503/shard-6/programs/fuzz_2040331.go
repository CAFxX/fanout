package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32 }

func main() {
    var v1 uint = 423
    v2 := make(map[int]bool)
    var v3 uint32 = 798
    v4 := make(map[string]int)
    v5 := make(map[string]bool)
    var v6 bool = ((-1 < -1) || false)
    switch 0 {
    case -649:
        var v7 float32 = -2.25
        _ = v7
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        fmt.Println(v1, fi8)
        var v9 int32 = 42
        v9 = v9
        _ = fi8
        _ = v9
    }
    v6 = (!(!(42 <= 2147483647)))
    v10 := []int{0, -649, 128, 4294967295}
    _ = len(v10)
    switch "foo bar" {
    case "\n\t":
        if (991 != -486) {
            var v11 float32 = 0.1
            var v12 uint16 = 455
            for fi13 := 0; fi13 < 14; fi13++ {
                if false {
                    v14 := GenS1{F0: -734}
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 3; fi15++ {
                        switch "" {
                        case "hello":
                            v16 := GenS1{F0: 1}
                            _ = v16
                        case "":
                            switch fi15 {
                            case 32767:
                                v11 = (v11 * v11)
                            }
                        default:
                            if len(v10) > 0 {
                                v10[0] = int(v12)
                            }
                        }
                        fmt.Println(v3, v11)
                        _ = fi15
                    }
                    v6 = v6
                    v17 := GenS2{F0: v3}
                    _ = v17
                }
                fmt.Println(v6, v3, v6)
                _ = fi13
            }
            _ = v11
            _ = v12
        }
    default:
        switch "foo bar" {
        case "foo bar":
            for fi18 := 0; fi18 < 6; fi18++ {
                for fi19 := 0; fi19 < 19; fi19++ {
                    if v6 {
                        fmt.Println(v1)
                        var v20 uint32 = (v3 / ((v3) | 1))
                        v1 = 122
                        _ = v20
                    } else {
                        v21 := GenS3{F0: v3}
                        if (0 != 806) {
                            if v6 {
                                delete(v4, "\n\t")
                                fmt.Println(v1, v6)
                                fmt.Println(fi19, fi19, fi19)
                            } else {
                                switch "a" {
                                case "foo bar":
                                    delete(v4, "foo bar")
                                }
                                v22 := GenS4{F0: "a"}
                                _ = v22
                            }
                            var v23 int = fi19
                            _ = v23
                        }
                        var v24 int16 = -156
                        var v25 uint16 = uint16(v1)
                        if v6 {
                            if v6 {
                                if (v6 || (1.5 != float64(3.14159))) {
                                    if len(v10) > 0 {
                                        v10[0] = 999
                                    }
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                } else {
                                    var v26 float64 = 0.1
                                    var v27 int16 = (v24 / ((1000) | 1))
                                    _ = v26
                                    _ = v27
                                }
                            } else {
                                if v6 {
                                    v25 += 42
                                } else {
                                    v3 -= (v3 - 42)
                                }
                                switch fi18 {
                                case -477:
                                    _ = len(v10)
                                case 222:
                                    v24 -= v24
                                }
                                v1 += v1
                            }
                            var v28 uint = (v1 * (v1 ^ 0))
                            _ = len(v2)
                            _ = v28
                        } else {
                            if true {
                                if v6 {
                                    var v29 int64 = -1
                                    _ = v29
                                }
                                var v30 int = -1
                                _ = v30
                            } else {
                                v6 = ((3.14159 == 1.5) && v6)
                                v31 := GenS1{F0: 2147483647}
                                _ = v31
                            }
                            var v32 int16 = v24
                            v24 = v24
                            _ = v32
                        }
                        _ = v21
                        _ = v24
                        _ = v25
                    }
                    v1 += 733
                    var v33 int = fi18
                    _ = fi19
                    _ = v33
                }
                v1 = v1
                _ = fi18
            }
        case "a":
            switch ((("foo bar" + (("hello" + "foo bar") + (("a" + "foo bar") + ("a" + "")))) + ("hello" + "foo bar")) + "a") {
            case "foo bar":
                var v34 uint16 = uint16(v3)
                _ = v34
            case "":
                var v35 float64 = (0.1 * (float64((0.0 * ((1.5 / 1.5) / 1.0))) - (-2.25 / float64(((1e10 / 1.0) / (3.14159 * 1e10))))))
                _ = v35
            case "\n\t":
                var v36 uint8 = 128
                _ = v36
            }
        case "\n\t":
            v10 = append(v10, 42)
        default:
            if v6 {
                v1 = v1
            }
        }
    }
    switch 127 {
    case 528:
        v6 = (v6 || true)
    case 4294967295:
        switch -2147483648 {
        case 42:
            v10 = append(v10, -1)
        default:
            var v37 uint64 = 4294967295
            _ = v37
        }
    }
    v10 = append(v10, -54)
    v1 += 0
    v38 := GenS5{F0: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
