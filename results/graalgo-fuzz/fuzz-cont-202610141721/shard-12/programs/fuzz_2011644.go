package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 uint64 = 4294967295
    var v2 int = 0
    var v3 int32 = -2147483648
    var v4 uint16 = 2
    v5 := []int{}
    if (596 > 10) {
        if true {
            var v6 int64 = -451
            if true {
                var v7 int16 = -91
                _ = v7
            }
            _ = v6
        }
        if false {
            var v8 int = v2
            _ = v8
        }
    } else {
        fmt.Println(v1)
        if (!true) {
            switch "hello" {
            case "hello":
                v9 := GenS1{F0: 100, F1: "hello"}
                _ = v9
            case "a":
                for fi10 := 0; fi10 < 14; fi10++ {
                    v2 = v2
                    _ = fi10
                }
            case "\n\t":
                var v11 float32 = float32((1e10 - 0.1))
                _ = v11
            }
        } else {
            var v12 string = "foo bar"
            if false {
                var v13 string = "hello"
                for fi14 := 0; fi14 < 13; fi14++ {
                    v15 := GenS1{F0: 0, F1: (v13 + "\n\t")}
                    _ = fi14
                    _ = v15
                }
                _ = v13
            } else {
                _ = len(v5)
                v12 = "foo bar"
            }
            switch (v2 >> ((675) & 63)) {
            case -801:
                var v16 int64 = -371
                _ = v16
            case 2147483647:
                v3 += v3
            default:
                var v17 string = "hello"
                _ = v17
            }
            _ = v12
        }
    }
    v1 = uint64(v4)
    var v18 uint32 = 1
    v19 := GenS2{F0: (v2 >> ((42) & 63)), F1: v18}
    var v20 uint16 = 0
    var v21 uint64 = (v1 % (((v1 + v1)) | 1))
    for fi22 := 0; fi22 < 15; fi22++ {
        if true {
            switch fi22 {
            case 132:
                fmt.Println(v1, v3, v3)
            case 32767:
                v18 = (^(v18))
            }
        } else {
            if false {
                switch "\n\t" {
                case "a":
                    if true {
                        var v23 float64 = 3.14159
                        v23 = (((v23 / v23) / (3.14159 - float64(0.0))) * v23)
                        _ = v23
                    }
                case "foo bar":
                    v24 := GenS2{F0: (v2 >> ((1) & 63)), F1: v18}
                    _ = v24
                case "\n\t":
                    v4 = v20
                }
                for fi25 := 0; fi25 < 16; fi25++ {
                    _ = len(v5)
                    fmt.Println(v2, v3)
                    _ = fi25
                }
            } else {
                if true {
                    var v26 uint16 = v4
                    _ = v26
                }
            }
            v27 := GenS1{F0: 127, F1: "a"}
            _ = v27
        }
        var v28 uint = 5
        var v29 float32 = 0.0
        _ = fi22
        _ = v28
        _ = v29
    }
    switch "hello" {
    case "foo bar":
        v30 := GenS1{F0: -128, F1: "\n\t"}
        _ = v30
    default:
        switch 9223372036854775807 {
        case -737:
            v2 = (v2 << ((701) & 63))
        case 128:
            v31 := GenS1{F0: 108, F1: "\n\t"}
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 19; fi32++ {
        var v33 int = (-(v2))
        _ = fi32
        _ = v33
    }
    var v34 int16 = 356
    v18 += 4294967295
    v35 := GenS3{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
