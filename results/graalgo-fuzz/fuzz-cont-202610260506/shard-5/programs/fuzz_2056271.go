package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint32 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint }
type GenS5 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint32, p1 uint) uint16 {
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := []int{fi4, fi4}
        _ = fi4
        _ = v5
    }
    p0 += 26
    p0 -= (p0 * (p0 ^ p0))
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = false
    v3 := make(map[int]int)
    v6 := GenS1{F0: 255, F1: 128, F2: 1}
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v7 uint16 = 33
        _ = v7
    case -322:
        delete(v3, 746)
    default:
        v8 := GenS2{F0: 42, F1: -377}
        _ = v8
    }
    v9 := []int{655, 42, -2147483648, -9223372036854775808}
    v9 = append(v9, 602)
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(v2, v2)
        _ = fi10
    }
    var v11 int8 = -128
    v12 := GenS1{F0: 2, F1: -892, F2: 719}
    fmt.Println(v11, v2)
    if false {
        if true {
            fmt.Println(v2, v11)
            v11 += (v11 + v11)
        } else {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
    } else {
        fmt.Println(v11, v11)
    }
    var v13 uint16 = 65535
    switch 32767 {
    case -1:
        switch (("" + "foo bar") + "hello") {
        case "foo bar":
            var v14 uint8 = 238
            _ = v14
        case "hello":
            v2 = v2
        }
    case 42:
        var v15 string = "\n\t"
        _ = v15
    case 32768:
        for fi16 := 0; fi16 < 13; fi16++ {
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 int8 = 111
                v19 := v1[fi16]
                if v2 {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    v18 = 4
                    v20 := GenS1{F0: v13, F1: fi16, F2: v19}
                    _ = v20
                } else {
                    var v21 string = "a"
                    var v22 uint16 = v13
                    _ = v21
                    _ = v22
                }
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = fi16
        }
    default:
        for fi23 := 0; fi23 < 13; fi23++ {
            var v24 int16 = -1
            var v25 int8 = (v11 - (v11 >> ((2) & 63)))
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    var v26 uint16 = (v13 + v13)
    fmt.Println(v26)
    fmt.Println(v26, v11, v13)
    v26 += v26
    for fi27 := 0; fi27 < 9; fi27++ {
        if v2 {
            v13 -= (^(v13))
        } else {
            var v28 uint64 = 2
            switch (fi27 * int(v13)) {
            case -2147483648:
                var v29 uint8 = 255
                _ = v29
            default:
                var v30 bool = v2
                _ = v30
            }
            v31 := GenS3{F0: fi27}
            _ = v28
            _ = v31
        }
        v32 := GenS4{F0: 0, F1: 1, F2: 65535}
        _ = fi27
        _ = v32
    }
    if v2 {
        v33 := GenS5{F0: (180 != 239), F1: 128}
        for fi34 := 0; fi34 < 2; fi34++ {
            for fi35 := 0; fi35 < 18; fi35++ {
                v36 := GenS1{F0: genFn0(1, 0), F1: fi35, F2: 0}
                _ = fi35
                _ = v36
            }
            var v37 float32 = -2.25
            v38 := v3[-637]
            _ = fi34
            _ = v37
            _ = v38
        }
        _ = v33
    } else {
        var v39 uint32 = 2147483647
        v13 -= (v13 * v26)
        _ = v39
    }
    v26 = v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v6
    _ = v9
}
