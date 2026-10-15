package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint16; F1 string; F2 uint16 }
type GenS3 struct { F0 uint64; F1 bool; F2 int }
func genFn0() uint32 {
    var v4 int32 = 1
    v5 := []int{9223372036854775807, -208, 127, 936}
    _ = v4
    _ = v5
    return 2147483647
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 255
    fmt.Println(v1, v3)
    if (3.14159 <= 3.14159) {
        v1 = v1
        v1 -= (v1 + v1)
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 uint64 = 2
            fmt.Println(fi6, fi6)
            _ = len(v2)
            _ = fi6
            _ = v7
        }
    } else {
        fmt.Println(v3, v1)
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        if (!false) {
            var v9 uint64 = 1
            fmt.Println(fi8)
            v10 := []int{fi8}
            _ = v10
            _ = v9
        } else {
            if (1 != 1) {
                for fi11 := 0; fi11 < 3; fi11++ {
                    fmt.Println(fi8)
                    _ = fi11
                }
            } else {
                v1 = v1
            }
            v12 := []int{920, (fi8 << ((v1) & 63))}
            _ = v12
        }
        _ = fi8
    }
    v1 = (^(v1))
    v13 := GenS1{F0: 65535, F1: 526, F2: 0}
    if false {
        if false {
            if true {
                for fi14 := 0; fi14 < 8; fi14++ {
                    v3 -= 0
                    fmt.Println(v3, v1, fi14)
                    _ = fi14
                }
                v1 = 255
                if (106 >= 1) {
                    switch -586 {
                    case -9223372036854775808:
                        for fi15 := 0; fi15 < 6; fi15++ {
                            var v16 uint16 = 2
                            var v17 int64 = -535
                            for fi18 := 0; fi18 < 6; fi18++ {
                                var v19 string = (("hello" + "\n\t") + "foo bar")
                                _ = fi18
                                _ = v19
                            }
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    case 128:
                        if true {
                            v20 := GenS1{F0: 557, F1: 1, F2: 18446744073709551615}
                            var v21 int32 = 1
                            v21 = 2147483647
                            _ = v20
                            _ = v21
                        } else {
                            fmt.Println(v1)
                        }
                    }
                } else {
                    for fi22 := 0; fi22 < 4; fi22++ {
                        var v23 float64 = (0.0 * (((((1.5 * 0.0) - float64(((float32(float64(-2.25)) * 0.0) - 0.1))) / -2.25) * -2.25) / 1.0))
                        _ = fi22
                        _ = v23
                    }
                }
            } else {
                var v24 int32 = -1
                _ = v24
            }
            _ = len(v2)
            var v25 string = (("\n\t" + ((("hello" + ("hello" + ("foo bar" + "hello"))) + "foo bar") + "foo bar")) + ("hello" + ""))
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        switch "\n\t" {
        case "hello":
            var v27 uint8 = (^(v3))
            _ = v27
        default:
            for fi28 := 0; fi28 < 8; fi28++ {
                if true {
                    v29 := GenS2{F0: 218, F1: "hello", F2: 32767}
                    fmt.Println(fi28)
                    v30 := GenS3{F0: 949, F1: (true || false), F2: -2147483648}
                    _ = v29
                    _ = v30
                } else {
                    if true {
                        switch "a" {
                        case "hello":
                            v31 := v2[genFn0()]
                            _ = v31
                        default:
                            fmt.Println(v3)
                        }
                    } else {
                        var v32 int = (fi26 + (^(fi26)))
                        var v33 float32 = 0.1
                        v34 := GenS1{F0: 0, F1: 941, F2: 53}
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    v35 := GenS1{F0: 42, F1: 0, F2: 840}
                    _ = v35
                }
                _ = fi28
            }
        }
        _ = fi26
    }
    var v36 float32 = (3.14159 * 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v36
}
