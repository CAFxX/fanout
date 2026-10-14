package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 string }
func genFn0(p0 int8) int16 {
    for fi7 := 0; fi7 < 2; fi7++ {
        if true {
            switch ("foo bar" + "a") {
            case "foo bar":
                p0 = (p0 % ((0) | 1))
            }
            var v8 int8 = p0
            _ = v8
        } else {
            var v9 uint8 = 46
            _ = v9
        }
        _ = fi7
    }
    p0 = p0
    _ = p0
    return 657
}
func genFn1(p0 uint, p1 uint64) uint64 {
    var v10 int8 = 52
    if false {
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 uint16 = 819
            _ = fi11
            _ = v12
        }
    } else {
        p1 = p1
        var v13 float64 = -2.25
        var v14 int8 = (v10 << ((p0) & 63))
        _ = v13
        _ = v14
    }
    v15 := []int{32767}
    v16 := GenS1{F0: -32768, F1: 2147483647}
    if len(v15) > 0 {
        v15[0] = 42
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v15
    _ = v16
    return 495
}

func main() {
    var v1 float64 = -2.25
    var v2 int32 = 541
    var v3 uint8 = 42
    var v4 uint32 = uint32(v3)
    var v5 int32 = (v2 + v2)
    v6 := []int{986, 128}
    switch "" {
    case "hello":
        var v17 uint = 65535
        _ = v17
    default:
        v3 += 2
    }
    fmt.Println(v2, v5, v4)
    if true {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    } else {
        fmt.Println(v3)
        if false {
            for fi18 := 0; fi18 < 6; fi18++ {
                if false {
                    fmt.Println(fi18, v3)
                } else {
                    switch ("a" + "a") {
                    case "\n\t":
                        for fi19 := 0; fi19 < 15; fi19++ {
                            var v20 int64 = 9223372036854775807
                            var v21 bool = (true || false)
                            switch ("foo bar" + "hello") {
                            case "hello":
                                v22 := GenS1{F0: int16(v4), F1: (v2 | -442)}
                                _ = v22
                            case "foo bar":
                                v3 = 2
                            case "a":
                                v20 = (v20 - v20)
                            }
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    case "a":
                        if false {
                            switch (fi18 >> ((206) & 63)) {
                            case -362:
                                var v23 bool = ((!false) && ((-593 <= (^(v5))) && true))
                                _ = v23
                            case -342:
                                var v24 uint = 2
                                _ = v24
                            default:
                                var v25 int16 = -286
                                _ = v25
                            }
                        }
                    }
                    v1 += v1
                }
                for fi26 := 0; fi26 < 4; fi26++ {
                    v27 := GenS1{F0: 1000, F1: 42}
                    var v28 int32 = v5
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                v29 := GenS1{F0: -32768, F1: -973}
                _ = fi18
                _ = v29
            }
            for fi30 := 0; fi30 < 10; fi30++ {
                for fi31 := 0; fi31 < 16; fi31++ {
                    v4 -= (v4 % (((^(v4))) | 1))
                    var v32 uint = 255
                    _ = fi31
                    _ = v32
                }
                v33 := GenS2{F0: 2}
                switch fi30 {
                case -1:
                    var v34 string = "foo bar"
                    _ = v34
                case 582:
                    v3 = v3
                case 128:
                    var v35 string = ""
                    _ = v35
                }
                _ = fi30
                _ = v33
            }
            var v36 float64 = v1
            _ = v36
        }
        for fi37 := 0; fi37 < 4; fi37++ {
            v38 := GenS3{F0: (v5 % (((v2 / ((v5) | 1))) | 1))}
            var v39 int16 = -32768
            _ = fi37
            _ = v38
            _ = v39
        }
    }
    v2 -= v5
    v4 = v4
    v3 += v3
    switch "foo bar" {
    case "hello":
        if (uint64(v3) >= 42) {
            if false {
                v40 := GenS3{F0: (v5 % (((v2 * v2)) | 1))}
                _ = v40
            }
            v41 := GenS3{F0: (v2 | int32(v4))}
            _ = len(v6)
            _ = v41
        } else {
            fmt.Println(v4)
        }
    }
    _ = len(v6)
    fmt.Println(v4, v2, v1)
    var v42 uint = 42
    v4 = (^(v4))
    v2 -= int32(v3)
    v42 -= v42
    v43 := GenS4{F0: -85, F1: "hello"}
    v44 := GenS4{F0: 0, F1: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
