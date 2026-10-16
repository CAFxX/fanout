package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint16; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int8) int32 {
    v5 := []int{725, 322}
    if false {
        p0 += 42
        var v6 bool = false
        if v6 {
            p0 += -118
        }
        _ = v6
    }
    _ = len(v5)
    if false {
        fmt.Println(p0)
    } else {
        p0 = (p0 | p0)
        switch "a" {
        case "foo bar":
            v7 := GenS1{F0: 255}
            _ = v7
        case "a":
            if false {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        case "\n\t":
            fmt.Println(p0)
        default:
            for fi8 := 0; fi8 < 13; fi8++ {
                fmt.Println(p0, fi8)
                _ = fi8
            }
        }
    }
    v9 := GenS1{F0: 2}
    _ = p0
    _ = v5
    _ = v9
    return -2147483648
}

func main() {
    var v1 int8 = -1
    var v2 float64 = 0.1
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    var v10 uint16 = 42
    v3[42] = "a"
    v2 = (3.14159 / 1.5)
    switch 42 {
    case -160:
        var v11 uint = 42
        _ = v11
    }
    v12 := GenS1{F0: 65535}
    switch "foo bar" {
    case "\n\t":
        if (0 <= 1) {
            fmt.Println(v2)
        } else {
            v3[42] = "a"
        }
    case "hello":
        v13 := GenS2{F0: v10, F1: 255, F2: 32767}
        _ = v13
    }
    v14 := GenS1{F0: 255}
    v15 := GenS3{F0: ("" + ("foo bar" + "hello"))}
    if true {
        for fi16 := 0; fi16 < 20; fi16++ {
            switch (("a" + "hello") + "a") {
            case "foo bar":
                for fi17 := 0; fi17 < 1; fi17++ {
                    var v18 uint16 = (v10 + (^(v10)))
                    var v19 int32 = 721
                    if false {
                        v18 -= (v18 * (v18 ^ (v18 - (^(v10)))))
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            case "a":
                for fi20 := 0; fi20 < 5; fi20++ {
                    v21 := GenS4{F0: -1, F1: "\n\t"}
                    _ = fi20
                    _ = v21
                }
            case "hello":
                var v22 uint64 = 0
                _ = v22
            }
            v10 -= v10
            fmt.Println(v2)
            _ = fi16
        }
        var v23 uint8 = 128
        _ = v23
    } else {
        for fi24 := 0; fi24 < 4; fi24++ {
            delete(v3, (fi24 >> ((255) & 63)))
            _ = fi24
        }
        if ((v1 == v1) && ("hello" == "a")) {
            v1 = v1
            fmt.Println(v2, v1, v2)
        } else {
            if true {
                delete(v4, 4294967295)
                var v25 float64 = (((v2 - v2) * (v2 + v2)) * v2)
                for fi26 := 0; fi26 < 3; fi26++ {
                    v27 := GenS5{F0: 355}
                    for fi28 := 0; fi28 < 17; fi28++ {
                        switch (fi28 & fi28) {
                        case 9223372036854775807:
                            fmt.Println(fi26, v10, v2)
                            v10 += (v10 ^ (v10 ^ (^(v10))))
                        }
                        _ = fi28
                    }
                    fmt.Println(v2)
                    _ = fi26
                    _ = v27
                }
                _ = v25
            }
            fmt.Println(v1)
        }
        switch 32767 {
        case 0:
            v2 -= -2.25
        case -477:
            var v29 float64 = (3.14159 + (-2.25 + float64((3.14159 / 1.0))))
            _ = v29
        case 465:
            var v30 float32 = 1.5
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
