package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int8; F2 int8 }
type GenS3 struct { F0 int; F1 uint8; F2 string }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    var v4 uint = 4294967295
    var v5 uint = (v4 & (v4 - v4))
    var v6 int16 = 304
    var v7 uint32 = 735
    v8 := GenS1{F0: 4294967295, F1: 561, F2: 65535}
    for fi9 := 0; fi9 < 16; fi9++ {
        fmt.Println(v4)
        for fi10 := 0; fi10 < 18; fi10++ {
            v4 -= 923
            v4 += v4
            var v11 int16 = 42
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v12 := GenS1{F0: 42, F1: 42, F2: 93}
    fmt.Println(v5, v5)
    var v13 int = -93
    if false {
        v14 := []int{}
        v14 = append(v14, 1)
        if (!false) {
            var v15 int16 = 0
            for fi16 := 0; fi16 < 15; fi16++ {
                fmt.Println(fi16, fi16)
                _ = fi16
            }
            if ((v6 << ((v4) & 63)) == (v15 / ((v6) | 1))) {
                for fi17 := 0; fi17 < 17; fi17++ {
                    if true {
                        for fi18 := 0; fi18 < 17; fi18++ {
                            v13 -= fi17
                            switch ("\n\t" + ("a" + "\n\t")) {
                            case "a":
                                var v19 float32 = (float32(((3.14159 - (0.0 * 1e10)) / (float64(1.5) * 1e10))) * -2.25)
                                _ = v19
                            case "foo bar":
                                if true {
                                    var v20 uint = v5
                                    _ = v20
                                } else {
                                    var v21 uint = 255
                                    fmt.Println(v7, v13, v15)
                                    v22 := GenS2{F0: (^(v7)), F1: 127, F2: 85}
                                    _ = v21
                                    _ = v22
                                }
                            }
                            v23 := GenS3{F0: fi17, F1: 2, F2: "foo bar"}
                            _ = fi18
                            _ = v23
                        }
                    }
                    for fi24 := 0; fi24 < 3; fi24++ {
                        v5 += v5
                        v6 += int16(v5)
                        var v25 uint8 = 255
                        _ = fi24
                        _ = v25
                    }
                    v13 = 607
                    _ = fi17
                }
                for fi26 := 0; fi26 < 6; fi26++ {
                    for fi27 := 0; fi27 < 15; fi27++ {
                        var v28 float64 = 1e10
                        if false {
                            v29 := GenS3{F0: (v13 & 9223372036854775807), F1: 39, F2: "\n\t"}
                            _ = v29
                        }
                        _ = fi27
                        _ = v28
                    }
                    _ = fi26
                }
            } else {
                if (true && (true && ((v15 >> ((255) & 63)) > -526))) {
                    var v30 uint = (^(v5))
                    _ = v30
                }
                var v31 uint64 = 42
                var v32 int16 = 1
                _ = v31
                _ = v32
            }
            _ = v15
        } else {
            var v33 int16 = 1
            v34 := GenS3{F0: v13, F1: 127, F2: "hello"}
            _ = v33
            _ = v34
        }
        _ = v14
    } else {
        if false {
            v35 := GenS1{F0: 18446744073709551615, F1: v13, F2: 0}
            var v36 int = v13
            var v37 uint16 = 0
            _ = v35
            _ = v36
            _ = v37
        } else {
            var v38 int32 = -1
            _ = v38
        }
    }
    var v39 uint16 = 191
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
