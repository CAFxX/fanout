package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint8; F1 int32; F2 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 int32; F2 bool }
type GenS5 struct { F0 int; F1 uint64; F2 uint8 }
func genFn0() int32 {
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: ("foo bar" + "hello"), F1: "hello"}
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 float64 = 1.5
            v8 += v8
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    v9 := GenS1{F0: "a", F1: "foo bar"}
    var v10 string = "\n\t"
    _ = v10
    _ = v9
    return -552
}

func main() {
    v1 := make(map[int]int)
    var v2 float32 = 1.5
    v3 := make(map[uint32]uint32)
    v4 := []string{(("foo bar" + "\n\t") + "\n\t"), "a", ("a" + "a")}
    v4 = append(v4, "\n\t")
    if true {
        var v11 int32 = 42
        fmt.Println(v11, v11)
        _ = v11
    }
    v12 := v3[4294967295]
    v2 += float32(0.0)
    v2 = v2
    _ = len(v1)
    for fi13 := 0; fi13 < 10; fi13++ {
        if len(v4) > 0 {
            v4[0] = ""
        }
        _ = fi13
    }
    switch 1 {
    case 4294967295:
        if false {
            v2 = (v2 + v2)
            if len(v4) > 0 {
                v4[0] = ""
            }
        } else {
            v2 -= float32(0.1)
        }
    case 566:
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
    case 128:
        v12 = v12
    }
    if true {
        v12 -= v12
        switch 2147483647 {
        case 2147483647:
            switch "hello" {
            case "foo bar":
                var v14 int16 = 0
                _ = v14
            case "":
                switch "a" {
                case "foo bar":
                    var v15 int8 = 100
                    _ = v15
                case "":
                    for fi16 := 0; fi16 < 17; fi16++ {
                        v12 -= v12
                        if false {
                            v12 = 375
                        }
                        v17 := GenS1{F0: "", F1: "hello"}
                        _ = fi16
                        _ = v17
                    }
                case "a":
                    if (!(float32((((((1.5 / (float64(-2.25) * ((0.0 + 0.0) + (1e10 - (1.5 + 3.14159))))) / 1.5) + ((((0.1 - 1e10) * 3.14159) - 1.5) + -2.25)) * (((0.0 - (0.0 - -2.25)) + (-2.25 * ((1e10 - 0.0) * 1.5))) * (-2.25 * 0.0))) + 0.1)) > v2)) {
                        var v18 int16 = 1
                        _ = v18
                    } else {
                        v12 = v12
                        fmt.Println(v2, v12)
                        v2 -= v2
                    }
                }
            case "\n\t":
                v12 += (v12 - v12)
            }
        default:
            switch "" {
            case "foo bar":
                for fi19 := 0; fi19 < 4; fi19++ {
                    v20 := GenS2{F0: 0, F1: 1, F2: ((((((("a" + "hello") + "foo bar") + "") + ("\n\t" + "hello")) + "") + ("foo bar" + "\n\t")) + "\n\t")}
                    if false {
                        for fi21 := 0; fi21 < 14; fi21++ {
                            var v22 uint32 = (^(v12))
                            _ = fi21
                            _ = v22
                        }
                    }
                    _ = fi19
                    _ = v20
                }
            case "hello":
                fmt.Println(v2, v12)
            case "a":
                for fi23 := 0; fi23 < 20; fi23++ {
                    var v24 uint64 = 18446744073709551615
                    v25 := GenS3{F0: true}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
        }
    }
    v26 := GenS4{F0: 0, F1: 0, F2: false}
    v27 := GenS3{F0: false}
    v2 = float32(3.14159)
    var v28 uint16 = 726
    fmt.Println(v12, v2, v28)
    v3[v12] = uint32(v28)
    if false {
        fmt.Println(v28, v2)
        if true {
            switch int(v12) {
            case 1:
                var v29 int8 = 106
                _ = v29
            }
            v30 := v1[127]
            _ = v30
        } else {
            v31 := GenS1{F0: "hello", F1: "\n\t"}
            _ = v31
        }
        fmt.Println(v12, v28, v12)
    } else {
        v32 := GenS5{F0: 128, F1: 4294967295, F2: 127}
        var v33 uint = 0
        var v34 string = ("a" + ("hello" + ""))
        _ = v32
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        delete(v1, fi35)
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        var v36 int = (fi35 + -2147483648)
        _ = fi35
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
