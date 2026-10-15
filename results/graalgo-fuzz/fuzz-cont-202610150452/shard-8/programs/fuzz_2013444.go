package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 string }
type GenS2 struct { F0 int32; F1 uint; F2 string }
type GenS3 struct { F0 string; F1 int64; F2 uint8 }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint }

func main() {
    var v1 string = (("hello" + ("" + "hello")) + "\n\t")
    var v2 int32 = 2147483647
    var v3 uint32 = 428
    v4 := make(map[uint32]bool)
    var v5 uint8 = 1
    v5 += v5
    var v6 int64 = 9223372036854775807
    v3 -= 1
    v7 := GenS1{F0: "hello", F1: 347, F2: ("hello" + ("a" + "hello"))}
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS2{F0: v2, F1: 2, F2: v1}
        fmt.Println(v6, v6)
        var v10 float64 = ((((-2.25 - (0.1 + ((1e10 * (0.0 - 1.5)) * (0.1 + (0.0 / 1.5))))) - ((((3.14159 + (3.14159 * 1e10)) / 0.1) / 0.1) - ((((-2.25 / 1.0) / 0.1) + -2.25) / 1.5))) * ((1e10 / 1.0) / ((1.5 / float64((0.1 + 0.1))) * 1e10))) * 1e10)
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        v6 = v6
        for fi12 := 0; fi12 < 14; fi12++ {
            if true {
                switch (("" + v1) + v1) {
                case "foo bar":
                    v13 := GenS3{F0: ((("\n\t" + (v1 + "hello")) + v1) + v1), F1: -33, F2: (v5 / (((^(v5))) | 1))}
                    _ = v13
                case "hello":
                    fmt.Println(v1)
                case "a":
                    v14 := GenS4{F0: 2147483647, F1: 18446744073709551615, F2: 18446744073709551615}
                    _ = v14
                }
                v5 = 0
            } else {
                var v15 int32 = v2
                _ = v15
            }
            switch fi12 {
            case 32768:
                fmt.Println(v3)
            }
            v16 := GenS4{F0: (v3 & uint32(v5)), F1: 42, F2: 65535}
            _ = fi12
            _ = v16
        }
        v17 := GenS3{F0: v1, F1: (v6 * -1), F2: (v5 / (((^(v5))) | 1))}
        _ = fi11
        _ = v17
    }
    fmt.Println(v3, v1)
    var v18 uint64 = 18446744073709551615
    v19 := GenS2{F0: v2, F1: 42, F2: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
