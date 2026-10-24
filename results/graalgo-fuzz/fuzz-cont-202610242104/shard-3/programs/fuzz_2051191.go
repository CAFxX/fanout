package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint }
type GenS2 struct { F0 int16; F1 string; F2 string }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"hello", "hello", (("\n\t" + "a") + "hello"), ("\n\t" + ("a" + "foo bar"))}
    var v3 uint32 = 863
    v4 := GenS1{F0: 42, F1: 513, F2: 81}
    var v5 bool = (!((1 >= 1) || (2 == 779)))
    v6 := GenS2{F0: -587, F1: "a", F2: ""}
    var v7 string = "a"
    v2 = append(v2, "")
    v5 = ((41 > 127) && v5)
    v5 = (0 > v3)
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(fi8)
        _ = fi8
    }
    var v9 bool = v5
    v10 := GenS3{F0: 32767, F1: 127}
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 uint32 = 531
        var v13 uint8 = 32
        if false {
            if true {
                v12 -= (^(v12))
                for fi14 := 0; fi14 < 13; fi14++ {
                    if (!false) {
                        _ = len(v1)
                        fmt.Println(v3, v5)
                        v7 = "a"
                        v5 = v5
                        v9 = ((0 == 0) && v5)
                    }
                    var v15 int32 = 42
                    _ = fi14
                    _ = v15
                }
            } else {
                switch (fi11 | 32768) {
                case 1:
                    var v16 bool = true
                    _ = v16
                }
                v9 = (true && v9)
                for fi17 := 0; fi17 < 3; fi17++ {
                    var v18 string = (v7 + "hello")
                    _ = fi17
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 20; fi19++ {
                delete(v1, ("" + (v7 + v7)))
                _ = fi19
            }
            var v20 int64 = -9223372036854775808
            _ = v20
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    switch 457 {
    case 128:
        fmt.Println(v5, v3, v5)
    case 766:
        v21 := GenS4{F0: 128, F1: 810}
        _ = v21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        var v23 int8 = -74
        _ = fi22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
