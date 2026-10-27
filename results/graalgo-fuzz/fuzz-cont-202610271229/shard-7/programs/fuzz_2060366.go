package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 bool }
type GenS2 struct { F0 int8; F1 int16; F2 int64 }
type GenS3 struct { F0 int16; F1 int32; F2 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 bool; F1 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{}
    v3 := make(map[string]string)
    if true {
        var v4 int64 = 42
        v4 = 243
        _ = v4
    }
    if ((769 >= 1) && true) {
        if false {
            for fi5 := 0; fi5 < 1; fi5++ {
                var v6 uint16 = 1
                var v7 float32 = 1.5
                var v8 uint16 = (^(v6))
                _ = fi5
                _ = v6
                _ = v7
                _ = v8
            }
            if ("hello" >= "") {
                var v9 int = -495
                _ = v9
            }
        } else {
            var v10 int32 = 42
            for fi11 := 0; fi11 < 3; fi11++ {
                v10 = (v10 >> ((4294967295) & 63))
                var v12 string = "a"
                if len(v2) > 0 {
                    v2[0] = (v12 + v12)
                }
                _ = fi11
                _ = v12
            }
            if ((v10 - v10) > -1) {
                v13 := GenS1{F0: 18446744073709551615, F1: 0, F2: true}
                _ = v13
            }
            _ = v10
        }
        var v14 uint8 = 128
        _ = v14
    } else {
        v15 := GenS2{F0: -32, F1: 1000, F2: -9223372036854775808}
        if (float64((0.1 + float32(-2.25))) <= (((1.5 * 1.5) * 3.14159) + (0.1 - -2.25))) {
            if ((455 < 65535) && (65535 <= 481)) {
                var v16 uint = 65535
                _ = v16
            } else {
                for fi17 := 0; fi17 < 4; fi17++ {
                    var v18 int64 = -1
                    fmt.Println(fi17, v18)
                    v19 := GenS2{F0: 0, F1: 42, F2: (v18 << ((255) & 63))}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                var v20 int = -958
                _ = v20
            }
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
        } else {
            var v21 uint32 = 42
            _ = v21
        }
        v22 := GenS1{F0: 2, F1: 255, F2: false}
        _ = v15
        _ = v22
    }
    var v23 string = "a"
    if (!(true && (-623 > 1))) {
        v24 := GenS3{F0: 1, F1: 1, F2: 0}
        v23 = ("" + (v23 + ""))
        _ = v24
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        delete(v1, 2)
        _ = fi25
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v23 = ((v23 + v23) + v23)
    var v26 uint8 = 128
    v26 += 255
    v26 = 1
    fmt.Println(v26, v23)
    var v27 int16 = -891
    var v28 uint16 = 42
    switch v23 {
    case "\n\t":
        var v29 float32 = ((((3.14159 * ((3.14159 + (3.14159 * (1.5 - 0.1))) / 1.5)) - ((0.0 - -2.25) * (((((1e10 * -2.25) - (float32((-2.25 - float64(0.0))) - (1e10 * ((((1.5 - 1.5) * 0.1) * 1e10) / 1.0)))) * (1.5 / (-2.25 * -2.25))) + 1.5) / (-2.25 / 0.1)))) / ((((float32((float64(3.14159) / 0.1)) - 3.14159) / -2.25) + 0.1) + 3.14159)) * 3.14159)
        _ = v29
    case "foo bar":
        v30 := GenS4{F0: (-15 >= -1)}
        _ = v30
    case "":
        v31 := GenS5{F0: 42, F1: true}
        _ = v31
    default:
        fmt.Println(v27)
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        v27 += 175
        var v33 float32 = 0.0
        _ = fi32
        _ = v33
    }
    v34 := GenS6{F0: true, F1: (v26 ^ v26)}
    v35 := GenS3{F0: v27, F1: 42, F2: -922}
    v36 := GenS1{F0: 4294967295, F1: v28, F2: true}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v36
}
