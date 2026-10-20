package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int32 }

func main() {
    v1 := make(map[string]string)
    var v2 bool = true
    var v3 float64 = float64(-2.25)
    v4 := []int{}
    v5 := make(map[int]bool)
    v6 := GenS1{F0: 65535}
    v7 := GenS2{F0: -1, F1: -2147483648}
    fmt.Println(v3)
    v3 -= 1.5
    v4 = append(v4, 4294967295)
    v2 = ((v2 && (42 == 0)) || (!v2))
    var v8 string = ("\n\t" + "\n\t")
    for fi9 := 0; fi9 < 17; fi9++ {
        fmt.Println(v3)
        fmt.Println(fi9, fi9)
        fmt.Println(v3, fi9)
        _ = fi9
    }
    fmt.Println(v8, v3, v2)
    v8 = (v8 + v8)
    v10 := GenS1{F0: 2}
    for fi11 := 0; fi11 < 15; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            switch v8 {
            case "":
                var v13 int = (fi12 ^ (-(fi11)))
                _ = v13
            }
            v2 = v2
            if len(v4) > 0 {
                v4[0] = fi11
            }
            _ = fi12
        }
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 string = (v8 + (("hello" + (v8 + v8)) + v8))
            v2 = (v2 || v2)
            switch v8 {
            case "hello":
                v3 -= (float64(1.5) + (v3 + 3.14159))
            default:
                v16 := GenS1{F0: 255}
                _ = v16
            }
            _ = fi14
            _ = v15
        }
        _ = fi11
    }
    if v2 {
        v2 = (!true)
        v17 := v1[((v8 + v8) + "hello")]
        fmt.Println(v8, v17)
        _ = v17
    } else {
        fmt.Println(v3)
        if (((!(v2 && false)) || v2) && v2) {
            v2 = (0.1 >= v3)
            fmt.Println(v3)
            v2 = ((!false) && v2)
        }
        delete(v1, ((((v8 + ((v8 + (((v8 + v8) + ((v8 + v8) + v8)) + (v8 + "\n\t"))) + v8)) + "") + ("a" + "a")) + v8))
    }
    v2 = v2
    var v18 uint32 = 42
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 uint8 = 0
        v18 = (v18 % ((26) | 1))
        _ = fi19
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
