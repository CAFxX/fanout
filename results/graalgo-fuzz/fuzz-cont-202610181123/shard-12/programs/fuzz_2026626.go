package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint16; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 int64 = 422
    var v2 int16 = -32768
    var v3 uint32 = 2
    v4 := []string{"", (("a" + "") + "")}
    var v5 uint = 331
    fmt.Println(v3)
    v1 = -966
    switch "foo bar" {
    case "\n\t":
        if ("\n\t" != "hello") {
            v6 := GenS1{F0: 0}
            _ = len(v4)
            fmt.Println(v2)
            _ = v6
        } else {
            var v7 int = -1
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch "hello" {
            case "a":
                _ = len(v4)
            case "\n\t":
                for fi8 := 0; fi8 < 18; fi8++ {
                    var v9 uint64 = 2
                    var v10 uint = v5
                    v7 = (^(fi8))
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            case "foo bar":
                var v11 int16 = v2
                _ = v11
            }
            _ = v7
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch ("\n\t" + "a") {
    case "foo bar":
        v12 := GenS1{F0: 9223372036854775807}
        _ = v12
    case "hello":
        fmt.Println(v3, v5)
    default:
        v5 = 18446744073709551615
    }
    if (false && (v5 >= (v5 * (v5 | (^(v5)))))) {
        v1 += 9223372036854775807
        var v13 float32 = -2.25
        v4 = append(v4, (((("hello" + "") + "a") + "") + ""))
        _ = v13
    } else {
        for fi14 := 0; fi14 < 20; fi14++ {
            switch "foo bar" {
            case "":
                v2 = (-(v2))
            case "foo bar":
                var v15 int64 = -9223372036854775808
                _ = v15
            case "\n\t":
                var v16 float32 = 0.1
                _ = v16
            }
            _ = fi14
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            if len(v4) > 0 {
                v4[0] = "hello"
            }
            _ = fi17
        }
        var v18 uint32 = (v3 >> ((2) & 63))
        _ = v18
    }
    v4 = append(v4, "\n\t")
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    switch 1 {
    case 4294967295:
        fmt.Println(v1)
    case -1:
        if (v1 <= (^(v1))) {
            var v19 int32 = 0
            for fi20 := 0; fi20 < 19; fi20++ {
                var v21 int64 = (v1 << (((^(v5))) & 63))
                _ = fi20
                _ = v21
            }
            _ = v19
        } else {
            var v22 float32 = 1e10
            v23 := GenS1{F0: -9223372036854775808}
            _ = v22
            _ = v23
        }
    case 296:
        if false {
            fmt.Println(v2, v1, v1)
            v24 := GenS2{F0: 503, F1: 42, F2: ("\n\t" + "hello")}
            _ = v24
        }
    }
    v3 = v3
    v25 := GenS1{F0: 1}
    switch "hello" {
    case "hello":
        v26 := GenS1{F0: -9223372036854775808}
        _ = v26
    case "":
        if false {
            if (!false) {
                var v27 bool = true
                var v28 uint8 = 18
                v28 = (v28 % ((v28) | 1))
                _ = v27
                _ = v28
            } else {
                v5 += v5
                v2 -= (^(v2))
            }
            v5 = (^(v5))
        } else {
            v29 := GenS3{F0: 127}
            _ = v29
        }
    case "\n\t":
        var v30 uint32 = 42
        _ = v30
    }
    for fi31 := 0; fi31 < 19; fi31++ {
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        _ = fi31
    }
    var v32 float64 = ((3.14159 * float64(1e10)) / 1.0)
    fmt.Println(v32)
    v33 := GenS3{F0: 127}
    if true {
        fmt.Println(v32, v1, v2)
    } else {
        switch ("hello" + "hello") {
        case "hello":
            v34 := GenS4{F0: v2}
            _ = v34
        default:
            v2 -= v2
        }
        var v35 float32 = (0.0 - 3.14159)
        _ = v35
    }
    if false {
        v2 -= int16(v1)
        var v36 uint32 = (v3 & v3)
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
