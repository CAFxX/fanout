package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0() uint8 {
    v5 := []int{42, 4294967295}
    switch 2147483647 {
    case 4294967295:
        var v6 uint8 = 231
        _ = v6
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if true {
        var v7 int16 = 0
        _ = v7
    }
    switch "hello" {
    case "hello":
        v8 := GenS1{F0: false, F1: 2147483648}
        _ = v8
    }
    _ = v5
    return 0
}

func main() {
    var v1 bool = true
    var v2 bool = (0.0 > ((-2.25 * -2.25) - (-2.25 / float32(((1e10 / 1e10) - 0.0)))))
    v3 := make(map[int]uint32)
    v4 := []int{-1, 127, 32768, 127}
    var v9 uint64 = 18446744073709551615
    fmt.Println(v1, v9, v9)
    var v10 bool = (0 > -2147483648)
    v1 = v10
    switch ("hello" + ("a" + ("foo bar" + "hello"))) {
    case "":
        switch (("\n\t" + "hello") + "foo bar") {
        case "foo bar":
            for fi11 := 0; fi11 < 11; fi11++ {
                for fi12 := 0; fi12 < 9; fi12++ {
                    v13 := GenS1{F0: ((v10 && true) && ("\n\t" == "")), F1: 2}
                    _ = fi12
                    _ = v13
                }
                _ = fi11
            }
        }
    case "a":
        v2 = v1
    case "foo bar":
        v14 := GenS2{F0: 782, F1: 2}
        _ = v14
    }
    v15 := GenS2{F0: 695, F1: 0}
    fmt.Println(v10, v1)
    var v16 bool = false
    if true {
        v3[-354] = 2147483647
    } else {
        var v17 int32 = -2147483648
        _ = v17
    }
    var v18 float64 = (((3.14159 + 3.14159) + 1.5) / 1.0)
    v16 = true
    for fi19 := 0; fi19 < 6; fi19++ {
        var v20 int = (-(fi19))
        v21 := GenS1{F0: v10, F1: 42}
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 uint16 = 1
        v18 += (float64(0.1) / v18)
        _ = fi22
        _ = v23
    }
    v24 := GenS2{F0: 65535, F1: uint16(v9)}
    var v25 bool = true
    v10 = v25
    if v10 {
        v26 := GenS3{F0: (-1 == 0), F1: 65535}
        v3[128] = 1
        _ = v26
    } else {
        v4 = append(v4, 128)
        var v27 uint64 = 219
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v9
}
