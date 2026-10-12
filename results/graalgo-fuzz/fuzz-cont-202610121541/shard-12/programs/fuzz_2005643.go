package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 uint64; F2 string }

func main() {
    var v1 float64 = 1.5
    var v2 uint = 0
    v3 := make(map[int]bool)
    var v4 string = "a"
    var v5 string = ("hello" + ("a" + v4))
    v2 = v2
    v2 = v2
    for fi6 := 0; fi6 < 14; fi6++ {
        v4 = "\n\t"
        _ = fi6
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        v5 = v4
        _ = fi7
    }
    v5 = (v5 + ("hello" + v4))
    v8 := GenS1{F0: 0, F1: 100, F2: 0}
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(v4, v4)
        fmt.Println(v4, v2, v2)
        _ = fi9
    }
    if true {
        var v10 int32 = 42
        fmt.Println(v5)
        var v11 int32 = v10
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: 1, F1: 1}
    v4 = (v5 + ("" + "hello"))
    v13 := GenS3{F0: -2147483648}
    for fi14 := 0; fi14 < 14; fi14++ {
        switch "a" {
        case "hello":
            for fi15 := 0; fi15 < 2; fi15++ {
                switch fi15 {
                case -304:
                    if (((!false) && ("a" == (v5 + ((("a" + (("" + ((v4 + v5) + v5)) + (((v5 + v5) + v4) + v5))) + "\n\t") + v5)))) && true) {
                        if false {
                            v5 = v4
                            var v16 uint16 = 2
                            _ = v16
                        }
                    } else {
                        v17 := GenS4{F0: 1, F1: 554, F2: ""}
                        v5 = (v5 + ("a" + ""))
                        fmt.Println(fi14, fi14)
                        _ = v17
                    }
                }
                _ = fi15
            }
        case "":
            if true {
                var v18 uint32 = 219
                _ = v18
            }
        default:
            var v19 int8 = 127
            _ = v19
        }
        _ = fi14
    }
    var v20 uint32 = 483
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
