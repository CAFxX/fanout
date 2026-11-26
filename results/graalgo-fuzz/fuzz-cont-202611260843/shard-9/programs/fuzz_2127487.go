package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 string; F1 int32; F2 uint8 }
type GenS3 struct { F0 int }
func genFn0(p0 bool, p1 uint64) uint8 {
    switch "hello" {
    case "foo bar":
        v6 := []int{-2147483648, 32767, 32767}
        _ = v6
    case "a":
        v7 := []int{-399}
        _ = v7
    case "":
        v8 := []int{643, 2147483647, -1, 9223372036854775807}
        _ = v8
    }
    v9 := []int{-708, 127, 32767, 127}
    p0 = ((-2.25 / 1.5) < 3.14159)
    var v10 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 42
}

func main() {
    v1 := []string{"\n\t", "a", "a", ("" + ("hello" + "a"))}
    var v2 int = 128
    var v3 string = ("foo bar" + "a")
    v4 := []string{v3}
    var v5 int = 4294967295
    var v11 float32 = ((1e10 + 3.14159) - 1.5)
    switch ("foo bar" + (v3 + "foo bar")) {
    case "foo bar":
        var v12 uint8 = 255
        _ = v12
    case "a":
        v3 = v3
    case "hello":
        v3 = "\n\t"
    }
    switch ("a" + "") {
    case "foo bar":
        v2 = 267
    case "hello":
        switch -583 {
        case 19:
            var v13 uint64 = 0
            _ = v13
        case -930:
            v14 := GenS1{F0: false, F1: 42}
            _ = v14
        case 42:
            var v15 uint64 = 899
            _ = v15
        }
    case "":
        v5 += (v5 & v2)
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if (4294967295 >= 42) {
        if (false || false) {
            if false {
                var v16 uint32 = 2147483648
                _ = v16
            }
        } else {
            v17 := GenS2{F0: "foo bar", F1: 1, F2: 128}
            var v18 int16 = -106
            v2 -= v2
            _ = v17
            _ = v18
        }
    } else {
        switch v2 {
        case 1:
            var v19 uint8 = genFn0((!("" <= ((("\n\t" + v3) + ((("foo bar" + (v3 + (v3 + v3))) + "foo bar") + "\n\t")) + "a"))), 4294967295)
            _ = v19
        case 235:
            var v20 int16 = -400
            _ = v20
        default:
            v21 := GenS1{F0: (2147483647 > 295), F1: 1}
            _ = v21
        }
    }
    v2 = v5
    for fi22 := 0; fi22 < 9; fi22++ {
        switch fi22 {
        case 2147483647:
            var v23 int8 = 78
            _ = v23
        case -2147483648:
            v24 := GenS3{F0: 596}
            _ = v24
        case 42:
            switch (v3 + (v3 + v3)) {
            case "a":
                v5 = fi22
            case "foo bar":
                var v25 int8 = -105
                _ = v25
            case "\n\t":
                var v26 float64 = 0.1
                _ = v26
            default:
                v3 = v3
            }
        default:
            v27 := GenS1{F0: false, F1: 34}
            _ = v27
        }
        _ = fi22
    }
    fmt.Println(v2, v5)
    for fi28 := 0; fi28 < 1; fi28++ {
        for fi29 := 0; fi29 < 11; fi29++ {
            v2 = fi29
            _ = fi29
        }
        _ = fi28
    }
    v1 = append(v1, (("" + "") + v3))
    v11 -= (1.5 * ((1.5 + v11) * (((v11 * (float32(-2.25) - v11)) + 0.0) / float32(1e10))))
    switch "foo bar" {
    case "foo bar":
        var v30 int = v2
        _ = v30
    }
    var v31 uint = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
