package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint16; F1 uint64; F2 bool }
type GenS4 struct { F0 int; F1 int8 }

func main() {
    var v1 int8 = 42
    var v2 int16 = 42
    var v3 string = "foo bar"
    v4 := GenS1{F0: 1, F1: 1}
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 int16 = -322
            v3 = v3
            _ = fi6
            _ = v7
        }
        if false {
            if ((int32(fi5) != -1) && false) {
                v3 = ("a" + v3)
                for fi8 := 0; fi8 < 6; fi8++ {
                    switch fi5 {
                    case 0:
                        if (!false) {
                            fmt.Println(fi8, v1, fi8)
                        } else {
                            if true {
                                var v9 int64 = 0
                                _ = v9
                            } else {
                                if (!false) {
                                    switch fi8 {
                                    case -7:
                                        v10 := GenS2{F0: 0, F1: 75, F2: 18446744073709551615}
                                        _ = v10
                                    case 1:
                                        var v11 int32 = 2147483647
                                        _ = v11
                                    }
                                    v12 := GenS3{F0: 0, F1: 783, F2: ((fi5 & fi8) <= 32767)}
                                    var v13 float32 = 0.0
                                    _ = v12
                                    _ = v13
                                }
                                v2 = 42
                                var v14 uint = 4294967295
                                _ = v14
                            }
                            v15 := []int{386, (fi5 - fi8), fi5}
                            switch ("\n\t" + v3) {
                            case "\n\t":
                                var v16 string = ((((v3 + v3) + v3) + ((v3 + "hello") + "foo bar")) + "hello")
                                _ = v16
                            case "":
                                v2 += 32767
                                switch v3 {
                                case "\n\t":
                                    fmt.Println(v3, v2)
                                case "":
                                    v17 := GenS1{F0: -9223372036854775808, F1: 42}
                                    _ = v17
                                default:
                                    var v18 int16 = (v2 - (v2 & -509))
                                    _ = v18
                                }
                                if len(v15) > 0 {
                                    v15[0] = fi5
                                }
                            default:
                                fmt.Println(fi8)
                            }
                            v3 = ("a" + ((v3 + v3) + ""))
                            _ = v15
                        }
                    }
                    var v19 bool = false
                    if v19 {
                        v20 := GenS3{F0: 17, F1: 4294967295, F2: (-749 <= 237)}
                        _ = v20
                    }
                    _ = fi8
                    _ = v19
                }
            } else {
                var v21 int = fi5
                var v22 int16 = v2
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 4; fi23++ {
                if true {
                    v24 := GenS3{F0: 206, F1: 18446744073709551615, F2: (!(fi23 < (fi23 >> ((18446744073709551615) & 63))))}
                    if false {
                        v2 = v2
                        if true {
                            var v25 int64 = -9223372036854775808
                            v26 := GenS1{F0: int64(v1), F1: 4294967295}
                            v2 = v2
                            v3 = ""
                            var v27 int16 = (v2 << ((65535) & 63))
                            v25 += v25
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                        var v28 int = -9223372036854775808
                        _ = v28
                    }
                    _ = v24
                } else {
                    var v29 string = v3
                    _ = v29
                }
                v2 = v2
                _ = fi23
            }
            switch "foo bar" {
            case "":
                var v30 int64 = 1
                _ = v30
            case "\n\t":
                fmt.Println(v2, v3)
            }
        }
        v31 := []int{fi5}
        _ = fi5
        _ = v31
    }
    switch 1 {
    case 127:
        v1 += (v1 % ((42) | 1))
    case 9223372036854775807:
        v32 := GenS4{F0: 2147483647, F1: 42}
        _ = v32
    case 1:
        var v33 float32 = 1.5
        _ = v33
    default:
        v34 := []int{}
        _ = v34
    }
    v3 = "foo bar"
    fmt.Println(v1, v3)
    for fi35 := 0; fi35 < 6; fi35++ {
        v3 = v3
        v3 = (("" + "\n\t") + v3)
        _ = fi35
    }
    fmt.Println(v2)
    v1 = v1
    if (v3 != "\n\t") {
        fmt.Println(v1, v2, v3)
    }
    v2 -= v2
    var v36 int32 = -2147483648
    var v37 int = -2147483648
    v38 := []int{v37, -1, 9223372036854775807}
    for fi39 := 0; fi39 < 9; fi39++ {
        if false {
            var v40 int64 = int64(fi39)
            _ = v40
        } else {
            v41 := GenS3{F0: 32767, F1: 452, F2: false}
            var v42 int32 = v36
            var v43 int8 = (v1 & (^(v1)))
            _ = v41
            _ = v42
            _ = v43
        }
        var v44 int = (-(v37))
        _ = fi39
        _ = v44
    }
    if len(v38) >= 2 {
        v38 = v38[1:len(v38)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
