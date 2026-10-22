package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 bool }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = (1e10 / 3.14159)
    var v3 int16 = 741
    var v4 string = (("a" + ("\n\t" + "a")) + "foo bar")
    for fi5 := 0; fi5 < 12; fi5++ {
        switch v4 {
        case "foo bar":
            var v6 int32 = -2147483648
            _ = v6
        }
        _ = fi5
    }
    var v7 uint32 = 2147483647
    v8 := GenS1{F0: (false || false), F1: (v7 & (^(v7))), F2: (v4 == "\n\t")}
    var v9 string = v4
    v10 := GenS1{F0: (false || false), F1: v7, F2: false}
    v11 := GenS1{F0: true, F1: v7, F2: (-859 >= -9223372036854775808)}
    if true {
        v12 := GenS1{F0: true, F1: 169, F2: (v2 == ((((v2 - (v2 + 0.1)) * v2) / v2) * 0.1))}
        for fi13 := 0; fi13 < 3; fi13++ {
            switch v9 {
            case "":
                var v14 string = ((v4 + (v9 + "a")) + (v9 + "hello"))
                _ = v14
            case "hello":
                var v15 int8 = 127
                _ = v15
            case "a":
                var v16 int32 = 42
                _ = v16
            default:
                v17 := []int{}
                _ = v17
            }
            _ = fi13
        }
        _ = v12
    }
    var v18 int64 = 9223372036854775807
    v3 += 32767
    delete(v1, 372)
    switch "foo bar" {
    case "\n\t":
        for fi19 := 0; fi19 < 4; fi19++ {
            switch (v4 + "\n\t") {
            case "":
                v20 := v1[32767]
                _ = v20
            default:
                var v21 int32 = -2147483648
                _ = v21
            }
            _ = fi19
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
