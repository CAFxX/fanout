package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS4 struct { F0 uint16; F1 uint32; F2 string }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 int32; F1 int8 }
func genFn0(p0 bool) uint {
    v4 := GenS1{F0: -125}
    var v5 int64 = -773
    v6 := GenS1{F0: 100}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = false
    v3 := []uint32{891, 305, 285}
    if true {
        fmt.Println(v2)
        for fi7 := 0; fi7 < 15; fi7++ {
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 float32 = 3.14159
                v9 = 0.0
                v10 := GenS1{F0: -1}
                _ = fi8
                _ = v10
                _ = v9
            }
            var v11 uint64 = 4294967295
            _ = fi7
            _ = v11
        }
    }
    v12 := GenS2{F0: "hello"}
    var v13 int64 = 0
    if (2147483648 != 463) {
        var v14 int16 = 870
        _ = v14
    } else {
        if v2 {
            v2 = (v2 && (!false))
            v15 := GenS2{F0: "\n\t"}
            _ = v15
        } else {
            var v16 float32 = (-2.25 * 0.0)
            var v17 float32 = v16
            _ = v16
            _ = v17
        }
        if (v13 != v13) {
            v2 = (127 == 255)
            var v18 bool = (32767 > 32767)
            v19 := GenS3{F0: 483, F1: (-9223372036854775808 == 9223372036854775807), F2: 32767}
            _ = v18
            _ = v19
        } else {
            var v20 uint16 = 65535
            _ = v20
        }
    }
    if (3.14159 != (1e10 * -2.25)) {
        v21 := GenS3{F0: 29, F1: (v2 || (0 != 128)), F2: 959}
        if v2 {
            switch 1 {
            case 128:
                var v22 uint16 = 32767
                _ = v22
            case 0:
                if len(v3) > 0 {
                    v3[0] = 65535
                }
            case 1:
                v23 := GenS4{F0: 291, F1: 2147483648, F2: "hello"}
                _ = v23
            }
        }
        _ = v21
    } else {
        var v24 float32 = 1e10
        v25 := GenS1{F0: 1}
        _ = v24
        _ = v25
    }
    var v26 int8 = -49
    v27 := GenS5{F0: 861, F1: "foo bar"}
    for fi28 := 0; fi28 < 4; fi28++ {
        if (v2 || (42 < 626)) {
            for fi29 := 0; fi29 < 13; fi29++ {
                if true {
                    var v30 bool = (0.0 < 1.5)
                    fmt.Println(fi28, v26, v30)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(fi29, v13)
                    v2 = (113 <= 2147483647)
                    _ = v30
                }
                v26 = 42
                _ = fi29
            }
        }
        _ = fi28
    }
    if (1 != 65535) {
        for fi31 := 0; fi31 < 1; fi31++ {
            fmt.Println(fi31, v2)
            _ = fi31
        }
    }
    if false {
        var v32 int8 = 1
        for fi33 := 0; fi33 < 12; fi33++ {
            fmt.Println(fi33)
            _ = fi33
        }
        if v2 {
            _ = len(v3)
            for fi34 := 0; fi34 < 8; fi34++ {
                v35 := v1[2147483648]
                var v36 int32 = 0
                _ = fi34
                _ = v35
                _ = v36
            }
        }
        _ = v32
    }
    var v37 int64 = (-(v13))
    for fi38 := 0; fi38 < 3; fi38++ {
        var v39 int = fi38
        var v40 int16 = 0
        _ = fi38
        _ = v39
        _ = v40
    }
    if ("hello" == "") {
        if (((1e10 * 0.0) + ((1.5 + 0.0) + ((float64((-2.25 / (float32(0.1) * 1.5))) - 3.14159) - float64(-2.25)))) <= 0.0) {
            for fi41 := 0; fi41 < 4; fi41++ {
                v42 := GenS6{F0: 1, F1: (^(v26))}
                _ = fi41
                _ = v42
            }
            v37 = v13
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v37
}
