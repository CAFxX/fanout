package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() string {
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS1{F0: 634}
        _ = fi7
        _ = v8
    }
    var v9 int8 = 78
    fmt.Println(v9)
    _ = v9
    return "a"
}

func main() {
    var v1 uint8 = 128
    v2 := []int{127}
    var v3 int64 = 9223372036854775807
    var v4 uint64 = 2
    v5 := []int{}
    v6 := []string{("foo bar" + ""), "foo bar", "\n\t", "\n\t"}
    var v10 uint16 = 0
    v11 := GenS1{F0: 65535}
    var v12 uint16 = v10
    for fi13 := 0; fi13 < 2; fi13++ {
        if (float32(3.14159) != -2.25) {
            var v14 int32 = -1
            var v15 uint8 = (v1 & v1)
            _ = v14
            _ = v15
        } else {
            v10 += 32767
            var v16 float32 = 0.0
            _ = v16
        }
        var v17 uint8 = 0
        _ = fi13
        _ = v17
    }
    var v18 int = 1
    v18 += 9223372036854775807
    v18 = v18
    for fi19 := 0; fi19 < 2; fi19++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = len(v2)
        _ = fi19
    }
    var v20 uint64 = v4
    var v21 int32 = 1
    if (("\n\t" + "\n\t") < "") {
        var v22 bool = true
        for fi23 := 0; fi23 < 15; fi23++ {
            var v24 uint = 82
            _ = fi23
            _ = v24
        }
        v21 += -2147483648
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
