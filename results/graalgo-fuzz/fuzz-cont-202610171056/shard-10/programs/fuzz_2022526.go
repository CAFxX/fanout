package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int16; F1 int32; F2 int64 }

func main() {
    var v1 bool = true
    v2 := make(map[uint32]uint32)
    v3 := []int{32768, 0, -9223372036854775808, 32768}
    var v4 int8 = 0
    var v5 int64 = 586
    if v1 {
        v6 := GenS1{F0: 255}
        var v7 uint8 = 0
        if true {
            var v8 float32 = 0.0
            _ = v8
        }
        _ = v6
        _ = v7
    } else {
        if v1 {
            var v9 uint16 = 255
            _ = v9
        } else {
            fmt.Println(v5)
            var v10 uint = 18446744073709551615
            var v11 int16 = -372
            _ = v10
            _ = v11
        }
        switch 127 {
        case 731:
            v4 -= (v4 << ((4294967295) & 63))
        case 42:
            v12 := GenS1{F0: 255}
            _ = v12
        case -2147483648:
            fmt.Println(v1)
        default:
            v13 := GenS2{F0: 1, F1: 0}
            _ = v13
        }
    }
    var v14 uint64 = 1
    fmt.Println(v5, v5, v1)
    _ = len(v2)
    fmt.Println(v1, v4, v4)
    v4 = v4
    for fi15 := 0; fi15 < 10; fi15++ {
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 uint32 = 2
            v18 := GenS1{F0: 1}
            if v1 {
                v17 = 0
            } else {
                fmt.Println(v1, v1)
                v19 := v2[340]
                var v20 int = (fi15 << ((18446744073709551615) & 63))
                var v21 int8 = v4
                var v22 int32 = 2147483647
                _ = v19
                _ = v20
                _ = v21
                _ = v22
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        for fi23 := 0; fi23 < 18; fi23++ {
            if true {
                switch "foo bar" {
                case "hello":
                    v5 = v5
                case "":
                    v4 += v4
                }
                fmt.Println(v1, fi15)
            }
            _ = fi23
        }
        _ = fi15
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        fmt.Println(fi24, fi24, v1)
        v3 = append(v3, (^(fi24)))
        fmt.Println(v1, v5)
        _ = fi24
    }
    if (v1 && v1) {
        fmt.Println(v5)
        for fi25 := 0; fi25 < 12; fi25++ {
            for fi26 := 0; fi26 < 12; fi26++ {
                if len(v3) > 0 {
                    v3[0] = (fi26 | fi25)
                }
                _ = fi26
            }
            _ = fi25
        }
    } else {
        fmt.Println(v14, v4, v1)
        for fi27 := 0; fi27 < 16; fi27++ {
            var v28 int64 = v5
            for fi29 := 0; fi29 < 4; fi29++ {
                v30 := GenS1{F0: 255}
                v14 = 2
                if (476 >= 2) {
                    v31 := GenS1{F0: 107}
                    fmt.Println(v14, fi29, v14)
                    _ = v31
                } else {
                    v32 := GenS3{F0: (-(fi29)), F1: 255}
                    if true {
                        switch ("hello" + ("foo bar" + "\n\t")) {
                        case "a":
                            var v33 int16 = 798
                            _ = v33
                        case "foo bar":
                            var v34 bool = true
                            _ = v34
                        default:
                            var v35 uint32 = 411
                            _ = v35
                        }
                        _ = len(v2)
                        var v36 int64 = (v5 * (v5 | 1))
                        _ = v36
                    }
                    fmt.Println(fi29, fi27)
                    _ = v32
                }
                _ = fi29
                _ = v30
            }
            _ = fi27
            _ = v28
        }
        var v37 uint64 = (v14 ^ v14)
        _ = v37
    }
    v38 := GenS1{F0: 11}
    for fi39 := 0; fi39 < 5; fi39++ {
        var v40 uint16 = 32767
        v14 = (v14 >> ((42) & 63))
        _ = fi39
        _ = v40
    }
    v2[42] = 2147483648
    if (true || v1) {
        for fi41 := 0; fi41 < 18; fi41++ {
            if (2 < 255) {
                switch 32768 {
                case 9223372036854775807:
                    var v42 int16 = 0
                    _ = v42
                case 106:
                    v5 += (-(v5))
                default:
                    v43 := GenS3{F0: fi41, F1: 2}
                    _ = v43
                }
                if (!v1) {
                    if false {
                        for fi44 := 0; fi44 < 10; fi44++ {
                            v45 := GenS4{F0: 42, F1: -952, F2: v5}
                            var v46 int16 = -1
                            v47 := GenS1{F0: 22}
                            _ = fi44
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                        var v48 int16 = 607
                        v3 = append(v3, (fi41 ^ (fi41 | fi41)))
                        _ = v48
                    }
                    v49 := GenS3{F0: -173, F1: 2}
                    _ = v49
                } else {
                    v5 = v5
                    v5 = (-(v5))
                    var v50 int8 = (v4 | (v4 | v4))
                    _ = v50
                }
            }
            v14 -= (v14 - (v14 | 4294967295))
            _ = fi41
        }
    }
    var v51 float64 = (1.5 + 1e10)
    fmt.Println(v4)
    if (true && (("" <= "a") || false)) {
        switch ("foo bar" + "hello") {
        case "hello":
            switch -9223372036854775808 {
            case -785:
                v52 := GenS1{F0: 0}
                _ = v52
            case 127:
                var v53 int32 = 889
                _ = v53
            default:
                v54 := GenS1{F0: 1}
                _ = v54
            }
        case "\n\t":
            var v55 uint8 = 1
            _ = v55
        }
    } else {
        var v56 uint8 = 0
        for fi57 := 0; fi57 < 14; fi57++ {
            v56 -= (v56 + v56)
            var v58 uint64 = 18446744073709551615
            _ = fi57
            _ = v58
        }
        _ = v56
    }
    v5 = (v5 | v5)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v51
}
