package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int; F2 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 int16 }
func genFn0() uint64 {
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := []int{fi5, fi5, fi5}
        _ = fi5
        _ = v6
    }
    switch 1 {
    case 276:
        v7 := GenS1{F0: 2147483647}
        _ = v7
    case 32767:
        var v8 uint64 = 18446744073709551615
        _ = v8
    default:
        if true {
            for fi9 := 0; fi9 < 17; fi9++ {
                v10 := GenS2{F0: (("foo bar" + ("\n\t" + "\n\t")) + ""), F1: fi9, F2: 32767}
                _ = fi9
                _ = v10
            }
            var v11 string = (("\n\t" + ("" + "a")) + ((((((("hello" + ("a" + "hello")) + "a") + "\n\t") + "") + "a") + "foo bar") + "foo bar"))
            _ = v11
        }
    }
    v12 := []int{}
    _ = v12
    return 468
}
func genFn1() uint64 {
    v13 := []int{}
    var v14 bool = true
    v15 := GenS3{F0: 599}
    v16 := GenS2{F0: ((("a" + "") + ("" + "\n\t")) + "\n\t"), F1: 42, F2: 0}
    if false {
        _ = len(v13)
        var v17 int16 = -984
        var v18 uint64 = 1
        _ = v17
        _ = v18
    }
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    return 1
}

func main() {
    var v1 int = 2147483647
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]int)
    var v4 int8 = -1
    switch "" {
    case "a":
        var v19 uint32 = 2
        _ = v19
    case "":
        var v20 uint = 255
        _ = v20
    case "hello":
        fmt.Println(v4)
    default:
        var v21 uint = 448
        _ = v21
    }
    fmt.Println(v4, v1)
    for fi22 := 0; fi22 < 8; fi22++ {
        if true {
            if (237 > 127) {
                var v23 int8 = (v4 - 100)
                v4 = (v4 | (v23 + (v4 << ((172) & 63))))
                v24 := GenS4{F0: 0}
                _ = v23
                _ = v24
            }
        } else {
            for fi25 := 0; fi25 < 16; fi25++ {
                if (796 > 740) {
                    delete(v2, 2)
                    var v26 float32 = 1e10
                    _ = v26
                } else {
                    v4 -= v4
                    switch fi25 {
                    case 128:
                        v1 = (-(fi22))
                    case 0:
                        v4 = v4
                    case 4294967295:
                        fmt.Println(fi25, fi22)
                    default:
                        var v27 uint32 = 568
                        _ = v27
                    }
                    var v28 float32 = -2.25
                    _ = v28
                }
                if false {
                    v4 += (v4 % ((1) | 1))
                } else {
                    var v29 uint16 = 2
                    fmt.Println(fi25)
                    _ = v29
                }
                _ = fi25
            }
            delete(v2, 175)
            v30 := GenS3{F0: 0}
            _ = v30
        }
        v3[65535] = (fi22 / ((860) | 1))
        v31 := GenS3{F0: 112}
        _ = fi22
        _ = v31
    }
    for fi32 := 0; fi32 < 20; fi32++ {
        v33 := GenS5{F0: -32768, F1: 42}
        var v34 float32 = 1e10
        var v35 uint = 474
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    switch ("hello" + "hello") {
    case "\n\t":
        v4 = (^(v4))
    default:
        v36 := GenS5{F0: -387, F1: -32768}
        _ = v36
    }
    if (false || false) {
        var v37 int = v1
        v38 := []int{v37, v37, v37}
        _ = v37
        _ = v38
    } else {
        for fi39 := 0; fi39 < 5; fi39++ {
            var v40 uint32 = 42
            if (-1 != -291) {
                for fi41 := 0; fi41 < 3; fi41++ {
                    switch ("a" + "hello") {
                    case "foo bar":
                        var v42 bool = ("hello" != "a")
                        var v43 uint64 = 442
                        fmt.Println(v42, v42)
                        _ = v42
                        _ = v43
                    case "hello":
                        v44 := GenS2{F0: ("hello" + ("a" + "a")), F1: (fi39 | v1), F2: int16(v40)}
                        _ = v44
                    }
                    _ = fi41
                }
            }
            var v45 int8 = v4
            _ = fi39
            _ = v40
            _ = v45
        }
        switch 32768 {
        case -9223372036854775808:
            fmt.Println(v4, v1)
        case 850:
            if false {
                var v46 uint16 = 652
                for fi47 := 0; fi47 < 9; fi47++ {
                    v1 -= v1
                    _ = fi47
                }
                for fi48 := 0; fi48 < 13; fi48++ {
                    v49 := []int{32767, (-(v1)), v1, 1}
                    switch fi48 {
                    case -9223372036854775808:
                        v50 := GenS4{F0: genFn0()}
                        _ = v50
                    }
                    if (v1 <= fi48) {
                        v46 += v46
                    }
                    _ = fi48
                    _ = v49
                }
                _ = v46
            }
        default:
            var v51 int16 = 1
            _ = v51
        }
    }
    for fi52 := 0; fi52 < 12; fi52++ {
        var v53 string = "\n\t"
        _ = fi52
        _ = v53
    }
    switch (v1 + (v1 % ((v1) | 1))) {
    case 2147483647:
        v54 := GenS3{F0: 1}
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
