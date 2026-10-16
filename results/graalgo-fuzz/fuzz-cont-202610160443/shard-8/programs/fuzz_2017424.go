package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0() uint32 {
    if (!true) {
        var v4 uint32 = 0
        for fi5 := 0; fi5 < 5; fi5++ {
            for fi6 := 0; fi6 < 4; fi6++ {
                var v7 uint64 = 1
                if false {
                    if false {
                        fmt.Println(v7)
                    } else {
                        v8 := GenS1{F0: 32767}
                        v9 := []int{}
                        var v10 int = (fi5 * (fi6 - fi5))
                        _ = v10
                        _ = v8
                        _ = v9
                    }
                    var v11 int16 = -1
                    _ = v11
                } else {
                    var v12 float32 = 1.5
                    fmt.Println(v12, fi5, fi6)
                    _ = v12
                }
                v7 -= (^(v7))
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
        _ = v4
    }
    var v13 uint8 = 1
    fmt.Println(v13)
    var v14 int16 = 42
    var v15 string = ("hello" + "a")
    _ = v13
    _ = v14
    _ = v15
    return 2147483648
}

func main() {
    var v1 uint8 = 128
    var v2 int = 127
    var v3 float32 = 1.5
    if true {
        var v16 bool = true
        v3 = 0.0
        _ = v16
    }
    var v17 int16 = 1000
    v2 -= v2
    fmt.Println(v3, v2, v17)
    var v18 string = "foo bar"
    var v19 uint8 = v1
    if (!(!true)) {
        var v20 float32 = (1e10 + v3)
        var v21 uint16 = 805
        v22 := GenS2{F0: 2147483647, F1: 512}
        _ = v20
        _ = v21
        _ = v22
    }
    var v23 bool = (v2 > (v2 ^ (v2 - -1)))
    v24 := GenS1{F0: (v17 | 1)}
    fmt.Println(v17, v18)
    var v25 uint8 = v19
    fmt.Println(v1)
    var v26 uint64 = 0
    fmt.Println(v3)
    fmt.Println(v23, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
