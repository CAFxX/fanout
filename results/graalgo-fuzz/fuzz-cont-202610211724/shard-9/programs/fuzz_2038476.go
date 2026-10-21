package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int }
type GenS4 struct { F0 uint; F1 int16 }

func main() {
    var v1 uint64 = 454
    var v2 int = 1
    var v3 int16 = 42
    var v4 bool = true
    v5 := make(map[string]bool)
    v3 -= v3
    var v6 int16 = v3
    fmt.Println(v2, v6, v1)
    var v7 uint32 = 65535
    var v8 uint64 = v1
    for fi9 := 0; fi9 < 19; fi9++ {
        v3 += v3
        _ = fi9
    }
    delete(v5, "\n\t")
    if v4 {
        v10 := GenS1{F0: int64(v2)}
        if v4 {
            fmt.Println(v4, v6)
            fmt.Println(v8)
            if v4 {
                var v11 uint = uint(v7)
                v12 := []int{int(v1)}
                fmt.Println(v1, v11, v7)
                _ = v11
                _ = v12
            } else {
                var v13 int32 = 1
                _ = v13
            }
        } else {
            if v4 {
                var v14 int8 = -42
                if v4 {
                    if false {
                        switch "foo bar" {
                        case "hello":
                            for fi15 := 0; fi15 < 1; fi15++ {
                                var v16 uint8 = 210
                                _ = fi15
                                _ = v16
                            }
                        }
                        for fi17 := 0; fi17 < 10; fi17++ {
                            fmt.Println(v14)
                            v18 := GenS2{F0: 42}
                            _ = fi17
                            _ = v18
                        }
                        v19 := GenS3{F0: v8, F1: 2, F2: (v2 ^ (v2 % ((v2) | 1)))}
                        _ = v19
                    } else {
                        v4 = v4
                    }
                    v20 := GenS2{F0: 145}
                    for fi21 := 0; fi21 < 20; fi21++ {
                        v14 = v14
                        var v22 bool = v4
                        _ = fi21
                        _ = v22
                    }
                    _ = v20
                }
                _ = v14
            }
            for fi23 := 0; fi23 < 9; fi23++ {
                v24 := GenS1{F0: 0}
                _ = fi23
                _ = v24
            }
            var v25 string = ("foo bar" + ("" + "foo bar"))
            _ = v25
        }
        _ = v10
    } else {
        v1 = 42
        v2 += v2
    }
    if (v4 || v4) {
        switch ("\n\t" + "a") {
        case "\n\t":
            var v26 bool = (((^(v2)) >= v2) && (-41 <= -128))
            _ = v26
        case "hello":
            switch v2 {
            case 9223372036854775807:
                v7 = v7
            case -9223372036854775808:
                fmt.Println(v6, v6)
            default:
                v2 += v2
            }
        case "":
            var v27 uint8 = 2
            _ = v27
        default:
            v6 -= 0
        }
        if (-9223372036854775808 >= 133) {
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 uint = 787
                for fi30 := 0; fi30 < 5; fi30++ {
                    switch "a" {
                    case "":
                        var v31 int64 = 42
                        _ = v31
                    case "foo bar":
                        switch "a" {
                        case "a":
                            var v32 int64 = 0
                            _ = v32
                        case "\n\t":
                            if v4 {
                                var v33 uint16 = uint16(v1)
                                _ = v33
                            }
                        }
                    }
                    _ = fi30
                }
                _ = fi28
                _ = v29
            }
            v34 := v5["\n\t"]
            _ = v34
        } else {
            var v35 uint32 = v7
            _ = v35
        }
        v36 := GenS4{F0: 42, F1: (v6 / ((v6) | 1))}
        _ = v36
    }
    v5["a"] = false
    for fi37 := 0; fi37 < 12; fi37++ {
        var v38 string = "a"
        v39 := GenS2{F0: 0}
        _ = fi37
        _ = v38
        _ = v39
    }
    for fi40 := 0; fi40 < 11; fi40++ {
        var v41 uint8 = 239
        v42 := v5["hello"]
        if false {
            v4 = (-9223372036854775808 > (v2 / ((128) | 1)))
        }
        _ = fi40
        _ = v41
        _ = v42
    }
    fmt.Println(v6, v6, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
