package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint32; F1 int; F2 int16 }

func main() {
    var v1 string = (("foo bar" + "foo bar") + "hello")
    v2 := []uint32{}
    v3 := []uint32{330}
    var v4 int64 = 1
    v5 := make(map[int]uint32)
    var v6 uint8 = 33
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 float64 = (((-2.25 * 1.5) - (3.14159 + (1.5 + (3.14159 - float64((-2.25 / 1e10)))))) / (float64(1.5) * -2.25))
        _ = fi7
        _ = v8
    }
    if true {
        switch v1 {
        case "foo bar":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    } else {
        var v9 uint64 = 932
        v9 += v9
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        if false {
            v1 = (v1 + v1)
            switch v1 {
            case "\n\t":
                fmt.Println(v1, v1)
            case "hello":
                v4 = -9223372036854775808
            case "foo bar":
                for fi11 := 0; fi11 < 9; fi11++ {
                    var v12 string = v1
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(fi11)
                    case "a":
                        var v13 float32 = 0.1
                        _ = v13
                    default:
                        if (false && true) {
                            v14 := v5[-861]
                            _ = v14
                        } else {
                            v15 := GenS1{F0: 235, F1: 1}
                            var v16 float32 = (1.5 + 1e10)
                            v6 = 42
                            _ = v15
                            _ = v16
                        }
                        if false {
                            v1 = "foo bar"
                            v17 := GenS1{F0: 47, F1: 2147483647}
                            var v18 bool = (42 >= 4294967295)
                            v19 := GenS2{F0: (-32768 <= 0), F1: -434}
                            _ = len(v5)
                            var v20 bool = false
                            v21 := GenS2{F0: (true || true), F1: 973}
                            _ = v17
                            _ = v18
                            _ = v19
                            _ = v20
                            _ = v21
                        } else {
                            v22 := GenS1{F0: 301, F1: 54}
                            v1 = v1
                            fmt.Println(v4, v4, v12)
                            v6 = 42
                            _ = v22
                        }
                        v1 = ("hello" + "\n\t")
                    }
                    var v23 float32 = 3.14159
                    _ = fi11
                    _ = v12
                    _ = v23
                }
            default:
                for fi24 := 0; fi24 < 9; fi24++ {
                    v6 += (v6 << ((255) & 63))
                    _ = fi24
                }
            }
        }
        var v25 float64 = float64(1e10)
        var v26 bool = ((fi10 % ((fi10) | 1)) < (fi10 % (((fi10 & fi10)) | 1)))
        _ = fi10
        _ = v25
        _ = v26
    }
    var v27 uint32 = 2147483648
    var v28 float64 = 1e10
    if true {
        for fi29 := 0; fi29 < 18; fi29++ {
            fmt.Println(v27, v28, v27)
            var v30 int = 9223372036854775807
            v31 := GenS1{F0: 108, F1: 42}
            _ = fi29
            _ = v30
            _ = v31
        }
        var v32 int64 = (v4 * 1)
        _ = v32
    } else {
        var v33 uint64 = 18446744073709551615
        v6 -= v6
        _ = v33
    }
    fmt.Println(v1)
    for fi34 := 0; fi34 < 18; fi34++ {
        if len(v3) > 0 {
            v3[0] = v27
        }
        v28 += v28
        var v35 int = fi34
        _ = fi34
        _ = v35
    }
    if (!true) {
        switch -824 {
        case 277:
            if false {
                v36 := GenS2{F0: true, F1: -1}
                _ = v36
            }
        case 9223372036854775807:
            if (true || true) {
                if false {
                    v27 -= (v27 + 0)
                }
                if false {
                    if len(v3) > 0 {
                        v3[0] = (v27 & (v27 | (v27 - v27)))
                    }
                }
            } else {
                if false {
                    var v37 int64 = v4
                    _ = v37
                }
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v38 int16 = 0
    var v39 int16 = v38
    switch v1 {
    case "":
        v39 += -458
    case "foo bar":
        var v40 uint = 428
        _ = v40
    default:
        switch v1 {
        case "hello":
            if (((((((0.0 + 0.1) / (3.14159 - 0.0)) - 1.5) * 0.0) * (float32(v28) - -2.25)) == ((-2.25 * 3.14159) - (1e10 * (1.5 - 0.0)))) && true) {
                for fi41 := 0; fi41 < 4; fi41++ {
                    var v42 int16 = v38
                    var v43 int = 32768
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                fmt.Println(v28)
            } else {
                var v44 int16 = (-(v38))
                _ = v44
            }
        case "\n\t":
            for fi45 := 0; fi45 < 15; fi45++ {
                v46 := GenS1{F0: 1, F1: 1}
                v39 = v39
                if (false && false) {
                    v47 := GenS2{F0: false, F1: (v39 / (((-(v38))) | 1))}
                    v48 := GenS3{F0: v27, F1: int(v39), F2: (^(v38))}
                    var v49 uint32 = uint32(v6)
                    _ = v47
                    _ = v48
                    _ = v49
                } else {
                    v50 := v5[(fi45 >> ((18446744073709551615) & 63))]
                    _ = v50
                }
                _ = fi45
                _ = v46
            }
        }
    }
    if len(v3) > 0 {
        v3[0] = (v27 << ((42) & 63))
    }
    switch 127 {
    case 42:
        v28 += v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
