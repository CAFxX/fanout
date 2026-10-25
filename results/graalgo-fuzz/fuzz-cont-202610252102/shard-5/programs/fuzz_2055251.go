package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint8 }
func genFn0(p0 uint8) uint {
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint16 = 2
        v8 += v8
        _ = fi7
        _ = v8
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(p0)
    }
    _ = p0
    return 417
}
func genFn1(p0 int64, p1 int16) bool {
    for fi9 := 0; fi9 < 20; fi9++ {
        switch (("hello" + ("a" + "foo bar")) + ("foo bar" + ("\n\t" + "a"))) {
        case "":
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 int16 = -7
                switch 127 {
                case -72:
                    var v12 int8 = -52
                    _ = v12
                case -962:
                    v13 := []int{4294967295}
                    _ = v13
                case 32767:
                    var v14 uint8 = 42
                    _ = v14
                }
                _ = fi10
                _ = v11
            }
        }
        var v15 int64 = 805
        var v16 uint = 0
        _ = fi9
        _ = v15
        _ = v16
    }
    var v17 int = -993
    for fi18 := 0; fi18 < 6; fi18++ {
        v19 := GenS1{F0: (fi18 | (v17 | fi18)), F1: v17, F2: 2}
        v20 := []int{1, 2147483647}
        var v21 uint64 = 1
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v17
    return (!(true || false))
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]bool)
    var v3 uint = 674
    var v4 int64 = 1
    v5 := []string{}
    var v6 uint = 0
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 int8 = 42
        _ = fi22
        _ = v23
    }
    v4 -= (v4 ^ v4)
    v3 = v6
    for fi24 := 0; fi24 < 10; fi24++ {
        v25 := GenS1{F0: fi24, F1: 732, F2: 128}
        v26 := v2[42]
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 uint16 = 1
    v28 := v2[2147483647]
    v29 := GenS1{F0: -422, F1: 369, F2: 2}
    fmt.Println(v4, v4, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
