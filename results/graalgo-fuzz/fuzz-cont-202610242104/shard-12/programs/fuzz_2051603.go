package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int16; F1 string }
func genFn0() int16 {
    var v5 bool = (false || false)
    fmt.Println(v5)
    fmt.Println(v5)
    v5 = false
    var v6 string = "hello"
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 1000
    v3 := []string{"a", "\n\t", "a", "a"}
    v4 := make(map[int]bool)
    if len(v3) > 0 {
        v3[0] = "a"
    }
    v7 := GenS1{F0: genFn0()}
    var v8 int64 = -844
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            v4[(fi10 & (fi10 * (fi9 | fi10)))] = true
            _ = fi10
        }
        _ = fi9
    }
    _ = len(v3)
    if false {
        v11 := GenS1{F0: v2}
        switch 32768 {
        case 771:
            if true {
                if len(v3) > 0 {
                    v3[0] = ("foo bar" + ((("hello" + "\n\t") + ("\n\t" + "a")) + ("" + ("foo bar" + "foo bar"))))
                }
                var v12 int32 = 0
                v13 := GenS2{F0: 632, F1: 446}
                _ = v12
                _ = v13
            } else {
                v3 = append(v3, ("hello" + ("\n\t" + ("hello" + "a"))))
                v14 := GenS1{F0: v2}
                _ = v14
            }
        case -550:
            fmt.Println(v8, v8)
        }
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 uint = 286
            _ = fi15
            _ = v16
        }
        _ = v11
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        switch fi17 {
        case 4294967295:
            delete(v1, "")
        case 0:
            v2 -= v2
        }
        v2 = (v2 / ((1) | 1))
        _ = fi17
    }
    v18 := GenS3{F0: 188, F1: "\n\t"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
