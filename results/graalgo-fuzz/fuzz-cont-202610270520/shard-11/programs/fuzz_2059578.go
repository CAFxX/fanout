package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 string = "foo bar"
    var v2 int8 = -50
    var v3 int16 = -683
    v3 = v3
    var v4 int32 = -1
    var v5 int16 = 541
    var v6 float32 = 0.0
    var v7 uint16 = 42
    v8 := GenS1{F0: v1}
    for fi9 := 0; fi9 < 14; fi9++ {
        v2 -= (v2 << ((105) & 63))
        v7 = v7
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := []int{698, -177, (^(fi10)), (fi10 ^ (fi9 << ((768) & 63)))}
            v12 := GenS1{F0: v1}
            var v13 int = (fi9 ^ -9223372036854775808)
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi9
    }
    switch 333 {
    case 341:
        var v14 bool = true
        _ = v14
    case -129:
        var v15 bool = (3.14159 != 1.5)
        _ = v15
    }
    if false {
        v16 := []int{-890}
        _ = v16
    }
    switch (((v1 + "hello") + ((v1 + v1) + v1)) + v1) {
    case "\n\t":
        v17 := []int{32768, 32767}
        _ = v17
    }
    switch 127 {
    case -1:
        v18 := []int{32768, int(v5), 127, 127}
        _ = v18
    case 127:
        v3 -= (v5 % ((-508) | 1))
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        switch (("\n\t" + (v1 + v1)) + ("\n\t" + "foo bar")) {
        case "a":
            var v20 float64 = ((1e10 - 1e10) - 3.14159)
            _ = v20
        case "\n\t":
            var v21 uint32 = 1
            _ = v21
        default:
            switch -9223372036854775808 {
            case 487:
                for fi22 := 0; fi22 < 18; fi22++ {
                    switch ("" + "foo bar") {
                    case "hello":
                        v23 := []int{0, fi19, (^(fi19)), fi22}
                        _ = v23
                    }
                    _ = fi22
                }
            case 128:
                fmt.Println(fi19, v7, v6)
            }
        }
        for fi24 := 0; fi24 < 10; fi24++ {
            if true {
                fmt.Println(fi19)
                var v25 uint = 18446744073709551615
                _ = v25
            }
            v4 = v4
            v26 := []int{(fi19 * fi24), (fi24 | fi24)}
            _ = fi24
            _ = v26
        }
        _ = fi19
    }
    var v27 bool = ((v4 & (^(v4))) != (^(v4)))
    v27 = false
    v2 -= 100
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
