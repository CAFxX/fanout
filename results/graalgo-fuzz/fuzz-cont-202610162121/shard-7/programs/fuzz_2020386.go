package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 997
    var v3 uint8 = 0
    v4 := make(map[int]string)
    var v5 bool = true
    var v6 uint64 = 2
    if true {
        for fi7 := 0; fi7 < 12; fi7++ {
            v6 = (v6 - v6)
            fmt.Println(v2)
            _ = fi7
        }
        var v8 int = 253
        _ = v8
    } else {
        var v9 uint16 = 32767
        fmt.Println(v3, v3, v6)
        _ = v9
    }
    var v10 uint16 = 1
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 uint32 = 2147483648
        var v13 uint64 = v6
        v12 = v12
        _ = fi11
        _ = v12
        _ = v13
    }
    switch ((("hello" + ((("" + "\n\t") + "hello") + (("foo bar" + "") + "hello"))) + ("hello" + "a")) + ((("a" + "a") + (("a" + ("foo bar" + "a")) + "\n\t")) + "")) {
    case "\n\t":
        v14 := []int{}
        _ = v14
    case "":
        v3 += (^(v3))
    case "a":
        var v15 bool = (false || false)
        _ = v15
    }
    if (!false) {
        for fi16 := 0; fi16 < 5; fi16++ {
            fmt.Println(v6)
            _ = fi16
        }
        for fi17 := 0; fi17 < 20; fi17++ {
            switch "a" {
            case "a":
                v5 = v5
            }
            v18 := []int{fi17, (fi17 ^ fi17)}
            for fi19 := 0; fi19 < 12; fi19++ {
                if true {
                    v20 := GenS1{F0: (fi19 / ((195) | 1))}
                    _ = v20
                }
                fmt.Println(fi17, v3)
                _ = fi19
            }
            _ = fi17
            _ = v18
        }
        v5 = v5
    }
    v21 := []int{-550, 0}
    for fi22 := 0; fi22 < 4; fi22++ {
        fmt.Println(v10, fi22, v10)
        _ = fi22
    }
    if (0 == 71) {
        if v5 {
            v3 -= v3
            var v23 int32 = v2
            _ = v23
        }
        v2 = v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
