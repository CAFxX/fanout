package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 string; F1 uint64; F2 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 uint64 = 675
    var v2 float64 = (((-2.25 * (3.14159 - (0.0 + ((((3.14159 * (0.0 * (float64(((0.1 - (-2.25 + 3.14159)) + 1e10)) / 1e10))) * 1e10) * 0.1) - ((float64(1.5) / 1.0) - (float64(float32(-2.25)) - (((3.14159 + (-2.25 - ((float64((1.5 - float32(3.14159))) + 3.14159) / 1e10))) * 3.14159) + ((1.5 * float64((-2.25 / 1.0))) - ((1e10 + (-2.25 + float64(1e10))) - ((1.5 - float64(0.1)) * -2.25)))))))))) - (0.0 * 3.14159)) * ((-2.25 - -2.25) - 1.5))
    v3 := []string{"a", "hello", ("foo bar" + "a")}
    var v4 int32 = 756
    var v5 string = "a"
    var v6 float64 = float64(3.14159)
    fmt.Println(v2, v5)
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint = 692
        var v9 int32 = (v4 >> (((v8 + v8)) & 63))
        _ = fi7
        _ = v8
        _ = v9
    }
    v2 = v2
    var v10 int16 = -573
    _ = len(v3)
    v4 -= (v4 / (((v4 ^ v4)) | 1))
    fmt.Println(v2, v6, v10)
    switch "a" {
    case "hello":
        if (127 >= 127) {
            for fi11 := 0; fi11 < 8; fi11++ {
                v10 -= v10
                var v12 bool = false
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 float64 = (1e10 * (v6 / 1e10))
                    _ = fi13
                    _ = v14
                }
                _ = fi11
                _ = v12
            }
        } else {
            var v15 uint8 = 255
            var v16 int32 = -541
            _ = v15
            _ = v16
        }
    }
    if true {
        fmt.Println(v10, v1)
    }
    v1 = 0
    v2 = 0.0
    var v17 int8 = -1
    for fi18 := 0; fi18 < 14; fi18++ {
        for fi19 := 0; fi19 < 19; fi19++ {
            if false {
                v17 -= v17
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if false {
                        var v20 uint64 = v1
                        _ = v20
                    } else {
                        v4 = -1
                        fmt.Println(v2, fi18)
                    }
                }
                var v21 int = fi19
                _ = v21
            }
            v22 := GenS1{F0: 255, F1: 503}
            _ = fi19
            _ = v22
        }
        fmt.Println(v10, v6)
        _ = fi18
    }
    var v23 float32 = 1e10
    if true {
        v6 += v6
    } else {
        var v24 uint32 = 1
        v10 = v10
        _ = v24
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        v3 = append(v3, "a")
        _ = fi25
    }
    fmt.Println(v17)
    if (true && true) {
        for fi26 := 0; fi26 < 12; fi26++ {
            if true {
                v5 = "hello"
            }
            _ = fi26
        }
        if ((false || true) || true) {
            v27 := GenS2{F0: v5, F1: v1, F2: int(v1)}
            v28 := GenS2{F0: ("hello" + ("a" + v5)), F1: v1, F2: 415}
            v29 := GenS3{F0: 744}
            _ = v27
            _ = v28
            _ = v29
        }
        for fi30 := 0; fi30 < 13; fi30++ {
            var v31 uint32 = 2
            var v32 uint = 4294967295
            _ = fi30
            _ = v31
            _ = v32
        }
    } else {
        switch v5 {
        case "hello":
            if false {
                v4 = (v4 | v4)
            }
        case "":
            v33 := GenS1{F0: 128, F1: 2}
            _ = v33
        default:
            if ((v10 % ((1000) | 1)) == (v10 - (^(v10)))) {
                var v34 string = v5
                _ = v34
            } else {
                v35 := GenS1{F0: 16, F1: 255}
                var v36 uint16 = 2
                _ = v35
                _ = v36
            }
        }
        v37 := GenS4{F0: (v1 + (v1 % ((v1) | 1)))}
        var v38 uint8 = 244
        _ = v37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
