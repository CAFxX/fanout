package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 string }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 uint64 }

func main() {
    var v1 uint = 2
    v2 := make(map[int]uint32)
    var v3 uint16 = 65535
    var v4 int64 = 42
    var v5 uint64 = 2
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := v2[fi6]
        var v8 int = 128
        var v9 int8 = 100
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: (v4 % (((v4 << (((v1 * v1)) & 63))) | 1)), F1: 332, F2: 181}
    if false {
        var v11 float32 = (3.14159 * (0.0 - -2.25))
        if (102 != 128) {
            v11 += 0.1
            var v12 bool = false
            _ = v12
        }
        _ = v11
    }
    var v13 int16 = 42
    v14 := GenS2{F0: v13, F1: 2, F2: ""}
    var v15 int32 = 103
    if (((("a" + "hello") + "\n\t") + "a") == "hello") {
        var v16 int64 = v4
        fmt.Println(v4, v3)
        _ = v16
    }
    switch 42 {
    case 0:
        var v17 uint16 = 41
        _ = v17
    case -341:
        var v18 uint16 = (^(v3))
        _ = v18
    default:
        for fi19 := 0; fi19 < 8; fi19++ {
            v20 := GenS3{F0: 2147483648, F1: ("" + "a")}
            v21 := GenS3{F0: 0, F1: ("hello" + "")}
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    var v22 int32 = (-(v15))
    fmt.Println(v13, v15, v13)
    for fi23 := 0; fi23 < 1; fi23++ {
        if false {
            v13 -= v13
        } else {
            v15 -= (v22 >> ((453) & 63))
        }
        fmt.Println(v13, v22, v3)
        _ = fi23
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        if false {
            var v25 int32 = (v15 >> ((v1) & 63))
            if false {
                v26 := v2[fi24]
                if true {
                    if false {
                        if false {
                            if (2 != (v26 ^ 2)) {
                                v27 := v2[(fi24 - fi24)]
                                if (int64(v5) >= v4) {
                                    if ((fi24 == fi24) || false) {
                                        fmt.Println(v25, v25, v5)
                                        v1 -= uint(v3)
                                        v5 -= (v5 | 1)
                                    }
                                    v22 += (v15 + v25)
                                } else {
                                    _ = len(v2)
                                }
                                _ = v27
                            } else {
                                switch (("\n\t" + "foo bar") + ("a" + "foo bar")) {
                                case "":
                                    v28 := GenS4{F0: int32(fi24)}
                                    _ = v28
                                default:
                                    if true {
                                        v15 = (^(v15))
                                        var v29 int8 = 42
                                        switch "foo bar" {
                                        case "\n\t":
                                            v22 -= (v25 + 42)
                                        case "":
                                            var v30 string = "hello"
                                            _ = v30
                                        }
                                        _ = v29
                                    }
                                }
                                var v31 float32 = ((0.1 - 0.0) + 0.1)
                                _ = v31
                            }
                        } else {
                            var v32 int32 = v25
                            if true {
                                var v33 int = 2147483647
                                v15 += (v32 - 2147483647)
                                _ = v33
                            }
                            _ = v32
                        }
                    }
                    if false {
                        for fi34 := 0; fi34 < 15; fi34++ {
                            var v35 uint = v1
                            switch 1 {
                            case 128:
                                v26 = 82
                            case 9223372036854775807:
                                var v36 int32 = v25
                                var v37 uint64 = (v5 % (((v5 - (v5 << (((^(v1))) & 63)))) | 1))
                                v3 -= (^(v3))
                                _ = v36
                                _ = v37
                            case -9223372036854775808:
                                v38 := GenS3{F0: (v26 % ((v26) | 1)), F1: ""}
                                _ = v38
                            }
                            var v39 int = (fi34 & -1)
                            _ = fi34
                            _ = v35
                            _ = v39
                        }
                        switch (((("foo bar" + "a") + "a") + "hello") + "") {
                        case "foo bar":
                            if false {
                                v26 -= (^(v26))
                                delete(v2, (fi24 & fi24))
                            } else {
                                v1 = (^(v1))
                                v26 += v26
                            }
                        case "hello":
                            if true {
                                v22 -= v22
                            }
                        }
                    }
                    if (((v4 * (v4 << (((v1 * (v1 & (v1 << ((v1) & 63))))) & 63))) >= -9223372036854775808) || false) {
                        if (!(0 != v26)) {
                            var v40 int16 = int16(v1)
                            v13 -= (-(v40))
                            _ = v40
                        } else {
                            v41 := GenS3{F0: v26, F1: "a"}
                            var v42 int64 = v4
                            _ = v41
                            _ = v42
                        }
                        for fi43 := 0; fi43 < 15; fi43++ {
                            if false {
                                var v44 int16 = 42
                                v45 := GenS1{F0: int64(v22), F1: v1, F2: 0}
                                fmt.Println(v25, v44)
                                _ = v44
                                _ = v45
                            } else {
                                var v46 string = ""
                                v46 = "foo bar"
                                v47 := GenS5{F0: 127}
                                _ = v46
                                _ = v47
                            }
                            v48 := GenS6{F0: (v13 / ((v13) | 1)), F1: uint64(v1)}
                            var v49 int = (fi43 % ((fi43) | 1))
                            _ = fi43
                            _ = v48
                            _ = v49
                        }
                    } else {
                        var v50 int64 = (v4 << ((834) & 63))
                        fmt.Println(v26, v4, v50)
                        _ = v50
                    }
                }
                fmt.Println(v5, v3, v13)
                _ = v26
            } else {
                var v51 int64 = (v4 & v4)
                for fi52 := 0; fi52 < 4; fi52++ {
                    v3 = 32767
                    if ((v15 * v22) < (v25 ^ 2147483647)) {
                        switch (fi52 * 390) {
                        case 42:
                            v5 = v5
                        case 32767:
                            v51 += v4
                        }
                        if false {
                            fmt.Println(v4, v51, v25)
                            var v53 float32 = 1.5
                            _ = v53
                        } else {
                            v54 := GenS4{F0: v25}
                            _ = v54
                        }
                        v1 += v1
                    } else {
                        v51 -= (v4 * (-(v4)))
                    }
                    var v55 uint32 = 0
                    _ = fi52
                    _ = v55
                }
                _ = v51
            }
            if false {
                switch int(v3) {
                case -2147483648:
                    if true {
                        v3 -= v3
                    }
                case -13:
                    v56 := []int{}
                    _ = v56
                default:
                    for fi57 := 0; fi57 < 3; fi57++ {
                        switch "hello" {
                        case "a":
                            _ = len(v2)
                        case "":
                            v3 = (v3 | v3)
                        }
                        var v58 float32 = ((((1e10 + (0.1 * (1e10 + (float32((0.0 + (0.0 - -2.25))) - 0.1)))) - (((1.5 * 0.1) + float32(1.5)) + (3.14159 / -2.25))) / float32(3.14159)) * (0.1 / (-2.25 - 1.5)))
                        _ = fi57
                        _ = v58
                    }
                }
                var v59 int8 = -42
                _ = v59
            } else {
                var v60 bool = (true && (!(!false)))
                _ = v60
            }
            _ = v25
        } else {
            _ = len(v2)
        }
        _ = fi24
    }
    var v61 uint = v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v61
}
