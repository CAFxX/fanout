package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 bool }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 uint }
func genFn0(p0 uint16, p1 uint, p2 uint16) uint64 {
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := []int{fi7}
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            p0 = 32767
            var v11 string = "\n\t"
            var v12 uint8 = 0
            if false {
                var v13 string = (v11 + "")
                var v14 int = fi10
                fmt.Println(v14, v14, v11)
                _ = v13
                _ = v14
            } else {
                var v15 float64 = 1e10
                var v16 int32 = 622
                switch fi9 {
                case 1:
                    fmt.Println(fi10, fi9)
                case 127:
                    v17 := []int{}
                    v18 := GenS1{F0: 42, F1: 0, F2: false}
                    fmt.Println(p1)
                    _ = v17
                    _ = v18
                case 9223372036854775807:
                    var v19 int = fi10
                    _ = v19
                }
                _ = v15
                _ = v16
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        var v20 uint32 = 1
        _ = fi9
        _ = v20
    }
    if true {
        var v21 uint64 = 355
        switch ("\n\t" + ("\n\t" + ("hello" + ("a" + ("hello" + ((("foo bar" + ("a" + "a")) + "\n\t") + "hello")))))) {
        case "hello":
            fmt.Println(p2)
        case "foo bar":
            p2 += (p2 | (p0 * 32767))
        default:
            var v22 int64 = 471
            _ = v22
        }
        _ = v21
    } else {
        v23 := []int{1, -1}
        v24 := GenS1{F0: -424, F1: 100, F2: (1 != (p0 | (^(p0))))}
        _ = v23
        _ = v24
    }
    v25 := []int{4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return 42
}
func genFn1() int32 {
    v26 := []int{-1, -9223372036854775808, 4294967295, -1}
    for fi27 := 0; fi27 < 3; fi27++ {
        if (false || true) {
            if (1 > 585) {
                v28 := GenS1{F0: 168, F1: -128, F2: (42 < -32768)}
                var v29 float64 = (1.5 / 0.1)
                _ = v28
                _ = v29
            }
            var v30 float64 = 3.14159
            _ = v30
        }
        fmt.Println(fi27)
        _ = fi27
    }
    _ = v26
    return 118
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{}
    var v3 int16 = 1
    v4 := make(map[uint32]bool)
    var v5 uint = 823
    var v31 int16 = (v3 & (v3 * -782))
    v32 := GenS1{F0: 0, F1: 116, F2: (true && (!(118 <= -9223372036854775808)))}
    v33 := GenS1{F0: 168, F1: 100, F2: true}
    for fi34 := 0; fi34 < 9; fi34++ {
        v35 := GenS1{F0: -618, F1: 1, F2: true}
        _ = fi34
        _ = v35
    }
    if false {
        v36 := GenS2{F0: 4294967295, F1: 42}
        _ = v36
    }
    v37 := GenS2{F0: 90, F1: 19}
    var v38 int32 = -657
    v39 := v1["\n\t"]
    v38 -= 0
    if ((^(v5)) >= (v5 & v5)) {
        for fi40 := 0; fi40 < 17; fi40++ {
            if ((-(fi40)) > (fi40 | (fi40 << ((42) & 63)))) {
                v41 := GenS3{F0: 1}
                for fi42 := 0; fi42 < 9; fi42++ {
                    v38 = -1
                    fmt.Println(v3)
                    _ = fi42
                }
                if true {
                    switch "foo bar" {
                    case "":
                        v5 = 42
                    case "hello":
                        for fi43 := 0; fi43 < 3; fi43++ {
                            v5 = (v5 % ((2) | 1))
                            _ = fi43
                        }
                    default:
                        var v44 string = (v39 + ((v39 + ((v39 + v39) + ("" + (v39 + "\n\t")))) + "foo bar"))
                        _ = v44
                    }
                    if false {
                        v4[2] = false
                        v45 := GenS1{F0: 9223372036854775807, F1: 127, F2: false}
                        _ = v45
                    }
                } else {
                    var v46 uint64 = 4294967295
                    _ = v46
                }
                _ = v41
            } else {
                var v47 uint32 = 65535
                _ = v47
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v5 = (v5 << ((v5) & 63))
            _ = fi40
        }
    }
    v48 := GenS2{F0: genFn0(42, (v5 % (((v5 << ((v5) & 63))) | 1)), 2), F1: 44}
    v5 += v5
    var v49 int64 = -1
    var v50 int = 32767
    var v51 int64 = (-(v49))
    fmt.Println(v51, v51)
    v52 := GenS4{F0: true, F1: (v5 * v5)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
}
