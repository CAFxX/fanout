package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint16, p1 uint32) int64 {
    for fi4 := 0; fi4 < 7; fi4++ {
        v5 := GenS1{F0: p1, F1: 0, F2: -128}
        v6 := GenS1{F0: p1, F1: uint32(p0), F2: -113}
        switch (fi4 / ((fi4) | 1)) {
        case 32767:
            switch "a" {
            case "a":
                fmt.Println(p0, fi4, fi4)
            case "\n\t":
                fmt.Println(p1, p0, p1)
            default:
                p0 = p0
            }
        case 0:
            if false {
                var v7 int16 = 42
                fmt.Println(p1)
                _ = v7
            }
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    switch 9223372036854775807 {
    case 2147483647:
        v8 := []int{-1, -2147483648, 32768}
        _ = v8
    case -2147483648:
        v9 := []int{-578}
        _ = v9
    case 32768:
        switch 42 {
        case 9223372036854775807:
            var v10 int32 = -1
            _ = v10
        case 42:
            v11 := []int{4294967295, 42}
            _ = v11
        }
    }
    if true {
        for fi12 := 0; fi12 < 10; fi12++ {
            fmt.Println(p0)
            switch fi12 {
            case -125:
                var v13 bool = (37 == 2)
                _ = v13
            default:
                for fi14 := 0; fi14 < 7; fi14++ {
                    var v15 string = ("foo bar" + ("foo bar" + "foo bar"))
                    var v16 int16 = 0
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            _ = fi12
        }
    } else {
        for fi17 := 0; fi17 < 7; fi17++ {
            if true {
                p0 += (p0 | 475)
                var v18 uint16 = p0
                _ = v18
            }
            _ = fi17
        }
    }
    v19 := GenS1{F0: (p1 ^ p1), F1: (p1 ^ (p1 * 42)), F2: 28}
    _ = p0
    _ = p1
    _ = v19
    return 1
}

func main() {
    v1 := []uint32{4294967295, 33}
    v2 := []string{("\n\t" + ("" + "\n\t")), "hello", ("hello" + "foo bar"), "hello"}
    var v3 string = ""
    var v20 int64 = -520
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        for fi22 := 0; fi22 < 12; fi22++ {
            if true {
                fmt.Println(v20, fi21)
                v3 = v3
            } else {
                var v23 int16 = -1
                var v24 float64 = ((1.5 * ((0.0 - -2.25) / 3.14159)) / 1.0)
                var v25 uint = 4294967295
                var v26 uint16 = 2
                _ = v23
                _ = v24
                _ = v25
                _ = v26
            }
            var v27 uint64 = 1
            _ = fi22
            _ = v27
        }
        _ = fi21
    }
    var v28 float32 = 0.0
    if (!true) {
        var v29 uint16 = 611
        switch (v3 + (("" + "hello") + v3)) {
        case "\n\t":
            fmt.Println(v3)
        case "hello":
            if (32768 != -660) {
                v30 := GenS2{F0: 496}
                v28 = (v28 * v28)
                v28 += float32(1e10)
                _ = v30
            } else {
                var v31 int = 0
                fmt.Println(v20, v29)
                v31 += (v31 << ((4294967295) & 63))
                _ = v31
            }
        case "":
            var v32 int = 1
            _ = v32
        default:
            v33 := GenS1{F0: 0, F1: 2147483648, F2: 100}
            _ = v33
        }
        if (18446744073709551615 < 2) {
            v34 := GenS2{F0: 833}
            _ = v34
        }
        _ = v29
    } else {
        v20 = -9223372036854775808
    }
    var v35 int32 = 1
    for fi36 := 0; fi36 < 16; fi36++ {
        if false {
            var v37 int16 = 1
            v20 = 1
            _ = v37
        }
        _ = fi36
    }
    var v38 string = "a"
    for fi39 := 0; fi39 < 15; fi39++ {
        v3 = (v38 + ("\n\t" + "foo bar"))
        _ = fi39
    }
    for fi40 := 0; fi40 < 17; fi40++ {
        v3 = (v38 + ((v38 + "foo bar") + v38))
        v28 = (v28 * v28)
        _ = fi40
    }
    var v41 uint32 = 89
    var v42 int16 = 1000
    var v43 uint8 = 240
    for fi44 := 0; fi44 < 9; fi44++ {
        var v45 uint64 = uint64(v41)
        _ = len(v1)
        var v46 uint8 = (v43 - (^(v43)))
        _ = fi44
        _ = v45
        _ = v46
    }
    for fi47 := 0; fi47 < 15; fi47++ {
        fmt.Println(fi47, fi47)
        fmt.Println(v41, v43)
        var v48 float32 = v28
        _ = fi47
        _ = v48
    }
    v38 = "hello"
    fmt.Println(v41, v42, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v35
    _ = v38
    _ = v41
    _ = v42
    _ = v43
}
