package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint32; F2 int }
type GenS4 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS5 struct { F0 int32; F1 int16; F2 uint }
type GenS6 struct { F0 string }
func genFn0(p0 uint8, p1 int8) int32 {
    var v4 int64 = 0
    var v5 float32 = (0.0 + float32(0.1))
    v6 := GenS1{F0: 157, F1: 634}
    fmt.Println(p0)
    v7 := []int{-1, -2147483648}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 982
}

func main() {
    var v1 bool = true
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    if v1 {
        switch "foo bar" {
        case "\n\t":
            v8 := GenS2{F0: ((3.14159 * (1e10 + (float32(0.0) + 0.0))) <= float32(float64(0.1)))}
            _ = v8
        }
        switch "a" {
        case "foo bar":
            for fi9 := 0; fi9 < 1; fi9++ {
                for fi10 := 0; fi10 < 11; fi10++ {
                    v11 := GenS3{F0: -2147483648, F1: 42, F2: fi9}
                    fmt.Println(v1)
                    v12 := []int{fi9, 32767}
                    var v13 int = -2147483648
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                _ = fi9
            }
        case "a":
            if (190 > 1) {
                v1 = (v1 || (-127 == -406))
                if (v1 || v1) {
                    delete(v2, "foo bar")
                    var v14 uint8 = 146
                    v15 := []int{1, 4294967295, -2147483648, 127}
                    _ = v14
                    _ = v15
                } else {
                    v1 = false
                    var v16 uint = 18446744073709551615
                    _ = v16
                }
                var v17 uint16 = 0
                _ = v17
            }
        case "hello":
            v18 := GenS4{F0: 842, F1: 898, F2: 49}
            _ = v18
        default:
            v1 = v1
        }
    } else {
        var v19 int = 127
        if (841 != 42) {
            switch -9223372036854775808 {
            case -9223372036854775808:
                delete(v2, "hello")
            case -184:
                v2[""] = (4294967295 > 2147483647)
            default:
                v20 := []int{}
                _ = v20
            }
            var v21 bool = v1
            _ = v21
        }
        _ = v19
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        _ = len(v2)
        var v23 float32 = 3.14159
        _ = fi22
        _ = v23
    }
    fmt.Println(v1)
    for fi24 := 0; fi24 < 15; fi24++ {
        var v25 bool = v1
        _ = fi24
        _ = v25
    }
    fmt.Println(v1)
    switch 2147483647 {
    case 127:
        v1 = true
    case 4294967295:
        var v26 float32 = 3.14159
        _ = v26
    }
    if (v1 && (v1 && (3.14159 != (-2.25 - ((3.14159 + (float32(1.5) + (0.0 / (0.1 - -2.25)))) - 3.14159))))) {
        if v1 {
            v1 = (0 != -9223372036854775808)
            v27 := GenS5{F0: -509, F1: 1, F2: 324}
            v28 := v3[4294967295]
            _ = v27
            _ = v28
        }
        if true {
            v29 := []int{}
            _ = v29
        }
        for fi30 := 0; fi30 < 5; fi30++ {
            var v31 uint64 = 2
            _ = fi30
            _ = v31
        }
    }
    v1 = v1
    for fi32 := 0; fi32 < 3; fi32++ {
        switch -30 {
        case 42:
            for fi33 := 0; fi33 < 5; fi33++ {
                v1 = v1
                _ = fi33
            }
        case 854:
            if true {
                var v34 int16 = 32767
                _ = v34
            }
        default:
            for fi35 := 0; fi35 < 5; fi35++ {
                v1 = false
                _ = fi35
            }
        }
        _ = fi32
    }
    v1 = (-2147483648 != 0)
    var v36 int = 128
    var v37 uint64 = 2
    v36 = v36
    if false {
        v1 = v1
    } else {
        if true {
            for fi38 := 0; fi38 < 1; fi38++ {
                v39 := GenS6{F0: ("a" + ("\n\t" + (("" + "foo bar") + "a")))}
                fmt.Println(v36)
                _ = fi38
                _ = v39
            }
        } else {
            v1 = (0.0 > 3.14159)
        }
    }
    var v40 int = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v40
}
