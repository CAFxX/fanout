package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 float64 = 3.14159
    var v2 float64 = (0.1 - ((v1 / float64(1.5)) - 0.0))
    var v3 string = ("a" + (("foo bar" + "a") + ("\n\t" + "\n\t")))
    var v4 int64 = 0
    v5 := make(map[string]bool)
    v6 := make(map[int]string)
    var v7 int64 = v4
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: ((("hello" + "hello") + (v3 + ("hello" + (v3 + v3)))) + "foo bar"), F1: 4294967295, F2: (-(v7))}
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 bool = (-128 > 0)
    for fi12 := 0; fi12 < 12; fi12++ {
        fmt.Println(v4, v11, v3)
        v13 := GenS1{F0: v3, F1: 42, F2: v7}
        for fi14 := 0; fi14 < 14; fi14++ {
            v5[((v3 + v3) + "a")] = v11
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    v15 := GenS2{F0: 42}
    if ((!(v2 > (((0.0 + (float64(0.1) / (float64(float32((v1 - v2))) / (v2 - 0.1)))) / (0.0 * (v1 + 1.5))) + (float64(-2.25) + v2)))) && v11) {
        v16 := []int{4294967295, -2147483648}
        v17 := GenS1{F0: "hello", F1: 954, F2: (-(v7))}
        _ = v16
        _ = v17
    }
    v18 := GenS3{F0: 239}
    fmt.Println(v2)
    switch v3 {
    case "":
        fmt.Println(v11, v1, v11)
    case "hello":
        var v19 int64 = 1
        _ = v19
    case "\n\t":
        delete(v5, (("\n\t" + v3) + ""))
    default:
        v20 := GenS2{F0: 2}
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
