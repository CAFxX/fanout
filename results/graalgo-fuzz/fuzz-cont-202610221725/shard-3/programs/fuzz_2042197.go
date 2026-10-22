package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint }
func genFn0(p0 int) uint {
    for fi6 := 0; fi6 < 15; fi6++ {
        p0 = (fi6 * p0)
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 uint32 = 42
            var v9 int16 = 56
            _ = fi7
            _ = v8
            _ = v9
        }
        if true {
            for fi10 := 0; fi10 < 8; fi10++ {
                v11 := []int{(p0 * (p0 << ((681) & 63))), 525}
                _ = fi10
                _ = v11
            }
            fmt.Println(fi6)
            p0 = (p0 + (p0 - (fi6 << ((650) & 63))))
        }
        _ = fi6
    }
    if false {
        v12 := []int{(p0 | (p0 + p0)), (p0 - p0), (-(p0))}
        var v13 int = -669
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 int16 = 32767
            _ = fi14
            _ = v15
        }
        _ = v12
        _ = v13
    } else {
        p0 = (p0 / ((-111) | 1))
        fmt.Println(p0)
        var v16 uint32 = 0
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        if (!false) {
            var v18 int = 9223372036854775807
            for fi19 := 0; fi19 < 5; fi19++ {
                v20 := []int{(^(fi19))}
                switch v18 {
                case -9223372036854775808:
                    p0 = 565
                case 314:
                    switch (-(fi17)) {
                    case 4294967295:
                        var v21 int32 = -2147483648
                        _ = v21
                    }
                default:
                    var v22 float64 = (-2.25 * -2.25)
                    _ = v22
                }
                v18 = v18
                _ = fi19
                _ = v20
            }
            _ = v18
        }
        _ = fi17
    }
    p0 = 42
    var v23 int = p0
    _ = p0
    _ = v23
    return 867
}
func genFn1(p0 int8) string {
    v24 := GenS1{F0: 1000, F1: 805}
    var v25 string = "\n\t"
    if (true && (610 != 7)) {
        if true {
            p0 = p0
        } else {
            v25 = v25
            if ((^(p0)) >= 1) {
                var v26 int64 = 863
                v27 := []int{777, -9223372036854775808, 42, -289}
                v25 = ("\n\t" + ("a" + v25))
                _ = v26
                _ = v27
            }
            var v28 uint16 = 255
            _ = v28
        }
        for fi29 := 0; fi29 < 15; fi29++ {
            v30 := []int{(fi29 ^ fi29)}
            for fi31 := 0; fi31 < 8; fi31++ {
                fmt.Println(p0)
                _ = fi31
            }
            p0 = p0
            _ = fi29
            _ = v30
        }
        var v32 int16 = 20
        _ = v32
    } else {
        v33 := []int{4294967295, 2147483647, 9223372036854775807}
        var v34 int8 = (^(p0))
        v25 = v25
        _ = v33
        _ = v34
    }
    var v35 string = ""
    _ = p0
    _ = v24
    _ = v25
    _ = v35
    return v35
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = -2.25
    var v3 bool = ((("" + "\n\t") + ("hello" + "a")) == (("" + "\n\t") + "foo bar"))
    var v4 int = 54
    var v5 float32 = 0.0
    for fi36 := 0; fi36 < 5; fi36++ {
        if (((0.1 * 1e10) / 1.5) >= 1.5) {
            if (true || v3) {
                delete(v1, 4294967295)
            } else {
                v37 := GenS2{F0: 0, F1: 42}
                _ = v37
            }
            switch "a" {
            case "\n\t":
                v38 := v1[42]
                _ = v38
            case "a":
                var v39 float64 = v2
                _ = v39
            }
        }
        if (1.5 != (v5 / (v5 / 0.1))) {
            fmt.Println(v2, v5, v5)
        } else {
            var v40 uint8 = 2
            v41 := []int{}
            _ = v40
            _ = v41
        }
        for fi42 := 0; fi42 < 13; fi42++ {
            if (829 < 32767) {
                var v43 uint32 = 42
                switch (v4 & (v4 | 1)) {
                case 9223372036854775807:
                    v5 = (-2.25 - ((v5 / float32(1e10)) / v5))
                default:
                    v3 = v3
                }
                _ = v43
            } else {
                if true {
                    v2 -= (v2 - 0.1)
                    var v44 string = ("a" + "hello")
                    switch (fi42 % ((552) | 1)) {
                    case 32767:
                        v44 = "hello"
                    case 0:
                        v45 := GenS3{F0: 0, F1: (!(!(v3 && (-549 > 42)))), F2: 580}
                        _ = v45
                    case 2147483647:
                        var v46 uint16 = 65535
                        _ = v46
                    }
                    fmt.Println(v2, fi36, v5)
                    delete(v1, 1)
                    var v47 uint8 = 255
                    _ = v44
                    _ = v47
                } else {
                    var v48 int = (fi36 | -2147483648)
                    _ = v48
                }
                var v49 int8 = 42
                if (0 > (v49 & (v49 & (v49 - (v49 * v49))))) {
                    var v50 uint32 = 175
                    fmt.Println(v2, v49, fi42)
                    delete(v1, (v50 | 20))
                    if v3 {
                        v1[4294967295] = "hello"
                    }
                    _ = v50
                } else {
                    var v51 uint32 = 65535
                    v1[(v51 % ((v51) | 1))] = "a"
                    var v52 uint64 = 353
                    _ = v51
                    _ = v52
                }
                _ = v49
            }
            switch "foo bar" {
            case "foo bar":
                var v53 int32 = -1
                fmt.Println(fi42, fi42)
                v54 := GenS2{F0: 4294967295, F1: 501}
                var v55 int32 = (-(v53))
                var v56 int8 = 127
                _ = v53
                _ = v54
                _ = v55
                _ = v56
            case "":
                var v57 uint32 = 2147483647
                _ = v57
            case "hello":
                v3 = (!v3)
            }
            _ = fi42
        }
        _ = fi36
    }
    v3 = true
    switch "a" {
    case "\n\t":
        v3 = v3
    }
    v4 = v4
    fmt.Println(v3, v3, v3)
    v58 := []int{v4, (v4 + 42), v4, 0}
    if (2 >= 42) {
        var v59 uint16 = 554
        _ = v59
    }
    var v60 uint8 = 128
    v5 += v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v58:", len(v58), v58)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v60
}
