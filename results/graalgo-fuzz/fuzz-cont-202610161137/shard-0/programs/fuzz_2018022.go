package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 string; F1 uint8; F2 int16 }
func genFn0(p0 uint32, p1 int, p2 uint16) uint8 {
    if true {
        v7 := GenS1{F0: "foo bar", F1: 96, F2: "a"}
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            fmt.Println(fi9)
            fmt.Println(p2, p1)
            p1 = (fi9 ^ (^(p1)))
            _ = fi9
        }
        fmt.Println(p1, p0, p2)
        var v10 int64 = 1
        _ = fi8
        _ = v10
    }
    p0 -= 42
    p1 -= p1
    v11 := GenS1{F0: "a", F1: 152, F2: ("\n\t" + "hello")}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 128
}

func main() {
    v1 := []int{746, -176, -421}
    var v2 uint64 = 4294967295
    v3 := make(map[int]uint32)
    v4 := make(map[int]int)
    var v5 uint32 = 65535
    var v6 uint32 = v5
    v1 = append(v1, -474)
    for fi12 := 0; fi12 < 11; fi12++ {
        if (255 > 32767) {
            v6 = (v5 % ((v6) | 1))
            v2 = v2
            switch "hello" {
            case "hello":
                if (76 < -1) {
                    switch "" {
                    case "a":
                        if false {
                            var v13 int32 = 42
                            _ = v13
                        }
                    case "":
                        fmt.Println(v2)
                    case "hello":
                        var v14 int32 = 0
                        _ = v14
                    }
                }
            case "foo bar":
                for fi15 := 0; fi15 < 11; fi15++ {
                    fmt.Println(fi15, fi15, v6)
                    _ = fi15
                }
            case "":
                switch "" {
                case "a":
                    v16 := GenS2{F0: -2147483648}
                    _ = v16
                }
            }
        } else {
            if (!((false || (true || false)) && true)) {
                v5 += v6
                var v17 uint8 = 1
                for fi18 := 0; fi18 < 9; fi18++ {
                    fmt.Println(v17, v2, v6)
                    var v19 string = "hello"
                    _ = fi18
                    _ = v19
                }
                _ = v17
            }
            switch ("a" + "\n\t") {
            case "foo bar":
                v20 := GenS3{F0: true, F1: (127 > 42)}
                _ = v20
            case "":
                switch 0 {
                case 9223372036854775807:
                    for fi21 := 0; fi21 < 18; fi21++ {
                        v2 += (^(v2))
                        var v22 uint64 = 2
                        _ = fi21
                        _ = v22
                    }
                }
            }
        }
        if false {
            var v23 uint64 = (v2 % (((v2 / ((v2) | 1))) | 1))
            v24 := GenS1{F0: ("hello" + ("foo bar" + "\n\t")), F1: 786, F2: ""}
            _ = v23
            _ = v24
        }
        v25 := GenS4{F0: "foo bar", F1: uint8(v5), F2: 42}
        _ = fi12
        _ = v25
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        var v27 int32 = 0
        v27 = 1
        if ((^(fi26)) != fi26) {
            v28 := v3[(fi26 >> ((18446744073709551615) & 63))]
            _ = v28
        }
        _ = fi26
        _ = v27
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        fmt.Println(v2, v6)
        for fi30 := 0; fi30 < 15; fi30++ {
            var v31 uint16 = uint16(v5)
            if false {
                _ = len(v1)
                var v32 int = (fi30 / ((fi29) | 1))
                _ = v32
            } else {
                v33 := GenS2{F0: -2147483648}
                v2 = 10
                _ = v33
            }
            v3[-2147483648] = v6
            v34 := GenS4{F0: "", F1: 0, F2: 678}
            _ = fi30
            _ = v31
            _ = v34
        }
        var v35 uint16 = 65535
        _ = fi29
        _ = v35
    }
    switch "\n\t" {
    case "a":
        v6 += (v5 ^ 68)
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        if len(v1) > 0 {
            v1[0] = fi36
        }
        var v37 int64 = 0
        fmt.Println(v37)
        _ = fi36
        _ = v37
    }
    var v38 bool = true
    v2 -= (^(v2))
    for fi39 := 0; fi39 < 18; fi39++ {
        var v40 int8 = -1
        fmt.Println(v40, v2, v6)
        _ = fi39
        _ = v40
    }
    var v41 float64 = -2.25
    var v42 uint32 = (v5 % ((686) | 1))
    for fi43 := 0; fi43 < 7; fi43++ {
        switch "" {
        case "foo bar":
            switch "hello" {
            case "":
                for fi44 := 0; fi44 < 16; fi44++ {
                    var v45 string = "foo bar"
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi44
                    _ = v45
                }
            }
        case "hello":
            delete(v3, (fi43 << ((uint(v5)) & 63)))
        case "a":
            if (1 >= 887) {
                v2 -= v2
            } else {
                v2 -= (v2 ^ 42)
                switch fi43 {
                case 1:
                    for fi46 := 0; fi46 < 13; fi46++ {
                        var v47 uint32 = 4294967295
                        _ = fi46
                        _ = v47
                    }
                default:
                    var v48 int = 2147483647
                    _ = v48
                }
                var v49 int64 = -1
                _ = v49
            }
        }
        _ = fi43
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
