package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 string }
func genFn0() uint64 {
    var v4 int16 = 32767
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 uint64 = 0
        v6 = v6
        _ = fi5
        _ = v6
    }
    v4 = (^(v4))
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(fi7, fi7)
        v4 += 32767
        switch "foo bar" {
        case "\n\t":
            v8 := []int{}
            _ = v8
        case "a":
            if (148 < 1) {
                v9 := []int{(fi7 + 435), fi7}
                var v10 uint = 529
                _ = v10
                _ = v9
            } else {
                if (false || (700 <= 376)) {
                    v11 := []int{fi7, (fi7 % ((fi7) | 1))}
                    var v12 int32 = int32(fi7)
                    _ = len(v11)
                    _ = v11
                    _ = v12
                }
            }
        case "hello":
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 uint16 = 2
                var v15 bool = false
                var v16 uint64 = 4294967295
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
        default:
            v17 := []int{fi7, fi7, fi7, (fi7 + -832)}
            _ = v17
        }
        _ = fi7
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        v4 -= 1
        switch (fi18 >> ((2) & 63)) {
        case 32768:
            v19 := []int{}
            _ = v19
        case 128:
            v20 := GenS1{F0: fi18, F1: 0, F2: "\n\t"}
            _ = v20
        case 42:
            v4 = (v4 + (v4 * -32768))
        default:
            for fi21 := 0; fi21 < 1; fi21++ {
                var v22 uint32 = 1
                _ = fi21
                _ = v22
            }
        }
        v23 := GenS1{F0: (fi18 - int(v4)), F1: 9223372036854775807, F2: "hello"}
        _ = fi18
        _ = v23
    }
    _ = v4
    return 248
}

func main() {
    var v1 int16 = 0
    v2 := make(map[int]uint32)
    var v3 int = 1
    var v24 uint = 18446744073709551615
    var v25 int = v3
    fmt.Println(v24)
    fmt.Println(v24, v3, v1)
    v2[v25] = 65535
    _ = len(v2)
    for fi26 := 0; fi26 < 8; fi26++ {
        fmt.Println(v25)
        v27 := GenS1{F0: (fi26 & 128), F1: 45, F2: "\n\t"}
        _ = fi26
        _ = v27
    }
    v3 = v3
    v1 = (v1 * v1)
    switch v25 {
    case 0:
        switch (v3 * (v3 - (v3 + v25))) {
        case 127:
            var v28 string = "\n\t"
            _ = v28
        case 4294967295:
            v25 -= v25
        }
    }
    v29 := GenS1{F0: v25, F1: 1, F2: "hello"}
    var v30 float64 = float64((0.1 / 3.14159))
    var v31 int8 = -128
    var v32 float32 = 1e10
    v30 = v30
    var v33 uint16 = 0
    var v34 float32 = ((-2.25 - v32) + v32)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}
