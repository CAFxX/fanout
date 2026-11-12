package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int64 }
func genFn0(p0 uint8) int32 {
    var v6 uint = 238
    var v7 uint16 = 32767
    _ = p0
    _ = v6
    _ = v7
    return 181
}

func main() {
    var v1 bool = true
    var v2 int = 127
    var v3 string = ("a" + "a")
    v4 := make(map[string]bool)
    v5 := make(map[string]string)
    delete(v4, v3)
    v8 := v4[(v3 + v3)]
    if ((v3 == "") || false) {
        v8 = (229 == 2)
    } else {
        var v9 uint16 = 42
        for fi10 := 0; fi10 < 1; fi10++ {
            delete(v4, "")
            for fi11 := 0; fi11 < 20; fi11++ {
                _ = len(v5)
                v12 := GenS1{F0: 2147483648, F1: 4294967295, F2: 9223372036854775807}
                v3 = (v3 + ((v3 + v3) + "hello"))
                v3 = v3
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
        _ = v9
    }
    v1 = ((!((!(!false)) && (2 >= 979))) || (2147483647 != 0))
    _ = len(v5)
    v8 = (2 > 4294967295)
    var v13 int8 = 109
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := GenS1{F0: 2, F1: 0, F2: 1}
        _ = fi14
        _ = v15
    }
    var v16 uint32 = 1
    v2 = v2
    var v17 bool = ((!(!v1)) || (0 >= 684))
    v18 := []int{0, 106, (v2 + -9223372036854775808), 404}
    var v19 uint8 = 127
    if v8 {
        var v20 float64 = 3.14159
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
