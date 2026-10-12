package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 string; F2 int }

func main() {
    v1 := make(map[string]string)
    var v2 int8 = -128
    var v3 uint16 = 2
    v4 := make(map[int]int)
    if true {
        for fi5 := 0; fi5 < 6; fi5++ {
            switch ("hello" + ("foo bar" + "\n\t")) {
            case "a":
                fmt.Println(v2, fi5)
            case "hello":
                fmt.Println(v3, v2)
            case "\n\t":
                var v6 int64 = 731
                _ = v6
            default:
                var v7 uint8 = 2
                _ = v7
            }
            v8 := GenS1{F0: 1, F1: (^(v3))}
            _ = fi5
            _ = v8
        }
        v9 := GenS2{F0: 42}
        fmt.Println(v3, v2)
        _ = v9
    }
    delete(v1, "a")
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 float32 = (0.1 * (float32(1e10) - 3.14159))
        v1[("hello" + "foo bar")] = (("" + (("\n\t" + "hello") + "hello")) + "a")
        var v12 string = "a"
        _ = fi10
        _ = v11
        _ = v12
    }
    v2 -= 100
    v13 := GenS2{F0: 128}
    v14 := GenS1{F0: -925, F1: 255}
    for fi15 := 0; fi15 < 3; fi15++ {
        switch (("\n\t" + "foo bar") + ("hello" + "a")) {
        case "a":
            fmt.Println(v2, v3)
        case "foo bar":
            var v16 uint8 = 0
            _ = v16
        default:
            switch (-(fi15)) {
            case 0:
                var v17 bool = true
                _ = v17
            }
        }
        switch "foo bar" {
        case "hello":
            if (!true) {
                fmt.Println(fi15)
                v2 += 0
                v18 := GenS2{F0: 255}
                _ = v18
            } else {
                fmt.Println(v3, fi15)
                v2 = v2
                if true {
                    for fi19 := 0; fi19 < 19; fi19++ {
                        v20 := []int{(-(fi15)), 2147483647, 42}
                        var v21 bool = false
                        v22 := GenS2{F0: 4}
                        _ = fi19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    v23 := GenS2{F0: 2}
                    _ = v23
                } else {
                    v2 = (v2 | v2)
                }
            }
        case "a":
            for fi24 := 0; fi24 < 19; fi24++ {
                v2 = -104
                _ = fi24
            }
        case "\n\t":
            var v25 int64 = 42
            _ = v25
        }
        _ = fi15
    }
    v3 = v3
    fmt.Println(v3)
    var v26 int8 = (v2 / ((v2) | 1))
    for fi27 := 0; fi27 < 7; fi27++ {
        for fi28 := 0; fi28 < 18; fi28++ {
            fmt.Println(v3, v3, fi27)
            if false {
                var v29 float32 = 0.1
                _ = v29
            } else {
                switch "foo bar" {
                case "\n\t":
                    v30 := GenS1{F0: -9223372036854775808, F1: 32767}
                    _ = v30
                case "a":
                    v4[(fi27 ^ (fi28 ^ fi27))] = fi28
                case "hello":
                    v31 := GenS1{F0: -155, F1: v3}
                    _ = v31
                }
                v32 := GenS3{F0: "foo bar"}
                var v33 int = fi27
                _ = v32
                _ = v33
            }
            _ = fi28
        }
        switch fi27 {
        case -590:
            if true {
                v3 = 381
            }
        case 9223372036854775807:
            v26 = v26
        case 42:
            v34 := GenS4{F0: (42 <= 464), F1: (("\n\t" + "foo bar") + "foo bar"), F2: fi27}
            _ = v34
        }
        _ = fi27
    }
    var v35 uint64 = 18446744073709551615
    var v36 int = 1
    v36 += v36
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
