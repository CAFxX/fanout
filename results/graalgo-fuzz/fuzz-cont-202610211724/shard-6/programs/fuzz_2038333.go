package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int8 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int64; F1 uint8; F2 int32 }
func genFn0() uint8 {
    if (true && false) {
        var v4 string = ("hello" + "hello")
        _ = v4
    } else {
        if true {
            v5 := []int{-1, 42}
            _ = v5
        }
        if (0.0 >= -2.25) {
            for fi6 := 0; fi6 < 2; fi6++ {
                var v7 uint32 = 42
                _ = fi6
                _ = v7
            }
            var v8 uint64 = 2
            v9 := GenS1{F0: 42, F1: -1, F2: 100}
            _ = v8
            _ = v9
        } else {
            v10 := []int{}
            _ = v10
        }
    }
    var v11 int = 2147483647
    _ = v11
    return 255
}
func genFn1() uint {
    if (1 > 0) {
        v12 := []int{127, 42, -2147483648}
        for fi13 := 0; fi13 < 10; fi13++ {
            if (0 != 60) {
                var v14 int = (fi13 - fi13)
                for fi15 := 0; fi15 < 19; fi15++ {
                    var v16 float64 = 0.0
                    var v17 int32 = 42
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                _ = v14
            }
            for fi18 := 0; fi18 < 1; fi18++ {
                _ = len(v12)
                v19 := GenS2{F0: 42, F1: 42}
                var v20 float32 = 0.0
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi13
        }
        switch (("foo bar" + "a") + ("" + "foo bar")) {
        case "a":
            var v21 int32 = -240
            _ = v21
        }
        _ = v12
    } else {
        var v22 uint = 65535
        v22 += 18446744073709551615
        v23 := GenS1{F0: 2, F1: 127, F2: 1}
        _ = v22
        _ = v23
    }
    var v24 uint16 = 507
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 uint = 255
        var v27 uint16 = (^(v24))
        var v28 string = ""
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    switch 805 {
    case -9223372036854775808:
        if false {
            v24 += 2
            for fi29 := 0; fi29 < 3; fi29++ {
                if false {
                    v30 := GenS2{F0: int8(v24), F1: 1}
                    _ = v30
                } else {
                    if false {
                        v24 += 151
                        var v31 uint8 = 234
                        var v32 int32 = 2147483647
                        _ = v31
                        _ = v32
                    } else {
                        var v33 int16 = 800
                        if true {
                            v34 := GenS2{F0: 75, F1: 150}
                            _ = v34
                        } else {
                            if true {
                                for fi35 := 0; fi35 < 3; fi35++ {
                                    v33 += v33
                                    _ = fi35
                                }
                            }
                            v33 -= -1
                        }
                        for fi36 := 0; fi36 < 10; fi36++ {
                            v24 -= (v24 - 0)
                            var v37 float32 = (((float32(3.14159) * 0.0) * (((((0.1 + (-2.25 / -2.25)) * 3.14159) * -2.25) + -2.25) * 3.14159)) / 1.5)
                            _ = fi36
                            _ = v37
                        }
                        _ = v33
                    }
                }
                v38 := []int{}
                _ = fi29
                _ = v38
            }
        }
    default:
        v24 += 526
    }
    v39 := GenS3{F0: 0, F1: uint8(v24), F2: 2147483647}
    _ = v24
    _ = v39
    return 4294967295
}

func main() {
    var v1 bool = false
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    var v40 uint = 687
    fmt.Println(v1, v40)
    var v41 float32 = -2.25
    var v42 uint32 = 2147483647
    for fi43 := 0; fi43 < 5; fi43++ {
        var v44 int16 = 1000
        v45 := GenS2{F0: -128, F1: v42}
        _ = fi43
        _ = v44
        _ = v45
    }
    v41 += v41
    v46 := []int{128, 0, 712}
    for fi47 := 0; fi47 < 16; fi47++ {
        fmt.Println(v42, v1, fi47)
        fmt.Println(v1, v41)
        for fi48 := 0; fi48 < 1; fi48++ {
            _ = len(v46)
            _ = fi48
        }
        _ = fi47
    }
    fmt.Println(v40, v42, v41)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v46:", len(v46), v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v42
    _ = v46
}
