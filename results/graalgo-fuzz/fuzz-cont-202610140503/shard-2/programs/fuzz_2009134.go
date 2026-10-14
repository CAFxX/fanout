package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 string }

func main() {
    var v1 int64 = -876
    var v2 int16 = -32768
    var v3 int64 = (-(v1))
    switch -265 {
    case 639:
        var v4 int32 = 561
        _ = v4
    }
    v5 := []int{}
    switch (((("hello" + "foo bar") + ("" + "foo bar")) + "a") + "foo bar") {
    case "hello":
        fmt.Println(v3, v3, v3)
    }
    if false {
        for fi6 := 0; fi6 < 5; fi6++ {
            v7 := GenS1{F0: 572, F1: v2, F2: ("foo bar" + ("hello" + ""))}
            fmt.Println(fi6, v1, v3)
            _ = fi6
            _ = v7
        }
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v8 float32 = 3.14159
    if false {
        for fi9 := 0; fi9 < 5; fi9++ {
            if true {
                for fi10 := 0; fi10 < 6; fi10++ {
                    v3 = (v3 * v3)
                    _ = fi10
                }
                var v11 int = 1
                v11 += (fi9 - -743)
                _ = v11
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi9
        }
    } else {
        v5 = append(v5, -785)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 int16 = v2
            _ = fi12
            _ = v13
        }
    }
    if (903 != 42) {
        switch "hello" {
        case "":
            var v14 int16 = v2
            _ = v14
        case "foo bar":
            switch ("hello" + "a") {
            case "\n\t":
                v2 = (v2 / ((42) | 1))
            case "hello":
                fmt.Println(v2)
            case "a":
                fmt.Println(v8, v8, v2)
            }
        }
        v15 := GenS1{F0: (v1 ^ (v3 / ((42) | 1))), F1: (-(v2)), F2: "hello"}
        _ = v15
    } else {
        for fi16 := 0; fi16 < 13; fi16++ {
            v2 = v2
            _ = fi16
        }
        var v17 uint8 = 13
        if true {
            v18 := GenS1{F0: (-(v1)), F1: v2, F2: "foo bar"}
            var v19 uint = 1
            _ = v18
            _ = v19
        }
        _ = v17
    }
    var v20 int64 = v3
    var v21 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v5
    _ = v8
}
