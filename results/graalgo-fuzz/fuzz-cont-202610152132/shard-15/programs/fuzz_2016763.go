package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 string = (("a" + "\n\t") + "a")
    var v2 float32 = float32(1.5)
    v3 := make(map[string]string)
    var v4 uint = 255
    if ((true && true) || true) {
        var v5 int = 32767
        var v6 uint = 18446744073709551615
        if true {
            var v7 uint64 = 1
            var v8 uint = v6
            var v9 string = (v1 + v1)
            _ = v7
            _ = v8
            _ = v9
        } else {
            if false {
                for fi10 := 0; fi10 < 19; fi10++ {
                    v6 = v4
                    var v11 uint64 = 220
                    v12 := []int{}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                var v13 uint = (v4 * (v6 % (((v6 & v6)) | 1)))
                _ = v13
            } else {
                v2 -= float32((-2.25 * 1e10))
            }
            v1 = v1
        }
        _ = v5
        _ = v6
    } else {
        fmt.Println(v2)
    }
    var v14 uint16 = 688
    if true {
        fmt.Println(v4, v4, v14)
    }
    v15 := []int{-66, -2147483648}
    for fi16 := 0; fi16 < 4; fi16++ {
        switch (fi16 - int(v14)) {
        case 2147483647:
            if false {
                for fi17 := 0; fi17 < 9; fi17++ {
                    v4 = (v4 % ((v4) | 1))
                    v4 = 1
                    _ = fi17
                }
                v2 = v2
                for fi18 := 0; fi18 < 13; fi18++ {
                    switch fi18 {
                    case 0:
                        _ = len(v15)
                    }
                    _ = fi18
                }
            }
        case -1:
            var v19 bool = false
            _ = v19
        default:
            var v20 string = v1
            _ = v20
        }
        v21 := GenS1{F0: 213, F1: (fi16 % ((708) | 1))}
        _ = fi16
        _ = v21
    }
    var v22 float32 = 1e10
    v22 -= v22
    v2 = 0.1
    switch "" {
    case "foo bar":
        for fi23 := 0; fi23 < 5; fi23++ {
            switch fi23 {
            case 9223372036854775807:
                var v24 uint16 = (v14 >> ((v4) & 63))
                _ = v24
            case -2147483648:
                v25 := GenS1{F0: 352, F1: -9223372036854775808}
                _ = v25
            default:
                v14 = v14
            }
            v26 := GenS1{F0: 4294967295, F1: -804}
            _ = fi23
            _ = v26
        }
    case "a":
        v1 = "\n\t"
    default:
        switch ((v1 + "") + ((v1 + v1) + v1)) {
        case "a":
            if ((!(1 >= 282)) || true) {
                var v27 bool = true
                var v28 uint64 = 69
                _ = v27
                _ = v28
            } else {
                v29 := GenS2{F0: (false || (177 < 644))}
                for fi30 := 0; fi30 < 7; fi30++ {
                    v22 = 3.14159
                    v31 := GenS1{F0: 2, F1: fi30}
                    _ = fi30
                    _ = v31
                }
                v32 := GenS1{F0: 2, F1: -2147483648}
                _ = v29
                _ = v32
            }
        case "foo bar":
            for fi33 := 0; fi33 < 18; fi33++ {
                fmt.Println(v14, fi33, fi33)
                v14 -= v14
                v34 := GenS3{F0: -32768, F1: 25}
                _ = fi33
                _ = v34
            }
        default:
            switch 32768 {
            case 9223372036854775807:
                var v35 uint8 = 0
                _ = v35
            case 819:
                v22 = 1.5
            default:
                v22 = (float32((-2.25 / 1.0)) + 1.5)
            }
        }
    }
    v22 += v2
    switch "\n\t" {
    case "a":
        delete(v3, ("foo bar" + ""))
    }
    var v36 uint64 = 0
    var v37 uint8 = 1
    v37 = v37
    if ((true && true) || false) {
        v14 -= (v14 - (^(v14)))
        if false {
            for fi38 := 0; fi38 < 6; fi38++ {
                var v39 uint32 = 1
                _ = fi38
                _ = v39
            }
            if (32767 <= 708) {
                for fi40 := 0; fi40 < 11; fi40++ {
                    for fi41 := 0; fi41 < 5; fi41++ {
                        v42 := GenS4{F0: v37}
                        _ = fi41
                        _ = v42
                    }
                    fmt.Println(v2)
                    fmt.Println(v14, v14, v36)
                    _ = fi40
                }
            }
            var v43 int16 = 679
            _ = v43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
