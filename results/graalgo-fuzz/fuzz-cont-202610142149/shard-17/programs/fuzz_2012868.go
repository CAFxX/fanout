package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int }
type GenS2 struct { F0 int16; F1 uint32; F2 string }
type GenS3 struct { F0 string; F1 string; F2 uint16 }

func main() {
    var v1 int64 = -347
    var v2 float32 = 0.0
    var v3 bool = (933 < 489)
    v4 := []uint32{114, 65535, 2147483648}
    v2 += (v2 + (v2 + v2))
    for fi5 := 0; fi5 < 19; fi5++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi5
    }
    var v6 float32 = (-2.25 / (((v2 - (v2 / 1e10)) + v2) * v2))
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int32 = -2147483648
        v9 := GenS1{F0: 100, F1: false, F2: (fi7 / (((fi7 | fi7)) | 1))}
        fmt.Println(fi7)
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint16 = 32767
    if false {
        for fi11 := 0; fi11 < 7; fi11++ {
            v12 := GenS1{F0: 21, F1: ((v1 * v1) == 499), F2: fi11}
            v13 := GenS2{F0: 1, F1: 389, F2: ((("a" + "\n\t") + "") + "foo bar")}
            v14 := GenS3{F0: ("" + ("a" + "a")), F1: "hello", F2: 0}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
    } else {
        switch "hello" {
        case "hello":
            var v15 uint = 18446744073709551615
            _ = v15
        case "foo bar":
            var v16 string = ("\n\t" + "foo bar")
            _ = v16
        }
        v17 := GenS2{F0: 42, F1: 42, F2: "foo bar"}
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        _ = v17
    }
    v3 = true
    if ("a" <= "a") {
        var v18 float32 = 1.5
        v10 = (v10 * v10)
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        v6 = 0.0
        v2 = 0.1
        _ = fi19
    }
    fmt.Println(v10, v6)
    v2 += v2
    fmt.Println(v2)
    var v20 uint = 42
    v21 := GenS2{F0: -229, F1: 490, F2: "hello"}
    var v22 string = ("foo bar" + "")
    fmt.Println(v10, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v6
}
