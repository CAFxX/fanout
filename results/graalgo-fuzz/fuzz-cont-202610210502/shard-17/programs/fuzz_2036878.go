package main

import (
    "fmt"
)

func genFn0(p0 uint64, p1 int8, p2 int32) string {
    p1 -= (p1 & p1)
    v6 := []int{int(p0), -211, 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "a"
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = true
    v3 := make(map[uint32]string)
    var v4 int = 827
    v5 := []string{}
    if (127 > 37) {
        switch ("hello" + "a") {
        case "\n\t":
            for fi7 := 0; fi7 < 5; fi7++ {
                var v8 uint64 = 4294967295
                _ = fi7
                _ = v8
            }
        }
    } else {
        v2 = (v2 && (!v2))
    }
    fmt.Println(v2, v2)
    var v9 bool = (((198 > 255) && (v2 && v2)) || v2)
    var v10 float32 = 0.1
    var v11 uint = 18446744073709551615
    fmt.Println(v11, v10)
    var v12 int8 = -1
    fmt.Println(v12, v11)
    v9 = (!v9)
    var v13 uint32 = 239
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
