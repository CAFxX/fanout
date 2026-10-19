package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 bool; F2 int16 }
type GenS4 struct { F0 int16; F1 uint8; F2 uint16 }

func main() {
    v1 := []int{}
    v2 := []string{"hello", ("a" + "foo bar")}
    v3 := make(map[uint32]string)
    v4 := make(map[string]int)
    if false {
        if (!(3.14159 != 1e10)) {
            var v5 uint32 = 792
            switch "a" {
            case "":
                v5 -= v5
            case "foo bar":
                v6 := GenS1{F0: 2, F1: 18446744073709551615}
                _ = v6
            case "a":
                if true {
                    v5 = (v5 & 4294967295)
                    var v7 float64 = 0.1
                    v8 := v4[("\n\t" + "")]
                    _ = v7
                    _ = v8
                } else {
                    fmt.Println(v5)
                }
            }
            _ = v5
        } else {
            var v9 int = 42
            v10 := GenS2{F0: 42}
            if ((!(0 == 32767)) || false) {
                v9 = 32767
            } else {
                for fi11 := 0; fi11 < 11; fi11++ {
                    v2 = append(v2, ("foo bar" + ""))
                    _ = fi11
                }
                v9 -= (^(v9))
            }
            _ = v10
            _ = v9
        }
    }
    if len(v2) > 0 {
        v2[0] = "a"
    }
    v12 := GenS2{F0: 1}
    v13 := GenS1{F0: 128, F1: 4294967295}
    var v14 uint8 = 127
    fmt.Println(v14)
    if false {
        if (-2.25 == ((0.1 + (-2.25 / -2.25)) / 1.5)) {
            v14 = v14
        } else {
            v14 = (v14 >> ((65535) & 63))
            v14 -= (v14 >> ((uint(v14)) & 63))
        }
        v14 -= 255
    } else {
        switch (("hello" + "hello") + ("a" + "a")) {
        case "\n\t":
            fmt.Println(v14)
        }
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := GenS3{F0: 0, F1: ((42 <= 65535) || true), F2: 728}
        for fi17 := 0; fi17 < 10; fi17++ {
            var v18 uint = 0
            var v19 uint8 = (v14 << ((v18) & 63))
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi15
        _ = v16
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        switch fi20 {
        case 2147483647:
            if (!(!(fi20 != (fi20 - 1)))) {
                v14 += v14
                fmt.Println(v14, fi20)
            } else {
                if (1 <= 1) {
                    if (((!true) || false) && (-1 >= 127)) {
                        v21 := GenS2{F0: v14}
                        v22 := GenS1{F0: v14, F1: 934}
                        v23 := GenS2{F0: (^(v14))}
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    v24 := v3[817]
                    _ = v24
                } else {
                    switch 32768 {
                    case 75:
                        if len(v2) > 0 {
                            v2[0] = ""
                        }
                    }
                    v25 := GenS2{F0: v14}
                    _ = v25
                }
            }
        case 4294967295:
            if ((-(fi20)) < fi20) {
                if false {
                    for fi26 := 0; fi26 < 11; fi26++ {
                        v27 := v4["\n\t"]
                        fmt.Println(v27, fi20, fi26)
                        _ = fi26
                        _ = v27
                    }
                } else {
                    switch fi20 {
                    case -846:
                        var v28 int16 = -1
                        _ = v28
                    default:
                        v29 := GenS4{F0: -1, F1: v14, F2: 32767}
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 18; fi30++ {
                        v14 = 255
                        fmt.Println(fi20)
                        v31 := GenS4{F0: 466, F1: v14, F2: 251}
                        _ = fi30
                        _ = v31
                    }
                }
            } else {
                var v32 uint64 = 4294967295
                _ = v32
            }
        case 128:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
