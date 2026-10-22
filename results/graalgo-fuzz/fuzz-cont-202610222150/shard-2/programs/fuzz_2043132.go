package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 bool }

func main() {
    var v1 float64 = 3.14159
    var v2 bool = (((0 > 1) && (295 <= 116)) || true)
    var v3 bool = (!false)
    v1 -= 0.1
    v4 := GenS1{F0: 42}
    v3 = v2
    fmt.Println(v1)
    v3 = v2
    var v5 uint64 = 285
    v6 := []int{32767}
    switch (("a" + "") + "\n\t") {
    case "\n\t":
        v2 = true
    }
    var v7 bool = (!false)
    if (v2 || false) {
        var v8 uint32 = 2
        _ = v8
    } else {
        v9 := GenS2{F0: 4294967295, F1: v5}
        _ = len(v6)
        fmt.Println(v3, v2)
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 uint64 = 2
            var v13 int16 = 32767
            var v14 uint64 = v12
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        v3 = v2
        _ = fi10
    }
    if (!(("hello" + "hello") < "\n\t")) {
        v2 = false
        v15 := GenS3{F0: 32767}
        if (82 > 42) {
            _ = len(v6)
        }
        _ = v15
    }
    v7 = v2
    if ((v2 || v3) && v2) {
        var v16 uint = 255
        for fi17 := 0; fi17 < 19; fi17++ {
            v18 := GenS3{F0: 408}
            if (!false) {
                var v19 float64 = (1.5 / ((((float64((3.14159 * float32(float64(((-2.25 / 1e10) + 0.0))))) * v1) + v1) * v1) - (v1 - v1)))
                _ = v19
            }
            _ = fi17
            _ = v18
        }
        _ = v16
    } else {
        if true {
            if (32767 != 71) {
                var v20 uint = 65535
                for fi21 := 0; fi21 < 14; fi21++ {
                    v20 -= v20
                    v2 = v3
                    _ = fi21
                }
                v20 += 815
                _ = v20
            }
            v6 = append(v6, int(v5))
            v1 = 3.14159
        } else {
            if v7 {
                fmt.Println(v7)
            } else {
                v3 = (42 > 2147483647)
            }
        }
        for fi22 := 0; fi22 < 11; fi22++ {
            fmt.Println(v7)
            var v23 int32 = 223
            switch "foo bar" {
            case "foo bar":
                v5 -= (v5 % ((v5) | 1))
            case "hello":
                switch fi22 {
                case 806:
                    if v3 {
                        var v24 uint32 = 899
                        v25 := GenS4{F0: (v23 % ((42) | 1)), F1: v7}
                        v6 = append(v6, fi22)
                        _ = v24
                        _ = v25
                    } else {
                        var v26 float32 = 3.14159
                        v27 := GenS2{F0: 350, F1: v5}
                        _ = v26
                        _ = v27
                    }
                case 4294967295:
                    var v28 int = 759
                    _ = v28
                default:
                    v23 = v23
                }
            case "":
                if true {
                    var v29 float32 = -2.25
                    v30 := GenS3{F0: 763}
                    var v31 uint64 = v5
                    _ = v29
                    _ = v30
                    _ = v31
                }
            }
            _ = fi22
            _ = v23
        }
        fmt.Println(v5, v3)
    }
    var v32 int32 = -1
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
