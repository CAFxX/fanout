package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int }
func genFn0() int8 {
    v5 := []int{693, 1, 2147483647}
    v6 := GenS1{F0: 127, F1: ""}
    _ = v5
    _ = v6
    return 127
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = ("foo bar" + ("" + ""))
    var v3 int = -1
    v4 := make(map[string]uint32)
    if (3.14159 == (((0.0 + 0.1) + ((-2.25 * ((((3.14159 / ((1e10 + 0.1) - (-2.25 * float64(float32(1e10))))) * float64(((1e10 + -2.25) * 0.0))) / float64(3.14159)) * 3.14159)) + 0.0)) + 1.5)) {
        v2 = v2
        switch "a" {
        case "hello":
            switch (v3 ^ v3) {
            case 2147483647:
                var v7 int32 = 1
                _ = v7
            }
        }
    } else {
        if false {
            switch (("" + v2) + v2) {
            case "":
                switch "" {
                case "a":
                    v3 = 186
                }
            case "foo bar":
                v8 := GenS2{F0: -1}
                _ = v8
            case "\n\t":
                for fi9 := 0; fi9 < 8; fi9++ {
                    v2 = v2
                    _ = fi9
                }
            default:
                v2 = (v2 + "a")
            }
            switch (v2 + "\n\t") {
            case "hello":
                v10 := GenS3{F0: 2147483647, F1: (^(v3))}
                _ = v10
            case "\n\t":
                if (true || (80 == 2)) {
                    v3 += v3
                } else {
                    var v11 float32 = float32(3.14159)
                    v12 := GenS1{F0: genFn0(), F1: v2}
                    _ = v11
                    _ = v12
                }
            case "foo bar":
                for fi13 := 0; fi13 < 2; fi13++ {
                    var v14 uint64 = 0
                    _ = fi13
                    _ = v14
                }
            }
            v2 = v2
        } else {
            var v15 uint8 = 128
            _ = v15
        }
        delete(v1, 9223372036854775807)
    }
    var v16 int = (^(v3))
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 int16 = -520
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 bool = false
            fmt.Println(v2, v20)
            _ = fi19
            _ = v20
        }
        var v21 int64 = -9223372036854775808
        _ = fi17
        _ = v18
        _ = v21
    }
    fmt.Println(v16)
    var v22 int = v3
    for fi23 := 0; fi23 < 18; fi23++ {
        var v24 string = (("foo bar" + "hello") + (v2 + "hello"))
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        if (2 > 620) {
            v16 -= (v3 ^ (v16 | 871))
        } else {
            for fi26 := 0; fi26 < 8; fi26++ {
                v2 = "foo bar"
                v2 = v2
                _ = fi26
            }
            for fi27 := 0; fi27 < 13; fi27++ {
                v28 := []int{-2147483648, (fi25 << ((65535) & 63))}
                var v29 uint8 = 15
                switch ("" + "hello") {
                case "a":
                    var v30 bool = true
                    _ = v30
                }
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        _ = fi25
    }
    var v31 int16 = 0
    v32 := v4[v2]
    var v33 int64 = 0
    v34 := GenS1{F0: 0, F1: "\n\t"}
    switch v3 {
    case 42:
        if false {
            fmt.Println(v3)
        } else {
            fmt.Println(v32, v22)
            v35 := GenS2{F0: (v33 << ((1) & 63))}
            v33 = v33
            _ = v35
        }
    case 128:
        delete(v1, v22)
    case 210:
        for fi36 := 0; fi36 < 2; fi36++ {
            for fi37 := 0; fi37 < 18; fi37++ {
                v22 = 902
                var v38 int = (v22 << ((1) & 63))
                v39 := []int{fi36}
                _ = fi37
                _ = v38
                _ = v39
            }
            _ = fi36
        }
    }
    if true {
        fmt.Println(v16)
        v40 := GenS3{F0: v32, F1: 0}
        _ = v40
    }
    v41 := GenS3{F0: v32, F1: (-(v22))}
    var v42 float64 = ((((0.1 * ((3.14159 * -2.25) * 0.1)) * (((float64(float32(0.1)) / float64(1e10)) / ((3.14159 - 1.5) / -2.25)) - float64(((float32((1e10 + (((0.0 * 1.5) * (((3.14159 / 3.14159) - (((1.5 - (-2.25 / 1.5)) - (1.5 - 0.0)) / ((0.1 - (0.0 + (float64((((0.0 - -2.25) - (3.14159 * (0.1 + float32((1.5 + 1.5))))) / 1.0)) * (1.5 / (((1.5 / 0.1) * (-2.25 * -2.25)) / 3.14159))))) + (0.0 / 0.1)))) * 0.0)) * 3.14159))) / 1.0) / 1.0)))) - (-2.25 + -2.25)) - -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v41
    _ = v42
}
