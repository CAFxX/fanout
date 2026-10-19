package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int16; F1 int8; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint; F2 int16 }
func genFn0(p0 uint, p1 uint32, p2 int16) int8 {
    p0 = uint(p1)
    var v7 uint8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 int16, p1 uint8, p2 uint) bool {
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint = 314
        fmt.Println(p2)
        _ = fi8
        _ = v9
    }
    p1 = (p1 & p1)
    v10 := []int{}
    v11 := GenS1{F0: 0}
    var v12 uint = (p2 - (^(p2)))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return false
}

func main() {
    var v1 uint = 255
    v2 := make(map[int]bool)
    var v3 uint8 = 0
    var v4 int8 = 100
    var v5 int64 = int64(v1)
    v6 := make(map[int]int)
    switch ("foo bar" + "foo bar") {
    case "hello":
        v13 := GenS1{F0: v4}
        _ = v13
    case "":
        v14 := GenS2{F0: 32767, F1: "a"}
        _ = v14
    }
    v15 := GenS3{F0: 42, F1: (^(v4)), F2: 351}
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 bool = ((2147483647 <= 42) || true)
        var v18 string = "foo bar"
        _ = len(v2)
        _ = fi16
        _ = v17
        _ = v18
    }
    switch "hello" {
    case "hello":
        var v19 uint = 255
        _ = v19
    case "a":
        v5 -= v5
    case "":
        v5 -= v5
    default:
        for fi20 := 0; fi20 < 10; fi20++ {
            v21 := GenS4{F0: (-872 != 1), F1: (v1 + 0), F2: 856}
            _ = fi20
            _ = v21
        }
    }
    switch "a" {
    case "":
        for fi22 := 0; fi22 < 6; fi22++ {
            v23 := GenS3{F0: 42, F1: (v4 * v4), F2: 782}
            _ = fi22
            _ = v23
        }
    case "a":
        var v24 uint32 = 4294967295
        _ = v24
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        var v26 string = ""
        _ = fi25
        _ = v26
    }
    v27 := v6[-97]
    v28 := GenS3{F0: -1, F1: genFn0((v1 << (((v1 ^ 0)) & 63)), 0, -1), F2: 0}
    if (!(!false)) {
        v29 := []int{}
        v30 := GenS1{F0: (v4 ^ 100)}
        _ = v29
        _ = v30
    } else {
        v27 += v27
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        var v32 uint64 = uint64(v3)
        if false {
            v33 := GenS1{F0: v4}
            for fi34 := 0; fi34 < 20; fi34++ {
                var v35 float32 = 1e10
                var v36 uint64 = (^(v32))
                v37 := GenS1{F0: (v4 << (((v1 + 948)) & 63))}
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
            }
            _ = v33
        }
        _ = fi31
        _ = v32
    }
    if ((3.14159 - (((-2.25 * -2.25) / ((0.1 + (0.0 + 1.5)) - 1.5)) - 1e10)) >= (float32(((-2.25 - 0.1) * 0.1)) - 3.14159)) {
        v38 := []int{}
        v39 := GenS1{F0: v4}
        _ = v38
        _ = v39
    }
    for fi40 := 0; fi40 < 9; fi40++ {
        switch ("hello" + (("\n\t" + ("foo bar" + "foo bar")) + "\n\t")) {
        case "":
            v41 := GenS1{F0: (v4 - (-(v4)))}
            _ = v41
        case "\n\t":
            v42 := []int{fi40, fi40}
            _ = v42
        default:
            v27 = v27
        }
        switch ("\n\t" + "\n\t") {
        case "\n\t":
            switch "\n\t" {
            case "hello":
                v43 := v2[4294967295]
                _ = v43
            }
        case "":
            if false {
                v44 := []int{(fi40 % ((fi40) | 1)), (fi40 | (fi40 * v27)), -2147483648, 4294967295}
                v5 = (-(v5))
                _ = v44
            } else {
                var v45 int16 = 882
                v46 := []int{(fi40 << ((1) & 63)), (fi40 - fi40), fi40}
                _ = v45
                _ = v46
            }
        case "a":
            var v47 int32 = 592
            _ = v47
        default:
            v1 = (v1 | v1)
        }
        var v48 uint64 = 42
        _ = fi40
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
