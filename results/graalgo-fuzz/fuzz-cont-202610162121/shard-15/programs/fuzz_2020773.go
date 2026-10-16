package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint16; F1 int; F2 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 bool }
func genFn0(p0 uint64) int64 {
    v6 := GenS1{F0: 42, F1: 42}
    var v7 int64 = -859
    _ = p0
    _ = v6
    _ = v7
    return v7
}

func main() {
    var v1 int = -1
    v2 := make(map[uint32]string)
    var v3 uint64 = 42
    var v4 string = "foo bar"
    v5 := make(map[string]int)
    v4 = v4
    if false {
        v3 = (v3 % ((v3) | 1))
    } else {
        v3 = v3
        var v8 int16 = -32768
        _ = v8
    }
    if false {
        v9 := GenS1{F0: 128, F1: 863}
        v10 := GenS1{F0: 42, F1: -945}
        fmt.Println(v4, v4, v3)
        _ = v10
        _ = v9
    } else {
        var v11 uint16 = 401
        if true {
            var v12 int8 = -45
            v13 := GenS2{F0: 4294967295, F1: 42}
            _ = v12
            _ = v13
        }
        _ = v11
    }
    if true {
        switch (v1 + v1) {
        case 32767:
            if (false && true) {
                var v14 float64 = (1.5 / 1.5)
                _ = v14
            }
        case 1:
            switch -9223372036854775808 {
            case -1:
                v3 += (v3 & v3)
            case -411:
                for fi15 := 0; fi15 < 15; fi15++ {
                    fmt.Println(v1)
                    _ = fi15
                }
            default:
                var v16 float64 = (0.1 * 0.1)
                _ = v16
            }
        case 127:
            var v17 bool = (1e10 < (1e10 - 1.5))
            _ = v17
        default:
            v18 := GenS3{F0: 1, F1: (v1 >> ((900) & 63)), F2: 486}
            _ = v18
        }
    } else {
        var v19 int = 1
        v20 := GenS4{F0: 42}
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        v22 := []int{fi21, v1}
        v23 := GenS4{F0: 349}
        if (743 == 0) {
            var v24 int = (fi21 * (v1 + v1))
            if ((v24 % ((v24) | 1)) < (v24 + v24)) {
                if false {
                    v3 += 18446744073709551615
                }
            }
            _ = v24
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        if (3.14159 == (0.0 / ((-2.25 / ((1.5 - 1e10) / -2.25)) + ((0.1 * float64(-2.25)) * 1.5)))) {
            v26 := []int{375, (^(fi25)), (v1 >> ((147) & 63))}
            var v27 int8 = 98
            if true {
                switch "foo bar" {
                case "":
                    for fi28 := 0; fi28 < 15; fi28++ {
                        v29 := GenS4{F0: 0}
                        _ = len(v26)
                        v4 = v4
                        _ = fi28
                        _ = v29
                    }
                case "hello":
                    v1 = fi25
                case "\n\t":
                    for fi30 := 0; fi30 < 20; fi30++ {
                        v31 := GenS2{F0: 0, F1: 0}
                        fmt.Println(fi30)
                        _ = fi30
                        _ = v31
                    }
                }
                v32 := GenS5{F0: 2}
                _ = v32
            }
            _ = v26
            _ = v27
        }
        var v33 uint8 = 220
        _ = fi25
        _ = v33
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        v4 = v4
        var v35 uint = 1
        var v36 uint = 4294967295
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 float32 = (0.1 - -2.25)
    v38 := []int{(v1 * (^(v1))), 42}
    for fi39 := 0; fi39 < 16; fi39++ {
        v4 = ("hello" + ((v4 + v4) + "hello"))
        for fi40 := 0; fi40 < 11; fi40++ {
            fmt.Println(v3, fi39)
            _ = fi40
        }
        v4 = ((v4 + v4) + v4)
        _ = fi39
    }
    switch -820 {
    case 781:
        var v41 uint16 = 255
        _ = v41
    case -1:
        var v42 int = -403
        _ = v42
    case -408:
        var v43 uint32 = uint32(v3)
        _ = v43
    default:
        for fi44 := 0; fi44 < 2; fi44++ {
            var v45 int16 = -32768
            _ = fi44
            _ = v45
        }
    }
    delete(v5, ("hello" + v4))
    for fi46 := 0; fi46 < 19; fi46++ {
        var v47 int = fi46
        if (v37 == float32(1.5)) {
            v48 := GenS6{F0: -454, F1: true}
            switch v47 {
            case 2147483647:
                v49 := GenS6{F0: 817, F1: true}
                _ = v49
            case 42:
                fmt.Println(fi46, fi46)
            }
            _ = v48
        } else {
            fmt.Println(v4)
            v3 = v3
        }
        v4 = (v4 + (((v4 + v4) + ((v4 + v4) + v4)) + v4))
        _ = fi46
        _ = v47
    }
    var v50 int16 = -702
    switch "foo bar" {
    case "a":
        var v51 uint64 = v3
        _ = v51
    case "foo bar":
        v37 -= v37
    case "":
        fmt.Println(v50, v3, v37)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v50
}
