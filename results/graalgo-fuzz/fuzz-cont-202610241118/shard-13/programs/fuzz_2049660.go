package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 string; F2 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32 }
func genFn0(p0 int16) int16 {
    p0 += 406
    if true {
        v7 := []int{-2147483648, 4294967295}
        _ = v7
    }
    var v8 int16 = (p0 / ((p0) | 1))
    var v9 int = 127
    _ = p0
    _ = v8
    _ = v9
    return v8
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]uint32)
    var v3 int8 = -2
    v4 := make(map[string]bool)
    v5 := make(map[uint32]bool)
    var v6 int16 = -1
    var v10 uint16 = 78
    for fi11 := 0; fi11 < 3; fi11++ {
        v3 = v3
        _ = fi11
    }
    var v12 uint64 = 2
    switch ("" + "a") {
    case "hello":
        v13 := GenS1{F0: 1, F1: ""}
        _ = v13
    case "":
        if false {
            if (65535 != 4294967295) {
                var v14 int64 = -307
                var v15 uint8 = 2
                _ = v14
                _ = v15
            }
            var v16 uint64 = v12
            _ = v16
        } else {
            switch "" {
            case "hello":
                var v17 uint = 1
                _ = v17
            case "\n\t":
                v18 := GenS2{F0: (v12 ^ 42)}
                _ = v18
            }
            var v19 uint32 = 600
            _ = v19
        }
    default:
        if (-9223372036854775808 > 0) {
            switch 2147483647 {
            case 128:
                v20 := GenS3{F0: "foo bar", F1: ((("\n\t" + "a") + "foo bar") + "foo bar"), F2: 0}
                _ = v20
            case -1:
                var v21 int8 = v3
                _ = v21
            case -423:
                var v22 float32 = (1e10 * -2.25)
                _ = v22
            }
            v10 = 255
            var v23 float32 = 1e10
            _ = v23
        } else {
            v6 += genFn0(v6)
            if (9223372036854775807 <= 4294967295) {
                for fi24 := 0; fi24 < 8; fi24++ {
                    v1 = append(v1, ("\n\t" + "a"))
                    _ = fi24
                }
                v10 = 255
                var v25 int32 = -2147483648
                _ = v25
            }
        }
    }
    switch "a" {
    case "":
        v26 := GenS3{F0: "hello", F1: "a", F2: (v12 / ((0) | 1))}
        _ = v26
    case "\n\t":
        for fi27 := 0; fi27 < 10; fi27++ {
            var v28 float32 = -2.25
            v3 -= (v3 - -15)
            fmt.Println(v3)
            _ = fi27
            _ = v28
        }
    default:
        for fi29 := 0; fi29 < 18; fi29++ {
            if false {
                fmt.Println(v10, v6)
            }
            fmt.Println(v12, v12, v6)
            switch (fi29 ^ fi29) {
            case -1:
                v30 := v4["a"]
                _ = v30
            default:
                _ = len(v1)
            }
            _ = fi29
        }
    }
    switch 42 {
    case -1:
        for fi31 := 0; fi31 < 7; fi31++ {
            v32 := GenS2{F0: v12}
            _ = fi31
            _ = v32
        }
    case -714:
        v12 -= v12
    default:
        if false {
            for fi33 := 0; fi33 < 18; fi33++ {
                v34 := GenS4{F0: "\n\t"}
                var v35 uint8 = 255
                _ = fi33
                _ = v34
                _ = v35
            }
        } else {
            v36 := GenS4{F0: "\n\t"}
            switch -126 {
            case 2147483647:
                var v37 uint = 65535
                _ = v37
            case 1:
                fmt.Println(v10, v3, v10)
            case 9223372036854775807:
                var v38 uint8 = 128
                _ = v38
            default:
                if false {
                    var v39 int8 = -72
                    var v40 uint64 = (v12 | v12)
                    _ = v39
                    _ = v40
                } else {
                    var v41 uint8 = 96
                    var v42 uint = 704
                    var v43 bool = true
                    _ = v41
                    _ = v42
                    _ = v43
                }
            }
            var v44 uint8 = 128
            _ = v36
            _ = v44
        }
    }
    _ = len(v1)
    fmt.Println(v12, v6, v10)
    v45 := GenS2{F0: v12}
    fmt.Println(v3, v6)
    for fi46 := 0; fi46 < 1; fi46++ {
        v6 = v6
        v47 := GenS5{F0: 1}
        v3 += (v3 + v3)
        _ = fi46
        _ = v47
    }
    v6 -= -477
    var v48 bool = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v48
    _ = v5
    _ = v6
}
