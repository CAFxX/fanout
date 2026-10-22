package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int }
func genFn0(p0 uint64) int32 {
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 float64 = ((((0.1 / (float64(3.14159) * (float64(((((float32(1e10) - 0.1) / 3.14159) * -2.25) / 1e10)) - 1.5))) / (1.5 - (((0.0 - 0.1) / 1.0) + 1e10))) - (float64((float32(1e10) * ((((float32((3.14159 - -2.25)) - (float32(float64((0.0 - 1e10))) / 1.5)) * 1e10) + 1e10) - (((3.14159 / 0.1) / 1e10) / 1.0)))) * (float64(((1e10 + 0.0) * 0.1)) + float64((((1e10 / 1.0) / 1.5) + 3.14159))))) + -2.25)
        _ = fi7
        _ = v8
    }
    var v9 float64 = 1.5
    switch -2147483648 {
    case 32768:
        fmt.Println(p0)
    }
    switch 128 {
    case 32768:
        var v10 int = -384
        _ = v10
    case -226:
        var v11 int = -9223372036854775808
        _ = v11
    }
    _ = p0
    _ = v9
    return 886
}
func genFn1(p0 bool, p1 uint8) uint32 {
    for fi12 := 0; fi12 < 1; fi12++ {
        switch fi12 {
        case 4294967295:
            var v13 uint16 = 32767
            _ = v13
        }
        _ = fi12
    }
    v14 := GenS1{F0: 0}
    var v15 float32 = 0.1
    p1 = 255
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 2147483647
}

func main() {
    var v1 int8 = 1
    v2 := []int{453, -2147483648, -2147483648}
    v3 := []uint32{42}
    var v4 int64 = -9223372036854775808
    var v5 float64 = (1e10 * ((3.14159 * 0.0) + ((1e10 + 0.1) / -2.25)))
    var v6 int16 = -1
    switch "\n\t" {
    case "foo bar":
        if ("" <= "\n\t") {
            v16 := GenS1{F0: 1}
            if ((v5 / v5) <= ((v5 - v5) - v5)) {
                v5 += v5
                v17 := GenS1{F0: 42}
                var v18 uint16 = 255
                _ = v17
                _ = v18
            }
            _ = v16
        }
    case "":
        for fi19 := 0; fi19 < 19; fi19++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v1 -= v1
            var v20 uint64 = 497
            _ = fi19
            _ = v20
        }
    default:
        v5 -= (v5 + 3.14159)
    }
    var v21 int64 = v4
    v22 := GenS2{F0: false, F1: 127}
    var v23 uint64 = 277
    v5 += 0.0
    v24 := GenS1{F0: 0}
    v4 -= (v4 >> ((0) & 63))
    var v25 uint32 = 976
    if (v1 >= int8(v23)) {
        v21 = v4
    }
    if (false && true) {
        switch "foo bar" {
        case "\n\t":
            switch -1 {
            case 99:
                switch 127 {
                case 2147483647:
                    var v26 bool = false
                    _ = v26
                default:
                    var v27 uint32 = v25
                    _ = v27
                }
            case 508:
                var v28 uint32 = 0
                _ = v28
            }
        }
    } else {
        var v29 int = 4294967295
        v6 = v6
        for fi30 := 0; fi30 < 3; fi30++ {
            v25 += (v25 + v25)
            if true {
                v29 -= fi30
                if len(v2) > 0 {
                    v2[0] = (fi30 * (fi30 << ((0) & 63)))
                }
            }
            _ = fi30
        }
        _ = v29
    }
    var v31 int = -854
    var v32 bool = (((v25 | v25) >= 591) || (!(false && false)))
    for fi33 := 0; fi33 < 4; fi33++ {
        v4 -= (v21 | (v4 % (((v21 & v21)) | 1)))
        v34 := GenS3{F0: 203, F1: v31}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi33
        _ = v34
    }
    fmt.Println(v5)
    var v35 string = "foo bar"
    for fi36 := 0; fi36 < 13; fi36++ {
        v31 += fi36
        _ = fi36
    }
    v37 := GenS3{F0: 127, F1: 32767}
    if len(v2) > 0 {
        v2[0] = int(v25)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
