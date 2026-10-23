package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint16, p1 int16, p2 string) uint32 {
    switch "\n\t" {
    case "a":
        p1 = p1
    case "":
        p1 += 492
    }
    v5 := []int{9223372036854775807, -1}
    var v6 float64 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}
func genFn1() string {
    v7 := GenS1{F0: "hello", F1: 128}
    v8 := GenS2{F0: 371}
    _ = v7
    _ = v8
    return ("" + "\n\t")
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    var v3 int32 = -1
    var v4 int64 = 1
    fmt.Println(v4, v4)
    for fi9 := 0; fi9 < 6; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 int16 = 32767
            fmt.Println(v3, fi10)
            v4 = (v4 & 9223372036854775807)
            _ = fi10
            _ = v11
        }
        v3 += (^(v3))
        var v12 uint32 = 272
        _ = fi9
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        switch "\n\t" {
        case "hello":
            for fi14 := 0; fi14 < 12; fi14++ {
                v15 := GenS3{F0: -266, F1: 65535, F2: 81}
                _ = fi14
                _ = v15
            }
        case "":
            var v16 int8 = -1
            _ = v16
        case "a":
            var v17 bool = false
            _ = v17
        default:
            var v18 uint32 = 1
            _ = v18
        }
        v19 := v1[fi13]
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 int16 = -89
            switch ("\n\t" + "") {
            case "":
                v22 := GenS2{F0: 2147483648}
                _ = v22
            case "a":
                switch fi13 {
                case -681:
                    v1[1] = true
                }
                _ = len(v1)
            }
            v23 := GenS4{F0: (v4 | v4)}
            _ = fi20
            _ = v21
            _ = v23
        }
        _ = fi13
        _ = v19
    }
    switch 0 {
    case 97:
        switch -2147483648 {
        case 1:
            var v24 int64 = v4
            _ = v24
        case 128:
            v3 += (v3 / ((v3) | 1))
        case -9223372036854775808:
            for fi25 := 0; fi25 < 16; fi25++ {
                var v26 int16 = -32768
                v3 += v3
                _ = fi25
                _ = v26
            }
        default:
            var v27 float32 = -2.25
            _ = v27
        }
    case -665:
        delete(v1, 42)
    }
    var v28 int32 = (v3 & (v3 ^ 125))
    fmt.Println(v28)
    var v29 uint32 = 2147483647
    if true {
        v30 := []int{-444, int(v29), 0}
        if (true && false) {
            for fi31 := 0; fi31 < 4; fi31++ {
                var v32 uint8 = 163
                _ = fi31
                _ = v32
            }
            var v33 uint32 = v29
            _ = v33
        } else {
            if ((!(v29 > genFn0(255, 32767, ""))) || ((!(255 != 127)) || false)) {
                var v34 uint8 = 1
                v34 -= (v34 << ((255) & 63))
                v35 := GenS3{F0: 0, F1: v29, F2: v34}
                _ = v34
                _ = v35
            }
            var v36 uint64 = 2
            _ = v36
        }
        _ = v30
    } else {
        v37 := []int{}
        var v38 int16 = -619
        v4 -= (v4 << ((1) & 63))
        _ = v37
        _ = v38
    }
    v39 := []int{103, 4294967295, -565}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v39
    _ = v4
}
