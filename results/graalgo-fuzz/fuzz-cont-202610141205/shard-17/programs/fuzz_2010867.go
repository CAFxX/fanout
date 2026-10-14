package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint64 }
func genFn0() string {
    v6 := GenS1{F0: 42, F1: 791}
    v7 := []int{}
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 float64 = float64(0.1)
        _ = fi8
        _ = v9
    }
    if (0.0 != 1e10) {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    switch ("\n\t" + "foo bar") {
    case "hello":
        if (!(921 == 4294967295)) {
            var v10 int64 = 42
            v10 = 1
            _ = v10
        }
    case "a":
        var v11 bool = (true || true)
        _ = v11
    case "":
        var v12 int = 32767
        _ = v12
    default:
        var v13 string = ""
        _ = v13
    }
    _ = v6
    _ = v7
    return (("" + (("hello" + "a") + "\n\t")) + ("hello" + "a"))
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 1
    v3 := make(map[string]int)
    var v4 int32 = -304
    var v5 string = ("hello" + "foo bar")
    var v14 int = 910
    v15 := GenS2{F0: (v2 + 605)}
    v16 := GenS1{F0: 0, F1: 77}
    v4 = 42
    v1[42] = uint32(v2)
    if false {
        var v17 float32 = (float32(1.5) + (1e10 - (1.5 / (0.1 - -2.25))))
        var v18 uint16 = 839
        _ = v17
        _ = v18
    }
    v19 := GenS2{F0: 799}
    switch int(v4) {
    case 0:
        var v20 uint16 = 443
        _ = v20
    case 1:
        for fi21 := 0; fi21 < 9; fi21++ {
            switch (fi21 ^ (fi21 | (v14 ^ fi21))) {
            case 1:
                switch fi21 {
                case -9223372036854775808:
                    v4 += (v4 % ((v4) | 1))
                case 1:
                    var v22 uint32 = 312
                    _ = v22
                case 4294967295:
                    switch 128 {
                    case 780:
                        for fi23 := 0; fi23 < 7; fi23++ {
                            if (v2 < 18446744073709551615) {
                                var v24 uint = uint(v2)
                                v5 = v5
                                var v25 string = "foo bar"
                                _ = v24
                                _ = v25
                            }
                            v26 := []int{-9223372036854775808}
                            v27 := GenS1{F0: 0, F1: 863}
                            _ = fi23
                            _ = v26
                            _ = v27
                        }
                    case 4294967295:
                        v1[472] = 634
                    case -782:
                        v2 -= (^(v2))
                    }
                }
            case 246:
                var v28 uint16 = 42
                _ = v28
            default:
                var v29 int32 = 0
                _ = v29
            }
            var v30 int64 = -9223372036854775808
            v5 = "hello"
            _ = fi21
            _ = v30
        }
    case 706:
        v31 := []int{1, v14, int(v4)}
        _ = v31
    default:
        v32 := []int{(^(v14)), v14}
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
