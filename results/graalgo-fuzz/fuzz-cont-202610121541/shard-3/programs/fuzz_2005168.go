package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint64 }
type GenS2 struct { F0 bool; F1 int64; F2 int }
type GenS3 struct { F0 int64; F1 int16; F2 uint16 }
func genFn0(p0 uint64, p1 uint, p2 int) uint64 {
    p0 -= (^(p0))
    v4 := []int{p2, p2}
    fmt.Println(p1)
    var v5 int16 = 32767
    switch "foo bar" {
    case "":
        for fi6 := 0; fi6 < 1; fi6++ {
            p0 = (p0 ^ (p0 >> (((p1 - (p1 % ((255) | 1)))) & 63)))
            _ = fi6
        }
    case "a":
        _ = len(v4)
    case "foo bar":
        v7 := GenS1{F0: p0, F1: (true || true), F2: 520}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p0
}
func genFn1(p0 uint16) int32 {
    p0 = p0
    if false {
        if true {
            v8 := GenS2{F0: false, F1: -1, F2: 128}
            v9 := GenS1{F0: 42, F1: (0 == 42), F2: 2}
            _ = v8
            _ = v9
        } else {
            p0 += p0
            v10 := []int{-1, 387, -459, 214}
            var v11 int16 = 975
            _ = v10
            _ = v11
        }
    }
    if true {
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := GenS1{F0: 0, F1: (1 < 2147483648), F2: 42}
            p0 = (p0 & p0)
            _ = fi12
            _ = v13
        }
        p0 += 42
    }
    if true {
        if true {
            var v14 uint16 = p0
            if true {
                v15 := GenS1{F0: genFn0(2, 798, 9223372036854775807), F1: (false && false), F2: 928}
                for fi16 := 0; fi16 < 16; fi16++ {
                    var v17 int = -2147483648
                    v18 := GenS2{F0: ((true || false) || (false && false)), F1: 1, F2: (^(v17))}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                var v19 int8 = -6
                _ = v15
                _ = v19
            } else {
                for fi20 := 0; fi20 < 4; fi20++ {
                    var v21 float64 = (((3.14159 * 1e10) / 0.1) - (float64(((1.5 / -2.25) * -2.25)) / 1.0))
                    _ = fi20
                    _ = v21
                }
                for fi22 := 0; fi22 < 15; fi22++ {
                    v14 += (v14 ^ 0)
                    v23 := []int{fi22, fi22, fi22}
                    p0 = p0
                    _ = fi22
                    _ = v23
                }
                var v24 int16 = 32767
                _ = v24
            }
            _ = v14
        }
        for fi25 := 0; fi25 < 8; fi25++ {
            switch ("foo bar" + "") {
            case "hello":
                var v26 uint64 = 131
                _ = v26
            case "a":
                for fi27 := 0; fi27 < 8; fi27++ {
                    v28 := GenS3{F0: -984, F1: 42, F2: p0}
                    v29 := []int{}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            case "":
                if true {
                    p0 -= 42
                }
            }
            if (genFn0(927, 2, (-(fi25))) == genFn0(genFn0(4294967295, 125, (fi25 + fi25)), 1, 32767)) {
                p0 = p0
            } else {
                for fi30 := 0; fi30 < 16; fi30++ {
                    v31 := GenS1{F0: 162, F1: true, F2: 42}
                    _ = fi30
                    _ = v31
                }
                fmt.Println(p0, fi25)
            }
            _ = fi25
        }
        switch "a" {
        case "hello":
            p0 = (p0 >> ((537) & 63))
        case "foo bar":
            p0 -= p0
        case "a":
            fmt.Println(p0)
        }
    } else {
        if (94 > -128) {
            var v32 int = -708
            v33 := GenS1{F0: 42, F1: ((100 == 127) || false), F2: genFn0(0, 190, (v32 * v32))}
            _ = v32
            _ = v33
        } else {
            var v34 uint16 = p0
            var v35 int = -1
            v34 += (v34 & (p0 ^ v34))
            _ = v34
            _ = v35
        }
    }
    if false {
        for fi36 := 0; fi36 < 18; fi36++ {
            p0 -= (^(p0))
            fmt.Println(p0, fi36)
            _ = fi36
        }
        v37 := []int{-998, 4294967295, 4294967295, 128}
        _ = v37
    } else {
        v38 := GenS1{F0: 793, F1: ((p0 / ((p0) | 1)) < (p0 - (p0 ^ 476))), F2: 1}
        var v39 int = 127
        p0 += p0
        _ = v38
        _ = v39
    }
    _ = p0
    return 2147483647
}

func main() {
    var v1 int16 = 42
    var v2 float64 = 0.1
    v3 := []string{"\n\t", "hello"}
    for fi40 := 0; fi40 < 19; fi40++ {
        switch (fi40 & 32767) {
        case 42:
            fmt.Println(v1, fi40)
        case 2147483647:
            if false {
                v41 := GenS2{F0: true, F1: -48, F2: 32767}
                _ = v41
            }
        }
        for fi42 := 0; fi42 < 5; fi42++ {
            var v43 int64 = 1
            _ = fi42
            _ = v43
        }
        _ = fi40
    }
    var v44 int16 = v1
    v45 := GenS2{F0: true, F1: 0, F2: 1}
    for fi46 := 0; fi46 < 13; fi46++ {
        fmt.Println(v2, v1, fi46)
        _ = fi46
    }
    for fi47 := 0; fi47 < 4; fi47++ {
        v3 = append(v3, "hello")
        v1 += 0
        _ = fi47
    }
    v44 = 32767
    var v48 int = 0
    fmt.Println(v48, v44, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v48
}
