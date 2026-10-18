package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint32; F1 uint16; F2 int8 }

func main() {
    v1 := []uint32{0, 0, 2147483648, 0}
    v2 := make(map[string]string)
    var v3 float64 = (((((float64(0.1) / 1.0) + 0.0) + 1e10) / 1.0) / (3.14159 / 1.0))
    v4 := []string{}
    v5 := []string{("\n\t" + (("\n\t" + "\n\t") + ""))}
    var v6 uint16 = 42
    var v7 int8 = -128
    switch -426 {
    case 9223372036854775807:
        var v8 int16 = 32767
        _ = v8
    case 464:
        v9 := GenS1{F0: "\n\t", F1: 1}
        _ = v9
    }
    if false {
        var v10 uint8 = 1
        if false {
            var v11 uint = 4294967295
            v12 := GenS1{F0: "hello", F1: 1}
            var v13 uint32 = uint32(v10)
            _ = v11
            _ = v12
            _ = v13
        } else {
            v14 := GenS1{F0: "a", F1: 4294967295}
            var v15 uint32 = 495
            _ = v14
            _ = v15
        }
        var v16 uint16 = 32767
        _ = v10
        _ = v16
    }
    v17 := GenS1{F0: ("a" + "foo bar"), F1: 200}
    var v18 bool = false
    switch 2147483647 {
    case 127:
        if ((-32768 != -32768) && v18) {
            fmt.Println(v6, v3)
            switch ("foo bar" + "a") {
            case "a":
                for fi19 := 0; fi19 < 17; fi19++ {
                    var v20 uint32 = 2147483647
                    _ = fi19
                    _ = v20
                }
            }
        } else {
            switch 32767 {
            case -1:
                var v21 uint = 847
                _ = v21
            default:
                if false {
                    var v22 float32 = (1.5 + -2.25)
                    _ = v22
                }
            }
        }
    case 4294967295:
        var v23 uint64 = 2
        _ = v23
    default:
        for fi24 := 0; fi24 < 19; fi24++ {
            _ = len(v2)
            fmt.Println(v7)
            _ = fi24
        }
    }
    v25 := GenS2{F0: 18446744073709551615, F1: 2}
    v26 := GenS3{F0: (v7 % (((v7 << ((2) & 63))) | 1)), F1: "foo bar"}
    switch ("foo bar" + "\n\t") {
    case "":
        for fi27 := 0; fi27 < 3; fi27++ {
            if v18 {
                var v28 float64 = -2.25
                v29 := GenS4{F0: 65535, F1: (v6 + 113), F2: v7}
                v30 := GenS3{F0: v7, F1: "a"}
                _ = v28
                _ = v29
                _ = v30
            } else {
                fmt.Println(v6, v6)
                var v31 bool = v18
                for fi32 := 0; fi32 < 19; fi32++ {
                    var v33 string = "\n\t"
                    _ = fi32
                    _ = v33
                }
                _ = v31
            }
            fmt.Println(v7, v7, v18)
            switch "a" {
            case "hello":
                v7 += 51
            case "\n\t":
                if v18 {
                    var v34 uint8 = 127
                    _ = v34
                }
            case "a":
                fmt.Println(v6, v3)
            }
            _ = fi27
        }
    case "\n\t":
        var v35 string = (("" + "") + ("a" + ""))
        _ = v35
    case "foo bar":
        var v36 int32 = -1
        _ = v36
    }
    v5 = append(v5, "\n\t")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
