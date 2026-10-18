package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint16; F1 uint16; F2 int }
type GenS4 struct { F0 uint16; F1 int64 }
func genFn0(p0 int, p1 int, p2 uint8) int8 {
    if (0 != -323) {
        if (!false) {
            var v7 uint32 = 181
            p2 -= 42
            var v8 float64 = 1.5
            _ = v7
            _ = v8
        }
        p0 = 2147483647
        p2 = p2
    } else {
        p1 = 449
    }
    var v9 uint = 573
    var v10 string = (("" + "\n\t") + "\n\t")
    var v11 uint64 = 2
    var v12 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 103
}

func main() {
    var v1 float32 = ((0.0 * ((-2.25 + 0.0) / (1e10 * 3.14159))) - 1.5)
    v2 := make(map[string]string)
    var v3 int64 = 9223372036854775807
    v4 := make(map[string]bool)
    v5 := make(map[uint32]string)
    v6 := make(map[string]uint32)
    fmt.Println(v3, v1)
    for fi13 := 0; fi13 < 4; fi13++ {
        switch 890 {
        case -648:
            v14 := []int{fi13, fi13, fi13, 4294967295}
            _ = v14
        default:
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 int8 = genFn0(fi13, fi15, 0)
                _ = fi15
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 float64 = float64((float32((-2.25 + 1e10)) * ((((v1 + 3.14159) + (v1 - (0.1 + v1))) - (v1 - v1)) * (v1 - -2.25))))
            if false {
                var v19 int32 = 2147483647
                _ = v19
            } else {
                switch -1 {
                case 32767:
                    v20 := []int{fi13, fi13, -669, fi17}
                    _ = v20
                case 42:
                    if (!true) {
                        v21 := GenS1{F0: 144}
                        _ = v21
                    } else {
                        v22 := GenS1{F0: 283}
                        _ = v22
                    }
                }
                var v23 uint8 = 32
                v24 := GenS1{F0: 65535}
                _ = v23
                _ = v24
            }
            var v25 uint = 2
            _ = fi17
            _ = v18
            _ = v25
        }
        _ = fi13
    }
    switch "a" {
    case "hello":
        v26 := GenS2{F0: 126, F1: 128}
        _ = v26
    case "\n\t":
        switch 315 {
        case -2147483648:
            switch -1 {
            case 0:
                fmt.Println(v3, v3)
            default:
                v1 = v1
            }
        }
    case "foo bar":
        v1 -= 0.0
    }
    v27 := GenS3{F0: 65535, F1: 0, F2: 32768}
    var v28 float64 = (-2.25 * 3.14159)
    v29 := []int{-832, 32768, -924, 127}
    if true {
        var v30 uint32 = 801
        _ = v30
    }
    v3 += v3
    for fi31 := 0; fi31 < 16; fi31++ {
        var v32 uint8 = 1
        var v33 uint32 = 4294967295
        switch fi31 {
        case 42:
            for fi34 := 0; fi34 < 2; fi34++ {
                var v35 int8 = 127
                _ = fi34
                _ = v35
            }
        case 832:
            var v36 uint16 = 1
            _ = v36
        case 32767:
            v37 := GenS1{F0: uint32(v32)}
            _ = v37
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    delete(v2, ("" + "hello"))
    if len(v29) >= 2 {
        v29 = v29[1:len(v29)]
    }
    var v38 bool = (((0 != 42) && false) || true)
    if ((v38 || ((!v38) && v38)) && (!(-32768 <= 0))) {
        fmt.Println(v38, v38, v1)
        v3 = (v3 / ((-9223372036854775808) | 1))
        var v39 string = "foo bar"
        _ = v39
    }
    v40 := GenS4{F0: 42, F1: 42}
    v41 := v6["\n\t"]
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v38
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
