package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 string }
func genFn0() uint64 {
    var v4 int16 = 42
    v5 := []int{32768}
    v6 := GenS1{F0: -49, F1: false, F2: 106}
    _ = len(v5)
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 2147483647
    var v3 int8 = 42
    fmt.Println(v3)
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS2{F0: false}
        fmt.Println(fi7, fi7)
        v9 := GenS1{F0: (v2 - v2), F1: true, F2: 583}
        _ = fi7
        _ = v8
        _ = v9
    }
    v3 = 10
    v2 = (v2 * (v2 >> ((82) & 63)))
    for fi10 := 0; fi10 < 1; fi10++ {
        delete(v1, "")
        v11 := GenS3{F0: false, F1: fi10}
        v12 := v1[("foo bar" + "\n\t")]
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int = -2147483648
    var v14 int16 = 1000
    switch v13 {
    case 0:
        var v15 uint8 = 1
        _ = v15
    case -316:
        var v16 bool = false
        _ = v16
    case 128:
        var v17 bool = false
        _ = v17
    }
    _ = len(v1)
    v18 := []int{v13, v13, v13, int(v2)}
    v19 := GenS4{F0: 2, F1: genFn0()}
    var v20 uint8 = 149
    if (-32768 == (v14 / (((v14 >> ((4294967295) & 63))) | 1))) {
        if false {
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        } else {
            if false {
                v21 := GenS1{F0: (v2 << ((558) & 63)), F1: (3.14159 > 0.1), F2: 65535}
                var v22 uint = 18446744073709551615
                _ = v21
                _ = v22
            }
        }
        v1[""] = (v13 & v13)
    } else {
        v13 += (-(v13))
        var v23 float32 = (0.1 + (0.1 / 1.0))
        for fi24 := 0; fi24 < 13; fi24++ {
            v14 += (v14 % (((v14 - int16(v2))) | 1))
            _ = fi24
        }
        _ = v23
    }
    var v25 uint16 = 167
    for fi26 := 0; fi26 < 1; fi26++ {
        v18 = append(v18, 127)
        _ = fi26
    }
    v27 := GenS5{F0: ""}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v27
    _ = v3
}
