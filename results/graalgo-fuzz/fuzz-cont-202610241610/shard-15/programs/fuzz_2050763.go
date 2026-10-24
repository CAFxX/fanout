package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint32; F2 bool }

func main() {
    v1 := []uint32{771}
    var v2 bool = true
    v3 := []int{1, 42}
    var v4 int32 = -432
    v5 := make(map[uint32]string)
    var v6 int64 = -9223372036854775808
    if v2 {
        v1 = append(v1, 0)
        fmt.Println(v4, v2)
        var v7 float64 = ((3.14159 / (1e10 + float64(0.0))) * 0.1)
        _ = v7
    }
    if ((3.14159 / -2.25) < (-2.25 - 3.14159)) {
        _ = len(v5)
        if (!true) {
            var v8 float32 = ((((3.14159 - (1e10 * float32(0.1))) + float32(((float64(0.1) * -2.25) + 1e10))) - 1e10) * (0.0 + 1.5))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v8
        }
        if true {
            if len(v3) > 0 {
                v3[0] = 9223372036854775807
            }
            switch "foo bar" {
            case "\n\t":
                if v2 {
                    v2 = (!(!v2))
                }
            }
        }
    } else {
        for fi9 := 0; fi9 < 2; fi9++ {
            fmt.Println(v2, v2)
            _ = fi9
        }
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        fmt.Println(v4)
        _ = fi10
    }
    delete(v5, 2147483648)
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 int16 = -1
        _ = fi11
        _ = v12
    }
    if true {
        switch 9223372036854775807 {
        case -2147483648:
            var v13 uint16 = 255
            _ = v13
        case 0:
            var v14 uint32 = 65535
            _ = v14
        }
        var v15 uint8 = 127
        _ = v15
    } else {
        for fi16 := 0; fi16 < 19; fi16++ {
            switch -600 {
            case 9223372036854775807:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case -2147483648:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                var v17 uint32 = 732
                _ = v17
            }
            _ = fi16
        }
        v1 = append(v1, 42)
    }
    v18 := GenS1{F0: 73, F1: v4, F2: 100}
    v19 := GenS2{F0: 69}
    v3 = append(v3, 2147483647)
    v6 += 42
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        _ = len(v3)
    case "a":
        fmt.Println(v4)
    case "\n\t":
        v6 += 1
    }
    var v20 float32 = (1e10 * 0.1)
    fmt.Println(v6, v20)
    switch "a" {
    case "a":
        var v21 float64 = (0.0 * -2.25)
        _ = v21
    case "hello":
        v22 := GenS3{F0: v6, F1: 5, F2: v2}
        _ = v22
    default:
        v2 = v2
    }
    v4 += 42
    _ = len(v5)
    var v23 uint64 = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
