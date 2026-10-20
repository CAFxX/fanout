package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint }

func main() {
    var v1 int32 = 2147483647
    var v2 int32 = v1
    v3 := make(map[string]bool)
    var v4 uint32 = 357
    var v5 int8 = 0
    var v6 int16 = 0
    v7 := GenS1{F0: v4, F1: v6}
    v8 := GenS1{F0: v4, F1: 0}
    for fi9 := 0; fi9 < 15; fi9++ {
        v5 += -1
        fmt.Println(v4)
        if (!(true && false)) {
            fmt.Println(v5, v6, fi9)
            var v10 string = "foo bar"
            var v11 string = ""
            _ = v10
            _ = v11
        }
        _ = fi9
    }
    fmt.Println(v5)
    switch 4294967295 {
    case 1:
        var v12 string = "hello"
        _ = v12
    case 42:
        var v13 int64 = -9223372036854775808
        _ = v13
    case -9223372036854775808:
        fmt.Println(v1)
    default:
        if false {
            var v14 uint64 = 57
            _ = v14
        }
    }
    if false {
        v15 := []int{955, 32767, -2147483648, -147}
        v5 = v5
        v16 := GenS2{F0: 570}
        _ = v15
        _ = v16
    } else {
        switch ("\n\t" + "a") {
        case "":
            if (1 >= 18446744073709551615) {
                if (!false) {
                    v2 = (v1 * v2)
                    fmt.Println(v5, v1)
                    v1 += v2
                } else {
                    for fi17 := 0; fi17 < 16; fi17++ {
                        v18 := []int{fi17}
                        var v19 uint16 = 65535
                        var v20 uint8 = 220
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    v21 := GenS1{F0: (v4 >> ((uint(v4)) & 63)), F1: 386}
                    _ = v21
                }
                var v22 int = 534
                _ = v22
            }
        case "\n\t":
            v23 := GenS3{F0: 42}
            _ = v23
        case "hello":
            fmt.Println(v5, v2, v4)
        }
        for fi24 := 0; fi24 < 15; fi24++ {
            var v25 int8 = (v5 << ((uint(v4)) & 63))
            var v26 float64 = 1e10
            switch 9223372036854775807 {
            case -187:
                var v27 uint16 = 716
                _ = v27
            case 0:
                v4 = v4
            case 752:
                v3["a"] = false
            }
            _ = fi24
            _ = v25
            _ = v26
        }
        fmt.Println(v2, v2, v4)
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        v6 = v6
        v29 := []int{fi28, 32768, 4294967295, (fi28 * (fi28 / (((fi28 & fi28)) | 1)))}
        _ = fi28
        _ = v29
    }
    if true {
        for fi30 := 0; fi30 < 15; fi30++ {
            for fi31 := 0; fi31 < 13; fi31++ {
                v2 += (v2 >> ((65535) & 63))
                _ = fi31
            }
            for fi32 := 0; fi32 < 20; fi32++ {
                if (2147483648 == 2147483648) {
                    var v33 int = (fi30 / (((fi30 | 127)) | 1))
                    _ = v33
                }
                if true {
                    var v34 uint64 = 42
                    fmt.Println(v6)
                    fmt.Println(v1, v34, v2)
                    _ = v34
                }
                var v35 uint64 = 42
                _ = fi32
                _ = v35
            }
            for fi36 := 0; fi36 < 2; fi36++ {
                v6 += 0
                v37 := GenS2{F0: 1}
                _ = fi36
                _ = v37
            }
            _ = fi30
        }
        for fi38 := 0; fi38 < 20; fi38++ {
            if (!(fi38 >= (fi38 % (((fi38 * (fi38 >> ((42) & 63)))) | 1)))) {
                v39 := GenS2{F0: uint64(v4)}
                switch ("a" + "") {
                case "":
                    var v40 uint64 = 42
                    _ = v40
                case "hello":
                    v2 = (v2 | (^(v1)))
                case "\n\t":
                    for fi41 := 0; fi41 < 6; fi41++ {
                        if ((((float64(0.1) * 1e10) * float64(float32(0.1))) / 1.0) >= 0.0) {
                            fmt.Println(v4, fi38, v5)
                            fmt.Println(v2, v2)
                            v42 := []int{(^(fi41)), 2147483647}
                            _ = v42
                        } else {
                            v5 += v5
                            fmt.Println(v6)
                            var v43 int16 = (v6 << ((1) & 63))
                            _ = v43
                        }
                        v44 := []int{(-(fi41)), fi41, fi38}
                        _ = fi41
                        _ = v44
                    }
                default:
                    v1 = v2
                }
                fmt.Println(v6, fi38)
                _ = v39
            } else {
                var v45 int16 = 42
                var v46 string = "\n\t"
                fmt.Println(v46, v46, v45)
                _ = v45
                _ = v46
            }
            _ = fi38
        }
        fmt.Println(v5, v5)
    } else {
        for fi47 := 0; fi47 < 12; fi47++ {
            fmt.Println(v5, v4)
            _ = fi47
        }
        delete(v3, ("\n\t" + "foo bar"))
    }
    fmt.Println(v5, v4)
    v48 := []int{128}
    v48 = append(v48, -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
