package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 string }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 2
    var v3 uint32 = 4294967295
    v4 := make(map[int]string)
    v5 := make(map[string]int)
    v6 := v5["foo bar"]
    if (!true) {
        v6 -= int(v3)
        var v7 uint16 = 828
        _ = v7
    }
    v2 -= (v2 % ((v2) | 1))
    var v8 float32 = 3.14159
    if true {
        var v9 int64 = 9223372036854775807
        v3 -= v3
        _ = v9
    }
    var v10 int16 = int16(v6)
    switch -9223372036854775808 {
    case 2147483647:
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 uint64 = 1
            v13 := GenS1{F0: 65535, F1: "", F2: ("hello" + "a")}
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    var v14 float64 = -2.25
    v15 := GenS1{F0: 1, F1: ("" + ("\n\t" + "")), F2: "\n\t"}
    if true {
        v16 := []int{4294967295, (v6 ^ 128), (-(v6)), 42}
        _ = v16
    } else {
        var v17 float64 = 0.1
        v18 := GenS1{F0: 681, F1: "hello", F2: (("\n\t" + "") + "\n\t")}
        _ = v17
        _ = v18
    }
    v10 = v10
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 string = "foo bar"
        _ = fi19
        _ = v20
    }
    v6 += (^(v6))
    v21 := v4[(^(v6))]
    for fi22 := 0; fi22 < 12; fi22++ {
        v23 := GenS1{F0: 392, F1: v21, F2: ""}
        fmt.Println(fi22)
        switch v21 {
        case "":
            v3 = (^(v3))
        default:
            var v24 string = (("hello" + v21) + "")
            _ = v24
        }
        _ = fi22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
