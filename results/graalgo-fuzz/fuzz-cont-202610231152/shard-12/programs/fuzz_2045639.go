package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int }
type GenS2 struct { F0 uint; F1 bool; F2 int32 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 int64; F2 int64 }
func genFn0(p0 int) int8 {
    v4 := GenS1{F0: 1, F1: 42, F2: p0}
    var v5 int = (p0 << ((734) & 63))
    _ = p0
    _ = v4
    _ = v5
    return 40
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{}
    var v3 int8 = 60
    for fi6 := 0; fi6 < 11; fi6++ {
        v3 = (^(v3))
        _ = fi6
    }
    var v7 uint16 = 2
    var v8 string = "foo bar"
    v3 += (v3 + -1)
    v3 = (v3 + (v3 >> ((18446744073709551615) & 63)))
    var v9 float32 = float32(-2.25)
    if false {
        var v10 string = "hello"
        _ = v10
    } else {
        switch ("\n\t" + v8) {
        case "foo bar":
            switch -1 {
            case 1:
                fmt.Println(v3, v3, v9)
            case 32767:
                var v11 int16 = 42
                _ = v11
            case 9223372036854775807:
                v9 -= v9
            }
        default:
            switch 42 {
            case 2147483647:
                v12 := GenS2{F0: 615, F1: false, F2: 2147483647}
                _ = v12
            case 42:
                for fi13 := 0; fi13 < 14; fi13++ {
                    fmt.Println(fi13)
                    _ = fi13
                }
            default:
                var v14 uint8 = uint8(v7)
                _ = v14
            }
        }
    }
    v15 := GenS3{F0: -321, F1: (-(v3))}
    v16 := GenS2{F0: 4294967295, F1: (42 == 205), F2: 2147483647}
    var v17 uint16 = 65535
    switch ((v8 + v8) + v8) {
    case "hello":
        if true {
            for fi18 := 0; fi18 < 9; fi18++ {
                if false {
                    v19 := GenS4{F0: 55}
                    _ = v19
                } else {
                    if true {
                        if len(v2) > 0 {
                            v2[0] = (fi18 / (((fi18 | -2147483648)) | 1))
                        }
                        v17 += v7
                        for fi20 := 0; fi20 < 9; fi20++ {
                            _ = len(v2)
                            v17 -= v17
                            _ = len(v2)
                            _ = fi20
                        }
                    }
                    for fi21 := 0; fi21 < 20; fi21++ {
                        if ((3.14159 + float64(0.1)) == (1e10 * -2.25)) {
                            var v22 uint32 = 871
                            var v23 int16 = -208
                            _ = v22
                            _ = v23
                        }
                        _ = fi21
                    }
                    v24 := v1[2147483647]
                    _ = v24
                }
                _ = fi18
            }
            v25 := GenS1{F0: 862, F1: 2, F2: 196}
            var v26 bool = false
            _ = v25
            _ = v26
        }
    case "a":
        v27 := GenS5{F0: 1, F1: -509, F2: 42}
        _ = v27
    }
    v28 := GenS3{F0: 0, F1: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
