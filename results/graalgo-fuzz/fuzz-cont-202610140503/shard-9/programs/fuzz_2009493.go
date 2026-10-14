package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int32 }
type GenS2 struct { F0 uint; F1 int8 }
func genFn0() int16 {
    var v7 uint = 255
    if false {
        if false {
            v8 := GenS1{F0: 1, F1: -623, F2: 1}
            v9 := []int{32767}
            _ = v8
            _ = v9
        } else {
            if false {
                fmt.Println(v7)
            } else {
                v10 := GenS1{F0: 42, F1: 0, F2: -882}
                _ = v10
            }
            v7 = (v7 % ((v7) | 1))
            for fi11 := 0; fi11 < 10; fi11++ {
                for fi12 := 0; fi12 < 8; fi12++ {
                    v13 := GenS1{F0: 0, F1: 2147483647, F2: 1}
                    fmt.Println(v7, v7, fi12)
                    v7 = v7
                    _ = fi12
                    _ = v13
                }
                v14 := []int{(fi11 >> (((v7 ^ (v7 << (((v7 * (v7 + (^(v7))))) & 63)))) & 63)), fi11}
                for fi15 := 0; fi15 < 4; fi15++ {
                    v7 += 255
                    _ = fi15
                }
                _ = fi11
                _ = v14
            }
        }
        v7 = v7
    }
    switch ("" + "hello") {
    case "foo bar":
        var v16 float32 = 1e10
        _ = v16
    }
    switch -954 {
    case -9223372036854775808:
        v17 := GenS1{F0: 1, F1: -2147483648, F2: 2147483647}
        _ = v17
    case 4294967295:
        v7 -= (^(v7))
    default:
        var v18 float64 = 0.1
        _ = v18
    }
    fmt.Println(v7)
    _ = v7
    return -763
}
func genFn1(p0 bool, p1 uint16) int {
    if p0 {
        p1 -= 430
        p0 = ((!(-1 == -99)) || true)
        v19 := GenS1{F0: 1000, F1: -226, F2: 349}
        _ = v19
    } else {
        p1 -= p1
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        var v21 int16 = -1
        p0 = p0
        for fi22 := 0; fi22 < 18; fi22++ {
            p0 = true
            p0 = p0
            v23 := GenS1{F0: genFn0(), F1: -1, F2: 1}
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    v24 := []int{1, 145}
    var v25 float64 = (0.1 / 3.14159)
    var v26 uint = 570
    _ = p0
    _ = p1
    _ = v24
    _ = v25
    _ = v26
    return 127
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"\n\t", "foo bar", (("hello" + ("\n\t" + (("a" + "a") + ""))) + (("\n\t" + ("hello" + "")) + "foo bar"))}
    var v3 uint = 737
    var v4 uint = (v3 + 299)
    v5 := []string{"foo bar", ("foo bar" + "\n\t"), "hello"}
    v6 := make(map[uint32]int)
    v27 := GenS1{F0: -918, F1: 1, F2: 0}
    var v28 int8 = 0
    v3 -= v3
    var v29 float64 = (1.5 - -2.25)
    for fi30 := 0; fi30 < 12; fi30++ {
        delete(v6, 752)
        _ = fi30
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        var v32 int64 = 651
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        v34 := GenS1{F0: 598, F1: -11, F2: 0}
        for fi35 := 0; fi35 < 18; fi35++ {
            switch "\n\t" {
            case "hello":
                v28 -= v28
            case "":
                var v36 int = genFn1(true, 653)
                _ = v36
            case "\n\t":
                v29 = (v29 + v29)
            }
            fmt.Println(fi35, v29, fi35)
            _ = fi35
        }
        var v37 uint64 = 618
        _ = fi33
        _ = v34
        _ = v37
    }
    v28 += v28
    var v38 int16 = -32768
    var v39 int = 32767
    var v40 float64 = (1.5 - v29)
    v41 := GenS2{F0: 4294967295, F1: -1}
    v40 = (1.5 / v29)
    fmt.Println(v29)
    delete(v6, 0)
    fmt.Println(v4)
    _ = len(v6)
    var v42 uint16 = 823
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
