package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 string, p1 uint16) int16 {
    var v4 int = 9223372036854775807
    v5 := GenS1{F0: -2147483648, F1: 4294967295, F2: 80}
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint = 42
        var v8 float64 = 0.1
        for fi9 := 0; fi9 < 20; fi9++ {
            if false {
                var v10 string = p0
                var v11 uint = 1
                fmt.Println(fi9)
                var v12 uint = (^(v7))
                _ = v10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v13 int32 = 2147483647
    var v14 int32 = (v13 & 0)
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v4
    _ = v5
    return 0
}

func main() {
    v1 := []string{"a", ("a" + ("foo bar" + "hello"))}
    v2 := make(map[int]string)
    v3 := []string{}
    v1 = append(v1, ("a" + ((("" + "a") + "") + "")))
    var v15 uint16 = 1
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int64 = 1
            var v19 int64 = (^(v18))
            _ = fi17
            _ = v18
            _ = v19
        }
        if true {
            fmt.Println(fi16)
            var v20 float32 = 1e10
            _ = v20
        } else {
            fmt.Println(v15, v15)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v21 := GenS1{F0: 571, F1: 2, F2: 168}
            _ = v21
        }
        delete(v2, -1)
        _ = fi16
    }
    var v22 bool = true
    var v23 string = "foo bar"
    for fi24 := 0; fi24 < 11; fi24++ {
        v22 = v22
        switch ("hello" + v23) {
        case "a":
            v25 := GenS1{F0: 0, F1: 1, F2: 16}
            _ = v25
        default:
            switch v23 {
            case "\n\t":
                if true {
                    var v26 string = "foo bar"
                    _ = v26
                }
            case "hello":
                for fi27 := 0; fi27 < 11; fi27++ {
                    _ = len(v2)
                    fmt.Println(fi24)
                    _ = fi27
                }
            case "a":
                var v28 float32 = (((1.5 + float32(((((0.0 + 1e10) * 1e10) / (((float64(-2.25) * ((float64(0.1) + ((1e10 - (float64((0.0 * (0.0 / -2.25))) * (0.1 + ((1.5 * 3.14159) / (-2.25 - 0.0))))) + 3.14159)) / 1e10)) - 1.5) + float64(0.1))) + float64(3.14159)))) / (((((0.0 + 3.14159) - 0.1) + float32(3.14159)) - float32(1.5)) + 3.14159)) - float32(1e10))
                _ = v28
            default:
                v23 = (v23 + v23)
            }
        }
        _ = fi24
    }
    v29 := GenS2{F0: (^(v15))}
    v23 = (((v23 + "foo bar") + "") + v23)
    var v30 float32 = 1e10
    v31 := GenS2{F0: (v15 >> ((165) & 63))}
    switch -137 {
    case 127:
        var v32 int64 = 1
        _ = v32
    case -1:
        fmt.Println(v30, v30, v23)
    case 965:
        v15 = (^(v15))
    }
    var v33 float64 = -2.25
    var v34 int32 = 42
    var v35 uint64 = 415
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v35
}
