package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int }
func genFn0(p0 uint32) uint32 {
    p0 = 65535
    var v6 uint32 = p0
    v7 := GenS1{F0: 0}
    _ = p0
    _ = v6
    _ = v7
    return (^(p0))
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = false
    v3 := make(map[string]uint32)
    v4 := make(map[int]uint32)
    v5 := make(map[string]uint32)
    if false {
        var v8 bool = (32768 >= -353)
        var v9 int8 = 0
        v10 := []int{-997, 0, 0}
        _ = v10
        _ = v8
        _ = v9
    }
    if v2 {
        for fi11 := 0; fi11 < 10; fi11++ {
            for fi12 := 0; fi12 < 8; fi12++ {
                v13 := GenS2{F0: "foo bar", F1: -2147483648}
                v2 = true
                var v14 int8 = -52
                fmt.Println(v14, fi11, v14)
                if false {
                    if (459 == 1) {
                        v2 = ((v14 - 127) != v14)
                    }
                    if (true && ((4294967295 >= 18446744073709551615) || true)) {
                        fmt.Println(fi11, v2)
                        fmt.Println(v14, fi11, v2)
                        var v15 uint8 = 51
                        _ = v15
                    } else {
                        v16 := []int{-2147483648, fi11}
                        var v17 int8 = v14
                        _ = v16
                        _ = v17
                    }
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            v2 = v2
            var v18 float64 = 0.1
            _ = fi11
            _ = v18
        }
        v19 := GenS3{F0: 128}
        _ = v19
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        for fi21 := 0; fi21 < 5; fi21++ {
            delete(v1, 2147483648)
            _ = fi21
        }
        _ = fi20
    }
    switch "a" {
    case "\n\t":
        if v2 {
            var v22 int16 = 782
            var v23 uint16 = 532
            fmt.Println(v23, v2, v23)
            _ = v22
            _ = v23
        }
    case "foo bar":
        var v24 int64 = 9223372036854775807
        _ = v24
    case "hello":
        for fi25 := 0; fi25 < 2; fi25++ {
            v2 = (-9223372036854775808 > -9223372036854775808)
            _ = fi25
        }
    }
    v2 = v2
    v2 = false
    v2 = true
    fmt.Println(v2)
    v1[2147483648] = v2
    switch -86 {
    case 9223372036854775807:
        var v26 int32 = 1
        _ = v26
    }
    switch 32768 {
    case 749:
        switch "\n\t" {
        case "hello":
            var v27 uint64 = 0
            _ = v27
        default:
            v28 := v1[533]
            _ = v28
        }
    case 2147483647:
        if v2 {
            delete(v3, ("\n\t" + ""))
            v29 := GenS2{F0: "", F1: -328}
            _ = v29
        } else {
            for fi30 := 0; fi30 < 12; fi30++ {
                var v31 int64 = 42
                fmt.Println(v31, v31)
                v32 := GenS2{F0: "", F1: 2147483647}
                _ = fi30
                _ = v31
                _ = v32
            }
            var v33 bool = (-2147483648 != 2147483647)
            v34 := v4[4294967295]
            _ = v33
            _ = v34
        }
    case 32767:
        var v35 uint8 = 0
        _ = v35
    }
    fmt.Println(v2)
    for fi36 := 0; fi36 < 1; fi36++ {
        if (!(v2 && false)) {
            v37 := GenS2{F0: "a", F1: 726}
            var v38 uint64 = 237
            _ = v37
            _ = v38
        } else {
            fmt.Println(v2, v2)
            for fi39 := 0; fi39 < 19; fi39++ {
                v40 := GenS2{F0: "hello", F1: 42}
                _ = fi39
                _ = v40
            }
        }
        for fi41 := 0; fi41 < 17; fi41++ {
            var v42 int = -1
            switch -142 {
            case -2147483648:
                fmt.Println(v42, v42, v42)
            }
            var v43 int16 = -1
            v2 = v2
            _ = fi41
            _ = v42
            _ = v43
        }
        var v44 uint16 = 552
        _ = fi36
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
