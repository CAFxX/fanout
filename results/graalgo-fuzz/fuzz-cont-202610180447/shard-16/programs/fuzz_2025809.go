package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int; F2 uint8 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int64; F1 int }
func genFn0(p0 uint8, p1 uint8) int16 {
    var v5 bool = (1 > 1)
    fmt.Println(v5)
    _ = p0
    _ = p1
    _ = v5
    return 1
}
func genFn1(p0 int32, p1 uint16) bool {
    p0 = (p0 + 1)
    if ("\n\t" > "\n\t") {
        for fi6 := 0; fi6 < 13; fi6++ {
            for fi7 := 0; fi7 < 1; fi7++ {
                if true {
                    p0 += -786
                }
                var v8 uint32 = 2
                var v9 string = ("foo bar" + "a")
                _ = fi7
                _ = v8
                _ = v9
            }
            if false {
                v10 := []int{32767}
                p0 = p0
                _ = v10
            } else {
                switch (fi6 & fi6) {
                case 923:
                    fmt.Println(p1, fi6)
                case 1:
                    v11 := []int{(fi6 / ((0) | 1)), -1, (fi6 ^ 32767), fi6}
                    _ = v11
                }
                for fi12 := 0; fi12 < 13; fi12++ {
                    v13 := GenS1{F0: true}
                    v14 := GenS2{F0: 922, F1: (fi12 & (^(fi6))), F2: 42}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
            fmt.Println(p0, fi6)
            _ = fi6
        }
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 uint = 190
            for fi17 := 0; fi17 < 11; fi17++ {
                if false {
                    fmt.Println(fi17, p1, fi17)
                } else {
                    var v18 int32 = -755
                    _ = v18
                }
                _ = fi17
            }
            _ = fi15
            _ = v16
        }
        fmt.Println(p0)
    } else {
        var v19 bool = (42 <= 4294967295)
        _ = v19
    }
    switch 4294967295 {
    case 42:
        p1 = p1
    case 0:
        if false {
            if false {
                switch ("a" + "foo bar") {
                case "a":
                    fmt.Println(p1, p1)
                default:
                    for fi20 := 0; fi20 < 16; fi20++ {
                        var v21 int32 = (p0 * (^(p0)))
                        _ = fi20
                        _ = v21
                    }
                }
            }
            var v22 uint8 = 94
            _ = v22
        }
    case 4294967295:
        var v23 float32 = (0.0 - 3.14159)
        _ = v23
    }
    if true {
        switch 128 {
        case 32767:
            for fi24 := 0; fi24 < 8; fi24++ {
                fmt.Println(p1, fi24, p1)
                var v25 bool = true
                _ = fi24
                _ = v25
            }
        case 944:
            fmt.Println(p1, p0)
        case -1:
            v26 := []int{4294967295, -9223372036854775808, 1, 431}
            _ = v26
        default:
            for fi27 := 0; fi27 < 6; fi27++ {
                p1 = p1
                var v28 int32 = 570
                _ = fi27
                _ = v28
            }
        }
    }
    v29 := []int{-227, 127, 824, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v29
    return true
}

func main() {
    var v1 uint64 = 793
    var v2 int16 = 0
    var v3 bool = false
    v4 := make(map[string]string)
    v30 := []int{1}
    for fi31 := 0; fi31 < 18; fi31++ {
        for fi32 := 0; fi32 < 8; fi32++ {
            v2 += v2
            v2 += v2
            v1 = 886
            var v33 uint = 255
            switch ("a" + "hello") {
            case "":
                if v3 {
                    var v34 bool = (340 == 789)
                    v35 := GenS3{F0: v2, F1: v33}
                    v34 = v34
                    _ = v34
                    _ = v35
                }
            }
            v1 = v1
            switch (fi32 * fi32) {
            case 2147483647:
                var v36 int8 = 22
                _ = v36
            case 4294967295:
                _ = len(v4)
            default:
                fmt.Println(v1, v2)
                v2 += genFn0(127, 0)
            }
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    v37 := GenS1{F0: (!(32767 == 42))}
    v3 = ((0.1 >= (float64(-2.25) + 1.5)) || (!v3))
    var v38 int8 = 88
    if len(v30) > 0 {
        v30[0] = 79
    }
    v39 := GenS3{F0: genFn0(255, 31), F1: 2}
    if false {
        v40 := GenS4{F0: -9223372036854775808, F1: 843}
        _ = v40
    } else {
        if (v2 >= v2) {
            var v41 int8 = -60
            var v42 string = "\n\t"
            _ = v41
            _ = v42
        }
        fmt.Println(v3)
    }
    var v43 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v43
}
