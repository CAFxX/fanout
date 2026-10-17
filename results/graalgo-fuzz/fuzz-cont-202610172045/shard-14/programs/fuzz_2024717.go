package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint32 }
func genFn0(p0 uint8) bool {
    p0 += p0
    v7 := GenS1{F0: 4294967295, F1: -77, F2: 2147483647}
    v8 := GenS1{F0: 976, F1: 111, F2: 4294967295}
    var v9 uint32 = 65535
    fmt.Println(v9, p0)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 string = ("a" + ("" + "hello"))
    v2 := make(map[int]uint32)
    var v3 float64 = 1e10
    v4 := []int{0, 729, 32767, 127}
    var v5 string = "a"
    v6 := []string{}
    v5 = ""
    v5 = "foo bar"
    for fi10 := 0; fi10 < 20; fi10++ {
        v3 -= 0.0
        if (18446744073709551615 > 1) {
            switch (fi10 + fi10) {
            case -266:
                var v11 float64 = v3
                _ = v11
            }
            fmt.Println(fi10)
        }
        v12 := GenS1{F0: 65535, F1: 1, F2: 2147483647}
        _ = fi10
        _ = v12
    }
    fmt.Println(v5)
    var v13 int64 = 1
    if len(v6) > 0 {
        v6[0] = ""
    }
    delete(v2, 32768)
    var v14 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
