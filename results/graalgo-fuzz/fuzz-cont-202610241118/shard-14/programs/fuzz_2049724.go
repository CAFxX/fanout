package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 uint16 }
func genFn0(p0 int, p1 uint64) int8 {
    v7 := GenS1{F0: 20, F1: 100, F2: 65535}
    var v8 int8 = 0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}
func genFn1() int8 {
    switch 32767 {
    case 1:
        v9 := GenS1{F0: 1000, F1: genFn0(4294967295, 610), F2: 698}
        _ = v9
    case 32768:
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := []int{-9223372036854775808}
            switch "foo bar" {
            case "\n\t":
                var v12 uint64 = 42
                _ = v12
            default:
                var v13 uint32 = 25
                _ = v13
            }
            fmt.Println(fi10)
            _ = fi10
            _ = v11
        }
    default:
        var v14 uint64 = 762
        _ = v14
    }
    var v15 int16 = 42
    v16 := []int{32767, 4294967295, 4294967295}
    v15 = (v15 * (v15 ^ v15))
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 int32 = -2147483648
        var v19 bool = true
        if ((v18 + -997) <= (v18 % ((v18) | 1))) {
            v19 = ((!v19) || false)
            v16 = append(v16, 42)
            v16 = append(v16, -372)
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = v15
    _ = v16
    return genFn0(2147483647, 42)
}

func main() {
    var v1 int64 = -962
    v2 := make(map[string]uint32)
    var v3 int8 = 1
    var v4 int16 = 0
    var v5 bool = false
    v6 := make(map[uint32]string)
    switch ("" + "\n\t") {
    case "":
        v4 = (-(v4))
    }
    if v5 {
        var v20 bool = v5
        v1 = v1
        v4 = (^(v4))
        _ = v20
    }
    v21 := []int{0, 429, 2147483647}
    for fi22 := 0; fi22 < 4; fi22++ {
        if true {
            if len(v21) > 0 {
                v21[0] = (fi22 + fi22)
            }
            v5 = v5
        }
        v23 := GenS2{F0: 18446744073709551615}
        _ = fi22
        _ = v23
    }
    v5 = v5
    v3 += (-(v3))
    var v24 string = ("\n\t" + ("hello" + ("" + "a")))
    v25 := GenS1{F0: v4, F1: (v3 / ((int8(v4)) | 1)), F2: 773}
    var v26 uint32 = 42
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    v27 := GenS3{F0: v4, F1: 573}
    for fi28 := 0; fi28 < 10; fi28++ {
        fmt.Println(fi28, v3, v5)
        _ = fi28
    }
    if false {
        var v29 uint = 594
        v24 = "hello"
        _ = v29
    } else {
        fmt.Println(v26, v4, v4)
        _ = len(v2)
        var v30 int32 = 0
        _ = v30
    }
    var v31 int = 42
    var v32 uint16 = 2
    fmt.Println(v3, v24, v1)
    var v33 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
