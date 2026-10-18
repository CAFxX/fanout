package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int8; F2 uint }
type GenS3 struct { F0 uint8; F1 int64; F2 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint; F1 uint8 }
func genFn0(p0 int8, p1 bool, p2 string) uint32 {
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 int = fi5
        v7 := GenS1{F0: 1000}
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 bool = p1
            p1 = (-2147483648 < (fi5 | (v6 >> ((18446744073709551615) & 63))))
            _ = fi8
            _ = v9
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    v10 := []int{420, 936, 9223372036854775807}
    for fi11 := 0; fi11 < 13; fi11++ {
        if true {
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 uint32 = 2
                _ = fi12
                _ = v13
            }
            v14 := GenS1{F0: 32767}
            p1 = true
            _ = v14
        }
        p1 = (0 >= 42)
        _ = fi11
    }
    switch 23 {
    case 4294967295:
        var v15 int32 = 1
        _ = v15
    default:
        if (0 > 1) {
            var v16 uint32 = 0
            p0 = 0
            fmt.Println(p2, p0, p2)
            _ = v16
        } else {
            for fi17 := 0; fi17 < 16; fi17++ {
                if p1 {
                    var v18 uint16 = 255
                    _ = v18
                } else {
                    _ = len(v10)
                    switch (fi17 ^ fi17) {
                    case 4294967295:
                        for fi19 := 0; fi19 < 18; fi19++ {
                            var v20 string = "hello"
                            _ = fi19
                            _ = v20
                        }
                    default:
                        switch p2 {
                        case "":
                            var v21 int = 127
                            _ = v21
                        }
                    }
                }
                _ = len(v10)
                _ = fi17
            }
            for fi22 := 0; fi22 < 6; fi22++ {
                for fi23 := 0; fi23 < 15; fi23++ {
                    switch (fi22 | (-(fi23))) {
                    case 2147483647:
                        if len(v10) > 0 {
                            v10[0] = (fi23 ^ (^(fi22)))
                        }
                    default:
                        p2 = p2
                    }
                    p1 = true
                    var v24 int16 = -32768
                    var v25 string = ""
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                v26 := GenS1{F0: -535}
                _ = fi22
                _ = v26
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 42
}
func genFn1(p0 int32, p1 bool) uint16 {
    switch 898 {
    case 128:
        v27 := GenS1{F0: 1000}
        _ = v27
    }
    if (1 >= 100) {
        p1 = (2 < 65535)
        switch ("foo bar" + "a") {
        case "":
            var v28 int64 = -1
            _ = v28
        case "hello":
            if (genFn0(100, p1, "a") >= 2) {
                var v29 uint64 = 284
                switch "\n\t" {
                case "a":
                    for fi30 := 0; fi30 < 18; fi30++ {
                        if (p0 != 0) {
                            var v31 float32 = 0.0
                            p0 = p0
                            _ = v31
                        } else {
                            v32 := []int{}
                            _ = v32
                        }
                        if (1.5 >= (1.5 + ((float64(1.5) * 3.14159) / 0.1))) {
                            switch fi30 {
                            case 989:
                                for fi33 := 0; fi33 < 9; fi33++ {
                                    v29 -= v29
                                    v34 := []int{}
                                    _ = fi33
                                    _ = v34
                                }
                            }
                            v35 := []int{fi30, (fi30 ^ fi30)}
                            _ = v35
                        } else {
                            v36 := GenS1{F0: 32767}
                            _ = v36
                        }
                        _ = fi30
                    }
                case "":
                    var v37 uint64 = 712
                    _ = v37
                case "foo bar":
                    var v38 bool = p1
                    _ = v38
                }
                _ = v29
            }
        case "foo bar":
            p0 = (p0 + 42)
        }
        v39 := []int{-895, 42, -9223372036854775808}
        _ = v39
    } else {
        switch "foo bar" {
        case "":
            v40 := []int{-1}
            _ = v40
        case "hello":
            switch 9223372036854775807 {
            case 568:
                fmt.Println(p0)
            case -992:
                if true {
                    v41 := []int{-357, -1, 127}
                    p0 = (p0 % (((p0 >> ((65535) & 63))) | 1))
                    _ = v41
                } else {
                    var v42 uint16 = 32767
                    _ = v42
                }
            case 2147483647:
                p0 -= p0
            default:
                for fi43 := 0; fi43 < 16; fi43++ {
                    var v44 string = (("hello" + "hello") + "foo bar")
                    var v45 int16 = -32768
                    v46 := GenS2{F0: false, F1: 127, F2: 42}
                    _ = fi43
                    _ = v44
                    _ = v45
                    _ = v46
                }
            }
        default:
            if p1 {
                v47 := GenS3{F0: 128, F1: 9223372036854775807, F2: 71}
                v48 := GenS2{F0: ((-(p0)) == (p0 * p0)), F1: -126, F2: 18446744073709551615}
                _ = v47
                _ = v48
            }
        }
        for fi49 := 0; fi49 < 5; fi49++ {
            for fi50 := 0; fi50 < 10; fi50++ {
                if (127 < 1) {
                    v51 := []int{fi49, (-(fi49))}
                    _ = v51
                }
                v52 := GenS4{F0: p0}
                _ = fi50
                _ = v52
            }
            var v53 float32 = 1.5
            p1 = ((3.14159 - 1e10) >= (-2.25 + 0.0))
            _ = fi49
            _ = v53
        }
    }
    p1 = (false || (!p1))
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 int = 0
    var v2 int32 = 141
    var v3 string = "\n\t"
    v4 := make(map[uint32]uint32)
    v2 += v2
    var v54 string = "\n\t"
    var v55 uint32 = 65535
    v56 := GenS2{F0: (-68 > -12), F1: -61, F2: 2}
    v57 := []int{(^(v1)), int(v55), 42, v1}
    switch v1 {
    case -9223372036854775808:
        v58 := GenS5{F0: 9223372036854775807}
        _ = v58
    case 128:
        if len(v57) >= 2 {
            v57 = v57[1:len(v57)]
        }
    case 42:
        var v59 int16 = 32767
        _ = v59
    default:
        if (!(v55 == v55)) {
            if false {
                if true {
                    _ = len(v4)
                    var v60 uint = 81
                    _ = v60
                } else {
                    v4[(v55 & v55)] = genFn0(127, true, ("hello" + (("a" + ("" + (v54 + v54))) + ("foo bar" + (v54 + (("" + v3) + "hello"))))))
                    v54 = "hello"
                    v2 = (v2 % ((v2) | 1))
                }
                v55 = genFn0(1, true, "a")
                switch 2147483647 {
                case -1:
                    v1 = v1
                case 378:
                    var v61 bool = (!true)
                    _ = v61
                }
            }
        } else {
            v62 := GenS6{F0: 42, F1: 25}
            _ = v62
        }
    }
    v63 := v4[0]
    fmt.Println(v3)
    for fi64 := 0; fi64 < 1; fi64++ {
        fmt.Println(v3, v54, v55)
        _ = fi64
    }
    var v65 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v63:", v63)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v63
    _ = v65
}
