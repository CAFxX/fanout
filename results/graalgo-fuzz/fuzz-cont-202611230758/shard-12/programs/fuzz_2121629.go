package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 bool }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 uint16; F1 bool }
func genFn0(p0 uint8) bool {
    fmt.Println(p0)
    v7 := []int{572}
    fmt.Println(p0)
    var v8 bool = (3.14159 != 1.5)
    _ = p0
    _ = v7
    _ = v8
    return v8
}
func genFn1() int32 {
    v9 := []int{643, -2147483648}
    if true {
        var v10 int64 = -891
        var v11 float32 = 3.14159
        _ = v10
        _ = v11
    } else {
        v12 := GenS1{F0: false, F1: -128, F2: true}
        v13 := GenS1{F0: (42 == 255), F1: -38, F2: true}
        var v14 uint16 = 561
        _ = v12
        _ = v13
        _ = v14
    }
    _ = v9
    return 0
}

func main() {
    v1 := []string{("foo bar" + ("" + ("hello" + ""))), ""}
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]string)
    var v5 int64 = 9223372036854775807
    var v6 uint32 = 2
    _ = len(v2)
    v15 := GenS1{F0: true, F1: 100, F2: false}
    var v16 int16 = 1
    fmt.Println(v16)
    var v17 uint64 = 689
    v18 := GenS1{F0: false, F1: -128, F2: true}
    if false {
        switch 32768 {
        case -2147483648:
            if (!true) {
                v19 := GenS1{F0: true, F1: -128, F2: true}
                _ = v19
            } else {
                if (!(false || true)) {
                    for fi20 := 0; fi20 < 6; fi20++ {
                        var v21 float32 = ((-2.25 + -2.25) / 1.5)
                        _ = fi20
                        _ = v21
                    }
                    fmt.Println(v17, v16)
                    v22 := GenS2{F0: (("" + "hello") + "\n\t"), F1: (v16 <= v16)}
                    _ = v22
                }
                v23 := GenS3{F0: v5, F1: ("a" + "")}
                _ = v23
            }
        case 2147483647:
            var v24 int64 = (v5 * 9223372036854775807)
            _ = v24
        }
        var v25 uint8 = 128
        _ = v25
    } else {
        v26 := GenS4{F0: 1, F1: ((v5 | v5) == v5)}
        v6 = 2
        var v27 bool = ((v17 + v17) > v17)
        _ = v26
        _ = v27
    }
    v28 := v3[-2147483648]
    var v29 uint = 2
    fmt.Println(v16, v6, v28)
    fmt.Println(v16)
    v28 = (3.14159 > float32((((-2.25 * 1e10) * 0.1) + 1e10)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
