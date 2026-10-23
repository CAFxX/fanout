package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint; F1 bool }
func genFn0(p0 uint16) int16 {
    var v5 int64 = -9223372036854775808
    var v6 int32 = 48
    var v7 float32 = ((1e10 - -2.25) - 0.1)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return -719
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -32768
    v3 := make(map[uint32]bool)
    var v4 int8 = 0
    for fi8 := 0; fi8 < 8; fi8++ {
        if false {
            v9 := GenS1{F0: -32768}
            if true {
                for fi10 := 0; fi10 < 12; fi10++ {
                    switch ("a" + "") {
                    case "foo bar":
                        v2 = (v2 | 971)
                    case "":
                        _ = len(v3)
                    }
                    _ = fi10
                }
                fmt.Println(v2)
            } else {
                v11 := GenS2{F0: 683}
                _ = v11
            }
            v1[1] = "a"
            _ = v9
        }
        if true {
            for fi12 := 0; fi12 < 18; fi12++ {
                v4 += (v4 & -128)
                _ = fi12
            }
            switch fi8 {
            case 0:
                var v13 float64 = (0.0 * ((0.1 * -2.25) * (float64(-2.25) + float64(3.14159))))
                _ = v13
            default:
                var v14 uint32 = 42
                _ = v14
            }
            v15 := GenS3{F0: (^(fi8)), F1: "foo bar"}
            _ = v15
        } else {
            var v16 bool = false
            v16 = ((v16 && v16) && ((1.5 - 1e10) > 1e10))
            _ = v16
        }
        v2 = 1
        _ = fi8
    }
    v17 := []int{}
    switch "a" {
    case "\n\t":
        if true {
            switch 128 {
            case 32768:
                var v18 float64 = (float64(1.5) * 1.5)
                _ = v18
            }
            switch 9223372036854775807 {
            case 32768:
                var v19 string = "foo bar"
                _ = v19
            default:
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
            }
        } else {
            v20 := GenS2{F0: 381}
            v2 = v2
            _ = v20
        }
    case "hello":
        v21 := GenS4{F0: 1, F1: int64(v4)}
        _ = v21
    }
    if (793 == 4294967295) {
        for fi22 := 0; fi22 < 9; fi22++ {
            switch ("\n\t" + "") {
            case "":
                v23 := GenS5{F0: 42}
                _ = v23
            case "a":
                v2 = (v2 & v2)
            case "foo bar":
                if (147 <= 1) {
                    var v24 int16 = genFn0(26)
                    _ = v24
                }
            }
            _ = fi22
        }
        _ = len(v17)
        _ = len(v1)
    }
    if len(v17) > 0 {
        v17[0] = 2147483647
    }
    if (2147483648 <= 42) {
        v25 := GenS4{F0: 1, F1: 42}
        for fi26 := 0; fi26 < 17; fi26++ {
            if true {
                v27 := v1[563]
                v17 = append(v17, fi26)
                _ = v27
            }
            _ = fi26
        }
        var v28 uint8 = 71
        _ = v25
        _ = v28
    } else {
        v29 := GenS2{F0: 65535}
        _ = v29
    }
    var v30 int = -817
    for fi31 := 0; fi31 < 3; fi31++ {
        if true {
            if true {
                var v32 float32 = 0.0
                delete(v3, 0)
                _ = v32
            }
        }
        _ = fi31
    }
    fmt.Println(v30, v2)
    fmt.Println(v2)
    for fi33 := 0; fi33 < 16; fi33++ {
        v4 = (-(v4))
        var v34 uint64 = 172
        var v35 uint16 = 2
        _ = fi33
        _ = v34
        _ = v35
    }
    v36 := GenS6{F0: 65535, F1: (true && ((^(v2)) >= genFn0(836)))}
    fmt.Println(v4)
    var v37 int16 = v2
    v38 := GenS4{F0: 0, F1: -9223372036854775808}
    var v39 uint32 = 0
    var v40 uint32 = v39
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
