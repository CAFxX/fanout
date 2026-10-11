package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int, p1 int64) uint16 {
    v5 := GenS1{F0: 311, F1: (p1 ^ p1), F2: int8(p0)}
    if (1 >= 173) {
        for fi6 := 0; fi6 < 14; fi6++ {
            v7 := []int{p0, p0, (p0 | (fi6 << ((255) & 63)))}
            if false {
                v8 := GenS1{F0: (p0 >> ((255) & 63)), F1: -1, F2: int8(p1)}
                v9 := GenS2{F0: -1}
                _ = v8
                _ = v9
            }
            p0 -= fi6
            _ = fi6
            _ = v7
        }
        var v10 uint16 = 0
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v5
    return 255
}
func genFn1() bool {
    switch 1 {
    case 0:
        if true {
            if true {
                var v11 uint64 = 4294967295
                _ = v11
            } else {
                var v12 uint = 18446744073709551615
                _ = v12
            }
            var v13 uint16 = genFn0(2147483647, -1)
            _ = v13
        } else {
            var v14 float32 = 0.0
            _ = v14
        }
    case 1:
        if false {
            var v15 uint16 = 42
            var v16 uint32 = 0
            v16 = 2147483648
            _ = v15
            _ = v16
        } else {
            var v17 float64 = 0.1
            v17 -= v17
            _ = v17
        }
    case -639:
        v18 := GenS1{F0: 32768, F1: 1, F2: 0}
        _ = v18
    default:
        for fi19 := 0; fi19 < 10; fi19++ {
            switch (fi19 - (fi19 + fi19)) {
            case 791:
                var v20 int64 = -9223372036854775808
                _ = v20
            case 42:
                v21 := []int{(fi19 << ((2) & 63)), 42}
                _ = v21
            default:
                var v22 int16 = -18
                _ = v22
            }
            var v23 string = "foo bar"
            _ = fi19
            _ = v23
        }
    }
    var v24 int32 = -1
    for fi25 := 0; fi25 < 3; fi25++ {
        v26 := []int{fi25, fi25, 128, 0}
        _ = fi25
        _ = v26
    }
    var v27 bool = false
    _ = v24
    _ = v27
    return (false || true)
}

func main() {
    var v1 float64 = 1.5
    var v2 int16 = 1000
    var v3 uint64 = 161
    v4 := make(map[string]string)
    v28 := []int{359}
    v29 := GenS1{F0: 32768, F1: -9223372036854775808, F2: -95}
    v30 := v4["a"]
    v3 += v3
    v30 = ((v30 + (v30 + "\n\t")) + "a")
    fmt.Println(v1)
    v3 -= (v3 * v3)
    switch (v30 + v30) {
    case "hello":
        v31 := GenS2{F0: (v2 / ((v2) | 1))}
        _ = v31
    case "foo bar":
        var v32 string = v30
        _ = v32
    case "":
        _ = len(v4)
    default:
        var v33 int = -9223372036854775808
        _ = v33
    }
    if true {
        var v34 int = 1
        _ = v34
    } else {
        v3 -= v3
        v28 = append(v28, -298)
        v3 = (v3 % ((18446744073709551615) | 1))
    }
    if (true || false) {
        if len(v28) >= 2 {
            v28 = v28[1:len(v28)]
        }
        v1 = 1.5
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
