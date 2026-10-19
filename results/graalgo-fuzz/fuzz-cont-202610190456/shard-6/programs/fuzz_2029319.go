package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint64, p1 uint32, p2 int8) uint16 {
    if false {
        p1 += (p1 * p1)
        if false {
            v4 := []int{794, 32768, 1, 0}
            switch (("foo bar" + "") + ("foo bar" + "hello")) {
            case "a":
                v5 := GenS1{F0: 9223372036854775807}
                _ = v5
            }
            if len(v4) > 0 {
                v4[0] = -616
            }
            _ = v4
        } else {
            for fi6 := 0; fi6 < 20; fi6++ {
                var v7 float32 = 1e10
                v7 = ((0.0 - v7) + (v7 / v7))
                _ = fi6
                _ = v7
            }
        }
    } else {
        var v8 uint32 = p1
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 float32 = -2.25
        var v11 uint8 = 1
        switch fi9 {
        case -564:
            var v12 uint = 327
            _ = v12
        case -581:
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 bool = (971 <= 0)
                _ = fi13
                _ = v14
            }
        case -9223372036854775808:
            if (true || false) {
                for fi15 := 0; fi15 < 2; fi15++ {
                    v10 += (0.0 / float32((1e10 + float64((float32(((0.0 / 1.0) + (1e10 - float64(float32(3.14159))))) - 1e10)))))
                    _ = fi15
                }
            }
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    var v16 uint64 = (p0 << ((553) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 1
}

func main() {
    var v1 int = -1
    v2 := []uint32{2147483648, 2147483647, 550, 42}
    var v3 float32 = (0.1 + ((float32(float64(((1e10 - float32(0.1)) + 0.1))) + float32(0.0)) - -2.25))
    v3 = v3
    fmt.Println(v1, v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v17 int8 = -75
    switch (v1 >> ((255) & 63)) {
    case -926:
        var v18 float64 = float64(float32(float64(float32(3.14159))))
        _ = v18
    case -2147483648:
        var v19 string = ("\n\t" + "a")
        _ = v19
    case 9223372036854775807:
        if (!false) {
            if len(v2) > 0 {
                v2[0] = 297
            }
            v17 += 42
        }
    default:
        var v20 int16 = 1
        _ = v20
    }
    v21 := GenS2{F0: -539}
    v2 = append(v2, 0)
    v17 -= (v17 ^ (v17 % ((v17) | 1)))
    v3 -= (float32(3.14159) - 1.5)
    var v22 int = v1
    fmt.Println(v17, v22)
    for fi23 := 0; fi23 < 17; fi23++ {
        if ((735 != 32767) || true) {
            if true {
                for fi24 := 0; fi24 < 4; fi24++ {
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                    var v25 int64 = 670
                    fmt.Println(v22)
                    v1 -= fi24
                    _ = fi24
                    _ = v25
                }
                var v26 int8 = (-(v17))
                _ = v26
            }
            v17 -= (v17 ^ 42)
            var v27 uint8 = 128
            _ = v27
        }
        _ = fi23
    }
    if ((!false) && (!false)) {
        switch (v1 ^ v1) {
        case 9223372036854775807:
            var v28 float64 = 0.0
            _ = v28
        case -1:
            var v29 float64 = (1.5 / (3.14159 / ((0.1 + 0.1) / (3.14159 * (((0.0 / float64((v3 - (float32((1.5 - 1.5)) - v3)))) + 1.5) - -2.25)))))
            _ = v29
        case 2147483647:
            if (false && true) {
                fmt.Println(v22, v1, v22)
                v17 -= 0
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        default:
            _ = len(v2)
        }
    } else {
        v1 -= v22
        switch "" {
        case "":
            var v30 float64 = ((((0.1 + (1.5 * (0.0 - 3.14159))) * (0.1 * (1.5 / 1e10))) + (0.0 - (1e10 - 1.5))) - 0.1)
            _ = v30
        case "a":
            v31 := GenS2{F0: 9223372036854775807}
            _ = v31
        case "hello":
            var v32 float64 = 0.0
            _ = v32
        default:
            v33 := GenS2{F0: 650}
            _ = v33
        }
    }
    var v34 bool = (("hello" + "a") > (("" + "hello") + ("" + (("hello" + "hello") + "hello"))))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v34
}
