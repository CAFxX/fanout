package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int; F1 int16; F2 int }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 string; F1 uint8 }

func main() {
    var v1 int = -986
    var v2 int16 = -32768
    v3 := make(map[int]string)
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]uint32)
    if (65535 != 2147483647) {
        v6 := []int{}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := v3[v1]
            v8 = (v8 + ("hello" + ""))
            for fi9 := 0; fi9 < 15; fi9++ {
                if true {
                    v10 := GenS1{F0: (v1 + fi7), F1: -2147483648}
                    v1 = fi7
                    v11 := GenS1{F0: (fi9 & fi7), F1: -518}
                    _ = v10
                    _ = v11
                }
                fmt.Println(fi7)
                v8 = ""
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
        var v12 uint32 = 331
        _ = v12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        v2 += 32767
        if (2147483647 <= 0) {
            v2 += (-(v2))
        }
        _ = fi13
    }
    var v14 uint = 255
    var v15 uint8 = 128
    _ = len(v3)
    v1 += v1
    if false {
        for fi16 := 0; fi16 < 11; fi16++ {
            v2 += v2
            v17 := GenS1{F0: (v1 / (((fi16 % (((fi16 + v1)) | 1))) | 1)), F1: 1}
            _ = fi16
            _ = v17
        }
        fmt.Println(v1, v2)
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        for fi19 := 0; fi19 < 8; fi19++ {
            v15 = v15
            _ = fi19
        }
        v15 = v15
        var v20 uint32 = 350
        _ = fi18
        _ = v20
    }
    if false {
        if true {
            for fi21 := 0; fi21 < 1; fi21++ {
                v14 -= v14
                switch "hello" {
                case "hello":
                    switch v1 {
                    case 1:
                        if ((((1e10 - (float32((1e10 * 1e10)) - 1.5)) + -2.25) * ((0.1 + (3.14159 + 0.1)) + 1.5)) != float32(1e10)) {
                            if (-340 != 42) {
                                fmt.Println(v2, v14, fi21)
                                v22 := GenS2{F0: 4294967295, F1: -400, F2: 42}
                                v23 := []int{}
                                _ = v22
                                _ = v23
                            } else {
                                v24 := GenS1{F0: v1, F1: -2147483648}
                                var v25 int = (v1 >> ((581) & 63))
                                _ = v24
                                _ = v25
                            }
                            var v26 float32 = ((((float32((0.1 - ((-2.25 * 0.0) * 0.1))) + (0.0 * (-2.25 * 0.1))) * float32(-2.25)) / float32((float64(-2.25) / -2.25))) / 1.0)
                            if true {
                                for fi27 := 0; fi27 < 1; fi27++ {
                                    v14 = 2
                                    fmt.Println(v2)
                                    var v28 int32 = 626
                                    _ = fi27
                                    _ = v28
                                }
                            } else {
                                var v29 uint = (^(v14))
                                var v30 int32 = -1
                                v31 := GenS1{F0: fi21, F1: v30}
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                            _ = v26
                        } else {
                            fmt.Println(v15, fi21, v1)
                        }
                    default:
                        v14 += v14
                    }
                case "foo bar":
                    v32 := GenS3{F0: -9223372036854775808, F1: "\n\t"}
                    _ = v32
                case "a":
                    v33 := []int{}
                    _ = v33
                }
                _ = fi21
            }
            v34 := []int{}
            var v35 int = int(v15)
            _ = v34
            _ = v35
        } else {
            for fi36 := 0; fi36 < 1; fi36++ {
                var v37 uint32 = 42
                _ = fi36
                _ = v37
            }
            v2 += 1
            v15 -= v15
        }
        fmt.Println(v14)
    } else {
        var v38 bool = false
        v39 := []int{}
        _ = v38
        _ = v39
    }
    for fi40 := 0; fi40 < 11; fi40++ {
        _ = len(v5)
        _ = fi40
    }
    var v41 int8 = 126
    var v42 int8 = v41
    v43 := v4[0]
    for fi44 := 0; fi44 < 14; fi44++ {
        fmt.Println(v14, v2)
        _ = fi44
    }
    for fi45 := 0; fi45 < 1; fi45++ {
        for fi46 := 0; fi46 < 4; fi46++ {
            v15 -= v15
            _ = fi46
        }
        for fi47 := 0; fi47 < 15; fi47++ {
            v2 += v2
            v48 := []int{42, (v1 ^ fi45), 127, int(v14)}
            _ = fi47
            _ = v48
        }
        v49 := GenS4{F0: "foo bar", F1: 2}
        _ = fi45
        _ = v49
    }
    var v50 int64 = 5
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v50
}
