package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []int{1, -2147483648, 273, 0}
    v2 := make(map[int]int)
    var v3 int8 = 42
    v4 := []int{828, 4294967295}
    var v5 uint64 = 197
    var v6 string = ("hello" + "hello")
    v6 = (v6 + v6)
    v7 := GenS1{F0: false}
    var v8 int = 42
    if false {
        for fi9 := 0; fi9 < 8; fi9++ {
            v1 = append(v1, -658)
            _ = fi9
        }
        v8 = 1
        var v10 float64 = 1.5
        _ = v10
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        switch "foo bar" {
        case "":
            for fi12 := 0; fi12 < 18; fi12++ {
                fmt.Println(v6)
                v13 := GenS1{F0: false}
                v8 = fi11
                var v14 string = v6
                v15 := GenS1{F0: (!(((v6 + ("\n\t" + v6)) + (((v6 + (v14 + (("a" + (v14 + v14)) + (v6 + "\n\t")))) + v14) + v14)) > v6))}
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
        case "\n\t":
            fmt.Println(v5)
        case "hello":
            v16 := GenS1{F0: true}
            _ = v16
        }
        _ = fi11
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 uint = 0
        if true {
            if (true || true) {
                var v19 float32 = 0.0
                fmt.Println(v3)
                for fi20 := 0; fi20 < 6; fi20++ {
                    if (0 > 255) {
                        var v21 uint8 = 95
                        if true {
                            v18 += (v18 * v18)
                            v19 = v19
                        }
                        fmt.Println(fi20)
                        _ = v21
                    } else {
                        v5 -= 4294967295
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = fi20
                }
                _ = v19
            } else {
                var v22 float64 = 3.14159
                _ = v22
            }
            fmt.Println(v5, fi17, fi17)
        } else {
            _ = len(v1)
        }
        for fi23 := 0; fi23 < 15; fi23++ {
            fmt.Println(v3)
            _ = fi23
        }
        _ = fi17
        _ = v18
    }
    v6 = (v6 + ((v6 + v6) + v6))
    var v24 float64 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
