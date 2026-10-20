package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int8; F1 uint32; F2 uint64 }

func main() {
    var v1 uint32 = 1
    var v2 int16 = -32768
    v3 := make(map[uint32]bool)
    var v4 int8 = 0
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := []int{-1, (fi5 & 32768)}
        v4 = (v4 << ((255) & 63))
        fmt.Println(v4, v4)
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 uint8 = 166
            if (!(1.5 != 1e10)) {
                var v10 uint64 = 42
                var v11 int64 = -794
                _ = v10
                _ = v11
            } else {
                var v12 bool = (0 == 724)
                v13 := GenS1{F0: (v12 || (v12 && v12)), F1: 0}
                _ = v12
                _ = v13
            }
            switch fi8 {
            case 1:
                v1 = uint32(v9)
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v14 := []int{0, 9223372036854775807}
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    var v15 uint32 = 1
    v16 := GenS1{F0: false, F1: 65535}
    v4 = 42
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 string = "\n\t"
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: true, F1: 324}
    for fi20 := 0; fi20 < 7; fi20++ {
        v21 := GenS1{F0: (v2 != (v2 << ((213) & 63))), F1: 0}
        v1 += v1
        _ = fi20
        _ = v21
    }
    v22 := GenS1{F0: true, F1: 644}
    for fi23 := 0; fi23 < 16; fi23++ {
        switch -536 {
        case -268:
            var v24 bool = false
            _ = v24
        case 0:
            v25 := GenS1{F0: (!true), F1: 841}
            _ = v25
        }
        _ = fi23
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        fmt.Println(v2, v2, fi26)
        if (42 != 0) {
            var v27 int64 = -373
            _ = v27
        }
        _ = fi26
    }
    if (true && (v2 != 42)) {
        for fi28 := 0; fi28 < 1; fi28++ {
            _ = len(v14)
            _ = fi28
        }
    } else {
        fmt.Println(v1, v15, v15)
    }
    var v29 int = -1
    switch "\n\t" {
    case "\n\t":
        v1 = v15
    case "":
        for fi30 := 0; fi30 < 7; fi30++ {
            _ = len(v3)
            if (42 <= (v2 * v2)) {
                fmt.Println(v4)
            } else {
                v1 += 452
            }
            if false {
                if true {
                    var v31 float32 = (-2.25 * (((0.1 + (0.1 - ((1.5 - 0.1) - ((3.14159 + float32(-2.25)) - 0.0)))) + 0.0) - (float32(3.14159) + float32(3.14159))))
                    _ = v31
                }
                v32 := GenS2{F0: v4, F1: 4294967295, F2: 358}
                _ = v32
            }
            _ = fi30
        }
    }
    v4 += v4
    if (((!false) && ((128 >= 255) || true)) && true) {
        fmt.Println(v15, v15, v4)
        v33 := v3[(v1 / (((^(v1))) | 1))]
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v4
}
