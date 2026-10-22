package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int64) uint32 {
    p0 += (p0 / ((42) | 1))
    p0 += (p0 / (((p0 >> ((131) & 63))) | 1))
    switch 127 {
    case -714:
        p0 = 978
    case 4294967295:
        var v5 int8 = 0
        _ = v5
    default:
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 int8 = 42
            fmt.Println(v7, v7, v7)
            _ = fi6
            _ = v7
        }
    }
    _ = p0
    return 179
}
func genFn1(p0 bool) string {
    var v8 string = "a"
    v9 := []int{0}
    var v10 int16 = -415
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return (("foo bar" + (v8 + "hello")) + "")
}

func main() {
    var v1 uint16 = 0
    var v2 uint64 = 177
    v3 := []string{"\n\t"}
    var v4 uint8 = 252
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := GenS1{F0: true}
            switch "" {
            case "\n\t":
                v14 := GenS1{F0: (0 <= -839)}
                _ = v14
            case "hello":
                v1 -= v1
            default:
                if true {
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                } else {
                    v1 = 0
                    fmt.Println(fi12, fi12)
                }
            }
            v1 -= 756
            var v15 uint64 = 18446744073709551615
            v15 = 1
            var v16 int32 = 0
            _ = fi12
            _ = v13
            _ = v15
            _ = v16
        }
        _ = fi11
    }
    var v17 uint8 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v18 int8 = 42
    if (!((v1 >> ((255) & 63)) < v1)) {
        var v19 float32 = 3.14159
        v3 = append(v3, ("hello" + ("hello" + "a")))
        _ = v19
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v1 -= v1
    v20 := GenS1{F0: false}
    v4 = v17
    v1 = v1
    fmt.Println(v4, v18)
    var v21 float32 = 1.5
    if (0 == 0) {
        var v22 int32 = -2147483648
        _ = v22
    } else {
        var v23 uint32 = 42
        var v24 uint = 42
        var v25 float32 = v21
        _ = v23
        _ = v24
        _ = v25
    }
    v1 = v1
    v17 -= v17
    v17 = uint8(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
