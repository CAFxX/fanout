package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint32) int {
    p0 += (p0 - 42)
    v6 := GenS1{F0: 100, F1: 1}
    p0 = 2147483648
    _ = p0
    _ = v6
    return 946
}

func main() {
    v1 := []int{467, 1, 42}
    v2 := make(map[string]string)
    var v3 uint16 = 65535
    var v4 float32 = 1.5
    v5 := []string{"hello", (("a" + (("foo bar" + ("hello" + "\n\t")) + "foo bar")) + ("foo bar" + "\n\t")), "\n\t"}
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(fi7)
        v4 -= 1e10
        _ = fi7
    }
    if ("" == "a") {
        v8 := GenS1{F0: 101, F1: 4294967295}
        v3 -= v3
        v9 := GenS1{F0: 1, F1: 4294967295}
        _ = v8
        _ = v9
    }
    var v10 float64 = float64(v4)
    v10 = (v10 - 1e10)
    fmt.Println(v4, v3)
    if false {
        var v11 int = 32767
        _ = v11
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = len(v5)
    var v12 int64 = 374
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
