package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 bool; F1 int16; F2 int8 }
type GenS3 struct { F0 uint; F1 bool }
func genFn0() string {
    var v7 int16 = -1
    switch 42 {
    case 32767:
        if true {
            v8 := []int{746}
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 int64 = 0
                v10 -= v10
                _ = fi9
                _ = v10
            }
            _ = v8
        } else {
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 float32 = 3.14159
                _ = fi11
                _ = v12
            }
            v13 := GenS1{F0: (187 >= 989), F1: "foo bar"}
            _ = v13
        }
    default:
        v14 := []int{0, -1, -1, 664}
        _ = v14
    }
    _ = v7
    return "hello"
}

func main() {
    var v1 uint8 = 42
    var v2 float32 = 1.5
    v3 := []string{"", "foo bar", "foo bar"}
    var v4 float32 = v2
    v5 := make(map[string]bool)
    v6 := []uint32{2, 42}
    var v15 int64 = -989
    var v16 uint16 = 255
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 uint16 = (v16 | v16)
        _ = fi17
        _ = v18
    }
    fmt.Println(v15)
    for fi19 := 0; fi19 < 4; fi19++ {
        switch "" {
        case "hello":
            var v20 float64 = 0.1
            _ = v20
        case "a":
            for fi21 := 0; fi21 < 16; fi21++ {
                switch 0 {
                case 32767:
                    var v22 bool = ((!false) || true)
                    v23 := GenS1{F0: (v22 && v22), F1: ("a" + "a")}
                    var v24 int32 = 1
                    _ = v22
                    _ = v23
                    _ = v24
                }
                if len(v3) > 0 {
                    v3[0] = "a"
                }
                var v25 int64 = v15
                _ = fi21
                _ = v25
            }
        default:
            var v26 uint64 = 2
            _ = v26
        }
        _ = fi19
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        v5["a"] = true
        _ = fi27
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        var v29 int64 = (v15 & (v15 >> ((65535) & 63)))
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v31 float64 = 0.1
        var v32 float32 = -2.25
        _ = fi30
        _ = v31
        _ = v32
    }
    switch 32768 {
    case 4294967295:
        for fi33 := 0; fi33 < 12; fi33++ {
            v15 += v15
            v2 = (0.0 + (0.0 / v4))
            _ = fi33
        }
    case 2147483647:
        var v34 int32 = -2147483648
        _ = v34
    case 9223372036854775807:
        v2 -= (0.0 * -2.25)
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        v36 := GenS2{F0: ((fi35 << ((859) & 63)) != (fi35 & 625)), F1: 0, F2: 100}
        var v37 int64 = (-(v15))
        var v38 int64 = v15
        _ = fi35
        _ = v36
        _ = v37
        _ = v38
    }
    v16 += (v16 / (((^(v16))) | 1))
    switch "foo bar" {
    case "hello":
        var v39 bool = true
        _ = v39
    case "foo bar":
        if true {
            fmt.Println(v16, v2)
        }
    case "a":
        var v40 int8 = 1
        _ = v40
    }
    switch -2147483648 {
    case 0:
        fmt.Println(v2, v1, v4)
    }
    v41 := GenS3{F0: 65535, F1: true}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
