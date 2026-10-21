package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int8, p1 bool) bool {
    p0 = 1
    var v4 float64 = -2.25
    _ = p0
    _ = p1
    _ = v4
    return (0.0 >= 0.1)
}
func genFn1() int {
    switch "a" {
    case "a":
        var v5 uint16 = 231
        _ = v5
    case "":
        var v6 float64 = 1.5
        _ = v6
    }
    var v7 int64 = -9223372036854775808
    _ = v7
    return -1
}

func main() {
    v1 := make(map[string]int)
    var v2 string = (((("foo bar" + "") + "hello") + ("foo bar" + "")) + ((("\n\t" + ("a" + "hello")) + ("foo bar" + ("hello" + "\n\t"))) + "foo bar"))
    var v3 int = -1
    v3 -= (v3 - v3)
    for fi8 := 0; fi8 < 19; fi8++ {
        _ = len(v1)
        if (128 > -23) {
            for fi9 := 0; fi9 < 10; fi9++ {
                var v10 float32 = 1.5
                _ = fi9
                _ = v10
            }
            switch (v2 + v2) {
            case "\n\t":
                fmt.Println(v2, v3, v2)
            case "hello":
                v3 += fi8
            }
        }
        _ = fi8
    }
    v11 := GenS1{F0: (3.14159 > ((-2.25 * 3.14159) + float32((0.1 + 0.1))))}
    if true {
        v12 := GenS1{F0: false}
        v13 := GenS1{F0: true}
        v14 := GenS1{F0: false}
        _ = v12
        _ = v13
        _ = v14
    } else {
        v1[v2] = (v3 >> ((0) & 63))
        v15 := GenS1{F0: true}
        var v16 int32 = -607
        _ = v15
        _ = v16
    }
    switch genFn1() {
    case 32768:
        v17 := []int{9223372036854775807, -416, genFn1()}
        _ = v17
    }
    if false {
        v2 = v2
        v18 := []int{}
        fmt.Println(v2)
        _ = v18
    }
    fmt.Println(v3)
    for fi19 := 0; fi19 < 12; fi19++ {
        for fi20 := 0; fi20 < 1; fi20++ {
            v3 = genFn1()
            v21 := GenS1{F0: true}
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        v3 = v3
        for fi23 := 0; fi23 < 10; fi23++ {
            fmt.Println(fi22, fi22)
            v2 = "hello"
            delete(v1, v2)
            if false {
                v3 = genFn1()
                if true {
                    if true {
                        var v24 uint8 = 226
                        v24 += (v24 ^ v24)
                        if true {
                            var v25 int16 = -32768
                            _ = v25
                        }
                        var v26 bool = true
                        _ = v24
                        _ = v26
                    }
                } else {
                    v27 := GenS1{F0: (float32(0.0) < 1e10)}
                    v28 := []int{v3}
                    _ = len(v1)
                    var v29 int32 = 2147483647
                    var v30 uint16 = 77
                    _ = v27
                    _ = v28
                    _ = v29
                    _ = v30
                }
            }
            var v31 string = "foo bar"
            _ = fi23
            _ = v31
        }
        _ = fi22
    }
    var v32 uint8 = 255
    v3 -= genFn1()
    v33 := GenS1{F0: true}
    v3 = v3
    fmt.Println(v3, v3)
    var v34 int64 = -9223372036854775808
    var v35 int16 = int16(v32)
    var v36 string = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
