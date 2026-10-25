package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint64, p1 uint16, p2 uint16) int64 {
    for fi5 := 0; fi5 < 13; fi5++ {
        if true {
            p1 = (p1 % (((p2 / (((^(p2))) | 1))) | 1))
            if (-32768 == 28) {
                p1 -= (p1 - (p1 << ((0) & 63)))
                p2 = (p2 % ((977) | 1))
                fmt.Println(p0, p0, p1)
            }
        }
        switch (fi5 / ((-1) | 1)) {
        case 127:
            fmt.Println(p0)
        case 32768:
            for fi6 := 0; fi6 < 7; fi6++ {
                var v7 uint8 = 136
                _ = fi6
                _ = v7
            }
        case -303:
            v8 := []int{(^(fi5))}
            _ = v8
        }
        _ = fi5
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int32 = 0
    v2 := []string{(("" + "a") + ("" + ("foo bar" + "\n\t"))), ("" + ""), "foo bar", ("\n\t" + "a")}
    var v3 bool = (4294967295 <= 2)
    v4 := make(map[string]int)
    switch "foo bar" {
    case "\n\t":
        v9 := GenS1{F0: v1}
        _ = v9
    case "a":
        v1 -= (v1 / (((v1 % ((-101) | 1))) | 1))
    }
    if ((!(v3 || v3)) && ((3.14159 + ((0.0 + (1e10 + (1.5 - 0.1))) + 1.5)) <= (-2.25 / ((-2.25 * ((0.0 * 1.5) + (-2.25 * 3.14159))) + 1e10)))) {
        v3 = true
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        switch "" {
        case "\n\t":
            var v11 int32 = v1
            _ = v11
        case "hello":
            v1 += v1
        }
        _ = fi10
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v12 := GenS1{F0: v1}
    for fi13 := 0; fi13 < 13; fi13++ {
        switch fi13 {
        case 2147483647:
            if ((v3 && (!v3)) || v3) {
                v14 := v4[("\n\t" + ("\n\t" + ("hello" + "\n\t")))]
                if v3 {
                    fmt.Println(v3, v14)
                    if v3 {
                        if len(v2) > 0 {
                            v2[0] = ""
                        }
                    } else {
                        var v15 uint16 = 1
                        _ = v15
                    }
                } else {
                    for fi16 := 0; fi16 < 5; fi16++ {
                        v1 = 0
                        v17 := GenS2{F0: 0, F1: 0, F2: 128}
                        var v18 int16 = 0
                        v1 -= 1
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    if false {
                        if v3 {
                            var v19 bool = true
                            _ = v19
                        } else {
                            switch ("a" + "foo bar") {
                            case "a":
                                fmt.Println(v14, fi13, v1)
                            case "\n\t":
                                if (128 >= 1) {
                                    if len(v2) > 0 {
                                        v2[0] = ("hello" + "hello")
                                    }
                                }
                            default:
                                v14 += v14
                            }
                            if (v3 || v3) {
                                var v20 uint8 = 1
                                v21 := GenS3{F0: 871, F1: 32767}
                                _ = v20
                                _ = v21
                            } else {
                                switch v14 {
                                case 2147483647:
                                    v3 = (true || (v14 > v14))
                                }
                                var v22 uint64 = 1
                                _ = v22
                            }
                        }
                        var v23 int32 = 1
                        _ = len(v2)
                        _ = v23
                    }
                    switch -9223372036854775808 {
                    case -720:
                        fmt.Println(fi13, v14)
                    }
                }
                var v24 float64 = 1e10
                _ = v14
                _ = v24
            }
        case -261:
            switch (fi13 % (((fi13 / ((fi13) | 1))) | 1)) {
            case 4294967295:
                for fi25 := 0; fi25 < 13; fi25++ {
                    var v26 string = ""
                    var v27 float32 = 1e10
                    fmt.Println(fi13, v3)
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            default:
                v1 += (v1 >> ((847) & 63))
            }
        case 1:
            fmt.Println(v3, fi13, fi13)
        }
        _ = fi13
    }
    var v28 uint8 = 42
    var v29 uint16 = 255
    v30 := GenS4{F0: 265}
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 string = ""
        var v33 uint32 = 4294967295
        v34 := GenS4{F0: 255}
        _ = fi31
        _ = v32
        _ = v33
        _ = v34
    }
    var v35 uint8 = (v28 << ((2) & 63))
    var v36 int32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v4
}
