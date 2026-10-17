package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint8; F1 string; F2 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]int)
    var v3 int64 = 0
    var v4 uint8 = 2
    var v5 int32 = -359
    v6 := GenS1{F0: v3, F1: v4}
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint = 434
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: v4, F1: "hello", F2: (v3 % (((v3 / ((v3) | 1))) | 1))}
    switch 32767 {
    case -2147483648:
        v10 := GenS3{F0: (true && false)}
        _ = v10
    case -1:
        switch 128 {
        case 858:
            v2[-387] = -943
        case 9223372036854775807:
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 uint16 = 42
                for fi13 := 0; fi13 < 6; fi13++ {
                    v14 := []int{(fi11 / ((2147483647) | 1)), fi11, fi13}
                    _ = fi13
                    _ = v14
                }
                var v15 int8 = -1
                _ = fi11
                _ = v12
                _ = v15
            }
        }
    case -9223372036854775808:
        v1 = 3.14159
    default:
        v16 := []int{0, 9223372036854775807, 408}
        _ = v16
    }
    v17 := []int{-2147483648, 32767, 259}
    switch 455 {
    case 4294967295:
        if false {
            v18 := GenS3{F0: false}
            for fi19 := 0; fi19 < 4; fi19++ {
                var v20 int = fi19
                var v21 uint32 = 2
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = v18
        }
    case -2147483648:
        _ = len(v2)
    case 996:
        if true {
            v4 = (v4 << ((1) & 63))
            var v22 int32 = (v5 % ((v5) | 1))
            _ = v22
        }
    default:
        fmt.Println(v3, v3)
    }
    switch 494 {
    case 9223372036854775807:
        for fi23 := 0; fi23 < 3; fi23++ {
            v5 -= (v5 - (v5 & (v5 << ((42) & 63))))
            if (!true) {
                switch (^(fi23)) {
                case 32767:
                    var v24 float64 = (float64(1e10) + float64(1.5))
                    _ = v24
                }
                _ = len(v2)
                var v25 int32 = v5
                _ = v25
            }
            v26 := GenS1{F0: (v3 & (v3 / ((v3) | 1))), F1: 107}
            _ = fi23
            _ = v26
        }
    case 4294967295:
        _ = len(v2)
    case 32768:
        if (29 <= 0) {
            if (false || false) {
                _ = len(v2)
            } else {
                var v27 uint = 2
                fmt.Println(v1, v5, v5)
                _ = v27
            }
        }
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        v4 = v4
        switch 2147483647 {
        case -731:
            switch "" {
            case "foo bar":
                if (61 == 2) {
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    for fi29 := 0; fi29 < 5; fi29++ {
                        switch "foo bar" {
                        case "foo bar":
                            fmt.Println(v4)
                        default:
                            v30 := GenS1{F0: v3, F1: v4}
                            _ = v30
                        }
                        v4 = v4
                        _ = fi29
                    }
                }
            case "\n\t":
                var v31 float32 = float32(((float64(0.0) - (((v1 + float64((1.5 / 1e10))) / 1.0) + v1)) - v1))
                _ = v31
            default:
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
            }
        }
        _ = fi28
    }
    var v32 int = 502
    for fi33 := 0; fi33 < 17; fi33++ {
        switch "a" {
        case "a":
            var v34 int8 = 127
            _ = v34
        case "":
            var v35 int16 = 32767
            _ = v35
        case "foo bar":
            v36 := GenS2{F0: (v4 - 44), F1: "", F2: v3}
            _ = v36
        default:
            fmt.Println(v4)
        }
        _ = len(v17)
        _ = fi33
    }
    if len(v17) > 0 {
        v17[0] = v32
    }
    fmt.Println(v3)
    var v37 int32 = (v5 + v5)
    switch "foo bar" {
    case "hello":
        v38 := GenS4{F0: false}
        _ = v38
    default:
        var v39 bool = (false && true)
        _ = v39
    }
    for fi40 := 0; fi40 < 18; fi40++ {
        if (!(3.14159 == (-2.25 - float32((v1 / (v1 + ((float64(0.0) + float64(-2.25)) - 1e10))))))) {
            v4 = 72
            var v41 int16 = 42
            _ = v41
        } else {
            var v42 uint32 = 775
            _ = v42
        }
        var v43 int = (v32 >> ((2) & 63))
        v44 := v2[v43]
        _ = fi40
        _ = v43
        _ = v44
    }
    for fi45 := 0; fi45 < 8; fi45++ {
        for fi46 := 0; fi46 < 17; fi46++ {
            var v47 bool = false
            v3 = (-(v3))
            _ = fi46
            _ = v47
        }
        _ = fi45
    }
    fmt.Println(v32, v37, v5)
    v17 = append(v17, v32)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
