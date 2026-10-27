package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int = 929
    var v2 int64 = 42
    var v3 float64 = -2.25
    for fi4 := 0; fi4 < 4; fi4++ {
        v3 -= ((float64(0.0) + ((v3 * float64(1e10)) + (0.0 - ((v3 / (v3 + (v3 / (((v3 + v3) / (v3 + (0.0 + ((1.5 - -2.25) / v3)))) + 0.1)))) - (float64(((0.0 - 3.14159) + (0.1 * (0.0 + 0.0)))) + (v3 + v3)))))) + v3)
        v5 := []int{fi4, fi4}
        _ = fi4
        _ = v5
    }
    var v6 uint16 = 70
    fmt.Println(v2)
    switch int(v6) {
    case 33:
        v7 := GenS1{F0: v6}
        _ = v7
    case -2147483648:
        if false {
            var v8 string = "hello"
            _ = v8
        }
    case -9223372036854775808:
        var v9 string = "\n\t"
        _ = v9
    default:
        var v10 string = ("\n\t" + ("foo bar" + (("a" + "\n\t") + (("\n\t" + "\n\t") + "a"))))
        _ = v10
    }
    var v11 float64 = (3.14159 / v3)
    fmt.Println(v1, v3, v2)
    if (-1 < -133) {
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := []int{4294967295, (v1 + int(v6))}
            v14 := GenS1{F0: (^(v6))}
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 uint64 = 493
        v16 := GenS1{F0: v6}
        _ = v15
        _ = v16
    }
    v1 += (^(v1))
    var v17 string = ("foo bar" + "")
    v2 = v2
    if true {
        v18 := GenS2{F0: ((true || true) || (("hello" + v17) < v17))}
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v6
}
