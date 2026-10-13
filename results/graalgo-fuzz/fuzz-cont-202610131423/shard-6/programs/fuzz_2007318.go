package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 uint; F1 int; F2 uint32 }
func genFn0(p0 int8) int8 {
    if (-518 >= 300) {
        for fi5 := 0; fi5 < 20; fi5++ {
            if true {
                var v6 float64 = 0.1
                _ = v6
            }
            _ = fi5
        }
        v7 := GenS1{F0: 777, F1: 2}
        p0 -= p0
        _ = v7
    } else {
        v8 := []int{}
        _ = v8
    }
    v9 := GenS1{F0: int(p0), F1: 65535}
    v10 := GenS2{F0: "foo bar", F1: 42}
    var v11 int64 = -1
    v12 := []int{}
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 56
}
func genFn1(p0 uint, p1 uint) string {
    v13 := []int{-9223372036854775808, 224}
    if (true && (!(!true))) {
        p0 -= 946
        p0 -= 0
    }
    switch "a" {
    case "hello":
        var v14 uint32 = 0
        _ = v14
    case "a":
        v15 := GenS3{F0: 2147483648, F1: 0, F2: 2}
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        if len(v13) > 0 {
            v13[0] = (fi16 & 671)
        }
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = v13
    return ((("" + "") + ("" + "foo bar")) + "")
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 128
    var v3 int16 = 0
    v4 := make(map[int]bool)
    switch (v2 - v2) {
    case 127:
        var v17 uint = 397
        _ = v17
    case 181:
        v2 += 42
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        v2 -= (fi18 / (((fi18 * v2)) | 1))
        for fi19 := 0; fi19 < 4; fi19++ {
            fmt.Println(v2, v3, fi19)
            var v20 float64 = -2.25
            var v21 int8 = int8(v3)
            var v22 int32 = 42
            v2 = -2147483648
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi18
    }
    v3 = (v3 ^ 42)
    v23 := GenS3{F0: 221, F1: 121, F2: 951}
    fmt.Println(v3, v2)
    for fi24 := 0; fi24 < 18; fi24++ {
        fmt.Println(fi24, fi24, v2)
        if (1 > 0) {
            v25 := GenS4{F0: 335, F1: true}
            _ = v25
        }
        _ = fi24
    }
    v2 = int(v3)
    fmt.Println(v2)
    if ((2 >= 840) && (!(v2 > 9223372036854775807))) {
        for fi26 := 0; fi26 < 16; fi26++ {
            if false {
                v3 += v3
                v3 += (^(v3))
                fmt.Println(v2, fi26)
            } else {
                fmt.Println(v3, v2, fi26)
                delete(v1, 4294967295)
                switch "\n\t" {
                case "":
                    var v27 int16 = (^(v3))
                    _ = v27
                }
            }
            v28 := GenS1{F0: v2, F1: 0}
            _ = fi26
            _ = v28
        }
        v2 = v2
        var v29 float64 = 0.0
        _ = v29
    } else {
        var v30 uint8 = 42
        _ = v30
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        v32 := []int{284}
        var v33 float64 = -2.25
        switch fi31 {
        case 2147483647:
            v3 = 1
        case -435:
            if len(v32) >= 2 {
                v32 = v32[1:len(v32)]
            }
        case -9223372036854775808:
            fmt.Println(v33, v2, v2)
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println(v2)
    var v34 int16 = v3
    _ = len(v4)
    for fi35 := 0; fi35 < 14; fi35++ {
        var v36 int8 = 21
        for fi37 := 0; fi37 < 13; fi37++ {
            var v38 float32 = float32(((1e10 + float64(-2.25)) + float64(1e10)))
            fmt.Println(fi37, v2)
            v39 := GenS5{F0: 963, F1: fi37, F2: 822}
            _ = fi37
            _ = v38
            _ = v39
        }
        var v40 float64 = ((-2.25 * 0.1) + float64(0.0))
        _ = fi35
        _ = v36
        _ = v40
    }
    for fi41 := 0; fi41 < 19; fi41++ {
        if ((-9223372036854775808 < 9223372036854775807) || (1 == 2)) {
            var v42 float32 = 0.0
            var v43 uint = 939
            if true {
                if false {
                    v3 = (v3 >> (((v43 & (v43 / (((v43 - (v43 & (v43 >> (((v43 | v43)) & 63))))) | 1)))) & 63))
                    var v44 uint = (v43 * (v43 * 255))
                    _ = v44
                }
            }
            _ = v42
            _ = v43
        }
        v45 := GenS1{F0: -129, F1: 198}
        fmt.Println(v3, v3)
        _ = fi41
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v34
    _ = v4
}
