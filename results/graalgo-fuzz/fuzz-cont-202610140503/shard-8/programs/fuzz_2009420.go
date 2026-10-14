package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 int32; F1 int8 }
func genFn0(p0 bool, p1 bool) uint16 {
    var v7 uint8 = 53
    var v8 int = 128
    v9 := GenS1{F0: 98}
    v10 := GenS1{F0: 2}
    fmt.Println(v8, v7, p0)
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 32767
}
func genFn1(p0 uint8, p1 uint32, p2 int) uint {
    for fi11 := 0; fi11 < 8; fi11++ {
        p1 = p1
        _ = fi11
    }
    var v12 uint64 = 42
    v13 := GenS1{F0: p0}
    var v14 uint64 = v12
    var v15 string = (((("a" + "hello") + "foo bar") + ("" + ("" + "\n\t"))) + (("\n\t" + ("foo bar" + ("\n\t" + "foo bar"))) + (("" + ("a" + (("hello" + "a") + (("hello" + ("hello" + "foo bar")) + "")))) + "hello")))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return 1
}

func main() {
    var v1 int16 = -1
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    var v4 float32 = -2.25
    v5 := make(map[uint32]bool)
    v6 := []string{"a", ""}
    v1 = v1
    var v16 uint = 2
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 float64 = 0.0
        fmt.Println(v18)
        var v19 uint16 = 485
        _ = fi17
        _ = v18
        _ = v19
    }
    if false {
        v1 = 0
    }
    var v20 int = -244
    fmt.Println(v20, v20, v16)
    v5[690] = (42 > -1)
    if ((v16 << (((v16 + 2)) & 63)) > (v16 + (v16 % (((v16 & (v16 | v16))) | 1)))) {
        v21 := GenS1{F0: 1}
        v22 := GenS1{F0: 183}
        v23 := v2[(v20 + (v20 ^ (v20 & v20)))]
        _ = v21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi24
    }
    var v25 bool = (859 != v16)
    var v26 string = "\n\t"
    fmt.Println(v16, v4)
    if ((v1 * 1) >= -32768) {
        switch v26 {
        case "\n\t":
            if len(v6) > 0 {
                v6[0] = "\n\t"
            }
        case "a":
            for fi27 := 0; fi27 < 1; fi27++ {
                var v28 string = v26
                switch "a" {
                case "a":
                    v20 = v20
                case "":
                    var v29 int8 = 100
                    _ = v29
                default:
                    var v30 uint = (v16 >> ((v16) & 63))
                    _ = v30
                }
                _ = fi27
                _ = v28
            }
        default:
            v4 = (float32(0.0) / v4)
        }
        var v31 uint8 = 165
        for fi32 := 0; fi32 < 18; fi32++ {
            var v33 string = v26
            v34 := GenS1{F0: v31}
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = v31
    } else {
        v35 := GenS2{F0: 42}
        _ = len(v6)
        v36 := GenS1{F0: 42}
        _ = v35
        _ = v36
    }
    v26 = v26
    if ("hello" < (((v26 + ("foo bar" + ("hello" + (("\n\t" + v26) + v26)))) + ("hello" + "hello")) + "")) {
        v37 := GenS3{F0: genFn1(255, 243, -767), F1: 1}
        _ = v37
    } else {
        v38 := GenS4{F0: int32(v1), F1: 127}
        _ = v38
    }
    var v39 int16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
