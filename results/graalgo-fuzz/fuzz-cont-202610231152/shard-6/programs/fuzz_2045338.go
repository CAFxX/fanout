package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 bool) int16 {
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(fi7, fi7)
        if (false || true) {
            p0 = (p0 && p0)
        } else {
            var v8 int8 = 31
            var v9 uint = 90
            fmt.Println(p0, v9, fi7)
            _ = v8
            _ = v9
        }
        v10 := []int{4294967295, 404, (-(fi7)), fi7}
        _ = fi7
        _ = v10
    }
    fmt.Println(p0)
    p0 = false
    v11 := GenS1{F0: 0}
    _ = p0
    _ = v11
    return -1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 42
    v3 := make(map[int]bool)
    var v4 int64 = 9223372036854775807
    var v5 int64 = v4
    v6 := []uint32{1}
    if (0.1 >= 3.14159) {
        var v12 float32 = -2.25
        _ = v12
    } else {
        fmt.Println(v5, v2, v2)
        if (-529 != 69) {
            for fi13 := 0; fi13 < 13; fi13++ {
                fmt.Println(fi13, v2, fi13)
                v4 -= v4
                _ = fi13
            }
            v4 = (v4 >> ((0) & 63))
        }
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        v2 += (v2 >> ((137) & 63))
        _ = fi14
    }
    fmt.Println(v5, v4)
    v5 = v4
    for fi15 := 0; fi15 < 11; fi15++ {
        fmt.Println(v2, v4, v4)
        _ = fi15
    }
    delete(v1, 641)
    var v16 string = ("foo bar" + (("hello" + "foo bar") + ("a" + ("foo bar" + "foo bar"))))
    if true {
        if true {
            v17 := GenS1{F0: 0}
            var v18 bool = (512 == 32767)
            var v19 uint = 748
            _ = v17
            _ = v18
            _ = v19
        }
    } else {
        v3[0] = false
        for fi20 := 0; fi20 < 5; fi20++ {
            fmt.Println(fi20, v4, v5)
            _ = len(v6)
            fmt.Println(v2)
            _ = fi20
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            switch "a" {
            case "a":
                v5 += v4
            case "":
                if false {
                    if true {
                        switch v16 {
                        case "hello":
                            if ((false && false) || false) {
                                for fi22 := 0; fi22 < 14; fi22++ {
                                    switch v16 {
                                    case "\n\t":
                                        var v23 string = v16
                                        _ = v23
                                    }
                                    _ = fi22
                                }
                            } else {
                                for fi24 := 0; fi24 < 9; fi24++ {
                                    if true {
                                        switch fi21 {
                                        case 384:
                                            if true {
                                                if (float64(0.0) <= 0.0) {
                                                    v16 = v16
                                                }
                                                v25 := GenS1{F0: 127}
                                                var v26 int = fi24
                                                _ = v25
                                                _ = v26
                                            } else {
                                                var v27 uint8 = 0
                                                v16 = v16
                                                if true {
                                                    if (!true) {
                                                        v28 := v1[1]
                                                        var v29 int64 = 60
                                                        v30 := v3[271]
                                                        _ = v28
                                                        _ = v29
                                                        _ = v30
                                                    } else {
                                                        var v31 int = 2147483647
                                                        var v32 uint = 65535
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                    fmt.Println(fi21, fi21)
                                                } else {
                                                    switch fi24 {
                                                    case -582:
                                                        v33 := GenS1{F0: (v27 / (((v27 + v27)) | 1))}
                                                        _ = v33
                                                    default:
                                                        v27 += v27
                                                    }
                                                    var v34 uint8 = v27
                                                    _ = v34
                                                }
                                                var v35 float64 = (0.0 - 0.1)
                                                _ = v27
                                                _ = v35
                                            }
                                        }
                                    }
                                    if (536 <= 2) {
                                        var v36 uint = 42
                                        v1[2147483647] = v16
                                        v37 := GenS1{F0: 84}
                                        _ = v36
                                        _ = v37
                                    }
                                    v2 -= (^(v2))
                                    _ = fi24
                                }
                                var v38 string = "foo bar"
                                _ = v38
                            }
                        }
                        _ = len(v6)
                    }
                }
            case "hello":
                var v39 int32 = -1
                _ = v39
            }
            var v40 string = ""
            fmt.Println(v40, v2)
            _ = fi21
            _ = v40
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
