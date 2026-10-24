package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 int }

func main() {
    v1 := make(map[int]string)
    v2 := []int{42, 1, -6}
    var v3 uint16 = 65535
    for fi4 := 0; fi4 < 17; fi4++ {
        v5 := GenS1{F0: 255}
        if (!(!true)) {
            delete(v1, fi4)
            v6 := v1[fi4]
            _ = v6
        } else {
            var v7 int64 = 28
            _ = v7
        }
        fmt.Println(fi4, v3)
        _ = fi4
        _ = v5
    }
    var v8 uint32 = 4294967295
    v3 = v3
    if false {
        var v9 int16 = 32767
        for fi10 := 0; fi10 < 1; fi10++ {
            switch (-(fi10)) {
            case 42:
                if (3.14159 < 1.5) {
                    v9 += v9
                    v9 -= -1
                }
            case -277:
                delete(v1, (fi10 ^ 32768))
            case 2147483647:
                var v11 bool = (100 == -1)
                _ = v11
            }
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 float64 = ((1e10 / (0.1 - 0.0)) / 1.0)
                _ = fi12
                _ = v13
            }
            _ = fi10
        }
        for fi14 := 0; fi14 < 17; fi14++ {
            for fi15 := 0; fi15 < 3; fi15++ {
                v9 = v9
                _ = fi15
            }
            if (255 < 42) {
                v9 = v9
            }
            v1[(fi14 % ((fi14) | 1))] = "foo bar"
            _ = fi14
        }
        _ = v9
    } else {
        switch -414 {
        case -1:
            var v16 bool = (((v8 * v8) < (v8 - (v8 & 130))) && (!true))
            _ = v16
        default:
            for fi17 := 0; fi17 < 19; fi17++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi17
            }
        }
        v8 = v8
        v18 := GenS2{F0: 125}
        _ = v18
    }
    fmt.Println(v3)
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 float32 = float32(float64((3.14159 / float32((0.1 * (-2.25 / 0.1))))))
        _ = fi19
        _ = v20
    }
    delete(v1, -2147483648)
    v3 = (v3 << ((0) & 63))
    delete(v1, 587)
    switch "foo bar" {
    case "a":
        _ = len(v1)
    default:
        switch "a" {
        case "":
            v21 := GenS1{F0: uint(v3)}
            _ = v21
        case "a":
            var v22 int = -9223372036854775808
            _ = v22
        case "foo bar":
            v23 := GenS1{F0: 601}
            _ = v23
        }
    }
    _ = len(v1)
    switch "a" {
    case "a":
        switch -9223372036854775808 {
        case 877:
            switch ("foo bar" + "hello") {
            case "\n\t":
                v8 += 1
            case "hello":
                for fi24 := 0; fi24 < 9; fi24++ {
                    v25 := GenS2{F0: -17}
                    var v26 bool = true
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            case "foo bar":
                var v27 string = ("foo bar" + "hello")
                _ = v27
            }
        }
    case "hello":
        v28 := GenS3{F0: 0, F1: 239}
        _ = v28
    default:
        for fi29 := 0; fi29 < 16; fi29++ {
            if ((3.14159 > -2.25) && false) {
                var v30 uint16 = 50
                switch fi29 {
                case -364:
                    if ("a" > "a") {
                        if false {
                            switch "\n\t" {
                            case "":
                                for fi31 := 0; fi31 < 13; fi31++ {
                                    v32 := v1[fi31]
                                    fmt.Println(v30, v8)
                                    if (!((v8 % (((v8 * 4294967295)) | 1)) >= (v8 % ((0) | 1)))) {
                                        var v33 uint = 4294967295
                                        _ = v33
                                    }
                                    _ = fi31
                                    _ = v32
                                }
                            case "foo bar":
                                v34 := GenS1{F0: 42}
                                _ = v34
                            default:
                                var v35 uint32 = (^(v8))
                                _ = v35
                            }
                        } else {
                            v36 := v1[(fi29 | (fi29 / (((fi29 | 1)) | 1)))]
                            _ = v36
                        }
                    }
                }
                _ = v30
            } else {
                _ = len(v1)
                var v37 uint64 = 2
                _ = v37
            }
            _ = fi29
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
