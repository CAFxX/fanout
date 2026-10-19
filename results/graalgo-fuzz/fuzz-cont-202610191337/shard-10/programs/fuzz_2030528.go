package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int16 }
type GenS2 struct { F0 string }

func main() {
    var v1 uint32 = 2147483647
    var v2 float64 = 3.14159
    var v3 int32 = -500
    var v4 float64 = (v2 - float64(-2.25))
    var v5 int32 = v3
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := []int{-1}
        fmt.Println(fi6)
        _ = fi6
        _ = v7
    }
    v1 = (v1 % ((v1) | 1))
    var v8 uint32 = (^(v1))
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        var v9 int8 = 57
        _ = v9
    case "":
        var v10 uint = 255
        _ = v10
    case "foo bar":
        fmt.Println(v4)
    }
    v11 := []int{}
    for fi12 := 0; fi12 < 3; fi12++ {
        v2 -= (v2 * (v2 / 1e10))
        if true {
            fmt.Println(v8)
            v13 := GenS1{F0: 1, F1: 89, F2: 1000}
            _ = v13
        } else {
            var v14 int16 = 32767
            if len(v11) > 0 {
                v11[0] = fi12
            }
            _ = v14
        }
        for fi15 := 0; fi15 < 3; fi15++ {
            switch "a" {
            case "foo bar":
                var v16 uint32 = (^(v1))
                _ = v16
            case "\n\t":
                if false {
                    switch fi12 {
                    case -212:
                        var v17 uint8 = 255
                        _ = v17
                    case -180:
                        if false {
                            v1 = 2
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                        }
                    }
                }
                var v18 int = fi12
                v19 := GenS1{F0: 218, F1: 100, F2: -1}
                _ = v18
                _ = v19
            }
            _ = fi15
        }
        _ = fi12
    }
    if true {
        var v20 uint64 = 42
        var v21 string = "hello"
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 5; fi22++ {
            v23 := GenS2{F0: "a"}
            for fi24 := 0; fi24 < 11; fi24++ {
                if (false || (fi22 == fi22)) {
                    if true {
                        v25 := GenS2{F0: "hello"}
                        _ = v25
                    } else {
                        v2 = v4
                        var v26 bool = false
                        var v27 int16 = -32768
                        _ = v26
                        _ = v27
                    }
                    fmt.Println(v1, v8)
                } else {
                    fmt.Println(v1)
                }
                v28 := GenS1{F0: 773, F1: 56, F2: -32768}
                _ = fi24
                _ = v28
            }
            _ = len(v11)
            _ = fi22
            _ = v23
        }
        if len(v11) > 0 {
            v11[0] = -2147483648
        }
        var v29 int = 32767
        _ = v29
    }
    var v30 float32 = (0.0 / (1e10 * 0.1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v8
}
