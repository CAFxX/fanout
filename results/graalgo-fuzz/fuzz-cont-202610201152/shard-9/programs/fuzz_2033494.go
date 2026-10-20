package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint; F1 uint32; F2 bool }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int16) uint16 {
    for fi4 := 0; fi4 < 16; fi4++ {
        p0 -= p0
        _ = fi4
    }
    var v5 int8 = 1
    _ = p0
    _ = v5
    return 32767
}

func main() {
    var v1 int8 = -49
    var v2 uint32 = 65535
    var v3 string = (("foo bar" + "") + "\n\t")
    v6 := GenS1{F0: v1, F1: 279}
    v7 := GenS1{F0: v1, F1: 32767}
    fmt.Println(v2, v2)
    fmt.Println(v3)
    var v8 uint8 = 185
    var v9 bool = true
    if (1 != 1) {
        var v10 uint64 = 4294967295
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        v12 := GenS1{F0: 42, F1: 1000}
        var v13 int8 = 127
        var v14 uint16 = 2
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 int16 = -32768
    if (3.14159 <= 0.0) {
        if false {
            switch ("a" + v3) {
            case "foo bar":
                v3 = "\n\t"
            case "a":
                v16 := GenS1{F0: 127, F1: (v15 | v15)}
                _ = v16
            case "\n\t":
                v17 := []int{}
                _ = v17
            }
        }
        v8 -= v8
    } else {
        v18 := GenS2{F0: 18446744073709551615, F1: v2, F2: false}
        if ((^(v1)) != v1) {
            var v19 uint32 = (v2 << ((2) & 63))
            var v20 int = 9223372036854775807
            v9 = v9
            _ = v19
            _ = v20
        }
        switch 42 {
        case 128:
            switch "\n\t" {
            case "foo bar":
                v2 = (v2 & (v2 & v2))
            }
        default:
            v1 += (v1 / ((v1) | 1))
        }
        _ = v18
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        switch (^(fi21)) {
        case -9223372036854775808:
            fmt.Println(v9)
        }
        if false {
            var v22 uint64 = 2
            var v23 uint = 0
            if v9 {
                if v9 {
                    switch 127 {
                    case -66:
                        v24 := GenS1{F0: (v1 << (((v23 ^ (v23 % ((v23) | 1)))) & 63)), F1: 32767}
                        _ = v24
                    case -796:
                        v8 = 0
                    default:
                        if ((v8 ^ (v8 << ((uint(v2)) & 63))) != v8) {
                            fmt.Println(v23)
                            for fi25 := 0; fi25 < 8; fi25++ {
                                fmt.Println(v15)
                                v15 -= -1
                                _ = fi25
                            }
                            v8 = 175
                        }
                    }
                    var v26 uint16 = 255
                    v27 := GenS2{F0: v23, F1: v2, F2: (v9 && v9)}
                    _ = v26
                    _ = v27
                }
                switch -921 {
                case 9223372036854775807:
                    v28 := []int{-1, 42, (fi21 ^ 9223372036854775807), int(v1)}
                    _ = v28
                default:
                    if (v22 < v22) {
                        var v29 uint16 = uint16(v23)
                        var v30 string = "a"
                        _ = v29
                        _ = v30
                    } else {
                        v31 := []int{fi21, fi21, int(v1)}
                        v9 = (v2 == (^(v2)))
                        if (128 == 4294967295) {
                            fmt.Println(v23, v2)
                            fmt.Println(fi21, v22, v15)
                        }
                        _ = v31
                    }
                }
            }
            _ = v22
            _ = v23
        } else {
            v32 := []int{(^(fi21)), (fi21 & fi21)}
            var v33 int32 = -1
            _ = v32
            _ = v33
        }
        v34 := GenS3{F0: 0}
        _ = fi21
        _ = v34
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        var v36 int16 = (v15 >> ((uint(v8)) & 63))
        for fi37 := 0; fi37 < 19; fi37++ {
            switch 32768 {
            case 1:
                v38 := GenS2{F0: 1, F1: 1, F2: v9}
                _ = v38
            case -1:
                v39 := []int{}
                _ = v39
            default:
                v9 = ((true || v9) && ((fi37 + fi35) >= fi37))
            }
            if ((v15 / ((v15) | 1)) > -32768) {
                var v40 float32 = float32(-2.25)
                fmt.Println(fi37)
                _ = v40
            }
            v3 = v3
            _ = fi37
        }
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
