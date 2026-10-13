package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64 }

func main() {
    var v1 uint8 = 0
    v2 := make(map[uint32]bool)
    var v3 int = 0
    v4 := make(map[uint32]string)
    v5 := make(map[string]int)
    var v6 int16 = 1000
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS1{F0: 0, F1: -550, F2: (v1 - (^(v1)))}
        _ = fi7
        _ = v8
    }
    v1 = v1
    v6 = 1000
    v3 += 127
    switch "\n\t" {
    case "":
        var v9 int32 = 0
        _ = v9
    case "a":
        var v10 uint = 255
        _ = v10
    default:
        for fi11 := 0; fi11 < 5; fi11++ {
            v12 := []int{(fi11 >> ((1) & 63)), fi11}
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 uint8 = v1
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    }
    if true {
        var v15 uint = 18446744073709551615
        if ((((("hello" + "") + "foo bar") + "hello") > "hello") && (("\n\t" > "\n\t") || ((v3 * (v3 % ((v3) | 1))) <= (v3 << ((v15) & 63))))) {
            v4[42] = ("" + "hello")
            var v16 int64 = -1
            _ = v16
        }
        _ = v15
    }
    if false {
        var v17 float32 = 3.14159
        v18 := GenS2{F0: 507, F1: (!false)}
        _ = v17
        _ = v18
    }
    fmt.Println(v3, v3)
    for fi19 := 0; fi19 < 11; fi19++ {
        v20 := GenS3{F0: ("" + ("a" + "\n\t")), F1: 707}
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 string = "a"
            _ = fi21
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 uint = 42
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    var v26 uint32 = 4294967295
    v27 := GenS4{F0: int(v1)}
    for fi28 := 0; fi28 < 2; fi28++ {
        for fi29 := 0; fi29 < 12; fi29++ {
            v30 := []int{v3}
            var v31 uint = 42
            fmt.Println(fi28, v31)
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi28
    }
    v6 += v6
    if false {
        var v32 int16 = v6
        switch ("hello" + ("hello" + "a")) {
        case "":
            v33 := GenS1{F0: 0, F1: 42, F2: (v1 * (v1 | 255))}
            _ = v33
        case "hello":
            if ((v1 - (v1 + v1)) != 255) {
                switch "" {
                case "hello":
                    var v34 float64 = 1.5
                    _ = v34
                default:
                    var v35 int8 = 23
                    _ = v35
                }
                if true {
                    var v36 uint32 = (v26 + 2147483647)
                    _ = v36
                }
            } else {
                delete(v4, (^(v26)))
            }
        default:
            fmt.Println(v26, v26)
        }
        if false {
            v3 = -9223372036854775808
            if false {
                v6 -= (v32 | (-(v6)))
            } else {
                if true {
                    v6 = (v6 >> ((42) & 63))
                    v6 -= -194
                }
            }
        }
        _ = v32
    }
    v37 := GenS5{F0: -9223372036854775808}
    for fi38 := 0; fi38 < 16; fi38++ {
        v26 = v26
        switch "hello" {
        case "\n\t":
            v39 := GenS4{F0: 32768}
            _ = v39
        }
        v40 := GenS3{F0: "", F1: 42}
        _ = fi38
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
