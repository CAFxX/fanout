package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int32 }
type GenS2 struct { F0 uint8; F1 int16; F2 int64 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 uint16; F1 uint; F2 uint32 }
func genFn0(p0 uint32, p1 int32, p2 uint) uint8 {
    p0 = (p0 % ((p0) | 1))
    fmt.Println(p0)
    p0 += p0
    switch "\n\t" {
    case "a":
        v7 := GenS1{F0: 32767, F1: -233, F2: p1}
        _ = v7
    case "foo bar":
        v8 := []int{-9223372036854775808, -9223372036854775808, 9223372036854775807, 127}
        _ = v8
    case "":
        for fi9 := 0; fi9 < 1; fi9++ {
            fmt.Println(fi9, p1)
            for fi10 := 0; fi10 < 13; fi10++ {
                fmt.Println(p2, p1, p0)
                v11 := GenS2{F0: 0, F1: 0, F2: -9223372036854775808}
                fmt.Println(fi9)
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 168
}
func genFn1() int64 {
    var v12 bool = (!(true || (65535 < 2147483647)))
    var v13 string = (("" + "") + ("" + (("foo bar" + "\n\t") + "hello")))
    var v14 uint8 = 42
    _ = v12
    _ = v13
    _ = v14
    return 9223372036854775807
}

func main() {
    var v1 float64 = 3.14159
    v2 := []uint32{947}
    v3 := []string{"foo bar", (("hello" + (("" + "\n\t") + "")) + "")}
    var v4 int = 67
    var v5 uint16 = 2
    var v6 float64 = -2.25
    if false {
        var v15 int64 = 0
        _ = v15
    }
    fmt.Println(v6, v1, v6)
    for fi16 := 0; fi16 < 12; fi16++ {
        if true {
            v6 = v6
        }
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 int32 = -2147483648
            fmt.Println(v6, fi17, v1)
            _ = fi17
            _ = v18
        }
        if (v5 >= (^(v5))) {
            fmt.Println(v5, v1, v5)
            var v19 float32 = (float32((float64((float32(v6) * (((1e10 * float32(v1)) - 1e10) * 0.1))) + ((((1e10 / -2.25) / (v6 * 1.5)) * v6) - float64(0.1)))) + 0.0)
            _ = v19
        }
        _ = fi16
    }
    switch v4 {
    case 128:
        fmt.Println(v4, v6)
    case -871:
        v5 = 65535
    case 32768:
        v20 := GenS2{F0: 50, F1: 32767, F2: 0}
        _ = v20
    default:
        v5 -= (v5 << ((uint(v5)) & 63))
    }
    var v21 bool = false
    for fi22 := 0; fi22 < 10; fi22++ {
        v1 = float64(1e10)
        switch "foo bar" {
        case "hello":
            switch (((("a" + "hello") + ("hello" + "a")) + (("" + "hello") + "foo bar")) + "a") {
            case "a":
                switch 42 {
                case 128:
                    switch v4 {
                    case -901:
                        fmt.Println(v21, fi22, v4)
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                }
            case "foo bar":
                v4 = v4
            case "\n\t":
                for fi23 := 0; fi23 < 18; fi23++ {
                    var v24 uint = 1
                    v25 := GenS3{F0: (^(v24)), F1: "a"}
                    v24 = v24
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
        default:
            var v26 uint16 = v5
            _ = v26
        }
        if v21 {
            var v27 uint16 = v5
            v2 = append(v2, 1)
            _ = v27
        } else {
            v1 -= float64(0.0)
        }
        _ = fi22
    }
    var v28 int64 = 9223372036854775807
    fmt.Println(v5, v21, v1)
    v6 = v6
    if false {
        var v29 float64 = 1.5
        _ = v29
    }
    v30 := GenS3{F0: 352, F1: ""}
    fmt.Println(v1, v28, v1)
    if (!false) {
        v31 := GenS4{F0: (v5 - 464), F1: 18446744073709551615, F2: 42}
        _ = len(v3)
        _ = v31
    }
    var v32 uint32 = 2147483647
    var v33 int32 = 0
    var v34 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
