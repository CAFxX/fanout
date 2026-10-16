package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint8 }
type GenS2 struct { F0 uint32; F1 string; F2 int64 }
type GenS3 struct { F0 uint64; F1 int16 }

func main() {
    var v1 int64 = 42
    var v2 uint64 = 2
    var v3 string = ""
    v4 := make(map[uint32]int)
    delete(v4, 42)
    for fi5 := 0; fi5 < 5; fi5++ {
        if true {
            switch "" {
            case "":
                var v6 float32 = float32((3.14159 - 1.5))
                _ = v6
            }
            var v7 bool = false
            _ = v7
        }
        _ = fi5
    }
    if false {
        fmt.Println(v3)
        switch 128 {
        case 42:
            _ = len(v4)
        case 9223372036854775807:
            var v8 uint = 18446744073709551615
            _ = v8
        case 46:
            if false {
                var v9 string = v3
                v10 := GenS1{F0: v1, F1: 32768, F2: 42}
                _ = v10
                _ = v9
            } else {
                v11 := []int{-2147483648, -9223372036854775808, 4294967295}
                _ = v11
            }
        default:
            fmt.Println(v1, v2, v3)
        }
        v2 = (v2 | (^(v2)))
    } else {
        v2 = v2
        switch "" {
        case "foo bar":
            switch (v3 + "") {
            case "hello":
                v12 := []int{314, 2147483647, 9223372036854775807}
                _ = v12
            default:
                for fi13 := 0; fi13 < 5; fi13++ {
                    v14 := GenS1{F0: v1, F1: fi13, F2: 42}
                    _ = fi13
                    _ = v14
                }
            }
        case "\n\t":
            if false {
                v15 := []int{}
                v16 := GenS1{F0: (v1 + (v1 << ((23) & 63))), F1: -1, F2: 96}
                _ = v15
                _ = v16
            }
        case "hello":
            for fi17 := 0; fi17 < 12; fi17++ {
                delete(v4, 2)
                var v18 float64 = -2.25
                v1 = 9223372036854775807
                _ = fi17
                _ = v18
            }
        }
    }
    if true {
        var v19 int = 1
        _ = v19
    } else {
        switch "a" {
        case "hello":
            v1 = (v1 + v1)
        case "":
            if false {
                var v20 uint64 = (v2 / (((v2 * (v2 + v2))) | 1))
                v21 := GenS2{F0: 4294967295, F1: "foo bar", F2: (v1 & (v1 % (((^(v1))) | 1)))}
                if (2 == 234) {
                    fmt.Println(v3, v2)
                }
                _ = v20
                _ = v21
            } else {
                var v22 string = "\n\t"
                v4[2147483647] = 0
                v23 := GenS3{F0: (v2 % ((1) | 1)), F1: 42}
                _ = v22
                _ = v23
            }
        }
        for fi24 := 0; fi24 < 2; fi24++ {
            var v25 uint64 = (^(v2))
            _ = fi24
            _ = v25
        }
    }
    var v26 int8 = 1
    for fi27 := 0; fi27 < 7; fi27++ {
        var v28 string = v3
        _ = fi27
        _ = v28
    }
    fmt.Println(v26)
    v1 -= v1
    v26 = (v26 / ((v26) | 1))
    v1 -= v1
    var v29 int16 = 32767
    v29 += (v29 * v29)
    switch "hello" {
    case "\n\t":
        switch -1 {
        case 1:
            var v30 uint32 = 0
            _ = v30
        case 9223372036854775807:
            _ = len(v4)
        }
    case "hello":
        var v31 float32 = 1.5
        _ = v31
    case "":
        switch v3 {
        case "\n\t":
            for fi32 := 0; fi32 < 20; fi32++ {
                var v33 bool = true
                _ = fi32
                _ = v33
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
}
