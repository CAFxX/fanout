package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }

func main() {
    var v1 uint64 = 4294967295
    var v2 bool = false
    var v3 string = "\n\t"
    var v4 uint32 = 65535
    v5 := make(map[string]uint32)
    var v6 uint64 = v1
    for fi7 := 0; fi7 < 9; fi7++ {
        v1 += uint64(v4)
        var v8 float32 = 3.14159
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 1
    if v2 {
        v1 = 18446744073709551615
        fmt.Println(v6, v9)
        if v2 {
            for fi10 := 0; fi10 < 14; fi10++ {
                fmt.Println(v9)
                v11 := GenS1{F0: -463, F1: 65535}
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 int = (fi12 << ((1) & 63))
                for fi14 := 0; fi14 < 9; fi14++ {
                    v13 += fi14
                    v15 := []int{(^(fi14)), 204}
                    v9 += 42
                    if (true || v2) {
                        v16 := v5["foo bar"]
                        v4 = 1
                        v13 = (fi14 << ((255) & 63))
                        _ = v16
                    }
                    var v17 int8 = 48
                    _ = fi14
                    _ = v15
                    _ = v17
                }
                v2 = (!(float64(1.5) > (((-2.25 - (0.1 / (float64(float32(1e10)) - (-2.25 / (1e10 + (0.1 + ((1.5 - (0.0 + ((((((1e10 + 0.1) + ((0.0 / 0.1) / ((3.14159 + (-2.25 - 1.5)) / 1.0))) - 0.0) + 1e10) * ((-2.25 / float64(-2.25)) + -2.25)) / 1e10))) * float64(0.1)))))))) + ((0.1 - (3.14159 - (1.5 + -2.25))) - ((((float64(1e10) - 0.0) + 0.1) - (float64(((float32(1e10) * -2.25) / ((((-2.25 + (3.14159 / 1e10)) * ((float32(0.1) - 1.5) / 1.5)) * float32(((0.0 + (3.14159 + 3.14159)) / 0.1))) / 0.1))) * (float64(float32(1e10)) / 3.14159))) * float64(0.0)))) - (0.1 / 0.1))))
                _ = fi12
                _ = v13
            }
            var v18 bool = ((!(!v2)) || false)
            _ = v18
        } else {
            switch v3 {
            case "hello":
                v3 = "foo bar"
            case "":
                v6 = 4294967295
            case "\n\t":
                v5[v3] = 4294967295
            }
        }
    }
    switch 2147483647 {
    case -2147483648:
        for fi19 := 0; fi19 < 19; fi19++ {
            v5["hello"] = 65535
            v20 := GenS1{F0: 42, F1: 554}
            switch v3 {
            case "\n\t":
                var v21 uint32 = v4
                _ = v21
            }
            _ = fi19
            _ = v20
        }
    }
    if v2 {
        var v22 int64 = -1
        switch -1 {
        case -9223372036854775808:
            v2 = false
        default:
            for fi23 := 0; fi23 < 8; fi23++ {
                v5[v3] = (v4 - v4)
                v2 = (true && v2)
                _ = fi23
            }
        }
        _ = v22
    } else {
        delete(v5, v3)
    }
    v24 := GenS1{F0: -659, F1: 65535}
    _ = len(v5)
    v25 := []int{0, 403, 9223372036854775807}
    v25 = append(v25, 9223372036854775807)
    var v26 bool = true
    v27 := v5[(("foo bar" + (v3 + ("a" + v3))) + v3)]
    var v28 int = 4294967295
    v29 := GenS1{F0: 0, F1: 1}
    v30 := v5["hello"]
    var v31 int64 = 975
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
