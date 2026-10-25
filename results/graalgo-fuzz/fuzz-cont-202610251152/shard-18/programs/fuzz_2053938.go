package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint }
type GenS2 struct { F0 uint16; F1 bool; F2 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int16, p1 uint, p2 int8) bool {
    v5 := []int{}
    var v6 uint32 = 478
    v7 := GenS1{F0: v6, F1: true, F2: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (9223372036854775807 == -9223372036854775808)
}
func genFn1() int8 {
    v8 := []int{428}
    if (-32768 != -32768) {
        _ = len(v8)
        if false {
            var v9 uint32 = 42
            var v10 uint32 = (v9 / (((^(v9))) | 1))
            if len(v8) > 0 {
                v8[0] = 32767
            }
            _ = v10
            _ = v9
        }
        switch 127 {
        case 128:
            var v11 uint64 = 42
            _ = v11
        case 0:
            v12 := GenS2{F0: 1, F1: true, F2: false}
            _ = v12
        case 2147483647:
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
    } else {
        var v13 bool = (!true)
        v13 = v13
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 uint = 4294967295
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    v16 := GenS1{F0: 240, F1: (255 >= 255), F2: 4294967295}
    var v17 int16 = 1000
    _ = v16
    _ = v17
    _ = v8
    return 42
}

func main() {
    var v1 string = "hello"
    var v2 uint16 = 251
    v3 := []int{2147483647, 81, -213, -683}
    var v4 uint8 = 2
    switch ("a" + v1) {
    case "\n\t":
        switch v1 {
        case "foo bar":
            fmt.Println(v2, v4, v2)
        case "":
            v18 := GenS2{F0: v2, F1: false, F2: (42 <= 0)}
            _ = v18
        case "\n\t":
            fmt.Println(v1, v2, v2)
        default:
            v19 := GenS2{F0: (v2 | (v2 | v2)), F1: true, F2: false}
            _ = v19
        }
    default:
        v20 := GenS2{F0: v2, F1: (true && false), F2: true}
        _ = v20
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 int = fi21
        _ = fi21
        _ = v22
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) > 0 {
        v3[0] = -954
    }
    v1 = ("a" + "hello")
    for fi23 := 0; fi23 < 10; fi23++ {
        v1 = (v1 + (v1 + v1))
        if (v4 < v4) {
            if len(v3) > 0 {
                v3[0] = -2147483648
            }
            fmt.Println(v2)
            v3 = append(v3, -9223372036854775808)
        }
        _ = fi23
    }
    var v24 float64 = (3.14159 / 1.0)
    for fi25 := 0; fi25 < 3; fi25++ {
        v26 := GenS3{F0: (false && true)}
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        if ((fi27 | (^(fi27))) > fi27) {
            var v28 int64 = 9223372036854775807
            if ((((float32(v24) * 1.5) * ((3.14159 - (float32(1e10) * (3.14159 * -2.25))) * (1.5 / 1.0))) / 3.14159) > (1.5 - 0.1)) {
                var v29 uint16 = v2
                _ = len(v3)
                for fi30 := 0; fi30 < 5; fi30++ {
                    var v31 float64 = v24
                    _ = fi30
                    _ = v31
                }
                _ = v29
            }
            var v32 int8 = 100
            _ = v28
            _ = v32
        } else {
            var v33 string = v1
            fmt.Println(v2)
            fmt.Println(v33, v33)
            _ = v33
        }
        _ = fi27
    }
    v34 := GenS4{F0: (v2 - v2)}
    switch "" {
    case "\n\t":
        v4 -= 78
    case "":
        for fi35 := 0; fi35 < 19; fi35++ {
            v1 = "foo bar"
            _ = fi35
        }
    default:
        v36 := GenS2{F0: 0, F1: true, F2: (!false)}
        _ = v36
    }
    var v37 int32 = -373
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v37
    _ = v4
}
