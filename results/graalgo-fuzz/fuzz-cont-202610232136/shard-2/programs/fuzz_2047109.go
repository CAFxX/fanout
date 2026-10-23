package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
func genFn0() uint {
    v6 := []int{127, 9223372036854775807, 32767}
    var v7 uint8 = 2
    switch 42 {
    case -9223372036854775808:
        var v8 bool = ((v7 * (^(v7))) < v7)
        _ = v8
    case -65:
        v7 += 42
    case 42:
        var v9 string = ""
        _ = v9
    default:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v7 += 127
    _ = v6
    _ = v7
    return 108
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[string]int)
    var v3 bool = true
    v4 := make(map[int]uint32)
    var v5 uint64 = 1
    if (!(!(v5 >= (v5 ^ v5)))) {
        for fi10 := 0; fi10 < 10; fi10++ {
            v5 -= v5
            _ = fi10
        }
        if (!v3) {
            for fi11 := 0; fi11 < 19; fi11++ {
                fmt.Println(fi11, v3)
                _ = fi11
            }
            if (!((!(1.5 >= 0.0)) || false)) {
                if true {
                    v5 += v5
                    v3 = v3
                } else {
                    var v12 int64 = 0
                    v5 = 1
                    _ = v12
                }
                for fi13 := 0; fi13 < 16; fi13++ {
                    v5 -= (v5 * v5)
                    _ = fi13
                }
                var v14 int8 = 42
                _ = v14
            } else {
                if v3 {
                    var v15 uint8 = 128
                    if (int64(v1) != -358) {
                        v16 := v2[""]
                        v1 = v16
                        _ = len(v4)
                        _ = v16
                    }
                    _ = v15
                }
                fmt.Println(v5, v5)
                var v17 int64 = 9223372036854775807
                _ = v17
            }
            var v18 int8 = -128
            _ = v18
        }
        for fi19 := 0; fi19 < 20; fi19++ {
            switch "" {
            case "":
                if (-32768 == 694) {
                    v2[""] = v1
                    v1 += (v1 >> ((genFn0()) & 63))
                    var v20 int64 = 0
                    _ = v20
                } else {
                    fmt.Println(v1, fi19, v3)
                }
            case "hello":
                var v21 float32 = ((3.14159 / (1e10 + (1e10 - (1e10 * float32(((3.14159 + ((3.14159 / ((-2.25 / 1e10) / 3.14159)) / 1.0)) * 0.1)))))) * (-2.25 + (1.5 * (float32(-2.25) + (3.14159 + 1.5)))))
                _ = v21
            case "a":
                switch "a" {
                case "":
                    v5 = 0
                default:
                    var v22 uint = 0
                    _ = v22
                }
            default:
                v1 += v1
            }
            var v23 float64 = -2.25
            switch "\n\t" {
            case "hello":
                v3 = true
            }
            _ = fi19
            _ = v23
        }
    } else {
        v5 = v5
        if v3 {
            v24 := []int{v1, v1, 42, v1}
            fmt.Println(v1)
            v5 = 18446744073709551615
            _ = v24
        }
    }
    var v25 bool = (!true)
    delete(v4, v1)
    v5 = (v5 / ((v5) | 1))
    var v26 int16 = -32768
    for fi27 := 0; fi27 < 12; fi27++ {
        var v28 uint16 = 42
        v5 += uint64(v28)
        var v29 int64 = 42
        _ = fi27
        _ = v28
        _ = v29
    }
    v30 := GenS1{F0: 0, F1: (^(v5))}
    v31 := GenS1{F0: -2147483648, F1: v5}
    _ = len(v4)
    v3 = (v25 && v3)
    if v3 {
        _ = len(v2)
        switch (((("\n\t" + (((("foo bar" + "hello") + "") + "foo bar") + "hello")) + (("foo bar" + (("a" + "\n\t") + "\n\t")) + "foo bar")) + ("foo bar" + "a")) + "") {
        case "\n\t":
            for fi32 := 0; fi32 < 3; fi32++ {
                for fi33 := 0; fi33 < 5; fi33++ {
                    v3 = false
                    var v34 uint8 = 0
                    _ = fi33
                    _ = v34
                }
                var v35 int16 = 35
                _ = fi32
                _ = v35
            }
        case "foo bar":
            v4[(v1 + 362)] = 4294967295
        case "hello":
            v36 := GenS1{F0: 606, F1: (v5 + 0)}
            _ = v36
        }
        _ = len(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
