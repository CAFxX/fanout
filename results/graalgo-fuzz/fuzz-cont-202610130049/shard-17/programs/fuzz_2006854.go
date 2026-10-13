package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 bool }
type GenS2 struct { F0 uint64; F1 int16; F2 bool }

func main() {
    var v1 string = ""
    v2 := make(map[int]bool)
    v3 := make(map[int]bool)
    v4 := []string{((v1 + (((("a" + (v1 + (v1 + (("hello" + "a") + "")))) + "a") + (v1 + v1)) + v1)) + "\n\t"), ("" + (v1 + v1)), ("a" + ("" + "a")), (v1 + v1)}
    v5 := []int{127, -752, 9223372036854775807}
    v6 := []string{v1, ((v1 + "\n\t") + v1), v1, v1}
    for fi7 := 0; fi7 < 15; fi7++ {
        if true {
            switch (-(fi7)) {
            case -2147483648:
                v8 := GenS1{F0: ((v1 + (("hello" + (v1 + v1)) + "\n\t")) + v1), F1: 42, F2: (-2.25 > ((1e10 * 3.14159) * 3.14159))}
                _ = v8
            case 2147483647:
                var v9 float32 = float32(0.1)
                _ = v9
            default:
                var v10 uint64 = 18446744073709551615
                _ = v10
            }
            v4 = append(v4, "foo bar")
        }
        fmt.Println(v1, fi7)
        var v11 string = "a"
        _ = fi7
        _ = v11
    }
    v1 = (v1 + (v1 + (v1 + (("a" + "hello") + v1))))
    for fi12 := 0; fi12 < 8; fi12++ {
        if (1 > 18446744073709551615) {
            for fi13 := 0; fi13 < 16; fi13++ {
                v2[(^(fi12))] = true
                v14 := GenS2{F0: 0, F1: 42, F2: true}
                _ = fi13
                _ = v14
            }
            if (("foo bar" >= "foo bar") && false) {
                v15 := GenS1{F0: (("foo bar" + v1) + (((v1 + v1) + v1) + (((v1 + v1) + v1) + v1))), F1: 1000, F2: (94 > 42)}
                _ = v15
            }
            if true {
                if (-765 != (fi12 % (((fi12 | fi12)) | 1))) {
                    var v16 int16 = -32768
                    var v17 int = -194
                    var v18 uint32 = 2
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    var v19 uint64 = 455
                    fmt.Println(fi12)
                    v1 = ("a" + "")
                    _ = v19
                }
                var v20 uint32 = 415
                var v21 uint64 = 1
                _ = v20
                _ = v21
            } else {
                var v22 uint32 = 65535
                _ = v22
            }
        }
        _ = fi12
    }
    v1 = v1
    var v23 float32 = 1.5
    v1 = ((v1 + v1) + v1)
    var v24 int16 = 558
    var v25 int64 = 42
    delete(v3, 128)
    v1 = ""
    if (false && true) {
        if (4294967295 == 1) {
            var v26 int64 = v25
            v5 = append(v5, 4294967295)
            _ = v26
        }
        v24 -= (v24 - -1)
    }
    if true {
        for fi27 := 0; fi27 < 5; fi27++ {
            switch fi27 {
            case -9223372036854775808:
                if (true || true) {
                    var v28 float64 = 1.5
                    v29 := v2[fi27]
                    _ = v28
                    _ = v29
                }
            }
            v23 = ((3.14159 * 3.14159) - 0.0)
            _ = fi27
        }
    }
    var v30 float32 = v23
    var v31 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
