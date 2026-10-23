package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
func genFn0(p0 uint32, p1 int8) int64 {
    for fi7 := 0; fi7 < 14; fi7++ {
        if false {
            p1 = p1
            v8 := []int{fi7, 127, 128, fi7}
            _ = v8
        }
        for fi9 := 0; fi9 < 16; fi9++ {
            p1 += (p1 + (p1 & int8(p0)))
            if false {
                v10 := GenS1{F0: 50, F1: ((("a" + "") + "hello") + ("foo bar" + ""))}
                _ = v10
            }
            _ = fi9
        }
        _ = fi7
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            fmt.Println(p1)
            switch "foo bar" {
            case "":
                var v13 int32 = -974
                _ = v13
            case "hello":
                p0 = (p0 << ((uint(p0)) & 63))
            case "a":
                var v14 uint8 = 42
                _ = v14
            default:
                v15 := GenS1{F0: 0, F1: "\n\t"}
                p1 -= (p1 / (((p1 - -128)) | 1))
                _ = v15
            }
            _ = fi12
        }
        _ = fi11
    }
    fmt.Println(p0, p0)
    var v16 int8 = 106
    _ = p0
    _ = p1
    _ = v16
    return int64(p0)
}
func genFn1() bool {
    var v17 bool = (true || false)
    var v18 string = ("\n\t" + "a")
    var v19 string = v18
    v20 := []int{-163, 128, 9223372036854775807, -1}
    _ = v17
    _ = v18
    _ = v19
    _ = v20
    return (-2147483648 >= -1)
}

func main() {
    v1 := []string{"a", "\n\t"}
    var v2 bool = true
    v3 := make(map[int]string)
    var v4 float32 = 3.14159
    var v5 int8 = -1
    v6 := make(map[string]string)
    switch 128 {
    case 773:
        var v21 int = -2147483648
        _ = v21
    default:
        for fi22 := 0; fi22 < 14; fi22++ {
            v5 += 100
            v3[fi22] = ("\n\t" + ("foo bar" + "a"))
            v23 := GenS1{F0: 636, F1: "hello"}
            _ = fi22
            _ = v23
        }
    }
    if (v5 != v5) {
        v24 := GenS1{F0: 507, F1: "\n\t"}
        _ = v24
    } else {
        fmt.Println(v5, v4)
    }
    if len(v1) > 0 {
        v1[0] = ("a" + "")
    }
    fmt.Println(v4)
    var v25 bool = (v2 || (42 != 140))
    var v26 bool = ((("hello" + ("foo bar" + ((("hello" + "a") + "a") + ("hello" + "a")))) == ("foo bar" + "")) && (float64(1e10) < 1e10))
    fmt.Println(v2, v26)
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
