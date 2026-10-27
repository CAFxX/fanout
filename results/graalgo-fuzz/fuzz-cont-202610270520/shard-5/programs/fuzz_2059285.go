package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64; F1 int64; F2 int16 }

func main() {
    var v1 float64 = 3.14159
    var v2 bool = true
    var v3 int64 = 0
    v4 := make(map[string]string)
    switch 127 {
    case 127:
        fmt.Println(v1, v2)
    case 1:
        switch "\n\t" {
        case "hello":
            fmt.Println(v3)
        case "\n\t":
            var v5 string = ("foo bar" + "\n\t")
            _ = v5
        }
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: 1000, F1: 255, F2: 304}
        if (false || (-2.25 != (1.5 - 1e10))) {
            for fi8 := 0; fi8 < 14; fi8++ {
                switch (fi8 + 992) {
                case 0:
                    v3 += -1
                case -2147483648:
                    v9 := GenS1{F0: -58, F1: 1, F2: 1}
                    _ = v9
                }
                _ = fi8
            }
        } else {
            v10 := []int{42, (fi6 & (-(fi6))), fi6, -958}
            _ = v10
        }
        v1 += v1
        _ = fi6
        _ = v7
    }
    switch 127 {
    case 629:
        v11 := v4[("hello" + ("a" + ((("a" + "hello") + ("\n\t" + "")) + "\n\t")))]
        _ = v11
    case 0:
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := v4["a"]
            v3 += (v3 % ((-637) | 1))
            _ = fi12
            _ = v13
        }
    case -268:
        v14 := []int{}
        _ = v14
    }
    fmt.Println(v1, v3)
    switch 4294967295 {
    case -51:
        for fi15 := 0; fi15 < 13; fi15++ {
            switch "hello" {
            case "hello":
                v16 := GenS2{F0: fi15}
                _ = v16
            case "foo bar":
                for fi17 := 0; fi17 < 11; fi17++ {
                    v1 += v1
                    v18 := GenS3{F0: 65535, F1: 542, F2: 2}
                    v1 += 1.5
                    _ = fi17
                    _ = v18
                }
            case "\n\t":
                for fi19 := 0; fi19 < 14; fi19++ {
                    if (v2 && ("foo bar" > "")) {
                        v20 := v4[""]
                        var v21 uint16 = 2
                        v4[v20] = (v20 + (v20 + "a"))
                        v1 = v1
                        _ = v20
                        _ = v21
                    } else {
                        if v2 {
                            v22 := []int{}
                            _ = v22
                        } else {
                            var v23 int8 = 97
                            v24 := GenS4{F0: 599, F1: 42}
                            var v25 uint16 = 42
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        delete(v4, ((("\n\t" + "hello") + ("a" + "")) + "hello"))
                        delete(v4, ("a" + "hello"))
                        v26 := GenS5{F0: 255}
                        v2 = true
                        _ = v26
                    }
                    _ = fi19
                }
            }
            _ = fi15
        }
    }
    v27 := []int{0}
    switch -1 {
    case -170:
        var v28 uint = 42
        _ = v28
    case 4294967295:
        for fi29 := 0; fi29 < 2; fi29++ {
            switch ("" + "") {
            case "":
                v30 := GenS6{F0: 983, F1: v3, F2: 32767}
                _ = v30
            case "hello":
                var v31 string = "\n\t"
                _ = v31
            }
            _ = fi29
        }
    default:
        var v32 uint8 = 90
        _ = v32
    }
    v2 = true
    v33 := GenS6{F0: 2, F1: 1, F2: 148}
    var v34 string = ("\n\t" + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
