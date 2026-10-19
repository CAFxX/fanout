package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 string }
func genFn0(p0 uint) int32 {
    p0 = p0
    v7 := GenS1{F0: 1, F1: 64}
    _ = p0
    _ = v7
    return 1
}

func main() {
    var v1 uint64 = 657
    var v2 int16 = 987
    v3 := make(map[string]bool)
    v4 := make(map[string]uint32)
    var v5 uint = 65535
    var v6 string = "hello"
    fmt.Println(v5, v2, v5)
    v2 = -304
    v6 = (v6 + "\n\t")
    for fi8 := 0; fi8 < 4; fi8++ {
        fmt.Println(v2)
        _ = fi8
    }
    v9 := GenS1{F0: genFn0((^(v5))), F1: 42}
    v5 = 255
    v10 := GenS1{F0: -1, F1: 167}
    for fi11 := 0; fi11 < 20; fi11++ {
        delete(v3, "hello")
        _ = fi11
    }
    delete(v3, v6)
    v1 += (v1 >> ((1) & 63))
    v6 = v6
    v1 -= v1
    var v12 float32 = (1.5 * 1.5)
    var v13 int = 0
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 bool = (float64(v12) == 1.5)
        for fi16 := 0; fi16 < 9; fi16++ {
            v13 += 0
            if (v5 >= (v5 - (v5 - v5))) {
                v17 := GenS1{F0: 84, F1: uint8(v1)}
                fmt.Println(v1)
                if v15 {
                    v18 := []int{(fi14 - (-(fi16))), v13}
                    _ = len(v4)
                    _ = v18
                } else {
                    var v19 int16 = v2
                    switch v13 {
                    case 42:
                        v20 := GenS2{F0: v6}
                        _ = v20
                    default:
                        v2 -= v19
                    }
                    _ = v19
                }
                _ = v17
            }
            delete(v3, (v6 + (v6 + (("a" + ((("foo bar" + ("a" + v6)) + ("foo bar" + (v6 + v6))) + ((v6 + v6) + "hello"))) + ("\n\t" + v6)))))
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    v21 := GenS1{F0: 42, F1: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
