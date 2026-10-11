package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint16; F2 int }
type GenS5 struct { F0 int16 }
func genFn0() uint {
    v6 := GenS1{F0: -669, F1: true}
    var v7 int16 = -32768
    switch "\n\t" {
    case "foo bar":
        v8 := []int{2147483647, 384, 272, 128}
        _ = v8
    case "\n\t":
        v9 := GenS2{F0: 726, F1: 65535}
        _ = v9
    default:
        var v10 int64 = 1
        _ = v10
    }
    _ = v6
    _ = v7
    return 905
}
func genFn1() bool {
    var v11 string = "hello"
    var v12 int8 = 0
    if false {
        switch 826 {
        case 128:
            v13 := []int{321}
            _ = v13
        case 742:
            v11 = v11
        case 32767:
            var v14 int8 = (v12 * (v12 << ((879) & 63)))
            _ = v14
        default:
            fmt.Println(v11, v12)
        }
        v15 := []int{32767}
        if (32767 > 1) {
            var v16 uint64 = 186
            if true {
                var v17 float64 = (0.0 / 3.14159)
                var v18 uint8 = 255
                var v19 float32 = 0.0
                _ = v17
                _ = v18
                _ = v19
            } else {
                v16 += v16
                if len(v15) > 0 {
                    v15[0] = 128
                }
            }
            _ = v16
        } else {
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
        }
        _ = v15
    }
    _ = v11
    _ = v12
    return true
}

func main() {
    var v1 uint = 42
    v2 := make(map[int]string)
    var v3 string = "hello"
    var v4 int16 = -32768
    var v5 string = "hello"
    fmt.Println(v3, v1)
    var v20 uint8 = 0
    var v21 int8 = 42
    v21 = -128
    switch ("hello" + "") {
    case "foo bar":
        for fi22 := 0; fi22 < 9; fi22++ {
            var v23 float32 = 0.1
            v24 := []int{(fi22 - fi22), (^(fi22)), fi22}
            _ = fi22
            _ = v23
            _ = v24
        }
    case "a":
        for fi25 := 0; fi25 < 14; fi25++ {
            v1 -= v1
            _ = fi25
        }
    }
    switch 952 {
    case 127:
        v26 := GenS2{F0: 1, F1: (v1 * 0)}
        _ = v26
    }
    var v27 bool = (42 <= -130)
    switch 1 {
    case 127:
        for fi28 := 0; fi28 < 1; fi28++ {
            v29 := GenS3{F0: -1}
            fmt.Println(v4, v3, v21)
            _ = fi28
            _ = v29
        }
    case -928:
        if (820 <= -218) {
            var v30 uint8 = v20
            if (((4294967295 == 32768) || true) || ((!v27) && v27)) {
                v2[-2147483648] = ((((("\n\t" + "hello") + v3) + "hello") + ("" + v3)) + (("hello" + ("\n\t" + v3)) + v3))
            } else {
                for fi31 := 0; fi31 < 19; fi31++ {
                    var v32 uint32 = 0
                    _ = fi31
                    _ = v32
                }
            }
            _ = v30
        } else {
            for fi33 := 0; fi33 < 13; fi33++ {
                v20 = v20
                _ = fi33
            }
        }
    case 42:
        v34 := v2[0]
        _ = v34
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        v36 := GenS1{F0: v4, F1: ((-173 < -327) && ((v20 + (v20 >> ((genFn0()) & 63))) != (v20 << (((v1 + v1)) & 63))))}
        v5 = (v3 + "hello")
        _ = fi35
        _ = v36
    }
    var v37 int8 = (^(v21))
    v38 := GenS3{F0: -9223372036854775808}
    v39 := []int{32768}
    v40 := GenS4{F0: 255, F1: 204, F2: 2147483647}
    if (v27 || v27) {
        var v41 uint = (v1 << ((v1) & 63))
        _ = v41
    } else {
        var v42 int = 150
        for fi43 := 0; fi43 < 11; fi43++ {
            if ((fi43 << ((genFn0()) & 63)) <= 127) {
                if len(v39) >= 2 {
                    v39 = v39[1:len(v39)]
                }
                if false {
                    v4 = (v4 % (((v4 * (v4 * 510))) | 1))
                } else {
                    if ((v37 + (v37 + -36)) == (v21 >> ((v1) & 63))) {
                        v44 := GenS5{F0: (v4 + (v4 - v4))}
                        _ = v44
                    }
                }
            } else {
                var v45 string = "hello"
                var v46 bool = (!v27)
                _ = v45
                _ = v46
            }
            if v27 {
                var v47 float64 = 1e10
                v2[(v42 + -757)] = v3
                v37 = (^(v37))
                _ = v47
            }
            _ = fi43
        }
        var v48 int = (-(v42))
        _ = v42
        _ = v48
    }
    v27 = v27
    if (v27 && false) {
        var v49 uint16 = 65535
        switch 4294967295 {
        case 127:
            v21 = (v21 << ((genFn0()) & 63))
        case 0:
            if v27 {
                var v50 uint8 = (^(v20))
                fmt.Println(v20)
                _ = v50
            }
        }
        _ = v49
    } else {
        _ = len(v2)
        fmt.Println(v27, v5, v3)
        switch v5 {
        case "hello":
            v2[-2147483648] = v5
        default:
            var v51 int16 = 1
            _ = v51
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
