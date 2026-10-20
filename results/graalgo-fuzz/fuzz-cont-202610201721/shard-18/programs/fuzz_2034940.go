package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint }
func genFn0(p0 int32, p1 uint64, p2 uint32) uint {
    var v7 uint16 = 1
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 65535
}

func main() {
    var v1 uint16 = 32767
    v2 := []int{34, 32768, 9223372036854775807}
    var v3 uint32 = 2
    var v4 bool = false
    var v5 int16 = 670
    var v6 float32 = 0.1
    v8 := GenS1{F0: (0 == 42)}
    for fi9 := 0; fi9 < 10; fi9++ {
        v3 -= (v3 % (((v3 * 4294967295)) | 1))
        v3 += 42
        _ = fi9
    }
    var v10 uint8 = 128
    if len(v2) > 0 {
        v2[0] = -334
    }
    _ = len(v2)
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            if len(v2) > 0 {
                v2[0] = (fi12 >> ((1) & 63))
            }
            v10 += (v10 ^ v10)
            if ((fi11 | 9223372036854775807) == 128) {
                var v13 uint = 2
                _ = v13
            }
            _ = fi12
        }
        _ = fi11
    }
    v14 := GenS2{F0: false, F1: (("hello" + "foo bar") + "\n\t")}
    if v4 {
        for fi15 := 0; fi15 < 1; fi15++ {
            v4 = v4
            _ = fi15
        }
        fmt.Println(v6)
        v16 := GenS3{F0: 18446744073709551615}
        _ = v16
    } else {
        var v17 int8 = -67
        var v18 uint8 = (v10 & uint8(v1))
        _ = v17
        _ = v18
    }
    if len(v2) > 0 {
        v2[0] = int(v3)
    }
    v10 -= v10
    v4 = (v4 || false)
    if len(v2) > 0 {
        v2[0] = -694
    }
    var v19 bool = (true && (false && (!(47 >= -1))))
    var v20 uint8 = v10
    for fi21 := 0; fi21 < 6; fi21++ {
        switch fi21 {
        case -9223372036854775808:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 42:
            var v22 bool = (genFn0(-2147483648, 0, (^(v3))) < 543)
            _ = v22
        case 833:
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 uint = 255
                _ = fi23
                _ = v24
            }
        default:
            switch "" {
            case "a":
                v25 := GenS1{F0: (((v1 - (v1 - (v1 * v1))) < 937) && (true || (v1 >= (v1 << ((42) & 63)))))}
                _ = v25
            }
        }
        if ((v6 * 1.5) <= (-2.25 / (v6 - (((v6 + (-2.25 - (0.1 * float32(((-2.25 / 0.1) - ((((float64(0.0) * 1e10) * 0.0) * (-2.25 + (1e10 + float64(1e10)))) * ((float64(float32(1e10)) * (3.14159 / 1.5)) * float64(((((v6 - v6) / 1.5) / 3.14159) / v6))))))))) - 1.5) / ((((v6 - ((-2.25 * (v6 * float32((1.5 + (float64(0.1) + 0.1))))) / (((v6 / 1e10) - 1.5) * ((float32((3.14159 * 3.14159)) + ((v6 * 3.14159) + 1.5)) / v6)))) - v6) * 0.0) * 1.5))))) {
            var v26 int8 = -128
            if len(v2) > 0 {
                v2[0] = 127
            }
            _ = v26
        } else {
            v3 = 2
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi21
    }
    var v27 int8 = 7
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
