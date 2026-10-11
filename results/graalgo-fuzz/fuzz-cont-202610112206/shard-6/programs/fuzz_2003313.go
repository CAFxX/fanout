package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
func genFn0() uint16 {
    if ((2 <= 306) || (127 <= 42)) {
        v5 := []int{4294967295}
        _ = v5
    }
    var v6 uint = 1
    if (true || (!(42 < 2))) {
        v6 = (^(v6))
    }
    _ = v6
    return 718
}

func main() {
    var v1 uint8 = 255
    var v2 string = "\n\t"
    v3 := []string{""}
    v4 := make(map[string]uint32)
    v7 := GenS1{F0: 397, F1: 4294967295}
    for fi8 := 0; fi8 < 1; fi8++ {
        if len(v3) > 0 {
            v3[0] = ((("a" + ("\n\t" + ((v2 + (v2 + ("hello" + (("foo bar" + v2) + v2)))) + v2))) + v2) + (((v2 + v2) + ("a" + v2)) + ""))
        }
        var v9 uint64 = 1
        _ = fi8
        _ = v9
    }
    switch "a" {
    case "foo bar":
        if (true || (!(((0.0 - 1.5) + 3.14159) < (1e10 - 3.14159)))) {
            var v10 uint = 2
            _ = v10
        } else {
            v4[v2] = 2147483648
            v11 := GenS1{F0: -1, F1: 4294967295}
            _ = v11
        }
    default:
        v12 := GenS1{F0: 0, F1: 4294967295}
        _ = v12
    }
    switch -325 {
    case 4294967295:
        if false {
            delete(v4, (v2 + ""))
        }
    }
    var v13 bool = false
    fmt.Println(v1, v13)
    for fi14 := 0; fi14 < 19; fi14++ {
        v15 := GenS1{F0: -2147483648, F1: 1}
        var v16 int16 = 32767
        v2 = v2
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 float64 = float64(1.5)
    var v18 uint = 461
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v19 := GenS1{F0: -1, F1: 4294967295}
    v1 -= v1
    switch 32767 {
    case 127:
        v17 += v17
    case 1:
        var v20 string = v2
        _ = v20
    default:
        v1 -= (v1 ^ (v1 & v1))
    }
    var v21 uint64 = 982
    v3 = append(v3, v2)
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    fmt.Println(v17, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v7
}
