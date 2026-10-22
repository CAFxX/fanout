package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() int32 {
    for fi4 := 0; fi4 < 16; fi4++ {
        switch "hello" {
        case "\n\t":
            var v5 bool = true
            _ = v5
        case "hello":
            if ((-2.25 / -2.25) <= (3.14159 + float32(1.5))) {
                var v6 uint = 2
                var v7 int = (fi4 % ((-1) | 1))
                _ = v6
                _ = v7
            } else {
                v8 := GenS1{F0: 255}
                switch "foo bar" {
                case "foo bar":
                    v9 := []int{fi4, (fi4 | (fi4 / ((42) | 1))), (fi4 << ((42) & 63))}
                    _ = v9
                }
                _ = v8
            }
        }
        if (9223372036854775807 < -585) {
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 uint32 = 1
                v11 = v11
                v11 = (v11 + (v11 | v11))
                _ = fi10
                _ = v11
            }
            if (true || false) {
                switch "hello" {
                case "foo bar":
                    v12 := GenS1{F0: 2}
                    _ = v12
                case "a":
                    for fi13 := 0; fi13 < 18; fi13++ {
                        var v14 uint16 = 2
                        _ = fi13
                        _ = v14
                    }
                case "":
                    for fi15 := 0; fi15 < 3; fi15++ {
                        var v16 uint16 = 32767
                        fmt.Println(fi4, v16)
                        fmt.Println(v16, fi15)
                        _ = fi15
                        _ = v16
                    }
                }
                var v17 float64 = (-2.25 / 0.1)
                v17 -= v17
                _ = v17
            }
        }
        _ = fi4
    }
    var v18 uint64 = 958
    v18 = 569
    _ = v18
    return -838
}

func main() {
    var v1 float64 = 0.1
    var v2 float32 = 1.5
    var v3 int = 2147483647
    var v19 bool = (1 > 1)
    v20 := []int{(v3 * v3), 9223372036854775807, (v3 / (((v3 ^ 127)) | 1)), v3}
    if v19 {
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
    }
    fmt.Println(v19, v2)
    v19 = false
    for fi21 := 0; fi21 < 14; fi21++ {
        v1 = (((v1 * float64((0.1 + 0.0))) / (((v1 + (((v1 - v1) - v1) * 0.0)) - float64((v2 / 1e10))) / 1.0)) + float64(v2))
        var v22 int32 = -2147483648
        var v23 float32 = ((0.1 / ((float32(float64((v2 - -2.25))) + float32(float64(1e10))) / (v2 / (3.14159 / v2)))) + (-2.25 / 0.1))
        _ = fi21
        _ = v22
        _ = v23
    }
    v1 = 3.14159
    var v24 uint = 42
    if len(v20) > 0 {
        v20[0] = 0
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        var v26 string = "a"
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
}
