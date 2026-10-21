package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint64 }
type GenS2 struct { F0 bool }

func main() {
    var v1 float32 = ((0.0 * float32(float64(1e10))) / 1.0)
    v2 := []int{-1, 32767, 128, 42}
    var v3 string = ""
    for fi4 := 0; fi4 < 12; fi4++ {
        v1 += v1
        var v5 string = (((v3 + v3) + ("hello" + (("foo bar" + (v3 + v3)) + v3))) + v3)
        _ = fi4
        _ = v5
    }
    switch ("hello" + v3) {
    case "\n\t":
        if (true && (82 < 2147483647)) {
            var v6 int = -2147483648
            var v7 int = (^(v6))
            v3 = v3
            _ = v6
            _ = v7
        }
    case "hello":
        var v8 uint32 = 378
        _ = v8
    case "":
        v9 := GenS1{F0: -1, F1: 42, F2: 18446744073709551615}
        _ = v9
    }
    v3 = v3
    v1 += (((v1 * (v1 - v1)) + 0.0) * (v1 - v1))
    v3 = "foo bar"
    fmt.Println(v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((!(1000 == 845)) || (-2.25 >= 3.14159)) {
        switch 9223372036854775807 {
        case -1:
            if true {
                var v10 int16 = -269
                switch "foo bar" {
                case "a":
                    v11 := GenS2{F0: (1.5 != 3.14159)}
                    _ = v11
                case "hello":
                    v12 := GenS1{F0: 9223372036854775807, F1: -1, F2: 4294967295}
                    _ = v12
                }
                _ = v10
            }
        case 128:
            v1 = v1
        case 32767:
            var v13 string = v3
            _ = v13
        }
    }
    fmt.Println(v3, v1)
    v14 := GenS1{F0: -237, F1: 2147483647, F2: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
}
