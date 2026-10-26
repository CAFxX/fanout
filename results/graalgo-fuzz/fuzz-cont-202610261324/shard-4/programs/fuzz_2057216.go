package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS5 struct { F0 uint32; F1 int32 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    v3 := []string{"", "\n\t", "\n\t"}
    var v4 int64 = 42
    var v5 uint8 = 168
    v6 := GenS1{F0: 100}
    for fi7 := 0; fi7 < 2; fi7++ {
        v4 = (-(v4))
        v4 -= v4
        _ = fi7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int32 = -2147483648
        if false {
            if (779 <= 65535) {
                var v10 int16 = -454
                switch "" {
                case "hello":
                    v11 := GenS2{F0: v10, F1: 42}
                    _ = v11
                default:
                    switch (-(fi8)) {
                    case 32767:
                        v12 := GenS2{F0: -97, F1: 780}
                        _ = v12
                    }
                }
                _ = v10
            }
        } else {
            for fi13 := 0; fi13 < 14; fi13++ {
                v5 = 215
                switch "foo bar" {
                case "hello":
                    v14 := GenS1{F0: -1}
                    _ = v14
                case "foo bar":
                    if false {
                        _ = len(v1)
                        var v15 int16 = -1
                        _ = v15
                    } else {
                        if true {
                            v9 -= 918
                            v16 := GenS2{F0: 1, F1: 18446744073709551615}
                            _ = v16
                        }
                        switch (-(fi8)) {
                        case 1:
                            var v17 int = (fi13 * fi8)
                            _ = v17
                        case 9223372036854775807:
                            var v18 uint = 152
                            var v19 uint8 = (v5 ^ v5)
                            var v20 string = "foo bar"
                            _ = v18
                            _ = v19
                            _ = v20
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                }
                if (255 == v5) {
                    v9 = (v9 / (((v9 ^ (v9 % (((-(v9))) | 1)))) | 1))
                    fmt.Println(fi13, v5)
                    var v21 bool = false
                    _ = v21
                }
                _ = fi13
            }
            v9 += v9
            v22 := v1[fi8]
            _ = v22
        }
        _ = fi8
        _ = v9
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        for fi24 := 0; fi24 < 10; fi24++ {
            var v25 int8 = 1
            var v26 uint64 = 2
            var v27 uint8 = v5
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        switch "hello" {
        case "\n\t":
            v28 := GenS2{F0: 32767, F1: 0}
            _ = v28
        default:
            if len(v3) > 0 {
                v3[0] = ""
            }
        }
        v5 -= 2
        _ = fi23
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        fmt.Println(v4, fi29, fi29)
        if (false && (0 == 4294967295)) {
            v4 += -9223372036854775808
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v5 += v5
        } else {
            switch "\n\t" {
            case "a":
                v3 = append(v3, "foo bar")
            case "hello":
                switch (fi29 / ((128) | 1)) {
                case 2147483647:
                    v30 := GenS3{F0: 1000}
                    _ = v30
                case 1:
                    _ = len(v3)
                case 0:
                    v5 = 18
                default:
                    var v31 uint8 = (v5 + (v5 >> ((42) & 63)))
                    _ = v31
                }
            default:
                v32 := GenS1{F0: 1}
                _ = v32
            }
            v33 := GenS4{F0: 2, F1: (v5 >> ((1) & 63)), F2: (v5 * (^(v5)))}
            var v34 int16 = -32768
            _ = v33
            _ = v34
        }
        v2["a"] = (fi29 / ((fi29) | 1))
        _ = fi29
    }
    if false {
        v35 := v2["foo bar"]
        v35 += (v35 + v35)
        for fi36 := 0; fi36 < 2; fi36++ {
            _ = len(v2)
            for fi37 := 0; fi37 < 15; fi37++ {
                var v38 uint64 = 42
                if ((-2.25 * 0.1) == 0.1) {
                    var v39 int8 = 27
                    v38 += (^(v38))
                    _ = v39
                } else {
                    v40 := GenS5{F0: 2, F1: -2147483648}
                    if (!true) {
                        fmt.Println(v5, v35, fi37)
                    }
                    v38 += v38
                    _ = v40
                }
                _ = fi37
                _ = v38
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi36
        }
        _ = v35
    } else {
        var v41 int = -534
        v42 := GenS2{F0: 1, F1: 4294967295}
        _ = v41
        _ = v42
    }
    for fi43 := 0; fi43 < 2; fi43++ {
        _ = len(v1)
        switch "a" {
        case "foo bar":
            for fi44 := 0; fi44 < 16; fi44++ {
                switch (-(fi43)) {
                case -9223372036854775808:
                    fmt.Println(v5, v5)
                case -2147483648:
                    var v45 uint8 = v5
                    _ = v45
                default:
                    v3 = append(v3, ("foo bar" + ("hello" + "a")))
                }
                var v46 uint32 = 1
                _ = fi44
                _ = v46
            }
        case "":
            fmt.Println(v5, v4, v4)
        default:
            v5 = (^(v5))
        }
        delete(v1, 127)
        _ = fi43
    }
    for fi47 := 0; fi47 < 10; fi47++ {
        v5 -= v5
        _ = fi47
    }
    fmt.Println(v4)
    var v48 int16 = 42
    var v49 int64 = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
