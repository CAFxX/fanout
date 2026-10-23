package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[string]string)
    var v2 uint = 293
    v3 := make(map[string]int)
    delete(v1, "hello")
    v2 = v2
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := GenS1{F0: false}
        _ = fi4
        _ = v5
    }
    var v6 uint8 = 1
    var v7 uint = (v2 * v2)
    var v8 uint = 0
    var v9 bool = (v6 > (v6 >> ((1) & 63)))
    if v9 {
        if true {
            var v10 uint8 = 255
            for fi11 := 0; fi11 < 3; fi11++ {
                if v9 {
                    if (v9 || ("hello" < "\n\t")) {
                        fmt.Println(v8)
                    }
                    var v12 uint8 = v10
                    for fi13 := 0; fi13 < 1; fi13++ {
                        v14 := GenS2{F0: v10}
                        _ = fi13
                        _ = v14
                    }
                    _ = v12
                } else {
                    fmt.Println(fi11, v10, v9)
                }
                var v15 float32 = (float32(0.0) + 3.14159)
                _ = fi11
                _ = v15
            }
            _ = v10
        }
    } else {
        v8 -= (v2 ^ 2)
        v16 := GenS1{F0: (!v9)}
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := GenS2{F0: 1}
            var v19 float32 = 1e10
            var v20 int16 = 42
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
