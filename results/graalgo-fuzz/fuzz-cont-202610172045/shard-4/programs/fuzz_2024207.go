package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8; F1 uint; F2 int8 }
func genFn0(p0 uint8) int8 {
    fmt.Println(p0)
    var v4 int64 = 202
    v5 := GenS1{F0: 2147483647}
    v4 -= v4
    _ = p0
    _ = v4
    _ = v5
    return int8(v4)
}

func main() {
    var v1 uint8 = 199
    var v2 int32 = 1
    var v3 int32 = v2
    v2 = v3
    v3 -= v3
    v6 := GenS2{F0: (v2 + (-(v2))), F1: 941}
    if true {
        v7 := GenS3{F0: 42, F1: 2}
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        fmt.Println(v3, v2)
        if ((128 <= 128) && true) {
            for fi9 := 0; fi9 < 14; fi9++ {
                v10 := []int{(-(fi8)), fi9}
                var v11 uint64 = 301
                _ = fi9
                _ = v10
                _ = v11
            }
            v3 = (v2 & (^(v3)))
            for fi12 := 0; fi12 < 17; fi12++ {
                v3 = (v2 | (v2 ^ v3))
                v2 += (v2 & (v3 * v2))
                switch ("hello" + ("foo bar" + ("" + "foo bar"))) {
                case "hello":
                    if true {
                        v2 -= v3
                        v13 := []int{}
                        v14 := GenS4{F0: -1, F1: 910, F2: (v1 ^ 42)}
                        _ = v13
                        _ = v14
                    } else {
                        fmt.Println(fi12, fi12, v3)
                    }
                    v15 := []int{(fi8 ^ fi12)}
                    _ = v15
                case "a":
                    v3 = (v2 >> ((255) & 63))
                case "":
                    var v16 int8 = genFn0((v1 - v1))
                    _ = v16
                default:
                    v17 := []int{32767, (fi12 | 255), (fi12 + -633)}
                    _ = v17
                }
                v18 := []int{(fi12 % (((^(fi12))) | 1)), 32768}
                v19 := GenS4{F0: -32768, F1: 42, F2: 1}
                _ = fi12
                _ = v18
                _ = v19
            }
        }
        _ = fi8
    }
    var v20 int64 = -1
    if ("hello" == ("a" + "a")) {
        var v21 int32 = -1
        for fi22 := 0; fi22 < 14; fi22++ {
            if (true || (true && false)) {
                v23 := GenS2{F0: (v21 << ((708) & 63)), F1: 0}
                var v24 int32 = (v2 | v21)
                v3 += (v24 >> ((2) & 63))
                _ = v23
                _ = v24
            } else {
                switch ((((("\n\t" + "hello") + ("hello" + "hello")) + "a") + "") + "a") {
                case "\n\t":
                    var v25 uint = 4294967295
                    _ = v25
                case "foo bar":
                    fmt.Println(v1)
                }
                for fi26 := 0; fi26 < 9; fi26++ {
                    var v27 uint8 = (v1 >> ((456) & 63))
                    fmt.Println(v1, v20, v3)
                    var v28 uint = 255
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                v29 := []int{(fi22 - 42), 42}
                _ = v29
            }
            for fi30 := 0; fi30 < 7; fi30++ {
                v31 := GenS4{F0: -32768, F1: 42, F2: v1}
                v32 := GenS5{F0: 2}
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi22
        }
        if (-2.25 <= 1.5) {
            for fi33 := 0; fi33 < 13; fi33++ {
                fmt.Println(v2, fi33, v3)
                v21 += (^(v21))
                _ = fi33
            }
            var v34 int32 = v2
            v34 += v34
            _ = v34
        } else {
            var v35 string = ("" + "foo bar")
            _ = v35
        }
        _ = v21
    } else {
        v36 := GenS6{F0: 5, F1: 42, F2: 74}
        for fi37 := 0; fi37 < 12; fi37++ {
            v1 = 255
            _ = fi37
        }
        _ = v36
    }
    v1 -= v1
    var v38 int16 = 1000
    v3 = 2147483647
    var v39 uint = 48
    if (false && false) {
        v40 := GenS4{F0: (-(v38)), F1: 60, F2: (v1 * v1)}
        _ = v40
    } else {
        switch 32767 {
        case -1:
            var v41 uint64 = 708
            _ = v41
        default:
            for fi42 := 0; fi42 < 19; fi42++ {
                var v43 uint32 = 4294967295
                switch "foo bar" {
                case "foo bar":
                    var v44 bool = false
                    _ = v44
                }
                _ = fi42
                _ = v43
            }
        }
        var v45 float32 = -2.25
        _ = v45
    }
    v1 += v1
    v46 := []int{}
    fmt.Println(v2, v3)
    _ = len(v46)
    v46 = append(v46, 955)
    var v47 int16 = (v38 | v38)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v38
    _ = v39
    _ = v46
    _ = v47
    _ = v6
}
