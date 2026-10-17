package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int16, p1 int8, p2 string) uint64 {
    fmt.Println(p2, p1)
    var v6 int = 809
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"", ("a" + "a"), "", "foo bar"}
    v3 := make(map[int]string)
    v4 := make(map[string]bool)
    v5 := make(map[string]string)
    if (!(-32768 == 32767)) {
        switch 32768 {
        case 182:
            var v7 int8 = 99
            _ = v7
        case -812:
            for fi8 := 0; fi8 < 11; fi8++ {
                v9 := GenS1{F0: 64}
                _ = fi8
                _ = v9
            }
        default:
            var v10 int8 = -56
            _ = v10
        }
        var v11 uint = 792
        switch "a" {
        case "\n\t":
            switch "\n\t" {
            case "hello":
                v12 := GenS2{F0: -9223372036854775808}
                _ = v12
            case "":
                var v13 int = 0
                _ = v13
            }
        default:
            var v14 int64 = 9223372036854775807
            _ = v14
        }
        _ = v11
    }
    v15 := GenS1{F0: 0}
    var v16 uint32 = 1
    var v17 uint = 4294967295
    if (!false) {
        v18 := GenS2{F0: 9223372036854775807}
        _ = v18
    }
    if (2147483647 < 415) {
        for fi19 := 0; fi19 < 13; fi19++ {
            var v20 float64 = float64(-2.25)
            var v21 bool = true
            _ = fi19
            _ = v20
            _ = v21
        }
    } else {
        fmt.Println(v16, v17)
    }
    if true {
        var v22 float64 = 0.1
        if (v17 >= v17) {
            v23 := GenS3{F0: 4294967295}
            for fi24 := 0; fi24 < 17; fi24++ {
                v16 -= (v16 & (v16 % ((2147483647) | 1)))
                var v25 int16 = -32768
                v16 = v16
                _ = fi24
                _ = v25
            }
            _ = v23
        }
        switch 262 {
        case 128:
            v17 = (v17 | (v17 + 1))
        }
        _ = v22
    } else {
        for fi26 := 0; fi26 < 10; fi26++ {
            if ((65535 < 2) || false) {
                if (100 != 0) {
                    var v27 int8 = -1
                    if false {
                        var v28 int32 = 986
                        switch "" {
                        case "":
                            if len(v2) > 0 {
                                v2[0] = "a"
                            }
                        case "\n\t":
                            var v29 float32 = -2.25
                            _ = v29
                        case "a":
                            v28 = (-(v28))
                        }
                        _ = v28
                    } else {
                        var v30 string = "hello"
                        var v31 uint16 = 2
                        _ = v30
                        _ = v31
                    }
                    _ = v27
                } else {
                    for fi32 := 0; fi32 < 9; fi32++ {
                        if (true || true) {
                            fmt.Println(v16)
                            v17 = (v17 << (((v17 + v17)) & 63))
                            v16 = v16
                        }
                        var v33 uint8 = 0
                        _ = fi32
                        _ = v33
                    }
                    var v34 int64 = -1
                    _ = v34
                }
            } else {
                var v35 int16 = 1000
                v16 = v16
                fmt.Println(v16, v35)
                _ = v35
            }
            for fi36 := 0; fi36 < 5; fi36++ {
                fmt.Println(fi26, v16, v17)
                _ = fi36
            }
            var v37 uint16 = 66
            _ = fi26
            _ = v37
        }
        var v38 int32 = -2147483648
        v2 = append(v2, "foo bar")
        _ = v38
    }
    var v39 float32 = 1.5
    v4["a"] = (59 != 939)
    var v40 uint8 = 134
    switch "\n\t" {
    case "":
        fmt.Println(v40, v16, v39)
    case "a":
        var v41 int32 = -2147483648
        _ = v41
    }
    for fi42 := 0; fi42 < 20; fi42++ {
        v40 = 115
        _ = fi42
    }
    switch 302 {
    case 9223372036854775807:
        delete(v1, (v16 / (((^(v16))) | 1)))
    }
    var v43 int8 = 100
    switch -547 {
    case 128:
        switch "" {
        case "":
            for fi44 := 0; fi44 < 18; fi44++ {
                var v45 uint8 = (v40 & uint8(v16))
                if true {
                    v46 := GenS1{F0: 899}
                    _ = v46
                }
                if len(v2) > 0 {
                    v2[0] = ("hello" + "foo bar")
                }
                _ = fi44
                _ = v45
            }
        case "a":
            var v47 uint8 = v40
            _ = v47
        case "hello":
            v43 += v43
        default:
            var v48 uint64 = 1
            _ = v48
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v5
}
