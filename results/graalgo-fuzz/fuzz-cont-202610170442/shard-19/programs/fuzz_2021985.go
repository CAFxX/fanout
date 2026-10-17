package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 bool }
type GenS2 struct { F0 int32; F1 bool; F2 uint16 }

func main() {
    var v1 uint16 = 0
    v2 := []string{}
    var v3 float32 = (float32((1.5 * 0.1)) / 1.0)
    var v4 int = 32768
    v5 := make(map[uint32]uint32)
    var v6 uint16 = 343
    v7 := GenS1{F0: 1, F1: 128, F2: false}
    v1 += 32767
    var v8 float64 = -2.25
    fmt.Println(v4, v4)
    switch (v4 | (v4 >> ((255) & 63))) {
    case 4294967295:
        _ = len(v2)
    default:
        var v9 float32 = v3
        _ = v9
    }
    v5[467] = 2147483648
    v3 += -2.25
    var v10 bool = ((!(0 >= 0)) || false)
    var v11 float64 = 0.0
    var v12 uint32 = 0
    for fi13 := 0; fi13 < 15; fi13++ {
        switch 32768 {
        case 127:
            if ((2 != uint64(v12)) || ("hello" != (("hello" + (("foo bar" + "\n\t") + "hello")) + "hello"))) {
                var v14 uint64 = 1
                v1 = (v6 >> ((521) & 63))
                v12 = (v12 & v12)
                _ = v14
            } else {
                v12 -= (^(v12))
                for fi15 := 0; fi15 < 9; fi15++ {
                    _ = len(v5)
                    _ = fi15
                }
                var v16 uint32 = (v12 ^ (^(v12)))
                _ = v16
            }
        default:
            v10 = (!((114 < v1) && false))
        }
        v1 -= (v1 - (v1 >> ((1) & 63)))
        _ = fi13
    }
    var v17 uint16 = 42
    for fi18 := 0; fi18 < 2; fi18++ {
        if v10 {
            if (2147483647 != -2147483648) {
                for fi19 := 0; fi19 < 4; fi19++ {
                    var v20 int64 = 535
                    v5[(v12 * 51)] = v12
                    switch (v4 + (-(v4))) {
                    case 2147483647:
                        _ = len(v5)
                    case 9223372036854775807:
                        fmt.Println(v8, v8)
                    default:
                        var v21 int16 = 0
                        _ = v21
                    }
                    _ = fi19
                    _ = v20
                }
                v22 := GenS1{F0: 42, F1: 2, F2: ((100 == -128) || false)}
                _ = v22
            }
        } else {
            _ = len(v2)
            for fi23 := 0; fi23 < 15; fi23++ {
                if false {
                    v24 := GenS1{F0: -947, F1: 42, F2: false}
                    _ = v24
                }
                _ = fi23
            }
        }
        for fi25 := 0; fi25 < 19; fi25++ {
            switch (^(fi18)) {
            case -797:
                fmt.Println(v12)
                var v26 float32 = v3
                v12 -= (v12 >> ((42) & 63))
                if len(v2) > 0 {
                    v2[0] = (("foo bar" + "foo bar") + ("hello" + "foo bar"))
                }
                v27 := GenS1{F0: 465, F1: 42, F2: v10}
                _ = v26
                _ = v27
            default:
                _ = len(v5)
                _ = len(v5)
            }
            _ = fi25
        }
        _ = fi18
    }
    if v10 {
        for fi28 := 0; fi28 < 8; fi28++ {
            for fi29 := 0; fi29 < 2; fi29++ {
                if (!(v6 <= 2)) {
                    var v30 uint32 = 2147483648
                    var v31 int64 = 204
                    fmt.Println(v6, v17)
                    _ = v30
                    _ = v31
                } else {
                    var v32 float32 = (v3 / 0.1)
                    switch v4 {
                    case -704:
                        v33 := GenS2{F0: 866, F1: ((v32 * (3.14159 - v3)) <= (((float32(v8) * (v32 - 3.14159)) / float32(v8)) / float32(3.14159))), F2: 217}
                        _ = v33
                    case 127:
                        v34 := GenS1{F0: 1, F1: 255, F2: true}
                        _ = v34
                    case 1:
                        switch (v4 & fi28) {
                        case 9223372036854775807:
                            switch "foo bar" {
                            case "\n\t":
                                v3 = (-2.25 + (float32(0.1) * v32))
                            default:
                                v5[(v12 ^ (v12 ^ (v12 - (v12 ^ v12))))] = (^(v12))
                            }
                        case -512:
                            if len(v2) > 0 {
                                v2[0] = "a"
                            }
                        }
                    default:
                        v17 = (v1 - v6)
                    }
                    _ = v32
                }
                _ = fi29
            }
            var v35 uint = uint(v6)
            _ = fi28
            _ = v35
        }
        v10 = v10
    }
    for fi36 := 0; fi36 < 18; fi36++ {
        if v10 {
            fmt.Println(v3, v6, v11)
            var v37 int8 = int8(fi36)
            _ = v37
        } else {
            switch ("hello" + ("\n\t" + "a")) {
            case "hello":
                var v38 float32 = (-2.25 * -2.25)
                _ = v38
            case "foo bar":
                v39 := GenS2{F0: 1, F1: false, F2: 42}
                _ = v39
            default:
                var v40 uint16 = (v1 * (v1 * (v1 - 0)))
                _ = v40
            }
            if v10 {
                var v41 int8 = 100
                _ = v41
            }
        }
        var v42 uint8 = 2
        v17 += 0
        _ = fi36
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
