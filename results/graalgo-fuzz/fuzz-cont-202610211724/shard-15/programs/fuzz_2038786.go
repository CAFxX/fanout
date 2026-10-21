package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint; F1 string; F2 int64 }
func genFn0() string {
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: 255, F1: 42}
        v8 := GenS1{F0: 21, F1: 0}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int32 = 2147483647
    _ = v9
    return "\n\t"
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = false
    var v3 int32 = -238
    v4 := make(map[string]string)
    v5 := []int{1, -330, -9223372036854775808, -696}
    var v10 uint64 = 42
    switch 328 {
    case 127:
        if len(v5) > 0 {
            v5[0] = 895
        }
    case 32768:
        v11 := v4[("" + "foo bar")]
        _ = v11
    default:
        v3 -= v3
    }
    var v12 int16 = 1000
    var v13 bool = (168 == 255)
    if v13 {
        var v14 float64 = 3.14159
        v12 -= v12
        _ = v14
    }
    var v15 int64 = -1
    for fi16 := 0; fi16 < 14; fi16++ {
        v3 = int32(v12)
        _ = fi16
    }
    if (!(623 != 4294967295)) {
        if v13 {
            v17 := v4["hello"]
            for fi18 := 0; fi18 < 17; fi18++ {
                for fi19 := 0; fi19 < 18; fi19++ {
                    v15 = (v15 << ((4294967295) & 63))
                    v20 := v1[v17]
                    delete(v4, ((v17 + "hello") + (("hello" + v17) + "foo bar")))
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
            _ = v17
        }
        v2 = (!false)
        v21 := GenS2{F0: 42, F1: ("a" + "\n\t"), F2: (v15 >> ((255) & 63))}
        _ = v21
    } else {
        if true {
            v12 -= v12
            _ = len(v4)
        }
        var v22 int8 = -1
        if len(v5) > 0 {
            v5[0] = 9223372036854775807
        }
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
