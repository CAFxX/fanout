package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 string; F2 uint }
func genFn0(p0 uint16, p1 string, p2 uint8) uint32 {
    var v4 uint8 = 0
    fmt.Println(p2, p1)
    switch (p1 + "") {
    case "":
        for fi5 := 0; fi5 < 6; fi5++ {
            var v6 uint8 = v4
            if true {
                switch (fi5 - (fi5 * (fi5 % ((fi5) | 1)))) {
                case 128:
                    var v7 int32 = -1
                    _ = v7
                case 399:
                    v8 := []int{(fi5 | (fi5 - fi5)), (fi5 * fi5)}
                    _ = v8
                case 4294967295:
                    var v9 uint8 = (p2 * (v6 - (v4 + 127)))
                    _ = v9
                default:
                    var v10 uint32 = uint32(p0)
                    _ = v10
                }
                fmt.Println(p0, p1)
            }
            _ = fi5
            _ = v6
        }
    }
    if ((false || ("" > ("\n\t" + (p1 + p1)))) || (false && (735 != 162))) {
        var v11 int = -2147483648
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := GenS1{F0: (!(-9223372036854775808 > -301)), F1: fi12}
        v4 -= (p2 ^ (v4 ^ 1))
        if (!(true && true)) {
            v14 := GenS1{F0: (-800 == 1), F1: fi12}
            var v15 uint32 = 42
            _ = v14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}
func genFn1(p0 int64) uint64 {
    if false {
        v16 := GenS2{F0: false}
        _ = v16
    } else {
        if ((false && (!false)) && true) {
            for fi17 := 0; fi17 < 12; fi17++ {
                p0 = p0
                p0 += (^(p0))
                p0 += (p0 ^ p0)
                _ = fi17
            }
        } else {
            if (127 >= 100) {
                for fi18 := 0; fi18 < 11; fi18++ {
                    p0 -= (^(p0))
                    _ = fi18
                }
                v19 := []int{}
                p0 = (p0 / (((^(p0))) | 1))
                _ = v19
            } else {
                var v20 uint8 = 1
                v21 := GenS1{F0: (685 >= 920), F1: 128}
                _ = v20
                _ = v21
            }
        }
        var v22 int64 = 0
        _ = v22
    }
    fmt.Println(p0)
    if true {
        for fi23 := 0; fi23 < 12; fi23++ {
            v24 := GenS1{F0: false, F1: 42}
            _ = fi23
            _ = v24
        }
        fmt.Println(p0)
    } else {
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 int16 = -544
            for fi27 := 0; fi27 < 6; fi27++ {
                if (2 != 0) {
                    var v28 float64 = ((-2.25 - 1e10) - 3.14159)
                    _ = v28
                }
                v29 := []int{}
                _ = fi27
                _ = v29
            }
            _ = fi25
            _ = v26
        }
        fmt.Println(p0)
    }
    _ = p0
    return 4294967295
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 65535
    var v30 int16 = 1000
    v31 := GenS3{F0: 0, F1: ("" + "hello"), F2: 209}
    var v32 uint64 = genFn1(229)
    v33 := GenS2{F0: true}
    for fi34 := 0; fi34 < 9; fi34++ {
        var v35 int16 = v30
        _ = fi34
        _ = v35
    }
    v32 = (v32 >> ((255) & 63))
    if (!false) {
        v30 += (v30 + (v30 << ((326) & 63)))
        v3 = (v3 - 745)
    } else {
        v30 = (v30 << ((42) & 63))
        if false {
            v3 = (v3 >> ((255) & 63))
        }
    }
    fmt.Println(v32, v3)
    v36 := GenS2{F0: true}
    if false {
        if (1 != 768) {
            fmt.Println(v3, v32)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v36
}
