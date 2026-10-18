package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint, p1 uint32) int8 {
    if false {
        p1 = (^(p1))
        v5 := GenS1{F0: (717 <= 0)}
        _ = v5
    } else {
        var v6 int16 = -243
        v7 := GenS1{F0: (0 < 1)}
        _ = v6
        _ = v7
    }
    p1 = 1
    var v8 string = ("hello" + ("" + "hello"))
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 bool = false
        p0 += (p0 - p0)
        for fi11 := 0; fi11 < 6; fi11++ {
            fmt.Println(fi9, fi11, v8)
            v12 := []int{}
            if (!(32767 == 0)) {
                var v13 bool = (1 != 0)
                _ = v13
            } else {
                var v14 string = (v8 + "foo bar")
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v8
    return 127
}
func genFn1(p0 int8, p1 int16) int {
    fmt.Println(p0)
    v15 := GenS1{F0: false}
    _ = p0
    _ = p1
    _ = v15
    return 287
}

func main() {
    var v1 bool = true
    var v2 bool = (1 < 65535)
    var v3 int32 = 1
    var v4 bool = false
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 uint = 1
        fmt.Println(v1, fi16, fi16)
        _ = fi16
        _ = v17
    }
    var v18 bool = v4
    var v19 uint64 = 78
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := GenS2{F0: v19, F1: 4294967295}
        for fi22 := 0; fi22 < 9; fi22++ {
            switch fi22 {
            case -569:
                if (463 <= 119) {
                    var v23 uint32 = 2147483647
                    v24 := []int{}
                    var v25 string = ""
                    _ = v23
                    _ = v24
                    _ = v25
                }
            }
            fmt.Println(v1, v1)
            _ = fi22
        }
        _ = fi20
        _ = v21
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        v18 = false
        switch fi26 {
        case 9223372036854775807:
            var v27 uint = 419
            _ = v27
        case 0:
            v3 = (-(v3))
        case -878:
            v18 = true
        }
        v2 = v18
        _ = fi26
    }
    if v4 {
        for fi28 := 0; fi28 < 4; fi28++ {
            fmt.Println(v18, v1, v3)
            v29 := GenS3{F0: 128, F1: 340}
            _ = fi28
            _ = v29
        }
        if v2 {
            var v30 uint16 = 65535
            if v18 {
                v4 = (((v30 >> ((2) & 63)) < v30) || true)
                var v31 int32 = (v3 | -140)
                _ = v31
            }
            v32 := []int{-297, 293, 2147483647}
            _ = v30
            _ = v32
        }
    } else {
        v1 = (0 == 2)
    }
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 float32 = -2.25
        switch ("foo bar" + "foo bar") {
        case "a":
            v35 := GenS2{F0: v19, F1: v19}
            _ = v35
        case "foo bar":
            var v36 int16 = 42
            _ = v36
        case "":
            var v37 int8 = 127
            _ = v37
        default:
            var v38 int = -863
            _ = v38
        }
        switch "\n\t" {
        case "hello":
            if true {
                var v39 float32 = (((-2.25 / v34) + v34) - (0.0 / (-2.25 * v34)))
                if (1 != 4294967295) {
                    v40 := []int{411}
                    _ = v40
                }
                _ = v39
            }
        case "":
            var v41 int = fi33
            _ = v41
        }
        _ = fi33
        _ = v34
    }
    switch 4294967295 {
    case -985:
        if (!v1) {
            for fi42 := 0; fi42 < 15; fi42++ {
                for fi43 := 0; fi43 < 3; fi43++ {
                    var v44 float32 = (0.0 - ((1e10 - -2.25) + (0.1 - (0.0 + 0.0))))
                    _ = fi43
                    _ = v44
                }
                _ = fi42
            }
            fmt.Println(v19)
        }
    case 128:
        v1 = (("foo bar" > "hello") || v18)
    case -2147483648:
        v45 := []int{177, genFn1(-1, 688)}
        _ = v45
    }
    var v46 float32 = -2.25
    var v47 uint64 = v19
    var v48 uint32 = 753
    v49 := GenS2{F0: v47, F1: 744}
    v2 = false
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
}
