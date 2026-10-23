package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint8; F2 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 int32 }
func genFn0(p0 string, p1 int, p2 int64) int16 {
    v6 := GenS1{F0: 931}
    var v7 float32 = ((float32(0.1) / float32((float64(3.14159) + 1.5))) - (1.5 - 0.0))
    var v8 int16 = -953
    v7 -= v7
    p1 -= 32768
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return v8
}

func main() {
    var v1 uint = 315
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    var v4 uint = (v1 * (^(v1)))
    var v5 float64 = -2.25
    v9 := v2[-9223372036854775808]
    switch "hello" {
    case "foo bar":
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v4, v4, v4)
        default:
            var v10 int8 = int8(v4)
            _ = v10
        }
    case "\n\t":
        if (!(42 <= -1)) {
            v1 += v4
            fmt.Println(v5, v5)
            v4 += v1
        }
    case "a":
        v9 -= 32767
    default:
        fmt.Println(v5)
    }
    if false {
        v5 += float64(-2.25)
    } else {
        switch ("foo bar" + ("hello" + "a")) {
        case "":
            var v11 int8 = -1
            _ = v11
        case "hello":
            for fi12 := 0; fi12 < 16; fi12++ {
                v1 = v4
                v1 -= (v4 * 4294967295)
                v4 += (v1 / (((^(v4))) | 1))
                _ = fi12
            }
        default:
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 uint = v1
                _ = fi13
                _ = v14
            }
        }
        for fi15 := 0; fi15 < 14; fi15++ {
            v16 := []int{9223372036854775807, (v9 % ((fi15) | 1))}
            for fi17 := 0; fi17 < 8; fi17++ {
                fmt.Println(v5, v9, v4)
                var v18 float64 = v5
                var v19 uint8 = 0
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = fi15
            _ = v16
        }
    }
    v4 = (v1 - (v1 ^ 4294967295))
    if (275 <= 798) {
        for fi20 := 0; fi20 < 8; fi20++ {
            for fi21 := 0; fi21 < 2; fi21++ {
                v22 := GenS2{F0: 1, F1: 178, F2: 42}
                v23 := []int{(fi21 / ((fi21) | 1)), -416, fi21, -290}
                v24 := GenS3{F0: "foo bar"}
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            v25 := []int{fi20, v9}
            var v26 int64 = 9223372036854775807
            _ = fi20
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 10; fi27++ {
            v1 = v1
            for fi28 := 0; fi28 < 9; fi28++ {
                v4 += (v4 | (v1 << (((v4 ^ v4)) & 63)))
                var v29 int8 = -1
                if false {
                    v30 := GenS1{F0: (^(v4))}
                    switch ("\n\t" + ("a" + "\n\t")) {
                    case "foo bar":
                        v29 = v29
                    }
                    _ = v30
                } else {
                    v31 := GenS4{F0: 77}
                    v29 += v29
                    v32 := []int{fi27, (v9 % (((fi27 & (fi27 * (fi27 << ((v1) & 63))))) | 1)), int(v1)}
                    _ = v31
                    _ = v32
                }
                v33 := v3[66]
                _ = fi28
                _ = v29
                _ = v33
            }
            _ = fi27
        }
    } else {
        var v34 int = v9
        _ = v34
    }
    if true {
        v35 := GenS5{F0: (v9 % ((v9) | 1)), F1: 1}
        _ = v35
    }
    switch "" {
    case "foo bar":
        v1 = v1
    case "a":
        for fi36 := 0; fi36 < 12; fi36++ {
            _ = len(v2)
            var v37 int32 = 2147483647
            v38 := v2[(fi36 & fi36)]
            _ = fi36
            _ = v37
            _ = v38
        }
    }
    fmt.Println(v1)
    v9 += (v9 >> (((v1 << ((908) & 63))) & 63))
    var v39 uint = 18446744073709551615
    if true {
        fmt.Println(v39)
    } else {
        switch (("hello" + ("hello" + "\n\t")) + "\n\t") {
        case "foo bar":
            v2[-1] = (-(v9))
        case "a":
            v4 = 545
        }
    }
    v40 := []int{v9, v9, (v9 & -985)}
    v5 = 1e10
    var v41 string = "a"
    fmt.Println(v9, v9)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v9
}
