package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int32 }
type GenS2 struct { F0 int8; F1 string }
func genFn0() uint8 {
    v7 := []int{788, -9223372036854775808}
    if false {
        v7 = append(v7, 128)
        var v8 uint = 2
        _ = v8
    }
    v9 := GenS1{F0: -95, F1: ("\n\t" + "\n\t"), F2: -1}
    var v10 uint64 = 18446744073709551615
    v10 += 95
    _ = v10
    _ = v7
    _ = v9
    return 127
}

func main() {
    var v1 bool = (false && (true || false))
    v2 := []int{453}
    v3 := []string{"hello", ("foo bar" + ("\n\t" + "foo bar")), "a"}
    v4 := make(map[string]int)
    var v5 int16 = 1
    var v6 uint = 1
    v11 := GenS1{F0: -128, F1: "a", F2: -113}
    v4[""] = 32768
    if true {
        var v12 int32 = -97
        _ = v12
    }
    v1 = ((1e10 + 0.1) != 0.0)
    v1 = (32767 <= 0)
    if (42 == 32767) {
        for fi13 := 0; fi13 < 10; fi13++ {
            for fi14 := 0; fi14 < 17; fi14++ {
                v1 = v1
                fmt.Println(v5)
                _ = fi14
            }
            var v15 int = fi13
            _ = fi13
            _ = v15
        }
        v6 = 44
    } else {
        v16 := GenS2{F0: -128, F1: "hello"}
        _ = v16
    }
    v17 := v4[("foo bar" + "")]
    v6 -= (v6 >> ((v6) & 63))
    v6 = v6
    v4["hello"] = -689
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
