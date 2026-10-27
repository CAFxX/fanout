package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint; F2 int8 }
type GenS3 struct { F0 int; F1 int8; F2 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]string)
    var v3 string = "a"
    v4 := GenS1{F0: 634, F1: 1, F2: 42}
    for fi5 := 0; fi5 < 13; fi5++ {
        switch v3 {
        case "":
            for fi6 := 0; fi6 < 9; fi6++ {
                var v7 int64 = -1
                _ = fi6
                _ = v7
            }
        case "foo bar":
            v8 := GenS1{F0: 32767, F1: 32767, F2: -1}
            _ = v8
        case "a":
            v3 = "hello"
        }
        v9 := GenS1{F0: 363, F1: 587, F2: -1}
        fmt.Println(v3)
        _ = fi5
        _ = v9
    }
    v3 = ((v3 + v3) + "hello")
    v3 = "a"
    for fi10 := 0; fi10 < 7; fi10++ {
        if true {
            for fi11 := 0; fi11 < 10; fi11++ {
                v12 := []int{fi11, fi11, 127}
                _ = fi11
                _ = v12
            }
            v13 := GenS1{F0: 1, F1: 1, F2: 0}
            switch 419 {
            case 42:
                var v14 float64 = (0.1 - float64(1e10))
                _ = v14
            default:
                v3 = ((v3 + ((v3 + "a") + v3)) + ("a" + ""))
            }
            _ = v13
        }
        _ = fi10
    }
    _ = len(v1)
    _ = len(v2)
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS1{F0: 336, F1: 774, F2: 0}
        v3 = "foo bar"
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        v3 = "foo bar"
        v18 := GenS2{F0: 65535, F1: 18446744073709551615, F2: 100}
        _ = fi17
        _ = v18
    }
    if ((!((((1.5 - 3.14159) + float32(((float64((1.5 - 1.5)) - float64(3.14159)) * 1e10))) + -2.25) == float32(float64(3.14159)))) && true) {
        v3 = (("" + v3) + "")
        if true {
            v19 := GenS3{F0: 4294967295, F1: 44, F2: -2147483648}
            if true {
                v20 := []int{0}
                var v21 uint64 = 1
                for fi22 := 0; fi22 < 15; fi22++ {
                    if false {
                        if true {
                            v20 = append(v20, fi22)
                        }
                    }
                    for fi23 := 0; fi23 < 13; fi23++ {
                        switch v3 {
                        case "a":
                            switch (((("" + ((("foo bar" + v3) + ((v3 + v3) + v3)) + v3)) + "foo bar") + ("\n\t" + (v3 + v3))) + "\n\t") {
                            case "a":
                                fmt.Println(fi23, v21)
                            case "hello":
                                if false {
                                    if (false || true) {
                                        v1[(fi23 ^ fi22)] = ("hello" > "hello")
                                    } else {
                                        v21 = v21
                                        delete(v1, 42)
                                    }
                                    v24 := GenS4{F0: 1}
                                    var v25 int64 = -1
                                    _ = v24
                                    _ = v25
                                } else {
                                    delete(v1, fi23)
                                }
                            case "":
                                if len(v20) > 0 {
                                    v20[0] = 869
                                }
                                if len(v20) > 0 {
                                    v20[0] = 127
                                }
                                if (!((1e10 + 0.1) > ((-2.25 + 1.5) * 1.5))) {
                                    v3 = v3
                                    var v26 float32 = 0.1
                                    var v27 int8 = int8(fi22)
                                    _ = v26
                                    _ = v27
                                }
                            }
                        case "foo bar":
                            switch "foo bar" {
                            case "hello":
                                var v28 uint64 = (v21 % ((v21) | 1))
                                _ = v28
                            }
                        case "":
                            var v29 int32 = int32(v21)
                            _ = v29
                        }
                        v21 = (v21 % (((v21 << ((2) & 63))) | 1))
                        v30 := GenS5{F0: 32767, F1: uint32(v21)}
                        var v31 string = v3
                        var v32 uint16 = 789
                        v31 = (((v3 + v31) + v3) + (v3 + v3))
                        var v33 bool = (32767 == -110)
                        fmt.Println(v32)
                        _ = fi23
                        _ = v30
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    _ = len(v20)
                    _ = fi22
                }
                _ = v20
                _ = v21
            } else {
                var v34 uint8 = 42
                _ = v34
            }
            _ = v19
        }
    }
    var v35 int32 = -98
    var v36 float32 = (3.14159 - (0.1 / ((0.1 - (0.1 - (1.5 - 3.14159))) - float32(-2.25))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
