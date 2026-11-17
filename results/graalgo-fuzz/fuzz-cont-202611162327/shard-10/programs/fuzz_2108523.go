package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int; F1 uint64; F2 int64 }
func genFn0(p0 uint64, p1 bool) int64 {
    v7 := []int{128, 0, -937, 32768}
    for fi8 := 0; fi8 < 12; fi8++ {
        switch "" {
        case "a":
            if p1 {
                p0 += (^(p0))
            } else {
                fmt.Println(p0)
                switch fi8 {
                case 4294967295:
                    var v9 float32 = 1e10
                    _ = v9
                case -165:
                    p0 = (p0 - 651)
                }
                if len(v7) > 0 {
                    v7[0] = fi8
                }
            }
        case "":
            var v10 int16 = 392
            _ = v10
        case "hello":
            var v11 int64 = -524
            _ = v11
        default:
            for fi12 := 0; fi12 < 10; fi12++ {
                p1 = (-20 == -2147483648)
                if ((42 != 4294967295) && (p1 || p1)) {
                    switch "hello" {
                    case "hello":
                        p1 = p1
                    case "":
                        var v13 int16 = 32767
                        _ = v13
                    }
                    fmt.Println(p0)
                    _ = len(v7)
                }
                v14 := GenS1{F0: -103}
                var v15 uint32 = 2147483647
                _ = fi12
                _ = v14
                _ = v15
            }
        }
        v16 := GenS2{F0: 1}
        var v17 int = 9223372036854775807
        _ = fi8
        _ = v16
        _ = v17
    }
    fmt.Println(p1, p1)
    _ = len(v7)
    _ = p0
    _ = p1
    _ = v7
    return 307
}

func main() {
    var v1 string = "a"
    var v2 float64 = 1e10
    v3 := []uint32{65535, 1}
    var v4 int64 = 1
    var v5 uint = 18446744073709551615
    v6 := make(map[string]int)
    for fi18 := 0; fi18 < 8; fi18++ {
        if len(v3) > 0 {
            v3[0] = 1
        }
        var v19 uint64 = 18446744073709551615
        var v20 uint32 = 300
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 uint8 = 42
    v22 := GenS3{F0: (v21 * (v21 | v21))}
    v23 := v6[(v1 + v1)]
    for fi24 := 0; fi24 < 18; fi24++ {
        if (false && true) {
            switch fi24 {
            case 713:
                for fi25 := 0; fi25 < 9; fi25++ {
                    v26 := GenS1{F0: 42}
                    if false {
                        v6[v1] = fi25
                    }
                    _ = fi25
                    _ = v26
                }
            case -874:
                switch v1 {
                case "":
                    v27 := GenS3{F0: 2}
                    _ = v27
                default:
                    for fi28 := 0; fi28 < 4; fi28++ {
                        var v29 int64 = v4
                        _ = fi28
                        _ = v29
                    }
                }
            case 32767:
                var v30 float32 = (0.1 * 0.1)
                _ = v30
            default:
                for fi31 := 0; fi31 < 19; fi31++ {
                    v32 := GenS3{F0: uint8(v5)}
                    v33 := GenS4{F0: 0}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            }
            v3 = append(v3, 4294967295)
        } else {
            fmt.Println(fi24)
        }
        _ = fi24
    }
    v21 += 147
    var v34 uint16 = 42
    fmt.Println(v1, v4, v5)
    if false {
        for fi35 := 0; fi35 < 1; fi35++ {
            var v36 uint8 = v21
            v37 := GenS5{F0: 174}
            _ = fi35
            _ = v36
            _ = v37
        }
        v1 = ""
        var v38 string = "\n\t"
        _ = v38
    } else {
        var v39 bool = false
        v40 := GenS6{F0: (^(v23)), F1: 689, F2: 9223372036854775807}
        _ = len(v6)
        _ = v39
        _ = v40
    }
    var v41 uint64 = 1
    var v42 uint64 = uint64(v34)
    switch v1 {
    case "hello":
        v43 := GenS6{F0: (v23 << ((v5) & 63)), F1: v42, F2: -9223372036854775808}
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
