package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
func genFn0(p0 int16, p1 uint8) int {
    p0 = (p0 & 726)
    p0 += p0
    fmt.Println(p0, p1)
    v4 := []int{810, 9223372036854775807, 127, 127}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v4
    return -1
}
func genFn1(p0 int16, p1 int64) uint16 {
    v5 := GenS1{F0: genFn0((p0 ^ p0), 42), F1: (false || false)}
    fmt.Println(p1, p0)
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int64 = p1
        var v8 float64 = 3.14159
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: 1, F1: 10, F2: 771}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return 7
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]int)
    var v3 float64 = -2.25
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := []int{}
            if false {
                var v13 int64 = -9223372036854775808
                _ = v13
            }
            _ = fi11
            _ = v12
        }
        var v14 bool = (111 > 42)
        _ = fi10
        _ = v14
    }
    v15 := GenS2{F0: 255, F1: 692, F2: 300}
    var v16 uint32 = 2147483648
    for fi17 := 0; fi17 < 3; fi17++ {
        v18 := GenS1{F0: (fi17 - (fi17 % ((fi17) | 1))), F1: false}
        v19 := GenS3{F0: 42}
        v20 := GenS3{F0: int8(v16)}
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        v22 := GenS3{F0: 42}
        var v23 int = 61
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        for fi25 := 0; fi25 < 1; fi25++ {
            switch (fi25 / ((fi24) | 1)) {
            case 9223372036854775807:
                var v26 uint32 = v16
                _ = v26
            case 4294967295:
                switch genFn0(453, 1) {
                case -9223372036854775808:
                    v16 += (^(v16))
                case 9223372036854775807:
                    v27 := GenS2{F0: 0, F1: 2, F2: 592}
                    _ = v27
                case 1:
                    var v28 int = (fi24 | genFn0(844, 128))
                    _ = v28
                default:
                    var v29 float32 = 0.1
                    _ = v29
                }
                var v30 int16 = 1000
                v31 := []int{fi24, fi25, (fi24 ^ 1), 0}
                _ = v30
                _ = v31
            case 1:
                var v32 uint64 = 209
                _ = v32
            }
            switch v1 {
            case "a":
                v33 := v2[v16]
                _ = v33
            }
            _ = fi25
        }
        if (true && true) {
            v3 -= 0.0
            if (530 == 255) {
                var v34 int8 = -69
                v35 := v2[v16]
                _ = v34
                _ = v35
            } else {
                var v36 uint = 65535
                for fi37 := 0; fi37 < 10; fi37++ {
                    v16 -= v16
                    var v38 uint16 = 42
                    _ = fi37
                    _ = v38
                }
                var v39 float32 = ((1.5 - 1e10) / 1.0)
                _ = v36
                _ = v39
            }
        }
        _ = fi24
    }
    if (((0.1 * (v3 * float64(((-2.25 * 0.1) - (1e10 * float32(3.14159)))))) * 0.1) < (0.1 * (v3 + v3))) {
        v1 = ("a" + "\n\t")
    } else {
        _ = len(v2)
        var v40 bool = (-554 > 32767)
        fmt.Println(v16, v40, v16)
        _ = v40
    }
    _ = len(v2)
    var v41 uint = 4294967295
    v42 := GenS2{F0: 93, F1: 355, F2: 2}
    if ((v1 + (v1 + v1)) == (v1 + (v1 + v1))) {
        var v43 string = v1
        _ = v43
    } else {
        v44 := GenS3{F0: -86}
        for fi45 := 0; fi45 < 20; fi45++ {
            var v46 string = ("a" + v1)
            _ = fi45
            _ = v46
        }
        v16 = (v16 % (((^(v16))) | 1))
        _ = v44
    }
    for fi47 := 0; fi47 < 18; fi47++ {
        fmt.Println(v3, v41, v41)
        if false {
            v48 := GenS4{F0: fi47}
            _ = v48
        }
        _ = fi47
    }
    switch 1 {
    case 32767:
        v41 = (v41 + (^(v41)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v41
    _ = v42
}
