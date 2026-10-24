package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint; F1 int16 }

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]uint32)
    var v3 int8 = 47
    var v4 uint16 = 932
    if true {
        for fi5 := 0; fi5 < 7; fi5++ {
            switch ("" + "") {
            case "hello":
                if (v3 > (v3 + (v3 * v3))) {
                    var v6 bool = (!(((0.0 + 1.5) + -2.25) >= (float32(0.0) * (3.14159 - ((0.0 / 0.1) + 0.1)))))
                    var v7 float32 = 1e10
                    _ = v6
                    _ = v7
                }
            case "foo bar":
                var v8 bool = (v4 == v4)
                _ = v8
            default:
                if true {
                    switch "foo bar" {
                    case "\n\t":
                        if true {
                            var v9 string = (("foo bar" + ("a" + ("foo bar" + "foo bar"))) + "")
                            v4 -= 42
                            _ = v9
                        }
                    case "foo bar":
                        v10 := GenS1{F0: 2, F1: v1}
                        _ = v10
                    case "a":
                        if false {
                            v3 = v3
                        } else {
                            switch fi5 {
                            case 32767:
                                var v11 uint8 = 255
                                _ = v11
                            }
                        }
                    }
                    var v12 float32 = 1.5
                    switch 2147483647 {
                    case -696:
                        v13 := GenS1{F0: 0, F1: v1}
                        _ = v13
                    }
                    _ = v12
                } else {
                    v2[1] = 173
                    fmt.Println(v1, fi5, v4)
                }
            }
            v14 := v2[0]
            _ = fi5
            _ = v14
        }
        if true {
            v15 := GenS2{F0: 4294967295, F1: 1}
            _ = v15
        }
        var v16 uint16 = v4
        _ = v16
    }
    var v17 uint16 = v4
    fmt.Println(v1)
    v3 = (-(v3))
    v1 = v1
    for fi18 := 0; fi18 < 2; fi18++ {
        v3 += (v3 / ((-128) | 1))
        switch ("a" + ("\n\t" + "hello")) {
        case "hello":
            fmt.Println(v3, v17)
        case "a":
            var v19 int16 = -1
            _ = v19
        }
        _ = fi18
    }
    v4 = v4
    var v20 int8 = v3
    v3 -= (v20 / ((v20) | 1))
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 uint8 = 72
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := GenS1{F0: 177, F1: int32(fi23)}
        _ = fi23
        _ = v24
    }
    v17 -= 1
    var v25 uint16 = (v17 + 65535)
    switch 886 {
    case 32767:
        v2[42] = 480
    default:
        var v26 uint64 = 4294967295
        _ = v26
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        var v28 uint32 = 1
        var v29 int16 = 32767
        _ = fi27
        _ = v28
        _ = v29
    }
    switch 9223372036854775807 {
    case -808:
        if true {
            for fi30 := 0; fi30 < 19; fi30++ {
                switch "\n\t" {
                case "":
                    v25 -= v17
                }
                fmt.Println(v4, v1)
                for fi31 := 0; fi31 < 15; fi31++ {
                    var v32 bool = true
                    _ = fi31
                    _ = v32
                }
                _ = fi30
            }
        }
    case 292:
        fmt.Println(v25, v3, v3)
    default:
        v4 += (^(v17))
    }
    v20 += (v3 + 100)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
}
