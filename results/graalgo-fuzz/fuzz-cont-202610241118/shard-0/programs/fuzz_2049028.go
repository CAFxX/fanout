package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint16; F1 int; F2 int8 }
type GenS3 struct { F0 uint; F1 uint16; F2 bool }
func genFn0(p0 int64) uint32 {
    v4 := []int{0, -2147483648, 127, 695}
    var v5 int64 = p0
    _ = p0
    _ = v4
    _ = v5
    return 4294967295
}

func main() {
    var v1 uint = 878
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    v6 := []int{2147483647, 32767, -395, 4294967295}
    var v7 float32 = -2.25
    v8 := GenS1{F0: 2, F1: -486}
    fmt.Println(v7, v1)
    for fi9 := 0; fi9 < 2; fi9++ {
        switch fi9 {
        case 4294967295:
            var v10 int8 = 2
            _ = v10
        case -9223372036854775808:
            fmt.Println(fi9, fi9, v1)
        }
        if ("hello" <= "") {
            v7 += ((v7 + 3.14159) * v7)
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 uint16 = 109
                if len(v6) > 0 {
                    v6[0] = 90
                }
                var v13 float32 = v7
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        _ = fi9
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v15 := GenS2{F0: 65535, F1: fi14, F2: 127}
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 127, F1: -1}
    if false {
        if (true || true) {
            switch 123 {
            case 646:
                for fi17 := 0; fi17 < 6; fi17++ {
                    for fi18 := 0; fi18 < 15; fi18++ {
                        fmt.Println(fi18)
                        _ = fi18
                    }
                    v19 := GenS2{F0: 65535, F1: (fi17 | fi17), F2: 100}
                    _ = fi17
                    _ = v19
                }
            default:
                if false {
                    v20 := v3[65535]
                    var v21 int = 2147483647
                    v22 := GenS2{F0: 255, F1: v21, F2: 42}
                    _ = v20
                    _ = v21
                    _ = v22
                }
            }
            var v23 float64 = float64(v7)
            var v24 int8 = 15
            _ = v23
            _ = v24
        }
    }
    switch -931 {
    case -9223372036854775808:
        fmt.Println(v1)
    case -1:
        v25 := GenS3{F0: v1, F1: 587, F2: false}
        _ = v25
    default:
        v1 = v1
    }
    v26 := GenS1{F0: 166, F1: 32767}
    v27 := GenS3{F0: v1, F1: 721, F2: false}
    var v28 uint = v1
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        v29 := GenS3{F0: (^(v28)), F1: 42, F2: false}
        _ = v29
    case "foo bar":
        _ = len(v2)
    case "":
        fmt.Println(v7, v7)
    }
    var v30 bool = (1 != 0)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v6
    _ = v7
    _ = v8
}
