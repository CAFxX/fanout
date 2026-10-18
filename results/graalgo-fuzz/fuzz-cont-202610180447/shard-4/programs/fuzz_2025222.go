package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 string; F2 int64 }
type GenS4 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint16, p1 int64) uint32 {
    v6 := []int{4294967295}
    var v7 int8 = -128
    if true {
        if (1.5 >= float64(0.0)) {
            if false {
                fmt.Println(v7, v7)
                if true {
                    var v8 int8 = v7
                    var v9 uint8 = 2
                    _ = v8
                    _ = v9
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
        }
        if len(v6) > 0 {
            v6[0] = -1
        }
        _ = len(v6)
    }
    v10 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 int = -1
    v2 := make(map[string]string)
    v3 := []string{("\n\t" + ("\n\t" + "a"))}
    var v4 uint = 255
    v5 := make(map[uint32]string)
    for fi11 := 0; fi11 < 12; fi11++ {
        if (!(!(42 > -617))) {
            v4 = v4
            var v12 int8 = -128
            _ = v12
        }
        v4 = (v4 % (((v4 + v4)) | 1))
        fmt.Println(v4, v4, fi11)
        _ = fi11
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := GenS1{F0: 106}
        _ = fi13
        _ = v14
    }
    v1 += (^(v1))
    if (656 <= 0) {
        if ((((-2.25 + -2.25) - (-2.25 * float64(float32(3.14159)))) != 1.5) && ((v1 % ((v1) | 1)) == 0)) {
            if len(v3) > 0 {
                v3[0] = (((("a" + "hello") + ("" + "")) + "foo bar") + "")
            }
            if ("hello" == ("hello" + ("" + ((("" + "a") + "hello") + "\n\t")))) {
                switch v1 {
                case 435:
                    v15 := GenS1{F0: 255}
                    _ = v15
                default:
                    v4 = v4
                }
                if (128 <= 1) {
                    v1 = -498
                } else {
                    switch (v1 << (((v4 * 1)) & 63)) {
                    case 128:
                        fmt.Println(v4)
                    case 580:
                        for fi16 := 0; fi16 < 7; fi16++ {
                            fmt.Println(v1)
                            fmt.Println(fi16, v4)
                            v4 += 4294967295
                            _ = fi16
                        }
                    }
                    fmt.Println(v1, v4)
                    v1 = 32768
                }
                v17 := GenS2{F0: 220}
                _ = v17
            }
            fmt.Println(v1, v4)
        }
        for fi18 := 0; fi18 < 17; fi18++ {
            v19 := GenS3{F0: -128, F1: "a", F2: -9223372036854775808}
            _ = fi18
            _ = v19
        }
    }
    switch (v1 / ((976) | 1)) {
    case -9223372036854775808:
        var v20 float64 = ((1e10 * float64((float32(((3.14159 * 1.5) * (((1.5 / float64((0.1 * (0.1 * 0.1)))) / float64(1.5)) * 0.1))) * (((0.1 * (1.5 / 1e10)) + 0.0) - (1.5 * -2.25))))) - float64((-2.25 - (((1e10 * 3.14159) - 0.1) - 1.5))))
        _ = v20
    case -2147483648:
        var v21 bool = (!(-121 < -128))
        _ = v21
    default:
        var v22 string = ("foo bar" + "hello")
        _ = v22
    }
    v1 = 32767
    switch 9223372036854775807 {
    case 128:
        v1 -= 128
    case -547:
        if false {
            for fi23 := 0; fi23 < 11; fi23++ {
                switch v1 {
                case 127:
                    switch fi23 {
                    case -9223372036854775808:
                        var v24 uint = (^(v4))
                        _ = v24
                    case 42:
                        fmt.Println(v1)
                    case 1:
                        v1 = v1
                    default:
                        var v25 int64 = 0
                        _ = v25
                    }
                case 32767:
                    if false {
                        var v26 string = "a"
                        var v27 int32 = 273
                        _ = v26
                        _ = v27
                    } else {
                        switch v1 {
                        case -9223372036854775808:
                            fmt.Println(v1, fi23)
                        default:
                            v1 = (-(v1))
                        }
                    }
                default:
                    v4 -= 42
                }
                _ = fi23
            }
        } else {
            v1 += 9223372036854775807
        }
    case -9223372036854775808:
        for fi28 := 0; fi28 < 18; fi28++ {
            if (genFn0(65535, 42) < 603) {
                fmt.Println(fi28, v4, v4)
                v4 = 255
            } else {
                fmt.Println(v1, v4, v1)
                var v29 int16 = -1
                _ = v29
            }
            _ = fi28
        }
    default:
        fmt.Println(v1)
    }
    delete(v5, 42)
    _ = len(v3)
    v30 := GenS2{F0: 1}
    switch -312 {
    case -9223372036854775808:
        var v31 float32 = 0.0
        _ = v31
    case 4294967295:
        var v32 uint8 = 1
        _ = v32
    case 42:
        if ((!true) || true) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v1 = v1
            v4 = (^(v4))
        }
    }
    fmt.Println(v1, v4)
    if true {
        v33 := GenS4{F0: 0, F1: 0}
        var v34 float32 = 0.0
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
