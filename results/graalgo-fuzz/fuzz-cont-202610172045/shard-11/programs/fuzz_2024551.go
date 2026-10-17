package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int8 = 96
    var v2 float32 = -2.25
    var v3 float32 = 3.14159
    var v4 float32 = (((v2 * v2) * v3) / v3)
    var v5 int32 = -2147483648
    if (!(true && true)) {
        for fi6 := 0; fi6 < 3; fi6++ {
            for fi7 := 0; fi7 < 19; fi7++ {
                switch "foo bar" {
                case "foo bar":
                    v3 -= v2
                case "a":
                    v4 -= 1e10
                case "hello":
                    v5 = 2147483647
                }
                _ = fi7
            }
            _ = fi6
        }
        v8 := GenS1{F0: 42}
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 string = "foo bar"
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 uint32 = 2147483647
                v12 = 767
                v13 := GenS1{F0: 4294967295}
                fmt.Println(v1, v12, fi11)
                v14 := []int{1, fi9}
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    v15 := []int{}
    v1 -= 38
    v4 = 0.1
    v3 = v3
    var v16 uint64 = 1
    fmt.Println(v2)
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    var v17 float32 = (v2 - 1.5)
    if (42 >= 32767) {
        v15 = append(v15, int(v5))
        _ = len(v15)
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
