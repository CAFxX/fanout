package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 int32; F2 uint }
func genFn0(p0 uint8, p1 string) uint {
    p0 = p0
    var v4 string = (((("foo bar" + "hello") + "a") + (p1 + "hello")) + (("hello" + p1) + p1))
    _ = p0
    _ = p1
    _ = v4
    return 229
}

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 1
    v3 := make(map[int]string)
    _ = len(v3)
    var v5 string = v1
    v5 = v1
    fmt.Println(v5, v1, v2)
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 float32 = 1e10
        _ = fi6
        _ = v7
    }
    v1 = ((("foo bar" + v1) + (v1 + v1)) + v5)
    for fi8 := 0; fi8 < 14; fi8++ {
        v1 = "a"
        _ = fi8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        v5 = (v5 + (v1 + v5))
        var v10 int8 = 100
        var v11 string = ("hello" + ("hello" + (v5 + (v5 + (v5 + (("foo bar" + v1) + (v1 + v5)))))))
        _ = fi9
        _ = v10
        _ = v11
    }
    if true {
        for fi12 := 0; fi12 < 4; fi12++ {
            v2 += v2
            var v13 uint8 = v2
            _ = fi12
            _ = v13
        }
    } else {
        if (!(!false)) {
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 bool = true
                v2 += (^(v2))
                _ = fi14
                _ = v15
            }
            var v16 uint16 = 31
            _ = v16
        } else {
            v2 -= (v2 << ((741) & 63))
        }
        v2 += 42
    }
    v17 := GenS1{F0: 42}
    if false {
        v1 = v5
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        v19 := GenS2{F0: "a", F1: fi18}
        if (3.14159 >= 1e10) {
            if true {
                v1 = (("foo bar" + "hello") + "foo bar")
                var v20 uint32 = 768
                _ = v20
            } else {
                v1 = v1
                v21 := GenS1{F0: (fi18 - (-(fi18)))}
                var v22 string = v5
                _ = v21
                _ = v22
            }
        }
        fmt.Println(v1, v1)
        _ = fi18
        _ = v19
    }
    if (42 > -69) {
        for fi23 := 0; fi23 < 9; fi23++ {
            v24 := GenS3{F0: 556, F1: 4294967295}
            var v25 float32 = 0.0
            for fi26 := 0; fi26 < 17; fi26++ {
                var v27 int8 = 100
                _ = fi26
                _ = v27
            }
            _ = fi23
            _ = v24
            _ = v25
        }
        if true {
            for fi28 := 0; fi28 < 8; fi28++ {
                switch "a" {
                case "foo bar":
                    var v29 float64 = -2.25
                    _ = v29
                case "hello":
                    v1 = (v5 + ((v5 + v1) + v1))
                default:
                    v5 = v1
                }
                v30 := []int{(^(fi28))}
                _ = fi28
                _ = v30
            }
            fmt.Println(v1)
            _ = len(v3)
        } else {
            v5 = v5
        }
    } else {
        for fi31 := 0; fi31 < 2; fi31++ {
            fmt.Println(v2, fi31, fi31)
            if (!true) {
                v32 := []int{(fi31 - (fi31 & 32768)), fi31, fi31}
                v1 = "hello"
                _ = v32
            }
            v33 := v3[fi31]
            _ = fi31
            _ = v33
        }
        var v34 int16 = -388
        v35 := GenS4{F0: 65535}
        _ = v34
        _ = v35
    }
    if true {
        v2 -= v2
        delete(v3, 2147483647)
        var v36 int32 = 1
        _ = v36
    }
    var v37 bool = false
    if v37 {
        v38 := GenS5{F0: 32767, F1: -1, F2: 265}
        v5 = "\n\t"
        var v39 uint32 = uint32(v2)
        _ = v38
        _ = v39
    }
    if v37 {
        if (331 > -1) {
            if (-149 <= 32767) {
                fmt.Println(v2)
            }
            fmt.Println(v37)
        } else {
            v40 := GenS2{F0: (((v5 + ("" + "foo bar")) + v1) + ("a" + "a")), F1: -9223372036854775808}
            var v41 int32 = 235
            var v42 int32 = v41
            _ = v40
            _ = v41
            _ = v42
        }
        v43 := []int{128, 4294967295, 590, -193}
        _ = v43
    }
    v44 := []int{-2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v37
    _ = v44
    _ = v5
}
