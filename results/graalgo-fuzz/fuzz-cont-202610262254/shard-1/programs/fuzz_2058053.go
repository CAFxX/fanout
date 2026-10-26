package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint8, p1 int32, p2 int8) string {
    p2 += 127
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := []int{fi6, fi6}
        v8 := GenS1{F0: 1}
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        switch "hello" {
        case "\n\t":
            var v10 uint32 = 42
            _ = v10
        }
        _ = fi9
    }
    switch "hello" {
    case "":
        var v11 int8 = -128
        _ = v11
    case "hello":
        switch (("a" + ("hello" + "a")) + "") {
        case "hello":
            v12 := GenS1{F0: 42}
            _ = v12
        case "\n\t":
            v13 := GenS2{F0: 42, F1: ("\n\t" + ("\n\t" + "foo bar"))}
            _ = v13
        case "":
            var v14 uint = 2
            _ = v14
        }
    default:
        fmt.Println(p2, p2)
    }
    fmt.Println(p1, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    return ("a" + ("a" + "hello"))
}
func genFn1(p0 uint8, p1 string, p2 int) int8 {
    fmt.Println(p2, p2, p0)
    var v15 string = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 1
}

func main() {
    var v1 bool = (!(65535 >= 65535))
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    v4 := make(map[uint32]string)
    var v5 uint8 = 42
    fmt.Println(v1, v1)
    if (true || v1) {
        v16 := GenS2{F0: 554, F1: "a"}
        _ = v16
    }
    v17 := GenS1{F0: uint16(v5)}
    v1 = (-125 > 32767)
    var v18 uint32 = 2
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := GenS3{F0: -597}
        _ = fi19
        _ = v20
    }
    delete(v3, ("" + "hello"))
    v21 := GenS3{F0: -481}
    v22 := GenS4{F0: 65535, F1: -43}
    if false {
        for fi23 := 0; fi23 < 17; fi23++ {
            v24 := []int{}
            _ = fi23
            _ = v24
        }
        var v25 uint8 = (v5 + 170)
        var v26 uint32 = v18
        _ = v25
        _ = v26
    }
    switch "a" {
    case "a":
        v1 = v1
    case "\n\t":
        fmt.Println(v18, v18, v18)
    case "foo bar":
        for fi27 := 0; fi27 < 5; fi27++ {
            fmt.Println(fi27)
            if v1 {
                v28 := []int{(fi27 | fi27), fi27, (fi27 ^ fi27)}
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
                if v1 {
                    var v29 uint64 = 0
                    _ = v29
                } else {
                    v30 := GenS3{F0: -2147483648}
                    if ((fi27 % ((fi27) | 1)) >= fi27) {
                        v18 = 65535
                        v1 = (false && (42 < 2147483647))
                    }
                    _ = v30
                }
                _ = v28
            } else {
                v5 += uint8(v18)
                for fi31 := 0; fi31 < 12; fi31++ {
                    v32 := []int{}
                    fmt.Println(v5, v5, fi31)
                    var v33 float64 = 0.0
                    var v34 int8 = 42
                    _ = fi31
                    _ = v32
                    _ = v33
                    _ = v34
                }
                v35 := GenS2{F0: 333, F1: ""}
                _ = v35
            }
            _ = fi27
        }
    }
    var v36 uint32 = (v18 + uint32(v5))
    var v37 uint8 = v5
    v5 = (v5 / ((v37) | 1))
    v36 = v18
    if false {
        fmt.Println(v18, v1, v36)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
