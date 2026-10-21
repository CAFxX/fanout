package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 string; F1 int16 }
func genFn0(p0 int8, p1 int8, p2 int32) bool {
    fmt.Println(p2)
    var v7 int = 42
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS1{F0: -1}
        v10 := GenS1{F0: -941}
        v11 := GenS1{F0: 42}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 uint16 = 278
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v7
    return (633 > 9223372036854775807)
}

func main() {
    var v1 uint8 = 2
    v2 := []string{}
    v3 := []string{"\n\t"}
    var v4 float32 = 1.5
    v5 := []int{1, -2147483648}
    var v6 float32 = ((v4 / v4) + (((((0.1 / ((v4 - (((-2.25 - (((v4 - v4) + float32(((1e10 + ((float64((float32(0.1) - (v4 * float32(float64((3.14159 - v4)))))) / 3.14159) - 0.0)) / 3.14159))) + -2.25)) / 3.14159) - v4)) * ((0.1 + v4) / 0.1))) / 3.14159) + float32(float64(0.1))) / v4) * v4))
    v13 := GenS2{F0: 127}
    fmt.Println(v6, v1, v4)
    v14 := GenS3{F0: 181}
    for fi15 := 0; fi15 < 5; fi15++ {
        v6 = v4
        if false {
            var v16 bool = ((255 > 255) && false)
            var v17 uint64 = 4294967295
            _ = v16
            _ = v17
        }
        if true {
            var v18 int16 = 1
            _ = v18
        }
        _ = fi15
    }
    var v19 float64 = 1e10
    v20 := GenS4{F0: (false && true), F1: -241}
    var v21 uint64 = 2
    v22 := GenS5{F0: "", F1: -1}
    v4 += float32(v19)
    var v23 uint16 = 0
    var v24 int16 = -32768
    if (!((42 != 371) && false)) {
        fmt.Println(v6)
        switch "a" {
        case "a":
            if (!(uint64(v23) != (v21 | v21))) {
                var v25 float32 = (v4 * (1e10 + ((0.0 * v6) * float32(v19))))
                var v26 uint32 = 692
                var v27 int32 = 0
                _ = v25
                _ = v26
                _ = v27
            }
        case "foo bar":
            var v28 bool = (true || false)
            _ = v28
        case "\n\t":
            var v29 float32 = (v6 + v4)
            _ = v29
        default:
            var v30 float64 = v19
            _ = v30
        }
        v19 = float64(v6)
    } else {
        fmt.Println(v1, v6)
        var v31 float32 = 1.5
        _ = v31
    }
    if true {
        var v32 int16 = v24
        for fi33 := 0; fi33 < 3; fi33++ {
            for fi34 := 0; fi34 < 16; fi34++ {
                var v35 float32 = v4
                _ = fi34
                _ = v35
            }
            for fi36 := 0; fi36 < 5; fi36++ {
                v2 = append(v2, "")
                var v37 uint64 = v21
                v38 := GenS1{F0: -1}
                _ = fi36
                _ = v37
                _ = v38
            }
            switch (-(fi33)) {
            case 127:
                for fi39 := 0; fi39 < 19; fi39++ {
                    v6 = float32((v19 / 1.5))
                    fmt.Println(v1)
                    _ = fi39
                }
            default:
                var v40 int = fi33
                _ = v40
            }
            _ = fi33
        }
        if true {
            if ("hello" == "hello") {
                v41 := GenS1{F0: 824}
                _ = v41
            }
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v42 := GenS5{F0: "hello", F1: (^(v32))}
            _ = v42
        }
        _ = v32
    }
    for fi43 := 0; fi43 < 12; fi43++ {
        var v44 float32 = float32((v19 - v19))
        _ = fi43
        _ = v44
    }
    v45 := GenS2{F0: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
