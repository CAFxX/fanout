package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint }
type GenS3 struct { F0 int; F1 int32 }
func genFn0() int16 {
    var v4 string = "\n\t"
    switch "a" {
    case "hello":
        var v5 float32 = 0.0
        _ = v5
    case "foo bar":
        v6 := GenS1{F0: -2147483648, F1: 1}
        _ = v6
    default:
        switch v4 {
        case "\n\t":
            if false {
                if false {
                    if false {
                        v4 = (("" + v4) + (v4 + ("hello" + v4)))
                        for fi7 := 0; fi7 < 12; fi7++ {
                            v8 := []int{fi7, 1, (fi7 & fi7)}
                            _ = fi7
                            _ = v8
                        }
                    } else {
                        switch "\n\t" {
                        case "hello":
                            v4 = "\n\t"
                        }
                        fmt.Println(v4)
                    }
                    v9 := GenS2{F0: 100, F1: 128, F2: 42}
                    _ = v9
                } else {
                    v10 := []int{-1, -229}
                    for fi11 := 0; fi11 < 16; fi11++ {
                        if true {
                            for fi12 := 0; fi12 < 15; fi12++ {
                                v4 = "a"
                                v13 := GenS1{F0: fi11, F1: 9223372036854775807}
                                var v14 float64 = float64(0.1)
                                v4 = (v4 + v4)
                                fmt.Println(v4)
                                _ = fi12
                                _ = v13
                                _ = v14
                            }
                        }
                        _ = fi11
                    }
                    v4 = v4
                    _ = v10
                }
                for fi15 := 0; fi15 < 5; fi15++ {
                    switch fi15 {
                    case 32767:
                        for fi16 := 0; fi16 < 19; fi16++ {
                            v17 := GenS3{F0: fi16, F1: -2147483648}
                            _ = fi16
                            _ = v17
                        }
                    case -332:
                        v4 = v4
                    }
                    var v18 int16 = 32767
                    for fi19 := 0; fi19 < 7; fi19++ {
                        v4 = v4
                        _ = fi19
                    }
                    _ = fi15
                    _ = v18
                }
            }
        case "foo bar":
            v20 := []int{}
            _ = v20
        }
    }
    var v21 uint16 = 650
    for fi22 := 0; fi22 < 5; fi22++ {
        v23 := GenS3{F0: fi22, F1: 2147483647}
        fmt.Println(v4, fi22, v21)
        _ = fi22
        _ = v23
    }
    _ = v21
    _ = v4
    return 660
}
func genFn1(p0 int) uint64 {
    v24 := GenS2{F0: -1, F1: 128, F2: 65535}
    if (2147483647 == 2147483647) {
        switch 9223372036854775807 {
        case 1:
            switch (("a" + "\n\t") + "") {
            case "":
                v25 := GenS1{F0: (p0 + 9223372036854775807), F1: 9223372036854775807}
                _ = v25
            case "\n\t":
                fmt.Println(p0)
            }
        }
        v26 := GenS1{F0: (-(p0)), F1: 1}
        _ = v26
    }
    _ = p0
    _ = v24
    return 58
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    v3 := []uint32{}
    var v27 string = "foo bar"
    for fi28 := 0; fi28 < 4; fi28++ {
        var v29 float32 = float32(0.1)
        var v30 uint8 = 70
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 bool = (!true)
    var v32 float32 = -2.25
    v33 := GenS1{F0: 2147483647, F1: -9223372036854775808}
    for fi34 := 0; fi34 < 8; fi34++ {
        var v35 int8 = 0
        _ = fi34
        _ = v35
    }
    var v36 float64 = (((float64(1.5) / -2.25) - 0.0) * 1e10)
    fmt.Println(v32)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v36
}
