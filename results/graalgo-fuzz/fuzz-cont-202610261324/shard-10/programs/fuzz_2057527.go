package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
func genFn0() uint {
    var v6 uint64 = 18446744073709551615
    v7 := GenS1{F0: 0}
    v8 := GenS2{F0: 401}
    if true {
        v9 := []int{}
        _ = v9
    } else {
        fmt.Println(v6)
    }
    _ = v6
    _ = v7
    _ = v8
    return 18446744073709551615
}
func genFn1(p0 uint64) uint64 {
    v10 := GenS2{F0: -942}
    var v11 bool = true
    v11 = false
    for fi12 := 0; fi12 < 20; fi12++ {
        fmt.Println(fi12)
        _ = fi12
    }
    _ = p0
    _ = v10
    _ = v11
    return (p0 % ((519) | 1))
}

func main() {
    var v1 float32 = 1.5
    v2 := []string{("hello" + "foo bar"), "foo bar"}
    var v3 uint8 = 168
    var v4 float64 = -2.25
    var v5 uint8 = (v3 / ((v3) | 1))
    var v13 uint32 = 1
    var v14 uint = 18446744073709551615
    for fi15 := 0; fi15 < 13; fi15++ {
        v4 -= (((v4 - v4) + 1e10) * (float64(0.1) * 0.1))
        var v16 uint = (v14 / (((v14 - 2)) | 1))
        var v17 string = (("hello" + "foo bar") + ((("foo bar" + "a") + ("" + "a")) + "a"))
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 int64 = -9223372036854775808
    v1 -= 0.0
    v3 = (v3 * 180)
    if false {
        for fi19 := 0; fi19 < 13; fi19++ {
            fmt.Println(v4)
            v2 = append(v2, "foo bar")
            _ = fi19
        }
        if ((v5 | 42) >= (v3 << ((0) & 63))) {
            var v20 float64 = ((float64(1.5) / (v4 / 1e10)) / v4)
            fmt.Println(v1, v20)
            var v21 int64 = -364
            _ = v20
            _ = v21
        }
    }
    var v22 float64 = ((v4 * (v4 + (v4 - (v4 * 1.5)))) * float64(float32((v4 * v4))))
    _ = len(v2)
    _ = len(v2)
    var v23 int64 = (v18 + (v18 | v18))
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 float64 = float64((v1 * 0.1))
        if true {
            var v26 float64 = 0.1
            _ = v26
        } else {
            var v27 float32 = v1
            switch "foo bar" {
            case "a":
                if (1 == -1) {
                    var v28 int32 = -658
                    var v29 float32 = 1e10
                    _ = v28
                    _ = v29
                }
            case "foo bar":
                for fi30 := 0; fi30 < 5; fi30++ {
                    v31 := GenS1{F0: 0}
                    var v32 int8 = 42
                    v4 -= 0.1
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            case "hello":
                v22 -= v25
            default:
                fmt.Println(v14, v14)
            }
            _ = v27
        }
        _ = fi24
        _ = v25
    }
    var v33 uint32 = 2147483648
    v34 := GenS1{F0: 547}
    switch "foo bar" {
    case "":
        v35 := GenS2{F0: 0}
        _ = v35
    case "\n\t":
        fmt.Println(v3, v23, v23)
    }
    var v36 int8 = -49
    var v37 uint32 = 2147483647
    for fi38 := 0; fi38 < 3; fi38++ {
        v39 := GenS2{F0: (fi38 * (fi38 & fi38))}
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
