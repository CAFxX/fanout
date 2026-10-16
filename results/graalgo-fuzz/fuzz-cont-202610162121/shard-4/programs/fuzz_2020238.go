package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int }
type GenS2 struct { F0 int8; F1 uint8; F2 uint64 }
func genFn0(p0 int16) int8 {
    switch ("hello" + "hello") {
    case "foo bar":
        if (p0 != p0) {
            var v4 bool = false
            v5 := GenS1{F0: 500, F1: 600, F2: -678}
            _ = v4
            _ = v5
        } else {
            if (42 == 413) {
                var v6 uint64 = 4294967295
                v6 -= 42
                _ = v6
            }
        }
    case "\n\t":
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := []int{162, 237, fi7}
            _ = fi7
            _ = v8
        }
    default:
        fmt.Println(p0)
    }
    p0 = p0
    _ = p0
    return 0
}
func genFn1() uint {
    var v9 uint16 = 42
    v10 := GenS2{F0: -128, F1: 255, F2: 18446744073709551615}
    for fi11 := 0; fi11 < 5; fi11++ {
        v9 = 255
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 uint8 = 1
            var v14 int = fi11
            _ = fi12
            _ = v13
            _ = v14
        }
        fmt.Println(fi11, fi11)
        _ = fi11
    }
    if true {
        for fi15 := 0; fi15 < 2; fi15++ {
            v9 = (v9 + (v9 & 32767))
            if (1e10 > -2.25) {
                v16 := []int{(fi15 % ((fi15) | 1))}
                var v17 bool = (4294967295 != 0)
                fmt.Println(v17, v17, fi15)
                _ = v16
                _ = v17
            }
            _ = fi15
        }
    } else {
        var v18 int64 = 671
        _ = v18
    }
    var v19 int = 15
    _ = v10
    _ = v19
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint16 = 512
    var v2 string = ""
    var v3 int32 = 42
    fmt.Println(v1)
    if (127 > -1) {
        var v20 uint8 = 1
        fmt.Println(v2, v2)
        v21 := []int{int(v20)}
        _ = v20
        _ = v21
    }
    switch -767 {
    case 1:
        fmt.Println(v3)
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        switch ("a" + "foo bar") {
        case "foo bar":
            v3 -= (v3 % (((v3 ^ (v3 >> ((4294967295) & 63)))) | 1))
        case "\n\t":
            v3 = v3
        case "hello":
            v23 := GenS1{F0: 18446744073709551615, F1: 164, F2: int(v1)}
            _ = v23
        default:
            v24 := []int{}
            _ = v24
        }
        var v25 float32 = 3.14159
        var v26 uint = 18446744073709551615
        _ = fi22
        _ = v25
        _ = v26
    }
    var v27 bool = true
    var v28 bool = (32767 < 42)
    v2 = (v2 + ((v2 + v2) + "hello"))
    v29 := []int{}
    if ((!((v2 + v2) != v2)) || ((214 >= 104) || false)) {
        fmt.Println(v3, v28, v3)
        v30 := GenS2{F0: 0, F1: 128, F2: 889}
        _ = v30
    }
    v31 := GenS1{F0: genFn1(), F1: (^(v1)), F2: 42}
    v28 = (0 == int64(v1))
    _ = len(v29)
    for fi32 := 0; fi32 < 17; fi32++ {
        _ = len(v29)
        _ = fi32
    }
    var v33 int32 = v3
    if v28 {
        fmt.Println(v28, v1)
    }
    v1 -= (v1 * (v1 & 32767))
    if v27 {
        v33 = (-(v3))
        var v34 int32 = int32(v1)
        var v35 int = 42
        _ = v34
        _ = v35
    }
    v33 = (v3 + (v33 << ((561) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v33
}
