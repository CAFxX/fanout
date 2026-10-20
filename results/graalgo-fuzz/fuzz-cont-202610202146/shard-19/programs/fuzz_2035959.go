package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int64; F2 uint8 }
func genFn0(p0 uint64, p1 string) int8 {
    fmt.Println(p1, p1)
    p1 = ""
    v7 := GenS1{F0: -1}
    v8 := GenS1{F0: -1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 127
}
func genFn1(p0 int) bool {
    var v9 int = p0
    v10 := GenS1{F0: 100}
    p0 = 9223372036854775807
    _ = p0
    _ = v10
    _ = v9
    return false
}

func main() {
    var v1 int16 = -1
    var v2 float64 = 1e10
    v3 := make(map[string]int)
    v4 := make(map[string]string)
    var v5 int32 = -144
    var v6 int64 = -954
    delete(v3, ("foo bar" + "hello"))
    var v11 uint64 = 4294967295
    v1 = (v1 * (v1 - -32768))
    v11 += v11
    v1 -= -322
    if ("hello" <= "foo bar") {
        fmt.Println(v2, v11, v2)
        var v12 int = -1
        var v13 int64 = v6
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 4; fi14++ {
            v15 := GenS1{F0: int8(v6)}
            _ = fi14
            _ = v15
        }
        var v16 uint32 = 1
        fmt.Println(v6, v1, v2)
        _ = v16
    }
    var v17 int = 179
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 float32 = float32((((v2 + float64(((1e10 + (((3.14159 * float32(v2)) + (0.1 + 0.0)) * (((3.14159 - ((((3.14159 / 1e10) - (0.1 / (float32((float64(1.5) * 1.5)) + 0.1))) - 1e10) / 0.1)) / 1.0) - -2.25))) + 3.14159))) - 1e10) / (v2 + (0.1 * (float64((float32(v2) * (((float32(0.1) - -2.25) + (3.14159 / 1.5)) + 3.14159))) * (v2 - (3.14159 + v2)))))))
        v20 := GenS1{F0: 127}
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 uint16 = 837
    switch 9223372036854775807 {
    case 9223372036854775807:
        switch ("" + "a") {
        case "a":
            v2 = v2
        case "":
            if (911 < 86) {
                var v22 float32 = (((-2.25 + 1e10) + -2.25) - 3.14159)
                for fi23 := 0; fi23 < 17; fi23++ {
                    fmt.Println(v17, v5, v5)
                    _ = fi23
                }
                _ = v22
            }
        }
    case 1:
        v11 -= (v11 + 0)
    case -828:
        if false {
            if (true || true) {
                var v24 bool = true
                v25 := GenS1{F0: 127}
                _ = v24
                _ = v25
            }
        } else {
            var v26 string = "\n\t"
            for fi27 := 0; fi27 < 16; fi27++ {
                v28 := GenS2{F0: 128, F1: 0, F2: 127}
                var v29 uint16 = v21
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = v26
        }
    }
    if true {
        var v30 int = (v17 >> ((1) & 63))
        _ = v30
    } else {
        v31 := GenS2{F0: 127, F1: 0, F2: 127}
        v1 = v1
        _ = v31
    }
    v32 := GenS1{F0: genFn0(v11, "foo bar")}
    if ((0.0 / 1e10) != (0.0 + (0.0 + ((0.1 - float32(((((0.0 / v2) + float64(0.0)) + (v2 * (v2 + (1.5 - -2.25)))) + (0.0 + 1e10)))) * float32((0.0 + (1.5 * float64(float32(float64((0.0 * 1.5))))))))))) {
        var v33 uint32 = 2
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
