package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
func genFn0() int {
    var v4 int32 = -2147483648
    var v5 string = "a"
    _ = v4
    _ = v5
    return 2147483647
}

func main() {
    v1 := []uint32{2, 1}
    var v2 string = "\n\t"
    v3 := make(map[string]uint32)
    var v6 uint = 884
    fmt.Println(v6, v2)
    for fi7 := 0; fi7 < 4; fi7++ {
        if (2 != 4294967295) {
            var v8 int = fi7
            v9 := GenS1{F0: 1, F1: 204}
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 uint16 = 255
                v8 = 4294967295
                _ = fi10
                _ = v11
            }
            _ = v8
            _ = v9
        }
        v2 = (v2 + v2)
        _ = fi7
    }
    v12 := GenS1{F0: uint16(v6), F1: 127}
    v13 := GenS1{F0: 434, F1: 255}
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 uint = v6
        var v16 int64 = 688
        _ = fi14
        _ = v15
        _ = v16
    }
    switch ("\n\t" + v2) {
    case "a":
        if true {
            var v17 string = ((("foo bar" + v2) + "") + "a")
            _ = len(v3)
            _ = v17
        }
    case "":
        var v18 int16 = 1
        _ = v18
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        _ = len(v1)
        var v20 int16 = -1
        v20 += v20
        _ = fi19
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v6
}
