package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint }
func genFn0(p0 uint8) uint {
    v5 := GenS1{F0: 65535, F1: 423}
    if true {
        switch -1 {
        case 1:
            p0 += p0
        case -1:
            var v6 int8 = 1
            _ = v6
        case 127:
            switch ("\n\t" + "") {
            case "foo bar":
                v7 := GenS2{F0: 1}
                _ = v7
            case "a":
                v8 := []int{619, -9223372036854775808, 32768}
                _ = v8
            case "\n\t":
                v9 := []int{42}
                _ = v9
            default:
                fmt.Println(p0)
            }
        default:
            var v10 int8 = -69
            _ = v10
        }
    }
    _ = p0
    _ = v5
    return 18446744073709551615
}
func genFn1(p0 uint8, p1 int16) int8 {
    switch -210 {
    case -1:
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 int = fi11
            if true {
                var v13 uint16 = 65535
                for fi14 := 0; fi14 < 17; fi14++ {
                    v13 -= v13
                    v15 := []int{fi14, v12, -265, 395}
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
            _ = fi11
            _ = v12
        }
    case 667:
        if (p0 < 127) {
            var v16 uint8 = p0
            fmt.Println(p1, p0)
            _ = v16
        } else {
            if false {
                v17 := GenS3{F0: 52}
                for fi18 := 0; fi18 < 3; fi18++ {
                    v19 := GenS4{F0: 2147483647, F1: "hello"}
                    p1 = p1
                    var v20 uint64 = 1
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                _ = v17
            }
            for fi21 := 0; fi21 < 11; fi21++ {
                for fi22 := 0; fi22 < 11; fi22++ {
                    p1 = (-(p1))
                    _ = fi22
                }
                _ = fi21
            }
        }
    }
    v23 := GenS5{F0: 42, F1: (p0 * p0)}
    v24 := []int{}
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    return 100
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 267
    v3 := make(map[int]uint32)
    var v4 int16 = 1
    v2 = 42
    fmt.Println(v2)
    for fi25 := 0; fi25 < 10; fi25++ {
        fmt.Println(v2, fi25, fi25)
        _ = fi25
    }
    var v26 int16 = v4
    v26 = v26
    for fi27 := 0; fi27 < 1; fi27++ {
        var v28 int32 = -1
        var v29 float32 = (0.1 - (0.0 - 0.1))
        _ = fi27
        _ = v28
        _ = v29
    }
    v30 := v3[42]
    v4 += v26
    for fi31 := 0; fi31 < 5; fi31++ {
        fmt.Println(fi31, v2, v4)
        if ((v26 / ((-32768) | 1)) <= v26) {
            var v32 int8 = genFn1(42, -32768)
            fmt.Println(v2)
            switch (v2 * 32767) {
            case -170:
                v4 -= v26
            case 128:
                if ((-2.25 / (1e10 - ((-2.25 / 1.0) / 3.14159))) != (float64(-2.25) * (0.0 / 1e10))) {
                    v32 = genFn1(255, -32768)
                    for fi33 := 0; fi33 < 2; fi33++ {
                        v34 := GenS1{F0: 477, F1: v30}
                        v35 := GenS6{F0: 684, F1: 867}
                        var v36 bool = false
                        var v37 uint32 = v30
                        _ = fi33
                        _ = v34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                } else {
                    v38 := GenS6{F0: 255, F1: 65535}
                    var v39 int64 = 9223372036854775807
                    fmt.Println(v32, v2, v39)
                    _ = v38
                    _ = v39
                }
            }
            _ = v32
        } else {
            v40 := GenS7{F0: 0}
            _ = v40
        }
        var v41 uint32 = (^(v30))
        _ = fi31
        _ = v41
    }
    v30 = (v30 % ((2147483647) | 1))
    if true {
        var v42 uint64 = 42
        var v43 bool = false
        for fi44 := 0; fi44 < 13; fi44++ {
            v30 -= 4294967295
            _ = len(v1)
            v43 = (!v43)
            _ = fi44
        }
        _ = v42
        _ = v43
    }
    if false {
        for fi45 := 0; fi45 < 19; fi45++ {
            v46 := GenS8{F0: 18446744073709551615}
            v47 := []int{v2, fi45, -2147483648}
            _ = fi45
            _ = v46
            _ = v47
        }
        var v48 int8 = 1
        _ = v48
    } else {
        v2 = v2
    }
    v26 -= 1000
    switch (((("" + ("foo bar" + "")) + "") + "a") + (("" + "\n\t") + ("\n\t" + ((("" + "") + (("" + "foo bar") + "a")) + ("foo bar" + ("" + "")))))) {
    case "foo bar":
        delete(v3, v2)
    }
    v49 := GenS8{F0: 665}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v49
}
