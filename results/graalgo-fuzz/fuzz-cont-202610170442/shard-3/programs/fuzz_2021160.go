package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS4 struct { F0 bool; F1 int64; F2 bool }

func main() {
    var v1 int64 = -561
    var v2 uint = 544
    var v3 uint64 = 0
    var v4 float32 = 3.14159
    v5 := make(map[string]int)
    for fi6 := 0; fi6 < 7; fi6++ {
        if true {
            fmt.Println(v4)
            v4 = (((v4 / 1.0) - v4) + v4)
            switch fi6 {
            case 2147483647:
                var v7 float64 = float64(0.0)
                _ = v7
            case 9223372036854775807:
                switch "hello" {
                case "hello":
                    switch "\n\t" {
                    case "hello":
                        v5[((("foo bar" + "\n\t") + "") + "a")] = fi6
                    case "a":
                        if (255 < 42) {
                            v2 += uint(v3)
                            v2 = v2
                            fmt.Println(v1, v4, v2)
                        }
                    }
                case "\n\t":
                    if true {
                        fmt.Println(v1, v4)
                        v8 := GenS1{F0: 127}
                        _ = v8
                    }
                case "a":
                    var v9 float32 = 3.14159
                    _ = v9
                }
            case -170:
                switch fi6 {
                case 447:
                    switch (fi6 & (fi6 + (fi6 - fi6))) {
                    case -9223372036854775808:
                        for fi10 := 0; fi10 < 9; fi10++ {
                            var v11 int32 = -525
                            v12 := GenS2{F0: ""}
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                    }
                case -64:
                    for fi13 := 0; fi13 < 7; fi13++ {
                        v4 += v4
                        v1 -= -279
                        _ = fi13
                    }
                case 42:
                    v4 = 3.14159
                default:
                    if false {
                        if true {
                            var v14 int64 = v1
                            _ = v14
                        }
                    }
                }
            }
        } else {
            v3 = v3
            v15 := []int{(fi6 & 32767), fi6}
            _ = v15
        }
        _ = fi6
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        v17 := GenS1{F0: 33}
        if false {
            v18 := GenS3{F0: 2, F1: (v2 + 279), F2: 4294967295}
            v3 -= (v3 / ((v3) | 1))
            _ = v18
        } else {
            fmt.Println(v4, v4)
            if ((-(v1)) > 882) {
                if false {
                    v19 := []int{}
                    if ((0.1 * 1.5) >= -2.25) {
                        v1 = v1
                        var v20 string = ""
                        _ = v20
                    }
                    _ = v19
                }
            } else {
                v21 := []int{fi16, 127, (fi16 * 2147483647), fi16}
                _ = v21
            }
        }
        var v22 uint32 = 1
        _ = fi16
        _ = v17
        _ = v22
    }
    var v23 uint64 = 1
    v2 += 255
    for fi24 := 0; fi24 < 6; fi24++ {
        v25 := GenS1{F0: 100}
        var v26 int32 = 42
        for fi27 := 0; fi27 < 1; fi27++ {
            v1 = 818
            _ = fi27
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    v1 -= (v1 ^ (v1 * v1))
    v28 := GenS4{F0: false, F1: (v1 - (-(v1))), F2: (-2.25 <= 1e10)}
    var v29 uint8 = 2
    switch ("\n\t" + ("\n\t" + "")) {
    case "hello":
        var v30 uint16 = 575
        _ = v30
    case "a":
        var v31 uint16 = 729
        _ = v31
    case "foo bar":
        var v32 int16 = -32768
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
