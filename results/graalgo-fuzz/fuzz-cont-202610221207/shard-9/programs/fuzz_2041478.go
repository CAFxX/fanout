package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 uint; F1 uint }

func main() {
    var v1 uint = 21
    var v2 int16 = -32768
    v3 := []string{}
    var v4 int32 = -2147483648
    var v5 int16 = (v2 >> ((v1) & 63))
    v6 := GenS1{F0: 4294967295, F1: 1, F2: -529}
    v7 := GenS2{F0: 0, F1: 42}
    fmt.Println(v5, v2)
    var v8 uint8 = uint8(v1)
    var v9 string = "\n\t"
    v8 -= (v8 & (v8 ^ (v8 / ((128) | 1))))
    v1 = (v1 >> ((v1) & 63))
    for fi10 := 0; fi10 < 15; fi10++ {
        if ((!(453 <= 18446744073709551615)) || (127 == 25)) {
            v11 := GenS3{F0: (v8 >> ((v1) & 63))}
            v2 += v5
            if (!(0 > 216)) {
                switch -906 {
                case -153:
                    for fi12 := 0; fi12 < 20; fi12++ {
                        v5 = 771
                        v3 = append(v3, "\n\t")
                        if (2 == 32767) {
                            _ = len(v3)
                            var v13 uint32 = 302
                            if (4294967295 <= 938) {
                                fmt.Println(v13, v9, v2)
                            } else {
                                if true {
                                    var v14 int32 = (v4 << ((v1) & 63))
                                    var v15 int = (fi10 | fi10)
                                    v16 := GenS4{F0: v15, F1: v9}
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                } else {
                                    _ = len(v3)
                                }
                                var v17 int16 = v5
                                _ = v17
                            }
                            _ = v13
                        } else {
                            var v18 int = 32767
                            switch int(v8) {
                            case -142:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            _ = v18
                        }
                        _ = fi12
                    }
                default:
                    switch 127 {
                    case 32767:
                        for fi19 := 0; fi19 < 2; fi19++ {
                            if (!true) {
                                v20 := GenS5{F0: (v1 * 65535), F1: v1}
                                v3 = append(v3, v9)
                                _ = v20
                            } else {
                                if false {
                                    var v21 int32 = v4
                                    v22 := GenS4{F0: 32767, F1: "foo bar"}
                                    v23 := GenS5{F0: (v1 ^ 0), F1: (v1 | v1)}
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                } else {
                                    v2 = (^(v5))
                                    switch (fi19 + (fi19 / ((fi10) | 1))) {
                                    case -2147483648:
                                        v24 := GenS4{F0: (fi10 / ((fi19) | 1)), F1: "\n\t"}
                                        _ = v24
                                    case 1:
                                        var v25 uint64 = 91
                                        _ = v25
                                    default:
                                        switch ("" + "foo bar") {
                                        case "":
                                            fmt.Println(v9)
                                        case "a":
                                            if true {
                                                fmt.Println(fi10)
                                                switch (fi19 - fi19) {
                                                case 9223372036854775807:
                                                    if (!false) {
                                                        v2 = v5
                                                        _ = len(v3)
                                                    }
                                                case 0:
                                                    v2 -= v5
                                                default:
                                                    v4 = (^(v4))
                                                }
                                                var v26 string = "\n\t"
                                                _ = v26
                                            } else {
                                                fmt.Println(fi10, v4)
                                            }
                                        }
                                    }
                                    if true {
                                        switch v9 {
                                        case "":
                                            var v27 int32 = (v4 | v4)
                                            _ = v27
                                        }
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    if true {
                                        var v28 uint8 = (v8 * v8)
                                        switch v9 {
                                        case "\n\t":
                                            if (!(v2 < (-(v5)))) {
                                                var v29 uint8 = uint8(v1)
                                                v2 -= v2
                                                v5 -= -1
                                                _ = v29
                                            }
                                        default:
                                            v30 := GenS4{F0: 4294967295, F1: v9}
                                            _ = v30
                                        }
                                        _ = v28
                                    }
                                }
                                if false {
                                    v1 -= (v1 - v1)
                                    var v31 uint = v1
                                    _ = v31
                                }
                                var v32 uint64 = 2
                                _ = v32
                            }
                            fmt.Println(fi10, v4)
                            _ = fi19
                        }
                    case -547:
                        v2 += v5
                    case -1:
                        var v33 uint16 = 65535
                        _ = v33
                    default:
                        fmt.Println(v2, v2, v4)
                    }
                }
                for fi34 := 0; fi34 < 7; fi34++ {
                    if true {
                        var v35 uint = v1
                        _ = v35
                    }
                    if (v9 == "\n\t") {
                        v2 = (^(v2))
                        v5 = (v2 & (v2 % ((-32768) | 1)))
                    } else {
                        v8 -= v8
                        v9 = v9
                    }
                    _ = fi34
                }
            } else {
                var v36 int8 = int8(fi10)
                _ = v36
            }
            _ = v11
        }
        _ = fi10
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
