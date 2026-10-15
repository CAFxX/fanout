package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }
func genFn0(p0 bool) uint64 {
    var v4 int64 = 1
    switch ("" + "foo bar") {
    case "hello":
        v5 := GenS1{F0: -1}
        _ = v5
    case "":
        v6 := GenS2{F0: -9223372036854775808}
        _ = v6
    default:
        var v7 bool = (((-2147483648 != 1) || false) || false)
        _ = v7
    }
    var v8 uint8 = 0
    v9 := []int{}
    _ = p0
    _ = v4
    _ = v8
    _ = v9
    return 4294967295
}
func genFn1(p0 uint32, p1 uint64, p2 int8) uint32 {
    p0 = 243
    p1 += genFn0(false)
    var v10 uint8 = 128
    var v11 int16 = 32767
    var v12 float64 = float64((((float32(-2.25) - -2.25) - ((((((0.1 - -2.25) / 0.1) / 1.0) + (((-2.25 + (((float32(0.1) + ((1e10 - (float32(-2.25) - (1e10 / 1.0))) + (1e10 / 1.5))) / 1.0) - (3.14159 - float32(1e10)))) * 0.1) - 1e10)) * 0.1) / float32(((3.14159 * -2.25) / float64(((3.14159 - (3.14159 * 1e10)) * (0.1 + float32(0.1)))))))) + 1e10))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return (p0 << ((18446744073709551615) & 63))
}

func main() {
    var v1 float64 = ((1.5 * -2.25) - ((((3.14159 - (0.0 - 0.1)) + 3.14159) * (0.1 - 3.14159)) / (float64(0.1) - 3.14159)))
    var v2 uint32 = 4294967295
    var v3 string = "a"
    var v13 float64 = v1
    if true {
        v3 = v3
    } else {
        fmt.Println(v2, v1, v13)
    }
    v14 := []int{}
    _ = len(v14)
    if len(v14) > 0 {
        v14[0] = -2147483648
    }
    fmt.Println(v13, v13, v2)
    v13 = v13
    switch "\n\t" {
    case "":
        switch -199 {
        case 549:
            fmt.Println(v3, v3, v2)
        }
    case "foo bar":
        var v15 float64 = 0.1
        _ = v15
    case "hello":
        v16 := GenS2{F0: 128}
        _ = v16
    }
    v13 -= float64((float32(v1) / -2.25))
    v13 += float64(1.5)
    switch 332 {
    case 32768:
        v1 = v1
    case 4294967295:
        var v17 uint8 = 194
        _ = v17
    }
    if false {
        if false {
            fmt.Println(v3)
            var v18 uint32 = 12
            var v19 uint = 0
            _ = v18
            _ = v19
        }
        fmt.Println(v13)
    }
    v20 := GenS3{F0: 1}
    if (false && true) {
        var v21 string = v3
        for fi22 := 0; fi22 < 11; fi22++ {
            v2 = v2
            v1 = 0.0
            _ = fi22
        }
        v21 = v3
        _ = v21
    } else {
        if len(v14) > 0 {
            v14[0] = 42
        }
        var v23 int = 32768
        _ = v23
    }
    v13 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
}
