package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int8 }
func genFn0(p0 uint64, p1 uint64, p2 uint64) uint32 {
    if false {
        if true {
            v6 := []int{}
            _ = v6
        }
    } else {
        var v7 int8 = -128
        _ = v7
    }
    var v8 float32 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2147483648
}

func main() {
    var v1 float64 = 1e10
    var v2 uint = 18446744073709551615
    var v3 uint = (v2 << ((255) & 63))
    v4 := make(map[string]int)
    var v5 string = "a"
    v3 = (v2 - (v2 % (((v3 | 4294967295)) | 1)))
    v5 = ("\n\t" + "a")
    switch v5 {
    case "":
        fmt.Println(v3, v3, v2)
    default:
        for fi9 := 0; fi9 < 7; fi9++ {
            fmt.Println(v3, v1)
            _ = fi9
        }
    }
    var v10 bool = false
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 19; fi12++ {
            v5 = v5
            var v13 uint32 = 42
            var v14 string = v5
            v14 = "hello"
            switch v14 {
            case "hello":
                switch (("hello" + ("a" + "foo bar")) + v5) {
                case "":
                    var v15 uint8 = 255
                    _ = v15
                case "hello":
                    var v16 int32 = -123
                    _ = v16
                default:
                    v17 := GenS1{F0: ((v14 + "hello") + (("foo bar" + "\n\t") + v5)), F1: v3, F2: -1}
                    _ = v17
                }
            case "foo bar":
                v10 = v10
            }
            if v10 {
                var v18 int16 = -32768
                _ = v18
            } else {
                if true {
                    var v19 uint16 = 654
                    _ = v19
                } else {
                    var v20 uint16 = 1
                    var v21 int32 = -2147483648
                    switch v5 {
                    case "hello":
                        v22 := GenS2{F0: 255}
                        _ = v22
                    default:
                        v2 = (^(v3))
                    }
                    _ = v20
                    _ = v21
                }
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        var v23 uint64 = 962
        _ = fi11
        _ = v23
    }
    if v10 {
        fmt.Println(v1, v2)
        if true {
            _ = len(v4)
        }
    } else {
        for fi24 := 0; fi24 < 5; fi24++ {
            v25 := []int{fi24, fi24, fi24, (^(fi24))}
            _ = fi24
            _ = v25
        }
        var v26 string = "\n\t"
        _ = v26
    }
    v2 += v3
    var v27 int32 = 171
    fmt.Println(v3, v5)
    if v10 {
        switch 128 {
        case 127:
            v2 += v2
        case 32768:
            for fi28 := 0; fi28 < 2; fi28++ {
                v10 = (v10 && v10)
                _ = fi28
            }
        case 128:
            switch int(v2) {
            case 128:
                v4[v5] = 2147483647
            }
        default:
            var v29 string = v5
            _ = v29
        }
        v4[("\n\t" + v5)] = 4294967295
    } else {
        for fi30 := 0; fi30 < 1; fi30++ {
            fmt.Println(v1, v3, fi30)
            var v31 bool = true
            v3 -= (v3 / ((v3) | 1))
            _ = fi30
            _ = v31
        }
    }
    switch v5 {
    case "a":
        v10 = ("a" != v5)
    case "foo bar":
        v27 = v27
    case "\n\t":
        for fi32 := 0; fi32 < 8; fi32++ {
            v5 = ("\n\t" + (("foo bar" + ("\n\t" + "foo bar")) + (v5 + (((("a" + v5) + "\n\t") + v5) + (("\n\t" + (v5 + "hello")) + "foo bar")))))
            _ = fi32
        }
    }
    fmt.Println(v27, v27)
    v10 = ((!v10) && v10)
    switch "a" {
    case "hello":
        var v33 int64 = -1
        _ = v33
    }
    v27 -= v27
    v34 := GenS3{F0: "\n\t", F1: -128}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
