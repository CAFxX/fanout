package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }
func genFn0() uint {
    switch "" {
    case "":
        if (!((982 != 42) || (0 < 1))) {
            switch ("\n\t" + "hello") {
            case "foo bar":
                var v5 uint8 = 127
                _ = v5
            case "a":
                v6 := []int{}
                _ = v6
            case "":
                var v7 string = "\n\t"
                _ = v7
            }
            switch "hello" {
            case "a":
                for fi8 := 0; fi8 < 1; fi8++ {
                    v9 := GenS1{F0: 1, F1: 2, F2: 663}
                    for fi10 := 0; fi10 < 1; fi10++ {
                        var v11 uint8 = 52
                        var v12 int64 = 379
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    _ = fi8
                    _ = v9
                }
            }
            v13 := GenS2{F0: 65535}
            _ = v13
        }
    }
    var v14 int32 = 1
    _ = v14
    return 404
}

func main() {
    var v1 bool = (255 <= 42)
    v2 := make(map[string]int)
    v3 := []int{}
    v4 := []string{"foo bar", ("\n\t" + ("hello" + ("a" + ("\n\t" + "a"))))}
    v1 = (363 == 18446744073709551615)
    var v15 int16 = -1
    v16 := GenS1{F0: 42, F1: 2, F2: 0}
    _ = len(v3)
    v17 := GenS3{F0: 0}
    var v18 int64 = 664
    v19 := GenS1{F0: -1, F1: 2147483648, F2: 2}
    var v20 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
