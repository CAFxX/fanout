package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS2 struct { F0 int32; F1 string; F2 int64 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int64; F1 uint64 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    var v3 int8 = -1
    var v4 uint32 = 298
    var v5 float32 = ((0.0 - 1e10) * 0.0)
    var v6 uint16 = 1
    var v7 float64 = float64((v5 - 1.5))
    v8 := GenS1{F0: (v6 % (((v6 * 759)) | 1)), F1: (v4 ^ v4), F2: (v3 + (v3 % ((100) | 1)))}
    fmt.Println(v4)
    v9 := GenS1{F0: v6, F1: v4, F2: (^(v3))}
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 uint8 = 127
        var v12 int64 = -9223372036854775808
        _ = fi10
        _ = v11
        _ = v12
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v4, v6)
    case "":
        fmt.Println(v7, v4, v4)
    default:
        var v13 uint32 = v4
        _ = v13
    }
    var v14 int8 = v3
    v15 := GenS1{F0: 1, F1: (^(v4)), F2: 127}
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 int16 = 0
        for fi18 := 0; fi18 < 16; fi18++ {
            v3 = v14
            v19 := GenS2{F0: 0, F1: "\n\t", F2: 0}
            switch (fi18 & (fi18 % (((fi18 & (fi18 / (((fi18 | (fi18 | fi16))) | 1)))) | 1))) {
            case 0:
                v20 := GenS1{F0: (^(v6)), F1: v4, F2: 1}
                _ = v20
            case 763:
                fmt.Println(v14, v4, v17)
            }
            _ = fi18
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    if (0 < -1) {
        v5 = 3.14159
        v1[""] = 32767
    } else {
        v6 = (v6 - v6)
        fmt.Println(v3)
    }
    var v21 bool = false
    if (!v21) {
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 int32 = 42
            var v24 int32 = 257
            _ = fi22
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 4; fi25++ {
            v26 := []int{(fi25 / ((fi25) | 1))}
            _ = fi25
            _ = v26
        }
    } else {
        for fi27 := 0; fi27 < 11; fi27++ {
            v14 = (-(v3))
            _ = fi27
        }
        if v21 {
            for fi28 := 0; fi28 < 19; fi28++ {
                v29 := []int{}
                for fi30 := 0; fi30 < 6; fi30++ {
                    switch fi30 {
                    case 546:
                        if ((fi30 << ((223) & 63)) > fi28) {
                            var v31 uint64 = 161
                            switch ("\n\t" + ("a" + "\n\t")) {
                            case "foo bar":
                                if true {
                                    var v32 int = fi28
                                    _ = v32
                                } else {
                                    v4 = (v4 & (^(v4)))
                                    v33 := GenS1{F0: (^(v6)), F1: v4, F2: 25}
                                    _ = v33
                                }
                            }
                            _ = v31
                        }
                    case -752:
                        fmt.Println(v6)
                    default:
                        fmt.Println(v3, v3)
                        v3 += -1
                        v5 -= v5
                    }
                    _ = fi30
                }
                v3 = -128
                _ = fi28
                _ = v29
            }
        } else {
            v34 := GenS3{F0: v6, F1: 0}
            _ = v34
        }
        var v35 int16 = 1000
        _ = v35
    }
    v21 = v21
    var v36 int64 = 0
    var v37 uint64 = 42
    switch "a" {
    case "a":
        if (v21 && v21) {
            switch 42 {
            case 2147483647:
                if (!(true && false)) {
                    var v38 float32 = v5
                    _ = v38
                }
            case 32768:
                fmt.Println(v7, v36)
            default:
                if (2 > 2) {
                    v36 -= v36
                    switch 2147483647 {
                    case 32767:
                        v39 := GenS4{F0: (^(v36)), F1: (^(v37))}
                        _ = v39
                    default:
                        v6 += (^(v6))
                    }
                } else {
                    v37 = (^(v37))
                }
            }
        } else {
            var v40 uint16 = v6
            var v41 int8 = -46
            _ = v40
            _ = v41
        }
    case "\n\t":
        var v42 int8 = (v14 >> ((uint(v6)) & 63))
        _ = v42
    case "":
        if false {
            switch "" {
            case "":
                if (!true) {
                    for fi43 := 0; fi43 < 19; fi43++ {
                        v37 += (^(v37))
                        _ = len(v2)
                        v44 := GenS1{F0: (v6 | v6), F1: (^(v4)), F2: v14}
                        _ = fi43
                        _ = v44
                    }
                }
            case "a":
                var v45 bool = v21
                _ = v45
            }
            if true {
                fmt.Println(v6, v4, v37)
                var v46 uint = 755
                _ = v46
            }
            if v21 {
                var v47 float32 = v5
                switch "" {
                case "foo bar":
                    var v48 int8 = 127
                    _ = v48
                case "":
                    v36 -= (v36 - v36)
                }
                _ = v47
            } else {
                v49 := GenS4{F0: v36, F1: (v37 * (v37 >> ((1) & 63)))}
                v50 := []int{-1, -970}
                v51 := GenS1{F0: v6, F1: 2147483648, F2: v14}
                _ = v49
                _ = v50
                _ = v51
            }
        }
    }
    for fi52 := 0; fi52 < 8; fi52++ {
        v53 := GenS3{F0: (^(v6)), F1: 2}
        _ = fi52
        _ = v53
    }
    v4 -= 227
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
