package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint64 = 2
    var v2 uint64 = 1
    var v3 int16 = -32768
    var v4 uint8 = 128
    var v5 int32 = -2147483648
    if false {
        switch 923 {
        case 4294967295:
            var v6 string = ("hello" + "foo bar")
            _ = v6
        }
    } else {
        for fi7 := 0; fi7 < 8; fi7++ {
            v8 := []int{fi7}
            _ = fi7
            _ = v8
        }
    }
    v9 := []int{}
    var v10 uint16 = 98
    for fi11 := 0; fi11 < 20; fi11++ {
        v10 = v10
        fmt.Println(v4, v1, fi11)
        v5 += v5
        _ = fi11
    }
    v12 := GenS1{F0: (v4 | (v4 + 128))}
    if len(v9) > 0 {
        v9[0] = 32768
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 string = "foo bar"
            var v16 int32 = v5
            fmt.Println(v3)
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
