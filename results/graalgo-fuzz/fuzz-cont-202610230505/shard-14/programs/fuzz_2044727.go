package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int16; F2 int64 }

func main() {
    var v1 string = "a"
    var v2 int8 = 42
    var v3 int16 = -32768
    var v4 int64 = 0
    var v5 string = ("\n\t" + "\n\t")
    if false {
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := GenS1{F0: 42}
            var v8 uint = 5
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 uint64 = 110
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := GenS2{F0: int8(fi10), F1: 65535, F2: 1}
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := []int{245, 32767, (fi10 % ((fi12) | 1)), int(v4)}
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        var v14 float32 = (((0.1 - 0.1) + 0.0) / 1.0)
        v15 := GenS3{F0: 915, F1: (v3 % ((0) | 1)), F2: 42}
        _ = v14
        _ = v15
    }
    fmt.Println(v3, v2, v1)
    v5 = v5
    var v16 uint32 = 2147483648
    var v17 uint64 = 0
    if (!false) {
        v5 = ((v5 + ((v1 + v1) + "foo bar")) + v5)
        var v18 int8 = (^(v2))
        v18 = (v2 & v18)
        _ = v18
    } else {
        var v19 uint16 = 313
        if ((3.14159 * 0.0) < (((0.1 / 1.0) + 0.0) - 0.0)) {
            v20 := GenS3{F0: v16, F1: v3, F2: 0}
            _ = v20
        }
        var v21 int = 32768
        _ = v19
        _ = v21
    }
    v17 = 1
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
