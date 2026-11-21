package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint16; F1 int32 }

func main() {
    var v1 int16 = 0
    v2 := []string{"a", ("a" + ("hello" + "\n\t"))}
    v3 := make(map[int]string)
    var v4 string = "a"
    var v5 uint32 = 162
    v6 := []uint32{(v5 / ((v5) | 1)), v5, v5, v5}
    v7 := GenS1{F0: (!false), F1: -260, F2: 1}
    if false {
        switch 2147483647 {
        case 4294967295:
            for fi8 := 0; fi8 < 14; fi8++ {
                var v9 float64 = 1e10
                if true {
                    if true {
                        var v10 float64 = ((v9 / v9) - ((0.0 * 0.1) / float64(1e10)))
                        var v11 float64 = (v10 / ((((v9 * v9) / (0.0 * v10)) * ((((v9 * v10) * (1.5 - (0.1 / (1.5 + (1.5 * v10))))) - (v10 + (v9 - v10))) * float64(3.14159))) + v10))
                        _ = v10
                        _ = v11
                    }
                    _ = len(v3)
                } else {
                    v5 = (v5 * (^(v5)))
                }
                _ = fi8
                _ = v9
            }
        case 0:
            var v12 int32 = 42
            _ = v12
        case 591:
            switch "a" {
            case "foo bar":
                v13 := GenS2{F0: -52}
                _ = v13
            default:
                v14 := GenS1{F0: true, F1: v1, F2: 563}
                _ = v14
            }
        }
    }
    if len(v2) > 0 {
        v2[0] = "a"
    }
    var v15 string = v4
    var v16 float64 = ((1e10 / 0.1) - 1e10)
    for fi17 := 0; fi17 < 13; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            _ = len(v3)
            fmt.Println(fi17, fi17, fi18)
            fmt.Println(fi17)
            _ = fi18
        }
        v6 = append(v6, 933)
        var v19 int = 128
        _ = fi17
        _ = v19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 float64 = 0.0
            fmt.Println(v1, v16)
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    var v23 string = "foo bar"
    var v24 uint = 42
    fmt.Println(v5)
    v16 = v16
    if (0 >= 178) {
        v25 := GenS1{F0: (127 >= 42), F1: v1, F2: 4294967295}
        var v26 int32 = 2147483647
        var v27 int8 = 49
        _ = v25
        _ = v26
        _ = v27
    } else {
        for fi28 := 0; fi28 < 7; fi28++ {
            delete(v3, 852)
            v29 := GenS3{F0: -1}
            _ = fi28
            _ = v29
        }
        if true {
            if (((!true) || true) || false) {
                v30 := GenS4{F0: 115, F1: 39, F2: 65535}
                for fi31 := 0; fi31 < 2; fi31++ {
                    for fi32 := 0; fi32 < 14; fi32++ {
                        _ = len(v3)
                        if true {
                            fmt.Println(fi32, v1)
                            v33 := GenS5{F0: 2, F1: 1}
                            v15 = "\n\t"
                            _ = v33
                        }
                        var v34 float64 = 0.0
                        _ = fi32
                        _ = v34
                    }
                    v35 := GenS4{F0: 1, F1: 100, F2: 32767}
                    _ = fi31
                    _ = v35
                }
                fmt.Println(v15, v16, v4)
                _ = v30
            }
            var v36 int = 42
            v4 = "foo bar"
            _ = v36
        }
    }
    var v37 uint8 = 1
    v23 = v23
    var v38 uint8 = (v37 << (((^(v24))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
