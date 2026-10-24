package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 bool }
func genFn0(p0 int32, p1 uint8) bool {
    if false {
        v4 := []int{-9223372036854775808}
        _ = v4
    }
    v5 := []int{-9223372036854775808, -657, 127, -692}
    var v6 uint8 = (p1 % (((p1 / (((p1 - 1)) | 1))) | 1))
    v5 = append(v5, 269)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return false
}

func main() {
    var v1 uint64 = 947
    var v2 int = 106
    var v3 bool = (1 > -32768)
    fmt.Println(v2, v2)
    switch v2 {
    case 128:
        var v7 int32 = -1
        _ = v7
    }
    var v8 int = (^(v2))
    v9 := []int{v2, v2, (v2 % ((v8) | 1))}
    if len(v9) > 0 {
        v9[0] = (v2 / ((v8) | 1))
    }
    v1 = v1
    var v10 uint32 = 2147483648
    if v3 {
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 float32 = (((1e10 / 1.5) + -2.25) + 0.1)
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            var v13 uint = 4294967295
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            if len(v9) > 0 {
                v9[0] = int(v10)
            }
            var v15 int16 = -226
            _ = fi14
            _ = v15
        }
    }
    v16 := GenS1{F0: 6, F1: 42, F2: v3}
    if v3 {
        v10 = (v10 << ((2) & 63))
        switch v8 {
        case 32767:
            switch ("\n\t" + "a") {
            case "":
                for fi17 := 0; fi17 < 12; fi17++ {
                    var v18 int = 1
                    _ = fi17
                    _ = v18
                }
            default:
                for fi19 := 0; fi19 < 20; fi19++ {
                    v3 = true
                    _ = fi19
                }
            }
        case 4294967295:
            v3 = (!(!v3))
        }
        var v20 uint32 = v10
        _ = v20
    }
    var v21 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v8
    _ = v9
}
