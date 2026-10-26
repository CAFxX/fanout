package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 uint }
func genFn0() uint8 {
    var v5 uint64 = 356
    v5 -= 2
    switch "hello" {
    case "":
        v6 := []int{-124, 42, 999}
        _ = v6
    case "\n\t":
        var v7 uint64 = v5
        _ = v7
    case "a":
        v8 := GenS1{F0: 255}
        _ = v8
    default:
        switch 9223372036854775807 {
        case -9223372036854775808:
            if true {
                fmt.Println(v5)
                fmt.Println(v5)
            }
        }
    }
    var v9 string = "hello"
    _ = v5
    _ = v9
    return 110
}
func genFn1() uint16 {
    switch 1 {
    case -1:
        var v10 int8 = 59
        _ = v10
    case -454:
        var v11 int8 = 1
        _ = v11
    }
    if true {
        var v12 uint64 = 0
        for fi13 := 0; fi13 < 18; fi13++ {
            switch (fi13 + (fi13 << ((18446744073709551615) & 63))) {
            case -946:
                for fi14 := 0; fi14 < 2; fi14++ {
                    switch int(v12) {
                    case 942:
                        if true {
                            fmt.Println(fi13, fi14, fi14)
                            v15 := []int{(fi13 % ((1) | 1)), 0, (fi13 % ((fi14) | 1))}
                            v16 := GenS2{F0: 100, F1: true}
                            _ = v15
                            _ = v16
                        }
                    case 42:
                        var v17 float32 = -2.25
                        _ = v17
                    case 1:
                        switch fi14 {
                        case 4294967295:
                            var v18 bool = true
                            _ = v18
                        case -192:
                            v12 = v12
                            v12 += 8
                        default:
                            var v19 int32 = 2147483647
                            _ = v19
                        }
                    default:
                        v20 := []int{32768, -922}
                        _ = v20
                    }
                    v21 := []int{fi14, fi14, fi13}
                    switch "hello" {
                    case "foo bar":
                        v22 := GenS1{F0: genFn0()}
                        _ = v22
                    default:
                        v12 += 0
                    }
                    fmt.Println(v12, v12)
                    _ = fi14
                    _ = v21
                }
            case 32768:
                v23 := []int{}
                _ = v23
            case -9223372036854775808:
                v24 := GenS3{F0: "foo bar", F1: 42}
                _ = v24
            }
            v12 = 42
            fmt.Println(v12)
            _ = fi13
        }
        fmt.Println(v12)
        _ = v12
    } else {
        var v25 uint32 = 309
        _ = v25
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        fmt.Println(fi26)
        var v27 uint64 = 18446744073709551615
        _ = fi26
        _ = v27
    }
    var v28 int = -1
    _ = v28
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 0
    var v3 uint16 = 42
    v4 := []int{9223372036854775807}
    for fi29 := 0; fi29 < 13; fi29++ {
        v2 = int8(fi29)
        _ = fi29
    }
    fmt.Println(v3)
    for fi30 := 0; fi30 < 19; fi30++ {
        for fi31 := 0; fi31 < 20; fi31++ {
            var v32 string = ("" + "hello")
            _ = len(v4)
            _ = fi31
            _ = v32
        }
        v3 += v3
        switch fi30 {
        case 785:
            for fi33 := 0; fi33 < 16; fi33++ {
                if (!false) {
                    v3 = v3
                    var v34 int32 = -294
                    switch "" {
                    case "":
                        v3 -= v3
                    case "hello":
                        var v35 int8 = v2
                        _ = v35
                    }
                    var v36 int = fi30
                    _ = v34
                    _ = v36
                }
                if ("\n\t" > ("hello" + "a")) {
                    _ = len(v4)
                    fmt.Println(v2)
                    var v37 uint64 = 4294967295
                    _ = v37
                }
                _ = fi33
            }
        }
        _ = fi30
    }
    var v38 int16 = 0
    fmt.Println(v38)
    for fi39 := 0; fi39 < 4; fi39++ {
        v40 := GenS2{F0: (v2 ^ v2), F1: true}
        v1["hello"] = ((("hello" + (("\n\t" + ("hello" + "\n\t")) + "")) + "") + ((("\n\t" + ("" + ("foo bar" + "a"))) + "a") + "foo bar"))
        var v41 int64 = 0
        _ = fi39
        _ = v40
        _ = v41
    }
    fmt.Println(v2, v38)
    v42 := GenS4{F0: v2}
    v43 := GenS2{F0: (v2 / (((v2 ^ v2)) | 1)), F1: true}
    fmt.Println(v2, v2)
    var v44 float64 = 1e10
    if false {
        var v45 uint64 = 71
        var v46 int16 = (v38 ^ (-(v38)))
        _ = v45
        _ = v46
    } else {
        for fi47 := 0; fi47 < 20; fi47++ {
            for fi48 := 0; fi48 < 17; fi48++ {
                var v49 uint8 = 29
                _ = fi48
                _ = v49
            }
            _ = fi47
        }
        if len(v4) > 0 {
            v4[0] = 32768
        }
    }
    v38 = v38
    v50 := GenS5{F0: (false || (2 == 4294967295)), F1: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v50
}
