package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int16 }
type GenS2 struct { F0 int8; F1 uint8; F2 string }
type GenS3 struct { F0 string; F1 uint16; F2 uint16 }
type GenS4 struct { F0 uint; F1 int8 }

func main() {
    var v1 uint16 = 65535
    var v2 int = -1
    var v3 bool = (-625 != 0)
    for fi4 := 0; fi4 < 10; fi4++ {
        v5 := []int{(v2 - fi4), int(v1), (-(fi4)), (v2 | v2)}
        fmt.Println(fi4, v3, v1)
        v6 := GenS1{F0: ("a" + "hello"), F1: 2, F2: 32767}
        _ = fi4
        _ = v5
        _ = v6
    }
    if v3 {
        v3 = v3
    } else {
        v7 := GenS1{F0: (("hello" + "\n\t") + (("" + "\n\t") + "hello")), F1: 888, F2: 1000}
        switch ("a" + ("\n\t" + "")) {
        case "\n\t":
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 int16 = -32768
                _ = fi8
                _ = v9
            }
        case "hello":
            v1 = v1
        }
        v3 = (!(1 < v2))
        _ = v7
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            switch "foo bar" {
            case "hello":
                v2 += (fi11 - fi11)
            default:
                var v12 uint = 65535
                _ = v12
            }
            v13 := GenS2{F0: -128, F1: 255, F2: (("\n\t" + "\n\t") + "")}
            var v14 float32 = (float32(0.1) / 3.14159)
            fmt.Println(fi11)
            v15 := []int{(fi11 ^ fi10), (^(fi10)), (fi11 ^ (-(fi11)))}
            _ = fi11
            _ = v13
            _ = v14
            _ = v15
        }
        switch "\n\t" {
        case "a":
            v16 := GenS1{F0: ((("" + (("foo bar" + ("\n\t" + "")) + ("" + ""))) + "foo bar") + ""), F1: 0, F2: -270}
            _ = v16
        }
        v17 := []int{}
        _ = fi10
        _ = v17
    }
    if (!v3) {
        var v18 uint16 = 886
        _ = v18
    } else {
        v19 := GenS3{F0: "a", F1: v1, F2: (^(v1))}
        v2 = (-(v2))
        _ = v19
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        for fi21 := 0; fi21 < 6; fi21++ {
            fmt.Println(fi21, fi21)
            var v22 int32 = 42
            _ = fi21
            _ = v22
        }
        switch "a" {
        case "\n\t":
            var v23 string = "foo bar"
            _ = v23
        case "":
            v3 = v3
        default:
            for fi24 := 0; fi24 < 6; fi24++ {
                var v25 int = (v2 * fi24)
                v26 := []int{(fi24 - (fi20 * fi20)), (fi20 % (((v2 | fi20)) | 1)), (v25 << ((160) & 63))}
                v25 -= (v25 % ((-1) | 1))
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        switch fi20 {
        case -9223372036854775808:
            v27 := []int{fi20, v2, -2147483648, (fi20 % (((v2 / ((v2) | 1))) | 1))}
            _ = v27
        case 9223372036854775807:
            var v28 int = 635
            _ = v28
        default:
            var v29 int8 = 77
            _ = v29
        }
        _ = fi20
    }
    switch -9223372036854775808 {
    case 0:
        var v30 int16 = -32768
        _ = v30
    case 42:
        var v31 uint32 = 114
        _ = v31
    case 1:
        for fi32 := 0; fi32 < 7; fi32++ {
            v2 -= (v2 << ((18446744073709551615) & 63))
            _ = fi32
        }
    }
    v33 := GenS4{F0: 255, F1: 1}
    switch "foo bar" {
    case "hello":
        var v34 float32 = 1e10
        _ = v34
    }
    if (false || (0 != 42)) {
        switch 0 {
        case -2147483648:
            var v35 float64 = -2.25
            _ = v35
        case 1:
            v36 := GenS3{F0: "", F1: (^(v1)), F2: 65535}
            _ = v36
        default:
            var v37 int = (v2 & 131)
            _ = v37
        }
        var v38 int = (v2 * (v2 >> ((4294967295) & 63)))
        v39 := GenS2{F0: -3, F1: 136, F2: ("\n\t" + "\n\t")}
        _ = v38
        _ = v39
    } else {
        for fi40 := 0; fi40 < 12; fi40++ {
            var v41 uint32 = 806
            switch (v2 & (fi40 >> ((1) & 63))) {
            case 4294967295:
                var v42 float32 = 1e10
                _ = v42
            }
            _ = fi40
            _ = v41
        }
        var v43 float32 = -2.25
        _ = v43
    }
    var v44 int = (v2 & (-(v2)))
    var v45 int32 = -2147483648
    var v46 int = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v44
    _ = v45
    _ = v46
}
