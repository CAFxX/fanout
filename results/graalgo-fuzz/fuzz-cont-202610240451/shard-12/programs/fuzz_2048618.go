package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 int32 }
func genFn0() string {
    var v7 uint64 = 0
    v8 := GenS1{F0: 9223372036854775807}
    fmt.Println(v7)
    _ = v7
    _ = v8
    return ("a" + (("\n\t" + "foo bar") + ("" + ("foo bar" + "hello"))))
}
func genFn1() uint16 {
    if false {
        switch ("foo bar" + ("\n\t" + "")) {
        case "\n\t":
            var v9 int16 = 32767
            _ = v9
        case "foo bar":
            var v10 float32 = 3.14159
            _ = v10
        default:
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 float32 = 0.1
                _ = fi11
                _ = v12
            }
        }
        var v13 float32 = 3.14159
        _ = v13
    } else {
        v14 := []int{}
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 uint8 = 2
            var v17 int32 = 0
            v17 = -1
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = v14
    }
    var v18 int16 = -32768
    var v19 int64 = -1
    var v20 int32 = 362
    var v21 uint16 = 42
    _ = v18
    _ = v19
    _ = v20
    _ = v21
    return (^(v21))
}

func main() {
    var v1 int16 = 423
    var v2 int64 = -1
    var v3 uint = 892
    v4 := make(map[int]string)
    var v5 uint32 = 1
    v6 := []string{"a", "", "", "foo bar"}
    fmt.Println(v2, v3, v2)
    switch (("" + "") + ("hello" + ("\n\t" + "hello"))) {
    case "foo bar":
        if false {
            if false {
                v22 := GenS1{F0: int64(v5)}
                v23 := GenS1{F0: (v2 / (((-(v2))) | 1))}
                v2 += (v2 ^ v2)
                _ = v22
                _ = v23
            } else {
                var v24 uint64 = 4294967295
                _ = v24
            }
            v25 := GenS2{F0: (v3 < (v3 + uint(v5)))}
            switch ((("" + ("a" + ("a" + ""))) + "hello") + "") {
            case "":
                v26 := GenS3{F0: -128}
                _ = v26
            }
            _ = v25
        } else {
            for fi27 := 0; fi27 < 2; fi27++ {
                v3 += v3
                _ = fi27
            }
            for fi28 := 0; fi28 < 17; fi28++ {
                for fi29 := 0; fi29 < 15; fi29++ {
                    if (!(0.0 > 1.5)) {
                        var v30 int = int(v1)
                        _ = v30
                    }
                    var v31 uint16 = 871
                    _ = fi29
                    _ = v31
                }
                _ = fi28
            }
        }
    case "":
        fmt.Println(v3, v1)
    case "\n\t":
        var v32 uint = 4294967295
        _ = v32
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        v34 := GenS4{F0: "foo bar", F1: (v1 ^ v1)}
        _ = fi33
        _ = v34
    }
    var v35 uint64 = 4294967295
    if false {
        switch 127 {
        case 0:
            v36 := GenS5{F0: uint32(v35), F1: v35}
            _ = v36
        case 9223372036854775807:
            v6 = append(v6, (("a" + "foo bar") + "a"))
        }
        var v37 float32 = -2.25
        _ = v37
    }
    switch "hello" {
    case "\n\t":
        var v38 int64 = v2
        _ = v38
    }
    var v39 uint64 = 42
    _ = len(v6)
    switch -272 {
    case 42:
        var v40 float64 = 3.14159
        _ = v40
    case 4294967295:
        v41 := GenS6{F0: 2147483647}
        _ = v41
    case 32768:
        v39 = v39
    }
    switch "a" {
    case "":
        fmt.Println(v3, v39)
    case "foo bar":
        v42 := GenS2{F0: (!true)}
        _ = v42
    default:
        v35 = 18446744073709551615
    }
    v3 = 643
    var v43 uint8 = 255
    if (v5 == (^(v5))) {
        switch 128 {
        case 127:
            var v44 uint32 = v5
            _ = v44
        case 0:
            var v45 uint = (v3 << ((v3) & 63))
            _ = v45
        default:
            fmt.Println(v5, v2, v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
