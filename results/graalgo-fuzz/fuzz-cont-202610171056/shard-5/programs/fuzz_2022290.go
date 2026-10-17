package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int }
type GenS2 struct { F0 uint32; F1 int32 }

func main() {
    v1 := make(map[int]string)
    var v2 string = "a"
    v3 := make(map[uint32]int)
    v4 := GenS1{F0: 42, F1: 380, F2: 128}
    var v5 float64 = 0.0
    var v6 bool = (32 < 0)
    v7 := []int{32767, 127}
    v8 := GenS1{F0: 2, F1: 222, F2: 2147483647}
    if (!((v2 + "hello") == "foo bar")) {
        v6 = v6
        if v6 {
            var v9 int16 = 32767
            switch v2 {
            case "foo bar":
                for fi10 := 0; fi10 < 9; fi10++ {
                    _ = len(v7)
                    switch 2147483647 {
                    case -740:
                        v2 = "hello"
                    case 128:
                        for fi11 := 0; fi11 < 19; fi11++ {
                            var v12 uint64 = 2
                            _ = fi11
                            _ = v12
                        }
                    case 0:
                        v13 := GenS1{F0: 128, F1: 42, F2: (fi10 & (fi10 + fi10))}
                        _ = v13
                    }
                    _ = fi10
                }
            case "a":
                v14 := GenS2{F0: 4294967295, F1: -2147483648}
                _ = v14
            case "":
                var v15 bool = v6
                _ = v15
            default:
                var v16 uint64 = 2
                _ = v16
            }
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 uint32 = 2147483648
                _ = fi17
                _ = v18
            }
            _ = v9
        }
    }
    var v19 int16 = -685
    v19 = v19
    var v20 int = -2147483648
    switch (v2 + v2) {
    case "hello":
        var v21 int16 = v19
        _ = v21
    }
    _ = len(v7)
    fmt.Println(v6)
    for fi22 := 0; fi22 < 10; fi22++ {
        fmt.Println(v2, fi22)
        _ = fi22
    }
    var v23 uint16 = 964
    var v24 int = v20
    fmt.Println(v19)
    fmt.Println(v6, v6)
    fmt.Println(v24, v2, v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
