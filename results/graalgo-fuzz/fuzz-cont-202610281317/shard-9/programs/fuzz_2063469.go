package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint32) int8 {
    var v5 int32 = 386
    if (0 == 657) {
        v6 := GenS1{F0: (2 == 961), F1: "hello"}
        _ = v6
    } else {
        var v7 int = -268
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{(fi8 << ((17) & 63)), fi8, (fi8 & 32768), 128}
        p0 = 65535
        _ = fi8
        _ = v9
    }
    if (v5 == (v5 - (^(v5)))) {
        if (45 <= 255) {
            p0 = p0
            fmt.Println(v5)
        } else {
            var v10 int8 = -128
            v11 := []int{42, 836, 32767}
            var v12 int32 = (v5 % (((v5 / ((-775) | 1))) | 1))
            _ = v10
            _ = v11
            _ = v12
        }
        p0 = 4294967295
        v5 = 2147483647
    } else {
        p0 -= (p0 + 902)
        var v13 uint16 = 836
        var v14 bool = ((!true) || (!false))
        _ = v13
        _ = v14
    }
    var v15 int = -2147483648
    _ = p0
    _ = v15
    _ = v5
    return -22
}
func genFn1(p0 string, p1 uint16, p2 uint64) uint8 {
    var v16 uint64 = p2
    v17 := GenS1{F0: (1.5 <= 0.1), F1: ("" + p0)}
    if true {
        var v18 int8 = 25
        _ = v18
    } else {
        if true {
            fmt.Println(p0, v16, p2)
        }
        p2 -= 4294967295
        var v19 int32 = 42
        _ = v19
    }
    var v20 int8 = 73
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v20
    return 0
}

func main() {
    v1 := []uint32{255}
    var v2 int8 = 36
    v3 := make(map[string]string)
    v4 := make(map[uint32]int)
    var v21 int = 2147483647
    var v22 uint = 433
    switch (v21 - (v21 | v21)) {
    case 32768:
        for fi23 := 0; fi23 < 4; fi23++ {
            var v24 float64 = 1.5
            _ = fi23
            _ = v24
        }
    case 4294967295:
        var v25 int = v21
        _ = v25
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        if len(v1) > 0 {
            v1[0] = 654
        }
        _ = fi26
    }
    if true {
        var v27 int8 = (^(v2))
        v28 := GenS1{F0: false, F1: "a"}
        _ = v27
        _ = v28
    }
    v29 := GenS2{F0: ((!((("\n\t" + "a") <= ("" + "")) && ((v21 >> ((v22) & 63)) < v21))) || (1 < -1))}
    for fi30 := 0; fi30 < 3; fi30++ {
        if (739 < 629) {
            v31 := GenS1{F0: true, F1: ""}
            v32 := GenS3{F0: 42}
            fmt.Println(fi30)
            _ = v31
            _ = v32
        }
        fmt.Println(v2)
        fmt.Println(v22)
        _ = fi30
    }
    _ = len(v1)
    v33 := GenS4{F0: (v22 * 804)}
    fmt.Println(v2, v2)
    switch "hello" {
    case "hello":
        v21 += (v21 >> ((v22) & 63))
    }
    v4[977] = (v21 << ((4294967295) & 63))
    switch v21 {
    case 254:
        for fi34 := 0; fi34 < 20; fi34++ {
            for fi35 := 0; fi35 < 20; fi35++ {
                v21 += v21
                _ = fi35
            }
            for fi36 := 0; fi36 < 6; fi36++ {
                var v37 uint16 = 2
                _ = fi36
                _ = v37
            }
            v2 += v2
            _ = fi34
        }
    case -521:
        if (false || (!(!("\n\t" != ("hello" + "a"))))) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v2 -= 1
        }
    }
    v21 += v21
    fmt.Println(v2, v21)
    for fi38 := 0; fi38 < 1; fi38++ {
        v1 = append(v1, 2147483648)
        _ = fi38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v33
    _ = v4
}
