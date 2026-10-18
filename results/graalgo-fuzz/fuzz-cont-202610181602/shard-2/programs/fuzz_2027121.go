package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int = 671
    var v2 float32 = 3.14159
    v3 := make(map[int]bool)
    v4 := []int{v1, v1, v1, 751}
    var v5 float64 = 1e10
    if (964 < -2147483648) {
        v1 = 32767
        v1 = (v1 / ((v1) | 1))
    } else {
        v5 -= 1e10
    }
    fmt.Println(v1, v1, v2)
    v6 := GenS1{F0: 1}
    delete(v3, v1)
    v7 := v3[818]
    v2 -= 0.0
    var v8 int16 = 125
    var v9 uint32 = 890
    var v10 float64 = ((v5 * 0.0) * -2.25)
    v8 = -409
    for fi11 := 0; fi11 < 18; fi11++ {
        switch (("\n\t" + "hello") + ("\n\t" + "")) {
        case "foo bar":
            var v12 uint32 = v9
            _ = v12
        }
        v10 -= 3.14159
        _ = fi11
    }
    v5 = (v5 * 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
