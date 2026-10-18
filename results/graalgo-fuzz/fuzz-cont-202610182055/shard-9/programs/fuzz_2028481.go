package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint16 = 0
    var v2 int16 = 543
    var v3 uint32 = 846
    var v4 string = "a"
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 float64 = 0.1
        fmt.Println(fi5, v3, v1)
        if (uint8(v3) >= 255) {
            var v7 int16 = 624
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            if (-13 == 42) {
                switch "hello" {
                case "foo bar":
                    switch (fi9 >> ((214) & 63)) {
                    case -748:
                        var v10 string = v4
                        _ = v10
                    case -497:
                        v11 := GenS1{F0: v3, F1: (v1 | v1)}
                        _ = v11
                    default:
                        if (true || (true || false)) {
                            v12 := []int{}
                            _ = v12
                        } else {
                            fmt.Println(fi8, v4, v2)
                        }
                    }
                case "a":
                    if false {
                        var v13 int64 = 42
                        fmt.Println(v4, v2)
                        var v14 uint16 = (v1 + v1)
                        _ = v13
                        _ = v14
                    } else {
                        if (!true) {
                            fmt.Println(v1)
                            v3 = (v3 * v3)
                            var v15 int = int(v3)
                            _ = v15
                        } else {
                            fmt.Println(v3, fi9, v4)
                        }
                    }
                default:
                    switch fi9 {
                    case -62:
                        fmt.Println(v2, v4)
                    case -651:
                        fmt.Println(v3)
                    }
                }
                var v16 int16 = -32768
                if false {
                    v17 := GenS1{F0: (v3 >> ((4294967295) & 63)), F1: v1}
                    var v18 int8 = -1
                    var v19 int8 = 122
                    v3 = (v3 >> ((0) & 63))
                    _ = v17
                    _ = v18
                    _ = v19
                }
                if false {
                    v2 = -32768
                    v2 = v2
                    var v20 float32 = float32(((1.5 - (-2.25 * 0.1)) * -2.25))
                    _ = v20
                } else {
                    fmt.Println(v16)
                }
                v21 := GenS1{F0: v3, F1: uint16(v3)}
                _ = v16
                _ = v21
            } else {
                v4 = "a"
                v22 := GenS1{F0: uint32(v1), F1: v1}
                var v23 string = v4
                _ = v22
                _ = v23
            }
            _ = fi9
        }
        _ = fi8
    }
    var v24 bool = true
    var v25 float64 = (0.1 / 0.1)
    var v26 int32 = 0
    v27 := []int{int(v26), 42}
    v26 = -2147483648
    if (((4294967295 != 18446744073709551615) && (42 <= -697)) && (false && ("foo bar" != v4))) {
        v25 = 3.14159
    } else {
        v25 += 0.0
        var v28 uint32 = 2147483648
        if (189 > (v28 / ((v3) | 1))) {
            v29 := GenS1{F0: (v28 - 867), F1: v1}
            _ = v29
        } else {
            if len(v27) >= 2 {
                v27 = v27[1:len(v27)]
            }
            switch 0 {
            case 127:
                v30 := GenS1{F0: v28, F1: 680}
                _ = v30
            case 9223372036854775807:
                for fi31 := 0; fi31 < 2; fi31++ {
                    v32 := GenS1{F0: v3, F1: (v1 | v1)}
                    if (200 == 255) {
                        var v33 uint64 = 166
                        _ = v33
                    } else {
                        fmt.Println(fi31)
                        if len(v27) > 0 {
                            v27[0] = fi31
                        }
                        switch 127 {
                        case 4294967295:
                            var v34 int64 = 423
                            _ = v34
                        default:
                            var v35 int = fi31
                            _ = v35
                        }
                    }
                    var v36 string = "foo bar"
                    _ = fi31
                    _ = v32
                    _ = v36
                }
            }
            v37 := GenS2{F0: 42}
            _ = v37
        }
        _ = v28
    }
    if v24 {
        v38 := GenS1{F0: v3, F1: 42}
        for fi39 := 0; fi39 < 16; fi39++ {
            var v40 string = ""
            v40 = v4
            _ = fi39
            _ = v40
        }
        _ = v38
    } else {
        var v41 uint = 65535
        fmt.Println(v4, v2)
        var v42 int32 = v26
        _ = v41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
