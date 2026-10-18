package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() uint32 {
    if true {
        var v4 float64 = 3.14159
        _ = v4
    }
    if (1 == 864) {
        var v5 uint32 = 0
        _ = v5
    }
    v6 := GenS1{F0: -9223372036854775808}
    _ = v6
    return 4294967295
}

func main() {
    var v1 int16 = 0
    var v2 uint16 = 357
    v3 := make(map[int]bool)
    v7 := GenS1{F0: 9223372036854775807}
    var v8 string = "hello"
    v2 -= v2
    delete(v3, int(v1))
    switch 4294967295 {
    case 0:
        var v9 float32 = 3.14159
        _ = v9
    case 72:
        v10 := []int{}
        _ = v10
    default:
        v8 = "hello"
    }
    var v11 uint64 = 0
    v1 -= 0
    for fi12 := 0; fi12 < 17; fi12++ {
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := []int{fi13, fi12}
            var v15 int32 = -1
            v1 -= v1
            var v16 uint64 = 1
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        v8 = v8
        _ = fi12
    }
    var v17 int8 = 1
    v8 = "foo bar"
    fmt.Println(v2, v17, v1)
    switch 32767 {
    case -925:
        var v18 int = 32767
        _ = v18
    case 32768:
        if ((v1 + v1) <= (v1 / (((v1 / ((int16(v11)) | 1))) | 1))) {
            v19 := []int{2147483647, int(v1), 32768}
            v3[881] = true
            if (v17 <= 127) {
                fmt.Println(v17)
            } else {
                fmt.Println(v1, v1, v17)
                if false {
                    for fi20 := 0; fi20 < 4; fi20++ {
                        var v21 uint16 = (^(v2))
                        v11 = 2
                        _ = fi20
                        _ = v21
                    }
                    v1 -= -78
                }
            }
            _ = v19
        } else {
            fmt.Println(v11, v2)
            for fi22 := 0; fi22 < 16; fi22++ {
                var v23 bool = true
                _ = fi22
                _ = v23
            }
        }
    case -2147483648:
        fmt.Println(v1, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
