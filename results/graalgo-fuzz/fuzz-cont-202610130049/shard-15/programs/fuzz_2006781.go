package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 float64 = 0.1
    var v2 string = "a"
    var v3 uint64 = 2
    v4 := make(map[uint32]int)
    var v5 uint8 = 230
    v1 -= 3.14159
    if false {
        var v6 int16 = 32767
        if false {
            var v7 uint32 = 2147483648
            for fi8 := 0; fi8 < 12; fi8++ {
                if (-1 == 127) {
                    v9 := GenS1{F0: (fi8 | fi8)}
                    _ = v9
                } else {
                    var v10 int32 = 42
                    _ = v10
                }
                fmt.Println(v1, v5)
                _ = fi8
            }
            _ = v7
        }
        _ = v6
    } else {
        for fi11 := 0; fi11 < 19; fi11++ {
            fmt.Println(v3, v3, v5)
            v5 += v5
            fmt.Println(fi11, v2)
            _ = fi11
        }
        var v12 float32 = float32(v1)
        var v13 uint = 255
        _ = v12
        _ = v13
    }
    switch 2147483647 {
    case 366:
        v4[4294967295] = 128
    default:
        fmt.Println(v1)
    }
    v5 -= (^(v5))
    for fi14 := 0; fi14 < 4; fi14++ {
        switch "a" {
        case "\n\t":
            var v15 float64 = (-2.25 / ((v1 * v1) * (v1 + 1e10)))
            _ = v15
        case "hello":
            if false {
                for fi16 := 0; fi16 < 20; fi16++ {
                    fmt.Println(v2)
                    _ = fi16
                }
                var v17 string = (v2 + v2)
                _ = v17
            }
        }
        _ = fi14
    }
    v5 += v5
    for fi18 := 0; fi18 < 4; fi18++ {
        v19 := GenS1{F0: (fi18 >> ((42) & 63))}
        v3 += v3
        if false {
            v20 := []int{}
            if true {
                if true {
                    v21 := GenS1{F0: (fi18 * fi18)}
                    var v22 uint16 = 255
                    var v23 float32 = ((-2.25 / 0.1) + float32(v1))
                    _ = v21
                    _ = v22
                    _ = v23
                }
                var v24 string = "foo bar"
                var v25 int64 = int64(v5)
                _ = v24
                _ = v25
            } else {
                _ = len(v20)
                v26 := v4[2]
                _ = v26
            }
            v27 := GenS2{F0: 2147483648, F1: v5}
            _ = v20
            _ = v27
        } else {
            v1 = (1e10 / (v1 * float64(0.1)))
        }
        _ = fi18
        _ = v19
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        v29 := []int{(fi28 % ((fi28) | 1)), (fi28 << ((42) & 63)), (fi28 | (fi28 * 42))}
        switch fi28 {
        case -391:
            switch v2 {
            case "foo bar":
                v3 = 2
            case "hello":
                fmt.Println(v5)
            case "\n\t":
                var v30 float32 = 0.1
                _ = v30
            default:
                v5 -= (v5 << ((18446744073709551615) & 63))
            }
        case -9223372036854775808:
            for fi31 := 0; fi31 < 2; fi31++ {
                var v32 int8 = int8(fi31)
                if true {
                    v3 -= (^(v3))
                } else {
                    var v33 int32 = 2147483647
                    _ = v33
                }
                _ = fi31
                _ = v32
            }
        }
        if (false || (true && (!(true || false)))) {
            if true {
                var v34 int = fi28
                for fi35 := 0; fi35 < 13; fi35++ {
                    var v36 uint16 = 65535
                    var v37 int = 876
                    v2 = (v2 + ("\n\t" + v2))
                    switch v2 {
                    case "hello":
                        var v38 int32 = 1
                        _ = v38
                    case "foo bar":
                        v39 := v4[1]
                        _ = v39
                    }
                    v36 -= (v36 / ((v36) | 1))
                    v1 = (v1 - (3.14159 - float64((float32(float64(1.5)) / 1.5))))
                    fmt.Println(v37, v37)
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                _ = v34
            } else {
                var v40 uint64 = (v3 >> ((255) & 63))
                _ = v40
            }
            var v41 int8 = 0
            v42 := GenS1{F0: int(v3)}
            _ = v41
            _ = v42
        } else {
            for fi43 := 0; fi43 < 6; fi43++ {
                if len(v29) > 0 {
                    v29[0] = (fi28 * (fi43 ^ 127))
                }
                _ = fi43
            }
            for fi44 := 0; fi44 < 15; fi44++ {
                _ = len(v29)
                v2 = (v2 + "foo bar")
                v45 := v4[1]
                _ = fi44
                _ = v45
            }
            v5 = v5
        }
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
