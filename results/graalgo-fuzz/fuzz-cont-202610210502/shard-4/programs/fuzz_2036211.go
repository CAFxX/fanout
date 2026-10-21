package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int16) uint32 {
    v5 := []int{-2147483648, 4294967295, 2147483647}
    var v6 int64 = -147
    v7 := GenS1{F0: -627, F1: 584}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 string, p1 uint64) int32 {
    switch p0 {
    case "a":
        p0 = (p0 + p0)
    case "\n\t":
        if (p0 > p0) {
            if false {
                p0 = p0
                var v8 float64 = float64(0.0)
                _ = v8
            } else {
                for fi9 := 0; fi9 < 20; fi9++ {
                    var v10 string = ("hello" + (p0 + "hello"))
                    v11 := []int{fi9, fi9}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
            fmt.Println(p1)
            v12 := []int{}
            _ = v12
        } else {
            var v13 int32 = 0
            v14 := GenS1{F0: -1, F1: 255}
            if (1000 == 1) {
                fmt.Println(p0)
                switch (((("a" + p0) + p0) + "\n\t") + (p0 + ("a" + ("\n\t" + "hello")))) {
                case "a":
                    fmt.Println(p0, p0, p0)
                case "hello":
                    if true {
                        fmt.Println(p0)
                        v15 := GenS2{F0: 42, F1: ((false || ((!true) || ((-181 > 224) && (279 >= 1000)))) || false)}
                        _ = v15
                    }
                case "":
                    v13 -= (v13 ^ (v13 | 42))
                }
            } else {
                v13 = v13
                var v16 int = -2147483648
                if true {
                    p0 = "foo bar"
                    var v17 float64 = ((float64(float32(((0.0 * 1.5) * 3.14159))) + 0.0) - 1.5)
                    _ = v17
                }
                _ = v16
            }
            _ = v13
            _ = v14
        }
    }
    v18 := []int{-383, 672}
    var v19 uint = 65535
    fmt.Println(p1)
    if (-1 == -128) {
        var v20 bool = (3.14159 >= 0.1)
        v21 := GenS1{F0: 658, F1: 2}
        for fi22 := 0; fi22 < 7; fi22++ {
            var v23 uint8 = 21
            var v24 uint64 = 1
            var v25 float32 = 0.1
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        _ = v20
        _ = v21
    } else {
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        if true {
            for fi26 := 0; fi26 < 15; fi26++ {
                fmt.Println(fi26, fi26)
                v19 = v19
                _ = fi26
            }
            v27 := GenS1{F0: 42, F1: 0}
            _ = v27
        }
        var v28 bool = (26 < 255)
        _ = v28
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return -2147483648
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 42
    v3 := make(map[int]string)
    var v4 int8 = 1
    switch 0 {
    case 0:
        v29 := []int{127, 0, -9223372036854775808}
        _ = v29
    }
    fmt.Println(v2, v4)
    for fi30 := 0; fi30 < 10; fi30++ {
        v4 = (v4 - (v4 / (((v4 / ((v4) | 1))) | 1)))
        v31 := GenS1{F0: -317, F1: 186}
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        v2 = (v2 | v2)
        _ = fi32
    }
    for fi33 := 0; fi33 < 3; fi33++ {
        v2 = 65535
        _ = fi33
    }
    if true {
        for fi34 := 0; fi34 < 8; fi34++ {
            v4 = -1
            if false {
                var v35 string = "foo bar"
                v36 := []int{}
                _ = v35
                _ = v36
            } else {
                v37 := GenS3{F0: 18446744073709551615}
                _ = v37
            }
            v38 := []int{(fi34 ^ 42), fi34, 32767, 496}
            _ = fi34
            _ = v38
        }
    }
    v4 += (-(v4))
    v2 += (v2 * (v2 / ((v2) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
