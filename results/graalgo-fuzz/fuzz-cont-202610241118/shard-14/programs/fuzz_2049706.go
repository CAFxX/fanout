package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint = 923
    v2 := make(map[string]int)
    var v3 int64 = -514
    v3 = v3
    for fi4 := 0; fi4 < 20; fi4++ {
        var v5 int32 = -2147483648
        var v6 uint64 = 2
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int = fi7
        fmt.Println(v1)
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: "\n\t"}
    v10 := []int{2147483647, -485, -1, 127}
    v1 = v1
    for fi11 := 0; fi11 < 19; fi11++ {
        v3 = v3
        _ = fi11
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 int8 = -84
        v1 = v1
        v14 := GenS1{F0: (("a" + "hello") + "hello")}
        _ = fi12
        _ = v13
        _ = v14
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v15 := GenS2{F0: -73}
    v1 -= v1
    if len(v10) > 0 {
        v10[0] = -100
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        switch (^(fi16)) {
        case 9223372036854775807:
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 string = "hello"
                _ = fi17
                _ = v18
            }
        case 0:
            var v19 float32 = float32(-2.25)
            _ = v19
        case 128:
            fmt.Println(fi16, fi16, v3)
        }
        _ = fi16
    }
    var v20 uint = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v9
}
