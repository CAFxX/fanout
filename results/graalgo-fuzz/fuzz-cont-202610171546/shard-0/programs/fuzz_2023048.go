package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint64; F2 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32 }
func genFn0() uint {
    switch -588 {
    case 4294967295:
        var v7 int = 127
        _ = v7
    default:
        for fi8 := 0; fi8 < 6; fi8++ {
            fmt.Println(fi8)
            var v9 int16 = 42
            _ = fi8
            _ = v9
        }
    }
    var v10 uint = 18446744073709551615
    v10 += (v10 ^ 1)
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 int8 = 100
        fmt.Println(fi11, v12, fi11)
        _ = fi11
        _ = v12
    }
    _ = v10
    return (v10 << (((v10 >> (((v10 + (^(v10)))) & 63))) & 63))
}

func main() {
    var v1 string = "foo bar"
    var v2 int64 = 0
    var v3 uint16 = 0
    var v4 int32 = 1
    v5 := make(map[uint32]uint32)
    var v6 uint16 = (v3 << ((2) & 63))
    for fi13 := 0; fi13 < 4; fi13++ {
        v4 += v4
        _ = fi13
    }
    if true {
        if true {
            var v14 int8 = 1
            var v15 float64 = (0.1 / (0.1 / 1.0))
            for fi16 := 0; fi16 < 6; fi16++ {
                if false {
                    v17 := []int{(fi16 % (((fi16 ^ 128)) | 1)), 1, (fi16 & fi16)}
                    _ = v17
                }
                var v18 float32 = 3.14159
                _ = fi16
                _ = v18
            }
            _ = v14
            _ = v15
        }
        var v19 uint64 = 1
        v2 = 9223372036854775807
        _ = v19
    } else {
        v2 += (v2 * 42)
        v4 += v4
    }
    if false {
        switch (v1 + ("foo bar" + v1)) {
        case "":
            v20 := v5[0]
            _ = v20
        }
    } else {
        v21 := GenS1{F0: 68}
        v1 = v1
        var v22 uint64 = 506
        _ = v21
        _ = v22
    }
    var v23 string = v1
    var v24 string = v23
    switch "hello" {
    case "":
        if false {
            var v25 int32 = -744
            v26 := GenS2{F0: genFn0()}
            _ = v25
            _ = v26
        }
    case "a":
        v23 = "foo bar"
    case "\n\t":
        v27 := GenS3{F0: 42, F1: 0, F2: v24}
        _ = v27
    default:
        v2 += (-(v2))
    }
    var v28 uint = 2
    v29 := GenS4{F0: 149}
    var v30 uint16 = (v6 ^ 2)
    if (((^(v4)) != 1) || false) {
        if (!(!true)) {
            v31 := []int{1, 42}
            v2 = (^(v2))
            v23 = ""
            _ = v31
        } else {
            if (-1 != -214) {
                if ((((0.0 + -2.25) - 0.0) / 0.1) == ((-2.25 * 0.1) / (float32(3.14159) + (1.5 / 1.0)))) {
                    if false {
                        var v32 uint32 = 0
                        _ = v32
                    }
                    v30 = v3
                } else {
                    for fi33 := 0; fi33 < 16; fi33++ {
                        var v34 int = 128
                        _ = fi33
                        _ = v34
                    }
                    var v35 int = 609
                    fmt.Println(v2, v35, v6)
                    _ = v35
                }
            } else {
                v28 += (v28 - uint(v6))
                if (!(v2 > (^(v2)))) {
                    for fi36 := 0; fi36 < 16; fi36++ {
                        if (-966 <= 42) {
                            fmt.Println(v23, fi36, fi36)
                            v37 := []int{-1, fi36, (fi36 ^ (fi36 >> ((v28) & 63))), fi36}
                            _ = v37
                        }
                        var v38 int16 = -1
                        var v39 uint8 = 128
                        _ = fi36
                        _ = v38
                        _ = v39
                    }
                    v23 = v1
                }
                v40 := v5[65535]
                _ = v40
            }
        }
        var v41 int16 = 32767
        var v42 uint64 = 619
        _ = v41
        _ = v42
    }
    var v43 int32 = -913
    v44 := GenS5{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
