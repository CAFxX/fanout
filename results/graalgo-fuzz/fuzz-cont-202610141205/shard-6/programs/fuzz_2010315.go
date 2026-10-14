package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]bool)
    var v3 int8 = -1
    var v4 float32 = (((((-2.25 / (1.5 / 0.1)) - 0.0) - 0.1) - 0.0) - (1e10 / (((0.1 + 0.0) - ((float32((((float64(((((3.14159 * float32(((1e10 * (-2.25 * ((1e10 + 1.5) * 0.1))) / float64(-2.25)))) / (((3.14159 - 0.0) * (0.0 - -2.25)) / (((1.5 / -2.25) + 1.5) - -2.25))) - (((0.0 * ((-2.25 / 1.0) - -2.25)) + (float32(((3.14159 / 1e10) * 1e10)) / 3.14159)) * (0.1 * 3.14159))) - 1.5)) - 3.14159) / ((0.1 - -2.25) / -2.25)) - ((1.5 + 1e10) * 1e10))) * (-2.25 - 3.14159)) + (1e10 / 3.14159))) + (-2.25 + 0.0))))
    var v5 int64 = 42
    var v6 string = "\n\t"
    delete(v1, 127)
    var v7 uint16 = 433
    var v8 uint = 65535
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := []int{fi9, (fi9 + 127), -9223372036854775808, fi9}
        v6 = "hello"
        for fi11 := 0; fi11 < 20; fi11++ {
            switch ("hello" + v6) {
            case "\n\t":
                var v12 int = (^(fi11))
                _ = v12
            }
            v6 = "hello"
            var v13 int32 = -1
            _ = fi11
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := []int{fi14, fi14, fi14}
        _ = fi14
        _ = v15
    }
    v5 = (v5 ^ (v5 >> (((v8 * v8)) & 63)))
    v16 := GenS1{F0: 2147483647, F1: true}
    v6 = "foo bar"
    var v17 bool = (false && ((487 <= 0) && true))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
