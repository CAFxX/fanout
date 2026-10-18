package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int32 }

func main() {
    var v1 uint32 = 4294967295
    var v2 int64 = -9223372036854775808
    v3 := make(map[string]bool)
    var v4 bool = (!(("hello" + ("\n\t" + "")) == "foo bar"))
    var v5 string = "foo bar"
    v1 += v1
    if (!false) {
        var v6 bool = (true || (42 == 59))
        v2 += v2
        switch 4294967295 {
        case -888:
            switch 32767 {
            case -361:
                v7 := []int{992, 1, -1, -883}
                _ = v7
            }
        case 474:
            fmt.Println(v6)
        case 933:
            for fi8 := 0; fi8 < 11; fi8++ {
                switch (fi8 + (-(fi8))) {
                case 32768:
                    var v9 float64 = 3.14159
                    _ = v9
                case 9223372036854775807:
                    var v10 float32 = (3.14159 - -2.25)
                    _ = v10
                case 1:
                    for fi11 := 0; fi11 < 16; fi11++ {
                        if v4 {
                            v1 = 65535
                            var v12 uint16 = 32767
                            var v13 int16 = 267
                            _ = v12
                            _ = v13
                        }
                        _ = fi11
                    }
                }
                for fi14 := 0; fi14 < 6; fi14++ {
                    var v15 string = v5
                    switch fi14 {
                    case 32767:
                        var v16 uint32 = (v1 >> ((42) & 63))
                        _ = v16
                    }
                    var v17 int16 = 0
                    _ = fi14
                    _ = v15
                    _ = v17
                }
                v18 := GenS1{F0: 100}
                _ = fi8
                _ = v18
            }
        default:
            var v19 int32 = -2147483648
            _ = v19
        }
        _ = v6
    } else {
        switch v5 {
        case "a":
            v1 = v1
        case "hello":
            var v20 float64 = 3.14159
            _ = v20
        default:
            var v21 int16 = -1
            _ = v21
        }
        var v22 bool = ((v4 && ((v1 | v1) < v1)) || (!(v1 <= (v1 / ((42) | 1)))))
        _ = v22
    }
    v3[("" + ("a" + "foo bar"))] = (42 >= -421)
    v1 -= 1
    switch v5 {
    case "a":
        var v23 string = v5
        _ = v23
    case "\n\t":
        v24 := GenS2{F0: 42}
        _ = v24
    }
    if true {
        var v25 float32 = 0.0
        var v26 uint = 0
        _ = v25
        _ = v26
    }
    var v27 uint8 = 0
    if false {
        v27 = 255
        var v28 int8 = -116
        _ = v28
    }
    v4 = v4
    v4 = v4
    v1 = 42
    v4 = false
    v1 = (v1 | (^(v1)))
    if (0 != 2) {
        v27 -= v27
        v29 := GenS3{F0: 0, F1: true}
        v30 := GenS4{F0: 0}
        _ = v29
        _ = v30
    } else {
        v1 = 773
    }
    var v31 bool = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
