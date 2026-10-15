package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint }

func main() {
    v1 := []string{}
    v2 := make(map[uint32]string)
    var v3 float32 = 0.1
    v4 := make(map[uint32]int)
    v3 -= 0.1
    v1 = append(v1, "")
    if len(v1) > 0 {
        v1[0] = ""
    }
    for fi5 := 0; fi5 < 1; fi5++ {
        if false {
            var v6 float32 = 3.14159
            _ = v6
        } else {
            if true {
                for fi7 := 0; fi7 < 9; fi7++ {
                    v3 += v3
                    v3 = v3
                    switch ("\n\t" + "\n\t") {
                    case "":
                        var v8 int64 = 0
                        _ = v8
                    case "\n\t":
                        v9 := GenS1{F0: 185, F1: -1, F2: 432}
                        _ = v9
                    case "hello":
                        fmt.Println(fi7, v3)
                    }
                    _ = fi7
                }
                fmt.Println(v3, v3)
                var v10 int16 = -579
                _ = v10
            } else {
                fmt.Println(fi5, v3)
            }
        }
        _ = fi5
    }
    v1 = append(v1, "hello")
    if false {
        if false {
            v3 = v3
        } else {
            switch "foo bar" {
            case "":
                v3 = 3.14159
            }
            fmt.Println(v3)
        }
        v3 = v3
    } else {
        v3 = v3
        for fi11 := 0; fi11 < 9; fi11++ {
            v3 += float32((float64(0.1) + 0.1))
            _ = fi11
        }
        v12 := GenS2{F0: 4294967295, F1: 4294967295}
        _ = v12
    }
    fmt.Println(v3)
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
