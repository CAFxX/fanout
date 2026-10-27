package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint32; F1 uint16 }
func genFn0(p0 uint16, p1 int8) string {
    v7 := GenS1{F0: p1, F1: 21, F2: p1}
    var v8 string = "foo bar"
    var v9 int32 = 0
    for fi10 := 0; fi10 < 3; fi10++ {
        fmt.Println(p1, fi10)
        var v11 int32 = 42
        if true {
            if true {
                fmt.Println(v8)
            }
            for fi12 := 0; fi12 < 14; fi12++ {
                fmt.Println(p0, fi12, v8)
                var v13 uint32 = 757
                v14 := GenS2{F0: (p0 - p0), F1: fi10}
                var v15 int8 = 100
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return v8
}
func genFn1(p0 string, p1 bool, p2 int64) bool {
    p0 = p0
    p2 = 9223372036854775807
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 uint64 = 42
        p1 = true
        v18 := []int{fi16, 872, (^(fi16))}
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS3{F0: 378, F1: 902}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return (2 > 1)
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]uint32)
    v3 := make(map[int]uint32)
    var v4 uint16 = 15
    var v5 float64 = -2.25
    v6 := make(map[int]string)
    fmt.Println(v5)
    for fi20 := 0; fi20 < 13; fi20++ {
        v5 += v5
        var v21 uint16 = 867
        switch (("foo bar" + "a") + "foo bar") {
        case "\n\t":
            var v22 uint16 = (v21 ^ v21)
            _ = v22
        case "foo bar":
            v4 += (v4 | v21)
        case "hello":
            v23 := []int{(^(fi20))}
            _ = v23
        default:
            for fi24 := 0; fi24 < 1; fi24++ {
                fmt.Println(v21)
                _ = fi24
            }
        }
        _ = fi20
        _ = v21
    }
    var v25 uint32 = 0
    if (2 < 56) {
        v4 = 65535
    } else {
        if ((v25 / ((1) | 1)) == v25) {
            v26 := GenS1{F0: 0, F1: 2, F2: 0}
            for fi27 := 0; fi27 < 19; fi27++ {
                v28 := v3[fi27]
                fmt.Println(v5, v5, v28)
                var v29 float64 = v5
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = v26
        } else {
            v30 := v2[("hello" + "foo bar")]
            var v31 int8 = 42
            _ = v30
            _ = v31
        }
        fmt.Println(v4, v5, v25)
        var v32 int16 = -1
        _ = v32
    }
    var v33 int64 = -893
    var v34 uint64 = 4294967295
    v35 := GenS2{F0: (v4 / ((42) | 1)), F1: 4294967295}
    if true {
        var v36 int32 = -1
        _ = v36
    }
    fmt.Println(v33)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
