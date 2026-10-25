package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint32 }
type GenS2 struct { F0 int16; F1 int8; F2 uint64 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 int32 = 42
    var v2 string = "\n\t"
    var v3 uint = 65535
    if (!(true || (1 == 24))) {
        for fi4 := 0; fi4 < 10; fi4++ {
            switch "foo bar" {
            case "hello":
                switch 32767 {
                case 9223372036854775807:
                    switch (^(fi4)) {
                    case 32768:
                        for fi5 := 0; fi5 < 8; fi5++ {
                            v1 = (v1 >> ((v3) & 63))
                            var v6 float32 = (0.1 - float32(0.1))
                            v2 = v2
                            _ = fi5
                            _ = v6
                        }
                    default:
                        var v7 int16 = 32767
                        _ = v7
                    }
                case -731:
                    var v8 int16 = -1
                    _ = v8
                case -797:
                    switch 127 {
                    case 719:
                        var v9 int = (^(fi4))
                        _ = v9
                    case -787:
                        for fi10 := 0; fi10 < 11; fi10++ {
                            v11 := GenS1{F0: v2, F1: v1, F2: 450}
                            _ = fi10
                            _ = v11
                        }
                    case 0:
                        switch v2 {
                        case "hello":
                            var v12 int8 = 0
                            _ = v12
                        case "foo bar":
                            v13 := GenS1{F0: v2, F1: (v1 >> ((311) & 63)), F2: 326}
                            _ = v13
                        default:
                            for fi14 := 0; fi14 < 15; fi14++ {
                                fmt.Println(fi14, v2)
                                v15 := []int{-1, fi14, fi4}
                                var v16 bool = false
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                        }
                    }
                default:
                    if true {
                        v17 := []int{(fi4 / ((32768) | 1))}
                        _ = v17
                    }
                }
            default:
                if false {
                    var v18 uint8 = 21
                    var v19 int64 = 1
                    v20 := GenS2{F0: 738, F1: 57, F2: 1}
                    _ = v18
                    _ = v19
                    _ = v20
                }
            }
            _ = fi4
        }
        fmt.Println(v1, v1, v3)
    } else {
        v21 := []int{-1, 42}
        _ = v21
    }
    v3 = (v3 % (((^(v3))) | 1))
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 int = (fi22 * (fi22 << ((65535) & 63)))
        for fi24 := 0; fi24 < 16; fi24++ {
            var v25 string = "\n\t"
            _ = fi24
            _ = v25
        }
        var v26 int64 = -9223372036854775808
        _ = fi22
        _ = v23
        _ = v26
    }
    v27 := []int{-9223372036854775808}
    _ = len(v27)
    var v28 bool = (false && ((true && (0 < -128)) || true))
    if v28 {
        var v29 int8 = -120
        v2 = "\n\t"
        v1 += 42
        _ = v29
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        switch -868 {
        case 4294967295:
            if len(v27) >= 2 {
                v27 = v27[1:len(v27)]
            }
        case -77:
            switch v2 {
            case "":
                var v31 float64 = 1.5
                _ = v31
            case "foo bar":
                for fi32 := 0; fi32 < 17; fi32++ {
                    fmt.Println(fi32, v3)
                    v33 := GenS1{F0: (("foo bar" + (v2 + (v2 + v2))) + v2), F1: -854, F2: 1}
                    _ = fi32
                    _ = v33
                }
            case "\n\t":
                _ = len(v27)
            default:
                v28 = v28
            }
        case -1:
            for fi34 := 0; fi34 < 2; fi34++ {
                v3 += v3
                var v35 float32 = (-2.25 / 0.1)
                v36 := GenS3{F0: 2147483647, F1: 42}
                var v37 uint16 = 42
                var v38 int = (^(fi34))
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
                _ = v38
            }
        }
        fmt.Println(v3, v3)
        if len(v27) > 0 {
            v27[0] = fi30
        }
        _ = fi30
    }
    switch 42 {
    case 1:
        if len(v27) > 0 {
            v27[0] = -9223372036854775808
        }
    case 128:
        fmt.Println(v3)
    case 4294967295:
        v1 -= 0
    }
    for fi39 := 0; fi39 < 6; fi39++ {
        var v40 int8 = -1
        _ = fi39
        _ = v40
    }
    var v41 uint16 = 255
    v3 += (^(v3))
    var v42 int8 = -47
    v43 := GenS4{F0: (v42 % ((v42) | 1))}
    switch 127 {
    case 2147483647:
        var v44 uint8 = 42
        _ = v44
    case 999:
        if ((v42 * v42) > v42) {
            fmt.Println(v28, v42)
        }
    }
    v28 = (v28 && (v1 > (v1 - (v1 ^ 42))))
    for fi45 := 0; fi45 < 12; fi45++ {
        v27 = append(v27, fi45)
        v27 = append(v27, 2147483647)
        _ = fi45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v41
    _ = v42
    _ = v43
}
