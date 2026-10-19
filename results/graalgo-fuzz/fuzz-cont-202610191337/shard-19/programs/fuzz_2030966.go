package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 uint32; F1 int32; F2 int64 }
func genFn0(p0 bool, p1 int16) uint16 {
    var v7 int8 = -2
    v8 := GenS1{F0: -1}
    v9 := []int{0}
    v10 := GenS2{F0: 4294967295}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 255
}
func genFn1(p0 int8, p1 int8) uint64 {
    switch 32767 {
    case 2147483647:
        var v11 float32 = 0.1
        _ = v11
    }
    var v12 uint16 = genFn0((42 == 255), 42)
    _ = p0
    _ = p1
    _ = v12
    return 666
}

func main() {
    var v1 uint = 0
    var v2 uint8 = 1
    var v3 uint8 = 165
    v4 := []string{}
    var v5 float32 = -2.25
    v6 := []uint32{}
    switch "foo bar" {
    case "foo bar":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case "hello":
        v13 := GenS2{F0: 128}
        _ = v13
    default:
        if len(v6) > 0 {
            v6[0] = 2
        }
    }
    if false {
        if false {
            v6 = append(v6, 543)
            v2 = (v2 & v2)
            var v14 string = "hello"
            _ = v14
        } else {
            v15 := GenS3{F0: -112, F1: 18446744073709551615, F2: 100}
            var v16 uint32 = 334
            if (42 == 0) {
                var v17 float32 = (float32(0.1) * float32((3.14159 - 0.0)))
                _ = v17
            }
            _ = v15
            _ = v16
        }
        v18 := GenS2{F0: 625}
        v19 := GenS2{F0: 42}
        _ = v18
        _ = v19
    } else {
        var v20 int8 = -39
        var v21 uint64 = 4294967295
        switch ("hello" + "hello") {
        case "foo bar":
            for fi22 := 0; fi22 < 16; fi22++ {
                for fi23 := 0; fi23 < 4; fi23++ {
                    fmt.Println(fi22)
                    _ = fi23
                }
                if len(v6) > 0 {
                    v6[0] = uint32(v21)
                }
                fmt.Println(v1, v3, v2)
                _ = fi22
            }
        }
        _ = v20
        _ = v21
    }
    var v24 float64 = (((0.0 * (3.14159 * 3.14159)) + 1.5) * -2.25)
    v25 := GenS3{F0: 0, F1: 279, F2: -6}
    var v26 int = 32767
    for fi27 := 0; fi27 < 12; fi27++ {
        v28 := GenS4{F0: 18446744073709551615, F1: 0}
        _ = fi27
        _ = v28
    }
    v29 := GenS1{F0: -528}
    v30 := GenS3{F0: -79, F1: 1, F2: 127}
    switch "a" {
    case "a":
        fmt.Println(v3, v1, v5)
    case "foo bar":
        var v31 string = (("a" + "") + "\n\t")
        _ = v31
    }
    var v32 int32 = -1
    if len(v4) > 0 {
        v4[0] = ((((("foo bar" + (("a" + "hello") + "")) + "foo bar") + "") + (("\n\t" + "a") + (((("" + "foo bar") + "") + ((("foo bar" + "\n\t") + (("\n\t" + "\n\t") + "\n\t")) + ("a" + ("a" + ("hello" + (("\n\t" + ("foo bar" + "foo bar")) + "foo bar")))))) + ("a" + ("foo bar" + "\n\t"))))) + "a")
    }
    v2 = (v2 * (v3 | v3))
    _ = len(v6)
    if false {
        var v33 int64 = 439
        _ = v33
    } else {
        v34 := GenS1{F0: -1}
        v35 := GenS5{F0: -913, F1: (v1 - v1)}
        _ = v34
        _ = v35
    }
    switch "foo bar" {
    case "a":
        v36 := GenS3{F0: -97, F1: 300, F2: -114}
        _ = v36
    default:
        var v37 float64 = float64(v5)
        _ = v37
    }
    fmt.Println(v24, v3, v26)
    v38 := GenS6{F0: 42, F1: v32, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
