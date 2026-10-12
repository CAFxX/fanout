package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
func genFn0(p0 int8, p1 bool) int16 {
    switch 735 {
    case 9223372036854775807:
        v5 := []int{32767, 1, 1}
        _ = v5
    default:
        for fi6 := 0; fi6 < 1; fi6++ {
            p0 = 127
            p1 = p1
            p0 += p0
            _ = fi6
        }
    }
    fmt.Println(p0, p0)
    p1 = p1
    _ = p0
    _ = p1
    return -802
}

func main() {
    var v1 uint16 = 278
    v2 := []uint32{1, 682, 0}
    var v3 bool = (!(-1 <= 42))
    var v4 int = 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (398 != 158) {
        for fi7 := 0; fi7 < 6; fi7++ {
            v4 -= v4
            var v8 float64 = 0.0
            _ = fi7
            _ = v8
        }
        _ = len(v2)
        fmt.Println(v1, v4, v1)
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v9 uint32 = 2147483648
        v10 := GenS1{F0: -465}
        _ = v10
        _ = v9
    }
    var v11 bool = (((float64(-2.25) / (3.14159 - (3.14159 / 3.14159))) + 1e10) == 3.14159)
    var v12 float32 = 0.1
    _ = len(v2)
    var v13 float32 = v12
    var v14 string = "hello"
    v15 := GenS2{F0: 1}
    var v16 uint16 = 0
    switch 4294967295 {
    case 47:
        v1 = v1
    case 4294967295:
        switch v14 {
        case "":
            v1 = v16
        case "hello":
            if len(v2) > 0 {
                v2[0] = 65535
            }
        case "a":
            var v17 float32 = float32((3.14159 - (0.0 * float64(3.14159))))
            _ = v17
        default:
            fmt.Println(v11)
        }
    case 9223372036854775807:
        if ((false && (((!((!true) || (!(v14 <= (((("\n\t" + ((v14 + (v14 + "hello")) + v14)) + "a") + (v14 + (("a" + v14) + v14))) + "\n\t"))))) || true) || (79 < 75))) || true) {
            var v18 bool = (v11 || v3)
            v13 = 0.0
            fmt.Println(v13)
            _ = v18
        }
    default:
        for fi19 := 0; fi19 < 1; fi19++ {
            v13 = float32(0.1)
            if (v3 || (!v11)) {
                var v20 int8 = 1
                var v21 uint16 = 42
                v1 = (v16 - v16)
                _ = v20
                _ = v21
            }
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
            _ = fi19
        }
    }
    var v22 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
