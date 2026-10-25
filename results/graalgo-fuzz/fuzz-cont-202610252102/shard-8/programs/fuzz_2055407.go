package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int64 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint }
func genFn0(p0 int16) int8 {
    if (127 >= 0) {
        p0 = p0
        for fi6 := 0; fi6 < 14; fi6++ {
            for fi7 := 0; fi7 < 7; fi7++ {
                v8 := GenS1{F0: 18446744073709551615, F1: 240, F2: 621}
                p0 = (p0 - (p0 % (((p0 ^ p0)) | 1)))
                fmt.Println(p0, fi6)
                if (!(p0 > (p0 / ((p0) | 1)))) {
                    v9 := GenS2{F0: ("a" + ""), F1: (true || false)}
                    _ = v9
                }
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
    }
    var v10 string = "\n\t"
    _ = p0
    _ = v10
    return 1
}
func genFn1() uint8 {
    switch "a" {
    case "a":
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := []int{}
            for fi13 := 0; fi13 < 15; fi13++ {
                fmt.Println(fi11)
                var v14 float32 = float32(float64(3.14159))
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    case "":
        if ("hello" <= "a") {
            var v15 int16 = -32768
            _ = v15
        } else {
            var v16 int = -333
            var v17 float64 = 3.14159
            _ = v16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        for fi19 := 0; fi19 < 15; fi19++ {
            var v20 uint32 = 2147483648
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 uint = 65535
        _ = fi21
        _ = v22
    }
    var v23 uint8 = 167
    v24 := []int{-462, 127, 9223372036854775807, 1}
    _ = v23
    _ = v24
    return (v23 / ((v23) | 1))
}

func main() {
    var v1 uint16 = 147
    var v2 uint16 = 32767
    var v3 int32 = 595
    var v4 int32 = v3
    var v5 uint32 = 701
    if true {
        var v25 bool = false
        fmt.Println(v2)
        v5 = (v5 << ((1) & 63))
        _ = v25
    }
    v5 = (v5 << ((1) & 63))
    fmt.Println(v4, v4)
    for fi26 := 0; fi26 < 14; fi26++ {
        switch fi26 {
        case 1:
            v1 = 42
        default:
            v27 := GenS3{F0: 2, F1: -946}
            _ = v27
        }
        switch 128 {
        case 127:
            v2 = (^(v1))
        default:
            fmt.Println(v5, v1, fi26)
        }
        fmt.Println(v4, v1, v3)
        _ = fi26
    }
    v28 := GenS2{F0: ("" + "a"), F1: false}
    fmt.Println(v1, v2)
    if ((711 != 1) && (1e10 == 0.1)) {
        fmt.Println(v1, v3)
    } else {
        v3 = -1
        var v29 uint16 = v1
        _ = v29
    }
    if true {
        var v30 int32 = 2147483647
        _ = v30
    }
    v31 := GenS3{F0: 42, F1: -9223372036854775808}
    if false {
        if true {
            for fi32 := 0; fi32 < 13; fi32++ {
                fmt.Println(v4, v2, v2)
                v33 := []int{(fi32 >> ((665) & 63)), fi32, 32768}
                for fi34 := 0; fi34 < 3; fi34++ {
                    if len(v33) >= 2 {
                        v33 = v33[1:len(v33)]
                    }
                    _ = fi34
                }
                _ = fi32
                _ = v33
            }
            v4 -= v4
        } else {
            v4 = v4
            if (true || true) {
                for fi35 := 0; fi35 < 13; fi35++ {
                    if false {
                        v2 -= v2
                        for fi36 := 0; fi36 < 8; fi36++ {
                            fmt.Println(v2, v2, v4)
                            v37 := GenS4{F0: v1, F1: 2, F2: 18446744073709551615}
                            _ = fi36
                            _ = v37
                        }
                        fmt.Println(fi35)
                    }
                    fmt.Println(v1, v4)
                    fmt.Println(v4)
                    _ = fi35
                }
            } else {
                v38 := []int{989, 2147483647}
                v39 := GenS5{F0: -329}
                var v40 int32 = (-(v3))
                _ = v38
                _ = v39
                _ = v40
            }
            var v41 uint8 = 0
            _ = v41
        }
    } else {
        v1 += v2
        var v42 uint32 = v5
        if (0 == -109) {
            for fi43 := 0; fi43 < 1; fi43++ {
                v2 = 2
                _ = fi43
            }
            var v44 int8 = -11
            _ = v44
        }
        _ = v42
    }
    v4 -= v3
    v45 := []int{-9223372036854775808}
    switch ("\n\t" + ("\n\t" + ("\n\t" + ""))) {
    case "":
        var v46 string = "\n\t"
        _ = v46
    }
    v47 := GenS5{F0: 9223372036854775807}
    v48 := GenS6{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v45
    _ = v47
    _ = v48
    _ = v5
}
