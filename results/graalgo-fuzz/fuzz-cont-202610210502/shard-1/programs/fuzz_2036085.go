package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int }
func genFn0(p0 uint8, p1 int8) uint64 {
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            p1 = (p1 - p1)
            if false {
                var v9 bool = false
                _ = v9
            } else {
                p1 += int8(p0)
                var v10 float64 = (1e10 * 0.0)
                _ = v10
            }
            p0 += 127
            switch ("hello" + "foo bar") {
            case "a":
                var v11 int = -170
                _ = v11
            }
            _ = fi8
        }
        _ = fi7
    }
    var v12 int8 = (p1 | p1)
    var v13 uint16 = 525
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 18446744073709551615
}

func main() {
    var v1 string = ("hello" + ("\n\t" + "a"))
    var v2 int = 4294967295
    v3 := []string{v1, (v1 + v1)}
    v4 := make(map[string]string)
    var v5 int = (v2 & (v2 & (-(v2))))
    var v6 float32 = 0.1
    v2 -= 31
    if len(v3) > 0 {
        v3[0] = ((v1 + ((v1 + "a") + "\n\t")) + "a")
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 int64 = 9223372036854775807
        var v16 int16 = -912
        _ = fi14
        _ = v15
        _ = v16
    }
    if true {
        v17 := GenS1{F0: 42, F1: 1}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v17
    }
    var v18 int = 855
    var v19 int64 = -1
    fmt.Println(v2)
    v5 += (v5 & (-(v2)))
    v20 := GenS1{F0: 572, F1: 100}
    if (742 != -1) {
        for fi21 := 0; fi21 < 4; fi21++ {
            var v22 string = v1
            _ = fi21
            _ = v22
        }
    } else {
        if len(v3) > 0 {
            v3[0] = (v1 + v1)
        }
        v23 := GenS1{F0: 1, F1: 127}
        _ = v23
    }
    switch v2 {
    case 128:
        v24 := GenS1{F0: 4294967295, F1: 0}
        _ = v24
    }
    v4[v1] = v1
    if (319 > 2) {
        for fi25 := 0; fi25 < 3; fi25++ {
            switch v1 {
            case "a":
                v4[(v1 + ((("" + ("a" + v1)) + v1) + v1))] = ((((v1 + "\n\t") + ("hello" + v1)) + "foo bar") + ("foo bar" + "a"))
            case "foo bar":
                var v26 float32 = float32(3.14159)
                _ = v26
            case "":
                var v27 int32 = 69
                _ = v27
            default:
                v28 := GenS1{F0: 1, F1: 127}
                _ = v28
            }
            _ = fi25
        }
        var v29 bool = true
        _ = v29
    } else {
        v30 := GenS2{F0: v2}
        _ = v30
    }
    var v31 int8 = 127
    var v32 uint32 = 42
    v32 = (v32 + (v32 - 0))
    var v33 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
