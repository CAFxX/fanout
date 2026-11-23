package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint32 = 2147483647
    var v2 int8 = -1
    var v3 int8 = v2
    var v4 int32 = 1
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 int16 = 1
        _ = fi5
        _ = v6
    }
    v3 -= v2
    var v7 float64 = (0.1 / 3.14159)
    var v8 uint = 1
    for fi9 := 0; fi9 < 18; fi9++ {
        v8 = (v8 * (v8 - v8))
        _ = fi9
    }
    v10 := GenS1{F0: 909}
    var v11 bool = (((("hello" + "hello") + "\n\t") + "foo bar") > (((("a" + "hello") + "") + "") + "hello"))
    fmt.Println(v3, v4)
    for fi12 := 0; fi12 < 2; fi12++ {
        switch ("\n\t" + ("" + "\n\t")) {
        case "\n\t":
            v11 = v11
        case "foo bar":
            v13 := []int{(-(fi12))}
            _ = v13
        default:
            v14 := []int{fi12, -2147483648}
            _ = v14
        }
        var v15 string = "foo bar"
        switch fi12 {
        case 2147483647:
            v16 := []int{fi12, fi12}
            _ = v16
        default:
            for fi17 := 0; fi17 < 16; fi17++ {
                v18 := []int{32767, fi12, fi17}
                v4 = v4
                _ = fi17
                _ = v18
            }
        }
        _ = fi12
        _ = v15
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        v8 -= v8
        v20 := []int{-814}
        var v21 int16 = 32767
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
