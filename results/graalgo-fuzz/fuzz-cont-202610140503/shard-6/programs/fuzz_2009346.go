package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int32 = 218
    var v2 string = "hello"
    var v3 int32 = 1
    v4 := make(map[uint32]int)
    v5 := make(map[string]int)
    var v6 int64 = 1
    if (!false) {
        fmt.Println(v6, v6)
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 float32 = 1.5
            _ = fi7
            _ = v8
        }
    } else {
        var v9 uint = 0
        if false {
            v4[42] = -699
            fmt.Println(v1, v1, v9)
        }
        _ = v9
    }
    v2 = v2
    if true {
        var v10 int64 = v6
        v6 -= v6
        var v11 bool = ((3.14159 - (float64((0.1 * 0.0)) + (((float64(1.5) + (1e10 / float64(1.5))) * (float64(1e10) + (-2.25 + (0.0 / 3.14159)))) / 1.0))) > 3.14159)
        _ = v10
        _ = v11
    }
    if false {
        if true {
            var v12 uint64 = 0
            _ = v12
        } else {
            v3 -= 239
        }
        var v13 uint8 = 250
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := GenS1{F0: 54}
            delete(v5, (((("\n\t" + v2) + ((v2 + (v2 + (v2 + ("foo bar" + (v2 + v2))))) + (v2 + v2))) + v2) + v2))
            v1 += v3
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        fmt.Println(v2)
    }
    var v16 uint = 667
    v16 -= 255
    v17 := GenS1{F0: -128}
    var v18 uint = v16
    v18 = 825
    for fi19 := 0; fi19 < 1; fi19++ {
        switch v2 {
        case "\n\t":
            v3 += (v1 & (v3 ^ 2147483647))
        case "foo bar":
            fmt.Println(v6, v16, v6)
        case "hello":
            switch (v2 + v2) {
            case "\n\t":
                for fi20 := 0; fi20 < 6; fi20++ {
                    v21 := GenS1{F0: 42}
                    _ = fi20
                    _ = v21
                }
            case "hello":
                v22 := []int{fi19, fi19, (-(fi19))}
                _ = v22
            }
        default:
            var v23 int32 = -255
            _ = v23
        }
        v16 = v18
        delete(v5, v2)
        _ = fi19
    }
    v3 -= (v3 ^ v3)
    var v24 string = ""
    v25 := []int{1, 609}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
