package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 int8 = -29
    var v2 float32 = 1.5
    v3 := make(map[string]int)
    var v4 bool = (1 <= 4294967295)
    v5 := make(map[uint32]int)
    switch "\n\t" {
    case "hello":
        var v6 int16 = -671
        _ = v6
    case "\n\t":
        v1 = (v1 >> ((2) & 63))
    default:
        v7 := GenS1{F0: "hello"}
        _ = v7
    }
    fmt.Println(v2, v4)
    delete(v3, ("\n\t" + ""))
    fmt.Println(v1, v4)
    v8 := GenS1{F0: "\n\t"}
    fmt.Println(v4, v1, v2)
    v4 = ((!(0.1 > 1.5)) && true)
    if false {
        for fi9 := 0; fi9 < 9; fi9++ {
            switch fi9 {
            case -71:
                var v10 int = fi9
                _ = v10
            case -9223372036854775808:
                switch (fi9 * (fi9 % (((fi9 ^ (fi9 + fi9))) | 1))) {
                case 4294967295:
                    switch (fi9 ^ 127) {
                    case 0:
                        for fi11 := 0; fi11 < 3; fi11++ {
                            v12 := []int{128, (fi9 ^ (^(fi9))), -1}
                            v13 := GenS1{F0: ("" + ("a" + ""))}
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    }
                }
            }
            v14 := GenS1{F0: "foo bar"}
            _ = fi9
            _ = v14
        }
        if v4 {
            v15 := GenS2{F0: -709, F1: (true && (!((!v4) || v4)))}
            _ = v15
        } else {
            for fi16 := 0; fi16 < 17; fi16++ {
                v1 -= (v1 ^ (v1 + (-(v1))))
                v2 += 0.0
                _ = len(v3)
                _ = fi16
            }
        }
        if (v4 || (((v2 * (0.0 + 0.0)) / v2) < v2)) {
            v17 := GenS2{F0: 480, F1: (!v4)}
            _ = v17
        } else {
            v18 := GenS2{F0: -1, F1: ((0 >= -1) || v4)}
            for fi19 := 0; fi19 < 14; fi19++ {
                if (-2.25 >= v2) {
                    var v20 int32 = 42
                    _ = v20
                }
                var v21 int = fi19
                v1 = (v1 + (v1 + (v1 / ((v1) | 1))))
                _ = fi19
                _ = v21
            }
            var v22 int16 = 0
            _ = v18
            _ = v22
        }
    }
    v23 := []int{2147483647, 32767, 0, 32767}
    var v24 uint = 255
    var v25 string = "hello"
    v26 := GenS1{F0: v25}
    switch 128 {
    case 0:
        v4 = true
    case -988:
        fmt.Println(v2, v4, v4)
    }
    fmt.Println(v2)
    if v4 {
        var v27 uint8 = 31
        v28 := GenS3{F0: v25, F1: (((2 >= 65535) && (("" + (((("" + (("foo bar" + "") + (v25 + v25))) + "hello") + ("hello" + v25)) + v25)) <= ("\n\t" + "hello"))) && false)}
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 uint = v24
            switch fi29 {
            case 4294967295:
                var v31 uint64 = 418
                _ = v31
            default:
                v32 := GenS4{F0: 2}
                _ = v32
            }
            _ = fi29
            _ = v30
        }
        _ = v27
        _ = v28
    }
    v1 = v1
    if len(v23) > 0 {
        v23[0] = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
