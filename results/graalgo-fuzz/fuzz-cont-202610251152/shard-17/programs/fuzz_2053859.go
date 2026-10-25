package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 bool = true
    var v2 float64 = float64(0.0)
    var v3 uint = 1
    var v4 int8 = 127
    switch "\n\t" {
    case "foo bar":
        v4 += (v4 | -128)
    case "\n\t":
        for fi5 := 0; fi5 < 17; fi5++ {
            v6 := []int{fi5, fi5, fi5}
            if v1 {
                var v7 int = fi5
                v8 := GenS1{F0: v7}
                _ = v7
                _ = v8
            }
            var v9 string = "hello"
            _ = fi5
            _ = v6
            _ = v9
        }
    default:
        v1 = v1
    }
    if false {
        v10 := GenS1{F0: -2147483648}
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := GenS1{F0: fi11}
            v4 = v4
            switch fi11 {
            case 9223372036854775807:
                v13 := []int{fi11, fi11, 1, fi11}
                _ = v13
            }
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    var v14 uint32 = 0
    v14 -= 2147483648
    v15 := GenS1{F0: 2147483647}
    switch (("a" + "") + (("\n\t" + "foo bar") + "a")) {
    case "a":
        var v16 uint8 = 113
        _ = v16
    case "foo bar":
        var v17 uint8 = 0
        _ = v17
    default:
        switch 9223372036854775807 {
        case 32767:
            v18 := GenS1{F0: -2147483648}
            _ = v18
        }
    }
    var v19 uint32 = (v14 >> ((65535) & 63))
    v20 := []int{451, int(v14)}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
