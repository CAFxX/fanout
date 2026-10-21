package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 uint8; F2 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 uint32 }
func genFn0(p0 int64, p1 int32, p2 int) int8 {
    v6 := []int{p2, p2}
    var v7 string = ""
    v7 = v7
    if true {
        if (-2.25 != 0.0) {
            var v8 uint32 = 1
            _ = v8
        } else {
            p2 -= -1
            var v9 int64 = p0
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 13; fi10++ {
            fmt.Println(p1, v7, v7)
            if len(v6) > 0 {
                v6[0] = (p2 / ((fi10) | 1))
            }
            _ = fi10
        }
        if (false || (2147483647 == 838)) {
            var v11 uint16 = 344
            v12 := GenS1{F0: "hello", F1: -1, F2: 599}
            _ = v11
            _ = v12
        }
        if (!false) {
            v13 := GenS2{F0: 42}
            fmt.Println(p0, p0)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v13
        }
    }
    v7 = ((v7 + v7) + (v7 + (("hello" + (v7 + v7)) + (v7 + v7))))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return int8(p1)
}
func genFn1(p0 int, p1 uint, p2 string) string {
    for fi14 := 0; fi14 < 4; fi14++ {
        p2 = p2
        _ = fi14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        switch p0 {
        case 32767:
            var v16 int32 = 42
            _ = v16
        case 518:
            fmt.Println(fi15, p2, fi15)
        case 128:
            switch (p2 + p2) {
            case "a":
                v17 := GenS2{F0: 4294967295}
                _ = v17
            }
        }
        p1 -= (p1 | (p1 | 255))
        p0 = -9223372036854775808
        _ = fi15
    }
    var v18 int64 = 260
    if (1.5 != 0.1) {
        v19 := []int{(p0 >> (((p1 % ((2) | 1))) & 63))}
        var v20 int16 = 0
        var v21 int = (p0 | (p0 | p0))
        _ = v19
        _ = v20
        _ = v21
    }
    fmt.Println(p2, v18, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return ""
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[int]bool)
    var v3 float64 = -2.25
    var v4 int = 0
    var v5 uint16 = 840
    v3 += float64((1e10 * (-2.25 + -2.25)))
    for fi22 := 0; fi22 < 16; fi22++ {
        switch ("a" + ("\n\t" + "\n\t")) {
        case "\n\t":
            var v23 uint16 = v5
            _ = v23
        case "hello":
            v1 = 42
        }
        _ = fi22
    }
    v24 := GenS3{F0: "a", F1: (v1 + 128), F2: 0}
    fmt.Println(v4)
    for fi25 := 0; fi25 < 18; fi25++ {
        var v26 int64 = 350
        v4 = 127
        var v27 string = "foo bar"
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 uint32 = 2147483648
    v29 := GenS4{F0: (v1 / (((v1 * (v1 >> ((4294967295) & 63)))) | 1))}
    var v30 bool = ((v1 < (v1 | (v1 / (((v1 & (^(v1)))) | 1)))) && false)
    for fi31 := 0; fi31 < 18; fi31++ {
        if ((v5 - (v5 + (v5 ^ 1))) != v5) {
            var v32 uint16 = v5
            v30 = (-9223372036854775808 != -1)
            _ = v32
        }
        _ = fi31
    }
    v33 := []int{(v4 - 283)}
    var v34 uint8 = v1
    var v35 uint32 = v28
    v36 := GenS5{F0: (^(v1)), F1: 190}
    if true {
        v30 = (((v30 && (!((v5 * v5) != v5))) && v30) || (1 != 18446744073709551615))
    } else {
        var v37 string = ""
        var v38 int64 = 1
        if (!v30) {
            var v39 float64 = 3.14159
            _ = v39
        } else {
            var v40 int16 = 0
            var v41 uint32 = v28
            _ = v40
            _ = v41
        }
        _ = v37
        _ = v38
    }
    var v42 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v42
    _ = v5
}
