package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int8; F2 int32 }
type GenS4 struct { F0 int16; F1 int32; F2 uint }
func genFn0() int8 {
    for fi4 := 0; fi4 < 17; fi4++ {
        for fi5 := 0; fi5 < 5; fi5++ {
            v6 := GenS1{F0: (188 <= 255), F1: 9223372036854775807}
            fmt.Println(fi4)
            _ = fi5
            _ = v6
        }
        v7 := []int{(^(fi4)), (fi4 >> ((2) & 63))}
        _ = fi4
        _ = v7
    }
    var v8 int16 = -32768
    for fi9 := 0; fi9 < 19; fi9++ {
        v8 = v8
        _ = fi9
    }
    _ = v8
    return 1
}
func genFn1(p0 string, p1 uint32) uint64 {
    if (454 != 0) {
        if (-1 <= 32767) {
            v10 := []int{-9223372036854775808}
            _ = v10
        }
        v11 := []int{-232}
        p0 = ("\n\t" + "")
        _ = v11
    }
    var v12 string = "hello"
    var v13 uint8 = 127
    for fi14 := 0; fi14 < 20; fi14++ {
        if (!((float32(1e10) / -2.25) != (0.0 / (0.1 - float32(1.5))))) {
            if (42 < 860) {
                var v15 uint64 = 42
                _ = v15
            }
        }
        if true {
            var v16 uint16 = 255
            _ = v16
        }
        _ = fi14
    }
    v17 := GenS1{F0: (false && true), F1: 42}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v17
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 0
    var v2 uint32 = 2147483647
    var v3 int64 = 1
    switch -468 {
    case -9223372036854775808:
        if true {
            v18 := GenS1{F0: true, F1: -981}
            fmt.Println(v1, v1, v1)
            fmt.Println(v3, v3)
            _ = v18
        } else {
            v19 := GenS1{F0: (false || (true && false)), F1: v3}
            _ = v19
        }
    case 32768:
        v20 := []int{128, -1}
        _ = v20
    default:
        v2 = 65535
    }
    var v21 int = 0
    v3 = int64(v21)
    var v22 uint16 = 361
    v23 := GenS2{F0: ("hello" + "\n\t")}
    var v24 uint64 = uint64(v22)
    switch 32767 {
    case 32768:
        v25 := GenS3{F0: 36, F1: 0, F2: 2147483647}
        _ = v25
    case 2147483647:
        var v26 int32 = 1
        _ = v26
    case 9223372036854775807:
        v27 := GenS2{F0: ("hello" + "foo bar")}
        _ = v27
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        v29 := GenS4{F0: 1000, F1: int32(v24), F2: 42}
        _ = fi28
        _ = v29
    }
    var v30 int32 = 42
    fmt.Println(v21, v3)
    v30 -= (v30 | v30)
    var v31 string = ""
    v32 := []int{(-(v21)), v21}
    var v33 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
}
