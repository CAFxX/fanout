package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint16; F1 int32 }

func main() {
    var v1 float32 = (3.14159 - float32(1e10))
    var v2 int16 = -555
    v3 := []uint32{0, 4294967295, 1, 191}
    fmt.Println(v1, v1)
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 int8 = 0
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
        _ = fi6
    }
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        for fi7 := 0; fi7 < 9; fi7++ {
            switch "" {
            case "a":
                v3 = append(v3, 2147483648)
            case "foo bar":
                switch "hello" {
                case "a":
                    v1 = v1
                case "\n\t":
                    switch (fi7 | fi7) {
                    case 4294967295:
                        for fi8 := 0; fi8 < 2; fi8++ {
                            fmt.Println(fi7, v2)
                            _ = fi8
                        }
                    default:
                        var v9 float32 = (v1 - 1e10)
                        _ = v9
                    }
                }
            case "hello":
                v1 += v1
            default:
                v1 = (v1 - (v1 * 0.1))
            }
            fmt.Println(v1)
            _ = len(v3)
            _ = fi7
        }
        v10 := GenS1{F0: ("\n\t" == "foo bar"), F1: 42}
        _ = v10
    }
    fmt.Println(v2)
    if true {
        for fi11 := 0; fi11 < 14; fi11++ {
            if len(v3) > 0 {
                v3[0] = 2
            }
            v2 += v2
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi11
        }
    }
    if false {
        v12 := GenS2{F0: true, F1: 42}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 12; fi13++ {
            v2 = -1
            _ = fi13
        }
        var v14 uint16 = 65535
        v14 += (v14 + 65535)
        _ = v14
    }
    var v15 float32 = v1
    switch 2147483647 {
    case -745:
        v16 := GenS3{F0: 2, F1: -423}
        _ = v16
    case 108:
        var v17 uint64 = 4294967295
        _ = v17
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (!(false && (false && false))) {
        for fi18 := 0; fi18 < 18; fi18++ {
            for fi19 := 0; fi19 < 7; fi19++ {
                var v20 bool = true
                v3 = append(v3, 4294967295)
                _ = fi19
                _ = v20
            }
            var v21 uint32 = 0
            _ = fi18
            _ = v21
        }
        var v22 uint32 = 65535
        v22 = v22
        _ = v22
    }
    switch -777 {
    case 32767:
        v1 -= 1.5
    case -1:
        if len(v3) > 0 {
            v3[0] = 467
        }
    }
    v15 += v1
    for fi23 := 0; fi23 < 16; fi23++ {
        v24 := GenS1{F0: (2 != 251), F1: 65535}
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        if len(v3) > 0 {
            v3[0] = 2
        }
        _ = fi25
    }
    v2 += (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
}
