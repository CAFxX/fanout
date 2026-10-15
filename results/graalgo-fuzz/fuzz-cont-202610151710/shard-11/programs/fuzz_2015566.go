package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 uint64; F2 bool }
type GenS5 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS6 struct { F0 int64 }

func main() {
    var v1 int16 = 872
    var v2 uint8 = 138
    var v3 string = "foo bar"
    v4 := make(map[uint32]string)
    v5 := make(map[string]uint32)
    fmt.Println(v2, v2, v3)
    v2 = v2
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 float64 = 0.0
        v1 += 801
        v1 -= v1
        _ = fi6
        _ = v7
    }
    var v8 float32 = 3.14159
    v9 := GenS1{F0: v2, F1: false}
    var v10 float64 = 1e10
    if (float64(v8) != -2.25) {
        var v11 uint64 = 4294967295
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 uint32 = 2147483648
            if false {
                var v14 bool = (239 < -2147483648)
                v15 := GenS2{F0: (v11 & v11)}
                _ = v14
                _ = v15
            } else {
                for fi16 := 0; fi16 < 14; fi16++ {
                    fmt.Println(v2, fi16, v1)
                    v2 -= v2
                    var v17 uint32 = (v13 << ((701) & 63))
                    switch "a" {
                    case "a":
                        v10 = ((1.5 * v10) - ((0.0 - v10) + v10))
                    case "foo bar":
                        v18 := GenS2{F0: uint64(v2)}
                        _ = v18
                    }
                    _ = fi16
                    _ = v17
                }
                v19 := GenS1{F0: 0, F1: true}
                _ = v19
            }
            switch int(v1) {
            case 4294967295:
                _ = len(v4)
            default:
                var v20 string = (v3 + v3)
                _ = v20
            }
            _ = fi12
            _ = v13
        }
        _ = v11
    } else {
        switch v3 {
        case "a":
            v21 := GenS1{F0: (^(v2)), F1: true}
            _ = v21
        case "":
            v22 := v4[2]
            _ = v22
        case "foo bar":
            v23 := GenS3{F0: 2}
            _ = v23
        }
        var v24 int = 9223372036854775807
        switch (v24 >> ((18446744073709551615) & 63)) {
        case -587:
            var v25 uint16 = 624
            _ = v25
        case -2147483648:
            fmt.Println(v2, v3, v24)
        default:
            switch v3 {
            case "hello":
                v26 := GenS4{F0: 668, F1: 42, F2: ((^(v2)) != v2)}
                _ = v26
            case "a":
                v27 := GenS5{F0: (v2 | v2), F1: 2, F2: v1}
                _ = v27
            default:
                v28 := []int{(v24 >> ((750) & 63)), int(v1), (v24 & v24)}
                _ = v28
            }
        }
        _ = v24
    }
    v10 = v10
    v1 -= -32768
    var v29 int16 = -1
    _ = len(v5)
    switch "foo bar" {
    case "hello":
        switch -643 {
        case 42:
            if true {
                if (!false) {
                    fmt.Println(v29)
                    if (0 == 362) {
                        v30 := GenS5{F0: v2, F1: 32767, F2: (v1 + v1)}
                        _ = v30
                    } else {
                        v8 -= -2.25
                    }
                    delete(v5, (v3 + (v3 + v3)))
                }
            } else {
                v1 = -510
                var v31 bool = false
                fmt.Println(v8, v2, v3)
                _ = v31
            }
        case -329:
            v1 = 42
        default:
            for fi32 := 0; fi32 < 12; fi32++ {
                var v33 float32 = v8
                _ = fi32
                _ = v33
            }
        }
    case "":
        for fi34 := 0; fi34 < 8; fi34++ {
            fmt.Println(fi34, v1)
            _ = fi34
        }
    case "\n\t":
        for fi35 := 0; fi35 < 5; fi35++ {
            v8 = float32((float64(0.1) + ((0.0 + (0.0 * float64(1e10))) + -2.25)))
            v5[(v3 + ((v3 + v3) + v3))] = 254
            var v36 uint = 42
            _ = fi35
            _ = v36
        }
    default:
        v37 := GenS1{F0: v2, F1: false}
        _ = v37
    }
    v1 -= v1
    v3 = ("foo bar" + v3)
    v38 := GenS6{F0: -1}
    var v39 uint8 = v2
    if true {
        for fi40 := 0; fi40 < 1; fi40++ {
            var v41 int32 = -1
            var v42 float32 = v8
            _ = fi40
            _ = v41
            _ = v42
        }
        v43 := GenS5{F0: (v2 - v2), F1: 65535, F2: v1}
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
