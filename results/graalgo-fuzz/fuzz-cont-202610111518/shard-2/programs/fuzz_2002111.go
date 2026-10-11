package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint; F2 string }

func main() {
    var v1 uint = 598
    var v2 uint8 = 164
    var v3 int16 = 1000
    var v4 uint64 = uint64(v1)
    fmt.Println(v4, v3)
    if false {
        var v5 uint8 = v2
        v6 := GenS1{F0: 255, F1: 42}
        v7 := GenS2{F0: -989}
        _ = v5
        _ = v6
        _ = v7
    } else {
        if (!true) {
            if false {
                v2 = 2
            } else {
                var v8 int32 = 2147483647
                v2 -= v2
                v2 -= (v2 - v2)
                _ = v8
            }
            switch 2147483647 {
            case 128:
                var v9 float64 = (3.14159 - 0.1)
                _ = v9
            }
        } else {
            fmt.Println(v1, v4, v3)
        }
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 bool = (v2 < (^(v2)))
        for fi12 := 0; fi12 < 4; fi12++ {
            fmt.Println(v3)
            _ = fi12
        }
        v3 = (v3 + 32767)
        _ = fi10
        _ = v11
    }
    switch -713 {
    case 127:
        if (1 < 104) {
            var v13 float64 = (0.0 / 1.0)
            if false {
                v13 = (v13 - (((3.14159 / v13) / (v13 - 1.5)) * float64(0.0)))
                v14 := []int{9223372036854775807, -113, 127, 570}
                fmt.Println(v3)
                _ = v14
            } else {
                var v15 uint32 = uint32(v4)
                v1 -= v1
                _ = v15
            }
            switch "hello" {
            case "foo bar":
                var v16 bool = (v3 > int16(v1))
                _ = v16
            case "hello":
                if true {
                    fmt.Println(v4, v2, v13)
                    var v17 uint32 = uint32(v1)
                    _ = v17
                } else {
                    switch int(v4) {
                    case 2147483647:
                        var v18 uint32 = 2
                        _ = v18
                    case 0:
                        if true {
                            if (42 >= (v4 + 1)) {
                                for fi19 := 0; fi19 < 7; fi19++ {
                                    for fi20 := 0; fi20 < 1; fi20++ {
                                        fmt.Println(v2)
                                        _ = fi20
                                    }
                                    var v21 float64 = v13
                                    v13 += v21
                                    _ = fi19
                                    _ = v21
                                }
                            } else {
                                switch ("" + "") {
                                case "a":
                                    v3 -= 1000
                                default:
                                    var v22 int32 = -117
                                    _ = v22
                                }
                            }
                            v23 := GenS1{F0: (v2 | v2), F1: 1}
                            for fi24 := 0; fi24 < 6; fi24++ {
                                v2 -= v2
                                v25 := GenS1{F0: 236, F1: -1}
                                v26 := []int{fi24, 1, (-(fi24)), fi24}
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                            _ = v23
                        } else {
                            fmt.Println(v1, v13, v1)
                            for fi27 := 0; fi27 < 11; fi27++ {
                                for fi28 := 0; fi28 < 11; fi28++ {
                                    v4 -= uint64(v2)
                                    switch (fi27 | fi28) {
                                    case 128:
                                        var v29 int16 = (v3 << (((^(v1))) & 63))
                                        _ = v29
                                    case 360:
                                        v30 := GenS2{F0: 1}
                                        _ = v30
                                    default:
                                        if (4294967295 >= 65535) {
                                            fmt.Println(v13, v4, v4)
                                            v31 := []int{-392}
                                            switch fi27 {
                                            case -2147483648:
                                                fmt.Println(v1)
                                            default:
                                                fmt.Println(v13, v2, v1)
                                            }
                                            v32 := GenS2{F0: 0}
                                            v2 -= 0
                                            var v33 float32 = 3.14159
                                            _ = v31
                                            _ = v32
                                            _ = v33
                                        } else {
                                            v1 += 0
                                            fmt.Println(v1, v4)
                                            var v34 uint16 = 32767
                                            _ = v34
                                        }
                                    }
                                    _ = fi28
                                }
                                var v35 uint = v1
                                if false {
                                    for fi36 := 0; fi36 < 19; fi36++ {
                                        v37 := GenS1{F0: v2, F1: 426}
                                        _ = fi36
                                        _ = v37
                                    }
                                    v38 := []int{(fi27 ^ (fi27 >> ((v35) & 63))), (fi27 | (fi27 >> ((18446744073709551615) & 63)))}
                                    _ = v38
                                } else {
                                    var v39 float32 = 1.5
                                    _ = v39
                                }
                                _ = fi27
                                _ = v35
                            }
                        }
                    }
                    v3 = v3
                }
            }
            _ = v13
        }
    case 822:
        var v40 string = "hello"
        _ = v40
    case 32767:
        for fi41 := 0; fi41 < 9; fi41++ {
            fmt.Println(v4, v2, v3)
            v4 -= 0
            _ = fi41
        }
    }
    v2 = (v2 - 90)
    v42 := []int{}
    if len(v42) > 0 {
        v42[0] = -973
    }
    var v43 uint = (v1 / (((v1 & 18446744073709551615)) | 1))
    var v44 string = ("\n\t" + "a")
    var v45 uint = v1
    for fi46 := 0; fi46 < 19; fi46++ {
        _ = len(v42)
        _ = fi46
    }
    v3 = (v3 ^ v3)
    var v47 int = 9223372036854775807
    fmt.Println(v44, v43)
    if false {
        v48 := GenS3{F0: (v43 % ((v1) | 1)), F1: (v43 >> ((v1) & 63)), F2: "a"}
        var v49 int32 = 401
        _ = v48
        _ = v49
    }
    v45 -= v45
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v47
}
