package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 int }

func main() {
    v1 := []int{32768}
    var v2 bool = true
    var v3 int = 9223372036854775807
    var v4 int = v3
    for fi5 := 0; fi5 < 17; fi5++ {
        switch "a" {
        case "hello":
            var v6 bool = true
            _ = v6
        case "foo bar":
            var v7 int16 = -264
            _ = v7
        default:
            if v2 {
                v1 = append(v1, -9223372036854775808)
                var v8 int8 = -1
                _ = v8
            } else {
                var v9 uint32 = 65535
                v3 = (fi5 & (fi5 * fi5))
                _ = v9
            }
        }
        var v10 string = ""
        _ = fi5
        _ = v10
    }
    _ = len(v1)
    switch (^(v3)) {
    case 32768:
        switch "a" {
        case "\n\t":
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 string = "\n\t"
                _ = fi11
                _ = v12
            }
        case "hello":
            fmt.Println(v3, v3)
        default:
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 uint8 = 1
                _ = fi13
                _ = v14
            }
        }
    }
    if v2 {
        _ = len(v1)
    } else {
        var v15 int8 = 0
        fmt.Println(v15)
        _ = v15
    }
    if v2 {
        var v16 uint16 = 178
        var v17 int64 = 0
        var v18 uint32 = 4294967295
        _ = v16
        _ = v17
        _ = v18
    }
    v4 = v4
    if (!(false && (v3 <= v3))) {
        if true {
            switch v4 {
            case 42:
                v3 = (v3 % ((-1) | 1))
            case 127:
                var v19 uint8 = 0
                _ = v19
            case -1:
                if true {
                    if v2 {
                        if (!v2) {
                            var v20 int16 = 1000
                            v21 := GenS1{F0: v2, F1: 2}
                            for fi22 := 0; fi22 < 17; fi22++ {
                                for fi23 := 0; fi23 < 7; fi23++ {
                                    var v24 int32 = -1
                                    _ = fi23
                                    _ = v24
                                }
                                _ = fi22
                            }
                            _ = v20
                            _ = v21
                        } else {
                            var v25 uint32 = 2
                            v2 = (!false)
                            if (uint16(v25) > 2) {
                                var v26 int8 = 0
                                _ = v26
                            }
                            _ = v25
                        }
                        v3 = (v4 / ((v4) | 1))
                    } else {
                        fmt.Println(v4)
                    }
                    v2 = v2
                    _ = len(v1)
                }
            default:
                var v27 int = 29
                _ = v27
            }
            if (291 < 1) {
                var v28 int16 = 0
                v29 := GenS1{F0: v2, F1: 42}
                var v30 uint = 185
                _ = v28
                _ = v29
                _ = v30
            }
        } else {
            fmt.Println(v2, v3)
            switch (v4 >> ((65535) & 63)) {
            case 1:
                var v31 float64 = 0.0
                _ = v31
            case 42:
                var v32 int64 = -722
                _ = v32
            default:
                for fi33 := 0; fi33 < 12; fi33++ {
                    for fi34 := 0; fi34 < 4; fi34++ {
                        v35 := GenS2{F0: 376, F1: 4294967295}
                        fmt.Println(v4)
                        if false {
                            v1 = append(v1, 837)
                            var v36 uint16 = 32767
                            v2 = v2
                            if v2 {
                                v4 -= 0
                            }
                            _ = v36
                        } else {
                            if true {
                                v37 := GenS2{F0: 701, F1: 2}
                                v38 := GenS3{F0: fi33}
                                var v39 uint16 = 2
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                        }
                        _ = fi34
                        _ = v35
                    }
                    _ = fi33
                }
            }
        }
        var v40 float64 = 0.1
        v4 = v4
        _ = v40
    } else {
        switch v3 {
        case 128:
            if false {
                v41 := GenS1{F0: (1 == 0), F1: 2}
                fmt.Println(v4)
                switch (v4 - (-(v3))) {
                case 953:
                    fmt.Println(v3, v4)
                case 730:
                    var v42 float32 = 0.1
                    _ = v42
                }
                _ = v41
            } else {
                var v43 float32 = float32(3.14159)
                _ = v43
            }
        default:
            switch v4 {
            case -914:
                v44 := GenS2{F0: 1, F1: 42}
                _ = v44
            default:
                v2 = false
            }
        }
    }
    if true {
        var v45 bool = (2 > 127)
        var v46 int64 = 42
        _ = v45
        _ = v46
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
