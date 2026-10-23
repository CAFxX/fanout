package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 bool; F1 int64 }

func main() {
    var v1 float64 = 3.14159
    v2 := []uint32{65535, 1}
    var v3 bool = (168 >= -1)
    var v4 uint32 = 0
    var v5 int = 0
    var v6 int = (v5 % ((v5) | 1))
    switch (v5 * v5) {
    case 786:
        for fi7 := 0; fi7 < 13; fi7++ {
            v6 -= (fi7 - int(v4))
            v3 = (v3 || (!v3))
            _ = fi7
        }
    case 32768:
        var v8 uint16 = 65535
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        v1 -= (1e10 / -2.25)
        for fi10 := 0; fi10 < 18; fi10++ {
            if false {
                v11 := GenS1{F0: 0, F1: false}
                switch (-(fi9)) {
                case 824:
                    v3 = (v3 || (0 > 1))
                case -513:
                    v12 := GenS2{F0: v4}
                    _ = v12
                }
                _ = v11
            } else {
                if (!false) {
                    v5 = fi10
                }
                if v3 {
                    v1 += 0.1
                    if (v3 && true) {
                        v13 := GenS3{F0: 30, F1: 2}
                        v14 := GenS1{F0: 908, F1: (127 != 1)}
                        v2 = append(v2, v4)
                        _ = v13
                        _ = v14
                    } else {
                        fmt.Println(v3)
                        fmt.Println(v5, fi9, v4)
                        fmt.Println(v3)
                    }
                    v15 := GenS4{F0: v3, F1: 559}
                    fmt.Println(v5, v3, fi9)
                    v2 = append(v2, 4294967295)
                    _ = v15
                }
                v6 += 127
                fmt.Println(v3, fi9, v1)
                var v16 string = "hello"
                _ = v16
            }
            var v17 int64 = 9223372036854775807
            var v18 uint32 = (v4 * v4)
            _ = fi10
            _ = v17
            _ = v18
        }
        _ = fi9
    }
    var v19 int = 128
    for fi20 := 0; fi20 < 10; fi20++ {
        fmt.Println(v4, fi20)
        _ = fi20
    }
    v21 := GenS3{F0: 1, F1: 2}
    switch "hello" {
    case "foo bar":
        fmt.Println(v1, v19, v6)
    case "":
        v4 = v4
    case "\n\t":
        v2 = append(v2, v4)
    }
    if false {
        if true {
            v6 = 32767
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        switch "hello" {
        case "\n\t":
            if v3 {
                for fi22 := 0; fi22 < 8; fi22++ {
                    var v23 uint32 = v4
                    v23 = 65535
                    _ = fi22
                    _ = v23
                }
                var v24 uint = 42
                var v25 int64 = 42
                _ = v24
                _ = v25
            }
        case "hello":
            var v26 float32 = float32(1.5)
            _ = v26
        case "":
            var v27 float64 = (v1 * 1.5)
            _ = v27
        }
    }
    var v28 uint16 = 42
    switch "" {
    case "":
        for fi29 := 0; fi29 < 16; fi29++ {
            var v30 uint64 = 1
            _ = fi29
            _ = v30
        }
    case "\n\t":
        v5 += (v19 * v19)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
