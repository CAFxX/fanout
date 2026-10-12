package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int32; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint64 }
func genFn0(p0 uint, p1 uint16, p2 bool) uint32 {
    p1 += p1
    switch -623 {
    case 128:
        if p2 {
            v4 := GenS1{F0: "a", F1: 0}
            _ = v4
        } else {
            fmt.Println(p0, p2, p1)
        }
    case 4294967295:
        switch (("a" + "foo bar") + "a") {
        case "a":
            p0 += 532
        default:
            p1 -= 65535
        }
    default:
        v5 := []int{-844}
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint8 = 127
    var v2 uint32 = 2147483647
    var v3 bool = true
    v1 = v1
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := []int{932, (fi6 % (((fi6 >> ((342) & 63))) | 1)), (fi6 << ((255) & 63)), (fi6 >> ((1) & 63))}
        var v8 uint = 255
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        switch "hello" {
        case "hello":
            v10 := []int{(^(fi9)), fi9, (fi9 / (((fi9 / (((-(fi9))) | 1))) | 1))}
            _ = v10
        case "a":
            for fi11 := 0; fi11 < 18; fi11++ {
                v12 := GenS2{F0: 348, F1: 792, F2: 575}
                _ = fi11
                _ = v12
            }
        default:
            var v13 uint64 = 2
            _ = v13
        }
        if ((32767 > -1) && (!v3)) {
            if false {
                var v14 string = "\n\t"
                v2 += (v2 & 42)
                var v15 int32 = -2147483648
                _ = v14
                _ = v15
            } else {
                if v3 {
                    v16 := []int{fi9, fi9}
                    var v17 float32 = 1.5
                    _ = v16
                    _ = v17
                }
            }
            if v3 {
                if (2 <= genFn0(2, 32767, (65535 <= 65535))) {
                    for fi18 := 0; fi18 < 16; fi18++ {
                        v19 := []int{}
                        v1 = (^(v1))
                        _ = fi18
                        _ = v19
                    }
                    fmt.Println(v2)
                }
            }
            v20 := []int{}
            _ = v20
        }
        _ = fi9
    }
    if (true || (v3 || (2 != 255))) {
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 uint16 = 65535
            v3 = true
            v23 := GenS2{F0: 1, F1: 2, F2: (v22 + 712)}
            _ = fi21
            _ = v22
            _ = v23
        }
        for fi24 := 0; fi24 < 2; fi24++ {
            v1 += 0
            var v25 uint8 = v1
            _ = fi24
            _ = v25
        }
        v3 = true
    }
    v26 := GenS3{F0: 2, F1: (v1 / (((^(v1))) | 1))}
    var v27 string = "hello"
    v28 := GenS2{F0: 3, F1: 1, F2: 624}
    v29 := []int{0, 0, 128, -1}
    var v30 bool = ((1e10 + 1e10) >= (1e10 - 3.14159))
    if true {
        if v3 {
            fmt.Println(v2, v3, v3)
        }
        for fi31 := 0; fi31 < 3; fi31++ {
            v32 := GenS1{F0: (v27 + ("a" + ((v27 + v27) + v27))), F1: 109}
            _ = fi31
            _ = v32
        }
        var v33 string = v27
        _ = v33
    } else {
        if v30 {
            for fi34 := 0; fi34 < 19; fi34++ {
                if v3 {
                    var v35 uint8 = 127
                    _ = v35
                }
                _ = len(v29)
                _ = fi34
            }
            switch 848 {
            case 296:
                v27 = (v27 + ("\n\t" + ((((v27 + v27) + "hello") + v27) + "a")))
            }
            if v3 {
                var v36 bool = v3
                v27 = (v27 + (v27 + ((v27 + "hello") + v27)))
                _ = v36
            }
        } else {
            var v37 string = v27
            if len(v29) >= 2 {
                v29 = v29[1:len(v29)]
            }
            _ = len(v29)
            _ = v37
        }
        v27 = ""
    }
    if len(v29) > 0 {
        v29[0] = 32768
    }
    for fi38 := 0; fi38 < 16; fi38++ {
        if (699 >= -32768) {
            if v30 {
                var v39 int64 = -419
                for fi40 := 0; fi40 < 13; fi40++ {
                    v39 += (v39 >> ((1) & 63))
                    _ = fi40
                }
                var v41 uint8 = (^(v1))
                _ = v39
                _ = v41
            }
        } else {
            if len(v29) > 0 {
                v29[0] = (fi38 | fi38)
            }
            v27 = v27
            var v42 uint32 = genFn0(42, 32767, (0 > 0))
            _ = v42
        }
        _ = fi38
    }
    v2 -= 52
    v43 := GenS4{F0: 925, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v43
}
