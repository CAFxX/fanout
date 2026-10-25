package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int8 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 string = "hello"
    v3 := make(map[string]bool)
    if false {
        for fi4 := 0; fi4 < 10; fi4++ {
            v5 := GenS1{F0: (!(2 == 127)), F1: 32767, F2: 16}
            _ = fi4
            _ = v5
        }
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 uint = 255
            v7 = v7
            fmt.Println(v7, fi6)
            _ = fi6
            _ = v7
        }
    } else {
        for fi8 := 0; fi8 < 2; fi8++ {
            v2 = ("hello" + "foo bar")
            _ = fi8
        }
    }
    v9 := GenS2{F0: -9223372036854775808}
    switch 9223372036854775807 {
    case -620:
        v2 = "\n\t"
    default:
        v2 = ((v2 + ("\n\t" + v2)) + ("hello" + v2))
    }
    var v10 uint = 2
    var v11 bool = true
    var v12 float64 = ((0.1 / 1e10) / ((-2.25 * float64(1e10)) / (3.14159 + float64(float32(float64((1.5 + 0.0)))))))
    var v13 bool = v11
    for fi14 := 0; fi14 < 4; fi14++ {
        switch (v2 + v2) {
        case "a":
            switch v2 {
            case "a":
                v15 := GenS2{F0: -9223372036854775808}
                _ = v15
            case "foo bar":
                fmt.Println(v12, v13)
            case "":
                for fi16 := 0; fi16 < 8; fi16++ {
                    var v17 string = ("" + "hello")
                    switch v2 {
                    case "":
                        v18 := GenS2{F0: -442}
                        _ = v18
                    }
                    fmt.Println(v2, fi16)
                    _ = fi16
                    _ = v17
                }
            }
        case "foo bar":
            v19 := []int{}
            _ = v19
        default:
            for fi20 := 0; fi20 < 2; fi20++ {
                v21 := []int{}
                var v22 int64 = -1
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        switch fi14 {
        case 128:
            var v23 int32 = 0
            _ = v23
        default:
            switch v2 {
            case "foo bar":
                var v24 int16 = -368
                _ = v24
            }
        }
        v25 := []int{}
        _ = fi14
        _ = v25
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        fmt.Println(v11, v12, v12)
        var v27 int64 = -1
        var v28 string = v2
        _ = fi26
        _ = v27
        _ = v28
    }
    if (177 <= 42) {
        var v29 int64 = -298
        _ = v29
    }
    for fi30 := 0; fi30 < 4; fi30++ {
        v31 := v1["a"]
        v10 -= (^(v10))
        _ = fi30
        _ = v31
    }
    v13 = v13
    if false {
        fmt.Println(v2, v13)
        var v32 bool = v13
        if true {
            v33 := []int{873}
            _ = v33
        }
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
