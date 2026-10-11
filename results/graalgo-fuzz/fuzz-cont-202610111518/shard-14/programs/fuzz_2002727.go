package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int }
func genFn0(p0 string, p1 uint, p2 int8) int16 {
    var v6 uint16 = 0
    if ((p0 + p0) != "hello") {
        var v7 int32 = -2147483648
        v7 = (^(v7))
        v8 := GenS1{F0: p0, F1: v7, F2: (p0 + "")}
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: ("\n\t" + p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return 42
}
func genFn1(p0 string, p1 string) uint64 {
    var v10 int8 = 127
    v11 := []int{573, 42}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 684
    var v3 float32 = (0.0 / 1e10)
    var v4 uint8 = 149
    var v5 float64 = (float64((v3 / v3)) * -2.25)
    v2 -= v2
    v5 -= float64(v3)
    if false {
        var v12 int8 = 100
        _ = v12
    } else {
        var v13 float64 = -2.25
        v14 := GenS2{F0: ""}
        _ = v13
        _ = v14
    }
    v2 = 2147483648
    var v15 uint16 = 706
    v16 := GenS1{F0: "foo bar", F1: -2147483648, F2: "a"}
    v2 -= v2
    v17 := GenS3{F0: 153, F1: -1}
    switch "foo bar" {
    case "\n\t":
        v3 = ((((3.14159 / v3) - (((((v3 / (float32(v5) / float32((((v5 / v5) - ((v5 / (float64(v3) / v5)) / (v5 * 3.14159))) * float64(0.0))))) * v3) - (3.14159 - ((v3 - (((-2.25 * (v3 * (((float32((v5 * (1e10 / v5))) - float32((1e10 - (3.14159 / v5)))) * (v3 - (float32(float64(((v3 + 3.14159) * 1.5))) + (((float32(float64(v3)) - float32((0.0 - v5))) / 1.0) + (v3 + v3))))) - 1.5))) * 0.0) * 3.14159)) + 1e10))) + (((v3 - ((((0.0 / (((1.5 * 1e10) / v3) - ((float32(v5) * v3) - -2.25))) * v3) + 1e10) + -2.25)) - (((v3 / (0.0 * ((1.5 * ((v3 + (((v3 - (1e10 * float32((((v5 * ((float64(v3) + ((1.5 * 1e10) * float64(((float32(3.14159) - v3) * float32((0.1 / v5)))))) + -2.25)) * v5) / (3.14159 * ((v5 / v5) - v5)))))) * (v3 - v3)) - 1.5)) - v3)) * ((0.1 * 0.1) * v3)))) - v3) / v3)) + v3)) + (v3 - float32(v5)))) + (v3 + 1e10)) - ((v3 - v3) * (float32(v5) - v3)))
    case "hello":
        for fi18 := 0; fi18 < 17; fi18++ {
            var v19 uint64 = 42
            v20 := GenS2{F0: "\n\t"}
            v21 := GenS2{F0: "foo bar"}
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
    case "a":
        v2 += (v2 * 294)
    default:
        if true {
            _ = len(v1)
            if (255 <= 657) {
                if false {
                    if true {
                        var v22 uint64 = 1
                        _ = v22
                    }
                }
                for fi23 := 0; fi23 < 16; fi23++ {
                    v24 := GenS3{F0: (v4 + (v4 ^ v4)), F1: 1}
                    _ = fi23
                    _ = v24
                }
            } else {
                fmt.Println(v2, v2, v5)
                if (v5 != (float64(v3) * v5)) {
                    v3 = v3
                    v25 := GenS3{F0: v4, F1: 53}
                    var v26 int = 1
                    _ = v25
                    _ = v26
                } else {
                    v27 := []int{1, 127}
                    _ = v27
                }
                switch 32767 {
                case 42:
                    var v28 int = 42
                    _ = v28
                case 734:
                    var v29 uint8 = (v4 | (^(v4)))
                    _ = v29
                default:
                    v30 := GenS2{F0: ""}
                    _ = v30
                }
            }
        }
    }
    var v31 int16 = -1
    v4 -= v4
    for fi32 := 0; fi32 < 11; fi32++ {
        var v33 bool = false
        var v34 int16 = v31
        for fi35 := 0; fi35 < 3; fi35++ {
            if (v33 && true) {
                v4 = v4
                v31 = (v34 << ((1) & 63))
                v36 := []int{2147483647, -9223372036854775808, fi35}
                _ = v36
            }
            _ = fi35
        }
        _ = fi32
        _ = v33
        _ = v34
    }
    for fi37 := 0; fi37 < 14; fi37++ {
        switch (fi37 / ((fi37) | 1)) {
        case 9223372036854775807:
            var v38 uint32 = (v2 << ((2) & 63))
            _ = v38
        case 42:
            var v39 string = "\n\t"
            _ = v39
        case -9223372036854775808:
            var v40 float64 = float64(v3)
            _ = v40
        }
        _ = fi37
    }
    for fi41 := 0; fi41 < 1; fi41++ {
        var v42 float32 = v3
        v15 += (v15 % ((v15) | 1))
        _ = fi41
        _ = v42
    }
    if false {
        switch 128 {
        case 9223372036854775807:
            for fi43 := 0; fi43 < 18; fi43++ {
                var v44 uint8 = (v4 * v4)
                v31 = v31
                _ = fi43
                _ = v44
            }
        }
        var v45 string = "\n\t"
        _ = v45
    } else {
        for fi46 := 0; fi46 < 14; fi46++ {
            fmt.Println(v15, v31)
            delete(v1, ("" + "foo bar"))
            v2 += v2
            _ = fi46
        }
    }
    for fi47 := 0; fi47 < 11; fi47++ {
        var v48 int16 = (v31 & (v31 | -1))
        fmt.Println(v5, v3, v48)
        v15 = 1
        _ = fi47
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
