package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 int64; F2 int8 }
type GenS4 struct { F0 uint64; F1 int32; F2 int8 }
type GenS5 struct { F0 uint32; F1 string; F2 uint }
type GenS6 struct { F0 int; F1 int }
type GenS7 struct { F0 uint16; F1 int16 }
func genFn0(p0 bool, p1 int64, p2 int) string {
    v4 := []int{p2, (p2 + 4294967295), (p2 - p2)}
    for fi5 := 0; fi5 < 2; fi5++ {
        p0 = p0
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 float32 = 1e10
            p0 = false
            if len(v4) > 0 {
                v4[0] = (fi6 - fi6)
            }
            _ = fi6
            _ = v7
        }
        var v8 uint = 108
        _ = fi5
        _ = v8
    }
    p0 = true
    var v9 uint8 = 2
    switch "\n\t" {
    case "a":
        for fi10 := 0; fi10 < 2; fi10++ {
            p0 = p0
            var v11 int8 = 42
            _ = fi10
            _ = v11
        }
    case "foo bar":
        p2 = p2
    default:
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 int32 = int32(p1)
            switch "" {
            case "hello":
                var v14 int32 = (v13 << ((255) & 63))
                _ = v14
            case "\n\t":
                switch (fi12 | (p2 + p2)) {
                case -9223372036854775808:
                    p1 += 1
                default:
                    v15 := GenS1{F0: (-(v13)), F1: false, F2: 547}
                    _ = v15
                }
            case "":
                if p0 {
                    switch ((("foo bar" + "\n\t") + "hello") + "hello") {
                    case "":
                        if len(v4) > 0 {
                            v4[0] = (-(p2))
                        }
                    case "\n\t":
                        v16 := GenS2{F0: 2147483648}
                        _ = v16
                    default:
                        var v17 float64 = 0.1
                        _ = v17
                    }
                    var v18 string = ""
                    _ = v18
                }
            default:
                for fi19 := 0; fi19 < 18; fi19++ {
                    v4 = append(v4, 32767)
                    var v20 uint = 253
                    _ = fi19
                    _ = v20
                }
            }
            _ = fi12
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v9
    return ("\n\t" + "hello")
}

func main() {
    var v1 int = -396
    var v2 int = v1
    v3 := []int{-2147483648}
    v1 = (-(v2))
    fmt.Println(v2)
    v21 := GenS1{F0: 1, F1: (!false), F2: 655}
    v1 -= v2
    fmt.Println(v2, v1)
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 bool = true
        v23 = false
        var v24 int8 = -1
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 int8 = 0
    v26 := GenS3{F0: 2, F1: -9223372036854775808, F2: 117}
    var v27 uint16 = 786
    switch (^(v1)) {
    case 128:
        switch (v1 - (v1 ^ v2)) {
        case 9223372036854775807:
            v28 := GenS4{F0: 2, F1: int32(v27), F2: v25}
            _ = v28
        }
    case 2147483647:
        v29 := GenS5{F0: 4294967295, F1: (("\n\t" + "foo bar") + "foo bar"), F2: 4294967295}
        _ = v29
    case 127:
        var v30 uint32 = 2147483647
        _ = v30
    }
    if false {
        var v31 uint32 = 928
        v32 := GenS6{F0: (v1 >> ((255) & 63)), F1: (v1 ^ (v2 ^ (v1 & v2)))}
        _ = v31
        _ = v32
    } else {
        var v33 uint64 = 147
        if true {
            if true {
                for fi34 := 0; fi34 < 20; fi34++ {
                    fmt.Println(v25, v27)
                    _ = fi34
                }
                for fi35 := 0; fi35 < 15; fi35++ {
                    for fi36 := 0; fi36 < 12; fi36++ {
                        if false {
                            var v37 float32 = 1.5
                            _ = v37
                        }
                        _ = fi36
                    }
                    fmt.Println(v27, v27, v25)
                    _ = fi35
                }
            } else {
                v27 = v27
                fmt.Println(v27)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
        fmt.Println(v33)
        _ = v33
    }
    _ = len(v3)
    if true {
        var v38 uint = 1
        v39 := GenS4{F0: 190, F1: 42, F2: v25}
        v40 := GenS4{F0: 1, F1: -870, F2: (v25 << ((v38) & 63))}
        _ = v38
        _ = v39
        _ = v40
    }
    v41 := GenS7{F0: 255, F1: -1}
    v42 := GenS1{F0: 1, F1: (-779 != 1), F2: 596}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v41
    _ = v42
}
