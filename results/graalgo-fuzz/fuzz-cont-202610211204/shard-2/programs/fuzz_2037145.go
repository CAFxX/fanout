package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
func genFn0(p0 bool) uint8 {
    v6 := GenS1{F0: 1, F1: true}
    v7 := GenS1{F0: 255, F1: (p0 && true)}
    v8 := GenS1{F0: 65535, F1: false}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 uint8 = 0
    var v2 bool = (!(false && (4294967295 != 2147483647)))
    v3 := []string{("" + ""), ("hello" + "foo bar"), "foo bar"}
    var v4 string = "\n\t"
    var v5 uint64 = 412
    v9 := GenS1{F0: 32767, F1: (!(0.0 < (1e10 + 3.14159)))}
    v10 := GenS1{F0: 65535, F1: v2}
    if ((!v2) && (2 < 65535)) {
        var v11 float32 = (0.1 * -2.25)
        _ = v11
    }
    fmt.Println(v2, v1, v4)
    var v12 uint16 = 32767
    fmt.Println(v12, v12, v4)
    _ = len(v3)
    var v13 uint32 = 42
    var v14 int32 = 850
    if (-2147483648 >= 42) {
        fmt.Println(v4)
    } else {
        v3 = append(v3, v4)
    }
    v3 = append(v3, v4)
    var v15 int16 = -32768
    if v2 {
        for fi16 := 0; fi16 < 2; fi16++ {
            v5 = 4294967295
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v15 -= v15
            _ = fi16
        }
    }
    var v17 uint32 = (v13 & (v13 / (((v13 * 273)) | 1)))
    fmt.Println(v17, v12, v4)
    v14 -= (v14 - v14)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
