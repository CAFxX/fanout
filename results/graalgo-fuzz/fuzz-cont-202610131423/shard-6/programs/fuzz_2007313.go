package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int32 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int16; F1 int; F2 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 int8 }
func genFn0(p0 uint, p1 uint8) int {
    p0 -= 4294967295
    p0 += p0
    p0 = p0
    v4 := GenS1{F0: -267, F1: 740, F2: -276}
    _ = p0
    _ = p1
    _ = v4
    return 9223372036854775807
}
func genFn1() uint16 {
    for fi5 := 0; fi5 < 10; fi5++ {
        fmt.Println(fi5)
        v6 := []int{(fi5 ^ (^(fi5)))}
        _ = len(v6)
        _ = fi5
        _ = v6
    }
    var v7 int = -923
    switch (^(v7)) {
    case 998:
        var v8 int16 = -1
        _ = v8
    case -497:
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 int64 = 169
            if false {
                v11 := []int{}
                for fi12 := 0; fi12 < 2; fi12++ {
                    if len(v11) > 0 {
                        v11[0] = v7
                    }
                    _ = fi12
                }
                switch "foo bar" {
                case "":
                    v10 = (v10 | (^(v10)))
                }
                _ = v11
            } else {
                switch (^(v7)) {
                case 691:
                    v13 := []int{genFn0(642, 42)}
                    _ = v13
                case 128:
                    v7 = 471
                case -1:
                    var v14 uint = 2
                    _ = v14
                default:
                    v15 := GenS2{F0: 127, F1: true}
                    _ = v15
                }
            }
            fmt.Println(fi9)
            _ = fi9
            _ = v10
        }
    }
    if (!true) {
        fmt.Println(v7)
    } else {
        v16 := GenS3{F0: 1, F1: (v7 << ((825) & 63)), F2: 1}
        var v17 int32 = -100
        var v18 int16 = -919
        _ = v16
        _ = v17
        _ = v18
    }
    switch v7 {
    case 42:
        var v19 bool = (1 < 4294967295)
        _ = v19
    default:
        if true {
            for fi20 := 0; fi20 < 15; fi20++ {
                v7 = -1
                switch ("foo bar" + "") {
                case "\n\t":
                    for fi21 := 0; fi21 < 12; fi21++ {
                        fmt.Println(v7, fi20)
                        _ = fi21
                    }
                }
                _ = fi20
            }
        } else {
            v22 := GenS2{F0: -45, F1: (222 > 32767)}
            var v23 uint = 0
            var v24 int16 = 598
            _ = v22
            _ = v23
            _ = v24
        }
    }
    _ = v7
    return 342
}

func main() {
    var v1 int32 = 2147483647
    var v2 int64 = -98
    var v3 uint = 65535
    v1 -= v1
    var v25 float32 = (1e10 + 3.14159)
    var v26 uint = (^(v3))
    v27 := GenS4{F0: v2}
    v26 = 65535
    if false {
        var v28 float32 = v25
        v29 := []int{-337, genFn0(v3, uint8(v3)), 591, 127}
        var v30 int32 = (v1 | v1)
        _ = v28
        _ = v29
        _ = v30
    } else {
        v31 := []int{-605, -1, 2147483647}
        for fi32 := 0; fi32 < 14; fi32++ {
            if len(v31) > 0 {
                v31[0] = fi32
            }
            var v33 uint16 = 2
            v34 := GenS3{F0: -1, F1: int(v26), F2: v1}
            _ = fi32
            _ = v33
            _ = v34
        }
        if (v26 >= (v26 << (((v3 ^ 270)) & 63))) {
            v3 += (v26 - v3)
            v1 += -1
        }
        _ = v31
    }
    v26 = v26
    v25 = (1e10 + v25)
    for fi35 := 0; fi35 < 15; fi35++ {
        v36 := []int{(fi35 - fi35), fi35}
        _ = fi35
        _ = v36
    }
    v37 := GenS5{F0: (v3 - v3), F1: 121}
    v25 = 3.14159
    for fi38 := 0; fi38 < 10; fi38++ {
        var v39 int8 = 77
        v40 := GenS4{F0: -9223372036854775808}
        for fi41 := 0; fi41 < 10; fi41++ {
            if false {
                var v42 int16 = 0
                _ = v42
            }
            v43 := []int{genFn0(v26, 255)}
            if len(v43) >= 2 {
                v43 = v43[1:len(v43)]
            }
            _ = fi41
            _ = v43
        }
        _ = fi38
        _ = v39
        _ = v40
    }
    v44 := GenS4{F0: (^(v2))}
    var v45 int16 = -684
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v37
    _ = v44
    _ = v45
}
