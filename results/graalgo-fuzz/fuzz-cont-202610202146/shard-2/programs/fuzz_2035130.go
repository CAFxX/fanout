package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint64, p1 uint64) bool {
    fmt.Println(p1, p0)
    v4 := GenS1{F0: true}
    p1 += p1
    var v5 uint8 = 127
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (!((^(v5)) < (v5 & (v5 << ((1) & 63)))))
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[int]string)
    var v3 int16 = -433
    var v6 uint8 = 42
    var v7 uint32 = 610
    var v8 bool = ((v7 >> ((239) & 63)) <= (v7 * 4294967295))
    switch 0 {
    case 0:
        v9 := GenS1{F0: v8}
        _ = v9
    case -9223372036854775808:
        var v10 uint = 18446744073709551615
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        v6 += 255
        v12 := GenS1{F0: ((!(v6 != v6)) && (false && false))}
        switch ((("foo bar" + ("hello" + (("\n\t" + "hello") + "hello"))) + "hello") + "foo bar") {
        case "a":
            v7 += uint32(v1)
        case "\n\t":
            for fi13 := 0; fi13 < 6; fi13++ {
                switch ("foo bar" + "\n\t") {
                case "a":
                    if (((((((float32((-2.25 * float64(0.1))) + 1e10) - float32(0.1)) + (1e10 / 1.0)) * 1.5) - 1.5) + 3.14159) > 1.5) {
                        v14 := GenS1{F0: (65535 != (^(v1)))}
                        _ = v14
                    }
                }
                v6 = (v6 * 2)
                var v15 uint32 = 2147483648
                _ = fi13
                _ = v15
            }
        case "":
            var v16 int = (fi11 / (((fi11 | (fi11 - fi11))) | 1))
            _ = v16
        default:
            var v17 uint16 = v1
            _ = v17
        }
        _ = fi11
        _ = v12
    }
    var v18 string = "a"
    fmt.Println(v7, v6)
    v7 = 4294967295
    v19 := GenS1{F0: v8}
    if (127 >= 128) {
        v20 := GenS2{F0: (v7 ^ v7), F1: 112}
        if v8 {
            var v21 uint32 = (^(v7))
            _ = v21
        } else {
            switch 9223372036854775807 {
            case 127:
                v8 = ((("" + ("foo bar" + "\n\t")) + v18) >= v18)
            }
            var v22 bool = (653 != 0)
            _ = v22
        }
        _ = v20
    } else {
        var v23 bool = v8
        var v24 int16 = (v3 & (v3 | v3))
        _ = v23
        _ = v24
    }
    fmt.Println(v18)
    v25 := GenS2{F0: 0, F1: -108}
    v26 := []int{108, 9223372036854775807, -2147483648}
    var v27 uint = 42
    for fi28 := 0; fi28 < 3; fi28++ {
        fmt.Println(v27)
        for fi29 := 0; fi29 < 1; fi29++ {
            if (!((-2147483648 == int32(v1)) || ((1e10 >= (float32(1e10) + (0.0 + 0.1))) && (((v18 <= v18) && v8) || ((v8 && v8) || v8))))) {
                var v30 int = fi28
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
                _ = v30
            } else {
                var v31 int16 = v3
                v31 = v3
                _ = v31
            }
            v32 := GenS2{F0: (v7 * uint32(v6)), F1: 127}
            _ = fi29
            _ = v32
        }
        var v33 int8 = 100
        _ = fi28
        _ = v33
    }
    switch "foo bar" {
    case "\n\t":
        if ((-2.25 / 1e10) >= 1e10) {
            var v34 float64 = 3.14159
            v1 = v1
            _ = v34
        } else {
            switch v18 {
            case "a":
                v35 := GenS2{F0: (^(v7)), F1: 0}
                _ = v35
            default:
                if (true && (-121 != -1)) {
                    var v36 float32 = float32(float64(0.1))
                    _ = v36
                }
            }
            var v37 uint64 = 4294967295
            v2[4294967295] = ("\n\t" + v18)
            _ = v37
        }
    case "a":
        fmt.Println(v7, v8)
    case "hello":
        var v38 int16 = v3
        _ = v38
    default:
        switch v18 {
        case "foo bar":
            var v39 int = 0
            _ = v39
        default:
            var v40 int32 = -2147483648
            _ = v40
        }
    }
    _ = len(v26)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
