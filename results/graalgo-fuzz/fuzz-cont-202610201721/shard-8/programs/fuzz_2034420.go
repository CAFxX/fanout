package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 string }
func genFn0(p0 uint64, p1 int16, p2 bool) int8 {
    switch "" {
    case "\n\t":
        for fi4 := 0; fi4 < 16; fi4++ {
            var v5 uint64 = 18446744073709551615
            v6 := []int{32767}
            _ = fi4
            _ = v5
            _ = v6
        }
    case "hello":
        switch 4294967295 {
        case 127:
            if false {
                if true {
                    var v7 int8 = 100
                    v8 := GenS1{F0: 42}
                    _ = v7
                    _ = v8
                } else {
                    var v9 int16 = 2
                    var v10 int = 127
                    _ = v10
                    _ = v9
                }
            } else {
                v11 := GenS1{F0: -479}
                v12 := []int{0, int(p0)}
                _ = v11
                _ = v12
            }
        }
    }
    v13 := GenS1{F0: 127}
    switch 9223372036854775807 {
    case 32767:
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := GenS1{F0: fi14}
            var v16 int32 = 42
            _ = fi14
            _ = v15
            _ = v16
        }
    case -687:
        var v17 uint64 = p0
        _ = v17
    }
    if (((((float32((((0.1 - (1e10 + float64(((1e10 / float32(3.14159)) / ((-2.25 - 1.5) / 1e10))))) + (3.14159 + -2.25)) * 3.14159)) * float32(3.14159)) * (((0.1 - 0.1) + (-2.25 * ((((1e10 * 0.1) + 1.5) + 1e10) * float32(1e10)))) / 1e10)) - 0.1) + (1.5 - ((-2.25 * 0.1) / (float32(0.1) - -2.25)))) <= 1.5) {
        var v18 int = -2147483648
        fmt.Println(p0)
        _ = v18
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 uint8 = 2
        _ = fi19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 1
}

func main() {
    v1 := []uint32{816, 418, 0}
    var v2 string = ("a" + "")
    var v3 int = 9223372036854775807
    v3 = (v3 & v3)
    for fi21 := 0; fi21 < 3; fi21++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v22 := GenS2{F0: 555}
        for fi23 := 0; fi23 < 4; fi23++ {
            v24 := GenS3{F0: 2147483647}
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    v2 = v2
    v1 = append(v1, 0)
    if false {
        switch "" {
        case "\n\t":
            fmt.Println(v2)
        case "":
            fmt.Println(v3, v2)
        default:
            v2 = "foo bar"
        }
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
        for fi25 := 0; fi25 < 20; fi25++ {
            v1 = append(v1, 1)
            _ = fi25
        }
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch v2 {
        case "foo bar":
            v2 = v2
        case "\n\t":
            for fi27 := 0; fi27 < 11; fi27++ {
                v28 := GenS1{F0: 9223372036854775807}
                switch 127 {
                case -9223372036854775808:
                    v29 := GenS4{F0: 440}
                    _ = v29
                case 141:
                    switch -837 {
                    case 1:
                        var v30 int = -348
                        _ = v30
                    case 4294967295:
                        v31 := GenS5{F0: 2, F1: v2}
                        _ = v31
                    default:
                        v2 = (("" + ("a" + "\n\t")) + ((v2 + "a") + "hello"))
                    }
                    v2 = "a"
                    v2 = (("" + (v2 + v2)) + ("" + ""))
                }
                _ = fi27
                _ = v28
            }
        case "a":
            var v32 int64 = 594
            _ = v32
        }
        _ = fi26
    }
    fmt.Println(v2, v3)
    fmt.Println(v2, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
