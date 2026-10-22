package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 int32; F1 uint32 }
func genFn0() string {
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint16 = 154
        v9 := []int{128, fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{128, -869}
    v11 := GenS1{F0: 146}
    _ = v10
    _ = v11
    return "foo bar"
}
func genFn1(p0 int8) int8 {
    switch "" {
    case "a":
        v12 := []int{}
        _ = v12
    case "":
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := GenS2{F0: 676, F1: 0}
            _ = fi13
            _ = v14
        }
    default:
        if (!true) {
            v15 := GenS1{F0: 65535}
            _ = v15
        }
    }
    v16 := GenS2{F0: 2147483648, F1: 4294967295}
    var v17 float32 = float32(((1.5 + -2.25) / -2.25))
    v17 -= ((1.5 + ((v17 + 0.1) + 3.14159)) / v17)
    _ = p0
    _ = v16
    _ = v17
    return (p0 >> ((65535) & 63))
}

func main() {
    var v1 float32 = 1e10
    v2 := []string{"\n\t"}
    v3 := make(map[int]string)
    v4 := make(map[string]bool)
    v5 := make(map[string]int)
    var v6 uint16 = 403
    v6 += v6
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        v18 := GenS2{F0: 42, F1: 817}
        if ((100 >= 127) && ((v6 > (^(v6))) && (true || true))) {
            var v19 int32 = 1
            var v20 int32 = (v19 ^ -2147483648)
            switch "foo bar" {
            case "hello":
                var v21 uint64 = 18446744073709551615
                _ = v21
            case "":
                var v22 int64 = 9223372036854775807
                _ = v22
            case "\n\t":
                v6 = 42
            default:
                switch ("" + "\n\t") {
                case "":
                    var v23 uint32 = 1
                    _ = v23
                default:
                    for fi24 := 0; fi24 < 8; fi24++ {
                        fmt.Println(fi24)
                        var v25 bool = (false || true)
                        _ = fi24
                        _ = v25
                    }
                }
            }
            _ = v19
            _ = v20
        }
        var v26 uint16 = v6
        _ = v18
        _ = v26
    }
    v6 = 328
    if false {
        var v27 int32 = 42
        _ = v27
    }
    _ = len(v2)
    for fi28 := 0; fi28 < 19; fi28++ {
        switch "\n\t" {
        case "":
            v29 := GenS2{F0: 4294967295, F1: 42}
            _ = v29
        }
        var v30 uint32 = 65535
        _ = fi28
        _ = v30
    }
    v31 := GenS2{F0: 961, F1: 626}
    var v32 int32 = -317
    if (42 < 298) {
        for fi33 := 0; fi33 < 18; fi33++ {
            var v34 uint32 = 2
            v34 += v34
            _ = fi33
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        if true {
            v36 := GenS1{F0: v6}
            v32 = 2147483647
            _ = v36
        } else {
            for fi37 := 0; fi37 < 20; fi37++ {
                v38 := GenS2{F0: 1, F1: 0}
                switch (fi35 % ((fi35) | 1)) {
                case 467:
                    v39 := GenS1{F0: v6}
                    v32 -= -2147483648
                    if (v6 <= v6) {
                        _ = len(v3)
                    }
                    _ = v39
                case 9223372036854775807:
                    v40 := GenS3{F0: (fi35 << ((4294967295) & 63)), F1: 18446744073709551615}
                    _ = v40
                case -834:
                    v6 = 65535
                default:
                    v2 = append(v2, "hello")
                }
                _ = fi37
                _ = v38
            }
            if true {
                var v41 uint32 = 65535
                for fi42 := 0; fi42 < 19; fi42++ {
                    v43 := GenS4{F0: 962, F1: v41}
                    var v44 int = (fi42 + fi42)
                    v2 = append(v2, "a")
                    _ = len(v4)
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                _ = v41
            }
            switch "\n\t" {
            case "a":
                var v45 int32 = (v32 / (((v32 & -104)) | 1))
                _ = v45
            case "hello":
                var v46 string = "a"
                _ = v46
            default:
                v1 += -2.25
            }
        }
        fmt.Println(v6, v1)
        _ = fi35
    }
    v32 = v32
    var v47 int8 = 110
    v6 += v6
    var v48 uint64 = 4294967295
    fmt.Println(v1, v6, v47)
    for fi49 := 0; fi49 < 18; fi49++ {
        var v50 uint = 18446744073709551615
        _ = fi49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
