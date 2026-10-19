package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int16, p1 int8) uint8 {
    v6 := []int{}
    for fi7 := 0; fi7 < 13; fi7++ {
        p1 = p1
        _ = fi7
    }
    if (42 > 65535) {
        fmt.Println(p1, p1)
    } else {
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = v6
    return 12
}

func main() {
    var v1 float64 = (0.1 + 1.5)
    var v2 int = 127
    v3 := make(map[string]string)
    var v4 float32 = 0.1
    var v5 int8 = 0
    v8 := v3["foo bar"]
    fmt.Println(v1, v2)
    switch (-(v2)) {
    case 0:
        delete(v3, v8)
    case 32768:
        var v9 uint = 18446744073709551615
        _ = v9
    default:
        switch v2 {
        case 128:
            v1 = 0.0
        }
    }
    switch (v2 % ((v2) | 1)) {
    case 1:
        var v10 uint16 = 234
        _ = v10
    case -1:
        fmt.Println(v1, v2)
    default:
        v11 := GenS1{F0: 127}
        _ = v11
    }
    if (-937 < 1) {
        switch ("foo bar" + v8) {
        case "":
            v12 := []int{(v2 / ((127) | 1)), v2, v2}
            _ = v12
        case "a":
            for fi13 := 0; fi13 < 6; fi13++ {
                for fi14 := 0; fi14 < 13; fi14++ {
                    if (!(1 < 71)) {
                        v2 -= fi14
                        var v15 int16 = -298
                        _ = v15
                    }
                    _ = fi14
                }
                v4 += (float32((3.14159 / (1e10 + v1))) / (v4 / -2.25))
                switch v8 {
                case "":
                    v16 := []int{fi13, (-(fi13))}
                    _ = v16
                default:
                    v3[v8] = v8
                }
                _ = fi13
            }
        }
        v17 := v3[(("foo bar" + v8) + v8)]
        if false {
            v17 = ((v8 + v8) + (v8 + "\n\t"))
            if true {
                fmt.Println(v17, v17, v5)
                for fi18 := 0; fi18 < 16; fi18++ {
                    fmt.Println(v1, v8)
                    if (1 != 964) {
                        fmt.Println(v17)
                    }
                    _ = fi18
                }
            } else {
                if false {
                    switch "hello" {
                    case "\n\t":
                        var v19 int32 = 1
                        _ = v19
                    case "foo bar":
                        var v20 bool = (!true)
                        _ = v20
                    default:
                        var v21 int = v2
                        _ = v21
                    }
                }
                v22 := GenS1{F0: 1}
                _ = v22
            }
            v4 -= (v4 + (v4 + 0.0))
        } else {
            var v23 float32 = v4
            _ = v23
        }
        _ = v17
    } else {
        var v24 int64 = 9223372036854775807
        v25 := []int{v2, v2, -238, (-(v2))}
        var v26 uint8 = genFn0(-255, (v5 >> ((18446744073709551615) & 63)))
        _ = v24
        _ = v25
        _ = v26
    }
    v27 := v3[v8]
    switch "" {
    case "\n\t":
        v28 := GenS2{F0: 852}
        _ = v28
    case "hello":
        var v29 uint8 = 127
        _ = v29
    case "a":
        v30 := GenS2{F0: 2}
        _ = v30
    default:
        v1 -= (v1 + 0.0)
    }
    var v31 int = (v2 - v2)
    var v32 uint64 = 1
    var v33 int16 = -404
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v8
}
