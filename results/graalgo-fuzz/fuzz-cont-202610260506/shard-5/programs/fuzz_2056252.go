package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := []uint32{65535, 4294967295, 2147483647, 233}
    var v2 float32 = float32((0.0 / 1e10))
    var v3 bool = true
    v4 := make(map[string]string)
    var v5 uint64 = 504
    fmt.Println(v3)
    switch "foo bar" {
    case "":
        delete(v4, ("foo bar" + "\n\t"))
    case "\n\t":
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 uint = 607
            var v8 float64 = ((-2.25 / float64(v2)) - -2.25)
            if (v3 && (uint32(v5) <= 893)) {
                var v9 int32 = 108
                _ = v9
            }
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    v3 = true
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v10 bool = false
    v10 = (!(("hello" + "") <= "hello"))
    v11 := GenS1{F0: "a"}
    var v12 int16 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
