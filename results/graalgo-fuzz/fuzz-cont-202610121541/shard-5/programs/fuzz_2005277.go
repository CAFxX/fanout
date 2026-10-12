package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := []int{}
    var v2 uint = 746
    v3 := make(map[string]string)
    delete(v3, "")
    fmt.Println(v2)
    v2 = v2
    var v4 bool = true
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 10; fi6++ {
            switch (("\n\t" + "") + "hello") {
            case "foo bar":
                v7 := GenS1{F0: 42}
                _ = v7
            case "a":
                var v8 float64 = 1e10
                _ = v8
            default:
                v9 := GenS1{F0: 42}
                _ = v9
            }
            _ = fi6
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            v3["\n\t"] = ""
            fmt.Println(fi5)
            _ = fi10
        }
        var v11 string = "\n\t"
        _ = fi5
        _ = v11
    }
    _ = len(v3)
    var v12 uint16 = 65535
    var v13 int32 = -78
    v12 = 539
    var v14 bool = (42 <= -107)
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS1{F0: v12}
        _ = fi15
        _ = v16
    }
    var v17 uint64 = 989
    v2 = v2
    v18 := v3["hello"]
    var v19 uint32 = 956
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
