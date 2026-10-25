package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int16 }
type GenS2 struct { F0 int32; F1 uint; F2 uint16 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int64; F1 int16; F2 bool }
type GenS5 struct { F0 uint64; F1 uint16; F2 int8 }
func genFn0(p0 bool, p1 uint8) uint {
    fmt.Println(p1)
    if (42 != p1) {
        if p0 {
            v5 := GenS1{F0: 0, F1: (("hello" + "\n\t") + ""), F2: 789}
            _ = v5
        } else {
            var v6 uint = 255
            _ = v6
        }
        v7 := GenS1{F0: 730, F1: "a", F2: 32767}
        p0 = true
        _ = v7
    }
    p0 = (4294967295 > -51)
    v8 := GenS2{F0: 1, F1: 761, F2: 32767}
    _ = p0
    _ = p1
    _ = v8
    return 290
}
func genFn1() int16 {
    v9 := []int{}
    var v10 int8 = 118
    _ = len(v9)
    for fi11 := 0; fi11 < 11; fi11++ {
        switch (^(fi11)) {
        case -329:
            if (12 != 36) {
                switch 4294967295 {
                case 9223372036854775807:
                    v10 = v10
                case 42:
                    for fi12 := 0; fi12 < 9; fi12++ {
                        v10 = v10
                        v10 = (-(v10))
                        _ = fi12
                    }
                case 0:
                    var v13 float64 = 1e10
                    _ = v13
                default:
                    var v14 int16 = 0
                    _ = v14
                }
                if (-2.25 < ((3.14159 + ((1e10 - (((0.0 - 1.5) + 3.14159) * ((0.1 * 1.5) - 0.1))) + 1.5)) * 3.14159)) {
                    for fi15 := 0; fi15 < 8; fi15++ {
                        v10 = (v10 * (v10 % ((v10) | 1)))
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        v16 := GenS2{F0: 42, F1: 4294967295, F2: 0}
                        _ = fi15
                        _ = v16
                    }
                } else {
                    if true {
                        if true {
                            switch fi11 {
                            case -23:
                                v10 -= v10
                            case 1:
                                var v17 float32 = 1e10
                                _ = v17
                            }
                            v10 = v10
                        }
                    }
                    fmt.Println(v10, v10)
                }
            }
        case 490:
            var v18 uint16 = 1
            _ = v18
        }
        _ = fi11
    }
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 float64 = (0.1 * 1e10)
    v2 := []string{"a"}
    v3 := make(map[uint32]string)
    var v4 int32 = 537
    _ = len(v3)
    v1 -= v1
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 float64 = v1
        if (!(!true)) {
            v21 := GenS1{F0: 2, F1: "\n\t", F2: 533}
            _ = len(v3)
            v20 += (v20 - v1)
            _ = v21
        }
        var v22 uint64 = 4294967295
        _ = fi19
        _ = v20
        _ = v22
    }
    switch 501 {
    case -1:
        if (0 != 1) {
            if (!(true && true)) {
                if (468 >= 2) {
                    v23 := GenS3{F0: 128, F1: 127}
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 11; fi24++ {
                        v25 := GenS4{F0: 597, F1: 32767, F2: true}
                        _ = fi24
                        _ = v25
                    }
                }
            }
            var v26 int = -9223372036854775808
            _ = v26
        } else {
            if false {
                v4 = (v4 & v4)
                switch (("" + "hello") + ("foo bar" + ("hello" + "a"))) {
                case "foo bar":
                    var v27 uint = 1
                    _ = v27
                }
            }
        }
    case 4294967295:
        v4 += (v4 - 1)
    case 9223372036854775807:
        v4 -= (v4 * (v4 & (^(v4))))
    }
    var v28 int = -1
    v29 := GenS5{F0: 42, F1: 42, F2: -128}
    v30 := GenS1{F0: 18446744073709551615, F1: ((("foo bar" + "\n\t") + ("\n\t" + "foo bar")) + "hello"), F2: 1}
    v4 += (v4 >> ((255) & 63))
    if (0.1 <= ((0.0 / 1.5) * (1e10 * 3.14159))) {
        if (true && false) {
            fmt.Println(v1)
            for fi31 := 0; fi31 < 16; fi31++ {
                fmt.Println(v28, v1, v1)
                var v32 int64 = 0
                v32 -= (-(v32))
                _ = fi31
                _ = v32
            }
            v33 := GenS3{F0: (v28 - -649), F1: 127}
            _ = v33
        }
        var v34 int64 = 1
        v35 := v3[4294967295]
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 15; fi36++ {
        delete(v3, 4294967295)
        if (611 < 2147483647) {
            var v37 uint64 = 442
            switch "foo bar" {
            case "foo bar":
                for fi38 := 0; fi38 < 9; fi38++ {
                    var v39 int16 = genFn1()
                    v40 := GenS3{F0: fi38, F1: 0}
                    _ = fi38
                    _ = v39
                    _ = v40
                }
            case "":
                v41 := GenS2{F0: (v4 << ((42) & 63)), F1: 2, F2: 82}
                _ = v41
            case "\n\t":
                var v42 int8 = 121
                _ = v42
            default:
                var v43 float32 = -2.25
                _ = v43
            }
            _ = v37
        } else {
            switch 82 {
            case -436:
                v28 -= -2147483648
            }
        }
        _ = fi36
    }
    fmt.Println(v4, v28, v1)
    var v44 uint64 = 419
    var v45 float32 = 1e10
    v46 := GenS1{F0: (v44 * 0), F1: "foo bar", F2: -766}
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v44
    _ = v45
    _ = v46
}
