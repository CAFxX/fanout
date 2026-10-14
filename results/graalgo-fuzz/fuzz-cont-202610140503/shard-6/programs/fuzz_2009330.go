package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
func genFn0() int16 {
    var v6 int = 4294967295
    var v7 int64 = -1
    _ = v6
    _ = v7
    return -336
}
func genFn1(p0 uint32, p1 int, p2 uint) bool {
    var v8 uint8 = 128
    v9 := GenS1{F0: 31}
    var v10 string = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return (100 < -50)
}

func main() {
    var v1 uint16 = 554
    var v2 float32 = (1e10 / 1e10)
    var v3 int16 = 1000
    v4 := make(map[string]uint32)
    var v5 uint32 = 661
    if ((v3 | genFn0()) == genFn0()) {
        v1 = (v1 % ((597) | 1))
    } else {
        var v11 bool = false
        v12 := GenS1{F0: 42}
        _ = v11
        _ = v12
    }
    v3 = v3
    var v13 bool = (!true)
    fmt.Println(v1, v3)
    var v14 uint8 = uint8(v1)
    switch 254 {
    case 1:
        switch (("hello" + "foo bar") + "foo bar") {
        case "foo bar":
            v1 = (v1 + 704)
        case "\n\t":
            v15 := []int{-361, 42, 2147483647, -132}
            _ = v15
        default:
            delete(v4, (("\n\t" + (("" + ("a" + "foo bar")) + "foo bar")) + "foo bar"))
        }
    }
    switch "" {
    case "\n\t":
        v5 = (v5 / (((v5 - v5)) | 1))
    case "":
        fmt.Println(v2, v1)
    case "hello":
        switch 621 {
        case -418:
            switch "hello" {
            case "":
                for fi16 := 0; fi16 < 13; fi16++ {
                    v1 -= v1
                    _ = fi16
                }
            case "foo bar":
                v17 := GenS2{F0: v13}
                _ = v17
            case "\n\t":
                v18 := []int{9223372036854775807}
                _ = v18
            }
        case 0:
            _ = len(v4)
        }
    default:
        if false {
            for fi19 := 0; fi19 < 5; fi19++ {
                if (v13 && (-11 != -128)) {
                    _ = len(v4)
                }
                _ = fi19
            }
            if false {
                var v20 uint64 = 4294967295
                if false {
                    fmt.Println(v13, v20)
                    if v13 {
                        for fi21 := 0; fi21 < 5; fi21++ {
                            v1 = 776
                            v13 = true
                            var v22 uint16 = v1
                            _ = fi21
                            _ = v22
                        }
                        var v23 float64 = float64(v2)
                        _ = v23
                    } else {
                        v14 = v14
                        v24 := GenS2{F0: v13}
                        v25 := v4["foo bar"]
                        _ = v24
                        _ = v25
                    }
                } else {
                    v26 := v4[("foo bar" + "")]
                    _ = v26
                }
                v5 -= 42
                _ = v20
            }
            v27 := GenS3{F0: (!(true && v13))}
            _ = v27
        } else {
            if (!false) {
                var v28 int64 = -275
                _ = v28
            } else {
                if true {
                    v1 = (v1 & 42)
                }
            }
            v1 += v1
            v1 = uint16(v5)
        }
    }
    var v29 uint64 = 494
    v29 -= v29
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
