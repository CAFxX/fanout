package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int32 }
func genFn0(p0 int8, p1 uint16) int {
    if (334 <= 34) {
        v4 := []int{91, 0, -9223372036854775808}
        _ = v4
    } else {
        p1 = p1
        v5 := GenS1{F0: "hello", F1: 65535}
        _ = v5
    }
    p0 += p0
    _ = p0
    _ = p1
    return 32767
}

func main() {
    var v1 uint8 = 127
    var v2 uint32 = 42
    var v3 uint16 = 618
    fmt.Println(v3, v2)
    fmt.Println(v3)
    var v6 bool = false
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: ("\n\t" + "foo bar"), F1: 255}
        v2 = (v2 / (((v2 << ((65535) & 63))) | 1))
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 127
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 float64 = (-2.25 / (3.14159 + 3.14159))
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: -1}
    var v13 bool = false
    for fi14 := 0; fi14 < 7; fi14++ {
        if v13 {
            var v15 float64 = 1e10
            var v16 uint = 2
            var v17 bool = ((fi14 & (^(fi14))) >= int(v16))
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi14
    }
    if (1e10 == 0.1) {
        for fi18 := 0; fi18 < 15; fi18++ {
            v19 := GenS2{F0: 42}
            _ = fi18
            _ = v19
        }
    }
    v9 = v9
    var v20 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v6
    _ = v9
}
