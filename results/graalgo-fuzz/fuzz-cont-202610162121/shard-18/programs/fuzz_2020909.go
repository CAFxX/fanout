package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int32 }
func genFn0() int {
    v5 := []int{}
    if (1 <= -935) {
        v6 := GenS1{F0: 42}
        _ = v6
    } else {
        var v7 int32 = -821
        v5 = append(v5, 118)
        for fi8 := 0; fi8 < 10; fi8++ {
            if false {
                for fi9 := 0; fi9 < 8; fi9++ {
                    v10 := GenS2{F0: int16(fi9), F1: 837}
                    _ = fi9
                    _ = v10
                }
                var v11 float64 = float64(((1e10 / ((float32(3.14159) + 0.0) - 0.1)) / 1.5))
                var v12 uint8 = 1
                _ = v11
                _ = v12
            }
            if true {
                v7 += v7
                fmt.Println(v7)
                switch (-(fi8)) {
                case 4294967295:
                    switch (fi8 >> ((4294967295) & 63)) {
                    case 9223372036854775807:
                        var v13 bool = false
                        _ = v13
                    }
                case 32768:
                    for fi14 := 0; fi14 < 11; fi14++ {
                        var v15 uint32 = 2147483647
                        if true {
                            v7 -= (-(v7))
                            v7 = v7
                            switch fi8 {
                            case 32768:
                                var v16 float32 = float32(0.0)
                                _ = v16
                            default:
                                v15 -= v15
                            }
                        } else {
                            var v17 int = (fi8 % ((fi14) | 1))
                            _ = v17
                        }
                        _ = fi14
                        _ = v15
                    }
                }
            }
            switch "\n\t" {
            case "hello":
                if (true && (((("\n\t" + "hello") < "") || false) || (255 >= 1))) {
                    var v18 float32 = 3.14159
                    if true {
                        var v19 uint8 = 128
                        v19 = v19
                        _ = v19
                    } else {
                        if true {
                            v7 = v7
                        }
                        if false {
                            if len(v5) > 0 {
                                v5[0] = (fi8 % ((fi8) | 1))
                            }
                            for fi20 := 0; fi20 < 3; fi20++ {
                                _ = len(v5)
                                _ = fi20
                            }
                        }
                    }
                    _ = len(v5)
                    _ = v18
                }
            case "":
                v7 = -234
            case "\n\t":
                for fi21 := 0; fi21 < 1; fi21++ {
                    fmt.Println(fi8)
                    var v22 string = "foo bar"
                    fmt.Println(v7)
                    if len(v5) > 0 {
                        v5[0] = fi8
                    }
                    _ = fi21
                    _ = v22
                }
            }
            _ = fi8
        }
        _ = v7
    }
    var v23 int16 = 100
    v23 = (v23 - (v23 & v23))
    var v24 string = "hello"
    _ = v23
    _ = v24
    _ = v5
    return 32767
}
func genFn1(p0 int16, p1 int) int16 {
    v25 := GenS2{F0: (-(p0)), F1: 1}
    var v26 float32 = (0.1 / 0.1)
    var v27 uint16 = 2
    var v28 bool = false
    v29 := []int{p1, (^(p1))}
    _ = p0
    _ = p1
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    return (-(p0))
}

func main() {
    v1 := []uint32{2, 657}
    var v2 uint16 = 255
    var v3 int16 = -161
    var v4 float32 = (-2.25 - 1e10)
    var v30 string = "a"
    switch 784 {
    case -350:
        for fi31 := 0; fi31 < 10; fi31++ {
            if false {
                v30 = (v30 + (((v30 + ((v30 + "") + v30)) + (v30 + "a")) + v30))
            } else {
                var v32 uint32 = uint32(v2)
                var v33 float32 = float32(float64(v4))
                _ = v32
                _ = v33
            }
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
            _ = fi31
        }
    case 338:
        fmt.Println(v4, v4, v30)
    default:
        v30 = (v30 + v30)
    }
    v2 = (v2 & 255)
    _ = len(v1)
    v34 := GenS3{F0: (v2 >> ((42) & 63)), F1: -2147483648}
    var v35 int64 = -828
    var v36 uint64 = 0
    v4 += v4
    for fi37 := 0; fi37 < 7; fi37++ {
        v38 := GenS1{F0: 1}
        _ = fi37
        _ = v38
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v39 := GenS3{F0: v2, F1: 0}
    var v40 string = v30
    v41 := GenS3{F0: (v2 | v2), F1: 1}
    v2 += (v2 % ((983) | 1))
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
