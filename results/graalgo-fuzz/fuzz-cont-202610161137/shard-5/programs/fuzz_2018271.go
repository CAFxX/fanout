package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int16; F1 string; F2 string }
type GenS3 struct { F0 uint16; F1 int16; F2 uint8 }
func genFn0(p0 uint16) uint64 {
    var v6 bool = false
    v7 := []int{0}
    var v8 int32 = 82
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"\n\t", "hello"}
    var v3 uint64 = 4294967295
    var v4 float32 = 1.5
    var v5 float32 = 1.5
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v9 := GenS1{F0: 2, F1: genFn0(427)}
    if false {
        fmt.Println(v4)
        v10 := GenS1{F0: 1, F1: (v3 << ((255) & 63))}
        v11 := GenS1{F0: 65535, F1: (^(v3))}
        _ = v10
        _ = v11
    }
    if (2147483647 < 567) {
        for fi12 := 0; fi12 < 12; fi12++ {
            fmt.Println(v5, fi12)
            var v13 float32 = v4
            _ = fi12
            _ = v13
        }
        var v14 uint16 = 756
        v15 := GenS1{F0: v14, F1: 1}
        _ = v14
        _ = v15
    } else {
        var v16 float32 = v5
        v3 = v3
        var v17 uint = 419
        _ = v16
        _ = v17
    }
    delete(v1, 128)
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 int8 = 15
        for fi20 := 0; fi20 < 1; fi20++ {
            v3 += 0
            v21 := GenS1{F0: 255, F1: v3}
            _ = fi20
            _ = v21
        }
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        _ = fi18
        _ = v19
    }
    var v22 int = 744
    for fi23 := 0; fi23 < 12; fi23++ {
        switch (-(fi23)) {
        case 32768:
            v22 = fi23
        default:
            if (-1 >= 4) {
                var v24 uint64 = (v3 + v3)
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 15; fi25++ {
            if (!true) {
                var v26 int64 = -48
                v27 := GenS1{F0: 63, F1: (v3 << ((0) & 63))}
                switch ("hello" + ("a" + "hello")) {
                case "hello":
                    v26 -= v26
                default:
                    v28 := GenS1{F0: 0, F1: v3}
                    _ = v28
                }
                _ = v26
                _ = v27
            } else {
                var v29 int16 = 32767
                _ = len(v1)
                _ = v29
            }
            fmt.Println(v22)
            var v30 string = ("a" + "foo bar")
            var v31 int32 = -2147483648
            v32 := GenS2{F0: 42, F1: (v30 + v30), F2: "\n\t"}
            _ = fi25
            _ = v30
            _ = v31
            _ = v32
        }
        _ = fi23
    }
    var v33 uint16 = 2
    v34 := GenS3{F0: (v33 - 324), F1: -704, F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v9
}
