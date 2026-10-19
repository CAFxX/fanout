package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint32; F1 string; F2 int8 }
type GenS4 struct { F0 int32; F1 uint64; F2 int32 }
func genFn0(p0 int16) int {
    var v4 bool = ((1 == 55) || (p0 >= (-(p0))))
    var v5 uint = 156
    _ = p0
    _ = v4
    _ = v5
    return -9223372036854775808
}

func main() {
    var v1 int = 2147483647
    var v2 string = "hello"
    var v3 string = ("" + ("" + (v2 + (v2 + v2))))
    switch (("a" + ("foo bar" + v3)) + (v3 + (v2 + "a"))) {
    case "":
        for fi6 := 0; fi6 < 5; fi6++ {
            v3 = "foo bar"
            if (-32768 >= 32767) {
                v7 := []int{genFn0(0)}
                var v8 uint = 0
                var v9 float32 = -2.25
                _ = v7
                _ = v8
                _ = v9
            }
            _ = fi6
        }
    case "\n\t":
        v10 := []int{2147483647}
        _ = v10
    case "a":
        v11 := []int{v1, (v1 >> ((1) & 63))}
        _ = v11
    default:
        if true {
            v1 = (v1 % ((v1) | 1))
            v1 = -197
        } else {
            var v12 int = (-(v1))
            v2 = v2
            switch v3 {
            case "foo bar":
                var v13 string = ((("" + ("hello" + "a")) + "foo bar") + v2)
                _ = v13
            default:
                var v14 uint8 = 127
                _ = v14
            }
            _ = v12
        }
    }
    if true {
        var v15 float32 = ((1e10 - float32(1e10)) * 1e10)
        if true {
            var v16 uint32 = 174
            var v17 uint64 = 55
            _ = v16
            _ = v17
        }
        _ = v15
    } else {
        v18 := GenS1{F0: 796}
        v1 -= (v1 >> ((65535) & 63))
        _ = v18
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        for fi20 := 0; fi20 < 18; fi20++ {
            v21 := GenS1{F0: 255}
            _ = fi20
            _ = v21
        }
        if (true || (42 <= 1)) {
            if true {
                v22 := []int{210, (fi19 / (((v1 & 1)) | 1)), (fi19 % ((v1) | 1))}
                if (127 != 0) {
                    if (!false) {
                        var v23 uint = 18446744073709551615
                        if (false || ((!true) || false)) {
                            var v24 uint32 = 2
                            _ = v24
                        }
                        _ = v23
                    } else {
                        var v25 int16 = -422
                        _ = v25
                    }
                } else {
                    var v26 int32 = 2147483647
                    _ = v26
                }
                var v27 uint32 = 4294967295
                _ = v22
                _ = v27
            }
            v28 := []int{fi19, (-(fi19)), 86}
            _ = v28
        }
        _ = fi19
    }
    for fi29 := 0; fi29 < 1; fi29++ {
        v30 := []int{(fi29 - (v1 % ((genFn0(0)) | 1))), v1}
        for fi31 := 0; fi31 < 8; fi31++ {
            v30 = append(v30, -992)
            fmt.Println(fi29, v2, fi29)
            fmt.Println(fi31)
            _ = fi31
        }
        _ = fi29
        _ = v30
    }
    if (127 > -99) {
        if true {
            for fi32 := 0; fi32 < 2; fi32++ {
                v2 = (v3 + ("\n\t" + v3))
                _ = fi32
            }
        }
        v2 = v3
        if false {
            v33 := []int{(v1 & genFn0(-1)), 1, 32768, v1}
            _ = v33
        } else {
            v34 := []int{}
            var v35 bool = (32767 != 936)
            _ = v34
            _ = v35
        }
    }
    v36 := GenS2{F0: 32767, F1: 42}
    var v37 int16 = -32768
    for fi38 := 0; fi38 < 14; fi38++ {
        fmt.Println(v37)
        for fi39 := 0; fi39 < 3; fi39++ {
            switch 4294967295 {
            case 4294967295:
                v37 -= (v37 >> ((1) & 63))
            case -845:
                v40 := GenS3{F0: 2147483648, F1: v3, F2: 106}
                _ = v40
            default:
                v41 := GenS4{F0: 482, F1: 4294967295, F2: 1}
                _ = v41
            }
            _ = fi39
        }
        var v42 uint32 = 4294967295
        _ = fi38
        _ = v42
    }
    switch v2 {
    case "a":
        switch ((v2 + v3) + v2) {
        case "":
            var v43 string = v3
            _ = v43
        case "\n\t":
            v3 = v2
        }
    }
    if true {
        for fi44 := 0; fi44 < 20; fi44++ {
            var v45 uint = 18446744073709551615
            v46 := GenS2{F0: v37, F1: 1}
            _ = fi44
            _ = v45
            _ = v46
        }
    } else {
        var v47 bool = false
        _ = v47
    }
    if false {
        var v48 string = ("" + (v3 + v3))
        v1 += genFn0((v37 + v37))
        v49 := GenS1{F0: 908}
        _ = v48
        _ = v49
    } else {
        v50 := []int{}
        v2 = "\n\t"
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
}
