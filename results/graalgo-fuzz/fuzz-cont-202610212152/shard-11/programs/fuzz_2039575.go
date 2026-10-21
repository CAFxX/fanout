package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 bool; F2 int8 }
type GenS3 struct { F0 string }
func genFn0(p0 int32) uint64 {
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 int32 = -1
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 int = fi7
            v8 = (^(fi5))
            _ = fi7
            _ = v8
        }
        p0 += (p0 << ((255) & 63))
        _ = fi5
        _ = v6
    }
    p0 += (p0 << ((0) & 63))
    if (!(3.14159 != 1.5)) {
        p0 -= p0
        var v9 int16 = 1000
        _ = v9
    } else {
        for fi10 := 0; fi10 < 15; fi10++ {
            p0 += 2147483647
            switch fi10 {
            case 9223372036854775807:
                v11 := []int{}
                _ = v11
            case 127:
                v12 := []int{-2147483648, (fi10 * (^(fi10)))}
                _ = v12
            case -9223372036854775808:
                var v13 int64 = 42
                _ = v13
            default:
                if (false || true) {
                    for fi14 := 0; fi14 < 12; fi14++ {
                        p0 -= p0
                        var v15 int64 = 1
                        var v16 int8 = -128
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                }
            }
            _ = fi10
        }
        v17 := []int{-462}
        v18 := GenS1{F0: 127}
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        p0 = (p0 | 1)
        _ = fi19
    }
    _ = p0
    return 396
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = (!false)
    var v3 string = ("a" + "hello")
    v4 := make(map[uint32]int)
    for fi20 := 0; fi20 < 12; fi20++ {
        v21 := GenS1{F0: 230}
        _ = fi20
        _ = v21
    }
    var v22 int8 = 0
    if (!(false || (!true))) {
        var v23 bool = v2
        for fi24 := 0; fi24 < 3; fi24++ {
            var v25 int = 9223372036854775807
            v25 = fi24
            v26 := []int{v25, (v25 | (fi24 >> ((0) & 63))), (v25 % ((v25) | 1)), fi24}
            _ = fi24
            _ = v25
            _ = v26
        }
        v2 = true
        _ = v23
    } else {
        if v2 {
            fmt.Println(v22)
        } else {
            fmt.Println(v2, v22, v3)
        }
        for fi27 := 0; fi27 < 14; fi27++ {
            for fi28 := 0; fi28 < 8; fi28++ {
                var v29 uint64 = 1
                v30 := GenS2{F0: 65535, F1: (v2 || v2), F2: 100}
                var v31 int32 = 42
                v32 := v4[253]
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
                _ = v32
            }
            fmt.Println(v2)
            var v33 float32 = float32((3.14159 + 3.14159))
            _ = fi27
            _ = v33
        }
    }
    _ = len(v1)
    for fi34 := 0; fi34 < 18; fi34++ {
        _ = len(v1)
        if (!(157 < -2147483648)) {
            v3 = v3
            var v35 uint32 = 76
            _ = v35
        } else {
            for fi36 := 0; fi36 < 15; fi36++ {
                var v37 uint8 = 127
                v38 := GenS3{F0: "hello"}
                v22 -= (v22 - v22)
                v39 := []int{233, fi34}
                if len(v39) >= 2 {
                    v39 = v39[1:len(v39)]
                }
                switch v3 {
                case "hello":
                    v40 := GenS3{F0: v3}
                    _ = v40
                }
                _ = fi36
                _ = v37
                _ = v38
                _ = v39
            }
            var v41 uint8 = 42
            _ = v41
        }
        for fi42 := 0; fi42 < 12; fi42++ {
            v43 := GenS2{F0: 2147483648, F1: (!(((!true) || v2) || v2)), F2: 42}
            v44 := []int{fi34}
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = fi34
    }
    v45 := []int{53, -9223372036854775808}
    if (true || (69 > 127)) {
        v2 = v2
        var v46 uint64 = 107
        v2 = v2
        _ = v46
    } else {
        v22 = (v22 / ((22) | 1))
    }
    if v2 {
        v47 := GenS1{F0: 109}
        fmt.Println(v3)
        _ = v47
    } else {
        switch v3 {
        case "":
            var v48 int16 = 1
            _ = v48
        default:
            for fi49 := 0; fi49 < 2; fi49++ {
                var v50 float64 = 1.5
                var v51 uint32 = 2147483647
                _ = fi49
                _ = v50
                _ = v51
            }
        }
        v4[0] = 753
        v3 = (v3 + v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v45:", len(v45), v45)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v45
}
