package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 uint32 }
func genFn0() uint {
    var v6 string = ("a" + ("hello" + "a"))
    v7 := []int{}
    var v8 int32 = -851
    var v9 uint64 = 18446744073709551615
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = false
    var v3 uint = 1
    var v4 int16 = 1
    var v5 int8 = 100
    if (v2 && false) {
        v10 := []int{}
        _ = v10
    }
    v11 := []int{int(v3)}
    v12 := GenS1{F0: 2147483647}
    v13 := GenS2{F0: ""}
    fmt.Println(v4, v2)
    if false {
        switch 0 {
        case 128:
            if v2 {
                var v14 string = "hello"
                var v15 int32 = -1
                _ = v14
                _ = v15
            }
        }
        v2 = v2
        var v16 uint64 = 1
        _ = v16
    } else {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        fmt.Println(v3, v2)
        v18 := GenS2{F0: ((("a" + ("foo bar" + ("" + ("\n\t" + "")))) + "\n\t") + ("\n\t" + "\n\t"))}
        _ = fi17
        _ = v18
    }
    fmt.Println(v2, v5, v3)
    if false {
        var v19 int16 = 42
        switch 127 {
        case 4294967295:
            var v20 float32 = 0.1
            _ = v20
        }
        _ = v19
    } else {
        if (v4 == v4) {
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 float64 = 0.0
                _ = fi21
                _ = v22
            }
        }
        var v23 uint = 42
        _ = v23
    }
    v24 := GenS2{F0: ("foo bar" + "")}
    switch ("" + "foo bar") {
    case "a":
        v25 := GenS2{F0: ("foo bar" + "")}
        _ = v25
    case "\n\t":
        for fi26 := 0; fi26 < 11; fi26++ {
            switch fi26 {
            case 858:
                v4 += v4
            case 1:
                fmt.Println(fi26, v3)
            default:
                var v27 string = "hello"
                _ = v27
            }
            _ = fi26
        }
    case "":
        fmt.Println(v3, v2, v3)
    default:
        for fi28 := 0; fi28 < 11; fi28++ {
            fmt.Println(v2, v3, v3)
            _ = fi28
        }
    }
    if (9223372036854775807 > 9223372036854775807) {
        for fi29 := 0; fi29 < 20; fi29++ {
            fmt.Println(v2, v3, v2)
            var v30 int64 = 565
            _ = fi29
            _ = v30
        }
        v3 = (v3 + 255)
        for fi31 := 0; fi31 < 5; fi31++ {
            fmt.Println(fi31)
            if v2 {
                var v32 int16 = 32767
                v1["\n\t"] = 2147483647
                _ = v32
            } else {
                v33 := GenS1{F0: 798}
                _ = v33
            }
            _ = fi31
        }
    }
    v2 = v2
    v3 = 18446744073709551615
    switch "hello" {
    case "foo bar":
        for fi34 := 0; fi34 < 4; fi34++ {
            var v35 uint64 = 1
            switch -9223372036854775808 {
            case 32767:
                var v36 float64 = (1.5 * (1e10 + 0.0))
                _ = v36
            case 0:
                fmt.Println(v3, v5)
            default:
                v37 := GenS3{F0: fi34, F1: 289}
                _ = v37
            }
            _ = fi34
            _ = v35
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
