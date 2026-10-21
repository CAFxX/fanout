package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{}
    var v3 int8 = 0
    v4 := make(map[int]string)
    v3 -= v3
    if false {
        switch 4294967295 {
        case 32768:
            fmt.Println(v3)
        case 187:
            for fi5 := 0; fi5 < 6; fi5++ {
                var v6 uint32 = 0
                _ = fi5
                _ = v6
            }
        }
        var v7 int16 = -372
        _ = v7
    }
    var v8 string = "foo bar"
    _ = len(v4)
    v8 = "hello"
    v8 = v8
    v3 = v3
    fmt.Println(v3, v8)
    var v9 float64 = 1.5
    var v10 bool = false
    switch 4294967295 {
    case 32768:
        if (("a" > ("foo bar" + ("\n\t" + (v8 + "hello")))) && ("" != (("" + (v8 + v8)) + "a"))) {
            var v11 string = "hello"
            var v12 uint8 = 127
            _ = v11
            _ = v12
        }
    case 42:
        var v13 string = "foo bar"
        _ = v13
    default:
        v9 = v9
    }
    var v14 int64 = -818
    if (0 >= 5) {
        switch v8 {
        case "\n\t":
            if ((^(v3)) != -70) {
                var v15 uint32 = 2
                var v16 string = v8
                v16 = ((((((v16 + (("\n\t" + (v8 + (v8 + "foo bar"))) + v16)) + (v16 + "\n\t")) + v16) + "\n\t") + "a") + (("hello" + (((v8 + v8) + v8) + v16)) + ""))
                _ = v15
                _ = v16
            }
        case "foo bar":
            var v17 uint8 = 2
            _ = v17
        }
        v18 := GenS1{F0: 65535, F1: 0}
        _ = v18
    }
    switch "\n\t" {
    case "a":
        v19 := GenS1{F0: 826, F1: 65535}
        _ = v19
    default:
        var v20 float64 = 1e10
        _ = v20
    }
    v21 := GenS1{F0: 65535, F1: 0}
    v14 -= 0
    v8 = v8
    switch "foo bar" {
    case "a":
        for fi22 := 0; fi22 < 12; fi22++ {
            v10 = false
            if true {
                var v23 string = v8
                _ = v23
            } else {
                switch v8 {
                case "\n\t":
                    v9 = -2.25
                }
                v9 = -2.25
                if v10 {
                    var v24 float64 = 3.14159
                    _ = v24
                }
            }
            v8 = "foo bar"
            _ = fi22
        }
    case "":
        var v25 uint = 4294967295
        _ = v25
    default:
        var v26 bool = ((32767 < -1) && v10)
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
