package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int }

func main() {
    var v1 uint64 = 0
    var v2 int16 = 1
    var v3 int = 178
    v4 := make(map[string]bool)
    fmt.Println(v3)
    if ((v3 | (v3 >> ((42) & 63))) > v3) {
        switch (v3 ^ 646) {
        case 2147483647:
            for fi5 := 0; fi5 < 2; fi5++ {
                v6 := GenS1{F0: ("foo bar" + "foo bar")}
                switch "a" {
                case "foo bar":
                    var v7 uint64 = 4294967295
                    _ = v7
                case "hello":
                    v1 = (^(v1))
                case "\n\t":
                    switch "\n\t" {
                    case "\n\t":
                        if false {
                            v2 = (v2 >> ((1) & 63))
                        }
                    case "a":
                        var v8 uint32 = 4294967295
                        _ = v8
                    }
                }
                _ = fi5
                _ = v6
            }
        default:
            v2 -= v2
        }
    } else {
        fmt.Println(v3, v3, v2)
        var v9 float64 = 0.1
        _ = v9
    }
    var v10 float64 = 0.1
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS1{F0: "\n\t"}
        var v13 int64 = -9223372036854775808
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := []int{v3}
    v2 = (-(v2))
    switch "\n\t" {
    case "foo bar":
        var v15 string = "hello"
        _ = v15
    default:
        for fi16 := 0; fi16 < 2; fi16++ {
            var v17 int16 = v2
            _ = fi16
            _ = v17
        }
    }
    fmt.Println(v1)
    v1 = v1
    v1 = v1
    if true {
        var v18 float64 = 1.5
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        for fi20 := 0; fi20 < 5; fi20++ {
            _ = len(v14)
            v2 = v2
            var v21 int64 = 0
            _ = fi20
            _ = v21
        }
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        if false {
            switch (^(v3)) {
            case -1:
                switch "" {
                case "\n\t":
                    for fi22 := 0; fi22 < 13; fi22++ {
                        var v23 int64 = 9223372036854775807
                        v24 := GenS1{F0: "a"}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                case "a":
                    _ = len(v4)
                default:
                    v25 := GenS1{F0: ("" + "foo bar")}
                    _ = v25
                }
            case 127:
                for fi26 := 0; fi26 < 1; fi26++ {
                    var v27 int64 = 272
                    v27 = (v27 ^ v27)
                    var v28 uint = 1
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            case 154:
                v2 = (v2 % ((42) | 1))
            }
            var v29 int64 = 9223372036854775807
            v30 := GenS1{F0: "\n\t"}
            _ = v29
            _ = v30
        } else {
            if (65535 >= 255) {
                var v31 int = (fi19 << ((42) & 63))
                _ = v31
            }
            switch "a" {
            case "":
                if true {
                    fmt.Println(v2)
                    var v32 float32 = 1e10
                    _ = v32
                }
            case "hello":
                v10 = float64(0.0)
            }
        }
        _ = fi19
    }
    if (((v2 / ((-680) | 1)) != (^(v2))) && true) {
        var v33 int = (v3 - v3)
        for fi34 := 0; fi34 < 7; fi34++ {
            delete(v4, "")
            if false {
                v35 := GenS2{F0: true, F1: 4}
                _ = v35
            }
            _ = fi34
        }
        _ = v33
    } else {
        var v36 int32 = -2147483648
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = v36
    }
    for fi37 := 0; fi37 < 16; fi37++ {
        fmt.Println(v2)
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v2)
        case "a":
            v10 -= (v10 / v10)
        case "hello":
            if true {
                var v38 int8 = 1
                v10 += float64(3.14159)
                _ = v38
            }
        default:
            var v39 int = -1
            _ = v39
        }
        _ = fi37
    }
    v40 := GenS3{F0: (^(v3))}
    var v41 float64 = 0.0
    v14 = append(v14, (v3 % ((1) | 1)))
    for fi42 := 0; fi42 < 9; fi42++ {
        v3 = (^(v3))
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
}
