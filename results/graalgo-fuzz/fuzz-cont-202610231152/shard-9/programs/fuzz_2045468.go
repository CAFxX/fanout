package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint32; F2 bool }

func main() {
    var v1 int = -63
    v2 := make(map[string]uint32)
    var v3 uint = 42
    v4 := make(map[int]bool)
    v5 := []string{"\n\t"}
    var v6 uint16 = 246
    switch v1 {
    case 9223372036854775807:
        v1 += (v1 | (v1 + v1))
    case -305:
        for fi7 := 0; fi7 < 20; fi7++ {
            v5 = append(v5, ("a" + "a"))
            _ = fi7
        }
    }
    if (!true) {
        for fi8 := 0; fi8 < 14; fi8++ {
            v3 -= (v3 | v3)
            if false {
                fmt.Println(v3, v6)
            }
            _ = fi8
        }
    } else {
        if true {
            v9 := GenS1{F0: (v6 << ((v3) & 63)), F1: 968}
            _ = v9
        } else {
            v10 := GenS2{F0: -1}
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := GenS2{F0: 1}
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v13 bool = (42 < 18446744073709551615)
                _ = fi11
                _ = v12
                _ = v13
            }
            v6 -= (v6 << ((4294967295) & 63))
            _ = v10
        }
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        v6 += 2
        _ = fi14
    }
    v1 -= (v1 & (-(v1)))
    var v15 int16 = int16(v3)
    fmt.Println(v6, v6)
    v15 += v15
    v15 -= (v15 / ((v15) | 1))
    if (uint64(v3) >= 4294967295) {
        v16 := GenS1{F0: v6, F1: -464}
        for fi17 := 0; fi17 < 13; fi17++ {
            v18 := GenS1{F0: v6, F1: 42}
            switch fi17 {
            case 840:
                v5 = append(v5, "")
            case 658:
                v6 = (v6 + v6)
            }
            for fi19 := 0; fi19 < 15; fi19++ {
                v20 := GenS1{F0: v6, F1: 2147483647}
                v21 := GenS1{F0: v6, F1: 1}
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi17
            _ = v18
        }
        _ = v16
    } else {
        _ = len(v5)
    }
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    if false {
        var v22 int32 = -268
        v23 := GenS2{F0: 0}
        _ = v22
        _ = v23
    } else {
        var v24 int = (v1 | (v1 & v1))
        var v25 int16 = v15
        v15 += v15
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        v27 := GenS2{F0: 9223372036854775807}
        _ = fi26
        _ = v27
    }
    v6 = 2
    var v28 int16 = v15
    for fi29 := 0; fi29 < 9; fi29++ {
        if (32768 > (^(fi29))) {
            v6 -= 0
            v15 += 42
        } else {
            v30 := GenS3{F0: "hello"}
            var v31 int64 = 42
            if false {
                switch fi29 {
                case 0:
                    fmt.Println(v31, fi29, v3)
                case -2147483648:
                    v1 += (-(v1))
                case 32767:
                    var v32 float32 = 1.5
                    _ = v32
                }
                _ = len(v5)
                var v33 float32 = 0.0
                _ = v33
            }
            _ = v30
            _ = v31
        }
        v6 += v6
        if true {
            _ = len(v2)
        }
        _ = fi29
    }
    for fi34 := 0; fi34 < 17; fi34++ {
        var v35 float32 = (3.14159 + 1.5)
        _ = fi34
        _ = v35
    }
    v36 := GenS4{F0: -728, F1: 2147483647, F2: (false || (false || (2 > 2)))}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
