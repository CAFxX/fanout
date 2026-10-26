package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint16 }
func genFn0(p0 int64) int16 {
    var v4 int32 = -170
    switch 42 {
    case 9223372036854775807:
        v4 = v4
    default:
        v5 := []int{4294967295, 32767}
        _ = v5
    }
    v6 := []int{475}
    _ = p0
    _ = v4
    _ = v6
    return -32768
}
func genFn1() uint16 {
    v7 := GenS1{F0: 32768, F1: genFn0(-608), F2: 234}
    var v8 string = "\n\t"
    var v9 int32 = -2147483648
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 float32 = 1e10
    v2 := []int{32767}
    v3 := []string{("a" + ("\n\t" + ("hello" + ""))), "foo bar", (("" + "") + ("foo bar" + ("" + "foo bar"))), "\n\t"}
    for fi10 := 0; fi10 < 5; fi10++ {
        if len(v3) > 0 {
            v3[0] = "a"
        }
        _ = fi10
    }
    var v11 uint32 = 640
    var v12 int = 9223372036854775807
    v13 := GenS1{F0: (v12 >> ((42) & 63)), F1: 32767, F2: 0}
    v12 -= v12
    _ = len(v2)
    for fi14 := 0; fi14 < 11; fi14++ {
        fmt.Println(v12)
        v15 := GenS1{F0: fi14, F1: -168, F2: genFn1()}
        _ = fi14
        _ = v15
    }
    fmt.Println(v12, v12)
    if (0.0 != float64(((0.0 * v1) * v1))) {
        switch "\n\t" {
        case "":
            fmt.Println(v1, v11, v1)
        }
    }
    v1 += v1
    v16 := GenS1{F0: (v12 ^ v12), F1: 42, F2: 866}
    fmt.Println(v12, v12)
    var v17 uint64 = 42
    v17 -= (v17 * 0)
    var v18 uint = 42
    var v19 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
