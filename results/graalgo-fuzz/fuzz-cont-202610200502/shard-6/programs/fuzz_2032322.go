package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[int]bool)
    var v2 bool = (false || false)
    var v3 int8 = 127
    var v4 uint = 635
    var v5 int = -32
    var v6 uint64 = 4294967295
    var v7 int16 = 42
    v8 := GenS1{F0: "a"}
    fmt.Println(v4, v2)
    if true {
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 uint8 = 0
            for fi11 := 0; fi11 < 2; fi11++ {
                v12 := GenS1{F0: "a"}
                _ = len(v1)
                v13 := []int{fi11, 2147483647}
                fmt.Println(v4, v4, v10)
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi9
            _ = v10
        }
        v4 += v4
        for fi14 := 0; fi14 < 11; fi14++ {
            v7 = (v7 - v7)
            var v15 int8 = 127
            _ = fi14
            _ = v15
        }
    } else {
        v3 += v3
        delete(v1, v5)
        v16 := v1[(v5 + -860)]
        _ = v16
    }
    fmt.Println(v4, v6)
    v2 = (false || ((v3 - (-(v3))) >= 70))
    v17 := GenS2{F0: (v7 + (v7 % (((v7 % (((v7 / ((v7) | 1))) | 1))) | 1)))}
    v1[944] = ((!true) && ("foo bar" <= "hello"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
