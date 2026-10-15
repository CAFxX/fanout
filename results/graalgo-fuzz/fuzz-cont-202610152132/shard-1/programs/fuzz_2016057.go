package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 string; F1 string; F2 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint8, p1 int8, p2 uint) int8 {
    var v7 uint64 = 18446744073709551615
    v8 := []int{}
    p1 -= (p1 | (p1 | (p1 + p1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p1
}
func genFn1() int16 {
    var v9 bool = (403 <= 1)
    var v10 uint = 0
    if v9 {
        var v11 int = 4294967295
        v12 := GenS1{F0: 2, F1: -1}
        _ = v11
        _ = v12
    }
    var v13 int16 = 42
    switch -782 {
    case -1:
        v14 := []int{}
        _ = v14
    }
    _ = v10
    _ = v13
    _ = v9
    return v13
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    var v3 string = ((("\n\t" + "\n\t") + "foo bar") + ("" + "\n\t"))
    var v4 uint64 = 0
    v5 := []uint32{}
    v6 := make(map[uint32]int)
    v15 := GenS2{F0: ("foo bar" + ""), F1: (v3 + "\n\t"), F2: -551}
    for fi16 := 0; fi16 < 1; fi16++ {
        v4 -= v4
        var v17 uint = 455
        if (0 == 853) {
            if false {
                var v18 int8 = -128
                _ = v18
            } else {
                fmt.Println(v17, v4)
            }
        } else {
            for fi19 := 0; fi19 < 2; fi19++ {
                v20 := GenS2{F0: v3, F1: v3, F2: 32767}
                v3 = (v3 + v3)
                _ = fi19
                _ = v20
            }
            var v21 int16 = 32767
            _ = v21
        }
        _ = fi16
        _ = v17
    }
    v22 := GenS1{F0: 0, F1: -861}
    for fi23 := 0; fi23 < 9; fi23++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi23
    }
    switch ((v3 + v3) + v3) {
    case "foo bar":
        for fi24 := 0; fi24 < 10; fi24++ {
            v3 = v3
            if (true || false) {
                v25 := GenS3{F0: v4}
                _ = v25
            } else {
                fmt.Println(v4)
                var v26 uint = 42
                _ = len(v5)
                _ = v26
            }
            v27 := GenS4{F0: (18446744073709551615 > (v4 & v4))}
            _ = fi24
            _ = v27
        }
    default:
        switch ("" + ("\n\t" + ("\n\t" + v3))) {
        case "hello":
            var v28 uint8 = 255
            _ = v28
        default:
            v4 -= 18446744073709551615
        }
    }
    var v29 uint16 = 65535
    v30 := GenS3{F0: v4}
    var v31 uint = 1
    switch v3 {
    case "hello":
        v31 = v31
    case "a":
        for fi32 := 0; fi32 < 10; fi32++ {
            var v33 int32 = 1
            _ = fi32
            _ = v33
        }
    case "\n\t":
        for fi34 := 0; fi34 < 12; fi34++ {
            fmt.Println(fi34, v4)
            _ = fi34
        }
    default:
        v3 = ((v3 + v3) + (v3 + v3))
    }
    v35 := GenS4{F0: true}
    var v36 bool = (false || (v3 > "foo bar"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
