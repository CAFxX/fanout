package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint16; F2 int8 }
type GenS3 struct { F0 string }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    v4 := GenS1{F0: 2147483648, F1: 840}
    v1 -= 1.5
    var v5 int = -9223372036854775808
    if false {
        v1 = v1
        v6 := GenS2{F0: 1, F1: 1, F2: 127}
        _ = v6
    } else {
        var v7 int16 = int16(v5)
        if ((-(v7)) > v7) {
            _ = len(v3)
        }
        _ = v7
    }
    v8 := GenS3{F0: "hello"}
    for fi9 := 0; fi9 < 6; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := []int{(v5 * 127), fi10, (^(fi10))}
            v1 = 1e10
            _ = len(v3)
            _ = fi10
            _ = v11
        }
        fmt.Println(v5)
        _ = fi9
    }
    v1 = 1e10
    for fi12 := 0; fi12 < 11; fi12++ {
        v13 := GenS2{F0: 42, F1: 2, F2: 127}
        var v14 int64 = -9223372036854775808
        var v15 uint32 = 1
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    if false {
        v5 = v5
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        if (0 >= 128) {
            switch "hello" {
            case "":
                for fi17 := 0; fi17 < 13; fi17++ {
                    var v18 bool = true
                    var v19 int32 = -699
                    fmt.Println(v5)
                    var v20 string = "hello"
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                }
            case "a":
                fmt.Println(fi16, fi16)
            default:
                if true {
                    v1 += v1
                    fmt.Println(fi16, v5)
                    var v21 int = v5
                    _ = v21
                }
            }
        }
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
