package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int64; F1 string; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS4 struct { F0 string }
func genFn0() uint32 {
    var v7 int16 = 1
    v8 := []int{}
    _ = v7
    _ = v8
    return 0
}
func genFn1() string {
    v9 := []int{-2147483648, 1, -9223372036854775808}
    if (!(1 != 42)) {
        v10 := GenS1{F0: 4294967295, F1: 127}
        _ = v10
    }
    if (1 < 0) {
        var v11 int16 = 1000
        _ = v11
    }
    var v12 uint = 4294967295
    _ = len(v9)
    _ = v12
    _ = v9
    return "a"
}

func main() {
    v1 := []int{32767}
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]int)
    v5 := make(map[string]uint32)
    v6 := make(map[uint32]int)
    v13 := GenS2{F0: 1, F1: "\n\t", F2: 56}
    if true {
        var v14 uint = 18446744073709551615
        var v15 uint16 = 927
        v15 -= v15
        _ = v14
        _ = v15
    } else {
        v16 := GenS3{F0: 1, F1: 735, F2: 32767}
        _ = v16
    }
    v17 := GenS4{F0: ("hello" + "foo bar")}
    var v18 int8 = 42
    var v19 uint64 = 2
    v20 := v3[4294967295]
    switch -1 {
    case 0:
        switch "\n\t" {
        case "hello":
            fmt.Println(v18, v18, v20)
        case "\n\t":
            v21 := GenS4{F0: "foo bar"}
            _ = v21
        default:
            for fi22 := 0; fi22 < 11; fi22++ {
                var v23 int8 = -61
                v19 = v19
                _ = fi22
                _ = v23
            }
        }
    case 42:
        if v20 {
            switch "\n\t" {
            case "a":
                var v24 int64 = 692
                _ = v24
            default:
                for fi25 := 0; fi25 < 10; fi25++ {
                    fmt.Println(v18, v20, v18)
                    var v26 string = "hello"
                    _ = fi25
                    _ = v26
                }
            }
        }
    default:
        delete(v6, genFn0())
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        var v28 uint = 828
        switch ((("\n\t" + "") + "a") + "foo bar") {
        case "\n\t":
            var v29 int16 = -487
            _ = v29
        case "a":
            var v30 uint32 = 2
            _ = v30
        }
        v18 = v18
        _ = fi27
        _ = v28
    }
    if len(v1) > 0 {
        v1[0] = -2147483648
    }
    switch -913 {
    case -9223372036854775808:
        var v31 float64 = (1.5 * 3.14159)
        _ = v31
    case 32768:
        for fi32 := 0; fi32 < 15; fi32++ {
            for fi33 := 0; fi33 < 1; fi33++ {
                switch "hello" {
                case "hello":
                    switch (fi32 & 0) {
                    case 32767:
                        v34 := GenS1{F0: v19, F1: v18}
                        _ = v34
                    default:
                        v35 := v2[""]
                        _ = v35
                    }
                case "\n\t":
                    var v36 int32 = 0
                    _ = v36
                case "a":
                    if v20 {
                        if (!v20) {
                            var v37 int16 = -392
                            _ = v37
                        }
                    }
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = fi33
            }
            v38 := GenS1{F0: v19, F1: v18}
            fmt.Println(v19, v18)
            _ = fi32
            _ = v38
        }
    case 42:
        v20 = true
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
