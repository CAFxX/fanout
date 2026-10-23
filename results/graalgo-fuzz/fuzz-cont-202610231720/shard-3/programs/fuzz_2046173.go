package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint8) int8 {
    v6 := GenS1{F0: 0, F1: p0}
    for fi7 := 0; fi7 < 16; fi7++ {
        switch "" {
        case "hello":
            v8 := []int{}
            _ = v8
        default:
            if true {
                if true {
                    var v9 string = ("" + "hello")
                    _ = v9
                } else {
                    v10 := []int{fi7}
                    fmt.Println(p0)
                    var v11 int8 = int8(p0)
                    _ = v10
                    _ = v11
                }
                var v12 float64 = ((((0.1 - float64(-2.25)) + 0.1) + 1.5) + -2.25)
                v13 := GenS1{F0: 42, F1: 1}
                _ = v12
                _ = v13
            }
        }
        v14 := GenS2{F0: ((("\n\t" + "\n\t") + "foo bar") + ("foo bar" + "hello"))}
        _ = fi7
        _ = v14
    }
    _ = p0
    _ = v6
    return -1
}
func genFn1(p0 uint, p1 int32) uint {
    switch (("a" + "hello") + ((("a" + ((("hello" + "a") + ("foo bar" + "foo bar")) + ((((("\n\t" + ("hello" + "hello")) + "foo bar") + "hello") + "\n\t") + (("\n\t" + (("\n\t" + "") + "hello")) + "a")))) + "foo bar") + "")) {
    case "hello":
        var v15 int = -9223372036854775808
        _ = v15
    case "\n\t":
        v16 := []int{2}
        _ = v16
    case "":
        fmt.Println(p1, p1)
    }
    v17 := []int{-875}
    for fi18 := 0; fi18 < 2; fi18++ {
        for fi19 := 0; fi19 < 6; fi19++ {
            if (694 <= 147) {
                var v20 uint32 = 0
                var v21 int32 = 42
                _ = v20
                _ = v21
            }
            v22 := GenS1{F0: -32768, F1: 0}
            _ = fi19
            _ = v22
        }
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 uint = 18446744073709551615
            fmt.Println(fi18, fi18)
            _ = fi23
            _ = v24
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = v17
    return (p0 << ((18446744073709551615) & 63))
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = (1 >= 18446744073709551615)
    var v3 int = -141
    var v4 int8 = -1
    var v5 int64 = 42
    var v25 int64 = (^(v5))
    v26 := GenS3{F0: 65535}
    v27 := GenS3{F0: 233}
    for fi28 := 0; fi28 < 9; fi28++ {
        v3 -= fi28
        if v2 {
            var v29 int32 = -2147483648
            _ = v29
        } else {
            var v30 uint32 = 1
            var v31 float32 = 0.1
            delete(v1, fi28)
            _ = v30
            _ = v31
        }
        _ = fi28
    }
    switch (v3 >> ((0) & 63)) {
    case 32768:
        var v32 int = -1
        _ = v32
    case -9223372036854775808:
        for fi33 := 0; fi33 < 5; fi33++ {
            v34 := GenS3{F0: 797}
            _ = fi33
            _ = v34
        }
    default:
        v35 := GenS2{F0: "\n\t"}
        _ = v35
    }
    delete(v1, v3)
    v5 = (v5 / ((v5) | 1))
    var v36 uint8 = 42
    delete(v1, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
