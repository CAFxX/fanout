package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]bool)
    var v3 string = "\n\t"
    var v4 float32 = (0.1 / 1e10)
    var v5 int32 = 42
    var v6 float32 = 1.5
    var v7 uint16 = 255
    v8 := []int{1, -781, -850, 4294967295}
    for fi9 := 0; fi9 < 1; fi9++ {
        v7 = (v7 + (v7 - (^(v7))))
        _ = fi9
    }
    switch ((v3 + v3) + ("hello" + v3)) {
    case "hello":
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 uint32 = 2
            _ = fi10
            _ = v11
        }
    }
    var v12 uint32 = uint32(v7)
    _ = len(v2)
    var v13 uint = v1
    if len(v8) > 0 {
        v8[0] = 42
    }
    if (!true) {
        switch -2147483648 {
        case -9223372036854775808:
            if (true || false) {
                v6 = 1.5
                fmt.Println(v6, v7)
            }
        case -786:
            fmt.Println(v3, v7, v6)
        default:
            v14 := GenS1{F0: 524, F1: (^(v12))}
            _ = v14
        }
    }
    switch v3 {
    case "":
        v15 := GenS1{F0: 1, F1: v12}
        _ = v15
    case "hello":
        for fi16 := 0; fi16 < 15; fi16++ {
            switch (-(fi16)) {
            case 1:
                fmt.Println(v6, v3, v7)
            case 32768:
                if ("a" != "a") {
                    v17 := GenS1{F0: 18446744073709551615, F1: v12}
                    v8 = append(v8, 32768)
                    _ = v17
                } else {
                    if (("a" + v3) != (((v3 + v3) + ((v3 + v3) + v3)) + v3)) {
                        v5 += (-(v5))
                        var v18 int64 = 1
                        if (2 >= 205) {
                            v19 := v2[795]
                            _ = v19
                        }
                        _ = v18
                    }
                    for fi20 := 0; fi20 < 19; fi20++ {
                        var v21 uint16 = v7
                        _ = fi20
                        _ = v21
                    }
                }
            }
            _ = fi16
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
