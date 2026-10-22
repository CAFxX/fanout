package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint; F1 bool; F2 string }
type GenS4 struct { F0 uint8 }
func genFn0() uint16 {
    switch -9223372036854775808 {
    case -2147483648:
        v4 := []int{1}
        _ = v4
    case -805:
        var v5 int16 = -1
        _ = v5
    }
    switch 0 {
    case -1:
        var v6 uint16 = 0
        _ = v6
    }
    return 42
}
func genFn1() uint16 {
    var v7 int = 1
    var v8 uint32 = 182
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int16 = -600
    v2 := []uint32{640}
    var v3 uint32 = 65535
    if true {
        if (false && false) {
            var v9 float32 = (float32(1.5) / float32((0.1 / float64(3.14159))))
            v1 = (v1 * v1)
            fmt.Println(v9, v9)
            _ = v9
        } else {
            var v10 uint = 522
            _ = v10
        }
        switch -9223372036854775808 {
        case 127:
            var v11 uint = 4294967295
            _ = v11
        case 4294967295:
            var v12 bool = true
            _ = v12
        case 2147483647:
            v13 := GenS1{F0: (v1 >> ((598) & 63))}
            _ = v13
        default:
            v1 = 1
        }
        v14 := GenS1{F0: (v1 << ((42) & 63))}
        _ = v14
    } else {
        v15 := GenS2{F0: 2147483647, F1: 42}
        var v16 string = "\n\t"
        _ = v15
        _ = v16
    }
    if (!(!false)) {
        v1 += v1
        for fi17 := 0; fi17 < 7; fi17++ {
            if ("foo bar" <= "foo bar") {
                _ = len(v2)
                var v18 int8 = 84
                fmt.Println(v1)
                _ = v18
            }
            _ = fi17
        }
        for fi19 := 0; fi19 < 19; fi19++ {
            if len(v2) > 0 {
                v2[0] = (v3 - v3)
            }
            _ = fi19
        }
    }
    if len(v2) > 0 {
        v2[0] = (v3 ^ 1)
    }
    fmt.Println(v3, v3)
    var v20 float32 = 0.1
    if len(v2) > 0 {
        v2[0] = v3
    }
    var v21 string = "hello"
    v2 = append(v2, (v3 * v3))
    switch "\n\t" {
    case "hello":
        if true {
            for fi22 := 0; fi22 < 3; fi22++ {
                var v23 int32 = -2147483648
                _ = fi22
                _ = v23
            }
            switch 9223372036854775807 {
            case 32767:
                switch -9223372036854775808 {
                case 42:
                    for fi24 := 0; fi24 < 3; fi24++ {
                        var v25 int32 = -2147483648
                        v20 = float32((1.5 / float64(-2.25)))
                        var v26 float32 = (float32(((float64(3.14159) / float64((0.1 / 1.0))) + 1.5)) / (((v20 - (((float32(((0.0 - 0.1) - (0.1 / (0.1 * 1.5)))) / ((v20 * v20) * 1.5)) + v20) / (0.1 - v20))) / v20) / v20))
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                case 127:
                    v3 -= v3
                }
            case 127:
                if false {
                    var v27 uint16 = 387
                    v28 := GenS3{F0: 255, F1: false, F2: ""}
                    _ = v27
                    _ = v28
                } else {
                    if true {
                        v20 -= ((v20 * float32(1e10)) / v20)
                        fmt.Println(v3, v1, v3)
                    } else {
                        if len(v2) > 0 {
                            v2[0] = v3
                        }
                        fmt.Println(v1)
                        var v29 uint64 = 624
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 17; fi30++ {
                        for fi31 := 0; fi31 < 1; fi31++ {
                            v20 = 1e10
                            switch (((v21 + (((v21 + (("" + "\n\t") + v21)) + v21) + v21)) + (("hello" + "\n\t") + ("hello" + v21))) + v21) {
                            case "foo bar":
                                var v32 string = ("hello" + (("" + "") + ("\n\t" + "foo bar")))
                                _ = v32
                            case "":
                                var v33 string = v21
                                _ = v33
                            case "hello":
                                v3 = (v3 / (((v3 % ((v3) | 1))) | 1))
                            default:
                                fmt.Println(v1)
                            }
                            v21 = v21
                            v20 += 3.14159
                            v3 = v3
                            _ = fi31
                        }
                        switch (v21 + v21) {
                        case "foo bar":
                            v34 := GenS1{F0: v1}
                            _ = v34
                        case "\n\t":
                            v35 := GenS3{F0: 255, F1: false, F2: v21}
                            _ = v35
                        case "":
                            v36 := GenS4{F0: 170}
                            _ = v36
                        default:
                            if (!(951 <= v3)) {
                                v2 = append(v2, 4294967295)
                            }
                        }
                        var v37 int = -118
                        _ = fi30
                        _ = v37
                    }
                }
            }
        }
    case "":
        for fi38 := 0; fi38 < 6; fi38++ {
            fmt.Println(fi38, v21, v1)
            _ = fi38
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
