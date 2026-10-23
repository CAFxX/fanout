package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
func genFn0() uint16 {
    v7 := []int{1, -1, 2147483647, 439}
    v8 := GenS1{F0: 2147483647, F1: -1}
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]int)
    var v3 uint16 = 0
    v4 := make(map[int]int)
    v5 := make(map[string]string)
    var v6 int32 = v1
    switch 4294967295 {
    case -1:
        var v9 uint32 = 65535
        _ = v9
    default:
        switch "\n\t" {
        case "\n\t":
            v1 = v1
        default:
            var v10 float32 = ((((-2.25 - -2.25) - 1e10) - 0.0) * ((0.0 * (0.0 + (((float32((0.0 / 1.0)) - 1e10) / ((1e10 - float32(1.5)) - 3.14159)) - 1.5))) * 0.0))
            _ = v10
        }
    }
    switch "\n\t" {
    case "a":
        v6 -= v1
    case "\n\t":
        var v11 int8 = 0
        _ = v11
    }
    var v12 uint32 = 42
    if true {
        if ((true && (!true)) && (true && (!((127 == 139) || false)))) {
            if false {
                v13 := GenS1{F0: (-(v6)), F1: 230}
                _ = v13
            }
            v14 := v5["hello"]
            var v15 uint32 = v12
            _ = v14
            _ = v15
        } else {
            switch "foo bar" {
            case "foo bar":
                _ = len(v5)
            case "\n\t":
                v3 -= (v3 & 0)
            }
            v12 = (v12 - v12)
        }
    } else {
        if false {
            v6 = 42
            switch "foo bar" {
            case "a":
                for fi16 := 0; fi16 < 16; fi16++ {
                    v17 := GenS1{F0: v6, F1: 42}
                    v1 -= 42
                    v18 := GenS1{F0: 2147483647, F1: 1}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            default:
                if false {
                    v19 := GenS1{F0: v1, F1: -963}
                    _ = v19
                }
            }
            if (0.0 == ((float64(0.1) / 0.1) / float64(1.5))) {
                for fi20 := 0; fi20 < 17; fi20++ {
                    for fi21 := 0; fi21 < 17; fi21++ {
                        v12 = 2
                        _ = fi21
                    }
                    switch "" {
                    case "\n\t":
                        var v22 string = ("a" + "")
                        _ = v22
                    case "foo bar":
                        var v23 string = "\n\t"
                        _ = v23
                    case "a":
                        var v24 uint16 = v3
                        _ = v24
                    default:
                        delete(v2, (v12 + (v12 & v12)))
                    }
                    v25 := GenS1{F0: int32(fi20), F1: -933}
                    _ = fi20
                    _ = v25
                }
            } else {
                var v26 int64 = -9223372036854775808
                var v27 uint32 = v12
                _ = v26
                _ = v27
            }
        }
        v6 = v6
    }
    var v28 float32 = (0.1 - ((0.0 / ((3.14159 + 1.5) + float32(0.0))) / 0.1))
    if (0 < 852) {
        v29 := v5[""]
        v12 = (v12 >> ((0) & 63))
        _ = v29
    }
    switch "" {
    case "hello":
        fmt.Println(v3)
    default:
        var v30 float64 = 0.1
        _ = v30
    }
    var v31 int = -357
    var v32 uint64 = 2
    fmt.Println(v12, v3)
    for fi33 := 0; fi33 < 2; fi33++ {
        fmt.Println(v3, v28, v1)
        _ = fi33
    }
    fmt.Println(v32, v31)
    for fi34 := 0; fi34 < 4; fi34++ {
        v35 := []int{}
        var v36 int = (fi34 >> ((1) & 63))
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 int = (v31 / ((v31) | 1))
    var v38 uint64 = v32
    var v39 float64 = ((float64((v28 * (v28 * float32((0.0 / 0.1))))) - (float64(((1.5 - -2.25) + v28)) + (1e10 / float64(v28)))) - 0.0)
    var v40 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
