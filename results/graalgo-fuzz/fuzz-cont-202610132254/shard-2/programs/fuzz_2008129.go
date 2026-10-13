package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
func genFn0(p0 int64, p1 uint16) uint8 {
    var v6 uint32 = 149
    v7 := []int{128, 9223372036854775807}
    v6 = (^(v6))
    p1 -= p1
    var v8 int8 = -19
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 1
    var v3 bool = (!(!true))
    var v4 int32 = 2147483647
    var v5 float32 = 1.5
    v9 := GenS1{F0: ("foo bar" + "foo bar"), F1: v3}
    v5 -= 1.5
    v3 = v3
    var v10 uint32 = (v2 | v2)
    if (42 > 123) {
        var v11 int32 = v4
        _ = v11
    } else {
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 uint16 = 1
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 float32 = v5
            for fi16 := 0; fi16 < 17; fi16++ {
                delete(v1, fi14)
                if v3 {
                    v2 -= 2
                } else {
                    var v17 string = "foo bar"
                    v18 := GenS1{F0: v17, F1: true}
                    v17 = "hello"
                    _ = v17
                    _ = v18
                }
                var v19 string = "a"
                _ = fi16
                _ = v19
            }
            var v20 int32 = -124
            _ = fi14
            _ = v15
            _ = v20
        }
    }
    v21 := GenS1{F0: "\n\t", F1: ("" < "hello")}
    v22 := []int{0, -949}
    v23 := GenS1{F0: (("hello" + "") + ""), F1: ((true && v3) || (42 <= 100))}
    v24 := GenS1{F0: (("hello" + ("\n\t" + "foo bar")) + "a"), F1: true}
    v10 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
