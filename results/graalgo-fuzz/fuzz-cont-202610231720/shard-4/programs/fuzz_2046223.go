package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int16 }
func genFn0(p0 int16, p1 int, p2 uint8) int8 {
    for fi4 := 0; fi4 < 12; fi4++ {
        fmt.Println(fi4, p2)
        _ = fi4
    }
    fmt.Println(p0)
    var v5 int32 = 1
    var v6 float32 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int64, p1 bool) uint32 {
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            if (1000 >= 32767) {
                v9 := GenS1{F0: 42}
                fmt.Println(p1, fi7)
                var v10 uint32 = 2147483648
                _ = v10
                _ = v9
            } else {
                var v11 uint64 = 974
                var v12 float32 = 1e10
                _ = v11
                _ = v12
            }
            v13 := GenS2{F0: (p0 & p0), F1: -932}
            var v14 int16 = 0
            _ = fi8
            _ = v13
            _ = v14
        }
        v15 := GenS1{F0: genFn0(0, fi7, 0)}
        fmt.Println(p1, p0, p1)
        _ = fi7
        _ = v15
    }
    p0 = 192
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    var v3 string = "foo bar"
    var v16 int8 = 1
    fmt.Println(v16, v3)
    fmt.Println(v3, v3)
    v16 += v16
    var v17 uint16 = 255
    v18 := GenS2{F0: -521, F1: -1}
    v17 = v17
    for fi19 := 0; fi19 < 19; fi19++ {
        v17 += 430
        if true {
            v20 := GenS1{F0: v16}
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 float64 = (0.1 * 1.5)
                _ = fi21
                _ = v22
            }
            var v23 uint32 = 65535
            _ = v20
            _ = v23
        }
        switch ((v3 + (v3 + ((("foo bar" + ("hello" + (v3 + ("\n\t" + ((v3 + ("" + v3)) + v3))))) + "\n\t") + v3))) + (v3 + v3)) {
        case "a":
            v17 += v17
        case "":
            switch (v3 + "hello") {
            case "":
                if ((0.0 / 3.14159) < (0.0 + (1e10 * 3.14159))) {
                    v24 := []int{(fi19 << ((2) & 63))}
                    v25 := v1[((("a" + "") + (v3 + v3)) + (v3 + v3))]
                    v1["hello"] = genFn1(9223372036854775807, false)
                    _ = v24
                    _ = v25
                }
            case "foo bar":
                delete(v2, v3)
            }
        default:
            var v26 bool = false
            _ = v26
        }
        _ = fi19
    }
    if true {
        v27 := GenS1{F0: genFn0(77, 9223372036854775807, 255)}
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
