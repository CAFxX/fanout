package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() bool {
    v6 := []int{1, 127, 812, -1}
    var v7 string = "foo bar"
    var v8 int8 = -123
    _ = v6
    _ = v7
    _ = v8
    return (float64(3.14159) <= 0.1)
}
func genFn1() uint64 {
    v9 := GenS1{F0: 42}
    v10 := GenS1{F0: 0}
    _ = v10
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float32 = 0.1
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]uint32)
    var v5 float32 = ((v2 / v2) / 1e10)
    switch 0 {
    case 281:
        for fi11 := 0; fi11 < 9; fi11++ {
            v2 = (((v2 + (v5 * ((3.14159 * v2) - v2))) - v2) / ((v2 * v5) * ((v5 - float32(1e10)) / (float32(-2.25) - ((1.5 / (1e10 / (v2 - ((v5 - v5) / (-2.25 * 3.14159))))) - -2.25)))))
            var v12 uint32 = 2147483648
            _ = fi11
            _ = v12
        }
    default:
        v3[-9223372036854775808] = 2147483647
    }
    var v13 uint = 0
    var v14 int64 = -1
    if false {
        delete(v4, 42)
        var v15 float64 = (((0.0 - 0.1) + float64(v2)) - 1.5)
        _ = v15
    }
    v16 := v4[65535]
    fmt.Println(v13)
    delete(v4, (v16 << ((4294967295) & 63)))
    var v17 string = "\n\t"
    var v18 uint64 = (v1 / (((v1 | (v1 + (v1 + v1)))) | 1))
    if false {
        switch "\n\t" {
        case "a":
            var v19 uint32 = v16
            _ = v19
        case "hello":
            var v20 bool = true
            _ = v20
        default:
            if (true && true) {
                if (v14 < v14) {
                    v21 := []int{-1, 1, 32767, -1}
                    _ = v21
                }
            }
        }
    } else {
        v2 += 1e10
    }
    var v22 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
