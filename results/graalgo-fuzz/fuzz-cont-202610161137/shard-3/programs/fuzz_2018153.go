package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
func genFn0(p0 int, p1 uint8, p2 string) uint32 {
    v4 := []int{128, p0}
    if false {
        if false {
            if true {
                p2 = ""
            } else {
                fmt.Println(p1)
            }
        }
        for fi5 := 0; fi5 < 13; fi5++ {
            p2 = p2
            _ = fi5
        }
    }
    switch 32767 {
    case 580:
        for fi6 := 0; fi6 < 6; fi6++ {
            p0 = fi6
            _ = fi6
        }
    case 32767:
        switch ("a" + "hello") {
        case "":
            if (true || true) {
                p0 -= p0
                v7 := GenS1{F0: 4294967295, F1: -844}
                p0 += (p0 / ((p0) | 1))
                _ = v7
            } else {
                var v8 uint16 = 32767
                _ = v8
            }
        case "hello":
            var v9 int8 = 100
            _ = v9
        case "\n\t":
            switch 4294967295 {
            case -1:
                var v10 uint16 = 446
                _ = v10
            case 9223372036854775807:
                v4 = append(v4, (^(p0)))
            case 0:
                var v11 uint16 = 0
                _ = v11
            }
        }
    }
    if false {
        var v12 uint16 = 255
        var v13 uint8 = (p1 % ((p1) | 1))
        var v14 bool = ("foo bar" > p2)
        _ = v12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        fmt.Println(p0, p0)
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 65535
}

func main() {
    var v1 uint = 255
    v2 := make(map[string]uint32)
    var v3 float32 = 1.5
    var v16 int8 = -126
    switch 0 {
    case 618:
        var v17 float64 = 0.0
        _ = v17
    }
    fmt.Println(v3, v1, v16)
    var v18 string = "foo bar"
    var v19 int16 = 32767
    v20 := []int{9223372036854775807, int(v1), 328}
    switch v18 {
    case "foo bar":
        var v21 int64 = 7
        _ = v21
    }
    switch ((v18 + ("a" + v18)) + v18) {
    case "a":
        if true {
            if (!false) {
                v22 := GenS1{F0: 2147483647, F1: (v19 ^ (v19 + (v19 | v19)))}
                var v23 uint8 = 127
                fmt.Println(v1, v23, v18)
                _ = v22
                _ = v23
            } else {
                fmt.Println(v1)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
