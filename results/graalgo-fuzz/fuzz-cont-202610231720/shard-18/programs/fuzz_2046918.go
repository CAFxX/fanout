package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint; F1 int; F2 string }
type GenS5 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS6 struct { F0 bool; F1 uint32; F2 int }
func genFn0() int8 {
    v5 := GenS1{F0: 128}
    v6 := GenS1{F0: 253}
    var v7 bool = true
    fmt.Println(v7)
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 float32 = 1.5
    var v2 int8 = 42
    var v3 uint16 = 0
    var v4 int8 = (^(v2))
    if false {
        v8 := GenS2{F0: 2147483647}
        _ = v8
    } else {
        for fi9 := 0; fi9 < 1; fi9++ {
            v3 += (v3 + 42)
            v3 -= (^(v3))
            switch "hello" {
            case "foo bar":
                var v10 int = (fi9 | 32767)
                _ = v10
            case "a":
                for fi11 := 0; fi11 < 11; fi11++ {
                    fmt.Println(v3, v3, v1)
                    _ = fi11
                }
            case "hello":
                var v12 uint8 = 128
                _ = v12
            }
            _ = fi9
        }
        v2 = v2
    }
    v13 := []int{127}
    var v14 int32 = 42
    switch 42 {
    case 2147483647:
        var v15 int16 = 32767
        _ = v15
    case -2147483648:
        var v16 uint32 = 4294967295
        _ = v16
    case 9223372036854775807:
        switch -991 {
        case 0:
            v1 = ((v1 * v1) - (float32(((0.0 * 0.1) - 1.5)) / -2.25))
        case -521:
            var v17 uint = 4294967295
            _ = v17
        }
    }
    if false {
        v2 -= 0
        v18 := GenS1{F0: 222}
        for fi19 := 0; fi19 < 10; fi19++ {
            for fi20 := 0; fi20 < 17; fi20++ {
                if true {
                    var v21 uint8 = 128
                    v22 := GenS3{F0: 42, F1: (v14 ^ (v14 | (v14 | v14))), F2: v3}
                    var v23 int = -564
                    var v24 float32 = v1
                    v25 := GenS1{F0: (v21 % (((v21 ^ (v21 ^ v21))) | 1))}
                    _ = v21
                    _ = v22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                switch "hello" {
                case "":
                    v3 -= 42
                case "foo bar":
                    v3 -= 42
                    v4 = 100
                    var v26 int8 = 100
                    _ = v26
                case "hello":
                    var v27 uint64 = 1
                    v27 -= v27
                    fmt.Println(v3, v3, fi20)
                    v28 := GenS4{F0: 883, F1: fi19, F2: ("" + "")}
                    if ((false && true) && false) {
                        v29 := GenS2{F0: 273}
                        var v30 float64 = (3.14159 * float64(v1))
                        _ = v29
                        _ = v30
                    }
                    _ = v27
                    _ = v28
                }
                _ = fi20
            }
            _ = fi19
        }
        _ = v18
    }
    if false {
        if (true && (((v3 != (^(v3))) && true) || false)) {
            fmt.Println(v4, v14)
        }
        var v31 bool = (!((!(!((-(v2)) <= (v4 & v4)))) || (!(("\n\t" + "") >= "foo bar"))))
        v1 -= (v1 - 0.0)
        _ = v31
    } else {
        var v32 int64 = 1
        for fi33 := 0; fi33 < 17; fi33++ {
            var v34 float32 = v1
            switch (fi33 / ((int(v3)) | 1)) {
            case 42:
                v2 = (v4 | (v2 | int8(v32)))
            default:
                var v35 uint8 = 56
                _ = v35
            }
            for fi36 := 0; fi36 < 14; fi36++ {
                v14 -= v14
                fmt.Println(fi36)
                var v37 int = fi33
                _ = fi36
                _ = v37
            }
            _ = fi33
            _ = v34
        }
        _ = v32
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        for fi39 := 0; fi39 < 15; fi39++ {
            if true {
                v2 += v4
                fmt.Println(fi39)
                v40 := GenS5{F0: v14, F1: 0, F2: 42}
                _ = v40
            } else {
                v14 -= v14
            }
            _ = fi39
        }
        _ = fi38
    }
    if (!false) {
        var v41 bool = true
        fmt.Println(v14, v3, v4)
        _ = v41
    }
    switch -2147483648 {
    case -898:
        v4 = v2
    }
    if (!(127 > 255)) {
        var v42 int = -608
        _ = v42
    }
    if len(v13) > 0 {
        v13[0] = -720
    }
    for fi43 := 0; fi43 < 2; fi43++ {
        if false {
            v2 = 42
            for fi44 := 0; fi44 < 9; fi44++ {
                v1 = v1
                _ = fi44
            }
            v45 := GenS6{F0: (614 < 150), F1: 2147483648, F2: (-(fi43))}
            _ = v45
        }
        var v46 float32 = v1
        _ = fi43
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
