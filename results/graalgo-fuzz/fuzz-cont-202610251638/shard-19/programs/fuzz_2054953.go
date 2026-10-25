package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32 }
func genFn0() uint16 {
    var v5 int8 = 1
    var v6 uint32 = 2147483648
    fmt.Println(v6, v6)
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint32) int8 {
    v7 := GenS1{F0: 9223372036854775807, F1: "foo bar"}
    var v8 uint64 = 1
    switch 42 {
    case 1:
        if false {
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := []int{fi9}
                var v11 bool = true
                _ = fi9
                _ = v10
                _ = v11
            }
        } else {
            if false {
                if (0.1 <= float64(-2.25)) {
                    var v12 int = -1
                    _ = v12
                } else {
                    var v13 float64 = ((3.14159 + ((((0.1 / 1.5) - 0.1) * float64(1e10)) / (((-2.25 + ((3.14159 + 0.0) * 0.0)) / 0.1) * 1e10))) * 3.14159)
                    _ = v13
                }
                var v14 int64 = 0
                _ = v14
            }
        }
    }
    switch "\n\t" {
    case "hello":
        p0 = uint32(v8)
    default:
        v15 := GenS2{F0: 255, F1: genFn0()}
        _ = v15
    }
    if false {
        v16 := GenS3{F0: 0}
        v8 += 269
        if false {
            fmt.Println(p0, v8)
        }
        _ = v16
    }
    _ = p0
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int8 = -65
    var v2 int16 = -32768
    var v3 string = ("a" + "foo bar")
    var v4 uint = 1
    v17 := []int{37, -1, 2147483647, 9223372036854775807}
    for fi18 := 0; fi18 < 5; fi18++ {
        v19 := GenS4{F0: 817}
        _ = fi18
        _ = v19
    }
    if false {
        v4 = v4
        for fi20 := 0; fi20 < 5; fi20++ {
            var v21 int8 = 100
            if len(v17) > 0 {
                v17[0] = (fi20 >> ((391) & 63))
            }
            _ = fi20
            _ = v21
        }
        fmt.Println(v2)
    }
    v4 = v4
    var v22 int = -745
    v23 := GenS1{F0: -9223372036854775808, F1: v3}
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        if true {
            var v25 int32 = -1
            switch "foo bar" {
            case "foo bar":
                var v26 int16 = (v2 << (((^(v4))) & 63))
                _ = v26
            case "":
                v3 = (v3 + (v3 + "\n\t"))
            case "\n\t":
                var v27 int32 = v25
                _ = v27
            }
            _ = v25
        }
        var v28 float64 = 0.1
        if false {
            switch v3 {
            case "":
                for fi29 := 0; fi29 < 17; fi29++ {
                    if ((v2 == v2) && (128 == 128)) {
                        v30 := GenS3{F0: 65535}
                        v28 -= ((((v28 - -2.25) + (v28 / (v28 / 1.5))) - 1.5) * (float64(0.1) / float64((0.1 * 0.1))))
                        _ = v30
                    }
                    _ = fi29
                }
            default:
                var v31 float64 = -2.25
                _ = v31
            }
            var v32 int32 = 42
            _ = v32
        } else {
            v33 := GenS3{F0: 570}
            var v34 int32 = -2147483648
            _ = v33
            _ = v34
        }
        _ = fi24
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
