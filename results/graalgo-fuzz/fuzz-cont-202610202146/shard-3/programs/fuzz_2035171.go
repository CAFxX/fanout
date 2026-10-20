package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint16 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint32) bool {
    var v4 int = 2147483647
    p0 = 2
    _ = p0
    _ = v4
    return (164 > 4294967295)
}

func main() {
    var v1 int64 = 1
    var v2 float32 = 0.1
    var v3 float32 = ((1e10 - 1.5) / v2)
    if ((-488 < 42) || ((!false) || (-1 > 1))) {
        v1 += (v1 / ((v1) | 1))
    }
    var v5 float64 = 0.0
    v2 += (v3 / (v2 * (3.14159 + 0.1)))
    v6 := []int{42, 42}
    if (!false) {
        v1 = v1
        v7 := GenS1{F0: -73, F1: -1, F2: 2}
        _ = v7
    }
    fmt.Println(v1, v1, v2)
    var v8 int = 0
    if false {
        var v9 string = ("" + "foo bar")
        _ = v9
    }
    if len(v6) > 0 {
        v6[0] = v8
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        if true {
            for fi11 := 0; fi11 < 12; fi11++ {
                var v12 float64 = (float64(((float32(0.0) * -2.25) - -2.25)) - v5)
                var v13 uint32 = 0
                _ = fi11
                _ = v12
                _ = v13
            }
            var v14 uint32 = 317
            switch ("\n\t" + "") {
            case "a":
                for fi15 := 0; fi15 < 9; fi15++ {
                    var v16 int = v8
                    _ = fi15
                    _ = v16
                }
            case "hello":
                v1 = 42
            case "foo bar":
                switch "\n\t" {
                case "":
                    for fi17 := 0; fi17 < 4; fi17++ {
                        fmt.Println(v1, v3, v5)
                        v18 := GenS1{F0: 0, F1: -102, F2: 255}
                        var v19 uint64 = 1
                        if true {
                            var v20 uint16 = 1
                            fmt.Println(fi10)
                            _ = v20
                        } else {
                            v2 -= v2
                        }
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
            default:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = v14
        }
        if false {
            v21 := GenS2{F0: v8, F1: fi10}
            if true {
                v5 = (v5 + 1e10)
                var v22 int32 = 0
                _ = v22
            } else {
                v23 := GenS1{F0: -117, F1: -128, F2: 255}
                switch ("foo bar" + "\n\t") {
                case "foo bar":
                    if false {
                        var v24 bool = false
                        var v25 int64 = (v1 >> ((2) & 63))
                        v3 += float32(1e10)
                        _ = v24
                        _ = v25
                    }
                case "":
                    v1 = v1
                case "a":
                    for fi26 := 0; fi26 < 18; fi26++ {
                        var v27 uint64 = 586
                        var v28 string = (("hello" + "") + "")
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                }
                _ = v23
            }
            _ = v21
        } else {
            var v29 float64 = v5
            _ = v29
        }
        _ = fi10
    }
    var v30 int16 = 0
    if (255 >= 127) {
        for fi31 := 0; fi31 < 15; fi31++ {
            switch (("hello" + "") + "hello") {
            case "":
                switch "" {
                case "foo bar":
                    var v32 float32 = v3
                    _ = v32
                default:
                    v33 := GenS3{F0: 255}
                    _ = v33
                }
            case "a":
                v1 += (v1 - (v1 * 1))
            default:
                for fi34 := 0; fi34 < 18; fi34++ {
                    var v35 int64 = v1
                    _ = fi34
                    _ = v35
                }
            }
            var v36 uint64 = 1
            _ = fi31
            _ = v36
        }
        v1 += v1
        v1 += 37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v5
    _ = v6
    _ = v8
}
