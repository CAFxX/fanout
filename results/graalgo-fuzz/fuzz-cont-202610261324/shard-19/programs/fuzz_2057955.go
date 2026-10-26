package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() int {
    v6 := GenS1{F0: (407 > 718)}
    var v7 uint16 = 710
    switch 32767 {
    case 4294967295:
        for fi8 := 0; fi8 < 2; fi8++ {
            for fi9 := 0; fi9 < 5; fi9++ {
                v10 := GenS1{F0: true}
                _ = fi9
                _ = v10
            }
            var v11 int16 = 32767
            _ = fi8
            _ = v11
        }
    case 9223372036854775807:
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := GenS1{F0: (!(false || (496 != 4294967295)))}
            _ = fi12
            _ = v13
        }
    case -683:
        if (!true) {
            var v14 int8 = 1
            v15 := GenS1{F0: ((v7 << ((18446744073709551615) & 63)) != (v7 & 42))}
            v7 += v7
            _ = v14
            _ = v15
        }
    default:
        var v16 float32 = (0.1 - 3.14159)
        _ = v16
    }
    fmt.Println(v7)
    _ = v6
    _ = v7
    return 32767
}

func main() {
    var v1 float32 = 0.0
    var v2 float32 = (v1 * ((v1 + v1) * ((3.14159 - 1e10) + (v1 * v1))))
    v3 := make(map[string]string)
    var v4 int = -9223372036854775808
    var v5 uint = 0
    var v17 uint16 = 665
    for fi18 := 0; fi18 < 10; fi18++ {
        v19 := []int{v4, (fi18 - 791), (v4 >> ((v5) & 63)), genFn0()}
        v20 := GenS1{F0: true}
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: (375 <= -1)}
    var v22 float64 = 1e10
    v23 := GenS1{F0: (1 > -2147483648)}
    var v24 float64 = (float64(((v2 - (1.5 - 3.14159)) * 1e10)) + (v22 - (v22 / float64(v2))))
    fmt.Println(v24)
    fmt.Println(v4, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
