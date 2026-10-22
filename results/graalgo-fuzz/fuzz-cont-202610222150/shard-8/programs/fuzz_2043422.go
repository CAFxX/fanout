package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    var v3 int16 = 0
    var v4 int32 = -534
    var v5 bool = false
    fmt.Println(v5, v4, v5)
    var v6 int16 = -510
    fmt.Println(v6)
    v6 -= v3
    v2[((((("foo bar" + ("\n\t" + ("\n\t" + ""))) + "") + "") + "hello") + "\n\t")] = v5
    v7 := v1["\n\t"]
    fmt.Println(v4, v4, v6)
    if (-9223372036854775808 >= 42) {
        for fi8 := 0; fi8 < 10; fi8++ {
            v3 = v6
            for fi9 := 0; fi9 < 3; fi9++ {
                v10 := GenS1{F0: 42, F1: 1}
                var v11 float64 = -2.25
                v6 = v3
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
        fmt.Println(v7, v6)
    } else {
        switch -1 {
        case -891:
            switch v7 {
            case "hello":
                v12 := GenS1{F0: 4294967295, F1: v4}
                _ = v12
            case "foo bar":
                v13 := []int{9223372036854775807, 128}
                _ = v13
            case "a":
                var v14 bool = v5
                _ = v14
            }
        case -9223372036854775808:
            if v5 {
                var v15 uint16 = 32767
                _ = v15
            } else {
                for fi16 := 0; fi16 < 3; fi16++ {
                    v17 := v1[v7]
                    _ = fi16
                    _ = v17
                }
            }
        case 327:
            v5 = (127 <= 20)
        }
    }
    v18 := GenS1{F0: 855, F1: (v4 * 133)}
    v19 := GenS1{F0: 0, F1: -2147483648}
    if false {
        v20 := GenS1{F0: 233, F1: 2147483647}
        v6 -= 32767
        _ = v20
    } else {
        switch 1 {
        case 0:
            v21 := GenS1{F0: 156, F1: (v4 % ((42) | 1))}
            _ = v21
        case 9223372036854775807:
            var v22 float64 = (3.14159 * 0.1)
            _ = v22
        case 1:
            var v23 int = -807
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        switch fi24 {
        case -1:
            for fi25 := 0; fi25 < 11; fi25++ {
                fmt.Println(v5)
                var v26 int8 = 84
                _ = fi25
                _ = v26
            }
        default:
            fmt.Println(fi24, v3, fi24)
        }
        v7 = v7
        _ = fi24
    }
    v27 := GenS1{F0: 909, F1: (v4 % ((v4) | 1))}
    fmt.Println(v5)
    var v28 int16 = v6
    v5 = (!v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
