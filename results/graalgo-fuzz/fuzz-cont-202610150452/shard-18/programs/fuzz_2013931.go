package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 int16; F2 bool }
func genFn0() uint16 {
    var v4 string = ("hello" + ("a" + "\n\t"))
    if true {
        fmt.Println(v4)
    }
    v5 := []int{}
    var v6 string = ""
    _ = v4
    _ = v5
    _ = v6
    return 32767
}
func genFn1(p0 bool) bool {
    switch "" {
    case "\n\t":
        p0 = p0
    case "hello":
        var v7 bool = (42 > 1)
        _ = v7
    }
    if (546 == 960) {
        var v8 uint8 = 255
        v8 -= 43
        var v9 int64 = 1
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 8; fi10++ {
            if p0 {
                v11 := GenS1{F0: 2147483648, F1: 1}
                p0 = (!false)
                v12 := []int{(fi10 ^ (fi10 ^ (^(fi10)))), (^(fi10))}
                _ = v11
                _ = v12
            }
            _ = fi10
        }
    }
    _ = p0
    return (p0 && p0)
}

func main() {
    var v1 uint16 = 255
    var v2 string = ("" + "\n\t")
    v3 := make(map[uint32]string)
    fmt.Println(v1, v1)
    for fi13 := 0; fi13 < 7; fi13++ {
        if (-749 == 298) {
            var v14 uint64 = 765
            _ = v14
        }
        v1 = (v1 | 0)
        var v15 int16 = 1
        _ = fi13
        _ = v15
    }
    v16 := []int{32767, 32767}
    for fi17 := 0; fi17 < 6; fi17++ {
        switch (fi17 * (^(fi17))) {
        case -1:
            for fi18 := 0; fi18 < 11; fi18++ {
                v19 := GenS2{F0: uint32(v1)}
                _ = fi18
                _ = v19
            }
        case 4294967295:
            var v20 uint = 14
            _ = v20
        default:
            switch (fi17 ^ fi17) {
            case 4294967295:
                if ((!(("\n\t" + "") == v2)) && (-9223372036854775808 <= 42)) {
                    for fi21 := 0; fi21 < 6; fi21++ {
                        v2 = "hello"
                        _ = fi21
                    }
                }
            case 9223372036854775807:
                v1 += v1
            case 1:
                var v22 uint8 = 0
                _ = v22
            default:
                v2 = (v2 + (("hello" + ("hello" + (v2 + v2))) + (v2 + v2)))
            }
        }
        v1 = v1
        switch (fi17 % (((fi17 ^ fi17)) | 1)) {
        case 4294967295:
            var v23 int32 = 2147483647
            _ = v23
        }
        _ = fi17
    }
    delete(v3, 742)
    switch "\n\t" {
    case "hello":
        v16 = append(v16, -661)
    case "":
        var v24 int = 492
        _ = v24
    case "foo bar":
        v25 := GenS3{F0: 798, F1: 1000, F2: true}
        _ = v25
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        v2 = (v2 + v2)
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 string = ("hello" + "a")
            if (!(0 > (^(fi26)))) {
                switch -308 {
                case 1:
                    var v29 int8 = 127
                    _ = v29
                case 42:
                    fmt.Println(v1, fi27, v2)
                }
            }
            v1 = v1
            _ = fi27
            _ = v28
        }
        switch ("hello" + (v2 + v2)) {
        case "\n\t":
            v30 := GenS3{F0: 42, F1: -1, F2: ((v1 ^ (v1 & (v1 | (v1 * genFn0())))) >= 255)}
            _ = v30
        case "hello":
            var v31 float32 = (1e10 * 0.0)
            _ = v31
        case "foo bar":
            v16 = append(v16, fi26)
        }
        _ = fi26
    }
    var v32 string = v2
    fmt.Println(v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v32
}
