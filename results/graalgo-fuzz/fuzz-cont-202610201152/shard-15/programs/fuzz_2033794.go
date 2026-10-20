package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint16 }
func genFn0() int32 {
    switch 9223372036854775807 {
    case -106:
        var v4 uint = 776
        _ = v4
    case 127:
        var v5 int32 = -393
        _ = v5
    default:
        var v6 uint8 = 0
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 bool = false
        if (v8 && (0.0 > (float64(-2.25) * 0.0))) {
            for fi9 := 0; fi9 < 10; fi9++ {
                v8 = (165 <= 4294967295)
                _ = fi9
            }
            switch (fi7 + 9223372036854775807) {
            case 283:
                v8 = (((-2.25 / 1.5) / 1.5) <= 3.14159)
            default:
                var v10 int8 = -109
                _ = v10
            }
            var v11 float32 = float32(0.0)
            _ = v11
        } else {
            switch "hello" {
            case "hello":
                fmt.Println(v8)
            default:
                var v12 float64 = (3.14159 / ((1e10 / 0.1) + 1.5))
                _ = v12
            }
            fmt.Println(v8, v8)
        }
        v13 := GenS1{F0: 315, F1: 139}
        _ = fi7
        _ = v13
        _ = v8
    }
    if ((false && (!(false && (1 >= -1)))) && (127 < 4294967295)) {
        for fi14 := 0; fi14 < 8; fi14++ {
            if true {
                var v15 uint8 = 187
                if (((v15 | v15) != (v15 - 83)) || (false && true)) {
                    v16 := []int{(fi14 << ((65535) & 63)), 1}
                    var v17 bool = (28 >= v15)
                    _ = v16
                    _ = v17
                } else {
                    v18 := []int{fi14}
                    _ = v18
                }
                switch "foo bar" {
                case "foo bar":
                    switch (fi14 & -604) {
                    case 630:
                        var v19 uint8 = 255
                        _ = v19
                    }
                case "\n\t":
                    v15 += v15
                }
                _ = v15
            }
            fmt.Println(fi14)
            var v20 int16 = 1
            _ = fi14
            _ = v20
        }
    } else {
        var v21 uint16 = 985
        _ = v21
    }
    var v22 int32 = -677
    _ = v22
    return v22
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 53
    var v3 uint32 = 2
    for fi23 := 0; fi23 < 11; fi23++ {
        v3 = (^(v3))
        _ = fi23
    }
    v2 = (v2 ^ v2)
    v24 := GenS1{F0: 65535, F1: 2}
    var v25 float32 = (0.0 / ((0.0 - (1.5 - 1e10)) + float32(-2.25)))
    for fi26 := 0; fi26 < 10; fi26++ {
        for fi27 := 0; fi27 < 5; fi27++ {
            v28 := GenS1{F0: 255, F1: 902}
            v29 := GenS1{F0: 879, F1: 362}
            var v30 uint8 = 127
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
        switch (fi26 & (fi26 ^ (fi26 & fi26))) {
        case 1:
            for fi31 := 0; fi31 < 1; fi31++ {
                v25 = float32(-2.25)
                v32 := []int{}
                v3 = v3
                v3 += v3
                v32 = append(v32, (^(fi26)))
                var v33 float64 = 0.1
                _ = fi31
                _ = v32
                _ = v33
            }
        case -556:
            v34 := []int{}
            _ = v34
        default:
            for fi35 := 0; fi35 < 8; fi35++ {
                v36 := GenS1{F0: 42, F1: 32767}
                _ = fi35
                _ = v36
            }
        }
        _ = fi26
    }
    v2 = v2
    if true {
        var v37 int = 128
        _ = v37
    } else {
        var v38 uint32 = (v3 - 65535)
        v38 -= v38
        _ = v38
    }
    switch 127 {
    case -2147483648:
        switch "\n\t" {
        case "":
            for fi39 := 0; fi39 < 2; fi39++ {
                v40 := GenS2{F0: -997, F1: 255}
                var v41 uint = 18446744073709551615
                fmt.Println(v25)
                _ = fi39
                _ = v40
                _ = v41
            }
        case "a":
            v25 = float32(1.5)
        case "foo bar":
            var v42 int8 = 0
            _ = v42
        default:
            v43 := []int{32768, 32768}
            _ = v43
        }
    case 0:
        v44 := v1[65535]
        _ = v44
    case 544:
        fmt.Println(v25, v3)
    }
    var v45 uint16 = 255
    v2 -= v2
    v25 = ((float32((1.5 - 1.5)) * (v25 + 1.5)) - v25)
    var v46 int = 42
    v47 := GenS1{F0: 65535, F1: (v45 & v45)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v45
    _ = v46
    _ = v47
}
