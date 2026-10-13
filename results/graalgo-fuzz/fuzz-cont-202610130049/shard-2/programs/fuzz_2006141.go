package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64 }
func genFn0(p0 uint64, p1 uint32, p2 uint64) uint32 {
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: "\n\t"}
        p0 -= p0
        _ = fi7
        _ = v8
    }
    p0 += (p0 * p2)
    if true {
        var v9 uint16 = uint16(p2)
        if false {
            v9 = v9
        }
        if false {
            var v10 int = -2147483648
            fmt.Println(v9, v10, v9)
            _ = v10
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 | (^(p1)))
}

func main() {
    var v1 uint = 1
    v2 := []string{"\n\t", "foo bar", "\n\t", "hello"}
    var v3 int32 = 42
    var v4 int64 = 0
    v5 := make(map[string]string)
    var v6 bool = false
    for fi11 := 0; fi11 < 10; fi11++ {
        switch "hello" {
        case "hello":
            var v12 float32 = ((0.1 / (float32(((((((float64(0.1) / 1.5) * 3.14159) + 1e10) * 3.14159) + (1.5 + 3.14159)) - (3.14159 / 1.5))) * -2.25)) + (3.14159 / -2.25))
            _ = v12
        case "foo bar":
            switch "foo bar" {
            case "":
                for fi13 := 0; fi13 < 12; fi13++ {
                    v3 = (v3 - (^(v3)))
                    fmt.Println(fi11, v3)
                    fmt.Println(v4)
                    _ = fi13
                }
            default:
                fmt.Println(v1)
            }
        case "\n\t":
            if false {
                v14 := GenS2{F0: (v1 - v1), F1: (v1 >> ((v1) & 63))}
                var v15 int16 = -489
                _ = v14
                _ = v15
            }
        }
        _ = fi11
    }
    _ = len(v5)
    var v16 uint = 65535
    for fi17 := 0; fi17 < 15; fi17++ {
        v18 := GenS2{F0: 1, F1: (v16 >> (((v16 << (((^(v1))) & 63))) & 63))}
        for fi19 := 0; fi19 < 19; fi19++ {
            if (!(true || v6)) {
                v20 := GenS3{F0: 42, F1: 1}
                switch (("hello" + "") + "hello") {
                case "hello":
                    if (485 == v3) {
                        fmt.Println(v1, fi17)
                        v1 -= v16
                    } else {
                        v21 := GenS4{F0: 1}
                        _ = v21
                    }
                case "a":
                    v22 := GenS5{F0: 1000}
                    _ = v22
                case "":
                    var v23 int = 886
                    _ = v23
                }
                _ = v20
            }
            fmt.Println(v1, fi19, v6)
            var v24 uint = v16
            _ = fi19
            _ = v24
        }
        _ = fi17
        _ = v18
    }
    if false {
        v4 = v4
    }
    fmt.Println(v16, v6)
    var v25 uint64 = 0
    for fi26 := 0; fi26 < 1; fi26++ {
        fmt.Println(v3)
        for fi27 := 0; fi27 < 2; fi27++ {
            var v28 uint16 = 65535
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        if v6 {
            for fi30 := 0; fi30 < 16; fi30++ {
                v4 -= 0
                if v6 {
                    v25 = (^(v25))
                    v5["\n\t"] = "foo bar"
                }
                v25 = v25
                _ = fi30
            }
            v31 := GenS6{F0: v4}
            _ = v31
        }
        _ = fi29
    }
    _ = len(v5)
    var v32 uint = (v16 - (v1 + 459))
    var v33 bool = ((v1 - 4294967295) != 255)
    v34 := GenS1{F0: "foo bar"}
    v1 = (^(v16))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
